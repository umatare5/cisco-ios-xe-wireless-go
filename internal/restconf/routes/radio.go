package routes

import "github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf"

// Radio Configuration Endpoints
//
// These constants define the RESTCONF API endpoints for radio configuration
// based on Cisco-IOS-XE-wireless-radio-cfg YANG model.

const (
	// RadioCfgBasePath defines the base path for Radio configuration endpoints
	RadioCfgBasePath = restconf.YANGModelPrefix + "radio-cfg:radio-cfg-data"
)

// Radio Configuration Endpoints
const (
	// RadioCfgEndpoint provides the endpoint for retrieving all radio configuration data
	RadioCfgEndpoint = RadioCfgBasePath

	// RadioCfgProfilesEndpoint provides the endpoint for retrieving radio profiles
	RadioCfgProfilesEndpoint = RadioCfgBasePath + "/radio-profiles"
)
