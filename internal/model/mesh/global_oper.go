package mesh

// MeshGlobalOper represents mesh global operational data response.
type MeshGlobalOper struct {
	MeshGlobalOperData struct {
		MeshGlobalStats MeshGlobalStats  `json:"mesh-global-stats"`           // Summary of mesh AP statistics (Live: IOS-XE 17.12.5)
		MeshApCacInfo   []MeshApCacInfo  `json:"mesh-ap-cac-info,omitempty"`  // Summary of mesh voice call statistics (YANG: IOS-XE 17.12.1)
		MeshApPathInfo  []MeshApPathInfo `json:"mesh-ap-path-info,omitempty"` // Mesh AP path from root AP to mesh AP (YANG: IOS-XE 17.12.1)
		MeshApTreeData  []MeshApTreeData `json:"mesh-ap-tree-data"`           // Mesh AP tree view (Live: IOS-XE 17.12.5)
	} `json:"Cisco-IOS-XE-wireless-mesh-global-oper:mesh-global-oper-data"` // Mesh global operational data (Live: IOS-XE 17.12.5)
}

// MeshGlobalStats represents statistics of mesh APs joined to the controller.
type MeshGlobalStats struct {
	NumOfBridgeAPs      uint32 `json:"num-of-bridge-aps"`       // Number of bridge mode APs (Live: IOS-XE 17.12.5)
	NumOfMAPs           uint32 `json:"num-of-maps"`             // Number of mesh APs (Live: IOS-XE 17.12.5)
	NumOfRAPs           uint32 `json:"num-of-raps"`             // Number of root APs (Live: IOS-XE 17.12.5)
	NumOfFlexBridgeAPs  uint32 `json:"num-of-flex-bridge-aps"`  // Number of flex bridge mode APs (Live: IOS-XE 17.12.5)
	NumOfFlexBridgeMAPs uint32 `json:"num-of-flex-bridge-maps"` // Number of flex mesh APs (Live: IOS-XE 17.12.5)
	NumOfFlexBridgeRAPs uint32 `json:"num-of-flex-bridge-raps"` // Number of flex root APs (Live: IOS-XE 17.12.5)
}

// MeshApCacInfo represents summary of mesh voice call statistics.
type MeshApCacInfo struct {
	WTPName string `json:"wtp-name"` // AP name (YANG: IOS-XE 17.12.1)
}

// MeshApPathInfo represents mesh AP path from root AP to mesh AP.
type MeshApPathInfo struct {
	WTPName string `json:"wtp-name"` // AP name (YANG: IOS-XE 17.12.1)
}

// MeshApTreeData represents mesh AP tree view.
type MeshApTreeData struct {
	SectorNumber uint32       `json:"sector-number"`          // Sector number (Live: IOS-XE 17.12.5)
	WTPMAC       string       `json:"wtp-mac"`                // MAC address of the mesh AP (Live: IOS-XE 17.12.5)
	MeshApCount  uint32       `json:"mesh-ap-count"`          // Number of bridge APs (Live: IOS-XE 17.12.5)
	RAPCount     uint32       `json:"rap-count"`              // Number of root APs (Live: IOS-XE 17.12.5)
	MAPCount     uint32       `json:"map-count"`              // Number of mesh APs (Live: IOS-XE 17.12.5)
	MeshApList   []MeshApInfo `json:"mesh-ap-list,omitempty"` // List of mesh APs with path info (YANG: IOS-XE 17.12.1)
}

// MeshApInfo represents mesh AP tree information.
type MeshApInfo struct {
	APName           string `json:"ap-name"`             // AP name (YANG: IOS-XE 17.12.1)
	APRole           string `json:"ap-role"`             // Mesh AP role (YANG: IOS-XE 17.12.1)
	BridgeGroupName  string `json:"bridge-group-name"`   // Bridge group name configured on this Access Point (YANG: IOS-XE 17.12.1)
	PrefParentAPName string `json:"pref-parent-ap-name"` // Mesh AP parent (YANG: IOS-XE 17.12.1)
}
