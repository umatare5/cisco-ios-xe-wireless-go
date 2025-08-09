# 📋 Scripts Reference

Central reference for all development scripts under `scripts/`. Each entry point focuses on a single concern and delegates shared logic to modular libraries in `scripts/lib/`. All scripts are: idempotent, side‑effect constrained, and exit non‑zero on failure.

- There are five main domains of entry point scripts:

  | Domain         | Entry Point Scripts                                                                                                                                                            |
  | -------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
  | Development    | [install_dependencies.sh](#install_dependencies.sh), [clean_artifacts.sh](#clean_artifacts.sh)                                                                                 |
  | Testing        | [test_unit.sh](#test_unit.sh), [test_integration.sh](#test_integration.sh), [test_coverage.sh](#test_coverage.sh), [generate_coverage_report.sh](#generate_coverage_report.sh) |
  | Quality        | [lint.sh](#lint.sh), [pre_commit_hook.sh](#pre_commit_hook.sh)                                                                                                                 |
  | YANG Operation | [get_yang_models.sh](#get_yang_models.sh), [get_yang_model_details.sh](#get_yang_model_details.sh), [get_yang_statement_details.sh](#get_yang_statement_details.sh)            |
  | Help           | [help.sh](#help.sh)                                                                                                                                                            |

- Makefile targets support all scripts, allowing you to run them.

## 🗂️ Structure

Scripts share a consistent bootstrap pattern:

- Source `lib/bootstrap.sh` in the entry script.
- Call `init_wnc_libraries(<script_dir>, <module_dir>)` to load the target module (e.g., `lib/testing`).
- Expose common predicates, formatters, and validators in the current shell.
- Invoke exactly one exported `run_*_operation` function.
- Keep entry points thin; centralize behavior under `scripts/lib/`.

```text
scripts/
├── <command>.sh            # Thin entry point(s)
└── lib/                    # Reusable modules (loaded via bootstrap)
    ├── bootstrap.sh        # Loader + init
    ├── artifacts/          # Cleanup operations
    ├── coverage/           # Coverage + HTML generation
    ├── dependencies/       # Dependency install/update
    ├── lint/               # Lint operations
    ├── output/             # Banners + formatting helpers
    ├── testing/            # go test orchestration
    ├── utils/              # generic predicates (jq detection, etc.)
    ├── validation/         # git / branch protection helpers
    └── yang/               # RESTCONF + YANG data utilities
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
✓ go
✓ golangci-lint
✓ gotestsum

✓ All 4 required CLI tools are available
======================================
        Cisco WNC Dependencies
          Module Management
======================================

ℹ Dependencies Info: Using Go version: go1.24.4
[2] Downloading dependencies...
✓ Dependencies Success: Dependencies tidied
✓ Dependencies Success: Dependencies downloaded

[✓] Dependencies management completed
```

</details>

### clean_artifacts.sh <a id="clean_artifacts.sh"></a> <!-- anchor for internal links -->

clean_artifacts.sh removes build artifacts, temporary files, and caches to restore a clean working tree. It supports granular flags or a single --all sweep.

#### Usage

```bash
❯ scripts/clean_artifacts.sh --help

USAGE: clean_artifacts [OPTIONS]

OPTIONS:
  -p, --project <DIR>  Project root directory [default: .]
  -v, --verbose        Enable verbose output
  -f, --force          Force removal without confirmation
      --go-cache       Clean Go build cache
      --go-modules     Clean Go module cache
      --temp-files     Clean temporary files (./tmp)
      --test-files     Clean test artifacts (.test binaries, coverage files)
      --all            Clean all artifacts [default: true]
      --dry-run        Show what would be cleaned without actually cleaning
      --no-color       Disable colored output
  -h, --help           Print help
  -V, --version        Print version
```

#### Sample Output

<details><summary>Click to expand sample output</summary>

```bash
❯ scripts/clean_artifacts.sh
Validating CLI tools (level: minimal)...
✓ go

✓ All 1 required CLI tools are available
======================================
         Cisco WNC Artifacts
           Cleanup Utility
======================================

[1] Cleaning Go build cache...
✓ Cleanup Success: Go build cache cleaned ( 12K freed)
[2] Cleaning Go module cache...
✓ Cleanup Success: Go module cache cleaned (330M freed)
[3] Cleaning temporary files...
ℹ Cleanup Info: No temporary directory found: ./tmp
[4] Cleaning test artifacts...
ℹ Cleanup Info: No test artifacts found to clean
[✓] Artifacts cleanup completed successfully
```

</details>

## 🧪 Testing Scripts

### test_unit.sh <a id="test_unit.sh"></a> <!-- anchor for internal links -->

Runs unit tests with optional short mode and coverage generation.

#### Usage

```bash
❯ scripts/test_unit.sh --help

USAGE: test_unit [OPTIONS]

OPTIONS:
  -p, --project <DIR>       Project root directory [default: .]
  -v, --verbose             Enable verbose test output
  -s, --short               Run tests in short mode (skip long-running tests)
  -c, --coverage            Generate coverage data
  -t, --timeout <DURATION>  Test timeout duration [default: 30s]
      --no-color            Disable colored output
  -h, --help                Print help
  -V, --version             Print version
```

#### Sample Output

<details><summary>Click to expand sample output</summary>

```bash
❯ scripts/test_unit.sh
Validating CLI tools (level: standard)...
✓ curl
✓ go
✓ golangci-lint
✓ gotestsum

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

DONE 1048 tests, 36 skipped in 12.783s

-----------------------------------------
✓ Unit tests completed successfully
ℹ Info: Duration: 29s
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
❯ scripts/test_integration.sh
Validating CLI tools (level: standard)...
✓ curl
✓ go
✓ golangci-lint
✓ gotestsum

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

DONE 1048 tests, 36 skipped in 12.783s

-----------------------------------------
✓ Unit tests completed successfully
ℹ Info: Duration: 29s
-----------------------------------------
```

</details>

### test_coverage.sh <a id="test_coverage.sh"></a> <!-- anchor for internal links -->

Runs all tests and writes a unified coverage profile to `./tmp/coverage.out` (overridable).

#### Usage

```bash
❯ scripts/test_coverage.sh --help

USAGE: test_coverage [OPTIONS]

OPTIONS:
  -p, --project <DIR>       Project root directory [default: .]
  -o, --output <FILE>       Coverage output file [default: ./tmp/coverage.out]
  -v, --verbose             Enable verbose test output
  -s, --short               Run tests in short mode (skip long-running tests)
  -t, --timeout <DURATION>  Test timeout duration [default: 30s]
      --no-color            Disable colored output
  -h, --help                Print help
  -V, --version             Print version
```

#### Sample Output

<details><summary>Click to expand sample output</summary>

```bash
❯ scripts/test_coverage.sh
Validating CLI tools (level: standard)...
✓ curl
✓ go
✓ golangci-lint
✓ gotestsum

✓ All 4 required CLI tools are available
======================================
       Cisco WNC Coverage Tests
         Go Testing Framework
======================================

→ Starting coverage tests...
PASS TestNewClient/ValidClient (0.00s)
PASS TestNewClient/ValidClientWithOptions (0.00s)
PASS TestNewClient/InvalidHost (0.00s)
<snip>

DONE 1048 tests, 36 skipped in 11.167s

-----------------------------------------
✓ Coverage tests completed successfully
ℹ Info: Duration: 13s
-----------------------------------------

ℹ Info: Coverage report generated: ././tmp/coverage.out
ℹ Info: Total coverage: 99.1%
```

</details>

### generate_coverage_report.sh <a id="generate_coverage_report.sh"></a> <!-- anchor for internal links -->

Generates an HTML coverage report from `coverage.out`.

#### Usage

```bash
❯ scripts/generate_coverage_report.sh --help

USAGE: generate_coverage_report [OPTIONS]

OPTIONS:
  -p, --project <DIR>  Project root directory [default: .]
  -i, --input <FILE>   Coverage input file [default: ./tmp/coverage.out]
  -o, --output <FILE>  HTML output file [default: ./tmp/coverage.html]
  -v, --verbose        Enable verbose output
      --no-color       Disable colored output
  -h, --help           Print help
  -V, --version        Print version
```

#### Sample Output

<details><summary>Click to expand sample output</summary>

```bash
❯ scripts/generate_coverage_report.sh
Validating CLI tools (level: standard)...
✓ curl
✓ go
✓ golangci-lint
✓ gotestsum

✓ All 4 required CLI tools are available
======================================
       Coverage HTML Generator
      Go Tool Cover Integration
======================================

→ Generating HTML coverage report...

✓ HTML coverage report generated successfully
ℹ Info: Report location: ././coverage/report.html
ℹ Info: Report size: 159374 bytes

ℹ Info: To view the report:
  open ././coverage/report.html
```

</details>

## ✅ Quality Scripts

### lint.sh <a id="lint.sh"></a> <!-- anchor for internal links -->

Runs golangci-lint using the repo configuration. Supports optional auto-fix.

#### Usage

`scripts/lint.sh` only supports execution with no arguments.

````bash

#### Sample Output

<details><summary>Click to expand sample output</summary>

```bash
❯ scripts/lint.sh
Validating CLI tools (level: standard)...
✓ curl
✓ go
✓ golangci-lint
✓ gotestsum

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

### pre_commit_hook.sh <a id="pre_commit_hook.sh"></a> <!-- anchor for internal links -->

Runs repository pre-commit validations (formatting, build, tests, coverage presence). Intended to be wired to git hooks or run ad-hoc.

#### Usage

`pre_commit_hook.sh` only supports execution with no arguments.

#### Sample Output

<details><summary>Click to expand sample output</summary>

```bash
❯ scripts/pre_commit_hook.sh
Validating CLI tools (level: minimal)...
✓ go

✓ All 1 required CLI tools are available
======================================
        Pre-commit Validation
          Branch Protection
======================================


ℹ Info: Current branch: umatare5/road_to_0.2.0
⚠ Warning: No staged changes found
ℹ Info: Use 'git add <files>' to stage changes before committing
✓ Success: Pre-commit validation passed
ℹ Info: Proceeding with commit on branch 'umatare5/road_to_0.2.0'
```

</details>

## 📡 YANG Operation Scripts

### get_yang_models.sh <a id="get_yang_models.sh"></a> <!-- anchor for internal links -->

Lists available Cisco wireless YANG models from the controller.

#### Usage

```bash
❯ scripts/get_yang_models.sh --help

USAGE: get_yang_models [OPTIONS]

OPTIONS:
  -c, --controller <HOST>    WNC controller hostname or IP (required unless WNC_CONTROLLER set)
  -t, --token <TOKEN>        Basic auth token (or use WNC_ACCESS_TOKEN env var)
  -p, --protocol <PROTOCOL>  Protocol: http or https [default: https] [choices: http,https]
  -k, --insecure             Skip TLS certificate verification
  -v, --verbose              Enable verbose output
      --no-color             Disable colored output
  -h, --help                 Print help
  -V, --version              Print version
```

#### Sample Output

<details><summary>Click to expand sample output</summary>

```bash
N/A
```

</details>

### get_yang_model_details.sh <a id="get_yang_model_details.sh"></a> <!-- anchor for internal links -->

Fetches and prints details for a specific YANG model.

#### Usage

```bash
❯ scripts/get_yang_model_details.sh --help

USAGE: get_yang_model_details [OPTIONS] <MODEL>

OPTIONS:
  -c, --controller <HOST>    WNC controller hostname or IP (required unless WNC_CONTROLLER set)
  -t, --token <TOKEN>        Basic auth token (or use WNC_ACCESS_TOKEN env var)
  -p, --protocol <PROTOCOL>  Protocol: http or https [default: https] [choices: http,https]
  -f, --format <FORMAT>      Output format: json or xml [default: json] [choices: json,xml]
  -k, --insecure             Skip TLS certificate verification
  -v, --verbose              Enable verbose output
  -r, --raw                  Output raw response without formatting
      --no-color             Disable colored output
  -h, --help                 Print help
  -V, --version              Print version
```

#### Sample Output

<details><summary>Click to expand sample output</summary>

```json
{
  "model": "Cisco-IOS-XE-wireless-ap-oper",
  "namespace": "http://cisco.com/ns/yang/Cisco-IOS-XE-wireless-ap-oper",
  "revision": "2023-10-01"
}
```

</details>

### get_yang_statement_details.sh <a id="get_yang_statement_details.sh"></a> <!-- anchor for internal links -->

Retrieves details for a specific statement under a given YANG model.

#### Usage

```bash
❯ scripts/get_yang_statement_details.sh --help

USAGE: get_yang_statement_details [OPTIONS] <model> <statement>

OPTIONS:
  -c, --controller <HOST>  Controller hostname or IP (or WNC_CONTROLLER)
  -t, --token <TOKEN>      Basic auth token (or WNC_ACCESS_TOKEN)
  -p, --protocol <PROTO>   http or https [default: https]
  -f, --format <FORMAT>    json or xml [default: json]
  -k, --insecure           Skip TLS certificate verification
  -v, --verbose            Enable verbose output
      --no-color           Disable colored output
  model                    YANG model name (required)
  statement                YANG statement name (required)
```

#### Sample Output

<details><summary>Click to expand sample output</summary>

```json
N/A
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

```text
✗ scripts/help.sh
🔧 Cisco WNC Development Scripts
-------------------------------

USAGE:
    make <target>                   # Use Makefile targets (recommended)
    ./scripts/<script>.sh [options] # Use scripts directly

COMMON DEVELOPMENT TARGETS:
    help                Show this help message
    clean               Clean build artifacts and temporary files
    deps                Install development dependencies
    lint                Run code linting tools
    build               Verify build compilation
    test-unit           Run unit tests only
    test-integration    Run integration tests (requires environment)
    test-coverage       Run tests with coverage analysis
    test-coverage-report Generate HTML coverage report

YANG MODEL DEVELOPMENT:
    yang-list           List all available YANG models
    yang-model          Get YANG model details (MODEL=model-name)
    yang-statement      Get YANG statement details (MODEL=model-name STATEMENT=statement-name)

ENVIRONMENT VARIABLES:
    WNC_CONTROLLER      Controller hostname/IP for integration tests
    WNC_ACCESS_TOKEN    Base64 encoded credentials for integration tests

EXAMPLES:
    # Basic development workflow
    make deps               # Install dependencies
    make lint               # Check code quality
    make test-unit          # Run unit tests
    make test-coverage      # Run tests with coverage
    make build              # Verify compilation

    # YANG development
    make yang-list                                    # List models
    make yang-model MODEL=wireless-access-point      # Get model details
    make yang-statement MODEL=wireless-client STATEMENT=active # Get statement details

    # Integration testing (requires environment setup)
    export WNC_CONTROLLER="<controller-hostname>"
    export WNC_ACCESS_TOKEN="YWRtaW46Y3l0WU43WVh4M2swc3piUnVhb1V1ZUx6"
    make test-integration

SCRIPT DETAILS:
    For specific script options and advanced usage:
    ./scripts/<script_name>.sh --help

    Available scripts:
    - clean_artifacts.sh      Clean build artifacts
    - install_dependencies.sh Install Go dependencies
    - lint.sh                Run golangci-lint
    - test_unit.sh           Run unit tests
    - test_integration.sh    Run integration tests
    - test_coverage.sh       Run coverage tests
    - generate_coverage_report.sh Generate HTML coverage
    - get_yang_models.sh     List YANG models
    - get_yang_model_details.sh Get model details
    - get_yang_statement_details.sh Get statement details

PROJECT STRUCTURE:
    scripts/                Script directory
    +-- lib/               Shared libraries
    |   +-- bootstrap.sh   Bootstrap library loader
    |   +-- coverage/      Coverage report functions
    |   +-- dependencies/  Dependency management
    |   +-- output/        Output formatting utilities
    |   +-- testing/       Test utilities
    |   +-- utils/         Utility functions
    |   +-- validation/    Git commit validation
    |   +-- yang/          YANG-specific functions
    +-- *.sh               Entry point scripts

This project uses a modular script architecture with shared libraries
for maintainability and consistency across all development operations.
````

</details>
