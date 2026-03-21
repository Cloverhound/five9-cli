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

# Upload a WAV file as a new prompt
five9 prompts upload --file greeting.wav --name "Main Greeting"

# Replace an existing prompt's audio
five9 prompts replace --file updated.wav --name "Main Greeting"

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
- **Per-folder defaults** — associate a user with a working directory for multi-tenant workflows
- **Environment variables** — set `FIVE9_USERNAME` and `FIVE9_PASSWORD` to skip the keyring

```bash
five9 login                                        # Login (prompts for credentials)
five9 logout                                       # Remove stored credentials
five9 auth status                                  # Show current user and folder default
five9 auth list                                    # List all authenticated users
five9 auth switch <username>                       # Switch default user
five9 auth set-folder-default <username>           # Set default user for current folder
five9 auth clear-folder-default                    # Remove folder default
```

Credential resolution order: `--username`/`--password` flags > `FIVE9_USERNAME`/`FIVE9_PASSWORD` env vars > `--user` flag > `FIVE9_USER` env var > folder default (`.five9-cli/config.json`) > global default > OS keyring.

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

### Automatic Setup

The installer and `five9 post-install` command will offer to install the skill for detected agents (Claude Code, Claude Cowork, OpenAI Codex, Cursor) via an interactive menu. Skills are also kept up to date when you run `five9 update`.

### Manual Setup

If you prefer to install manually, download the skill into the correct folder for your coding agent:

#### Claude Code

```bash
mkdir -p ~/.claude/skills/five9-cli
curl -fsSL https://raw.githubusercontent.com/Cloverhound/five9-cli/main/skill/SKILL.md \
  -o ~/.claude/skills/five9-cli/SKILL.md
```

#### OpenAI Codex

```bash
mkdir -p ~/.codex/skills/five9-cli
curl -fsSL https://raw.githubusercontent.com/Cloverhound/five9-cli/main/skill/SKILL.md \
  -o ~/.codex/skills/five9-cli/SKILL.md
```

#### Cursor

```bash
mkdir -p ~/.cursor/skills/five9-cli
curl -fsSL https://raw.githubusercontent.com/Cloverhound/five9-cli/main/skill/SKILL.md \
  -o ~/.cursor/skills/five9-cli/SKILL.md
```

#### Claude Cowork

Cowork requires uploading a ZIP file through the Claude Desktop app. Run `five9 post-install` to generate the ZIP, or create it manually:

```bash
mkdir -p /tmp/five9-cli && cp ~/.claude/skills/five9-cli/SKILL.md /tmp/five9-cli/
cd /tmp && zip -r ~/Downloads/five9-cli-skill.zip five9-cli/
```

Then upload at: Claude Desktop → Cowork tab → Customize → Skills → + → Upload a skill.

> These commands install the skill globally (user-level). You can also install per-project by placing the `five9-cli/SKILL.md` folder inside your project's `.claude/skills/`, `.codex/skills/`, or `.cursor/skills/` directory instead.

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
