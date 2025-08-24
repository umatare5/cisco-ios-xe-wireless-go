package routes

import "github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf"

// RRM (Radio Resource Management) Configuration and Operational Endpoints
//
// These constants define the RESTCONF API endpoints for RRM configuration
// and operational data based on Cisco-IOS-XE-wireless-rrm YANG models.

// RRM Base Paths
const (
	// RRMCfgBasePath defines the base path for RRM configuration endpoints
	RRMCfgBasePath = restconf.YANGModelPrefix + "rrm-cfg:rrm-cfg-data"

	// RRMOperBasePath defines the base path for RRM operational data endpoints
	RRMOperBasePath = restconf.YANGModelPrefix + "rrm-oper:rrm-oper-data"

	// RRMGlobalOperBasePath defines the base path for RRM global operational data endpoints
	RRMGlobalOperBasePath = restconf.YANGModelPrefix + "rrm-global-oper:rrm-global-oper-data"

	// RRMEmulOperBasePath defines the base path for RRM emulation operational data endpoints
	RRMEmulOperBasePath = restconf.YANGModelPrefix + "rrm-emul-oper:rrm-emul-oper-data"
)

// RRM Configuration Endpoints
const (
	// RRMCfgEndpoint retrieves complete RRM configuration data
	RRMCfgEndpoint = RRMCfgBasePath

	// RRMByBandEndpoint retrieves RRM configuration by band
	RRMByBandEndpoint = RRMCfgBasePath + "/rrms/rrm"

	// RRMMgrByBandEndpoint retrieves RRM manager configuration by band
	RRMMgrByBandEndpoint = RRMCfgBasePath + "/rrm-mgr-cfg-entries/rrm-mgr-cfg-entry"
)

// RRM Operational Endpoints
const (
	// RRMOperEndpoint retrieves RRM operational data
	RRMOperEndpoint = RRMOperBasePath

	// SpectrumDeviceTableEndpoint retrieves spectrum device table data
	SpectrumDeviceTableEndpoint = RRMOperBasePath + "/spectrum-device-table"

	// MainDataEndpoint retrieves RRM main operational data
	MainDataEndpoint = RRMOperBasePath + "/main-data"

	// ApAutoRfDot11DataEndpoint retrieves AP auto RF 802.11 data
	ApAutoRfDot11DataEndpoint = RRMOperBasePath + "/ap-auto-rf-dot11-data"
)

// RRM Global Operational Endpoints
const (
	// RRMGlobalOperEndpoint retrieves RRM global operational data
	RRMGlobalOperEndpoint = RRMGlobalOperBasePath

	// RadioOperData5GEndpoint retrieves 5G radio operational data
	RadioOperData5GEndpoint = RRMGlobalOperBasePath + "/radio-oper-data-5g"

	// RadioOperData6GhzEndpoint retrieves 6GHz radio operational data
	RadioOperData6GhzEndpoint = RRMGlobalOperBasePath + "/radio-oper-data-6ghz"

	// RadioOperData24GEndpoint retrieves 2.4G radio operational data
	RadioOperData24GEndpoint = RRMGlobalOperBasePath + "/radio-oper-data-24g"

	// SpectrumBandConfigDataEndpoint retrieves spectrum band configuration data
	SpectrumBandConfigDataEndpoint = RRMGlobalOperBasePath + "/spectrum-band-config-data"

	// SpectrumAqWorstTableEndpoint retrieves spectrum air quality worst table
	SpectrumAqWorstTableEndpoint = RRMGlobalOperBasePath + "/spectrum-aq-worst-table"

	// RRMChannelParamsEndpoint retrieves RRM channel parameters
	RRMChannelParamsEndpoint = RRMGlobalOperBasePath + "/rrm-channel-params"

	// RRMOneShotCountersEndpoint retrieves RRM one-shot counters
	RRMOneShotCountersEndpoint = RRMGlobalOperBasePath + "/rrm-one-shot-counters"
)

// RRM Emulation Operational Endpoints
const (
	// RRMEmulOperEndpoint retrieves RRM emulation operational data
	RRMEmulOperEndpoint = RRMEmulOperBasePath
)
