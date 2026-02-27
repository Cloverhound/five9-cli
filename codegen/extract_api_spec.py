#!/usr/bin/env python3
"""Parse scraped Five9 API docs into a structured JSON spec."""

import json
import os
import re
import sys

DOCS_DIR = os.path.join(os.path.dirname(__file__), "..", "docs", "scrape")


def parse_method_file(filepath):
    """Parse a single method markdown file into a structured dict."""
    with open(filepath) as f:
        content = f.read()

    lines = content.strip().split("\n")
    if not lines:
        return None

    # Extract method name from first heading (## or ####)
    method_name = None
    for line in lines:
        m = re.match(r"^#{2,4}\s+(\w+)", line)
        if m:
            method_name = m.group(1)
            break

    if not method_name:
        return None

    # Extract description (text before first ####)
    desc_lines = []
    for line in lines[1:]:
        if line.startswith("####"):
            break
        stripped = line.strip()
        if stripped and not stripped.startswith("|") and not stripped.startswith("["):
            desc_lines.append(stripped)
    description = " ".join(desc_lines).strip()
    # Clean up markdown formatting
    description = re.sub(r"\*\*[^*]*\*\*:?\s*", "", description).strip()
    description = re.sub(r"\[([^\]]+)\]\([^)]+\)", r"\1", description)

    # Extract request parameters
    params = []
    in_request_table = False
    in_response_section = False

    for i, line in enumerate(lines):
        # Detect response section
        if re.match(r"^####.*[Rr]esponse", line):
            in_response_section = True
            in_request_table = False
            continue

        # Detect request section (#### with method name, not response)
        if line.startswith("####") and not in_response_section:
            in_request_table = True
            continue

        if line.strip() == "Contains no parameters.":
            in_request_table = False
            continue

        stripped_line = line.strip()
        if in_request_table and stripped_line.startswith("|") and "---" not in stripped_line and "Parameter" not in stripped_line:
            parts = [p.strip() for p in stripped_line.split("|")]
            parts = [p for p in parts if p]
            if len(parts) >= 3:
                param_name = parts[0]
                param_type = parts[1]
                param_desc = parts[2] if len(parts) > 2 else ""
                # Clean param name — take only the first word (the identifier)
                # Some scraped docs have extra text like "Version 3" appended
                if " " in param_name:
                    param_name = param_name.split()[0]
                # Clean param type similarly
                param_type = param_type.strip()
                # Clean description — remove markdown artifacts
                param_desc = re.sub(r"\*\*[^*]*\*\*:?\s*", "", param_desc).strip()
                params.append({
                    "name": param_name,
                    "type": param_type,
                    "description": param_desc,
                })

        # Stop request table at empty line (but not lines with whitespace that may precede table rows)
        if in_request_table and not in_response_section and line.strip() == "" and not (i + 1 < len(lines) and lines[i + 1].strip().startswith("|")):
            in_request_table = False

    return {
        "method": method_name,
        "description": description,
        "params": params,
    }


def parse_type_file(filepath):
    """Parse a single data type markdown file into a structured dict."""
    with open(filepath) as f:
        content = f.read()

    lines = content.strip().split("\n")
    if not lines:
        return None

    # Extract description from text before the table
    desc_lines = []
    for line in lines:
        stripped = line.strip()
        if stripped.startswith("|") or stripped.startswith("#"):
            break
        if stripped and not stripped.startswith("[") and not stripped.startswith("!"):
            desc_lines.append(stripped)
    description = " ".join(desc_lines).strip()
    # Clean markdown formatting
    description = re.sub(r"\*\*[^*]*\*\*:?\s*", "", description).strip()
    description = re.sub(r"\[([^\]]+)\]\([^)]+\)", r"\1", description)

    # Extract fields from the Name/Type/Description table
    fields = []
    in_table = False
    for i, line in enumerate(lines):
        stripped = line.strip()
        # Look for the specific header row with Name | Type | Description
        if stripped.startswith("|") and "Name" in stripped and "Type" in stripped and "Description" in stripped:
            # Skip the separator line after the header
            in_table = False  # Will be set True by the --- line
            continue
        if stripped.startswith("|") and "---" in stripped:
            # Only start table parsing if we just saw the Name/Type/Description header
            # Check previous non-empty line was the header
            for j in range(i - 1, max(i - 3, -1), -1):
                prev = lines[j].strip()
                if prev.startswith("|") and "Name" in prev and "Type" in prev:
                    in_table = True
                    break
            continue
        if in_table and stripped.startswith("|"):
            parts = [p.strip() for p in stripped.split("|")]
            parts = [p for p in parts if p]
            if len(parts) >= 3:
                name = parts[0]
                ftype = parts[1]
                fdesc = parts[2] if len(parts) > 2 else ""
                # Clean name — take first word only
                if " " in name:
                    name = name.split()[0]
                # Skip section dividers (empty type + description)
                if not ftype.strip() and not fdesc.strip():
                    continue
                # Detect arrays
                is_array = "[0..unbounded]" in ftype or "[]" in ftype
                # Clean type — remove array notation
                base_type = re.sub(r"\s*\[.*?\]", "", ftype).strip()
                # Clean description
                fdesc = re.sub(r"\*\*[^*]*\*\*:?\s*", "", fdesc).strip()
                fields.append({
                    "name": name,
                    "type": base_type,
                    "is_array": is_array,
                    "description": fdesc,
                })
        elif in_table and stripped == "":
            break

    if not fields:
        return None

    # Detect enums: all fields have type "string" and 3+ fields
    all_string = all(f["type"] == "string" for f in fields)
    if all_string and len(fields) >= 3:
        return {
            "description": description,
            "is_enum": True,
            "enum_values": [f["name"] for f in fields],
        }

    return {
        "description": description,
        "is_enum": False,
        "fields": fields,
    }


def parse_all_types():
    """Parse all data type files and return a dict keyed by type name."""
    types = {}
    for fname in sorted(os.listdir(DOCS_DIR)):
        if not fname.startswith("config-web-services_data-types_"):
            continue
        if "__ch-" in fname:
            continue
        filepath = os.path.join(DOCS_DIR, fname)
        # Extract type name from filename
        base = fname.replace("config-web-services_data-types_", "").replace(".md", "")
        # For nested filenames like baseOutboundcampaign_outboundCampaign, use last segment
        if "_" in base:
            type_name = base.split("_")[-1]
        else:
            type_name = base

        parsed = parse_type_file(filepath)
        if parsed is None:
            continue

        # Prefer the entry with more fields when duplicates exist
        if type_name in types:
            existing = types[type_name]
            new_fields = len(parsed.get("fields", parsed.get("enum_values", [])))
            old_fields = len(existing.get("fields", existing.get("enum_values", [])))
            if new_fields <= old_fields:
                continue
        types[type_name] = parsed

    return types


def main():
    # Find all method files
    method_files = []
    for fname in sorted(os.listdir(DOCS_DIR)):
        if not fname.startswith("config-web-services_methods_"):
            continue
        if "__ch-" in fname:
            continue
        method_files.append(os.path.join(DOCS_DIR, fname))

    # Group by doc group (second-to-last segment of filename)
    groups = {}
    for filepath in method_files:
        fname = os.path.basename(filepath)
        # Pattern: config-web-services_methods_<group>_<method>.md
        parts = fname.replace(".md", "").split("_")
        # Find the group name: everything between "methods" and the last segment
        methods_idx = parts.index("methods")
        group_parts = parts[methods_idx + 1 : -1]
        group = "-".join(group_parts)
        method_name = parts[-1]

        parsed = parse_method_file(filepath)
        if parsed is None:
            print(f"Warning: could not parse {filepath}", file=sys.stderr)
            continue

        parsed["doc_group"] = group

        if group not in groups:
            groups[group] = []
        groups[group].append(parsed)

    # Map doc groups to CLI commands
    cli_mapping = {
        "agent-groups": "agent-groups",
        "call-variables": "call-variables",
        "campaign-configuration": "campaigns",
        "campaign-management": "campaigns",
        "campaign-profiles": "campaign-profiles",
        "connectors": "web-connectors",
        "contact-fields": "contacts",
        "contact-management": "contacts",
        "dialing-rules": "dialing-rules",
        "disposition-configuration": "dispositions",
        "domain-limits": "domain-limits",
        "IVR-script": "ivr-scripts",
        "list-management": "lists",
        "locales-and-languages": "locales",
        "prompt-management": "prompts",
        "reason-codes": "reason-codes",
        "reports": "reports",
        "session-information": "session",
        "skill-management": "skills",
        "speed-dial-information": "speed-dial",
        "user-management": "users",
        "user-profiles": "user-profiles",
        "VCC-configuration": "vcc",
    }

    # Build final spec grouped by CLI command
    spec = {}
    for doc_group, methods in groups.items():
        cli_cmd = cli_mapping.get(doc_group, doc_group)
        if cli_cmd not in spec:
            spec[cli_cmd] = []
        spec[cli_cmd].extend(methods)

    # Parse data types
    types = parse_all_types()

    output = {
        "commands": spec,
        "types": types,
    }

    output_path = os.path.join(os.path.dirname(__file__), "five9_api_spec.json")
    with open(output_path, "w") as f:
        json.dump(output, f, indent=2)

    # Print summary
    total = sum(len(m) for m in spec.values())
    print(f"Extracted {total} methods across {len(spec)} CLI commands")
    for cmd, methods in sorted(spec.items()):
        print(f"  {cmd}: {len(methods)} methods")
    enum_count = sum(1 for t in types.values() if t.get("is_enum"))
    print(f"Extracted {len(types)} data types ({enum_count} enums, {len(types) - enum_count} structs)")


if __name__ == "__main__":
    main()
