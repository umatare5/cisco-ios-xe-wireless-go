package client_test

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/pkg/testutil"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/client"
)

// TestClientServiceUnit_Constructor_Success tests service constructor functionality.
func TestClientServiceUnit_Constructor_Success(t *testing.T) {
	t.Run("NewServiceWithValidClient", func(t *testing.T) {
		// Create mock server and test client using public API
		responses := map[string]string{
			"test-endpoint": `{"status": "success"}`,
		}
		mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(responses))
		defer mockServer.Close()

		testClient := testutil.NewTestClient(mockServer)
		service := client.NewService(testClient.Core().(*core.Client))

		// Test that service can be created and has a client
		if service.Client() == nil {
			t.Error("Expected service to have a client")
		}
	})

	t.Run("NewServiceWithNilClient", func(t *testing.T) {
		service := client.NewService(nil)

		// Service should still be created even with nil client
		if service.Client() != nil {
			t.Error("Expected service with nil client to return nil from Client()")
		}
	})
}

// TestClientServiceUnit_GetOperations_MockSuccess tests Get operations using mock server with live WNC data.
func TestClientServiceUnit_GetOperations_MockSuccess(t *testing.T) {
	// Create mock RESTCONF server with Client endpoints using real WNC data structure
	responses := map[string]string{
		"Cisco-IOS-XE-wireless-client-oper:client-oper-data": `{
			"Cisco-IOS-XE-wireless-client-oper:client-oper-data": {
				"common-oper-data": [{
					"client-mac": "02:40:f1:f7:f7:87",
					"ap-name": "TEST-AP01",
					"ms-ap-slot-id": 0,
					"ms-radio-type": "client-dot11ax-24ghz-prot",
					"wlan-id": 1,
					"client-type": "dot11-client-normal",
					"co-state": "client-status-run",
					"aaa-override-passphrase": false,
					"is-tvi-enabled": false,
					"wlan-policy": {
						"current-switching-mode": "local",
						"wlan-switching-mode": "local",
						"central-authentication": "client-authentication-type-local",
						"central-dhcp": false,
						"central-assoc-enable": false,
						"vlan-central-switching": false,
						"is-fabric-client": false,
						"is-guest-fabric-client": false,
						"upn-bit-flag": ""
					},
					"username": "",
					"guest-lan-client-info": {
						"wired-vlan": 0,
						"phy-ifid": 0,
						"idle-time-seconds": 0
					},
					"method-id": "no-method-id",
					"l3-vlan-override-received": false,
					"upn-id": 0,
					"is-locally-administered-mac": true,
					"idle-timeout": 0,
					"idle-timestamp": "1970-01-01T00:00:00+00:00",
					"client-duid": "",
					"vrf-name": ""
				}]
			}
		}`,
		"Cisco-IOS-XE-wireless-client-oper:client-oper-data/common-oper-data": `{
			"Cisco-IOS-XE-wireless-client-oper:common-oper-data": [{
				"client-mac": "02:40:f1:f7:f7:87",
				"ap-name": "TEST-AP01",
				"ms-ap-slot-id": 0,
				"ms-radio-type": "client-dot11ax-24ghz-prot",
				"wlan-id": 1,
				"client-type": "dot11-client-normal",
				"co-state": "client-status-run",
				"aaa-override-passphrase": false,
				"is-tvi-enabled": false,
				"wlan-policy": {
					"current-switching-mode": "local",
					"wlan-switching-mode": "local",
					"central-authentication": "client-authentication-type-local",
					"central-dhcp": false,
					"central-assoc-enable": false,
					"vlan-central-switching": false,
					"is-fabric-client": false,
					"is-guest-fabric-client": false,
					"upn-bit-flag": ""
				},
				"username": "",
				"guest-lan-client-info": {
					"wired-vlan": 0,
					"phy-ifid": 0,
					"idle-time-seconds": 0
				},
				"method-id": "no-method-id",
				"l3-vlan-override-received": false,
				"upn-id": 0,
				"is-locally-administered-mac": true,
				"idle-timeout": 0,
				"idle-timestamp": "1970-01-01T00:00:00+00:00",
				"client-duid": "",
				"vrf-name": ""
			}]
		}`,
		"Cisco-IOS-XE-wireless-client-oper:client-oper-data/common-oper-data=02:40:f1:f7:f7:87": `{
			"Cisco-IOS-XE-wireless-client-oper:common-oper-data": [{
				"client-mac": "02:40:f1:f7:f7:87",
				"ap-name": "TEST-AP01",
				"ms-ap-slot-id": 0,
				"ms-radio-type": "client-dot11ax-24ghz-prot",
				"wlan-id": 1,
				"client-type": "dot11-client-normal",
				"co-state": "client-status-run"
			}]
		}`,
		"Cisco-IOS-XE-wireless-client-oper:client-oper-data/dc-info": `{
			"Cisco-IOS-XE-wireless-client-oper:dc-info": []
		}`,
		"Cisco-IOS-XE-wireless-client-oper:client-oper-data/dot11-oper-data": `{
			"Cisco-IOS-XE-wireless-client-oper:dot11-oper-data": [{
				"ms-mac-address": "02:40:f1:f7:f7:87",
				"dot11-state": "associated",
				"ms-bssid": "f0:d8:05:2c:41:21",
				"ap-mac-address": "c4:14:a2:c9:02:70",
				"current-channel": 11,
				"ms-wlan-id": 1,
				"vap-ssid": "labo-wlan",
				"policy-profile": "labo-wlan-profile",
				"ms-ap-slot-id": 0,
				"radio-type": "dot11-radio-type-bg",
				"ms-association-id": 8,
				"ms-auth-alg-num": "open-system",
				"ms-reason-code": "reason-none",
				"ms-assoc-time": "2025-09-17T10:50:37.41636+00:00",
				"is-11g-client": true,
				"ms-supported-rates-str": "54.0",
				"ms-wifi": {
					"wpa-version": "wpa2",
					"cipher-suite": "ccmp-aes",
					"auth-key-mgmt": "psk",
					"group-mgmt-cipher-suite": "rsn-cipher-suite-use-group",
					"group-cipher-suite": "rsn-cipher-suite-use-group",
					"pwe-mode": "sae-pwe-mode-none"
				},
				"ms-wme-enabled": true,
				"dot11w-enabled": false,
				"ewlc-ms-phy-type": "client-dot11ax-24ghz-prot"
			}]
		}`,
		"Cisco-IOS-XE-wireless-client-oper:client-oper-data/mm-if-client-history": `{
			"Cisco-IOS-XE-wireless-client-oper:mm-if-client-history": []
		}`,
		"Cisco-IOS-XE-wireless-client-oper:client-oper-data/mm-if-client-stats": `{
			"Cisco-IOS-XE-wireless-client-oper:mm-if-client-stats": []
		}`,
		"Cisco-IOS-XE-wireless-client-oper:client-oper-data/mobility-oper-data": `{
			"Cisco-IOS-XE-wireless-client-oper:mobility-oper-data": [{
				"ms-mac-addr": "02:40:f1:f7:f7:87",
				"mm-client-role": "mm-client-role-local",
				"mm-client-roam-type": "mm-roam-type-none",
				"mm-instance": 0,
				"mm-complete-timestamp": "2025-09-17T10:50:37+00:00",
				"mm-remote-tunnel-ip": "0.0.0.0",
				"mm-remote-tunnel-sec-ip": "0.0.0.0",
				"mm-remote-platform-id": 0,
				"mm-remote-tunnel-id": 0,
				"mm-anchor-ip": "0.0.0.0"
			}]
		}`,
		"Cisco-IOS-XE-wireless-client-oper:client-oper-data/policy-data": `{
			"Cisco-IOS-XE-wireless-client-oper:policy-data": [{
				"mac": "02:40:f1:f7:f7:87",
				"res-vlan-id": 800,
				"res-vlan-name": "LAB-INTERNAL"
			}]
		}`,
		"Cisco-IOS-XE-wireless-client-oper:client-oper-data/sisf-db-mac": `{
			"Cisco-IOS-XE-wireless-client-oper:sisf-db-mac": [{
				"mac-addr": "02:40:f1:f7:f7:87",
				"ipv4-binding": {
					"ip-key": {
						"zone-id": 0,
						"ip-addr": "192.168.0.37"
					}
				},
				"ipv6-binding": [{
					"ip-key": {
						"zone-id": 2147484448,
						"ip-addr": "fe80::40:f1ff:fef7:f787"
					}
				}]
			}]
		}`,
		"Cisco-IOS-XE-wireless-client-oper:client-oper-data/traffic-stats": `{
			"Cisco-IOS-XE-wireless-client-oper:traffic-stats": [{
				"ms-mac-address": "02:40:f1:f7:f7:87",
				"bytes-rx": "37085614",
				"bytes-tx": "291727367",
				"policy-errs": "0",
				"pkts-rx": "160344",
				"pkts-tx": "260841",
				"data-retries": "33530",
				"rts-retries": "0",
				"duplicate-rcv": "0",
				"decrypt-failed": "0",
				"mic-mismatch": "0",
				"mic-missing": "0",
				"most-recent-rssi": -42,
				"most-recent-snr": 57,
				"tx-excessive-retries": "0",
				"tx-retries": "0",
				"power-save-state": 1,
				"current-rate": "m11 ss2",
				"speed": 287,
				"spatial-stream": 2,
				"client-active": true,
				"glan-stats-update-timestamp": "1970-01-01T00:00:00+00:00",
				"glan-idle-update-timestamp": "1970-01-01T00:00:00+00:00",
				"rx-group-counter": "0",
				"tx-total-drops": "577"
			}]
		}`,
		// Add MAC query responses for all *ByMAC functions with real WNC data
		"Cisco-IOS-XE-wireless-client-oper:client-oper-data/dc-info=02:40:f1:f7:f7:87": `{
			"Cisco-IOS-XE-wireless-client-oper:dc-info": []
		}`,
		"Cisco-IOS-XE-wireless-client-oper:client-oper-data/dot11-oper-data=02:40:f1:f7:f7:87": `{
			"Cisco-IOS-XE-wireless-client-oper:dot11-oper-data": [{
				"ms-mac-address": "02:40:f1:f7:f7:87",
				"dot11-state": "associated",
				"current-channel": 11,
				"ms-wlan-id": 1,
				"vap-ssid": "labo-wlan",
				"policy-profile": "labo-wlan-profile"
			}]
		}`,
		"Cisco-IOS-XE-wireless-client-oper:client-oper-data/mm-if-client-history=02:40:f1:f7:f7:87": `{
			"Cisco-IOS-XE-wireless-client-oper:mm-if-client-history": []
		}`,
		"Cisco-IOS-XE-wireless-client-oper:client-oper-data/mm-if-client-stats=02:40:f1:f7:f7:87": `{
			"Cisco-IOS-XE-wireless-client-oper:mm-if-client-stats": []
		}`,
		"Cisco-IOS-XE-wireless-client-oper:client-oper-data/mobility-oper-data=02:40:f1:f7:f7:87": `{
			"Cisco-IOS-XE-wireless-client-oper:mobility-oper-data": [{
				"ms-mac-addr": "02:40:f1:f7:f7:87",
				"mm-client-role": "mm-client-role-local",
				"mm-client-roam-type": "mm-roam-type-none"
			}]
		}`,
		"Cisco-IOS-XE-wireless-client-oper:client-oper-data/policy-data=02:40:f1:f7:f7:87": `{
			"Cisco-IOS-XE-wireless-client-oper:policy-data": [{
				"mac": "02:40:f1:f7:f7:87",
				"res-vlan-id": 800,
				"res-vlan-name": "LAB-INTERNAL"
			}]
		}`,
		"Cisco-IOS-XE-wireless-client-oper:client-oper-data/sisf-db-mac=02:40:f1:f7:f7:87": `{
			"Cisco-IOS-XE-wireless-client-oper:sisf-db-mac": [{
				"mac-addr": "02:40:f1:f7:f7:87",
				"ipv4-binding": {
					"ip-key": {
						"zone-id": 0,
						"ip-addr": "192.168.0.37"
					}
				}
			}]
		}`,
		"Cisco-IOS-XE-wireless-client-oper:client-oper-data/traffic-stats=02:40:f1:f7:f7:87": `{
			"Cisco-IOS-XE-wireless-client-oper:traffic-stats": [{
				"ms-mac-address": "02:40:f1:f7:f7:87",
				"bytes-rx": "37085614",
				"bytes-tx": "291727367",
				"most-recent-rssi": -42,
				"most-recent-snr": 57,
				"current-rate": "m11 ss2",
				"speed": 287,
				"spatial-stream": 2,
				"client-active": true
			}]
		}`,
	}
	mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(responses))
	defer mockServer.Close()

	// Create test client configured for the mock server
	testClient := testutil.NewTestClient(mockServer)
	service := client.NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	// Test basic operations
	result, err := service.GetOperational(ctx)
	if err != nil {
		t.Errorf("GetOperational failed: %v", err)
	}
	if result == nil {
		t.Error("Expected non-nil result from GetOperational")
	}

	// Test ListCommonInfo
	commonResult, err := service.ListCommonInfo(ctx)
	if err != nil {
		t.Errorf("ListCommonInfo failed: %v", err)
	}
	if commonResult == nil {
		t.Error("Expected non-nil result from ListCommonInfo")
	}

	// Test GetCommonInfoByMAC with real WNC MAC address
	commonByMAC, err := service.GetCommonInfoByMAC(ctx, "02:40:f1:f7:f7:87")
	if err != nil {
		t.Errorf("GetCommonInfoByMAC failed: %v", err)
	}
	if commonByMAC == nil {
		t.Error("Expected non-nil result from GetCommonInfoByMAC")
	}

	// Test remaining List functions
	dcResult, err := service.ListDCInfo(ctx)
	if err != nil {
		t.Errorf("ListDCInfo failed: %v", err)
	}
	if dcResult == nil {
		t.Error("Expected non-nil result from ListDCInfo")
	}

	dot11Result, err := service.ListDot11Info(ctx)
	if err != nil {
		t.Errorf("ListDot11Info failed: %v", err)
	}
	if dot11Result == nil {
		t.Error("Expected non-nil result from ListDot11Info")
	}

	mmifHistoryResult, err := service.ListMMIFClientHistory(ctx)
	if err != nil {
		t.Errorf("ListMMIFClientHistory failed: %v", err)
	}
	if mmifHistoryResult == nil {
		t.Error("Expected non-nil result from ListMMIFClientHistory")
	}

	mmifStatsResult, err := service.ListMMIFClientStats(ctx)
	if err != nil {
		t.Errorf("ListMMIFClientStats failed: %v", err)
	}
	if mmifStatsResult == nil {
		t.Error("Expected non-nil result from ListMMIFClientStats")
	}

	mobilityResult, err := service.ListMobilityInfo(ctx)
	if err != nil {
		t.Errorf("ListMobilityInfo failed: %v", err)
	}
	if mobilityResult == nil {
		t.Error("Expected non-nil result from ListMobilityInfo")
	}

	policyResult, err := service.ListPolicyInfo(ctx)
	if err != nil {
		t.Errorf("ListPolicyInfo failed: %v", err)
	}
	if policyResult == nil {
		t.Error("Expected non-nil result from ListPolicyInfo")
	}

	sisfResult, err := service.ListSISFDB(ctx)
	if err != nil {
		t.Errorf("ListSISFDB failed: %v", err)
	}
	if sisfResult == nil {
		t.Error("Expected non-nil result from ListSISFDB")
	}

	trafficResult, err := service.ListTrafficStats(ctx)
	if err != nil {
		t.Errorf("ListTrafficStats failed: %v", err)
	}
	if trafficResult == nil {
		t.Error("Expected non-nil result from ListTrafficStats")
	}

	// Test all *ByMAC functions with real WNC MAC address
	dcByMAC, err := service.GetDCInfoByMAC(ctx, "02:40:f1:f7:f7:87")
	if err != nil {
		t.Errorf("GetDCInfoByMAC failed: %v", err)
	}
	if dcByMAC == nil {
		t.Error("Expected non-nil result from GetDCInfoByMAC")
	}

	dot11ByMAC, err := service.GetDot11InfoByMAC(ctx, "02:40:f1:f7:f7:87")
	if err != nil {
		t.Errorf("GetDot11InfoByMAC failed: %v", err)
	}
	if dot11ByMAC == nil {
		t.Error("Expected non-nil result from GetDot11InfoByMAC")
	}

	mmifHistoryByMAC, err := service.GetMMIFClientHistoryByMAC(ctx, "02:40:f1:f7:f7:87")
	if err != nil {
		t.Errorf("GetMMIFClientHistoryByMAC failed: %v", err)
	}
	if mmifHistoryByMAC == nil {
		t.Error("Expected non-nil result from GetMMIFClientHistoryByMAC")
	}

	mmifStatsByMAC, err := service.GetMMIFClientStatsByMAC(ctx, "02:40:f1:f7:f7:87")
	if err != nil {
		t.Errorf("GetMMIFClientStatsByMAC failed: %v", err)
	}
	if mmifStatsByMAC == nil {
		t.Error("Expected non-nil result from GetMMIFClientStatsByMAC")
	}

	mobilityByMAC, err := service.GetMobilityInfoByMAC(ctx, "02:40:f1:f7:f7:87")
	if err != nil {
		t.Errorf("GetMobilityInfoByMAC failed: %v", err)
	}
	if mobilityByMAC == nil {
		t.Error("Expected non-nil result from GetMobilityInfoByMAC")
	}

	policyByMAC, err := service.GetPolicyInfoByMAC(ctx, "02:40:f1:f7:f7:87")
	if err != nil {
		t.Errorf("GetPolicyInfoByMAC failed: %v", err)
	}
	if policyByMAC == nil {
		t.Error("Expected non-nil result from GetPolicyInfoByMAC")
	}

	sisfByMAC, err := service.GetSISFDBByMAC(ctx, "02:40:f1:f7:f7:87")
	if err != nil {
		t.Errorf("GetSISFDBByMAC failed: %v", err)
	}
	if sisfByMAC == nil {
		t.Error("Expected non-nil result from GetSISFDBByMAC")
	}

	trafficByMAC, err := service.GetTrafficStatsByMAC(ctx, "02:40:f1:f7:f7:87")
	if err != nil {
		t.Errorf("GetTrafficStatsByMAC failed: %v", err)
	}
	if trafficByMAC == nil {
		t.Error("Expected non-nil result from GetTrafficStatsByMAC")
	}
}

// TestClientServiceUnit_GetOperations_ErrorHandling tests error scenarios using mock server.
func TestClientServiceUnit_GetOperations_ErrorHandling(t *testing.T) {
	// Create mock server that returns 404 for Client endpoints
	errorPaths := []string{
		"Cisco-IOS-XE-wireless-client-oper:client-oper-data",
	}
	mockServer := testutil.NewMockServer(testutil.WithErrorResponses(errorPaths, 404))
	defer mockServer.Close()

	testClient := testutil.NewTestClient(mockServer)
	service := client.NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	// Test that GetOperational properly handles 404 errors
	_, err := service.GetOperational(ctx)
	if err == nil {
		t.Error("Expected error for 404 response, got nil")
	}
}

// TestClientServiceUnit_ValidationErrors_EmptyMAC tests validation error scenarios.
func TestClientServiceUnit_ValidationErrors_EmptyMAC(t *testing.T) {
	t.Parallel()

	server := testutil.NewMockServer(testutil.WithSuccessResponses(map[string]string{}))
	defer server.Close()
	testClient := testutil.NewTestClient(server)
	service := client.NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	// Test GetCommonInfoByMAC with empty MAC
	_, err := service.GetCommonInfoByMAC(ctx, "")
	if err == nil {
		t.Error("Expected error for empty MAC in GetCommonInfoByMAC")
	}

	// Test GetDCInfoByMAC with empty MAC
	_, err = service.GetDCInfoByMAC(ctx, "")
	if err == nil {
		t.Error("Expected error for empty MAC in GetDCInfoByMAC")
	}

	// Test GetDot11InfoByMAC with empty MAC
	_, err = service.GetDot11InfoByMAC(ctx, "")
	if err == nil {
		t.Error("Expected error for empty MAC in GetDot11InfoByMAC")
	}

	// Test GetMMIFClientHistoryByMAC with empty MAC
	_, err = service.GetMMIFClientHistoryByMAC(ctx, "")
	if err == nil {
		t.Error("Expected error for empty MAC in GetMMIFClientHistoryByMAC")
	}

	// Test GetMMIFClientStatsByMAC with empty MAC
	_, err = service.GetMMIFClientStatsByMAC(ctx, "")
	if err == nil {
		t.Error("Expected error for empty MAC in GetMMIFClientStatsByMAC")
	}

	// Test GetMobilityInfoByMAC with empty MAC
	_, err = service.GetMobilityInfoByMAC(ctx, "")
	if err == nil {
		t.Error("Expected error for empty MAC in GetMobilityInfoByMAC")
	}

	// Test GetPolicyInfoByMAC with empty MAC
	_, err = service.GetPolicyInfoByMAC(ctx, "")
	if err == nil {
		t.Error("Expected error for empty MAC in GetPolicyInfoByMAC")
	}

	// Test GetSISFDBByMAC with empty MAC
	_, err = service.GetSISFDBByMAC(ctx, "")
	if err == nil {
		t.Error("Expected error for empty MAC in GetSISFDBByMAC")
	}

	// Test GetTrafficStatsByMAC with empty MAC
	_, err = service.GetTrafficStatsByMAC(ctx, "")
	if err == nil {
		t.Error("Expected error for empty MAC in GetTrafficStatsByMAC")
	}
}

// TestClientServiceUnit_KnownIssueHandling_Success tests known issue error handling.
func TestClientServiceUnit_KnownIssueHandling_Success(t *testing.T) {
	t.Parallel()

	// Create mock server with normal responses
	responses := map[string]string{
		"Cisco-IOS-XE-wireless-client-oper:client-oper-data": `{
			"Cisco-IOS-XE-wireless-client-oper:client-oper-data": {
				"common-oper-data": [{"ms-mac": "aa:bb:cc:dd:ee:ff"}]
			}
		}`,
		"Cisco-IOS-XE-wireless-client-oper:client-oper-data/dot11-oper-data": `{
			"Cisco-IOS-XE-wireless-client-oper:dot11-oper-data": [{"ms-mac": "aa:bb:cc:dd:ee:ff"}]
		}`,
	}

	mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(responses))
	defer mockServer.Close()

	testClient := testutil.NewTestClient(mockServer)
	service := client.NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	// Test ListDot11Info normal operation
	result, err := service.ListDot11Info(ctx)
	if err != nil {
		t.Errorf("ListDot11Info returned unexpected error: %v", err)
	}
	if result == nil {
		t.Error("Expected result for ListDot11Info, got nil")
	}
}

// TestClientServiceUnit_KnownIssueHandling_Dot11Errors tests known Dot11 operational data issue handling.
func TestClientServiceUnit_KnownIssueHandling_Dot11Errors(t *testing.T) {
	t.Parallel()

	// Create mock server with specific error message for known Dot11 issues using pkg/testutil
	server := testutil.NewMockServer(
		testutil.WithTesting(t),
		testutil.WithCustomResponse("dot11-oper-data", testutil.ResponseConfig{
			StatusCode: 500,
			Body:       `{"ietf-restconf:errors": {"error": [{"error-message": "failed to retrieve table cursor"}]}}`,
		}),
	)
	defer server.Close()

	testClient := testutil.NewTestClient(server)
	service := client.NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	// Test ListDot11Info with known cursor issue - should return empty result
	result, err := service.ListDot11Info(ctx)
	if err != nil {
		t.Errorf("Expected ListDot11Info to handle known issue gracefully, got error: %v", err)
	}
	if result == nil {
		t.Error("Expected empty result for known issue, got nil")
	}

	// Test GetDot11InfoByMAC with known cursor issue - should return empty result
	resultByMAC, err := service.GetDot11InfoByMAC(ctx, "02:40:f1:f7:f7:87")
	if err != nil {
		t.Errorf("Expected GetDot11InfoByMAC to handle known issue gracefully, got error: %v", err)
	}
	if resultByMAC == nil {
		t.Error("Expected empty result for known issue, got nil")
	}
}

// TestClientServiceUnit_KnownIssueHandling_GetOperationalErrors tests GetOperational known issue handling.
func TestClientServiceUnit_KnownIssueHandling_GetOperationalErrors(t *testing.T) {
	t.Parallel()

	// Create mock server with specific error message for known GetOperational issues using pkg/testutil
	server := testutil.NewMockServer(
		testutil.WithTesting(t),
		testutil.WithCustomResponse("client-oper-data", testutil.ResponseConfig{
			StatusCode: 500,
			Body:       `{"ietf-restconf:errors": {"error": [{"error-message": "unexpected EOF"}]}}`,
		}),
	)
	defer server.Close()

	testClient := testutil.NewTestClient(server)
	service := client.NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	// Test GetOperational with known EOF issue - should return empty result
	result, err := service.GetOperational(ctx)
	if err != nil {
		t.Errorf("Expected GetOperational to handle known issue gracefully, got error: %v", err)
	}
	if result == nil {
		t.Error("Expected empty result for known issue, got nil")
	}
}

// TestClientServiceUnit_KnownIssueHandling_AdditionalDot11Errors tests additional Dot11 error scenarios.
func TestClientServiceUnit_KnownIssueHandling_AdditionalDot11Errors(t *testing.T) {
	t.Parallel()

	// Create mock server for additional known Dot11 error: "Process DBAL response failed" using pkg/testutil
	server := testutil.NewMockServer(
		testutil.WithTesting(t),
		testutil.WithCustomResponse("dot11-oper-data", testutil.ResponseConfig{
			StatusCode: 500,
			Body:       `{"ietf-restconf:errors": {"error": [{"error-message": "Process DBAL response failed"}]}}`,
		}),
	)
	defer server.Close()

	testClient := testutil.NewTestClient(server)
	service := client.NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	// Test ListDot11Info with DBAL error - should return empty result
	result, err := service.ListDot11Info(ctx)
	if err != nil {
		t.Errorf("Expected ListDot11Info to handle DBAL error gracefully, got error: %v", err)
	}
	if result == nil {
		t.Error("Expected empty result for DBAL error, got nil")
	}

	// Test GetDot11InfoByMAC with DBAL error - should return empty result
	resultByMAC, err := service.GetDot11InfoByMAC(ctx, "02:40:f1:f7:f7:87")
	if err != nil {
		t.Errorf("Expected GetDot11InfoByMAC to handle DBAL error gracefully, got error: %v", err)
	}
	if resultByMAC == nil {
		t.Error("Expected empty result for DBAL error, got nil")
	}
}

// TestClientServiceUnit_KnownIssueHandling_UnknownErrors tests error scenarios that are not known issues.
func TestClientServiceUnit_KnownIssueHandling_UnknownErrors(t *testing.T) {
	t.Parallel()

	// Create mock server with unknown error messages that should not be handled gracefully
	server := testutil.NewMockServer(
		testutil.WithTesting(t),
		testutil.WithCustomResponse("dot11-oper-data", testutil.ResponseConfig{
			StatusCode: 500,
			Body:       `{"ietf-restconf:errors": {"error": [{"error-message": "unknown database error"}]}}`,
		}),
	)
	defer server.Close()

	testClient := testutil.NewTestClient(server)
	service := client.NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	// Test ListDot11Info with unknown error - should return error (not gracefully handled)
	_, err := service.ListDot11Info(ctx)
	if err == nil {
		t.Error("Expected error for unknown database error, got nil")
	}

	// Test GetDot11InfoByMAC with unknown error - should return error (not gracefully handled)
	_, err = service.GetDot11InfoByMAC(ctx, "02:40:f1:f7:f7:87")
	if err == nil {
		t.Error("Expected error for unknown database error, got nil")
	}
}

// TestClientServiceUnit_KnownIssueHandling_GetOperationalUnknownErrors tests GetOperational with unknown errors.
func TestClientServiceUnit_KnownIssueHandling_GetOperationalUnknownErrors(t *testing.T) {
	t.Parallel()

	// Create mock server with unknown error message for GetOperational
	server := testutil.NewMockServer(
		testutil.WithTesting(t),
		testutil.WithCustomResponse("client-oper-data", testutil.ResponseConfig{
			StatusCode: 500,
			Body:       `{"ietf-restconf:errors": {"error": [{"error-message": "unknown system error"}]}}`,
		}),
	)
	defer server.Close()

	testClient := testutil.NewTestClient(server)
	service := client.NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	// Test GetOperational with unknown error - should return error (not gracefully handled)
	_, err := service.GetOperational(ctx)
	if err == nil {
		t.Error("Expected error for unknown system error, got nil")
	}
}
