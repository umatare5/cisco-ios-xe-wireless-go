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
)

// AWIPS Operational Paths.
const (
	// AWIPSOperPath provides the path for AWIPS operational data.
	AWIPSOperPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-awips-oper:awips-oper-data"

	// AWIPSPerApInfoPath provides the path for AWIPS per-AP info.
	AWIPSPerApInfoPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-awips-oper:awips-oper-data/awips-per-ap-info"

	// AWIPSApDownloadStatusPath provides the path for AWIPS AP download status.
	AWIPSApDownloadStatusPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-awips-oper:awips-oper-data/awips-ap-dwld-status"

	// AWIPSDwldStatusPath provides the path for AWIPS download status.
	AWIPSDwldStatusPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-awips-oper:awips-oper-data/awips-dwld-status"

	// AWIPSPerSignStatsPath provides the path for AWIPS per signature statistics.
	AWIPSPerSignStatsPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-awips-oper:awips-oper-data/awips-per-sign-stats"

	// AWIPSGlobStatsPath provides the path for AWIPS global statistics.
	AWIPSGlobStatsPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-awips-oper:awips-oper-data/awips-glob-stats"

	// AWIPSDwldStatusWncdPath provides the path for AWIPS download status for WNCD.
	AWIPSDwldStatusWncdPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-awips-oper:awips-oper-data/awips-dwld-status-wncd"
)
