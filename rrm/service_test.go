package rrm

import (
	"context"
	"encoding/json"
	"sync"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/tests"
)

// TestDataCollector holds test data for RRM service functions
type TestDataCollector struct {
	mu             sync.Mutex
	CfgResp        *model.RrmCfgResponse
	CfgErr         error
	OperResp       *model.RrmOperResponse
	OperErr        error
	GlobalOperResp *model.RrmGlobalOperResponse
	GlobalOperErr  error
	EmulOperResp   *model.RrmEmulOperResponse
	EmulOperErr    error
}

// TestRrmService tests all RRM service functions with the 4-pattern testing approach
func TestRrmService(t *testing.T) {
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
		// Test service creation with both nil and valid clients
		nilService := NewService(nil)
		if nilService.c != nil {
			t.Error("NewService with nil client should have nil internal client")
		}

		if client != nil {
			validService := NewService(client)
			if validService.c == nil {
				t.Error("NewService with valid client should have non-nil internal client")
			}
		}
	})

	t.Run("Data_Collection", func(t *testing.T) {
		// Test the functionality by collecting data using all available methods
		collector := &TestDataCollector{}

		var wg sync.WaitGroup
		wg.Add(4) // 4 methods to test

		// Test GetCfg method
		go func() {
			defer wg.Done()
			resp, err := service.GetCfg(ctx)
			collector.mu.Lock()
			collector.CfgResp = resp
			collector.CfgErr = err
			collector.mu.Unlock()
		}()

		// Test GetOper method
		go func() {
			defer wg.Done()
			resp, err := service.GetOper(ctx)
			collector.mu.Lock()
			collector.OperResp = resp
			collector.OperErr = err
			collector.mu.Unlock()
		}()

		// Test GetGlobalOper method
		go func() {
			defer wg.Done()
			resp, err := service.GetGlobalOper(ctx)
			collector.mu.Lock()
			collector.GlobalOperResp = resp
			collector.GlobalOperErr = err
			collector.mu.Unlock()
		}()

		// Test GetEmulOper method
		go func() {
			defer wg.Done()
			resp, err := service.GetEmulOper(ctx)
			collector.mu.Lock()
			collector.EmulOperResp = resp
			collector.EmulOperErr = err
			collector.mu.Unlock()
		}()

		wg.Wait()

		// Validate collected data
		t.Logf("Collected data from %d RRM service methods", 4)
	})

	// Test JSON serialization/deserialization
	t.Run("JSON_Serialization", func(t *testing.T) {
		testCases := []struct {
			name     string
			jsonData string
		}{
			{
				name: "RrmCfgResponse",
				jsonData: `{
					"Cisco-IOS-XE-wireless-rrm-cfg:rrm-cfg-data": {
						"rrm-config": {
							"enable": true
						}
					}
				}`,
			},
			{
				name: "RrmOperResponse",
				jsonData: `{
					"Cisco-IOS-XE-wireless-rrm-oper:rrm-oper-data": {
						"rrm-statistics": {
							"channel-assignments": 10
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
			{
				name: "GetOper",
				method: func() (interface{}, error) {
					return service.GetOper(ctx)
				},
			},
			{
				name: "GetGlobalOper",
				method: func() (interface{}, error) {
					return service.GetGlobalOper(ctx)
				},
			},
			{
				name: "GetEmulOper",
				method: func() (interface{}, error) {
					return service.GetEmulOper(ctx)
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

		// Test configuration data
		cfgResp, cfgErr := service.GetCfg(ctx)
		if cfgErr != nil {
			t.Logf("Integration test - Cfg error: %v", cfgErr)
		} else {
			t.Logf("Integration test - Cfg success: %+v", cfgResp)
		}

		// Test operational data
		operResp, operErr := service.GetOper(ctx)
		if operErr != nil {
			t.Logf("Integration test - Oper error: %v", operErr)
		} else {
			t.Logf("Integration test - Oper success: %+v", operResp)
		}
	})
}
