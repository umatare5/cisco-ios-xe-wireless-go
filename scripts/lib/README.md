# 🧱 Scripts Library

Modular shell library powering all `scripts/*.sh` entry points.

## 📂 Layout

```text
scripts/lib/
    bootstrap.sh
    core/ (predicates, constants, argc helpers, args)
    utils/ (validation, files, cli, build, deps)
    network/ (http_client, auth, yang_common)
    output/ (output_formatter)
    modules/ (testing, lint, deps, yang, coverage, validation, artifacts)
```

## 🚀 Init Patterns

| Use Case | Code |
|----------|------|
| Full | `init_wnc_libraries "$SCRIPT_DIR" "$SCRIPT_DIR/lib/testing"` |
| Basic | `init_wnc_basic` |
| Network | `init_wnc_network` |

Always: `source "$SCRIPT_DIR/lib/bootstrap.sh"` first.

## 🔌 Key Functions

| Func | Purpose |
|------|---------|
| `init_wnc_libraries` | Load all + module path |
| `init_wnc_basic` | Core + utils only |
| `init_wnc_network` | Core + network + output |

## 🔍 Categories (Collapsed)

<details><summary>Library roles</summary>

Core: predicates, constants, argc, arg parsing.
Utils: validation, fs, tool checks, build, deps.
Network: HTTP, auth, YANG helpers.
Output: formatting.
Modules: feature focused (testing, lint, coverage, etc.).

</details>

## 🔄 Migration

Old (deprecated): `init_script_libraries` + manual sourcing.

New:

```bash
source "$SCRIPT_DIR/lib/bootstrap.sh"
init_wnc_libraries "$SCRIPT_DIR" "$SCRIPT_DIR/lib/testing"
```

## 💡 Benefits

Separation, lower coupling, selective load (faster), clearer maintenance.

## 🚨 Deprecated (Collapsed)

<details><summary>Legacy APIs</summary>

`init_script_libraries`, `source_wnc_libraries` → remove in future; verbose mode warns.

</details>

## 🧪 Notes

Deterministic: strict `set -euo pipefail`. Call init early; avoid ad‑hoc sourcing.

