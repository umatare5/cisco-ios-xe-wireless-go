package routes

// RRM (Radio Resource Management) Configuration and Operational Paths
//
// These constants define the RESTCONF API paths for RRM configuration
// and operational data based on Cisco-IOS-XE-wireless-rrm YANG models.

// RRM Configuration Paths.
const (
	// RRMCfgPath provides the path for RRM configuration data.
	RRMCfgPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-rrm-cfg:rrm-cfg-data"
)

// RRM Operational Paths.
const (
	// RRMOperPath provides the path for RRM operational data.
	RRMOperPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-rrm-oper:rrm-oper-data"

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

	// RRMGlobalStatsPath provides the path for RRM global statistics.
	RRMGlobalStatsPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-rrm-global-oper:rrm-global-oper-data/rrm-global-stats"
)

// RRM Emulation Operational Paths.
const (
	// RRMEmulOperPath provides the path for RRM emulation operational data.
	RRMEmulOperPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-rrm-emul-oper:rrm-emul-oper-data"

	// RRMEmulApDataPath provides the path for RRM emulation AP data.
	RRMEmulApDataPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-rrm-emul-oper:rrm-emul-oper-data/rrm-emul-ap-data"
)
