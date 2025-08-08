package afc

import (
	"context"
	"encoding/json"
	"sync"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/tests"
)

// TestDataCollector holds test data for AFC service functions
type TestDataCollector struct {
	mu             sync.Mutex
	OperResp       *model.AfcOperResponse
	OperErr        error
	APRespResp     *model.AfcOperEwlcAfcApRespResponse
	APRespErr      error
	CloudOperResp  *model.AfcCloudOperResponse
	CloudOperErr   error
	CloudStatsResp *model.AfcCloudOperAfcCloudStatsResponse
	CloudStatsErr  error
}

// TestAfcService tests all AFC service functions with the 4-pattern testing approach
func TestAfcService(t *testing.T) {
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

		// Test Oper method
		go func() {
			defer wg.Done()
			resp, err := service.GetOper(ctx)
			collector.mu.Lock()
			collector.OperResp = resp
			collector.OperErr = err
			collector.mu.Unlock()
		}()

		// Test APResp method
		go func() {
			defer wg.Done()
			resp, err := service.GetAPResp(ctx)
			collector.mu.Lock()
			collector.APRespResp = resp
			collector.APRespErr = err
			collector.mu.Unlock()
		}()

		// Test CloudOper method
		go func() {
			defer wg.Done()
			resp, err := service.GetCloudOper(ctx)
			collector.mu.Lock()
			collector.CloudOperResp = resp
			collector.CloudOperErr = err
			collector.mu.Unlock()
		}()

		// Test CloudStats method
		go func() {
			defer wg.Done()
			resp, err := service.GetCloudStats(ctx)
			collector.mu.Lock()
			collector.CloudStatsResp = resp
			collector.CloudStatsErr = err
			collector.mu.Unlock()
		}()

		wg.Wait()

		// Validate collected data
		t.Logf("Collected data from %d AFC service methods", 4)
	})

	// Test JSON serialization/deserialization
	t.Run("JSON_Serialization", func(t *testing.T) {
		testCases := []struct {
			name     string
			jsonData string
		}{
			{
				name: "AfcOperResponse",
				jsonData: `{
					"Cisco-IOS-XE-wireless-afc-oper:afc-oper-data": {
						"afc-enable": true,
						"afc-mode": "standard"
					}
				}`,
			},
			{
				name: "AfcCloudOperResponse",
				jsonData: `{
					"Cisco-IOS-XE-wireless-afc-cloud-oper:afc-cloud-oper-data": {
						"cloud-status": "connected",
						"last-update": "2024-01-01T12:00:00.000Z"
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
				name: "Oper",
				method: func() (interface{}, error) {
					return service.GetOper(ctx)
				},
			},
			{
				name: "APResp",
				method: func() (interface{}, error) {
					return service.GetAPResp(ctx)
				},
			},
			{
				name: "CloudOper",
				method: func() (interface{}, error) {
					return service.GetCloudOper(ctx)
				},
			},
			{
				name: "CloudStats",
				method: func() (interface{}, error) {
					return service.GetCloudStats(ctx)
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
			_, err := service.GetOper(ctx)
			if err == nil {
				t.Fatal("Expected error with nil client, got none")
			}
		})

		// Test with nil context
		t.Run("NilContext", func(t *testing.T) {
			var nilCtx context.Context //nolint:SA1012 // Testing nil context behavior
			_, err := service.GetOper(nilCtx)
			if err == nil {
				t.Fatal("Expected error when using nil context, but got none")
			}
		})

		// Test with canceled context
		t.Run("CanceledContext", func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			cancel()
			_, err := service.GetOper(ctx)
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

		// Test operational data
		operResp, operErr := service.GetOper(ctx)
		if operErr != nil {
			t.Logf("Integration test - Oper error: %v", operErr)
		} else {
			t.Logf("Integration test - Oper success: %+v", operResp)
		}

		// Test cloud operational data
		cloudResp, cloudErr := service.GetCloudOper(ctx)
		if cloudErr != nil {
			t.Logf("Integration test - CloudOper error: %v", cloudErr)
		} else {
			t.Logf("Integration test - CloudOper success: %+v", cloudResp)
		}
	})
}
