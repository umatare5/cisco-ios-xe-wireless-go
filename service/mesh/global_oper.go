package mesh

// MeshGlobalOper represents mesh global operational data response.
type MeshGlobalOper struct {
	MeshGlobalOperData struct {
		MeshGlobalStats MeshGlobalStats  `json:"mesh-global-stats"`           // Summary of mesh AP statistics (Live: IOS-XE 17.12.6a)
		MeshApCacInfo   []MeshApCacInfo  `json:"mesh-ap-cac-info,omitempty"`  // Summary of mesh voice call statistics (YANG: IOS-XE 17.12.1)
		MeshApPathInfo  []MeshApPathInfo `json:"mesh-ap-path-info,omitempty"` // Mesh AP path from root AP to mesh AP (YANG: IOS-XE 17.12.1)
		MeshApTreeData  []MeshApTreeData `json:"mesh-ap-tree-data"`           // Mesh AP tree view (Live: IOS-XE 17.12.6a)
	} `json:"Cisco-IOS-XE-wireless-mesh-global-oper:mesh-global-oper-data"` // Mesh global operational data (Live: IOS-XE 17.12.6a)
}

// CiscoIOSXEWirelessMeshOperMeshGlobalStats represents mesh global statistics wrapper.
type CiscoIOSXEWirelessMeshOperMeshGlobalStats struct {
	MeshGlobalStats MeshGlobalStats `json:"Cisco-IOS-XE-wireless-mesh-global-oper:mesh-global-stats"` // Summary of mesh AP statistics (Live: IOS-XE 17.12.6a)
}

// CiscoIOSXEWirelessMeshOperMeshApCacInfo represents mesh AP CAC information wrapper.
type CiscoIOSXEWirelessMeshOperMeshApCacInfo struct {
	MeshApCacInfo []MeshApCacInfo `json:"Cisco-IOS-XE-wireless-mesh-global-oper:mesh-ap-cac-info,omitempty"` // Summary of mesh voice call statistics (YANG: IOS-XE 17.12.1)
}

// CiscoIOSXEWirelessMeshOperMeshApPathInfo represents mesh AP path information wrapper.
type CiscoIOSXEWirelessMeshOperMeshApPathInfo struct {
	MeshApPathInfo []MeshApPathInfo `json:"Cisco-IOS-XE-wireless-mesh-global-oper:mesh-ap-path-info,omitempty"` // Mesh AP path from root AP to mesh AP (YANG: IOS-XE 17.12.1)
}

// CiscoIOSXEWirelessMeshOperMeshApTreeData represents mesh AP tree data wrapper.
type CiscoIOSXEWirelessMeshOperMeshApTreeData struct {
	MeshApTreeData []MeshApTreeData `json:"Cisco-IOS-XE-wireless-mesh-global-oper:mesh-ap-tree-data"` // Mesh AP tree view (Live: IOS-XE 17.12.6a)
}

// MeshGlobalStats represents statistics of mesh APs joined to the controller.
type MeshGlobalStats struct {
	NumOfBridgeAPs      uint32 `json:"num-of-bridge-aps"`       // Number of bridge mode APs (Live: IOS-XE 17.12.6a)
	NumOfMAPs           uint32 `json:"num-of-maps"`             // Number of mesh APs (Live: IOS-XE 17.12.6a)
	NumOfRAPs           uint32 `json:"num-of-raps"`             // Number of root APs (Live: IOS-XE 17.12.6a)
	NumOfFlexBridgeAPs  uint32 `json:"num-of-flex-bridge-aps"`  // Number of flex bridge mode APs (Live: IOS-XE 17.12.6a)
	NumOfFlexBridgeMAPs uint32 `json:"num-of-flex-bridge-maps"` // Number of flex mesh APs (Live: IOS-XE 17.12.6a)
	NumOfFlexBridgeRAPs uint32 `json:"num-of-flex-bridge-raps"` // Number of flex root APs (Live: IOS-XE 17.12.6a)
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
	SectorNumber uint32       `json:"sector-number"`          // Sector number (Live: IOS-XE 17.12.6a)
	WTPMAC       string       `json:"wtp-mac"`                // MAC address of the mesh AP (Live: IOS-XE 17.12.6a)
	MeshApCount  uint32       `json:"mesh-ap-count"`          // Number of bridge APs (Live: IOS-XE 17.12.6a)
	RAPCount     uint32       `json:"rap-count"`              // Number of root APs (Live: IOS-XE 17.12.6a)
	MAPCount     uint32       `json:"map-count"`              // Number of mesh APs (Live: IOS-XE 17.12.6a)
	MeshApList   []MeshApInfo `json:"mesh-ap-list,omitempty"` // List of mesh APs with path info (YANG: IOS-XE 17.12.1)
}

// MeshApInfo represents mesh AP tree information.
type MeshApInfo struct {
	APName           string `json:"ap-name"`             // AP name (YANG: IOS-XE 17.12.1)
	APRole           string `json:"ap-role"`             // Mesh AP role (YANG: IOS-XE 17.12.1)
	BridgeGroupName  string `json:"bridge-group-name"`   // Bridge group name configured on this Access Point (YANG: IOS-XE 17.12.1)
	PrefParentAPName string `json:"pref-parent-ap-name"` // Mesh AP parent (YANG: IOS-XE 17.12.1)
}
