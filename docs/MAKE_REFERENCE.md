# 📋 Make Command Reference

Primary reference for Make targets used in this repository. Make is the recommended entry point; each
target delegates to a thin shell script under `scripts/` that provides consistent banners, validation,
and output formatting.

## 🚀 Quick start

```bash
# Optional (only for integration/coverage that hit a live controller)
export WNC_CONTROLLER=<controller-host-or-ip>
export WNC_ACCESS_TOKEN=<base64-username:password>

make deps                  # install/update dev tools
make lint                  # static analysis
make test-unit             # unit + table + fail-fast after lint
make test-integration      # live controller tests (env required)
make test-coverage         # unified coverage (unit+integration)
make test-coverage-report  # HTML coverage export
```

## 🧰 Commands

| Target                 | Purpose                                 | Underlying script                       |
| ---------------------- | --------------------------------------- | --------------------------------------- |
| `help`                 | Show command help overview              | `scripts/help.sh`                       |
| `deps`                 | Install / update dev tools              | `scripts/install_dependencies.sh`       |
| `clean`                | Remove caches / temp / coverage files   | `scripts/clean_artifacts.sh`            |
| `lint`                 | Run golangci-lint                       | `scripts/lint.sh`                       |
| `build`                | Verify build (go build ./...)           | (inline)                                |
| `test-unit`            | Unit + table + fail-fast                | `scripts/test_unit.sh`                  |
| `test-integration`     | Integration against live controller     | `scripts/test_integration.sh`           |
| `test-coverage`        | Unified coverage (unit+integration)     | `scripts/test_coverage.sh`              |
| `test-coverage-report` | Generate HTML coverage report           | `scripts/generate_coverage_report.sh`   |
| `yang-list`            | List available YANG models              | `scripts/get_yang_models.sh`            |
| `yang-model`           | Fetch a YANG module definition          | `scripts/get_yang_model_details.sh`     |
| `yang-statement`       | Fetch a YANG subtree (RESTCONF)         | `scripts/get_yang_statement_details.sh` |
| `pre-commit-install`   | Install pre-commit hook                 | `scripts/pre_commit_hook.sh`            |
| `pre-commit-test`      | Test pre-commit hook without installing | `scripts/pre_commit_hook.sh`            |
| `pre-commit-uninstall` | Remove pre-commit hook                  | `scripts/pre_commit_hook.sh`            |

For detailed script usage and examples, see the Scripts Reference: `docs/SCRIPT_REFERENCE.md`.

> [!NOTE]
> Integration-related targets require a reachable Cisco C9800 controller. following variables before running them:
>
> - `WNC_CONTROLLER`
> - `WNC_ACCESS_TOKEN` (Base64 `user:pass`)

> [!TIP]
> YANG targets accept variables: `make yang-model MODEL=<name>` and `make yang-statement MODEL=<name> STATEMENT=<stmt>`.

## 🔍 Use Cases

### Development loop

```bash
make deps && make lint && make build
```

### Testing & coverage

```bash
make test-unit
make test-integration   # requires WNC_* env
make test-coverage
make test-coverage-report
```

### YANG exploration

```bash
make yang-list
make yang-model MODEL=Cisco-IOS-XE-wireless-access-point-oper
make yang-statement MODEL=Cisco-IOS-XE-wireless-access-point-oper STATEMENT=access-point-oper-data
```
