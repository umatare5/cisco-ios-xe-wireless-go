// Package site provides site configuration test functionality for the Cisco Wireless Network Controller API.
package site

import (
	"encoding/json"
	"testing"
)

// =============================================================================
// 1. UNIT TESTS (Structure/Type Validation & JSON Serialization/Deserialization)
// =============================================================================

func TestSiteCfgDataStructures(t *testing.T) {
	// Test SiteCfgResponse structure
	t.Run("SiteCfgResponse", func(t *testing.T) {
		sampleJSON := `{
			"Cisco-IOS-XE-wireless-site-cfg:site-cfg-data": {
				"ap-cfg-profiles": [
					{
						"profile-name": "campus-ap",
						"description": "Campus access point profile",
						"stats-timer": {
							"stats-timer": 300
						},
						"user-mgmt": {
							"username": "admin",
							"password": "password123",
							"password-type": "clear"
						}
					}
				],
				"site-tag-configs": [
					{
						"site-tag-name": "building-a",
						"description": "Building A site configuration",
						"is-local-site": true
					}
				]
			}
		}`

		var response SiteCfgResponse
		err := json.Unmarshal([]byte(sampleJSON), &response)
		if err != nil {
			t.Fatalf("Failed to unmarshal SiteCfgResponse: %v", err)
		}

		if len(response.CiscoIOSXEWirelessSiteCfgData.ApCfgProfiles) != 1 {
			t.Errorf("Expected 1 AP config profile, got %d",
				len(response.CiscoIOSXEWirelessSiteCfgData.ApCfgProfiles))
		}

		profile := response.CiscoIOSXEWirelessSiteCfgData.ApCfgProfiles[0]
		if profile.ProfileName != "campus-ap" {
			t.Errorf("Expected profile name 'campus-ap', got '%s'", profile.ProfileName)
		}

		if profile.StatsTimer == nil || profile.StatsTimer.StatsTimer != 300 {
			t.Error("Expected stats timer to be 300")
		}

		if len(response.CiscoIOSXEWirelessSiteCfgData.SiteTagConfigs) != 1 {
			t.Errorf("Expected 1 site tag config, got %d",
				len(response.CiscoIOSXEWirelessSiteCfgData.SiteTagConfigs))
		}

		siteTag := response.CiscoIOSXEWirelessSiteCfgData.SiteTagConfigs[0]
		if siteTag.SiteTagName != "building-a" {
			t.Errorf("Expected site tag name 'building-a', got '%s'", siteTag.SiteTagName)
		}
	})

	// Test SiteApCfgProfilesResponse structure
	t.Run("SiteApCfgProfilesResponse", func(t *testing.T) {
		sampleJSON := `{
			"Cisco-IOS-XE-wireless-site-cfg:ap-cfg-profiles": [
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
		}`

		var response SiteApCfgProfilesResponse
		err := json.Unmarshal([]byte(sampleJSON), &response)
		if err != nil {
			t.Fatalf("Failed to unmarshal SiteApCfgProfilesResponse: %v", err)
		}

		if len(response.ApCfgProfiles) != 1 {
			t.Errorf("Expected 1 AP config profile, got %d", len(response.ApCfgProfiles))
		}

		profile := response.ApCfgProfiles[0]
		if profile.ProfileName != "outdoor-ap" {
			t.Errorf("Expected profile name 'outdoor-ap', got '%s'", profile.ProfileName)
		}

		if profile.Tunnel == nil || profile.Tunnel.PreferredMode != "capwap" {
			t.Error("Expected tunnel preferred mode to be 'capwap'")
		}

		if profile.Hyperlocation == nil || !profile.Hyperlocation.HyperlocationEnable {
			t.Error("Expected hyperlocation to be enabled")
		}
	})

	// Test SiteTagConfigsResponse structure
	t.Run("SiteTagConfigsResponse", func(t *testing.T) {
		sampleJSON := `{
			"Cisco-IOS-XE-wireless-site-cfg:site-tag-configs": [
				{
					"site-tag-name": "warehouse",
					"description": "Warehouse site configuration",
					"flex-profile": "warehouse-flex",
					"ap-join-profile": "warehouse-join",
					"is-local-site": false
				}
			]
		}`

		var response SiteTagConfigsResponse
		err := json.Unmarshal([]byte(sampleJSON), &response)
		if err != nil {
			t.Fatalf("Failed to unmarshal SiteTagConfigsResponse: %v", err)
		}

		if len(response.SiteTagConfigs) != 1 {
			t.Errorf("Expected 1 site tag config, got %d", len(response.SiteTagConfigs))
		}

		siteTag := response.SiteTagConfigs[0]
		if siteTag.SiteTagName != "warehouse" {
			t.Errorf("Expected site tag name 'warehouse', got '%s'", siteTag.SiteTagName)
		}

		if siteTag.FlexProfile != "warehouse-flex" {
			t.Errorf("Expected flex profile 'warehouse-flex', got '%s'", siteTag.FlexProfile)
		}

		if siteTag.IsLocalSite {
			t.Error("Expected is-local-site to be false")
		}
	})

	// Test ApCfgProfile structure
	t.Run("ApCfgProfile", func(t *testing.T) {
		sampleJSON := `{
			"profile-name": "retail-ap",
			"description": "Retail store AP profile",
			"rogue-detection": {
				"ap-rogue-detection-min-rssi": -70
			},
			"reporting-interval": {
				"radio-24ghz": 60,
				"radio-5ghz": 60
			},
			"device-mgmt": {
				"ssh": true
			}
		}`

		var profile ApCfgProfile
		err := json.Unmarshal([]byte(sampleJSON), &profile)
		if err != nil {
			t.Fatalf("Failed to unmarshal ApCfgProfile: %v", err)
		}

		if profile.ProfileName != "retail-ap" {
			t.Errorf("Expected profile name 'retail-ap', got '%s'", profile.ProfileName)
		}

		if profile.RogueDetection == nil || profile.RogueDetection.ApRogueDetectionMinRssi != -70 {
			t.Error("Expected rogue detection min RSSI to be -70")
		}

		if profile.ReportingInterval == nil || profile.ReportingInterval.Radio24Ghz != 60 {
			t.Error("Expected 2.4GHz reporting interval to be 60")
		}

		if profile.DeviceMgmt == nil || !profile.DeviceMgmt.SSH {
			t.Error("Expected SSH to be enabled in device management")
		}
	})

	// Test SiteTagConfig structure
	t.Run("SiteTagConfig", func(t *testing.T) {
		sampleJSON := `{
			"site-tag-name": "main-office",
			"description": "Main office site tag",
			"flex-profile": "office-flex",
			"ap-join-profile": "office-join",
			"is-local-site": true
		}`

		var siteTag SiteTagConfig
		err := json.Unmarshal([]byte(sampleJSON), &siteTag)
		if err != nil {
			t.Fatalf("Failed to unmarshal SiteTagConfig: %v", err)
		}

		if siteTag.SiteTagName != "main-office" {
			t.Errorf("Expected site tag name 'main-office', got '%s'", siteTag.SiteTagName)
		}

		if siteTag.Description != "Main office site tag" {
			t.Errorf("Expected description 'Main office site tag', got '%s'", siteTag.Description)
		}

		if !siteTag.IsLocalSite {
			t.Error("Expected is-local-site to be true")
		}
	})
}
