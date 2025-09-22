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
		// Base configuration data
		"Cisco-IOS-XE-wireless-general-cfg:general-cfg-data": `{
			"Cisco-IOS-XE-wireless-general-cfg:general-cfg-data": {
				"mewlc-config": {"enable": true},
				"cac-config": {"voice-sip-bandwidth": 64},
				"mfp": {"client-protection": "optional"},
				"fips-cfg": {"enable": false},
				"wsa-ap-client-event": {"enable": true},
				"sim-l3-interface-cache-data": {"vlan-id": 100},
				"wlc-management-data": {"cert-type": "manufacturing"},
				"laginfo": {"lag-support": true},
				"multicast-config": {"multicast-mode": "unicast"},
				"feature-usage-cfg": {"enable": true},
				"threshold-warn-cfg": {"memory-threshold": 80},
				"ap-loc-ranging-cfg": {"enable": false},
				"geolocation-cfg": {"enable": false}
			}
		}`,

		// Base operational data
		"Cisco-IOS-XE-wireless-general-oper:general-oper-data": `{
			"Cisco-IOS-XE-wireless-general-oper:general-oper-data": {
				"mgmt-intf-data": {
					"intf-name": "GigabitEthernet0",
					"intf-type": "ethernet",
					"intf-id": 0,
					"mgmt-ip": "192.168.1.100",
					"net-mask": "255.255.255.0",
					"mgmt-mac": "aa:bb:cc:dd:ee:ff"
				}
			}
		}`,

		// Individual configuration endpoints
		"Cisco-IOS-XE-wireless-general-cfg:general-cfg-data/mewlc-config": `{
			"Cisco-IOS-XE-wireless-general-cfg:mewlc-config": {"enable": true}
		}`,

		"Cisco-IOS-XE-wireless-general-cfg:general-cfg-data/cac-config": `{
			"Cisco-IOS-XE-wireless-general-cfg:cac-config": {"voice-sip-bandwidth": 64}
		}`,

		"Cisco-IOS-XE-wireless-general-cfg:general-cfg-data/mfp": `{
			"Cisco-IOS-XE-wireless-general-cfg:mfp": {"client-protection": "optional"}
		}`,

		"Cisco-IOS-XE-wireless-general-cfg:general-cfg-data/fips-cfg": `{
			"Cisco-IOS-XE-wireless-general-cfg:fips-cfg": {"enable": false}
		}`,

		"Cisco-IOS-XE-wireless-general-cfg:general-cfg-data/wsa-ap-client-event": `{
			"Cisco-IOS-XE-wireless-general-cfg:wsa-ap-client-event": {"enable": true}
		}`,

		"Cisco-IOS-XE-wireless-general-cfg:general-cfg-data/sim-l3-interface-cache-data": `{
			"Cisco-IOS-XE-wireless-general-cfg:sim-l3-interface-cache-data": {"vlan-id": 100}
		}`,

		"Cisco-IOS-XE-wireless-general-cfg:general-cfg-data/wlc-management-data": `{
			"Cisco-IOS-XE-wireless-general-cfg:wlc-management-data": {"cert-type": "manufacturing"}
		}`,

		"Cisco-IOS-XE-wireless-general-cfg:general-cfg-data/laginfo": `{
			"Cisco-IOS-XE-wireless-general-cfg:laginfo": {"lag-support": true}
		}`,

		"Cisco-IOS-XE-wireless-general-cfg:general-cfg-data/multicast-config": `{
			"Cisco-IOS-XE-wireless-general-cfg:multicast-config": {"multicast-mode": "unicast"}
		}`,

		"Cisco-IOS-XE-wireless-general-cfg:general-cfg-data/feature-usage-cfg": `{
			"Cisco-IOS-XE-wireless-general-cfg:feature-usage-cfg": {"enable": true}
		}`,

		"Cisco-IOS-XE-wireless-general-cfg:general-cfg-data/threshold-warn-cfg": `{
			"Cisco-IOS-XE-wireless-general-cfg:threshold-warn-cfg": {"memory-threshold": 80}
		}`,

		"Cisco-IOS-XE-wireless-general-cfg:general-cfg-data/ap-loc-ranging-cfg": `{
			"Cisco-IOS-XE-wireless-general-cfg:ap-loc-ranging-cfg": {"enable": false}
		}`,

		"Cisco-IOS-XE-wireless-general-cfg:general-cfg-data/geolocation-cfg": `{
			"Cisco-IOS-XE-wireless-general-cfg:geolocation-cfg": {"enable": false}
		}`,

		// Individual operational endpoints
		"Cisco-IOS-XE-wireless-general-oper:general-oper-data/mgmt-intf-data": `{
			"Cisco-IOS-XE-wireless-general-oper:mgmt-intf-data": {
				"intf-name": "GigabitEthernet0",
				"intf-type": "ethernet",
				"intf-id": 0,
				"mgmt-ip": "192.168.1.100",
				"net-mask": "255.255.255.0",
				"mgmt-mac": "aa:bb:cc:dd:ee:ff"
			}
		}`,
	}

	mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(responses))
	defer mockServer.Close()

	testClient := testutil.NewTestClient(mockServer)
	service := general.NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	// Test base Get* functions (operational)
	t.Run("GetOperational", func(t *testing.T) {
		result, err := service.GetOperational(ctx)
		if err != nil {
			t.Fatalf("GetOperational failed: %v", err)
		}
		if result == nil {
			t.Fatal("GetOperational returned nil result")
		}
	})

	t.Run("GetManagementInterfaceState", func(t *testing.T) {
		result, err := service.GetManagementInterfaceState(ctx)
		if err != nil {
			t.Fatalf("GetManagementInterfaceState failed: %v", err)
		}
		if result == nil {
			t.Fatal("GetManagementInterfaceState returned nil result")
		}
	})

	// Test base Get* functions (configuration)
	t.Run("GetConfig", func(t *testing.T) {
		result, err := service.GetConfig(ctx)
		if err != nil {
			t.Fatalf("GetConfig failed: %v", err)
		}
		if result == nil {
			t.Fatal("GetConfig returned nil result")
		}
	})

	t.Run("GetAPLocationRangingConfig", func(t *testing.T) {
		result, err := service.GetAPLocationRangingConfig(ctx)
		if err != nil {
			t.Fatalf("GetAPLocationRangingConfig failed: %v", err)
		}
		if result == nil {
			t.Fatal("GetAPLocationRangingConfig returned nil result")
		}
	})

	t.Run("GetCACConfig", func(t *testing.T) {
		result, err := service.GetCACConfig(ctx)
		if err != nil {
			t.Fatalf("GetCACConfig failed: %v", err)
		}
		if result == nil {
			t.Fatal("GetCACConfig returned nil result")
		}
	})

	t.Run("GetFeatureUsageConfig", func(t *testing.T) {
		result, err := service.GetFeatureUsageConfig(ctx)
		if err != nil {
			t.Fatalf("GetFeatureUsageConfig failed: %v", err)
		}
		if result == nil {
			t.Fatal("GetFeatureUsageConfig returned nil result")
		}
	})

	t.Run("GetFIPSConfig", func(t *testing.T) {
		result, err := service.GetFIPSConfig(ctx)
		if err != nil {
			t.Fatalf("GetFIPSConfig failed: %v", err)
		}
		if result == nil {
			t.Fatal("GetFIPSConfig returned nil result")
		}
	})

	t.Run("GetGeolocationConfig", func(t *testing.T) {
		result, err := service.GetGeolocationConfig(ctx)
		if err != nil {
			t.Fatalf("GetGeolocationConfig failed: %v", err)
		}
		if result == nil {
			t.Fatal("GetGeolocationConfig returned nil result")
		}
	})

	t.Run("GetLAGInfo", func(t *testing.T) {
		result, err := service.GetLAGInfo(ctx)
		if err != nil {
			t.Fatalf("GetLAGInfo failed: %v", err)
		}
		if result == nil {
			t.Fatal("GetLAGInfo returned nil result")
		}
	})

	t.Run("GetMEWLCConfig", func(t *testing.T) {
		result, err := service.GetMEWLCConfig(ctx)
		if err != nil {
			t.Fatalf("GetMEWLCConfig failed: %v", err)
		}
		if result == nil {
			t.Fatal("GetMEWLCConfig returned nil result")
		}
	})

	t.Run("GetMFPConfig", func(t *testing.T) {
		result, err := service.GetMFPConfig(ctx)
		if err != nil {
			t.Fatalf("GetMFPConfig failed: %v", err)
		}
		if result == nil {
			t.Fatal("GetMFPConfig returned nil result")
		}
	})

	t.Run("GetMulticastConfig", func(t *testing.T) {
		result, err := service.GetMulticastConfig(ctx)
		if err != nil {
			t.Fatalf("GetMulticastConfig failed: %v", err)
		}
		if result == nil {
			t.Fatal("GetMulticastConfig returned nil result")
		}
	})

	t.Run("GetThresholdWarningConfig", func(t *testing.T) {
		result, err := service.GetThresholdWarningConfig(ctx)
		if err != nil {
			t.Fatalf("GetThresholdWarningConfig failed: %v", err)
		}
		if result == nil {
			t.Fatal("GetThresholdWarningConfig returned nil result")
		}
	})

	t.Run("GetWLCManagementInfo", func(t *testing.T) {
		result, err := service.GetWLCManagementInfo(ctx)
		if err != nil {
			t.Fatalf("GetWLCManagementInfo failed: %v", err)
		}
		if result == nil {
			t.Fatal("GetWLCManagementInfo returned nil result")
		}
	})

	t.Run("GetWSAAPClientEventConfig", func(t *testing.T) {
		result, err := service.GetWSAAPClientEventConfig(ctx)
		if err != nil {
			t.Fatalf("GetWSAAPClientEventConfig failed: %v", err)
		}
		if result == nil {
			t.Fatal("GetWSAAPClientEventConfig returned nil result")
		}
	})

	// Test existing List* function
	t.Run("ListSIML3InterfaceCache", func(t *testing.T) {
		result, err := service.ListSIML3InterfaceCache(ctx)
		if err != nil {
			t.Fatalf("ListSIML3InterfaceCache failed: %v", err)
		}
		if result == nil {
			t.Fatal("ListSIML3InterfaceCache returned nil result")
		}
	})

	t.Run("ListCfgMewlcConfig", func(t *testing.T) {
		result, err := service.ListCfgMewlcConfig(ctx)
		if err != nil {
			t.Fatalf("ListCfgMewlcConfig failed: %v", err)
		}
		if result == nil {
			t.Fatal("ListCfgMewlcConfig returned nil result")
		}
	})

	t.Run("ListCfgCacConfig", func(t *testing.T) {
		result, err := service.ListCfgCacConfig(ctx)
		if err != nil {
			t.Fatalf("ListCfgCacConfig failed: %v", err)
		}
		if result == nil {
			t.Fatal("ListCfgCacConfig returned nil result")
		}
	})

	t.Run("ListCfgMfp", func(t *testing.T) {
		result, err := service.ListCfgMfp(ctx)
		if err != nil {
			t.Fatalf("ListCfgMfp failed: %v", err)
		}
		if result == nil {
			t.Fatal("ListCfgMfp returned nil result")
		}
	})

	t.Run("ListCfgFipsCfg", func(t *testing.T) {
		result, err := service.ListCfgFipsCfg(ctx)
		if err != nil {
			t.Fatalf("ListCfgFipsCfg failed: %v", err)
		}
		if result == nil {
			t.Fatal("ListCfgFipsCfg returned nil result")
		}
	})

	t.Run("ListCfgWsaApClientEvent", func(t *testing.T) {
		result, err := service.ListCfgWsaApClientEvent(ctx)
		if err != nil {
			t.Fatalf("ListCfgWsaApClientEvent failed: %v", err)
		}
		if result == nil {
			t.Fatal("ListCfgWsaApClientEvent returned nil result")
		}
	})

	t.Run("ListCfgSimL3InterfaceCacheData", func(t *testing.T) {
		result, err := service.ListCfgSimL3InterfaceCacheData(ctx)
		if err != nil {
			t.Fatalf("ListCfgSimL3InterfaceCacheData failed: %v", err)
		}
		if result == nil {
			t.Fatal("ListCfgSimL3InterfaceCacheData returned nil result")
		}
	})

	t.Run("ListCfgWlcManagementData", func(t *testing.T) {
		result, err := service.ListCfgWlcManagementData(ctx)
		if err != nil {
			t.Fatalf("ListCfgWlcManagementData failed: %v", err)
		}
		if result == nil {
			t.Fatal("ListCfgWlcManagementData returned nil result")
		}
	})

	t.Run("ListCfgLaginfo", func(t *testing.T) {
		result, err := service.ListCfgLaginfo(ctx)
		if err != nil {
			t.Fatalf("ListCfgLaginfo failed: %v", err)
		}
		if result == nil {
			t.Fatal("ListCfgLaginfo returned nil result")
		}
	})

	t.Run("ListCfgMulticastConfig", func(t *testing.T) {
		result, err := service.ListCfgMulticastConfig(ctx)
		if err != nil {
			t.Fatalf("ListCfgMulticastConfig failed: %v", err)
		}
		if result == nil {
			t.Fatal("ListCfgMulticastConfig returned nil result")
		}
	})

	t.Run("ListCfgFeatureUsageCfg", func(t *testing.T) {
		result, err := service.ListCfgFeatureUsageCfg(ctx)
		if err != nil {
			t.Fatalf("ListCfgFeatureUsageCfg failed: %v", err)
		}
		if result == nil {
			t.Fatal("ListCfgFeatureUsageCfg returned nil result")
		}
	})

	t.Run("ListCfgThresholdWarnCfg", func(t *testing.T) {
		result, err := service.ListCfgThresholdWarnCfg(ctx)
		if err != nil {
			t.Fatalf("ListCfgThresholdWarnCfg failed: %v", err)
		}
		if result == nil {
			t.Fatal("ListCfgThresholdWarnCfg returned nil result")
		}
	})

	t.Run("ListCfgApLocRangingCfg", func(t *testing.T) {
		result, err := service.ListCfgApLocRangingCfg(ctx)
		if err != nil {
			t.Fatalf("ListCfgApLocRangingCfg failed: %v", err)
		}
		if result == nil {
			t.Fatal("ListCfgApLocRangingCfg returned nil result")
		}
	})

	t.Run("ListCfgGeolocationCfg", func(t *testing.T) {
		result, err := service.ListCfgGeolocationCfg(ctx)
		if err != nil {
			t.Fatalf("ListCfgGeolocationCfg failed: %v", err)
		}
		if result == nil {
			t.Fatal("ListCfgGeolocationCfg returned nil result")
		}
	})

	t.Run("ListOperMgmtIntfData", func(t *testing.T) {
		result, err := service.ListOperMgmtIntfData(ctx)
		if err != nil {
			t.Fatalf("ListOperMgmtIntfData failed: %v", err)
		}
		if result == nil {
			t.Fatal("ListOperMgmtIntfData returned nil result")
		}
	})

	t.Run("ListMewlcConfig", func(t *testing.T) {
		result, err := service.ListMewlcConfig(ctx)
		if err != nil {
			t.Fatalf("ListMewlcConfig failed: %v", err)
		}
		if result == nil {
			t.Fatal("ListMewlcConfig returned nil result")
		}
	})

	t.Run("ListCacConfig", func(t *testing.T) {
		result, err := service.ListCacConfig(ctx)
		if err != nil {
			t.Fatalf("ListCacConfig failed: %v", err)
		}
		if result == nil {
			t.Fatal("ListCacConfig returned nil result")
		}
	})

	t.Run("ListMfp", func(t *testing.T) {
		result, err := service.ListMfp(ctx)
		if err != nil {
			t.Fatalf("ListMfp failed: %v", err)
		}
		if result == nil {
			t.Fatal("ListMfp returned nil result")
		}
	})

	t.Run("ListFipsCfg", func(t *testing.T) {
		result, err := service.ListFipsCfg(ctx)
		if err != nil {
			t.Fatalf("ListFipsCfg failed: %v", err)
		}
		if result == nil {
			t.Fatal("ListFipsCfg returned nil result")
		}
	})

	t.Run("ListWsaApClientEvent", func(t *testing.T) {
		result, err := service.ListWsaApClientEvent(ctx)
		if err != nil {
			t.Fatalf("ListWsaApClientEvent failed: %v", err)
		}
		if result == nil {
			t.Fatal("ListWsaApClientEvent returned nil result")
		}
	})

	t.Run("ListWlcManagementData", func(t *testing.T) {
		result, err := service.ListWlcManagementData(ctx)
		if err != nil {
			t.Fatalf("ListWlcManagementData failed: %v", err)
		}
		if result == nil {
			t.Fatal("ListWlcManagementData returned nil result")
		}
	})

	t.Run("ListLaginfo", func(t *testing.T) {
		result, err := service.ListLaginfo(ctx)
		if err != nil {
			t.Fatalf("ListLaginfo failed: %v", err)
		}
		if result == nil {
			t.Fatal("ListLaginfo returned nil result")
		}
	})

	t.Run("ListMulticastConfig", func(t *testing.T) {
		result, err := service.ListMulticastConfig(ctx)
		if err != nil {
			t.Fatalf("ListMulticastConfig failed: %v", err)
		}
		if result == nil {
			t.Fatal("ListMulticastConfig returned nil result")
		}
	})

	t.Run("ListFeatureUsageCfg", func(t *testing.T) {
		result, err := service.ListFeatureUsageCfg(ctx)
		if err != nil {
			t.Fatalf("ListFeatureUsageCfg failed: %v", err)
		}
		if result == nil {
			t.Fatal("ListFeatureUsageCfg returned nil result")
		}
	})

	t.Run("ListThresholdWarnCfg", func(t *testing.T) {
		result, err := service.ListThresholdWarnCfg(ctx)
		if err != nil {
			t.Fatalf("ListThresholdWarnCfg failed: %v", err)
		}
		if result == nil {
			t.Fatal("ListThresholdWarnCfg returned nil result")
		}
	})

	t.Run("ListApLocRangingCfg", func(t *testing.T) {
		result, err := service.ListApLocRangingCfg(ctx)
		if err != nil {
			t.Fatalf("ListApLocRangingCfg failed: %v", err)
		}
		if result == nil {
			t.Fatal("ListApLocRangingCfg returned nil result")
		}
	})

	t.Run("ListGeolocationCfg", func(t *testing.T) {
		result, err := service.ListGeolocationCfg(ctx)
		if err != nil {
			t.Fatalf("ListGeolocationCfg failed: %v", err)
		}
		if result == nil {
			t.Fatal("ListGeolocationCfg returned nil result")
		}
	})
}

// TestGeneralServiceUnit_GetOperations_ErrorHandling tests error scenarios using mock server.
func TestGeneralServiceUnit_GetOperations_ErrorHandling(t *testing.T) {
	// Create mock server that returns 404 for General endpoints
	errorPaths := []string{
		"Cisco-IOS-XE-wireless-general-oper:general-oper-data",
		"Cisco-IOS-XE-wireless-general-cfg:general-cfg-data",
	}
	mockServer := testutil.NewMockServer(testutil.WithErrorResponses(errorPaths, 404))
	defer mockServer.Close()

	testClient := testutil.NewTestClient(mockServer)
	service := general.NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	// Test that GetOperational properly handles 404 errors
	t.Run("GetOperational_404Error", func(t *testing.T) {
		_, err := service.GetOperational(ctx)
		if err == nil {
			t.Error("Expected error for 404 response, got nil")
		}
	})

	// Test representative configuration functions error handling
	t.Run("GetConfig_404Error", func(t *testing.T) {
		_, err := service.GetConfig(ctx)
		if err == nil {
			t.Error("Expected error for 404 response, got nil")
		}
	})

	t.Run("ListCfgMewlcConfig_404Error", func(t *testing.T) {
		_, err := service.ListCfgMewlcConfig(ctx)
		if err == nil {
			t.Error("Expected error for 404 response, got nil")
		}
	})

	t.Run("ListMewlcConfig_404Error", func(t *testing.T) {
		_, err := service.ListMewlcConfig(ctx)
		if err == nil {
			t.Error("Expected error for 404 response, got nil")
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

	t.Run("ListCfgMewlcConfig_NilClient", func(t *testing.T) {
		service := general.NewService(nil)
		ctx := testutil.TestContext(t)

		result, err := service.ListCfgMewlcConfig(ctx)
		if err == nil {
			t.Error("Expected error for nil client")
		}
		if result != nil {
			t.Error("Expected nil result for error case")
		}
	})

	t.Run("ListMewlcConfig_NilClient", func(t *testing.T) {
		service := general.NewService(nil)
		ctx := testutil.TestContext(t)

		result, err := service.ListMewlcConfig(ctx)
		if err == nil {
			t.Error("Expected error for nil client")
		}
		if result != nil {
			t.Error("Expected nil result for error case")
		}
	})
}
