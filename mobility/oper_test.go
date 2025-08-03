// Package mobility provides mobility operational data test functionality for the Cisco Wireless Network Controller API.
package mobility

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil"
)

// =============================================================================
// 1. UNIT TESTS (Structure/Type Validation & JSON Serialization/Deserialization)
// =============================================================================

// getTestClient creates a test client using environment variables
func getTestClient(t *testing.T) *wnc.Client {
	return testutil.CreateTestClientFromEnv(t)
}

// =============================================================================
// 2. INTEGRATION TESTS (API Endpoint Testing with Live Data Validation)
// =============================================================================

func TestMobilityOperGetMobilityOper(t *testing.T) {
	client := getTestClient(t)
	ctx, cancel := testutil.CreateDefaultTestContext()
	defer cancel()

	result, err := GetMobilityOper(client, ctx)
	if err != nil {
		t.Fatalf("GetMobilityOper failed: %v", err)
	}

	if result == nil {
		t.Fatal("GetMobilityOper returned nil result")
	}

	// Save result to JSON file
	filename := fmt.Sprintf("mobility_oper_data_%d.json", time.Now().Unix())
	if err := testutil.SaveTestDataToFile(filename, result); err != nil {
		t.Logf("Warning: Failed to save data to %s: %v", filename, err)
	} else {
		t.Logf("Data saved to %s", filename)
	}

	t.Logf("GetMobilityOper successful, collected mobility operational data")
}

func TestMobilityOperGetMobilityMmIfGlobalStats(t *testing.T) {
	client := getTestClient(t)
	ctx, cancel := testutil.CreateDefaultTestContext()
	defer cancel()

	result, err := GetMobilityMmIfGlobalStats(client, ctx)
	if err != nil {
		t.Fatalf("GetMobilityMmIfGlobalStats failed: %v", err)
	}

	if result == nil {
		t.Fatal("GetMobilityMmIfGlobalStats returned nil result")
	}

	// Save result to JSON file
	filename := fmt.Sprintf("mobility_mm_if_global_stats_data_%d.json", time.Now().Unix())
	if err := testutil.SaveTestDataToFile(filename, result); err != nil {
		t.Logf("Warning: Failed to save data to %s: %v", filename, err)
	} else {
		t.Logf("Data saved to %s", filename)
	}

	t.Logf("GetMobilityMmIfGlobalStats successful")
}

func TestMobilityOperGetMobilityMmIfGlobalMsgStats(t *testing.T) {
	client := getTestClient(t)
	ctx, cancel := testutil.CreateDefaultTestContext()
	defer cancel()

	result, err := GetMobilityMmIfGlobalMsgStats(client, ctx)
	if err != nil {
		t.Fatalf("GetMobilityMmIfGlobalMsgStats failed: %v", err)
	}

	if result == nil {
		t.Fatal("GetMobilityMmIfGlobalMsgStats returned nil result")
	}

	// Save result to JSON file
	filename := fmt.Sprintf("mobility_mm_if_global_msg_stats_data_%d.json", time.Now().Unix())
	if err := testutil.SaveTestDataToFile(filename, result); err != nil {
		t.Logf("Warning: Failed to save data to %s: %v", filename, err)
	} else {
		t.Logf("Data saved to %s", filename)
	}

	t.Logf("GetMobilityMmIfGlobalMsgStats successful")
}

func TestMobilityOperGetMobilityGlobalStats(t *testing.T) {
	client := getTestClient(t)
	ctx, cancel := testutil.CreateDefaultTestContext()
	defer cancel()

	result, err := GetMobilityGlobalStats(client, ctx)
	if err != nil {
		t.Fatalf("GetMobilityGlobalStats failed: %v", err)
	}

	if result == nil {
		t.Fatal("GetMobilityGlobalStats returned nil result")
	}

	// Save result to JSON file
	filename := fmt.Sprintf("mobility_global_stats_data_%d.json", time.Now().Unix())
	if err := testutil.SaveTestDataToFile(filename, result); err != nil {
		t.Logf("Warning: Failed to save data to %s: %v", filename, err)
	} else {
		t.Logf("Data saved to %s", filename)
	}

	t.Logf("GetMobilityGlobalStats successful")
}

func TestMobilityOperGetMobilityMmGlobalData(t *testing.T) {
	client := getTestClient(t)
	ctx, cancel := testutil.CreateDefaultTestContext()
	defer cancel()

	result, err := GetMobilityMmGlobalData(client, ctx)
	if err != nil {
		t.Fatalf("GetMobilityMmGlobalData failed: %v", err)
	}

	if result == nil {
		t.Fatal("GetMobilityMmGlobalData returned nil result")
	}

	// Save result to JSON file
	filename := fmt.Sprintf("mobility_mm_global_data_%d.json", time.Now().Unix())
	if err := testutil.SaveTestDataToFile(filename, result); err != nil {
		t.Logf("Warning: Failed to save data to %s: %v", filename, err)
	} else {
		t.Logf("Data saved to %s", filename)
	}

	t.Logf("GetMobilityMmGlobalData successful")
}

func TestMobilityOperGetMobilityGlobalMsgStats(t *testing.T) {
	client := getTestClient(t)
	ctx, cancel := testutil.CreateDefaultTestContext()
	defer cancel()

	result, err := GetMobilityGlobalMsgStats(client, ctx)
	if err != nil {
		t.Fatalf("GetMobilityGlobalMsgStats failed: %v", err)
	}

	if result == nil {
		t.Fatal("GetMobilityGlobalMsgStats returned nil result")
	}

	// Save result to JSON file
	filename := fmt.Sprintf("mobility_global_msg_stats_data_%d.json", time.Now().Unix())
	if err := testutil.SaveTestDataToFile(filename, result); err != nil {
		t.Logf("Warning: Failed to save data to %s: %v", filename, err)
	} else {
		t.Logf("Data saved to %s", filename)
	}

	t.Logf("GetMobilityGlobalMsgStats successful")
}

func TestMobilityOperGetMobilityClientData(t *testing.T) {
	client := getTestClient(t)
	ctx, cancel := testutil.CreateDefaultTestContext()
	defer cancel()

	result, err := GetMobilityClientData(client, ctx)
	if err != nil {
		t.Fatalf("GetMobilityClientData failed: %v", err)
	}

	if result == nil {
		t.Fatal("GetMobilityClientData returned nil result")
	}

	// Save result to JSON file
	filename := fmt.Sprintf("mobility_client_data_%d.json", time.Now().Unix())
	if err := testutil.SaveTestDataToFile(filename, result); err != nil {
		t.Logf("Warning: Failed to save data to %s: %v", filename, err)
	} else {
		t.Logf("Data saved to %s", filename)
	}

	t.Logf("GetMobilityClientData successful")
}

func TestMobilityOperGetMobilityApCache(t *testing.T) {
	client := getTestClient(t)
	ctx, cancel := testutil.CreateDefaultTestContext()
	defer cancel()

	result, err := GetMobilityApCache(client, ctx)
	if err != nil {
		t.Fatalf("GetMobilityApCache failed: %v", err)
	}

	if result == nil {
		t.Fatal("GetMobilityApCache returned nil result")
	}

	// Save result to JSON file
	filename := fmt.Sprintf("mobility_ap_cache_data_%d.json", time.Now().Unix())
	if err := testutil.SaveTestDataToFile(filename, result); err != nil {
		t.Logf("Warning: Failed to save data to %s: %v", filename, err)
	} else {
		t.Logf("Data saved to %s", filename)
	}

	t.Logf("GetMobilityApCache successful")
}

func TestMobilityOperGetMobilityApPeerList(t *testing.T) {
	client := getTestClient(t)
	ctx, cancel := testutil.CreateDefaultTestContext()
	defer cancel()

	result, err := GetMobilityApPeerList(client, ctx)
	if err != nil {
		t.Fatalf("GetMobilityApPeerList failed: %v", err)
	}

	if result == nil {
		t.Fatal("GetMobilityApPeerList returned nil result")
	}

	// Save result to JSON file
	filename := fmt.Sprintf("mobility_ap_peer_list_data_%d.json", time.Now().Unix())
	if err := testutil.SaveTestDataToFile(filename, result); err != nil {
		t.Logf("Warning: Failed to save data to %s: %v", filename, err)
	} else {
		t.Logf("Data saved to %s", filename)
	}

	t.Logf("GetMobilityApPeerList successful")
}

func TestMobilityOperGetMobilityClientStats(t *testing.T) {
	client := getTestClient(t)
	ctx, cancel := testutil.CreateDefaultTestContext()
	defer cancel()

	result, err := GetMobilityClientStats(client, ctx)
	if err != nil {
		t.Fatalf("GetMobilityClientStats failed: %v", err)
	}

	if result == nil {
		t.Fatal("GetMobilityClientStats returned nil result")
	}

	// Save result to JSON file
	filename := fmt.Sprintf("mobility_client_stats_data_%d.json", time.Now().Unix())
	if err := testutil.SaveTestDataToFile(filename, result); err != nil {
		t.Logf("Warning: Failed to save data to %s: %v", filename, err)
	} else {
		t.Logf("Data saved to %s", filename)
	}

	t.Logf("GetMobilityClientStats successful")
}

func TestMobilityOperGetMobilityWlanClientLimit(t *testing.T) {
	client := getTestClient(t)
	ctx, cancel := testutil.CreateDefaultTestContext()
	defer cancel()

	result, err := GetMobilityWlanClientLimit(client, ctx)
	if err != nil {
		t.Fatalf("GetMobilityWlanClientLimit failed: %v", err)
	}

	if result == nil {
		t.Fatal("GetMobilityWlanClientLimit returned nil result")
	}

	// Save result to JSON file
	filename := fmt.Sprintf("mobility_wlan_client_limit_data_%d.json", time.Now().Unix())
	if err := testutil.SaveTestDataToFile(filename, result); err != nil {
		t.Logf("Warning: Failed to save data to %s: %v", filename, err)
	} else {
		t.Logf("Data saved to %s", filename)
	}

	t.Logf("GetMobilityWlanClientLimit successful")
}

func TestMobilityOperGetMobilityGlobalDTLSStats(t *testing.T) {
	client := getTestClient(t)
	ctx, cancel := testutil.CreateDefaultTestContext()
	defer cancel()

	result, err := GetMobilityGlobalDTLSStats(client, ctx)
	if err != nil {
		t.Fatalf("GetMobilityGlobalDTLSStats failed: %v", err)
	}

	if result == nil {
		t.Fatal("GetMobilityGlobalDTLSStats returned nil result")
	}

	// Save result to JSON file
	filename := fmt.Sprintf("mobility_global_dtls_stats_data_%d.json", time.Now().Unix())
	if err := testutil.SaveTestDataToFile(filename, result); err != nil {
		t.Logf("Warning: Failed to save data to %s: %v", filename, err)
	} else {
		t.Logf("Data saved to %s", filename)
	}

	t.Logf("GetMobilityGlobalDTLSStats successful")
}

func TestMobilityOperCollectAllData(t *testing.T) {
	client := getTestClient(t)
	ctx, cancel := testutil.CreateDefaultTestContext()
	defer cancel()

	allData := make(map[string]interface{})

	// Collect data from all mobility operational endpoints
	tests := []struct {
		name string
		fn   func() (interface{}, error)
	}{
		{"GetMobilityOper", func() (interface{}, error) { return GetMobilityOper(client, ctx) }},
		{"GetMobilityMmIfGlobalStats", func() (interface{}, error) { return GetMobilityMmIfGlobalStats(client, ctx) }},
		{"GetMobilityMmIfGlobalMsgStats", func() (interface{}, error) { return GetMobilityMmIfGlobalMsgStats(client, ctx) }},
		{"GetMobilityGlobalStats", func() (interface{}, error) { return GetMobilityGlobalStats(client, ctx) }},
		{"GetMobilityMmGlobalData", func() (interface{}, error) { return GetMobilityMmGlobalData(client, ctx) }},
		{"GetMobilityGlobalMsgStats", func() (interface{}, error) { return GetMobilityGlobalMsgStats(client, ctx) }},
		{"GetMobilityClientData", func() (interface{}, error) { return GetMobilityClientData(client, ctx) }},
		{"GetMobilityApCache", func() (interface{}, error) { return GetMobilityApCache(client, ctx) }},
		{"GetMobilityApPeerList", func() (interface{}, error) { return GetMobilityApPeerList(client, ctx) }},
		{"GetMobilityClientStats", func() (interface{}, error) { return GetMobilityClientStats(client, ctx) }},
		{"GetMobilityWlanClientLimit", func() (interface{}, error) { return GetMobilityWlanClientLimit(client, ctx) }},
		{"GetMobilityGlobalDTLSStats", func() (interface{}, error) { return GetMobilityGlobalDTLSStats(client, ctx) }},
	}

	for _, test := range tests {
		result, err := test.fn()
		if err != nil {
			t.Logf("Warning: %s failed: %v", test.name, err)
			allData[test.name] = map[string]string{"error": err.Error()}
		} else {
			allData[test.name] = result
			t.Logf("%s successful", test.name)
		}
	}

	// Save all collected data to a comprehensive JSON file
	filename := fmt.Sprintf("mobility_oper_comprehensive_data_%d.json", time.Now().Unix())
	if err := testutil.SaveTestDataToFile(filename, allData); err != nil {
		t.Logf("Warning: Failed to save comprehensive data to %s: %v", filename, err)
	} else {
		t.Logf("Comprehensive mobility operational data saved to %s", filename)
	}
}

// TestMobilityOperDataStructures tests the basic structure of mobility operational data types
func TestMobilityOperDataStructures(t *testing.T) {
	tests := []struct {
		name     string
		jsonData string
		dataType interface{}
	}{
		{
			name: "MobilityOperMmIfGlobalStatsResponse",
			jsonData: `{
				"Cisco-IOS-XE-wireless-mobility-oper:mm-if-global-stats": {
					"total-handoffs": 150,
					"successful-handoffs": 145,
					"failed-handoffs": 5,
					"inter-controller-handoffs": 75,
					"intra-controller-handoffs": 70,
					"anchor-requests": 50,
					"foreign-requests": 25
				}
			}`,
			dataType: &MmIfGlobalStatsResponse{},
		},
		{
			name: "MobilityOperMobilityGlobalStatsResponse",
			jsonData: `{
				"Cisco-IOS-XE-wireless-mobility-oper:mobility-global-stats": {
					"total-clients": 500,
					"local-clients": 450,
					"anchor-clients": 30,
					"foreign-clients": 20,
					"mobility-tunnels": 15,
					"active-sessions": 485
				}
			}`,
			dataType: &MobilityGlobalStatsResponse{},
		},
		{
			name: "MobilityOperMobilityClientDataResponse",
			jsonData: `{
				"Cisco-IOS-XE-wireless-mobility-oper:mobility-client-data": [
					{
						"client-mac": "aa:bb:cc:dd:ee:ff",
						"mobility-state": "local",
						"anchor-controller": "192.168.1.10",
						"foreign-controller": "",
						"tunnel-id": 0,
						"session-timeout": 3600
					}
				]
			}`,
			dataType: &MobilityClientDataResponse{},
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
