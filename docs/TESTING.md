# 🧪 Testing

This guide explains the testing strategy, conventions, and execution procedures for the Cisco IOS-XE Wireless Go SDK.

> [!NOTE]
> Integration tests require an accessible Cisco C9800 and these variables: `WNC_CONTROLLER`, `WNC_ACCESS_TOKEN`, `WNC_AP_MAC_ADDR`.

## 🎯 Testing Strategy

### Test Categories

The SDK implements **standardized test patterns** using direct `pkg/testutil` integration:

| Category          | Purpose                                   | Implementation Pattern       | Coverage Target |
| ----------------- | ----------------------------------------- | ---------------------------- | --------------- |
| **Service Tests** | Service construction and lifecycle        | Direct service instantiation | 100%            |
| **Get Tests**     | Mock-based GET operations with validation | `testutil.NewMockServer`     | Get/List: 100%  |
| **Set Tests**     | Mock-based SET/RPC operations             | `MockErrorServer` patterns   | Set/Admin: 90%+ |
| **Integration**   | Live WNC operations (GET only)            | Direct test implementation   | N/A             |
| **Scenario/E2E**  | Non-disruptive CRUD against live WNC      | Custom scenario suites       | N/A             |

### IOS-XE Version Support

| IOS-XE Version | Service Support                            | Test Strategy                                 |
| -------------- | ------------------------------------------ | --------------------------------------------- |
| **17.12.x**    | Core services (AP, WLAN, Client, RF, etc.) | Full mock + integration testing               |
| **17.18.x**    | Advanced services (WAT, URWB, Spaces)      | 404 error expectation tests + real data mocks |

> [!NOTE]
> WAT (Wireless Assurance Testing), URWB (Ultra-Reliable Wireless Backhaul), and Spaces services require IOS-XE 17.18.1+. Tests expect 404 responses when services are not configured and use real WNC data structure for mock responses.

### Coverage Requirements

- **Service package** (`service/`): **90% minimum**
- **Repository overall**: **80% minimum**
- **Critical paths** (Get/List methods): **100% required**

## 📂 Test Organization

### Directory Structure

```text
cisco-ios-xe-wireless-go/
├── service/
│   └── {service}/
│       ├── service_test.go         # Direct service tests using pkg/testutil
│       └── errors.go              # Service-specific error constants
├── tests/
│   ├── integration/
│   │   ├── {service}_service_test.go  # Live WNC integration tests per service
│   │   └── ...                       # Additional integration tests
│   └── scenario/                     # E2E scenario tests
└── pkg/
    └── testutil/
        ├── client.go               # Test client utilities
        ├── mock.go                 # Mock server implementations
        └── helper.go              # Test helper functions
```

### Naming Conventions

**Test Functions (New Unified Naming Convention):**

```go
// Unit tests (service/ directory)
TestXServiceUnit_Constructor_Success        // Service construction tests
TestXServiceUnit_GetOperations_MockSuccess  // GET operations with mock server
TestXServiceUnit_GetOperations_ValidationErrors // GET validation and edge cases
TestXServiceUnit_SetOperations_MockSuccess  // SET/RPC operations with mock server
TestXServiceUnit_SetOperations_ValidationErrors // SET validation and edge cases

// Integration tests (tests/integration/ directory)
TestXServiceIntegration_GetOperations_Success // Live WNC GET operations
TestXServiceIntegration_GetConfig_Success     // Live WNC configuration retrieval
```

**Examples:**

- `TestAPServiceUnit_Constructor_Success` - AP service construction using direct instantiation
- `TestWLANServiceUnit_GetOperations_MockSuccess` - WLAN GET operations with mock server
- `TestRFServiceIntegration_GetConfig_Success` - RF configuration retrieval with live controller

## 🧰 Prerequisites

### For Unit Tests (Layers 1-2)

Unit tests require no special configuration and can be run in any Go development environment.

| Requirement | Version | Notes                              |
| ----------- | ------- | ---------------------------------- |
| Go          | 1.25+   | Uses stdlib testing + pkg/testutil |
| make        | Latest  | Convenience targets                |

### For Integration/E2E Tests (Layers 3-4)

#### 1. Cisco Catalyst 9800 Wireless Network Controller

Integration and E2E tests require a real Cisco Catalyst 9800 WNC. Please refer to [References](#references).

#### 2. Environment Variables

| Variable           | Description         | Example                 |
| ------------------ | ------------------- | ----------------------- |
| `WNC_CONTROLLER`   | Controller host/IP  | `wnc1.example.internal` |
| `WNC_ACCESS_TOKEN` | Base64 `user:pass`  | `YWRtaW46cGFzc3dvcmQ=`  |
| `WNC_AP_MAC_ADDR`  | Test AP's Radio MAC | `00:11:22:33:dd:ee:f0`  |

<details><summary>Environment setup</summary>

```bash
export WNC_CONTROLLER="<controller-host-or-ip>"
export WNC_ACCESS_TOKEN="<base64-username:password>"
export WNC_AP_MAC_ADDR="<test-ap-radio-mac-address>"
```

</details>

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

<details><summary><strong>Example:</strong> Service Constructor Test</summary>

```go
// Service constructor test using direct instantiation
func TestWatServiceUnit_Constructor_Success(t *testing.T) {
    service := wat.NewService(nil)
    if service.Client() != nil {
        t.Error("Expected nil client service")
    }

    // Test with valid client
    mockServer := testutil.NewMockServer(map[string]string{
        "test": `{"data": {}}`,
    })
    defer mockServer.Close()

    client := testutil.NewTestClient(mockServer)
    service = wat.NewService(client.Core().(*core.Client))
    if service.Client() == nil {
        t.Error("Expected service to have client, got nil")
    }
}
```

```bash
go test ./service/ap -run "TestAPServiceUnit_Constructor"
```

</details>

#### Layer 2: Mock-based Method Tests

Tests all operations with **mock server** using direct `pkg/testutil` implementation.

<details><summary><strong>Example:</strong> Mock-based Method Tests</summary>

```go
// Real WNC data-based mock test (Core services)
func TestClientServiceUnit_GetOperations_MockSuccess(t *testing.T) {
    // Mock responses based on real IOS-XE 17.12.x WNC data
    responses := map[string]string{
        "Cisco-IOS-XE-wireless-client-oper:client-oper-data": `{
            "Cisco-IOS-XE-wireless-client-oper:client-oper-data": {
                "common-oper-data": [{
                    "client-mac": "02:40:f1:f7:f7:87",
                    "ap-name": "TEST-AP01",
                    "wlan-id": 1
                }]
            }
        }`,
    }
    mockServer := testutil.NewMockServer(responses)
    defer mockServer.Close()

    testClient := testutil.NewTestClient(mockServer)
    service := client.NewService(testClient.Core().(*core.Client))
    ctx := testutil.TestContext(t)

    result, err := service.GetOperational(ctx)
    if err != nil {
        t.Errorf("GetOperational failed: %v", err)
    }
}

// IOS-XE 17.18.1+ service with 404 expectation test
func TestUrwbServiceUnit_GetConfig_ErrorExpected(t *testing.T) {
    mockServer := testutil.NewMockErrorServer([]string{
        "Cisco-IOS-XE-wireless-urwbnet-cfg:urwbnet-cfg-data",
    }, 404)
    defer mockServer.Close()

    client := testutil.NewTestClient(mockServer)
    service := urwb.NewService(client.Core().(*core.Client))
    ctx := testutil.TestContext(t)

    _, err := service.GetConfig(ctx)
    if err == nil {
        t.Error("Expected 404 error for unconfigured URWB service")
    }
}
```

```bash
go test ./service/ap -run "TestAPServiceUnit.*Mock"
```

</details>

#### Layer 3: Integration Tests

Tests only GET operations with **live WNC**.

```bash
go test ./tests/integration -tags=integration
```

**Example:**

- [`tests/integration/client_service_test.go`](../tests/integration/client_service_test.go) - `TestClientServiceIntegration_GetOperations_Success`
- [`tests/integration/rrm_service_test.go`](../tests/integration/rrm_service_test.go) - `TestRrmServiceIntegration_GetConfig_Success`

#### Layer 4: E2E Scenario Tests

Non-disruptive CRUD operations with **live WNC**.

```bash
go test ./tests/scenario/ap/ -tags=scenario -v
go test ./tests/scenario/rf/ -tags=scenario -v
go test ./tests/scenario/site/ -tags=scenario -v
go test ./tests/scenario/wlan/ -tags=scenario -v
```

**Example:**

- [`tests/scenario/ap/service_test.go#TestAPServiceScenario_AdminStateManagement_Success`](../tests/scenario/ap/service_test.go)
- [`tests/scenario/site/tag_service_test.go#TestSiteTagServiceScenario_TagLifecycleManagement_Success`](../tests/scenario/site/tag_service_test.go)

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

## 📊 Test Data Management

### Test Fixtures

Tests use real WNC data for accurate mock responses.

| Location                      | Purpose                          |
| ----------------------------- | -------------------------------- |
| `testdata/*.json`             | Raw controller responses         |
| `pkg/testutil/`               | Mock server implementations      |
| `tests/integration/testdata/` | Integration test snapshots       |
| Service test files            | Inline real WNC data with source |

<details><summary>Example structure</summary>

```text
testdata/
├── ap_oper_response.json
├── client_oper_response.json
├── general_cfg_response.json
└── rrm_global_oper_response.json
```

</details>

## 📚 Appendix

### Testing Tips

1. **Start with unit tests** - Validate basic functionality first
2. **Use real WNC data** - Base mock responses on actual controller data
3. **Test 404 scenarios** - IOS-XE 17.18.1+ services may not be configured
4. **Follow TN-001 naming** - Use standardized test function names
5. **Direct pkg/testutil usage** - No framework abstractions needed
6. **Coverage-driven development** - Write tests to meet coverage targets

### Troubleshooting

| Issue                  | Solution                                                                  |
| ---------------------- | ------------------------------------------------------------------------- |
| Missing env vars       | Set `WNC_CONTROLLER`, `WNC_ACCESS_TOKEN` and `WNC_AP_MAC_ADDR`            |
| Unreachable controller | Verify DNS/IP connectivity                                                |
| TLS errors             | Check certificate validity; use `WithInsecureSkipVerify` for testing only |
| Auth failures          | Ensure token is Base64 of `user:pass`                                     |

### References

- 📖 [Cisco Catalyst 9800-CL Wireless Controller for Cloud Deployment Guide](https://www.cisco.com/c/en/us/td/docs/wireless/controller/9800/technical-reference/c9800-cl-dg.html)
- 📖 [Cisco Catalyst 9800 Series Wireless Controller Programmability Guide](https://www.cisco.com/c/en/us/td/docs/wireless/controller/9800/programmability-guide/b_c9800_programmability_cg/cisco-catalyst-9800-series-wireless-controller-programmability-guide.html)
- 📖 [YANG Models and Platform Capabilities for Cisco IOS XE 17.12.1](https://github.com/YangModels/yang/tree/main/vendor/cisco/xe/17121#readme)
