# 🧪 Testing

This guide explains how to run unit and integration tests, manage test data, and generate coverage reports.

> [!NOTE]
> Integration tests require an accessible Cisco C9800 and these variables: `WNC_CONTROLLER`, `WNC_ACCESS_TOKEN`.

## 🎯 Scopes

There are four main layers of tests:

- **Unit tests**: These tests validate serialization and deserialization between JSON and Go structs.
- **Table-driven tests**: Multiple test cases are efficiently executed using a table-driven approach.
- **Fail-fast error detection tests**: These tests fail immediately if an unexpected error occurs during execution.
- **Integration tests**: These tests interact with multiple API endpoints to verify API communication and overall functionality.

## 🧰 Prerequisites

Tools and environment required to run unit and integration tests.

### For Unit / Table / Fail-fast Tests

Unit tests require no special configuration and can be run in any Go development environment.

| Requirement | Version | Notes                    |
| ----------- | ------- | ------------------------ |
| Go          | 1.24+   | Uses only stdlib testing |
| make        | Latest  | Convenience targets      |

### For Integration Tests

#### 1. Cisco Catalyst 9800 Wireless Network Controller

Integration tests require a real Cisco Catalyst 9800 WNC. Please refer to [#References](#references).

#### 2. Environment Variables

Integration tests also require the following environment variables:

| Variable           | Description        | Example                 |
| ------------------ | ------------------ | ----------------------- |
| `WNC_CONTROLLER`   | Controller host/IP | `wnc1.example.internal` |
| `WNC_ACCESS_TOKEN` | Base64 `user:pass` | `YWRtaW46cGFzc3dvcmQ=`  |

<details><summary>Environment setup</summary>

```bash
export WNC_CONTROLLER="<controller-host-or-ip>"
export WNC_ACCESS_TOKEN="<base64-username:password>"
```

</details>

> [!CAUTION]
> Never commit real tokens or `.env` files. Please refer to [SECURITY.md](./SECURITY.md).

## 🚀 Running tests

Run unit and integration suites via Make targets for consistent, reproducible results.

Primary Make targets:

| Command                          | Description                           |
| -------------------------------- | ------------------------------------- |
| `make test-unit`                 | Run unit + table + fail-fast suites   |
| `make test-integration`          | Run integration (skips if env unset)  |
| `make test-unit-coverage`        | Run unit tests with coverage analysis |
| `make test-integration-coverage` | Run integration tests with coverage   |
| `make test-coverage-report`      | Generate HTML coverage report         |

> [!NOTE]
> Lint runs automatically where configured; see Make and Scripts references.

## 📊 Test Data Collection

Integration tests persist controller responses to support regression and offline inspection.

| Aspect   | Detail                                 |
| -------- | -------------------------------------- |
| Location | `*/test_data/*.json`                   |
| Format   | Raw controller JSON (pretty if stable) |
| Use      | Schema drift and offline analysis      |

<details><summary>Example layout</summary>

```text
test_data/
├── ap_oper_response.json
├── client_oper_response.json
├── general_cfg_response.json
└── rrm_global_oper_response.json
```

</details>

## 📈 Coverage Reports

Generate coverage summaries and an HTML report to assess tested code paths.

| Command                          | Notes                                                         |
| -------------------------------- | ------------------------------------------------------------- |
| `make test-unit-coverage`        | Writes `./tmp/coverage.out` from unit tests.                  |
| `make test-integration-coverage` | Writes `./tmp/coverage.out` from integration tests.           |
| `make test-coverage-report`      | Generates `report.out` and `report.html` under `./coverage/.` |

> [!NOTE]
> CI publish a coverage badge from `coverage/report.out` when present.

## 📚️ Appendix

### Testing Tips

For efficient testing workflow, start with unit tests and gradually move to integration tests.

1. **Install Dependencies**: `make deps` - Install gotestsum and other development tools.
2. **Unit Tests First**: `make test-unit` - Ensure basic functionality with enhanced output.
3. **Environment Setup**: Configure environment variables for integration tests.
4. **Environment Verification**: Check controller access to verify connectivity and credentials.
5. **Coverage Analysis**: `make test-unit-coverage` or `make test-integration-coverage` - Run tests with coverage analysis.
6. **HTML Coverage Report**: `make test-coverage-report` - Generate detailed HTML coverage report.
7. **Test Data Review**: Examine generated JSON files to understand API response structures for debugging.
8. **Incremental Testing**: Test individual modules to target specific functionality when debugging.
9. **Run Integration Tests**: `make test-integration` - Ensure all functionality works as expected.

> [!TIP]
> For comprehensive testing, run both `make test-unit` and `make test-integration` sequentially to validate all functionality.

### Development Dependencies

The project uses several tools to enhance the testing experience. Install all dependencies with: `make deps`

### Troubleshooting

Common issues and concise fixes for failed or flaky test runs.

- Missing env vars: set `WNC_CONTROLLER` and `WNC_ACCESS_TOKEN` for integration.
- Unreachable controller: verify DNS or use `wnc1.example.internal` or an IP.
- TLS errors: see Security → TLS Verification and avoid disabling checks in prod.
- Auth failures: ensure the token is Base64 of `user:pass` and not expired.
- Flaky tests: re-run with verbose logs; isolate by package using `go test ./pkg`.

### References

These references provide additional information on Cisco Catalyst 9800 WNC and related technologies:

- 📖 [Cisco Catalyst 9800-CL Wireless Controller for Cloud Deployment Guide](https://www.cisco.com/c/en/us/td/docs/wireless/controller/9800/technical-reference/c9800-cl-dg.html)
  - A comprehensive guide for deploying Cisco Catalyst 9800-CL WNC in cloud environments.
  - This includes setup instructions, configuration examples, and best practices.
- 📖 [Cisco Catalyst 9800 Series Wireless Controller Programmability Guide](https://www.cisco.com/c/en/us/td/docs/wireless/controller/9800/programmability-guide/b_c9800_programmability_cg/cisco-catalyst-9800-series-wireless-controller-programmability-guide.html)
  - A guide for programming and automating Cisco Catalyst 9800 WNC.
  - This includes information on RESTCONF APIs, YANG models, and more.
- 📖 [YANG Models and Platform Capabilities for Cisco IOS XE 17.12.1](https://github.com/YangModels/yang/tree/main/vendor/cisco/xe/17121#readme)
  - A repository containing YANG models and platform capabilities for Cisco IOS XE 17.12.1.
  - This is useful for understanding the data structures used in the API.
