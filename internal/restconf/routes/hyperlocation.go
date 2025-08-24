package routes

import "github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf"

// Hyperlocation Operational Endpoints
//
// These constants define the RESTCONF API endpoints for hyperlocation operational
// data based on Cisco-IOS-XE-wireless-hyperlocation-oper YANG model.

const (
	// HyperlocationOperBasePath defines the base path for hyperlocation operational data endpoints
	HyperlocationOperBasePath = restconf.YANGModelPrefix + "hyperlocation-oper:hyperlocation-oper-data"
)

// Hyperlocation Operational Endpoints
const (
	// HyperlocationOperEndpoint defines the endpoint for hyperlocation operational data
	HyperlocationOperEndpoint = HyperlocationOperBasePath

	// HyperlocationProfilesEndpoint defines the endpoint for hyperlocation profiles
	HyperlocationProfilesEndpoint = HyperlocationOperBasePath + "/ewlc-hyperlocation-profile"
)
