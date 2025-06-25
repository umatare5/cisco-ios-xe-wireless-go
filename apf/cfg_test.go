// Package apf provides Access Point Filter configuration test functionality for the Cisco Wireless Network Controller API.
package apf

import (
	"encoding/json"
	"testing"
)

// =============================================================================
// 1. UNIT TESTS (Structure/Type Validation & JSON Serialization/Deserialization)
// =============================================================================

func TestApfCfgDataStructures(t *testing.T) {
	// Test ApfCfgResponse structure
	t.Run("ApfCfgResponse", func(t *testing.T) {
		sampleJSON := `{
			"Cisco-IOS-XE-wireless-apf-cfg:apf-cfg-data": {
				"apf": {
					"system-mgmt-via-wireless": true,
					"network-name": "corporate-network"
				}
			}
		}`

		var response ApfCfgResponse
		err := json.Unmarshal([]byte(sampleJSON), &response)
		if err != nil {
			t.Fatalf("Failed to unmarshal ApfCfgResponse: %v", err)
		}

		if !response.CiscoIOSXEWirelessApfCfgApfCfgData.Apf.SystemMgmtViaWireless {
			t.Error("Expected system-mgmt-via-wireless to be true")
		}

		if response.CiscoIOSXEWirelessApfCfgApfCfgData.Apf.NetworkName != "corporate-network" {
			t.Errorf("Expected network name 'corporate-network', got '%s'",
				response.CiscoIOSXEWirelessApfCfgApfCfgData.Apf.NetworkName)
		}
	})

	// Test ApfCfgApfResponse structure
	t.Run("ApfCfgApfResponse", func(t *testing.T) {
		sampleJSON := `{
			"Cisco-IOS-XE-wireless-apf-cfg:apf": {
				"system-mgmt-via-wireless": false,
				"network-name": "guest-network"
			}
		}`

		var response ApfCfgApfResponse
		err := json.Unmarshal([]byte(sampleJSON), &response)
		if err != nil {
			t.Fatalf("Failed to unmarshal ApfCfgApfResponse: %v", err)
		}

		if response.Apf.SystemMgmtViaWireless {
			t.Error("Expected system-mgmt-via-wireless to be false")
		}

		if response.Apf.NetworkName != "guest-network" {
			t.Errorf("Expected network name 'guest-network', got '%s'", response.Apf.NetworkName)
		}
	})

	// Test Apf structure
	t.Run("Apf", func(t *testing.T) {
		sampleJSON := `{
			"system-mgmt-via-wireless": true,
			"network-name": "production-network"
		}`

		var apf Apf
		err := json.Unmarshal([]byte(sampleJSON), &apf)
		if err != nil {
			t.Fatalf("Failed to unmarshal Apf: %v", err)
		}

		if !apf.SystemMgmtViaWireless {
			t.Error("Expected system-mgmt-via-wireless to be true")
		}

		if apf.NetworkName != "production-network" {
			t.Errorf("Expected network name 'production-network', got '%s'", apf.NetworkName)
		}
	})
}
