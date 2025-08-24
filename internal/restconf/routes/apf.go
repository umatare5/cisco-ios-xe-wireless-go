package routes

import "github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf"

// APF (Application Policy Framework) Configuration Endpoints
//
// These constants define the RESTCONF API endpoints for Application Policy Framework
// configuration based on Cisco-IOS-XE-wireless-apf-cfg YANG model.

const (
	// APFCfgBasePath defines the base path for Application Policy Framework configuration endpoints
	APFCfgBasePath = restconf.YANGModelPrefix + "apf-cfg:apf-cfg-data"
)

// APF Configuration Endpoints
const (
	// APFCfgEndpoint retrieves complete Application Policy Framework configuration data
	APFCfgEndpoint = APFCfgBasePath
)
