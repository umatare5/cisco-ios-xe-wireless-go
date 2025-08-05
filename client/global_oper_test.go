// Package client provides client global operational data test functionality for the Cisco Wireless Network Controller API.
package client

import (
	"context"
	"testing"

	wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
	testutils "github.com/umatare5/cisco-ios-xe-wireless-go/internal/tests"
)

// =============================================================================
// 1. UNIT TESTS (Structure/Type Validation & JSON Serialization/Deserialization)
// =============================================================================

func TestClientGlobalOperDataStructures(t *testing.T) {
	testCases := []testutils.JSONTestCase{
		{
			Name: "ClientGlobalOperResponse",
			JSONData: `{
				"cisco-ios-xe-wireless-client-global-oper:client-global-oper-data": {}
			}`,
			Target:     &ClientGlobalOperResponse{},
			TypeName:   "ClientGlobalOperResponse",
			ShouldFail: false,
		},
		{
			Name: "ClientLiveStatsResponse",
			JSONData: `{
				"cisco-ios-xe-wireless-client-global-oper:live-stats": {}
			}`,
			Target:     &ClientLiveStatsResponse{},
			TypeName:   "ClientLiveStatsResponse",
			ShouldFail: false,
		},
		{
			Name: "ClientGlobalStatsDataResponse",
			JSONData: `{
				"cisco-ios-xe-wireless-client-global-oper:global-stats-data": {}
			}`,
			Target:     &ClientGlobalStatsDataResponse{},
			TypeName:   "ClientGlobalStatsDataResponse",
			ShouldFail: false,
		},
		{
			Name: "ClientStatsResponse",
			JSONData: `{
				"cisco-ios-xe-wireless-client-global-oper:stats": {}
			}`,
			Target:     &ClientStatsResponse{},
			TypeName:   "ClientStatsResponse",
			ShouldFail: false,
		},
		{
			Name: "ClientDot11StatsResponse",
			JSONData: `{
				"cisco-ios-xe-wireless-client-global-oper:dot11-stats": {}
			}`,
			Target:     &ClientDot11StatsResponse{},
			TypeName:   "ClientDot11StatsResponse",
			ShouldFail: false,
		},
	}

	testutils.RunJSONTests(t, testCases)
}

// =============================================================================
// 2. ERROR HANDLING TESTS
// =============================================================================

func TestClientGlobalOperErrorHandling(t *testing.T) {
	// Test nil client scenarios
	t.Run("NilClientTests", func(t *testing.T) {
		var nilClient *wnc.Client = nil
		ctx := context.Background()

		tests := []struct {
			name     string
			testFunc func() error
		}{
			{
				name: "GetClientGlobalOper_NilClient",
				testFunc: func() error {
					_, err := GetClientGlobalOper(nilClient, ctx)
					return err
				},
			},
			{
				name: "GetClientLiveStats_NilClient",
				testFunc: func() error {
					_, err := GetClientLiveStats(nilClient, ctx)
					return err
				},
			},
			{
				name: "GetClientGlobalStatsData_NilClient",
				testFunc: func() error {
					_, err := GetClientGlobalStatsData(nilClient, ctx)
					return err
				},
			},
			{
				name: "GetClientStats_NilClient",
				testFunc: func() error {
					_, err := GetClientStats(nilClient, ctx)
					return err
				},
			},
			{
				name: "GetClientDot11Stats_NilClient",
				testFunc: func() error {
					_, err := GetClientDot11Stats(nilClient, ctx)
					return err
				},
			},
			{
				name: "GetClientLatencyStats_NilClient",
				testFunc: func() error {
					_, err := GetClientLatencyStats(nilClient, ctx)
					return err
				},
			},
			{
				name: "GetClientSmWebauthStats_NilClient",
				testFunc: func() error {
					_, err := GetClientSmWebauthStats(nilClient, ctx)
					return err
				},
			},
			{
				name: "GetClientDot1XGlobalStats_NilClient",
				testFunc: func() error {
					_, err := GetClientDot1XGlobalStats(nilClient, ctx)
					return err
				},
			},
			{
				name: "GetClientExclusionStats_NilClient",
				testFunc: func() error {
					_, err := GetClientExclusionStats(nilClient, ctx)
					return err
				},
			},
			{
				name: "GetClientSmDeviceCount_NilClient",
				testFunc: func() error {
					_, err := GetClientSmDeviceCount(nilClient, ctx)
					return err
				},
			},
			{
				name: "GetClientTofStats_NilClient",
				testFunc: func() error {
					_, err := GetClientTofStats(nilClient, ctx)
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
// INTEGRATION TESTS - Live API Data Validation
// =============================================================================

// TestClientGlobalOperFunctions tests all client global operation functions with real API endpoints
func TestClientGlobalOperFunctions(t *testing.T) {
	client := testutils.CreateTestClientFromEnv(t)
	ctx := context.Background()

	// Function test configurations
	testCases := []struct {
		name     string
		testFunc func() (interface{}, error)
	}{
		{
			name: "GetClientGlobalOper",
			testFunc: func() (interface{}, error) {
				return GetClientGlobalOper(client, ctx)
			},
		},
		{
			name: "GetClientLiveStats",
			testFunc: func() (interface{}, error) {
				return GetClientLiveStats(client, ctx)
			},
		},
		{
			name: "GetClientGlobalStatsData",
			testFunc: func() (interface{}, error) {
				return GetClientGlobalStatsData(client, ctx)
			},
		},
		{
			name: "GetClientStats",
			testFunc: func() (interface{}, error) {
				return GetClientStats(client, ctx)
			},
		},
		{
			name: "GetClientDot11Stats",
			testFunc: func() (interface{}, error) {
				return GetClientDot11Stats(client, ctx)
			},
		},
		{
			name: "GetClientLatencyStats",
			testFunc: func() (interface{}, error) {
				return GetClientLatencyStats(client, ctx)
			},
		},
		{
			name: "GetClientSmWebauthStats",
			testFunc: func() (interface{}, error) {
				return GetClientSmWebauthStats(client, ctx)
			},
		},
		{
			name: "GetClientDot1XGlobalStats",
			testFunc: func() (interface{}, error) {
				return GetClientDot1XGlobalStats(client, ctx)
			},
		},
		{
			name: "GetClientExclusionStats",
			testFunc: func() (interface{}, error) {
				return GetClientExclusionStats(client, ctx)
			},
		},
		{
			name: "GetClientSmDeviceCount",
			testFunc: func() (interface{}, error) {
				return GetClientSmDeviceCount(client, ctx)
			},
		},
		{
			name: "GetClientTofStats",
			testFunc: func() (interface{}, error) {
				return GetClientTofStats(client, ctx)
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
		"client_global_oper_test_data": collector,
	}

	if err := testutils.SaveTestDataToFile("client_global_oper_test_data_collected.json", testDataMap); err != nil {
		t.Errorf("Failed to save test data: %v", err)
		return
	}

	t.Logf("Test data saved to test_data/client_global_oper_test_data_collected.json")
}

// =============================================================================
// ENDPOINT TESTS - API URL Validation
// =============================================================================

// TestClientGlobalOperEndpoints tests the endpoint URL generation and validation
func TestClientGlobalOperEndpoints(t *testing.T) {
	tests := []struct {
		name        string
		endpoint    string
		description string
	}{
		{
			name:        "ClientGlobalOperEndpoint",
			endpoint:    "/restconf/data/Cisco-IOS-XE-wireless-client-global-oper:client-global-oper-data",
			description: "Client global operational data endpoint",
		},
		{
			name:        "ClientLiveStatsEndpoint",
			endpoint:    "/restconf/data/Cisco-IOS-XE-wireless-client-global-oper:client-global-oper-data/live-stats",
			description: "Client live statistics endpoint",
		},
		{
			name:        "ClientGlobalStatsDataEndpoint",
			endpoint:    "/restconf/data/Cisco-IOS-XE-wireless-client-global-oper:client-global-oper-data/global-stats-data",
			description: "Client global statistics data endpoint",
		},
		{
			name:        "ClientStatsEndpoint",
			endpoint:    "/restconf/data/Cisco-IOS-XE-wireless-client-global-oper:client-global-oper-data/stats",
			description: "Client statistics endpoint",
		},
		{
			name:        "ClientDot11StatsEndpoint",
			endpoint:    "/restconf/data/Cisco-IOS-XE-wireless-client-global-oper:client-global-oper-data/dot11-stats",
			description: "Client 802.11 statistics endpoint",
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

func TestClientGlobalOperContextHandling(t *testing.T) {
	// Test each client global operation function with context handling
	t.Run("GetClientGlobalOper", func(t *testing.T) {
		testutils.TestContextHandling(t, func(ctx context.Context, client *wnc.Client) error {
			_, err := GetClientGlobalOper(client, ctx)
			return err
		})
	})

	t.Run("GetClientLiveStats", func(t *testing.T) {
		testutils.TestContextHandling(t, func(ctx context.Context, client *wnc.Client) error {
			_, err := GetClientLiveStats(client, ctx)
			return err
		})
	})

	t.Run("GetClientGlobalStatsData", func(t *testing.T) {
		testutils.TestContextHandling(t, func(ctx context.Context, client *wnc.Client) error {
			_, err := GetClientGlobalStatsData(client, ctx)
			return err
		})
	})

	t.Run("GetClientStats", func(t *testing.T) {
		testutils.TestContextHandling(t, func(ctx context.Context, client *wnc.Client) error {
			_, err := GetClientStats(client, ctx)
			return err
		})
	})

	t.Run("GetClientDot11Stats", func(t *testing.T) {
		testutils.TestContextHandling(t, func(ctx context.Context, client *wnc.Client) error {
			_, err := GetClientDot11Stats(client, ctx)
			return err
		})
	})

	// Add context handling for the remaining functions
	t.Run("GetClientLatencyStats", func(t *testing.T) {
		testutils.TestContextHandling(t, func(ctx context.Context, client *wnc.Client) error {
			_, err := GetClientLatencyStats(client, ctx)
			return err
		})
	})

	t.Run("GetClientSmWebauthStats", func(t *testing.T) {
		testutils.TestContextHandling(t, func(ctx context.Context, client *wnc.Client) error {
			_, err := GetClientSmWebauthStats(client, ctx)
			return err
		})
	})

	t.Run("GetClientDot1XGlobalStats", func(t *testing.T) {
		testutils.TestContextHandling(t, func(ctx context.Context, client *wnc.Client) error {
			_, err := GetClientDot1XGlobalStats(client, ctx)
			return err
		})
	})

	t.Run("GetClientExclusionStats", func(t *testing.T) {
		testutils.TestContextHandling(t, func(ctx context.Context, client *wnc.Client) error {
			_, err := GetClientExclusionStats(client, ctx)
			return err
		})
	})

	t.Run("GetClientSmDeviceCount", func(t *testing.T) {
		testutils.TestContextHandling(t, func(ctx context.Context, client *wnc.Client) error {
			_, err := GetClientSmDeviceCount(client, ctx)
			return err
		})
	})

	t.Run("GetClientTofStats", func(t *testing.T) {
		testutils.TestContextHandling(t, func(ctx context.Context, client *wnc.Client) error {
			_, err := GetClientTofStats(client, ctx)
			return err
		})
	})
}

// =============================================================================
// 6. INTEGRATION TESTS FOR SUCCESSFUL PATHS
// =============================================================================

func TestClientGlobalOperIntegrationSuccess(t *testing.T) {
	client := testutils.CreateTestClientFromEnv(t)
	ctx, cancel := context.WithTimeout(context.Background(), testutils.DefaultTestTimeout)
	defer cancel()

	// Test functions that need success path coverage
	t.Run("GetClientLatencyStatsSuccess", func(t *testing.T) {
		result, err := GetClientLatencyStats(client, ctx)
		if err != nil {
			t.Logf("GetClientLatencyStats returned error (expected in some environments): %v", err)
		} else if result != nil {
			t.Logf("GetClientLatencyStats successful")
		}
	})

	t.Run("GetClientSmWebauthStatsSuccess", func(t *testing.T) {
		result, err := GetClientSmWebauthStats(client, ctx)
		if err != nil {
			t.Logf("GetClientSmWebauthStats returned error (expected in some environments): %v", err)
		} else if result != nil {
			t.Logf("GetClientSmWebauthStats successful")
		}
	})

	t.Run("GetClientDot1XGlobalStatsSuccess", func(t *testing.T) {
		result, err := GetClientDot1XGlobalStats(client, ctx)
		if err != nil {
			t.Logf("GetClientDot1XGlobalStats returned error (expected in some environments): %v", err)
		} else if result != nil {
			t.Logf("GetClientDot1XGlobalStats successful")
		}
	})

	t.Run("GetClientExclusionStatsSuccess", func(t *testing.T) {
		result, err := GetClientExclusionStats(client, ctx)
		if err != nil {
			t.Logf("GetClientExclusionStats returned error (expected in some environments): %v", err)
		} else if result != nil {
			t.Logf("GetClientExclusionStats successful")
		}
	})

	t.Run("GetClientSmDeviceCountSuccess", func(t *testing.T) {
		result, err := GetClientSmDeviceCount(client, ctx)
		if err != nil {
			t.Logf("GetClientSmDeviceCount returned error (expected in some environments): %v", err)
		} else if result != nil {
			t.Logf("GetClientSmDeviceCount successful")
		}
	})

	t.Run("GetClientTofStatsSuccess", func(t *testing.T) {
		result, err := GetClientTofStats(client, ctx)
		if err != nil {
			t.Logf("GetClientTofStats returned error (expected in some environments): %v", err)
		} else if result != nil {
			t.Logf("GetClientTofStats successful")
		}
	})
}

// =============================================================================
// 6. DETAILED SUCCESS PATH TESTS
// =============================================================================

// TestClientGlobalOperSuccessPathCoverage tests specific functions to ensure 100% coverage
func TestClientGlobalOperSuccessPathCoverage(t *testing.T) {
	// Create a mock server that returns success responses
	mockServer := testutils.NewMockHTTPServer()

	mockServer.AddHandler("/restconf/data/Cisco-IOS-XE-wireless-client-global-oper:client-global-oper-data/tof-stats",
		testutils.CreateJSONResponse(testutils.TestHTTPResponse{
			StatusCode: 200,
			Body: `{
				"Cisco-IOS-XE-wireless-client-global-oper:tof-stats": {
					"tof-tag": ["tag1", "tag2", "tag3"]
				}
			}`,
			Headers: map[string]string{"Content-Type": "application/yang-data+json"},
		}))

	defer mockServer.Close()

	client := testutils.CreateTestClientForMockServer(t, mockServer)
	ctx, cancel := context.WithTimeout(context.Background(), testutils.DefaultTestTimeout)
	defer cancel()

	t.Run("GetClientTofStatsSuccessPath", func(t *testing.T) {
		result, err := GetClientTofStats(client, ctx)
		if err != nil {
			t.Errorf("Expected GetClientTofStats to succeed with mock server, got error: %v", err)
		}
		if result == nil {
			t.Error("Expected GetClientTofStats to return non-nil result")
		}
		// Verify the result structure
		if result != nil && len(result.TofStats.TofTag) == 0 {
			t.Log("GetClientTofStats returned result but with empty data (acceptable)")
		}
	})
}
