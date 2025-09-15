package routes

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
)
