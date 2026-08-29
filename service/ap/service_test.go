package ap_test

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
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
	t.Run("GetGlobalOperational", func(t *testing.T) {
		result, err := service.GetGlobalOperational(ctx)
		if err != nil {
			t.Errorf("Expected no error for GetGlobalOperational, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for GetGlobalOperational, got nil")
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
			"Cisco-IOS-XE-wireless-ap-cfg:ap-tag": [{
				"ap-mac": "28:ac:9e:11:48:10",
				"policy-tag": "labo-wlan-flex",
				"site-tag": "labo-site-flex",
				"rf-tag": "labo-inside"
			}]
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
		"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/ethernet-if-stats=aa:bb:cc:dd:ee:ff,GigabitEthernet0": `{
			"Cisco-IOS-XE-wireless-access-point-oper:ethernet-if-stats": [{
				"wtp-mac": "aa:bb:cc:dd:ee:ff",
				"if-id": "GigabitEthernet0",
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
		"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/lldp-neigh=aa:bb:cc:dd:ee:ff,11:22:33:44:55:66": `{
			"Cisco-IOS-XE-wireless-access-point-oper:lldp-neigh": [{
				"wtp-mac": "aa:bb:cc:dd:ee:ff",
				"neigh-mac": "11:22:33:44:55:66",
				"local-port": "GigabitEthernet0"
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

	t.Run("GetLldpNeighByWTPMACAndNeighMAC", func(t *testing.T) {
		result, err := service.GetLldpNeighByWTPMACAndNeighMAC(ctx, "aa:bb:cc:dd:ee:ff", "11:22:33:44:55:66")
		if err != nil {
			t.Errorf("Expected no error for GetLldpNeighByWTPMACAndNeighMAC, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for GetLldpNeighByWTPMACAndNeighMAC, got nil")
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
		result, err := service.GetEthernetIfStatsByWTPMACAndInterfaceID(ctx, "aa:bb:cc:dd:ee:ff", "GigabitEthernet0")
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

	t.Run("GetLldpNeighByWTPMACAndNeighMAC_EmptyWTPMAC", func(t *testing.T) {
		_, err := service.GetLldpNeighByWTPMACAndNeighMAC(ctx, "", "11:22:33:44:55:66")
		if err == nil {
			t.Error("Expected error for empty MAC address, got nil")
		}
	})

	t.Run("GetLldpNeighByWTPMACAndNeighMAC_EmptyNeighMAC", func(t *testing.T) {
		_, err := service.GetLldpNeighByWTPMACAndNeighMAC(ctx, "aa:bb:cc:dd:ee:ff", "")
		if err == nil {
			t.Error("Expected error for empty neighbor MAC address, got nil")
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
		"Cisco-IOS-XE-wireless-ap-cfg:ap-cfg-data/ap-tags/ap-tag=aa:bb:cc:dd:ee:ff": `{
			"Cisco-IOS-XE-wireless-ap-cfg:ap-tag": [{
				"ap-mac": "aa:bb:cc:dd:ee:ff",
				"site-tag": "existing-site",
				"policy-tag": "existing-policy",
				"rf-tag": "existing-rf"
			}]
		}`,
		"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/capwap-data=aa:bb:cc:dd:ee:ff": `{
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
	t.Run("EnableAPByMAC", func(t *testing.T) {
		err := service.EnableAPByMAC(ctx, "aa:bb:cc:dd:ee:ff")
		if err != nil {
			t.Errorf("Expected no error for EnableAPByMAC, got: %v", err)
		}
	})

	t.Run("DisableAPByMAC", func(t *testing.T) {
		err := service.DisableAPByMAC(ctx, "aa:bb:cc:dd:ee:ff")
		if err != nil {
			t.Errorf("Expected no error for DisableAPByMAC, got: %v", err)
		}
	})

	t.Run("EnableAPByName", func(t *testing.T) {
		err := service.EnableAPByName(ctx, "TEST-AP01")
		if err != nil {
			t.Errorf("Expected no error for EnableAPByName, got: %v", err)
		}
	})

	t.Run("DisableAPByName", func(t *testing.T) {
		err := service.DisableAPByName(ctx, "TEST-AP01")
		if err != nil {
			t.Errorf("Expected no error for DisableAPByName, got: %v", err)
		}
	})

	// Test radio state operations
	t.Run("EnableRadioByMAC", func(t *testing.T) {
		err := service.EnableRadioByMAC(ctx, "aa:bb:cc:dd:ee:ff", 0, ap.RadioType80211BG)
		if err != nil {
			t.Errorf("Expected no error for EnableRadioByMAC, got: %v", err)
		}
	})

	t.Run("DisableRadioByMAC", func(t *testing.T) {
		err := service.DisableRadioByMAC(ctx, "aa:bb:cc:dd:ee:ff", 1, ap.RadioType80211A)
		if err != nil {
			t.Errorf("Expected no error for DisableRadioByMAC, got: %v", err)
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
	t.Run("ReloadByMAC", func(t *testing.T) {
		err := service.ReloadByMAC(ctx, "aa:bb:cc:dd:ee:ff")
		if err != nil {
			t.Errorf("Expected no error for ReloadByMAC, got: %v", err)
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
	t.Run("EnableAPByMAC_InvalidMAC", func(t *testing.T) {
		err := service.EnableAPByMAC(ctx, "invalid-mac")
		if err == nil {
			t.Error("Expected error for invalid MAC address, got nil")
		}
	})

	t.Run("DisableAPByMAC_InvalidMAC", func(t *testing.T) {
		err := service.DisableAPByMAC(ctx, "invalid-mac")
		if err == nil {
			t.Error("Expected error for invalid MAC address, got nil")
		}
	})

	t.Run("EnableAPByName_BlankName", func(t *testing.T) {
		err := service.EnableAPByName(ctx, "")
		if !errors.Is(err, core.ErrResourceNotFound) {
			t.Errorf("Expected core.ErrResourceNotFound for a blank AP name, got: %v", err)
		}
	})

	t.Run("DisableAPByName_BlankName", func(t *testing.T) {
		err := service.DisableAPByName(ctx, " ")
		if !errors.Is(err, core.ErrResourceNotFound) {
			t.Errorf("Expected core.ErrResourceNotFound for a blank AP name, got: %v", err)
		}
	})

	// Test invalid MAC validation for radio operations
	t.Run("EnableRadioByMAC_InvalidMAC", func(t *testing.T) {
		err := service.EnableRadioByMAC(ctx, "invalid-mac", 0, ap.RadioType80211BG)
		if err == nil {
			t.Error("Expected error for invalid MAC address, got nil")
		}
	})

	t.Run("DisableRadioByMAC_InvalidMAC", func(t *testing.T) {
		err := service.DisableRadioByMAC(ctx, "invalid-mac", 1, ap.RadioType80211A)
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
	t.Run("ReloadByMAC_InvalidMAC", func(t *testing.T) {
		err := service.ReloadByMAC(ctx, "invalid-mac")
		if err == nil {
			t.Error("Expected error for invalid MAC address, got nil")
		}
	})

	t.Run("ReloadByName_BlankName", func(t *testing.T) {
		err := service.ReloadByName(ctx, "\t")
		if !errors.Is(err, core.ErrResourceNotFound) {
			t.Errorf("Expected core.ErrResourceNotFound for a blank AP name, got: %v", err)
		}
	})
}

// TestApServiceUnit_DoOperations_MockSuccess tests edge cases and error branches.
func TestApServiceUnit_DoOperations_MockSuccess(t *testing.T) {
	// Create mock server with specific responses for edge cases
	responses := map[string]string{
		"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/capwap-data=aa:bb:cc:dd:ee:ff": `{
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
	t.Run("ReloadByMAC_EmptyCAPWAPData", func(t *testing.T) {
		err := service.ReloadByMAC(ctx, "aa:bb:cc:dd:ee:ff")
		if err == nil {
			t.Error("Expected error for AP not found in CAPWAP data, got nil")
		}
	})
}

// TestApServiceUnit_DoOperations_ErrorHandling tests nil CAPWAP data handling.
func TestApServiceUnit_DoOperations_ErrorHandling(t *testing.T) {
	mockServer := testutil.NewMockServer(testutil.WithErrorResponses(
		[]string{"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/capwap-data=aa:bb:cc:dd:ee:ff"},
		500,
	))
	defer mockServer.Close()

	testClient := testutil.NewTestClient(mockServer)
	service := ap.NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	// Test reload with failed CAPWAP data retrieval
	t.Run("ReloadByMAC_FailedCAPWAPDataRetrieval", func(t *testing.T) {
		err := service.ReloadByMAC(ctx, "aa:bb:cc:dd:ee:ff")
		if err == nil {
			t.Error("Expected error for failed CAPWAP data retrieval, got nil")
		}
	})
}

// TestApServiceUnit_Reload_EdgeCases tests specific edge cases for Reload function to achieve 100% coverage.
func TestApServiceUnit_Reload_EdgeCases(t *testing.T) {
	// Test Reload with nil CAPWAP data response
	t.Run("ReloadByMAC_NilCAPWAPResponse", func(t *testing.T) {
		responses := map[string]string{
			"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/capwap-data=aa:bb:cc:dd:ee:ff": `null`,
		}
		mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(responses))
		defer mockServer.Close()

		testClient := testutil.NewTestClient(mockServer)
		service := ap.NewService(testClient.Core().(*core.Client))
		ctx := testutil.TestContext(t)

		err := service.ReloadByMAC(ctx, "aa:bb:cc:dd:ee:ff")
		if err == nil {
			t.Error("Expected error for nil CAPWAP response, got nil")
		}
	})

	// The record IS the one this address keys, so nothing here exercises an absence: the mock
	// serves no ap-reset node, and the refusal comes from the write the resolve leads to. The
	// genuine absence is covered by ReloadByMAC_EmptyCAPWAPData and by the 404 arm.
	t.Run("ReloadByMAC_ResetNotServed", func(t *testing.T) {
		responses := map[string]string{
			"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/capwap-data=aa:bb:cc:dd:ee:ff": `{
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

		err := service.ReloadByMAC(ctx, "aa:bb:cc:dd:ee:ff")
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
			},
		),
		testutil.WithCustomResponse(
			"Cisco-IOS-XE-wireless-access-point-cfg-rpc:set-ap-slot-admin-state", testutil.ResponseConfig{
				StatusCode: 400,
				Body:       "Invalid request",
			},
		),
		testutil.WithCustomResponse(
			"Cisco-IOS-XE-wireless-ap-cfg:ap-cfg-data/ap-tags/ap-tag=aa:bb:cc:dd:ee:ff", testutil.ResponseConfig{
				StatusCode: 400,
				Body:       "Invalid request",
			},
		),
	)
	defer mockServer.Close()

	testClient := testutil.NewTestClient(mockServer)
	service := ap.NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	// Test updateAPState error handling
	t.Run("UpdateAPState_RPCError", func(t *testing.T) {
		err := service.EnableAPByMAC(ctx, "aa:bb:cc:dd:ee:ff")
		if err == nil {
			t.Error("Expected error for failed RPC call, got nil")
		}
	})

	// Test updateRadioState error handling
	t.Run("UpdateRadioState_RPCError", func(t *testing.T) {
		err := service.EnableRadioByMAC(ctx, "aa:bb:cc:dd:ee:ff", 0, ap.RadioType80211BG)
		if err == nil {
			t.Error("Expected error for failed radio RPC call, got nil")
		}
	})

	// Test updateRadioState with a radio type the RPC has no band number for
	t.Run("UpdateRadioState_UnnumberedRadioType", func(t *testing.T) {
		err := service.EnableRadioByMAC(ctx, "aa:bb:cc:dd:ee:ff", 0, ap.RadioTypeUWB)
		if err == nil {
			t.Error("Expected error for an unnumbered radio type, got nil")
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

// apTagNode is the node both the tag read and the tag write end at, which is what the mock
// server matches a handler on. The list key the URL appends is not part of the match.
const apTagNode = "Cisco-IOS-XE-wireless-ap-cfg:ap-cfg-data/ap-tags/ap-tag"

// slotAdminNode is the node the radio admin RPC posts to.
const slotAdminNode = "Cisco-IOS-XE-wireless-access-point-cfg-rpc:set-ap-slot-admin-state"

// apAdminNode is the node the AP admin RPC posts to.
const apAdminNode = "Cisco-IOS-XE-wireless-access-point-cfg-rpc:set-ap-admin-state"

// capwapResetNode is the node the CAPWAP reset RPC posts to.
const capwapResetNode = "Cisco-IOS-XE-wireless-access-point-cmd-rpc:set-rad-capwap-reset"

// writtenBody returns the body of the first recorded request made with method, and fails if
// none was. The read that precedes a write is recorded too, so the write is selected by
// method rather than by position.
func writtenBody(t *testing.T, server *testutil.RESTCONFServer, method string) string {
	t.Helper()

	for _, recorded := range server.Requests() {
		if recorded.Method == method {
			return recorded.Body
		}
	}

	t.Fatalf("no %s reached the server", method)

	return ""
}

// decodeWrittenTag decodes the body of the write a tag assignment sent into payload.
func decodeWrittenTag(t *testing.T, server *testutil.RESTCONFServer, payload any) {
	t.Helper()

	if err := json.Unmarshal([]byte(writtenBody(t, server, http.MethodPut)), payload); err != nil {
		t.Fatalf("Failed to decode the recorded write: %v", err)
	}
}

// writtenTags returns the three tags the recorded write carried.
func writtenTags(t *testing.T, server *testutil.RESTCONFServer) (siteTag, policyTag, rfTag string) {
	t.Helper()

	var payload struct {
		ApTag struct {
			SiteTag   string `json:"site-tag"`
			PolicyTag string `json:"policy-tag"`
			RFTag     string `json:"rf-tag"`
		} `json:"Cisco-IOS-XE-wireless-ap-cfg:ap-tag"`
	}
	decodeWrittenTag(t, server, &payload)

	return payload.ApTag.SiteTag, payload.ApTag.PolicyTag, payload.ApTag.RFTag
}

// writtenPrimingProfile returns the priming profile the recorded write carried.
func writtenPrimingProfile(t *testing.T, server *testutil.RESTCONFServer) string {
	t.Helper()

	var payload struct {
		ApTag struct {
			PrimingProfile string `json:"priming-profile"`
		} `json:"Cisco-IOS-XE-wireless-ap-cfg:ap-tag"`
	}
	decodeWrittenTag(t, server, &payload)

	return payload.ApTag.PrimingProfile
}

// writtenRPCInput returns the leaves the recorded RPC input carried, keyed by leaf name. The
// input is decoded into a map so a leaf the caller never named is visible: a mandatory choice
// is answered by the arm that is PRESENT, and a struct with a missing omitempty puts the other
// arm on the wire at its zero value, which an assertion on the wanted arm alone cannot see.
func writtenRPCInput(t *testing.T, server *testutil.RESTCONFServer) map[string]any {
	t.Helper()

	var payload map[string]map[string]any
	if err := json.Unmarshal([]byte(writtenBody(t, server, http.MethodPost)), &payload); err != nil {
		t.Fatalf("Failed to decode the recorded RPC: %v", err)
	}
	if len(payload) != 1 {
		t.Fatalf("the RPC body carried %d top-level keys, want 1", len(payload))
	}

	for _, input := range payload {
		return input
	}

	return nil
}

// assertRPCInputLeaves fails unless the recorded RPC input carried exactly want.
func assertRPCInputLeaves(t *testing.T, server *testutil.RESTCONFServer, want map[string]any) {
	t.Helper()

	got := writtenRPCInput(t, server)
	for leaf, value := range want {
		if got[leaf] != value {
			t.Errorf("input[%q] = %v, want %v", leaf, got[leaf], value)
		}
	}
	for leaf := range got {
		if _, ok := want[leaf]; !ok {
			t.Errorf("input carried %q = %v, which the caller never named", leaf, got[leaf])
		}
	}
}

// newTagRecorderService answers a tag read with entry and lets the server record the write that
// follows. An empty entry makes the read a 404, which is the AP-has-no-entry case.
func newTagRecorderService(t *testing.T, entry string) (ap.Service, *testutil.RESTCONFServer) {
	t.Helper()

	server := testutil.NewRESTCONFServer(t)
	t.Cleanup(server.Close)

	if entry != "" {
		server.AddHandler(http.MethodGet, apTagNode, func() (int, string) {
			return http.StatusOK, entry
		})
	}
	server.AddHandler(http.MethodPut, apTagNode, func() (int, string) {
		return http.StatusNoContent, ""
	})

	testClient := testutil.NewTestClient(testutil.NewMockServerFromHTTP(server.Server))

	return ap.NewService(testClient.Core().(*core.Client)), server
}

// newRPCService answers node with 204 and records what reached it.
func newRPCService(t *testing.T, node string) (ap.Service, *testutil.RESTCONFServer) {
	t.Helper()

	server := testutil.NewRESTCONFServer(t)
	t.Cleanup(server.Close)
	server.AddHandler(http.MethodPost, node, func() (int, string) {
		return http.StatusNoContent, ""
	})

	testClient := testutil.NewTestClient(testutil.NewMockServerFromHTTP(server.Server))

	return ap.NewService(testClient.Core().(*core.Client)), server
}

const testAPTagEntry = `{
	"Cisco-IOS-XE-wireless-ap-cfg:ap-tag": [{
		"ap-mac": "aa:bb:cc:dd:ee:ff",
		"site-tag": "existing-site",
		"policy-tag": "existing-policy",
		"rf-tag": "existing-rf"
	}]
}`

// testAPTagEntryAtDefault is the entry of an AP whose policy and RF tags hold their default, which
// the controller omits from the read.
const testAPTagEntryAtDefault = `{
	"Cisco-IOS-XE-wireless-ap-cfg:ap-tag": [{
		"ap-mac": "aa:bb:cc:dd:ee:ff",
		"site-tag": "existing-site"
	}]
}`

// TestApServiceUnit_AssignTags_MergeSemantics tests that a tag assignment preserves the tags
// the caller did not name, because the write replaces the whole entry.
func TestApServiceUnit_AssignTags_MergeSemantics(t *testing.T) {
	const apMAC = "aa:bb:cc:dd:ee:ff"

	tests := []struct {
		name              string
		entry             string
		assign            func(ctx context.Context, s ap.Service) error
		expectedSiteTag   string
		expectedPolicyTag string
		expectedRFTag     string
	}{
		{
			name:  "site tag assignment keeps the other two",
			entry: testAPTagEntry,
			assign: func(ctx context.Context, s ap.Service) error {
				return s.AssignSiteTag(ctx, apMAC, "new-site")
			},
			expectedSiteTag:   "new-site",
			expectedPolicyTag: "existing-policy",
			expectedRFTag:     "existing-rf",
		},
		{
			name:  "policy tag assignment keeps the other two",
			entry: testAPTagEntry,
			assign: func(ctx context.Context, s ap.Service) error {
				return s.AssignPolicyTag(ctx, apMAC, "new-policy")
			},
			expectedSiteTag:   "existing-site",
			expectedPolicyTag: "new-policy",
			expectedRFTag:     "existing-rf",
		},
		{
			name:  "RF tag assignment keeps the other two",
			entry: testAPTagEntry,
			assign: func(ctx context.Context, s ap.Service) error {
				return s.AssignRFTag(ctx, apMAC, "new-rf")
			},
			expectedSiteTag:   "existing-site",
			expectedPolicyTag: "existing-policy",
			expectedRFTag:     "new-rf",
		},
		{
			name:  "a tag the read omits falls back to its default, not an empty leaf",
			entry: testAPTagEntryAtDefault,
			assign: func(ctx context.Context, s ap.Service) error {
				return s.AssignPolicyTag(ctx, apMAC, "new-policy")
			},
			expectedSiteTag:   "existing-site",
			expectedPolicyTag: "new-policy",
			expectedRFTag:     "default-rf-tag",
		},
		{
			name:  "no entry falls back to the defaults",
			entry: "",
			assign: func(ctx context.Context, s ap.Service) error {
				return s.AssignSiteTag(ctx, apMAC, "new-site")
			},
			expectedSiteTag:   "new-site",
			expectedPolicyTag: "default-policy-tag",
			expectedRFTag:     "default-rf-tag",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service, server := newTagRecorderService(t, tt.entry)

			if err := tt.assign(context.Background(), service); err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			siteTag, policyTag, rfTag := writtenTags(t, server)
			if siteTag != tt.expectedSiteTag {
				t.Errorf("site-tag = %q, want %q", siteTag, tt.expectedSiteTag)
			}
			if policyTag != tt.expectedPolicyTag {
				t.Errorf("policy-tag = %q, want %q", policyTag, tt.expectedPolicyTag)
			}
			if rfTag != tt.expectedRFTag {
				t.Errorf("rf-tag = %q, want %q", rfTag, tt.expectedRFTag)
			}
		})
	}
}

// TestApServiceUnit_AssignTags_KeepsPrimingProfile tests that the replacing write carries a
// priming profile the caller never named, which the write body did not declare before.
func TestApServiceUnit_AssignTags_KeepsPrimingProfile(t *testing.T) {
	const apMAC = "aa:bb:cc:dd:ee:ff"
	const entry = `{
		"Cisco-IOS-XE-wireless-ap-cfg:ap-tag": [{
			"ap-mac": "aa:bb:cc:dd:ee:ff",
			"site-tag": "existing-site",
			"policy-tag": "existing-policy",
			"rf-tag": "existing-rf",
			"priming-profile": "existing-priming"
		}]
	}`

	t.Run("an existing priming profile survives", func(t *testing.T) {
		service, server := newTagRecorderService(t, entry)

		if err := service.AssignSiteTag(context.Background(), apMAC, "new-site"); err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		if got := writtenPrimingProfile(t, server); got != "existing-priming" {
			t.Errorf("priming-profile = %q, want %q", got, "existing-priming")
		}
	})

	t.Run("no priming profile is invented", func(t *testing.T) {
		service, server := newTagRecorderService(t, testAPTagEntry)

		if err := service.AssignSiteTag(context.Background(), apMAC, "new-site"); err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		if got := writtenPrimingProfile(t, server); got != "" {
			t.Errorf("priming-profile = %q, want it absent", got)
		}
	})
}

// TestApServiceUnit_SetRadioAdminState_RPCInput_Success pins every leaf the radio-admin RPC puts
// on the wire for each arm, and the band each radio type maps to. The band is what the previous
// shape derived from the slot: an XOR radio in slot 0 takes 3, not the 1 its served band implies.
func TestApServiceUnit_SetRadioAdminState_RPCInput_Success(t *testing.T) {
	tests := []struct {
		name      string
		radioType ap.RadioType
		slotID    int
		wantBand  string
	}{
		{name: "dedicated 2.4 GHz", radioType: ap.RadioType80211BG, slotID: 0, wantBand: "1"},
		{name: "dedicated 5 GHz", radioType: ap.RadioType80211A, slotID: 1, wantBand: "2"},
		{name: "2.4/5 GHz XOR in slot 0", radioType: ap.RadioType80211ABGN, slotID: 0, wantBand: "3"},
		{name: "5/6 GHz XOR in slot 2", radioType: ap.RadioTypeXOR5And6GHz, slotID: 2, wantBand: "3"},
		{name: "dedicated 6 GHz in slot 3", radioType: ap.RadioType6GHz, slotID: 3, wantBand: "4"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service, server := newRPCService(t, slotAdminNode)
			err := service.DisableRadioByMAC(t.Context(), "aa:bb:cc:dd:ee:ff", tt.slotID, tt.radioType)
			if err != nil {
				t.Fatalf("DisableRadioByMAC() error = %v", err)
			}

			assertRPCInputLeaves(t, server, map[string]any{
				"mode":     "admin-state-disabled",
				"slot-id":  float64(tt.slotID),
				"band":     tt.wantBand,
				"mac-addr": "aa:bb:cc:dd:ee:ff",
			})
		})
	}

	t.Run("ByName", func(t *testing.T) {
		service, server := newRPCService(t, slotAdminNode)
		err := service.EnableRadioByName(t.Context(), "TEST-AP01", 2, ap.RadioTypeXOR5And6GHz)
		if err != nil {
			t.Fatalf("EnableRadioByName() error = %v", err)
		}

		assertRPCInputLeaves(t, server, map[string]any{
			"mode":    "admin-state-enabled",
			"slot-id": float64(2),
			"band":    "3",
			"ap-name": "TEST-AP01",
		})
	})
}

// TestApServiceUnit_RadioBandNumber_UnnumberedTypes_Error holds the four members the RPC's 1..4
// band domain has no number for, so a table row added without a measurement fails here.
func TestApServiceUnit_RadioBandNumber_UnnumberedTypes_Error(t *testing.T) {
	unnumbered := []ap.RadioType{
		ap.RadioTypeInvalid,
		ap.RadioTypeUWB,
		ap.RadioTypeRemoteLAN,
		ap.RadioTypeXOR24And6GHz,
	}

	for _, radioType := range unnumbered {
		t.Run(string(radioType), func(t *testing.T) {
			service, _ := newRPCService(t, slotAdminNode)
			if err := service.EnableRadioByMAC(t.Context(), "aa:bb:cc:dd:ee:ff", 0, radioType); err == nil {
				t.Errorf("EnableRadioByMAC(%s) error = nil, want a refusal", radioType)
			}
		})
	}
}

// TestApServiceUnit_SetAPAdminState_RPCInput_Success pins every leaf the admin-state RPC puts on
// the wire for each arm, so an arm the caller never named is a failure here rather than a 400.
func TestApServiceUnit_SetAPAdminState_RPCInput_Success(t *testing.T) {
	t.Run("ByMAC", func(t *testing.T) {
		service, server := newRPCService(t, apAdminNode)
		if err := service.DisableAPByMAC(t.Context(), "AA-BB-CC-DD-EE-FF"); err != nil {
			t.Fatalf("DisableAPByMAC() error = %v", err)
		}

		assertRPCInputLeaves(t, server, map[string]any{
			"mode":     "admin-state-disabled",
			"mac-addr": "aa:bb:cc:dd:ee:ff",
		})
	})

	t.Run("ByName", func(t *testing.T) {
		service, server := newRPCService(t, apAdminNode)
		if err := service.EnableAPByName(t.Context(), "TEST-AP01"); err != nil {
			t.Fatalf("EnableAPByName() error = %v", err)
		}

		assertRPCInputLeaves(t, server, map[string]any{
			"mode":    "admin-state-enabled",
			"ap-name": "TEST-AP01",
		})
	})
}

// TestApServiceUnit_ResetCAPWAP_RPCInput_Success pins every leaf the CAPWAP-reset RPC puts on the
// wire for each arm. The input's choice is mandatory, so the arm the caller did not name must be
// absent rather than present at an empty string.
func TestApServiceUnit_ResetCAPWAP_RPCInput_Success(t *testing.T) {
	t.Run("ByName", func(t *testing.T) {
		service, server := newRPCService(t, capwapResetNode)
		if err := service.ResetCAPWAPByName(t.Context(), "TEST-AP01"); err != nil {
			t.Fatalf("ResetCAPWAPByName() error = %v", err)
		}

		assertRPCInputLeaves(t, server, map[string]any{"ap-name": "TEST-AP01"})
	})

	t.Run("ByMAC", func(t *testing.T) {
		service, server := newRPCService(t, capwapResetNode)
		if err := service.ResetCAPWAPByMAC(t.Context(), "AABBCCDDEEFF"); err != nil {
			t.Fatalf("ResetCAPWAPByMAC() error = %v", err)
		}

		assertRPCInputLeaves(t, server, map[string]any{"mac-addr": "aa:bb:cc:dd:ee:ff"})
	})

	t.Run("BlankArguments", func(t *testing.T) {
		service, _ := newRPCService(t, capwapResetNode)
		if err := service.ResetCAPWAPByName(t.Context(), " "); !errors.Is(err, core.ErrResourceNotFound) {
			t.Errorf("ResetCAPWAPByName(blank) error = %v, want core.ErrResourceNotFound", err)
		}
		if err := service.ResetCAPWAPByMAC(t.Context(), ""); !errors.Is(err, core.ErrResourceNotFound) {
			t.Errorf("ResetCAPWAPByMAC(blank) error = %v, want core.ErrResourceNotFound", err)
		}
	})
}

func TestApServiceUnit_StateVocabularies_MockSuccess(t *testing.T) {
	t.Parallel()

	body := `{
		"Cisco-IOS-XE-wireless-access-point-oper:capwap-data": [
			{
				"wtp-mac": "aa:bb:cc:dd:ee:ff",
				"ap-state": {"ap-admin-state": "adminstate-disabled", "ap-operation-state": "registered"},
				"ap-mode-data": {"wtp-mode": "local-mode", "ap-sub-mode": "not-configured"}
			}
		]
	}`

	mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(map[string]string{
		"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/capwap-data": body,
	}))
	defer mockServer.Close()

	testClient := testutil.NewTestClient(mockServer)
	service := ap.NewService(testClient.Core().(*core.Client))

	result, err := service.ListCAPWAPData(testutil.TestContext(t))
	if err != nil {
		t.Fatalf("ListCAPWAPData returned unexpected error: %v", err)
	}
	if len(result.CAPWAPData) != 1 {
		t.Fatalf("decoded %d records, want 1", len(result.CAPWAPData))
	}

	record := result.CAPWAPData[0]

	if record.ApModeData.ApSubMode != "not-configured" {
		t.Fatalf("ap-sub-mode = %q, so ap-mode-data was not decoded", record.ApModeData.ApSubMode)
	}
	if record.ApState.ApAdminState != ap.APAdminStateDisabled {
		t.Errorf("ap-admin-state = %q, want %q", record.ApState.ApAdminState, ap.APAdminStateDisabled)
	}
	if record.ApModeData.WtpMode != ap.WtpModeLocal {
		t.Errorf("wtp-mode = %q, want %q", record.ApModeData.WtpMode, ap.WtpModeLocal)
	}
}
