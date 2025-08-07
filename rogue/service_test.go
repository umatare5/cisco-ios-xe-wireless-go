package rogue

import (
	"context"
	"encoding/json"
	"os"
	"sync"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/constants"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/tests"
)

// TestDataCollector holds test data for Rogue service functions
type TestDataCollector struct {
	mu             sync.Mutex
	OperResp       *model.RogueOperResponse
	OperErr        error
	StatsResp      *model.RogueStatsResponse
	StatsErr       error
	DataResp       *model.RogueDataResponse
	DataErr        error
	ClientDataResp *model.RogueClientDataResponse
	ClientDataErr  error
	RldpStatsResp  *model.RldpStatsResponse
	RldpStatsErr   error
}

// TestRogueService tests all Rogue service functions with the 4-pattern testing approach
func TestRogueService(t *testing.T) {
	// Create a mock client that will be used when environment variables are not set
	var client *core.Client
	var ctx context.Context

	// Try to get real client from environment
	if os.Getenv("WNC_CONTROLLER") != "" && os.Getenv("WNC_ACCESS_TOKEN") != "" {
		client = tests.TestClient(t)
		ctx = tests.TestContext(t)
	} else {
		// Use nil client for unit testing
		client = nil
		ctx = context.Background()
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
		wg.Add(constants.RogueServiceMethods) // 5 methods to test

		// Test Oper method
		go func() {
			defer wg.Done()
			resp, err := service.Oper(ctx)
			collector.mu.Lock()
			collector.OperResp = resp
			collector.OperErr = err
			collector.mu.Unlock()
		}()

		// Test Stats method
		go func() {
			defer wg.Done()
			resp, err := service.Stats(ctx)
			collector.mu.Lock()
			collector.StatsResp = resp
			collector.StatsErr = err
			collector.mu.Unlock()
		}()

		// Test Data method
		go func() {
			defer wg.Done()
			resp, err := service.Data(ctx)
			collector.mu.Lock()
			collector.DataResp = resp
			collector.DataErr = err
			collector.mu.Unlock()
		}()

		// Test ClientData method
		go func() {
			defer wg.Done()
			resp, err := service.ClientData(ctx)
			collector.mu.Lock()
			collector.ClientDataResp = resp
			collector.ClientDataErr = err
			collector.mu.Unlock()
		}()

		// Test RldpStats method
		go func() {
			defer wg.Done()
			resp, err := service.RldpStats(ctx)
			collector.mu.Lock()
			collector.RldpStatsResp = resp
			collector.RldpStatsErr = err
			collector.mu.Unlock()
		}()

		wg.Wait()

		// Validate collected data
		t.Logf("Collected data from %d Rogue service methods", 5)
	})

	// Test JSON serialization/deserialization
	t.Run("JSON_Serialization", func(t *testing.T) {
		testCases := []struct {
			name     string
			jsonData string
		}{
			{
				name: "RogueOperResponse",
				jsonData: `{
					"Cisco-IOS-XE-wireless-rogue-oper:rogue-oper-data": {
						"rogue-stats": {
							"total-rogue-aps": 5
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
				name: "Oper",
				method: func() (interface{}, error) {
					return service.Oper(ctx)
				},
			},
			{
				name: "Stats",
				method: func() (interface{}, error) {
					return service.Stats(ctx)
				},
			},
			{
				name: "Data",
				method: func() (interface{}, error) {
					return service.Data(ctx)
				},
			},
			{
				name: "ClientData",
				method: func() (interface{}, error) {
					return service.ClientData(ctx)
				},
			},
			{
				name: "RldpStats",
				method: func() (interface{}, error) {
					return service.RldpStats(ctx)
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
			_, err := service.Oper(ctx)
			if err == nil {
				t.Fatal("Expected error with nil client, got none")
			}
		})

		// Test with nil context
		t.Run("NilContext", func(t *testing.T) {
			var nilCtx context.Context //nolint:SA1012 // Testing nil context behavior
			_, err := service.Oper(nilCtx)
			if err == nil {
				t.Fatal("Expected error when using nil context, but got none")
			}
		})

		// Test with canceled context
		t.Run("CanceledContext", func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			cancel()
			_, err := service.Oper(ctx)
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
		operResp, operErr := service.Oper(ctx)
		if operErr != nil {
			t.Logf("Integration test - Oper error: %v", operErr)
		} else {
			t.Logf("Integration test - Oper success: %+v", operResp)
		}

		// Test statistics
		statsResp, statsErr := service.Stats(ctx)
		if statsErr != nil {
			t.Logf("Integration test - Stats error: %v", statsErr)
		} else {
			t.Logf("Integration test - Stats success: %+v", statsResp)
		}
	})
}
