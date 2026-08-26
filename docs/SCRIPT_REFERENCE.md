# 📜 Scripts Reference

This document provides an overview of the development scripts available in this repository.

> [!NOTE]
> Integration tests require an accessible Cisco C9800 and these variables: See [TESTING.md - Prerequisites](./TESTING.md#-prerequisites)

## 🧰 Scripts

Following is a summary of available scripts:

| Script                                                          | Purpose                               | Upstream Make Target |
| --------------------------------------------------------------- | ------------------------------------- | -------------------- |
| [help.sh](#help.sh)                                             | Show command help overview            | `help`               |
| [install_dependencies.sh](#install_dependencies.sh)             | Install / update dev tools            | `deps`               |
| [lint.sh](#lint.sh)                                             | Run golangci-lint                     | `lint`               |
| [test_unit.sh](#test_unit.sh)                                   | Run unit tests with unified coverage  | `test-unit`          |
| [test_integration.sh](#test_integration.sh)                     | Run integration tests with coverage   | `test-integration`   |

## 🗂️ Structure

Scripts share a consistent bootstrap pattern:

- Source `lib/bootstrap.sh` in the entry script.
- Call `init_wnc_libraries(<module_dir>)` to load the target module (e.g., `lib/testing`).
- Expose common predicates, formatters, and validators in the current shell.
- Invoke exactly one exported `run_*_operation` function.
- Keep entry points thin — behavior is centralized under `scripts/lib/`.
- Output goes through the shared `show_*` helpers and `printf`, never `echo -e`.

```plaintext
scripts/
├── <command>.sh            # Thin entry point(s)
└── lib/                    # Reusable modules (loaded via bootstrap)
    ├── bootstrap.sh        # Loader + init
    ├── coverage/           # Coverage + HTML generation
    ├── dependencies/       # Dependency install/update
    ├── lint/               # Lint operations
    ├── output/             # Banners + formatting helpers
    ├── share/              # Shared libraries across modules
    │   └── testing/        # Unified testing operations (core.sh)
    ├── testing/            # go test orchestration
    └── utils/              # Tool install, build, and CLI validation
```

## 📦 Development Scripts

### install_dependencies.sh <a id="install_dependencies.sh"></a> <!-- anchor for internal links -->

install_dependencies.sh installs or updates development tools required for the project. It checks for necessary CLI tools, downloads dependencies, and ensures the environment is ready for development.

#### Usage

```bash
❯ scripts/install_dependencies.sh --help

USAGE: install_dependencies [OPTIONS]

OPTIONS:
  -p, --project <DIR>            Project root directory [default: .]
      --golangci-lint <VERSION>  golangci-lint version [default: latest]
      --gotestsum <VERSION>      gotestsum version [default: latest]
  -v, --verbose                  Enable verbose output
  -c, --clean                    Clean module cache before installing
  -u, --update                   Update all dependencies to latest versions
      --force                    Force reinstall even if exists
      --download-only            Download dependencies without installing
      --verify                   Verify dependencies after installation
      --no-color                 Disable colored output
  -h, --help                     Print help
  -V, --version                  Print version
```

#### Sample Output

<details><summary>Click to expand sample output</summary>

```bash
❯ scripts/install_dependencies.sh
Validating CLI tools (level: standard)...
✓ curl
<snip>

✓ All 4 required CLI tools are available
======================================
        Cisco WNC Dependencies
          Module Management
======================================

ℹ Dependencies Info: Using Go version: go1.25.1
[2] Downloading dependencies...
✓ Dependencies Success: Dependencies tidied
✓ Dependencies Success: Dependencies downloaded

[✓] Dependencies management completed
```

</details>

## 🧪 Testing Scripts

### test_unit.sh <a id="test_unit.sh"></a> <!-- anchor for internal links -->

Runs unit tests with unified coverage support.

#### Usage

```bash
❯ scripts/test_unit.sh --help

USAGE: test_unit [OPTIONS]

OPTIONS:
  -p, --project <DIR>       Project root directory [default: .]
  -v, --verbose             Enable verbose test output
  -s, --short               Run tests in short mode (skip long-running tests)
  -c, --coverage            Generate coverage data
  -o, --output <FILE>       Coverage output file [default: ./tmp/coverage.out]
  -t, --timeout <DURATION>  Test timeout duration [default: 30s]
      --no-color            Disable colored output
  -h, --help                Print help
  -V, --version             Print version
```

#### Sample Output

<details><summary>Click to expand sample output</summary>

```bash
❯ export WNC_CONTROLLER=""
❯ export WNC_ACCESS_TOKEN=""
❯ scripts/test_unit.sh
Validating CLI tools (level: standard)...
✓ curl
<snip>

✓ All 4 required CLI tools are available
======================================
         Cisco WNC Unit Tests
         Go Testing Framework
======================================

→ Starting unit tests...
PASS TestNewClient/ValidClient (0.00s)
PASS TestNewClient/ValidClientWithOptions (0.00s)
PASS TestNewClient/InvalidHost (0.00s)
<snip>

DONE 932 tests, 77 skipped in 8.463s

-----------------------------------------
✓ Unit tests completed successfully
ℹ Info: Duration: 9s
-----------------------------------------
```

</details>

### test_integration.sh <a id="test_integration.sh"></a> <!-- anchor for internal links -->

Runs integration tests against a live Cisco C9800 controller. Requires `WNC_CONTROLLER` and `WNC_ACCESS_TOKEN`.

#### Usage

```bash
❯ scripts/test_integration.sh --help

USAGE: test_integration [OPTIONS]

OPTIONS:
  -p, --project <DIR>       Project root directory [default: .]
  -v, --verbose             Enable verbose test output
      --race                Enable race detection [default: true]
  -o, --output <FILE>       Coverage output file [default: ./tmp/coverage.out]
  -t, --timeout <DURATION>  Test timeout [default: 10m]
      --package <PATTERN>   Package pattern to test [default: ./...]
      --check-env-only      Only check environment without running tests
      --no-color            Disable colored output
  -h, --help                Print help
  -V, --version             Print version
```

#### Sample Output

<details><summary>Click to expand sample output</summary>

```bash
❯ export WNC_CONTROLLER="<controller-host-or-ip>"
❯ export WNC_ACCESS_TOKEN="<base64-username:password>"
❯ scripts/test_integration.sh
Validating CLI tools (level: standard)...
✓ curl
<snip>

✓ All 4 required CLI tools are available
======================================
     Cisco WNC Integration Tests
         Go Testing Framework
======================================

→ Starting integration tests...
PASS afc.TestAfcService/Service_Creation (0.00s)
PASS afc.TestAfcService/Data_Collection (0.08s)
PASS afc.TestAfcService/JSON_Serialization/AfcOperResponse (0.00s)
<snip>

DONE 1004 tests, 21 skipped in 4.215s

-----------------------------------------
✓ Integration tests completed successfully
ℹ Info: Duration: 5s
-----------------------------------------
```

</details>

## ✅ Quality Scripts

### lint.sh <a id="lint.sh"></a> <!-- anchor for internal links -->

Runs golangci-lint using the repo configuration. Supports optional auto-fix.

#### Usage

`scripts/lint.sh` only supports execution with no arguments.

````plaintext

#### Sample Output

<details><summary>Click to expand sample output</summary>

```bash
❯ scripts/lint.sh
Validating CLI tools (level: standard)...
✓ curl
<snip>

✓ All 4 required CLI tools are available
======================================
        Cisco WNC Code Linter
      golangci-lint Integration
======================================

ℹ Info: Starting code linting...
0 issues.

✓ Code linting completed successfully
````

</details>

## 🆘 Help Script

### help.sh <a id="help.sh"></a> <!-- anchor for internal links -->

Prints a consolidated help guide covering common Make targets, environment variables, and examples.

#### Usage

`scripts/help.sh` only supports execution with no arguments.

````bash

#### Sample Output

<details><summary>Click to expand sample output</summary>

```plaintext
Cisco WNC Development Scripts
-------------------------------

USAGE:
    make <target>                   # Use Makefile targets (recommended)
    ./scripts/<script>.sh [options] # Use scripts directly

COMMON DEVELOPMENT TARGETS:
    help                Show this help message
    deps                Install development dependencies
    lint                Run code linting tools
    build               Verify build compilation
    test-unit           Run unit tests only
    test-integration    Run integration tests (requires environment)
    test-unit-coverage  Run unit tests with coverage analysis

ENVIRONMENT VARIABLES:
    WNC_CONTROLLER      Controller hostname/IP for integration tests
    WNC_ACCESS_TOKEN    Base64 encoded credentials for integration tests

EXAMPLES:
    # Basic development workflow
    make deps               # Install dependencies
    make lint               # Check code quality
    make test-unit          # Run unit tests
    make test-unit-coverage # Run unit tests with coverage
    make build              # Verify compilation

    # Integration testing (requires environment setup)
    export WNC_CONTROLLER="<controller-host-or-ip>"
    export WNC_ACCESS_TOKEN="<base64-username:password>"
    make test-integration

SCRIPT DETAILS:
    For specific script options and advanced usage:
    ./scripts/<script_name>.sh --help

    Available scripts:
    - install_dependencies.sh Install Go dependencies
    - lint.sh                Run golangci-lint
    - test_unit.sh           Run unit tests (supports --coverage)
    - test_integration.sh    Run integration tests

PROJECT STRUCTURE:
    scripts/                Script directory
    +-- lib/               Shared libraries
    |   +-- bootstrap.sh   Bootstrap library loader
    |   +-- coverage/      Coverage report functions
    |   +-- dependencies/  Dependency management
    |   +-- output/        Output formatting utilities
    |   +-- testing/       Test utilities
    |   +-- utils/         Utility functions
````

</details>
