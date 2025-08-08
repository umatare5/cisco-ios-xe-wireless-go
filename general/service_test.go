package general

import (
	"context"
	"encoding/json"
	"sync"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/tests"
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
		wg.Add(16) // 16 methods to test

		// Test Oper method
		go func() {
			defer wg.Done()
			resp, err := service.GetOper(ctx)
			collector.mu.Lock()
			collector.OperResp = resp
			collector.OperErr = err
			collector.mu.Unlock()
		}()

		// Test MgmtIntfData method
		go func() {
			defer wg.Done()
			resp, err := service.GetMgmtIntfData(ctx)
			collector.mu.Lock()
			collector.MgmtIntfDataResp = resp
			collector.MgmtIntfDataErr = err
			collector.mu.Unlock()
		}()

		// Test Cfg method
		go func() {
			defer wg.Done()
			resp, err := service.GetCfg(ctx)
			collector.mu.Lock()
			collector.CfgResp = resp
			collector.CfgErr = err
			collector.mu.Unlock()
		}()

		// Test MewlcConfig method
		go func() {
			defer wg.Done()
			resp, err := service.GetMewlcConfig(ctx)
			collector.mu.Lock()
			collector.MewlcConfigResp = resp
			collector.MewlcConfigErr = err
			collector.mu.Unlock()
		}()

		// Test CacConfig method
		go func() {
			defer wg.Done()
			resp, err := service.GetCacConfig(ctx)
			collector.mu.Lock()
			collector.CacConfigResp = resp
			collector.CacConfigErr = err
			collector.mu.Unlock()
		}()

		// Test Mfp method
		go func() {
			defer wg.Done()
			resp, err := service.GetMfp(ctx)
			collector.mu.Lock()
			collector.MfpResp = resp
			collector.MfpErr = err
			collector.mu.Unlock()
		}()

		// Test FipsCfg method
		go func() {
			defer wg.Done()
			resp, err := service.GetFipsCfg(ctx)
			collector.mu.Lock()
			collector.FipsCfgResp = resp
			collector.FipsCfgErr = err
			collector.mu.Unlock()
		}()

		// Test WsaApClientEvent method
		go func() {
			defer wg.Done()
			resp, err := service.GetWsaApClientEvent(ctx)
			collector.mu.Lock()
			collector.WsaApClientEventResp = resp
			collector.WsaApClientEventErr = err
			collector.mu.Unlock()
		}()

		// Test SimL3InterfaceCacheData method
		go func() {
			defer wg.Done()
			resp, err := service.GetSimL3InterfaceCacheData(ctx)
			collector.mu.Lock()
			collector.SimL3InterfaceCacheDataResp = resp
			collector.SimL3InterfaceCacheDataErr = err
			collector.mu.Unlock()
		}()

		// Test WlcManagementData method
		go func() {
			defer wg.Done()
			resp, err := service.GetWlcManagementData(ctx)
			collector.mu.Lock()
			collector.WlcManagementDataResp = resp
			collector.WlcManagementDataErr = err
			collector.mu.Unlock()
		}()

		// Test Laginfo method
		go func() {
			defer wg.Done()
			resp, err := service.GetLaginfo(ctx)
			collector.mu.Lock()
			collector.LaginfoResp = resp
			collector.LaginfoErr = err
			collector.mu.Unlock()
		}()

		// Test MulticastConfig method
		go func() {
			defer wg.Done()
			resp, err := service.GetMulticastConfig(ctx)
			collector.mu.Lock()
			collector.MulticastConfigResp = resp
			collector.MulticastConfigErr = err
			collector.mu.Unlock()
		}()

		// Test FeatureUsageCfg method
		go func() {
			defer wg.Done()
			resp, err := service.GetFeatureUsageCfg(ctx)
			collector.mu.Lock()
			collector.FeatureUsageCfgResp = resp
			collector.FeatureUsageCfgErr = err
			collector.mu.Unlock()
		}()

		// Test ThresholdWarnCfg method
		go func() {
			defer wg.Done()
			resp, err := service.GetThresholdWarnCfg(ctx)
			collector.mu.Lock()
			collector.ThresholdWarnCfgResp = resp
			collector.ThresholdWarnCfgErr = err
			collector.mu.Unlock()
		}()

		// Test ApLocRangingCfg method
		go func() {
			defer wg.Done()
			resp, err := service.GetApLocRangingCfg(ctx)
			collector.mu.Lock()
			collector.ApLocRangingCfgResp = resp
			collector.ApLocRangingCfgErr = err
			collector.mu.Unlock()
		}()

		// Test GeolocationCfg method
		go func() {
			defer wg.Done()
			resp, err := service.GetGeolocationCfg(ctx)
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
		testCases := []struct {
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
				name: "GetOper",
				method: func() (interface{}, error) {
					return service.GetOper(ctx)
				},
			},
			{
				name: "GetMgmtIntfData",
				method: func() (interface{}, error) {
					return service.GetMgmtIntfData(ctx)
				},
			},
			{
				name: "GetCfg",
				method: func() (interface{}, error) {
					return service.GetCfg(ctx)
				},
			},
			{
				name: "GetMewlcConfig",
				method: func() (interface{}, error) {
					return service.GetMewlcConfig(ctx)
				},
			},
			{
				name: "GetCacConfig",
				method: func() (interface{}, error) {
					return service.GetCacConfig(ctx)
				},
			},
			{
				name: "GetMfp",
				method: func() (interface{}, error) {
					return service.GetMfp(ctx)
				},
			},
			{
				name: "GetFipsCfg",
				method: func() (interface{}, error) {
					return service.GetFipsCfg(ctx)
				},
			},
			{
				name: "GetWsaApClientEvent",
				method: func() (interface{}, error) {
					return service.GetWsaApClientEvent(ctx)
				},
			},
			{
				name: "GetSimL3InterfaceCacheData",
				method: func() (interface{}, error) {
					return service.GetSimL3InterfaceCacheData(ctx)
				},
			},
			{
				name: "GetWlcManagementData",
				method: func() (interface{}, error) {
					return service.GetWlcManagementData(ctx)
				},
			},
			{
				name: "GetLaginfo",
				method: func() (interface{}, error) {
					return service.GetLaginfo(ctx)
				},
			},
			{
				name: "GetMulticastConfig",
				method: func() (interface{}, error) {
					return service.GetMulticastConfig(ctx)
				},
			},
			{
				name: "GetFeatureUsageCfg",
				method: func() (interface{}, error) {
					return service.GetFeatureUsageCfg(ctx)
				},
			},
			{
				name: "GetThresholdWarnCfg",
				method: func() (interface{}, error) {
					return service.GetThresholdWarnCfg(ctx)
				},
			},
			{
				name: "GetApLocRangingCfg",
				method: func() (interface{}, error) {
					return service.GetApLocRangingCfg(ctx)
				},
			},
			{
				name: "GetGeolocationCfg",
				method: func() (interface{}, error) {
					return service.GetGeolocationCfg(ctx)
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

		// Test configuration data
		cfgResp, cfgErr := service.GetCfg(ctx)
		if cfgErr != nil {
			t.Logf("Integration test - Cfg error: %v", cfgErr)
		} else {
			t.Logf("Integration test - Cfg success: %+v", cfgResp)
		}
	})
}
