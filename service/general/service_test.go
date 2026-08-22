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

	t.Run("GetMEWLCConfig_404Error", func(t *testing.T) {
		_, err := service.GetMEWLCConfig(ctx)
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

	t.Run("GetMEWLCConfig_NilClient", func(t *testing.T) {
		service := general.NewService(nil)
		ctx := testutil.TestContext(t)

		result, err := service.GetMEWLCConfig(ctx)
		if err == nil {
			t.Error("Expected error for nil client")
		}
		if result != nil {
			t.Error("Expected nil result for error case")
		}
	})
}
