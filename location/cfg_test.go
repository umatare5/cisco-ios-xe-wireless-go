// Package location provides location configuration test functionality for the Cisco Wireless Network Controller API.
package location

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

func TestLocationCfgDataStructures(t *testing.T) {
	// Test LocationCfgResponse structure
	t.Run("LocationCfgResponse", func(t *testing.T) {
		sampleJSON := `{
			"Cisco-IOS-XE-wireless-location-cfg:location-cfg-data": {
				"nmsp-config": {}
			}
		}`

		var response LocationCfgResponse
		err := json.Unmarshal([]byte(sampleJSON), &response)
		if err != nil {
			t.Fatalf("Failed to unmarshal LocationCfgResponse: %v", err)
		}

		// Since nmsp-config is an empty struct, just verify it unmarshals successfully
	})

	// Test LocationNmspConfigResponse structure
	t.Run("LocationNmspConfigResponse", func(t *testing.T) {
		sampleJSON := `{
			"Cisco-IOS-XE-wireless-location-cfg:nmsp-config": {}
		}`

		var response LocationNmspConfigResponse
		err := json.Unmarshal([]byte(sampleJSON), &response)
		if err != nil {
			t.Fatalf("Failed to unmarshal LocationNmspConfigResponse: %v", err)
		}

		// Since nmsp-config is an empty struct, just verify it unmarshals successfully
	})
}

// =============================================================================
// 2. INTEGRATION TESTS (Actual API Calls to Live Controller)
// =============================================================================

// TestLocationConfigurationFunctions tests all location configuration functions with a live controller
func TestLocationConfigurationFunctions(t *testing.T) {
	client := testutil.CreateTestClientFromEnv(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test GetLocationCfg function
	t.Run("GetLocationCfg", func(t *testing.T) {
		data, err := GetLocationCfg(client, ctx)
		if err != nil {
			t.Fatalf("GetLocationCfg failed: %v", err)
		}

		// Validate basic structure
		if data == nil {
			t.Fatal("GetLocationCfg returned nil data")
		}

		// Save test data for analysis
		if err := testutil.SaveTestDataToFile("location_cfg_data.json", data); err != nil {
			t.Logf("Warning: Could not save test data: %v", err)
		} else {
			t.Logf("Location config data saved to test_data/location_cfg_data.json")
		}

		// Validate endpoint was constructed correctly
		endpoint := LocationCfgEndpoint
		if endpoint == "" {
			t.Error("LocationCfgEndpoint should not be empty")
		}
		if endpoint != "/restconf/data/Cisco-IOS-XE-wireless-location-cfg:location-cfg-data" {
			t.Errorf("LocationCfgEndpoint unexpected value: got %s", endpoint)
		}
	})

	// Test GetLocationNmspConfig function
	t.Run("GetLocationNmspConfig", func(t *testing.T) {
		data, err := GetLocationNmspConfig(client, ctx)
		if err != nil {
			t.Fatalf("GetLocationNmspConfig failed: %v", err)
		}

		// Validate basic structure
		if data == nil {
			t.Fatal("GetLocationNmspConfig returned nil data")
		}

		// Save test data for analysis
		if err := testutil.SaveTestDataToFile("location_nmsp_config_data.json", data); err != nil {
			t.Logf("Warning: Could not save test data: %v", err)
		} else {
			t.Logf("Location NMSP config data saved to test_data/location_nmsp_config_data.json")
		}

		// Validate endpoint was constructed correctly
		endpoint := LocationCfgNmspConfigEndpoint
		if endpoint == "" {
			t.Error("LocationCfgNmspConfigEndpoint should not be empty")
		}
		expectedEndpoint := "/restconf/data/Cisco-IOS-XE-wireless-location-cfg:location-cfg-data/nmsp-config"
		if endpoint != expectedEndpoint {
			t.Errorf("LocationCfgNmspConfigEndpoint unexpected value: expected %s, got %s", expectedEndpoint, endpoint)
		}
	})
}

// TestLocationConfigurationEndpoints validates location configuration endpoint constants
func TestLocationConfigurationEndpoints(t *testing.T) {
	// Test base path validation
	t.Run("Validate_LocationCfgBasePath", func(t *testing.T) {
		expectedBasePath := "/restconf/data/Cisco-IOS-XE-wireless-location-cfg:location-cfg-data"
		if LocationCfgBasePath != expectedBasePath {
			t.Errorf("LocationCfgBasePath mismatch: expected %s, got %s", expectedBasePath, LocationCfgBasePath)
		}
	})

	// Test endpoint validation
	t.Run("Validate_LocationCfgEndpoint", func(t *testing.T) {
		if LocationCfgEndpoint != LocationCfgBasePath {
			t.Errorf("LocationCfgEndpoint should equal LocationCfgBasePath: expected %s, got %s", LocationCfgBasePath, LocationCfgEndpoint)
		}
	})

	// Test NMSP config endpoint validation
	t.Run("Validate_LocationCfgNmspConfigEndpoint", func(t *testing.T) {
		expectedEndpoint := LocationCfgBasePath + "/nmsp-config"
		if LocationCfgNmspConfigEndpoint != expectedEndpoint {
			t.Errorf("LocationCfgNmspConfigEndpoint mismatch: expected %s, got %s", expectedEndpoint, LocationCfgNmspConfigEndpoint)
		}
	})
}
