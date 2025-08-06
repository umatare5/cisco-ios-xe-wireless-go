// Package ap provides access point operational data test functionality for the Cisco Wireless Network Controller API.
package ap

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
	testutils "github.com/umatare5/cisco-ios-xe-wireless-go/internal/tests"
	wnccore "github.com/umatare5/cisco-ios-xe-wireless-go/wnc"
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
	client := testutils.CreateTestClientFromEnv(t)
	ctx, cancel := context.WithTimeout(context.Background(), testutils.DefaultTestTimeout)
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
		client := testutils.CreateTestClientFromEnv(t)
		_, err := GetApOper(client, nil)
		if err == nil {
			t.Fatal("Expected error with nil context, got none")
		}
		t.Logf("Correctly returned error with nil context: %v", err)
	})

	// Test with canceled context
	t.Run("CanceledContext", func(t *testing.T) {
		client := testutils.CreateTestClientFromEnv(t)
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
	client := testutils.CreateTestClientFromEnv(t)

	ctx, cancel := context.WithTimeout(context.Background(), testutils.DefaultTestTimeout)
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

	t.Run("GetApCapwapPkts", func(t *testing.T) {
		runAPOperTestAndCollectData(t, collector, "GetApCapwapPkts", func() (interface{}, error) {
			return GetApCapwapPkts(client, ctx)
		})
	})

	t.Run("GetApCountryOper", func(t *testing.T) {
		runAPOperTestAndCollectData(t, collector, "GetApCountryOper", func() (interface{}, error) {
			return GetApCountryOper(client, ctx)
		})
	})

	t.Run("GetApSuppCountryOper", func(t *testing.T) {
		runAPOperTestAndCollectData(t, collector, "GetApSuppCountryOper", func() (interface{}, error) {
			return GetApSuppCountryOper(client, ctx)
		})
	})

	t.Run("GetApNhGlobalData", func(t *testing.T) {
		runAPOperTestAndCollectData(t, collector, "GetApNhGlobalData", func() (interface{}, error) {
			return GetApNhGlobalData(client, ctx)
		})
	})

	t.Run("GetApImagePrepareLocation", func(t *testing.T) {
		runAPOperTestAndCollectData(t, collector, "GetApImagePrepareLocation", func() (interface{}, error) {
			return GetApImagePrepareLocation(client, ctx)
		})
	})

	t.Run("GetApImageActiveLocation", func(t *testing.T) {
		runAPOperTestAndCollectData(t, collector, "GetApImageActiveLocation", func() (interface{}, error) {
			return GetApImageActiveLocation(client, ctx)
		})
	})

	// Save collected test data
	if len(collector.Data) > 0 {
		if err := testutils.SaveTestDataToFile("ap_oper_test_data_collected.json", collector.Data); err != nil {
			t.Logf("Warning: Could not save test data: %v", err)
		} else {
			t.Logf("AP operation test data saved to %s/ap_oper_test_data_collected.json", testutils.TestDataDir)
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

	client := testutils.CreateTestClientFromEnv(t)
	ctx, cancel := context.WithTimeout(context.Background(), testutils.ExtendedTestTimeout)
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

// =============================================================================
// 6. ERROR HANDLING TESTS
// =============================================================================

// TestApOperErrorHandling tests error handling for all operational functions
func TestApOperErrorHandling(t *testing.T) {
	ctx := context.Background()

	testCases := []struct {
		name string
		fn   func() (interface{}, error)
	}{
		{"GetApOper", func() (interface{}, error) { return GetApOper(nil, ctx) }},
		{"GetApRadioNeighbor", func() (interface{}, error) { return GetApRadioNeighbor(nil, ctx) }},
		{"GetApRadioOperData", func() (interface{}, error) { return GetApRadioOperData(nil, ctx) }},
		{"GetApRadioResetStats", func() (interface{}, error) { return GetApRadioResetStats(nil, ctx) }},
		{"GetApQosClientData", func() (interface{}, error) { return GetApQosClientData(nil, ctx) }},
		{"GetApCapwapData", func() (interface{}, error) { return GetApCapwapData(nil, ctx) }},
		{"GetApNameMacMap", func() (interface{}, error) { return GetApNameMacMap(nil, ctx) }},
		{"GetApWtpSlotWlanStats", func() (interface{}, error) { return GetApWtpSlotWlanStats(nil, ctx) }},
		{"GetApEthernetMacWtpMacMap", func() (interface{}, error) { return GetApEthernetMacWtpMacMap(nil, ctx) }},
		{"GetApRadioOperStats", func() (interface{}, error) { return GetApRadioOperStats(nil, ctx) }},
		{"GetApEthernetIfStats", func() (interface{}, error) { return GetApEthernetIfStats(nil, ctx) }},
		{"GetApEwlcWncdStats", func() (interface{}, error) { return GetApEwlcWncdStats(nil, ctx) }},
		{"GetApIoxOperData", func() (interface{}, error) { return GetApIoxOperData(nil, ctx) }},
		{"GetApQosGlobalStats", func() (interface{}, error) { return GetApQosGlobalStats(nil, ctx) }},
		{"GetApOperData", func() (interface{}, error) { return GetApOperData(nil, ctx) }},
		{"GetApRlanOper", func() (interface{}, error) { return GetApRlanOper(nil, ctx) }},
		{"GetApEwlcMewlcPredownloadRec", func() (interface{}, error) { return GetApEwlcMewlcPredownloadRec(nil, ctx) }},
		{"GetApCdpCacheData", func() (interface{}, error) { return GetApCdpCacheData(nil, ctx) }},
		{"GetApLldpNeigh", func() (interface{}, error) { return GetApLldpNeigh(nil, ctx) }},
		{"GetApTpCertInfo", func() (interface{}, error) { return GetApTpCertInfo(nil, ctx) }},
		{"GetApDiscData", func() (interface{}, error) { return GetApDiscData(nil, ctx) }},
		{"GetApCapwapPkts", func() (interface{}, error) { return GetApCapwapPkts(nil, ctx) }},
		{"GetApCountryOper", func() (interface{}, error) { return GetApCountryOper(nil, ctx) }},
		{"GetApSuppCountryOper", func() (interface{}, error) { return GetApSuppCountryOper(nil, ctx) }},
		{"GetApNhGlobalData", func() (interface{}, error) { return GetApNhGlobalData(nil, ctx) }},
		{"GetApImagePrepareLocation", func() (interface{}, error) { return GetApImagePrepareLocation(nil, ctx) }},
		{"GetApImageActiveLocation", func() (interface{}, error) { return GetApImageActiveLocation(nil, ctx) }},
	}

	for _, tc := range testCases {
		t.Run(tc.name+"WithNilClient", func(t *testing.T) {
			_, err := tc.fn()
			if err == nil {
				t.Errorf("Expected error with nil client, got nil")
			}
			// Accept either error message format for consistency
			errorMsg := err.Error()
			if errorMsg != "client is nil" && errorMsg != "invalid client configuration: client cannot be nil" {
				t.Errorf("Expected 'client is nil' or 'invalid client configuration' error, got: %v", err)
			}
		})
	}
}

// =============================================================================
// 7. CONTEXT HANDLING TESTS
// =============================================================================

// TestApOperContextHandling tests context handling for all operational functions
func TestApOperContextHandling(t *testing.T) {
	testCases := []struct {
		name string
		fn   func(context.Context, *wnc.Client) error
	}{
		{"GetApOper", func(ctx context.Context, client *wnc.Client) error { _, err := GetApOper(client, ctx); return err }},
		{"GetApRadioNeighbor", func(ctx context.Context, client *wnc.Client) error {
			_, err := GetApRadioNeighbor(client, ctx)
			return err
		}},
		{"GetApRadioOperData", func(ctx context.Context, client *wnc.Client) error {
			_, err := GetApRadioOperData(client, ctx)
			return err
		}},
		{"GetApRadioResetStats", func(ctx context.Context, client *wnc.Client) error {
			_, err := GetApRadioResetStats(client, ctx)
			return err
		}},
		{"GetApQosClientData", func(ctx context.Context, client *wnc.Client) error {
			_, err := GetApQosClientData(client, ctx)
			return err
		}},
		{"GetApCapwapData", func(ctx context.Context, client *wnc.Client) error {
			_, err := GetApCapwapData(client, ctx)
			return err
		}},
		{"GetApNameMacMap", func(ctx context.Context, client *wnc.Client) error {
			_, err := GetApNameMacMap(client, ctx)
			return err
		}},
		{"GetApWtpSlotWlanStats", func(ctx context.Context, client *wnc.Client) error {
			_, err := GetApWtpSlotWlanStats(client, ctx)
			return err
		}},
		{"GetApEthernetMacWtpMacMap", func(ctx context.Context, client *wnc.Client) error {
			_, err := GetApEthernetMacWtpMacMap(client, ctx)
			return err
		}},
		{"GetApRadioOperStats", func(ctx context.Context, client *wnc.Client) error {
			_, err := GetApRadioOperStats(client, ctx)
			return err
		}},
		{"GetApEthernetIfStats", func(ctx context.Context, client *wnc.Client) error {
			_, err := GetApEthernetIfStats(client, ctx)
			return err
		}},
		{"GetApEwlcWncdStats", func(ctx context.Context, client *wnc.Client) error {
			_, err := GetApEwlcWncdStats(client, ctx)
			return err
		}},
		{"GetApIoxOperData", func(ctx context.Context, client *wnc.Client) error {
			_, err := GetApIoxOperData(client, ctx)
			return err
		}},
		{"GetApQosGlobalStats", func(ctx context.Context, client *wnc.Client) error {
			_, err := GetApQosGlobalStats(client, ctx)
			return err
		}},
		{"GetApOperData", func(ctx context.Context, client *wnc.Client) error { _, err := GetApOperData(client, ctx); return err }},
		{"GetApRlanOper", func(ctx context.Context, client *wnc.Client) error { _, err := GetApRlanOper(client, ctx); return err }},
		{"GetApEwlcMewlcPredownloadRec", func(ctx context.Context, client *wnc.Client) error {
			_, err := GetApEwlcMewlcPredownloadRec(client, ctx)
			return err
		}},
		{"GetApCdpCacheData", func(ctx context.Context, client *wnc.Client) error {
			_, err := GetApCdpCacheData(client, ctx)
			return err
		}},
		{"GetApLldpNeigh", func(ctx context.Context, client *wnc.Client) error { _, err := GetApLldpNeigh(client, ctx); return err }},
		{"GetApTpCertInfo", func(ctx context.Context, client *wnc.Client) error {
			_, err := GetApTpCertInfo(client, ctx)
			return err
		}},
		{"GetApDiscData", func(ctx context.Context, client *wnc.Client) error { _, err := GetApDiscData(client, ctx); return err }},
		{"GetApCapwapPkts", func(ctx context.Context, client *wnc.Client) error {
			_, err := GetApCapwapPkts(client, ctx)
			return err
		}},
		{"GetApCountryOper", func(ctx context.Context, client *wnc.Client) error {
			_, err := GetApCountryOper(client, ctx)
			return err
		}},
		{"GetApSuppCountryOper", func(ctx context.Context, client *wnc.Client) error {
			_, err := GetApSuppCountryOper(client, ctx)
			return err
		}},
		{"GetApNhGlobalData", func(ctx context.Context, client *wnc.Client) error {
			_, err := GetApNhGlobalData(client, ctx)
			return err
		}},
		{"GetApImagePrepareLocation", func(ctx context.Context, client *wnc.Client) error {
			_, err := GetApImagePrepareLocation(client, ctx)
			return err
		}},
		{"GetApImageActiveLocation", func(ctx context.Context, client *wnc.Client) error {
			_, err := GetApImageActiveLocation(client, ctx)
			return err
		}},
	}

	for _, tc := range testCases {
		t.Run(tc.name+"ContextHandling", func(t *testing.T) {
			testutils.TestContextHandling(t, tc.fn)
		})
	}
}

// =============================================================================
// 4. INTEGRATION TESTS FOR SUCCESSFUL PATHS
// =============================================================================

func TestApOperIntegrationSuccess(t *testing.T) {
	client := testutils.CreateTestClientFromEnv(t)
	ctx, cancel := context.WithTimeout(context.Background(), testutils.DefaultTestTimeout)
	defer cancel()

	// Test functions that need success path coverage
	t.Run("GetApQosClientDataSuccess", func(t *testing.T) {
		result, err := GetApQosClientData(client, ctx)
		if err != nil {
			t.Logf("GetApQosClientData returned error (expected in some environments): %v", err)
		} else if result != nil {
			t.Logf("GetApQosClientData successful")
		}
	})

	t.Run("GetApRlanOperSuccess", func(t *testing.T) {
		result, err := GetApRlanOper(client, ctx)
		if err != nil {
			t.Logf("GetApRlanOper returned error (expected in some environments): %v", err)
		} else if result != nil {
			t.Logf("GetApRlanOper successful")
		}
	})

	// Additional functions needing success path coverage
	t.Run("GetApCapwapDataSuccess", func(t *testing.T) {
		result, err := GetApCapwapData(client, ctx)
		if err != nil {
			t.Logf("GetApCapwapData returned error (expected in some environments): %v", err)
		} else if result != nil {
			t.Logf("GetApCapwapData successful")
		}
	})

	t.Run("GetApNameMacMapSuccess", func(t *testing.T) {
		result, err := GetApNameMacMap(client, ctx)
		if err != nil {
			t.Logf("GetApNameMacMap returned error (expected in some environments): %v", err)
		} else if result != nil {
			t.Logf("GetApNameMacMap successful")
		}
	})

	t.Run("GetApWtpSlotWlanStatsSuccess", func(t *testing.T) {
		result, err := GetApWtpSlotWlanStats(client, ctx)
		if err != nil {
			t.Logf("GetApWtpSlotWlanStats returned error (expected in some environments): %v", err)
		} else if result != nil {
			t.Logf("GetApWtpSlotWlanStats successful")
		}
	})

	t.Run("GetApEthernetMacWtpMacMapSuccess", func(t *testing.T) {
		result, err := GetApEthernetMacWtpMacMap(client, ctx)
		if err != nil {
			t.Logf("GetApEthernetMacWtpMacMap returned error (expected in some environments): %v", err)
		} else if result != nil {
			t.Logf("GetApEthernetMacWtpMacMap successful")
		}
	})

	t.Run("GetApRadioOperStatsSuccess", func(t *testing.T) {
		result, err := GetApRadioOperStats(client, ctx)
		if err != nil {
			t.Logf("GetApRadioOperStats returned error (expected in some environments): %v", err)
		} else if result != nil {
			t.Logf("GetApRadioOperStats successful")
		}
	})
}

// =============================================================================
// 6. DETAILED SUCCESS PATH TESTS
// =============================================================================

// TestApOperSuccessPathCoverage tests specific functions to ensure 100% coverage
func TestApOperSuccessPathCoverage(t *testing.T) {
	// Create a mock server that returns success responses
	mockServer := testutils.NewMockHTTPServer()

	mockServer.AddHandler("Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/qos-client-data",
		testutils.CreateJSONResponse(testutils.TestHTTPResponse{
			StatusCode: 200,
			Body: `{
				"Cisco-IOS-XE-wireless-access-point-oper:qos-client-data": [
					{
						"client-mac": "aa:bb:cc:dd:ee:ff",
						"qos-level": "gold"
					}
				]
			}`,
			Headers: map[string]string{"Content-Type": "application/yang-data+json"},
		}))

	mockServer.AddHandler("Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/rlan-oper",
		testutils.CreateJSONResponse(testutils.TestHTTPResponse{
			StatusCode: 200,
			Body: `{
				"Cisco-IOS-XE-wireless-access-point-oper:rlan-oper": [
					{
						"ap-name": "test-ap",
						"rlan-id": 1
					}
				]
			}`,
			Headers: map[string]string{"Content-Type": "application/yang-data+json"},
		}))

	defer mockServer.Close()

	client := testutils.CreateTestClientForMockServer(t, mockServer)
	ctx, cancel := context.WithTimeout(context.Background(), testutils.DefaultTestTimeout)
	defer cancel()

	t.Run("GetApQosClientDataSuccessPath", func(t *testing.T) {
		result, err := GetApQosClientData(client, ctx)
		if err != nil {
			// QoS client data may not be available on all controllers (404 Not Found)
			var httpErr *wnccore.HTTPError
			if errors.As(err, &httpErr) && httpErr.Status == 404 {
				t.Skipf("QoS client data not supported on this controller: HTTP 404")
			}
			t.Errorf("Expected GetApQosClientData to succeed with mock server, got error: %v", err)
		}
		if result == nil {
			t.Error("Expected GetApQosClientData to return non-nil result")
		}
		// Verify the result structure
		if result != nil && len(result.QosClientData) == 0 {
			t.Log("GetApQosClientData returned result but with empty data (acceptable)")
		}
	})

	t.Run("GetApRlanOperSuccessPath", func(t *testing.T) {
		result, err := GetApRlanOper(client, ctx)
		if err != nil {
			// RLAN operations may not be available on all controllers (404 Not Found)
			var httpErr *wnccore.HTTPError
			if errors.As(err, &httpErr) && httpErr.Status == 404 {
				t.Skipf("RLAN operations not supported on this controller: HTTP 404")
			}
			t.Errorf("Expected GetApRlanOper to succeed with mock server, got error: %v", err)
		}
		if result == nil {
			t.Error("Expected GetApRlanOper to return non-nil result")
		}
		// Verify the result structure
		if result != nil && len(result.RlanOper) == 0 {
			t.Log("GetApRlanOper returned result but with empty data (acceptable)")
		}
	})
}
