// Package location provides location configuration test functionality for the Cisco Wireless Network Controller API.
package location

import (
	"encoding/json"
	"testing"
)

// =============================================================================
// 1. UNIT TESTS (Structure/Type Validation & JSON Serialization/Deserialization)
// =============================================================================

func TestLocationCfgDataStructures(t *testing.T) {
	// Test LocationCfgResponse structure
	t.Run("LocationCfgResponse", func(t *testing.T) {
		sampleJSON := `{
			"Cisco-IOS-XE-wireless-location-cfg:location-cfg-data": {
				"nmsp-config": {}
			}
		}`

		var response LocationCfgResponse
		err := json.Unmarshal([]byte(sampleJSON), &response)
		if err != nil {
			t.Fatalf("Failed to unmarshal LocationCfgResponse: %v", err)
		}

		// Since nmsp-config is an empty struct, just verify it unmarshals successfully
	})

	// Test LocationNmspConfigResponse structure
	t.Run("LocationNmspConfigResponse", func(t *testing.T) {
		sampleJSON := `{
			"Cisco-IOS-XE-wireless-location-cfg:nmsp-config": {}
		}`

		var response LocationNmspConfigResponse
		err := json.Unmarshal([]byte(sampleJSON), &response)
		if err != nil {
			t.Fatalf("Failed to unmarshal LocationNmspConfigResponse: %v", err)
		}

		// Since nmsp-config is an empty struct, just verify it unmarshals successfully
	})
}
