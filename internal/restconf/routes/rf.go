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

	// RFTagsPath provides the path for retrieving RF tags.
	RFTagsPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-rf-cfg:rf-cfg-data/rf-tags"

	// RFTagByNamePath provides the path for retrieving RF tag by name.
	RFTagByNamePath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-rf-cfg:rf-cfg-data/rf-tags/rf-tag"

	// MultiBssidProfilesPath provides the path for retrieving Multi-BSSID profiles.
	MultiBssidProfilesPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-rf-cfg:rf-cfg-data/multi-bssid-profiles"

	// AtfPoliciesPath provides the path for retrieving ATF policies.
	AtfPoliciesPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-rf-cfg:rf-cfg-data/atf-policies"

	// RFProfileDefaultEntriesPath provides the path for retrieving RF profile default entries.
	RFProfileDefaultEntriesPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-rf-cfg:rf-cfg-data/rf-profile-default-entries"
)
