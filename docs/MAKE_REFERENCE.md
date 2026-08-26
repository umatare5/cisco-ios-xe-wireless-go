# 📋 Make Command Reference

This document provides an overview of Make targets used in this repository.

> [!NOTE]
> Integration tests require an accessible Cisco C9800 and these variables: See [TESTING.md - Prerequisites](./TESTING.md#-prerequisites)

## 🧰 Commands

Following is a summary of available Make targets:

| Make Target            | Purpose                                    | Underlying Script                       |
| ---------------------- | ------------------------------------------ | --------------------------------------- |
| `help`                 | Show command help overview                 | `scripts/help.sh`                       |
| `deps`                 | Install / update dev tools                 | `scripts/install_dependencies.sh`       |
| `lint`                 | Run golangci-lint                          | `scripts/lint.sh`                       |
| `build`                | Verify build (`go build ./...`)            | (inline)                                |
| `test-unit`            | Unit + table + fail-fast                   | `scripts/test_unit.sh`                  |
| `test-integration`     | Integration tests (tests/integration only) | `scripts/test_integration.sh`           |
| `test-unit-coverage`   | Unit tests with coverage analysis          | `scripts/test_coverage.sh`              |
| `pre-commit-install`   | Install pre-commit hook with symlink       | (inline)                                |
| `pre-commit-test`      | Test pre-commit hook without installing    | `.githooks/pre-commit`                  |
| `pre-commit-uninstall` | Remove a symlink to pre-commit hook        | (inline)                                |

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

## 📜 About Underlying Scripts

For detailed script usage and examples, please see the [SCRIPT_REFERENCE.md](./SCRIPT_REFERENCE.md).
