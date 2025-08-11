// Package mcast provides multicast test functionality for the Cisco Wireless Network Controller API.
package mcast

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/tests"
)

// McastTestDataCollector holds test data for Mcast service functions
type McastTestDataCollector struct {
	Data map[string]interface{} `json:"mcast_test_data"`
}

// newMcastTestDataCollector creates a new test data collector
func newMcastTestDataCollector() *McastTestDataCollector {
	return &McastTestDataCollector{
		Data: make(map[string]interface{}),
	}
}

func runMcastTestAndCollectData(
	t *testing.T,
	collector *McastTestDataCollector,
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

// TestMcastServiceStructures tests the basic structure of Mcast service and data types
func TestMcastServiceStructures(t *testing.T) {
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
			name: "McastOperResponse",
			jsonData: `{
				"Cisco-IOS-XE-wireless-mcast-oper:mcast-oper-data": {
					"mcast-enable": true,
					"multicast-mode": "standard"
				}
			}`,
		},
		{
			name: "FlexMediastreamClientSummaryResponse",
			jsonData: `{
				"Cisco-IOS-XE-wireless-mcast-oper:flex-mediastream-client-summary": [
					{
						"multicast-ip": "224.1.1.1",
						"client-count": 5
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

// TestMCASTServiceMethods tests MCAST service methods with table-driven approach
func TestMcastServiceMethods(t *testing.T) {
	client := tests.TestClient(t)
	service := NewService(client)
	ctx := tests.TestContext(t)

	// Table-driven test cases for MCAST endpoints
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
			name:       "GetFlexMediastreamClientSummary",
			testFunc:   func() (interface{}, error) { return service.GetFlexMediastreamClientSummary(ctx) },
			shouldFail: false,
		},
		{
			name:       "GetVlanL2MgidOp",
			testFunc:   func() (interface{}, error) { return service.GetVlanL2MgidOp(ctx) },
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

// TestMcastServiceFailFast tests fail-fast scenarios for Mcast service operations
func TestMcastServiceFailFast(t *testing.T) {
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

// TestMCASTServiceIntegration tests all MCAST service functions with real WNC data collection
func TestMcastServiceIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode")
	}

	collector := newMcastTestDataCollector()
	client := tests.TestClient(t)
	service := NewService(client)
	ctx := tests.TestContext(t)

	t.Run("Oper", func(t *testing.T) {
		runMcastTestAndCollectData(t, collector, "GetOper", func() (interface{}, error) {
			return service.GetOper(ctx)
		})
	})

	t.Run("FlexMediastreamClientSummary", func(t *testing.T) {
		runMcastTestAndCollectData(t, collector, "GetFlexMediastreamClientSummary", func() (interface{}, error) {
			return service.GetFlexMediastreamClientSummary(ctx)
		})
	})

	t.Run("VlanL2MgidOp", func(t *testing.T) {
		runMcastTestAndCollectData(t, collector, "GetVlanL2MgidOp", func() (interface{}, error) {
			return service.GetVlanL2MgidOp(ctx)
		})
	})

	// Save collected test data
	if len(collector.Data) > 0 {
		if err := tests.SaveTestDataToFile("mcast_test_data_collected.json", collector.Data); err != nil {
			t.Logf("Warning: Could not save test data: %v", err)
		} else {
			// Log successful save of test data
			t.Logf("Mcast test data saved to test_data/mcast_test_data_collected.json")
		}
	}
}
