// Package apf provides Access Point Filter configuration test functionality for the Cisco Wireless Network Controller API.
package apf

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

// =============================================================================
// 2. INTEGRATION TESTS (Actual API Calls to Live Controller)
// =============================================================================

// TestApfConfigurationFunctions tests all APF configuration functions with a live controller
func TestApfConfigurationFunctions(t *testing.T) {
	client := testutil.CreateTestClientFromEnv(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test GetApfCfg function
	t.Run("GetApfCfg", func(t *testing.T) {
		data, err := GetApfCfg(client, ctx)
		if err != nil {
			t.Fatalf("GetApfCfg failed: %v", err)
		}

		// Validate basic structure
		if data == nil {
			t.Fatal("GetApfCfg returned nil data")
		}

		// Save test data for analysis
		if err := testutil.SaveTestDataToFile("apf_cfg_data.json", data); err != nil {
			t.Logf("Warning: Could not save test data: %v", err)
		} else {
			t.Logf("APF config data saved to test_data/apf_cfg_data.json")
		}

		// Validate endpoint was constructed correctly
		endpoint := ApfCfgEndpoint
		if endpoint == "" {
			t.Error("ApfCfgEndpoint should not be empty")
		}
		if endpoint != "/restconf/data/Cisco-IOS-XE-wireless-apf-cfg:apf-cfg-data" {
			t.Errorf("ApfCfgEndpoint unexpected value: got %s", endpoint)
		}
	})

	// Test GetApf function
	t.Run("GetApf", func(t *testing.T) {
		data, err := GetApf(client, ctx)
		if err != nil {
			t.Fatalf("GetApf failed: %v", err)
		}

		// Validate basic structure
		if data == nil {
			t.Fatal("GetApf returned nil data")
		}

		// Save test data for analysis
		if err := testutil.SaveTestDataToFile("apf_data.json", data); err != nil {
			t.Logf("Warning: Could not save test data: %v", err)
		} else {
			t.Logf("APF data saved to test_data/apf_data.json")
		}

		// Validate endpoint was constructed correctly
		endpoint := ApfEndpoint
		if endpoint == "" {
			t.Error("ApfEndpoint should not be empty")
		}
		expectedEndpoint := "/restconf/data/Cisco-IOS-XE-wireless-apf-cfg:apf-cfg-data/apf"
		if endpoint != expectedEndpoint {
			t.Errorf("ApfEndpoint unexpected value: expected %s, got %s", expectedEndpoint, endpoint)
		}
	})
}

// TestApfConfigurationEndpoints validates APF configuration endpoint constants
func TestApfConfigurationEndpoints(t *testing.T) {
	// Test base path validation
	t.Run("Validate_ApfCfgBasePath", func(t *testing.T) {
		expectedBasePath := "/restconf/data/Cisco-IOS-XE-wireless-apf-cfg:apf-cfg-data"
		if ApfCfgBasePath != expectedBasePath {
			t.Errorf("ApfCfgBasePath mismatch: expected %s, got %s", expectedBasePath, ApfCfgBasePath)
		}
	})

	// Test endpoint validation
	t.Run("Validate_ApfCfgEndpoint", func(t *testing.T) {
		if ApfCfgEndpoint != ApfCfgBasePath {
			t.Errorf("ApfCfgEndpoint should equal ApfCfgBasePath: expected %s, got %s", ApfCfgBasePath, ApfCfgEndpoint)
		}
	})

	// Test APF specific endpoint validation
	t.Run("Validate_ApfEndpoint", func(t *testing.T) {
		expectedEndpoint := ApfCfgBasePath + "/apf"
		if ApfEndpoint != expectedEndpoint {
			t.Errorf("ApfEndpoint mismatch: expected %s, got %s", expectedEndpoint, ApfEndpoint)
		}
	})
}
