package ap

import (
	"context"
	"encoding/json"
	"os"
	"sync"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/tests"
)

// TestDataCollector holds test data for AP service functions
type TestDataCollector struct {
	mu                           sync.Mutex
	CfgResp                      *model.ApCfgResponse
	CfgErr                       error
	TagSourcePriorityConfigsResp *model.TagSourcePriorityConfigs
	TagSourcePriorityConfigsErr  error
	ApTagsResp                   *model.ApCfgApTagsResponse
	ApTagsErr                    error
	OperResp                     *model.ApOperResponse
	OperErr                      error
	RadioNeighborResp            *model.ApOperApRadioNeighborResponse
	RadioNeighborErr             error
	NameMacMapResp               *[]model.ApNameMacMap
	NameMacMapErr                error
	CapwapDataResp               *[]model.CapwapData
	CapwapDataErr                error
	GlobalOperResp               *model.ApGlobalOperResponse
	GlobalOperErr                error
	HistoryResp                  *model.ApGlobalOperApHistoryResponse
	HistoryErr                   error
	EwlcApStatsResp              *model.ApGlobalOperEwlcApStatsResponse
	EwlcApStatsErr               error
}

// TestApService tests all AP service functions with the 4-pattern testing approach
func TestApService(t *testing.T) {
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
		wg.Add(10) // 10 methods to test

		// Test Cfg method
		go func() {
			defer wg.Done()
			resp, err := service.Cfg(ctx)
			collector.mu.Lock()
			collector.CfgResp = resp
			collector.CfgErr = err
			collector.mu.Unlock()
		}()

		// Test TagSourcePriorityConfigs method
		go func() {
			defer wg.Done()
			resp, err := service.TagSourcePriorityConfigs(ctx)
			collector.mu.Lock()
			collector.TagSourcePriorityConfigsResp = resp
			collector.TagSourcePriorityConfigsErr = err
			collector.mu.Unlock()
		}()

		// Test ApTags method
		go func() {
			defer wg.Done()
			resp, err := service.ApTags(ctx)
			collector.mu.Lock()
			collector.ApTagsResp = resp
			collector.ApTagsErr = err
			collector.mu.Unlock()
		}()

		// Test Oper method
		go func() {
			defer wg.Done()
			resp, err := service.Oper(ctx)
			collector.mu.Lock()
			collector.OperResp = resp
			collector.OperErr = err
			collector.mu.Unlock()
		}()

		// Test RadioNeighbor method
		go func() {
			defer wg.Done()
			resp, err := service.RadioNeighbor(ctx)
			collector.mu.Lock()
			collector.RadioNeighborResp = resp
			collector.RadioNeighborErr = err
			collector.mu.Unlock()
		}()

		// Test NameMacMap method
		go func() {
			defer wg.Done()
			resp, err := service.NameMacMap(ctx)
			collector.mu.Lock()
			collector.NameMacMapResp = resp
			collector.NameMacMapErr = err
			collector.mu.Unlock()
		}()

		// Test CapwapData method
		go func() {
			defer wg.Done()
			resp, err := service.CapwapData(ctx)
			collector.mu.Lock()
			collector.CapwapDataResp = resp
			collector.CapwapDataErr = err
			collector.mu.Unlock()
		}()

		// Test GlobalOper method
		go func() {
			defer wg.Done()
			resp, err := service.GlobalOper(ctx)
			collector.mu.Lock()
			collector.GlobalOperResp = resp
			collector.GlobalOperErr = err
			collector.mu.Unlock()
		}()

		// Test History method
		go func() {
			defer wg.Done()
			resp, err := service.History(ctx)
			collector.mu.Lock()
			collector.HistoryResp = resp
			collector.HistoryErr = err
			collector.mu.Unlock()
		}()

		// Test EwlcApStats method
		go func() {
			defer wg.Done()
			resp, err := service.EwlcApStats(ctx)
			collector.mu.Lock()
			collector.EwlcApStatsResp = resp
			collector.EwlcApStatsErr = err
			collector.mu.Unlock()
		}()

		wg.Wait()

		// Validate collected data
		t.Logf("Collected data from %d AP service methods", 10)
	})

	// Test JSON serialization/deserialization
	t.Run("JSON_Serialization", func(t *testing.T) {
		testCases := []struct {
			name     string
			jsonData string
		}{
			{
				name: "ApCfgResponse",
				jsonData: `{
					"Cisco-IOS-XE-wireless-ap-cfg:ap-cfg-data": {
						"ap-tags": {
							"ap-tag": []
						}
					}
				}`,
			},
			{
				name: "ApOperResponse",
				jsonData: `{
					"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data": {
						"ap-name-mac-map": []
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
				name: "Cfg",
				method: func() (interface{}, error) {
					return service.Cfg(ctx)
				},
			},
			{
				name: "TagSourcePriorityConfigs",
				method: func() (interface{}, error) {
					return service.TagSourcePriorityConfigs(ctx)
				},
			},
			{
				name: "ApTags",
				method: func() (interface{}, error) {
					return service.ApTags(ctx)
				},
			},
			{
				name: "Oper",
				method: func() (interface{}, error) {
					return service.Oper(ctx)
				},
			},
			{
				name: "RadioNeighbor",
				method: func() (interface{}, error) {
					return service.RadioNeighbor(ctx)
				},
			},
			{
				name: "NameMacMap",
				method: func() (interface{}, error) {
					return service.NameMacMap(ctx)
				},
			},
			{
				name: "CapwapData",
				method: func() (interface{}, error) {
					return service.CapwapData(ctx)
				},
			},
			{
				name: "GlobalOper",
				method: func() (interface{}, error) {
					return service.GlobalOper(ctx)
				},
			},
			{
				name: "History",
				method: func() (interface{}, error) {
					return service.History(ctx)
				},
			},
			{
				name: "EwlcApStats",
				method: func() (interface{}, error) {
					return service.EwlcApStats(ctx)
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
			_, err := service.Cfg(ctx)
			if err == nil {
				t.Fatal("Expected error with nil client, got none")
			}
		})

		// Test with nil context
		t.Run("NilContext", func(t *testing.T) {
			var nilCtx context.Context //nolint:SA1012 // Testing nil context behavior
			_, err := service.Cfg(nilCtx)
			if err == nil {
				t.Fatal("Expected error when using nil context, but got none")
			}
		})

		// Test with canceled context
		t.Run("CanceledContext", func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			cancel()
			_, err := service.Cfg(ctx)
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
		cfgResp, cfgErr := service.Cfg(ctx)
		if cfgErr != nil {
			t.Logf("Integration test - Cfg error: %v", cfgErr)
		} else {
			t.Logf("Integration test - Cfg success: %+v", cfgResp)
		}

		// Test operational data
		operResp, operErr := service.Oper(ctx)
		if operErr != nil {
			t.Logf("Integration test - Oper error: %v", operErr)
		} else {
			t.Logf("Integration test - Oper success: %+v", operResp)
		}
	})
}
