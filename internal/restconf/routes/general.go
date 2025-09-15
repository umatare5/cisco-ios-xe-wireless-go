package routes

// General Configuration and Operational Paths
//
// These constants define the RESTCONF API paths for general wireless
// configuration and operational data based on Cisco-IOS-XE-wireless-general YANG models.

// General Operational Paths.
const (
	// GeneralOperPath defines the path for general operational data.
	GeneralOperPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-general-oper:general-oper-data"

	// GeneralMgmtIntfDataPath defines the path for management interface data.
	GeneralMgmtIntfDataPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-general-oper:general-oper-data/mgmt-intf-data"
)

// General Configuration Paths.
const (
	// GeneralCfgPath defines the path for general configuration data.
	GeneralCfgPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-general-cfg:general-cfg-data"

	// GeneralMewlcConfigPath defines the path for MEWLC configuration.
	GeneralMewlcConfigPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-general-cfg:general-cfg-data/mewlc-config"

	// GeneralCacConfigPath defines the path for CAC configuration.
	GeneralCacConfigPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-general-cfg:general-cfg-data/cac-config"

	// GeneralMfpPath defines the path for MFP configuration.
	GeneralMfpPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-general-cfg:general-cfg-data/mfp"

	// GeneralFipsCfgPath defines the path for FIPS configuration.
	GeneralFipsCfgPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-general-cfg:general-cfg-data/fips-cfg"

	// GeneralWsaApClientEventPath defines the path for WSA AP client event configuration.
	GeneralWsaApClientEventPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-general-cfg:general-cfg-data/wsa-ap-client-event"

	// GeneralSimL3InterfaceCacheDataPath defines the path for SIM L3 interface cache data.
	GeneralSimL3InterfaceCacheDataPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-general-cfg:general-cfg-data/sim-l3-interface-cache-data"

	// GeneralWlcManagementDataPath defines the path for WLC management data.
	GeneralWlcManagementDataPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-general-cfg:general-cfg-data/wlc-management-data"

	// GeneralLaginfoPath defines the path for LAG information.
	GeneralLaginfoPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-general-cfg:general-cfg-data/laginfo"

	// GeneralMulticastConfigPath defines the path for multicast configuration.
	GeneralMulticastConfigPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-general-cfg:general-cfg-data/multicast-config"

	// GeneralFeatureUsageCfgPath defines the path for feature usage configuration.
	GeneralFeatureUsageCfgPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-general-cfg:general-cfg-data/feature-usage-cfg"

	// GeneralThresholdWarnCfgPath defines the path for threshold warning configuration.
	GeneralThresholdWarnCfgPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-general-cfg:general-cfg-data/threshold-warn-cfg"

	// GeneralApLocRangingCfgPath defines the path for AP location ranging configuration.
	GeneralApLocRangingCfgPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-general-cfg:general-cfg-data/ap-loc-ranging-cfg"

	// GeneralGeolocationCfgPath defines the path for geolocation configuration.
	GeneralGeolocationCfgPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-general-cfg:general-cfg-data/geolocation-cfg"
)
