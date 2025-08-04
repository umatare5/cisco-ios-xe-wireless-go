// Package location provides location configuration test functionality for the Cisco Wireless Network Controller API.
package location

import (
	"context"
	"testing"

	testutils "github.com/umatare5/cisco-ios-xe-wireless-go/internal/tests"
)

// =============================================================================
// 1. UNIT TESTS (Structure/Type Validation & JSON Serialization/Deserialization)
// =============================================================================

func TestLocationCfgDataStructures(t *testing.T) {
	testCases := []testutils.JSONTestCase{
		{
			Name: "LocationCfgResponse",
			JSONData: `{
				"Cisco-IOS-XE-wireless-location-cfg:location-cfg-data": {
					"nmsp-config": {
						"nmsp-enable": true,
						"data-collection-enable": true
					}
				}
			}`,
			Target:     &LocationCfgResponse{},
			TypeName:   "LocationCfgResponse",
			ShouldFail: false,
		},
		{
			Name: "LocationNmspConfigResponse",
			JSONData: `{
				"Cisco-IOS-XE-wireless-location-cfg:nmsp-config": {
					"nmsp-enable": false,
					"data-collection-enable": false
				}
			}`,
			Target:     &LocationNmspConfigResponse{},
			TypeName:   "LocationNmspConfigResponse",
			ShouldFail: false,
		},
	}

	testutils.RunJSONTests(t, testCases)
}

// =============================================================================
// 2. ERROR HANDLING TESTS
// =============================================================================

func TestLocationErrorHandling(t *testing.T) {
	ctx := context.Background()

	t.Run("GetLocationCfgWithNilClient", func(t *testing.T) {
		_, err := GetLocationCfg(nil, ctx)
		if err == nil {
			t.Error("Expected error with nil client, got nil")
		}
		testutils.ValidateErrorContains(t, err, "client is nil")
	})

	t.Run("GetLocationNmspConfigWithNilClient", func(t *testing.T) {
		_, err := GetLocationNmspConfig(nil, ctx)
		if err == nil {
			t.Error("Expected error with nil client, got nil")
		}
		testutils.ValidateErrorContains(t, err, "client is nil")
	})
}

// =============================================================================
// 3. INTEGRATION TESTS
// =============================================================================

func TestLocationConfigurationFunctions(t *testing.T) {
	client := testutils.CreateTestClientFromEnv(t)
	ctx := context.Background()

	// Function test configurations
	testCases := []struct {
		name     string
		testFunc func() (interface{}, error)
	}{
		{
			name: "GetLocationCfg",
			testFunc: func() (interface{}, error) {
				return GetLocationCfg(client, ctx)
			},
		},
		{
			name: "GetLocationNmspConfig",
			testFunc: func() (interface{}, error) {
				return GetLocationNmspConfig(client, ctx)
			},
		},
	}

	// Execute tests and collect data
	collector := make(map[string]interface{})

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			data, err := tc.testFunc()
			if err != nil {
				t.Logf("%s returned error: %v", tc.name, err)
				collector[tc.name] = map[string]interface{}{
					"error": err.Error(),
				}
				return
			}

			if data == nil {
				t.Logf("%s returned nil data", tc.name)
				collector[tc.name] = nil
				return
			}

			t.Logf("%s executed successfully", tc.name)
			collector[tc.name] = data
		})
	}

	// Save collected data to file
	testDataMap := map[string]interface{}{
		"location_configuration_test_data": collector,
	}

	if err := testutils.SaveTestDataToFile("location_configuration_test_data_collected.json", testDataMap); err != nil {
		t.Errorf("Failed to save test data: %v", err)
		return
	}

	t.Logf("Test data saved to test_data/location_configuration_test_data_collected.json")
}

// =============================================================================
// 4. ENDPOINT TESTS - API URL Validation
// =============================================================================

func TestLocationConfigurationEndpoints(t *testing.T) {
	tests := []struct {
		name        string
		endpoint    string
		description string
	}{
		{
			name:        "LocationCfgBasePath",
			endpoint:    LocationCfgBasePath,
			description: "Location configuration base path",
		},
		{
			name:        "LocationCfgEndpoint",
			endpoint:    LocationCfgEndpoint,
			description: "Location configuration endpoint",
		},
		{
			name:        "LocationCfgNmspConfigEndpoint",
			endpoint:    LocationCfgNmspConfigEndpoint,
			description: "Location NMSP configuration endpoint",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testutils.EndpointValidationTest(t, tt.endpoint, tt.endpoint)
		})
	}
}
