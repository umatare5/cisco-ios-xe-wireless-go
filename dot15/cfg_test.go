// Package dot15 provides 802.15.4 configuration test functionality for the Cisco Wireless Network Controller API.
package dot15

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

func TestDot15CfgDataStructures(t *testing.T) {
	// Test Dot15CfgResponse structure
	t.Run("Dot15CfgResponse", func(t *testing.T) {
		sampleJSON := `{
			"Cisco-IOS-XE-wireless-dot15-cfg:dot15-cfg-data": {
				"dot15-global-config": {}
			}
		}`

		var response Dot15CfgResponse
		err := json.Unmarshal([]byte(sampleJSON), &response)
		if err != nil {
			t.Fatalf("Failed to unmarshal Dot15CfgResponse: %v", err)
		}

		// Since dot15-global-config is an empty struct, just verify it unmarshals successfully
	})

	// Test Dot15GlobalConfigResponse structure
	t.Run("Dot15GlobalConfigResponse", func(t *testing.T) {
		sampleJSON := `{
			"Cisco-IOS-XE-wireless-dot15-cfg:dot15-global-config": {}
		}`

		var response Dot15GlobalConfigResponse
		err := json.Unmarshal([]byte(sampleJSON), &response)
		if err != nil {
			t.Fatalf("Failed to unmarshal Dot15GlobalConfigResponse: %v", err)
		}

		// Since dot15-global-config is an empty struct, just verify it unmarshals successfully
	})
}

// =============================================================================
// 2. INTEGRATION TESTS (Actual API Calls to Live Controller)
// =============================================================================

// TestDot15ConfigurationFunctions tests all 802.15 configuration functions with a live controller
func TestDot15ConfigurationFunctions(t *testing.T) {
	client := testutil.CreateTestClientFromEnv(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test GetDot15Cfg function
	t.Run("GetDot15Cfg", func(t *testing.T) {
		data, err := GetDot15Cfg(client, ctx)
		if err != nil {
			t.Fatalf("GetDot15Cfg failed: %v", err)
		}

		// Validate basic structure
		if data == nil {
			t.Fatal("GetDot15Cfg returned nil data")
		}

		// Save test data for analysis
		if err := testutil.SaveTestDataToFile("dot15_cfg_data.json", data); err != nil {
			t.Logf("Warning: Could not save test data: %v", err)
		} else {
			t.Logf("Dot15 config data saved to test_data/dot15_cfg_data.json")
		}

		// Validate endpoint was constructed correctly
		endpoint := Dot15CfgEndpoint
		if endpoint == "" {
			t.Error("Dot15CfgEndpoint should not be empty")
		}
		if endpoint != "/restconf/data/Cisco-IOS-XE-wireless-dot15-cfg:dot15-cfg-data" {
			t.Errorf("Dot15CfgEndpoint unexpected value: got %s", endpoint)
		}
	})

	// Test GetDot15GlobalConfig function
	t.Run("GetDot15GlobalConfig", func(t *testing.T) {
		data, err := GetDot15GlobalConfig(client, ctx)
		if err != nil {
			t.Fatalf("GetDot15GlobalConfig failed: %v", err)
		}

		// Validate basic structure
		if data == nil {
			t.Fatal("GetDot15GlobalConfig returned nil data")
		}

		// Save test data for analysis
		if err := testutil.SaveTestDataToFile("dot15_global_config_data.json", data); err != nil {
			t.Logf("Warning: Could not save test data: %v", err)
		} else {
			t.Logf("Dot15 global config data saved to test_data/dot15_global_config_data.json")
		}
	})
}

// TestDot15ConfigurationEndpoints validates 802.15 configuration endpoint constants
func TestDot15ConfigurationEndpoints(t *testing.T) {
	// Test base path validation
	t.Run("Validate_Dot15CfgBasePath", func(t *testing.T) {
		expectedBasePath := "/restconf/data/Cisco-IOS-XE-wireless-dot15-cfg:dot15-cfg-data"
		if Dot15CfgBasePath != expectedBasePath {
			t.Errorf("Dot15CfgBasePath mismatch: expected %s, got %s", expectedBasePath, Dot15CfgBasePath)
		}
	})

	// Test endpoint validation
	t.Run("Validate_Dot15CfgEndpoint", func(t *testing.T) {
		if Dot15CfgEndpoint != Dot15CfgBasePath {
			t.Errorf("Dot15CfgEndpoint should equal Dot15CfgBasePath: expected %s, got %s", Dot15CfgBasePath, Dot15CfgEndpoint)
		}
	})
}
