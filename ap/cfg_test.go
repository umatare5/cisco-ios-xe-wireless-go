// Package ap provides access point configuration test functionality for the Cisco Wireless Network Controller API.
package ap

import (
	"encoding/json"
	"testing"

	"github.com/umatare5/cisco-xe-wireless-restconf-go/internal/testutil"
)

// =============================================================================
// 1. UNIT TESTS (Structure/Type Validation & JSON Serialization/Deserialization)
// =============================================================================

// APCfgTestDataCollector holds test data for AP configuration functions
type APCfgTestDataCollector struct {
	Data map[string]interface{} `json:"ap_cfg_test_data"`
}

func newAPCfgTestDataCollector() *APCfgTestDataCollector {
	return &APCfgTestDataCollector{
		Data: make(map[string]interface{}),
	}
}

func (collector *APCfgTestDataCollector) runTestAndCollectData(t *testing.T, testName string, testFunc func() (interface{}, error)) {
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

func TestAPConfigurationFunctions(t *testing.T) {
	client := testutil.CreateTestClientFromEnv(t)

	ctx, cancel := testutil.CreateDefaultTestContext()
	defer cancel()

	collector := newAPCfgTestDataCollector()

	t.Run("GetApCfg", func(t *testing.T) {
		collector.runTestAndCollectData(t, "GetApCfg", func() (interface{}, error) {
			return GetApCfg(client, ctx)
		})
	})

	t.Run("GetTagSourcePriorityConfigs", func(t *testing.T) {
		collector.runTestAndCollectData(t, "GetTagSourcePriorityConfigs", func() (interface{}, error) {
			return GetTagSourcePriorityConfigs(client, ctx)
		})
	})

	t.Run("GetApTagSourcePriorityConfigs", func(t *testing.T) {
		collector.runTestAndCollectData(t, "GetApTagSourcePriorityConfigs", func() (interface{}, error) {
			return GetApTagSourcePriorityConfigs(client, ctx)
		})
	})

	t.Run("GetApApTags", func(t *testing.T) {
		collector.runTestAndCollectData(t, "GetApApTags", func() (interface{}, error) {
			return GetApApTags(client, ctx)
		})
	})

	// Save collected test data to file
	if len(collector.Data) > 0 {
		if err := testutil.SaveTestDataToFile("ap_cfg_test_data_collected.json", collector.Data); err != nil {
			t.Logf("Warning: Could not save test data: %v", err)
		} else {
			t.Logf("AP configuration test data saved to %s/ap_cfg_test_data_collected.json", testutil.TestDataDir)
		}
	}
}

// TestApCfgDataStructures tests the basic structure of AP configuration data types
func TestApCfgDataStructures(t *testing.T) {
	tests := []struct {
		name     string
		jsonData string
		dataType interface{}
	}{
		{
			name: "ApCfgResponse",
			jsonData: `{
				"Cisco-IOS-XE-wireless-ap-cfg:ap-cfg-data": {
					"tag-source-priority-configs": {
						"tag-source-priority-config": [
							{
								"priority": 1,
								"tag-src": "local"
							},
							{
								"priority": 2,
								"tag-src": "radius"
							}
						]
					},
					"ap-tags": {
						"ap-tag": [
							{
								"ap-mac": "aa:bb:cc:dd:ee:ff",
								"policy-tag": "default-policy",
								"site-tag": "site-01",
								"rf-tag": "rf-default"
							}
						]
					}
				}
			}`,
			dataType: &ApCfgResponse{},
		},
		{
			name: "ApCfgTagSourcePriorityConfigsResponse",
			jsonData: `{
				"Cisco-IOS-XE-wireless-ap-cfg:tag-source-priority-configs": {
					"tag-source-priority-config": [
						{
							"priority": 1,
							"tag-src": "local"
						},
						{
							"priority": 2,
							"tag-src": "radius"
						},
						{
							"priority": 3,
							"tag-src": "mac-address"
						}
					]
				}
			}`,
			dataType: &ApCfgTagSourcePriorityConfigsResponse{},
		},
		{
			name: "ApCfgApTagsResponse",
			jsonData: `{
				"Cisco-IOS-XE-wireless-ap-cfg:ap-tags": {
					"ap-tag": [
						{
							"ap-mac": "aa:bb:cc:dd:ee:ff",
							"policy-tag": "default-policy",
							"site-tag": "site-01",
							"rf-tag": "rf-default"
						},
						{
							"ap-mac": "11:22:33:44:55:66",
							"policy-tag": "guest-policy",
							"site-tag": "site-02"
						}
					]
				}
			}`,
			dataType: &ApCfgApTagsResponse{},
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
