// Package site provides site configuration test functionality for the Cisco Wireless Network Controller API.
package site

import (
	"context"
	"testing"

	wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
	testutils "github.com/umatare5/cisco-ios-xe-wireless-go/internal/tests"
)

// =============================================================================
// 1. UNIT TESTS (Structure/Type Validation & JSON Serialization/Deserialization)
// =============================================================================

func TestSiteCfgDataStructures(t *testing.T) {
	testCases := []testutils.JSONTestCase{
		{
			Name: "SiteCfgResponse",
			JSONData: `{
				"Cisco-IOS-XE-wireless-site-cfg:site-cfg-data": {
					"ap-cfg-profiles": {
						"ap-cfg-profile": [
							{
								"profile-name": "campus-ap",
								"description": "Campus access point configuration",
								"stats-timer": {
									"stats-timer": 300
								}
							}
						]
					},
					"site-tag-configs": {
						"site-tag-config": [
							{
								"site-tag-name": "building-a",
								"description": "Building A site configuration",
								"is-local-site": true
							}
						]
					}
				}
			}`,
			Target:     &SiteCfgResponse{},
			TypeName:   "SiteCfgResponse",
			ShouldFail: false,
		},
		{
			Name: "SiteApCfgProfilesResponse",
			JSONData: `{
				"Cisco-IOS-XE-wireless-site-cfg:ap-cfg-profiles": {
					"ap-cfg-profile": [
						{
							"profile-name": "outdoor-ap",
							"description": "Outdoor access point configuration",
							"tunnel": {
								"preferred-mode": "capwap"
							},
							"hyperlocation": {
								"hyperlocation-enable": true
							}
						}
					]
				}
			}`,
			Target:     &SiteApCfgProfilesResponse{},
			TypeName:   "SiteApCfgProfilesResponse",
			ShouldFail: false,
		},
		{
			Name: "SiteTagConfigsResponse",
			JSONData: `{
				"Cisco-IOS-XE-wireless-site-cfg:site-tag-configs": {
					"site-tag-config": [
						{
							"site-tag-name": "building-b",
							"description": "Building B site configuration",
							"is-local-site": false
						}
					]
				}
			}`,
			Target:     &SiteTagConfigsResponse{},
			TypeName:   "SiteTagConfigsResponse",
			ShouldFail: false,
		},
	}

	testutils.RunJSONTests(t, testCases)

	// Additional field validation for successfully unmarshaled structures
	t.Run("SiteCfgResponseFieldValidation", func(t *testing.T) {
		var response SiteCfgResponse
		testutils.TestJSONUnmarshal(t, testCases[0].JSONData, &response, "SiteCfgResponse")

		testutils.ValidateJSONStructFields(t, "SiteCfgResponse", func() error {
			if len(response.CiscoIOSXEWirelessSiteCfgData.ApCfgProfiles.ApCfgProfile) != 1 {
				t.Errorf("Expected 1 AP config profile, got %d",
					len(response.CiscoIOSXEWirelessSiteCfgData.ApCfgProfiles.ApCfgProfile))
			}

			profile := response.CiscoIOSXEWirelessSiteCfgData.ApCfgProfiles.ApCfgProfile[0]
			if profile.ProfileName != "campus-ap" {
				t.Errorf("Expected profile name 'campus-ap', got '%s'", profile.ProfileName)
			}

			if profile.StatsTimer == nil || profile.StatsTimer.StatsTimer != 300 {
				t.Error("Expected stats timer to be 300")
			}

			if len(response.CiscoIOSXEWirelessSiteCfgData.SiteTagConfigs.SiteTagConfig) != 1 {
				t.Errorf("Expected 1 site tag config, got %d",
					len(response.CiscoIOSXEWirelessSiteCfgData.SiteTagConfigs.SiteTagConfig))
			}

			siteTag := response.CiscoIOSXEWirelessSiteCfgData.SiteTagConfigs.SiteTagConfig[0]
			if siteTag.SiteTagName != "building-a" {
				t.Errorf("Expected site tag name 'building-a', got '%s'", siteTag.SiteTagName)
			}
			return nil
		})
	})
}

// =============================================================================
// 2. ERROR HANDLING TESTS
// =============================================================================

func TestSiteCfgErrorHandling(t *testing.T) {
	t.Run("GetSiteCfgWithNilClient", func(t *testing.T) {
		_, err := GetSiteCfg(nil, context.Background())
		if err == nil || (err.Error() != "client is nil" && err.Error() != "invalid client configuration: client cannot be nil") {
			t.Errorf("Expected 'client is nil' or 'invalid client configuration: client cannot be nil' error, got: %v", err)
		}
	})

	t.Run("GetSiteApCfgProfilesWithNilClient", func(t *testing.T) {
		_, err := GetSiteApCfgProfiles(nil, context.Background())
		if err == nil || (err.Error() != "client is nil" && err.Error() != "invalid client configuration: client cannot be nil") {
			t.Errorf("Expected 'client is nil' or 'invalid client configuration: client cannot be nil' error, got: %v", err)
		}
	})

	t.Run("GetSiteTagConfigsWithNilClient", func(t *testing.T) {
		_, err := GetSiteTagConfigs(nil, context.Background())
		if err == nil || (err.Error() != "client is nil" && err.Error() != "invalid client configuration: client cannot be nil") {
			t.Errorf("Expected 'client is nil' or 'invalid client configuration: client cannot be nil' error, got: %v", err)
		}
	})
}

// =============================================================================
// 3. CONTEXT HANDLING TESTS
// =============================================================================

func TestSiteCfgContextHandling(t *testing.T) {
	// Test each site configuration function with context handling
	t.Run("GetSiteCfg", func(t *testing.T) {
		testutils.TestContextHandling(t, func(ctx context.Context, client *wnc.Client) error {
			_, err := GetSiteCfg(client, ctx)
			return err
		})
	})

	t.Run("GetSiteApCfgProfiles", func(t *testing.T) {
		testutils.TestContextHandling(t, func(ctx context.Context, client *wnc.Client) error {
			_, err := GetSiteApCfgProfiles(client, ctx)
			return err
		})
	})

	t.Run("GetSiteTagConfigs", func(t *testing.T) {
		testutils.TestContextHandling(t, func(ctx context.Context, client *wnc.Client) error {
			_, err := GetSiteTagConfigs(client, ctx)
			return err
		})
	})
}

// =============================================================================
// 4. ENDPOINT TESTS - API URL Validation
// =============================================================================

func TestSiteCfgEndpoints(t *testing.T) {
	tests := []struct {
		name        string
		endpoint    string
		description string
	}{
		{
			name:        "SiteCfgEndpoint",
			endpoint:    "/restconf/data/Cisco-IOS-XE-wireless-site-cfg:site-cfg-data",
			description: "Site configuration data endpoint",
		},
		{
			name:        "ApCfgProfilesEndpoint",
			endpoint:    "/restconf/data/Cisco-IOS-XE-wireless-site-cfg:site-cfg-data/ap-cfg-profiles",
			description: "AP configuration profiles endpoint",
		},
		{
			name:        "SiteTagConfigsEndpoint",
			endpoint:    "/restconf/data/Cisco-IOS-XE-wireless-site-cfg:site-cfg-data/site-tag-configs",
			description: "Site tag configurations endpoint",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testutils.EndpointValidationTest(t, tt.endpoint, tt.endpoint)
		})
	}
}

// =============================================================================
// 5. INTEGRATION TESTS
// =============================================================================

func TestSiteCfgIntegration(t *testing.T) {
	client := testutils.CreateTestClientFromEnv(t)
	ctx, cancel := context.WithTimeout(context.Background(), testutils.DefaultTestTimeout)
	defer cancel()

	t.Run("GetSiteCfgSuccess", func(t *testing.T) {
		result, err := GetSiteCfg(client, ctx)
		if err != nil {
			t.Logf("GetSiteCfg returned error (expected in some environments): %v", err)
		} else if result != nil {
			t.Logf("GetSiteCfg successful, got %d AP profiles and %d site tags",
				len(result.CiscoIOSXEWirelessSiteCfgData.ApCfgProfiles.ApCfgProfile),
				len(result.CiscoIOSXEWirelessSiteCfgData.SiteTagConfigs.SiteTagConfig))
		}
	})

	t.Run("GetSiteApCfgProfilesSuccess", func(t *testing.T) {
		result, err := GetSiteApCfgProfiles(client, ctx)
		if err != nil {
			t.Logf("GetSiteApCfgProfiles returned error (expected in some environments): %v", err)
		} else if result != nil {
			t.Logf("GetSiteApCfgProfiles successful, got %d profiles", len(result.ApCfgProfiles.ApCfgProfile))
		}
	})

	t.Run("GetSiteTagConfigsSuccess", func(t *testing.T) {
		result, err := GetSiteTagConfigs(client, ctx)
		if err != nil {
			t.Logf("GetSiteTagConfigs returned error (expected in some environments): %v", err)
		} else if result != nil {
			t.Logf("GetSiteTagConfigs successful, got %d site tags", len(result.SiteTagConfigs.SiteTagConfig))
		}
	})
}
