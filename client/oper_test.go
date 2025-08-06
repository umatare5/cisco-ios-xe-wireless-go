// Package client provides client operational data test functionality for the Cisco Wireless Network Controller API.
package client

import (
	"context"
	"testing"

	wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
	testutils "github.com/umatare5/cisco-ios-xe-wireless-go/internal/tests"
)

// =============================================================================
// 1. UNIT TESTS (Structure/Type Validation & JSON Serialization/Deserialization)
// =============================================================================

func TestClientOperDataStructures(t *testing.T) {
	testCases := []testutils.JSONTestCase{
		{
			Name: "ClientOperCommonOperDataResponse",
			JSONData: `{
				"Cisco-IOS-XE-wireless-client-oper:common-oper-data": []
			}`,
			Target:     &model.ClientOperCommonOperDataResponse{},
			TypeName:   "ClientOperCommonOperDataResponse",
			ShouldFail: false,
		},
		{
			Name: "ClientOperDot11OperDataResponse",
			JSONData: `{
				"Cisco-IOS-XE-wireless-client-oper:dot11-oper-data": []
			}`,
			Target:     &model.ClientOperDot11OperDataResponse{},
			TypeName:   "ClientOperDot11OperDataResponse",
			ShouldFail: false,
		},
		{
			Name: "ClientOperMobilityOperDataResponse",
			JSONData: `{
				"Cisco-IOS-XE-wireless-client-oper:mobility-oper-data": []
			}`,
			Target:     &model.ClientOperMobilityOperDataResponse{},
			TypeName:   "ClientOperMobilityOperDataResponse",
			ShouldFail: false,
		},
		{
			Name: "ClientOperTrafficStatsResponse",
			JSONData: `{
				"Cisco-IOS-XE-wireless-client-oper:traffic-stats": []
			}`,
			Target:     &model.ClientOperTrafficStatsResponse{},
			TypeName:   "ClientOperTrafficStatsResponse",
			ShouldFail: false,
		},
	}

	testutils.RunJSONTests(t, testCases)
}

// =============================================================================
// 2. ERROR HANDLING TESTS
// =============================================================================

func TestClientOperErrorHandling(t *testing.T) {
	// Test nil client scenarios
	t.Run("NilClientTests", func(t *testing.T) {
		var nilClient *wnc.Client = nil
		ctx := context.Background()

		tests := []struct {
			name     string
			testFunc func() error
		}{
			{
				name: "GetClientOper_NilClient",
				testFunc: func() error {
					_, err := GetClientOper(nilClient, ctx)
					return err
				},
			},
			{
				name: "GetClientOperCommonOperData_NilClient",
				testFunc: func() error {
					_, err := GetClientOperCommonOperData(nilClient, ctx)
					return err
				},
			},
			{
				name: "GetClientOperDot11OperData_NilClient",
				testFunc: func() error {
					_, err := GetClientOperDot11OperData(nilClient, ctx)
					return err
				},
			},
			{
				name: "GetClientOperMobilityOperData_NilClient",
				testFunc: func() error {
					_, err := GetClientOperMobilityOperData(nilClient, ctx)
					return err
				},
			},
			{
				name: "GetClientOperMmIfClientStats_NilClient",
				testFunc: func() error {
					_, err := GetClientOperMmIfClientStats(nilClient, ctx)
					return err
				},
			},
			{
				name: "GetClientOperMmIfClientHistory_NilClient",
				testFunc: func() error {
					_, err := GetClientOperMmIfClientHistory(nilClient, ctx)
					return err
				},
			},
			{
				name: "GetClientOperTrafficStats_NilClient",
				testFunc: func() error {
					_, err := GetClientOperTrafficStats(nilClient, ctx)
					return err
				},
			},
			{
				name: "GetClientOperPolicyData_NilClient",
				testFunc: func() error {
					_, err := GetClientOperPolicyData(nilClient, ctx)
					return err
				},
			},
			{
				name: "GetClientOperSisfDbMac_NilClient",
				testFunc: func() error {
					_, err := GetClientOperSisfDbMac(nilClient, ctx)
					return err
				},
			},
			{
				name: "GetClientOperDcInfo_NilClient",
				testFunc: func() error {
					_, err := GetClientOperDcInfo(nilClient, ctx)
					return err
				},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				err := tt.testFunc()
				if err == nil {
					t.Errorf("Expected error for nil client, but got nil")
					return
				}
				testutils.ValidateErrorContains(t, err, "client is nil")
			})
		}
	})
}

// =============================================================================
// 3. INTEGRATION TESTS - Live API Data Validation
// =============================================================================

func TestClientOperFunctions(t *testing.T) {
	client := testutils.CreateTestClientFromEnv(t)
	ctx := context.Background()

	// Function test configurations
	testCases := []struct {
		name     string
		testFunc func() (interface{}, error)
	}{
		{
			name: "GetClientOper",
			testFunc: func() (interface{}, error) {
				return GetClientOper(client, ctx)
			},
		},
		{
			name: "GetClientOperCommonOperData",
			testFunc: func() (interface{}, error) {
				return GetClientOperCommonOperData(client, ctx)
			},
		},
		{
			name: "GetClientOperDot11OperData",
			testFunc: func() (interface{}, error) {
				return GetClientOperDot11OperData(client, ctx)
			},
		},
		{
			name: "GetClientOperMobilityOperData",
			testFunc: func() (interface{}, error) {
				return GetClientOperMobilityOperData(client, ctx)
			},
		},
		{
			name: "GetClientOperMmIfClientStats",
			testFunc: func() (interface{}, error) {
				return GetClientOperMmIfClientStats(client, ctx)
			},
		},
		{
			name: "GetClientOperMmIfClientHistory",
			testFunc: func() (interface{}, error) {
				return GetClientOperMmIfClientHistory(client, ctx)
			},
		},
		{
			name: "GetClientOperTrafficStats",
			testFunc: func() (interface{}, error) {
				return GetClientOperTrafficStats(client, ctx)
			},
		},
		{
			name: "GetClientOperPolicyData",
			testFunc: func() (interface{}, error) {
				return GetClientOperPolicyData(client, ctx)
			},
		},
		{
			name: "GetClientOperSisfDbMac",
			testFunc: func() (interface{}, error) {
				return GetClientOperSisfDbMac(client, ctx)
			},
		},
		{
			name: "GetClientOperDcInfo",
			testFunc: func() (interface{}, error) {
				return GetClientOperDcInfo(client, ctx)
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
		"client_oper_test_data": collector,
	}

	if err := testutils.SaveTestDataToFile("client_oper_test_data_collected.json", testDataMap); err != nil {
		t.Errorf("Failed to save test data: %v", err)
		return
	}

	t.Logf("Test data saved to test_data/client_oper_test_data_collected.json")
}

// =============================================================================
// 4. ENDPOINT TESTS - API URL Validation
// =============================================================================

func TestClientOperEndpoints(t *testing.T) {
	tests := []struct {
		name        string
		endpoint    string
		description string
	}{
		{
			name:        "ClientOperEndpoint",
			endpoint:    "/restconf/data/Cisco-IOS-XE-wireless-client-oper:client-oper-data",
			description: "Client operational data endpoint",
		},
		{
			name:        "MobilityOperDataEndpoint",
			endpoint:    "/restconf/data/Cisco-IOS-XE-wireless-client-oper:client-oper-data/mobility-oper-data",
			description: "Client mobility operational data endpoint",
		},
		{
			name:        "PolicyDataEndpoint",
			endpoint:    "/restconf/data/Cisco-IOS-XE-wireless-client-oper:client-oper-data/policy-data",
			description: "Client policy data endpoint",
		},
		{
			name:        "SisfDbMacEndpoint",
			endpoint:    "/restconf/data/Cisco-IOS-XE-wireless-client-oper:client-oper-data/sisf-db-mac",
			description: "Client SISF DB MAC endpoint",
		},
		{
			name:        "DcInfoEndpoint",
			endpoint:    "/restconf/data/Cisco-IOS-XE-wireless-client-oper:client-oper-data/dc-info",
			description: "Client DC info endpoint",
		},
		{
			name:        "CommonOperDataEndpoint",
			endpoint:    "/restconf/data/Cisco-IOS-XE-wireless-client-oper:client-oper-data/common-oper-data",
			description: "Client common operational data endpoint",
		},
		{
			name:        "Dot11OperDataEndpoint",
			endpoint:    "/restconf/data/Cisco-IOS-XE-wireless-client-oper:client-oper-data/dot11-oper-data",
			description: "Client 802.11 operational data endpoint",
		},
		{
			name:        "MmIfClientStatsEndpoint",
			endpoint:    "/restconf/data/Cisco-IOS-XE-wireless-client-oper:client-oper-data/mm-if-client-stats",
			description: "Client MM interface statistics endpoint",
		},
		{
			name:        "MmIfClientHistoryEndpoint",
			endpoint:    "/restconf/data/Cisco-IOS-XE-wireless-client-oper:client-oper-data/mm-if-client-history",
			description: "Client MM interface history endpoint",
		},
		{
			name:        "TrafficStatsEndpoint",
			endpoint:    "/restconf/data/Cisco-IOS-XE-wireless-client-oper:client-oper-data/traffic-stats",
			description: "Client traffic statistics endpoint",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testutils.EndpointValidationTest(t, tt.endpoint, tt.endpoint)
		})
	}
}

// =============================================================================
// 5. CONTEXT HANDLING TESTS
// =============================================================================

func TestClientOperContextHandling(t *testing.T) {
	// Test each client operation function with context handling
	t.Run("GetClientOper", func(t *testing.T) {
		testutils.TestContextHandling(t, func(ctx context.Context, client *wnc.Client) error {
			_, err := GetClientOper(client, ctx)
			return err
		})
	})

	t.Run("GetClientOperCommonOperData", func(t *testing.T) {
		testutils.TestContextHandling(t, func(ctx context.Context, client *wnc.Client) error {
			_, err := GetClientOperCommonOperData(client, ctx)
			return err
		})
	})

	t.Run("GetClientOperDot11OperData", func(t *testing.T) {
		testutils.TestContextHandling(t, func(ctx context.Context, client *wnc.Client) error {
			_, err := GetClientOperDot11OperData(client, ctx)
			return err
		})
	})

	t.Run("GetClientOperMobilityOperData", func(t *testing.T) {
		testutils.TestContextHandling(t, func(ctx context.Context, client *wnc.Client) error {
			_, err := GetClientOperMobilityOperData(client, ctx)
			return err
		})
	})

	t.Run("GetClientOperTrafficStats", func(t *testing.T) {
		testutils.TestContextHandling(t, func(ctx context.Context, client *wnc.Client) error {
			_, err := GetClientOperTrafficStats(client, ctx)
			return err
		})
	})

	t.Run("GetClientOperPolicyData", func(t *testing.T) {
		testutils.TestContextHandling(t, func(ctx context.Context, client *wnc.Client) error {
			_, err := GetClientOperPolicyData(client, ctx)
			return err
		})
	})

	t.Run("GetClientOperSisfDbMac", func(t *testing.T) {
		testutils.TestContextHandling(t, func(ctx context.Context, client *wnc.Client) error {
			_, err := GetClientOperSisfDbMac(client, ctx)
			return err
		})
	})

	t.Run("GetClientOperDcInfo", func(t *testing.T) {
		testutils.TestContextHandling(t, func(ctx context.Context, client *wnc.Client) error {
			_, err := GetClientOperDcInfo(client, ctx)
			return err
		})
	})

	t.Run("GetClientOperMmIfClientStats", func(t *testing.T) {
		testutils.TestContextHandling(t, func(ctx context.Context, client *wnc.Client) error {
			_, err := GetClientOperMmIfClientStats(client, ctx)
			return err
		})
	})

	t.Run("GetClientOperMmIfClientHistory", func(t *testing.T) {
		testutils.TestContextHandling(t, func(ctx context.Context, client *wnc.Client) error {
			_, err := GetClientOperMmIfClientHistory(client, ctx)
			return err
		})
	})
}
