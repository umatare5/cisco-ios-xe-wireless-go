package routes

import "github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf"

// CTS (Cisco TrustSec) Configuration Endpoints
//
// These constants define the RESTCONF API endpoints for CTS SXP configuration
// based on Cisco-IOS-XE-wireless-cts-sxp-cfg YANG model.

const (
	// CTSCfgBasePath defines the base path for CTS configuration endpoints
	CTSCfgBasePath = restconf.YANGModelPrefix + "Cisco-IOS-XE-wireless-cts-sxp-cfg:cts-sxp-cfg-data"
)

// CTS Configuration Endpoints
const (
	// CTSCfgEndpoint retrieves complete CTS configuration data
	CTSCfgEndpoint = CTSCfgBasePath
)
