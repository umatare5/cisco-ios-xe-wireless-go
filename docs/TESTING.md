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

<details><summary>Env Example</summary>

```bash
export WNC_CONTROLLER="192.168.1.100"          # Your WNC IP address
export WNC_ACCESS_TOKEN="YWRtaW46cGFzc3dvcmQ=" # Base64 encoded username:password
```

</details>

## 🚀 Running Tests

Use provided Make targets (lint runs automatically where defined):

| Command                 | Description                          |
| ----------------------- | ------------------------------------ |
| `make test-unit`        | Run unit + table + fail-fast suites  |
| `make test-integration` | Run integration (skips if env unset) |
| `make test-coverage`    | Combined coverage generation         |

<details><summary>Sample Output (Truncated)</summary>

```text
📦 github.com/umatare5/cisco-ios-xe-wireless-go (42.9% coverage)
  ✅ TestClientConfig (0s)
  ✅ TestClientFunctions (10.67s)
  ✅ TestClientFunctions/GET_APOper (5.63s)
    client_test.go:399: GET AP oper request successful
  🚧 TestClientFunctions (0s)
    client_test.go:364: WNC_CONTROLLER and WNC_ACCESS_TOKEN environment variables must be set for integration tests

📦 github.com/umatare5/cisco-ios-xe-wireless-go/ap (1.1% coverage)
  ✅ TestApOperationFailFast/NilClient (0s)
    oper_test.go:210: Correctly returned error with nil client: invalid client configuration: client cannot be nil
  🚧 TestAPConfigurationFunctions (0s)
    cfg_test.go:48: Required environment variables not set - skipping test

📦 github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil
    coverage: 0.0% of statements
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

<details><summary>Sample Coverage Output</summary>

```text
Coverage report generated at ./tmp/coverage.out
total: (statements) 6.1%

📦 github.com/umatare5/cisco-ios-xe-wireless-go (42.9% coverage)
📦 github.com/umatare5/cisco-ios-xe-wireless-go/awips (75% coverage)
📦 github.com/umatare5/cisco-ios-xe-wireless-go/ap (1.1% coverage)
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

### References

| Resource                            | Focus                 |
| ----------------------------------- | --------------------- |
| Catalyst 9800 Programmability Guide | RESTCONF & automation |
| YANG Models (17.12.1)               | Data model reference  |

---

**Back to:** [API Reference](API_REFERENCE.md) | [Security](SECURITY.md)
