// Package flex provides FlexConnect configuration test functionality for the Cisco Wireless Network Controller API.
package flex

import (
	"encoding/json"
	"testing"
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
