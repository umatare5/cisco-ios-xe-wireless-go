package client_test

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/pkg/testutil"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/client"
)

// TestClientServiceUnit_RadioVocabularies_MockSuccess tests that the two client radio vocabularies
// decode into their own named types from one record.
//
// The inversion is the reason this test exists: ms-radio-type carries a PHY generation and
// radio-type carries a band, so a constant whose spelling drifts to the other domain would compile
// and read empty. The bare siblings are the control — they prove the record was decoded.
func TestClientServiceUnit_RadioVocabularies_MockSuccess(t *testing.T) {
	t.Parallel()

	body := `{
		"Cisco-IOS-XE-wireless-client-oper:client-oper-data": {
			"common-oper-data": [
				{
					"ap-name": "TEST-AP01",
					"ms-radio-type": "client-dot11be-6ghz-prot"
				}
			],
			"dot11-oper-data": [
				{
					"radio-type": "dot11-radio-type-6ghz",
					"ewlc-ms-phy-type": "client-dot11be-6ghz-prot",
					"multilink-info": [
						{
							"sta-mac": "aa:bb:cc:dd:ee:ff",
							"band": "dot11-6-ghz-band",
							"radio-type": "dot11-radio-type-6ghz"
						}
					]
				}
			]
		}
	}`

	mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(map[string]string{
		"Cisco-IOS-XE-wireless-client-oper:client-oper-data": body,
	}))
	defer mockServer.Close()

	testClient := testutil.NewTestClient(mockServer)
	service := client.NewService(testClient.Core().(*core.Client))

	result, err := service.GetOperational(testutil.TestContext(t))
	if err != nil {
		t.Fatalf("GetOperational returned unexpected error: %v", err)
	}

	data := result.CiscoIOSXEWirelessClientOperData
	if data == nil || len(data.CommonOperData) != 1 || len(data.Dot11OperData) != 1 {
		t.Fatalf("client-oper-data did not decode one common and one dot11 record")
	}

	common := data.CommonOperData[0]
	dot11 := data.Dot11OperData[0]

	if common.ApName != "TEST-AP01" {
		t.Fatalf("ap-name = %q, so common-oper-data was not decoded", common.ApName)
	}
	if len(dot11.MultilinkInfo) != 1 {
		t.Fatalf("decoded %d multilink-info entries, want 1", len(dot11.MultilinkInfo))
	}
	if dot11.MultilinkInfo[0].Band != "dot11-6-ghz-band" {
		t.Fatalf("band = %q, so multilink-info was not decoded", dot11.MultilinkInfo[0].Band)
	}

	if common.MsRadioType != client.PHYRadioTypeDot11be6GHz {
		t.Errorf("ms-radio-type = %q, want %q", common.MsRadioType, client.PHYRadioTypeDot11be6GHz)
	}
	if dot11.EwlcMsPhyType != client.PHYRadioTypeDot11be6GHz {
		t.Errorf("ewlc-ms-phy-type = %q, want %q", dot11.EwlcMsPhyType, client.PHYRadioTypeDot11be6GHz)
	}
	if dot11.RadioType != client.RadioBandType6GHz {
		t.Errorf("radio-type = %q, want %q", dot11.RadioType, client.RadioBandType6GHz)
	}
	if dot11.MultilinkInfo[0].RadioType != client.RadioBandType6GHz {
		t.Errorf("multilink-info radio-type = %q, want %q",
			dot11.MultilinkInfo[0].RadioType, client.RadioBandType6GHz)
	}
	if dot11.MultilinkInfo[0].StaMAC == "" {
		t.Errorf("multilink-info sta-mac is empty, so LinkInfo dropped the link's own MAC")
	}
}

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

// byMACLookup names one of the nine list lookups this service keys by MAC address,
// together with the list node it reads, so the wire form can be asserted for each.
type byMACLookup struct {
	name string
	node string
	call func(ctx context.Context, s client.Service, mac string) error
}

// byMACLookups enumerates every lookup that rejects an unusable MAC address, so the
// sentinel, the malformed-address rejection and the wire form are asserted for all
// nine rather than for one.
func byMACLookups() []byMACLookup {
	return []byMACLookup{
		{
			name: "GetCommonInfoByMAC",
			node: "common-oper-data",
			call: func(ctx context.Context, s client.Service, mac string) error {
				_, err := s.GetCommonInfoByMAC(ctx, mac)
				return err
			},
		},
		{
			name: "GetDCInfoByMAC",
			node: "dc-info",
			call: func(ctx context.Context, s client.Service, mac string) error {
				_, err := s.GetDCInfoByMAC(ctx, mac)
				return err
			},
		},
		{
			name: "GetDot11InfoByMAC",
			node: "dot11-oper-data",
			call: func(ctx context.Context, s client.Service, mac string) error {
				_, err := s.GetDot11InfoByMAC(ctx, mac)
				return err
			},
		},
		{
			name: "GetMMIFClientHistoryByMAC",
			node: "mm-if-client-history",
			call: func(ctx context.Context, s client.Service, mac string) error {
				_, err := s.GetMMIFClientHistoryByMAC(ctx, mac)
				return err
			},
		},
		{
			name: "GetMMIFClientStatsByMAC",
			node: "mm-if-client-stats",
			call: func(ctx context.Context, s client.Service, mac string) error {
				_, err := s.GetMMIFClientStatsByMAC(ctx, mac)
				return err
			},
		},
		{
			name: "GetMobilityInfoByMAC",
			node: "mobility-oper-data",
			call: func(ctx context.Context, s client.Service, mac string) error {
				_, err := s.GetMobilityInfoByMAC(ctx, mac)
				return err
			},
		},
		{
			name: "GetPolicyInfoByMAC",
			node: "policy-data",
			call: func(ctx context.Context, s client.Service, mac string) error {
				_, err := s.GetPolicyInfoByMAC(ctx, mac)
				return err
			},
		},
		{
			name: "GetSISFDBByMAC",
			node: "sisf-db-mac",
			call: func(ctx context.Context, s client.Service, mac string) error {
				_, err := s.GetSISFDBByMAC(ctx, mac)
				return err
			},
		},
		{
			name: "GetTrafficStatsByMAC",
			node: "traffic-stats",
			call: func(ctx context.Context, s client.Service, mac string) error {
				_, err := s.GetTrafficStatsByMAC(ctx, mac)
				return err
			},
		},
	}
}

// TestClientServiceUnit_ValidationErrors_EmptyMAC pins the sentinel an empty list key
// carries. It is core.ErrResourceNotFound, which core.IsNotFoundError reports true for:
// asking for a record by an address that cannot name one is an absence, not a client
// misconfiguration.
func TestClientServiceUnit_ValidationErrors_EmptyMAC(t *testing.T) {
	t.Parallel()

	server := testutil.NewMockServer(testutil.WithSuccessResponses(map[string]string{}))
	defer server.Close()
	testClient := testutil.NewTestClient(server)
	service := client.NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	lookups := byMACLookups()
	if len(lookups) != 9 {
		t.Fatalf("Enumerated %d MAC lookups, want 9", len(lookups))
	}

	for _, lookup := range lookups {
		for _, mac := range []string{"", "   "} {
			err := lookup.call(ctx, service, mac)
			if !errors.Is(err, core.ErrResourceNotFound) {
				t.Errorf("%s(%q) error = %v, want core.ErrResourceNotFound", lookup.name, mac, err)
			}
			if errors.Is(err, core.ErrInvalidConfiguration) {
				t.Errorf("%s(%q) still reports the pre-v0.6.0 sentinel", lookup.name, mac)
			}
		}
	}
}

// TestClientServiceUnit_ValidationErrors_MalformedMAC pins that a malformed address is
// refused before a request is built, and that it is not reported as an absent record: a
// caller testing with core.IsNotFoundError must not read its own typo as an empty list.
func TestClientServiceUnit_ValidationErrors_MalformedMAC(t *testing.T) {
	t.Parallel()

	server := testutil.NewMockServer(testutil.WithSuccessResponses(map[string]string{}))
	defer server.Close()
	testClient := testutil.NewTestClient(server)
	service := client.NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	for _, lookup := range byMACLookups() {
		err := lookup.call(ctx, service, "not-a-mac")
		if err == nil {
			t.Errorf("%s(\"not-a-mac\") returned no error", lookup.name)
			continue
		}
		if core.IsNotFoundError(err) {
			t.Errorf("%s(\"not-a-mac\") error = %v, want a validation error", lookup.name, err)
		}
	}
}

// TestClientServiceUnit_ByMAC_WireForm pins the address form that reaches the controller,
// for all nine lookups rather than one. Every separator spelling and either case has to
// arrive as the lower-case colon form the controller keys its lists by, so a dashed
// address reads the same record.
func TestClientServiceUnit_ByMAC_WireForm(t *testing.T) {
	t.Parallel()

	const prefix = "/restconf/data/Cisco-IOS-XE-wireless-client-oper:client-oper-data/"

	server := testutil.NewRESTCONFServer(t)
	defer server.Close()

	lookups := byMACLookups()
	for _, lookup := range lookups {
		// The sole-key check holds the body against the node requested, so each node
		// needs its own envelope.
		body := `{"Cisco-IOS-XE-wireless-client-oper:` + lookup.node + `":[]}`
		server.AddHandler(http.MethodGet, "client-oper-data/"+lookup.node, func() (int, string) {
			return http.StatusOK, body
		})
	}

	testClient := testutil.NewTestClient(testutil.NewMockServerFromHTTP(server.Server))
	service := client.NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	spellings := []string{
		"00:11:22:33:44:55",
		"00-11-22-33-44-55",
		"0011.2233.4455",
		"001122334455",
		"AA:BB:CC:DD:EE:55",
	}
	wants := map[string]string{
		"00:11:22:33:44:55": "00:11:22:33:44:55",
		"00-11-22-33-44-55": "00:11:22:33:44:55",
		"0011.2233.4455":    "00:11:22:33:44:55",
		"001122334455":      "00:11:22:33:44:55",
		"AA:BB:CC:DD:EE:55": "aa:bb:cc:dd:ee:55",
	}

	for _, lookup := range lookups {
		for _, mac := range spellings {
			if err := lookup.call(ctx, service, mac); err != nil {
				t.Fatalf("%s(%q) unexpected error: %v", lookup.name, mac, err)
			}
		}
	}

	recorded := server.Requests()
	if want := len(lookups) * len(spellings); len(recorded) != want {
		t.Fatalf("Recorded %d requests, want %d", len(recorded), want)
	}

	i := 0
	for _, lookup := range lookups {
		for _, mac := range spellings {
			want := prefix + lookup.node + "=" + wants[mac]
			if got := recorded[i].Path; got != want {
				t.Errorf("%s(%q) wire path = %q, want %q", lookup.name, mac, got, want)
			}
			i++
		}
	}
}

// TestClientServiceUnit_Dot11Operations_MockSuccess tests ListDot11Info against a well-formed response.
func TestClientServiceUnit_Dot11Operations_MockSuccess(t *testing.T) {
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

// TestClientServiceUnit_ReadFailure_ReturnsError tests that a failed read reaches the caller
// instead of an empty client table.
func TestClientServiceUnit_ReadFailure_ReturnsError(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		response testutil.ResponseConfig
	}{
		{
			name: "UnexpectedEOF",
			response: testutil.ResponseConfig{
				StatusCode: 500,
				Body:       `{"ietf-restconf:errors": {"error": [{"error-message": "unexpected EOF"}]}}`,
			},
		},
		{
			name: "TableCursorFailure",
			response: testutil.ResponseConfig{
				StatusCode: 500,
				Body:       `{"ietf-restconf:errors": {"error": [{"error-message": "failed to retrieve table cursor"}]}}`,
			},
		},
		{
			name: "DBALResponseFailure",
			response: testutil.ResponseConfig{
				StatusCode: 500,
				Body:       `{"ietf-restconf:errors": {"error": [{"error-message": "Process DBAL response failed"}]}}`,
			},
		},
		{
			name: "UnknownFailure",
			response: testutil.ResponseConfig{
				StatusCode: 500,
				Body:       `{"ietf-restconf:errors": {"error": [{"error-message": "unknown database error"}]}}`,
			},
		},
		{
			name: "TruncatedBody",
			response: testutil.ResponseConfig{
				StatusCode: 200,
				Body:       `{"Cisco-IOS-XE-wireless-client-oper:client-oper-data":`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// The path of every dot11-oper-data read contains client-oper-data, so one
			// handler covers all three reads under test.
			server := testutil.NewMockServer(
				testutil.WithTesting(t),
				testutil.WithCustomResponse("client-oper-data", tt.response),
			)
			defer server.Close()

			testClient := testutil.NewTestClient(server)
			service := client.NewService(testClient.Core().(*core.Client))
			ctx := testutil.TestContext(t)

			oper, err := service.GetOperational(ctx)
			if err == nil {
				t.Error("Expected error from GetOperational, got nil")
			}
			if oper != nil {
				t.Error("Expected nil result from GetOperational")
			}

			dot11, err := service.ListDot11Info(ctx)
			if err == nil {
				t.Error("Expected error from ListDot11Info, got nil")
			}
			if dot11 != nil {
				t.Error("Expected nil result from ListDot11Info")
			}

			dot11ByMAC, err := service.GetDot11InfoByMAC(ctx, "aa:bb:cc:dd:ee:ff")
			if err == nil {
				t.Error("Expected error from GetDot11InfoByMAC, got nil")
			}
			if dot11ByMAC != nil {
				t.Error("Expected nil result from GetDot11InfoByMAC")
			}
		})
	}
}
