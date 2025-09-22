package routes

// CTS (Cisco TrustSec) Configuration Paths
//
// These constants define the RESTCONF API paths for CTS SXP configuration
// based on Cisco-IOS-XE-wireless-cts-sxp-cfg YANG model.

// CTS Configuration Paths.
const (
	// CTSCfgPath retrieves complete CTS configuration data.
	CTSCfgPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-cts-sxp-cfg:cts-sxp-cfg-data"
)

// CTS Operational Paths.
const (
	// CTSOperPath retrieves complete CTS operational data.
	CTSOperPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-cts-sxp-oper:cts-sxp-oper-data"

	// CTSFlexModeApSxpConnectionStatusPath retrieves FlexConnect AP SXP connection status.
	CTSFlexModeApSxpConnectionStatusPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-cts-sxp-oper:cts-sxp-oper-data/flex-mode-ap-sxp-connection-status"
)
