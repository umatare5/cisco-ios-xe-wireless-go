package apf

import (
	"context"
	"encoding/json"
	"sync"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/tests"
)

// TestDataCollector holds test data for APF service functions
type TestDataCollector struct {
	mu      sync.Mutex
	CfgResp *model.ApfCfgResponse
	CfgErr  error
}

// TestAPFService tests all APF service functions with the 4-pattern testing approach
func TestAPFService(t *testing.T) {
	client := tests.OptionalTestClient(t)
	ctx := context.Background()
	if client != nil {
		ctx = tests.TestContext(t)
	}

	service := NewService(client)

	// ========================================
	// 1. UNIT TESTS (Structure/Type Validation & JSON Serialization/Deserialization)
	// ========================================

	t.Run("Service_Creation", func(t *testing.T) {
		// Test service creation with nil client
		nilService := NewService(nil)
		// With value receiver, we check if the internal client is properly set
		if nilService.c != nil {
			t.Error("NewService should set internal client to nil when nil client passed")
		}

		if client != nil {
			validService := NewService(client)
			if validService.c != client {
				t.Error("NewService should set internal client correctly with valid client")
			}
		}
	})

	t.Run("Data_Collection", func(t *testing.T) {
		// Test the functionality by collecting data using available methods
		collector := &TestDataCollector{}

		var wg sync.WaitGroup
		wg.Add(1) // Only one method to test

		// Test GetCfg method
		go func() {
			defer wg.Done()
			resp, err := service.GetCfg(ctx)
			collector.mu.Lock()
			collector.CfgResp = resp
			collector.CfgErr = err
			collector.mu.Unlock()
		}()

		wg.Wait()

		// Validate collected data
		if collector.CfgErr != nil {
			t.Logf("Cfg method returned error: %v", collector.CfgErr)
		}
		if collector.CfgResp != nil {
			t.Logf("Cfg method returned data successfully")
		}
	})

	// Test JSON serialization/deserialization
	t.Run("JSON_Serialization", func(t *testing.T) {
		testCases := []struct {
			name     string
			jsonData string
		}{
			{
				name: "ApfCfgResponse",
				jsonData: `{
					"Cisco-IOS-XE-wireless-apf-cfg:apf-cfg-data": {
						"apf-config": {
							"enable": true
						}
					}
				}`,
			},
		}

		for _, tt := range testCases {
			t.Run(tt.name, func(t *testing.T) {
				var data interface{}
				err := json.Unmarshal([]byte(tt.jsonData), &data)
				if err != nil {
					t.Errorf("Failed to unmarshal %s: %v", tt.name, err)
				}

				_, err = json.Marshal(data)
				if err != nil {
					t.Errorf("Failed to marshal %s: %v", tt.name, err)
				}
			})
		}
	})

	// ========================================
	// 2. TABLE-DRIVEN TEST PATTERNS
	// ========================================

	t.Run("Method_Tests", func(t *testing.T) {
		methodTests := []struct {
			name   string
			method func() (interface{}, error)
		}{
			{
				name: "GetCfg",
				method: func() (interface{}, error) {
					return service.GetCfg(ctx)
				},
			},
		}

		for _, tt := range methodTests {
			t.Run(tt.name, func(t *testing.T) {
				result, err := tt.method()
				if err != nil {
					t.Logf("Method %s returned error: %v", tt.name, err)
				}
				if result != nil {
					t.Logf("Method %s returned result", tt.name)
				}
			})
		}
	})

	// ========================================
	// 3. FAIL-FAST ERROR DETECTION (t.Fatalf/t.Fatal)
	// ========================================

	t.Run("Critical_Validations", func(t *testing.T) {
		// Test with nil client
		t.Run("NilClient", func(t *testing.T) {
			service := NewService(nil)
			_, err := service.GetCfg(ctx)
			if err == nil {
				t.Fatal("Expected error with nil client, got none")
			}
		})

		// Test with nil context
		t.Run("NilContext", func(t *testing.T) {
			var nilCtx context.Context //nolint:SA1012 // Testing nil context behavior
			_, err := service.GetCfg(nilCtx)
			if err == nil {
				t.Fatal("Expected error when using nil context, but got none")
			}
		})

		// Test with canceled context
		t.Run("CanceledContext", func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			cancel()
			_, err := service.GetCfg(ctx)
			if err == nil {
				t.Fatal("Expected error with canceled context, got none")
			}
		})
	})

	// ========================================
	// 4. INTEGRATION TESTS (API Endpoint, Real Controller)
	// ========================================

	t.Run("Integration_Test", func(t *testing.T) {
		// This test requires actual WNC environment
		// Skip if running in unit test mode
		if testing.Short() {
			t.Skip("Skipping integration test in short mode")
		}

		// Test with real service
		resp, err := service.GetCfg(ctx)
		if err != nil {
			t.Logf("Integration test - Cfg error: %v", err)
		} else {
			t.Logf("Integration test - Cfg success: %+v", resp)
		}
	})
}
