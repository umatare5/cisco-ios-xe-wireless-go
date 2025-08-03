// Package client provides client operational data test functionality for the Cisco Wireless Network Controller API.
package client

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil"
)

// ClientOperTestDataCollector holds test data for client operation functions
type ClientOperTestDataCollector struct {
	Data map[string]interface{} `json:"client_oper_test_data"`
}

func newClientOperTestDataCollector() *ClientOperTestDataCollector {
	return &ClientOperTestDataCollector{
		Data: make(map[string]interface{}),
	}
}

func (collector *ClientOperTestDataCollector) runTestAndCollectData(t *testing.T, testName string, testFunc func() (interface{}, error)) {
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
// 1. UNIT TESTS (Structure/Type Validation & JSON Serialization/Deserialization)
// =============================================================================

// TestClientOperationEndpoints tests that all client operation endpoints are correctly defined
func TestClientOperationEndpoints(t *testing.T) {
	expectedEndpoints := map[string]string{
		"ClientOperEndpoint":        "/restconf/data/Cisco-IOS-XE-wireless-client-oper:client-oper-data",
		"CommonOperDataEndpoint":    "/restconf/data/Cisco-IOS-XE-wireless-client-oper:client-oper-data/common-oper-data",
		"Dot11OperDataEndpoint":     "/restconf/data/Cisco-IOS-XE-wireless-client-oper:client-oper-data/dot11-oper-data",
		"MobilityOperDataEndpoint":  "/restconf/data/Cisco-IOS-XE-wireless-client-oper:client-oper-data/mobility-oper-data",
		"MmIfClientStatsEndpoint":   "/restconf/data/Cisco-IOS-XE-wireless-client-oper:client-oper-data/mm-if-client-stats",
		"MmIfClientHistoryEndpoint": "/restconf/data/Cisco-IOS-XE-wireless-client-oper:client-oper-data/mm-if-client-history",
		"TrafficStatsEndpoint":      "/restconf/data/Cisco-IOS-XE-wireless-client-oper:client-oper-data/traffic-stats",
		"PolicyDataEndpoint":        "/restconf/data/Cisco-IOS-XE-wireless-client-oper:client-oper-data/policy-data",
		"SisfDbMacEndpoint":         "/restconf/data/Cisco-IOS-XE-wireless-client-oper:client-oper-data/sisf-db-mac",
		"DcInfoEndpoint":            "/restconf/data/Cisco-IOS-XE-wireless-client-oper:client-oper-data/dc-info",
	}

	for name, expected := range expectedEndpoints {
		t.Run(name, func(t *testing.T) {
			switch name {
			case "ClientOperEndpoint":
				if ClientOperEndpoint != expected {
					t.Errorf("Expected %s = %s, got %s", name, expected, ClientOperEndpoint)
				}
			case "CommonOperDataEndpoint":
				if CommonOperDataEndpoint != expected {
					t.Errorf("Expected %s = %s, got %s", name, expected, CommonOperDataEndpoint)
				}
			case "Dot11OperDataEndpoint":
				if Dot11OperDataEndpoint != expected {
					t.Errorf("Expected %s = %s, got %s", name, expected, Dot11OperDataEndpoint)
				}
			case "MobilityOperDataEndpoint":
				if MobilityOperDataEndpoint != expected {
					t.Errorf("Expected %s = %s, got %s", name, expected, MobilityOperDataEndpoint)
				}
			case "MmIfClientStatsEndpoint":
				if MmIfClientStatsEndpoint != expected {
					t.Errorf("Expected %s = %s, got %s", name, expected, MmIfClientStatsEndpoint)
				}
			case "MmIfClientHistoryEndpoint":
				if MmIfClientHistoryEndpoint != expected {
					t.Errorf("Expected %s = %s, got %s", name, expected, MmIfClientHistoryEndpoint)
				}
			case "TrafficStatsEndpoint":
				if TrafficStatsEndpoint != expected {
					t.Errorf("Expected %s = %s, got %s", name, expected, TrafficStatsEndpoint)
				}
			case "PolicyDataEndpoint":
				if PolicyDataEndpoint != expected {
					t.Errorf("Expected %s = %s, got %s", name, expected, PolicyDataEndpoint)
				}
			case "SisfDbMacEndpoint":
				if SisfDbMacEndpoint != expected {
					t.Errorf("Expected %s = %s, got %s", name, expected, SisfDbMacEndpoint)
				}
			case "DcInfoEndpoint":
				if DcInfoEndpoint != expected {
					t.Errorf("Expected %s = %s, got %s", name, expected, DcInfoEndpoint)
				}
			}
		})
	}
}

// TestClientOperDataStructures tests the basic structure of client operational data types
func TestClientOperDataStructures(t *testing.T) {
	tests := []struct {
		name     string
		jsonData string
		dataType interface{}
	}{
		{
			name: "ClientOperCommonOperDataResponse",
			jsonData: `{
				"Cisco-IOS-XE-wireless-client-oper:common-oper-data": [
					{
						"ms-mac-address": "aa:bb:cc:dd:ee:ff",
						"ap-mac-address": "11:22:33:44:55:66",
						"wtp-mac": "11:22:33:44:55:66",
						"station-type": "associated",
						"client-state": "associated",
						"policy-profile": "default-policy",
						"ssid": "TestSSID",
						"vlan-id": 100,
						"ap-name": "AP-Floor1-001",
						"connected-time": 3600,
						"client-username": "testuser"
					}
				]
			}`,
			dataType: &ClientOperCommonOperDataResponse{},
		},
		{
			name: "ClientOperDot11OperDataResponse",
			jsonData: `{
				"Cisco-IOS-XE-wireless-client-oper:dot11-oper-data": [
					{
						"ms-mac-address": "aa:bb:cc:dd:ee:ff",
						"ap-mac-address": "11:22:33:44:55:66",
						"radio-type": "dot11ac",
						"channel": 36,
						"rssi": -45,
						"snr": 35,
						"data-rate": "866.7Mbps",
						"tx-rate": "433.3Mbps",
						"rx-rate": "433.3Mbps"
					}
				]
			}`,
			dataType: &ClientOperDot11OperDataResponse{},
		},
		{
			name: "ClientOperMobilityOperDataResponse",
			jsonData: `{
				"Cisco-IOS-XE-wireless-client-oper:mobility-oper-data": [
					{
						"ms-mac-address": "aa:bb:cc:dd:ee:ff",
						"mobility-state": "local",
						"anchor-controller": "192.168.1.10",
						"foreign-controller": "",
						"mobility-role": "anchor",
						"handoff-count": 0
					}
				]
			}`,
			dataType: &ClientOperMobilityOperDataResponse{},
		},
		{
			name: "ClientOperTrafficStatsResponse",
			jsonData: `{
				"Cisco-IOS-XE-wireless-client-oper:traffic-stats": [
					{
						"ms-mac-address": "aa:bb:cc:dd:ee:ff",
						"bytes-tx": "1024000",
						"bytes-rx": "2048000",
						"pkts-tx": "1000",
						"pkts-rx": "2000",
						"data-retries": "5",
						"rts-retries": "2",
						"policy-errs": "0",
						"duplicate-rcv": "0",
						"decrypt-failed": "0",
						"mic-mismatch": "0",
						"mic-missing": "0",
						"most-recent-rssi": -45,
						"most-recent-snr": 55,
						"tx-excessive-retries": "0",
						"tx-retries": "0",
						"power-save-state": 0,
						"current-rate": "m7",
						"speed": 72,
						"spatial-stream": 1,
						"client-active": true,
						"glan-stats-update-timestamp": "1970-01-01T00:00:00+00:00",
						"glan-idle-update-timestamp": "1970-01-01T00:00:00+00:00",
						"rx-group-counter": "0",
						"tx-total-drops": "0"
					}
				]
			}`,
			dataType: &ClientOperTrafficStatsResponse{},
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
// 2. TABLE-DRIVEN TEST PATTERNS
// =============================================================================

// Currently no table-driven tests specific to client operations

// =============================================================================
// 3. FAIL-FAST ERROR DETECTION TESTS
// =============================================================================

// Currently no fail-fast error detection tests specific to client operations

// =============================================================================
// 4. INTEGRATION TESTS (API Communication & Full Workflow Tests)
// =============================================================================

// TestClientOperationFunctions tests all client operation functions with real WNC data collection
func TestClientOperationFunctions(t *testing.T) {
	client := testutil.CreateTestClientFromEnv(t)

	ctx, cancel := context.WithTimeout(context.Background(), testutil.DefaultTestTimeout)
	defer cancel()

	collector := newClientOperTestDataCollector()

	t.Run("GetClientOper", func(t *testing.T) {
		collector.runTestAndCollectData(t, "GetClientOper", func() (interface{}, error) {
			return GetClientOper(client, ctx)
		})
	})

	t.Run("GetClientOperCommonOperData", func(t *testing.T) {
		resp, err := GetClientOperCommonOperData(client, ctx)
		if err != nil {
			t.Logf("GetClientOperCommonOperData returned error (may be expected): %v", err)
		} else {
			t.Log("GetClientOperCommonOperData executed successfully")
			if resp != nil {
				t.Logf("Response contains client common operation data")
			}
		}
	})

	t.Run("GetClientOperDot11OperData", func(t *testing.T) {
		resp, err := GetClientOperDot11OperData(client, ctx)
		if err != nil {
			t.Logf("GetClientOperDot11OperData returned error (may be expected): %v", err)
		} else {
			t.Log("GetClientOperDot11OperData executed successfully")
			if resp != nil {
				t.Logf("Response contains client 802.11 operation data")
			}
		}
	})

	t.Run("GetClientOperMobilityOperData", func(t *testing.T) {
		resp, err := GetClientOperMobilityOperData(client, ctx)
		if err != nil {
			t.Logf("GetClientOperMobilityOperData returned error (may be expected): %v", err)
		} else {
			t.Log("GetClientOperMobilityOperData executed successfully")
			if resp != nil {
				t.Logf("Response contains client mobility operation data")
			}
		}
	})

	t.Run("GetClientOperMmIfClientStats", func(t *testing.T) {
		resp, err := GetClientOperMmIfClientStats(client, ctx)
		if err != nil {
			t.Logf("GetClientOperMmIfClientStats returned error (may be expected): %v", err)
		} else {
			t.Log("GetClientOperMmIfClientStats executed successfully")
			if resp != nil {
				t.Logf("Response contains client MM interface stats")
			}
		}
	})

	t.Run("GetClientOperMmIfClientHistory", func(t *testing.T) {
		resp, err := GetClientOperMmIfClientHistory(client, ctx)
		if err != nil {
			t.Logf("GetClientOperMmIfClientHistory returned error (may be expected): %v", err)
		} else {
			t.Log("GetClientOperMmIfClientHistory executed successfully")
			if resp != nil {
				t.Logf("Response contains client MM interface history")
			}
		}
	})

	t.Run("GetClientOperTrafficStats", func(t *testing.T) {
		resp, err := GetClientOperTrafficStats(client, ctx)
		if err != nil {
			t.Logf("GetClientOperTrafficStats returned error (may be expected): %v", err)
		} else {
			t.Log("GetClientOperTrafficStats executed successfully")
			if resp != nil {
				t.Logf("Response contains client traffic stats")
			}
		}
	})

	t.Run("GetClientOperPolicyData", func(t *testing.T) {
		resp, err := GetClientOperPolicyData(client, ctx)
		if err != nil {
			t.Logf("GetClientOperPolicyData returned error (may be expected): %v", err)
		} else {
			t.Log("GetClientOperPolicyData executed successfully")
			if resp != nil {
				t.Logf("Response contains client policy data")
			}
		}
	})

	t.Run("GetClientOperSisfDbMac", func(t *testing.T) {
		resp, err := GetClientOperSisfDbMac(client, ctx)
		if err != nil {
			t.Logf("GetClientOperSisfDbMac returned error (may be expected): %v", err)
		} else {
			t.Log("GetClientOperSisfDbMac executed successfully")
			if resp != nil {
				t.Logf("Response contains client SISF DB MAC data")
			}
		}
	})

	t.Run("GetClientOperDcInfo", func(t *testing.T) {
		resp, err := GetClientOperDcInfo(client, ctx)
		if err != nil {
			t.Logf("GetClientOperDcInfo returned error (may be expected): %v", err)
		} else {
			t.Log("GetClientOperDcInfo executed successfully")
			if resp != nil {
				t.Logf("Response contains client DC info")
			}
		}
	})

	// Save collected test data to file
	if len(collector.Data) > 0 {
		if err := testutil.SaveTestDataToFile("client_oper_test_data_collected.json", collector.Data); err != nil {
			t.Logf("Warning: Could not save test data: %v", err)
		} else {
			t.Logf("Test data saved to %s/client_oper_test_data_collected.json", testutil.TestDataDir)
		}
	}
}

// =============================================================================
// 5. OTHER TESTS
// =============================================================================

// Currently no other tests specific to client operations
