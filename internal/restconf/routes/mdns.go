package routes

// mDNS (Multicast DNS) Operational Paths
//
// These constants define the RESTCONF API paths for mDNS operational
// data based on Cisco-IOS-XE-wireless-mdns-oper YANG model.

// mDNS Operational Paths.
const (
	// MDNSOperPath provides the base path for mDNS operational data.
	MDNSOperPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-mdns-oper:mdns-oper-data"

	// MDNSGlobalStatsPath provides the path for mDNS global statistics.
	MDNSGlobalStatsPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-mdns-oper:mdns-oper-data/mdns-global-stats"

	// MDNSWlanStatsPath provides the path for mDNS WLAN statistics.
	MDNSWlanStatsPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-mdns-oper:mdns-oper-data/mdns-wlan-stats"
)
