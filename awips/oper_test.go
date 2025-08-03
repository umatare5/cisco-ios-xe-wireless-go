// Package awips provides AWIPS (Adaptive Wireless Intrusion Prevention System) operational data test functionality for the Cisco Wireless Network Controller API.
package awips

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

// TestAWIPSOperEndpoints tests that all AWIPS operation endpoint constants are correctly defined
func TestAWIPSOperEndpoints(t *testing.T) {
	expectedEndpoints := map[string]string{
		"AwipsOperBasePath":         "/restconf/data/Cisco-IOS-XE-wireless-awips-oper:awips-oper-data",
		"AwipsOperEndpoint":         "/restconf/data/Cisco-IOS-XE-wireless-awips-oper:awips-oper-data",
		"AwipsPerApInfoEndpoint":    "/restconf/data/Cisco-IOS-XE-wireless-awips-oper:awips-oper-data/awips-per-ap-info",
		"AwipsDwldStatusEndpoint":   "/restconf/data/Cisco-IOS-XE-wireless-awips-oper:awips-oper-data/awips-dwld-status",
		"AwipsApDwldStatusEndpoint": "/restconf/data/Cisco-IOS-XE-wireless-awips-oper:awips-oper-data/awips-ap-dwld-status",
	}

	actualEndpoints := map[string]string{
		"AwipsOperBasePath":         AwipsOperBasePath,
		"AwipsOperEndpoint":         AwipsOperEndpoint,
		"AwipsPerApInfoEndpoint":    AwipsPerApInfoEndpoint,
		"AwipsDwldStatusEndpoint":   AwipsDwldStatusEndpoint,
		"AwipsApDwldStatusEndpoint": AwipsApDwldStatusEndpoint,
	}

	testutil.GenerateEndpointValidationTest(t, expectedEndpoints, actualEndpoints)
}

// TestAWIPSOperDataStructures tests the basic structure of AWIPS operational data types
func TestAWIPSOperDataStructures(t *testing.T) {
	tests := []struct {
		name     string
		jsonData string
		dataType interface{}
	}{
		{
			name: "AwipsOperResponse",
			jsonData: `{
				"Cisco-IOS-XE-wireless-awips-oper:awips-oper-data": {
					"awips-per-ap-info": [],
					"awips-dwld-status": {
						"last-success-timestamp": "2024-01-01T00:00:00.000Z",
						"last-failed-timestamp": "",
						"num-of-failure-attempts": 0,
						"last-failure-reason": 0,
						"wlc-version": "17.12.01",
						"max-file-ver": 1,
						"latest-file-version": 1,
						"download-status": "success",
						"file-hash": "abc123"
					},
					"awips-ap-dwld-status": []
				}
			}`,
			dataType: &AwipsOperResponse{},
		},
		{
			name: "AwipsOperPerApInfoResponse",
			jsonData: `{
				"Cisco-IOS-XE-wireless-awips-oper:awips-per-ap-info": [
					{
						"ap-mac": "aa:bb:cc:dd:ee:ff",
						"awips-status": "enabled",
						"alarm-count": "0",
						"forensic-capture-status": "disabled"
					}
				]
			}`,
			dataType: &AwipsOperPerApInfoResponse{},
		},
		{
			name: "AwipsOperDwldStatusResponse",
			jsonData: `{
				"Cisco-IOS-XE-wireless-awips-oper:awips-dwld-status": {
					"last-success-timestamp": "2024-01-01T00:00:00.000Z",
					"last-failed-timestamp": "",
					"num-of-failure-attempts": 0,
					"last-failure-reason": 0,
					"wlc-version": "17.12.01",
					"max-file-ver": 1,
					"latest-file-version": 1,
					"download-status": "success",
					"file-hash": "abc123"
				}
			}`,
			dataType: &AwipsOperDwldStatusResponse{},
		},
		{
			name: "AwipsOperApDwldStatusResponse",
			jsonData: `{
				"Cisco-IOS-XE-wireless-awips-oper:awips-ap-dwld-status": [
					{
						"ap-mac": "aa:bb:cc:dd:ee:ff",
						"dwld-status": "success",
						"file-version": 1,
						"file-hash": "abc123"
					}
				]
			}`,
			dataType: &AwipsOperApDwldStatusResponse{},
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

// getTestClient creates a test client using environment variables
func getTestClient(t *testing.T) *wnc.Client {
	return testutil.CreateTestClientFromEnv(t)
}

// TestGetAwipsOper tests the GetAwipsOper method
func TestGetAwipsOper(t *testing.T) {
	client := getTestClient(t)
	ctx, cancel := testutil.CreateDefaultTestContext()
	defer cancel()

	t.Run("GetAwipsOper", func(t *testing.T) {
		result, err := GetAwipsOper(client, ctx)
		if err != nil {
			t.Logf("GetAwipsOper failed (expected if AWIPS not configured): %v", err)
			return
		}

		if result == nil {
			t.Error("Expected non-nil result")
			return
		}

		// Log the response for debugging
		t.Logf("GetAwipsOper response received")

		// Save to file for analysis
		if err := testutil.SaveTestDataToFile(fmt.Sprintf("test_data_%s.json", "awips_oper_data"), result); err != nil {
			t.Logf("Failed to save test data: %v", err)
		}
	})

}

// TestGetAwipsPerApInfo tests the GetAwipsPerApInfo method
func TestGetAwipsPerApInfo(t *testing.T) {
	client := getTestClient(t)
	ctx, cancel := testutil.CreateDefaultTestContext()
	defer cancel()

	t.Run("GetAwipsPerApInfo", func(t *testing.T) {
		result, err := GetAwipsPerApInfo(client, ctx)
		if err != nil {
			t.Logf("GetAwipsPerApInfo failed (expected if AWIPS not configured): %v", err)
			return
		}

		if result == nil {
			t.Error("Expected non-nil result")
			return
		}

		t.Logf("GetAwipsPerApInfo response received")
		if err := testutil.SaveTestDataToFile(fmt.Sprintf("test_data_%s.json", "awips_per_ap_info_data"), result); err != nil {
			t.Logf("Failed to save test data: %v", err)
		}
	})

}

// TestGetAwipsDwldStatus tests the GetAwipsDwldStatus method
func TestGetAwipsDwldStatus(t *testing.T) {
	client := getTestClient(t)
	ctx, cancel := testutil.CreateDefaultTestContext()
	defer cancel()

	t.Run("GetAwipsDwldStatus", func(t *testing.T) {
		result, err := GetAwipsDwldStatus(client, ctx)
		if err != nil {
			t.Logf("GetAwipsDwldStatus failed (expected if AWIPS not configured): %v", err)
			return
		}

		if result == nil {
			t.Error("Expected non-nil result")
			return
		}

		t.Logf("GetAwipsDwldStatus response received")
		if err := testutil.SaveTestDataToFile(fmt.Sprintf("test_data_%s.json", "awips_dwld_status_data"), result); err != nil {
			t.Logf("Failed to save test data: %v", err)
		}
	})

}

// TestGetAwipsApDwldStatus tests the GetAwipsApDwldStatus method
func TestGetAwipsApDwldStatus(t *testing.T) {
	client := getTestClient(t)
	ctx, cancel := testutil.CreateDefaultTestContext()
	defer cancel()

	t.Run("GetAwipsApDwldStatus", func(t *testing.T) {
		result, err := GetAwipsApDwldStatus(client, ctx)
		if err != nil {
			t.Logf("GetAwipsApDwldStatus failed (expected if AWIPS not configured): %v", err)
			return
		}

		if result == nil {
			t.Error("Expected non-nil result")
			return
		}

		t.Logf("GetAwipsApDwldStatus response received")
		if err := testutil.SaveTestDataToFile(fmt.Sprintf("test_data_%s.json", "awips_ap_dwld_status_data"), result); err != nil {
			t.Logf("Failed to save test data: %v", err)
		}
	})

}

// TestAWIPSComprehensiveOperations tests all AWIPS operations comprehensively
func TestAWIPSComprehensiveOperations(t *testing.T) {
	client := getTestClient(t)

	ctx, cancel := context.WithTimeout(context.Background(), 125*time.Second)
	defer cancel()

	operations := map[string]func() (interface{}, error){
		"GetAwipsOper": func() (interface{}, error) {
			return GetAwipsOper(client, ctx)
		},
		"GetAwipsPerApInfo": func() (interface{}, error) {
			return GetAwipsPerApInfo(client, ctx)
		},
		"GetAwipsDwldStatus": func() (interface{}, error) {
			return GetAwipsDwldStatus(client, ctx)
		},
		"GetAwipsApDwldStatus": func() (interface{}, error) {
			return GetAwipsApDwldStatus(client, ctx)
		},
	}

	allResults := make(map[string]interface{})

	for operationName, operation := range operations {
		t.Run(operationName, func(t *testing.T) {
			start := time.Now()
			result, err := operation()
			duration := time.Since(start)

			if err != nil {
				t.Logf("%s failed after %v (may be expected if AWIPS not configured): %v", operationName, duration, err)
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
	if err := testutil.SaveTestDataToFile("test_data_awips_comprehensive_test_results.json", allResults); err != nil {
		t.Logf("Failed to save comprehensive test results: %v", err)
	}
}

// TestAWIPSOperClientInterfaceCompliance verifies that Client implements all AWIPS methods
func TestAWIPSOperClientInterfaceCompliance(t *testing.T) {
	// Create a properly initialized client for interface compliance testing
	config := wnc.Config{
		Controller:  "test.local",
		AccessToken: "test-token",
	}
	client, err := wnc.NewClient(config)
	if err != nil {
		t.Fatalf("Failed to create test client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Test that all methods exist and can be called (they will fail without real config, but should compile)
	t.Run("MethodExistence", func(t *testing.T) {
		// These calls will fail but should compile, proving the methods exist
		_, _ = GetAwipsOper(client, ctx)
		_, _ = GetAwipsPerApInfo(client, ctx)
		_, _ = GetAwipsDwldStatus(client, ctx)
		_, _ = GetAwipsApDwldStatus(client, ctx)
	})
}
