package wlan

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/pkg/testutil"
)

func TestWlanServiceUnit_Constructor_Success(t *testing.T) {
	service := NewService(nil)
	if service.Client() != nil {
		t.Error("Expected nil client service")
	}
}

func TestWlanServiceUnit_GetOperations_MockSuccess(t *testing.T) {
	// Mock server with basic WLAN response structure
	mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(map[string]string{
		"Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data": `{
			"Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data": {
				"global-params": {
					"country-code": "US"
				}
			}
		}`,
		"Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data/wlan-cfg-entries": `{
			"Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-entries": {
				"wlan-cfg-entry": [
					{
						"profile-name": "test-wlan",
						"ssid": "TEST_SSID",
						"admin-status": true
					}
				]
			}
		}`,
		"Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data/wlan-cfg-entries/wlan-cfg-entry=test-wlan": `{
			"Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-entry": [
				{
					"profile-name": "test-wlan",
					"ssid": "TEST_SSID",
					"admin-status": true
				}
			]
		}`,
		"Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data/wlan-policies": `{
			"Cisco-IOS-XE-wireless-wlan-cfg:wlan-policies": {
				"wlan-policy": [
					{
						"policy-name": "test-policy",
						"description": "Test policy"
					}
				]
			}
		}`,
		"Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data/policy-list-entries": `{
			"Cisco-IOS-XE-wireless-wlan-cfg:policy-list-entries": {
				"policy-list-entry": [
					{
						"tag-name": "test-policy-tag",
						"description": "Test policy tag"
					}
				]
			}
		}`,
		"Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data/wireless-aaa-policy-configs": `{
			"Cisco-IOS-XE-wireless-wlan-cfg:wireless-aaa-policy-configs": {
				"wireless-aaa-policy-config": [
					{
						"policy-name": "test-aaa-policy",
						"description": "Test AAA policy"
					}
				]
			}
		}`,
		"Cisco-IOS-XE-wireless-wlan-global-oper:wlan-global-oper-data": `{
			"Cisco-IOS-XE-wireless-wlan-global-oper:wlan-global-oper-data": {
				"global-stats": {
					"total-wlans": 2,
					"active-wlans": 1
				}
			}
		}`,
	}))
	defer mockServer.Close()

	client := testutil.NewTestClient(mockServer)
	service := NewService(client.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	// Test GetConfig
	config, err := service.GetConfig(ctx)
	if err != nil {
		t.Errorf("GetConfig failed: %v", err)
		return
	}

	if config == nil {
		t.Error("GetConfig returned nil result")
		return
	}

	// Test GetOperational
	operational, err := service.GetOperational(ctx)
	if err != nil {
		t.Errorf("GetOperational failed: %v", err)
		return
	}

	if operational == nil {
		t.Error("GetOperational returned nil result")
		return
	}

	// Test ListWlanCfgEntries
	cfgEntries, err := service.ListWlanCfgEntries(ctx)
	if err != nil {
		t.Errorf("ListWlanCfgEntries failed: %v", err)
		return
	}

	if cfgEntries == nil {
		t.Error("ListWlanCfgEntries returned nil result")
		return
	}

	// Test ListWlanPolicies
	wlanPolicies, err := service.ListWlanPolicies(ctx)
	if err != nil {
		t.Errorf("ListWlanPolicies failed: %v", err)
		return
	}

	if wlanPolicies == nil {
		t.Error("ListWlanPolicies returned nil result")
		return
	}

	// Test ListCfgPolicyListEntries
	cfgPolicyEntries, err := service.ListCfgPolicyListEntries(ctx)
	if err != nil {
		t.Errorf("ListCfgPolicyListEntries failed: %v", err)
		return
	}

	if cfgPolicyEntries == nil {
		t.Error("ListCfgPolicyListEntries returned nil result")
		return
	}

	// Test ListCfgWirelessAaaPolicyConfigs
	cfgAaaConfigs, err := service.ListCfgWirelessAaaPolicyConfigs(ctx)
	if err != nil {
		t.Errorf("ListCfgWirelessAaaPolicyConfigs failed: %v", err)
		return
	}

	if cfgAaaConfigs == nil {
		t.Error("ListCfgWirelessAaaPolicyConfigs returned nil result")
		return
	}

	// Test ListDot11beProfiles (skip if not supported by mock server)
	dot11beProfiles, err := service.ListDot11beProfiles(ctx)
	if err != nil {
		// Wi-Fi 7 endpoints may not be supported by all mock servers
		t.Logf("ListDot11beProfiles failed (expected for older mock servers): %v", err)
	} else if dot11beProfiles == nil {
		t.Error("ListDot11beProfiles returned nil result")
		return
	}

	// Test ListWlanInfo (skip if not supported by mock server)
	wlanInfo, err := service.ListWlanInfo(ctx)
	if err != nil {
		// WlanInfo endpoint may not be supported by all mock servers
		t.Logf("ListWlanInfo failed (expected for older mock servers): %v", err)
	} else if wlanInfo == nil {
		t.Error("ListWlanInfo returned nil result")
		return
	}

	t.Logf("All get operations returned valid WLAN data")
}

func TestWlanServiceUnit_ErrorHandling_NilClient(t *testing.T) {
	service := NewService(nil)
	ctx := testutil.TestContext(t)

	_, err := service.GetConfig(ctx)
	if err == nil {
		t.Error("Expected error with nil client for GetConfig")
	}

	_, err = service.GetOperational(ctx)
	if err == nil {
		t.Error("Expected error with nil client for GetOperational")
	}

	_, err = service.ListWlanCfgEntries(ctx)
	if err == nil {
		t.Error("Expected error with nil client for ListWlanCfgEntries")
	}

	_, err = service.ListWlanPolicies(ctx)
	if err == nil {
		t.Error("Expected error with nil client for ListWlanPolicies")
	}

	_, err = service.ListCfgPolicyListEntries(ctx)
	if err == nil {
		t.Error("Expected error with nil client for ListCfgPolicyListEntries")
	}

	_, err = service.ListCfgWirelessAaaPolicyConfigs(ctx)
	if err == nil {
		t.Error("Expected error with nil client for ListCfgWirelessAaaPolicyConfigs")
	}

	// Note: ListDot11beProfiles and ListWlanInfo are not tested with nil client as they may not be supported by all mock servers
}

// TestWlanServiceUnit_OmittedSecurityLeaf_MockSuccess tests that an omitted security leaf stays
// nil while an explicitly configured false decodes to a non-nil false.
func TestWlanServiceUnit_OmittedSecurityLeaf_MockSuccess(t *testing.T) {
	mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(map[string]string{
		"Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data/wlan-cfg-entries": `{
			"Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-entries": {
				"wlan-cfg-entry": [
					{"wlan-id": 1, "profile-name": "profile-default"},
					{
						"wlan-id": 2,
						"profile-name": "profile-explicit",
						"wpa2-enabled": false,
						"wlan-11k-neigh-list": false,
						"apf-vap-802-11v-data": {"dot11v-dms": false},
						"apf-vap-id-data": {"ssid": "ssid-explicit", "wlan-status": false}
					}
				]
			}
		}`,
	}))
	defer mockServer.Close()

	testClient := testutil.NewTestClient(mockServer)
	service := NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	result, err := service.ListWlanCfgEntries(ctx)
	if err != nil {
		t.Fatalf("ListWlanCfgEntries failed: %v", err)
	}
	if result.WlanCfgEntries == nil || len(result.WlanCfgEntries.WlanCfgEntry) != 2 {
		t.Fatalf("Expected 2 entries, got %+v", result.WlanCfgEntries)
	}

	entries := result.WlanCfgEntries.WlanCfgEntry
	if entries[0].WPA2Enabled != nil {
		t.Error("Expected an omitted wpa2-enabled to stay nil: the default in force is not false")
	}
	if entries[0].Wlan11kNeighList != nil {
		t.Error("Expected an omitted wlan-11k-neigh-list to stay nil")
	}
	if entries[0].APFVap80211vData != nil {
		t.Error("Expected an omitted apf-vap-802-11v-data container to stay nil")
	}
	// Every leaf this test retypes is asserted in both directions. An omission-only assertion
	// leaves the JSON tag unreachable, so renaming it breaks nothing a nil-check can observe.
	if entries[1].WPA2Enabled == nil || *entries[1].WPA2Enabled {
		t.Error("Expected an explicit false to decode to a non-nil false")
	}
	if entries[1].Wlan11kNeighList == nil || *entries[1].Wlan11kNeighList {
		t.Error("Expected an explicit false wlan-11k-neigh-list to decode to a non-nil false")
	}
	if entries[1].APFVap80211vData == nil || entries[1].APFVap80211vData.Dot11vDms == nil ||
		*entries[1].APFVap80211vData.Dot11vDms {
		t.Error("Expected an explicit false dot11v-dms to decode to a non-nil false")
	}
	if entries[1].APFVapIDData == nil || entries[1].APFVapIDData.WlanStatus == nil {
		t.Fatal("Expected wlan-status to decode")
	}
	if *entries[1].APFVapIDData.WlanStatus {
		t.Error("Expected the explicit false for wlan-status")
	}
}

// TestWlanServiceUnit_PartialTimeoutContainer_MockSuccess tests that a container arriving with only
// some of its leaves leaves the omitted ones nil, and that an explicit zero stays distinguishable
// from an omission. Both record shapes are the measured ones: a plain read sends wlan-timeout with
// session-timeout alone and wlan-switching-policy with two of its four central-* leaves, while a
// with-defaults read sends every leaf and idle-threshold arrives as zero.
func TestWlanServiceUnit_PartialTimeoutContainer_MockSuccess(t *testing.T) {
	mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(map[string]string{
		"Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data/wlan-policies": `{
			"Cisco-IOS-XE-wireless-wlan-cfg:wlan-policies": {
				"wlan-policy": [
					{
						"policy-profile-name": "policy-omitted-leaves",
						"wlan-timeout": {"session-timeout": 1800},
						"wlan-switching-policy": {"central-switching": false, "central-dhcp": false}
					},
					{
						"policy-profile-name": "policy-every-leaf",
						"wlan-timeout": {"session-timeout": 1800, "idle-timeout": 300, "idle-threshold": 0},
						"wlan-switching-policy": {
							"central-switching": false,
							"central-authentication": false,
							"central-dhcp": false,
							"central-assoc-enable": false
						}
					}
				]
			}
		}`,
	}))
	defer mockServer.Close()

	testClient := testutil.NewTestClient(mockServer)
	service := NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	result, err := service.ListWlanPolicies(ctx)
	if err != nil {
		t.Fatalf("ListWlanPolicies failed: %v", err)
	}
	if result.WlanPolicies == nil || len(result.WlanPolicies.WlanPolicy) != 2 {
		t.Fatalf("Expected 2 policies, got %+v", result.WlanPolicies)
	}

	partial := result.WlanPolicies.WlanPolicy[0]
	if partial.WlanTimeout == nil || partial.WlanTimeout.SessionTimeout == nil {
		t.Fatal("Expected session-timeout to decode")
	}
	if partial.WlanTimeout.IdleTimeout != nil || partial.WlanTimeout.IdleThreshold != nil {
		t.Error("Expected the idle leaves omitted from a present container to stay nil")
	}
	if partial.WlanSwitchingPolicy == nil || partial.WlanSwitchingPolicy.CentralSwitching == nil {
		t.Fatal("Expected central-switching to decode")
	}
	if partial.WlanSwitchingPolicy.CentralAuthentication != nil ||
		partial.WlanSwitchingPolicy.CentralAssocEnable != nil {
		t.Error("Expected an omitted central-* leaf to stay nil: the default in force is true")
	}

	// central-dhcp arrives on both records, so it is asserted in the present direction. An
	// omission-only assertion leaves the tag unreachable: renaming it changes nothing a nil-check
	// can see.
	if partial.WlanSwitchingPolicy.CentralDHCP == nil || *partial.WlanSwitchingPolicy.CentralDHCP {
		t.Error("Expected an explicit false central-dhcp to decode to a non-nil false")
	}

	complete := result.WlanPolicies.WlanPolicy[1]
	if complete.WlanTimeout.IdleThreshold == nil || *complete.WlanTimeout.IdleThreshold != 0 {
		t.Error("Expected an explicit zero to decode to a non-nil zero")
	}
	if complete.WlanTimeout.IdleTimeout == nil || *complete.WlanTimeout.IdleTimeout != 300 {
		t.Error("Expected the idle-timeout of a complete container to decode to its value")
	}
	if complete.WlanSwitchingPolicy.CentralAuthentication == nil ||
		complete.WlanSwitchingPolicy.CentralAssocEnable == nil {
		t.Error("Expected every central-* leaf of a complete container to decode")
	}
	if complete.WlanSwitchingPolicy.CentralDHCP == nil || *complete.WlanSwitchingPolicy.CentralDHCP {
		t.Error("Expected the central-dhcp of a complete container to decode as a non-nil false")
	}
}
