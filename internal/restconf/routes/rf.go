package routes

// RF (Radio Frequency) Profile Configuration Paths
//
// These constants define the RESTCONF API paths for RF profile configuration
// based on Cisco-IOS-XE-wireless-rf-cfg YANG model.

// RF Configuration Paths.
const (
	// RFCfgPath provides the path for retrieving all RF configuration data.
	RFCfgPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-rf-cfg:rf-cfg-data"

	// RFProfilesPath provides the path for retrieving RF profiles.
	RFProfilesPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-rf-cfg:rf-cfg-data/rf-profiles"

	// RFProfileByNamePath provides the path for retrieving RF profile by name.
	RFProfileByNamePath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-rf-cfg:rf-cfg-data/rf-profiles/rf-profile"

	// RFTagsPath provides the path for retrieving RF tags.
	RFTagsPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-rf-cfg:rf-cfg-data/rf-tags"
)
