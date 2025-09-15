package routes

// Rogue Detection Operational Paths
//
// These constants define the RESTCONF API paths for rogue detection
// operational data based on Cisco-IOS-XE-wireless-rogue-oper YANG model.

// Rogue Operational Paths.
const (
	// RogueOperPath provides the path for rogue operational data.
	RogueOperPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-rogue-oper:rogue-oper-data"

	// RogueStatsPath provides the path for rogue statistics.
	RogueStatsPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-rogue-oper:rogue-oper-data/rogue-stats"

	// RogueDataPath provides the path for rogue data.
	RogueDataPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-rogue-oper:rogue-oper-data/rogue-data"

	// RogueClientDataPath provides the path for rogue client data.
	RogueClientDataPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-rogue-oper:rogue-oper-data/rogue-client-data"
)
