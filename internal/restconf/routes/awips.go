package routes

// AWIPS (Advanced Wireless Intrusion Prevention System) Configuration and Operational Paths
//
// These constants define the RESTCONF API paths for AWIPS configuration and operational
// data based on Cisco-IOS-XE-wireless-awips YANG models.

// AWIPS Configuration Paths.
const (
	// AWIPSCfgPath provides the path for AWIPS configuration data.
	AWIPSCfgPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-awips-cfg:awips-cfg-data"

	// AWIPSProfilesPath provides the path for AWIPS profiles.
	AWIPSProfilesPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-awips-cfg:awips-cfg-data/awips-profiles"

	// AWIPSProfileByNamePath provides the path for AWIPS profile by name.
	AWIPSProfileByNamePath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-awips-cfg:awips-cfg-data/awips-profiles/awips-profile"
)

// AWIPS Operational Paths.
const (
	// AWIPSOperPath provides the path for AWIPS operational data.
	AWIPSOperPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-awips-oper:awips-oper-data"

	// AWIPSPerApInfoPath provides the path for AWIPS per-AP info.
	AWIPSPerApInfoPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-awips-oper:awips-oper-data/awips-per-ap-info"

	// AWIPSApDownloadStatusPath provides the path for AWIPS AP download status.
	AWIPSApDownloadStatusPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-awips-oper:awips-oper-data/awips-ap-dwld-status"
)
