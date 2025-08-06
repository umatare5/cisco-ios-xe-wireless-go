// Package rogue provides rogue access point detection operational data test functionality for the Cisco Wireless Network Controller API.
package rogue

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
	testutils "github.com/umatare5/cisco-ios-xe-wireless-go/internal/tests"
)

// getTestClient creates a test client using environment variables
func getTestClient(t *testing.T) *wnc.Client {
	return testutils.CreateTestClientFromEnv(t)
}

// =============================================================================
// 1. UNIT TESTS (Structure/Type Validation & JSON Serialization/Deserialization)
// =============================================================================

// TestRogueOperEndpoints tests that all Rogue operation endpoint constants are correctly defined
func TestRogueOperEndpoints(t *testing.T) {
	expectedEndpoints := map[string]string{
		"RogueOperBasePath":       "Cisco-IOS-XE-wireless-rogue-oper:rogue-oper-data",
		"RogueOperEndpoint":       "Cisco-IOS-XE-wireless-rogue-oper:rogue-oper-data",
		"RogueStatsEndpoint":      "Cisco-IOS-XE-wireless-rogue-oper:rogue-oper-data/rogue-stats",
		"RogueDataEndpoint":       "Cisco-IOS-XE-wireless-rogue-oper:rogue-oper-data/rogue-data",
		"RogueClientDataEndpoint": "Cisco-IOS-XE-wireless-rogue-oper:rogue-oper-data/rogue-client-data",
		"RldpStatsEndpoint":       "Cisco-IOS-XE-wireless-rogue-oper:rogue-oper-data/rldp-stats",
	}

	for name, expected := range expectedEndpoints {
		t.Run(name, func(t *testing.T) {
			actualEndpoints := map[string]string{
				"RogueOperBasePath":       RogueOperBasePath,
				"RogueOperEndpoint":       RogueOperEndpoint,
				"RogueStatsEndpoint":      RogueStatsEndpoint,
				"RogueDataEndpoint":       RogueDataEndpoint,
				"RogueClientDataEndpoint": RogueClientDataEndpoint,
				"RldpStatsEndpoint":       RldpStatsEndpoint,
			}

			actual, exists := actualEndpoints[name]
			if !exists {
				t.Errorf("Endpoint constant %s not found", name)
				return
			}
			if actual != expected {
				t.Errorf(wnc.EndpointMismatchErrorTemplate, name, expected, actual)
			}
		})
	}
}

// TestRogueOperDataStructures tests the basic structure of Rogue operational data types
func TestRogueOperDataStructures(t *testing.T) {
	tests := []struct {
		name     string
		jsonData string
		dataType interface{}
	}{
		{
			name: "RogueOperResponse",
			jsonData: `{
				"Cisco-IOS-XE-wireless-rogue-oper:rogue-oper-data": {
					"rogue-stats": {
						"rogue-count": 5,
						"rogue-enabled": true
					},
					"rogue-data": [
						{
							"rogue-key": {
								"mac-addr": "aa:bb:cc:dd:ee:ff"
							},
							"rogue-first-timestamp": "2024-01-01T00:00:00Z",
							"rogue-last-timestamp": "2024-01-01T01:00:00Z",
							"contained": false
						}
					],
					"rogue-client-data": [
						{
							"rogue-client-key": {
								"rogue-client-mac": "aa:bb:cc:dd:ee:00"
							},
							"rogue-client-first-timestamp": "2024-01-01T00:00:00Z",
							"rogue-client-last-timestamp": "2024-01-01T01:00:00Z",
							"contained": false
						}
					],
					"rldp-stats": {
						"num-rldp-started": 0,
						"connected": 0,
						"not-connected": 0
					}
				}
			}`,
			dataType: &RogueOperResponse{},
		},
		{
			name: "RogueStatsResponse",
			jsonData: `{
				"Cisco-IOS-XE-wireless-rogue-oper:rogue-stats": {
					"rogue-count": 5,
					"rogue-enabled": true
				}
			}`,
			dataType: &RogueStatsResponse{},
		},
		{
			name: "RogueDataResponse",
			jsonData: `{
				"Cisco-IOS-XE-wireless-rogue-oper:rogue-data": [
					{
						"rogue-key": {
							"mac-addr": "aa:bb:cc:dd:ee:ff"
						},
						"rogue-first-timestamp": "2024-01-01T00:00:00Z",
						"rogue-last-timestamp": "2024-01-01T01:00:00Z",
						"contained": false
					}
				]
			}`,
			dataType: &RogueDataResponse{},
		},
		{
			name: "RogueClientDataResponse",
			jsonData: `{
				"Cisco-IOS-XE-wireless-rogue-oper:rogue-client-data": [
					{
						"rogue-client-key": {
							"rogue-client-mac": "aa:bb:cc:dd:ee:00"
						},
						"rogue-client-first-timestamp": "2024-01-01T00:00:00Z",
						"rogue-client-last-timestamp": "2024-01-01T01:00:00Z",
						"contained": false
					}
				]
			}`,
			dataType: &RogueClientDataResponse{},
		},
		{
			name: "RldpStatsResponse",
			jsonData: `{
				"Cisco-IOS-XE-wireless-rogue-oper:rldp-stats": {
					"num-rldp-started": 0,
					"connected": 0,
					"not-connected": 0
				}
			}`,
			dataType: &RldpStatsResponse{},
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
// 2. TABLE-DRIVEN TEST PATTERNS
// =============================================================================

// Currently no table-driven tests specific to Rogue operations

// =============================================================================
// 3. FAIL-FAST ERROR DETECTION TESTS
// =============================================================================

// Currently no fail-fast error detection tests specific to Rogue operations

// =============================================================================
// 4. INTEGRATION TESTS (API Communication & Full Workflow Tests)
// =============================================================================

// TestGetRogueOper tests the GetRogueOper method
func TestGetRogueOper(t *testing.T) {
	client := getTestClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	t.Run("GetRogueOper", func(t *testing.T) {
		result, err := GetRogueOper(client, ctx)
		if err != nil {
			t.Logf("GetRogueOper failed (expected if Rogue detection not configured): %v", err)
			return
		}

		if result == nil {
			t.Error("Expected non-nil result")
			return
		}

		// Log the response for debugging
		t.Logf("GetRogueOper response received")

		// Save to file for analysis
		saveRogueTestData("rogue_oper_data", result)
	})

}

// TestGetRogueStats tests the GetRogueStats method
func TestGetRogueStats(t *testing.T) {
	client := getTestClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	t.Run("GetRogueStats", func(t *testing.T) {
		result, err := GetRogueStats(client, ctx)
		if err != nil {
			t.Logf("GetRogue* failed (expected if Rogue detection not configured): %v", err)
			return
		}

		if result == nil {
			t.Error("Expected non-nil result")
			return
		}

		t.Logf("GetRogueStats response received")
		saveRogueTestData("rogue_stats_data", result)
	})

}

// TestGetRogueData tests the GetRogueData method
func TestGetRogueData(t *testing.T) {
	client := getTestClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	t.Run("GetRogueData", func(t *testing.T) {
		result, err := GetRogueData(client, ctx)
		if err != nil {
			t.Logf("GetRogue* failed (expected if Rogue detection not configured): %v", err)
			return
		}

		if result == nil {
			t.Error("Expected non-nil result")
			return
		}

		t.Logf("GetRogueData response received")
		saveRogueTestData("rogue_data", result)
	})

}

// TestGetRogueClientData tests the GetRogueClientData method
func TestGetRogueClientData(t *testing.T) {
	client := getTestClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	t.Run("GetRogueClientData", func(t *testing.T) {
		result, err := GetRogueClientData(client, ctx)
		if err != nil {
			t.Logf("GetRogue* failed (expected if Rogue detection not configured): %v", err)
			return
		}

		if result == nil {
			t.Error("Expected non-nil result")
			return
		}

		t.Logf("GetRogueClientData response received")
		saveRogueTestData("rogue_client_data", result)
	})

}

// TestGetRldpStats tests the GetRldpStats method
func TestGetRldpStats(t *testing.T) {
	client := getTestClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	t.Run("GetRldpStats", func(t *testing.T) {
		result, err := GetRldpStats(client, ctx)
		if err != nil {
			t.Logf("GetRogue* failed (expected if Rogue detection not configured): %v", err)
			return
		}

		if result == nil {
			t.Error("Expected non-nil result")
			return
		}

		t.Logf("GetRldpStats response received")
		saveRogueTestData("rldp_stats_data", result)
	})

}

// TestRogueComprehensiveOperations tests all Rogue operations comprehensively
func TestRogueComprehensiveOperations(t *testing.T) {
	client := getTestClient(t)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	operations := map[string]func() (interface{}, error){
		"GetRogueOper": func() (interface{}, error) {
			return GetRogueOper(client, ctx)
		},
		"GetRogueStats": func() (interface{}, error) {
			return GetRogueStats(client, ctx)
		},
		"GetRogueData": func() (interface{}, error) {
			return GetRogueData(client, ctx)
		},
		"GetRogueClientData": func() (interface{}, error) {
			return GetRogueClientData(client, ctx)
		},
		"GetRldpStats": func() (interface{}, error) {
			return GetRldpStats(client, ctx)
		},
	}

	allResults := make(map[string]interface{})

	for operationName, operation := range operations {
		t.Run(operationName, func(t *testing.T) {
			start := time.Now()
			result, err := operation()
			duration := time.Since(start)

			if err != nil {
				t.Logf("%s failed after %v (may be expected if Rogue detection not configured): %v", operationName, duration, err)
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
	saveRogueTestData("rogue_comprehensive_test_results", allResults)
}

// =============================================================================
// 5. OTHER TESTS
// =============================================================================

// TestRogueOperClientInterfaceCompliance verifies that Client implements all Rogue methods
func TestRogueOperClientInterfaceCompliance(t *testing.T) {
	client := getTestClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Test that all methods exist and can be called (they will fail without real config, but should compile)
	t.Run("MethodExistence", func(t *testing.T) {
		// These calls will fail but should compile, proving the methods exist
		_, _ = GetRogueOper(client, ctx)
		_, _ = GetRogueStats(client, ctx)
		_, _ = GetRogueData(client, ctx)
		_, _ = GetRogueClientData(client, ctx)
		_, _ = GetRldpStats(client, ctx)
	})
}

func saveRogueTestData(filename string, data interface{}) {
	if data == nil {
		return
	}

	if err := testutils.SaveTestDataToFile(fmt.Sprintf("test_data_%s.json", filename), data); err != nil {
		fmt.Printf("Error saving test data for %s: %v\n", filename, err)
	} else {
		fmt.Printf("Test data saved to %s/test_data_%s.json\n", testutils.TestDataDir, filename)
	}
}

// TestRogueOperErrorHandling tests error handling for all rogue functions.
func TestRogueOperErrorHandling(t *testing.T) {
	t.Run("GetRogueOperWithNilClient", func(t *testing.T) {
		_, err := GetRogueOper(nil, context.Background())
		if err == nil || err.Error() != "invalid client configuration: client cannot be nil" {
			t.Errorf("Expected 'client is nil' error, got: %v", err)
		}
	})

	t.Run("GetRogueStatsWithNilClient", func(t *testing.T) {
		_, err := GetRogueStats(nil, context.Background())
		if err == nil || err.Error() != "invalid client configuration: client cannot be nil" {
			t.Errorf("Expected 'client is nil' error, got: %v", err)
		}
	})

	t.Run("GetRogueDataWithNilClient", func(t *testing.T) {
		_, err := GetRogueData(nil, context.Background())
		if err == nil || err.Error() != "invalid client configuration: client cannot be nil" {
			t.Errorf("Expected 'client is nil' error, got: %v", err)
		}
	})

	t.Run("GetRogueClientDataWithNilClient", func(t *testing.T) {
		_, err := GetRogueClientData(nil, context.Background())
		if err == nil || err.Error() != "invalid client configuration: client cannot be nil" {
			t.Errorf("Expected 'client is nil' error, got: %v", err)
		}
	})

	t.Run("GetRldpStatsWithNilClient", func(t *testing.T) {
		_, err := GetRldpStats(nil, context.Background())
		if err == nil || err.Error() != "invalid client configuration: client cannot be nil" {
			t.Errorf("Expected 'client is nil' error, got: %v", err)
		}
	})
}

// TestRogueOperContextHandling tests context handling for all rogue functions.
func TestRogueOperContextHandling(t *testing.T) {
	t.Run("GetRogueOperContextHandling", func(t *testing.T) {
		testutils.TestContextHandling(t, func(ctx context.Context, client *wnc.Client) error {
			_, err := GetRogueOper(client, ctx)
			return err
		})
	})

	t.Run("GetRogueStatsContextHandling", func(t *testing.T) {
		testutils.TestContextHandling(t, func(ctx context.Context, client *wnc.Client) error {
			_, err := GetRogueStats(client, ctx)
			return err
		})
	})

	t.Run("GetRogueDataContextHandling", func(t *testing.T) {
		testutils.TestContextHandling(t, func(ctx context.Context, client *wnc.Client) error {
			_, err := GetRogueData(client, ctx)
			return err
		})
	})

	t.Run("GetRogueClientDataContextHandling", func(t *testing.T) {
		testutils.TestContextHandling(t, func(ctx context.Context, client *wnc.Client) error {
			_, err := GetRogueClientData(client, ctx)
			return err
		})
	})

	t.Run("GetRldpStatsContextHandling", func(t *testing.T) {
		testutils.TestContextHandling(t, func(ctx context.Context, client *wnc.Client) error {
			_, err := GetRldpStats(client, ctx)
			return err
		})
	})
}
