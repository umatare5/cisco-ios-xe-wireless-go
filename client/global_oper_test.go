// Package client provides client global operational data test functionality for the Cisco Wireless Network Controller API.
package client

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil"
)

// =============================================================================
// 1. UNIT TESTS (Structure/Type Validation & JSON Serialization/Deserialization)
// =============================================================================

// ClientGlobalOperTestDataCollector holds test data for client global operation functions
type ClientGlobalOperTestDataCollector struct {
	Data map[string]interface{} `json:"client_global_oper_test_data"`
}

func newClientGlobalOperTestDataCollector() *ClientGlobalOperTestDataCollector {
	return &ClientGlobalOperTestDataCollector{
		Data: make(map[string]interface{}),
	}
}

func (collector *ClientGlobalOperTestDataCollector) runTestAndCollectData(t *testing.T, testName string, testFunc func() (interface{}, error)) {
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

// TestClientGlobalOperationFunctions tests all client global operation functions with real WNC data collection
func TestClientGlobalOperationFunctions(t *testing.T) {
	client := testutil.CreateTestClientFromEnv(t)

	ctx, cancel := context.WithTimeout(context.Background(), testutil.DefaultTestTimeout)
	defer cancel()

	collector := newClientGlobalOperTestDataCollector()

	t.Run("GetClientGlobalOper", func(t *testing.T) {
		collector.runTestAndCollectData(t, "GetClientGlobalOper", func() (interface{}, error) {
			return GetClientGlobalOper(client, ctx)
		})
	})

	t.Run("GetClientLiveStats", func(t *testing.T) {
		collector.runTestAndCollectData(t, "GetClientLiveStats", func() (interface{}, error) {
			return GetClientLiveStats(client, ctx)
		})
	})

	t.Run("GetClientGlobalStatsData", func(t *testing.T) {
		collector.runTestAndCollectData(t, "GetClientGlobalStatsData", func() (interface{}, error) {
			return GetClientGlobalStatsData(client, ctx)
		})
	})

	t.Run("GetClientStats", func(t *testing.T) {
		collector.runTestAndCollectData(t, "GetClientStats", func() (interface{}, error) {
			return GetClientStats(client, ctx)
		})
	})

	t.Run("GetClientDot11Stats", func(t *testing.T) {
		collector.runTestAndCollectData(t, "GetClientDot11Stats", func() (interface{}, error) {
			return GetClientDot11Stats(client, ctx)
		})
	})

	t.Run("GetClientLatencyStats", func(t *testing.T) {
		collector.runTestAndCollectData(t, "GetClientLatencyStats", func() (interface{}, error) {
			return GetClientLatencyStats(client, ctx)
		})
	})

	t.Run("GetClientSmWebauthStats", func(t *testing.T) {
		collector.runTestAndCollectData(t, "GetClientSmWebauthStats", func() (interface{}, error) {
			return GetClientSmWebauthStats(client, ctx)
		})
	})

	t.Run("GetClientDot1XGlobalStats", func(t *testing.T) {
		collector.runTestAndCollectData(t, "GetClientDot1XGlobalStats", func() (interface{}, error) {
			return GetClientDot1XGlobalStats(client, ctx)
		})
	})

	t.Run("GetClientExclusionStats", func(t *testing.T) {
		collector.runTestAndCollectData(t, "GetClientExclusionStats", func() (interface{}, error) {
			return GetClientExclusionStats(client, ctx)
		})
	})

	t.Run("GetClientSmDeviceCount", func(t *testing.T) {
		collector.runTestAndCollectData(t, "GetClientSmDeviceCount", func() (interface{}, error) {
			return GetClientSmDeviceCount(client, ctx)
		})
	})

	t.Run("GetClientTofStats", func(t *testing.T) {
		collector.runTestAndCollectData(t, "GetClientTofStats", func() (interface{}, error) {
			return GetClientTofStats(client, ctx)
		})
	})

	// Save collected test data to file
	if len(collector.Data) > 0 {
		if err := testutil.SaveTestDataToFile("client_global_oper_test_data_collected.json", collector.Data); err != nil {
			t.Logf("Warning: Could not save test data: %v", err)
		} else {
			t.Logf("Test data saved to %s/client_global_oper_test_data_collected.json", testutil.TestDataDir)
		}
	}
}

// TestClientGlobalOperDataStructures tests the basic structure of client global operational data types
func TestClientGlobalOperDataStructures(t *testing.T) {
	tests := []struct {
		name     string
		jsonData string
		dataType interface{}
	}{
		{
			name: "ClientGlobalOperResponse",
			jsonData: `{
				"Cisco-IOS-XE-wireless-client-global-oper:client-global-oper-data": {
					"client-live-stats": {
						"assoc-count": 10,
						"auth-count": 8,
						"mobility-count": 2,
						"iplearn-count": 7,
						"webauth-pending-count": 1,
						"run-count": 6,
						"delete-count": 0
					},
					"client-global-stats-data": {
						"current-clients": 25,
						"excluded-clients": 2,
						"reauthentication-clients": 1,
						"assisted-roaming-stats": {
							"ar-11k-attempts": 5,
							"ar-11k-success": 4,
							"ar-11v-attempts": 3,
							"ar-11v-success": 2
						}
					}
				}
			}`,
			dataType: &ClientGlobalOperResponse{},
		},
		{
			name: "ClientLiveStatsResponse",
			jsonData: `{
				"Cisco-IOS-XE-wireless-client-global-oper:client-live-stats": {
					"assoc-count": 15,
					"auth-count": 12,
					"mobility-count": 3,
					"iplearn-count": 10,
					"webauth-pending-count": 2,
					"run-count": 9,
					"delete-count": 1
				}
			}`,
			dataType: &ClientLiveStatsResponse{},
		},
		{
			name: "ClientGlobalStatsDataResponse",
			jsonData: `{
				"Cisco-IOS-XE-wireless-client-global-oper:client-global-stats-data": {
					"current-clients": 50,
					"excluded-clients": 5,
					"reauthentication-clients": 3,
					"assisted-roaming-stats": {
						"ar-11k-attempts": 10,
						"ar-11k-success": 8,
						"ar-11v-attempts": 6,
						"ar-11v-success": 4,
						"ar-ota-attempts": 2,
						"ar-ota-success": 1
					}
				}
			}`,
			dataType: &ClientGlobalStatsDataResponse{},
		},
		{
			name: "ClientStatsResponse",
			jsonData: `{
				"Cisco-IOS-XE-wireless-client-global-oper:client-stats": {
					"no-of-clients": 100,
					"client-summary": [
						{
							"wtp-mac": "aa:bb:cc:dd:ee:01",
							"client-mac": "11:22:33:44:55:66",
							"ap-name": "ap-01",
							"client-state": "associated",
							"client-username": "user1"
						}
					]
				}
			}`,
			dataType: &ClientStatsResponse{},
		},
		{
			name: "ClientDot11StatsResponse",
			jsonData: `{
				"Cisco-IOS-XE-wireless-client-global-oper:client-dot11-stats": {
					"tx-bytes": 1024000,
					"rx-bytes": 2048000,
					"tx-packets": 1000,
					"rx-packets": 2000,
					"tx-data-packets": 950,
					"rx-data-packets": 1900
				}
			}`,
			dataType: &ClientDot11StatsResponse{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := json.Unmarshal([]byte(tt.jsonData), tt.dataType)
			if err != nil {
				t.Errorf("Failed to unmarshal %s: %v", tt.name, err)
			}
		})
	}
}
