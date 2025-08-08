# 🧾 Scripts Overview

Utilities for developing, testing, inspecting YANG / RESTCONF data.

## 🚀 Quick Start

```bash
export WNC_CONTROLLER=<host>
export WNC_ACCESS_TOKEN=$(echo -n 'user:pass' | base64)
./scripts/list_yang_models.sh -k       # enumerate
./scripts/get_yang_model_details.sh -m Cisco-IOS-XE-wireless-general-oper -k
```

## 📂 Key Scripts

| Script | Purpose |
|--------|---------|
| `list_yang_models.sh` | List modules |
| `get_yang_model_details.sh` | Fetch module text |
| `get_yang_statement_details.sh` | Query subtree |
| `test_unit.sh` | Unit tests + lint |
| `test_integration.sh` | Live tests |
| `test_coverage.sh` | Coverage aggregate |
| `clean_artifacts.sh` | Remove temp outputs |
| `install_dependencies.sh` | Go mod hygiene |

## ⚙️ Common Flags

| Flag | Meaning |
|------|---------|
| `-c` / `--controller` | Controller host/IP |
| `-t` / `--token` | Base64 creds |
| `-k` / `--insecure` | Skip TLS verify (dev) |
| `-f` / `--format` | Output fmt (pretty/json/raw) |
| `-v` / `--verbose` | Verbose output |
| `--no-color` | Disable color |

Missing required env/flags ⇒ exit 1.

## 🛠 Dev Flow

```bash
./scripts/install_dependencies.sh --clean
./scripts/test_unit.sh --coverage
./scripts/test_integration.sh
./scripts/test_coverage.sh
```

## 🔐 Env Vars

| Var | Purpose |
|-----|---------|
| `WNC_CONTROLLER` | Host/IP |
| `WNC_ACCESS_TOKEN` | Base64 creds |

## 🧱 Architecture (Collapsed)

<details><summary>Structure & features</summary>

Modular loader in `scripts/lib/` (bootstrap + core + network + output). Predicate helpers (`is_verbose_enabled`) keep conditions readable. Strict `set -euo pipefail` and explicit validation guard misuse.

Structure (simplified):

- `scripts/*.sh` entry points (thin)
- `scripts/lib/bootstrap.sh` loader
- `scripts/lib/core` predicates / constants / argc
- `scripts/lib/network` http_client, auth, yang helpers
- `scripts/lib/output` formatting
- `scripts/lib/modules` (testing, lint, coverage, yang, deps, artifacts)

</details>

## 🐞 Troubleshooting (Collapsed)

<details><summary>Common issues</summary>

| Symptom | Fix |
|---------|-----|
| 401 | Recreate token |
| TLS fail | Provide CA or use `-k` (dev) |
| Empty list | Check controller version |
| Permission | `chmod +x scripts/*.sh` |

</details>

## 📚 Links

`docs/scripts/` · `docs/api/` · `docs/testing/` · `docs/security/`

Run any script with `--help` for full options.
