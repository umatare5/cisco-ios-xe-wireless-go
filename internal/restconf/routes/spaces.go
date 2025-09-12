package routes

// Spaces Configuration Paths
//
// These constants define the RESTCONF API paths for Spaces configuration
// based on Cisco-IOS-XE-wireless-spaces-cfg YANG model.

// Spaces Configuration Paths.
const (
	// SpacesCfgPath provides the path for Spaces configuration data.
	SpacesCfgPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-spaces-cfg:spaces-cfg-data"

	// SpacesProfilesPath provides the path for Spaces profiles.
	SpacesProfilesPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-spaces-cfg:spaces-cfg-data/spaces-profiles"

	// SpacesProfileByNamePath provides the path for Spaces profile by name.
	SpacesProfileByNamePath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-spaces-cfg:spaces-cfg-data/spaces-profiles/spaces-profile"
)

// Spaces Operational Paths.
const (
	// SpacesOperPath provides the path for Spaces operational data.
	SpacesOperPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-cisco-spaces-oper:cisco-spaces-oper-data"

	// SpacesConnectionDetailPath provides the path for Spaces connection details.
	SpacesConnectionDetailPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-cisco-spaces-oper:cisco-spaces-oper-data/spaces-connection-detail"
)
