package rrm_test

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/pkg/testutil"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/rrm"
)

// TestRrmServiceUnit_Constructor_Success tests service constructor functionality.
func TestRrmServiceUnit_Constructor_Success(t *testing.T) {
	t.Parallel()

	t.Run("NewServiceWithValidClient", func(t *testing.T) {
		// Create mock server and test client
		mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(map[string]string{
			"test": `{"data": {}}`,
		}))
		defer mockServer.Close()

		client := testutil.NewTestClient(mockServer)
		service := rrm.NewService(client.Core().(*core.Client))
		if service.Client() == nil {
			t.Error("Expected service to have client, got nil")
		}
	})

	t.Run("NewServiceWithNilClient", func(t *testing.T) {
		service := rrm.NewService(nil)
		if service.Client() != nil {
			t.Error("Expected nil client service")
		}
	})
}

// TestRrmServiceUnit_GetConfigOperations_MockSuccess tests Get configuration operations using mock server.
func TestRrmServiceUnit_GetConfigOperations_MockSuccess(t *testing.T) {
	t.Parallel()

	// Mock responses based on real WNC RRM data structure (IOS-XE 17.12.5)
	responses := map[string]string{
		"Cisco-IOS-XE-wireless-rrm-cfg:rrm-cfg-data": `{
			"Cisco-IOS-XE-wireless-rrm-cfg:rrm-cfg-data": {
				"rrms": {
					"rrm": [
						{
							"band": "dot11-2-dot-4-ghz-band",
							"rrm": {
								"roaming-en": true,
								"data-rate-threshold": "optroam-rate-12-m"
							}
						},
						{
							"band": "dot11-5-ghz-band",
							"rrm": {
								"roaming-en": true,
								"data-rate-threshold": "optroam-rate-24-m"
							}
						},
						{
							"band": "dot11-6-ghz-band",
							"rrm": {
								"roaming-en": true,
								"data-rate-threshold": "optroam-rate-24-m",
								"measurement-interval": 600
							}
						}
					]
				},
				"rrm-mgr-cfg-entries": {
					"rrm-mgr-cfg-entry": [
						{
							"band": "dot11-2-dot-4-ghz-band"
						},
						{
							"band": "dot11-5-ghz-band"
						},
						{
							"band": "dot11-6-ghz-band"
						}
					]
				}
			}
		}`,
		"Cisco-IOS-XE-wireless-rrm-oper:rrm-oper-data": `{
			"Cisco-IOS-XE-wireless-rrm-oper:rrm-oper-data": {
				"ap-auto-rf-dot11-data": [
					{
						"wtp-mac": "aa:bb:cc:dd:ee:ff",
						"radio-slot-id": 0,
						"neighbor-radio-info": {
							"neighbor-radio-list": [
								{
									"neighbor-radio-info": {
										"neighbor-radio-mac": "aa:bb:cc:dd:ee:ff",
										"neighbor-radio-slot-id": 0,
										"rssi": -19,
										"snr": 62,
										"channel": 11,
										"power": 18,
										"group-leader-ip": "192.168.255.1"
									}
								}
							]
						}
					}
				]
			}
		}`,
		"Cisco-IOS-XE-wireless-rrm-global-oper:rrm-global-oper-data": `{
			"Cisco-IOS-XE-wireless-rrm-global-oper:rrm-global-oper-data": {
				"rrm-one-shot-counters": [
					{
						"phy-type": "rrm-phy-80211b",
						"power-counter": 0
					},
					{
						"phy-type": "rrm-phy-80211a",
						"power-counter": 0
					}
				]
			}
		}`,
		"Cisco-IOS-XE-wireless-rrm-emul-oper:rrm-emul-oper-data": `{
			"Cisco-IOS-XE-wireless-rrm-emul-oper:rrm-emul-oper-data": {
				"rrm-fra-stats": {
					"dual-band-monitor-to-24ghz": 0,
					"dual-band-monitor-to-5ghz": 0,
					"dual-band-24ghz-to-5ghz": 0,
					"dual-band-24ghz-to-monitor": 0,
					"dual-band-5ghz-to-24ghz": 0,
					"dual-band-5ghz-to-monitor": 0
				}
			}
		}`,
	}

	mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(responses))
	defer mockServer.Close()

	testClient := testutil.NewTestClient(mockServer)
	service := rrm.NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	t.Run("GetConfig", func(t *testing.T) {
		result, err := service.GetConfig(ctx)
		if err != nil {
			t.Errorf("GetConfig returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("GetConfig returned nil result")
		}
	})

	t.Run("GetOperational", func(t *testing.T) {
		result, err := service.GetOperational(ctx)
		if err != nil {
			t.Errorf("GetOperational returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("GetOperational returned nil result")
		}
	})

	t.Run("GetGlobalOperational", func(t *testing.T) {
		result, err := service.GetGlobalOperational(ctx)
		if err != nil {
			t.Errorf("GetGlobalOperational returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("GetGlobalOperational returned nil result")
		}
	})

	t.Run("GetEmulationOperational", func(t *testing.T) {
		result, err := service.GetEmulationOperational(ctx)
		if err != nil {
			t.Errorf("GetEmulationOperational returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("GetEmulationOperational returned nil result")
		}
	})

	// Test List operations
	t.Run("ListRrms", func(t *testing.T) {
		result, err := service.ListRrms(ctx)
		if err != nil {
			t.Errorf("ListRrms returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("ListRrms returned nil result")
		}
	})

	t.Run("ListRRMMgrCfgEntries", func(t *testing.T) {
		result, err := service.ListRRMMgrCfgEntries(ctx)
		if err != nil {
			t.Errorf("ListRRMMgrCfgEntries returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("ListRRMMgrCfgEntries returned nil result")
		}
	})
}

// TestRrmServiceUnit_GetOperations_ErrorHandling tests error scenarios for operations.
func TestRrmServiceUnit_GetOperations_ErrorHandling(t *testing.T) {
	t.Parallel()

	// Create test server and service
	server := testutil.NewMockServer(testutil.WithSuccessResponses(map[string]string{}))
	defer server.Close()

	// Create test client configured for the mock server
	testClient := testutil.NewTestClient(server)
	service := rrm.NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	t.Run("GetConfig_404Error", func(t *testing.T) {
		result, err := service.GetConfig(ctx)
		if err == nil {
			t.Error("Expected error for GetConfig, got nil")
		}
		if result != nil {
			t.Error("Expected nil result on error, got non-nil result")
		}
	})

	t.Run("GetOperational_404Error", func(t *testing.T) {
		result, err := service.GetOperational(ctx)
		if err == nil {
			t.Error("Expected error for GetOperational, got nil")
		}
		if result != nil {
			t.Error("Expected nil result on error, got non-nil result")
		}
	})

	t.Run("GetGlobalOperational_404Error", func(t *testing.T) {
		result, err := service.GetGlobalOperational(ctx)
		if err == nil {
			t.Error("Expected error for GetGlobalOperational, got nil")
		}
		if result != nil {
			t.Error("Expected nil result on error, got non-nil result")
		}
	})

	t.Run("GetEmulationOperational_404Error", func(t *testing.T) {
		result, err := service.GetEmulationOperational(ctx)
		if err == nil {
			t.Error("Expected error for GetEmulationOperational, got nil")
		}
		if result != nil {
			t.Error("Expected nil result on error, got non-nil result")
		}
	})

	t.Run("ListRrms_404Error", func(t *testing.T) {
		result, err := service.ListRrms(ctx)
		if err == nil {
			t.Error("Expected error for ListRrms, got nil")
		}
		if result != nil {
			t.Error("Expected nil result on error, got non-nil result")
		}
	})

	t.Run("ListRRMMgrCfgEntries_404Error", func(t *testing.T) {
		result, err := service.ListRRMMgrCfgEntries(ctx)
		if err == nil {
			t.Error("Expected error for ListRRMMgrCfgEntries, got nil")
		}
		if result != nil {
			t.Error("Expected nil result on error, got non-nil result")
		}
	})
}

// TestRrmServiceUnit_ErrorHandling_NilClient tests error handling with nil client.
func TestRrmServiceUnit_ErrorHandling_NilClient(t *testing.T) {
	t.Parallel()

	service := rrm.NewService(nil)
	ctx := testutil.TestContext(t)

	t.Run("GetConfig_NilClient", func(t *testing.T) {
		result, err := service.GetConfig(ctx)
		if err == nil {
			t.Error("Expected error for nil client")
		}
		if result != nil {
			t.Error("Expected nil result for error case")
		}
	})

	t.Run("GetOperational_NilClient", func(t *testing.T) {
		result, err := service.GetOperational(ctx)
		if err == nil {
			t.Error("Expected error for nil client")
		}
		if result != nil {
			t.Error("Expected nil result for error case")
		}
	})

	t.Run("GetGlobalOperational_NilClient", func(t *testing.T) {
		result, err := service.GetGlobalOperational(ctx)
		if err == nil {
			t.Error("Expected error for nil client")
		}
		if result != nil {
			t.Error("Expected nil result for error case")
		}
	})

	t.Run("GetEmulationOperational_NilClient", func(t *testing.T) {
		result, err := service.GetEmulationOperational(ctx)
		if err == nil {
			t.Error("Expected error for nil client")
		}
		if result != nil {
			t.Error("Expected nil result for error case")
		}
	})

	t.Run("ListRrms_NilClient", func(t *testing.T) {
		result, err := service.ListRrms(ctx)
		if err == nil {
			t.Error("Expected error for nil client")
		}
		if result != nil {
			t.Error("Expected nil result for error case")
		}
	})

	t.Run("ListRRMMgrCfgEntries_NilClient", func(t *testing.T) {
		result, err := service.ListRRMMgrCfgEntries(ctx)
		if err == nil {
			t.Error("Expected error for nil client")
		}
		if result != nil {
			t.Error("Expected nil result for error case")
		}
	})
}
