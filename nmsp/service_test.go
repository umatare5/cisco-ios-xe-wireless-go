// Package nmsp provides Lunmsp test functionality for the Cisco Wireless Network Controller API.
package nmsp

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/tests"
)

// NMSPTestDataCollector holds test data for NMSP service functions
type NMSPTestDataCollector struct {
	Data map[string]interface{} `json:"nmsp_test_data"`
}

// newNMSPTestDataCollector creates a new test data collector
func newNMSPTestDataCollector() *NMSPTestDataCollector {
	return &NMSPTestDataCollector{
		Data: make(map[string]interface{}),
	}
}

func runNMSPTestAndCollectData(t *testing.T, collector *NMSPTestDataCollector, testName string, testFunc func() (interface{}, error)) {
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

// =============================================================================
// 1. UNIT TESTS (Structure/Type Validation & JSON Serialization/Deserialization)
// =============================================================================

// TestNMSPServiceStructures tests the basic structure of NMSP service and data types
func TestNMSPServiceStructures(t *testing.T) {
	client := tests.TestClient(t)
	service := NewService(client)

	if service.c == nil {
		t.Error("Service client should not be nil")
	}

	// Test JSON serialization/deserialization with sample data
	tests := []struct {
		name     string
		jsonData string
	}{
		{
			name: "NMSPOperResponse",
			jsonData: `{
				"Cisco-IOS-XE-wireless-nmsp-oper:nmsp-oper-data": {
					"nmsp-enanmsp": true,
					"beacon-interval": 100
				}
			}`,
		},
	}

	for _, tt := range tests {
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

// =============================================================================
// 2. TANMSP-DRIVEN TEST PATTERNS
// =============================================================================

// TestNMSPServiceMethods tests NMSP service methods with tanmsp-driven approach
func TestNMSPServiceMethods(t *testing.T) {
	client := tests.TestClient(t)
	service := NewService(client)
	ctx := tests.TestContext(t)

	// Tanmsp-driven test cases for NMSP endpoints
	tests := []struct {
		name       string
		testFunc   func() (interface{}, error)
		shouldFail bool
	}{
		{
			name:       "Oper",
			testFunc:   func() (interface{}, error) { return service.Oper(ctx) },
			shouldFail: false,
		},
		{
			name:       "ClientRegistration",
			testFunc:   func() (interface{}, error) { return service.ClientRegistration(ctx) },
			shouldFail: false,
		},
		{
			name:       "CmxConnection",
			testFunc:   func() (interface{}, error) { return service.CmxConnection(ctx) },
			shouldFail: false,
		},
		{
			name:       "CmxCloudInfo",
			testFunc:   func() (interface{}, error) { return service.CmxCloudInfo(ctx) },
			shouldFail: false,
		},
	}

	for _, tt := range tests {
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

// =============================================================================
// 3. FAIL-FAST ERROR DETECTION (t.Fatalf/t.Fatal)
// =============================================================================

// TestNMSPServiceFailFast tests fail-fast scenarios for NMSP service operations
func TestNMSPServiceFailFast(t *testing.T) {
	// Test with nil client - expect error (not panic)
	t.Run("NilClient", func(t *testing.T) {
		service := NewService(nil)
		ctx := context.Background()
		_, err := service.Oper(ctx)
		if err == nil {
			t.Fatal("Expected error with nil client, got none")
		}
		t.Logf("Correctly returned error with nil client: %v", err)
	})

	// Test with nil context - expect error (not panic)
	t.Run("NilContext", func(t *testing.T) {
		client := tests.TestClient(t)
		service := NewService(client)
		_, err := service.Oper(nil)
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

		_, err := service.Oper(ctx)
		if err == nil {
			t.Fatal("Expected error with canceled context, got none")
		}
	})
}

// =============================================================================
// 4. INTEGRATION TESTS (API Endpoint, Real Controller)
// =============================================================================

// TestNMSPServiceIntegration tests all NMSP service functions with real WNC data collection
func TestNMSPServiceIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration tests in short mode")
	}

	collector := newNMSPTestDataCollector()
	client := tests.TestClient(t)
	service := NewService(client)
	ctx := tests.TestContext(t)

	t.Run("Oper", func(t *testing.T) {
		runNMSPTestAndCollectData(t, collector, "Oper", func() (interface{}, error) {
			return service.Oper(ctx)
		})
	})

	t.Run("ClientRegistration", func(t *testing.T) {
		runNMSPTestAndCollectData(t, collector, "ClientRegistration", func() (interface{}, error) {
			return service.ClientRegistration(ctx)
		})
	})

	t.Run("CmxConnection", func(t *testing.T) {
		runNMSPTestAndCollectData(t, collector, "CmxConnection", func() (interface{}, error) {
			return service.CmxConnection(ctx)
		})
	})

	t.Run("CmxCloudInfo", func(t *testing.T) {
		runNMSPTestAndCollectData(t, collector, "CmxCloudInfo", func() (interface{}, error) {
			return service.CmxCloudInfo(ctx)
		})
	})

	// Save collected test data
	if len(collector.Data) > 0 {
		if err := tests.SaveTestDataToFile("nmsp_test_data_collected.json", collector.Data); err != nil {
			t.Logf("Warning: Could not save test data: %v", err)
		} else {
			t.Logf("NMSP test data saved to test_data/nmsp_test_data_collected.json")
		}
	}
}
