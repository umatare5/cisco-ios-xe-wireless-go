// Package ap provides access point global operational data test functionality for the Cisco Wireless Network Controller API.
package ap

import (
	"context"
	"encoding/json"
	"testing"

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

// APGlobalOperTestDataCollector holds test data for AP global operation functions
type APGlobalOperTestDataCollector struct {
	Data map[string]interface{} `json:"ap_global_oper_test_data"`
}

// newAPGlobalOperTestDataCollector creates a new test data collector
func newAPGlobalOperTestDataCollector() *APGlobalOperTestDataCollector {
	return &APGlobalOperTestDataCollector{
		Data: make(map[string]interface{}),
	}
}

func runAPGlobalOperTestAndCollectData(t *testing.T, collector *APGlobalOperTestDataCollector, testName string, testFunc func() (interface{}, error)) {
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
// 2. INTEGRATION TESTS (API Endpoint Testing with Live Data Validation)
// =============================================================================

func TestAPGlobalOperationFunctions(t *testing.T) {
	collector := newAPGlobalOperTestDataCollector()
	client := getTestClient(t)

	ctx, cancel := context.WithTimeout(context.Background(), testutil.DefaultTestTimeout)
	defer cancel()

	t.Run("GetApGlobalOper", func(t *testing.T) {
		runAPGlobalOperTestAndCollectData(t, collector, "GetApGlobalOper", func() (interface{}, error) {
			return GetApGlobalOper(client, ctx)
		})
	})

	t.Run("GetApHistory", func(t *testing.T) {
		runAPGlobalOperTestAndCollectData(t, collector, "GetApHistory", func() (interface{}, error) {
			return GetApHistory(client, ctx)
		})
	})

	t.Run("GetApEwlcApStats", func(t *testing.T) {
		runAPGlobalOperTestAndCollectData(t, collector, "GetApEwlcApStats", func() (interface{}, error) {
			return GetApEwlcApStats(client, ctx)
		})
	})

	t.Run("GetApImgPredownloadStats", func(t *testing.T) {
		runAPGlobalOperTestAndCollectData(t, collector, "GetApImgPredownloadStats", func() (interface{}, error) {
			return GetApImgPredownloadStats(client, ctx)
		})
	})

	t.Run("GetApJoinStats", func(t *testing.T) {
		runAPGlobalOperTestAndCollectData(t, collector, "GetApJoinStats", func() (interface{}, error) {
			return GetApJoinStats(client, ctx)
		})
	})

	t.Run("GetApWlanClientStats", func(t *testing.T) {
		runAPGlobalOperTestAndCollectData(t, collector, "GetApWlanClientStats", func() (interface{}, error) {
			return GetApWlanClientStats(client, ctx)
		})
	})

	t.Run("GetApEmltdJoinCountStat", func(t *testing.T) {
		runAPGlobalOperTestAndCollectData(t, collector, "GetApEmltdJoinCountStat", func() (interface{}, error) {
			return GetApEmltdJoinCountStat(client, ctx)
		})
	})

	// Save collected test data to file
	if len(collector.Data) > 0 {
		if err := testutil.SaveTestDataToFile("ap_global_oper_test_data_collected.json", collector.Data); err != nil {
			t.Logf("Warning: Could not save test data: %v", err)
		} else {
			t.Logf("AP global operation test data saved to %s/ap_global_oper_test_data_collected.json", testutil.TestDataDir)
		}
	}
}

// TestAPGlobalOperEndpoints tests the AP global operation endpoint constants
func TestAPGlobalOperEndpoints(t *testing.T) {
	expectedEndpoints := map[string]string{
		"ApGlobalOperBasePath":          "/restconf/data/Cisco-IOS-XE-wireless-ap-global-oper:ap-global-oper-data",
		"ApGlobalOperEndpoint":          "/restconf/data/Cisco-IOS-XE-wireless-ap-global-oper:ap-global-oper-data",
		"ApHistoryEndpoint":             "/restconf/data/Cisco-IOS-XE-wireless-ap-global-oper:ap-global-oper-data/ap-history",
		"EwlcApStatsEndpoint":           "/restconf/data/Cisco-IOS-XE-wireless-ap-global-oper:ap-global-oper-data/ewlc-ap-stats",
		"ApImgPredownloadStatsEndpoint": "/restconf/data/Cisco-IOS-XE-wireless-ap-global-oper:ap-global-oper-data/ap-img-predownload-stats",
		"ApJoinStatsEndpoint":           "/restconf/data/Cisco-IOS-XE-wireless-ap-global-oper:ap-global-oper-data/ap-join-stats",
		"WlanClientStatsEndpoint":       "/restconf/data/Cisco-IOS-XE-wireless-ap-global-oper:ap-global-oper-data/wlan-client-stats",
		"EmltdJoinCountStatEndpoint":    "/restconf/data/Cisco-IOS-XE-wireless-ap-global-oper:ap-global-oper-data/emltd-join-count-stat",
	}

	for name, expected := range expectedEndpoints {
		t.Run(name, func(t *testing.T) {
			switch name {
			case "ApGlobalOperBasePath":
				if ApGlobalOperBasePath != expected {
					t.Errorf("Expected %s = %s, got %s", name, expected, ApGlobalOperBasePath)
				}
			case "ApGlobalOperEndpoint":
				if ApGlobalOperEndpoint != expected {
					t.Errorf("Expected %s = %s, got %s", name, expected, ApGlobalOperEndpoint)
				}
			case "ApHistoryEndpoint":
				if ApHistoryEndpoint != expected {
					t.Errorf("Expected %s = %s, got %s", name, expected, ApHistoryEndpoint)
				}
			case "EwlcApStatsEndpoint":
				if EwlcApStatsEndpoint != expected {
					t.Errorf("Expected %s = %s, got %s", name, expected, EwlcApStatsEndpoint)
				}
			case "ApImgPredownloadStatsEndpoint":
				if ApImgPredownloadStatsEndpoint != expected {
					t.Errorf("Expected %s = %s, got %s", name, expected, ApImgPredownloadStatsEndpoint)
				}
			case "ApJoinStatsEndpoint":
				if ApJoinStatsEndpoint != expected {
					t.Errorf("Expected %s = %s, got %s", name, expected, ApJoinStatsEndpoint)
				}
			case "WlanClientStatsEndpoint":
				if WlanClientStatsEndpoint != expected {
					t.Errorf("Expected %s = %s, got %s", name, expected, WlanClientStatsEndpoint)
				}
			case "EmltdJoinCountStatEndpoint":
				if EmltdJoinCountStatEndpoint != expected {
					t.Errorf("Expected %s = %s, got %s", name, expected, EmltdJoinCountStatEndpoint)
				}
			}
		})
	}
}

// TestApGlobalOperDataStructures tests the basic structure of AP Global operational data types
func TestApGlobalOperDataStructures(t *testing.T) {
	tests := []struct {
		name     string
		jsonData string
		dataType interface{}
	}{
		{
			name: "ApGlobalOperResponse",
			jsonData: `{
				"queryResponse": {
					"@type": "ApGlobalOperResponse",
					"ap-global-oper-data": {
						"ap-history": [],
						"ewlc-ap-stats": {
							"ap-count": 0,
							"ap-joined": 0,
							"ap-disconnected": 0
						    },
						"ap-img-predownload-stats": [],
						"ap-join-stats": {
							"ap-join-count": 0,
							"ap-join-fail-count": 0
						    },
						"wlan-client-stats": [],
						"emltd-join-count-stat": []
					    }
				    }
			    }`,
			dataType: &ApGlobalOperResponse{},
		},
		{
			name: "ApGlobalOperApHistoryResponse",
			jsonData: `{
				"queryResponse": {
					"@type": "ApHistoryResponse",
					"ap-history": [
						{
							"wtp-mac": "aa:bb:cc:dd:ee:ff",
							"ap-name": "` + wnc.TestAPName + `",
							"event-type": "join",
							"timestamp": "` + wnc.TestTimestamp + `"
						    }
					]
				    }
			    }`,
			dataType: &ApGlobalOperApHistoryResponse{},
		},
		{
			name: "ApGlobalOperEwlcApStatsResponse",
			jsonData: `{
				"Cisco-IOS-XE-wireless-ap-global-oper:ewlc-ap-stats": {
					"stats-80211-a-rad": {
						"total-radios": 2,
						"radios-up": 2,
						"radios-down": 0
					    },
					"stats-80211-bg-rad": {
						"total-radios": 2,
						"radios-up": 2,
						"radios-down": 0
					    },
					"stats-misconfigured-aps": 0
				    }
			    }`,
			dataType: &ApGlobalOperEwlcApStatsResponse{},
		},
		{
			name: "ApImgPredownloadStatsResponse",
			jsonData: `{
				"Cisco-IOS-XE-wireless-ap-global-oper:ap-img-predownload-stats": {
					"predownload-stats": {
						"num-initiated": 0,
						"num-in-progress": 0,
						"num-complete": 0,
						"num-unsupported": 0,
						"num-failed": 0,
						"is-predownload-in-progress": false,
						"num-total": 0
					    },
					"downloads-in-progress": 0,
					"downloads-complete": 0,
					"wlc-predownload-stats": {
						"num-initiated": 0,
						"num-in-progress": 0,
						"num-complete": 0,
						"num-unsupported": 0,
						"num-failed": 0,
						"is-predownload-in-progress": false,
						"num-total": 0
					    }
				    }
			    }`,
			dataType: &ApGlobalOperApImgPredownloadStatsResponse{},
		},
		{
			name: "ApJoinStatsResponse",
			jsonData: `{
				"Cisco-IOS-XE-wireless-ap-global-oper:ap-join-stats": [
					{
						"wtp-mac": "28:ac:9e:bb:3c:80",
						"ap-join-info": {
							"ap-ip-addr": "192.168.255.11",
							"ap-ethernet-mac": "28:ac:9e:11:48:10",
							"ap-name": "lab2-ap1815-06f-02",
							"is-joined": true,
							"num-join-req-recvd": 2,
							"num-config-req-recvd": 6,
							"last-join-failure-type": "jf-none",
							"last-config-failure-type": "cf-none",
							"last-error-type": "ap-con-failure-run",
							"last-error-time": "2025-06-14T06:12:11.467356+00:00",
							"last-msg-decr-fail-reason": "",
							"num-succ-join-resp-sent": 2,
							"num-unsucc-join-req-procn": 0,
							"num-succ-conf-resp-sent": 6,
							"num-unsucc-conf-req-procn": 0,
							"last-succ-join-atmpt-time": "2025-06-14T06:16:43.721684+00:00",
							"last-fail-join-atmpt-time": "1970-01-01T00:00:00+00:00",
							"last-succ-conf-atmpt-time": "2025-06-14T06:16:45.410469+00:00",
							"last-fail-conf-atmpt-time": "1970-01-01T00:00:00+00:00"
						    },
						"ap-disconnect-reason": "Wtp reset config cmd sent",
						"reboot-reason": "ap-reboot-reason-reboot-cmd",
						"disconnect-reason": "wtp-controller-initiated-reason"
					    }
				]
			    }`,
			dataType: &ApGlobalOperApJoinStatsResponse{},
		},
		{
			name: "WlanClientStatsResponse",
			jsonData: `{
				"Cisco-IOS-XE-wireless-ap-global-oper:wlan-client-stats": [
					{
						"wlan-id": 1,
						"wlan-name": "test-wlan",
						"client-count": 25,
						"associated-clients": 20
					    }
				]
			    }`,
			dataType: &ApGlobalOperWlanClientStatsResponse{},
		},
		{
			name: "EmltdJoinCountStatResponse",
			jsonData: `{
				"Cisco-IOS-XE-wireless-ap-global-oper:emltd-join-count-stat": {
					"joined-aps-count": 2
				    }
			    }`,
			dataType: &ApGlobalOperEmltdJoinCountStatResponse{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := json.Unmarshal([]byte(tt.jsonData), tt.dataType)
			if err != nil {
				t.Errorf("Failed to unmarshal %s: %v", tt.name, err)
			}
		})
	}
}
