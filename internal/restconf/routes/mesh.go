package routes

import "github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf"

// Mesh Configuration and Operational Endpoints
//
// These constants define the RESTCONF API endpoints for mesh configuration
// and operational data based on Cisco-IOS-XE-wireless-mesh YANG models.

// Mesh Base Paths
const (
	// MeshOperBasePath defines the base path for Mesh operational data endpoints
	MeshOperBasePath = restconf.YANGModelPrefix + "mesh-oper:mesh-oper-data"

	// MeshCfgBasePath defines the base path for Mesh configuration endpoints
	MeshCfgBasePath = restconf.YANGModelPrefix + "mesh-cfg:mesh-cfg-data"
)

// Mesh Operational Endpoints
const (
	// MeshOperEndpoint provides the endpoint for retrieving all mesh operational data
	MeshOperEndpoint = MeshOperBasePath

	// MeshOperNodesEndpoint provides the endpoint for retrieving mesh nodes
	MeshOperNodesEndpoint = MeshOperBasePath + "/mesh-nodes"

	// MeshOperStatsEndpoint provides the endpoint for retrieving mesh statistics
	MeshOperStatsEndpoint = MeshOperBasePath + "/mesh-stats"
)

// Mesh Configuration Endpoints
const (
	// MeshCfgEndpoint provides the endpoint for retrieving mesh configuration
	MeshCfgEndpoint = MeshCfgBasePath

	// MeshCfgProfilesEndpoint provides the endpoint for retrieving mesh profiles
	MeshCfgProfilesEndpoint = MeshCfgBasePath + "/mesh-profiles"
)
