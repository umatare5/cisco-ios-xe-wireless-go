// Package dot15 provides 802.15.4 configuration test functionality for the Cisco Wireless Network Controller API.
package dot15

import (
	"context"
	"testing"
	"time"

	testutils "github.com/umatare5/cisco-ios-xe-wireless-go/internal/tests"
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
		testutils.TestJSONUnmarshal(t, sampleJSON, &response, "Dot15CfgResponse")
	})

	// Test Dot15GlobalConfigResponse structure
	t.Run("Dot15GlobalConfigResponse", func(t *testing.T) {
		sampleJSON := `{
			"Cisco-IOS-XE-wireless-dot15-cfg:dot15-global-config": {}
		}`

		var response Dot15GlobalConfigResponse
		testutils.TestJSONUnmarshal(t, sampleJSON, &response, "Dot15GlobalConfigResponse")
	})
}

// =============================================================================
// 2. ERROR HANDLING TESTS
// =============================================================================

func TestDot15NilClientHandling(t *testing.T) {
	ctx, cancel := testutils.CreateStandardTestContext()
	defer cancel()

	// Test GetDot15Cfg with nil client
	t.Run("GetDot15Cfg with nil client", func(t *testing.T) {
		_, err := GetDot15Cfg(nil, ctx)
		if err == nil {
			t.Error("Expected error when client is nil")
		}
		if err.Error() != "invalid client configuration: client cannot be nil" {
			t.Errorf("Expected 'client is nil', got %v", err)
		}
	})

	// Test GetDot15GlobalConfig with nil client
	t.Run("GetDot15GlobalConfig with nil client", func(t *testing.T) {
		_, err := GetDot15GlobalConfig(nil, ctx)
		if err == nil {
			t.Error("Expected error when client is nil")
		}
		if err.Error() != "invalid client configuration: client cannot be nil" {
			t.Errorf("Expected 'client is nil', got %v", err)
		}
	})
}

func TestDot15ContextHandling(t *testing.T) {
	client := testutils.GetTestClient(t)

	// Test with cancelled context
	cancelledCtx, cancel := context.WithCancel(context.Background())
	cancel() // Cancel immediately

	t.Run("GetDot15Cfg with cancelled context", func(t *testing.T) {
		_, err := GetDot15Cfg(client, cancelledCtx)
		if err == nil {
			t.Error("Expected error with cancelled context")
		}
	})

	// Test with timeout context
	timeoutCtx, cancel := context.WithTimeout(context.Background(), 1*time.Nanosecond)
	defer cancel()
	time.Sleep(10 * time.Millisecond) // Ensure context times out

	t.Run("GetDot15GlobalConfig with timeout context", func(t *testing.T) {
		_, err := GetDot15GlobalConfig(client, timeoutCtx)
		if err == nil {
			t.Error("Expected error with timeout context")
		}
	})
}

// =============================================================================
// 3. INTEGRATION TESTS (Actual API Calls to Live Controller)
// =============================================================================

// TestDot15ConfigurationFunctions tests all 802.15 configuration functions with a live controller
func TestDot15ConfigurationFunctions(t *testing.T) {
	client := testutils.CreateTestClientFromEnv(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test GetDot15Cfg function
	t.Run("GetDot15Cfg", func(t *testing.T) {
		data, err := GetDot15Cfg(client, ctx)
		if err != nil {
			t.Fatalf("GetDot15Cfg failed: %v", err)
		}

		// Validate basic structure
		if data == nil {
			t.Fatal("GetDot15Cfg returned nil data")
		}

		// Save test data for analysis
		if err := testutils.SaveTestDataToFile("dot15_cfg_data.json", data); err != nil {
			t.Logf("Warning: Could not save test data: %v", err)
		} else {
			t.Logf("Dot15 config data saved to test_data/dot15_cfg_data.json")
		}
	})

	// Test GetDot15GlobalConfig function
	t.Run("GetDot15GlobalConfig", func(t *testing.T) {
		data, err := GetDot15GlobalConfig(client, ctx)
		if err != nil {
			t.Fatalf("GetDot15GlobalConfig failed: %v", err)
		}

		// Validate basic structure
		if data == nil {
			t.Fatal("GetDot15GlobalConfig returned nil data")
		}

		// Save test data for analysis
		if err := testutils.SaveTestDataToFile("dot15_global_config_data.json", data); err != nil {
			t.Logf("Warning: Could not save test data: %v", err)
		} else {
			t.Logf("Dot15 global config data saved to test_data/dot15_global_config_data.json")
		}
	})
}

// TestDot15ConfigurationEndpoints validates 802.15 configuration endpoint constants
func TestDot15ConfigurationEndpoints(t *testing.T) {
	// Test base path validation
	t.Run("Validate_Dot15CfgBasePath", func(t *testing.T) {
		expectedBasePath := "Cisco-IOS-XE-wireless-dot15-cfg:dot15-cfg-data"
		testutils.EndpointValidationTest(t, Dot15CfgBasePath, expectedBasePath)
	})

	// Test endpoint validation
	t.Run("Validate_Dot15CfgEndpoint", func(t *testing.T) {
		testutils.EndpointValidationTest(t, Dot15CfgEndpoint, Dot15CfgBasePath)
	})
}

// =============================================================================
// 4. DATA STRUCTURE VALIDATION TESTS
// =============================================================================

func TestDot15DataStructureValidation(t *testing.T) {
	// Test data structure fields
	t.Run("Dot15CfgResponse structure validation", func(t *testing.T) {
		response := Dot15CfgResponse{}
		expectedFields := []string{"CiscoIOSXEWirelessDot15CfgDot15CfgData"}
		testutils.DataStructureValidationTest(t, response, expectedFields)
	})

	t.Run("Dot15GlobalConfigResponse structure validation", func(t *testing.T) {
		response := Dot15GlobalConfigResponse{}
		expectedFields := []string{"Dot15GlobalConfig"}
		testutils.DataStructureValidationTest(t, response, expectedFields)
	})
}
