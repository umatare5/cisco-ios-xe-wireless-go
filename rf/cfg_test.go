// Package rf provides radio frequency configuration test functionality for the Cisco Wireless Network Controller API.
package rf

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

func TestRfCfgDataStructures(t *testing.T) {
	testCases := []testutils.JSONTestCase{
		{
			Name: "RfCfgResponse",
			JSONData: `{
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
						"atf-policy": [
							{
								"policy-id": 1,
								"atfpolicy-name": "default"
							}
						]
					},
					"rf-tags": {
						"rf-tag": [
							{
								"tag-name": "default-rf-tag",
								"description": "Default RF tag"
							}
						]
					},
					"rf-profiles": {
						"rf-profile": [
							{
								"name": "default-rf-profile",
								"description": "Default RF profile",
								"status": true,
								"band": "2.4GHz",
								"data-rate-6m": "mandatory",
								"data-rate-12m": "supported",
								"data-rate-24m": "supported"
							}
						]
					},
					"rf-profile-default-entries": {
						"rf-profile-default-entry": [
							{
								"band": "2.4GHz",
								"name": "default-rf-profile-2g",
								"description": "Default 2.4GHz RF profile",
								"data-rate-6m": "mandatory",
								"data-rate-12m": "supported",
								"data-rate-24m": "supported",
								"rf-mcs-default-entries": {
									"rf-mcs-default-entry": []
								}
							}
						]
					}
				}
			}`,
			Target:     &RfCfgResponse{},
			TypeName:   "RfCfgResponse",
			ShouldFail: false,
		},
		{
			Name: "MultiBssidProfilesResponse",
			JSONData: `{
				"Cisco-IOS-XE-wireless-rf-cfg:multi-bssid-profiles": {
					"multi-bssid-profile": [
						{
							"profile-name": "test-profile",
							"description": "Test multi-BSSID profile"
						}
					]
				}
			}`,
			Target:     &MultiBssidProfilesResponse{},
			TypeName:   "MultiBssidProfilesResponse",
			ShouldFail: false,
		},
		{
			Name: "AtfPoliciesResponse",
			JSONData: `{
				"Cisco-IOS-XE-wireless-rf-cfg:atf-policies": {
					"atf-policy": [
						{
							"policy-id": 2,
							"atfpolicy-name": "priority-policy"
						}
					]
				}
			}`,
			Target:     &AtfPoliciesResponse{},
			TypeName:   "AtfPoliciesResponse",
			ShouldFail: false,
		},
		{
			Name: "RfTagsResponse",
			JSONData: `{
				"Cisco-IOS-XE-wireless-rf-cfg:rf-tags": {
					"rf-tag": [
						{
							"tag-name": "test-rf-tag",
							"description": "Test RF tag",
							"rf-tag-radio-profiles": {
								"rf-tag-radio-profile": [
									{
										"slot-id": "0",
										"band-id": "2.4GHz"
									}
								]
							}
						}
					]
				}
			}`,
			Target:     &RfTagsResponse{},
			TypeName:   "RfTagsResponse",
			ShouldFail: false,
		},
		{
			Name: "RfProfilesResponse",
			JSONData: `{
				"Cisco-IOS-XE-wireless-rf-cfg:rf-profiles": {
					"rf-profile": [
						{
							"name": "high-power",
							"description": "High power RF profile",
							"status": true,
							"band": "5GHz",
							"data-rate-6m": "mandatory",
							"data-rate-12m": "supported",
							"data-rate-24m": "supported"
						}
					]
				}
			}`,
			Target:     &RfProfilesResponse{},
			TypeName:   "RfProfilesResponse",
			ShouldFail: false,
		},
	}

	testutils.RunJSONTests(t, testCases)

	// Additional field validation for successfully unmarshaled structures
	t.Run("RfCfgResponseFieldValidation", func(t *testing.T) {
		var response RfCfgResponse
		testutils.TestJSONUnmarshal(t, testCases[0].JSONData, &response, "RfCfgResponse")

		testutils.ValidateJSONStructFields(t, "RfCfgResponse", func() error {
			if len(response.CiscoIOSXEWirelessRfCfgRfCfgData.MultiBssidProfiles.MultiBssidProfile) != 1 {
				t.Errorf("Expected 1 multi-BSSID profile, got %d",
					len(response.CiscoIOSXEWirelessRfCfgRfCfgData.MultiBssidProfiles.MultiBssidProfile))
			}
			if len(response.CiscoIOSXEWirelessRfCfgRfCfgData.AtfPolicies.AtfPolicy) != 1 {
				t.Errorf("Expected 1 ATF policy, got %d",
					len(response.CiscoIOSXEWirelessRfCfgRfCfgData.AtfPolicies.AtfPolicy))
			}
			if len(response.CiscoIOSXEWirelessRfCfgRfCfgData.RfTags.RfTag) != 1 {
				t.Errorf("Expected 1 RF tag, got %d",
					len(response.CiscoIOSXEWirelessRfCfgRfCfgData.RfTags.RfTag))
			}
			return nil
		})
	})
}

const (
	TestDataDir = "test_data"
)

// getTestClient creates a test client using environment variables
func getTestClient(t *testing.T) *wnc.Client {
	return testutils.CreateTestClientFromEnv(t)
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
		if err := testutils.SaveTestDataToFile("rf_cfg_test_data_collected.json", rfCfgTestDataCollector.Data); err != nil {
			t.Logf("Warning: Could not save test data: %v", err)
		} else {
			t.Logf("Test data saved to %s/rf_cfg_test_data_collected.json", TestDataDir)
		}
	}
}

// TestRfCfgErrorHandling tests error handling for all RF configuration functions.
func TestRfCfgErrorHandling(t *testing.T) {
	t.Run("GetRfCfgWithNilClient", func(t *testing.T) {
		_, err := GetRfCfg(nil, context.Background())
		if err == nil || err.Error() != "invalid client configuration: client cannot be nil" {
			t.Errorf("Expected 'client is nil' error, got: %v", err)
		}
	})

	t.Run("GetRfMultiBssidProfilesWithNilClient", func(t *testing.T) {
		_, err := GetRfMultiBssidProfiles(nil, context.Background())
		if err == nil || err.Error() != "invalid client configuration: client cannot be nil" {
			t.Errorf("Expected 'client is nil' error, got: %v", err)
		}
	})

	t.Run("GetRfAtfPoliciesWithNilClient", func(t *testing.T) {
		_, err := GetRfAtfPolicies(nil, context.Background())
		if err == nil || err.Error() != "invalid client configuration: client cannot be nil" {
			t.Errorf("Expected 'client is nil' error, got: %v", err)
		}
	})

	t.Run("GetRfTagsWithNilClient", func(t *testing.T) {
		_, err := GetRfTags(nil, context.Background())
		if err == nil || err.Error() != "invalid client configuration: client cannot be nil" {
			t.Errorf("Expected 'client is nil' error, got: %v", err)
		}
	})

	t.Run("GetRfProfilesWithNilClient", func(t *testing.T) {
		_, err := GetRfProfiles(nil, context.Background())
		if err == nil || err.Error() != "invalid client configuration: client cannot be nil" {
			t.Errorf("Expected 'client is nil' error, got: %v", err)
		}
	})

	t.Run("GetRfProfileDefaultEntriesWithNilClient", func(t *testing.T) {
		_, err := GetRfProfileDefaultEntries(nil, context.Background())
		if err == nil || err.Error() != "invalid client configuration: client cannot be nil" {
			t.Errorf("Expected 'client is nil' error, got: %v", err)
		}
	})
}

// TestRfCfgContextHandling tests context handling for all RF configuration functions.
func TestRfCfgContextHandling(t *testing.T) {
	t.Run("GetRfCfgContextHandling", func(t *testing.T) {
		testutils.TestContextHandling(t, func(ctx context.Context, client *wnc.Client) error {
			_, err := GetRfCfg(client, ctx)
			return err
		})
	})

	t.Run("GetRfMultiBssidProfilesContextHandling", func(t *testing.T) {
		testutils.TestContextHandling(t, func(ctx context.Context, client *wnc.Client) error {
			_, err := GetRfMultiBssidProfiles(client, ctx)
			return err
		})
	})

	t.Run("GetRfAtfPoliciesContextHandling", func(t *testing.T) {
		testutils.TestContextHandling(t, func(ctx context.Context, client *wnc.Client) error {
			_, err := GetRfAtfPolicies(client, ctx)
			return err
		})
	})

	t.Run("GetRfTagsContextHandling", func(t *testing.T) {
		testutils.TestContextHandling(t, func(ctx context.Context, client *wnc.Client) error {
			_, err := GetRfTags(client, ctx)
			return err
		})
	})

	t.Run("GetRfProfilesContextHandling", func(t *testing.T) {
		testutils.TestContextHandling(t, func(ctx context.Context, client *wnc.Client) error {
			_, err := GetRfProfiles(client, ctx)
			return err
		})
	})

	t.Run("GetRfProfileDefaultEntriesContextHandling", func(t *testing.T) {
		testutils.TestContextHandling(t, func(ctx context.Context, client *wnc.Client) error {
			_, err := GetRfProfileDefaultEntries(client, ctx)
			return err
		})
	})
}
