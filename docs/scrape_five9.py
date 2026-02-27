#!/usr/bin/env python3
"""
Parallel Five9 documentation scraper using Playwright.
Scrapes all Configuration Web Services API pages to markdown files.
"""

import asyncio
import json
import os
import re
import sys
import time
from pathlib import Path

SCRAPE_DIR = Path(__file__).parent / "scrape"
SCRAPE_DIR.mkdir(exist_ok=True)

CONCURRENCY = 8  # parallel browser tabs
LANDING_URL = "https://documentation.five9.com/bundle/api-config-web/page/config-web-services/landing-config-web-services.htm"

LOGIN_EMAIL = "eumansky@cloverhound.com"
LOGIN_PASSWORD = "CC4life!"


async def login(context):
    """Log into Five9 documentation portal via the VCC/Cloud Contact Center login flow."""
    print("  Logging in...")
    page = await context.new_page()
    await page.goto(LANDING_URL, wait_until="networkidle", timeout=60000)

    # Check if we're already logged in
    try:
        await page.wait_for_selector(".body-container", timeout=8000)
        print("  Already logged in!")
        await page.close()
        return
    except Exception:
        pass

    # Step 1: Click "Five9 Customers" on the docs login chooser page
    try:
        await page.click("text=Five9 Customers", timeout=10000)
        print("  Clicked 'Five9 Customers'...")
    except Exception as e:
        await page.screenshot(path=str(SCRAPE_DIR / "_debug_login_step1.png"))
        print(f"  ERROR: Could not find 'Five9 Customers' link: {e}")
        await page.close()
        return

    # Step 2: Wait for the community.five9.com login page with two panels
    # Right panel = "Five9 Support Community Login" (this is the correct one for docs)
    await page.wait_for_timeout(3000)
    await page.wait_for_load_state("networkidle")
    current_url = page.url
    print(f"  Redirected to: {current_url}")

    # Step 3: Fill the Support Community Login form (RIGHT panel on community.five9.com)
    try:
        # The community login form has email + password fields
        # Look for the input fields on the community.five9.com page
        email_sel = 'input[placeholder*="company.com"], input[name="username"], input[type="text"]'
        password_sel = 'input[type="password"]'

        await page.wait_for_selector(email_sel, timeout=15000)
        await page.fill(email_sel, LOGIN_EMAIL)
        print(f"  Filled community email: {LOGIN_EMAIL}")

        await page.fill(password_sel, LOGIN_PASSWORD)
        print("  Filled community password")

        # Click the community "Log in" button (in the dark navy panel)
        # Try submit button in the form, or the specific community login button
        try:
            await page.click('input[type="submit"][value="Log in"], button[type="submit"]', timeout=3000)
        except Exception:
            try:
                # Try finding Log in button that's NOT the VCC one
                # The community panel log in button
                login_buttons = page.locator('input[value*="Log"], button:has-text("Log")')
                count = await login_buttons.count()
                print(f"  Found {count} login buttons, clicking last one...")
                if count > 1:
                    await login_buttons.nth(count - 1).click(timeout=3000)
                else:
                    await login_buttons.first.click(timeout=3000)
            except Exception:
                await page.click('text="Log in"', timeout=3000)
        print("  Clicked community 'Log in', waiting for redirect...")

    except Exception as e:
        await page.screenshot(path=str(SCRAPE_DIR / "_debug_login_fill_error.png"))
        print(f"  ERROR filling login form: {e}")
        await page.close()
        return

    # Step 6: Wait for redirect back to documentation.five9.com
    await page.wait_for_timeout(8000)
    await page.wait_for_load_state("networkidle")

    final_url = page.url
    print(f"  Final URL after login: {final_url}")
    await page.screenshot(path=str(SCRAPE_DIR / "_debug_login_final.png"))

    if "documentation.five9.com" in final_url:
        try:
            await page.wait_for_selector(".body-container", timeout=15000)
            print("  Login successful!")
        except Exception:
            print("  Warning: on docs site but .body-container not found")
    else:
        print(f"  Warning: not redirected to docs site yet. Will try navigating...")
        # Try navigating to the landing page directly now that we're logged in
        await page.goto(LANDING_URL, wait_until="networkidle", timeout=60000)
        await page.wait_for_timeout(3000)
        final_url2 = page.url
        print(f"  After manual nav: {final_url2}")
        await page.screenshot(path=str(SCRAPE_DIR / "_debug_login_final2.png"))
        try:
            await page.wait_for_selector(".body-container", timeout=15000)
            print("  Login successful after manual navigation!")
        except Exception:
            print("  ERROR: Could not reach docs after login.")

    await page.close()


# JS to extract all unique page URLs from the landing page sidebar/TOC
EXTRACT_URLS_JS = """
() => {
    const allLinks = document.querySelectorAll('a[href]');
    const seen = new Set();
    const urls = [];
    const skip = new Set(['Skip to main content','Skip to search','English (United States)',
        'Français (Canada)','Deutsch (Germany)','Español (Latinoamérica)','Português (Brasil)','Logout','Zoomin']);
    const landing = window.location.href.split('#')[0].split('?')[0];
    allLinks.forEach(a => {
        const href = a.href.split('#')[0].split('?')[0];
        const text = a.textContent.trim();
        if (href.includes('/bundle/api-config-web/page/') && !seen.has(href) && text && !skip.has(text) && href !== landing) {
            seen.add(href);
            urls.push(href);
        }
    });
    return urls;
}
"""

# JS to extract main content as Markdown from a page
EXTRACT_MD_JS = """
() => {
    const container = document.querySelector('.body-container');
    if (!container) return null;
    const clone = container.cloneNode(true);
    clone.querySelectorAll('script, style, nav, .zDocsTopicActions, [class*="feedback"], [class*="share"], [class*="PrevTopic"], [class*="NextTopic"], [class*="breadcrumb"], .zDocsRelatedContent, .zDocsTopicPageBodyRelatedContent, button').forEach(el => el.remove());

    let md = '';
    function walk(node, ld) {
        if (node.nodeType === 3) { md += node.textContent; return; }
        if (node.nodeType !== 1) return;
        const t = node.tagName.toLowerCase();
        if (['script','style','nav','button','svg'].includes(t)) return;
        if (t === 'h1') { md += '\\n# '; node.childNodes.forEach(c => walk(c, ld)); md += '\\n\\n'; return; }
        if (t === 'h2') { md += '\\n## '; node.childNodes.forEach(c => walk(c, ld)); md += '\\n\\n'; return; }
        if (t === 'h3') { md += '\\n### '; node.childNodes.forEach(c => walk(c, ld)); md += '\\n\\n'; return; }
        if (t === 'h4') { md += '\\n#### '; node.childNodes.forEach(c => walk(c, ld)); md += '\\n\\n'; return; }
        if (t === 'h5') { md += '\\n##### '; node.childNodes.forEach(c => walk(c, ld)); md += '\\n\\n'; return; }
        if (t === 'p') { md += '\\n'; node.childNodes.forEach(c => walk(c, ld)); md += '\\n'; return; }
        if (t === 'br') { md += '\\n'; return; }
        if (t === 'li') { md += '\\n' + '  '.repeat(ld) + '- '; node.childNodes.forEach(c => walk(c, ld)); return; }
        if (t === 'ul' || t === 'ol') { node.childNodes.forEach(c => walk(c, ld + 1)); md += '\\n'; return; }
        if (t === 'pre') { md += '\\n```\\n' + node.textContent + '\\n```\\n'; return; }
        if (t === 'code') { md += '`'; node.childNodes.forEach(c => walk(c, ld)); md += '`'; return; }
        if (t === 'strong' || t === 'b') { md += '**'; node.childNodes.forEach(c => walk(c, ld)); md += '**'; return; }
        if (t === 'em' || t === 'i') { md += '*'; node.childNodes.forEach(c => walk(c, ld)); md += '*'; return; }
        if (t === 'a') { md += '['; node.childNodes.forEach(c => walk(c, ld)); md += '](' + (node.getAttribute('href')||'') + ')'; return; }
        if (t === 'table') {
            const rows = node.querySelectorAll('tr');
            rows.forEach((row, ri) => {
                const cells = row.querySelectorAll('th, td');
                md += '| ';
                cells.forEach(cell => { md += cell.textContent.trim().replace(/\\s+/g,' ') + ' | '; });
                md += '\\n';
                if (ri === 0) { md += '| '; cells.forEach(() => { md += '--- | '; }); md += '\\n'; }
            });
            md += '\\n'; return;
        }
        if (t === 'img') { md += '![' + (node.getAttribute('alt')||'') + '](' + (node.getAttribute('src')||'') + ')'; return; }
        node.childNodes.forEach(c => walk(c, ld));
    }
    walk(clone, 0);
    return md.replace(/[ \\t]+\\n/g, '\\n').replace(/\\n{3,}/g, '\\n\\n').replace(/^[ \\t]+/gm, m => m.replace(/[ \\t]{2,}/g, ' ')).trim();
}
"""


def url_to_filename(url):
    """Convert a URL to a safe markdown filename."""
    path = url.split("/page/")[1] if "/page/" in url else url.split("/")[-1]
    name = path.replace("/", "_").replace(".htm", "")
    name = re.sub(r'[^\w\-.]', '_', name)
    return name + ".md"


async def scrape_page(context, semaphore, url, index, total, stats):
    """Scrape a single page using a new tab in the shared context."""
    async with semaphore:
        page = await context.new_page()
        filename = url_to_filename(url)
        filepath = SCRAPE_DIR / filename
        try:
            await page.goto(url, wait_until="networkidle", timeout=30000)
            await page.wait_for_selector(".body-container", timeout=15000)
            await page.wait_for_timeout(500)

            md = await page.evaluate(EXTRACT_MD_JS)
            if md:
                filepath.write_text(md, encoding="utf-8")
                stats["ok"] += 1
                print(f"  [{stats['ok']+stats['fail']}/{total}] ✓ {filename} ({len(md)} chars)")
            else:
                stats["fail"] += 1
                print(f"  [{stats['ok']+stats['fail']}/{total}] ⚠ {filename} (no content)")
        except Exception as e:
            stats["fail"] += 1
            print(f"  [{stats['ok']+stats['fail']}/{total}] ✗ {filename} — {e}")
        finally:
            await page.close()


async def main():
    from playwright.async_api import async_playwright

    start = time.time()

    async with async_playwright() as p:
        browser = await p.chromium.launch(headless=True)
        context = await browser.new_context()

        # Step 0: Log in — use headful browser for manual login
        print("Step 0: Authenticating (manual login in visible browser)...")
        await login_manual(context)

        # Step 1: Discover all URLs from landing page
        print(f"\nStep 1: Loading landing page to discover all URLs...")
        page = await context.new_page()
        await page.goto(LANDING_URL, wait_until="networkidle", timeout=60000)

        try:
            await page.wait_for_selector(".body-container", timeout=20000)
        except Exception:
            await page.screenshot(path=str(SCRAPE_DIR / "_debug_landing.png"))
            print(f"  .body-container not found after login. Debug screenshot saved.")
            await page.wait_for_timeout(3000)

        await page.wait_for_timeout(2000)  # let TOC fully render

        urls = await page.evaluate(EXTRACT_URLS_JS)
        await page.close()
        print(f"  Found {len(urls)} pages to scrape.\n")

        if not urls:
            print("ERROR: No URLs found. Cookies may have expired.")
            await browser.close()
            sys.exit(1)

        # Save URL list for reference
        urls_file = Path(__file__).parent / "urls.json"
        urls_file.write_text(json.dumps(urls, indent=2))

        # Step 2: Scrape all pages in parallel
        print(f"Step 2: Scraping {len(urls)} pages ({CONCURRENCY} parallel tabs)...")
        stats = {"ok": 0, "fail": 0}
        semaphore = asyncio.Semaphore(CONCURRENCY)
        tasks = [
            scrape_page(context, semaphore, url, i, len(urls), stats)
            for i, url in enumerate(urls)
        ]
        await asyncio.gather(*tasks)

        await browser.close()

    elapsed = time.time() - start
    scraped = list(SCRAPE_DIR.glob("*.md"))
    print(f"\n{'='*50}")
    print(f"Done in {elapsed:.1f}s!")
    print(f"  Scraped: {stats['ok']}/{len(urls)} pages")
    print(f"  Failed:  {stats['fail']}/{len(urls)} pages")
    print(f"  Output:  {SCRAPE_DIR}")


if __name__ == "__main__":
    asyncio.run(main())
