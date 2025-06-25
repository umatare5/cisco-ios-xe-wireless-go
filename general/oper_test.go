// Package general provides general operational data test functionality for the Cisco Wireless Network Controller API.
package general

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	wnc "github.com/umatare5/cisco-xe-wireless-restconf-go"
	"github.com/umatare5/cisco-xe-wireless-restconf-go/internal/testutil"
)

// GeneralOperTestDataCollector holds test data for general operation functions
type GeneralOperTestDataCollector struct {
	Data map[string]interface{} `json:"general_oper_test_data"`
}

func newGeneralOperTestDataCollector() *GeneralOperTestDataCollector {
	return &GeneralOperTestDataCollector{
		Data: make(map[string]interface{}),
	}
}

func (collector *GeneralOperTestDataCollector) runTestAndCollectData(t *testing.T, testName string, testFunc func() (interface{}, error)) {
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

// getTestClient creates a test client using environment variables
func getTestClient(t *testing.T) *wnc.Client {
	return testutil.CreateTestClientFromEnv(t)
}

// =============================================================================
// 1. UNIT TESTS (Structure/Type Validation & JSON Serialization/Deserialization)
// =============================================================================

// TestGeneralOperDataStructures tests the basic structure of general operational data types
func TestGeneralOperDataStructures(t *testing.T) {
	tests := []struct {
		name     string
		jsonData string
		dataType interface{}
	}{
		{
			name: "GeneralOperResponse",
			jsonData: `{
				"Cisco-IOS-XE-wireless-general-oper:general-oper-data": {
					"mgmt-intf-data": {
						"intf-name": "Vlan100",
						"intf-type": "vlan",
						"intf-id": 100,
						"mgmt-ip": "192.168.1.10",
						"net-mask": "255.255.255.0",
						"mgmt-mac": "aa:bb:cc:dd:ee:ff"
					}
				}
			}`,
			dataType: &GeneralOperResponse{},
		},
		{
			name: "GeneralOperMgmtIntfDataResponse",
			jsonData: `{
				"Cisco-IOS-XE-wireless-general-oper:mgmt-intf-data": {
					"intf-name": "Vlan100",
					"intf-type": "vlan",
					"intf-id": 100,
					"mgmt-ip": "192.168.1.10",
					"net-mask": "255.255.255.0",
					"mgmt-mac": "aa:bb:cc:dd:ee:ff"
				}
			}`,
			dataType: &GeneralOperMgmtIntfDataResponse{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := json.Unmarshal([]byte(tt.jsonData), tt.dataType)
			if err != nil {
				t.Errorf("Failed to unmarshal %s: %v", tt.name, err)
			}

			_, err = json.Marshal(tt.dataType)
			if err != nil {
				t.Errorf("Failed to marshal %s: %v", tt.name, err)
			}
		})
	}
}

// =============================================================================
// 2. TABLE-DRIVEN TEST PATTERNS
// =============================================================================

// TestGeneralOperEndpoints tests general operation endpoints with table-driven approach
func TestGeneralOperEndpoints(t *testing.T) {
	client := getTestClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), testutil.DefaultTestTimeout)
	defer cancel()

	// Table-driven test cases for various general endpoints
	tests := []struct {
		name       string
		testFunc   func() (interface{}, error)
		shouldFail bool
	}{
		{
			name:       "GetGeneralOper",
			testFunc:   func() (interface{}, error) { return GetGeneralOper(client, ctx) },
			shouldFail: false,
		},
		{
			name:       "GetGeneralOperMgmtIntfData",
			testFunc:   func() (interface{}, error) { return GetGeneralOperMgmtIntfData(client, ctx) },
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

// TestGeneralOperFailFast tests fail-fast scenarios for general operations
func TestGeneralOperFailFast(t *testing.T) {
	// Test with nil client - expect error (not panic)
	t.Run("NilClient", func(t *testing.T) {
		ctx := context.Background()
		_, err := GetGeneralOper(nil, ctx)
		if err == nil {
			t.Fatal("Expected error with nil client, got none")
		}
		t.Logf("Correctly returned error with nil client: %v", err)
	})

	// Test with nil context - expect error (not panic)
	t.Run("NilContext", func(t *testing.T) {
		client := getTestClient(t)
		_, err := GetGeneralOper(client, nil)
		if err == nil {
			t.Fatal("Expected error with nil context, got none")
		}
		t.Logf("Correctly returned error with nil context: %v", err)
	})

	// Test with canceled context
	t.Run("CanceledContext", func(t *testing.T) {
		client := getTestClient(t)
		ctx, cancel := context.WithCancel(context.Background())
		cancel() // Cancel immediately

		_, err := GetGeneralOper(client, ctx)
		if err == nil {
			t.Fatal("Expected error with canceled context, got none")
		}
	})
}

// =============================================================================
// 4. INTEGRATION TESTS (API Endpoint, Real Controller)
// =============================================================================

// TestGeneralOperGetGeneralOper tests GetGeneralOper with real WNC data collection
func TestGeneralOperGetGeneralOper(t *testing.T) {
	client := getTestClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), testutil.DefaultTestTimeout)
	defer cancel()

	result, err := GetGeneralOper(client, ctx)
	if err != nil {
		t.Fatalf("GetGeneralOper failed: %v", err)
	}

	if result == nil {
		t.Fatal("GetGeneralOper returned nil result")
	}

	// Save result to JSON file
	filename := fmt.Sprintf("general_oper_data_%d.json", time.Now().Unix())
	if err := testutil.SaveTestDataToFile(filename, result); err != nil {
		t.Logf("Warning: Failed to save data to %s: %v", filename, err)
	} else {
		t.Logf("Data saved to %s", filename)
	}

	t.Logf("GetGeneralOper successful, collected general operational data")
}

// TestGeneralOperGetGeneralOperMgmtIntfData tests GetGeneralOperMgmtIntfData with real WNC data collection
func TestGeneralOperGetGeneralOperMgmtIntfData(t *testing.T) {
	client := getTestClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), testutil.DefaultTestTimeout)
	defer cancel()

	result, err := GetGeneralOperMgmtIntfData(client, ctx)
	if err != nil {
		t.Fatalf("GetGeneralOperMgmtIntfData failed: %v", err)
	}

	if result == nil {
		t.Fatal("GetGeneralOperMgmtIntfData returned nil result")
	}

	// Save result to JSON file
	filename := fmt.Sprintf("general_oper_mgmt_intf_data_%d.json", time.Now().Unix())
	if err := testutil.SaveTestDataToFile(filename, result); err != nil {
		t.Logf("Warning: Failed to save data to %s: %v", filename, err)
	} else {
		t.Logf("Data saved to %s", filename)
	}

	t.Logf("GetGeneralOperMgmtIntfData successful")
}

// TestGeneralOperCollectAllData runs all general operational tests and collects comprehensive data
func TestGeneralOperCollectAllData(t *testing.T) {
	client := getTestClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), testutil.DefaultTestTimeout)
	defer cancel()
	allData := make(map[string]interface{})

	// Collect data from all general operational endpoints
	tests := []struct {
		name string
		fn   func() (interface{}, error)
	}{
		{"GetGeneralOper", func() (interface{}, error) {
			return GetGeneralOper(client, ctx)
		}},
		{"GetGeneralOperMgmtIntfData", func() (interface{}, error) {
			return GetGeneralOperMgmtIntfData(client, ctx)
		}},
	}

	for _, test := range tests {
		result, err := test.fn()
		if err != nil {
			t.Logf("Warning: %s failed: %v", test.name, err)
			allData[test.name] = map[string]string{"error": err.Error()}
		} else {
			allData[test.name] = result
			t.Logf("Successfully collected data from %s", test.name)
		}
	}

	// Save all collected data to a comprehensive JSON file
	filename := fmt.Sprintf("general_oper_comprehensive_data_%d.json", time.Now().Unix())
	if err := testutil.SaveTestDataToFile(filename, allData); err != nil {
		t.Logf("Warning: Failed to save comprehensive data to %s: %v", filename, err)
	} else {
		t.Logf("Comprehensive data saved to %s", filename)
	}

	t.Logf("GeneralOper comprehensive test completed")
}

// =============================================================================
// 5. OTHER TESTS (Performance, Edge Cases, etc.)
// =============================================================================

// TestGeneralOperPerformance tests performance characteristics of general operations
func TestGeneralOperPerformance(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping performance test in short mode")
	}

	client := getTestClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), testutil.ExtendedTestTimeout)
	defer cancel()

	// Test concurrent requests
	t.Run("ConcurrentRequests", func(t *testing.T) {
		const numGoroutines = 3
		done := make(chan bool, numGoroutines)

		for i := 0; i < numGoroutines; i++ {
			go func() {
				_, err := GetGeneralOper(client, ctx)
				if err != nil {
					t.Logf("Concurrent request error (may be expected): %v", err)
				}
				done <- true
			}()
		}

		// Wait for all goroutines to complete
		for i := 0; i < numGoroutines; i++ {
			<-done
		}
	})
}
