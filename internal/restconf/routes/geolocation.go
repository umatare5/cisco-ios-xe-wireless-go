package routes

import "github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf"

// Geolocation Operational Endpoints
//
// These constants define the RESTCONF API endpoints for geolocation operational
// data based on Cisco-IOS-XE-wireless-geolocation-oper YANG model.

const (
	// GeolocationOperBasePath defines the base path for geolocation operational endpoints
	GeolocationOperBasePath = restconf.YANGModelPrefix + "geolocation-oper:geolocation-oper-data"
)

// Geolocation Operational Endpoints
const (
	// GeolocationOperEndpoint defines the endpoint for geolocation operational data
	GeolocationOperEndpoint = GeolocationOperBasePath

	// ApGeoLocStatsEndpoint defines the endpoint for AP geolocation statistics
	ApGeoLocStatsEndpoint = GeolocationOperBasePath + "/ap-geo-loc-stats"
)
