// Package mcast provides multicast operational data test functionality for the Cisco Wireless Network Controller API.
package mcast

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

// TestMcastOperEndpoints tests that all Multicast operation endpoint constants are correctly defined
func TestMcastOperEndpoints(t *testing.T) {
	expectedEndpoints := map[string]string{
		"McastOperBasePath":                    "/restconf/data/Cisco-IOS-XE-wireless-mcast-oper:mcast-oper-data",
		"McastOperEndpoint":                    "/restconf/data/Cisco-IOS-XE-wireless-mcast-oper:mcast-oper-data",
		"FlexMediastreamClientSummaryEndpoint": "/restconf/data/Cisco-IOS-XE-wireless-mcast-oper:mcast-oper-data/flex-mediastream-client-summary",
		"VlanL2MgidOpEndpoint":                 "/restconf/data/Cisco-IOS-XE-wireless-mcast-oper:mcast-oper-data/vlan-l2-mgid-op",
	}

	for name, expected := range expectedEndpoints {
		t.Run(name, func(t *testing.T) {
			switch name {
			case "McastOperBasePath":
				if McastOperBasePath != expected {
					t.Errorf(wnc.EndpointMismatchErrorTemplate, name, expected, McastOperBasePath)
				}
			case "McastOperEndpoint":
				if McastOperEndpoint != expected {
					t.Errorf(wnc.EndpointMismatchErrorTemplate, name, expected, McastOperEndpoint)
				}
			case "FlexMediastreamClientSummaryEndpoint":
				if FlexMediastreamClientSummaryEndpoint != expected {
					t.Errorf(wnc.EndpointMismatchErrorTemplate, name, expected, FlexMediastreamClientSummaryEndpoint)
				}
			case "VlanL2MgidOpEndpoint":
				if VlanL2MgidOpEndpoint != expected {
					t.Errorf(wnc.EndpointMismatchErrorTemplate, name, expected, VlanL2MgidOpEndpoint)
				}
			}
		})
	}
}

// TestMcastOperDataStructures tests the basic structure of Multicast operational data types
func TestMcastOperDataStructures(t *testing.T) {
	tests := []struct {
		name     string
		jsonData string
		dataType interface{}
	}{
		{
			name: "McastOperResponse",
			jsonData: `{
				"Cisco-IOS-XE-wireless-mcast-oper:mcast-oper-data": {
					"flex-mediastream-client-summary": [
						{
							"wlan-id": 1,
							"client-count": 10,
							"is-nonip-multicast-enabled": true,
							"is-broadcast-enable": true
						}
					],
					"vlan-l2-mgid-op": [
						{
							"vlan-id": 100,
							"mgid": 1000,
							"multicast-group": "224.1.1.1"
						}
					]
				}
			}`,
			dataType: &McastOperResponse{},
		},
		{
			name: "McastOperFlexMediastreamClientSummaryResponse",
			jsonData: `{
				"Cisco-IOS-XE-wireless-mcast-oper:flex-mediastream-client-summary": [
					{
						"wlan-id": 1,
						"client-count": 10,
						"is-nonip-multicast-enabled": true,
						"is-broadcast-enable": true
					}
				]
			}`,
			dataType: &McastOperFlexMediastreamClientSummaryResponse{},
		},
		{
			name: "McastOperVlanL2MgidOpResponse",
			jsonData: `{
				"Cisco-IOS-XE-wireless-mcast-oper:vlan-l2-mgid-op": [
					{
						"vlan-id": 100,
						"mgid": 1000,
						"multicast-group": "224.1.1.1"
					}
				]
			}`,
			dataType: &McastOperVlanL2MgidOpResponse{},
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

// TestGetMcastOper tests the GetMcastOper method
func TestGetMcastOper(t *testing.T) {
	client := getTestClient(t)

	ctx, cancel := context.WithTimeout(context.Background(), testutil.DefaultTestTimeout)
	defer cancel()

	t.Run("GetMcastOper", func(t *testing.T) {
		result, err := GetMcastOper(client, ctx)
		if err != nil {
			t.Logf("GetMcastOper failed (expected if Multicast not configured): %v", err)
			return
		}

		if result == nil {
			t.Error("Expected non-nil result")
			return
		}

		// Log the response for debugging
		t.Logf("GetMcastOper response received")

		// Save to file for analysis
		saveMcastTestData("mcast_oper_data", result)
	})

}

// TestGetMcastFlexMediastreamClientSummary tests the GetMcastFlexMediastreamClientSummary method
func TestGetMcastFlexMediastreamClientSummary(t *testing.T) {
	client := getTestClient(t)

	ctx, cancel := context.WithTimeout(context.Background(), testutil.DefaultTestTimeout)
	defer cancel()

	t.Run("GetMcastFlexMediastreamClientSummary", func(t *testing.T) {
		result, err := GetMcastFlexMediastreamClientSummary(client, ctx)
		if err != nil {
			t.Logf("GetMcastFlexMediastreamClientSummary failed (expected if Multicast not configured): %v", err)
			return
		}

		if result == nil {
			t.Error("Expected non-nil result")
			return
		}

		t.Logf("GetMcastFlexMediastreamClientSummary response received")
		saveMcastTestData("mcast_flex_mediastream_client_summary_data", result)
	})

}

// TestGetMcastVlanL2MgidOp tests the GetMcastVlanL2MgidOp method
func TestGetMcastVlanL2MgidOp(t *testing.T) {
	client := getTestClient(t)

	ctx, cancel := context.WithTimeout(context.Background(), testutil.DefaultTestTimeout)
	defer cancel()

	t.Run("GetMcastVlanL2MgidOp", func(t *testing.T) {
		result, err := GetMcastVlanL2MgidOp(client, ctx)
		if err != nil {
			t.Logf("GetMcastVlanL2MgidOp failed (expected if Multicast not configured): %v", err)
			return
		}

		if result == nil {
			t.Error("Expected non-nil result")
			return
		}

		t.Logf("GetMcastVlanL2MgidOp response received")
		saveMcastTestData("mcast_vlan_l2_mgid_op_data", result)
	})

}

// TestMcastComprehensiveOperations tests all Multicast operations comprehensively
func TestMcastComprehensiveOperations(t *testing.T) {
	client := getTestClient(t)

	ctx, cancel := context.WithTimeout(context.Background(), testutil.ExtendedTestTimeout)
	defer cancel()

	operations := map[string]func() (interface{}, error){
		"GetMcastOper": func() (interface{}, error) {
			return GetMcastOper(client, ctx)
		},
		"GetMcastFlexMediastreamClientSummary": func() (interface{}, error) {
			return GetMcastFlexMediastreamClientSummary(client, ctx)
		},
		"GetMcastVlanL2MgidOp": func() (interface{}, error) {
			return GetMcastVlanL2MgidOp(client, ctx)
		},
	}

	allResults := make(map[string]interface{})

	for operationName, operation := range operations {
		t.Run(operationName, func(t *testing.T) {
			start := time.Now()
			result, err := operation()
			duration := time.Since(start)

			if err != nil {
				t.Logf("%s failed after %v (may be expected if Multicast not configured): %v", operationName, duration, err)
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
	saveMcastTestData("mcast_comprehensive_test_results", allResults)
}

func saveMcastTestData(filename string, data interface{}) {
	if data == nil {
		return
	}

	if err := testutil.SaveTestDataToFile(fmt.Sprintf("test_data_%s.json", filename), data); err != nil {
		fmt.Printf("Error saving test data for %s: %v\n", filename, err)
	} else {
		fmt.Printf("Test data saved to %s/test_data_%s.json\n", testutil.TestDataDir, filename)
	}
}

// TestMcastOperClientInterfaceCompliance verifies that Client implements all Multicast methods
func TestMcastOperClientInterfaceCompliance(t *testing.T) {
	client := getTestClient(t)

	ctx, cancel := context.WithTimeout(context.Background(), testutil.QuickTestTimeout)
	defer cancel()

	// Test that all methods exist and can be called
	t.Run("MethodExistence", func(t *testing.T) {
		// These calls should succeed with real client
		_, _ = GetMcastOper(client, ctx)
		_, _ = GetMcastFlexMediastreamClientSummary(client, ctx)
		_, _ = GetMcastVlanL2MgidOp(client, ctx)
		t.Log("All Mcast operation methods exist and are callable")
	})
}
