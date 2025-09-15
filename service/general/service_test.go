package general_test

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/pkg/testutil"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/general"
)

// TestGeneralServiceUnit_Constructor_Success tests service constructor functionality.
func TestGeneralServiceUnit_Constructor_Success(t *testing.T) {
	t.Run("NewServiceWithValidClient", func(t *testing.T) {
		// Create mock server and test client using public API
		responses := map[string]string{
			"test-endpoint": `{"status": "success"}`,
		}
		mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(responses))
		defer mockServer.Close()

		testClient := testutil.NewTestClient(mockServer)
		service := general.NewService(testClient.Core().(*core.Client))

		// Test that service can be created and has a client
		if service.Client() == nil {
			t.Error("Expected service to have a client")
		}
	})

	t.Run("NewServiceWithNilClient", func(t *testing.T) {
		service := general.NewService(nil)

		// Service should still be created even with nil client
		if service.Client() != nil {
			t.Error("Expected service with nil client to return nil from Client()")
		}
	})
}

// TestGeneralServiceUnit_GetOperations_MockSuccess tests Get operations using mock server
// This is essential for CI environments where actual Cisco controllers are not available.
func TestGeneralServiceUnit_GetOperations_MockSuccess(t *testing.T) {
	// Create mock RESTCONF server with General endpoints
	responses := map[string]string{
		"Cisco-IOS-XE-wireless-general-oper:general-oper-data": `{
			"Cisco-IOS-XE-wireless-general-oper:general-oper-data": {
				"general-summary": {
					"total-aps": 100,
					"enabled-interfaces": 4
				}
			}
		}`,
	}
	mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(responses))
	defer mockServer.Close()

	// Create test client configured for the mock server
	testClient := testutil.NewTestClient(mockServer)
	service := general.NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	// Test GetOperational operation
	result, err := service.GetOperational(ctx)
	if err != nil {
		t.Errorf("Expected no error for mock GetOperational, got: %v", err)
	}
	if result == nil {
		t.Error("Expected result for mock GetOperational, got nil")
	}
}

// TestGeneralServiceUnit_GetOperations_ErrorHandling tests error scenarios using mock server.
func TestGeneralServiceUnit_GetOperations_ErrorHandling(t *testing.T) {
	// Create mock server that returns 404 for General endpoints
	errorPaths := []string{
		"Cisco-IOS-XE-wireless-general-oper:general-oper-data",
	}
	mockServer := testutil.NewMockServer(testutil.WithErrorResponses(errorPaths, 404))
	defer mockServer.Close()

	testClient := testutil.NewTestClient(mockServer)
	service := general.NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	// Test that GetOperational properly handles 404 errors
	_, err := service.GetOperational(ctx)
	if err == nil {
		t.Error("Expected error for 404 response, got nil")
	}

	// Verify error contains expected information
	if !core.IsNotFoundError(err) {
		t.Errorf("Expected NotFound error, got: %v", err)
	}
}

// TestGeneralServiceUnit_GetConfigOperations_MockSuccess tests configuration Get operations using mock server.
func TestGeneralServiceUnit_GetConfigOperations_MockSuccess(t *testing.T) {
	t.Parallel()

	// Mock responses based on real WNC data structure for general configuration
	responses := map[string]string{
		"Cisco-IOS-XE-wireless-general-cfg:general-cfg-data": `{
			"Cisco-IOS-XE-wireless-general-cfg:general-cfg-data": {
				"mewlc-config": {},
				"cac-config": {},
				"mfp": {},
				"fips-cfg": {},
				"wsa-ap-client-event": {},
				"sim-l3-interface-cache-data": {
					"interface-name": "Vlan801"
				},
				"wlc-management-data": {
					"pki-trustpoint-name": "WNC1_WLC_TP"
				},
				"laginfo": {},
				"multicast-config": {
					"is-mdns-enabled": false
				},
				"feature-usage-cfg": {},
				"threshold-warn-cfg": {},
				"ap-loc-ranging-cfg": {},
				"geolocation-cfg": {}
			}
		}`,
		"Cisco-IOS-XE-wireless-general-oper:general-oper-data/mgmt-intf-data": `{
			"Cisco-IOS-XE-wireless-general-oper:mgmt-intf-data": {
				"intf-name": "Vlan801",
				"intf-type": "Management",
				"intf-id": 801,
				"mgmt-ip": "192.168.255.1",
				"net-mask": "255.255.255.0",
				"mgmt-mac": "00:1e:49:96:4c:ff"
			}
		}`,
		"Cisco-IOS-XE-wireless-general-cfg:general-cfg-data/mfp": `{
			"Cisco-IOS-XE-wireless-general-cfg:mfp": {
				"enable": false
			}
		}`,
		"Cisco-IOS-XE-wireless-general-cfg:general-cfg-data/laginfo": `{
			"Cisco-IOS-XE-wireless-general-cfg:laginfo": {
				"lag-enable": false,
				"lag-mode": "active"
			}
		}`,
		"Cisco-IOS-XE-wireless-general-cfg:general-cfg-data/cac-config": `{
			"Cisco-IOS-XE-wireless-general-cfg:cac-config": {
				"cac-enable": false
			}
		}`,
		"Cisco-IOS-XE-wireless-general-cfg:general-cfg-data/fips-cfg": `{
			"Cisco-IOS-XE-wireless-general-cfg:fips-cfg": {
				"fips-enable": false
			}
		}`,
		"Cisco-IOS-XE-wireless-general-cfg:general-cfg-data/feature-usage-cfg": `{
			"Cisco-IOS-XE-wireless-general-cfg:feature-usage-cfg": {
				"enable": false
			}
		}`,
		"Cisco-IOS-XE-wireless-general-cfg:general-cfg-data/mewlc-config": `{
			"Cisco-IOS-XE-wireless-general-cfg:mewlc-config": {
				"enable": false
			}
		}`,
		"Cisco-IOS-XE-wireless-general-cfg:general-cfg-data/multicast-config": `{
			"Cisco-IOS-XE-wireless-general-cfg:multicast-config": {
				"igmp-snooping": false
			}
		}`,
		"Cisco-IOS-XE-wireless-general-cfg:general-cfg-data/sim-l3-interface-cache-data": `{
			"Cisco-IOS-XE-wireless-general-cfg:sim-l3-interface-cache-data": {
				"enable": false
			}
		}`,
		"Cisco-IOS-XE-wireless-general-cfg:general-cfg-data/threshold-warn-cfg": `{
			"Cisco-IOS-XE-wireless-general-cfg:threshold-warn-cfg": {
				"enable": false
			}
		}`,
		"Cisco-IOS-XE-wireless-general-cfg:general-cfg-data/wlc-management-data": `{
			"Cisco-IOS-XE-wireless-general-cfg:wlc-management-data": {
				"country": "US"
			}
		}`,
		"Cisco-IOS-XE-wireless-general-cfg:general-cfg-data/wsa-ap-client-event": `{
			"Cisco-IOS-XE-wireless-general-cfg:wsa-ap-client-event": {
				"enable": false
			}
		}`,
		"Cisco-IOS-XE-wireless-general-cfg:general-cfg-data/ap-loc-ranging-cfg": `{
			"Cisco-IOS-XE-wireless-general-cfg:ap-loc-ranging-cfg": {
				"enable": false
			}
		}`,
		"Cisco-IOS-XE-wireless-general-cfg:general-cfg-data/geolocation-cfg": `{
			"Cisco-IOS-XE-wireless-general-cfg:geolocation-cfg": {
				"enable": false
			}
		}`,
	}

	mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(responses))
	defer mockServer.Close()

	testClient := testutil.NewTestClient(mockServer)
	service := general.NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	t.Run("GetConfig", func(t *testing.T) {
		result, err := service.GetConfig(ctx)
		if err != nil {
			t.Errorf("GetConfig returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("GetConfig returned nil result")
			return
		}

		// Validate that the main configuration structure is properly populated
		cfgData := result.CiscoIOSXEWirelessGeneralCfgGeneralCfgData

		// Check that all expected configuration sections exist
		if cfgData.SimL3InterfaceCacheData == nil {
			t.Error("Expected sim-l3-interface-cache-data to be present")
		}
		if cfgData.WlcManagementData == nil {
			t.Error("Expected wlc-management-data to be present")
		}
		if cfgData.MulticastConfig == nil {
			t.Error("Expected multicast-config to be present")
		}
	})

	t.Run("GetManagementInterfaceState", func(t *testing.T) {
		result, err := service.GetManagementInterfaceState(ctx)
		if err != nil {
			t.Errorf("GetManagementInterfaceState returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("GetManagementInterfaceState returned nil result")
		}
	})

	t.Run("GetMFPConfig", func(t *testing.T) {
		result, err := service.GetMFPConfig(ctx)
		if err != nil {
			t.Errorf("GetMFPConfig returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("GetMFPConfig returned nil result")
		}
	})

	t.Run("GetLAGInfo", func(t *testing.T) {
		result, err := service.GetLAGInfo(ctx)
		if err != nil {
			t.Errorf("GetLAGInfo returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("GetLAGInfo returned nil result")
		}
	})

	t.Run("GetCACConfig", func(t *testing.T) {
		result, err := service.GetCACConfig(ctx)
		if err != nil {
			t.Errorf("GetCACConfig returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("GetCACConfig returned nil result")
		}
	})

	t.Run("GetFIPSConfig", func(t *testing.T) {
		result, err := service.GetFIPSConfig(ctx)
		if err != nil {
			t.Errorf("GetFIPSConfig returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("GetFIPSConfig returned nil result")
		}
	})

	t.Run("GetFeatureUsageConfig", func(t *testing.T) {
		result, err := service.GetFeatureUsageConfig(ctx)
		if err != nil {
			t.Errorf("GetFeatureUsageConfig returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("GetFeatureUsageConfig returned nil result")
		}
	})

	t.Run("GetMEWLCConfig", func(t *testing.T) {
		result, err := service.GetMEWLCConfig(ctx)
		if err != nil {
			t.Errorf("GetMEWLCConfig returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("GetMEWLCConfig returned nil result")
		}
	})

	t.Run("GetMulticastConfig", func(t *testing.T) {
		result, err := service.GetMulticastConfig(ctx)
		if err != nil {
			t.Errorf("GetMulticastConfig returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("GetMulticastConfig returned nil result")
		}
	})

	t.Run("ListSIML3InterfaceCache", func(t *testing.T) {
		result, err := service.ListSIML3InterfaceCache(ctx)
		if err != nil {
			t.Errorf("ListSIML3InterfaceCache returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("ListSIML3InterfaceCache returned nil result")
		}
	})

	t.Run("GetThresholdWarningConfig", func(t *testing.T) {
		result, err := service.GetThresholdWarningConfig(ctx)
		if err != nil {
			t.Errorf("GetThresholdWarningConfig returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("GetThresholdWarningConfig returned nil result")
		}
	})

	t.Run("GetWLCManagementInfo", func(t *testing.T) {
		result, err := service.GetWLCManagementInfo(ctx)
		if err != nil {
			t.Errorf("GetWLCManagementInfo returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("GetWLCManagementInfo returned nil result")
		}
	})

	t.Run("GetWSAAPClientEventConfig", func(t *testing.T) {
		result, err := service.GetWSAAPClientEventConfig(ctx)
		if err != nil {
			t.Errorf("GetWSAAPClientEventConfig returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("GetWSAAPClientEventConfig returned nil result")
		}
	})

	t.Run("GetAPLocationRangingConfig", func(t *testing.T) {
		result, err := service.GetAPLocationRangingConfig(ctx)
		if err != nil {
			t.Errorf("GetAPLocationRangingConfig returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("GetAPLocationRangingConfig returned nil result")
		}
	})

	t.Run("GetGeolocationConfig", func(t *testing.T) {
		result, err := service.GetGeolocationConfig(ctx)
		if err != nil {
			t.Errorf("GetGeolocationConfig returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("GetGeolocationConfig returned nil result")
		}
	})
}

// TestGeneralServiceUnit_ErrorHandling_NilClient tests error handling with nil client.
func TestGeneralServiceUnit_ErrorHandling_NilClient(t *testing.T) {
	t.Parallel()

	t.Run("GetConfig_NilClient", func(t *testing.T) {
		service := general.NewService(nil)
		ctx := testutil.TestContext(t)

		result, err := service.GetConfig(ctx)
		if err == nil {
			t.Error("Expected error for nil client")
		}
		if result != nil {
			t.Error("Expected nil result for error case")
		}
	})

	t.Run("GetManagementInterfaceState_NilClient", func(t *testing.T) {
		service := general.NewService(nil)
		ctx := testutil.TestContext(t)

		result, err := service.GetManagementInterfaceState(ctx)
		if err == nil {
			t.Error("Expected error for nil client")
		}
		if result != nil {
			t.Error("Expected nil result for error case")
		}
	})

	t.Run("GetMFPConfig_NilClient", func(t *testing.T) {
		service := general.NewService(nil)
		ctx := testutil.TestContext(t)

		result, err := service.GetMFPConfig(ctx)
		if err == nil {
			t.Error("Expected error for nil client")
		}
		if result != nil {
			t.Error("Expected nil result for error case")
		}
	})

	t.Run("GetLAGInfo_NilClient", func(t *testing.T) {
		service := general.NewService(nil)
		ctx := testutil.TestContext(t)

		result, err := service.GetLAGInfo(ctx)
		if err == nil {
			t.Error("Expected error for nil client")
		}
		if result != nil {
			t.Error("Expected nil result for error case")
		}
	})
}
