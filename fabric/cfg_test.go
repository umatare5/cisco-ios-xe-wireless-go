// Package fabric provides fabric configuration test functionality for the Cisco Wireless Network Controller API.
package fabric

import (
	"encoding/json"
	"testing"
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
