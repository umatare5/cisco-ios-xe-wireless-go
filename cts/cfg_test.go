// Package cts provides CTS (Cisco TrustSec) configuration test functionality for the Cisco Wireless Network Controller API.
package cts

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

// =============================================================================
// 2. INTEGRATION TESTS (Actual API Calls to Live Controller)
// =============================================================================

// TestCtsConfigurationFunctions tests all CTS configuration functions with a live controller
func TestCtsConfigurationFunctions(t *testing.T) {
	client := testutil.CreateTestClientFromEnv(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test GetCtsSxpCfg function
	t.Run("GetCtsSxpCfg", func(t *testing.T) {
		data, err := GetCtsSxpCfg(client, ctx)
		if err != nil {
			t.Fatalf("GetCtsSxpCfg failed: %v", err)
		}

		// Validate basic structure
		if data == nil {
			t.Fatal("GetCtsSxpCfg returned nil data")
		}

		// Save test data for analysis
		if err := testutil.SaveTestDataToFile("cts_cfg_data.json", data); err != nil {
			t.Logf("Warning: Could not save test data: %v", err)
		} else {
			t.Logf("CTS config data saved to test_data/cts_cfg_data.json")
		}

		// Validate endpoint was constructed correctly
		endpoint := CtsSxpCfgEndpoint
		if endpoint == "" {
			t.Error("CtsSxpCfgEndpoint should not be empty")
		}
		if endpoint != "/restconf/data/Cisco-IOS-XE-wireless-cts-sxp-cfg:cts-sxp-cfg-data" {
			t.Errorf("CtsSxpCfgEndpoint unexpected value: got %s", endpoint)
		}
	})

	// Test GetCtsSxpConfiguration function
	t.Run("GetCtsSxpConfiguration", func(t *testing.T) {
		data, err := GetCtsSxpConfiguration(client, ctx)
		if err != nil {
			t.Fatalf("GetCtsSxpConfiguration failed: %v", err)
		}

		// Validate basic structure
		if data == nil {
			t.Fatal("GetCtsSxpConfiguration returned nil data")
		}

		// Save test data for analysis
		if err := testutil.SaveTestDataToFile("cts_sxp_configuration_data.json", data); err != nil {
			t.Logf("Warning: Could not save test data: %v", err)
		} else {
			t.Logf("CTS SXP configuration data saved to test_data/cts_sxp_configuration_data.json")
		}

		// Validate endpoint was constructed correctly
		endpoint := CtsSxpConfigurationEndpoint
		if endpoint == "" {
			t.Error("CtsSxpConfigurationEndpoint should not be empty")
		}
		expectedEndpoint := "/restconf/data/Cisco-IOS-XE-wireless-cts-sxp-cfg:cts-sxp-cfg-data/cts-sxp-configuration"
		if endpoint != expectedEndpoint {
			t.Errorf("CtsSxpConfigurationEndpoint unexpected value: expected %s, got %s", expectedEndpoint, endpoint)
		}
	})
}

// TestCtsConfigurationEndpoints validates CTS configuration endpoint constants
func TestCtsConfigurationEndpoints(t *testing.T) {
	// Test base path validation
	t.Run("Validate_CtsSxpCfgBasePath", func(t *testing.T) {
		expectedBasePath := "/restconf/data/Cisco-IOS-XE-wireless-cts-sxp-cfg:cts-sxp-cfg-data"
		if CtsSxpCfgBasePath != expectedBasePath {
			t.Errorf("CtsSxpCfgBasePath mismatch: expected %s, got %s", expectedBasePath, CtsSxpCfgBasePath)
		}
	})

	// Test endpoint validation
	t.Run("Validate_CtsSxpCfgEndpoint", func(t *testing.T) {
		if CtsSxpCfgEndpoint != CtsSxpCfgBasePath {
			t.Errorf("CtsSxpCfgEndpoint should equal CtsSxpCfgBasePath: expected %s, got %s", CtsSxpCfgBasePath, CtsSxpCfgEndpoint)
		}
	})

	// Test specific configuration endpoint validation
	t.Run("Validate_CtsSxpConfigurationEndpoint", func(t *testing.T) {
		expectedEndpoint := CtsSxpCfgBasePath + "/cts-sxp-configuration"
		if CtsSxpConfigurationEndpoint != expectedEndpoint {
			t.Errorf("CtsSxpConfigurationEndpoint mismatch: expected %s, got %s", expectedEndpoint, CtsSxpConfigurationEndpoint)
		}
	})
}
