// Package rrm provides Radio Resource Management global operational data test functionality for the Cisco Wireless Network Controller API.
package rrm

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/umatare5/cisco-xe-wireless-restconf-go/internal/testutil"
)

// =============================================================================
// 1. UNIT TESTS (Structure/Type Validation & JSON Serialization/Deserialization)
// =============================================================================

// RrmGlobalOperTestDataCollector holds test data for RRM global operation functions
type RrmGlobalOperTestDataCollector struct {
	Data map[string]interface{} `json:"rrm_global_oper_test_data"`
}

var rrmGlobalOperTestDataCollector = RrmGlobalOperTestDataCollector{
	Data: make(map[string]interface{}),
}

// =============================================================================
// 2. INTEGRATION TESTS (API Endpoint Testing with Live Data Validation)
// =============================================================================

func TestRrmGlobalOperGetRrmGlobalOper(t *testing.T) {
	client := testutil.CreateTestClientFromEnv(t)
	ctx, cancel := context.WithTimeout(context.Background(), testutil.DefaultTestTimeout)
	defer cancel()

	result, err := GetRrmGlobalOper(client, ctx)
	if err != nil {
		t.Fatalf("GetRrmGlobalOper failed: %v", err)
	}

	if result == nil {
		t.Fatal("GetRrmGlobalOper returned nil result")
	}

	// Save result to JSON file
	filename := fmt.Sprintf("rrm_global_oper_data_%d.json", time.Now().Unix())
	if err := testutil.SaveTestDataToFile(filename, result); err != nil {
		t.Logf("Warning: Failed to save data to %s: %v", filename, err)
	} else {
		t.Logf("Data saved to %s", filename)
	}

	t.Logf("GetRrmGlobalOper successful")
}

func TestRrmGlobalOperGetRrmGlobalOneShotCounters(t *testing.T) {
	client := testutil.CreateTestClientFromEnv(t)
	ctx, cancel := context.WithTimeout(context.Background(), testutil.DefaultTestTimeout)
	defer cancel()

	result, err := GetRrmGlobalOneShotCounters(client, ctx)
	if err != nil {
		t.Fatalf("GetRrmGlobalOneShotCounters failed: %v", err)
	}

	if result == nil {
		t.Fatal("GetRrmGlobalOneShotCounters returned nil result")
	}

	// Save result to JSON file
	filename := fmt.Sprintf("rrm_global_one_shot_counters_data_%d.json", time.Now().Unix())
	if err := testutil.SaveTestDataToFile(filename, result); err != nil {
		t.Logf("Warning: Failed to save data to %s: %v", filename, err)
	} else {
		t.Logf("Data saved to %s", filename)
	}

	t.Logf("GetRrmGlobalOneShotCounters successful")
}

func TestRrmGlobalOperGetRrmGlobalChannelParams(t *testing.T) {
	client := testutil.CreateTestClientFromEnv(t)
	ctx, cancel := context.WithTimeout(context.Background(), testutil.DefaultTestTimeout)
	defer cancel()

	result, err := GetRrmGlobalChannelParams(client, ctx)
	if err != nil {
		t.Fatalf("GetRrmGlobalChannelParams failed: %v", err)
	}

	if result == nil {
		t.Fatal("GetRrmGlobalChannelParams returned nil result")
	}

	// Save result to JSON file
	filename := fmt.Sprintf("rrm_global_channel_params_data_%d.json", time.Now().Unix())
	if err := testutil.SaveTestDataToFile(filename, result); err != nil {
		t.Logf("Warning: Failed to save data to %s: %v", filename, err)
	} else {
		t.Logf("Data saved to %s", filename)
	}

	t.Logf("GetRrmGlobalChannelParams successful")
}

func TestRrmGlobalOperGetRrmGlobalSpectrumAqWorstTable(t *testing.T) {
	client := testutil.CreateTestClientFromEnv(t)
	ctx, cancel := context.WithTimeout(context.Background(), testutil.DefaultTestTimeout)
	defer cancel()

	result, err := GetRrmGlobalSpectrumAqWorstTable(client, ctx)
	if err != nil {
		t.Fatalf("GetRrmGlobalSpectrumAqWorstTable failed: %v", err)
	}

	if result == nil {
		t.Fatal("GetRrmGlobalSpectrumAqWorstTable returned nil result")
	}

	// Save result to JSON file
	filename := fmt.Sprintf("rrm_global_spectrum_aq_worst_table_data_%d.json", time.Now().Unix())
	if err := testutil.SaveTestDataToFile(filename, result); err != nil {
		t.Logf("Warning: Failed to save data to %s: %v", filename, err)
	} else {
		t.Logf("Data saved to %s", filename)
	}

	t.Logf("GetRrmGlobalSpectrumAqWorstTable successful")
}

func TestRrmGlobalOperGetRrmGlobalRadioOperData24G(t *testing.T) {
	client := testutil.CreateTestClientFromEnv(t)
	ctx, cancel := context.WithTimeout(context.Background(), testutil.DefaultTestTimeout)
	defer cancel()

	result, err := GetRrmGlobalRadioOperData24G(client, ctx)
	if err != nil {
		t.Fatalf("GetRrmGlobalRadioOperData24G failed: %v", err)
	}

	if result == nil {
		t.Fatal("GetRrmGlobalRadioOperData24G returned nil result")
	}

	// Save result to JSON file
	filename := fmt.Sprintf("rrm_global_radio_oper_data_24g_%d.json", time.Now().Unix())
	if err := testutil.SaveTestDataToFile(filename, result); err != nil {
		t.Logf("Warning: Failed to save data to %s: %v", filename, err)
	} else {
		t.Logf("Data saved to %s", filename)
	}

	t.Logf("GetRrmGlobalRadioOperData24G successful")
}

func TestRrmGlobalOperGetRrmGlobalRadioOperData5G(t *testing.T) {
	client := testutil.CreateTestClientFromEnv(t)
	ctx, cancel := context.WithTimeout(context.Background(), testutil.DefaultTestTimeout)
	defer cancel()

	result, err := GetRrmGlobalRadioOperData5G(client, ctx)
	if err != nil {
		t.Fatalf("GetRrmGlobalRadioOperData5G failed: %v", err)
	}

	if result == nil {
		t.Fatal("GetRrmGlobalRadioOperData5G returned nil result")
	}

	// Save result to JSON file
	filename := fmt.Sprintf("rrm_global_radio_oper_data_5g_%d.json", time.Now().Unix())
	if err := testutil.SaveTestDataToFile(filename, result); err != nil {
		t.Logf("Warning: Failed to save data to %s: %v", filename, err)
	} else {
		t.Logf("Data saved to %s", filename)
	}

	t.Logf("GetRrmGlobalRadioOperData5G successful")
}

func TestRrmGlobalOperGetRrmGlobalRadioOperData6G(t *testing.T) {
	client := testutil.CreateTestClientFromEnv(t)
	ctx, cancel := context.WithTimeout(context.Background(), testutil.DefaultTestTimeout)
	defer cancel()

	result, err := GetRrmGlobalRadioOperData6G(client, ctx)
	if err != nil {
		t.Fatalf("GetRrmGlobalRadioOperData6G failed: %v", err)
	}

	if result == nil {
		t.Fatal("GetRrmGlobalRadioOperData6G returned nil result")
	}

	// Save result to JSON file
	filename := fmt.Sprintf("rrm_global_radio_oper_data_6g_%d.json", time.Now().Unix())
	if err := testutil.SaveTestDataToFile(filename, result); err != nil {
		t.Logf("Warning: Failed to save data to %s: %v", filename, err)
	} else {
		t.Logf("Data saved to %s", filename)
	}

	t.Logf("GetRrmGlobalRadioOperData6G successful")
}

func TestRrmGlobalOperGetRrmGlobalSpectrumBandConfigData(t *testing.T) {
	client := testutil.CreateTestClientFromEnv(t)
	ctx, cancel := context.WithTimeout(context.Background(), testutil.DefaultTestTimeout)
	defer cancel()

	result, err := GetRrmGlobalSpectrumBandConfigData(client, ctx)
	if err != nil {
		t.Fatalf("GetRrmGlobalSpectrumBandConfigData failed: %v", err)
	}

	if result == nil {
		t.Fatal("GetRrmGlobalSpectrumBandConfigData returned nil result")
	}

	// Save result to JSON file
	filename := fmt.Sprintf("rrm_global_spectrum_band_config_data_%d.json", time.Now().Unix())
	if err := testutil.SaveTestDataToFile(filename, result); err != nil {
		t.Logf("Warning: Failed to save data to %s: %v", filename, err)
	} else {
		t.Logf("Data saved to %s", filename)
	}

	t.Logf("GetRrmGlobalSpectrumBandConfigData successful")
}

func TestRrmGlobalOperGetRrmGlobalRadioOperDataDualband(t *testing.T) {
	client := testutil.CreateTestClientFromEnv(t)
	ctx, cancel := context.WithTimeout(context.Background(), testutil.DefaultTestTimeout)
	defer cancel()

	result, err := GetRrmGlobalRadioOperDataDualband(client, ctx)
	if err != nil {
		t.Fatalf("GetRrmGlobalRadioOperDataDualband failed: %v", err)
	}

	if result == nil {
		t.Fatal("GetRrmGlobalRadioOperDataDualband returned nil result")
	}

	// Save result to JSON file
	filename := fmt.Sprintf("rrm_global_radio_oper_data_dualband_%d.json", time.Now().Unix())
	if err := testutil.SaveTestDataToFile(filename, result); err != nil {
		t.Logf("Warning: Failed to save data to %s: %v", filename, err)
	} else {
		t.Logf("Data saved to %s", filename)
	}

	t.Logf("GetRrmGlobalRadioOperDataDualband successful")
}

func TestRrmGlobalOperGetRrmGlobalClientData(t *testing.T) {
	client := testutil.CreateTestClientFromEnv(t)
	ctx, cancel := context.WithTimeout(context.Background(), testutil.DefaultTestTimeout)
	defer cancel()

	result, err := GetRrmGlobalClientData(client, ctx)
	if err != nil {
		t.Fatalf("GetRrmGlobalClientData failed: %v", err)
	}

	if result == nil {
		t.Fatal("GetRrmGlobalClientData returned nil result")
	}

	// Save result to JSON file
	filename := fmt.Sprintf("rrm_global_client_data_%d.json", time.Now().Unix())
	if err := testutil.SaveTestDataToFile(filename, result); err != nil {
		t.Logf("Warning: Failed to save data to %s: %v", filename, err)
	} else {
		t.Logf("Data saved to %s", filename)
	}

	t.Logf("GetRrmGlobalClientData successful")
}

func TestRrmGlobalOperGetRrmGlobalFraStats(t *testing.T) {
	client := testutil.CreateTestClientFromEnv(t)
	ctx, cancel := context.WithTimeout(context.Background(), testutil.DefaultTestTimeout)
	defer cancel()

	result, err := GetRrmGlobalFraStats(client, ctx)
	if err != nil {
		t.Fatalf("GetRrmGlobalFraStats failed: %v", err)
	}

	if result == nil {
		t.Fatal("GetRrmGlobalFraStats returned nil result")
	}

	// Save result to JSON file
	filename := fmt.Sprintf("rrm_global_fra_stats_%d.json", time.Now().Unix())
	if err := testutil.SaveTestDataToFile(filename, result); err != nil {
		t.Logf("Warning: Failed to save data to %s: %v", filename, err)
	} else {
		t.Logf("Data saved to %s", filename)
	}

	t.Logf("GetRrmGlobalFraStats successful")
}

func TestRrmGlobalOperGetRrmGlobalCoverage(t *testing.T) {
	client := testutil.CreateTestClientFromEnv(t)
	ctx, cancel := context.WithTimeout(context.Background(), testutil.DefaultTestTimeout)
	defer cancel()

	result, err := GetRrmGlobalCoverage(client, ctx)
	if err != nil {
		t.Fatalf("GetRrmGlobalCoverage failed: %v", err)
	}

	if result == nil {
		t.Fatal("GetRrmGlobalCoverage returned nil result")
	}

	// Save result to JSON file
	filename := fmt.Sprintf("rrm_global_coverage_%d.json", time.Now().Unix())
	if err := testutil.SaveTestDataToFile(filename, result); err != nil {
		t.Logf("Warning: Failed to save data to %s: %v", filename, err)
	} else {
		t.Logf("Data saved to %s", filename)
	}

	t.Logf("GetRrmGlobalCoverage successful")
}

func TestRrmGlobalOperCollectAllData(t *testing.T) {
	client := testutil.CreateTestClientFromEnv(t)
	ctx, cancel := context.WithTimeout(context.Background(), testutil.DefaultTestTimeout)
	defer cancel()

	allData := make(map[string]interface{})

	// Collect data from all RRM global operational endpoints
	tests := []struct {
		name string
		fn   func() (interface{}, error)
	}{
		{"GetRrmGlobalOper", func() (interface{}, error) { return GetRrmGlobalOper(client, ctx) }},
		{"GetRrmGlobalOneShotCounters", func() (interface{}, error) { return GetRrmGlobalOneShotCounters(client, ctx) }},
		{"GetRrmGlobalChannelParams", func() (interface{}, error) { return GetRrmGlobalChannelParams(client, ctx) }},
		{"GetRrmGlobalSpectrumAqWorstTable", func() (interface{}, error) { return GetRrmGlobalSpectrumAqWorstTable(client, ctx) }},
		{"GetRrmGlobalRadioOperData24G", func() (interface{}, error) { return GetRrmGlobalRadioOperData24G(client, ctx) }},
		{"GetRrmGlobalRadioOperData5G", func() (interface{}, error) { return GetRrmGlobalRadioOperData5G(client, ctx) }},
		{"GetRrmGlobalRadioOperData6G", func() (interface{}, error) { return GetRrmGlobalRadioOperData6G(client, ctx) }},
		{"GetRrmGlobalSpectrumBandConfigData", func() (interface{}, error) { return GetRrmGlobalSpectrumBandConfigData(client, ctx) }},
		{"GetRrmGlobalRadioOperDataDualband", func() (interface{}, error) { return GetRrmGlobalRadioOperDataDualband(client, ctx) }},
		{"GetRrmGlobalClientData", func() (interface{}, error) { return GetRrmGlobalClientData(client, ctx) }},
		{"GetRrmGlobalFraStats", func() (interface{}, error) { return GetRrmGlobalFraStats(client, ctx) }},
		{"GetRrmGlobalCoverage", func() (interface{}, error) { return GetRrmGlobalCoverage(client, ctx) }},
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
	filename := fmt.Sprintf("rrm_global_oper_comprehensive_data_%d.json", time.Now().Unix())
	if err := testutil.SaveTestDataToFile(filename, allData); err != nil {
		t.Logf("Warning: Failed to save comprehensive data to %s: %v", filename, err)
	} else {
		t.Logf("Comprehensive RRM global operational data saved to %s", filename)
	}
}

// TestRrmGlobalOperDataStructures tests the basic structure of RRM global operational data types
func TestRrmGlobalOperDataStructures(t *testing.T) {
	// Sample RRM global operational data based on real WNC response structure
	sampleJSON := `{
		"Cisco-IOS-XE-wireless-rrm-global-oper:rrm-global-oper-data": {
			"rrm-one-shot-counters": [
				{
					"phy-type": "802.11a",
					"power-counter": 1250
				    },
				{
					"phy-type": "802.11b/g",
					"power-counter": 980
				    }
			],
			"rrm-channel-params": [
				{
					"phy-type": "802.11a",
					"min-dwell": 100,
					"avg-dwell": 250,
					"max-dwell": 500,
					"min-rssi": -85,
					"max-rssi": -35,
					"avg-rssi": -60,
					"channel-counter": 36
				    }
			],
			"spectrum-aq-worst-table": [
				{
					"band-id": 5,
					"detecting-ap-name": "AP001",
					"channel-num": 36,
					"min-aqi": 65,
					"aqi": 75,
					"total-intf-device-count": 3,
					"wtp-ca-si-capable": "yes",
					"scan-radio-type": "5GHz"
				    }
			],
			"radio-oper-data-24g": [
				{
					"wtp-mac": "00:11:22:33:44:55",
					"radio-slot-id": 0,
					"ap-mac": "aa:bb:cc:dd:ee:ff",
					"slot-id": 0,
					"name": "2.4GHz-Radio",
					"spectrum-capable": ["CleanAir"],
					"num-slots": 2,
					"mesh-radio-role": "access",
					"ap-up-time": "2024-01-15T10:30:00Z",
					"capwap-up-time": "2024-01-15T10:30:15Z"
				    }
			],
			"radio-oper-data-5g": [
				{
					"wtp-mac": "00:11:22:33:44:55",
					"radio-slot-id": 1,
					"ap-mac": "aa:bb:cc:dd:ee:ff",
					"slot-id": 1,
					"name": "5GHz-Radio",
					"spectrum-capable": ["CleanAir"],
					"num-slots": 2,
					"mesh-radio-role": "access",
					"ap-up-time": "2024-01-15T10:30:00Z",
					"capwap-up-time": "2024-01-15T10:30:15Z"
				    }
			],
			"radio-oper-data-6g": [],
			"spectrum-band-config-data": [
				{
					"band-id": 5,
					"spectrum-enabled": true,
					"spectrum-mode": "both",
					"duty-cycle": 50
				    }
			],
			"radio-oper-data-dualband": [],
			"rrm-client-data": [
				{
					"client-mac": "11:22:33:44:55:66",
					"ap-mac": "aa:bb:cc:dd:ee:ff",
					"slot-id": 1,
					"rssi": -45,
					"snr": 35
				    }
			],
			"rrm-fra-stats": {
				"dual-band-monitor-to-24ghz": 145,
				"dual-band-monitor-to-5ghz": 267,
				"dual-band-24ghz-to-5ghz": 89,
				"dual-band-24ghz-to-monitor": 73,
				"dual-band-5ghz-to-24ghz": 56,
				"dual-band-5ghz-to-monitor": 184,
				"sec-radio-monitor-to-5ghz": 92,
				"sec-radio-5ghz-to-monitor": 108
			    },
			"rrm-coverage": [
				{
					"wtp-mac": "aa:bb:cc:dd:ee:ff",
					"radio-slot-id": 1,
					"failed-client-count": 2,
					"snr-info": []
				    }
			]
		    }
	    }`

	// Test unmarshaling into RrmGlobalOperResponse
	var globalOper RrmGlobalOperResponse
	err := json.Unmarshal([]byte(sampleJSON), &globalOper)
	if err != nil {
		t.Fatalf("Failed to unmarshal RrmGlobalOperResponse: %v", err)
	}

	operData := globalOper.CiscoIOSXEWirelessRrmGlobalOperData

	// Test one-shot counters
	if len(operData.RrmOneShotCounters) == 0 {
		t.Error("Expected at least one RRM one-shot counter")
	} else {
		counter := operData.RrmOneShotCounters[0]
		if counter.PhyType != "802.11a" {
			t.Errorf("Expected phy type '802.11a', got '%s'", counter.PhyType)
		}
		if counter.PowerCounter != 1250 {
			t.Errorf("Expected power counter 1250, got %d", counter.PowerCounter)
		}
	}

	// Test channel parameters
	if len(operData.RrmChannelParams) == 0 {
		t.Error("Expected at least one RRM channel parameter")
	} else {
		param := operData.RrmChannelParams[0]
		if param.PhyType != "802.11a" {
			t.Errorf("Expected phy type '802.11a', got '%s'", param.PhyType)
		}
		if param.AvgDwell != 250 {
			t.Errorf("Expected avg dwell 250, got %d", param.AvgDwell)
		}
		if param.ChannelCounter != 36 {
			t.Errorf("Expected channel counter 36, got %d", param.ChannelCounter)
		}
	}

	// Test spectrum air quality worst table
	if len(operData.SpectrumAqWorstTable) == 0 {
		t.Error("Expected at least one spectrum AQ worst table entry")
	} else {
		entry := operData.SpectrumAqWorstTable[0]
		if entry.BandID != 5 {
			t.Errorf("Expected band ID 5, got %d", entry.BandID)
		}
		if entry.DetectingApName != "AP001" {
			t.Errorf("Expected detecting AP name 'AP001', got '%s'", entry.DetectingApName)
		}
		if entry.Aqi != 75 {
			t.Errorf("Expected AQI 75, got %d", entry.Aqi)
		}
	}

	// Test 2.4GHz radio operational data
	if len(operData.RadioOperData24G) == 0 {
		t.Error("Expected at least one 2.4GHz radio operational data entry")
	} else {
		radio := operData.RadioOperData24G[0]
		if radio.WtpMac != "00:11:22:33:44:55" {
			t.Errorf("Expected WTP MAC '00:11:22:33:44:55', got '%s'", radio.WtpMac)
		}
		if radio.RadioSlotID != 0 {
			t.Errorf("Expected radio slot ID 0, got %d", radio.RadioSlotID)
		}
		if radio.MeshRadioRole != "access" {
			t.Errorf("Expected mesh radio role 'access', got '%s'", radio.MeshRadioRole)
		}
	}

	// Test 5GHz radio operational data
	if len(operData.RadioOperData5G) == 0 {
		t.Error("Expected at least one 5GHz radio operational data entry")
	} else {
		radio := operData.RadioOperData5G[0]
		if radio.RadioSlotID != 1 {
			t.Errorf("Expected radio slot ID 1, got %d", radio.RadioSlotID)
		}
	}

	// Test RRM FRA stats
	fraStats := operData.RrmFraStats
	if fraStats.DualBandMonitorTo24Ghz != 145 {
		t.Errorf("Expected DualBandMonitorTo24Ghz 145, got %d", fraStats.DualBandMonitorTo24Ghz)
	}
	if fraStats.DualBandMonitorTo5Ghz != 267 {
		t.Errorf("Expected DualBandMonitorTo5Ghz 267, got %d", fraStats.DualBandMonitorTo5Ghz)
	}

	// Test RRM coverage
	if len(operData.RrmCoverage) == 0 {
		t.Error("Expected at least one RRM coverage entry")
	} else {
		coverage := operData.RrmCoverage[0]
		if coverage.WtpMac != "aa:bb:cc:dd:ee:ff" {
			t.Errorf("Expected WTP MAC 'aa:bb:cc:dd:ee:ff', got '%s'", coverage.WtpMac)
		}
		if coverage.FailedClientCount != 2 {
			t.Errorf("Expected failed client count 2, got %d", coverage.FailedClientCount)
		}
	}

	// Test individual response structures
	sampleOneShotJSON := `{
		"Cisco-IOS-XE-wireless-rrm-global-oper:rrm-one-shot-counters": [
			{
				"phy-type": "802.11b/g",
				"power-counter": 750
			    }
		]
	    }`

	var oneShotResp RrmOneShotCountersResponse
	err = json.Unmarshal([]byte(sampleOneShotJSON), &oneShotResp)
	if err != nil {
		t.Fatalf("Failed to unmarshal RrmOneShotCountersResponse: %v", err)
	}

	if len(oneShotResp.RrmOneShotCounters) == 0 {
		t.Error("Expected at least one one-shot counter in response")
	}

	_, err = json.Marshal(globalOper)
	if err != nil {
		t.Errorf("Failed to marshal RrmGlobalOperResponse back to JSON: %v", err)
	}

	_, err = json.Marshal(oneShotResp)
	if err != nil {
		t.Errorf("Failed to marshal RrmOneShotCountersResponse back to JSON: %v", err)
	}
}
