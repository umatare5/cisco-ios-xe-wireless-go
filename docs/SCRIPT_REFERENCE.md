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
| [test_unit.sh](#test_unit.sh)                                   | Run unit tests                        | `test-unit`          |
| [test_coverage.sh](#test_coverage.sh)                           | Unit tests with coverage and a report | `test-unit-coverage` |
| [test_integration.sh](#test_integration.sh)                     | Run integration tests with coverage   | `test-integration`   |

## 🗂️ Structure

One entry script per Make target, each owning the work only it performs:

- Each entry script `source`s the `lib/` modules it needs, by name — there is no loader.
- A module lands in `lib/` only when more than one command calls into it.
- Log lines go through `lib/log.sh`, the only file that spells an ANSI escape.
- No script parses flags. Only `test_coverage.sh` takes arguments, and they are positional.

```plaintext
scripts/
├── <command>.sh            # One entry point per Make target
└── lib/
    ├── log.sh              # Colour decision, log lines, banner
    ├── env.sh              # Toolchain and project-directory checks
    └── gotest.sh           # The gotestsum runs behind the test targets
```

## 📦 Development Scripts

### install_dependencies.sh <a id="install_dependencies.sh"></a> <!-- anchor for internal links -->

install_dependencies.sh installs or updates development tools required for the project. It checks for necessary CLI tools, downloads dependencies, and ensures the environment is ready for development.

#### Usage

Takes no arguments. Every tool is installed at `@latest`.

#### Sample Output

<details><summary>Click to expand sample output</summary>

```bash
❯ scripts/install_dependencies.sh
Validating CLI tools (level: minimal)...
✓ go

✓ All 1 required CLI tools are available
======================================
        Cisco WNC Dependencies
          Module Management
======================================

ℹ Dependencies Info: Using Go version: go1.25.1
[1] Downloading dependencies...
✓ Dependencies Success: Dependencies tidied
✓ Dependencies Success: Dependencies downloaded

[✓] Dependencies management completed
```

</details>

## 🧪 Testing Scripts

### test_unit.sh <a id="test_unit.sh"></a> <!-- anchor for internal links -->

Runs the unit tests. Coverage and the HTML report belong to `test_coverage.sh`.

#### Usage

Takes no arguments. `lib/gotest.sh` sets the 30s timeout.

#### Sample Output

<details><summary>Click to expand sample output</summary>

```bash
❯ export WNC_CONTROLLER=""
❯ export WNC_ACCESS_TOKEN=""
❯ scripts/test_unit.sh
Validating CLI tools (level: standard)...
<snip>

✓ All 5 required CLI tools are available
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

### test_coverage.sh <a id="test_coverage.sh"></a> <!-- anchor for internal links -->

Runs the unit tests with a coverprofile, then renders the HTML report and the artifact
`.octocov.yml` reads for the README badge.

#### Usage

Three optional positional arguments, so a verification run can redirect every output away
from the tracked defaults.

```bash
scripts/test_coverage.sh [coverprofile] [html] [report]
```

| Position | Default                 |
| -------- | ----------------------- |
| 1        | `./tmp/coverage.out`    |
| 2        | `./coverage/report.html`|
| 3        | `./coverage/report.out` |

#### Sample Output

<details><summary>Click to expand sample output</summary>

```bash
❯ scripts/test_coverage.sh
Validating CLI tools (level: standard)...
<snip>

✓ All 5 required CLI tools are available
======================================
       Cisco WNC Coverage Tests
         Go Testing Framework
======================================

→ Starting coverage tests...
<snip>

-----------------------------------------
✓ Coverage tests completed successfully
ℹ Info: Duration: 4s
-----------------------------------------

ℹ Info: Coverage report generated: ./tmp/coverage.out
ℹ Info: Total coverage: 91.0%
→ Generating HTML coverage report...

✓ HTML coverage report generated successfully
ℹ Info: Report location: ./coverage/report.html
ℹ Info: Report size: 322732 bytes

ℹ Info: To view the report:
  open ./coverage/report.html
✓ Coverage report generated successfully
```

</details>

### test_integration.sh <a id="test_integration.sh"></a> <!-- anchor for internal links -->

Runs integration tests against a live Cisco C9800 controller. Requires `WNC_CONTROLLER` and `WNC_ACCESS_TOKEN`.

#### Usage

Takes no arguments. `lib/gotest.sh` sets the 10m timeout.

#### Sample Output

<details><summary>Click to expand sample output</summary>

```bash
❯ export WNC_CONTROLLER="<controller-host-or-ip>"
❯ export WNC_ACCESS_TOKEN="<base64-username:password>"
❯ scripts/test_integration.sh
Validating CLI tools (level: standard)...
<snip>

✓ All 5 required CLI tools are available
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

Formats the working tree with goimports and gofumpt, then lints Go, shell and Markdown.

#### Usage

Takes no arguments. The project root is the directory above `scripts/`.

#### Sample Output

<details><summary>Click to expand sample output</summary>

```bash
❯ scripts/lint.sh
Validating CLI tools (level: standard)...
<snip>

✓ All 5 required CLI tools are available
======================================
        Cisco WNC Code Linter
      golangci-lint Integration
======================================

Validating Go module...
✓ Go module validated
ℹ Info: Starting code linting...
Running formatting operations...
<snip>

--- Formatting Summary ---
All formatting operations completed successfully

Running linting operations...
<snip>

--- Linting Summary ---
All linting checks passed successfully

✓ Success: Lint passed for .
```

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
    Every script takes no arguments.

    Available scripts:
    - install_dependencies.sh Install Go dependencies
    - lint.sh                Run golangci-lint
    - test_unit.sh           Run unit tests
    - test_coverage.sh       Run unit tests with coverage and an HTML report
    - test_integration.sh    Run integration tests

PROJECT STRUCTURE:
    scripts/                Script directory
    +-- <command>.sh       One entry point per Make target
    +-- lib/               Sourced explicitly by the entry points
        +-- log.sh         Colour, log lines and the banner
        +-- env.sh         Toolchain and project-directory checks
        +-- gotest.sh      The gotestsum runs behind the test targets
````

</details>
