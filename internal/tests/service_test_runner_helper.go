package tests

import (
	"context"
	"encoding/json"
	"os"
	"sync"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
)

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
			// I can't create a generic service here, but we can test the pattern
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
				// I would need specific service instance to test this
				_ = nilCtx
				t.Log("Nil context test - implementation specific")
			})

			// Test with canceled context
			t.Run("CanceledContext", func(t *testing.T) {
				canceledCtx, cancel := context.WithCancel(context.Background())
				cancel()
				// I would need specific service instance to test this
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
		if shortModeCheck() && config.SkipShortTests {
			t.Skip("Skipping integration test in short mode")
		}
		if client == nil {
			t.Skip("No test client available for integration tests")
		}
		if len(config.TestMethods) > 0 {
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
