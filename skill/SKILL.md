---
name: five9-cli
description: "Five9 CLI: query and manage the Five9 Configuration Web Services (SOAP) API via the `five9` command-line tool. Use for listing resources, checking configurations, managing campaigns, skills, users, and administering Five9 environments."
argument-hint: "[command or resource-name]"
allowed-tools: Bash, Read, Grep, Glob
user-invocable: true
---

# Five9 CLI Skill

This skill uses the `five9` CLI tool to interact with the Five9 Configuration Web Services SOAP API.

**Setup:** Install the CLI via `curl -fsSL https://raw.githubusercontent.com/Cloverhound/five9-cli/main/install.sh | sh`, or set the path to your local build below.

**Binary path** (update to match your installation):
```bash
five9
```

## Authentication

The CLI uses HTTP Basic auth with credentials stored in the OS keyring.

```bash
five9 auth status          # Show current user
five9 auth list            # List all stored users
five9 auth switch <user>   # Change default user
five9 login                # Prompt for username and password
five9 logout [username]    # Remove stored credentials
```

If not logged in, use `--username` and `--password` flags, or set `FIVE9_USERNAME` and `FIVE9_PASSWORD` environment variables.

Credential resolution order: `--username`/`--password` flags > `FIVE9_USERNAME`/`FIVE9_PASSWORD` env vars > OS keyring.

## Command Structure

```
five9 <resource-group> <action> [flags]
```

### Command Groups

| Command | Description |
|---------|-------------|
| `agent-groups` | Manage agent groups |
| `call-variables` | Manage call variables and variable groups |
| `campaign-profiles` | Manage campaign profiles |
| `campaigns` | Manage campaigns (inbound, outbound, autodial) |
| `contacts` | Manage contacts and contact fields |
| `dialing-rules` | Manage dialing rules |
| `dispositions` | Manage dispositions |
| `domain-limits` | View domain limits |
| `ivr-scripts` | Manage IVR scripts |
| `lists` | Manage contact lists |
| `locales` | Manage locales and languages |
| `prompts` | Manage audio prompts |
| `reason-codes` | Manage reason codes |
| `reports` | Run and retrieve reports |
| `session` | Manage API sessions |
| `skills` | Manage skills |
| `speed-dial` | Manage speed dial numbers |
| `user-profiles` | Manage user profiles |
| `users` | Manage users |
| `vcc` | VCC configuration |
| `web-connectors` | Manage web connectors |

### Global Flags

- `--username <user>` — Override stored username
- `--password <pass>` — Override stored password
- `--user <name>` — Use a specific authenticated user from the keyring
- `--endpoint <url>` — Override Five9 API endpoint
- `--output json|table|csv|raw` — Output format (default: json)
- `--debug` — Show SOAP request/response details
- `--dry-run` — Print write requests without executing them

## Common Examples

### Skills

```bash
five9 skills list
five9 skills get --skill-name "Sales"
five9 skills create --body-file skill.json
five9 skills delete --skill-name "Old Skill"
```

### Campaigns

```bash
five9 campaigns list
five9 campaigns get --campaign-name "Outbound Sales"
five9 campaigns start --campaign-name "Outbound Sales"
five9 campaigns stop --campaign-name "Outbound Sales"
five9 campaigns reset --campaign-name "Outbound Sales"
```

### Users

```bash
five9 users list
five9 users get --user-name "jdoe@example.com"
five9 users info
five9 users create --body-file user.json
```

### Agent Groups

```bash
five9 agent-groups list
five9 agent-groups get --group-name "Support Team"
five9 agent-groups create --body-file group.json
five9 agent-groups modify --body-file group-update.json
```

### Dispositions

```bash
five9 dispositions list
five9 dispositions get --disposition-name "Sale"
five9 dispositions create --body-file disposition.json
```

### Reports

```bash
# Run a report and wait for results (composite command)
five9 reports run-and-wait --folder-name "My Reports" --report-name "Call Log"

# Or run manually
five9 reports run --folder-name "My Reports" --report-name "Call Log"
five9 reports is-running --report-id <id>
five9 reports get-result --report-id <id>
```

### Lists

```bash
five9 lists list-info
five9 lists create --list-name "New List"
five9 lists delete --list-name "Old List"
five9 lists add-record --list-name "My List" --body-file record.json
```

### IVR Scripts

```bash
five9 ivr-scripts list
five9 ivr-scripts get --script-name "Main IVR"
```

### Contact Management

```bash
five9 contacts get-fields
five9 contacts lookup --body-file lookup.json
```

## Complex Parameters

Some commands accept complex parameters via `--body-file <path.json>`. Use `--sample` to print a sample JSON body:

```bash
five9 users create --sample          # Print sample JSON
five9 users create --body-file user.json
```

## Output Handling

```bash
# Pretty JSON (default)
five9 skills list

# Table format
five9 skills list --output table

# CSV — great for exports
five9 users list --output csv > users.csv

# Raw XML response
five9 skills list --output raw

# Debug mode — see SOAP envelopes
five9 skills list --debug
```

## Shell Usage Best Practices

1. **Check `--help` before guessing** subcommand names:
   ```bash
   five9 campaigns --help
   ```

2. **Use `--sample`** to discover JSON body format for create/modify commands:
   ```bash
   five9 campaigns create --sample
   ```

3. **Write output to temp files** when gathering data, then read the files to analyze:
   ```bash
   five9 skills list > /tmp/skills.json
   ```

4. **Use `--dry-run`** to preview write operations before executing:
   ```bash
   five9 campaigns start --campaign-name "Sales" --dry-run
   ```

## When Answering Questions

1. **Check auth first** with `five9 auth status` to confirm the user is logged in
2. **Use `--help`** on any command you're not sure about before running it
3. **Write to temp files** when gathering data, then read the files to analyze
4. **Use `--sample`** to discover JSON body formats for create/modify commands
