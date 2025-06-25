// Package mesh provides mesh networking configuration test functionality for the Cisco Wireless Network Controller API.
package mesh

import (
	"encoding/json"
	"testing"
)

// =============================================================================
// 1. UNIT TESTS (Structure/Type Validation & JSON Serialization/Deserialization)
// =============================================================================

func TestMeshCfgDataStructures(t *testing.T) {
	// Test MeshCfgResponse structure
	t.Run("MeshCfgResponse", func(t *testing.T) {
		sampleJSON := `{
			"Cisco-IOS-XE-wireless-mesh-cfg:mesh-cfg-data": {
				"mesh": {},
				"mesh-profiles": {
					"mesh-profile": [
						{
							"profile-name": "outdoor-mesh",
							"description": "Outdoor mesh profile for campus deployment"
						},
						{
							"profile-name": "indoor-mesh"
						}
					]
				}
			}
		}`

		var response MeshCfgResponse
		err := json.Unmarshal([]byte(sampleJSON), &response)
		if err != nil {
			t.Fatalf("Failed to unmarshal MeshCfgResponse: %v", err)
		}

		if len(response.CiscoIOSXEWirelessMeshCfgMeshCfgData.MeshProfiles.MeshProfile) != 2 {
			t.Errorf("Expected 2 mesh profiles, got %d",
				len(response.CiscoIOSXEWirelessMeshCfgMeshCfgData.MeshProfiles.MeshProfile))
		}

		profile := response.CiscoIOSXEWirelessMeshCfgMeshCfgData.MeshProfiles.MeshProfile[0]
		if profile.ProfileName != "outdoor-mesh" {
			t.Errorf("Expected profile name 'outdoor-mesh', got '%s'", profile.ProfileName)
		}

		if profile.Description != "Outdoor mesh profile for campus deployment" {
			t.Errorf("Expected description 'Outdoor mesh profile for campus deployment', got '%s'", profile.Description)
		}
	})

	// Test MeshResponse structure
	t.Run("MeshResponse", func(t *testing.T) {
		sampleJSON := `{
			"Cisco-IOS-XE-wireless-mesh-cfg:mesh": {}
		}`

		var response MeshResponse
		err := json.Unmarshal([]byte(sampleJSON), &response)
		if err != nil {
			t.Fatalf("Failed to unmarshal MeshResponse: %v", err)
		}

		// Since mesh is an empty struct, just verify it unmarshals successfully
	})

	// Test MeshProfilesResponse structure
	t.Run("MeshProfilesResponse", func(t *testing.T) {
		sampleJSON := `{
			"Cisco-IOS-XE-wireless-mesh-cfg:mesh-profiles": {
				"mesh-profile": [
					{
						"profile-name": "warehouse-mesh",
						"description": "High-density warehouse mesh configuration"
					},
					{
						"profile-name": "bridge-mesh",
						"description": "Point-to-point bridge mesh profile"
					}
				]
			}
		}`

		var response MeshProfilesResponse
		err := json.Unmarshal([]byte(sampleJSON), &response)
		if err != nil {
			t.Fatalf("Failed to unmarshal MeshProfilesResponse: %v", err)
		}

		if len(response.MeshProfiles.MeshProfile) != 2 {
			t.Errorf("Expected 2 mesh profiles, got %d", len(response.MeshProfiles.MeshProfile))
		}

		profile := response.MeshProfiles.MeshProfile[0]
		if profile.ProfileName != "warehouse-mesh" {
			t.Errorf("Expected profile name 'warehouse-mesh', got '%s'", profile.ProfileName)
		}

		if profile.Description != "High-density warehouse mesh configuration" {
			t.Errorf("Expected description 'High-density warehouse mesh configuration', got '%s'", profile.Description)
		}
	})

	// Test MeshProfile structure
	t.Run("MeshProfile", func(t *testing.T) {
		sampleJSON := `{
			"profile-name": "retail-mesh",
			"description": "Retail store mesh configuration"
		}`

		var profile MeshProfile
		err := json.Unmarshal([]byte(sampleJSON), &profile)
		if err != nil {
			t.Fatalf("Failed to unmarshal MeshProfile: %v", err)
		}

		if profile.ProfileName != "retail-mesh" {
			t.Errorf("Expected profile name 'retail-mesh', got '%s'", profile.ProfileName)
		}

		if profile.Description != "Retail store mesh configuration" {
			t.Errorf("Expected description 'Retail store mesh configuration', got '%s'", profile.Description)
		}
	})
}
