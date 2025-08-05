// Package mesh provides mesh networking configuration test functionality for the Cisco Wireless Network Controller API.
package mesh

import (
	"context"
	"encoding/json"
	"strings"
	"testing"
	"time"

	wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
	testutils "github.com/umatare5/cisco-ios-xe-wireless-go/internal/tests"
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

// =============================================================================
// 2. INTEGRATION TESTS (Live API Endpoint Testing)
// =============================================================================

func TestMeshConfigurationFunctions(t *testing.T) {
	client := testutils.CreateTestClientFromEnv(t)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Create a comprehensive test data collection
	collector := testutils.NewTestDataCollector()
	endpointMapping := map[string]string{
		"MeshCfgEndpoint":      "/restconf/data/Cisco-IOS-XE-wireless-mesh-cfg:mesh-cfg-data",
		"MeshMeshEndpoint":     "/restconf/data/Cisco-IOS-XE-wireless-mesh-cfg:mesh-cfg-data/mesh",
		"MeshProfilesEndpoint": "/restconf/data/Cisco-IOS-XE-wireless-mesh-cfg:mesh-cfg-data/mesh-profiles",
	}

	t.Run("GetMeshCfg", func(t *testing.T) {
		result, err := GetMeshCfg(client, ctx)
		testutils.CollectTestResult(collector, "GetMeshCfg", endpointMapping["MeshCfgEndpoint"], result, err)
		if err != nil {
			t.Logf("GetMeshCfg failed: %v", err)
		}
	})

	t.Run("GetMesh", func(t *testing.T) {
		result, err := GetMesh(client, ctx)
		testutils.CollectTestResult(collector, "GetMesh", endpointMapping["MeshMeshEndpoint"], result, err)
		if err != nil {
			t.Logf("GetMesh failed: %v", err)
		}
	})

	t.Run("GetMeshProfiles", func(t *testing.T) {
		result, err := GetMeshProfiles(client, ctx)
		testutils.CollectTestResult(collector, "GetMeshProfiles", endpointMapping["MeshProfilesEndpoint"], result, err)
		if err != nil {
			t.Logf("GetMeshProfiles failed: %v", err)
		}
	})

	// Save collected test data to JSON file
	testutils.SaveCollectedTestData(t, collector, "mesh_cfg_test_data_collected.json")
}

// TestMeshConfigurationEndpoints tests endpoint validation
func TestMeshConfigurationEndpoints(t *testing.T) {
	t.Run("Validate_MeshCfgBasePath", func(t *testing.T) {
		if MeshCfgBasePath == "" {
			t.Error("MeshCfgBasePath is empty")
		}
		if !strings.HasPrefix(MeshCfgBasePath, "/restconf/data/") {
			t.Errorf("MeshCfgBasePath should start with '/restconf/data/', got: %s", MeshCfgBasePath)
		}
	})

	t.Run("Validate_MeshCfgEndpoint", func(t *testing.T) {
		if MeshCfgEndpoint == "" {
			t.Error("MeshCfgEndpoint is empty")
		}
		if MeshCfgEndpoint != MeshCfgBasePath {
			t.Errorf("MeshCfgEndpoint should equal MeshCfgBasePath, got: %s", MeshCfgEndpoint)
		}
	})

	t.Run("Validate_MeshMeshEndpoint", func(t *testing.T) {
		if MeshMeshEndpoint == "" {
			t.Error("MeshMeshEndpoint is empty")
		}
		if !strings.HasSuffix(MeshMeshEndpoint, "/mesh") {
			t.Errorf("MeshMeshEndpoint should end with '/mesh', got: %s", MeshMeshEndpoint)
		}
	})

	t.Run("Validate_MeshProfilesEndpoint", func(t *testing.T) {
		if MeshProfilesEndpoint == "" {
			t.Error("MeshProfilesEndpoint is empty")
		}
		if !strings.HasSuffix(MeshProfilesEndpoint, "/mesh-profiles") {
			t.Errorf("MeshProfilesEndpoint should end with '/mesh-profiles', got: %s", MeshProfilesEndpoint)
		}
	})
}

// =============================================================================
// 4. ERROR HANDLING TESTS
// =============================================================================

func TestMeshCfgErrorHandling(t *testing.T) {
	ctx := context.Background()

	t.Run("GetMeshCfg_NilClient", func(t *testing.T) {
		result, err := GetMeshCfg(nil, ctx)
		if err == nil {
			t.Error("Expected error for nil client")
		}
		if result != nil {
			t.Error("Expected nil result for nil client")
		}
		if err.Error() != "client is nil" {
			t.Errorf("Expected 'client is nil' error, got: %s", err.Error())
		}
	})

	t.Run("GetMesh_NilClient", func(t *testing.T) {
		result, err := GetMesh(nil, ctx)
		if err == nil {
			t.Error("Expected error for nil client")
		}
		if result != nil {
			t.Error("Expected nil result for nil client")
		}
		if err.Error() != "client is nil" {
			t.Errorf("Expected 'client is nil' error, got: %s", err.Error())
		}
	})

	t.Run("GetMeshProfiles_NilClient", func(t *testing.T) {
		result, err := GetMeshProfiles(nil, ctx)
		if err == nil {
			t.Error("Expected error for nil client")
		}
		if result != nil {
			t.Error("Expected nil result for nil client")
		}
		if err.Error() != "client is nil" {
			t.Errorf("Expected 'client is nil' error, got: %s", err.Error())
		}
	})
}

// =============================================================================
// 5. CONTEXT HANDLING TESTS
// =============================================================================

func TestMeshCfgContextHandling(t *testing.T) {
	testutils.TestContextHandling(t, func(ctx context.Context, client *wnc.Client) error {
		_, err := GetMeshCfg(client, ctx)
		return err
	})

	testutils.TestContextHandling(t, func(ctx context.Context, client *wnc.Client) error {
		_, err := GetMesh(client, ctx)
		return err
	})

	testutils.TestContextHandling(t, func(ctx context.Context, client *wnc.Client) error {
		_, err := GetMeshProfiles(client, ctx)
		return err
	})
}
