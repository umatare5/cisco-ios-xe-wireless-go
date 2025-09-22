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
		// Wrapper function endpoints - ListRrms
		"Cisco-IOS-XE-wireless-rrm-cfg:rrm-cfg-data/rrms": `{
			"Cisco-IOS-XE-wireless-rrm-cfg:rrms": {
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
			}
		}`,
		// Wrapper function endpoints - ListRRMMgrCfgEntries
		"Cisco-IOS-XE-wireless-rrm-cfg:rrm-cfg-data/rrm-mgr-cfg-entries": `{
			"Cisco-IOS-XE-wireless-rrm-cfg:rrm-mgr-cfg-entries": {
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
		// Additional wrapper endpoints for global operational
		"Cisco-IOS-XE-wireless-rrm-global-oper:rrm-global-oper-data/rrm-one-shot-counters": `{
			"Cisco-IOS-XE-wireless-rrm-global-oper:rrm-one-shot-counters": [
				{
					"phy-type": "rrm-phy-80211b",
					"power-counter": 0
				},
				{
					"phy-type": "rrm-phy-80211a",
					"power-counter": 0
				}
			]
		}`,
		"Cisco-IOS-XE-wireless-rrm-global-oper:rrm-global-oper-data/radio-oper-data-24g": `{
			"Cisco-IOS-XE-wireless-rrm-global-oper:radio-oper-data-24g": [
				{
					"wtp-mac": "aa:bb:cc:dd:ee:ff",
					"radio-slot-id": 0
				}
			]
		}`,
		"Cisco-IOS-XE-wireless-rrm-global-oper:rrm-global-oper-data/radio-oper-data-5g": `{
			"Cisco-IOS-XE-wireless-rrm-global-oper:radio-oper-data-5g": [
				{
					"wtp-mac": "aa:bb:cc:dd:ee:ff",
					"radio-slot-id": 1
				}
			]
		}`,
		"Cisco-IOS-XE-wireless-rrm-global-oper:rrm-global-oper-data/radio-oper-data-6ghz": `{
			"Cisco-IOS-XE-wireless-rrm-global-oper:radio-oper-data-6ghz": [
				{
					"wtp-mac": "aa:bb:cc:dd:ee:ff",
					"radio-slot-id": 2
				}
			]
		}`,
		// Additional operational endpoints
		"Cisco-IOS-XE-wireless-rrm-oper:rrm-oper-data/ap-auto-rf-dot11-data": `{
			"Cisco-IOS-XE-wireless-rrm-oper:ap-auto-rf-dot11-data": [
				{
					"wtp-mac": "aa:bb:cc:dd:ee:ff",
					"radio-slot-id": 0
				}
			]
		}`,
		"Cisco-IOS-XE-wireless-rrm-oper:rrm-oper-data/ap-dot11-radar-data": `{
			"Cisco-IOS-XE-wireless-rrm-oper:ap-dot11-radar-data": []
		}`,
		"Cisco-IOS-XE-wireless-rrm-oper:rrm-oper-data/ap-dot11-spectrum-data": `{
			"Cisco-IOS-XE-wireless-rrm-oper:ap-dot11-spectrum-data": []
		}`,
		"Cisco-IOS-XE-wireless-rrm-oper:rrm-oper-data/rrm-measurement": `{
			"Cisco-IOS-XE-wireless-rrm-oper:rrm-measurement": []
		}`,
		"Cisco-IOS-XE-wireless-rrm-oper:rrm-oper-data/radio-slot": `{
			"Cisco-IOS-XE-wireless-rrm-oper:radio-slot": []
		}`,
		"Cisco-IOS-XE-wireless-rrm-oper:rrm-oper-data/main-data": `{
			"Cisco-IOS-XE-wireless-rrm-oper:main-data": [
				{
					"phy-type": "rrm-phy-80211a",
					"country-code": "US"
				}
			]
		}`,
		"Cisco-IOS-XE-wireless-rrm-oper:rrm-oper-data/reg-domain-oper": `{
			"Cisco-IOS-XE-wireless-rrm-oper:reg-domain-oper": {
				"country-list": "US CA"
			}
		}`,
		"Cisco-IOS-XE-wireless-rrm-oper:rrm-oper-data/spectrum-device-table": `{
			"Cisco-IOS-XE-wireless-rrm-oper:spectrum-device-table": []
		}`,
		"Cisco-IOS-XE-wireless-rrm-oper:rrm-oper-data/spectrum-aq-table": `{
			"Cisco-IOS-XE-wireless-rrm-oper:spectrum-aq-table": []
		}`,
		// Additional global operational endpoints
		"Cisco-IOS-XE-wireless-rrm-global-oper:rrm-global-oper-data/rrm-channel-params": `{
			"Cisco-IOS-XE-wireless-rrm-global-oper:rrm-channel-params": []
		}`,
		"Cisco-IOS-XE-wireless-rrm-global-oper:rrm-global-oper-data/radio-oper-data-dualband": `{
			"Cisco-IOS-XE-wireless-rrm-global-oper:radio-oper-data-dualband": []
		}`,
		"Cisco-IOS-XE-wireless-rrm-global-oper:rrm-global-oper-data/spectrum-band-config-data": `{
			"Cisco-IOS-XE-wireless-rrm-global-oper:spectrum-band-config-data": []
		}`,
		"Cisco-IOS-XE-wireless-rrm-global-oper:rrm-global-oper-data/rrm-client-data": `{
			"Cisco-IOS-XE-wireless-rrm-global-oper:rrm-client-data": []
		}`,
		"Cisco-IOS-XE-wireless-rrm-global-oper:rrm-global-oper-data/rrm-fra-stats": `{
			"Cisco-IOS-XE-wireless-rrm-global-oper:rrm-fra-stats": {
				"dual-band-monitor-to-24ghz": 0,
				"dual-band-monitor-to-5ghz": 0
			}
		}`,
		"Cisco-IOS-XE-wireless-rrm-global-oper:rrm-global-oper-data/rrm-coverage": `{
			"Cisco-IOS-XE-wireless-rrm-global-oper:rrm-coverage": []
		}`,
		"Cisco-IOS-XE-wireless-rrm-global-oper:rrm-global-oper-data/spectrum-aq-worst-table": `{
			"Cisco-IOS-XE-wireless-rrm-global-oper:spectrum-aq-worst-table": []
		}`,
		// Emulation endpoints
		"Cisco-IOS-XE-wireless-rrm-emul-oper:rrm-emul-oper-data/rrm-fra-stats": `{
			"Cisco-IOS-XE-wireless-rrm-emul-oper:rrm-fra-stats": {
				"dual-band-monitor-to-24ghz": 0,
				"dual-band-monitor-to-5ghz": 0
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

	// === Operational Data Tests ===
	t.Run("ListApAutoRFDot11Data", func(t *testing.T) {
		result, err := service.ListApAutoRFDot11Data(ctx)
		if err != nil {
			t.Errorf("ListApAutoRFDot11Data returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("ListApAutoRFDot11Data returned nil result")
		}
	})

	t.Run("ListApDot11RadarData", func(t *testing.T) {
		result, err := service.ListApDot11RadarData(ctx)
		if err != nil {
			t.Errorf("ListApDot11RadarData returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("ListApDot11RadarData returned nil result")
		}
	})

	t.Run("ListApDot11SpectrumData", func(t *testing.T) {
		result, err := service.ListApDot11SpectrumData(ctx)
		if err != nil {
			t.Errorf("ListApDot11SpectrumData returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("ListApDot11SpectrumData returned nil result")
		}
	})

	t.Run("ListRRMMeasurement", func(t *testing.T) {
		result, err := service.ListRRMMeasurement(ctx)
		if err != nil {
			t.Errorf("ListRRMMeasurement returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("ListRRMMeasurement returned nil result")
		}
	})

	t.Run("ListRadioSlot", func(t *testing.T) {
		result, err := service.ListRadioSlot(ctx)
		if err != nil {
			t.Errorf("ListRadioSlot returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("ListRadioSlot returned nil result")
		}
	})

	t.Run("ListMainData", func(t *testing.T) {
		result, err := service.ListMainData(ctx)
		if err != nil {
			t.Errorf("ListMainData returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("ListMainData returned nil result")
		}
	})

	t.Run("ListRegDomainOper", func(t *testing.T) {
		result, err := service.ListRegDomainOper(ctx)
		if err != nil {
			t.Errorf("ListRegDomainOper returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("ListRegDomainOper returned nil result")
		}
	})

	t.Run("ListSpectrumDeviceTable", func(t *testing.T) {
		result, err := service.ListSpectrumDeviceTable(ctx)
		if err != nil {
			t.Errorf("ListSpectrumDeviceTable returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("ListSpectrumDeviceTable returned nil result")
		}
	})

	t.Run("ListSpectrumAqTable", func(t *testing.T) {
		result, err := service.ListSpectrumAqTable(ctx)
		if err != nil {
			t.Errorf("ListSpectrumAqTable returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("ListSpectrumAqTable returned nil result")
		}
	})

	// === Global Operational Data Tests ===
	t.Run("ListRRMOneShotCounters", func(t *testing.T) {
		result, err := service.ListRRMOneShotCounters(ctx)
		if err != nil {
			t.Errorf("ListRRMOneShotCounters returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("ListRRMOneShotCounters returned nil result")
		}
	})

	t.Run("ListRRMChannelParams", func(t *testing.T) {
		result, err := service.ListRRMChannelParams(ctx)
		if err != nil {
			t.Errorf("ListRRMChannelParams returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("ListRRMChannelParams returned nil result")
		}
	})

	t.Run("ListRadioOperData24g", func(t *testing.T) {
		result, err := service.ListRadioOperData24g(ctx)
		if err != nil {
			t.Errorf("ListRadioOperData24g returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("ListRadioOperData24g returned nil result")
		}
	})

	t.Run("ListRadioOperData5g", func(t *testing.T) {
		result, err := service.ListRadioOperData5g(ctx)
		if err != nil {
			t.Errorf("ListRadioOperData5g returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("ListRadioOperData5g returned nil result")
		}
	})

	t.Run("ListRadioOperData6ghz", func(t *testing.T) {
		result, err := service.ListRadioOperData6ghz(ctx)
		if err != nil {
			t.Errorf("ListRadioOperData6ghz returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("ListRadioOperData6ghz returned nil result")
		}
	})

	t.Run("ListRadioOperDataDualband", func(t *testing.T) {
		result, err := service.ListRadioOperDataDualband(ctx)
		if err != nil {
			t.Errorf("ListRadioOperDataDualband returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("ListRadioOperDataDualband returned nil result")
		}
	})

	t.Run("ListSpectrumBandConfigData", func(t *testing.T) {
		result, err := service.ListSpectrumBandConfigData(ctx)
		if err != nil {
			t.Errorf("ListSpectrumBandConfigData returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("ListSpectrumBandConfigData returned nil result")
		}
	})

	t.Run("ListRRMClientData", func(t *testing.T) {
		result, err := service.ListRRMClientData(ctx)
		if err != nil {
			t.Errorf("ListRRMClientData returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("ListRRMClientData returned nil result")
		}
	})

	t.Run("ListRRMFraStats", func(t *testing.T) {
		result, err := service.ListRRMFraStats(ctx)
		if err != nil {
			t.Errorf("ListRRMFraStats returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("ListRRMFraStats returned nil result")
		}
	})

	t.Run("ListRRMCoverage", func(t *testing.T) {
		result, err := service.ListRRMCoverage(ctx)
		if err != nil {
			t.Errorf("ListRRMCoverage returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("ListRRMCoverage returned nil result")
		}
	})

	t.Run("ListSpectrumAqWorstTable", func(t *testing.T) {
		result, err := service.ListSpectrumAqWorstTable(ctx)
		if err != nil {
			t.Errorf("ListSpectrumAqWorstTable returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("ListSpectrumAqWorstTable returned nil result")
		}
	})

	// === Emulation Operational Data Tests ===
	t.Run("ListRRMFraStatsFromEmul", func(t *testing.T) {
		result, err := service.ListRRMFraStatsFromEmul(ctx)
		if err != nil {
			t.Errorf("ListRRMFraStatsFromEmul returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("ListRRMFraStatsFromEmul returned nil result")
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

	// Operational data error tests
	t.Run("ListApAutoRFDot11Data_404Error", func(t *testing.T) {
		result, err := service.ListApAutoRFDot11Data(ctx)
		if err == nil {
			t.Error("Expected error for ListApAutoRFDot11Data, got nil")
		}
		if result != nil {
			t.Error("Expected nil result on error, got non-nil result")
		}
	})

	t.Run("ListApDot11RadarData_404Error", func(t *testing.T) {
		result, err := service.ListApDot11RadarData(ctx)
		if err == nil {
			t.Error("Expected error for ListApDot11RadarData, got nil")
		}
		if result != nil {
			t.Error("Expected nil result on error, got non-nil result")
		}
	})

	t.Run("ListRRMOneShotCounters_404Error", func(t *testing.T) {
		result, err := service.ListRRMOneShotCounters(ctx)
		if err == nil {
			t.Error("Expected error for ListRRMOneShotCounters, got nil")
		}
		if result != nil {
			t.Error("Expected nil result on error, got non-nil result")
		}
	})

	t.Run("ListRadioOperData24g_404Error", func(t *testing.T) {
		result, err := service.ListRadioOperData24g(ctx)
		if err == nil {
			t.Error("Expected error for ListRadioOperData24g, got nil")
		}
		if result != nil {
			t.Error("Expected nil result on error, got non-nil result")
		}
	})

	t.Run("ListRRMFraStatsFromEmul_404Error", func(t *testing.T) {
		result, err := service.ListRRMFraStatsFromEmul(ctx)
		if err == nil {
			t.Error("Expected error for ListRRMFraStatsFromEmul, got nil")
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

	// Operational data nil client tests
	t.Run("ListApAutoRFDot11Data_NilClient", func(t *testing.T) {
		result, err := service.ListApAutoRFDot11Data(ctx)
		if err == nil {
			t.Error("Expected error for nil client")
		}
		if result != nil {
			t.Error("Expected nil result for error case")
		}
	})

	t.Run("ListApDot11RadarData_NilClient", func(t *testing.T) {
		result, err := service.ListApDot11RadarData(ctx)
		if err == nil {
			t.Error("Expected error for nil client")
		}
		if result != nil {
			t.Error("Expected nil result for error case")
		}
	})

	t.Run("ListRRMOneShotCounters_NilClient", func(t *testing.T) {
		result, err := service.ListRRMOneShotCounters(ctx)
		if err == nil {
			t.Error("Expected error for nil client")
		}
		if result != nil {
			t.Error("Expected nil result for error case")
		}
	})

	t.Run("ListRadioOperData24g_NilClient", func(t *testing.T) {
		result, err := service.ListRadioOperData24g(ctx)
		if err == nil {
			t.Error("Expected error for nil client")
		}
		if result != nil {
			t.Error("Expected nil result for error case")
		}
	})

	t.Run("ListRRMFraStatsFromEmul_NilClient", func(t *testing.T) {
		result, err := service.ListRRMFraStatsFromEmul(ctx)
		if err == nil {
			t.Error("Expected error for nil client")
		}
		if result != nil {
			t.Error("Expected nil result for error case")
		}
	})
}
