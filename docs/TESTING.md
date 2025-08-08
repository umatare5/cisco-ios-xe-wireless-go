# 🧪 Testing

Test suite focuses on correctness of serialization, basic domain wiring, and live integration against a real controller (when env vars set). Only deterministic, read‑only operations are executed.

| Layer        | Purpose                                | Trigger                 |
| ------------ | -------------------------------------- | ----------------------- |
| Unit         | JSON ↔ Go struct fidelity              | `make test-unit`        |
| Table-driven | Multiple scenario assertions           | `make test-unit`        |
| Fail-fast    | Immediate surface of unexpected errors | `make test-unit`        |
| Integration  | Live controller data shape sanity      | `make test-integration` |

## 🎯 Prerequisites

### Unit / Table / Fail-fast

| Requirement | Version | Notes                    |
| ----------- | ------- | ------------------------ |
| Go          | 1.24+   | Uses only stdlib testing |
| make        | Latest  | Convenience targets      |

### Integration

Requires reachable Catalyst 9800 controller + credentials:

| Variable           | Description        | Example                 |
| ------------------ | ------------------ | ----------------------- |
| `WNC_CONTROLLER`   | Controller host/IP | `wnc1.example.internal` |
| `WNC_ACCESS_TOKEN` | Base64 `user:pass` | `YWRtaW46cGFzc3dvcmQ=`  |

<details><summary>Environment Variable Example</summary>

```bash
export WNC_CONTROLLER="192.168.1.100"          # Your WNC IP address
export WNC_ACCESS_TOKEN="YWRtaW46cGFzc3dvcmQ=" # Base64 encoded username:password
```

</details>

## 🚀 Running Tests

Use provided Make targets:

| Command                 | Description                          |
| ----------------------- | ------------------------------------ |
| `make test-unit`        | Run unit + table + fail-fast suites  |
| `make test-integration` | Run integration (skips if env unset) |

> [!NOTE]
> lint runs automatically where defined.

<details><summary>Sample Output: `test-integration`</summary>

```bash
❯ make test-unit
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
DONE 1048 tests, 36 skipped in 4.194s

-----------------------------------------
✓ Unit tests completed successfully
ℹ Info: Duration: 4s
-----------------------------------------
```

</details>

<details><summary>Sample Output: `test-unit`</summary>

```bash
❯ make test-integration
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
PASS TestNewClient/ValidClient (0.00s)
PASS TestNewClient/ValidClientWithOptions (0.00s)
PASS TestNewClient/InvalidHost (0.00s)
<snip>
DONE 1048 tests, 36 skipped in 9.386s

-----------------------------------------
✓ Integration tests completed successfully
ℹ Info: Duration: 10s
-----------------------------------------
```

</details>

<details><summary>Sample Output: `test-integration`</summary>

```bash
❯ make test-integration
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
PASS TestNewClient/ValidClient (0.00s)
PASS TestNewClient/ValidClientWithOptions (0.00s)
PASS TestNewClient/InvalidHost (0.00s)
<snip>
DONE 1048 tests, 36 skipped in 9.386s

-----------------------------------------
✓ Integration tests completed successfully
ℹ Info: Duration: 10s
-----------------------------------------
```

</details>

## 📊 Test Data Artifacts

Integration tests persist sample controller responses for regression & offline parsing.

| Aspect   | Detail                                     |
| -------- | ------------------------------------------ |
| Location | `*/test_data/*.json`                       |
| Format   | Raw controller JSON (pretty if stable)     |
| Use      | Schema drift detection, offline inspection |

<details><summary>Example Layout</summary>

```text
test_data/
├── ap_oper_response.json
├── client_oper_response.json
├── general_cfg_response.json
└── rrm_global_oper_response.json
```

</details>

## 📈 Coverage

### Reports

| Target  | Command                     | Notes                                 |
| ------- | --------------------------- | ------------------------------------- |
| Summary | `make test-coverage`        | Outputs aggregate %                   |
| HTML    | `make test-coverage-report` | Generates browsable HTML under `tmp/` |

<details><summary>Sample Output: test-coverage</summary>

```bash
❯ make test-coverage
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
DONE 1048 tests, 36 skipped in 10.136s

-----------------------------------------
✓ Coverage tests completed successfully
ℹ Info: Duration: 11s
-----------------------------------------

ℹ Info: Coverage report generated: ././tmp/coverage.out
ℹ Info: Total coverage: 99.1%
```

</details>

<details><summary>Sample Output: test-coverage-report</summary>

```bash
❯ make test-coverage-report
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
ℹ Info: Report location: ././tmp/coverage.html
ℹ Info: Report size: 159374 bytes

ℹ Info: To view the report:
  open ././tmp/coverage.html
```

</details>

## � Tips

1. Run unit first (`make test-unit`)
2. Export env vars only when needed for integration
3. Use `grep` over `test_data/` to compare schema drift
4. Keep JSON fixtures minimal—avoid giant blobs
5. Fail fast: add explicit error checks in new tests

> [!TIP]
> CI can run unit tests without controller; integration can be opt‑in nightly.

### 🔧 Tooling

The project uses several tools to enhance the testing experience:

| Tool            | Purpose                            |
| --------------- | ---------------------------------- |
| `golangci-lint` | Static analysis                    |
| `gotestsum`     | (Optional) nicer local test output |
| `markdownlint`  | Doc lint (indirect via scripts)    |

> [!NOTE]
> Install dev helpers: `make deps`
