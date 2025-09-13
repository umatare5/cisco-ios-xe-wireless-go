# 📋 Make Command Reference

This document provides an overview of Make targets used in this repository.

> [!NOTE]
> Some scripts require an accessible Cisco C9800 and the following variables:
>
> - `WNC_CONTROLLER`
> - `WNC_ACCESS_TOKEN`
> - `WNC_AP_MAC_ADDR`
> - `WNC_CLIENT_MAC_ADDR` (optional for enhanced client testing)

## 🧰 Commands

Following is a summary of available Make targets:

| Make Target            | Purpose                                    | Underlying Script                       |
| ---------------------- | ------------------------------------------ | --------------------------------------- |
| `help`                 | Show command help overview                 | `scripts/help.sh`                       |
| `deps`                 | Install / update dev tools                 | `scripts/install_dependencies.sh`       |
| `clean`                | Remove caches / temp / coverage files      | `scripts/clean_artifacts.sh`            |
| `lint`                 | Run golangci-lint                          | `scripts/lint.sh`                       |
| `build`                | Verify build (`go build ./...`)            | (inline)                                |
| `test-unit`            | Unit + table + fail-fast                   | `scripts/test_unit.sh`                  |
| `test-integration`     | Integration tests (tests/integration only) | `scripts/test_integration.sh`           |
| `test-unit-coverage`   | Unit tests with coverage analysis          | `scripts/test_unit.sh --coverage`       |
| `yang-list`            | List available YANG models                 | `scripts/get_yang_models.sh`            |
| `yang-model`           | Fetch a YANG module definition             | `scripts/get_yang_model_details.sh`     |
| `yang-statement`       | Fetch a YANG subtree (RESTCONF)            | `scripts/get_yang_statement_details.sh` |
| `pre-commit-install`   | Install pre-commit hook with symlink       | (inline)                                |
| `pre-commit-test`      | Test pre-commit hook without installing    | `scripts/pre_commit_hook.sh`            |
| `pre-commit-uninstall` | Remove a symlink to pre-commit hook        | (inline)                                |

> [!TIP]
> YANG targets accept variables: `make yang-model MODEL=<name>` and `make yang-statement MODEL=<name> STATEMENT=<stmt>`.
>
> You can pass extra flags via `ARGS`, for example: `make yang-list ARGS="--insecure"` for lab certs.

## 🔍 Examples

### Development loop

Install tools, lint the code, and verify the build in one step.

```bash
make deps && make lint && make build
```

### Testing & coverage

Execute unit/integration tests and generate coverage plus an HTML report.

```bash
make test-unit            # unit + table + fail-fast after lint
make test-unit-coverage   # unit tests with coverage analysis
make test-integration     # integration tests (tests/integration only)
```

### YANG Exploration

List available YANG models and inspect a model or a specific statement.

```bash
make yang-list
make yang-model MODEL=Cisco-IOS-XE-wireless-access-point-oper
make yang-statement MODEL=Cisco-IOS-XE-wireless-access-point-oper STATEMENT=access-point-oper-data
```

## 📜 About Underlying Scripts

For detailed script usage and examples, please see the [SCRIPT_REFERENCE.md](./SCRIPT_REFERENCE.md).
