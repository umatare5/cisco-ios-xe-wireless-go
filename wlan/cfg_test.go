// Package wlan provides WLAN configuration test functionality for the Cisco Wireless Network Controller API.
package wlan

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
	testutils "github.com/umatare5/cisco-ios-xe-wireless-go/internal/tests"
)

// WlanCfgTestDataCollector holds test data for WLAN configuration functions
type WlanCfgTestDataCollector struct {
	Data map[string]interface{} `json:"wlan_cfg_test_data"`
}

var wlanCfgTestDataCollector = WlanCfgTestDataCollector{
	Data: make(map[string]interface{}),
}

func runWlanCfgTestAndCollectData(t *testing.T, testName string, testFunc func() (interface{}, error)) {
	data, err := testFunc()
	if err != nil {
		t.Logf("%s returned error: %v", testName, err)
		wlanCfgTestDataCollector.Data[testName] = map[string]interface{}{
			"error":   err.Error(),
			"success": false,
		}
	} else {
		t.Logf("%s executed successfully", testName)
		wlanCfgTestDataCollector.Data[testName] = map[string]interface{}{
			"data":    data,
			"success": true,
		}
	}
}

// =============================================================================
// 1. UNIT TESTS (Structure/Type Validation & JSON Serialization/Deserialization)
// =============================================================================

// TestWlanConfigurationEndpoints tests that all WLAN endpoints are correctly defined
func TestWlanConfigurationEndpoints(t *testing.T) {
	expectedEndpoints := map[string]string{
		"WlanCfgEndpoint":                  "Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data",
		"WlanCfgEntriesEndpoint":           "Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data/wlan-cfg-entries",
		"WlanPoliciesEndpoint":             "Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data/wlan-policies",
		"PolicyListEntriesEndpoint":        "Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data/policy-list-entries",
		"WirelessAaaPolicyConfigsEndpoint": "Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data/wireless-aaa-policy-configs",
	}

	for name, expected := range expectedEndpoints {
		t.Run(name, func(t *testing.T) {
			switch name {
			case "WlanCfgEndpoint":
				if WlanCfgEndpoint != expected {
					t.Errorf("Expected %s = %s, got %s", name, expected, WlanCfgEndpoint)
				}
			case "WlanCfgEntriesEndpoint":
				if WlanCfgEntriesEndpoint != expected {
					t.Errorf("Expected %s = %s, got %s", name, expected, WlanCfgEntriesEndpoint)
				}
			case "WlanPoliciesEndpoint":
				if WlanPoliciesEndpoint != expected {
					t.Errorf("Expected %s = %s, got %s", name, expected, WlanPoliciesEndpoint)
				}
			case "PolicyListEntriesEndpoint":
				if PolicyListEntriesEndpoint != expected {
					t.Errorf("Expected %s = %s, got %s", name, expected, PolicyListEntriesEndpoint)
				}
			case "WirelessAaaPolicyConfigsEndpoint":
				if WirelessAaaPolicyConfigsEndpoint != expected {
					t.Errorf("Expected %s = %s, got %s", name, expected, WirelessAaaPolicyConfigsEndpoint)
				}
			}
		})
	}
}

// TestWlanCfgDataStructures tests the basic structure of WLAN configuration data types
func TestWlanCfgDataStructures(t *testing.T) {
	// Sample WLAN configuration data based on real WNC response structure
	sampleJSON := `{
		"Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data": {
			"wlan-cfg-entries": {
				"wlan-cfg-entry": [
					{
						"profile-name": "guest-wlan",
						"ssid": "Guest-Network",
						"wlan-id": 10,
						"admin-status": true,
						"broadcast-ssid": true,
						"security": {
							"ft-security": false,
							"ft-reassociation-timeout": 20,
							"pmf-assoc-comeback-timeout": 1,
							"pmf-sa-query-retry-timeout": 200
						},
						"no-auth": false,
						"wpa-wpa2": {
							"wpa-wpa2-type": "wpa2",
							"wpa-wpa2-passphrase": "GuestPass123"
						},
						"exclusionlist": [],
						"session-timeout": 3600,
						"idle-timeout": 300,
						"mobility": {
							"anchor": false
						}
					}
				]
			},
			"wlan-policies": {
				"wlan-policy": [
					{
						"policy-name": "default-policy",
						"central-switching": true,
						"central-dhcp": true,
						"aaa-policy-name": "default-aaa",
						"session-timeout": 7200
					}
				]
			},
			"policy-list-entries": {
				"policy-list-entry": [
					{
						"policy-name": "acl-policy",
						"acl-name": "guest-acl",
						"priority": 1,
						"action": "permit"
					}
				]
			},
			"wireless-aaa-policy-configs": {
				"wireless-aaa-policy-config": [
					{
						"policy-name": "default-aaa",
						"auth-method": "dot1x",
						"auth-server-group": "radius-group",
						"accounting": true,
						"accounting-server-group": "radius-group"
					}
				]
			}
		}
	}`

	// Test unmarshaling into WlanCfgResponse
	var wlanCfg WlanCfgResponse
	err := json.Unmarshal([]byte(sampleJSON), &wlanCfg)
	if err != nil {
		t.Fatalf("Failed to unmarshal WlanCfgResponse: %v", err)
	}

	cfgData := wlanCfg.CiscoIOSXEWirelessWlanCfgWlanCfgData

	// Test WLAN configuration entries
	wlanEntries := cfgData.WlanCfgEntries.WlanCfgEntry
	if len(wlanEntries) == 0 {
		t.Error("Expected at least one WLAN configuration entry")
	} else {
		entry := wlanEntries[0]
		if entry.ProfileName != "guest-wlan" {
			t.Errorf("Expected profile name 'guest-wlan', got '%s'", entry.ProfileName)
		}
		if entry.WlanID != 10 {
			t.Errorf("Expected WLAN ID 10, got %d", entry.WlanID)
		}
	}

	// Test WLAN policies
	wlanPolicies := cfgData.WlanPolicies.WlanPolicy
	if len(wlanPolicies) == 0 {
		t.Error("Expected at least one WLAN policy")
	}

	// Test policy list entries
	policyEntries := cfgData.PolicyListEntries.PolicyListEntry
	if len(policyEntries) == 0 {
		t.Error("Expected at least one policy list entry")
	}

	// Test wireless AAA policy configurations
	aaaPolicies := cfgData.WirelessAaaPolicyConfigs.WirelessAaaPolicyConfig
	if len(aaaPolicies) == 0 {
		t.Error("Expected at least one wireless AAA policy configuration")
	} else {
		aaaPolicy := aaaPolicies[0]
		if aaaPolicy.PolicyName != "default-aaa" {
			t.Errorf("Expected policy name 'default-aaa', got '%s'", aaaPolicy.PolicyName)
		}
	}

	// Test individual response structures
	sampleWlanEntriesJSON := `{
		"Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-entries": {
			"wlan-cfg-entry": [
				{
					"profile-name": "corporate-wlan",
					"ssid": "Corporate-Network",
					"wlan-id": 20,
					"admin-status": true,
					"broadcast-ssid": false,
					"security": {
						"ft-security": true,
						"ft-reassociation-timeout": 20,
						"pmf-assoc-comeback-timeout": 1,
						"pmf-sa-query-retry-timeout": 200
					},
					"no-auth": false,
					"exclusionlist": [],
					"session-timeout": 28800,
					"idle-timeout": 600
				}
			]
		}
	}`

	var wlanEntriesResp WlanCfgEntriesResponse
	err = json.Unmarshal([]byte(sampleWlanEntriesJSON), &wlanEntriesResp)
	if err != nil {
		t.Fatalf("Failed to unmarshal WlanCfgEntriesResponse: %v", err)
	}

	if len(wlanEntriesResp.WlanCfgEntries.WlanCfgEntry) == 0 {
		t.Error("Expected at least one WLAN entry in entries response")
	}

	_, err = json.Marshal(wlanCfg)
	if err != nil {
		t.Errorf("Failed to marshal WlanCfgResponse back to JSON: %v", err)
	}

	_, err = json.Marshal(wlanEntriesResp)
	if err != nil {
		t.Errorf("Failed to marshal WlanCfgEntriesResponse back to JSON: %v", err)
	}
}

// =============================================================================
// 2. TABLE-DRIVEN TEST PATTERNS
// =============================================================================

// Currently no table-driven tests specific to WLAN configuration

// =============================================================================
// 3. FAIL-FAST ERROR DETECTION TESTS
// =============================================================================

// Currently no fail-fast error detection tests specific to WLAN configuration

// =============================================================================
// 4. INTEGRATION TESTS (API Communication & Full Workflow Tests)
// =============================================================================

// TestWlanConfigurationFunctions tests all WLAN configuration functions with real WNC data collection
func TestWlanConfigurationFunctions(t *testing.T) {
	client := testutils.CreateTestClientFromEnv(t)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	t.Run("GetWlanCfg", func(t *testing.T) {
		runWlanCfgTestAndCollectData(t, "GetWlanCfg", func() (interface{}, error) {
			return GetWlanCfg(client, ctx)
		})
	})

	t.Run("GetWlanCfgEntries", func(t *testing.T) {
		runWlanCfgTestAndCollectData(t, "GetWlanCfgEntries", func() (interface{}, error) {
			return GetWlanCfgEntries(client, ctx)
		})
	})

	t.Run("GetWlanPolicies", func(t *testing.T) {
		runWlanCfgTestAndCollectData(t, "GetWlanPolicies", func() (interface{}, error) {
			return GetWlanPolicies(client, ctx)
		})
	})

	t.Run("GetPolicyListEntries", func(t *testing.T) {
		runWlanCfgTestAndCollectData(t, "GetPolicyListEntries", func() (interface{}, error) {
			return GetPolicyListEntries(client, ctx)
		})
	})

	t.Run("GetWirelessAaaPolicyConfigs", func(t *testing.T) {
		runWlanCfgTestAndCollectData(t, "GetWirelessAaaPolicyConfigs", func() (interface{}, error) {
			return GetWirelessAaaPolicyConfigs(client, ctx)
		})
	})

	// Save collected test data to file
	if len(wlanCfgTestDataCollector.Data) > 0 {
		if err := testutils.SaveTestDataToFile("wlan_cfg_test_data_collected.json", wlanCfgTestDataCollector.Data); err != nil {
			t.Logf("Warning: Could not save test data: %v", err)
		} else {
			t.Logf("Test data saved to %s/wlan_cfg_test_data_collected.json", testutils.TestDataDir)
		}
	}
}

// TestWlanGlobalOperationFunctions tests all WLAN global operation functions
func TestWlanGlobalOperationFunctions(t *testing.T) {
	client := testutils.CreateTestClientFromEnv(t)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	t.Run("GetWlanGlobalOper", func(t *testing.T) {
		resp, err := GetWlanGlobalOper(client, ctx)
		if err != nil {
			t.Logf("GetWlanGlobalOper returned error (may be expected): %v", err)
		} else {
			t.Log("GetWlanGlobalOper executed successfully")
			if resp != nil {
				t.Logf("Response contains WLAN global operation data")
			}
		}
	})

	t.Run("GetWlanGlobalOperWlanInfo", func(t *testing.T) {
		resp, err := GetWlanGlobalOperWlanInfo(client, ctx)
		if err != nil {
			t.Logf("GetWlanGlobalOperWlanInfo returned error (may be expected): %v", err)
		} else {
			t.Log("GetWlanGlobalOperWlanInfo executed successfully")
			if resp != nil {
				t.Logf("Response contains WLAN info data")
			}
		}
	})
}

// =============================================================================
// 5. ERROR HANDLING TESTS
// =============================================================================

// TestWlanCfgErrorHandling tests error handling for all configuration functions
func TestWlanCfgErrorHandling(t *testing.T) {
	ctx := context.Background()

	// Test GetWlanCfg with nil client
	_, err := GetWlanCfg(nil, ctx)
	if err == nil || err.Error() != "client is nil" {
		t.Errorf("Expected 'client is nil' error, got: %v", err)
	}

	// Test GetWlanCfgEntries with nil client
	_, err = GetWlanCfgEntries(nil, ctx)
	if err == nil || err.Error() != "client is nil" {
		t.Errorf("Expected 'client is nil' error, got: %v", err)
	}

	// Test GetWlanPolicies with nil client
	_, err = GetWlanPolicies(nil, ctx)
	if err == nil || err.Error() != "client is nil" {
		t.Errorf("Expected 'client is nil' error, got: %v", err)
	}

	// Test GetPolicyListEntries with nil client
	_, err = GetPolicyListEntries(nil, ctx)
	if err == nil || err.Error() != "client is nil" {
		t.Errorf("Expected 'client is nil' error, got: %v", err)
	}

	// Test GetWirelessAaaPolicyConfigs with nil client
	_, err = GetWirelessAaaPolicyConfigs(nil, ctx)
	if err == nil || err.Error() != "client is nil" {
		t.Errorf("Expected 'client is nil' error, got: %v", err)
	}
}

// =============================================================================
// 6. CONTEXT HANDLING TESTS
// =============================================================================

// TestWlanCfgContextHandling tests context handling for all configuration functions
func TestWlanCfgContextHandling(t *testing.T) {
	// Test GetWlanCfg with context handling
	testutils.TestContextHandling(t, func(ctx context.Context, client *wnc.Client) error {
		_, err := GetWlanCfg(client, ctx)
		return err
	})

	// Test GetWlanCfgEntries with context handling
	testutils.TestContextHandling(t, func(ctx context.Context, client *wnc.Client) error {
		_, err := GetWlanCfgEntries(client, ctx)
		return err
	})

	// Test GetWlanPolicies with context handling
	testutils.TestContextHandling(t, func(ctx context.Context, client *wnc.Client) error {
		_, err := GetWlanPolicies(client, ctx)
		return err
	})

	// Test GetPolicyListEntries with context handling
	testutils.TestContextHandling(t, func(ctx context.Context, client *wnc.Client) error {
		_, err := GetPolicyListEntries(client, ctx)
		return err
	})

	// Test GetWirelessAaaPolicyConfigs with context handling
	testutils.TestContextHandling(t, func(ctx context.Context, client *wnc.Client) error {
		_, err := GetWirelessAaaPolicyConfigs(client, ctx)
		return err
	})
}

// =============================================================================
// 7. SERVICE TESTS
// =============================================================================

func TestWLANServiceConfiguration(t *testing.T) {
	client := testutils.GetTestClient(t)
	if client == nil {
		t.Skip("Skipping service tests: no test client available")
	}

	ctx := context.Background()
	service := NewService(client.CoreClient())

	// Test all configuration service methods
	t.Run("Service_Cfg", func(t *testing.T) {
		testutils.TestServiceMethod(t, func() error {
			_, err := service.Cfg(ctx)
			return err
		})
	})

	t.Run("Service_CfgEntries", func(t *testing.T) {
		testutils.TestServiceMethod(t, func() error {
			_, err := service.CfgEntries(ctx)
			return err
		})
	})

	t.Run("Service_Policies", func(t *testing.T) {
		testutils.TestServiceMethod(t, func() error {
			_, err := service.Policies(ctx)
			return err
		})
	})

	t.Run("Service_PolicyListEntries", func(t *testing.T) {
		testutils.TestServiceMethod(t, func() error {
			_, err := service.PolicyListEntries(ctx)
			return err
		})
	})

	t.Run("Service_WirelessAaaPolicyConfigs", func(t *testing.T) {
		testutils.TestServiceMethod(t, func() error {
			_, err := service.WirelessAaaPolicyConfigs(ctx)
			return err
		})
	})

	t.Run("Service_GlobalOper", func(t *testing.T) {
		testutils.TestServiceMethod(t, func() error {
			_, err := service.GlobalOper(ctx)
			return err
		})
	})
}

// =============================================================================
// 8. OTHER TESTS
// =============================================================================

// Currently no other tests specific to WLAN configuration
