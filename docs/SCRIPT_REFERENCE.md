# 📋 Scripts Reference

Central reference for all development scripts under `scripts/`. Each entry point focuses on a single concern and delegates shared logic to modular libraries in `scripts/lib/`. All scripts are: idempotent, side‑effect constrained, and exit non‑zero on failure.

- There are five main domains of scripts:

  | Domain         | Scripts (entry points)                                                                   |
  | -------------- | ---------------------------------------------------------------------------------------- |
  | Development    | `install_dependencies.sh`, `clean_artifacts.sh`                                          |
  | Testing        | `test_unit.sh`, `test_integration.sh`, `test_coverage.sh`, `generate_coverage_report.sh` |
  | Quality        | `lint.sh`, `pre_commit_hook.sh`                                                          |
  | YANG Operation | `get_yang_models.sh`, `get_yang_model_details.sh`, `get_yang_statement_details.sh`       |
  | Help           | `help.sh`                                                                                |

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

### install_dependencies.sh

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
❯ make deps
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
ℹ Dependencies Info: Direct dependencies: 1
ℹ Dependencies Info: Total dependencies: 0
```

</details>

### clean_artifacts.sh

clean_artifacts.sh removes build artifacts, temporary files, and caches to restore a clean working tree. It supports granular flags or a single --all sweep.

#### Usage

```bash
❯ scripts/clean_artifacts.sh --help

USAGE: clean_artifacts [OPTIONS]

OPTIONS:
  -p, --project <DIR>   Project root directory [default: .]
  -v, --verbose         Enable verbose output
  -f, --force           Force removal without confirmation
      --go-cache        Clean Go build cache
      --go-modules      Clean Go module cache
      --temp-files      Clean temporary files (./tmp)
      --test-files      Clean test artifacts (.test, coverage files)
      --all             Clean all artifacts [default: true]
      --dry-run         Show what would be cleaned
      --no-color        Disable colored output
```

#### Sample Output

<details><summary>Click to expand sample output</summary>

```bash
❯ make clean
🧹 Cleaning artifacts...
• Removing ./tmp/*
• Removing coverage/*.html
• go clean -testcache
✓ Cleanup completed
```

</details>

## 🧪 Testing Scripts

### test_unit.sh

Runs unit tests with optional short mode and coverage generation.

#### Usage

```bash
❯ scripts/test_unit.sh --help

USAGE: test_unit [OPTIONS]

OPTIONS:
  -p, --project <DIR>   Project root directory [default: .]
  -v, --verbose         Enable verbose test output
  -s, --short           Run tests in short mode
  -c, --coverage        Generate coverage data
  -t, --timeout <DUR>   Test timeout [default: 30s]
      --no-color        Disable colored output
```

#### Sample Output

<details><summary>Click to expand sample output</summary>

```bash
❯ make test-unit
ok   github.com/umatare5/cisco-ios-xe-wireless-go/radio  0.42s  coverage: 78.3% of statements
ok   github.com/umatare5/cisco-ios-xe-wireless-go/wlan   0.56s  coverage: 81.2% of statements
✓ Unit tests passed
```

</details>

### test_integration.sh

Runs integration tests against a live Cisco C9800 controller. Requires `WNC_CONTROLLER` and `WNC_ACCESS_TOKEN`.

#### Usage

```bash
❯ scripts/test_integration.sh --help

USAGE: test_integration [OPTIONS]

OPTIONS:
  -p, --project <DIR>   Project root directory [default: .]
  -v, --verbose         Enable verbose test output
      --race            Enable race detector [default: true]
  -t, --timeout <DUR>   Test timeout [default: 10m]
      --package <PAT>   Package pattern [default: ./...]
      --check-env-only  Check environment and exit
      --no-color        Disable colored output
```

#### Sample Output

<details><summary>Click to expand sample output</summary>

```bash
❯ export WNC_CONTROLLER=198.51.100.10
❯ export WNC_ACCESS_TOKEN=YWRtaW46cGFzcw==
❯ make test-integration
🔌 Checking controller reachability... OK
running (race enabled) ./...
ok   github.com/umatare5/cisco-ios-xe-wireless-go/client  2.31s
ok   github.com/umatare5/cisco-ios-xe-wireless-go/ap      3.02s
✓ Integration tests passed
```

</details>

### test_coverage.sh

Runs all tests and writes a unified coverage profile to `./tmp/coverage.out` (overridable).

#### Usage

```bash
❯ scripts/test_coverage.sh --help

USAGE: test_coverage [OPTIONS]

OPTIONS:
  -p, --project <DIR>   Project root directory [default: .]
  -o, --output <FILE>   Coverage output [default: ./tmp/coverage.out]
  -v, --verbose         Enable verbose test output
  -s, --short           Run tests in short mode
  -t, --timeout <DUR>   Test timeout [default: 30s]
      --no-color        Disable colored output
```

#### Sample Output

<details><summary>Click to expand sample output</summary>

```bash
❯ make test-coverage
running unit + integration tests...
profile: ./tmp/coverage.out
mode: atomic
✓ Coverage profile generated
```

</details>

### generate_coverage_report.sh

Generates an HTML coverage report from `coverage.out`.

#### Usage

```bash
❯ scripts/generate_coverage_report.sh --help

USAGE: generate_coverage_report [OPTIONS]

OPTIONS:
  -p, --project <DIR>   Project root directory [default: .]
  -i, --input <FILE>    Coverage input [default: ./tmp/coverage.out]
  -o, --output <FILE>   HTML output [default: ./tmp/coverage.html]
  -v, --verbose         Enable verbose output
      --no-color        Disable colored output
```

#### Sample Output

<details><summary>Click to expand sample output</summary>

```bash
❯ make test-coverage-report
Converting ./tmp/coverage.out -> ./tmp/coverage.html
✓ HTML report generated: ./tmp/coverage.html
```

</details>

## ✅ Quality Scripts

### lint.sh

Runs golangci-lint using the repo configuration. Supports optional auto-fix.

#### Usage

```bash
❯ scripts/lint.sh --help

USAGE: lint [OPTIONS]

OPTIONS:
  -p, --project <DIR>   Project root directory [default: .]
  -v, --verbose         Enable verbose output
      --fix             Automatically fix issues where possible
      --config <FILE>   Custom golangci-lint config
      --no-color        Disable colored output
```

#### Sample Output

<details><summary>Click to expand sample output</summary>

```bash
❯ make lint
golangci-lint run ./...
✓ No issues found
```

</details>

### pre_commit_hook.sh

Runs repository pre-commit validations (formatting, build, tests, coverage presence). Intended to be wired to git hooks or run ad-hoc.

#### Usage

```bash
❯ scripts/pre_commit_hook.sh --help

USAGE: pre_commit_hook [OPTIONS]

OPTIONS:
  -v, --verbose         Enable verbose output
      --no-color        Disable colored output
```

#### Sample Output

<details><summary>Click to expand sample output</summary>

```bash
❯ ./scripts/pre_commit_hook.sh
🔎 Checking workspace state
🔧 Running lint
🧪 Running unit tests
✓ Pre-commit validation passed
```

</details>

## 📡 YANG Operation Scripts

### get_yang_models.sh

Lists available Cisco wireless YANG models from the controller.

#### Usage

```bash
❯ scripts/get_yang_models.sh --help

USAGE: get_yang_models [OPTIONS]

OPTIONS:
  -c, --controller <HOST>  Controller hostname or IP (or WNC_CONTROLLER)
  -t, --token <TOKEN>      Basic auth token (or WNC_ACCESS_TOKEN)
  -p, --protocol <PROTO>   http or https [default: https]
  -k, --insecure           Skip TLS certificate verification
  -v, --verbose            Enable verbose output
      --no-color           Disable colored output
```

#### Sample Output

<details><summary>Click to expand sample output</summary>

```bash
❯ make yang-list
Cisco-IOS-XE-wireless-ap-oper
Cisco-IOS-XE-wireless-client-oper
Cisco-IOS-XE-wireless-rrm-oper
```

</details>

### get_yang_model_details.sh

Fetches and prints details for a specific YANG model.

#### Usage

```bash
❯ scripts/get_yang_model_details.sh --help

USAGE: get_yang_model_details [OPTIONS] <model>

OPTIONS:
  -c, --controller <HOST>  Controller hostname or IP (or WNC_CONTROLLER)
  -t, --token <TOKEN>      Basic auth token (or WNC_ACCESS_TOKEN)
  -p, --protocol <PROTO>   http or https [default: https]
  -f, --format <FORMAT>    json or xml [default: json]
  -k, --insecure           Skip TLS certificate verification
  -v, --verbose            Enable verbose output
  -r, --raw                Output raw response
      --no-color           Disable colored output
  model                    YANG model name (required)
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

### get_yang_statement_details.sh

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
{
  "model": "Cisco-IOS-XE-wireless-ap-oper",
  "statement": "access-point-oper-data",
  "children": ["ap-oper-data", "ap-oper-state"]
}
```

</details>

## 🆘 Help Script

### help.sh

Prints a consolidated help guide covering common Make targets, environment variables, and examples.

#### Usage

```bash
❯ scripts/help.sh
```

#### Sample Output

<details><summary>Click to expand sample output</summary>

```text
🔧 Cisco WNC Development Scripts
USAGE:
  make <target>
  ./scripts/<script>.sh [options]

COMMON DEVELOPMENT TARGETS:
  help  clean  deps  lint  build  test-unit  test-integration  test-coverage  test-coverage-report
```

</details>
