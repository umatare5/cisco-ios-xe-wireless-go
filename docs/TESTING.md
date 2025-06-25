# 🧪 Testing

This library includes following tests that validate API functionality and client behavior:

- **Unit tests**: These tests validate serialization and deserialization between JSON and Go structs.
- **Table-driven tests**: Multiple test cases are efficiently executed using a table-driven approach.
- **Fail-fast error detection tests**: These tests fail immediately if an unexpected error occurs during execution.
- **Integration tests**: These tests interact with multiple API endpoints to verify API communication and overall functionality.

> [!Note]
> Currently, the test coverage is insufficient. All tests will be covered in the future release `v0.3.0`.

## 🎯 Prerequisites

### For Unit, Table-driven and Fail-fast Tests

Unit tests require no special configuration and can be run in any Go development environment.

| Requirement   | Version/Details  | Description                                          |
| ------------- | ---------------- | ---------------------------------------------------- |
| Go            | 1.24 or later    | Required for running tests and building the project. |
| Testing Tools | Standard library | Built-in Go testing framework.                       |

### For Integration Tests

#### 1. Cisco Catalyst 9800 Wireless Network Controller

Integration tests require a real Cisco Catalyst 9800 WNC. For instructions on setting up WNC, please refer to [References Section](#references).

#### 2. Environment Variables

Integration tests also require the following environment variables:

| Variable           | Description                | Example                |
| ------------------ | -------------------------- | ---------------------- |
| `WNC_CONTROLLER`   | WNC IP address or hostname | `192.168.1.100`        |
| `WNC_ACCESS_TOKEN` | Base64 encoded credentials | `YWRtaW46cGFzc3dvcmQ=` |

<details><summary>Example of Shell Configuration</summary>

```bash
export WNC_CONTROLLER="192.168.1.100"          # Your WNC IP address
export WNC_ACCESS_TOKEN="YWRtaW46cGFzc3dvcmQ=" # Base64 encoded username:password
```

</details>

## 🚀 Running Tests

The project includes convenient Makefile targets for testing:

| Command                 | Description                                          |
| ----------------------- | ---------------------------------------------------- |
| `make test-unit`        | Run unit tests only.                                 |
| `make test-integration` | Run integration tests. \* Requires the access to WNC |

<details><summary>Example of Test Output</summary>

```bash
# Successful test run
PASS: TestNewClient (0.01s)
PASS: TestGetApOper_Integration (2.34s)
PASS: TestClientOperations_Integration (1.87s)

# Skipped integration tests (missing environment)
SKIP: TestGetApOper_Integration (environment variables not set)
SKIP: TestClientOperations_Integration (environment variables not set)

# Failed integration test
FAIL: TestGetApOper_Integration (2.34s)
    client_test.go:45: Failed to connect to controller: dial timeout
```

</details>

## 📊 Test Data Collection

Integration tests automatically collect and save real WNC data to JSON files for validation and debugging purposes.

- **Location**: `test_data/` directory in each module
- **Format**: JSON files with descriptive names (e.g., `ap_oper_response.json`)
- **Purpose**: Verify API response structure and enable offline debugging

<details><summary>Example of test data tree structure</summary>

```text
test_data/
├── ap_oper_response.json
├── client_oper_response.json
├── general_cfg_response.json
└── rrm_global_oper_response.json
```

</details>

## 📚️ Appendix

### Testing Tips

For efficient testing workflow, start with unit tests and gradually move to integration tests.

1. **Unit Tests First**: `make test-unit` - Ensure basic functionality.
2. **Environment Verification**: Check controller access to verify connectivity and credentials.
3. **Test Data Review**: Examine generated JSON files to understand API response structures for debugging.
4. **Incremental Testing**: Test individual modules to target specific functionality when debugging.
5. **Run Integration Tests**: `make test-integration` - Ensure all functionality works as expected.
6. **Run All Tests**: `make test` - Run all tests including unit and integration tests.

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
