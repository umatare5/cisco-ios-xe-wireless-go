// Package afc provides Automated Frequency Coordination operational data test functionality for the Cisco Wireless Network Controller API.
package afc

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil"
)

// =============================================================================
// 1. UNIT TESTS (Structure/Type Validation & JSON Serialization/Deserialization)
// =============================================================================

// TestAfcOperEndpoints tests endpoint constants validation
func TestAfcOperEndpoints(t *testing.T) {
	tests := []struct {
		name     string
		endpoint string
		expected string
	}{
		{
			name:     "Validate_AfcOperBasePath",
			endpoint: AfcOperBasePath,
			expected: "/restconf/data/Cisco-IOS-XE-wireless-afc-oper:afc-oper-data",
		},
		{
			name:     "Validate_AfcOperEndpoint",
			endpoint: "/restconf/data/Cisco-IOS-XE-wireless-afc-oper:afc-oper-data",
			expected: "/restconf/data/Cisco-IOS-XE-wireless-afc-oper:afc-oper-data",
		},
		{
			name:     "Validate_AfcEwlcAfcApRespEndpoint",
			endpoint: "/restconf/data/Cisco-IOS-XE-wireless-afc-oper:afc-oper-data/ewlc-afc-ap-resp",
			expected: "/restconf/data/Cisco-IOS-XE-wireless-afc-oper:afc-oper-data/ewlc-afc-ap-resp",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.endpoint != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, tt.endpoint)
			}
		})
	}
}

// TestAfcOperDataStructures tests the basic structure of AFC operational data types
func TestAfcOperDataStructures(t *testing.T) {
	tests := []struct {
		name     string
		jsonData string
		dataType interface{}
	}{
		{
			name: "AfcOperResponse",
			jsonData: `{
				"Cisco-IOS-XE-wireless-afc-oper:afc-oper-data": {
					"ewlc-afc-ap-resp": [
						{
							"ap-mac": "aa:bb:cc:dd:ee:ff",
							"resp-data": {
								"request-id": "req-12345",
								"ruleset-id": "rule-67890",
								"resp-code": {
									"code": 0,
									"description": "Success",
									"supplemental-info": "AFC response processed successfully"
								},
								"band20": {
									"global-oper-class": 115
								},
								"band40": {
									"global-oper-class": 116
								},
								"band80": {
									"global-oper-class": 125
								},
								"band160": {
									"global-oper-class": 126
								},
								"band80plus": {
									"global-oper-class": 127
								},
								"expire-time": "2024-12-31T23:59:59.000Z",
								"resp-rcvd-timestamp": "2024-01-01T12:00:00.000Z"
							},
							"slot": 0
						}
					]
				}
			}`,
			dataType: &AfcOperResponse{},
		},
		{
			name: "AfcOperEwlcAfcApRespResponse",
			jsonData: `{
				"Cisco-IOS-XE-wireless-afc-oper:ewlc-afc-ap-resp": [
					{
						"ap-mac": "aa:bb:cc:dd:ee:ff",
						"resp-data": {
							"request-id": "req-12345",
							"ruleset-id": "rule-67890",
							"resp-code": {
								"code": 0,
								"description": "Success",
								"supplemental-info": "AFC response processed successfully"
							},
							"band20": {
								"global-oper-class": 115
							},
							"band40": {
								"global-oper-class": 116
							},
							"band80": {
								"global-oper-class": 125
							},
							"band160": {
								"global-oper-class": 126
							},
							"band80plus": {
								"global-oper-class": 127
							},
							"expire-time": "2024-12-31T23:59:59.000Z",
							"resp-rcvd-timestamp": "2024-01-01T12:00:00.000Z"
						},
						"slot": 0
					}
				]
			}`,
			dataType: &AfcOperEwlcAfcApRespResponse{},
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

// Currently no table-driven tests specific to AFC operations

// =============================================================================
// 3. FAIL-FAST ERROR DETECTION TESTS
// =============================================================================

// Currently no fail-fast error detection tests specific to AFC operations

// =============================================================================
// 4. INTEGRATION TESTS (API Communication & Full Workflow Tests)
// =============================================================================

// TestAfcOperMethods tests AFC operational methods with integration
func TestAfcOperMethods(t *testing.T) {
	client := testutil.CreateTestClientFromEnv(t)

	ctx := context.Background()

	// Create a comprehensive test data collection
	collector := testutil.NewTestDataCollector()
	endpointMapping := map[string]string{
		"AfcOperEndpoint":          "/restconf/data/Cisco-IOS-XE-wireless-afc-oper:afc-oper-data",
		"AfcEwlcAfcApRespEndpoint": "/restconf/data/Cisco-IOS-XE-wireless-afc-oper:afc-oper-data/ewlc-afc-ap-resp",
	}

	t.Run("GetAfcOper", func(t *testing.T) {
		result, err := GetAfcOper(client, ctx)
		testutil.CollectTestResult(collector, "GetAfcOper", endpointMapping["AfcOperEndpoint"], result, err)
		if err != nil {
			t.Logf("GetAfcOper failed: %v", err)
		}
	})

	t.Run("GetAfcEwlcAfcApResp", func(t *testing.T) {
		result, err := GetAfcEwlcAfcApResp(client, ctx)
		testutil.CollectTestResult(collector, "GetAfcEwlcAfcApResp", endpointMapping["AfcEwlcAfcApRespEndpoint"], result, err)
		if err != nil {
			t.Logf("GetAfcEwlcAfcApResp failed: %v", err)
		}
	})

	// Save collected test data to JSON file
	testutil.SaveCollectedTestData(t, collector, "afc_oper_test_data_collected.json")
}

// =============================================================================
// 5. OTHER TESTS
// =============================================================================

// Currently no other tests specific to AFC operations
