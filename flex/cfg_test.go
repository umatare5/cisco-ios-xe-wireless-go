// Package flex provides FlexConnect configuration test functionality for the Cisco Wireless Network Controller API.
package flex

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

func TestFlexCfgDataStructures(t *testing.T) {
	// Test FlexCfgResponse structure
	t.Run("FlexCfgResponse", func(t *testing.T) {
		sampleJSON := `{
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
									},
									{
										"interface-name": "GigabitEthernet0/0/2",
										"vlan-id": 200
									}
								]
							}
						}
					]
				}
			}
		}`

		var response FlexCfgResponse
		err := json.Unmarshal([]byte(sampleJSON), &response)
		if err != nil {
			t.Fatalf("Failed to unmarshal FlexCfgResponse: %v", err)
		}

		if len(response.CiscoIOSXEWirelessFlexCfgData.FlexCfgData.FlexPolicyEntry) != 1 {
			t.Errorf("Expected 1 flex policy entry, got %d",
				len(response.CiscoIOSXEWirelessFlexCfgData.FlexCfgData.FlexPolicyEntry))
		}

		policy := response.CiscoIOSXEWirelessFlexCfgData.FlexCfgData.FlexPolicyEntry[0]
		if policy.PolicyName != "corporate-flex" {
			t.Errorf("Expected policy name 'corporate-flex', got '%s'", policy.PolicyName)
		}

		if policy.Description != "Corporate FlexConnect policy" {
			t.Errorf("Expected description 'Corporate FlexConnect policy', got '%s'", policy.Description)
		}

		if policy.IfNameVlanIds == nil || len(policy.IfNameVlanIds.IfNameVlanId) != 2 {
			t.Error("Expected 2 interface-vlan mappings")
		} else {
			if policy.IfNameVlanIds.IfNameVlanId[0].InterfaceName != "GigabitEthernet0/0/1" {
				t.Errorf("Expected interface name 'GigabitEthernet0/0/1', got '%s'",
					policy.IfNameVlanIds.IfNameVlanId[0].InterfaceName)
			}
			if policy.IfNameVlanIds.IfNameVlanId[0].VlanID != 100 {
				t.Errorf("Expected VLAN ID 100, got %d", policy.IfNameVlanIds.IfNameVlanId[0].VlanID)
			}
		}
	})

	// Test FlexCfgDataResponse structure
	t.Run("FlexCfgDataResponse", func(t *testing.T) {
		sampleJSON := `{
			"Cisco-IOS-XE-wireless-flex-cfg:flex-policy-entries": {
				"flex-policy-entry": [
					{
						"policy-name": "guest-flex",
						"description": "Guest FlexConnect policy"
					}
				]
			}
		}`

		var response FlexCfgDataResponse
		err := json.Unmarshal([]byte(sampleJSON), &response)
		if err != nil {
			t.Fatalf("Failed to unmarshal FlexCfgDataResponse: %v", err)
		}

		if len(response.FlexCfgData.FlexPolicyEntry) != 1 {
			t.Errorf("Expected 1 flex policy entry, got %d", len(response.FlexCfgData.FlexPolicyEntry))
		}

		policy := response.FlexCfgData.FlexPolicyEntry[0]
		if policy.PolicyName != "guest-flex" {
			t.Errorf("Expected policy name 'guest-flex', got '%s'", policy.PolicyName)
		}

		if policy.Description != "Guest FlexConnect policy" {
			t.Errorf("Expected description 'Guest FlexConnect policy', got '%s'", policy.Description)
		}
	})

	// Test FlexPolicyEntry structure
	t.Run("FlexPolicyEntry", func(t *testing.T) {
		sampleJSON := `{
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
		}`

		var policy FlexPolicyEntry
		err := json.Unmarshal([]byte(sampleJSON), &policy)
		if err != nil {
			t.Fatalf("Failed to unmarshal FlexPolicyEntry: %v", err)
		}

		if policy.PolicyName != "test-flex-policy" {
			t.Errorf("Expected policy name 'test-flex-policy', got '%s'", policy.PolicyName)
		}

		if policy.Description != "Test FlexConnect policy" {
			t.Errorf("Expected description 'Test FlexConnect policy', got '%s'", policy.Description)
		}

		if policy.IfNameVlanIds == nil || len(policy.IfNameVlanIds.IfNameVlanId) != 1 {
			t.Error("Expected 1 interface-vlan mapping")
		} else {
			ifVlan := policy.IfNameVlanIds.IfNameVlanId[0]
			if ifVlan.InterfaceName != "FastEthernet0/0/1" {
				t.Errorf("Expected interface name 'FastEthernet0/0/1', got '%s'", ifVlan.InterfaceName)
			}
			if ifVlan.VlanID != 50 {
				t.Errorf("Expected VLAN ID 50, got %d", ifVlan.VlanID)
			}
		}
	})
}

// =============================================================================
// 2. INTEGRATION TESTS (Actual API Calls to Live Controller)
// =============================================================================

// TestFlexConfigurationFunctions tests all FlexConnect configuration functions with a live controller
func TestFlexConfigurationFunctions(t *testing.T) {
	client := testutil.CreateTestClientFromEnv(t)
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
		if endpoint != "/restconf/data/Cisco-IOS-XE-wireless-flex-cfg:flex-cfg-data" {
			t.Errorf("FlexCfgEndpoint unexpected value: got %s", endpoint)
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
		expectedEndpoint := "/restconf/data/Cisco-IOS-XE-wireless-flex-cfg:flex-cfg-data/flex-policy-entries"
		if endpoint != expectedEndpoint {
			t.Errorf("FlexCfgDataEndpoint unexpected value: expected %s, got %s", expectedEndpoint, endpoint)
		}
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
