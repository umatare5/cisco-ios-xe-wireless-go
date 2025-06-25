// Package afc provides Automated Frequency Coordination cloud operational data test functionality for the Cisco Wireless Network Controller API.
package afc

import (
	"context"
	"encoding/json"
	"testing"

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

func TestAfcCloudOperMethods(t *testing.T) {
	client := getTestClient(t)

	ctx := context.Background()

	// Create a comprehensive test data collection
	collector := testutil.NewTestDataCollector()
	endpointMapping := map[string]string{
		"AfcCloudOperEndpoint":  "/restconf/data/Cisco-IOS-XE-wireless-afc-cloud-oper:afc-cloud-oper-data",
		"AfcCloudStatsEndpoint": "/restconf/data/Cisco-IOS-XE-wireless-afc-cloud-oper:afc-cloud-oper-data/afc-cloud-stats",
	}

	t.Run("GetAfcCloudOper", func(t *testing.T) {
		result, err := GetAfcCloudOper(client, ctx)
		testutil.CollectTestResult(collector, "GetAfcCloudOper", endpointMapping["AfcCloudOperEndpoint"], result, err)
		if err != nil {
			t.Logf("GetAfcCloudOper failed: %v", err)
		}
	})

	t.Run("GetAfcCloudStats", func(t *testing.T) {
		result, err := GetAfcCloudStats(client, ctx)
		testutil.CollectTestResult(collector, "GetAfcCloudStats", endpointMapping["AfcCloudStatsEndpoint"], result, err)
		if err != nil {
			t.Logf("GetAfcCloudStats failed: %v", err)
		}
	})

	// Save collected test data to JSON file
	testutil.SaveCollectedTestData(t, collector, "afc_cloud_oper_test_data_collected.json")
}

func TestAfcCloudOperEndpoints(t *testing.T) {
	// Test endpoint validation
	endpoints := map[string]string{
		"AfcCloudOperBasePath":  AfcCloudOperBasePath,
		"AfcCloudOperEndpoint":  "/restconf/data/Cisco-IOS-XE-wireless-afc-cloud-oper:afc-cloud-oper-data",
		"AfcCloudStatsEndpoint": "/restconf/data/Cisco-IOS-XE-wireless-afc-cloud-oper:afc-cloud-oper-data/afc-cloud-stats",
	}

	testutil.ValidateEndpoints(t, endpoints)
}

// TestAfcCloudOperDataStructures tests the basic structure of AFC cloud operational data types
func TestAfcCloudOperDataStructures(t *testing.T) {
	tests := []struct {
		name     string
		jsonData string
		dataType interface{}
	}{
		{
			name: "AfcCloudOperResponse",
			jsonData: `{
				"Cisco-IOS-XE-wireless-afc-cloud-oper:afc-cloud-oper-data": {
					"afc-cloud-stats": {
						"num-afc-ap": 5,
						"afc-msg-sent": "123",
						"afc-msg-rcvd": "120",
						"afc-msg-err": "3",
						"afc-msg-pending": 2,
						"last-msg-sent": {
							"request-id": "req-12345",
							"ap-mac": "aa:bb:cc:dd:ee:ff",
							"msg-timestamp": "2024-01-01T12:00:00.000Z"
						},
						"last-msg-rcvd": {
							"request-id": "req-12344",
							"ap-mac": "aa:bb:cc:dd:ee:fe",
							"msg-timestamp": "2024-01-01T12:00:01.000Z"
						},
						"min-msg-rtt": "50ms",
						"max-msg-rtt": "500ms",
						"avg-rtt": "150ms",
						"healthcheck": {
							"hc-timestamp": "2024-01-01T12:05:00.000Z",
							"query-in-progress": false,
							"country-not-supported": false,
							"num-hc-down": 0,
							"hc-error-status": {
								"not-otp-upgraded": false
							}
						},
						"num-6ghz-ap": 3
					}
				}
			}`,
			dataType: &AfcCloudOperResponse{},
		},
		{
			name: "AfcCloudOperAfcCloudStatsResponse",
			jsonData: `{
				"Cisco-IOS-XE-wireless-afc-cloud-stats": {
					"num-afc-ap": 5,
					"afc-msg-sent": "123",
					"afc-msg-rcvd": "120",
					"afc-msg-err": "3",
					"afc-msg-pending": 2,
					"last-msg-sent": {
						"request-id": "req-12345",
						"ap-mac": "aa:bb:cc:dd:ee:ff",
						"msg-timestamp": "2024-01-01T12:00:00.000Z"
					},
					"last-msg-rcvd": {
						"request-id": "req-12344",
						"ap-mac": "aa:bb:cc:dd:ee:fe",
						"msg-timestamp": "2024-01-01T12:00:01.000Z"
					},
					"min-msg-rtt": "50ms",
					"max-msg-rtt": "500ms",
					"avg-rtt": "150ms",
					"healthcheck": {
						"hc-timestamp": "2024-01-01T12:05:00.000Z",
						"query-in-progress": false,
						"country-not-supported": false,
						"num-hc-down": 0,
						"hc-error-status": {
							"not-otp-upgraded": false
						}
					},
					"num-6ghz-ap": 3
				}
			}`,
			dataType: &AfcCloudOperAfcCloudStatsResponse{},
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
