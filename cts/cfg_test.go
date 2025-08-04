// Package cts provides Cisco TrustSec configuration test functionality for the Cisco Wireless Network Controller API.
package cts

import (
	"context"
	"testing"
	"time"

	wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
	testutils "github.com/umatare5/cisco-ios-xe-wireless-go/internal/tests"
)

// =============================================================================
// 1. UNIT TESTS (Structure/Type Validation & JSON Serialization/Deserialization)
// =============================================================================

func TestCtsSxpCfgDataStructures(t *testing.T) {
	testCases := []testutils.JSONTestCase{
		{
			Name: "CtsSxpCfgResponse",
			JSONData: `{
				"Cisco-IOS-XE-wireless-cts-sxp-cfg:cts-sxp-cfg-data": {
					"cts-sxp-configuration": {
						"cts-sxp-config": [
							{
								"sxp-profile-name": "corporate-sxp"
							},
							{
								"sxp-profile-name": "guest-sxp"
							}
						]
					}
				}
			}`,
			Target:     &CtsSxpCfgResponse{},
			TypeName:   "CtsSxpCfgResponse",
			ShouldFail: false,
		},
		{
			Name: "CtsSxpConfigurationResponse",
			JSONData: `{
				"Cisco-IOS-XE-wireless-cts-sxp-cfg:cts-sxp-configuration": {
					"cts-sxp-config": [
						{
							"sxp-profile-name": "production-sxp"
						}
					]
				}
			}`,
			Target:     &CtsSxpConfigurationResponse{},
			TypeName:   "CtsSxpConfigurationResponse",
			ShouldFail: false,
		},
		{
			Name: "CtsSxpConfig",
			JSONData: `{
				"sxp-profile-name": "test-sxp"
			}`,
			Target:     &CtsSxpConfig{},
			TypeName:   "CtsSxpConfig",
			ShouldFail: false,
		},
	}

	testutils.RunJSONTests(t, testCases)

	// Additional field validation for successfully unmarshaled structures
	t.Run("CtsSxpCfgResponseFieldValidation", func(t *testing.T) {
		var response CtsSxpCfgResponse
		testutils.TestJSONUnmarshal(t, testCases[0].JSONData, &response, "CtsSxpCfgResponse")

		testutils.ValidateJSONStructFields(t, "CtsSxpCfgResponse", func() error {
			if len(response.CiscoIOSXEWirelessCtsSxpCfgCtsSxpCfgData.CtsSxpConfiguration.CtsSxpConfig) != 2 {
				t.Errorf("Expected 2 cts-sxp-config entries, got %d",
					len(response.CiscoIOSXEWirelessCtsSxpCfgCtsSxpCfgData.CtsSxpConfiguration.CtsSxpConfig))
			}
			if response.CiscoIOSXEWirelessCtsSxpCfgCtsSxpCfgData.CtsSxpConfiguration.CtsSxpConfig[0].SxpProfileName != "corporate-sxp" {
				t.Errorf("Expected first profile name 'corporate-sxp', got '%s'",
					response.CiscoIOSXEWirelessCtsSxpCfgCtsSxpCfgData.CtsSxpConfiguration.CtsSxpConfig[0].SxpProfileName)
			}
			return nil
		})
	})

	t.Run("CtsSxpConfigurationResponseFieldValidation", func(t *testing.T) {
		var response CtsSxpConfigurationResponse
		testutils.TestJSONUnmarshal(t, testCases[1].JSONData, &response, "CtsSxpConfigurationResponse")

		testutils.ValidateJSONStructFields(t, "CtsSxpConfigurationResponse", func() error {
			if len(response.CtsSxpConfiguration.CtsSxpConfig) != 1 {
				t.Errorf("Expected 1 cts-sxp-config entry, got %d", len(response.CtsSxpConfiguration.CtsSxpConfig))
			}
			if response.CtsSxpConfiguration.CtsSxpConfig[0].SxpProfileName != "production-sxp" {
				t.Errorf("Expected profile name 'production-sxp', got '%s'", response.CtsSxpConfiguration.CtsSxpConfig[0].SxpProfileName)
			}
			return nil
		})
	})

	t.Run("CtsSxpConfigFieldValidation", func(t *testing.T) {
		var config CtsSxpConfig
		testutils.TestJSONUnmarshal(t, testCases[2].JSONData, &config, "CtsSxpConfig")

		testutils.ValidateJSONStructFields(t, "CtsSxpConfig", func() error {
			if config.SxpProfileName != "test-sxp" {
				t.Errorf("Expected profile name 'test-sxp', got '%s'", config.SxpProfileName)
			}
			return nil
		})
	})
}

// =============================================================================
// 2. INTEGRATION TESTS (Actual API Calls to Live Controller)
// =============================================================================

// TestCtsConfigurationFunctions tests all CTS configuration functions with a live controller
func TestCtsConfigurationFunctions(t *testing.T) {
	client := testutils.GetTestClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test GetCtsSxpCfg function
	t.Run("GetCtsSxpCfg", func(t *testing.T) {
		data, err := GetCtsSxpCfg(client, ctx)
		if err != nil {
			t.Fatalf("GetCtsSxpCfg failed: %v", err)
		}

		// Validate basic structure
		if data == nil {
			t.Fatal("GetCtsSxpCfg returned nil data")
		}

		// Save test data for analysis
		if err := testutils.SaveTestDataToFile("cts_sxp_cfg_data.json", data); err != nil {
			t.Logf("Warning: Could not save test data: %v", err)
		} else {
			t.Logf("CTS SXP config data saved to test_data/cts_sxp_cfg_data.json")
		}

		// Validate endpoint was constructed correctly
		endpoint := CtsSxpCfgEndpoint
		if endpoint == "" {
			t.Error("CtsSxpCfgEndpoint should not be empty")
		}
	})

	// Test GetCtsSxpConfiguration function
	t.Run("GetCtsSxpConfiguration", func(t *testing.T) {
		data, err := GetCtsSxpConfiguration(client, ctx)
		if err != nil {
			t.Fatalf("GetCtsSxpConfiguration failed: %v", err)
		}

		// Validate basic structure
		if data == nil {
			t.Fatal("GetCtsSxpConfiguration returned nil data")
		}

		// Save test data for analysis
		if err := testutils.SaveTestDataToFile("cts_sxp_configuration_data.json", data); err != nil {
			t.Logf("Warning: Could not save test data: %v", err)
		} else {
			t.Logf("CTS SXP configuration data saved to test_data/cts_sxp_configuration_data.json")
		}

		// Validate endpoint was constructed correctly
		endpoint := CtsSxpConfigurationEndpoint
		if endpoint == "" {
			t.Error("CtsSxpConfigurationEndpoint should not be empty")
		}
	})

	// Test error handling with common error patterns
	testutils.RunCommonErrorTests(t, "CtsErrorHandling", []testutils.ErrorTestCase{
		{
			Name: "GetCtsSxpCfgWithNilClient",
			TestFunc: func(client *wnc.Client) error {
				_, err := GetCtsSxpCfg(nil, ctx)
				return err
			},
			ExpectedError: "client is nil",
		},
		{
			Name: "GetCtsSxpConfigurationWithNilClient",
			TestFunc: func(client *wnc.Client) error {
				_, err := GetCtsSxpConfiguration(nil, ctx)
				return err
			},
			ExpectedError: "client is nil",
		},
	})

	// Test context handling
	t.Run("ContextHandling", func(t *testing.T) {
		testutils.TestContextHandling(t, func(ctx context.Context, client *wnc.Client) error {
			_, err := GetCtsSxpCfg(client, ctx)
			return err
		})

		testutils.TestContextHandling(t, func(ctx context.Context, client *wnc.Client) error {
			_, err := GetCtsSxpConfiguration(client, ctx)
			return err
		})
	})
}

// TestCtsConfigurationEndpoints validates CTS configuration endpoint constants
func TestCtsConfigurationEndpoints(t *testing.T) {
	// Test base path validation
	t.Run("Validate_CtsSxpCfgBasePath", func(t *testing.T) {
		expectedBasePath := "/restconf/data/Cisco-IOS-XE-wireless-cts-sxp-cfg:cts-sxp-cfg-data"
		if CtsSxpCfgBasePath != expectedBasePath {
			t.Errorf("CtsSxpCfgBasePath mismatch: expected %s, got %s", expectedBasePath, CtsSxpCfgBasePath)
		}
	})

	// Test endpoint validation
	t.Run("Validate_CtsSxpCfgEndpoint", func(t *testing.T) {
		if CtsSxpCfgEndpoint != CtsSxpCfgBasePath {
			t.Errorf("CtsSxpCfgEndpoint should equal CtsSxpCfgBasePath: expected %s, got %s", CtsSxpCfgBasePath, CtsSxpCfgEndpoint)
		}
	})

	// Test CTS SXP configuration endpoint validation
	t.Run("Validate_CtsSxpConfigurationEndpoint", func(t *testing.T) {
		expectedEndpoint := CtsSxpCfgBasePath + "/cts-sxp-configuration"
		if CtsSxpConfigurationEndpoint != expectedEndpoint {
			t.Errorf("CtsSxpConfigurationEndpoint mismatch: expected %s, got %s", expectedEndpoint, CtsSxpConfigurationEndpoint)
		}
	})
}
