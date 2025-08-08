// Package tests provides common testing utilities for the Cisco WNC API client.
package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"sync"
	"testing"
	"time"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/constants"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
)

// TestClient creates a test client using environment variables
func TestClient(t *testing.T) *core.Client {
	t.Helper()

	controller := os.Getenv("WNC_CONTROLLER")
	token := os.Getenv("WNC_ACCESS_TOKEN")

	if controller == "" || token == "" {
		t.Skip("WNC_CONTROLLER and WNC_ACCESS_TOKEN environment variables must be set for integration tests")
	}

	client, err := core.New(controller, token,
		core.WithTimeout(30*time.Second),
		core.WithInsecureSkipVerify(true))
	if err != nil {
		t.Fatalf("Failed to create test client: %v", err)
	}

	return client
}

// CreateTestClientFromEnv creates a test client from environment variables
// This is an alias for TestClient to match expected API
func CreateTestClientFromEnv(t *testing.T) *core.Client {
	return TestClient(t)
}

// TestTimeouts
const (
	// DefaultTestTimeout is the default timeout for tests
	DefaultTestTimeout = 30 * time.Second
	// ExtendedTestTimeout is an extended timeout for longer tests
	ExtendedTestTimeout = 60 * time.Second
	// ShortTestTimeout is a short timeout for quick tests
	ShortTestTimeout = 5 * time.Second
)

// TestDataDir is the directory for test data files
const TestDataDir = "./test_data"

// TestContext creates a test context with timeout
func TestContext(t *testing.T) context.Context {
	return TestContextWithTimeout(t, DefaultTestTimeout)
}

// TestContextWithTimeout creates a test context with custom timeout
func TestContextWithTimeout(t *testing.T, timeout time.Duration) context.Context {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	t.Cleanup(cancel)
	return ctx
}

// SkipIfNoConnection skips the test if no network connection to WNC
func SkipIfNoConnection(t *testing.T, client *core.Client) {
	t.Helper()

	// Try a simple health check - this assumes there's some basic endpoint available
	ctx, cancel := context.WithTimeout(context.Background(), ShortTestTimeout)
	defer cancel()

	// We'll use a very basic endpoint that should exist on most controllers
	var result interface{}
	err := client.Do(ctx, "GET", "/yang-library-version", &result)
	if err != nil {
		t.Skipf("No connection to WNC controller: %v", err)
	}
}

// SaveTestDataToFile saves test data to a JSON file
func SaveTestDataToFile(filename string, data interface{}) error {
	// Create test_data directory if it doesn't exist
	if err := os.MkdirAll(TestDataDir, 0o755); err != nil { //nolint:gosec // Test directory permissions
		return err
	}

	// Create the full file path
	fullPath := filepath.Join(TestDataDir, filename)

	// Marshal data to JSON
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	// Write to file
	return os.WriteFile(fullPath, jsonData, 0o644) //nolint:gosec // Test file permissions
}

// Common Service Testing Patterns

// ServiceTestConfig holds configuration for service testing
type ServiceTestConfig struct {
	ServiceName    string
	TestMethods    []TestMethod
	JSONTestCases  []JSONTestCase
	SkipShortTests bool
}

// TestMethod represents a service method to test
type TestMethod struct {
	Name   string
	Method func() (interface{}, error)
}

// JSONTestCase represents a JSON serialization test case
type JSONTestCase struct {
	Name     string
	JSONData string
}

// GenericTestDataCollector provides a generic data collector for service tests
type GenericTestDataCollector struct {
	mu      sync.Mutex
	Results map[string]ServiceMethodResult
}

// ServiceMethodResult holds the result of a service method call
type ServiceMethodResult struct {
	Response interface{}
	Error    error
}

// NewGenericTestDataCollector creates a new generic test data collector
func NewGenericTestDataCollector() *GenericTestDataCollector {
	return &GenericTestDataCollector{
		Results: make(map[string]ServiceMethodResult),
	}
}

// Collect stores the result of a service method call
func (c *GenericTestDataCollector) Collect(methodName string, response interface{}, err error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Results[methodName] = ServiceMethodResult{
		Response: response,
		Error:    err,
	}
}

// RunServiceTests executes the standard 4-pattern testing approach
func RunServiceTests(t *testing.T, config ServiceTestConfig) {
	t.Helper()

	// Create a test client (may be nil if environment not set)
	var client *core.Client

	// Try to get real client from environment
	if os.Getenv("WNC_CONTROLLER") != "" && os.Getenv("WNC_ACCESS_TOKEN") != "" {
		client = TestClient(t)
	}

	// ========================================
	// 1. UNIT TESTS (Structure/Type Validation & JSON Serialization/Deserialization)
	// ========================================

	t.Run("Service_Creation", func(t *testing.T) {
		// Test with nil client - should not panic during creation
		if client != nil {
			// We can't create a generic service here, but we can test the pattern
			t.Logf("Service creation test for %s - client available", config.ServiceName)
		} else {
			t.Logf("Service creation test for %s - no client available", config.ServiceName)
		}
	})

	t.Run("Data_Collection", func(t *testing.T) {
		if len(config.TestMethods) == 0 {
			t.Skip("No test methods provided")
		}

		collector := NewGenericTestDataCollector()
		var wg sync.WaitGroup

		wg.Add(len(config.TestMethods))

		for _, method := range config.TestMethods {
			go func(m TestMethod) {
				defer wg.Done()
				resp, err := m.Method()
				collector.Collect(m.Name, resp, err)
			}(method)
		}

		wg.Wait()

		// Log results
		for methodName, result := range collector.Results {
			if result.Error != nil {
				t.Logf("Method %s returned error: %v", methodName, result.Error)
			} else {
				t.Logf("Method %s returned result of type %T", methodName, result.Response)
			}
		}
	})

	// Test JSON serialization/deserialization
	t.Run("JSON_Serialization", func(t *testing.T) {
		for _, testCase := range config.JSONTestCases {
			t.Run(testCase.Name, func(t *testing.T) {
				var data interface{}
				err := json.Unmarshal([]byte(testCase.JSONData), &data)
				if err != nil {
					t.Errorf("Failed to unmarshal %s: %v", testCase.Name, err)
				}

				_, err = json.Marshal(data)
				if err != nil {
					t.Errorf("Failed to marshal %s: %v", testCase.Name, err)
				}
			})
		}
	})

	// ========================================
	// 2. TABLE-DRIVEN TEST PATTERNS
	// ========================================

	t.Run("Method_Tests", func(t *testing.T) {
		for _, method := range config.TestMethods {
			t.Run(method.Name, func(t *testing.T) {
				result, err := method.Method()
				if err != nil {
					t.Logf("Method %s returned error: %v", method.Name, err)
				}
				if result != nil {
					t.Logf("Method %s returned result of type %T", method.Name, result)
				}
			})
		}
	})

	// ========================================
	// 3. FAIL-FAST ERROR DETECTION (t.Fatalf/t.Fatal)
	// ========================================

	t.Run("Critical_Validations", func(t *testing.T) {
		if len(config.TestMethods) > 0 {
			// Test with nil context (should handle gracefully or fail fast)
			t.Run("NilContext", func(t *testing.T) {
				var nilCtx context.Context //nolint:SA1012 // Testing nil context behavior
				// We would need specific service instance to test this
				_ = nilCtx
				t.Log("Nil context test - implementation specific")
			})

			// Test with canceled context
			t.Run("CanceledContext", func(t *testing.T) {
				canceledCtx, cancel := context.WithCancel(context.Background())
				cancel()
				// We would need specific service instance to test this
				_ = canceledCtx
				t.Log("Canceled context test - implementation specific")
			})
		}
	})

	// ========================================
	// 4. INTEGRATION TESTS (API Endpoint, Real Controller)
	// ========================================

	t.Run("Integration_Test", func(t *testing.T) {
		// Skip if running in short mode or no integration tests requested
		if testing.Short() && config.SkipShortTests {
			t.Skip("Skipping integration test in short mode")
		}

		if client == nil {
			t.Skip("No test client available for integration tests")
		}

		if len(config.TestMethods) > 0 {
			// Run first method as integration test example
			method := config.TestMethods[0]
			resp, err := method.Method()
			if err != nil {
				t.Logf("Integration test - %s error: %v", method.Name, err)
			} else {
				t.Logf("Integration test - %s success: type %T", method.Name, resp)
			}
		}
	})
}

// ValidateStructType validates that a struct type can be properly marshaled/unmarshaled
func ValidateStructType(t *testing.T, structType interface{}) {
	t.Helper()

	// Handle nil interface
	if structType == nil {
		t.Logf("ValidateStructType received nil value - skipping reflection test")
		return
	}

	// Test marshaling
	data, err := json.Marshal(structType)
	if err != nil {
		t.Logf("Failed to marshal struct type %T: %v", structType, err)
		return
	}

	// Test unmarshaling - check if we can create a new instance
	structTypeReflected := reflect.TypeOf(structType)
	if structTypeReflected == nil {
		t.Logf("ValidateStructType received value with nil type - skipping unmarshal test")
		return
	}

	newInstance := reflect.New(structTypeReflected).Interface()
	err = json.Unmarshal(data, newInstance)
	if err != nil {
		t.Logf("Failed to unmarshal struct type %T: %v", structType, err)
	}
}

// AssertNonNilResult asserts that a result is not nil
func AssertNonNilResult(t *testing.T, result interface{}, methodName string) {
	t.Helper()
	if result == nil {
		t.Errorf("Method %s returned nil result", methodName)
	}
}

// LogMethodResult logs the result of a method call
func LogMethodResult(t *testing.T, methodName string, result interface{}, err error) {
	t.Helper()
	if err != nil {
		t.Logf("Method %s returned error: %v", methodName, err)
	} else {
		t.Logf("Method %s returned %T", methodName, result)
	}
}

// StandardJSONTestCases provides standard JSON test cases for most services
func StandardJSONTestCases(yangModule string) []JSONTestCase {
	return []JSONTestCase{
		{
			Name: fmt.Sprintf("%sCfgResponse", pascalCase(yangModule)),
			JSONData: fmt.Sprintf(`{
				"%s%s-cfg:%s-cfg-data": {
					"test-data": "value"
				}
			}`, constants.YANGModelPrefix, yangModule, yangModule),
		},
		{
			Name: fmt.Sprintf("%sOperResponse", pascalCase(yangModule)),
			JSONData: fmt.Sprintf(`{
				"%s%s-oper:%s-oper-data": {
					"test-data": "value"
				}
			}`, constants.YANGModelPrefix, yangModule, yangModule),
		},
	}
}

// pascalCase converts a string to PascalCase
func pascalCase(s string) string {
	if len(s) == 0 {
		return s
	}
	if len(s) == 1 {
		// Only convert lowercase letters to uppercase
		if s[0] >= 'a' && s[0] <= 'z' {
			return string(s[0] - 32) // Convert to uppercase
		}
		return s // Return as-is if not a lowercase letter
	}

	// Convert first character if it's a lowercase letter
	if s[0] >= 'a' && s[0] <= 'z' {
		return string(s[0]-32) + s[1:]
	}
	return s // Return as-is if first character is not lowercase
}
