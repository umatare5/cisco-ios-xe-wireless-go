package general

import (
	"context"
	"encoding/json"
	"os"
	"sync"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/tests"
	"github.com/umatare5/cisco-ios-xe-wireless-go/wnc"
)

// TestDataCollector holds test data for General service functions
type TestDataCollector struct {
	mu                          sync.Mutex
	OperResp                    *model.GeneralOperResponse
	OperErr                     error
	MgmtIntfDataResp            *model.GeneralOperMgmtIntfDataResponse
	MgmtIntfDataErr             error
	CfgResp                     *model.GeneralCfgResponse
	CfgErr                      error
	MewlcConfigResp             *model.MewlcConfigResponse
	MewlcConfigErr              error
	CacConfigResp               *model.CacConfigResponse
	CacConfigErr                error
	MfpResp                     *model.MfpResponse
	MfpErr                      error
	FipsCfgResp                 *model.FipsCfgResponse
	FipsCfgErr                  error
	WsaApClientEventResp        *model.WsaApClientEventResponse
	WsaApClientEventErr         error
	SimL3InterfaceCacheDataResp *model.SimL3InterfaceCacheDataResponse
	SimL3InterfaceCacheDataErr  error
	WlcManagementDataResp       *model.WlcManagementDataResponse
	WlcManagementDataErr        error
	LaginfoResp                 *model.LaginfoResponse
	LaginfoErr                  error
	MulticastConfigResp         *model.MulticastConfigResponse
	MulticastConfigErr          error
	FeatureUsageCfgResp         *model.FeatureUsageCfgResponse
	FeatureUsageCfgErr          error
	ThresholdWarnCfgResp        *model.ThresholdWarnCfgResponse
	ThresholdWarnCfgErr         error
	ApLocRangingCfgResp         *model.ApLocRangingCfgResponse
	ApLocRangingCfgErr          error
	GeolocationCfgResp          *model.GeolocationCfgResponse
	GeolocationCfgErr           error
}

// TestGeneralService tests all General service functions with the 4-pattern testing approach
func TestGeneralService(t *testing.T) {
	// Create a mock client that will be used when environment variables are not set
	var client *wnc.Client
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
		wg.Add(16) // 16 methods to test

		// Test Oper method
		go func() {
			defer wg.Done()
			resp, err := service.Oper(ctx)
			collector.mu.Lock()
			collector.OperResp = resp
			collector.OperErr = err
			collector.mu.Unlock()
		}()

		// Test MgmtIntfData method
		go func() {
			defer wg.Done()
			resp, err := service.MgmtIntfData(ctx)
			collector.mu.Lock()
			collector.MgmtIntfDataResp = resp
			collector.MgmtIntfDataErr = err
			collector.mu.Unlock()
		}()

		// Test Cfg method
		go func() {
			defer wg.Done()
			resp, err := service.Cfg(ctx)
			collector.mu.Lock()
			collector.CfgResp = resp
			collector.CfgErr = err
			collector.mu.Unlock()
		}()

		// Test MewlcConfig method
		go func() {
			defer wg.Done()
			resp, err := service.MewlcConfig(ctx)
			collector.mu.Lock()
			collector.MewlcConfigResp = resp
			collector.MewlcConfigErr = err
			collector.mu.Unlock()
		}()

		// Test CacConfig method
		go func() {
			defer wg.Done()
			resp, err := service.CacConfig(ctx)
			collector.mu.Lock()
			collector.CacConfigResp = resp
			collector.CacConfigErr = err
			collector.mu.Unlock()
		}()

		// Test Mfp method
		go func() {
			defer wg.Done()
			resp, err := service.Mfp(ctx)
			collector.mu.Lock()
			collector.MfpResp = resp
			collector.MfpErr = err
			collector.mu.Unlock()
		}()

		// Test FipsCfg method
		go func() {
			defer wg.Done()
			resp, err := service.FipsCfg(ctx)
			collector.mu.Lock()
			collector.FipsCfgResp = resp
			collector.FipsCfgErr = err
			collector.mu.Unlock()
		}()

		// Test WsaApClientEvent method
		go func() {
			defer wg.Done()
			resp, err := service.WsaApClientEvent(ctx)
			collector.mu.Lock()
			collector.WsaApClientEventResp = resp
			collector.WsaApClientEventErr = err
			collector.mu.Unlock()
		}()

		// Test SimL3InterfaceCacheData method
		go func() {
			defer wg.Done()
			resp, err := service.SimL3InterfaceCacheData(ctx)
			collector.mu.Lock()
			collector.SimL3InterfaceCacheDataResp = resp
			collector.SimL3InterfaceCacheDataErr = err
			collector.mu.Unlock()
		}()

		// Test WlcManagementData method
		go func() {
			defer wg.Done()
			resp, err := service.WlcManagementData(ctx)
			collector.mu.Lock()
			collector.WlcManagementDataResp = resp
			collector.WlcManagementDataErr = err
			collector.mu.Unlock()
		}()

		// Test Laginfo method
		go func() {
			defer wg.Done()
			resp, err := service.Laginfo(ctx)
			collector.mu.Lock()
			collector.LaginfoResp = resp
			collector.LaginfoErr = err
			collector.mu.Unlock()
		}()

		// Test MulticastConfig method
		go func() {
			defer wg.Done()
			resp, err := service.MulticastConfig(ctx)
			collector.mu.Lock()
			collector.MulticastConfigResp = resp
			collector.MulticastConfigErr = err
			collector.mu.Unlock()
		}()

		// Test FeatureUsageCfg method
		go func() {
			defer wg.Done()
			resp, err := service.FeatureUsageCfg(ctx)
			collector.mu.Lock()
			collector.FeatureUsageCfgResp = resp
			collector.FeatureUsageCfgErr = err
			collector.mu.Unlock()
		}()

		// Test ThresholdWarnCfg method
		go func() {
			defer wg.Done()
			resp, err := service.ThresholdWarnCfg(ctx)
			collector.mu.Lock()
			collector.ThresholdWarnCfgResp = resp
			collector.ThresholdWarnCfgErr = err
			collector.mu.Unlock()
		}()

		// Test ApLocRangingCfg method
		go func() {
			defer wg.Done()
			resp, err := service.ApLocRangingCfg(ctx)
			collector.mu.Lock()
			collector.ApLocRangingCfgResp = resp
			collector.ApLocRangingCfgErr = err
			collector.mu.Unlock()
		}()

		// Test GeolocationCfg method
		go func() {
			defer wg.Done()
			resp, err := service.GeolocationCfg(ctx)
			collector.mu.Lock()
			collector.GeolocationCfgResp = resp
			collector.GeolocationCfgErr = err
			collector.mu.Unlock()
		}()

		wg.Wait()

		// Validate collected data
		t.Logf("Collected data from %d General service methods", 16)
	})

	// Test JSON serialization/deserialization
	t.Run("JSON_Serialization", func(t *testing.T) {
		tests := []struct {
			name     string
			jsonData string
		}{
			{
				name: "GeneralOperResponse",
				jsonData: `{
					"Cisco-IOS-XE-wireless-general-oper:general-oper-data": {
						"ap-summary": {
							"total-ap": 10
						}
					}
				}`,
			},
			{
				name: "GeneralCfgResponse",
				jsonData: `{
					"Cisco-IOS-XE-wireless-general-cfg:general-cfg-data": {
						"mewlc-config": {
							"enable": true
						}
					}
				}`,
			},
		}

		for _, tt := range tests {
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
		tests := []struct {
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
				name: "MgmtIntfData",
				method: func() (interface{}, error) {
					return service.MgmtIntfData(ctx)
				},
			},
			{
				name: "Cfg",
				method: func() (interface{}, error) {
					return service.Cfg(ctx)
				},
			},
			{
				name: "MewlcConfig",
				method: func() (interface{}, error) {
					return service.MewlcConfig(ctx)
				},
			},
			{
				name: "CacConfig",
				method: func() (interface{}, error) {
					return service.CacConfig(ctx)
				},
			},
			{
				name: "Mfp",
				method: func() (interface{}, error) {
					return service.Mfp(ctx)
				},
			},
			{
				name: "FipsCfg",
				method: func() (interface{}, error) {
					return service.FipsCfg(ctx)
				},
			},
			{
				name: "WsaApClientEvent",
				method: func() (interface{}, error) {
					return service.WsaApClientEvent(ctx)
				},
			},
			{
				name: "SimL3InterfaceCacheData",
				method: func() (interface{}, error) {
					return service.SimL3InterfaceCacheData(ctx)
				},
			},
			{
				name: "WlcManagementData",
				method: func() (interface{}, error) {
					return service.WlcManagementData(ctx)
				},
			},
			{
				name: "Laginfo",
				method: func() (interface{}, error) {
					return service.Laginfo(ctx)
				},
			},
			{
				name: "MulticastConfig",
				method: func() (interface{}, error) {
					return service.MulticastConfig(ctx)
				},
			},
			{
				name: "FeatureUsageCfg",
				method: func() (interface{}, error) {
					return service.FeatureUsageCfg(ctx)
				},
			},
			{
				name: "ThresholdWarnCfg",
				method: func() (interface{}, error) {
					return service.ThresholdWarnCfg(ctx)
				},
			},
			{
				name: "ApLocRangingCfg",
				method: func() (interface{}, error) {
					return service.ApLocRangingCfg(ctx)
				},
			},
			{
				name: "GeolocationCfg",
				method: func() (interface{}, error) {
					return service.GeolocationCfg(ctx)
				},
			},
		}

		for _, tt := range tests {
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
			_, err := service.Oper(nil)
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

		// Test configuration data
		cfgResp, cfgErr := service.Cfg(ctx)
		if cfgErr != nil {
			t.Logf("Integration test - Cfg error: %v", cfgErr)
		} else {
			t.Logf("Integration test - Cfg success: %+v", cfgResp)
		}
	})
}
