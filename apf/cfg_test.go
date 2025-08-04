// Package apf provides Access Point Filter configuration test functionality for the Cisco Wireless Network Controller API.
package apf

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

func TestApfCfgDataStructures(t *testing.T) {
	testCases := []testutils.JSONTestCase{
		{
			Name: "ApfCfgResponse",
			JSONData: `{
				"Cisco-IOS-XE-wireless-apf-cfg:apf-cfg-data": {
					"apf": {
						"system-mgmt-via-wireless": true,
						"network-name": "corporate-network"
					}
				}
			}`,
			Target:     &ApfCfgResponse{},
			TypeName:   "ApfCfgResponse",
			ShouldFail: false,
		},
		{
			Name: "ApfCfgApfResponse",
			JSONData: `{
				"Cisco-IOS-XE-wireless-apf-cfg:apf": {
					"system-mgmt-via-wireless": false,
					"network-name": "guest-network"
				}
			}`,
			Target:     &ApfCfgApfResponse{},
			TypeName:   "ApfCfgApfResponse",
			ShouldFail: false,
		},
		{
			Name: "Apf",
			JSONData: `{
				"system-mgmt-via-wireless": true,
				"network-name": "production-network"
			}`,
			Target:     &Apf{},
			TypeName:   "Apf",
			ShouldFail: false,
		},
		{
			Name: "InvalidApfCfgResponse",
			JSONData: `{
				"invalid-field": "invalid-value"
			}`,
			Target:     &ApfCfgResponse{},
			TypeName:   "ApfCfgResponse",
			ShouldFail: false, // JSON unmarshaling is lenient with extra fields
		},
	}

	testutils.RunJSONTests(t, testCases)

	// Additional field validation for successfully unmarshaled structures
	t.Run("ApfCfgResponseFieldValidation", func(t *testing.T) {
		var response ApfCfgResponse
		testutils.TestJSONUnmarshal(t, testCases[0].JSONData, &response, "ApfCfgResponse")

		testutils.ValidateJSONStructFields(t, "ApfCfgResponse", func() error {
			if !response.CiscoIOSXEWirelessApfCfgApfCfgData.Apf.SystemMgmtViaWireless {
				t.Error("Expected system-mgmt-via-wireless to be true")
			}
			if response.CiscoIOSXEWirelessApfCfgApfCfgData.Apf.NetworkName != "corporate-network" {
				t.Errorf("Expected network name 'corporate-network', got '%s'",
					response.CiscoIOSXEWirelessApfCfgApfCfgData.Apf.NetworkName)
			}
			return nil
		})
	})

	t.Run("ApfCfgApfResponseFieldValidation", func(t *testing.T) {
		var response ApfCfgApfResponse
		testutils.TestJSONUnmarshal(t, testCases[1].JSONData, &response, "ApfCfgApfResponse")

		testutils.ValidateJSONStructFields(t, "ApfCfgApfResponse", func() error {
			if response.Apf.SystemMgmtViaWireless {
				t.Error("Expected system-mgmt-via-wireless to be false")
			}
			if response.Apf.NetworkName != "guest-network" {
				t.Errorf("Expected network name 'guest-network', got '%s'", response.Apf.NetworkName)
			}
			return nil
		})
	})

	t.Run("ApfFieldValidation", func(t *testing.T) {
		var apf Apf
		testutils.TestJSONUnmarshal(t, testCases[2].JSONData, &apf, "Apf")

		testutils.ValidateJSONStructFields(t, "Apf", func() error {
			if !apf.SystemMgmtViaWireless {
				t.Error("Expected system-mgmt-via-wireless to be true")
			}
			if apf.NetworkName != "production-network" {
				t.Errorf("Expected network name 'production-network', got '%s'", apf.NetworkName)
			}
			return nil
		})
	})
}

// =============================================================================
// 2. INTEGRATION TESTS (Actual API Calls to Live Controller)
// =============================================================================

// TestApfConfigurationFunctions tests all APF configuration functions with a live controller
func TestApfConfigurationFunctions(t *testing.T) {
	client := testutils.GetTestClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test GetApfCfg function
	t.Run("GetApfCfg", func(t *testing.T) {
		data, err := GetApfCfg(client, ctx)
		if err != nil {
			t.Fatalf("GetApfCfg failed: %v", err)
		}

		// Validate basic structure
		if data == nil {
			t.Fatal("GetApfCfg returned nil data")
		}

		// Save test data for analysis
		if err := testutils.SaveTestDataToFile("apf_cfg_data.json", data); err != nil {
			t.Logf("Warning: Could not save test data: %v", err)
		} else {
			t.Logf("APF config data saved to test_data/apf_cfg_data.json")
		}

		// Validate endpoint was constructed correctly
		endpoint := ApfCfgEndpoint
		if endpoint == "" {
			t.Error("ApfCfgEndpoint should not be empty")
		}
		if endpoint != "/restconf/data/Cisco-IOS-XE-wireless-apf-cfg:apf-cfg-data" {
			t.Errorf("ApfCfgEndpoint unexpected value: got %s", endpoint)
		}
	})

	// Test GetApf function
	t.Run("GetApf", func(t *testing.T) {
		data, err := GetApf(client, ctx)
		if err != nil {
			t.Fatalf("GetApf failed: %v", err)
		}

		// Validate basic structure
		if data == nil {
			t.Fatal("GetApf returned nil data")
		}

		// Save test data for analysis
		if err := testutils.SaveTestDataToFile("apf_data.json", data); err != nil {
			t.Logf("Warning: Could not save test data: %v", err)
		} else {
			t.Logf("APF data saved to test_data/apf_data.json")
		}

		// Validate endpoint was constructed correctly
		endpoint := ApfEndpoint
		if endpoint == "" {
			t.Error("ApfEndpoint should not be empty")
		}
		expectedEndpoint := "/restconf/data/Cisco-IOS-XE-wireless-apf-cfg:apf-cfg-data/apf"
		if endpoint != expectedEndpoint {
			t.Errorf("ApfEndpoint unexpected value: expected %s, got %s", expectedEndpoint, endpoint)
		}
	})

	// Test error handling with common error patterns
	testutils.RunCommonErrorTests(t, "ApfErrorHandling", []testutils.ErrorTestCase{
		{
			Name: "GetApfCfgWithNilClient",
			TestFunc: func(client *wnc.Client) error {
				_, err := GetApfCfg(nil, ctx)
				return err
			},
			ExpectedError: "client is nil",
		},
		{
			Name: "GetApfWithNilClient",
			TestFunc: func(client *wnc.Client) error {
				_, err := GetApf(nil, ctx)
				return err
			},
			ExpectedError: "client is nil",
		},
	})

	// Test context handling
	t.Run("ContextHandling", func(t *testing.T) {
		testutils.TestContextHandling(t, func(ctx context.Context, client *wnc.Client) error {
			_, err := GetApfCfg(client, ctx)
			return err
		})

		testutils.TestContextHandling(t, func(ctx context.Context, client *wnc.Client) error {
			_, err := GetApf(client, ctx)
			return err
		})
	})
}

// TestApfConfigurationEndpoints validates APF configuration endpoint constants
func TestApfConfigurationEndpoints(t *testing.T) {
	// Test base path validation
	t.Run("Validate_ApfCfgBasePath", func(t *testing.T) {
		expectedBasePath := "/restconf/data/Cisco-IOS-XE-wireless-apf-cfg:apf-cfg-data"
		if ApfCfgBasePath != expectedBasePath {
			t.Errorf("ApfCfgBasePath mismatch: expected %s, got %s", expectedBasePath, ApfCfgBasePath)
		}
	})

	// Test endpoint validation
	t.Run("Validate_ApfCfgEndpoint", func(t *testing.T) {
		if ApfCfgEndpoint != ApfCfgBasePath {
			t.Errorf("ApfCfgEndpoint should equal ApfCfgBasePath: expected %s, got %s", ApfCfgBasePath, ApfCfgEndpoint)
		}
	})

	// Test APF specific endpoint validation
	t.Run("Validate_ApfEndpoint", func(t *testing.T) {
		expectedEndpoint := ApfCfgBasePath + "/apf"
		if ApfEndpoint != expectedEndpoint {
			t.Errorf("ApfEndpoint mismatch: expected %s, got %s", expectedEndpoint, ApfEndpoint)
		}
	})
}
