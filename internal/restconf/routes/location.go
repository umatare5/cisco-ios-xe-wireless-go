package routes

import "github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf"

// Location Configuration and Operational Endpoints
//
// These constants define the RESTCONF API endpoints for location configuration
// and operational data based on Cisco-IOS-XE-wireless-location YANG models.

// Location Base Paths
const (
	// LocationCfgBasePath defines the base path for Location configuration endpoints
	LocationCfgBasePath = restconf.YANGModelPrefix + "location-cfg:location-cfg-data"

	// LocationOperBasePath defines the base path for Location operational endpoints
	LocationOperBasePath = restconf.YANGModelPrefix + "location-oper:location-oper-data"
)

// Location Configuration Endpoints
const (
	// LocationCfgEndpoint provides the endpoint for retrieving all location configuration data
	LocationCfgEndpoint = LocationCfgBasePath

	// LocationCfgProfilesEndpoint provides the endpoint for retrieving location profiles
	LocationCfgProfilesEndpoint = LocationCfgBasePath + "/location-cfg/profiles"

	// LocationCfgServersEndpoint provides the endpoint for retrieving location servers
	LocationCfgServersEndpoint = LocationCfgBasePath + "/location-cfg/servers"

	// LocationCfgSettingsEndpoint provides the endpoint for retrieving location settings
	LocationCfgSettingsEndpoint = LocationCfgBasePath + "/location-cfg/settings"
)

// Location Operational Endpoints
const (
	// LocationOperEndpoint provides the endpoint for retrieving location operational data
	LocationOperEndpoint = LocationOperBasePath

	// LocationOperStatsEndpoint provides the endpoint for retrieving location statistics
	LocationOperStatsEndpoint = LocationOperBasePath + "/location-oper/statistics"
)
