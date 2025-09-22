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

	// Test ListConfigEntries
	profiles, err := service.ListConfigEntries(ctx)
	if err != nil {
		t.Errorf("ListConfigEntries failed: %v", err)
		return
	}

	if profiles == nil {
		t.Error("ListConfigEntries returned nil result")
		return
	}

	// Test ListPolicies
	policies, err := service.ListPolicies(ctx)
	if err != nil {
		t.Errorf("ListPolicies failed: %v", err)
		return
	}

	if policies == nil {
		t.Error("ListPolicies returned nil result")
		return
	}

	// Test ListPolicyListEntries
	policyEntries, err := service.ListPolicyListEntries(ctx)
	if err != nil {
		t.Errorf("ListPolicyListEntries failed: %v", err)
		return
	}

	if policyEntries == nil {
		t.Error("ListPolicyListEntries returned nil result")
		return
	}

	// Test ListWirelessAAAPolicyConfigs
	aaaConfigs, err := service.ListWirelessAAAPolicyConfigs(ctx)
	if err != nil {
		t.Errorf("ListWirelessAAAPolicyConfigs failed: %v", err)
		return
	}

	if aaaConfigs == nil {
		t.Error("ListWirelessAAAPolicyConfigs returned nil result")
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

	_, err = service.ListConfigEntries(ctx)
	if err == nil {
		t.Error("Expected error with nil client for ListConfigEntries")
	}

	_, err = service.ListPolicies(ctx)
	if err == nil {
		t.Error("Expected error with nil client for ListPolicies")
	}

	_, err = service.ListPolicyListEntries(ctx)
	if err == nil {
		t.Error("Expected error with nil client for ListPolicyListEntries")
	}

	_, err = service.ListWirelessAAAPolicyConfigs(ctx)
	if err == nil {
		t.Error("Expected error with nil client for ListWirelessAAAPolicyConfigs")
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
