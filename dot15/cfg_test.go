// Package dot15 provides 802.15.4 configuration test functionality for the Cisco Wireless Network Controller API.
package dot15

import (
	"encoding/json"
	"testing"
)

// =============================================================================
// 1. UNIT TESTS (Structure/Type Validation & JSON Serialization/Deserialization)
// =============================================================================

func TestDot15CfgDataStructures(t *testing.T) {
	// Test Dot15CfgResponse structure
	t.Run("Dot15CfgResponse", func(t *testing.T) {
		sampleJSON := `{
			"Cisco-IOS-XE-wireless-dot15-cfg:dot15-cfg-data": {
				"dot15-global-config": {}
			}
		}`

		var response Dot15CfgResponse
		err := json.Unmarshal([]byte(sampleJSON), &response)
		if err != nil {
			t.Fatalf("Failed to unmarshal Dot15CfgResponse: %v", err)
		}

		// Since dot15-global-config is an empty struct, just verify it unmarshals successfully
	})

	// Test Dot15GlobalConfigResponse structure
	t.Run("Dot15GlobalConfigResponse", func(t *testing.T) {
		sampleJSON := `{
			"Cisco-IOS-XE-wireless-dot15-cfg:dot15-global-config": {}
		}`

		var response Dot15GlobalConfigResponse
		err := json.Unmarshal([]byte(sampleJSON), &response)
		if err != nil {
			t.Fatalf("Failed to unmarshal Dot15GlobalConfigResponse: %v", err)
		}

		// Since dot15-global-config is an empty struct, just verify it unmarshals successfully
	})
}
