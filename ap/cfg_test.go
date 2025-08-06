// Package ap provides access point configuration test functionality for the Cisco Wireless Network Controller API.
package ap

import (
	"context"
	"testing"

	wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
	testutils "github.com/umatare5/cisco-ios-xe-wireless-go/internal/tests"
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
	client := testutils.CreateTestClientFromEnv(t)

	ctx, cancel := testutils.CreateDefaultTestContext()
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
		if err := testutils.SaveTestDataToFile("ap_cfg_test_data_collected.json", collector.Data); err != nil {
			t.Logf("Warning: Could not save test data: %v", err)
		} else {
			t.Logf("AP configuration test data saved to %s/ap_cfg_test_data_collected.json", testutils.TestDataDir)
		}
	}
}

// TestApCfgDataStructures tests the basic structure of AP configuration data types
func TestApCfgDataStructures(t *testing.T) {
	testCases := []testutils.JSONTestCase{
		{
			Name: "ApCfgResponse",
			JSONData: `{
				"Cisco-IOS-XE-wireless-ap-cfg:ap-cfg-data": {
					"tag-source-priority-configs": {
						"tag-source-priority-config": [
							{
								"priority": 1,
								"tag-src": "local"
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
			Target:     &ApCfgResponse{},
			TypeName:   "ApCfgResponse",
			ShouldFail: false,
		},
		{
			Name: "ApCfgTagSourcePriorityConfigsResponse",
			JSONData: `{
				"Cisco-IOS-XE-wireless-ap-cfg:tag-source-priority-configs": {
					"tag-source-priority-config": [
						{
							"priority": 2,
							"tag-src": "radius"
						}
					]
				}
			}`,
			Target:     &ApCfgTagSourcePriorityConfigsResponse{},
			TypeName:   "ApCfgTagSourcePriorityConfigsResponse",
			ShouldFail: false,
		},
		{
			Name: "ApCfgApTagsResponse",
			JSONData: `{
				"Cisco-IOS-XE-wireless-ap-cfg:ap-tags": {
					"ap-tag": [
						{
							"ap-mac": "11:22:33:44:55:66",
							"policy-tag": "guest-policy",
							"site-tag": "site-02"
						}
					]
				}
			}`,
			Target:     &ApCfgApTagsResponse{},
			TypeName:   "ApCfgApTagsResponse",
			ShouldFail: false,
		},
	}

	testutils.RunJSONTests(t, testCases)

	// Additional field validation for successfully unmarshaled structures
	t.Run("ApCfgResponseFieldValidation", func(t *testing.T) {
		var response ApCfgResponse
		testutils.TestJSONUnmarshal(t, testCases[0].JSONData, &response, "ApCfgResponse")

		testutils.ValidateJSONStructFields(t, "ApCfgResponse", func() error {
			if len(response.CiscoIOSXEWirelessApCfgApCfgData.TagSourcePriorityConfigs.TagSourcePriorityConfig) != 1 {
				t.Errorf("Expected 1 tag source priority config, got %d",
					len(response.CiscoIOSXEWirelessApCfgApCfgData.TagSourcePriorityConfigs.TagSourcePriorityConfig))
			}
			if response.CiscoIOSXEWirelessApCfgApCfgData.TagSourcePriorityConfigs.TagSourcePriorityConfig[0].Priority != 1 {
				t.Errorf("Expected priority 1, got %d",
					response.CiscoIOSXEWirelessApCfgApCfgData.TagSourcePriorityConfigs.TagSourcePriorityConfig[0].Priority)
			}
			return nil
		})
	})

	t.Run("ApCfgTagSourcePriorityConfigsResponseFieldValidation", func(t *testing.T) {
		var response ApCfgTagSourcePriorityConfigsResponse
		testutils.TestJSONUnmarshal(t, testCases[1].JSONData, &response, "ApCfgTagSourcePriorityConfigsResponse")

		testutils.ValidateJSONStructFields(t, "ApCfgTagSourcePriorityConfigsResponse", func() error {
			if len(response.TagSourcePriorityConfigs.TagSourcePriorityConfig) != 1 {
				t.Errorf("Expected 1 tag source priority config, got %d", len(response.TagSourcePriorityConfigs.TagSourcePriorityConfig))
			}
			return nil
		})
	})

	t.Run("ApCfgApTagsResponseFieldValidation", func(t *testing.T) {
		var response ApCfgApTagsResponse
		testutils.TestJSONUnmarshal(t, testCases[2].JSONData, &response, "ApCfgApTagsResponse")

		testutils.ValidateJSONStructFields(t, "ApCfgApTagsResponse", func() error {
			if len(response.ApTags.ApTag) != 1 {
				t.Errorf("Expected 1 AP tag, got %d", len(response.ApTags.ApTag))
			}
			return nil
		})
	})
}

// =============================================================================
// 3. ERROR HANDLING TESTS
// =============================================================================

// TestApCfgErrorHandling tests error handling for all configuration functions
func TestApCfgErrorHandling(t *testing.T) {
	ctx := context.Background()

	testCases := []struct {
		name string
		fn   func() (interface{}, error)
	}{
		{"GetApCfg", func() (interface{}, error) { return GetApCfg(nil, ctx) }},
		{"GetTagSourcePriorityConfigs", func() (interface{}, error) { return GetTagSourcePriorityConfigs(nil, ctx) }},
		{"GetApTagSourcePriorityConfigs", func() (interface{}, error) { return GetApTagSourcePriorityConfigs(nil, ctx) }},
		{"GetApApTags", func() (interface{}, error) { return GetApApTags(nil, ctx) }},
	}

	for _, tc := range testCases {
		t.Run(tc.name+"WithNilClient", func(t *testing.T) {
			_, err := tc.fn()
			if err == nil {
				t.Errorf("Expected error with nil client, got nil")
			}
			// Accept either error message format for consistency
			errorMsg := err.Error()
			if errorMsg != "invalid client configuration: client cannot be nil" {
				t.Errorf("Expected 'client is nil' or 'invalid client configuration' error, got: %v", err)
			}
		})
	}
}

// =============================================================================
// 4. CONTEXT HANDLING TESTS
// =============================================================================

// TestApCfgContextHandling tests context handling for all configuration functions
func TestApCfgContextHandling(t *testing.T) {
	testCases := []struct {
		name string
		fn   func(context.Context, *wnc.Client) error
	}{
		{"GetApCfg", func(ctx context.Context, client *wnc.Client) error { _, err := GetApCfg(client, ctx); return err }},
		{"GetTagSourcePriorityConfigs", func(ctx context.Context, client *wnc.Client) error {
			_, err := GetTagSourcePriorityConfigs(client, ctx)
			return err
		}},
		{"GetApTagSourcePriorityConfigs", func(ctx context.Context, client *wnc.Client) error {
			_, err := GetApTagSourcePriorityConfigs(client, ctx)
			return err
		}},
		{"GetApApTags", func(ctx context.Context, client *wnc.Client) error { _, err := GetApApTags(client, ctx); return err }},
	}

	for _, tc := range testCases {
		t.Run(tc.name+"ContextHandling", func(t *testing.T) {
			testutils.TestContextHandling(t, tc.fn)
		})
	}
}

// =============================================================================
// 5. ENDPOINT VALIDATION TESTS
// =============================================================================

// TestApCfgEndpoints validates access point configuration endpoint constants
func TestApCfgEndpoints(t *testing.T) {
	// Test base path validation
	t.Run("Validate_ApCfgBasePath", func(t *testing.T) {
		expectedBasePath := "Cisco-IOS-XE-wireless-ap-cfg:ap-cfg-data"
		if ApCfgBasePath != expectedBasePath {
			t.Errorf("ApCfgBasePath mismatch: expected %s, got %s", expectedBasePath, ApCfgBasePath)
		}
	})

	// Test endpoint validation
	t.Run("Validate_ApCfgEndpoint", func(t *testing.T) {
		if ApCfgEndpoint != ApCfgBasePath {
			t.Errorf("ApCfgEndpoint should equal ApCfgBasePath: expected %s, got %s", ApCfgBasePath, ApCfgEndpoint)
		}
	})

	// Test tag source priority configs endpoint validation
	t.Run("Validate_TagSourcePriorityConfigsEndpoint", func(t *testing.T) {
		expectedEndpoint := ApCfgBasePath + "/tag-source-priority-configs"
		if TagSourcePriorityConfigsEndpoint != expectedEndpoint {
			t.Errorf("TagSourcePriorityConfigsEndpoint mismatch: expected %s, got %s", expectedEndpoint, TagSourcePriorityConfigsEndpoint)
		}
	})

	// Test AP tags endpoint validation
	t.Run("Validate_ApTagsEndpoint", func(t *testing.T) {
		expectedEndpoint := ApCfgBasePath + "/ap-tags"
		if ApTagsEndpoint != expectedEndpoint {
			t.Errorf("ApTagsEndpoint mismatch: expected %s, got %s", expectedEndpoint, ApTagsEndpoint)
		}
	})
}

// =============================================================================
// 6. SERVICE TESTS
// =============================================================================

func TestAPService(t *testing.T) {
	client := testutils.GetTestClient(t)
	if client == nil {
		t.Skip("Skipping service tests: no test client available")
	}

	ctx := context.Background()
	service := NewService(client.CoreClient())

	// Test configuration methods
	t.Run("Service_Cfg", func(t *testing.T) {
		testutils.TestServiceMethod(t, func() error {
			_, err := service.Cfg(ctx)
			return err
		})
	})

	t.Run("Service_TagSourcePriorityConfigs", func(t *testing.T) {
		testutils.TestServiceMethod(t, func() error {
			_, err := service.TagSourcePriorityConfigs(ctx)
			return err
		})
	})

	t.Run("Service_ApTags", func(t *testing.T) {
		testutils.TestServiceMethod(t, func() error {
			_, err := service.ApTags(ctx)
			return err
		})
	})
}
