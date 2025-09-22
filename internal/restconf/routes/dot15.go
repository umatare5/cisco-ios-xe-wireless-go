package routes

// 802.15 Configuration Paths
//
// These constants define the RESTCONF API paths for 802.15 standard
// configuration based on Cisco-IOS-XE-wireless-dot15-cfg YANG model.

// 802.15 Configuration Paths.
const (
	// Dot15CfgPath retrieves complete 802.15 configuration data.
	Dot15CfgPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-dot15-cfg:dot15-cfg-data"

	// Dot15GlobalConfigPath retrieves 802.15 global configuration.
	Dot15GlobalConfigPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-dot15-cfg:dot15-cfg-data/dot15-global-config"
)
