# 🧪 Testing Guide

Comprehensive testing framework for the Cisco IOS-XE Wireless Go SDK, featuring automated test execution, coverage analysis, and integration with live controllers.

> [!TIP]
> This guide covers unit testing, integration testing, and code coverage analysis for both developers and contributors.

## 🎯 Test Architecture

### Test Types

- **Unit Tests**: Validate core functionality, serialization, and error handling
- **Integration Tests**: Interact with real Cisco Catalyst 9800 controllers
- **Coverage Tests**: Analyze code coverage with detailed reporting
- **Service Tests**: Validate all 25+ service implementations with unified patterns

### Service Testing Pattern

All services follow a consistent 4-stage testing approach:

1. **Service Creation**: Constructor and configuration validation
2. **Data Collection**: Method execution and response handling
3. **JSON Serialization**: Type validation and data marshaling
4. **Integration Testing**: Live controller interaction

## 📊 Coverage Analysis

> [!NOTE]
> The project maintains high code coverage standards with automated reporting.

| Coverage Type  | Target | Current |
|---------------|--------|---------|
| Core Client   | ≥98%   | 98.9%   |
| Services      | ≥95%   | 99.2%   |
| Total Project | ≥92%   | 96.8%   |

<details>
<summary>View coverage commands</summary>

```bash
# Generate coverage report
make test-coverage

# Create HTML coverage report
make test-coverage-html

# Analyze specific functions
go tool cover -func=./tmp/coverage.out | grep -v 100.0%
```

</details>

## 🚀 Prerequisites

### Development Environment

| Requirement | Version | Description |
|------------|---------|-------------|
| Go         | 1.24+   | Latest Go toolchain |
| Make       | Any     | Build automation |

### Integration Testing

Integration tests require a live Cisco Catalyst 9800 controller:

| Variable           | Description                | Example               |
|-------------------|---------------------------|-----------------------|
| `WNC_CONTROLLER`  | Controller IP or hostname | `192.168.1.100`      |
| `WNC_ACCESS_TOKEN`| Base64 encoded credentials| `YWRtaW46cGFzc3dvcmQ=`|

#### Environment Setup

```bash
export WNC_CONTROLLER="your-controller-ip"
export WNC_ACCESS_TOKEN="$(echo -n 'username:password' | base64)"
```

## 🛠️ Running Tests

### Quick Commands

| Command | Description | Requirements |
|---------|-------------|--------------|
| `make test-unit` | Unit tests only | None |
| `make test-integration` | Integration tests | WNC access |
| `make test-coverage` | Coverage analysis | None |
| `make build` | Compilation check | None |

### Detailed Test Execution

#### Unit Tests

```bash
# Basic unit tests
make test-unit

# Verbose output
./scripts/test_unit.sh --verbose

# Short mode (skip slow tests)
./scripts/test_unit.sh --short
```

#### Integration Tests

```bash
# Full integration test suite
make test-integration

# Check environment only
./scripts/test_integration.sh --check-env-only

# Verbose integration testing
./scripts/test_integration.sh --verbose
```

#### Coverage Analysis

```bash
# Generate coverage data
make test-coverage

# Create HTML report and open in browser
make test-coverage-html
```

## 📁 Test Data Management

### Automatic Data Collection

Integration tests automatically save API responses:

- **Location**: `./test_data/` in each service directory
- **Format**: JSON files with descriptive names
- **Purpose**: Validation and offline debugging

#### Example Test Data Structure

```text
afc/test_data/
├── afc_oper_data.json              # Live API response
├── afc_cloud_oper_data.json        # Cloud operations data
└── afc_test_data_collected.json    # Test run collection
```

### Data Validation

Test data files serve multiple purposes:

- Response structure validation
- Offline development and debugging
- Regression testing
- API change detection

## 📋 Testing Best Practices

### Unit Test Standards

```go
// ✅ Correct pattern
func TestServiceMethod(t *testing.T) {
    client := createTestClient(t)
    service := NewService(client)

    ctx := context.Background()
    result, err := service.Method(ctx)

    if err != nil {
        t.Fatalf("Expected no error, got %v", err)
    }

    // Validate result structure and content
}
```

### Error Handling Validation

```go
// ✅ SA1012 compliant nil context testing
func TestMethodWithNilContext(t *testing.T) {
    client := createTestClient(t)
    service := NewService(client)

    var nilCtx context.Context
    _, err := service.Method(nilCtx)

    if err == nil {
        t.Fatal("Expected error for nil context")
    }
}
```

### Integration Test Patterns

```go
func TestServiceIntegration(t *testing.T) {
    client := tests.TestClient(t) // Handles env var validation
    ctx := tests.TestContext(t)   // Creates timeout context

    service := NewService(client)
    result, err := service.Method(ctx)

    if err != nil {
        t.Logf("Integration test error: %v", err)
        return
    }

    // Save test data for analysis
    tests.SaveTestDataToFile("service_response.json", result)
    t.Logf("Integration test success: %+v", result)
}
```

## 🔧 Development Workflow

### Recommended Testing Sequence

1. **Setup**: `make deps` - Install development tools
2. **Unit Testing**: `make test-unit` - Validate core functionality
3. **Environment**: Configure WNC environment variables
4. **Integration**: `make test-integration` - Test live functionality
5. **Coverage**: `make test-coverage-html` - Analyze coverage
6. **Build**: `make build` - Verify compilation

### Debugging Tips

#### Coverage Investigation

When coverage targets aren't met:

```bash
# Find uncovered functions
go tool cover -func=./tmp/coverage.out | grep -v 100.0%

# Analyze specific packages
go test -coverprofile=./tmp/package.out ./package/...
go tool cover -html=./tmp/package.out
```

#### Integration Test Debugging

```bash
# Test specific service
go test -v ./service/... -run TestServiceIntegration

# Check environment setup
./scripts/test_integration.sh --check-env-only

# Verbose output for troubleshooting
go test -v ./... -tags=integration
```

## 🎯 Quality Standards

### Code Quality Requirements

- **SA1012 Compliance**: Use `var nilCtx context.Context` instead of `nil`
- **Error Handling**: Standardized error patterns across all functions
- **Type Safety**: Strict typing for all API responses
- **Context Usage**: Proper context handling in all operations

### Performance Benchmarks

| Test Type | Target Duration | Actual |
|-----------|----------------|--------|
| Unit Tests | < 30s | ~25s |
| Integration | < 2m | ~90s |
| Coverage | < 60s | ~45s |

## 📚 References

### Testing Tools

- [Go Testing Package](https://pkg.go.dev/testing) - Standard library testing
- [Testify](https://github.com/stretchr/testify) - Testing toolkit
- [Go Tool Cover](https://pkg.go.dev/cmd/cover) - Coverage analysis

### Cisco Documentation

- [Catalyst 9800 Programmability Guide](https://www.cisco.com/c/en/us/td/docs/wireless/controller/9800/programmability-guide/b_c9800_programmability_cg.html)
- [YANG Models for IOS-XE 17.12](https://github.com/YangModels/yang/tree/main/vendor/cisco/xe/17121)
- [RESTCONF API Reference](https://www.cisco.com/c/en/us/td/docs/wireless/controller/9800/technical-reference/c9800-cl-dg.html)

## 🚀 Running Tests

The project includes convenient Makefile targets for testing:

| Command                 | Description                                        |
| ----------------------- | -------------------------------------------------- |
| `make test-unit`        | Run unit tests. WNC access is not required.        |
| `make test-integration` | Run integration tests. **WNC access is required.** |

<details><summary>Example of command result</summary>

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

## 📈 Coverage Analysis

The project supports comprehensive test coverage analysis:

### Coverage Reports

| Output Type     | Command                   | Description                                  |
| --------------- | ------------------------- | -------------------------------------------- |
| Terminal Output | `make test-coverage`      | Run tests with coverage analysis.            |
| HTML Report     | `make test-coverage-html` | Run tests and generate HTML coverage report. |

<details><summary>Example of Coverage Output</summary>

```text
Coverage report generated at ./tmp/coverage.out
total: (statements) 6.1%

📦 github.com/umatare5/cisco-ios-xe-wireless-go (42.9% coverage)
📦 github.com/umatare5/cisco-ios-xe-wireless-go/awips (75% coverage)
📦 github.com/umatare5/cisco-ios-xe-wireless-go/ap (1.1% coverage)
```

</details>

## 📚️ Appendix

### Code Quality & Error Handling

The project enforces strict code quality standards:

#### SA1012 Compliance Testing

All nil context tests must follow static analysis requirements:

```go
func TestFunctionWithNilContext(t *testing.T) {
    client, _ := wnc.New("192.168.1.100", "token")

    // ✅ Correct - SA1012 compliant
    var nilCtx context.Context
    _, err := SomeFunction(nilCtx, client)

    // ❌ Incorrect - SA1012 violation
    // _, err := SomeFunction(nil, client)

    if err == nil {
        t.Fatal("Expected error for nil context")
    }
}
```

#### Service Architecture Testing

All services must be tested with consistent patterns:

```go
func TestServiceConstructor(t *testing.T) {
    client, _ := wnc.New("192.168.1.100", "token")
    service := domain.NewService(client)

    if service == nil {
        t.Fatal("Expected non-nil service")
    }
}

func TestServiceMethod(t *testing.T) {
    // Test typed method signature: Method(ctx) (*model.Type, error)
    ctx := context.Background()
    result, err := service.Method(ctx)

    // Validate return types and error handling
}
```

#### Error Handling Standards

All client validation follows standardized error patterns:

```go
// ✅ Correct - standardized error handling
if client == nil {
    return nil, fmt.Errorf("%w: client cannot be nil", wnc.ErrInvalidConfiguration)
}

// ❌ Incorrect - basic error handling (deprecated)
if client == nil {
    return nil, errors.New("client is nil")
}
```

#### Test Error Expectations

When testing error conditions, expect the standardized error message:

```go
func TestFunction(t *testing.T) {
    _, err := SomeFunction(nil)
    if err == nil {
        t.Fatal("Expected error for nil client")
    }

    expectedError := "invalid client configuration: client cannot be nil"
    if err.Error() != expectedError {
        t.Errorf("Expected %q, got %q", expectedError, err.Error())
    }
}
```

### Testing Tips

For efficient testing workflow, start with unit tests and gradually move to integration tests.

1. **Install Dependencies**: `make deps` - Install gotestsum and other development tools.
2. **Unit Tests First**: `make test-unit` - Ensure basic functionality with enhanced output.
3. **Environment Setup**: Configure environment variables for integration tests.
4. **Environment Verification**: Check controller access to verify connectivity and credentials.
5. **Coverage Analysis**: `make test-coverage` - Run tests with coverage analysis.
6. **HTML Coverage Report**: `make test-coverage-html` - Generate detailed HTML coverage report.
7. **Test Data Review**: Examine generated JSON files to understand API response structures for debugging.
8. **Incremental Testing**: Test individual modules to target specific functionality when debugging.
9. **Run Integration Tests**: `make test-integration` - Ensure all functionality works as expected.

> [!TIP]
> For comprehensive testing, run both `make test-unit` and `make test-integration` sequentially to validate all functionality.

### 🔧 Development Dependencies

The project uses several tools to enhance the testing experience:

- **gotestsum**: Provides enhanced, human-readable test output with various formats
- **golangci-lint**: Code linting and static analysis
- **goreleaser**: Release automation

> [!Note]
> Install all dependencies with: `make deps`

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
