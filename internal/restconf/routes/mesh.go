package routes

// Mesh Configuration and Operational Paths
//
// These constants define the RESTCONF API paths for mesh configuration
// and operational data based on Cisco-IOS-XE-wireless-mesh YANG models.

// Mesh Operational Paths.
const (
	// MeshOperPath provides the path for retrieving all mesh operational data.
	MeshOperPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-mesh-oper:mesh-oper-data"

	// MeshQueueStatsPath provides the path for retrieving mesh queue statistics.
	MeshQueueStatsPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-mesh-oper:mesh-oper-data/mesh-q-stats"

	// MeshDataRateStatsPath provides the path for retrieving mesh data rate statistics.
	MeshDataRateStatsPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-mesh-oper:mesh-oper-data/mesh-dr-stats"

	// MeshSecurityStatsPath provides the path for retrieving mesh security statistics.
	MeshSecurityStatsPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-mesh-oper:mesh-oper-data/mesh-sec-stats"

	// MeshOperationalDataPath provides the path for retrieving mesh operational data.
	MeshOperationalDataPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-mesh-oper:mesh-oper-data/mesh-oper-data"

	// Legacy paths (compatibility with mesh-global-oper)
	// MeshGlobalStatsPath provides the path for retrieving mesh nodes (using mesh-global-stats).
	MeshGlobalStatsPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-mesh-global-oper:mesh-global-oper-data/mesh-global-stats"

	// MeshApTreeDataPath provides the path for retrieving mesh statistics (using mesh-ap-tree-data).
	MeshApTreeDataPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-mesh-global-oper:mesh-global-oper-data/mesh-ap-tree-data"

	// MeshApCacInfoPath provides the path for retrieving mesh AP CAC information.
	MeshApCacInfoPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-mesh-global-oper:mesh-global-oper-data/mesh-ap-cac-info"

	// MeshApPathInfoPath provides the path for retrieving mesh AP path information.
	MeshApPathInfoPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-mesh-global-oper:mesh-global-oper-data/mesh-ap-path-info"
)

// Mesh Configuration Paths.
const (
	// MeshCfgPath provides the path for retrieving mesh configuration.
	MeshCfgPath = RESTCONFDataPath + "/Cisco-IOS-XE-wireless-mesh-cfg:mesh-cfg-data"
)
