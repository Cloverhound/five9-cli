# Five9 CLI — Agent Instructions

## Project Overview

Go CLI for the Five9 SOAP Admin API, using Cobra. Modeled after `../webex-cli`.

## Codegen

Files in `cmd/` are **auto-generated** by `codegen/generate_cli.py` — never edit them by hand.
They have a comment at the top of each file confirming this.

To regenerate after changing the spec or codegen:

```bash
python3 codegen/extract_api_spec.py   # Parse docs → codegen/five9_api_spec.json
python3 codegen/generate_cli.py       # JSON spec → cmd/*.go
go build -o five9 .                   # Build
```

### Key codegen files

- `codegen/extract_api_spec.py` — parses `docs/scrape/*.md` into `codegen/five9_api_spec.json`
- `codegen/generate_cli.py` — generates `cmd/*.go` from spec
- `cmd/reports.go` — has a hand-written `run-and-wait` composite command (also generated, but with custom logic)

## Architecture

- **SOAP API** (not REST) — endpoint: `https://api.five9.com/wsadmin/v14/AdminWebService`
- **HTTP Basic auth** — credentials stored in OS keyring via `internal/auth/keyring.go`
- **SOAP client** — `internal/soap/` (envelope builder, XML→JSON parser, fault handling)
- **Output formatting** — `internal/output/output.go` (json/table/csv/raw)

## Dependencies

- `tablewriter` pinned to v0.0.5 (v1.x has breaking API changes)
- `go-keyring` for OS credential storage
- `cobra` for CLI framework

## Conventions

- Pattern parameters (e.g. `--group-name-pattern`) default to `".*"` so list commands work without requiring the flag.
- CLI command groups map 1:1 to generated `cmd/*.go` files (21 groups, ~181 SOAP methods).
