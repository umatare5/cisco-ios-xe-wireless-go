// Package cts provides CTS (Cisco TrustSec) configuration test functionality for the Cisco Wireless Network Controller API.
package cts

import (
	"encoding/json"
	"testing"
)

// =============================================================================
// 1. UNIT TESTS (Structure/Type Validation & JSON Serialization/Deserialization)
// =============================================================================

func TestCtsSxpCfgDataStructures(t *testing.T) {
	// Test CtsSxpCfgResponse structure
	t.Run("CtsSxpCfgResponse", func(t *testing.T) {
		sampleJSON := `{
			"Cisco-IOS-XE-wireless-cts-sxp-cfg:cts-sxp-cfg-data": {
				"cts-sxp-configuration": {
					"cts-sxp-config": [
						{
							"sxp-profile-name": "corporate-sxp"
						},
						{
							"sxp-profile-name": "guest-sxp"
						}
					]
				}
			}
		}`

		var response CtsSxpCfgResponse
		err := json.Unmarshal([]byte(sampleJSON), &response)
		if err != nil {
			t.Fatalf("Failed to unmarshal CtsSxpCfgResponse: %v", err)
		}

		if len(response.CiscoIOSXEWirelessCtsSxpCfgCtsSxpCfgData.CtsSxpConfiguration.CtsSxpConfig) != 2 {
			t.Errorf("Expected 2 cts-sxp-config entries, got %d",
				len(response.CiscoIOSXEWirelessCtsSxpCfgCtsSxpCfgData.CtsSxpConfiguration.CtsSxpConfig))
		}

		if response.CiscoIOSXEWirelessCtsSxpCfgCtsSxpCfgData.CtsSxpConfiguration.CtsSxpConfig[0].SxpProfileName != "corporate-sxp" {
			t.Errorf("Expected first profile name 'corporate-sxp', got '%s'",
				response.CiscoIOSXEWirelessCtsSxpCfgCtsSxpCfgData.CtsSxpConfiguration.CtsSxpConfig[0].SxpProfileName)
		}
	})

	// Test CtsSxpConfigurationResponse structure
	t.Run("CtsSxpConfigurationResponse", func(t *testing.T) {
		sampleJSON := `{
			"Cisco-IOS-XE-wireless-cts-sxp-cfg:cts-sxp-configuration": {
				"cts-sxp-config": [
					{
						"sxp-profile-name": "production-sxp"
					}
				]
			}
		}`

		var response CtsSxpConfigurationResponse
		err := json.Unmarshal([]byte(sampleJSON), &response)
		if err != nil {
			t.Fatalf("Failed to unmarshal CtsSxpConfigurationResponse: %v", err)
		}

		if len(response.CtsSxpConfiguration.CtsSxpConfig) != 1 {
			t.Errorf("Expected 1 cts-sxp-config entry, got %d",
				len(response.CtsSxpConfiguration.CtsSxpConfig))
		}

		if response.CtsSxpConfiguration.CtsSxpConfig[0].SxpProfileName != "production-sxp" {
			t.Errorf("Expected profile name 'production-sxp', got '%s'",
				response.CtsSxpConfiguration.CtsSxpConfig[0].SxpProfileName)
		}
	})

	// Test CtsSxpConfig structure
	t.Run("CtsSxpConfig", func(t *testing.T) {
		sampleJSON := `{
			"sxp-profile-name": "test-sxp-profile"
		}`

		var config CtsSxpConfig
		err := json.Unmarshal([]byte(sampleJSON), &config)
		if err != nil {
			t.Fatalf("Failed to unmarshal CtsSxpConfig: %v", err)
		}

		if config.SxpProfileName != "test-sxp-profile" {
			t.Errorf("Expected profile name 'test-sxp-profile', got '%s'", config.SxpProfileName)
		}
	})
}
