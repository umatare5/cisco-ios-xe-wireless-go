# 📋 Scripts Overview

Utilities to inspect wireless YANG + live data.

## Env

`WNC_CONTROLLER`, `WNC_ACCESS_TOKEN` required (or flags). Missing => exit 1.

## Common Flags

| Flag | Meaning |
|------|---------|
| `-c` | Controller host/IP |
| `-t` | Base64 token |
| `-p` | Protocol (https/http) |
| `-k` | Insecure TLS (dev only) |
| `-f` | Output format |
| `-v` | Verbose |

## Scripts

| Script | Purpose |
|--------|---------|
| `list_yang_models.sh` | Enumerate modules |
| `get_yang_model_details.sh` | Fetch module text |
| `get_yang_statement_details.sh` | Query data subtree |

## Output Formats

`pretty`, `json`, `raw` (details script only).

## Banners

Unified color banners (auto disable with `--no-color`).

## Discovery

```bash
./scripts/list_yang_models.sh -k
```

## Model Text

```bash
./scripts/get_yang_model_details.sh -m Cisco-IOS-XE-wireless-general-oper -k
```

## Operational Data

```bash
./scripts/get_yang_statement_details.sh -m Cisco-IOS-XE-wireless-general-oper -i general-oper-data -k
```

## Troubleshooting

| Issue | Fix |
|-------|-----|
| Auth fail | Recreate token |
| TLS | Use proper CA / `-k` for dev |
| Empty list | Version mismatch |
| Bad model | Copy exact name |

## Next

See `models.md` for common module names.
