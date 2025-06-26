// Package hyperlocation provides hyperlocation operational data test functionality for the Cisco Wireless Network Controller API.
package hyperlocation

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil"
)

// =============================================================================
// 1. UNIT TESTS (Structure/Type Validation & JSON Serialization/Deserialization)
// =============================================================================

// getTestClient creates a test client using environment variables
func getTestClient(t *testing.T) *wnc.Client {
	return testutil.CreateTestClientFromEnv(t)
}

// TestHyperlocationOperEndpoints tests that all Hyperlocation operation endpoint constants are correctly defined
func TestHyperlocationOperEndpoints(t *testing.T) {
	expectedEndpoints := map[string]string{
		"HyperlocationOperBasePath":     "/restconf/data/Cisco-IOS-XE-wireless-hyperlocation-oper:hyperlocation-oper-data",
		"HyperlocationOperEndpoint":     "/restconf/data/Cisco-IOS-XE-wireless-hyperlocation-oper:hyperlocation-oper-data",
		"HyperlocationProfilesEndpoint": "/restconf/data/Cisco-IOS-XE-wireless-hyperlocation-oper:hyperlocation-oper-data/ewlc-hyperlocation-profile",
	}

	for name, expected := range expectedEndpoints {
		t.Run(name, func(t *testing.T) {
			switch name {
			case "HyperlocationOperBasePath":
				if HyperlocationOperBasePath != expected {
					t.Errorf(wnc.EndpointMismatchErrorTemplate, name, expected, HyperlocationOperBasePath)
				}
			case "HyperlocationOperEndpoint":
				if HyperlocationOperEndpoint != expected {
					t.Errorf(wnc.EndpointMismatchErrorTemplate, name, expected, HyperlocationOperEndpoint)
				}
			case "HyperlocationProfilesEndpoint":
				if HyperlocationProfilesEndpoint != expected {
					t.Errorf(wnc.EndpointMismatchErrorTemplate, name, expected, HyperlocationProfilesEndpoint)
				}
			}
		})
	}
}

// TestHyperlocationOperDataStructures tests the basic structure of Hyperlocation operational data types
func TestHyperlocationOperDataStructures(t *testing.T) {
	tests := []struct {
		name     string
		jsonData string
		dataType interface{}
	}{
		{
			name: "HyperlocationOperResponse",
			jsonData: `{
				"Cisco-IOS-XE-wireless-hyperlocation-oper:hyperlocation-oper-data": {
					"ewlc-hyperlocation-profile": [
						{
							"name": "hyperlocation-profile-1",
							"hyperlocation-data": {
								"hyperlocation-enable": true,
								"pak-rssi-threshold-detection": -50,
								"pak-rssi-threshold-trigger": -45,
								"pak-rssi-threshold-reset": -55
							},
							"ntp-server": "192.168.1.1",
							"status": true,
							"reason-down": ""
						}
					]
				}
			}`,
			dataType: &HyperlocationOperResponse{},
		},
		{
			name: "HyperlocationProfilesResponse",
			jsonData: `{
				"Cisco-IOS-XE-wireless-hyperlocation-oper:ewlc-hyperlocation-profile": [
					{
						"name": "hyperlocation-profile-1",
						"hyperlocation-data": {
							"hyperlocation-enable": true,
							"pak-rssi-threshold-detection": -50,
							"pak-rssi-threshold-trigger": -45,
							"pak-rssi-threshold-reset": -55
						},
						"ntp-server": "192.168.1.1",
						"status": true,
						"reason-down": ""
					}
				]
			}`,
			dataType: &HyperlocationProfilesResponse{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := json.Unmarshal([]byte(tt.jsonData), tt.dataType)
			if err != nil {
				t.Errorf("Failed to unmarshal %s: %v", tt.name, err)
			}
		})
	}
}

// =============================================================================
// 2. INTEGRATION TESTS (API Endpoint Testing with Live Data Validation)
// =============================================================================

// TestGetHyperlocationOper tests the GetHyperlocationOper method
func TestGetHyperlocationOper(t *testing.T) {
	client := getTestClient(t)
	ctx, cancel := testutil.CreateDefaultTestContext()
	defer cancel()

	t.Run("GetHyperlocationOper", func(t *testing.T) {
		result, err := GetHyperlocationOper(client, ctx)
		if err != nil {
			t.Logf("GetHyperlocationOper failed (expected if Hyperlocation not configured): %v", err)
			return
		}

		if result == nil {
			t.Error("Expected non-nil result")
			return
		}

		// Log the response for debugging
		t.Logf("GetHyperlocationOper response received")

		// Save to file for analysis
		testutil.SaveTestDataWithLogging(fmt.Sprintf("test_data_%s.json", "hyperlocation_oper_data"), result)
	})

}

// TestGetHyperlocationProfiles tests the GetHyperlocationProfiles method
func TestGetHyperlocationProfiles(t *testing.T) {
	client := getTestClient(t)
	ctx, cancel := testutil.CreateDefaultTestContext()
	defer cancel()

	t.Run("GetHyperlocationProfiles", func(t *testing.T) {
		result, err := GetHyperlocationProfiles(client, ctx)
		if err != nil {
			t.Logf("GetHyperlocationProfiles failed (expected if Hyperlocation not configured): %v", err)
			return
		}

		if result == nil {
			t.Error("Expected non-nil result")
			return
		}

		t.Logf("GetHyperlocationProfiles response received")
		testutil.SaveTestDataWithLogging(fmt.Sprintf("test_data_%s.json", "hyperlocation_profiles_data"), result)
	})

}

// TestHyperlocationComprehensiveOperations tests all Hyperlocation operations comprehensively
func TestHyperlocationComprehensiveOperations(t *testing.T) {
	client := getTestClient(t)

	ctx, cancel := context.WithTimeout(context.Background(), 125*time.Second)
	defer cancel()

	operations := map[string]func() (interface{}, error){
		"GetHyperlocationOper": func() (interface{}, error) {
			return GetHyperlocationOper(client, ctx)
		},
		"GetHyperlocationProfiles": func() (interface{}, error) {
			return GetHyperlocationProfiles(client, ctx)
		},
	}

	allResults := make(map[string]interface{})

	for operationName, operation := range operations {
		t.Run(operationName, func(t *testing.T) {
			start := time.Now()
			result, err := operation()
			duration := time.Since(start)

			if err != nil {
				t.Logf("%s failed after %v (may be expected if Hyperlocation not configured): %v", operationName, duration, err)
				allResults[operationName] = map[string]interface{}{
					"error":    err.Error(),
					"duration": duration.String(),
				}
				return
			}

			if result == nil {
				t.Errorf("%s returned nil result", operationName)
				return
			}

			t.Logf("%s completed successfully in %v", operationName, duration)
			allResults[operationName] = map[string]interface{}{
				"success":  true,
				"duration": duration.String(),
				"data":     result,
			}
		})
	}

	// Save comprehensive results
	testutil.SaveTestDataWithLogging("test_data_hyperlocation_comprehensive_test_results.json", allResults)
}

// TestHyperlocationOperClientInterfaceCompliance verifies that Client implements all Hyperlocation methods
func TestHyperlocationOperClientInterfaceCompliance(t *testing.T) {
	client := getTestClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Test that all methods exist and can be called (they will fail without real config, but should compile)
	t.Run("MethodExistence", func(t *testing.T) {
		// These calls will fail but should compile, proving the methods exist
		_, _ = GetHyperlocationOper(client, ctx)
		_, _ = GetHyperlocationProfiles(client, ctx)
	})
}
