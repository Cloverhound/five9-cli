# Five9 CLI

A command-line tool for the Five9 Configuration Web Services (SOAP) API.

## Install

**macOS / Linux:**
```bash
curl -fsSL https://raw.githubusercontent.com/Cloverhound/five9-cli/main/install.sh | sh
```

**Windows (PowerShell):**
```powershell
irm https://raw.githubusercontent.com/Cloverhound/five9-cli/main/install.ps1 | iex
```

Or download from [Releases](https://github.com/Cloverhound/five9-cli/releases).

## Update

```bash
five9 update
```

## Quick Start

```bash
# Login (stores credentials in OS keyring)
five9 login

# List resources
five9 skills list
five9 campaigns list
five9 users list
five9 agent-groups list
five9 dispositions list

# Get details
five9 skills get --skill-name "My Skill"
five9 users get --user-name "jdoe@example.com"

# Run a report
five9 reports run-and-wait --folder-name "My Reports" --report-name "Call Log"

# Different output formats
five9 skills list --output table
five9 campaigns list --output csv > campaigns.csv
```

## API Coverage

21 command groups covering 181 SOAP methods:

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

## Authentication

- **HTTP Basic auth** over SOAP — `five9 login` prompts for username and password
- **OS keyring storage** — credentials stored securely in macOS Keychain / Linux keyring / Windows Credential Manager
- **Multi-user** — log in with multiple Five9 accounts and switch between them
- **Environment variables** — set `FIVE9_USERNAME` and `FIVE9_PASSWORD` to skip the keyring

```bash
five9 login                    # Login (prompts for credentials)
five9 logout                   # Remove stored credentials
five9 auth status              # Show current user
five9 auth list                # List all authenticated users
five9 auth switch <username>   # Switch default user
```

Credential resolution order: `--username`/`--password` flags > `FIVE9_USERNAME`/`FIVE9_PASSWORD` env vars > OS keyring.

## Output Formats

Control output with `--output`:

| Format | Description |
|--------|-------------|
| `json` | Pretty-printed JSON (default) |
| `table` | ASCII table with auto-detected columns |
| `csv` | CSV with headers |
| `raw` | Raw API response |

## Global Flags

| Flag | Description |
|------|-------------|
| `--username <user>` | Override stored username |
| `--password <pass>` | Override stored password |
| `--user <name>` | Use a specific authenticated user |
| `--endpoint <url>` | Override Five9 API endpoint |
| `--output json\|table\|csv\|raw` | Output format (default: json) |
| `--debug` | Show SOAP request/response details |
| `--dry-run` | Print write requests without executing them |

## Coding Agent Skill

The Five9 CLI includes a [skill](https://agentskills.io) that enables AI coding agents to query and manage your Five9 environment. It works with any agent that supports the skills standard — [Claude Code](https://claude.com/claude-code), [OpenAI Codex](https://openai.com/codex/), [Cursor](https://cursor.com), and others.

### Setup

1. Install the Five9 CLI (see [Install](#install))
2. Login: `five9 login`
3. Download the skill into the correct folder for your coding agent:

#### Claude Code

**macOS / Linux:**
```bash
mkdir -p ~/.claude/skills/five9-cli
curl -fsSL https://raw.githubusercontent.com/Cloverhound/five9-cli/main/skill/SKILL.md \
  -o ~/.claude/skills/five9-cli/SKILL.md
```

**Windows (PowerShell):**
```powershell
New-Item -ItemType Directory -Force -Path "$env:USERPROFILE\.claude\skills\five9-cli"
Invoke-WebRequest -Uri "https://raw.githubusercontent.com/Cloverhound/five9-cli/main/skill/SKILL.md" `
  -OutFile "$env:USERPROFILE\.claude\skills\five9-cli\SKILL.md"
```

#### OpenAI Codex

**macOS / Linux:**
```bash
mkdir -p ~/.agents/skills/five9-cli
curl -fsSL https://raw.githubusercontent.com/Cloverhound/five9-cli/main/skill/SKILL.md \
  -o ~/.agents/skills/five9-cli/SKILL.md
```

**Windows (PowerShell):**
```powershell
New-Item -ItemType Directory -Force -Path "$env:USERPROFILE\.agents\skills\five9-cli"
Invoke-WebRequest -Uri "https://raw.githubusercontent.com/Cloverhound/five9-cli/main/skill/SKILL.md" `
  -OutFile "$env:USERPROFILE\.agents\skills\five9-cli\SKILL.md"
```

#### Cursor

**macOS / Linux:**
```bash
mkdir -p ~/.cursor/skills/five9-cli
curl -fsSL https://raw.githubusercontent.com/Cloverhound/five9-cli/main/skill/SKILL.md \
  -o ~/.cursor/skills/five9-cli/SKILL.md
```

**Windows (PowerShell):**
```powershell
New-Item -ItemType Directory -Force -Path "$env:USERPROFILE\.cursor\skills\five9-cli"
Invoke-WebRequest -Uri "https://raw.githubusercontent.com/Cloverhound/five9-cli/main/skill/SKILL.md" `
  -OutFile "$env:USERPROFILE\.cursor\skills\five9-cli\SKILL.md"
```

> These commands install the skill globally (user-level). You can also install per-project by placing the `five9-cli/SKILL.md` folder inside your project's `.claude/skills/`, `.agents/skills/`, or `.cursor/skills/` directory instead.

4. If the `five9` binary is not in your `$PATH`, ask your coding agent to update the binary path in the skill file.

### Example Prompts

```
/five9-cli list all skills and show which campaigns they're assigned to

/five9-cli how many users are in each agent group?

/five9-cli check auth status and list all dispositions
```

## Development

See [AGENTS.md](AGENTS.md) for project structure and development workflow.

Commands in `cmd/` are **auto-generated** from the API spec — do not edit by hand. See the codegen pipeline in `AGENTS.md` for details.

## License

MIT
