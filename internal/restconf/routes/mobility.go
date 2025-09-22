package routes

// Mobility Configuration and Operational Paths
//
// These constants define the RESTCONF API paths for mobility configuration
// and operational data based on Cisco-IOS-XE-wireless-mobility YANG models.

// Mobility Configuration Paths.
const (
	// MobilityCfgPath provides the path for retrieving all mobility configuration data.
	MobilityCfgPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-mobility-cfg:mobility-cfg-data"

	// MobilityConfigPath provides the path for retrieving mobility configuration.
	MobilityConfigPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-mobility-cfg:mobility-cfg-data/mobility-config"
)

// Mobility Operational Paths
//
// These constants define the RESTCONF API paths for mobility operational
// data based on Cisco-IOS-XE-wireless-mobility-oper YANG model.

// Mobility Operational Paths.
const (
	// MobilityOperPath provides the path for retrieving all mobility operational data.
	MobilityOperPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-mobility-oper:mobility-oper-data"

	// MobilityApCachePath provides the path for retrieving AP cache data.
	MobilityApCachePath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-mobility-oper:mobility-oper-data/ap-cache"

	// MobilityApPeerListPath provides the path for retrieving AP peer list.
	MobilityApPeerListPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-mobility-oper:mobility-oper-data/ap-peer-list"

	// MobilityMmGlobalDataPath provides the path for retrieving MM global data.
	MobilityMmGlobalDataPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-mobility-oper:mobility-oper-data/mm-global-data"

	// MobilityMmIfGlobalStatsPath provides the path for retrieving MM interface global statistics.
	MobilityMmIfGlobalStatsPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-mobility-oper:mobility-oper-data/mm-if-global-stats"

	// MobilityClientDataPath provides the path for retrieving mobility client data.
	MobilityClientDataPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-mobility-oper:mobility-oper-data/mobility-client-data"

	// MobilityGlobalStatsPath provides the path for retrieving mobility global statistics.
	MobilityGlobalStatsPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-mobility-oper:mobility-oper-data/mobility-global-stats"

	// MobilityMmIfGlobalMsgStatsPath provides the path for retrieving MM interface global message statistics.
	MobilityMmIfGlobalMsgStatsPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-mobility-oper:mobility-oper-data/mm-if-global-msg-stats"

	// MobilityClientStatsPath provides the path for retrieving mobility client statistics.
	MobilityClientStatsPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-mobility-oper:mobility-oper-data/mobility-client-stats"

	// MobilityGlobalDTLSStatsPath provides the path for retrieving mobility global DTLS statistics.
	MobilityGlobalDTLSStatsPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-mobility-oper:mobility-oper-data/mobility-global-dtls-stats"

	// MobilityGlobalMsgStatsPath provides the path for retrieving mobility global message statistics.
	MobilityGlobalMsgStatsPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-mobility-oper:mobility-oper-data/mobility-global-msg-stats"

	// MobilityWlanClientLimitPath provides the path for retrieving WLAN client limit data.
	MobilityWlanClientLimitPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-mobility-oper:mobility-oper-data/wlan-client-limit"
)
