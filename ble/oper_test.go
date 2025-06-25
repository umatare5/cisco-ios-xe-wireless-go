// Package ble provides Bluetooth Low Energy operational data test functionality for the Cisco Wireless Network Controller API.
package ble

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	wnc "github.com/umatare5/cisco-xe-wireless-restconf-go"
	"github.com/umatare5/cisco-xe-wireless-restconf-go/internal/testutil"
)

// =============================================================================
// 1. UNIT TESTS (Structure/Type Validation & JSON Serialization/Deserialization)
// =============================================================================

// getTestClient creates a test client using environment variables
func getTestClient(t *testing.T) *wnc.Client {
	return testutil.CreateTestClientFromEnv(t)
}

// =============================================================================
// 2. INTEGRATION TESTS (API Endpoint Testing with Live Data Validation)
// =============================================================================

func TestBleLtxOperMethods(t *testing.T) {
	client := getTestClient(t)

	ctx := context.Background()

	// Create a comprehensive test data collection
	testResults := make(map[string]interface{})
	endpointMapping := map[string]string{
		"BleLtxOperEndpoint":      "/restconf/data/Cisco-IOS-XE-wireless-ble-ltx-oper:ble-ltx-oper-data",
		"BleLtxApAntennaEndpoint": "/restconf/data/Cisco-IOS-XE-wireless-ble-ltx-oper:ble-ltx-oper-data/ap-antenna",
		"BleLtxApEndpoint":        "/restconf/data/Cisco-IOS-XE-wireless-ble-ltx-oper:ble-ltx-oper-data/ap",
	}

	t.Run("GetBleLtxOper", func(t *testing.T) {
		result, err := GetBleLtxOper(client, ctx)
		collectBleLtxTestResult(testResults, "GetBleLtxOper", endpointMapping["BleLtxOperEndpoint"], result, err)
		if err != nil {
			t.Logf("GetBleLtxOper failed: %v", err)
		}
	})

	t.Run("GetBleLtxApAntenna", func(t *testing.T) {
		result, err := GetBleLtxApAntenna(client, ctx)
		collectBleLtxTestResult(testResults, "GetBleLtxApAntenna", endpointMapping["BleLtxApAntennaEndpoint"], result, err)
		if err != nil {
			t.Logf("GetBleLtxApAntenna failed: %v", err)
		}
	})

	t.Run("GetBleLtxAp", func(t *testing.T) {
		result, err := GetBleLtxAp(client, ctx)
		collectBleLtxTestResult(testResults, "GetBleLtxAp", endpointMapping["BleLtxApEndpoint"], result, err)
		if err != nil {
			t.Logf("GetBleLtxAp failed: %v", err)
		}
	})

	// Save collected test data to JSON file
	saveBleLtxTestData(t, testResults, "ble_ltx_oper_test_data_collected.json")
}

// collectBleLtxTestResult helper function to collect test results
func collectBleLtxTestResult(testResults map[string]interface{}, methodName, endpoint string, result interface{}, err error) {
	testData := map[string]interface{}{
		"method":    methodName,
		"endpoint":  endpoint,
		"timestamp": time.Now().Format(time.RFC3339),
	}

	if err != nil {
		testData["error"] = err.Error()
		testData["success"] = false
	} else {
		testData["success"] = true
		testData["response"] = result
	}

	testResults[methodName] = testData
}

// saveBleLtxTestData helper function to save test data to JSON file
func saveBleLtxTestData(t *testing.T, testResults map[string]interface{}, filename string) {
	if err := testutil.SaveTestDataToFile(filename, testResults); err != nil {
		t.Logf("Failed to save test data to %s: %v", filename, err)
	} else {
		t.Logf("Test data saved to %s/%s", testutil.TestDataDir, filename)
	}
}

func TestBleLtxOperEndpoints(t *testing.T) {
	// Test endpoint validation
	endpoints := map[string]string{
		"BleLtxOperBasePath":      BleLtxOperBasePath,
		"BleLtxOperEndpoint":      "/restconf/data/Cisco-IOS-XE-wireless-ble-ltx-oper:ble-ltx-oper-data",
		"BleLtxApAntennaEndpoint": "/restconf/data/Cisco-IOS-XE-wireless-ble-ltx-oper:ble-ltx-oper-data/ap-antenna",
		"BleLtxApEndpoint":        "/restconf/data/Cisco-IOS-XE-wireless-ble-ltx-oper:ble-ltx-oper-data/ap",
	}

	for name, endpoint := range endpoints {
		t.Run(fmt.Sprintf("Validate_%s", name), func(t *testing.T) {
			if endpoint == "" {
				t.Errorf(wnc.EmptyEndpointErrorTemplate, name)
			}
			if len(endpoint) < 10 {
				t.Errorf(wnc.ShortEndpointErrorTemplate, name, endpoint)
			}
		})
	}
}

// TestBleLtxOperDataStructures tests the basic structure of BLE LTX operational data types
func TestBleLtxOperDataStructures(t *testing.T) {
	tests := []struct {
		name     string
		jsonData string
		dataType interface{}
	}{
		{
			name: "BleLtxOperResponse",
			jsonData: `{
				"Cisco-IOS-XE-wireless-ble-ltx-oper:ble-ltx-oper-data": {
					"ble-ltx-ap-antenna": [
						{
							"ap-mac": "aa:bb:cc:dd:ee:ff",
							"ble-slot-id": 2,
							"ble-antenna-id": 1,
							"is-ble-antenna-present": true,
							"ble-antenna-pid": "BLE-ANT-2.4G",
							"ble-antenna-gain": 2,
							"ble-antenna-type": "omnidirectional",
							"ble-antenna-mode": "diversity",
							"ble-antenna-diversity": "enabled",
							"ble-antenna-options": "standard"
						}
					],
					"ble-ltx-ap": [
						{
							"ap-mac": "aa:bb:cc:dd:ee:ff",
							"ble-slot-id": 2,
							"ble-admin-state": "enabled",
							"ble-oper-state": "up",
							"ble-beacon-interval": 100,
							"ble-tx-power": 0,
							"ble-status": "active"
						}
					]
				}
			}`,
			dataType: &BleLtxOperResponse{},
		},
		{
			name: "BleLtxApAntennaResponse",
			jsonData: `{
				"Cisco-IOS-XE-wireless-ble-ltx-oper:ble-ltx-ap-antenna": [
					{
						"ap-mac": "aa:bb:cc:dd:ee:ff",
						"ble-slot-id": 2,
						"ble-antenna-id": 1,
						"is-ble-antenna-present": true,
						"ble-antenna-pid": "BLE-ANT-2.4G",
						"ble-antenna-gain": 2,
						"ble-antenna-type": "omnidirectional",
						"ble-antenna-mode": "diversity",
						"ble-antenna-diversity": "enabled",
						"ble-antenna-options": "standard"
					}
				]
			}`,
			dataType: &BleLtxApAntennaResponse{},
		},
		{
			name: "BleLtxApResponse",
			jsonData: `{
				"Cisco-IOS-XE-wireless-ble-ltx-oper:ble-ltx-ap": [
					{
						"ap-mac": "aa:bb:cc:dd:ee:ff",
						"ble-slot-id": 2,
						"ble-admin-state": "enabled",
						"ble-oper-state": "up",
						"ble-beacon-interval": 100,
						"ble-tx-power": 0,
						"ble-status": "active"
					}
				]
			}`,
			dataType: &BleLtxApResponse{},
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
