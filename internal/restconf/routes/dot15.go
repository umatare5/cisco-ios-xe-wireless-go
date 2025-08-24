package routes

import "github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf"

// 802.15 Configuration Endpoints
//
// These constants define the RESTCONF API endpoints for 802.15 standard
// configuration based on Cisco-IOS-XE-wireless-dot15-cfg YANG model.

const (
	// Dot15CfgBasePath defines the base path for 802.15 configuration endpoints
	Dot15CfgBasePath = restconf.YANGModelPrefix + "dot15-cfg:dot15-cfg-data"
)

// 802.15 Configuration Endpoints
const (
	// Dot15CfgEndpoint retrieves complete 802.15 configuration data
	Dot15CfgEndpoint = Dot15CfgBasePath
)
