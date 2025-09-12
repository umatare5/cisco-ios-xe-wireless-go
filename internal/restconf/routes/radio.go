package routes

// Radio Configuration Paths
//
// These constants define the RESTCONF API paths for radio configuration
// based on Cisco-IOS-XE-wireless-radio-cfg YANG model.

// Radio Configuration Paths.
const (
	// RadioCfgPath provides the path for retrieving all radio configuration data.
	RadioCfgPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-radio-cfg:radio-cfg-data"

	// RadioProfilesPath provides the path for retrieving radio profiles.
	RadioProfilesPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-radio-cfg:radio-cfg-data/radio-profiles"

	// RadioProfileByNamePath provides the path for retrieving radio profile by name.
	RadioProfileByNamePath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-radio-cfg:radio-cfg-data/radio-profiles/radio-profile"
)
