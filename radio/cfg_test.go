// Package radio provides radio configuration test functionality for the Cisco Wireless Network Controller API.
package radio

import (
	"context"
	"testing"
	"time"

	wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
	testutils "github.com/umatare5/cisco-ios-xe-wireless-go/internal/tests"
)

// =============================================================================
// 1. UNIT TESTS (Structure/Type Validation & JSON Serialization/Deserialization)
// =============================================================================

func TestRadioCfgDataStructures(t *testing.T) {
	testCases := []testutils.JSONTestCase{
		{
			Name: "RadioCfgResponse",
			JSONData: `{
				"Cisco-IOS-XE-wireless-radio-cfg:radio-cfg-data": {
					"radio-profiles": {
						"radio-profile": [
							{
								"name": "corporate-radio",
								"desc": "Corporate radio profile",
								"mesh-backhaul": true
							}
						]
					}
				}
			}`,
			Target:     &RadioCfgResponse{},
			TypeName:   "RadioCfgResponse",
			ShouldFail: false,
		},
		{
			Name: "RadioProfilesResponse",
			JSONData: `{
				"Cisco-IOS-XE-wireless-radio-cfg:radio-profiles": {
					"radio-profile": [
						{
							"name": "guest-radio",
							"desc": "Guest radio profile",
							"mesh-backhaul": false
						}
					]
				}
			}`,
			Target:     &RadioProfilesResponse{},
			TypeName:   "RadioProfilesResponse",
			ShouldFail: false,
		},
		{
			Name: "RadioProfile",
			JSONData: `{
				"name": "test-radio",
				"desc": "Test radio profile",
				"mesh-backhaul": true
			}`,
			Target:     &RadioProfile{},
			TypeName:   "RadioProfile",
			ShouldFail: false,
		},
	}

	testutils.RunJSONTests(t, testCases)

	// Additional field validation for successfully unmarshaled structures
	t.Run("RadioCfgResponseFieldValidation", func(t *testing.T) {
		var response RadioCfgResponse
		testutils.TestJSONUnmarshal(t, testCases[0].JSONData, &response, "RadioCfgResponse")

		testutils.ValidateJSONStructFields(t, "RadioCfgResponse", func() error {
			if len(response.CiscoIOSXEWirelessRadioCfgData.RadioProfiles.RadioProfile) != 1 {
				t.Errorf("Expected 1 radio profile, got %d",
					len(response.CiscoIOSXEWirelessRadioCfgData.RadioProfiles.RadioProfile))
			}
			if response.CiscoIOSXEWirelessRadioCfgData.RadioProfiles.RadioProfile[0].Name != "corporate-radio" {
				t.Errorf("Expected profile name 'corporate-radio', got '%s'",
					response.CiscoIOSXEWirelessRadioCfgData.RadioProfiles.RadioProfile[0].Name)
			}
			return nil
		})
	})

	t.Run("RadioProfilesResponseFieldValidation", func(t *testing.T) {
		var response RadioProfilesResponse
		testutils.TestJSONUnmarshal(t, testCases[1].JSONData, &response, "RadioProfilesResponse")

		testutils.ValidateJSONStructFields(t, "RadioProfilesResponse", func() error {
			if len(response.RadioProfiles.RadioProfile) != 1 {
				t.Errorf("Expected 1 radio profile, got %d", len(response.RadioProfiles.RadioProfile))
			}
			if response.RadioProfiles.RadioProfile[0].Name != "guest-radio" {
				t.Errorf("Expected profile name 'guest-radio', got '%s'", response.RadioProfiles.RadioProfile[0].Name)
			}
			return nil
		})
	})

	t.Run("RadioProfileFieldValidation", func(t *testing.T) {
		var profile RadioProfile
		testutils.TestJSONUnmarshal(t, testCases[2].JSONData, &profile, "RadioProfile")

		testutils.ValidateJSONStructFields(t, "RadioProfile", func() error {
			if profile.Name != "test-radio" {
				t.Errorf("Expected profile name 'test-radio', got '%s'", profile.Name)
			}
			return nil
		})
	})
}

// =============================================================================
// 2. INTEGRATION TESTS (Actual API Calls to Live Controller)
// =============================================================================

// TestRadioConfigurationFunctions tests all radio configuration functions with a live controller
func TestRadioConfigurationFunctions(t *testing.T) {
	client := testutils.GetTestClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test GetRadioCfg function
	t.Run("GetRadioCfg", func(t *testing.T) {
		data, err := GetRadioCfg(client, ctx)
		if err != nil {
			t.Fatalf("GetRadioCfg failed: %v", err)
		}

		// Validate basic structure
		if data == nil {
			t.Fatal("GetRadioCfg returned nil data")
		}

		// Save test data for analysis
		if err := testutils.SaveTestDataToFile("radio_cfg_data.json", data); err != nil {
			t.Logf("Warning: Could not save test data: %v", err)
		} else {
			t.Logf("Radio config data saved to test_data/radio_cfg_data.json")
		}

		// Validate endpoint was constructed correctly
		endpoint := RadioCfgEndpoint
		if endpoint == "" {
			t.Error("RadioCfgEndpoint should not be empty")
		}
	})

	// Test GetRadioProfiles function
	t.Run("GetRadioProfiles", func(t *testing.T) {
		data, err := GetRadioProfiles(client, ctx)
		if err != nil {
			t.Fatalf("GetRadioProfiles failed: %v", err)
		}

		// Validate basic structure
		if data == nil {
			t.Fatal("GetRadioProfiles returned nil data")
		}

		// Save test data for analysis
		if err := testutils.SaveTestDataToFile("radio_profiles_data.json", data); err != nil {
			t.Logf("Warning: Could not save test data: %v", err)
		} else {
			t.Logf("Radio profiles data saved to test_data/radio_profiles_data.json")
		}

		// Validate endpoint was constructed correctly
		endpoint := RadioProfilesEndpoint
		if endpoint == "" {
			t.Error("RadioProfilesEndpoint should not be empty")
		}
	})

	// Test error handling with nil client
	t.Run("RadioErrorHandling", func(t *testing.T) {
		t.Run("GetRadioCfgWithNilClient", func(t *testing.T) {
			_, err := GetRadioCfg(nil, ctx)
			if err == nil {
				t.Error("Expected error with nil client, got nil")
			}
			if err.Error() != "client is nil" {
				t.Errorf("Expected 'client is nil' error, got: %v", err)
			}
		})

		t.Run("GetRadioProfilesWithNilClient", func(t *testing.T) {
			_, err := GetRadioProfiles(nil, ctx)
			if err == nil {
				t.Error("Expected error with nil client, got nil")
			}
			if err.Error() != "client is nil" {
				t.Errorf("Expected 'client is nil' error, got: %v", err)
			}
		})
	})

	// Test context handling
	t.Run("ContextHandling", func(t *testing.T) {
		testutils.TestContextHandling(t, func(ctx context.Context, client *wnc.Client) error {
			_, err := GetRadioCfg(client, ctx)
			return err
		})

		testutils.TestContextHandling(t, func(ctx context.Context, client *wnc.Client) error {
			_, err := GetRadioProfiles(client, ctx)
			return err
		})
	})
}

// TestRadioConfigurationEndpoints validates radio configuration endpoint constants
func TestRadioConfigurationEndpoints(t *testing.T) {
	// Test base path validation
	t.Run("Validate_RadioCfgBasePath", func(t *testing.T) {
		expectedBasePath := "/restconf/data/Cisco-IOS-XE-wireless-radio-cfg:radio-cfg-data"
		if RadioCfgBasePath != expectedBasePath {
			t.Errorf("RadioCfgBasePath mismatch: expected %s, got %s", expectedBasePath, RadioCfgBasePath)
		}
	})

	// Test endpoint validation
	t.Run("Validate_RadioCfgEndpoint", func(t *testing.T) {
		if RadioCfgEndpoint != RadioCfgBasePath {
			t.Errorf("RadioCfgEndpoint should equal RadioCfgBasePath: expected %s, got %s", RadioCfgBasePath, RadioCfgEndpoint)
		}
	})

	// Test radio profiles endpoint validation
	t.Run("Validate_RadioProfilesEndpoint", func(t *testing.T) {
		expectedEndpoint := RadioCfgBasePath + "/radio-profiles"
		if RadioProfilesEndpoint != expectedEndpoint {
			t.Errorf("RadioProfilesEndpoint mismatch: expected %s, got %s", expectedEndpoint, RadioProfilesEndpoint)
		}
	})
}
