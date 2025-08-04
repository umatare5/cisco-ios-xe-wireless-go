// Package flex provides FlexConnect configuration test functionality for the Cisco Wireless Network Controller API.
package flex

import (
	"context"
	"testing"
	"time"

	wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil"
	testutils "github.com/umatare5/cisco-ios-xe-wireless-go/tests/utils"
)

// =============================================================================
// 1. UNIT TESTS (Structure/Type Validation & JSON Serialization/Deserialization)
// =============================================================================

func TestFlexCfgDataStructures(t *testing.T) {
	testCases := []testutils.JSONTestCase{
		{
			Name: "FlexCfgResponse",
			JSONData: `{
				"Cisco-IOS-XE-wireless-flex-cfg:flex-cfg-data": {
					"flex-policy-entries": {
						"flex-policy-entry": [
							{
								"policy-name": "corporate-flex",
								"description": "Corporate FlexConnect policy",
								"if-name-vlan-ids": {
									"if-name-vlan-id": [
										{
											"interface-name": "GigabitEthernet0/0/1",
											"vlan-id": 100
										}
									]
								}
							}
						]
					}
				}
			}`,
			Target:     &FlexCfgResponse{},
			TypeName:   "FlexCfgResponse",
			ShouldFail: false,
		},
		{
			Name: "FlexCfgDataResponse",
			JSONData: `{
				"Cisco-IOS-XE-wireless-flex-cfg:flex-policy-entries": {
					"flex-policy-entry": [
						{
							"policy-name": "guest-flex",
							"description": "Guest FlexConnect policy"
						}
					]
				}
			}`,
			Target:     &FlexCfgDataResponse{},
			TypeName:   "FlexCfgDataResponse",
			ShouldFail: false,
		},
		{
			Name: "FlexPolicyEntry",
			JSONData: `{
				"policy-name": "test-flex-policy",
				"description": "Test FlexConnect policy",
				"if-name-vlan-ids": {
					"if-name-vlan-id": [
						{
							"interface-name": "FastEthernet0/0/1",
							"vlan-id": 50
						}
					]
				}
			}`,
			Target:     &FlexPolicyEntry{},
			TypeName:   "FlexPolicyEntry",
			ShouldFail: false,
		},
	}

	testutils.RunJSONTests(t, testCases)

	// Additional field validation for successfully unmarshaled structures
	t.Run("FlexCfgResponseFieldValidation", func(t *testing.T) {
		var response FlexCfgResponse
		testutils.TestJSONUnmarshal(t, testCases[0].JSONData, &response, "FlexCfgResponse")

		testutils.ValidateJSONStructFields(t, "FlexCfgResponse", func() error {
			if len(response.CiscoIOSXEWirelessFlexCfgData.FlexCfgData.FlexPolicyEntry) != 1 {
				t.Errorf("Expected 1 flex policy entry, got %d",
					len(response.CiscoIOSXEWirelessFlexCfgData.FlexCfgData.FlexPolicyEntry))
			}
			policy := response.CiscoIOSXEWirelessFlexCfgData.FlexCfgData.FlexPolicyEntry[0]
			if policy.PolicyName != "corporate-flex" {
				t.Errorf("Expected policy name 'corporate-flex', got '%s'", policy.PolicyName)
			}
			return nil
		})
	})

	t.Run("FlexCfgDataResponseFieldValidation", func(t *testing.T) {
		var response FlexCfgDataResponse
		testutils.TestJSONUnmarshal(t, testCases[1].JSONData, &response, "FlexCfgDataResponse")

		testutils.ValidateJSONStructFields(t, "FlexCfgDataResponse", func() error {
			if len(response.FlexCfgData.FlexPolicyEntry) != 1 {
				t.Errorf("Expected 1 flex policy entry, got %d", len(response.FlexCfgData.FlexPolicyEntry))
			}
			policy := response.FlexCfgData.FlexPolicyEntry[0]
			if policy.PolicyName != "guest-flex" {
				t.Errorf("Expected policy name 'guest-flex', got '%s'", policy.PolicyName)
			}
			return nil
		})
	})

	t.Run("FlexPolicyEntryFieldValidation", func(t *testing.T) {
		var policy FlexPolicyEntry
		testutils.TestJSONUnmarshal(t, testCases[2].JSONData, &policy, "FlexPolicyEntry")

		testutils.ValidateJSONStructFields(t, "FlexPolicyEntry", func() error {
			if policy.PolicyName != "test-flex-policy" {
				t.Errorf("Expected policy name 'test-flex-policy', got '%s'", policy.PolicyName)
			}
			if policy.Description != "Test FlexConnect policy" {
				t.Errorf("Expected description 'Test FlexConnect policy', got '%s'", policy.Description)
			}
			return nil
		})
	})
}

// =============================================================================
// 2. INTEGRATION TESTS (Actual API Calls to Live Controller)
// =============================================================================

// TestFlexConfigurationFunctions tests all FlexConnect configuration functions with a live controller
func TestFlexConfigurationFunctions(t *testing.T) {
	client := testutils.GetTestClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test GetFlexCfg function
	t.Run("GetFlexCfg", func(t *testing.T) {
		data, err := GetFlexCfg(client, ctx)
		if err != nil {
			t.Fatalf("GetFlexCfg failed: %v", err)
		}

		// Validate basic structure
		if data == nil {
			t.Fatal("GetFlexCfg returned nil data")
		}

		// Save test data for analysis
		if err := testutil.SaveTestDataToFile("flex_cfg_data.json", data); err != nil {
			t.Logf("Warning: Could not save test data: %v", err)
		} else {
			t.Logf("Flex config data saved to test_data/flex_cfg_data.json")
		}

		// Validate endpoint was constructed correctly
		endpoint := FlexCfgEndpoint
		if endpoint == "" {
			t.Error("FlexCfgEndpoint should not be empty")
		}
	})

	// Test GetFlexCfgData function
	t.Run("GetFlexCfgData", func(t *testing.T) {
		data, err := GetFlexCfgData(client, ctx)
		if err != nil {
			t.Fatalf("GetFlexCfgData failed: %v", err)
		}

		// Validate basic structure
		if data == nil {
			t.Fatal("GetFlexCfgData returned nil data")
		}

		// Save test data for analysis
		if err := testutil.SaveTestDataToFile("flex_cfg_policy_entries_data.json", data); err != nil {
			t.Logf("Warning: Could not save test data: %v", err)
		} else {
			t.Logf("Flex policy entries data saved to test_data/flex_cfg_policy_entries_data.json")
		}

		// Validate endpoint was constructed correctly
		endpoint := FlexCfgDataEndpoint
		if endpoint == "" {
			t.Error("FlexCfgDataEndpoint should not be empty")
		}
	})

	// Test error handling with common error patterns
	testutils.RunCommonErrorTests(t, "FlexErrorHandling", []testutils.ErrorTestCase{
		{
			Name: "GetFlexCfgWithNilClient",
			TestFunc: func(client *wnc.Client) error {
				_, err := GetFlexCfg(nil, ctx)
				return err
			},
			ExpectedError: "client is nil",
		},
		{
			Name: "GetFlexCfgDataWithNilClient",
			TestFunc: func(client *wnc.Client) error {
				_, err := GetFlexCfgData(nil, ctx)
				return err
			},
			ExpectedError: "client is nil",
		},
	})

	// Test context handling
	t.Run("ContextHandling", func(t *testing.T) {
		testutils.TestContextHandling(t, func(ctx context.Context, client *wnc.Client) error {
			_, err := GetFlexCfg(client, ctx)
			return err
		})

		testutils.TestContextHandling(t, func(ctx context.Context, client *wnc.Client) error {
			_, err := GetFlexCfgData(client, ctx)
			return err
		})
	})
}

// TestFlexConfigurationEndpoints validates FlexConnect configuration endpoint constants
func TestFlexConfigurationEndpoints(t *testing.T) {
	// Test base path validation
	t.Run("Validate_FlexCfgBasePath", func(t *testing.T) {
		expectedBasePath := "/restconf/data/Cisco-IOS-XE-wireless-flex-cfg:flex-cfg-data"
		if FlexCfgBasePath != expectedBasePath {
			t.Errorf("FlexCfgBasePath mismatch: expected %s, got %s", expectedBasePath, FlexCfgBasePath)
		}
	})

	// Test endpoint validation
	t.Run("Validate_FlexCfgEndpoint", func(t *testing.T) {
		if FlexCfgEndpoint != FlexCfgBasePath {
			t.Errorf("FlexCfgEndpoint should equal FlexCfgBasePath: expected %s, got %s", FlexCfgBasePath, FlexCfgEndpoint)
		}
	})

	// Test policy entries endpoint validation
	t.Run("Validate_FlexCfgDataEndpoint", func(t *testing.T) {
		expectedEndpoint := FlexCfgBasePath + "/flex-policy-entries"
		if FlexCfgDataEndpoint != expectedEndpoint {
			t.Errorf("FlexCfgDataEndpoint mismatch: expected %s, got %s", expectedEndpoint, FlexCfgDataEndpoint)
		}
	})
}
