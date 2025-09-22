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
		mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(responses))
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
				"oper-data": [{
					"wtp-mac": "aa:bb:cc:dd:ee:ff",
					"radio-id": 4,
					"ap-antenna-band-mode": "ant-band-mode-unknown",
					"link-encryption-enabled": false,
					"ap-remote-debug-mode": false,
					"ap-ip-data": {
						"ap-prefix": 0,
						"mtu": 1485,
						"is-static-ap-ipaddr": true,
						"domain-name": "",
						"ap-ip-addr": "192.168.255.11",
						"ap-ipv6-addr": "::",
						"ap-ip-netmask": "255.255.255.0",
						"ap-ip-gateway": "192.168.255.1",
						"ap-ipv6-gateway": "::",
						"ap-name-server-type": "unknown",
						"ap-ipv6-method": "unknown-method",
						"static-ip": "192.168.255.11",
						"static-gw-ip": "192.168.255.1",
						"static-netmask": "255.255.255.0",
						"static-prefix": 0
					},
					"ap-prime-info": {
						"primary-controller-name": "WNC1",
						"secondary-controller-name": "",
						"primary-controller-ip-addr": "192.168.255.1",
						"secondary-controller-ip-addr": "0.0.0.0",
						"tertiary-controller-name": "",
						"tertiary-controller-ip-addr": "0.0.0.0",
						"ap-fallback-ip": "0.0.0.0",
						"fallback-enabled": true
					},
					"ap-pow": {
						"power-injector-sel": "pwrinj-selection-unknown",
						"power-injector-macaddr": "00:00:00:00:00:00",
						"pre-std-switch-enabled": false,
						"power-injector-enabled": false,
						"power-type": "pwr-src-poe-plus",
						"power-mode": "dot11-set-high-pwr"
					},
					"ap-indoor-mode": false,
					"is-local-net": false
				}],
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
		"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/oper-data": `{
			"Cisco-IOS-XE-wireless-access-point-oper:oper-data": [{
				"wtp-mac": "aa:bb:cc:dd:ee:ff",
				"radio-id": 4,
				"ap-antenna-band-mode": "ant-band-mode-unknown",
				"link-encryption-enabled": false,
				"ap-remote-debug-mode": false,
				"ap-ip-data": {
					"ap-prefix": 0,
					"mtu": 1485,
					"is-static-ap-ipaddr": true,
					"domain-name": "",
					"ap-ip-addr": "192.168.255.11",
					"ap-ipv6-addr": "::",
					"ap-ip-netmask": "255.255.255.0",
					"ap-ip-gateway": "192.168.255.1",
					"ap-ipv6-gateway": "::",
					"ap-name-server-type": "unknown",
					"ap-ipv6-method": "unknown-method",
					"static-ip": "192.168.255.11",
					"static-gw-ip": "192.168.255.1",
					"static-netmask": "255.255.255.0",
					"static-prefix": 0
				},
				"ap-prime-info": {
					"primary-controller-name": "WNC1",
					"secondary-controller-name": "",
					"primary-controller-ip-addr": "192.168.255.1",
					"secondary-controller-ip-addr": "0.0.0.0",
					"tertiary-controller-name": "",
					"tertiary-controller-ip-addr": "0.0.0.0",
					"ap-fallback-ip": "0.0.0.0",
					"fallback-enabled": true
				},
				"ap-pow": {
					"power-injector-sel": "pwrinj-selection-unknown",
					"power-injector-macaddr": "00:00:00:00:00:00",
					"pre-std-switch-enabled": false,
					"power-injector-enabled": false,
					"power-type": "pwr-src-poe-plus",
					"power-mode": "dot11-set-high-pwr"
				},
				"ap-indoor-mode": false,
				"is-local-net": false
			}]
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
		"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/iot-firmware": `{
			"Cisco-IOS-XE-wireless-access-point-oper:iot-firmware": [
				{
					"ap-mac": "aa:bb:cc:dd:ee:ff",
					"if-name": "ttyiot0",
					"is-default": [null],
					"version": "2.7.21",
					"vendor-name": "Cisco Systems Inc",
					"type": "iot-radio-fw-ble",
					"desc": "Firmware developed by Cisco for IoT use"
				},
				{
					"ap-mac": "aa:bb:cc:dd:ee:ff",
					"if-name": "ttyiot0",
					"is-default": [null],
					"version": "3.1.0",
					"vendor-name": "Cisco Systems Inc",
					"type": "iot-radio-fw-ble",
					"desc": "Firmware developed by Cisco for IoT use"
				}
			]
		}`,
		// New operational endpoints
		"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/radio-reset-stats": `{
			"Cisco-IOS-XE-wireless-access-point-oper:radio-reset-stats": [
				{
					"ap-mac": "aa:bb:cc:dd:ee:ff",
					"radio-id": 0,
					"cause": "none",
					"detail-cause": "none",
					"count": 0
				}
			]
		}`,
		"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/qos-client-data": `{
			"Cisco-IOS-XE-wireless-access-point-oper:qos-client-data": [
				{
					"client-mac": "80:7d:3a:77:19:a9",
					"aaa-qos-params": {
						"aaa-avgdtus": 0,
						"aaa-avgrtdtus": 0,
						"aaa-bstdtus": 0,
						"aaa-bstrtdtus": 0,
						"aaa-avgdtds": 0,
						"aaa-avgrtdtds": 0,
						"aaa-bstdtds": 0,
						"aaa-bstrtdtds": 0
					}
				}
			]
		}`,
		"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/wtp-slot-wlan-stats": `{
			"Cisco-IOS-XE-wireless-access-point-oper:wtp-slot-wlan-stats": [
				{
					"wtp-mac": "aa:bb:cc:dd:ee:ff",
					"slot-id": 0,
					"wlan-id": 1,
					"tx-bytes": 123456,
					"rx-bytes": 654321
				}
			]
		}`,
		"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/ethernet-mac-wtp-mac-map": `{
			"Cisco-IOS-XE-wireless-access-point-oper:ethernet-mac-wtp-mac-map": [
				{
					"ethernet-mac": "aa:bb:cc:dd:ee:ff",
					"wtp-mac": "bb:cc:dd:ee:ff:aa"
				}
			]
		}`,
		"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/radio-oper-stats": `{
			"Cisco-IOS-XE-wireless-access-point-oper:radio-oper-stats": [
				{
					"wtp-mac": "aa:bb:cc:dd:ee:ff",
					"slot-id": 0,
					"tx-frames": 100,
					"rx-frames": 200
				}
			]
		}`,
		"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/ethernet-if-stats": `{
			"Cisco-IOS-XE-wireless-access-point-oper:ethernet-if-stats": [
				{
					"wtp-mac": "aa:bb:cc:dd:ee:ff",
					"if-id": "GigabitEthernet0",
					"tx-bytes": 987654,
					"rx-bytes": 456789
				}
			]
		}`,
		"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/ewlc-wncd-stats": `{
			"Cisco-IOS-XE-wireless-access-point-oper:ewlc-wncd-stats": {
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
				"downloads-complete": 0
			}
		}`,
		"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/ap-iox-oper-data": `{
			"Cisco-IOS-XE-wireless-access-point-oper:ap-iox-oper-data": [
				{
					"wtp-mac": "aa:bb:cc:dd:ee:ff",
					"app-name": "test-app",
					"state": "running"
				}
			]
		}`,
		"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/qos-global-stats": `{
			"Cisco-IOS-XE-wireless-access-point-oper:qos-global-stats": {
				"qos-client-voice-stats": {
					"total-num-of-tspec-rcvd": 0,
					"new-tspec-from-assoc-req": 0,
					"tspec-renewal-from-assoc-req": 0,
					"new-tspec-as-add-ts": 0,
					"tspec-renewal-from-add-ts": 0,
					"tspec-process-failed-get-rec": 0,
					"total-sip-invite-on-caller": 0,
					"total-sip-invite-on-callee": 0,
					"total-num-of-call-report": 0,
					"total-sip-failure-trap-send": 0,
					"num-of-calls-accepted": 0,
					"num-of-calls-rejected-insuf-bw": 0,
					"num-of-calls-rejected-qos": 0,
					"num-of-calls-rejected-phy-rate": 0,
					"num-of-calls-rej-invalid-tspec": 0,
					"num-of-roam-calls-accepted": 0,
					"num-of-roam-calls-rejected": 0,
					"num-of-active-sip-calls": 0,
					"num-of-active-tspec-calls": 0
				}
			}
		}`,
		"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/rlan-oper": `{
			"Cisco-IOS-XE-wireless-access-point-oper:rlan-oper": [
				{
					"wtp-mac": "aa:bb:cc:dd:ee:ff",
					"rlan-port-id": 1,
					"rlan-oper-state": true,
					"rlan-port-status": true
				}
			]
		}`,
		"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/ewlc-mewlc-predownload-rec": `{
			"Cisco-IOS-XE-wireless-access-point-oper:ewlc-mewlc-predownload-rec": {
				"num-initiated": 0,
				"num-in-progress": 0,
				"num-complete": 0
			}
		}`,
		"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/cdp-cache-data": `{
			"Cisco-IOS-XE-wireless-access-point-oper:cdp-cache-data": [
				{
					"wtp-mac": "aa:bb:cc:dd:ee:ff",
					"local-intf-name": "GigabitEthernet0",
					"device-id": "Switch1"
				}
			]
		}`,
		"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/lldp-neigh": `{
			"Cisco-IOS-XE-wireless-access-point-oper:lldp-neigh": [
				{
					"wtp-mac": "aa:bb:cc:dd:ee:ff",
					"local-intf-name": "GigabitEthernet0",
					"device-id": "Switch1"
				}
			]
		}`,
		"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/tp-cert-info": `{
			"Cisco-IOS-XE-wireless-access-point-oper:tp-cert-info": {
				"trustpoint": {
					"trustpoint-name": "WNC1_WLC_TP",
					"is-cert-available": true,
					"is-privkey-available": true,
					"cert-hash": "1d35399409f0dd2274c49bbec14142b67f8f9a96",
					"cert-type": "trustpoint-ssc",
					"fips-suitability": "fips-na"
				}
			}
		}`,
		"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/disc-data": `{
			"Cisco-IOS-XE-wireless-access-point-oper:disc-data": [
				{
					"wtp-mac": "aa:bb:cc:dd:ee:ff",
					"disc-req": 10,
					"disc-rsp": 10
				}
			]
		}`,
		"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/country-oper": `{
			"Cisco-IOS-XE-wireless-access-point-oper:country-oper": [
				{
					"wtp-mac": "aa:bb:cc:dd:ee:ff",
					"radio-id": 0,
					"country-code": "US",
					"regulatory-domain": "FCC"
				}
			]
		}`,
		"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/supp-country-oper": `{
			"Cisco-IOS-XE-wireless-access-point-oper:supp-country-oper": [
				{
					"wtp-mac": "aa:bb:cc:dd:ee:ff",
					"radio-id": 0,
					"country-code": "US",
					"supported-channels": "1,6,11"
				}
			]
		}`,
		"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/ap-nh-global-data": `{
			"Cisco-IOS-XE-wireless-access-point-oper:ap-nh-global-data": {
				"algorithm-running": false,
				"algorithm-itr-count": 0,
				"ideal-capacity-per-rg": 0,
				"num-of-neighborhood": 0
			}
		}`,
	}
	mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(responses))
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

	t.Run("ListApOperData", func(t *testing.T) {
		result, err := service.ListApOperData(ctx)
		if err != nil {
			t.Errorf("Expected no error for ListApOperData, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for ListApOperData, got nil")
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

	t.Run("ListRadioData", func(t *testing.T) {
		result, err := service.ListRadioData(ctx)
		if err != nil {
			t.Errorf("Expected no error for ListRadioData, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for ListRadioData, got nil")
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

	t.Run("ListIotFirmware", func(t *testing.T) {
		result, err := service.ListIotFirmware(ctx)
		if err != nil {
			t.Errorf("Expected no error for ListIotFirmware, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for ListIotFirmware, got nil")
		}
	})

	// Test newly added AP operational data operations
	t.Run("ListRadioResetStats", func(t *testing.T) {
		result, err := service.ListRadioResetStats(ctx)
		if err != nil {
			t.Errorf("Expected no error for ListRadioResetStats, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for ListRadioResetStats, got nil")
		}
	})

	t.Run("ListQosClientData", func(t *testing.T) {
		result, err := service.ListQosClientData(ctx)
		if err != nil {
			t.Errorf("Expected no error for ListQosClientData, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for ListQosClientData, got nil")
		}
	})

	t.Run("ListWtpSlotWlanStats", func(t *testing.T) {
		result, err := service.ListWtpSlotWlanStats(ctx)
		if err != nil {
			t.Errorf("Expected no error for ListWtpSlotWlanStats, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for ListWtpSlotWlanStats, got nil")
		}
	})

	t.Run("ListEthernetMACWtpMACMaps", func(t *testing.T) {
		result, err := service.ListEthernetMACWtpMACMaps(ctx)
		if err != nil {
			t.Errorf("Expected no error for ListEthernetMACWtpMACMaps, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for ListEthernetMACWtpMACMaps, got nil")
		}
	})

	t.Run("ListRadioOperStats", func(t *testing.T) {
		result, err := service.ListRadioOperStats(ctx)
		if err != nil {
			t.Errorf("Expected no error for ListRadioOperStats, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for ListRadioOperStats, got nil")
		}
	})

	t.Run("ListEthernetIfStats", func(t *testing.T) {
		result, err := service.ListEthernetIfStats(ctx)
		if err != nil {
			t.Errorf("Expected no error for ListEthernetIfStats, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for ListEthernetIfStats, got nil")
		}
	})

	t.Run("ListEwlcWncdStats", func(t *testing.T) {
		result, err := service.ListEwlcWncdStats(ctx)
		if err != nil {
			t.Errorf("Expected no error for ListEwlcWncdStats, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for ListEwlcWncdStats, got nil")
		}
	})

	t.Run("ListApIoxOperData", func(t *testing.T) {
		result, err := service.ListApIoxOperData(ctx)
		if err != nil {
			t.Errorf("Expected no error for ListApIoxOperData, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for ListApIoxOperData, got nil")
		}
	})

	t.Run("ListQosGlobalStats", func(t *testing.T) {
		result, err := service.ListQosGlobalStats(ctx)
		if err != nil {
			t.Errorf("Expected no error for ListQosGlobalStats, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for ListQosGlobalStats, got nil")
		}
	})

	t.Run("ListRlanOper", func(t *testing.T) {
		result, err := service.ListRlanOper(ctx)
		if err != nil {
			t.Errorf("Expected no error for ListRlanOper, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for ListRlanOper, got nil")
		}
	})

	t.Run("ListEwlcMewlcPredownloadRec", func(t *testing.T) {
		result, err := service.ListEwlcMewlcPredownloadRec(ctx)
		if err != nil {
			t.Errorf("Expected no error for ListEwlcMewlcPredownloadRec, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for ListEwlcMewlcPredownloadRec, got nil")
		}
	})

	t.Run("ListCdpCacheData", func(t *testing.T) {
		result, err := service.ListCdpCacheData(ctx)
		if err != nil {
			t.Errorf("Expected no error for ListCdpCacheData, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for ListCdpCacheData, got nil")
		}
	})

	t.Run("ListLldpNeigh", func(t *testing.T) {
		result, err := service.ListLldpNeigh(ctx)
		if err != nil {
			t.Errorf("Expected no error for ListLldpNeigh, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for ListLldpNeigh, got nil")
		}
	})

	t.Run("ListTpCertInfo", func(t *testing.T) {
		result, err := service.ListTpCertInfo(ctx)
		if err != nil {
			t.Errorf("Expected no error for ListTpCertInfo, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for ListTpCertInfo, got nil")
		}
	})

	t.Run("ListDiscData", func(t *testing.T) {
		result, err := service.ListDiscData(ctx)
		if err != nil {
			t.Errorf("Expected no error for ListDiscData, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for ListDiscData, got nil")
		}
	})

	t.Run("ListCountryOper", func(t *testing.T) {
		result, err := service.ListCountryOper(ctx)
		if err != nil {
			t.Errorf("Expected no error for ListCountryOper, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for ListCountryOper, got nil")
		}
	})

	t.Run("ListSuppCountryOper", func(t *testing.T) {
		result, err := service.ListSuppCountryOper(ctx)
		if err != nil {
			t.Errorf("Expected no error for ListSuppCountryOper, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for ListSuppCountryOper, got nil")
		}
	})

	t.Run("ListApNhGlobalData", func(t *testing.T) {
		result, err := service.ListApNhGlobalData(ctx)
		if err != nil {
			t.Errorf("Expected no error for ListApNhGlobalData, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for ListApNhGlobalData, got nil")
		}
	})
}

// TestApServiceUnit_GetOperations_ErrorHandling tests error scenarios using mock server.
func TestApServiceUnit_GetOperations_ErrorHandling(t *testing.T) {
	// Create mock server that returns 404 for AP endpoints
	errorPaths := []string{
		"Cisco-IOS-XE-wireless-ap-cfg:ap-cfg-data",
		"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data",
		"Cisco-IOS-XE-wireless-ap-global-oper:ap-global-oper-data",
		// New endpoint paths for error testing
		"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/radio-reset-stats",
		"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/qos-client-data",
		"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/wtp-slot-wlan-stats",
		"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/ethernet-mac-wtp-mac-map",
		"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/radio-oper-stats",
		"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/ethernet-if-stats",
		"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/ewlc-wncd-stats",
		"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/ap-iox-oper-data",
		"Cisco-IOS-XE-wireless-ap-global-oper:ap-global-oper-data/qos-global-stats",
		"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/rlan-oper",
		"Cisco-IOS-XE-wireless-ap-global-oper:ap-global-oper-data/ewlc-mewlc-predownload-rec",
		"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/cdp-cache-data",
		"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/lldp-neigh",
		"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/tp-cert-info",
		"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/disc-data",
		"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/country-oper",
		"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/supp-country-oper",
		"Cisco-IOS-XE-wireless-ap-global-oper:ap-global-oper-data/ap-nh-global-data",
	}
	mockServer := testutil.NewMockServer(testutil.WithErrorResponses(errorPaths, 404))
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

	// Test newly added List functions error handling
	t.Run("ListRadioResetStats", func(t *testing.T) {
		_, err := service.ListRadioResetStats(ctx)
		if err == nil {
			t.Error("Expected error for 404 response, got nil")
		}
		if !core.IsNotFoundError(err) {
			t.Errorf("Expected NotFound error, got: %v", err)
		}
	})

	t.Run("ListQosClientData", func(t *testing.T) {
		_, err := service.ListQosClientData(ctx)
		if err == nil {
			t.Error("Expected error for 404 response, got nil")
		}
		if !core.IsNotFoundError(err) {
			t.Errorf("Expected NotFound error, got: %v", err)
		}
	})

	t.Run("ListWtpSlotWlanStats", func(t *testing.T) {
		_, err := service.ListWtpSlotWlanStats(ctx)
		if err == nil {
			t.Error("Expected error for 404 response, got nil")
		}
		if !core.IsNotFoundError(err) {
			t.Errorf("Expected NotFound error, got: %v", err)
		}
	})

	t.Run("ListEthernetMACWtpMACMaps", func(t *testing.T) {
		_, err := service.ListEthernetMACWtpMACMaps(ctx)
		if err == nil {
			t.Error("Expected error for 404 response, got nil")
		}
		if !core.IsNotFoundError(err) {
			t.Errorf("Expected NotFound error, got: %v", err)
		}
	})

	t.Run("ListRadioOperStats", func(t *testing.T) {
		_, err := service.ListRadioOperStats(ctx)
		if err == nil {
			t.Error("Expected error for 404 response, got nil")
		}
		if !core.IsNotFoundError(err) {
			t.Errorf("Expected NotFound error, got: %v", err)
		}
	})

	t.Run("ListEthernetIfStats", func(t *testing.T) {
		_, err := service.ListEthernetIfStats(ctx)
		if err == nil {
			t.Error("Expected error for 404 response, got nil")
		}
		if !core.IsNotFoundError(err) {
			t.Errorf("Expected NotFound error, got: %v", err)
		}
	})

	t.Run("ListEwlcWncdStats", func(t *testing.T) {
		_, err := service.ListEwlcWncdStats(ctx)
		if err == nil {
			t.Error("Expected error for 404 response, got nil")
		}
		if !core.IsNotFoundError(err) {
			t.Errorf("Expected NotFound error, got: %v", err)
		}
	})

	t.Run("ListApIoxOperData", func(t *testing.T) {
		_, err := service.ListApIoxOperData(ctx)
		if err == nil {
			t.Error("Expected error for 404 response, got nil")
		}
		if !core.IsNotFoundError(err) {
			t.Errorf("Expected NotFound error, got: %v", err)
		}
	})

	t.Run("ListQosGlobalStats", func(t *testing.T) {
		_, err := service.ListQosGlobalStats(ctx)
		if err == nil {
			t.Error("Expected error for 404 response, got nil")
		}
		if !core.IsNotFoundError(err) {
			t.Errorf("Expected NotFound error, got: %v", err)
		}
	})

	t.Run("ListRlanOper", func(t *testing.T) {
		_, err := service.ListRlanOper(ctx)
		if err == nil {
			t.Error("Expected error for 404 response, got nil")
		}
		if !core.IsNotFoundError(err) {
			t.Errorf("Expected NotFound error, got: %v", err)
		}
	})

	t.Run("ListEwlcMewlcPredownloadRec", func(t *testing.T) {
		_, err := service.ListEwlcMewlcPredownloadRec(ctx)
		if err == nil {
			t.Error("Expected error for 404 response, got nil")
		}
		if !core.IsNotFoundError(err) {
			t.Errorf("Expected NotFound error, got: %v", err)
		}
	})

	t.Run("ListCdpCacheData", func(t *testing.T) {
		_, err := service.ListCdpCacheData(ctx)
		if err == nil {
			t.Error("Expected error for 404 response, got nil")
		}
		if !core.IsNotFoundError(err) {
			t.Errorf("Expected NotFound error, got: %v", err)
		}
	})

	t.Run("ListLldpNeigh", func(t *testing.T) {
		_, err := service.ListLldpNeigh(ctx)
		if err == nil {
			t.Error("Expected error for 404 response, got nil")
		}
		if !core.IsNotFoundError(err) {
			t.Errorf("Expected NotFound error, got: %v", err)
		}
	})

	t.Run("ListTpCertInfo", func(t *testing.T) {
		_, err := service.ListTpCertInfo(ctx)
		if err == nil {
			t.Error("Expected error for 404 response, got nil")
		}
		if !core.IsNotFoundError(err) {
			t.Errorf("Expected NotFound error, got: %v", err)
		}
	})

	t.Run("ListDiscData", func(t *testing.T) {
		_, err := service.ListDiscData(ctx)
		if err == nil {
			t.Error("Expected error for 404 response, got nil")
		}
		if !core.IsNotFoundError(err) {
			t.Errorf("Expected NotFound error, got: %v", err)
		}
	})

	t.Run("ListCountryOper", func(t *testing.T) {
		_, err := service.ListCountryOper(ctx)
		if err == nil {
			t.Error("Expected error for 404 response, got nil")
		}
		if !core.IsNotFoundError(err) {
			t.Errorf("Expected NotFound error, got: %v", err)
		}
	})

	t.Run("ListSuppCountryOper", func(t *testing.T) {
		_, err := service.ListSuppCountryOper(ctx)
		if err == nil {
			t.Error("Expected error for 404 response, got nil")
		}
		if !core.IsNotFoundError(err) {
			t.Errorf("Expected NotFound error, got: %v", err)
		}
	})

	t.Run("ListApNhGlobalData", func(t *testing.T) {
		_, err := service.ListApNhGlobalData(ctx)
		if err == nil {
			t.Error("Expected error for 404 response, got nil")
		}
		if !core.IsNotFoundError(err) {
			t.Errorf("Expected NotFound error, got: %v", err)
		}
	})

	// Test GetBy* filtered functions error handling
	t.Run("GetRadioResetStatsByAPMACAndRadioID", func(t *testing.T) {
		_, err := service.GetRadioResetStatsByAPMACAndRadioID(ctx, "aa:bb:cc:dd:ee:ff", 0)
		if err == nil {
			t.Error("Expected error for 404 response, got nil")
		}
		if !core.IsNotFoundError(err) {
			t.Errorf("Expected NotFound error, got: %v", err)
		}
	})

	t.Run("GetQosClientDataByClientMAC", func(t *testing.T) {
		_, err := service.GetQosClientDataByClientMAC(ctx, "80:7d:3a:77:19:a9")
		if err == nil {
			t.Error("Expected error for 404 response, got nil")
		}
		if !core.IsNotFoundError(err) {
			t.Errorf("Expected NotFound error, got: %v", err)
		}
	})
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
		// New filtered endpoints
		"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/radio-reset-stats=aa:bb:cc:dd:ee:ff,0": `{
			"Cisco-IOS-XE-wireless-access-point-oper:radio-reset-stats": [{
				"ap-mac": "aa:bb:cc:dd:ee:ff",
				"radio-id": 0,
				"cause": "none",
				"detail-cause": "none",
				"count": 0
			}]
		}`,
		"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/qos-client-data=80:7d:3a:77:19:a9": `{
			"Cisco-IOS-XE-wireless-access-point-oper:qos-client-data": [{
				"client-mac": "80:7d:3a:77:19:a9",
				"aaa-qos-params": {
					"aaa-avgdtus": 0,
					"aaa-avgrtdtus": 0,
					"aaa-bstdtus": 0,
					"aaa-bstrtdtus": 0,
					"aaa-avgdtds": 0,
					"aaa-avgrtdtds": 0,
					"aaa-bstdtds": 0,
					"aaa-bstrtdtds": 0
				}
			}]
		}`,
		"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/wtp-slot-wlan-stats=aa:bb:cc:dd:ee:ff,0,1": `{
			"Cisco-IOS-XE-wireless-access-point-oper:wtp-slot-wlan-stats": [{
				"wtp-mac": "aa:bb:cc:dd:ee:ff",
				"slot-id": 0,
				"wlan-id": 1,
				"tx-bytes": 123456,
				"rx-bytes": 654321
			}]
		}`,
		"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/ethernet-mac-wtp-mac-map=aa:bb:cc:dd:ee:ff": `{
			"Cisco-IOS-XE-wireless-access-point-oper:ethernet-mac-wtp-mac-map": [{
				"ethernet-mac": "aa:bb:cc:dd:ee:ff",
				"wtp-mac": "bb:cc:dd:ee:ff:aa"
			}]
		}`,
		"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/radio-oper-stats=aa:bb:cc:dd:ee:ff,0": `{
			"Cisco-IOS-XE-wireless-access-point-oper:radio-oper-stats": [{
				"wtp-mac": "aa:bb:cc:dd:ee:ff",
				"slot-id": 0,
				"tx-frames": 100,
				"rx-frames": 200
			}]
		}`,
		"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/ethernet-if-stats=aa:bb:cc:dd:ee:ff,0": `{
			"Cisco-IOS-XE-wireless-access-point-oper:ethernet-if-stats": [{
				"wtp-mac": "aa:bb:cc:dd:ee:ff",
				"if-id": "0",
				"tx-bytes": 987654,
				"rx-bytes": 456789
			}]
		}`,
		"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/ap-iox-oper-data=aa:bb:cc:dd:ee:ff": `{
			"Cisco-IOS-XE-wireless-access-point-oper:ap-iox-oper-data": [{
				"wtp-mac": "aa:bb:cc:dd:ee:ff",
				"app-name": "test-app",
				"state": "running"
			}]
		}`,
		"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/rlan-oper=aa:bb:cc:dd:ee:ff,1": `{
			"Cisco-IOS-XE-wireless-access-point-oper:rlan-oper": [{
				"wtp-mac": "aa:bb:cc:dd:ee:ff",
				"rlan-port-id": 1,
				"rlan-oper-state": true,
				"rlan-port-status": true
			}]
		}`,
		"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/cdp-cache-data=aa:bb:cc:dd:ee:ff": `{
			"Cisco-IOS-XE-wireless-access-point-oper:cdp-cache-data": [{
				"wtp-mac": "aa:bb:cc:dd:ee:ff",
				"local-intf-name": "GigabitEthernet0",
				"device-id": "Switch1"
			}]
		}`,
		"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/lldp-neigh=aa:bb:cc:dd:ee:ff": `{
			"Cisco-IOS-XE-wireless-access-point-oper:lldp-neigh": [{
				"wtp-mac": "aa:bb:cc:dd:ee:ff",
				"local-intf-name": "GigabitEthernet0",
				"device-id": "Switch1"
			}]
		}`,
		"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/disc-data=aa:bb:cc:dd:ee:ff": `{
			"Cisco-IOS-XE-wireless-access-point-oper:disc-data": [{
				"wtp-mac": "aa:bb:cc:dd:ee:ff",
				"disc-req": 10,
				"disc-rsp": 10
			}]
		}`,
		"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/country-oper=aa:bb:cc:dd:ee:ff,0": `{
			"Cisco-IOS-XE-wireless-access-point-oper:country-oper": [{
				"wtp-mac": "aa:bb:cc:dd:ee:ff",
				"radio-id": 0,
				"country-code": "US",
				"regulatory-domain": "FCC"
			}]
		}`,
		"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/supp-country-oper=aa:bb:cc:dd:ee:ff,0": `{
			"Cisco-IOS-XE-wireless-access-point-oper:supp-country-oper": [{
				"wtp-mac": "aa:bb:cc:dd:ee:ff",
				"radio-id": 0,
				"country-code": "US",
				"supported-channels": "1,6,11"
			}]
		}`,
	}
	mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(responses))
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

	// Test individual filtered operations that were missing
	t.Run("GetDiscDataByWTPMAC", func(t *testing.T) {
		result, err := service.GetDiscDataByWTPMAC(ctx, "aa:bb:cc:dd:ee:ff")
		if err != nil {
			t.Errorf("Expected no error for GetDiscDataByWTPMAC, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for GetDiscDataByWTPMAC, got nil")
		}
	})

	t.Run("GetCountryOperByWTPMACAndRadioID", func(t *testing.T) {
		result, err := service.GetCountryOperByWTPMACAndRadioID(ctx, "aa:bb:cc:dd:ee:ff", 0)
		if err != nil {
			t.Errorf("Expected no error for GetCountryOperByWTPMACAndRadioID, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for GetCountryOperByWTPMACAndRadioID, got nil")
		}
	})

	t.Run("GetSuppCountryOperByWTPMACAndRadioID", func(t *testing.T) {
		result, err := service.GetSuppCountryOperByWTPMACAndRadioID(ctx, "aa:bb:cc:dd:ee:ff", 0)
		if err != nil {
			t.Errorf("Expected no error for GetSuppCountryOperByWTPMACAndRadioID, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for GetSuppCountryOperByWTPMACAndRadioID, got nil")
		}
	})

	t.Run("GetLldpNeighByWTPMAC", func(t *testing.T) {
		result, err := service.GetLldpNeighByWTPMAC(ctx, "aa:bb:cc:dd:ee:ff")
		if err != nil {
			t.Errorf("Expected no error for GetLldpNeighByWTPMAC, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for GetLldpNeighByWTPMAC, got nil")
		}
	})

	t.Run("GetCdpCacheDataByWTPMAC", func(t *testing.T) {
		result, err := service.GetCdpCacheDataByWTPMAC(ctx, "aa:bb:cc:dd:ee:ff")
		if err != nil {
			t.Errorf("Expected no error for GetCdpCacheDataByWTPMAC, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for GetCdpCacheDataByWTPMAC, got nil")
		}
	})

	t.Run("GetApIoxOperDataByWTPMAC", func(t *testing.T) {
		result, err := service.GetApIoxOperDataByWTPMAC(ctx, "aa:bb:cc:dd:ee:ff")
		if err != nil {
			t.Errorf("Expected no error for GetApIoxOperDataByWTPMAC, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for GetApIoxOperDataByWTPMAC, got nil")
		}
	})

	t.Run("GetRadioOperStatsByWTPMACAndSlot", func(t *testing.T) {
		result, err := service.GetRadioOperStatsByWTPMACAndSlot(ctx, "aa:bb:cc:dd:ee:ff", 0)
		if err != nil {
			t.Errorf("Expected no error for GetRadioOperStatsByWTPMACAndSlot, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for GetRadioOperStatsByWTPMACAndSlot, got nil")
		}
	})

	t.Run("GetEthernetMACWtpMACMapByEthernetMAC", func(t *testing.T) {
		result, err := service.GetEthernetMACWtpMACMapByEthernetMAC(ctx, "aa:bb:cc:dd:ee:ff")
		if err != nil {
			t.Errorf("Expected no error for GetEthernetMACWtpMACMapByEthernetMAC, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for GetEthernetMACWtpMACMapByEthernetMAC, got nil")
		}
	})

	t.Run("GetEthernetIfStatsByWTPMACAndInterfaceID", func(t *testing.T) {
		result, err := service.GetEthernetIfStatsByWTPMACAndInterfaceID(ctx, "aa:bb:cc:dd:ee:ff", "0")
		if err != nil {
			t.Errorf("Expected no error for GetEthernetIfStatsByWTPMACAndInterfaceID, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for GetEthernetIfStatsByWTPMACAndInterfaceID, got nil")
		}
	})
}

// TestApServiceUnit_GetOperations_ValidationErrors tests input validation scenarios.
func TestApServiceUnit_GetOperations_ValidationErrors(t *testing.T) {
	// Use minimal mock server since we're testing validation before network calls
	responses := map[string]string{}
	mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(responses))
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

	// Test validation for newly added GetBy* functions
	t.Run("GetRadioResetStatsByAPMACAndRadioID_EmptyMAC", func(t *testing.T) {
		_, err := service.GetRadioResetStatsByAPMACAndRadioID(ctx, "", 0)
		if err == nil {
			t.Error("Expected error for empty MAC address, got nil")
		}
	})

	t.Run("GetRadioResetStatsByAPMACAndRadioID_InvalidMAC", func(t *testing.T) {
		_, err := service.GetRadioResetStatsByAPMACAndRadioID(ctx, "invalid-mac", 0)
		if err == nil {
			t.Error("Expected error for invalid MAC address, got nil")
		}
	})

	t.Run("GetQosClientDataByClientMAC_EmptyMAC", func(t *testing.T) {
		_, err := service.GetQosClientDataByClientMAC(ctx, "")
		if err == nil {
			t.Error("Expected error for empty MAC address, got nil")
		}
	})

	t.Run("GetQosClientDataByClientMAC_InvalidMAC", func(t *testing.T) {
		_, err := service.GetQosClientDataByClientMAC(ctx, "invalid-mac")
		if err == nil {
			t.Error("Expected error for invalid MAC address, got nil")
		}
	})

	t.Run("GetWtpSlotWlanStatsByWTPMACSlotAndWLANID_EmptyMAC", func(t *testing.T) {
		_, err := service.GetWtpSlotWlanStatsByWTPMACSlotAndWLANID(ctx, "", 0, 1)
		if err == nil {
			t.Error("Expected error for empty MAC address, got nil")
		}
	})

	t.Run("GetWtpSlotWlanStatsByWTPMACSlotAndWLANID_InvalidWLANID", func(t *testing.T) {
		_, err := service.GetWtpSlotWlanStatsByWTPMACSlotAndWLANID(ctx, "aa:bb:cc:dd:ee:ff", 0, 0)
		if err == nil {
			t.Error("Expected error for invalid WLAN ID, got nil")
		}
	})

	t.Run("GetEthernetMACWtpMACMapByEthernetMAC_EmptyMAC", func(t *testing.T) {
		_, err := service.GetEthernetMACWtpMACMapByEthernetMAC(ctx, "")
		if err == nil {
			t.Error("Expected error for empty MAC address, got nil")
		}
	})

	t.Run("GetRadioOperStatsByWTPMACAndSlot_EmptyMAC", func(t *testing.T) {
		_, err := service.GetRadioOperStatsByWTPMACAndSlot(ctx, "", 0)
		if err == nil {
			t.Error("Expected error for empty MAC address, got nil")
		}
	})

	t.Run("GetEthernetIfStatsByWTPMACAndInterfaceID_EmptyMAC", func(t *testing.T) {
		_, err := service.GetEthernetIfStatsByWTPMACAndInterfaceID(ctx, "", "GigabitEthernet0")
		if err == nil {
			t.Error("Expected error for empty MAC address, got nil")
		}
	})

	t.Run("GetEthernetIfStatsByWTPMACAndInterfaceID_EmptyInterfaceID", func(t *testing.T) {
		_, err := service.GetEthernetIfStatsByWTPMACAndInterfaceID(ctx, "aa:bb:cc:dd:ee:ff", "")
		if err == nil {
			t.Error("Expected error for empty interface ID, got nil")
		}
	})

	t.Run("GetApIoxOperDataByWTPMAC_EmptyMAC", func(t *testing.T) {
		_, err := service.GetApIoxOperDataByWTPMAC(ctx, "")
		if err == nil {
			t.Error("Expected error for empty MAC address, got nil")
		}
	})

	t.Run("GetRlanOperByWTPMACAndPortID_EmptyMAC", func(t *testing.T) {
		_, err := service.GetRlanOperByWTPMACAndPortID(ctx, "", 1)
		if err == nil {
			t.Error("Expected error for empty MAC address, got nil")
		}
	})

	t.Run("GetRlanOperByWTPMACAndPortID_InvalidPortID", func(t *testing.T) {
		_, err := service.GetRlanOperByWTPMACAndPortID(ctx, "aa:bb:cc:dd:ee:ff", 0)
		if err == nil {
			t.Error("Expected error for invalid port ID, got nil")
		}
	})

	t.Run("GetCdpCacheDataByWTPMAC_EmptyMAC", func(t *testing.T) {
		_, err := service.GetCdpCacheDataByWTPMAC(ctx, "")
		if err == nil {
			t.Error("Expected error for empty MAC address, got nil")
		}
	})

	t.Run("GetLldpNeighByWTPMAC_EmptyMAC", func(t *testing.T) {
		_, err := service.GetLldpNeighByWTPMAC(ctx, "")
		if err == nil {
			t.Error("Expected error for empty MAC address, got nil")
		}
	})

	t.Run("GetDiscDataByWTPMAC_EmptyMAC", func(t *testing.T) {
		_, err := service.GetDiscDataByWTPMAC(ctx, "")
		if err == nil {
			t.Error("Expected error for empty MAC address, got nil")
		}
	})

	t.Run("GetCountryOperByWTPMACAndRadioID_EmptyMAC", func(t *testing.T) {
		_, err := service.GetCountryOperByWTPMACAndRadioID(ctx, "", 0)
		if err == nil {
			t.Error("Expected error for empty MAC address, got nil")
		}
	})

	t.Run("GetSuppCountryOperByWTPMACAndRadioID_EmptyMAC", func(t *testing.T) {
		_, err := service.GetSuppCountryOperByWTPMACAndRadioID(ctx, "", 0)
		if err == nil {
			t.Error("Expected error for empty MAC address, got nil")
		}
	})
}

// TestApServiceUnit_GetOperations_EdgeCaseValidation tests additional validation edge cases.
func TestApServiceUnit_GetOperations_EdgeCaseValidation(t *testing.T) {
	responses := map[string]string{}
	mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(responses))
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
	mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(responses))
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
	mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(responses))
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

// TestApServiceUnit_DoOperations_MockSuccess tests edge cases and error branches.
func TestApServiceUnit_DoOperations_MockSuccess(t *testing.T) {
	// Create mock server with specific responses for edge cases
	responses := map[string]string{
		"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/capwap-data": `{
			"Cisco-IOS-XE-wireless-access-point-oper:capwap-data": []
		}`,
		"Cisco-IOS-XE-wireless-access-point-cmd-rpc:ap-reset": `{"status": "success"}`,
	}
	mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(responses))
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

// TestApServiceUnit_DoOperations_ErrorHandling tests nil CAPWAP data handling.
func TestApServiceUnit_DoOperations_ErrorHandling(t *testing.T) {
	mockServer := testutil.NewMockServer(testutil.WithErrorResponses(
		[]string{"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/capwap-data"},
		500,
	))
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

// TestApServiceUnit_Reload_EdgeCases tests specific edge cases for Reload function to achieve 100% coverage.
func TestApServiceUnit_Reload_EdgeCases(t *testing.T) {
	// Test Reload with nil CAPWAP data response
	t.Run("Reload_NilCAPWAPResponse", func(t *testing.T) {
		responses := map[string]string{
			"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/capwap-data": `null`,
		}
		mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(responses))
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
		mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(responses))
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

// TestApTagServiceUnit_SetOperations_ErrorHandling tests additional error handling scenarios for 100% coverage.
func TestApTagServiceUnit_SetOperations_ErrorHandling(t *testing.T) {
	// Mock server with specific error responses for edge cases
	mockServer := testutil.NewMockServer(
		testutil.WithTesting(t),
		testutil.WithCustomResponse(
			"Cisco-IOS-XE-wireless-access-point-cfg-rpc:set-ap-admin-state", testutil.ResponseConfig{
				StatusCode: 400,
				Body:       "Invalid request",
			}),
		testutil.WithCustomResponse(
			"Cisco-IOS-XE-wireless-access-point-cfg-rpc:set-ap-slot-admin-state", testutil.ResponseConfig{
				StatusCode: 400,
				Body:       "Invalid request",
			}),
		testutil.WithCustomResponse(
			"Cisco-IOS-XE-wireless-ap-cfg:ap-cfg-data/ap-tag=aa:bb:cc:dd:ee:ff", testutil.ResponseConfig{
				StatusCode: 400,
				Body:       "Invalid request",
			}),
	)
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
