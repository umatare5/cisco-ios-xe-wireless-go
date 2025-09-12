package routes

// Hyperlocation Operational Paths
//
// These constants define the RESTCONF API paths for hyperlocation operational
// data based on Cisco-IOS-XE-wireless-hyperlocation-oper YANG model.

// Hyperlocation Operational Paths.
const (
	// HyperlocationOperPath defines the path for hyperlocation operational data.
	HyperlocationOperPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-hyperlocation-oper:hyperlocation-oper-data"

	// HyperlocationProfilesPath defines the path for hyperlocation profiles.
	HyperlocationProfilesPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-hyperlocation-oper:hyperlocation-oper-data/ewlc-hyperlocation-profile"

	// HyperlocationProfileByNamePath defines the path for hyperlocation profile by name.
	HyperlocationProfileByNamePath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-hyperlocation-oper:hyperlocation-oper-data/ewlc-hyperlocation-profile"
)
