// Package hyperlocation provides hyperlocation test functionality for the Cisco Wireless Network Controller API.
package hyperlocation

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/tests"
)

// HyperlocationTestDataCollector holds test data for Hyperlocation service functions
type HyperlocationTestDataCollector struct {
	Data map[string]interface{} `json:"hyperlocation_test_data"`
}

// newHyperlocationTestDataCollector creates a new test data collector
func newHyperlocationTestDataCollector() *HyperlocationTestDataCollector {
	return &HyperlocationTestDataCollector{
		Data: make(map[string]interface{}),
	}
}

func runHyperlocationTestAndCollectData(
	t *testing.T,
	collector *HyperlocationTestDataCollector,
	testName string,
	testFunc func() (interface{}, error),
) {
	data, err := testFunc()
	if err != nil {
		t.Logf("%s returned error: %v", testName, err)
		collector.Data[testName] = map[string]interface{}{
			"error":   err.Error(),
			"success": false,
		}
	} else {
		t.Logf("%s executed successfully", testName)
		collector.Data[testName] = map[string]interface{}{
			"data":    data,
			"success": true,
		}
	}
}

// ========================================
// 1. UNIT TESTS (Structure/Type Validation & JSON Serialization/Deserialization)
// ========================================

// TestHyperlocationServiceStructures tests the basic structure of Hyperlocation service and data types
func TestHyperlocationServiceStructures(t *testing.T) {
	client := tests.TestClient(t)
	service := NewService(client)

	if service.c == nil {
		t.Error("Service client should not be nil")
	}

	// Test JSON serialization/deserialization with sample data
	testCases := []struct {
		name     string
		jsonData string
	}{
		{
			name: "HyperlocationOperResponse",
			jsonData: `{
				"Cisco-IOS-XE-wireless-hyperlocation-oper:hyperlocation-oper-data": {
					"hyperlocation-enable": true,
					"precision-mode": "high"
				}
			}`,
		},
		{
			name: "HyperlocationProfilesResponse",
			jsonData: `{
				"Cisco-IOS-XE-wireless-hyperlocation-oper:ewlc-hyperlocation-profile": [
					{
						"profile-name": "default-profile",
						"accuracy": 1.0
					}
				]
			}`,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			var data interface{}
			err := json.Unmarshal([]byte(tt.jsonData), &data)
			if err != nil {
				t.Errorf("Failed to unmarshal %s: %v", tt.name, err)
			}

			_, err = json.Marshal(data)
			if err != nil {
				t.Errorf("Failed to marshal %s: %v", tt.name, err)
			}
		})
	}
}

// ========================================
// 2. TABLE-DRIVEN TEST PATTERNS
// ========================================

// TestHyperlocationServiceMethods tests Hyperlocation service methods with table-driven approach
func TestHyperlocationServiceMethods(t *testing.T) {
	client := tests.TestClient(t)
	service := NewService(client)
	ctx := tests.TestContext(t)

	// Table-driven test cases for Hyperlocation endpoints
	testCases := []struct {
		name       string
		testFunc   func() (interface{}, error)
		shouldFail bool
	}{
		{
			name:       "GetOper",
			testFunc:   func() (interface{}, error) { return service.GetOper(ctx) },
			shouldFail: false,
		},
		{
			name:       "GetProfiles",
			testFunc:   func() (interface{}, error) { return service.GetProfiles(ctx) },
			shouldFail: false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			data, err := tt.testFunc()

			if tt.shouldFail && err == nil {
				t.Errorf("Expected %s to fail, but it succeeded", tt.name)
			}

			if !tt.shouldFail && err != nil {
				t.Logf("%s returned error (may be expected in test environment): %v", tt.name, err)
			} else if !tt.shouldFail && data != nil {
				t.Logf("%s executed successfully", tt.name)
			}
		})
	}
}

// ========================================
// 3. FAIL-FAST ERROR DETECTION (t.Fatalf/t.Fatal)
// ========================================

// TestHyperlocationServiceFailFast tests fail-fast scenarios for Hyperlocation service operations
func TestHyperlocationServiceFailFast(t *testing.T) {
	// Test with nil client - expect error (not panic)
	t.Run("NilClient", func(t *testing.T) {
		service := NewService(nil)
		ctx := context.Background()
		_, err := service.GetOper(ctx)
		if err == nil {
			t.Fatal("Expected error with nil client, got none")
		}
		t.Logf("Correctly returned error with nil client: %v", err)
	})

	// Test with nil context - expect error (not panic)
	t.Run("NilContext", func(t *testing.T) {
		client := tests.TestClient(t)
		service := NewService(client)
		var nilCtx context.Context //nolint:SA1012 // Testing nil context behavior
		_, err := service.GetOper(nilCtx)
		if err == nil {
			t.Fatal("Expected error with nil context, got none")
		}
		t.Logf("Correctly returned error with nil context: %v", err)
	})

	// Test with canceled context
	t.Run("CanceledContext", func(t *testing.T) {
		client := tests.TestClient(t)
		service := NewService(client)
		ctx, cancel := context.WithCancel(context.Background())
		cancel() // Cancel immediately

		_, err := service.GetOper(ctx)
		if err == nil {
			t.Fatal("Expected error with canceled context, got none")
		}
	})
}

// ========================================
// 4. INTEGRATION TESTS (API Endpoint, Real Controller)
// ========================================

// TestHyperlocationServiceIntegration tests all Hyperlocation service functions with real WNC data collection
func TestHyperlocationServiceIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode")
	}

	collector := newHyperlocationTestDataCollector()
	client := tests.TestClient(t)
	service := NewService(client)
	ctx := tests.TestContext(t)

	t.Run("Oper", func(t *testing.T) {
		runHyperlocationTestAndCollectData(t, collector, "GetOper", func() (interface{}, error) {
			return service.GetOper(ctx)
		})
	})

	t.Run("Profiles", func(t *testing.T) {
		runHyperlocationTestAndCollectData(t, collector, "GetProfiles", func() (interface{}, error) {
			return service.GetProfiles(ctx)
		})
	})

	// Save collected test data
	if len(collector.Data) > 0 {
		if err := tests.SaveTestDataToFile("hyperlocation_test_data_collected.json", collector.Data); err != nil {
			t.Logf("Warning: Could not save test data: %v", err)
		} else {
			t.Logf("Hyperlocation test data saved to test_data/hyperlocation_test_data_collected.json")
		}
	}
}
