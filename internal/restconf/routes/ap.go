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

	// APApNameMACMapPath retrieves AP name to MAC address mapping.
	APApNameMACMapPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/ap-name-mac-map"

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

	// APRadioResetStatsPath retrieves radio reset statistics information.
	APRadioResetStatsPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/radio-reset-stats"

	// APQosClientDataPath retrieves QoS client data information.
	APQosClientDataPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/qos-client-data"

	// APWtpSlotWlanStatsPath retrieves WTP slot WLAN statistics information.
	APWtpSlotWlanStatsPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/wtp-slot-wlan-stats"

	// APEthernetMACWtpMACMapPath retrieves Ethernet MAC to WTP MAC mapping information.
	APEthernetMACWtpMACMapPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/ethernet-mac-wtp-mac-map"

	// APRadioOperStatsPath retrieves radio operational statistics information.
	APRadioOperStatsPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/radio-oper-stats"

	// APEthernetIfStatsPath retrieves Ethernet interface statistics information.
	APEthernetIfStatsPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/ethernet-if-stats"

	// APEwlcWncdStatsPath retrieves EWLC WNCD statistics information.
	APEwlcWncdStatsPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/ewlc-wncd-stats"

	// APApIoxOperDataPath retrieves AP IOx operational data information.
	APApIoxOperDataPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/ap-iox-oper-data"

	// APQosGlobalStatsPath retrieves QoS global statistics information.
	APQosGlobalStatsPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/qos-global-stats"

	// APRlanOperPath retrieves RLAN operational data information.
	APRlanOperPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/rlan-oper"

	// APEwlcMewlcPredownloadRecPath retrieves EWLC MEWLC predownload record information.
	APEwlcMewlcPredownloadRecPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/ewlc-mewlc-predownload-rec"

	// APCdpCacheDataPath retrieves CDP cache data information.
	APCdpCacheDataPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/cdp-cache-data"

	// APLldpNeighPath retrieves LLDP neighbor information.
	APLldpNeighPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/lldp-neigh"

	// APTpCertInfoPath retrieves trustpoint certificate info information.
	APTpCertInfoPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/tp-cert-info"

	// APDiscDataPath retrieves discovery data information.
	APDiscDataPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/disc-data"

	// APCountryOperPath retrieves country operational data information.
	APCountryOperPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/country-oper"

	// APSuppCountryOperPath retrieves supported country operational data information.
	APSuppCountryOperPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/supp-country-oper"

	// APApNhGlobalDataPath retrieves AP neighborhood global data information.
	APApNhGlobalDataPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/ap-nh-global-data"
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

// AP Query Paths.
const (
	// APTagQueryPath provides the path for querying AP tag by MAC address.
	APTagQueryPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-ap-cfg:ap-cfg-data/ap-tags/ap-tag"

	// APTagSourcePriorityConfigQueryPath provides the path for querying tag source priority config by priority.
	APTagSourcePriorityConfigQueryPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-ap-cfg:ap-cfg-data/tag-source-priority-configs/tag-source-priority-config"

	// APHistoryQueryPath provides the path for querying AP history by ethernet MAC.
	APHistoryQueryPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-ap-global-oper:ap-global-oper-data/ap-history"

	// APJoinStatsQueryPath provides the path for querying AP join statistics by WTP MAC.
	APJoinStatsQueryPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-ap-global-oper:ap-global-oper-data/ap-join-stats"

	// APWlanClientStatsQueryPath provides the path for querying WLAN client statistics by WTP MAC.
	APWlanClientStatsQueryPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-ap-global-oper:ap-global-oper-data/wlan-client-stats"
)
