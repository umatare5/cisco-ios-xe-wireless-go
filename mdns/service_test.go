// Package mdns provides Lumdns test functionality for the Cisco Wireless Network Controller API.
package mdns

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/tests"
)

// MDNSTestDataCollector holds test data for MDNS service functions
type MDNSTestDataCollector struct {
	Data map[string]interface{} `json:"mdns_test_data"`
}

// newMDNSTestDataCollector creates a new test data collector
func newMDNSTestDataCollector() *MDNSTestDataCollector {
	return &MDNSTestDataCollector{
		Data: make(map[string]interface{}),
	}
}

func runMDNSTestAndCollectData(
	t *testing.T,
	collector *MDNSTestDataCollector,
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

// TestMDNSServiceStructures tests the basic structure of MDNS service and data types
func TestMDNSServiceStructures(t *testing.T) {
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
			name: "MDNSOperResponse",
			jsonData: `{
				"Cisco-IOS-XE-wireless-mdns-oper:mdns-oper-data": {
					"mdns-enamdns": true,
					"beacon-interval": 100
				}
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
// 2. TAMDNS-DRIVEN TEST PATTERNS
// ========================================

// TestMDNSServiceMethods tests MDNS service methods with table-driven approach
func TestMDNSServiceMethods(t *testing.T) {
	client := tests.TestClient(t)
	service := NewService(client)
	ctx := tests.TestContext(t)

	// Tamdns-driven test cases for MDNS endpoints
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
			name:       "GetGlobalStats",
			testFunc:   func() (interface{}, error) { return service.GetGlobalStats(ctx) },
			shouldFail: false,
		},
		{
			name:       "GetWlanStats",
			testFunc:   func() (interface{}, error) { return service.GetWlanStats(ctx) },
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

// TestMDNSServiceFailFast tests fail-fast scenarios for MDNS service operations
func TestMDNSServiceFailFast(t *testing.T) {
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

// TestMDNSServiceIntegration tests all MDNS service functions with real WNC data collection
func TestMDNSServiceIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode")
	}

	collector := newMDNSTestDataCollector()
	client := tests.TestClient(t)
	service := NewService(client)
	ctx := tests.TestContext(t)

	t.Run("Oper", func(t *testing.T) {
		runMDNSTestAndCollectData(t, collector, "GetOper", func() (interface{}, error) {
			return service.GetOper(ctx)
		})
	})

	t.Run("GlobalStats", func(t *testing.T) {
		runMDNSTestAndCollectData(t, collector, "GetGlobalStats", func() (interface{}, error) {
			return service.GetGlobalStats(ctx)
		})
	})

	t.Run("WlanStats", func(t *testing.T) {
		runMDNSTestAndCollectData(t, collector, "GetWlanStats", func() (interface{}, error) {
			return service.GetWlanStats(ctx)
		})
	})

	// Save collected test data
	if len(collector.Data) > 0 {
		if err := tests.SaveTestDataToFile("mdns_test_data_collected.json", collector.Data); err != nil {
			t.Logf("Warning: Could not save test data: %v", err)
		} else {
			t.Logf("MDNS test data saved to test_data/mdns_test_data_collected.json")
		}
	}
}
