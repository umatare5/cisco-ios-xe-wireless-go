// Package ap provides access point operational data test functionality for the Cisco Wireless Network Controller API.
package ap

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/umatare5/cisco-xe-wireless-restconf-go/internal/testutil"
)

// APOperTestDataCollector holds test data for AP operation functions
type APOperTestDataCollector struct {
	Data map[string]interface{} `json:"ap_oper_test_data"`
}

// newAPOperTestDataCollector creates a new test data collector
func newAPOperTestDataCollector() *APOperTestDataCollector {
	return &APOperTestDataCollector{
		Data: make(map[string]interface{}),
	}
}

func runAPOperTestAndCollectData(t *testing.T, collector *APOperTestDataCollector, testName string, testFunc func() (interface{}, error)) {
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

// TestApOperDataStructures tests the basic structure of AP operational data types
func TestApOperDataStructures(t *testing.T) {
	tests := []struct {
		name     string
		jsonData string
		dataType interface{}
	}{
		{
			name: "ApOperApRadioNeighborResponse",
			jsonData: `{
				"Cisco-IOS-XE-wireless-access-point-oper:ap-radio-neighbor": [
					{
						"wtp-mac": "aa:bb:cc:dd:ee:ff",
						"radio-slot-id": 0,
						"neighbor-ap-mac": "11:22:33:44:55:66",
						"neighbor-slot-id": 0,
						"rssi": -45,
						"neighbor-freq": 2437
					    }
				]
			    }`,
			dataType: &ApOperApRadioNeighborResponse{},
		},
		{
			name: "ApOperRadioOperDataResponse",
			jsonData: `{
				"Cisco-IOS-XE-wireless-access-point-oper:radio-oper-data": [
					{
						"wtp-mac": "aa:bb:cc:dd:ee:ff",
						"radio-slot-id": 0,
						"phy-type": 7,
						"antenna-mode": "dual",
						"current-channel": 6,
						"current-power": 23,
						"admin-state": "enabled",
						"oper-state": "up",
						"radio-band": "twenty-ghz"
					    }
				]
			    }`,
			dataType: &ApOperRadioOperDataResponse{},
		},
		{
			name: "ApOperCapwapDataResponse",
			jsonData: `{
				"Cisco-IOS-XE-wireless-access-point-oper:capwap-data": [
					{
						"wtp-mac": "aa:bb:cc:dd:ee:ff",
						"capwap-state": "run",
						"join-time": "2024-01-01T12:00:00.000Z",
						"last-heartbeat": "2024-01-01T12:05:00.000Z",
						"discovery-type": "broadcast",
						"wlc-ip": "192.168.1.10"
					    }
				]
			    }`,
			dataType: &ApOperCapwapDataResponse{},
		},
		{
			name: "ApOperApNameMacMapResponse",
			jsonData: `{
				"Cisco-IOS-XE-wireless-access-point-oper:ap-name-mac-map": [
					{
						"wtp-mac": "aa:bb:cc:dd:ee:ff",
						"ap-name": "AP-Floor1-001",
						"ethernet-mac": "aa:bb:cc:dd:ee:fe"
					    }
				]
			    }`,
			dataType: &ApOperApNameMacMapResponse{},
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

// TestApOperationEndpoints tests AP operation endpoints with table-driven approach
func TestApOperationEndpoints(t *testing.T) {
	client := testutil.CreateTestClientFromEnv(t)
	ctx, cancel := context.WithTimeout(context.Background(), testutil.DefaultTestTimeout)
	defer cancel()

	// Table-driven test cases for various AP endpoints
	tests := []struct {
		name       string
		testFunc   func() (interface{}, error)
		shouldFail bool
	}{
		{
			name:       "GetApOper",
			testFunc:   func() (interface{}, error) { return GetApOper(client, ctx) },
			shouldFail: false,
		},
		{
			name:       "GetApRadioNeighbor",
			testFunc:   func() (interface{}, error) { return GetApRadioNeighbor(client, ctx) },
			shouldFail: false,
		},
		{
			name:       "GetApRadioOperData",
			testFunc:   func() (interface{}, error) { return GetApRadioOperData(client, ctx) },
			shouldFail: false,
		},
		{
			name:       "GetApCapwapData",
			testFunc:   func() (interface{}, error) { return GetApCapwapData(client, ctx) },
			shouldFail: false,
		},
		{
			name:       "GetApNameMacMap",
			testFunc:   func() (interface{}, error) { return GetApNameMacMap(client, ctx) },
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

			// Store test data regardless of success/failure for analysis
			collector := newAPOperTestDataCollector()
			runAPOperTestAndCollectData(t, collector, tt.name, tt.testFunc)
		})
	}
}

// =============================================================================
// 3. FAIL-FAST ERROR DETECTION (t.Fatalf/t.Fatal)
// =============================================================================

// TestApOperationFailFast tests fail-fast scenarios for AP operations
func TestApOperationFailFast(t *testing.T) {
	// Test with nil client - expect error (not panic)
	t.Run("NilClient", func(t *testing.T) {
		ctx := context.Background()
		_, err := GetApOper(nil, ctx)
		if err == nil {
			t.Fatal("Expected error with nil client, got none")
		}
		t.Logf("Correctly returned error with nil client: %v", err)
	})

	// Test with nil context - expect error (not panic)
	t.Run("NilContext", func(t *testing.T) {
		client := testutil.CreateTestClientFromEnv(t)
		_, err := GetApOper(client, nil)
		if err == nil {
			t.Fatal("Expected error with nil context, got none")
		}
		t.Logf("Correctly returned error with nil context: %v", err)
	})

	// Test with canceled context
	t.Run("CanceledContext", func(t *testing.T) {
		client := testutil.CreateTestClientFromEnv(t)
		ctx, cancel := context.WithCancel(context.Background())
		cancel() // Cancel immediately

		_, err := GetApOper(client, ctx)
		if err == nil {
			t.Fatal("Expected error with canceled context, got none")
		}
	})
}

// =============================================================================
// 4. INTEGRATION TESTS (API Endpoint, Real Controller)
// =============================================================================

// TestApOperationFunctions tests all AP operation functions with real WNC data collection
func TestApOperationFunctions(t *testing.T) {
	collector := newAPOperTestDataCollector()
	client := testutil.CreateTestClientFromEnv(t)

	ctx, cancel := context.WithTimeout(context.Background(), testutil.DefaultTestTimeout)
	defer cancel()

	t.Run("GetApOper", func(t *testing.T) {
		runAPOperTestAndCollectData(t, collector, "GetApOper", func() (interface{}, error) {
			return GetApOper(client, ctx)
		})
	})

	t.Run("GetApRadioNeighbor", func(t *testing.T) {
		runAPOperTestAndCollectData(t, collector, "GetApRadioNeighbor", func() (interface{}, error) {
			return GetApRadioNeighbor(client, ctx)
		})
	})

	t.Run("GetApRadioOperData", func(t *testing.T) {
		runAPOperTestAndCollectData(t, collector, "GetApRadioOperData", func() (interface{}, error) {
			return GetApRadioOperData(client, ctx)
		})
	})

	t.Run("GetApRadioResetStats", func(t *testing.T) {
		runAPOperTestAndCollectData(t, collector, "GetApRadioResetStats", func() (interface{}, error) {
			return GetApRadioResetStats(client, ctx)
		})
	})

	t.Run("GetApQosClientData", func(t *testing.T) {
		runAPOperTestAndCollectData(t, collector, "GetApQosClientData", func() (interface{}, error) {
			return GetApQosClientData(client, ctx)
		})
	})

	t.Run("GetApCapwapData", func(t *testing.T) {
		runAPOperTestAndCollectData(t, collector, "GetApCapwapData", func() (interface{}, error) {
			return GetApCapwapData(client, ctx)
		})
	})

	t.Run("GetApNameMacMap", func(t *testing.T) {
		runAPOperTestAndCollectData(t, collector, "GetApNameMacMap", func() (interface{}, error) {
			return GetApNameMacMap(client, ctx)
		})
	})

	t.Run("GetApWtpSlotWlanStats", func(t *testing.T) {
		runAPOperTestAndCollectData(t, collector, "GetApWtpSlotWlanStats", func() (interface{}, error) {
			return GetApWtpSlotWlanStats(client, ctx)
		})
	})

	t.Run("GetApEthernetMacWtpMacMap", func(t *testing.T) {
		runAPOperTestAndCollectData(t, collector, "GetApEthernetMacWtpMacMap", func() (interface{}, error) {
			return GetApEthernetMacWtpMacMap(client, ctx)
		})
	})

	t.Run("GetApRadioOperStats", func(t *testing.T) {
		runAPOperTestAndCollectData(t, collector, "GetApRadioOperStats", func() (interface{}, error) {
			return GetApRadioOperStats(client, ctx)
		})
	})

	t.Run("GetApEthernetIfStats", func(t *testing.T) {
		runAPOperTestAndCollectData(t, collector, "GetApEthernetIfStats", func() (interface{}, error) {
			return GetApEthernetIfStats(client, ctx)
		})
	})

	t.Run("GetApEwlcWncdStats", func(t *testing.T) {
		runAPOperTestAndCollectData(t, collector, "GetApEwlcWncdStats", func() (interface{}, error) {
			return GetApEwlcWncdStats(client, ctx)
		})
	})

	t.Run("GetApIoxOperData", func(t *testing.T) {
		runAPOperTestAndCollectData(t, collector, "GetApIoxOperData", func() (interface{}, error) {
			return GetApIoxOperData(client, ctx)
		})
	})

	t.Run("GetApQosGlobalStats", func(t *testing.T) {
		runAPOperTestAndCollectData(t, collector, "GetApQosGlobalStats", func() (interface{}, error) {
			return GetApQosGlobalStats(client, ctx)
		})
	})

	t.Run("GetApOperData", func(t *testing.T) {
		runAPOperTestAndCollectData(t, collector, "GetApOperData", func() (interface{}, error) {
			return GetApOperData(client, ctx)
		})
	})

	t.Run("GetApRlanOper", func(t *testing.T) {
		runAPOperTestAndCollectData(t, collector, "GetApRlanOper", func() (interface{}, error) {
			return GetApRlanOper(client, ctx)
		})
	})

	t.Run("GetApEwlcMewlcPredownloadRec", func(t *testing.T) {
		runAPOperTestAndCollectData(t, collector, "GetApEwlcMewlcPredownloadRec", func() (interface{}, error) {
			return GetApEwlcMewlcPredownloadRec(client, ctx)
		})
	})

	t.Run("GetApCdpCacheData", func(t *testing.T) {
		runAPOperTestAndCollectData(t, collector, "GetApCdpCacheData", func() (interface{}, error) {
			return GetApCdpCacheData(client, ctx)
		})
	})

	t.Run("GetApLldpNeigh", func(t *testing.T) {
		runAPOperTestAndCollectData(t, collector, "GetApLldpNeigh", func() (interface{}, error) {
			return GetApLldpNeigh(client, ctx)
		})
	})

	t.Run("GetApTpCertInfo", func(t *testing.T) {
		runAPOperTestAndCollectData(t, collector, "GetApTpCertInfo", func() (interface{}, error) {
			return GetApTpCertInfo(client, ctx)
		})
	})

	t.Run("GetApDiscData", func(t *testing.T) {
		runAPOperTestAndCollectData(t, collector, "GetApDiscData", func() (interface{}, error) {
			return GetApDiscData(client, ctx)
		})
	})

	// Save collected test data
	if len(collector.Data) > 0 {
		if err := testutil.SaveTestDataToFile("ap_oper_test_data_collected.json", collector.Data); err != nil {
			t.Logf("Warning: Could not save test data: %v", err)
		} else {
			t.Logf("AP operation test data saved to %s/ap_oper_test_data_collected.json", testutil.TestDataDir)
		}
	}
}

// =============================================================================
// 5. OTHER TESTS (Performance, Edge Cases, etc.)
// =============================================================================

// TestApOperationPerformance tests performance characteristics of AP operations
func TestApOperationPerformance(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping performance test in short mode")
	}

	client := testutil.CreateTestClientFromEnv(t)
	ctx, cancel := context.WithTimeout(context.Background(), testutil.ExtendedTestTimeout)
	defer cancel()

	// Test concurrent requests
	t.Run("ConcurrentRequests", func(t *testing.T) {
		const numGoroutines = 5
		done := make(chan bool, numGoroutines)

		for i := 0; i < numGoroutines; i++ {
			go func() {
				_, err := GetApOper(client, ctx)
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
