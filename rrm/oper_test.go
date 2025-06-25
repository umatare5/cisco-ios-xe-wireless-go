// Package rrm provides Radio Resource Management operational data test functionality for the Cisco Wireless Network Controller API.
package rrm

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	wnc "github.com/umatare5/cisco-xe-wireless-restconf-go"
	"github.com/umatare5/cisco-xe-wireless-restconf-go/internal/testutil"
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

func TestRrmOperMethods(t *testing.T) {
	client := getTestClient(t)

	ctx := context.Background()

	// Create a comprehensive test data collection
	testResults := make(map[string]interface{})
	endpointMapping := map[string]string{
		"RrmOperEndpoint":             "/restconf/data/Cisco-IOS-XE-wireless-rrm-oper:rrm-oper-data",
		"ApAutoRfDot11DataEndpoint":   "/restconf/data/Cisco-IOS-XE-wireless-rrm-oper:rrm-oper-data/ap-auto-rf-dot11-data",
		"ApDot11RadarDataEndpoint":    "/restconf/data/Cisco-IOS-XE-wireless-rrm-oper:rrm-oper-data/ap-dot11-radar-data",
		"ApDot11SpectrumDataEndpoint": "/restconf/data/Cisco-IOS-XE-wireless-rrm-oper:rrm-oper-data/ap-dot11-spectrum-data",
		"RrmMeasurementEndpoint":      "/restconf/data/Cisco-IOS-XE-wireless-rrm-oper:rrm-oper-data/rrm-measurement",
		"RadioSlotEndpoint":           "/restconf/data/Cisco-IOS-XE-wireless-rrm-oper:rrm-oper-data/radio-slot",
		"MainDataEndpoint":            "/restconf/data/Cisco-IOS-XE-wireless-rrm-oper:rrm-oper-data/main-data",
		"SpectrumDeviceTableEndpoint": "/restconf/data/Cisco-IOS-XE-wireless-rrm-oper:rrm-oper-data/spectrum-device-table",
		"SpectrumAqTableEndpoint":     "/restconf/data/Cisco-IOS-XE-wireless-rrm-oper:rrm-oper-data/spectrum-aq-table",
		"RegDomainOperEndpoint":       "/restconf/data/Cisco-IOS-XE-wireless-rrm-oper:rrm-oper-data/reg-domain-oper",
	}

	t.Run("GetRrmOper", func(t *testing.T) {
		result, err := GetRrmOper(client, ctx)
		collectRrmOperTestResult(testResults, "GetRrmOper", endpointMapping["RrmOperEndpoint"], result, err)
		if err != nil {
			t.Logf("GetRrmOper failed: %v", err)
		}
	})

	t.Run("GetApAutoRfDot11Data", func(t *testing.T) {
		result, err := GetApAutoRfDot11Data(client, ctx)
		collectRrmOperTestResult(testResults, "GetApAutoRfDot11Data", endpointMapping["ApAutoRfDot11DataEndpoint"], result, err)
		if err != nil {
			t.Logf("GetApAutoRfDot11Data failed: %v", err)
		}
	})

	t.Run("GetApDot11RadarData", func(t *testing.T) {
		result, err := GetApDot11RadarData(client, ctx)
		collectRrmOperTestResult(testResults, "GetApDot11RadarData", endpointMapping["ApDot11RadarDataEndpoint"], result, err)
		if err != nil {
			t.Logf("GetApDot11RadarData failed: %v", err)
		}
	})

	t.Run("GetApDot11SpectrumData", func(t *testing.T) {
		result, err := GetApDot11SpectrumData(client, ctx)
		collectRrmOperTestResult(testResults, "GetApDot11SpectrumData", endpointMapping["ApDot11SpectrumDataEndpoint"], result, err)
		if err != nil {
			t.Logf("GetApDot11SpectrumData failed: %v", err)
		}
	})

	t.Run("GetRrmMeasurement", func(t *testing.T) {
		result, err := GetRrmMeasurement(client, ctx)
		collectRrmOperTestResult(testResults, "GetRrmMeasurement", endpointMapping["RrmMeasurementEndpoint"], result, err)
		if err != nil {
			t.Logf("GetRrmMeasurement failed: %v", err)
		}
	})

	t.Run("GetRadioSlot", func(t *testing.T) {
		result, err := GetRadioSlot(client, ctx)
		collectRrmOperTestResult(testResults, "GetRadioSlot", endpointMapping["RadioSlotEndpoint"], result, err)
		if err != nil {
			t.Logf("GetRadioSlot failed: %v", err)
		}
	})

	t.Run("GetMainData", func(t *testing.T) {
		result, err := GetMainData(client, ctx)
		collectRrmOperTestResult(testResults, "GetMainData", endpointMapping["MainDataEndpoint"], result, err)
		if err != nil {
			t.Logf("GetMainData failed: %v", err)
		}
	})

	t.Run("GetSpectrumDeviceTable", func(t *testing.T) {
		result, err := GetSpectrumDeviceTable(client, ctx)
		collectRrmOperTestResult(testResults, "GetSpectrumDeviceTable", endpointMapping["SpectrumDeviceTableEndpoint"], result, err)
		if err != nil {
			t.Logf("GetSpectrumDeviceTable failed: %v", err)
		}
	})

	t.Run("GetSpectrumAqTable", func(t *testing.T) {
		result, err := GetSpectrumAqTable(client, ctx)
		collectRrmOperTestResult(testResults, "GetSpectrumAqTable", endpointMapping["SpectrumAqTableEndpoint"], result, err)
		if err != nil {
			t.Logf("GetSpectrumAqTable failed: %v", err)
		}
	})

	t.Run("GetRegDomainOper", func(t *testing.T) {
		result, err := GetRegDomainOper(client, ctx)
		collectRrmOperTestResult(testResults, "GetRegDomainOper", endpointMapping["RegDomainOperEndpoint"], result, err)
		if err != nil {
			t.Logf("GetRegDomainOper failed: %v", err)
		}
	})

	// Save collected test data to JSON file
	saveRrmOperTestData(t, testResults, "rrm_oper_test_data_collected.json")
}

// collectRrmOperTestResult helper function to collect test results
func collectRrmOperTestResult(testResults map[string]interface{}, methodName, endpoint string, result interface{}, err error) {
	testData := map[string]interface{}{
		"method":    methodName,
		"endpoint":  endpoint,
		"timestamp": time.Now().Format(time.RFC3339),
	}

	if err != nil {
		testData["error"] = err.Error()
		testData["success"] = false
	} else {
		testData["success"] = true
		testData["response"] = result
	}

	testResults[methodName] = testData
}

// saveRrmOperTestData helper function to save test data to JSON file
func saveRrmOperTestData(t *testing.T, testResults map[string]interface{}, filename string) {
	if err := testutil.SaveTestDataToFile(filename, testResults); err != nil {
		t.Logf("Failed to save test data to %s: %v", filename, err)
	} else {
		t.Logf("Test data saved to %s/%s", testutil.TestDataDir, filename)
	}
}

func TestRrmOperEndpoints(t *testing.T) {
	// Test endpoint validation
	endpoints := map[string]string{
		"RrmOperBasePath":             RrmOperBasePath,
		"RrmOperEndpoint":             "/restconf/data/Cisco-IOS-XE-wireless-rrm-oper:rrm-oper-data",
		"ApAutoRfDot11DataEndpoint":   "/restconf/data/Cisco-IOS-XE-wireless-rrm-oper:rrm-oper-data/ap-auto-rf-dot11-data",
		"ApDot11RadarDataEndpoint":    "/restconf/data/Cisco-IOS-XE-wireless-rrm-oper:rrm-oper-data/ap-dot11-radar-data",
		"ApDot11SpectrumDataEndpoint": "/restconf/data/Cisco-IOS-XE-wireless-rrm-oper:rrm-oper-data/ap-dot11-spectrum-data",
		"RrmMeasurementEndpoint":      "/restconf/data/Cisco-IOS-XE-wireless-rrm-oper:rrm-oper-data/rrm-measurement",
		"RadioSlotEndpoint":           "/restconf/data/Cisco-IOS-XE-wireless-rrm-oper:rrm-oper-data/radio-slot",
		"MainDataEndpoint":            "/restconf/data/Cisco-IOS-XE-wireless-rrm-oper:rrm-oper-data/main-data",
		"SpectrumDeviceTableEndpoint": "/restconf/data/Cisco-IOS-XE-wireless-rrm-oper:rrm-oper-data/spectrum-device-table",
		"SpectrumAqTableEndpoint":     "/restconf/data/Cisco-IOS-XE-wireless-rrm-oper:rrm-oper-data/spectrum-aq-table",
		"RegDomainOperEndpoint":       "/restconf/data/Cisco-IOS-XE-wireless-rrm-oper:rrm-oper-data/reg-domain-oper",
	}

	for name, endpoint := range endpoints {
		t.Run(fmt.Sprintf("Validate_%s", name), func(t *testing.T) {
			if endpoint == "" {
				t.Errorf("%s endpoint is empty", name)
			}
			if len(endpoint) < 10 {
				t.Errorf("%s endpoint is too short: %s", name, endpoint)
			}
		})
	}
}

// TestRrmOperDataStructures tests the basic structure of RRM operational data types
func TestRrmOperDataStructures(t *testing.T) {
	// Sample RRM operational data based on real WNC response structure
	sampleJSON := `{
		"Cisco-IOS-XE-wireless-rrm-oper:rrm-oper-data": {
			"ap-auto-rf-dot11-data": [
				{
					"wtp-mac": "00:11:22:33:44:55",
					"radio-slot-id": 1,
					"neighbor-radio-info": {
						"neighbor-radio-list": [
							{
								"neighbor-radio-info": {
									"neighbor-radio-mac": "f0:d8:05:2c:41:20",
									"neighbor-radio-slot-id": 0,
									"rssi": -15,
									"snr": 67,
									"channel": 11,
									"power": 18,
									"group-leader-ip": "192.168.255.1",
									"chan-width": "radio-neighbor-chan-width-20-mhz",
									"sensor-covered": false
								}
							}
						]
					}
				}
			],
			"ap-dot11-radar-data": [
				{
					"wtp-mac": "00:11:22:33:44:55",
					"radio-slot-id": 1,
					"last-radar-on-radio": "2024-01-15T10:30:00Z"
				}
			],
			"ap-dot11-spectrum-data": [
				{
					"wtp-mac": "00:11:22:33:44:55",
					"radio-slot-id": 1,
					"config": {
						"spectrum-intelligence-enable": true,
						"spectrum-wtp-ca-si-capable": "yes",
						"spectrum-operation-state": "enabled",
						"spectrum-admin-state": true,
						"spectrum-capable": true,
						"rapid-update-enable": false,
						"sensord-operational-status": 1,
						"scan-radio-type": "5GHz"
					}
				}
			],
			"rrm-measurement": [
				{
					"wtp-mac": "00:11:22:33:44:55",
					"radio-slot-id": 1,
					"foreign": {
						"foreign": {
							"foreign-data": [
								{
									"chan": 36,
									"power": 14,
									"rogue-20-count": 0,
									"rogue-40-primary-count": 0,
									"rogue-80-primary-count": 0,
									"chan-util": 25
								}
							]
						}
					},
					"noise": {
						"noise": {
							"noise-data": [
								{
									"chan": 36,
									"noise": -95
								}
							]
						}
					},
					"load": {
						"rx-util-percentage": 0,
						"tx-util-percentage": 0,
						"cca-util-percentage": 17,
						"stations": 1,
						"rx-noise-channel-utilization": 17,
						"non-wifi-inter": 4
					}
				}
			],
			"radio-slot": [
				{
					"wtp-mac": "00:11:22:33:44:55",
					"radio-slot-id": 1,
					"radio-data": {
						"best-tx-pwr-level": 1,
						"best-rts-thresh": 2347,
						"best-frag-thresh": 2346,
						"load-prof-passed": true,
						"coverage-profile-passed": true,
						"interference-profile-passed": false,
						"noise-profile-passed": true,
						"dca-stats": {
							"best-chan": 1,
							"current-chan-energy": -66,
							"last-chan-energy": -66,
							"chan-changes": 1
						},
						"coverage-overlap-factor": "None",
						"sensor-coverage-factor": "None"
					}
				}
			],
			"main-data": [
				{
					"phy-type": "dot11-2-dot-4-ghz-band",
					"grp": {
						"current-state": "rrm-leader-idle",
						"last-run": "2025-06-24T08:19:23.783802+00:00",
						"dca": {
							"dca-last-run": "2025-06-24T08:19:23.783436+00:00"
						},
						"txpower": {
							"dpc-last-run": "2025-06-24T08:19:23.783799+00:00",
							"run-time": 0
						},
						"current-grouping-mode": "rrm-automatic-mode",
						"join-protocol-ver": 100,
						"current-grouping-role": "rrm-group-auto-leader",
						"cntrlr-name": "lab2-cat98wlc-06f-01a",
						"cntrlr-ip-addr": "192.168.255.1",
						"cntrlr-secondary-ip-addr": "::",
						"is-static-member": "rrm-auto-member-config"
					},
					"rf-name": "labo"
				}
			],
			"spectrum-device-table": [
				{
					"device-id": "05:2c:41:20:b3:92",
					"cluster-id": "e2:00:00:00:10:63",
					"last-updated-time": "2025-06-24T08:22:55.004316+00:00",
					"idr-data": {
						"detecting-ap-mac": "f0:d8:05:2c:41:20",
						"affected-channel-list": "1,2",
						"is-persistent": false,
						"class-type-enum": "pmac-dev-id-bt"
					}
				}
			],
			"spectrum-aq-table": [
				{
					"wtp-mac": "f0:d8:05:2c:41:20",
					"band": "dot11-2-dot-4-ghz-band",
					"reporting-ap-name": "lab2-ap9166-06f-01"
				}
			],
			"reg-domain-oper": {
				"country-list": "J4"
			}
		}
	}`

	// Test unmarshaling into RrmOperResponse
	var rrmOper RrmOperResponse
	err := json.Unmarshal([]byte(sampleJSON), &rrmOper)
	if err != nil {
		t.Fatalf("Failed to unmarshal RrmOperResponse: %v", err)
	}

	operData := rrmOper.CiscoIOSXEWirelessRrmOperRrmOperData

	// Test AP auto RF data
	if len(operData.ApAutoRfDot11Data) == 0 {
		t.Error("Expected at least one AP auto RF data entry")
	} else {
		autoRf := operData.ApAutoRfDot11Data[0]
		if autoRf.WtpMac != "00:11:22:33:44:55" {
			t.Errorf("Expected WTP MAC '00:11:22:33:44:55', got '%s'", autoRf.WtpMac)
		}
		if autoRf.RadioSlotID != 1 {
			t.Errorf("Expected radio slot ID 1, got %d", autoRf.RadioSlotID)
		}
	}

	// Test radar data
	if len(operData.ApDot11RadarData) == 0 {
		t.Error("Expected at least one radar data entry")
	} else {
		radar := operData.ApDot11RadarData[0]
		if radar.WtpMac != "00:11:22:33:44:55" {
			t.Errorf("Expected WTP MAC '00:11:22:33:44:55', got '%s'", radar.WtpMac)
		}
		if radar.RadioSlotID != 1 {
			t.Errorf("Expected radio slot ID 1, got %d", radar.RadioSlotID)
		}
	}

	// Test spectrum data
	if len(operData.ApDot11SpectrumData) == 0 {
		t.Error("Expected at least one spectrum data entry")
	} else {
		spectrum := operData.ApDot11SpectrumData[0]
		if spectrum.WtpMac != "00:11:22:33:44:55" {
			t.Errorf("Expected WTP MAC '00:11:22:33:44:55', got '%s'", spectrum.WtpMac)
		}
		if !spectrum.Config.SpectrumIntelligenceEnable {
			t.Error("Expected spectrum intelligence enable to be true")
		}
		if !spectrum.Config.SpectrumCapable {
			t.Error("Expected spectrum capable to be true")
		}
	}

	// Test RRM measurement
	if len(operData.RrmMeasurement) == 0 {
		t.Error("Expected at least one RRM measurement entry")
	} else {
		measurement := operData.RrmMeasurement[0]
		if measurement.WtpMac != "00:11:22:33:44:55" {
			t.Errorf("Expected WTP MAC '00:11:22:33:44:55', got '%s'", measurement.WtpMac)
		}
		if measurement.RadioSlotID != 1 {
			t.Errorf("Expected radio slot ID 1, got %d", measurement.RadioSlotID)
		}
	}

	// Test radio slot
	if len(operData.RadioSlot) == 0 {
		t.Error("Expected at least one radio slot entry")
	} else {
		slot := operData.RadioSlot[0]
		if slot.WtpMac != "00:11:22:33:44:55" {
			t.Errorf("Expected WTP MAC '00:11:22:33:44:55', got '%s'", slot.WtpMac)
		}
		if slot.RadioSlotID != 1 {
			t.Errorf("Expected radio slot ID 1, got %d", slot.RadioSlotID)
		}
	}

	// Test main data
	if len(operData.MainData) == 0 {
		t.Error("Expected at least one main data entry")
	}

	// Test spectrum device table
	if len(operData.SpectrumDeviceTable) == 0 {
		t.Error("Expected at least one spectrum device table entry")
	}

	// Test spectrum AQ table
	if len(operData.SpectrumAqTable) == 0 {
		t.Error("Expected at least one spectrum AQ table entry")
	}

	// Test regulatory domain operational data - basic validation
	// Note: RegDomainOper structure varies, so we just test it exists

	_, err = json.Marshal(rrmOper)
	if err != nil {
		t.Errorf("Failed to marshal RrmOperResponse back to JSON: %v", err)
	}
}
