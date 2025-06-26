// Package radio provides radio configuration test functionality for the Cisco Wireless Network Controller API.
package radio

import (
	"encoding/json"
	"testing"

	wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil"
)

// =============================================================================
// 1. UNIT TESTS (Structure/Type Validation & JSON Serialization/Deserialization)
// =============================================================================

// getTestClient creates a test client using environment variables
func getTestClient(t *testing.T) *wnc.Client {
	return testutil.CreateTestClientFromEnv(t)
}

// RadioCfgTestDataCollector holds test data for radio configuration functions
type RadioCfgTestDataCollector struct {
	Data map[string]interface{} `json:"radio_cfg_test_data"`
}

var radioCfgTestDataCollector = RadioCfgTestDataCollector{
	Data: make(map[string]interface{}),
}

func runRadioCfgTestAndCollectData(t *testing.T, testName string, testFunc func() (interface{}, error)) {
	data, err := testFunc()
	if err != nil {
		t.Logf("%s returned error: %v", testName, err)
		radioCfgTestDataCollector.Data[testName] = map[string]interface{}{
			"error":   err.Error(),
			"success": false,
		}
	} else {
		t.Logf("%s executed successfully", testName)
		radioCfgTestDataCollector.Data[testName] = map[string]interface{}{
			"data":    data,
			"success": true,
		}
	}
}

// =============================================================================
// 2. INTEGRATION TESTS (API Endpoint Testing with Live Data Validation)
// =============================================================================

// TestRadioConfigurationFunctions tests all radio configuration functions with real WNC data collection
func TestRadioConfigurationFunctions(t *testing.T) {
	client := getTestClient(t)
	ctx, cancel := testutil.CreateDefaultTestContext()
	defer cancel()

	t.Run("GetRadioCfg", func(t *testing.T) {
		runRadioCfgTestAndCollectData(t, "GetRadioCfg", func() (interface{}, error) {
			return GetRadioCfg(client, ctx)
		})
	})

	t.Run("GetRadioProfiles", func(t *testing.T) {
		runRadioCfgTestAndCollectData(t, "GetRadioProfiles", func() (interface{}, error) {
			return GetRadioProfiles(client, ctx)
		})
	})

	// Save collected test data to file
	if len(radioCfgTestDataCollector.Data) > 0 {
		if err := testutil.SaveTestDataToFile("radio_cfg_test_data_collected.json", radioCfgTestDataCollector.Data); err != nil {
			t.Logf("Warning: Could not save test data: %v", err)
		} else {
			t.Logf("Test data saved to %s/radio_cfg_test_data_collected.json", testutil.TestDataDir)
		}
	}
}

// TestRadioCfgDataStructures tests the basic structure of radio configuration data types
func TestRadioCfgDataStructures(t *testing.T) {
	// Sample radio configuration data based on real WNC response structure
	sampleJSON := `{
		"Cisco-IOS-XE-wireless-radio-cfg:radio-cfg-data": {
			"radio-profiles": {
				"radio-profile": [
					{
						"name": "default-radio-profile",
						"desc": "Default radio profile for wireless operations",
						"mesh-backhaul": false
					    },
					{
						"name": "high-density-profile",
						"desc": "High density radio profile",
						"mesh-backhaul": true
					    }
				]
			    }
		    }
	    }`

	// Test unmarshaling into RadioCfgResponse
	var radioCfg RadioCfgResponse
	err := json.Unmarshal([]byte(sampleJSON), &radioCfg)
	if err != nil {
		t.Fatalf("Failed to unmarshal RadioCfgResponse: %v", err)
	}

	// Test that data was properly unmarshaled
	profiles := radioCfg.CiscoIOSXEWirelessRadioCfgData.RadioProfiles.RadioProfile
	if len(profiles) == 0 {
		t.Error("Expected at least one radio profile")
	}

	// Validate first radio profile
	if len(profiles) > 0 {
		profile := profiles[0]
		if profile.Name != "default-radio-profile" {
			t.Errorf("Expected profile name 'default-radio-profile', got '%s'", profile.Name)
		}

		if profile.Desc != "Default radio profile for wireless operations" {
			t.Errorf("Expected profile description to match, got '%s'", profile.Desc)
		}

		if profile.MeshBackhaul {
			t.Error("Expected mesh backhaul to be false for default profile")
		}
	}

	// Validate second radio profile if present
	if len(profiles) > 1 {
		profile := profiles[1]
		if profile.Name != "high-density-profile" {
			t.Errorf("Expected profile name 'high-density-profile', got '%s'", profile.Name)
		}

		if !profile.MeshBackhaul {
			t.Error("Expected mesh backhaul to be true for high-density profile")
		}
	}

	// Test RadioProfilesResponse structure
	sampleProfilesJSON := `{
		"Cisco-IOS-XE-wireless-radio-cfg:radio-profiles": {
			"radio-profile": [
				{
					"name": "test-profile",
					"desc": "Test radio profile",
					"mesh-backhaul": true
				    }
			]
		    }
	    }`

	var profilesResp RadioProfilesResponse
	err = json.Unmarshal([]byte(sampleProfilesJSON), &profilesResp)
	if err != nil {
		t.Fatalf("Failed to unmarshal RadioProfilesResponse: %v", err)
	}

	if len(profilesResp.RadioProfiles.RadioProfile) == 0 {
		t.Error("Expected at least one radio profile in profiles response")
	}

	_, err = json.Marshal(radioCfg)
	if err != nil {
		t.Errorf("Failed to marshal RadioCfgResponse back to JSON: %v", err)
	}

	_, err = json.Marshal(profilesResp)
	if err != nil {
		t.Errorf("Failed to marshal RadioProfilesResponse back to JSON: %v", err)
	}
}
