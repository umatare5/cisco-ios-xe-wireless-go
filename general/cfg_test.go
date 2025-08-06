// Package general provides general configuration test functionality for the Cisco Wireless Network Controller API.
package general

import (
	"context"
	"encoding/json"
	"testing"

	wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
	testutils "github.com/umatare5/cisco-ios-xe-wireless-go/internal/tests"
)

// =============================================================================
// 1. UNIT TESTS (Structure/Type Validation & JSON Serialization/Deserialization)
// =============================================================================

// getTestClientCfg creates a test client using environment variables
func getTestClientCfg(t *testing.T) *wnc.Client {
	return testutils.CreateTestClientFromEnv(t)
}

// GeneralCfgTestDataCollector holds test data for general configuration functions
type GeneralCfgTestDataCollector struct {
	Data map[string]interface{} `json:"general_cfg_test_data"`
}

// newGeneralCfgTestDataCollector creates a new test data collector
func newGeneralCfgTestDataCollector() *GeneralCfgTestDataCollector {
	return &GeneralCfgTestDataCollector{
		Data: make(map[string]interface{}),
	}
}

func runGeneralCfgTestAndCollectData(t *testing.T, collector *GeneralCfgTestDataCollector, testName string, testFunc func() (interface{}, error)) {
	data, err := testFunc()
	if err != nil {
		t.Logf("%s returned error: %v", testName, err)
		collector.Data[testName] = map[string]interface{}{
			"error":   err.Error(),
			"success": false,
		}
	} else {
		t.Logf("%s executed successfully", testName)
		collector.Data[testName] = map[string]interface{}{
			"data":    data,
			"success": true,
		}
	}
}

// =============================================================================
// 2. INTEGRATION TESTS (API Endpoint Testing with Live Data Validation)
// =============================================================================

// TestGeneralConfigurationFunctions tests all general configuration functions with real WNC data collection
func TestGeneralConfigurationFunctions(t *testing.T) {
	client := getTestClientCfg(t)
	collector := newGeneralCfgTestDataCollector()

	ctx, cancel := testutils.CreateDefaultTestContext()
	defer cancel()

	t.Run("GetGeneralCfg", func(t *testing.T) {
		runGeneralCfgTestAndCollectData(t, collector, "GetGeneralCfg", func() (interface{}, error) {
			return GetGeneralCfg(client, ctx)
		})
	})

	t.Run("GetGeneralMewlcConfig", func(t *testing.T) {
		resp, err := GetGeneralMewlcConfig(client, ctx)
		if err != nil {
			t.Logf("GetGeneralMewlcConfig returned error (may be expected): %v", err)
		} else {
			t.Log("GetGeneralMewlcConfig executed successfully")
			if resp != nil {
				t.Logf("Response contains MEWLC configuration data")
			}
		}
	})

	t.Run("GetGeneralCacConfig", func(t *testing.T) {
		resp, err := GetGeneralCacConfig(client, ctx)
		if err != nil {
			t.Logf("GetGeneralCacConfig returned error (may be expected): %v", err)
		} else {
			t.Log("GetGeneralCacConfig executed successfully")
			if resp != nil {
				t.Logf("Response contains CAC configuration data")
			}
		}
	})

	t.Run("GetGeneralMfp", func(t *testing.T) {
		resp, err := GetGeneralMfp(client, ctx)
		if err != nil {
			t.Logf("GetGeneralMfp returned error (may be expected): %v", err)
		} else {
			t.Log("GetGeneralMfp executed successfully")
			if resp != nil {
				t.Logf("Response contains MFP configuration data")
			}
		}
	})

	t.Run("GetGeneralFipsCfg", func(t *testing.T) {
		resp, err := GetGeneralFipsCfg(client, ctx)
		if err != nil {
			t.Logf("GetGeneralFipsCfg returned error (may be expected): %v", err)
		} else {
			t.Log("GetGeneralFipsCfg executed successfully")
			if resp != nil {
				t.Logf("Response contains FIPS configuration data")
			}
		}
	})

	t.Run("GetGeneralWsaApClientEvent", func(t *testing.T) {
		resp, err := GetGeneralWsaApClientEvent(client, ctx)
		if err != nil {
			t.Logf("GetGeneralWsaApClientEvent returned error (may be expected): %v", err)
		} else {
			t.Log("GetGeneralWsaApClientEvent executed successfully")
			if resp != nil {
				t.Logf("Response contains WSA AP client event data")
			}
		}
	})

	t.Run("GetGeneralSimL3InterfaceCacheData", func(t *testing.T) {
		resp, err := GetGeneralSimL3InterfaceCacheData(client, ctx)
		if err != nil {
			t.Logf("GetGeneralSimL3InterfaceCacheData returned error (may be expected): %v", err)
		} else {
			t.Log("GetGeneralSimL3InterfaceCacheData executed successfully")
			if resp != nil {
				t.Logf("Response contains SIM L3 interface cache data")
			}
		}
	})

	t.Run("GetGeneralWlcManagementData", func(t *testing.T) {
		resp, err := GetGeneralWlcManagementData(client, ctx)
		if err != nil {
			t.Logf("GetGeneralWlcManagementData returned error (may be expected): %v", err)
		} else {
			t.Log("GetGeneralWlcManagementData executed successfully")
			if resp != nil {
				t.Logf("Response contains WLC management data")
			}
		}
	})

	t.Run("GetGeneralLaginfo", func(t *testing.T) {
		resp, err := GetGeneralLaginfo(client, ctx)
		if err != nil {
			t.Logf("GetGeneralLaginfo returned error (may be expected): %v", err)
		} else {
			t.Log("GetGeneralLaginfo executed successfully")
			if resp != nil {
				t.Logf("Response contains LAG information")
			}
		}
	})

	t.Run("GetGeneralMulticastConfig", func(t *testing.T) {
		resp, err := GetGeneralMulticastConfig(client, ctx)
		if err != nil {
			t.Logf("GetGeneralMulticastConfig returned error (may be expected): %v", err)
		} else {
			t.Log("GetGeneralMulticastConfig executed successfully")
			if resp != nil {
				t.Logf("Response contains multicast configuration")
			}
		}
	})

	t.Run("GetGeneralFeatureUsageCfg", func(t *testing.T) {
		resp, err := GetGeneralFeatureUsageCfg(client, ctx)
		if err != nil {
			t.Logf("GetGeneralFeatureUsageCfg returned error (may be expected): %v", err)
		} else {
			t.Log("GetGeneralFeatureUsageCfg executed successfully")
			if resp != nil {
				t.Logf("Response contains feature usage configuration")
			}
		}
	})

	t.Run("GetGeneralThresholdWarnCfg", func(t *testing.T) {
		resp, err := GetGeneralThresholdWarnCfg(client, ctx)
		if err != nil {
			t.Logf("GetGeneralThresholdWarnCfg returned error (may be expected): %v", err)
		} else {
			t.Log("GetGeneralThresholdWarnCfg executed successfully")
			if resp != nil {
				t.Logf("Response contains threshold warning configuration")
			}
		}
	})

	t.Run("GetGeneralApLocRangingCfg", func(t *testing.T) {
		resp, err := GetGeneralApLocRangingCfg(client, ctx)
		if err != nil {
			t.Logf("GetGeneralApLocRangingCfg returned error (may be expected): %v", err)
		} else {
			t.Log("GetGeneralApLocRangingCfg executed successfully")
			if resp != nil {
				t.Logf("Response contains AP location ranging configuration")
			}
		}
	})

	t.Run("GetGeneralGeolocationCfg", func(t *testing.T) {
		resp, err := GetGeneralGeolocationCfg(client, ctx)
		if err != nil {
			t.Logf("GetGeneralGeolocationCfg returned error (may be expected): %v", err)
		} else {
			t.Log("GetGeneralGeolocationCfg executed successfully")
			if resp != nil {
				t.Logf("Response contains geolocation configuration")
			}
		}
	})

	// Save collected test data to file
	if len(collector.Data) > 0 {
		testutils.SaveTestDataWithLogging("general_cfg_test_data_collected.json", collector.Data)
	}
}

// TestGeneralConfigurationEndpoints tests that all endpoints are correctly defined
func TestGeneralConfigurationEndpoints(t *testing.T) {
	expectedEndpoints := map[string]string{
		"GeneralCfgEndpoint":              "Cisco-IOS-XE-wireless-general-cfg:general-cfg-data",
		"MewlcConfigEndpoint":             "Cisco-IOS-XE-wireless-general-cfg:general-cfg-data/mewlc-config",
		"CacConfigEndpoint":               "Cisco-IOS-XE-wireless-general-cfg:general-cfg-data/cac-config",
		"MfpEndpoint":                     "Cisco-IOS-XE-wireless-general-cfg:general-cfg-data/mfp",
		"FipsCfgEndpoint":                 "Cisco-IOS-XE-wireless-general-cfg:general-cfg-data/fips-cfg",
		"WsaApClientEventEndpoint":        "Cisco-IOS-XE-wireless-general-cfg:general-cfg-data/wsa-ap-client-event",
		"SimL3InterfaceCacheDataEndpoint": "Cisco-IOS-XE-wireless-general-cfg:general-cfg-data/sim-l3-interface-cache-data",
		"WlcManagementDataEndpoint":       "Cisco-IOS-XE-wireless-general-cfg:general-cfg-data/wlc-management-data",
		"LaginfoEndpoint":                 "Cisco-IOS-XE-wireless-general-cfg:general-cfg-data/laginfo",
		"MulticastConfigEndpoint":         "Cisco-IOS-XE-wireless-general-cfg:general-cfg-data/multicast-config",
		"FeatureUsageCfgEndpoint":         "Cisco-IOS-XE-wireless-general-cfg:general-cfg-data/feature-usage-cfg",
		"ThresholdWarnCfgEndpoint":        "Cisco-IOS-XE-wireless-general-cfg:general-cfg-data/threshold-warn-cfg",
		"ApLocRangingCfgEndpoint":         "Cisco-IOS-XE-wireless-general-cfg:general-cfg-data/ap-loc-ranging-cfg",
		"GeolocationCfgEndpoint":          "Cisco-IOS-XE-wireless-general-cfg:general-cfg-data/geolocation-cfg",
	}

	for name, expected := range expectedEndpoints {
		t.Run(name, func(t *testing.T) {
			switch name {
			case "GeneralCfgEndpoint":
				if GeneralCfgEndpoint != expected {
					t.Errorf("Expected %s = %s, got %s", name, expected, GeneralCfgEndpoint)
				}
			case "MewlcConfigEndpoint":
				if MewlcConfigEndpoint != expected {
					t.Errorf("Expected %s = %s, got %s", name, expected, MewlcConfigEndpoint)
				}
			case "CacConfigEndpoint":
				if CacConfigEndpoint != expected {
					t.Errorf("Expected %s = %s, got %s", name, expected, CacConfigEndpoint)
				}
			case "MfpEndpoint":
				if MfpEndpoint != expected {
					t.Errorf("Expected %s = %s, got %s", name, expected, MfpEndpoint)
				}
			case "FipsCfgEndpoint":
				if FipsCfgEndpoint != expected {
					t.Errorf("Expected %s = %s, got %s", name, expected, FipsCfgEndpoint)
				}
			case "WsaApClientEventEndpoint":
				if WsaApClientEventEndpoint != expected {
					t.Errorf("Expected %s = %s, got %s", name, expected, WsaApClientEventEndpoint)
				}
			case "SimL3InterfaceCacheDataEndpoint":
				if SimL3InterfaceCacheDataEndpoint != expected {
					t.Errorf("Expected %s = %s, got %s", name, expected, SimL3InterfaceCacheDataEndpoint)
				}
			case "WlcManagementDataEndpoint":
				if WlcManagementDataEndpoint != expected {
					t.Errorf("Expected %s = %s, got %s", name, expected, WlcManagementDataEndpoint)
				}
			case "LaginfoEndpoint":
				if LaginfoEndpoint != expected {
					t.Errorf("Expected %s = %s, got %s", name, expected, LaginfoEndpoint)
				}
			case "MulticastConfigEndpoint":
				if MulticastConfigEndpoint != expected {
					t.Errorf("Expected %s = %s, got %s", name, expected, MulticastConfigEndpoint)
				}
			case "FeatureUsageCfgEndpoint":
				if FeatureUsageCfgEndpoint != expected {
					t.Errorf("Expected %s = %s, got %s", name, expected, FeatureUsageCfgEndpoint)
				}
			case "ThresholdWarnCfgEndpoint":
				if ThresholdWarnCfgEndpoint != expected {
					t.Errorf("Expected %s = %s, got %s", name, expected, ThresholdWarnCfgEndpoint)
				}
			case "ApLocRangingCfgEndpoint":
				if ApLocRangingCfgEndpoint != expected {
					t.Errorf("Expected %s = %s, got %s", name, expected, ApLocRangingCfgEndpoint)
				}
			case "GeolocationCfgEndpoint":
				if GeolocationCfgEndpoint != expected {
					t.Errorf("Expected %s = %s, got %s", name, expected, GeolocationCfgEndpoint)
				}
			}
		})
	}
}

// TestGeneralCfgDataStructures tests the basic structure of general configuration data types
func TestGeneralCfgDataStructures(t *testing.T) {
	tests := []struct {
		name     string
		jsonData string
		dataType interface{}
	}{
		{
			name: "MewlcConfigResponse",
			jsonData: `{
				"Cisco-IOS-XE-wireless-general-cfg:mewlc-config": {
					"ewlc-enable": true,
					"ewlc-tunnel-mode": "capwap",
					"mewlc-heartbeat-interval": 30,
					"mewlc-discovery-interval": 10,
					"mewlc-retry-count": 3,
					"mewlc-stats-interval": 60
				}
			}`,
			dataType: &MewlcConfigResponse{},
		},
		{
			name: "CacConfigResponse",
			jsonData: `{
				"Cisco-IOS-XE-wireless-general-cfg:cac-config": {
					"cac-enable": true,
					"cac-voice-sip-bandwidth": 64,
					"cac-voice-bandwidth": 23437,
					"cac-video-bandwidth": 0,
					"cac-video-max-bandwidth": 0,
					"cac-voice-roam-bandwidth": 6,
					"load-balancing-window": 5
				}
			}`,
			dataType: &CacConfigResponse{},
		},
		{
			name: "MfpResponse",
			jsonData: `{
				"Cisco-IOS-XE-wireless-general-cfg:mfp": {
					"mfp-client-protection": "optional",
					"mfp-infrastructure-protection": "optional",
					"mfp-comeback-timer": 300,
					"mfp-sa-query-timeout": 200,
					"mfp-sa-query-retry-timeout": 200
				}
			}`,
			dataType: &MfpResponse{},
		},
		{
			name: "FipsCfgResponse",
			jsonData: `{
				"Cisco-IOS-XE-wireless-general-cfg:fips-cfg": {
					"fips-enable": false,
					"fips-auth-type": "none",
					"fips-encryption": "none",
					"fips-data-encryption": "none"
				}
			}`,
			dataType: &FipsCfgResponse{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := json.Unmarshal([]byte(tt.jsonData), tt.dataType)
			if err != nil {
				t.Errorf("Failed to unmarshal %s: %v", tt.name, err)
			}

			_, err = json.Marshal(tt.dataType)
			if err != nil {
				t.Errorf("Failed to marshal %s: %v", tt.name, err)
			}
		})
	}
}

// =============================================================================
// 4. ERROR HANDLING TESTS
// =============================================================================

// TestGeneralCfgErrorHandling tests error handling for all general configuration functions
func TestGeneralCfgErrorHandling(t *testing.T) {
	ctx := context.Background()

	// Test all 14 functions with nil client
	testCases := []struct {
		name string
		fn   func() (interface{}, error)
	}{
		{"GetGeneralCfg", func() (interface{}, error) { return GetGeneralCfg(nil, ctx) }},
		{"GetGeneralMewlcConfig", func() (interface{}, error) { return GetGeneralMewlcConfig(nil, ctx) }},
		{"GetGeneralCacConfig", func() (interface{}, error) { return GetGeneralCacConfig(nil, ctx) }},
		{"GetGeneralMfp", func() (interface{}, error) { return GetGeneralMfp(nil, ctx) }},
		{"GetGeneralFipsCfg", func() (interface{}, error) { return GetGeneralFipsCfg(nil, ctx) }},
		{"GetGeneralWsaApClientEvent", func() (interface{}, error) { return GetGeneralWsaApClientEvent(nil, ctx) }},
		{"GetGeneralSimL3InterfaceCacheData", func() (interface{}, error) { return GetGeneralSimL3InterfaceCacheData(nil, ctx) }},
		{"GetGeneralWlcManagementData", func() (interface{}, error) { return GetGeneralWlcManagementData(nil, ctx) }},
		{"GetGeneralLaginfo", func() (interface{}, error) { return GetGeneralLaginfo(nil, ctx) }},
		{"GetGeneralMulticastConfig", func() (interface{}, error) { return GetGeneralMulticastConfig(nil, ctx) }},
		{"GetGeneralFeatureUsageCfg", func() (interface{}, error) { return GetGeneralFeatureUsageCfg(nil, ctx) }},
		{"GetGeneralThresholdWarnCfg", func() (interface{}, error) { return GetGeneralThresholdWarnCfg(nil, ctx) }},
		{"GetGeneralApLocRangingCfg", func() (interface{}, error) { return GetGeneralApLocRangingCfg(nil, ctx) }},
		{"GetGeneralGeolocationCfg", func() (interface{}, error) { return GetGeneralGeolocationCfg(nil, ctx) }},
	}

	for _, tc := range testCases {
		t.Run(tc.name+"WithNilClient", func(t *testing.T) {
			_, err := tc.fn()
			if err == nil || err.Error() != "invalid client configuration: client cannot be nil" {
				t.Errorf("Expected 'client is nil' error, got: %v", err)
			}
		})
	}
}

// =============================================================================
// 5. CONTEXT HANDLING TESTS
// =============================================================================

// TestGeneralCfgContextHandling tests context handling for all general configuration functions
func TestGeneralCfgContextHandling(t *testing.T) {
	// Test all 14 functions with context handling
	testCases := []struct {
		name string
		fn   func(context.Context, *wnc.Client) error
	}{
		{"GetGeneralCfg", func(ctx context.Context, client *wnc.Client) error { _, err := GetGeneralCfg(client, ctx); return err }},
		{"GetGeneralMewlcConfig", func(ctx context.Context, client *wnc.Client) error {
			_, err := GetGeneralMewlcConfig(client, ctx)
			return err
		}},
		{"GetGeneralCacConfig", func(ctx context.Context, client *wnc.Client) error {
			_, err := GetGeneralCacConfig(client, ctx)
			return err
		}},
		{"GetGeneralMfp", func(ctx context.Context, client *wnc.Client) error { _, err := GetGeneralMfp(client, ctx); return err }},
		{"GetGeneralFipsCfg", func(ctx context.Context, client *wnc.Client) error {
			_, err := GetGeneralFipsCfg(client, ctx)
			return err
		}},
		{"GetGeneralWsaApClientEvent", func(ctx context.Context, client *wnc.Client) error {
			_, err := GetGeneralWsaApClientEvent(client, ctx)
			return err
		}},
		{"GetGeneralSimL3InterfaceCacheData", func(ctx context.Context, client *wnc.Client) error {
			_, err := GetGeneralSimL3InterfaceCacheData(client, ctx)
			return err
		}},
		{"GetGeneralWlcManagementData", func(ctx context.Context, client *wnc.Client) error {
			_, err := GetGeneralWlcManagementData(client, ctx)
			return err
		}},
		{"GetGeneralLaginfo", func(ctx context.Context, client *wnc.Client) error {
			_, err := GetGeneralLaginfo(client, ctx)
			return err
		}},
		{"GetGeneralMulticastConfig", func(ctx context.Context, client *wnc.Client) error {
			_, err := GetGeneralMulticastConfig(client, ctx)
			return err
		}},
		{"GetGeneralFeatureUsageCfg", func(ctx context.Context, client *wnc.Client) error {
			_, err := GetGeneralFeatureUsageCfg(client, ctx)
			return err
		}},
		{"GetGeneralThresholdWarnCfg", func(ctx context.Context, client *wnc.Client) error {
			_, err := GetGeneralThresholdWarnCfg(client, ctx)
			return err
		}},
		{"GetGeneralApLocRangingCfg", func(ctx context.Context, client *wnc.Client) error {
			_, err := GetGeneralApLocRangingCfg(client, ctx)
			return err
		}},
		{"GetGeneralGeolocationCfg", func(ctx context.Context, client *wnc.Client) error {
			_, err := GetGeneralGeolocationCfg(client, ctx)
			return err
		}},
	}

	for _, tc := range testCases {
		t.Run(tc.name+"ContextHandling", func(t *testing.T) {
			testutils.TestContextHandling(t, tc.fn)
		})
	}
}

// =============================================================================
// 7. SERVICE TESTS
// =============================================================================

func TestGeneralServiceConfiguration(t *testing.T) {
	client := testutils.GetTestClient(t)
	if client == nil {
		t.Skip("Skipping service tests: no test client available")
	}

	ctx := context.Background()
	service := NewService(client.CoreClient())

	// Test all configuration service methods
	t.Run("Service_Cfg", func(t *testing.T) {
		testutils.TestServiceMethod(t, func() error {
			_, err := service.Cfg(ctx)
			return err
		})
	})

	t.Run("Service_MewlcConfig", func(t *testing.T) {
		testutils.TestServiceMethod(t, func() error {
			_, err := service.MewlcConfig(ctx)
			return err
		})
	})

	t.Run("Service_CacConfig", func(t *testing.T) {
		testutils.TestServiceMethod(t, func() error {
			_, err := service.CacConfig(ctx)
			return err
		})
	})

	t.Run("Service_Mfp", func(t *testing.T) {
		testutils.TestServiceMethod(t, func() error {
			_, err := service.Mfp(ctx)
			return err
		})
	})

	t.Run("Service_FipsCfg", func(t *testing.T) {
		testutils.TestServiceMethod(t, func() error {
			_, err := service.FipsCfg(ctx)
			return err
		})
	})

	t.Run("Service_WsaApClientEvent", func(t *testing.T) {
		testutils.TestServiceMethod(t, func() error {
			_, err := service.WsaApClientEvent(ctx)
			return err
		})
	})

	t.Run("Service_SimL3InterfaceCacheData", func(t *testing.T) {
		testutils.TestServiceMethod(t, func() error {
			_, err := service.SimL3InterfaceCacheData(ctx)
			return err
		})
	})

	t.Run("Service_WlcManagementData", func(t *testing.T) {
		testutils.TestServiceMethod(t, func() error {
			_, err := service.WlcManagementData(ctx)
			return err
		})
	})

	t.Run("Service_Laginfo", func(t *testing.T) {
		testutils.TestServiceMethod(t, func() error {
			_, err := service.Laginfo(ctx)
			return err
		})
	})

	t.Run("Service_MulticastConfig", func(t *testing.T) {
		testutils.TestServiceMethod(t, func() error {
			_, err := service.MulticastConfig(ctx)
			return err
		})
	})

	t.Run("Service_FeatureUsageCfg", func(t *testing.T) {
		testutils.TestServiceMethod(t, func() error {
			_, err := service.FeatureUsageCfg(ctx)
			return err
		})
	})

	t.Run("Service_ThresholdWarnCfg", func(t *testing.T) {
		testutils.TestServiceMethod(t, func() error {
			_, err := service.ThresholdWarnCfg(ctx)
			return err
		})
	})

	t.Run("Service_ApLocRangingCfg", func(t *testing.T) {
		testutils.TestServiceMethod(t, func() error {
			_, err := service.ApLocRangingCfg(ctx)
			return err
		})
	})

	t.Run("Service_GeolocationCfg", func(t *testing.T) {
		testutils.TestServiceMethod(t, func() error {
			_, err := service.GeolocationCfg(ctx)
			return err
		})
	})
}
