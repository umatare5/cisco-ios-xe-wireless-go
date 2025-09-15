package routes

// Mesh Configuration and Operational Paths
//
// These constants define the RESTCONF API paths for mesh configuration
// and operational data based on Cisco-IOS-XE-wireless-mesh YANG models.

// Mesh Operational Paths.
const (
	// MeshOperPath provides the path for retrieving all mesh operational data.
	MeshOperPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-mesh-global-oper:mesh-global-oper-data"

	// MeshGlobalStatsPath provides the path for retrieving mesh nodes (using mesh-global-stats).
	MeshGlobalStatsPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-mesh-global-oper:mesh-global-oper-data/mesh-global-stats"

	// MeshApTreeDataPath provides the path for retrieving mesh statistics (using mesh-ap-tree-data).
	MeshApTreeDataPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-mesh-global-oper:mesh-global-oper-data/mesh-ap-tree-data"
)

// Mesh Configuration Paths.
const (
	// MeshCfgPath provides the path for retrieving mesh configuration.
	MeshCfgPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-mesh-cfg:mesh-cfg-data"

	// MeshProfilesPath provides the path for retrieving mesh profiles.
	MeshProfilesPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-mesh-cfg:mesh-cfg-data/mesh-profiles/mesh-profile"
)
