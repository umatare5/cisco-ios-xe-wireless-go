// Package rf provides radio frequency configuration test functionality for the Cisco Wireless Network Controller API.
package rf

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	wnc "github.com/umatare5/cisco-xe-wireless-restconf-go"
	"github.com/umatare5/cisco-xe-wireless-restconf-go/internal/testutil"
)

// =============================================================================
// 1. UNIT TESTS (Structure/Type Validation & JSON Serialization/Deserialization)
// =============================================================================

const (
	TestDataDir = "test_data"
)

// getTestClient creates a test client using environment variables
func getTestClient(t *testing.T) *wnc.Client {
	return testutil.CreateTestClientFromEnv(t)
}

// RfCfgTestDataCollector holds test data for RF configuration functions
type RfCfgTestDataCollector struct {
	Data map[string]interface{} `json:"rf_cfg_test_data"`
}

var rfCfgTestDataCollector = RfCfgTestDataCollector{
	Data: make(map[string]interface{}),
}

func runRfCfgTestAndCollectData(t *testing.T, testName string, testFunc func() (interface{}, error)) {
	data, err := testFunc()
	if err != nil {
		t.Logf("%s returned error: %v", testName, err)
		rfCfgTestDataCollector.Data[testName] = map[string]interface{}{
			"error":   err.Error(),
			"success": false,
		}
	} else {
		t.Logf("%s executed successfully", testName)
		rfCfgTestDataCollector.Data[testName] = map[string]interface{}{
			"data":    data,
			"success": true,
		}
	}
}

// =============================================================================
// 2. INTEGRATION TESTS (API Endpoint Testing with Live Data Validation)
// =============================================================================

// TestRfConfigurationFunctions tests all RF configuration functions with real WNC data collection
func TestRfConfigurationFunctions(t *testing.T) {
	client := getTestClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	t.Run("GetRfCfg", func(t *testing.T) {
		runRfCfgTestAndCollectData(t, "GetRfCfg", func() (interface{}, error) {
			return GetRfCfg(client, ctx)
		})
	})

	t.Run("GetRfMultiBssidProfiles", func(t *testing.T) {
		runRfCfgTestAndCollectData(t, "GetRfMultiBssidProfiles", func() (interface{}, error) {
			return GetRfMultiBssidProfiles(client, ctx)
		})
	})

	t.Run("GetRfAtfPolicies", func(t *testing.T) {
		runRfCfgTestAndCollectData(t, "GetRfAtfPolicies", func() (interface{}, error) {
			return GetRfAtfPolicies(client, ctx)
		})
	})

	t.Run("GetRfTags", func(t *testing.T) {
		runRfCfgTestAndCollectData(t, "GetRfTags", func() (interface{}, error) {
			return GetRfTags(client, ctx)
		})
	})

	t.Run("GetRfProfiles", func(t *testing.T) {
		runRfCfgTestAndCollectData(t, "GetRfProfiles", func() (interface{}, error) {
			return GetRfProfiles(client, ctx)
		})
	})

	t.Run("GetRfProfileDefaultEntries", func(t *testing.T) {
		runRfCfgTestAndCollectData(t, "GetRfProfileDefaultEntries", func() (interface{}, error) {
			return GetRfProfileDefaultEntries(client, ctx)
		})
	})

	// Save collected test data to file
	if len(rfCfgTestDataCollector.Data) > 0 {
		if err := testutil.SaveTestDataToFile("rf_cfg_test_data_collected.json", rfCfgTestDataCollector.Data); err != nil {
			t.Logf("Warning: Could not save test data: %v", err)
		} else {
			t.Logf("Test data saved to %s/rf_cfg_test_data_collected.json", TestDataDir)
		}
	}
}

// TestRfCfgDataStructures tests the basic structure of RF configuration data types
func TestRfCfgDataStructures(t *testing.T) {
	// Sample RF configuration data based on real WNC response structure
	sampleJSON := `{
		"Cisco-IOS-XE-wireless-rf-cfg:rf-cfg-data": {
			"multi-bssid-profiles": {
				"multi-bssid-profile": [
					{
						"profile-name": "default-multi-bssid",
						"description": "Default multi-BSSID profile"
					    }
				]
			    },
			"atf-policies": {
				"atf-policy": []
			    },
			"rf-tags": {
				"rf-tag": [
					{
						"tag-name": "default-rf-tag",
						"dot11a-rf-profile-name": "default-rf-profile-5g",
						"dot11b-rf-profile-name": "default-rf-profile-2g",
						"description": "Default RF tag configuration",
						"rf-tag-radio-profiles": {
							"rf-tag-radio-profile": [
								{
									"slot-id": "0",
									"band-id": "2.4GHz"
								    },
								{
									"slot-id": "1",
									"band-id": "5GHz"
								    }
							]
						    }
					    }
				]
			    },
			"rf-profiles": {
				"rf-profile": [
					{
						"name": "default-rf-profile-2g",
						"description": "Default 2.4GHz RF profile",
						"status": true,
						"band": "2.4GHz",
						"data-rate-6m": "mandatory",
						"data-rate-12m": "supported",
						"data-rate-24m": "supported",
						"rx-sen-sop-custom": -80,
						"rf-dca-chan-width": "20MHz",
						"tx-power-v1-threshold": 6,
						"coverage-data-packet-rssi-threshold": -80,
						"min-num-clients": 5,
						"coverage-voice-packet-rssi-threshold": -75
					    }
				]
			    },
			"rf-profile-default-entries": {
				"rf-profile-default-entry": [
					{
						"rf-profile-name": "default-rf-profile-2g",
						"band": "2.4GHz"
					    }
				]
			    }
		    }
	    }`

	// Test unmarshaling into RfCfgResponse
	var rfCfg RfCfgResponse
	err := json.Unmarshal([]byte(sampleJSON), &rfCfg)
	if err != nil {
		t.Fatalf("Failed to unmarshal RfCfgResponse: %v", err)
	}

	// Test Multi-BSSID profiles
	multiBssidProfiles := rfCfg.CiscoIOSXEWirelessRfCfgRfCfgData.MultiBssidProfiles.MultiBssidProfile
	if len(multiBssidProfiles) == 0 {
		t.Error("Expected at least one multi-BSSID profile")
	} else {
		profile := multiBssidProfiles[0]
		if profile.ProfileName != "default-multi-bssid" {
			t.Errorf("Expected profile name 'default-multi-bssid', got '%s'", profile.ProfileName)
		}
	}

	// Test RF tags
	rfTags := rfCfg.CiscoIOSXEWirelessRfCfgRfCfgData.RfTags.RfTag
	if len(rfTags) == 0 {
		t.Error("Expected at least one RF tag")
	} else {
		tag := rfTags[0]
		if tag.TagName != "default-rf-tag" {
			t.Errorf("Expected tag name 'default-rf-tag', got '%s'", tag.TagName)
		}

		if tag.Dot11ARfProfileName != "default-rf-profile-5g" {
			t.Errorf("Expected 5G profile name 'default-rf-profile-5g', got '%s'", tag.Dot11ARfProfileName)
		}

		if len(tag.RfTagRadioProfiles.RfTagRadioProfile) < 2 {
			t.Error("Expected at least 2 radio profiles in RF tag")
		}
	}

	// Test RF profiles
	rfProfiles := rfCfg.CiscoIOSXEWirelessRfCfgRfCfgData.RfProfiles.RfProfile
	if len(rfProfiles) == 0 {
		t.Error("Expected at least one RF profile")
	} else {
		profile := rfProfiles[0]
		if profile.Name != "default-rf-profile-2g" {
			t.Errorf("Expected profile name 'default-rf-profile-2g', got '%s'", profile.Name)
		}

		if !profile.Status {
			t.Error("Expected RF profile status to be true")
		}

		if profile.Band != "2.4GHz" {
			t.Errorf("Expected band '2.4GHz', got '%s'", profile.Band)
		}

		if profile.DataRate6M != "mandatory" {
			t.Errorf("Expected data rate 6M 'mandatory', got '%s'", profile.DataRate6M)
		}
	}

	// Test RfTagsResponse structure separately
	sampleTagsJSON := `{
		"Cisco-IOS-XE-wireless-rf-cfg:rf-tags": {
			"rf-tag": [
				{
					"tag-name": "test-rf-tag",
					"description": "Test RF tag",
					"rf-tag-radio-profiles": {
						"rf-tag-radio-profile": []
					    }
				    }
			]
		    }
	    }`

	var tagsResp RfTagsResponse
	err = json.Unmarshal([]byte(sampleTagsJSON), &tagsResp)
	if err != nil {
		t.Fatalf("Failed to unmarshal RfTagsResponse: %v", err)
	}

	if len(tagsResp.RfTags.RfTag) == 0 {
		t.Error("Expected at least one RF tag in tags response")
	}

	_, err = json.Marshal(rfCfg)
	if err != nil {
		t.Errorf("Failed to marshal RfCfgResponse back to JSON: %v", err)
	}

	_, err = json.Marshal(tagsResp)
	if err != nil {
		t.Errorf("Failed to marshal RfTagsResponse back to JSON: %v", err)
	}
}
