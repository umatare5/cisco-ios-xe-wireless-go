package routes

// URWB (Ultra-Reliable Wireless Backhaul) Configuration Paths
//
// These constants define the RESTCONF API paths for URWB configuration
// based on Cisco-IOS-XE-wireless-urwb-cfg YANG model.

// URWB Configuration Paths.
const (
	// URWBCfgPath provides the path for URWB configuration data.
	URWBCfgPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-urwb-cfg:urwb-cfg-data"

	// URWBProfilesPath provides the path for URWB profiles.
	URWBProfilesPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-urwb-cfg:urwb-cfg-data/urwb-profiles"

	// URWBProfileByNamePath provides the path template for URWB profile by name.
	URWBProfileByNamePath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-urwb-cfg:urwb-cfg-data/urwb-profiles/urwb-profile=%s"
)

// URWB Operational Paths.
const (
	// URWBOperPath provides the path for URWB operational data.
	URWBOperPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-urwbnet-oper:urwbnet-oper-data"

	// URWBStatsPath provides the path for URWB statistics.
	URWBStatsPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-urwbnet-oper:urwbnet-oper-data/urwbnet-stats"

	// URWBNodeGroupPath provides the path for URWB node groups.
	URWBNodeGroupPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-urwbnet-oper:urwbnet-oper-data/urwbnet-node-g"
)
