package routes

import "github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf"

// Mobility Operational Endpoints
//
// These constants define the RESTCONF API endpoints for mobility operational
// data based on Cisco-IOS-XE-wireless-mobility-oper YANG model.

const (
	// MobilityOperBasePath defines the base path for Mobility operational data endpoints
	MobilityOperBasePath = restconf.YANGModelPrefix + "mobility-oper:mobility-oper-data"
)

// Mobility Operational Endpoints
const (
	// MobilityOperEndpoint provides the endpoint for retrieving all mobility operational data
	MobilityOperEndpoint = MobilityOperBasePath

	// MobilityOperApCacheEndpoint provides the endpoint for retrieving AP cache data
	MobilityOperApCacheEndpoint = MobilityOperBasePath + "/ap-cache"

	// MobilityOperApPeerListEndpoint provides the endpoint for retrieving AP peer list
	MobilityOperApPeerListEndpoint = MobilityOperBasePath + "/ap-peer-list"

	// MobilityOperMmGlobalDataEndpoint provides the endpoint for retrieving MM global data
	MobilityOperMmGlobalDataEndpoint = MobilityOperBasePath + "/mm-global-data"

	// MobilityOperMmIfGlobalStatsEndpoint provides the endpoint for retrieving MM interface global statistics
	MobilityOperMmIfGlobalStatsEndpoint = MobilityOperBasePath + "/mm-if-global-stats"

	// MobilityOperMobilityClientDataEndpoint provides the endpoint for retrieving mobility client data
	MobilityOperMobilityClientDataEndpoint = MobilityOperBasePath + "/mobility-client-data"

	// MobilityOperMobilityGlobalStatsEndpoint provides the endpoint for retrieving mobility global statistics
	MobilityOperMobilityGlobalStatsEndpoint = MobilityOperBasePath + "/mobility-global-stats"
)
