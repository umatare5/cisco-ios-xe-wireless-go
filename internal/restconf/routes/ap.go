package routes

// AP (Access Point) Configuration Paths
//
// These constants define the RESTCONF API paths for access point configuration
// operations based on the Cisco-IOS-XE-wireless-ap-cfg YANG model.

// AP Configuration Paths.
const (
	// APCfgPath retrieves complete access point configuration data.
	APCfgPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-ap-cfg:ap-cfg-data"

	// APTagsPath retrieves access point tag configurations.
	APTagsPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-ap-cfg:ap-cfg-data/ap-tags"

	// APTagPath retrieves access point tag configurations.
	APTagPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-ap-cfg:ap-cfg-data/ap-tag"

	// TagSourcePriorityConfigsPath retrieves tag source priority configurations.
	APTagSourcePriorityConfigsPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-ap-cfg:ap-cfg-data/tag-source-priority-configs"
)

// AP Global Operational Paths.
const (
	// APGlobalOperPath retrieves complete AP global operational data.
	APGlobalOperPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-ap-global-oper:ap-global-oper-data"

	// APEwlcApStatsPath retrieves EWLC AP statistics.
	APEwlcApStatsPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-ap-global-oper:ap-global-oper-data/ewlc-ap-stats"

	// APHistoryPath retrieves AP history data.
	APHistoryPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-ap-global-oper:ap-global-oper-data/ap-history"

	// APJoinStatsPath retrieves AP join statistics data.
	APJoinStatsPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-ap-global-oper:ap-global-oper-data/ap-join-stats"

	// APWlanClientStatsPath retrieves WLAN client statistics data.
	APWlanClientStatsPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-ap-global-oper:ap-global-oper-data/wlan-client-stats"
)

// AP Operational Paths.
const (
	// APOperPath retrieves complete access point operational data.
	APOperPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data"

	// APRadioOperDataPath retrieves radio operational data for access points.
	APRadioOperDataPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/radio-oper-data"

	// APOperDataPath retrieves CAPWAP data for access points.
	APOperDataPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/oper-data"

	// APCapwapDataPath retrieves CAPWAP data for access points.
	APCapwapDataPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/capwap-data"

	// APApNameMacMapPath retrieves AP name to MAC address mapping.
	APApNameMacMapPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/ap-name-mac-map"

	// APRadioNeighborPath retrieves AP radio neighbor information.
	APRadioNeighborPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/ap-radio-neighbor"

	// APImageActiveLocationPath retrieves AP image active location information.
	APImageActiveLocationPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/ap-image-active-location"

	// APImagePrepareLocationPath retrieves AP image prepare location information.
	APImagePrepareLocationPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/ap-image-prepare-location"

	// APPwrInfoPath retrieves AP power information.
	APPwrInfoPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/ap-pwr-info"

	// APSensorStatusPath retrieves AP sensor status information.
	APSensorStatusPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/ap-sensor-status"

	// APCapwapPktsPath retrieves CAPWAP packets information.
	APCapwapPktsPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/capwap-pkts"

	// APIotFirmwarePath retrieves IoT firmware information for access points.
	APIotFirmwarePath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/iot-firmware"
)

// AP RPC Operations
//
// These constants define RPC operations for access point administrative
// state changes and reset operations.

const (
	// APSetApSlotAdminStateRPC defines the RPC for setting AP slot (radio) administrative state.
	APSetApSlotAdminStateRPC = RESTCONFOperationsPath + "/Cisco-IOS-XE-wireless-access-point-cfg-rpc:set-ap-slot-admin-state"

	// APSetApAdminStateRPC defines the RPC for setting AP administrative state.
	APSetApAdminStateRPC = RESTCONFOperationsPath + "/Cisco-IOS-XE-wireless-access-point-cfg-rpc:set-ap-admin-state"

	// APApResetRPC defines the RPC for AP reset operations.
	APApResetRPC = RESTCONFOperationsPath + "/Cisco-IOS-XE-wireless-access-point-cmd-rpc:ap-reset"
)
