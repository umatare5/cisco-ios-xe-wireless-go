package routes

// RRM (Radio Resource Management) Configuration and Operational Paths
//
// These constants define the RESTCONF API paths for RRM configuration
// and operational data based on Cisco-IOS-XE-wireless-rrm YANG models.

// RRM Configuration Paths.
const (
	// RRMCfgPath provides the path for RRM configuration data.
	RRMCfgPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-rrm-cfg:rrm-cfg-data"

	// RRMCfgRrmsPath provides the path for RRM configurations by band.
	RRMCfgRrmsPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-rrm-cfg:rrm-cfg-data/rrms"

	// RRMCfgRRMMgrCfgEntriesPath provides the path for RRM manager configuration entries.
	RRMCfgRRMMgrCfgEntriesPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-rrm-cfg:rrm-cfg-data/rrm-mgr-cfg-entries"
)

// RRM Operational Paths.
const (
	// RRMOperPath provides the path for RRM operational data.
	RRMOperPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-rrm-oper:rrm-oper-data"

	// RRMOperApAutoRFDot11DataPath provides the path for AP auto RF 802.11 data.
	RRMOperApAutoRFDot11DataPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-rrm-oper:rrm-oper-data/ap-auto-rf-dot11-data"

	// RRMOperApDot11RadarDataPath provides the path for AP radar detection data.
	RRMOperApDot11RadarDataPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-rrm-oper:rrm-oper-data/ap-dot11-radar-data"

	// RRMOperApDot11SpectrumDataPath provides the path for AP spectrum analysis data.
	RRMOperApDot11SpectrumDataPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-rrm-oper:rrm-oper-data/ap-dot11-spectrum-data"

	// RRMOperRRMMeasurementPath provides the path for RRM measurement data.
	RRMOperRRMMeasurementPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-rrm-oper:rrm-oper-data/rrm-measurement"

	// RRMOperRadioSlotPath provides the path for radio slot operational data.
	RRMOperRadioSlotPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-rrm-oper:rrm-oper-data/radio-slot"

	// RRMOperMainDataPath provides the path for main RRM data by PHY type.
	RRMOperMainDataPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-rrm-oper:rrm-oper-data/main-data"

	// RRMOperRegDomainOperPath provides the path for regulatory domain operational data.
	RRMOperRegDomainOperPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-rrm-oper:rrm-oper-data/reg-domain-oper"

	// RRMOperSpectrumDeviceTablePath provides the path for spectrum device detection table.
	RRMOperSpectrumDeviceTablePath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-rrm-oper:rrm-oper-data/spectrum-device-table"

	// RRMOperSpectrumAqTablePath provides the path for spectrum air quality table.
	RRMOperSpectrumAqTablePath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-rrm-oper:rrm-oper-data/spectrum-aq-table"

	// RRMRadioStatsPath provides the path for RRM radio statistics.
	RRMRadioStatsPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-rrm-oper:rrm-oper-data/rrm-radio-stats"

	// RRMLoadStatsPath provides the path for RRM load statistics.
	RRMLoadStatsPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-rrm-oper:rrm-oper-data/rrm-load-stats"

	// RRMNoiseStatsPath provides the path for RRM noise statistics.
	RRMNoiseStatsPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-rrm-oper:rrm-oper-data/rrm-noise-stats"

	// RRMNeighborStatsPath provides the path for RRM neighbor statistics.
	RRMNeighborStatsPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-rrm-oper:rrm-oper-data/rrm-neighbor-stats"

	// RRM24GhzConfigPath provides the path for RRM 2.4GHz configuration.
	RRM24GhzConfigPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-rrm-oper:rrm-oper-data/rrm-24ghz-config"

	// RRM5GhzConfigPath provides the path for RRM 5GHz configuration.
	RRM5GhzConfigPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-rrm-oper:rrm-oper-data/rrm-5ghz-config"
)

// RRM Global Operational Paths.
const (
	// RRMGlobalOperPath provides the path for RRM global operational data.
	RRMGlobalOperPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-rrm-global-oper:rrm-global-oper-data"

	// RRMGlobalOperRRMOneShotCountersPath provides the path for RRM one-shot counters.
	RRMGlobalOperRRMOneShotCountersPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-rrm-global-oper:rrm-global-oper-data/rrm-one-shot-counters"

	// RRMGlobalOperRRMChannelParamsPath provides the path for RRM channel parameters.
	RRMGlobalOperRRMChannelParamsPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-rrm-global-oper:rrm-global-oper-data/rrm-channel-params"

	// RRMGlobalOperRadioOperData24gPath provides the path for 2.4GHz radio operational data.
	RRMGlobalOperRadioOperData24gPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-rrm-global-oper:rrm-global-oper-data/radio-oper-data-24g"

	// RRMGlobalOperRadioOperData5gPath provides the path for 5GHz radio operational data.
	RRMGlobalOperRadioOperData5gPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-rrm-global-oper:rrm-global-oper-data/radio-oper-data-5g"

	// RRMGlobalOperRadioOperData6ghzPath provides the path for 6GHz radio operational data.
	RRMGlobalOperRadioOperData6ghzPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-rrm-global-oper:rrm-global-oper-data/radio-oper-data-6ghz"

	// RRMGlobalOperRadioOperDataDualbandPath provides the path for dual-band radio operational data.
	RRMGlobalOperRadioOperDataDualbandPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-rrm-global-oper:rrm-global-oper-data/radio-oper-data-dualband"

	// RRMGlobalOperSpectrumBandConfigDataPath provides the path for spectrum band configuration data.
	RRMGlobalOperSpectrumBandConfigDataPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-rrm-global-oper:rrm-global-oper-data/spectrum-band-config-data"

	// RRMGlobalOperRRMClientDataPath provides the path for RRM client data.
	RRMGlobalOperRRMClientDataPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-rrm-global-oper:rrm-global-oper-data/rrm-client-data"

	// RRMGlobalOperRRMFraStatsPath provides the path for RRM flexible radio assignment statistics.
	RRMGlobalOperRRMFraStatsPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-rrm-global-oper:rrm-global-oper-data/rrm-fra-stats"

	// RRMGlobalOperRRMCoveragePath provides the path for RRM coverage information.
	RRMGlobalOperRRMCoveragePath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-rrm-global-oper:rrm-global-oper-data/rrm-coverage"

	// RRMGlobalOperSpectrumAqWorstTablePath provides the path for spectrum air quality worst table.
	RRMGlobalOperSpectrumAqWorstTablePath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-rrm-global-oper:rrm-global-oper-data/spectrum-aq-worst-table"

	// RRMGlobalStatsPath provides the path for RRM global statistics.
	RRMGlobalStatsPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-rrm-global-oper:rrm-global-oper-data/rrm-global-stats"
)

// RRM Emulation Operational Paths.
const (
	// RRMEmulOperPath provides the path for RRM emulation operational data.
	RRMEmulOperPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-rrm-emul-oper:rrm-emul-oper-data"

	// RRMEmulOperRRMFraStatsPath provides the path for RRM flexible radio assignment statistics.
	RRMEmulOperRRMFraStatsPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-rrm-emul-oper:rrm-emul-oper-data/rrm-fra-stats"

	// RRMEmulApDataPath provides the path for RRM emulation AP data.
	RRMEmulApDataPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-rrm-emul-oper:rrm-emul-oper-data/rrm-emul-ap-data"
)
