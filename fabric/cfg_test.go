// Package fabric provides fabric configuration test functionality for the Cisco Wireless Network Controller API.
package fabric

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

func TestFabricCfgDataStructures(t *testing.T) {
	// Test FabricCfgResponse structure
	t.Run("FabricCfgResponse", func(t *testing.T) {
		sampleJSON := `{
			"Cisco-IOS-XE-wireless-fabric-cfg:fabric-cfg-data": {
				"fabric": {},
				"fabric-controlplane-names": {
					"fabric-controlplane-name": [
						{
							"control-plane-name": "cp-primary",
							"description": "Primary control plane"
						},
						{
							"control-plane-name": "cp-secondary"
						}
					]
				}
			}
		}`

		var response FabricCfgResponse
		err := json.Unmarshal([]byte(sampleJSON), &response)
		if err != nil {
			t.Fatalf("Failed to unmarshal FabricCfgResponse: %v", err)
		}

		if len(response.CiscoIOSXEWirelessFabricCfgFabricCfgData.FabricControlplaneNames.FabricControlplaneName) != 2 {
			t.Errorf("Expected 2 fabric control plane names, got %d",
				len(response.CiscoIOSXEWirelessFabricCfgFabricCfgData.FabricControlplaneNames.FabricControlplaneName))
		}

		firstCP := response.CiscoIOSXEWirelessFabricCfgFabricCfgData.FabricControlplaneNames.FabricControlplaneName[0]
		if firstCP.ControlPlaneName != "cp-primary" {
			t.Errorf("Expected first control plane name 'cp-primary', got '%s'", firstCP.ControlPlaneName)
		}

		if firstCP.Description != "Primary control plane" {
			t.Errorf("Expected description 'Primary control plane', got '%s'", firstCP.Description)
		}
	})

	// Test FabricControlplaneNamesResponse structure
	t.Run("FabricControlplaneNamesResponse", func(t *testing.T) {
		sampleJSON := `{
			"Cisco-IOS-XE-wireless-fabric-cfg:fabric-controlplane-names": {
				"fabric-controlplane-name": [
					{
						"control-plane-name": "backup-cp",
						"description": "Backup control plane instance"
					}
				]
			}
		}`

		var response FabricControlplaneNamesResponse
		err := json.Unmarshal([]byte(sampleJSON), &response)
		if err != nil {
			t.Fatalf("Failed to unmarshal FabricControlplaneNamesResponse: %v", err)
		}

		if len(response.FabricControlplaneNames.FabricControlplaneName) != 1 {
			t.Errorf("Expected 1 fabric control plane name, got %d",
				len(response.FabricControlplaneNames.FabricControlplaneName))
		}

		cp := response.FabricControlplaneNames.FabricControlplaneName[0]
		if cp.ControlPlaneName != "backup-cp" {
			t.Errorf("Expected control plane name 'backup-cp', got '%s'", cp.ControlPlaneName)
		}

		if cp.Description != "Backup control plane instance" {
			t.Errorf("Expected description 'Backup control plane instance', got '%s'", cp.Description)
		}
	})

	// Test FabricControlplaneName structure
	t.Run("FabricControlplaneName", func(t *testing.T) {
		sampleJSON := `{
			"control-plane-name": "test-cp",
			"description": "Test control plane"
		}`

		var cp FabricControlplaneName
		err := json.Unmarshal([]byte(sampleJSON), &cp)
		if err != nil {
			t.Fatalf("Failed to unmarshal FabricControlplaneName: %v", err)
		}

		if cp.ControlPlaneName != "test-cp" {
			t.Errorf("Expected control plane name 'test-cp', got '%s'", cp.ControlPlaneName)
		}

		if cp.Description != "Test control plane" {
			t.Errorf("Expected description 'Test control plane', got '%s'", cp.Description)
		}
	})

	// Test FabricResponse structure
	t.Run("FabricResponse", func(t *testing.T) {
		sampleJSON := `{
			"Cisco-IOS-XE-wireless-fabric-cfg:fabric": {}
		}`

		var response FabricResponse
		err := json.Unmarshal([]byte(sampleJSON), &response)
		if err != nil {
			t.Fatalf("Failed to unmarshal FabricResponse: %v", err)
		}

		// Since fabric is an empty struct, just verify it unmarshals successfully
	})
}

// =============================================================================
// 2. INTEGRATION TESTS (Actual API Calls to Live Controller)
// =============================================================================

// TestFabricConfigurationFunctions tests all fabric configuration functions with a live controller
func TestFabricConfigurationFunctions(t *testing.T) {
	client := testutil.CreateTestClientFromEnv(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test GetFabricCfg function
	t.Run("GetFabricCfg", func(t *testing.T) {
		data, err := GetFabricCfg(client, ctx)
		if err != nil {
			t.Fatalf("GetFabricCfg failed: %v", err)
		}

		// Validate basic structure
		if data == nil {
			t.Fatal("GetFabricCfg returned nil data")
		}

		// Save test data for analysis
		if err := testutil.SaveTestDataToFile("fabric_cfg_data.json", data); err != nil {
			t.Logf("Warning: Could not save test data: %v", err)
		} else {
			t.Logf("Fabric config data saved to test_data/fabric_cfg_data.json")
		}

		// Validate endpoint was constructed correctly
		endpoint := FabricCfgEndpoint
		if endpoint == "" {
			t.Error("FabricCfgEndpoint should not be empty")
		}
		if endpoint != "/restconf/data/Cisco-IOS-XE-wireless-fabric-cfg:fabric-cfg-data" {
			t.Errorf("FabricCfgEndpoint unexpected value: got %s", endpoint)
		}
	})

	// Test GetFabricControlplaneNames function
	t.Run("GetFabricControlplaneNames", func(t *testing.T) {
		data, err := GetFabricControlplaneNames(client, ctx)
		if err != nil {
			t.Fatalf("GetFabricControlplaneNames failed: %v", err)
		}

		// Validate basic structure
		if data == nil {
			t.Fatal("GetFabricControlplaneNames returned nil data")
		}

		// Save test data for analysis
		if err := testutil.SaveTestDataToFile("fabric_controlplane_names_data.json", data); err != nil {
			t.Logf("Warning: Could not save test data: %v", err)
		} else {
			t.Logf("Fabric controlplane names data saved to test_data/fabric_controlplane_names_data.json")
		}

		// Validate endpoint was constructed correctly
		endpoint := FabricControlplaneNamesEndpoint
		if endpoint == "" {
			t.Error("FabricControlplaneNamesEndpoint should not be empty")
		}
		expectedEndpoint := "/restconf/data/Cisco-IOS-XE-wireless-fabric-cfg:fabric-cfg-data/fabric-controlplane-names"
		if endpoint != expectedEndpoint {
			t.Errorf("FabricControlplaneNamesEndpoint unexpected value: expected %s, got %s", expectedEndpoint, endpoint)
		}
	})

	// Test GetFabric function
	t.Run("GetFabric", func(t *testing.T) {
		data, err := GetFabric(client, ctx)
		if err != nil {
			t.Fatalf("GetFabric failed: %v", err)
		}

		// Validate basic structure
		if data == nil {
			t.Fatal("GetFabric returned nil data")
		}

		// Save test data for analysis
		if err := testutil.SaveTestDataToFile("fabric_data.json", data); err != nil {
			t.Logf("Warning: Could not save test data: %v", err)
		} else {
			t.Logf("Fabric data saved to test_data/fabric_data.json")
		}
	})
}

// TestFabricConfigurationEndpoints validates fabric configuration endpoint constants
func TestFabricConfigurationEndpoints(t *testing.T) {
	// Test base path validation
	t.Run("Validate_FabricCfgBasePath", func(t *testing.T) {
		expectedBasePath := "/restconf/data/Cisco-IOS-XE-wireless-fabric-cfg:fabric-cfg-data"
		if FabricCfgBasePath != expectedBasePath {
			t.Errorf("FabricCfgBasePath mismatch: expected %s, got %s", expectedBasePath, FabricCfgBasePath)
		}
	})

	// Test endpoint validation
	t.Run("Validate_FabricCfgEndpoint", func(t *testing.T) {
		if FabricCfgEndpoint != FabricCfgBasePath {
			t.Errorf("FabricCfgEndpoint should equal FabricCfgBasePath: expected %s, got %s", FabricCfgBasePath, FabricCfgEndpoint)
		}
	})

	// Test controlplane names endpoint validation
	t.Run("Validate_FabricControlplaneNamesEndpoint", func(t *testing.T) {
		expectedEndpoint := FabricCfgBasePath + "/fabric-controlplane-names"
		if FabricControlplaneNamesEndpoint != expectedEndpoint {
			t.Errorf("FabricControlplaneNamesEndpoint mismatch: expected %s, got %s", expectedEndpoint, FabricControlplaneNamesEndpoint)
		}
	})
}
