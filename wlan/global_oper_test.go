// Package wlan provides WLAN global operational data test functionality for the Cisco Wireless Network Controller API.
package wlan

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil"
)

// =============================================================================
// 1. UNIT TESTS (Structure/Type Validation & JSON Serialization/Deserialization)
// =============================================================================

// WlanGlobalOperTestDataCollector holds test data for WLAN global operation functions
type WlanGlobalOperTestDataCollector struct {
	Data map[string]interface{} `json:"wlan_global_oper_test_data"`
}

var wlanGlobalOperTestDataCollector = WlanGlobalOperTestDataCollector{
	Data: make(map[string]interface{}),
}

func runWlanGlobalOperTestAndCollectData(t *testing.T, testName string, testFunc func() (interface{}, error)) {
	data, err := testFunc()
	if err != nil {
		t.Logf("%s returned error: %v", testName, err)
		wlanGlobalOperTestDataCollector.Data[testName] = map[string]interface{}{
			"error":   err.Error(),
			"success": false,
		}
	} else {
		t.Logf("%s executed successfully", testName)
		wlanGlobalOperTestDataCollector.Data[testName] = map[string]interface{}{
			"data":    data,
			"success": true,
		}
	}
}

// =============================================================================
// 2. INTEGRATION TESTS (API Endpoint Testing with Live Data Validation)
// =============================================================================

// TestWlanGlobalOperFunctions tests all WLAN global operation functions with real WNC data collection
func TestWlanGlobalOperFunctions(t *testing.T) {
	client := testutil.CreateTestClientFromEnv(t)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	t.Run("GetWlanGlobalOper", func(t *testing.T) {
		runWlanGlobalOperTestAndCollectData(t, "GetWlanGlobalOper", func() (interface{}, error) {
			return GetWlanGlobalOper(client, ctx)
		})
	})

	t.Run("GetWlanGlobalOperWlanInfo", func(t *testing.T) {
		runWlanGlobalOperTestAndCollectData(t, "GetWlanGlobalOperWlanInfo", func() (interface{}, error) {
			return GetWlanGlobalOperWlanInfo(client, ctx)
		})
	})

	// Save collected test data to file
	if len(wlanGlobalOperTestDataCollector.Data) > 0 {
		if err := testutil.SaveTestDataToFile("wlan_global_oper_test_data_collected.json", wlanGlobalOperTestDataCollector.Data); err != nil {
			t.Logf("Warning: Could not save test data: %v", err)
		} else {
			t.Logf("Test data saved to %s/wlan_global_oper_test_data_collected.json", testutil.TestDataDir)
		}
	}
}

// TestWlanGlobalOperDataStructures tests the basic structure of WLAN global operational data types
func TestWlanGlobalOperDataStructures(t *testing.T) {
	// Sample WLAN global operational data based on real WNC response structure
	sampleJSON := `{
		"Cisco-IOS-XE-wireless-wlan-global-oper:wlan-global-oper-data": {
			"wlan-info": [
				{
					"wlan-profile": "guest-wlan",
					"curr-clients-count": 25,
					"per-wlan-max-client-syslog": true
				    },
				{
					"wlan-profile": "corporate-wlan",
					"curr-clients-count": 150,
					"per-wlan-max-client-syslog": false
				    },
				{
					"wlan-profile": "iot-wlan",
					"curr-clients-count": 8,
					"per-wlan-max-client-syslog": true
				    }
			]
		    }
	    }`

	// Test unmarshaling into WlanGlobalOperResponse
	var wlanGlobalOper WlanGlobalOperResponse
	err := json.Unmarshal([]byte(sampleJSON), &wlanGlobalOper)
	if err != nil {
		t.Fatalf("Failed to unmarshal WlanGlobalOperResponse: %v", err)
	}

	// Test that data was properly unmarshaled
	wlanInfo := wlanGlobalOper.CiscoIOSXEWirelessWlanGlobalOperData.WlanInfo
	if len(wlanInfo) == 0 {
		t.Error("Expected at least one WLAN info entry")
	}

	// Validate first WLAN info entry
	if len(wlanInfo) > 0 {
		info := wlanInfo[0]
		if info.WlanProfile != "guest-wlan" {
			t.Errorf("Expected WLAN profile 'guest-wlan', got '%s'", info.WlanProfile)
		}

		if info.CurrClientsCount != 25 {
			t.Errorf("Expected current clients count 25, got %d", info.CurrClientsCount)
		}

		if !info.PerWlanMaxClientSyslog {
			t.Error("Expected per-WLAN max client syslog to be true for guest-wlan")
		}
	}

	// Validate second WLAN info entry
	if len(wlanInfo) > 1 {
		info := wlanInfo[1]
		if info.WlanProfile != "corporate-wlan" {
			t.Errorf("Expected WLAN profile 'corporate-wlan', got '%s'", info.WlanProfile)
		}

		if info.CurrClientsCount != 150 {
			t.Errorf("Expected current clients count 150, got %d", info.CurrClientsCount)
		}

		if info.PerWlanMaxClientSyslog {
			t.Error("Expected per-WLAN max client syslog to be false for corporate-wlan")
		}
	}

	// Validate third WLAN info entry
	if len(wlanInfo) > 2 {
		info := wlanInfo[2]
		if info.WlanProfile != "iot-wlan" {
			t.Errorf("Expected WLAN profile 'iot-wlan', got '%s'", info.WlanProfile)
		}

		if info.CurrClientsCount != 8 {
			t.Errorf("Expected current clients count 8, got %d", info.CurrClientsCount)
		}
	}

	// Test WlanGlobalOperWlanInfoResponse structure separately
	sampleWlanInfoJSON := `{
		"Cisco-IOS-XE-wireless-wlan-global-oper:wlan-info": [
			{
				"wlan-profile": "test-wlan",
				"curr-clients-count": 42,
				"per-wlan-max-client-syslog": false
			    }
		]
	    }`

	var wlanInfoResp WlanGlobalOperWlanInfoResponse
	err = json.Unmarshal([]byte(sampleWlanInfoJSON), &wlanInfoResp)
	if err != nil {
		t.Fatalf("Failed to unmarshal WlanGlobalOperWlanInfoResponse: %v", err)
	}

	if len(wlanInfoResp.WlanInfo) == 0 {
		t.Error("Expected at least one WLAN info entry in info response")
	} else {
		info := wlanInfoResp.WlanInfo[0]
		if info.WlanProfile != "test-wlan" {
			t.Errorf("Expected WLAN profile 'test-wlan', got '%s'", info.WlanProfile)
		}

		if info.CurrClientsCount != 42 {
			t.Errorf("Expected current clients count 42, got %d", info.CurrClientsCount)
		}
	}

	_, err = json.Marshal(wlanGlobalOper)
	if err != nil {
		t.Errorf("Failed to marshal WlanGlobalOperResponse back to JSON: %v", err)
	}

	_, err = json.Marshal(wlanInfoResp)
	if err != nil {
		t.Errorf("Failed to marshal WlanGlobalOperWlanInfoResponse back to JSON: %v", err)
	}
}
