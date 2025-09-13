package ap_test

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/pkg/testutil"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/ap"
)

// TestApServiceUnit_Constructor_Success tests service constructor functionality.
func TestApServiceUnit_Constructor_Success(t *testing.T) {
	t.Run("NewServiceWithValidClient", func(t *testing.T) {
		// Create mock server and test client using public API
		responses := map[string]string{
			"test-endpoint": `{"status": "success"}`,
		}
		mockServer := testutil.NewMockServer(responses)
		defer mockServer.Close()

		testClient := testutil.NewTestClient(mockServer)
		service := ap.NewService(testClient.Core().(*core.Client))

		// Test that service can be created and has a client
		if service.Client() == nil {
			t.Error("Expected service to have a client")
		}
	})

	t.Run("NewServiceWithNilClient", func(t *testing.T) {
		service := ap.NewService(nil)

		// Service should still be created even with nil client
		if service.Client() != nil {
			t.Error("Expected service with nil client to return nil from Client()")
		}
	})
}

// TestApServiceUnit_GetOperations_MockSuccess tests Get operations using mock server
// This is essential for CI environments where actual Cisco controllers are not available.
func TestApServiceUnit_GetOperations_MockSuccess(t *testing.T) {
	// Create comprehensive mock RESTCONF server with all AP endpoints
	responses := map[string]string{
		"Cisco-IOS-XE-wireless-ap-cfg:ap-cfg-data": `{
			"Cisco-IOS-XE-wireless-ap-cfg:ap-cfg-data": {
				"ap-tags": {
					"ap-tag": [{
						"ap-mac": "28:ac:9e:11:48:10",
						"policy-tag": "labo-wlan-flex",
						"site-tag": "labo-site-flex",
						"rf-tag": "labo-inside"
					}]
				},
				"tag-source-priority-configs": {
					"tag-source-priority-config": [{
						"priority": 1,
						"source": "filter"
					}]
				}
			}
		}`,
		"Cisco-IOS-XE-wireless-ap-cfg:ap-cfg-data/ap-tags": `{
			"Cisco-IOS-XE-wireless-ap-cfg:ap-tags": {
				"ap-tag": [{
					"ap-mac": "28:ac:9e:11:48:10",
					"policy-tag": "labo-wlan-flex",
					"site-tag": "labo-site-flex",
					"rf-tag": "labo-inside"
				}, {
					"ap-mac": "c4:14:a2:c9:02:70",
					"policy-tag": "labo-wlan-flex",
					"site-tag": "labo-site-flex",
					"rf-tag": "labo-inside"
				}]
			}
		}`,
		"Cisco-IOS-XE-wireless-ap-cfg:ap-cfg-data/tag-source-priority-configs": `{
			"Cisco-IOS-XE-wireless-ap-cfg:tag-source-priority-configs": {
				"tag-source-priority-config": [{
					"priority": 1,
					"source": "filter"
				}]
			}
		}`,
		"Cisco-IOS-XE-wireless-ap-global-oper:ap-global-oper-data": `{
			"Cisco-IOS-XE-wireless-ap-global-oper:ap-global-oper-data": {
				"ewlc-ap-stats": {
					"ap-count": 2,
					"ap-up": 2,
					"ap-down": 0
				},
				"ap-history": [{
					"ethernet-mac": "28:ac:9e:11:48:10",
					"ap-name": "TEST-AP01",
					"wtp-mac": "aa:bb:cc:dd:ee:ff"
				}],
				"ap-join-stats": [{
					"wtp-mac": "aa:bb:cc:dd:ee:ff",
					"ap-join-info": {
						"ap-ip-addr": "192.168.255.11",
						"ap-ethernet-mac": "28:ac:9e:11:48:10",
						"ap-name": "TEST-AP01",
						"is-joined": true
					}
				}],
				"wlan-client-stats": [{
					"wlan-id": 1,
					"client-count": 0
				}]
			}
		}`,
		"Cisco-IOS-XE-wireless-ap-global-oper:ap-global-oper-data/ewlc-ap-stats": `{
			"Cisco-IOS-XE-wireless-ap-global-oper:ewlc-ap-stats": {
				"ap-count": 10,
				"ap-up": 8,
				"ap-down": 2
			}
		}`,
		"Cisco-IOS-XE-wireless-ap-global-oper:ap-global-oper-data/ap-history": `{
			"Cisco-IOS-XE-wireless-ap-global-oper:ap-history": [{
				"ethernet-mac": "28:ac:9e:11:48:10",
				"ap-name": "TEST-AP01",
				"wtp-mac": "aa:bb:cc:dd:ee:ff"
			}]
		}`,
		"Cisco-IOS-XE-wireless-ap-global-oper:ap-global-oper-data/ap-join-stats": `{
			"Cisco-IOS-XE-wireless-ap-global-oper:ap-join-stats": [{
				"wtp-mac": "aa:bb:cc:dd:ee:ff",
				"ap-join-info": {
					"ap-ip-addr": "192.168.255.11",
					"ap-ethernet-mac": "28:ac:9e:11:48:10",
					"ap-name": "TEST-AP01",
					"is-joined": true
				}
			}]
		}`,
		"Cisco-IOS-XE-wireless-ap-global-oper:ap-global-oper-data/wlan-client-stats": `{
			"Cisco-IOS-XE-wireless-ap-global-oper:wlan-client-stats": [{
				"wlan-id": 1,
				"client-count": 5
			}]
		}`,
		"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data": `{
			"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data": {
				"capwap-data": [{
					"wtp-mac": "aa:bb:cc:dd:ee:ff",
					"ip-addr": "192.168.255.11",
					"name": "TEST-AP01",
					"device-detail": {
						"static-info": {
							"board-data": {
								"wtp-serial-num": "FGL2209B05T",
								"wtp-enet-mac": "28:ac:9e:11:48:10"
							}
						}
					}
				}],
				"ap-name-mac-map": [{
					"wtp-name": "TEST-AP01",
					"wtp-mac": "aa:bb:cc:dd:ee:ff",
					"eth-mac": "28:ac:9e:11:48:10"
				}],
				"radio-oper-data": [{
					"wtp-mac": "aa:bb:cc:dd:ee:ff",
					"radio-slot-id": 0,
					"slot-id": 0,
					"radio-type": "radio-80211bg",
					"admin-state": "enabled",
					"oper-state": "radio-up"
				}],
				"ap-radio-neighbor": [{
					"wtp-mac": "aa:bb:cc:dd:ee:ff",
					"radio-slot-id": 0,
					"bssid": "aa:bb:cc:dd:ee:ff"
				}],
				"ap-image-active-location": [{
					"wtp-mac": "aa:bb:cc:dd:ee:ff",
					"image-location": "flash:/c9800-universal-k9.16.12.07.SPA.bin"
				}],
				"ap-image-prepare-location": [{
					"wtp-mac": "aa:bb:cc:dd:ee:ff",
					"image-location": "flash:/c9800-universal-k9.16.12.07.SPA.bin"
				}],
				"ap-pwr-info": [{
					"wtp-mac": "aa:bb:cc:dd:ee:ff",
					"power-consumption": 20.5
				}],
				"ap-sensor-status": [{
					"wtp-mac": "aa:bb:cc:dd:ee:ff",
					"temperature": 42
				}],
				"capwap-pkts": [{
					"wtp-mac": "aa:bb:cc:dd:ee:ff",
					"tx-pkts": 1000,
					"rx-pkts": 950
				}]
			}
		}`,
		"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/capwap-data": `{
			"Cisco-IOS-XE-wireless-access-point-oper:capwap-data": [{
				"wtp-mac": "aa:bb:cc:dd:ee:ff",
				"ip-addr": "192.168.255.11",
				"name": "TEST-AP01",
				"device-detail": {
					"static-info": {
						"board-data": {
							"wtp-serial-num": "FGL2209B05T",
							"wtp-enet-mac": "28:ac:9e:11:48:10"
						}
					}
				}
			}]
		}`,
		"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/ap-name-mac-map": `{
			"Cisco-IOS-XE-wireless-access-point-oper:ap-name-mac-map": [{
				"wtp-name": "AP-Test-01",
				"wtp-mac": "aa:bb:cc:dd:ee:ff"
			}]
		}`,
		"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/radio-oper-data": `{
			"Cisco-IOS-XE-wireless-access-point-oper:radio-oper-data": [{
				"wtp-mac": "aa:bb:cc:dd:ee:ff",
				"radio-slot-id": 0,
				"oper-state": "up"
			}]
		}`,
		"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/ap-radio-neighbor": `{
			"Cisco-IOS-XE-wireless-access-point-oper:ap-radio-neighbor": [{
				"wtp-mac": "aa:bb:cc:dd:ee:ff",
				"radio-slot-id": 0,
				"bssid": "bb:cc:dd:ee:ff:aa"
			}]
		}`,
		"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/ap-image-active-location": `{
			"Cisco-IOS-XE-wireless-access-point-oper:ap-image-active-location": [{
				"wtp-mac": "aa:bb:cc:dd:ee:ff",
				"image-location": "flash:ap3g2-k9w8-mx.152-4.JB6"
			}]
		}`,
		"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/ap-image-prepare-location": `{
			"Cisco-IOS-XE-wireless-access-point-oper:ap-image-prepare-location": [{
				"wtp-mac": "aa:bb:cc:dd:ee:ff",
				"image-location": "flash:ap3g2-k9w8-mx.152-4.JB6"
			}]
		}`,
		"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/ap-pwr-info": `{
			"Cisco-IOS-XE-wireless-access-point-oper:ap-pwr-info": [{
				"wtp-mac": "aa:bb:cc:dd:ee:ff",
				"power-consumption": 15.5
			}]
		}`,
		"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/ap-sensor-status": `{
			"Cisco-IOS-XE-wireless-access-point-oper:ap-sensor-status": [{
				"wtp-mac": "aa:bb:cc:dd:ee:ff",
				"temperature": 45
			}]
		}`,
		"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/capwap-pkts": `{
			"Cisco-IOS-XE-wireless-access-point-oper:capwap-pkts": [{
				"wtp-mac": "aa:bb:cc:dd:ee:ff",
				"tx-pkts": 1000,
				"rx-pkts": 2000
			}]
		}`,
	}
	mockServer := testutil.NewMockServer(responses)
	defer mockServer.Close()

	// Create test client configured for the mock server
	testClient := testutil.NewTestClient(mockServer)
	service := ap.NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	// Test basic configuration operations
	t.Run("GetConfig", func(t *testing.T) {
		result, err := service.GetConfig(ctx)
		if err != nil {
			t.Errorf("Expected no error for GetConfig, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for GetConfig, got nil")
		}
	})

	t.Run("ListTagConfigs", func(t *testing.T) {
		result, err := service.ListTagConfigs(ctx)
		if err != nil {
			t.Errorf("Expected no error for ListTagConfigs, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for ListTagConfigs, got nil")
		}
	})

	t.Run("ListTagSourcePriorityConfigs", func(t *testing.T) {
		result, err := service.ListTagSourcePriorityConfigs(ctx)
		if err != nil {
			t.Errorf("Expected no error for ListTagSourcePriorityConfigs, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for ListTagSourcePriorityConfigs, got nil")
		}
	})

	// Test global operational data operations
	t.Run("GetGlobalInfo", func(t *testing.T) {
		result, err := service.GetGlobalInfo(ctx)
		if err != nil {
			t.Errorf("Expected no error for GetGlobalInfo, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for GetGlobalInfo, got nil")
		}
	})

	t.Run("GetEWLCAPStats", func(t *testing.T) {
		result, err := service.GetEWLCAPStats(ctx)
		if err != nil {
			t.Errorf("Expected no error for GetEWLCAPStats, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for GetEWLCAPStats, got nil")
		}
	})

	t.Run("ListAPHistory", func(t *testing.T) {
		result, err := service.ListAPHistory(ctx)
		if err != nil {
			t.Errorf("Expected no error for ListAPHistory, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for ListAPHistory, got nil")
		}
	})

	t.Run("ListAPJoinStats", func(t *testing.T) {
		result, err := service.ListAPJoinStats(ctx)
		if err != nil {
			t.Errorf("Expected no error for ListAPJoinStats, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for ListAPJoinStats, got nil")
		}
	})

	t.Run("ListWLANClientStats", func(t *testing.T) {
		result, err := service.ListWLANClientStats(ctx)
		if err != nil {
			t.Errorf("Expected no error for ListWLANClientStats, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for ListWLANClientStats, got nil")
		}
	})

	// Test AP operational data operations
	t.Run("GetOperational", func(t *testing.T) {
		result, err := service.GetOperational(ctx)
		if err != nil {
			t.Errorf("Expected no error for GetOperational, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for GetOperational, got nil")
		}
	})

	t.Run("ListCAPWAPData", func(t *testing.T) {
		result, err := service.ListCAPWAPData(ctx)
		if err != nil {
			t.Errorf("Expected no error for ListCAPWAPData, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for ListCAPWAPData, got nil")
		}
	})

	t.Run("ListNameMACMaps", func(t *testing.T) {
		result, err := service.ListNameMACMaps(ctx)
		if err != nil {
			t.Errorf("Expected no error for ListNameMACMaps, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for ListNameMACMaps, got nil")
		}
	})

	t.Run("ListRadioStatus", func(t *testing.T) {
		result, err := service.ListRadioStatus(ctx)
		if err != nil {
			t.Errorf("Expected no error for ListRadioStatus, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for ListRadioStatus, got nil")
		}
	})

	t.Run("ListRadioNeighbors", func(t *testing.T) {
		result, err := service.ListRadioNeighbors(ctx)
		if err != nil {
			t.Errorf("Expected no error for ListRadioNeighbors, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for ListRadioNeighbors, got nil")
		}
	})

	t.Run("ListActiveImageLocations", func(t *testing.T) {
		result, err := service.ListActiveImageLocations(ctx)
		if err != nil {
			t.Errorf("Expected no error for ListActiveImageLocations, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for ListActiveImageLocations, got nil")
		}
	})

	t.Run("ListPreparedImageLocations", func(t *testing.T) {
		result, err := service.ListPreparedImageLocations(ctx)
		if err != nil {
			t.Errorf("Expected no error for ListPreparedImageLocations, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for ListPreparedImageLocations, got nil")
		}
	})

	t.Run("ListPowerInfo", func(t *testing.T) {
		result, err := service.ListPowerInfo(ctx)
		if err != nil {
			t.Errorf("Expected no error for ListPowerInfo, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for ListPowerInfo, got nil")
		}
	})

	t.Run("ListSensorStatus", func(t *testing.T) {
		result, err := service.ListSensorStatus(ctx)
		if err != nil {
			t.Errorf("Expected no error for ListSensorStatus, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for ListSensorStatus, got nil")
		}
	})

	t.Run("ListCAPWAPPackets", func(t *testing.T) {
		result, err := service.ListCAPWAPPackets(ctx)
		if err != nil {
			t.Errorf("Expected no error for ListCAPWAPPackets, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for ListCAPWAPPackets, got nil")
		}
	})
}

// TestApServiceUnit_GetOperations_ErrorHandling tests error scenarios using mock server.
func TestApServiceUnit_GetOperations_ErrorHandling(t *testing.T) {
	// Create mock server that returns 404 for AP endpoints
	errorPaths := []string{
		"Cisco-IOS-XE-wireless-ap-cfg:ap-cfg-data",
		"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data",
	}
	mockServer := testutil.NewMockErrorServer(errorPaths, 404)
	defer mockServer.Close()

	testClient := testutil.NewTestClient(mockServer)
	service := ap.NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	// Test that GetConfig properly handles 404 errors
	_, err := service.GetConfig(ctx)
	if err == nil {
		t.Error("Expected error for 404 response, got nil")
	}

	// Verify error contains expected information
	if !core.IsNotFoundError(err) {
		t.Errorf("Expected NotFound error, got: %v", err)
	}
}

// TestApServiceUnit_GetOperations_FilteredSuccess tests filtered Get operations.
func TestApServiceUnit_GetOperations_FilteredSuccess(t *testing.T) {
	// Create mock server with query-based responses
	responses := map[string]string{
		"Cisco-IOS-XE-wireless-ap-cfg:ap-cfg-data/ap-tags/ap-tag=aa%3Abb%3Acc%3Add%3Aee%3Aff": `{
			"Cisco-IOS-XE-wireless-ap-cfg:ap-tag": {
				"ap-mac": "aa:bb:cc:dd:ee:ff",
				"site-tag": "building1",
				"policy-tag": "default-policy",
				"rf-tag": "typical"
			}
		}`,
		"Cisco-IOS-XE-wireless-ap-cfg:ap-cfg-data/tag-source-priority-configs/tag-source-priority-config=1": `{
			"Cisco-IOS-XE-wireless-ap-cfg:tag-source-priority-config": {
				"priority": 1,
				"source": "filter"
			}
		}`,
		"Cisco-IOS-XE-wireless-ap-global-oper:ap-global-oper-data?content=config&fields=ap-history(ethernet-mac;ip-addr)&ap-history=aa%3Abb%3Acc%3Add%3Aee%3Aff": `{
			"Cisco-IOS-XE-wireless-ap-global-oper:ap-history": [{
				"ethernet-mac": "aa:bb:cc:dd:ee:ff",
				"ip-addr": "192.168.1.100"
			}]
		}`,
		"Cisco-IOS-XE-wireless-ap-global-oper:ap-global-oper-data/ap-join-stats=aa%3Abb%3Acc%3Add%3Aee%3Aff": `{
			"Cisco-IOS-XE-wireless-ap-global-oper:ap-join-stats": [{
				"wtp-mac": "aa:bb:cc:dd:ee:ff",
				"join-time": "2023-01-01T00:00:00Z"
			}]
		}`,
		"Cisco-IOS-XE-wireless-ap-global-oper:ap-global-oper-data?content=config&fields=wlan-client-stats(wlan-id;client-count)&wlan-client-stats=1": `{
			"Cisco-IOS-XE-wireless-ap-global-oper:wlan-client-stats": [{
				"wlan-id": 1,
				"client-count": 5
			}]
		}`,
		"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/capwap-data?content=config&fields=capwap-data(wtp-mac;name;ip-addr)&capwap-data=aa%3Abb%3Acc%3Add%3Aee%3Aff": `{
			"Cisco-IOS-XE-wireless-access-point-oper:capwap-data": [{
				"wtp-mac": "aa:bb:cc:dd:ee:ff",
				"ip-addr": "192.168.1.100",
				"name": "AP-Test-01"
			}]
		}`,
		"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/ap-name-mac-map?content=config&fields=ap-name-mac-map(wtp-name;wtp-mac)&ap-name-mac-map=AP-Test-01": `{
			"Cisco-IOS-XE-wireless-access-point-oper:ap-name-mac-map": [{
				"wtp-name": "AP-Test-01",
				"wtp-mac": "aa:bb:cc:dd:ee:ff"
			}]
		}`,
		"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/radio-oper-data?content=config&fields=radio-oper-data(wtp-mac;radio-slot-id;oper-state)&radio-oper-data=aa%3Abb%3Acc%3Add%3Aee%3Aff%2C0": `{
			"Cisco-IOS-XE-wireless-access-point-oper:radio-oper-data": [{
				"wtp-mac": "aa:bb:cc:dd:ee:ff",
				"radio-slot-id": 0,
				"oper-state": "up"
			}]
		}`,
		"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/ap-radio-neighbor?content=config&fields=ap-radio-neighbor(wtp-mac;radio-slot-id;bssid)&ap-radio-neighbor=aa%3Abb%3Acc%3Add%3Aee%3Aff%2C0": `{
			"Cisco-IOS-XE-wireless-access-point-oper:ap-radio-neighbor": [{
				"wtp-mac": "aa:bb:cc:dd:ee:ff",
				"radio-slot-id": 0,
				"bssid": "bb:cc:dd:ee:ff:aa"
			}]
		}`,
		"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/ap-radio-neighbor?content=config&fields=ap-radio-neighbor(wtp-mac;radio-slot-id;bssid)&ap-radio-neighbor=aa%3Abb%3Acc%3Add%3Aee%3Aff%2C0%2Cbb%3Acc%3Add%3Aee%3Aff%3Aaa": `{
			"Cisco-IOS-XE-wireless-access-point-oper:ap-radio-neighbor": [{
				"wtp-mac": "aa:bb:cc:dd:ee:ff",
				"radio-slot-id": 0,
				"bssid": "bb:cc:dd:ee:ff:aa"
			}]
		}`,
		// AP tag query endpoints
		"Cisco-IOS-XE-wireless-ap-cfg:ap-cfg-data/ap-tags?ap-tag=28%3Aac%3A9e%3A11%3A48%3A10": `{
			"Cisco-IOS-XE-wireless-ap-cfg:ap-tags": {
				"ap-tag": [{
					"ap-mac": "28:ac:9e:11:48:10",
					"policy-tag": "labo-wlan-flex",
					"site-tag": "labo-site-flex",
					"rf-tag": "labo-inside"
				}]
			}
		}`,
		"Cisco-IOS-XE-wireless-ap-cfg:ap-cfg-data/ap-tags/ap-tag=28:ac:9e:11:48:10": `{
			"Cisco-IOS-XE-wireless-ap-cfg:ap-tags": {
				"ap-tag": [{
					"ap-mac": "28:ac:9e:11:48:10",
					"policy-tag": "labo-wlan-flex",
					"site-tag": "labo-site-flex",
					"rf-tag": "labo-inside"
				}]
			}
		}`,
		"Cisco-IOS-XE-wireless-ap-global-oper:ap-global-oper-data/ap-history=28:ac:9e:11:48:10": `{
			"Cisco-IOS-XE-wireless-ap-global-oper:ap-history": [{
				"ethernet-mac": "28:ac:9e:11:48:10",
				"ap-name": "TEST-AP01",
				"wtp-mac": "aa:bb:cc:dd:ee:ff"
			}]
		}`,
		"Cisco-IOS-XE-wireless-ap-global-oper:ap-global-oper-data/ap-join-stats=aa:bb:cc:dd:ee:ff": `{
			"Cisco-IOS-XE-wireless-ap-global-oper:ap-join-stats": [{
				"wtp-mac": "aa:bb:cc:dd:ee:ff",
				"ap-join-info": {
					"ap-ip-addr": "192.168.255.11",
					"ap-ethernet-mac": "28:ac:9e:11:48:10",
					"ap-name": "TEST-AP01",
					"is-joined": true
				}
			}]
		}`,
		"Cisco-IOS-XE-wireless-ap-global-oper:ap-global-oper-data/wlan-client-stats=1": `{
			"Cisco-IOS-XE-wireless-ap-global-oper:wlan-client-stats": [{
				"wlan-id": 1,
				"client-count": 0
			}]
		}`,
		"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/capwap-data=aa:bb:cc:dd:ee:ff": `{
			"Cisco-IOS-XE-wireless-access-point-oper:capwap-data": [{
				"wtp-mac": "aa:bb:cc:dd:ee:ff",
				"ip-addr": "192.168.255.11",
				"name": "TEST-AP01"
			}]
		}`,
		"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/ap-name-mac-map=TEST-AP01": `{
			"Cisco-IOS-XE-wireless-access-point-oper:ap-name-mac-map": [{
				"wtp-name": "TEST-AP01",
				"wtp-mac": "aa:bb:cc:dd:ee:ff",
				"eth-mac": "28:ac:9e:11:48:10"
			}]
		}`,
		"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/radio-oper-data=aa:bb:cc:dd:ee:ff,0": `{
			"Cisco-IOS-XE-wireless-access-point-oper:radio-oper-data": [{
				"wtp-mac": "aa:bb:cc:dd:ee:ff",
				"radio-slot-id": 0,
				"slot-id": 0,
				"radio-type": "radio-80211bg",
				"admin-state": "enabled",
				"oper-state": "radio-up"
			}]
		}`,
		"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/ap-radio-neighbor=aa:bb:cc:dd:ee:ff,0": `{
			"Cisco-IOS-XE-wireless-access-point-oper:ap-radio-neighbor": [{
				"wtp-mac": "aa:bb:cc:dd:ee:ff",
				"radio-slot-id": 0,
				"bssid": "aa:bb:cc:dd:ee:ff"
			}]
		}`,
		"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/ap-radio-neighbor=aa:bb:cc:dd:ee:ff,0,aa:bb:cc:dd:ee:ff": `{
			"Cisco-IOS-XE-wireless-access-point-oper:ap-radio-neighbor": [{
				"wtp-mac": "aa:bb:cc:dd:ee:ff",
				"radio-slot-id": 0,
				"bssid": "aa:bb:cc:dd:ee:ff"
			}]
		}`,
		// PUT/POST endpoints for tag assignment
		"Cisco-IOS-XE-wireless-ap-cfg:ap-cfg-data/ap-tag=aa:bb:cc:dd:ee:ff": `{}`,
	}
	mockServer := testutil.NewMockServer(responses)
	defer mockServer.Close()

	testClient := testutil.NewTestClient(mockServer)
	service := ap.NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	// Test filtered configuration operations
	t.Run("GetTagConfigByMAC", func(t *testing.T) {
		result, err := service.GetTagConfigByMAC(ctx, "28:ac:9e:11:48:10")
		if err != nil {
			t.Errorf("Expected no error for GetTagConfigByMAC, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for GetTagConfigByMAC, got nil")
		}
	})

	t.Run("GetTagSourcePriorityConfigByPriority", func(t *testing.T) {
		result, err := service.GetTagSourcePriorityConfigByPriority(ctx, 1)
		if err != nil {
			t.Errorf("Expected no error for GetTagSourcePriorityConfigByPriority, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for GetTagSourcePriorityConfigByPriority, got nil")
		}
	})

	// Test filtered global operational operations
	t.Run("ListAPHistoryByEthernetMAC", func(t *testing.T) {
		result, err := service.ListAPHistoryByEthernetMAC(ctx, "28:ac:9e:11:48:10")
		if err != nil {
			t.Errorf("Expected no error for ListAPHistoryByEthernetMAC, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for ListAPHistoryByEthernetMAC, got nil")
		}
	})

	t.Run("GetAPJoinStatsByWTPMAC", func(t *testing.T) {
		result, err := service.GetAPJoinStatsByWTPMAC(ctx, "aa:bb:cc:dd:ee:ff")
		if err != nil {
			t.Errorf("Expected no error for GetAPJoinStatsByWTPMAC, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for GetAPJoinStatsByWTPMAC, got nil")
		}
	})

	t.Run("GetWLANClientStatsByWLANID", func(t *testing.T) {
		result, err := service.GetWLANClientStatsByWLANID(ctx, 1)
		if err != nil {
			t.Errorf("Expected no error for GetWLANClientStatsByWLANID, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for GetWLANClientStatsByWLANID, got nil")
		}
	})

	// Test filtered operational operations
	t.Run("GetCAPWAPDataByWTPMAC", func(t *testing.T) {
		result, err := service.GetCAPWAPDataByWTPMAC(ctx, "aa:bb:cc:dd:ee:ff")
		if err != nil {
			t.Errorf("Expected no error for GetCAPWAPDataByWTPMAC, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for GetCAPWAPDataByWTPMAC, got nil")
		}
	})

	t.Run("GetNameMACMapByWTPName", func(t *testing.T) {
		result, err := service.GetNameMACMapByWTPName(ctx, "TEST-AP01")
		if err != nil {
			t.Errorf("Expected no error for GetNameMACMapByWTPName, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for GetNameMACMapByWTPName, got nil")
		}
	})

	t.Run("GetRadioStatusByWTPMACAndSlot", func(t *testing.T) {
		result, err := service.GetRadioStatusByWTPMACAndSlot(ctx, "aa:bb:cc:dd:ee:ff", 0)
		if err != nil {
			t.Errorf("Expected no error for GetRadioStatusByWTPMACAndSlot, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for GetRadioStatusByWTPMACAndSlot, got nil")
		}
	})

	t.Run("GetRadioNeighborByAPMACSlotAndBSSID", func(t *testing.T) {
		result, err := service.GetRadioNeighborByAPMACSlotAndBSSID(ctx, "aa:bb:cc:dd:ee:ff", 0, "aa:bb:cc:dd:ee:ff")
		if err != nil {
			t.Errorf("Expected no error for GetRadioNeighborByAPMACSlotAndBSSID, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for GetRadioNeighborByAPMACSlotAndBSSID, got nil")
		}
	})
}

// TestApServiceUnit_ValidationErrors tests input validation scenarios.
func TestApServiceUnit_ValidationErrors(t *testing.T) {
	// Use minimal mock server since we're testing validation before network calls
	responses := map[string]string{}
	mockServer := testutil.NewMockServer(responses)
	defer mockServer.Close()

	testClient := testutil.NewTestClient(mockServer)
	service := ap.NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	// Test invalid MAC address validation
	t.Run("GetTagConfigByMAC_InvalidMAC", func(t *testing.T) {
		_, err := service.GetTagConfigByMAC(ctx, "invalid-mac")
		if err == nil {
			t.Error("Expected error for invalid MAC address, got nil")
		}
	})

	// Test empty parameter validation
	t.Run("ListAPHistoryByEthernetMAC_EmptyMAC", func(t *testing.T) {
		_, err := service.ListAPHistoryByEthernetMAC(ctx, "")
		if err == nil {
			t.Error("Expected error for empty MAC address, got nil")
		}
	})

	t.Run("ListAPHistoryByEthernetMAC_WhitespaceMAC", func(t *testing.T) {
		_, err := service.ListAPHistoryByEthernetMAC(ctx, "   ")
		if err == nil {
			t.Error("Expected error for whitespace MAC address, got nil")
		}
	})

	t.Run("GetAPJoinStatsByWTPMAC_EmptyMAC", func(t *testing.T) {
		_, err := service.GetAPJoinStatsByWTPMAC(ctx, "")
		if err == nil {
			t.Error("Expected error for empty MAC address, got nil")
		}
	})

	t.Run("GetWLANClientStatsByWLANID_InvalidWLANID", func(t *testing.T) {
		_, err := service.GetWLANClientStatsByWLANID(ctx, 0)
		if err == nil {
			t.Error("Expected error for invalid WLAN ID, got nil")
		}
		_, err = service.GetWLANClientStatsByWLANID(ctx, -1)
		if err == nil {
			t.Error("Expected error for negative WLAN ID, got nil")
		}
	})

	t.Run("GetCAPWAPDataByWTPMAC_EmptyMAC", func(t *testing.T) {
		_, err := service.GetCAPWAPDataByWTPMAC(ctx, "")
		if err == nil {
			t.Error("Expected error for empty MAC address, got nil")
		}
	})

	t.Run("GetNameMACMapByWTPName_EmptyName", func(t *testing.T) {
		_, err := service.GetNameMACMapByWTPName(ctx, "")
		if err == nil {
			t.Error("Expected error for empty WTP name, got nil")
		}
	})

	t.Run("GetRadioStatusByWTPMACAndSlot_EmptyMAC", func(t *testing.T) {
		_, err := service.GetRadioStatusByWTPMACAndSlot(ctx, "", 0)
		if err == nil {
			t.Error("Expected error for empty MAC address, got nil")
		}
	})

	t.Run("GetRadioNeighborByAPMACSlotAndBSSID_EmptyMAC", func(t *testing.T) {
		_, err := service.GetRadioNeighborByAPMACSlotAndBSSID(ctx, "", 0, "bb:cc:dd:ee:ff:aa")
		if err == nil {
			t.Error("Expected error for empty MAC address, got nil")
		}
	})

	t.Run("GetRadioNeighborByAPMACSlotAndBSSID_EmptyBSSID", func(t *testing.T) {
		_, err := service.GetRadioNeighborByAPMACSlotAndBSSID(ctx, "aa:bb:cc:dd:ee:ff", 0, "")
		if err == nil {
			t.Error("Expected error for empty BSSID, got nil")
		}
	})

	t.Run("GetRadioNeighborByAPMACSlotAndBSSID_InvalidMAC", func(t *testing.T) {
		_, err := service.GetRadioNeighborByAPMACSlotAndBSSID(ctx, "invalid-mac", 0, "bb:cc:dd:ee:ff:aa")
		if err == nil {
			t.Error("Expected error for invalid MAC address, got nil")
		}
	})
}

// TestApServiceUnit_SetOperations_MockSuccess tests RPC and state change operations.
func TestApServiceUnit_SetOperations_MockSuccess(t *testing.T) {
	// Create mock server for RPC operations
	responses := map[string]string{
		"Cisco-IOS-XE-wireless-access-point-cfg-rpc:set-ap-admin-state":      `{"status": "success"}`,
		"Cisco-IOS-XE-wireless-access-point-cfg-rpc:set-ap-slot-admin-state": `{"status": "success"}`,
		"Cisco-IOS-XE-wireless-access-point-cmd-rpc:ap-reset":                `{"status": "success"}`,
		"Cisco-IOS-XE-wireless-ap-cfg:ap-cfg-data/ap-tag=aa:bb:cc:dd:ee:ff":  `{"status": "success"}`,
		"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/capwap-data": `{
			"Cisco-IOS-XE-wireless-access-point-oper:capwap-data": [{
				"wtp-mac": "aa:bb:cc:dd:ee:ff",
				"name": "TEST-AP01"
			}]
		}`,
	}
	mockServer := testutil.NewMockServer(responses)
	defer mockServer.Close()

	testClient := testutil.NewTestClient(mockServer)
	service := ap.NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	// Test AP admin state operations
	t.Run("EnableAP", func(t *testing.T) {
		err := service.EnableAP(ctx, "aa:bb:cc:dd:ee:ff")
		if err != nil {
			t.Errorf("Expected no error for EnableAP, got: %v", err)
		}
	})

	t.Run("DisableAP", func(t *testing.T) {
		err := service.DisableAP(ctx, "aa:bb:cc:dd:ee:ff")
		if err != nil {
			t.Errorf("Expected no error for DisableAP, got: %v", err)
		}
	})

	// Test radio state operations (using core.RadioBand constants)
	t.Run("EnableRadio", func(t *testing.T) {
		err := service.EnableRadio(ctx, "aa:bb:cc:dd:ee:ff", core.RadioBand24GHz)
		if err != nil {
			t.Errorf("Expected no error for EnableRadio, got: %v", err)
		}
	})

	t.Run("DisableRadio", func(t *testing.T) {
		err := service.DisableRadio(ctx, "aa:bb:cc:dd:ee:ff", core.RadioBand5GHz)
		if err != nil {
			t.Errorf("Expected no error for DisableRadio, got: %v", err)
		}
	})

	// Test tag assignment operations
	t.Run("AssignSiteTag", func(t *testing.T) {
		err := service.AssignSiteTag(ctx, "aa:bb:cc:dd:ee:ff", "labo-site-flex")
		if err != nil {
			t.Errorf("Expected no error for AssignSiteTag, got: %v", err)
		}
	})

	t.Run("AssignPolicyTag", func(t *testing.T) {
		err := service.AssignPolicyTag(ctx, "aa:bb:cc:dd:ee:ff", "labo-wlan-flex")
		if err != nil {
			t.Errorf("Expected no error for AssignPolicyTag, got: %v", err)
		}
	})

	t.Run("AssignRFTag", func(t *testing.T) {
		err := service.AssignRFTag(ctx, "aa:bb:cc:dd:ee:ff", "labo-inside")
		if err != nil {
			t.Errorf("Expected no error for AssignRFTag, got: %v", err)
		}
	})

	// Test AP reload operation
	t.Run("Reload", func(t *testing.T) {
		err := service.Reload(ctx, "aa:bb:cc:dd:ee:ff")
		if err != nil {
			t.Errorf("Expected no error for Reload, got: %v", err)
		}
	})
}

// TestApServiceUnit_SetOperations_ValidationErrors tests validation for state change operations.
func TestApServiceUnit_SetOperations_ValidationErrors(t *testing.T) {
	responses := map[string]string{}
	mockServer := testutil.NewMockServer(responses)
	defer mockServer.Close()

	testClient := testutil.NewTestClient(mockServer)
	service := ap.NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	// Test invalid MAC validation for AP state operations
	t.Run("EnableAP_InvalidMAC", func(t *testing.T) {
		err := service.EnableAP(ctx, "invalid-mac")
		if err == nil {
			t.Error("Expected error for invalid MAC address, got nil")
		}
	})

	t.Run("DisableAP_InvalidMAC", func(t *testing.T) {
		err := service.DisableAP(ctx, "invalid-mac")
		if err == nil {
			t.Error("Expected error for invalid MAC address, got nil")
		}
	})

	// Test invalid MAC validation for radio operations
	t.Run("EnableRadio_InvalidMAC", func(t *testing.T) {
		err := service.EnableRadio(ctx, "invalid-mac", core.RadioBand24GHz)
		if err == nil {
			t.Error("Expected error for invalid MAC address, got nil")
		}
	})

	t.Run("DisableRadio_InvalidMAC", func(t *testing.T) {
		err := service.DisableRadio(ctx, "invalid-mac", core.RadioBand5GHz)
		if err == nil {
			t.Error("Expected error for invalid MAC address, got nil")
		}
	})

	// Test empty tag validation
	t.Run("AssignSiteTag_EmptyTag", func(t *testing.T) {
		err := service.AssignSiteTag(ctx, "aa:bb:cc:dd:ee:ff", "")
		if err == nil {
			t.Error("Expected error for empty site tag, got nil")
		}
	})

	t.Run("AssignPolicyTag_EmptyTag", func(t *testing.T) {
		err := service.AssignPolicyTag(ctx, "aa:bb:cc:dd:ee:ff", "")
		if err == nil {
			t.Error("Expected error for empty policy tag, got nil")
		}
	})

	t.Run("AssignRFTag_EmptyTag", func(t *testing.T) {
		err := service.AssignRFTag(ctx, "aa:bb:cc:dd:ee:ff", "")
		if err == nil {
			t.Error("Expected error for empty RF tag, got nil")
		}
	})

	// Test invalid MAC validation for tag assignment
	t.Run("AssignSiteTag_InvalidMAC", func(t *testing.T) {
		err := service.AssignSiteTag(ctx, "invalid-mac", "building1")
		if err == nil {
			t.Error("Expected error for invalid MAC address, got nil")
		}
	})

	// Test reload with invalid MAC
	t.Run("Reload_InvalidMAC", func(t *testing.T) {
		err := service.Reload(ctx, "invalid-mac")
		if err == nil {
			t.Error("Expected error for invalid MAC address, got nil")
		}
	})
}

// TestApServiceUnit_EdgeCases_MockSuccess tests edge cases and error branches.
func TestApServiceUnit_EdgeCases_MockSuccess(t *testing.T) {
	// Create mock server with specific responses for edge cases
	responses := map[string]string{
		"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/capwap-data": `{
			"Cisco-IOS-XE-wireless-access-point-oper:capwap-data": []
		}`,
		"Cisco-IOS-XE-wireless-access-point-cmd-rpc:ap-reset": `{"status": "success"}`,
	}
	mockServer := testutil.NewMockServer(responses)
	defer mockServer.Close()

	testClient := testutil.NewTestClient(mockServer)
	service := ap.NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	// Test reload with empty CAPWAP data
	t.Run("Reload_EmptyCAPWAPData", func(t *testing.T) {
		err := service.Reload(ctx, "aa:bb:cc:dd:ee:ff")
		if err == nil {
			t.Error("Expected error for AP not found in CAPWAP data, got nil")
		}
	})
}

// TestApServiceUnit_NilCAPWAPData tests nil CAPWAP data handling.
func TestApServiceUnit_NilCAPWAPData(t *testing.T) {
	mockServer := testutil.NewMockErrorServer(
		[]string{"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/capwap-data"},
		500,
	)
	defer mockServer.Close()

	testClient := testutil.NewTestClient(mockServer)
	service := ap.NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	// Test reload with failed CAPWAP data retrieval
	t.Run("Reload_FailedCAPWAPDataRetrieval", func(t *testing.T) {
		err := service.Reload(ctx, "aa:bb:cc:dd:ee:ff")
		if err == nil {
			t.Error("Expected error for failed CAPWAP data retrieval, got nil")
		}
	})
}

// TestApServiceUnit_AdditionalErrorCases tests additional error handling scenarios for 100% coverage.
func TestApServiceUnit_AdditionalErrorCases(t *testing.T) {
	// Mock server with specific error responses for edge cases
	errorConfig := map[string]testutil.ErrorConfig{
		"Cisco-IOS-XE-wireless-access-point-cfg-rpc:set-ap-admin-state": {
			StatusCode:   400,
			ErrorMessage: "Invalid request",
		},
		"Cisco-IOS-XE-wireless-access-point-cfg-rpc:set-ap-slot-admin-state": {
			StatusCode:   400,
			ErrorMessage: "Invalid request",
		},
		"Cisco-IOS-XE-wireless-ap-cfg:ap-cfg-data/ap-tag=aa:bb:cc:dd:ee:ff": {
			StatusCode:   400,
			ErrorMessage: "Invalid request",
		},
	}

	mockServer := testutil.NewMockServerWithCustomErrors(t, errorConfig)
	defer mockServer.Close()

	testClient := testutil.NewTestClient(mockServer)
	service := ap.NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	// Test updateAPState error handling
	t.Run("UpdateAPState_RPCError", func(t *testing.T) {
		err := service.EnableAP(ctx, "aa:bb:cc:dd:ee:ff")
		if err == nil {
			t.Error("Expected error for failed RPC call, got nil")
		}
	})

	// Test updateRadioState error handling
	t.Run("UpdateRadioState_RPCError", func(t *testing.T) {
		err := service.EnableRadio(ctx, "aa:bb:cc:dd:ee:ff", core.RadioBand24GHz)
		if err == nil {
			t.Error("Expected error for failed radio RPC call, got nil")
		}
	})

	// Test updateRadioState with invalid radio band
	t.Run("UpdateRadioState_InvalidRadioBand", func(t *testing.T) {
		err := service.EnableRadio(ctx, "aa:bb:cc:dd:ee:ff", core.RadioBand(999)) // Invalid band
		if err == nil {
			t.Error("Expected error for invalid radio band, got nil")
		}
	})

	// Test assignTags error handling
	t.Run("AssignTags_RPCError", func(t *testing.T) {
		err := service.AssignSiteTag(ctx, "aa:bb:cc:dd:ee:ff", "test-site")
		if err == nil {
			t.Error("Expected error for failed tag assignment RPC call, got nil")
		}
	})
}

// TestApServiceUnit_EdgeCaseValidation tests additional validation edge cases.
func TestApServiceUnit_EdgeCaseValidation(t *testing.T) {
	responses := map[string]string{}
	mockServer := testutil.NewMockServer(responses)
	defer mockServer.Close()

	testClient := testutil.NewTestClient(mockServer)
	service := ap.NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	// Test GetAPJoinStatsByWTPMAC with whitespace MAC
	t.Run("GetAPJoinStatsByWTPMAC_WhitespaceMAC", func(t *testing.T) {
		_, err := service.GetAPJoinStatsByWTPMAC(ctx, "   ")
		if err == nil {
			t.Error("Expected error for whitespace-only MAC, got nil")
		}
	})

	// Test GetNameMACMapByWTPName with whitespace name
	t.Run("GetNameMACMapByWTPName_WhitespaceName", func(t *testing.T) {
		_, err := service.GetNameMACMapByWTPName(ctx, "   ")
		if err == nil {
			t.Error("Expected error for whitespace-only name, got nil")
		}
	})

	// Test GetRadioStatusByWTPMACAndSlot with whitespace MAC
	t.Run("GetRadioStatusByWTPMACAndSlot_WhitespaceMAC", func(t *testing.T) {
		_, err := service.GetRadioStatusByWTPMACAndSlot(ctx, "   ", 0)
		if err == nil {
			t.Error("Expected error for whitespace-only MAC, got nil")
		}
	})
}

// TestApServiceUnit_ReloadEdgeCases tests specific edge cases for Reload function to achieve 100% coverage.
func TestApServiceUnit_ReloadEdgeCases(t *testing.T) {
	// Test Reload with nil CAPWAP data response
	t.Run("Reload_NilCAPWAPResponse", func(t *testing.T) {
		responses := map[string]string{
			"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/capwap-data": `null`,
		}
		mockServer := testutil.NewMockServer(responses)
		defer mockServer.Close()

		testClient := testutil.NewTestClient(mockServer)
		service := ap.NewService(testClient.Core().(*core.Client))
		ctx := testutil.TestContext(t)

		err := service.Reload(ctx, "aa:bb:cc:dd:ee:ff")
		if err == nil {
			t.Error("Expected error for nil CAPWAP response, got nil")
		}
	})

	// Test Reload with AP not found in CAPWAP data
	t.Run("Reload_APNotFoundInCAPWAP", func(t *testing.T) {
		responses := map[string]string{
			"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/capwap-data": `{
				"Cisco-IOS-XE-wireless-access-point-oper:capwap-data": [{
					"wtp-mac": "aa:bb:cc:dd:ee:ff",
					"name": "Different-AP"
				}]
			}`,
		}
		mockServer := testutil.NewMockServer(responses)
		defer mockServer.Close()

		testClient := testutil.NewTestClient(mockServer)
		service := ap.NewService(testClient.Core().(*core.Client))
		ctx := testutil.TestContext(t)

		err := service.Reload(ctx, "aa:bb:cc:dd:ee:ff")
		if err == nil {
			t.Error("Expected error for AP not found in CAPWAP data, got nil")
		}
	})
}
