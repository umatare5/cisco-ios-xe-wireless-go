# 🧪 Testing

This guide explains the testing strategy, conventions, and execution procedures for the Cisco IOS-XE Wireless Go SDK.

> [!NOTE]
> Integration tests require an accessible Cisco C9800 and these variables: See [Prerequisites](#-prerequisites)

## 🎯 Testing Strategy

### Test Categories

The SDK implements **standardized test patterns** using the unified `pkg/testutil` API:

| Category             | Purpose                   | Implementation Pattern        | Coverage Target |
| -------------------- | ------------------------- | ----------------------------- | --------------- |
| **1. Service Tests** | Service construction      | Direct service instantiation  | 100%            |
| **2. Get Tests**     | Mock-based GET operations | `testutil.NewMockServer()`    | Get/List: 100%  |
| **3. Set Tests**     | Mock-based RPC operations | `testutil.NewMockServer()`    | Set/Admin: 90%+ |
| **4. Integration**   | Live WNC GET operations   | Integration test suites       | N/A             |
| **5. Scenario/E2E**  | Live WNC RPC operations   | Scenario-based test workflows | N/A             |

### IOS-XE Version Support

| IOS-XE Version | Service Support                            | Test Strategy                                 |
| -------------- | ------------------------------------------ | --------------------------------------------- |
| **17.12.x**    | Core services (AP, WLAN, Client, RF, etc.) | Full mock + integration testing               |
| **17.18.x**    | Advanced services (WAT, URWB, Spaces)      | 404 error expectation tests + real data mocks |

> [!NOTE]
> WAT (Wireless Assurance Testing), URWB (Ultra-Reliable Wireless Backhaul), and Spaces services require IOS-XE 17.18.1+. Tests expect 404 responses when services are not configured and use real WNC data structure for mock responses.

### Coverage Requirements

- **Repository overall**: **80% minimum**
- **Service package** (`service/`): **90% minimum**

## 📂 Test Organization

### Directory Structure

```text
cisco-ios-xe-wireless-go/
├── service/
│   └── {service}/
│       ├── service_test.go         # Direct service tests using pkg/testutil
│       ├── service.go              # Service implementation
│       ├── errors.go              # Service-specific error constants
│       └── doc.go                 # Package documentation
├── tests/
│   ├── integration/
│   │   ├── {service}_service_test.go  # Live WNC integration tests per service
│   │   └── ...                       # Additional integration tests
│   └── scenario/                     # E2E scenario tests
│       ├── ap/                       # AP scenario tests
│       ├── rf/                       # RF scenario tests
│       ├── site/                     # Site scenario tests
│       └── wlan/                     # WLAN scenario tests
└── pkg/
    └── testutil/
        ├── testing.go              # Main testing utilities and mock server
        ├── context.go              # Test context management
        └── doc.go                  # Package documentation
```

### Naming Conventions

**Test Functions (New Unified Naming Convention):**

```go
// Unit tests (service/ directory)
TestXServiceUnit_Constructor_Success            // Service construction tests
TestXServiceUnit_GetOperations_MockSuccess     // GET operations with mock server
TestXServiceUnit_GetOperations_ErrorHandling   // GET error scenarios and edge cases
TestXServiceUnit_SetOperations_MockSuccess     // SET/RPC operations with mock server
TestXServiceUnit_SetOperations_ValidationErrors // SET validation and edge cases
TestXServiceUnit_GetOperations_FilteredSuccess // Filtered GET operations
TestXServiceUnit_ValidationErrors              // Input validation tests
TestXServiceUnit_EdgeCases_MockSuccess         // Edge cases and error branches

// Integration tests (tests/integration/ directory)
TestXServiceIntegration_GetOperationalOperations_Success // Live WNC GET operations
TestXServiceIntegration_GetConfigurationOperations_Success // Live WNC configuration retrieval
```

**Examples:**

- `TestApServiceUnit_Constructor_Success` - AP service construction using direct instantiation
- `TestApServiceUnit_GetOperations_MockSuccess` - AP GET operations with mock server
- `TestApServiceUnit_SetOperations_ValidationErrors` - AP SET validation and edge cases
- `TestClientServiceIntegration_GetOperationalOperations_Success` - Client operational data retrieval with live controller

## 🧰 Prerequisites

### For Unit Tests (Layers 1-3)

Unit tests require no special configuration and can be run in any Go development environment.

| Requirement | Version | Notes                              |
| ----------- | ------- | ---------------------------------- |
| Go          | 1.25+   | Uses stdlib testing + pkg/testutil |
| make        | Latest  | Convenience targets                |

### For Integration/E2E Tests (Layers 4-5)

#### 1. Cisco Catalyst 9800 Wireless Network Controller

Integration and E2E tests require a real Cisco Catalyst 9800 WNC. Please refer to [References](#references).

#### 2. Environment Variables

| Variable                | Description            | Example                 |
| ----------------------- | ---------------------- | ----------------------- |
| `WNC_CONTROLLER`        | Controller host/IP     | `wnc1.example.internal` |
| `WNC_ACCESS_TOKEN`      | Base64 `user:pass`     | `YWRtaW46cGFzc3dvcmQ=`  |
| `WNC_AP_MAC_ADDR`       | Test AP's Radio MAC    | `aa:bb:cc:dd:ee:f0`     |
| `WNC_CLIENT_MAC_ADDR`   | Test Client MAC        | `11:22:33:aa:bb:cc`     |
| `WNC_AP_WLAN_BSSID`     | Test AP WLAN BSSID     | `aa:bb:cc:dd:ee:f1`     |
| `WNC_AP_NEIGHBOR_BSSID` | Test AP Neighbor BSSID | `11:22:33:dd:ee:ff`     |

<details><summary>Environment setup</summary>

```bash
export WNC_CONTROLLER="<controller-host-or-ip>"
export WNC_ACCESS_TOKEN="<base64-username:password>"
export WNC_AP_MAC_ADDR="<test-ap-radio-mac-address>"
export WNC_CLIENT_MAC_ADDR="<test-client-mac-address>"
export WNC_AP_WLAN_BSSID="<test-ap-wlan-bssid>"
export WNC_AP_NEIGHBOR_BSSID="<test-ap-neighbor-bssid>"
```

</details>

> [!TIP]
> Environment variables such as `WNC_AP_MAC_ADDR` and `WNC_CLIENT_MAC_ADDR` can be discovered by running the example commands listed in the [README.md - Usecases](../README.md?#-usecases) section.

> [!CAUTION]
> Never commit real tokens or `.env` files. Please refer to [SECURITY.md](./SECURITY.md).

## 🚀 Running Tests

### Quick Start

```bash
# Run all unit tests
make test-unit

# Run integration tests (requires WNC)
make test-integration
```

### Detailed Test Execution

#### Layer 1: Service Construction Tests

Tests service construction and lifecycle using direct service instantiation.

```bash
go test ./service/ap -run "TestApServiceUnit_Constructor" -v
```

**Example:**

- [`service/ap/service_test.go`](../service/ap/service_test.go)

#### Layer 2 and 3: Mock-based Method Tests

Tests all operations with **unified mock server API** using functional options.

```bash
go test ./service/ap -run "TestApServiceUnit_GetOperations_Mock" -v
```

**Examples:**

- [`service/ap/service_test.go`](../service/ap/service_test.go)
- [`service/wat/service_test.go`](../service/wat/service_test.go)

#### Layer 3: Integration Tests

Tests only GET operations with **live WNC**.

```bash
go test ./tests/integration -tags=integration -v
```

**Example:**

- [`tests/integration/client_service_test.go`](../tests/integration/client_service_test.go)
- [`tests/integration/ap_service_test.go`](../tests/integration/ap_service_test.go)
- [`tests/integration/rrm_service_test.go`](../tests/integration/rrm_service_test.go)

#### Layer 4: E2E Scenario Tests

Non-disruptive CRUD operations with **live WNC**.

```bash
go test ./tests/scenario/ap/ -tags=scenario -v
go test ./tests/scenario/rf/ -tags=scenario -v
go test ./tests/scenario/site/ -tags=scenario -v
go test ./tests/scenario/wlan/ -tags=scenario -v
```

**Example:**

- [`tests/scenario/ap/service_test.go`](../tests/scenario/ap/service_test.go) - AP admin, and radio operations
- [`tests/scenario/site/tag_service_test.go`](../tests/scenario/site/tag_service_test.go) - Site tag operations
- [`tests/scenario/rf/service_test.go`](../tests/scenario/rf/service_test.go) - RF tag operations
- [`tests/scenario/wlan/service_test.go`](../tests/scenario/wlan/service_test.go) - Poliy tag operations

> [!NOTE]
> Tag operations in scenario tests **MUST** use newly created tags to avoid communication impact.

## 📈 Coverage Reports

### Coverage Analysis

Generates coverage reports for unit tests.

```bash
make test-unit-coverage
```

### Coverage Requirements Validation

```bash
# Check service package coverage (must be ≥90%)
go test -cover ./service/...

# Check repository coverage (must be ≥80%)
go test -cover ./...
```

## 📚 Appendix

### Testing Tips

1. **Start with unit tests** - Validate basic functionality first using unified MockServer API
2. **Use real WNC data** - Base mock responses on actual controller data from IOS-XE 17.12.x
3. **Test error scenarios** - IOS-XE 17.18.1+ services may return 404 when not configured
4. **Follow naming conventions** - Use standardized test function names (e.g., `TestXServiceUnit_*`)
5. **Use unified API** - Use `testutil.NewMockServer()` with functional options.
6. **Leverage options** - Combine multiple `MockServerOption`s for complex test scenarios
7. **Coverage-driven development** - Write comprehensive tests to meet coverage targets
8. **Parallel-safe integration** - Mark integration tests with `t.Parallel()` for GET-only operations
9. **Scenario isolation** - Use newly created resources in E2E scenarios to avoid impact

### Troubleshooting

| Issue                  | Solution                                                                                 |
| ---------------------- | ---------------------------------------------------------------------------------------- |
| Missing env vars       | Ensure all required `WNC_*` variables are set                                            |
| Unreachable controller | Verify DNS/IP connectivity                                                               |
| TLS errors             | Check certificate validity; use `WithInsecureSkipVerify` for testing only                |
| Auth failures          | Ensure token is Base64 of `user:pass`                                                    |
| TestClient creation    | Use `testutil.NewTestClient(mockServer)` to create test clients for service construction |

### References

- 📖 [Cisco Catalyst 9800-CL Wireless Controller for Cloud Deployment Guide](https://www.cisco.com/c/en/us/td/docs/wireless/controller/9800/technical-reference/c9800-cl-dg.html)
- 📖 [Cisco Catalyst 9800 Series Wireless Controller Programmability Guide](https://www.cisco.com/c/en/us/td/docs/wireless/controller/9800/programmability-guide/b_c9800_programmability_cg/cisco-catalyst-9800-series-wireless-controller-programmability-guide.html)
- 📖 [YANG Models and Platform Capabilities for Cisco IOS XE 17.12.1](https://github.com/YangModels/yang/tree/main/vendor/cisco/xe/17121#readme)
