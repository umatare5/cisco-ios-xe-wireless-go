package routes

// CTS (Cisco TrustSec) Configuration Paths
//
// These constants define the RESTCONF API paths for CTS SXP configuration
// based on Cisco-IOS-XE-wireless-cts-sxp-cfg YANG model.

// CTS Configuration Paths.
const (
	// CTSCfgPath retrieves complete CTS configuration data.
	CTSCfgPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-cts-sxp-cfg:cts-sxp-cfg-data"

	// CTSSxpConfigPath retrieves CTS SXP configuration entries.
	CTSSxpConfigPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-cts-sxp-cfg:cts-sxp-cfg-data/cts-sxp-configuration/cts-sxp-config"
)
