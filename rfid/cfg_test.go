// Package rfid provides RFID configuration test functionality for the Cisco Wireless Network Controller API.
package rfid

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
	testutils "github.com/umatare5/cisco-ios-xe-wireless-go/internal/tests"
)

// =============================================================================
// 1. UNIT TESTS (Structure/Type Validation & JSON Serialization/Deserialization)
// =============================================================================

// TestRfidCfgDataStructures tests the basic structure of RFID configuration data types
func TestRfidCfgDataStructures(t *testing.T) {
	// Test RfidCfgResponse structure
	t.Run("RfidCfgResponse", func(t *testing.T) {
		sampleJSON := `{
			"Cisco-IOS-XE-wireless-rfid-cfg:rfid-cfg-data": {
				"rfid": {}
			}
		}`

		var response RfidCfgResponse
		err := json.Unmarshal([]byte(sampleJSON), &response)
		if err != nil {
			t.Fatalf("Failed to unmarshal RfidCfgResponse: %v", err)
		}

		// Since rfid is an empty struct, just verify it unmarshals successfully
	})

	// Test RfidResponse structure
	t.Run("RfidResponse", func(t *testing.T) {
		sampleJSON := `{
			"Cisco-IOS-XE-wireless-rfid-cfg:rfid": {}
		}`

		var response RfidResponse
		err := json.Unmarshal([]byte(sampleJSON), &response)
		if err != nil {
			t.Fatalf("Failed to unmarshal RfidResponse: %v", err)
		}

		// Since rfid is an empty struct, just verify it unmarshals successfully
	})
}

// =============================================================================
// 2. INTEGRATION TESTS (Actual API Calls to Live Controller)
// =============================================================================

// TestRfidConfigurationFunctions tests all RFID configuration functions with a live controller
func TestRfidConfigurationFunctions(t *testing.T) {
	client := testutils.CreateTestClientFromEnv(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test GetRfidCfg function
	t.Run("GetRfidCfg", func(t *testing.T) {
		data, err := GetRfidCfg(client, ctx)
		if err != nil {
			t.Fatalf("GetRfidCfg failed: %v", err)
		}

		// Validate basic structure
		if data == nil {
			t.Fatal("GetRfidCfg returned nil data")
		}

		// Save test data for analysis
		if err := testutils.SaveTestDataToFile("rfid_cfg_data.json", data); err != nil {
			t.Logf("Warning: Could not save test data: %v", err)
		} else {
			t.Logf("RFID config data saved to test_data/rfid_cfg_data.json")
		}

		// Validate endpoint was constructed correctly
		endpoint := RfidCfgEndpoint
		if endpoint == "" {
			t.Error("RfidCfgEndpoint should not be empty")
		}
		if endpoint != "Cisco-IOS-XE-wireless-rfid-cfg:rfid-cfg-data" {
			t.Errorf("RfidCfgEndpoint unexpected value: got %s", endpoint)
		}
	})
}

// TestRfidConfigurationEndpoints validates RFID configuration endpoint constants
func TestRfidConfigurationEndpoints(t *testing.T) {
	// Test base path validation
	t.Run("Validate_RfidCfgBasePath", func(t *testing.T) {
		expectedBasePath := "Cisco-IOS-XE-wireless-rfid-cfg:rfid-cfg-data"
		if RfidCfgBasePath != expectedBasePath {
			t.Errorf("RfidCfgBasePath mismatch: expected %s, got %s", expectedBasePath, RfidCfgBasePath)
		}
	})

	// Test endpoint validation
	t.Run("Validate_RfidCfgEndpoint", func(t *testing.T) {
		if RfidCfgEndpoint != RfidCfgBasePath {
			t.Errorf("RfidCfgEndpoint should equal RfidCfgBasePath: expected %s, got %s", RfidCfgBasePath, RfidCfgEndpoint)
		}
	})
}

// =============================================================================
// 3. TABLE-DRIVEN TEST PATTERNS
// =============================================================================

// Currently no table-driven tests specific to RFID configuration

// =============================================================================
// 4. FAIL-FAST ERROR DETECTION TESTS
// =============================================================================

// Currently no fail-fast error detection tests specific to RFID configuration

// =============================================================================
// 5. OTHER TESTS
// =============================================================================

// TestRfidCfgErrorHandling tests error handling for all RFID configuration functions.
func TestRfidCfgErrorHandling(t *testing.T) {
	t.Run("GetRfidCfgWithNilClient", func(t *testing.T) {
		_, err := GetRfidCfg(nil, context.Background())
		if err == nil || err.Error() != "invalid client configuration: client cannot be nil" {
			t.Errorf("Expected 'client is nil' error, got: %v", err)
		}
	})
}

// TestRfidCfgContextHandling tests context handling for all RFID configuration functions.
func TestRfidCfgContextHandling(t *testing.T) {
	t.Run("GetRfidCfgContextHandling", func(t *testing.T) {
		testutils.TestContextHandling(t, func(ctx context.Context, client *wnc.Client) error {
			_, err := GetRfidCfg(client, ctx)
			return err
		})
	})
}
