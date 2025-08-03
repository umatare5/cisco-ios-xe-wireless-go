// Package general provides general configuration test functionality for the Cisco Wireless Network Controller API.
package general

import (
	"encoding/json"
	"testing"

	wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil"
)

// =============================================================================
// 1. UNIT TESTS (Structure/Type Validation & JSON Serialization/Deserialization)
// =============================================================================

// getTestClientCfg creates a test client using environment variables
func getTestClientCfg(t *testing.T) *wnc.Client {
	return testutil.CreateTestClientFromEnv(t)
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

	ctx, cancel := testutil.CreateDefaultTestContext()
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
		testutil.SaveTestDataWithLogging("general_cfg_test_data_collected.json", collector.Data)
	}
}

// TestGeneralConfigurationEndpoints tests that all endpoints are correctly defined
func TestGeneralConfigurationEndpoints(t *testing.T) {
	expectedEndpoints := map[string]string{
		"GeneralCfgEndpoint":              "/restconf/data/Cisco-IOS-XE-wireless-general-cfg:general-cfg-data",
		"MewlcConfigEndpoint":             "/restconf/data/Cisco-IOS-XE-wireless-general-cfg:general-cfg-data/mewlc-config",
		"CacConfigEndpoint":               "/restconf/data/Cisco-IOS-XE-wireless-general-cfg:general-cfg-data/cac-config",
		"MfpEndpoint":                     "/restconf/data/Cisco-IOS-XE-wireless-general-cfg:general-cfg-data/mfp",
		"FipsCfgEndpoint":                 "/restconf/data/Cisco-IOS-XE-wireless-general-cfg:general-cfg-data/fips-cfg",
		"WsaApClientEventEndpoint":        "/restconf/data/Cisco-IOS-XE-wireless-general-cfg:general-cfg-data/wsa-ap-client-event",
		"SimL3InterfaceCacheDataEndpoint": "/restconf/data/Cisco-IOS-XE-wireless-general-cfg:general-cfg-data/sim-l3-interface-cache-data",
		"WlcManagementDataEndpoint":       "/restconf/data/Cisco-IOS-XE-wireless-general-cfg:general-cfg-data/wlc-management-data",
		"LaginfoEndpoint":                 "/restconf/data/Cisco-IOS-XE-wireless-general-cfg:general-cfg-data/laginfo",
		"MulticastConfigEndpoint":         "/restconf/data/Cisco-IOS-XE-wireless-general-cfg:general-cfg-data/multicast-config",
		"FeatureUsageCfgEndpoint":         "/restconf/data/Cisco-IOS-XE-wireless-general-cfg:general-cfg-data/feature-usage-cfg",
		"ThresholdWarnCfgEndpoint":        "/restconf/data/Cisco-IOS-XE-wireless-general-cfg:general-cfg-data/threshold-warn-cfg",
		"ApLocRangingCfgEndpoint":         "/restconf/data/Cisco-IOS-XE-wireless-general-cfg:general-cfg-data/ap-loc-ranging-cfg",
		"GeolocationCfgEndpoint":          "/restconf/data/Cisco-IOS-XE-wireless-general-cfg:general-cfg-data/geolocation-cfg",
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
