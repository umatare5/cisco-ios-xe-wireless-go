package model

// MeshOper  represents the structure for Mesh operational data.
type MeshOper struct {
	MeshOperData MeshOperData `json:"Cisco-IOS-XE-wireless-mesh-oper:mesh-oper-data"`
}

// MeshOperMeshNodes  represents the structure for Mesh nodes.
type MeshOperMeshNodes struct {
	MeshNodes []MeshNode `json:"Cisco-IOS-XE-wireless-mesh-oper:mesh-nodes"`
}

// MeshOperMeshStats  represents the structure for Mesh statistics.
type MeshOperMeshStats struct {
	MeshStats MeshStats `json:"Cisco-IOS-XE-wireless-mesh-oper:mesh-stats"`
}

type MeshOperData struct {
	MeshNodes []MeshNode `json:"mesh-nodes"`
	MeshStats MeshStats  `json:"mesh-stats"`
}

type MeshNode struct {
	MacAddress    string   `json:"mac-address"`
	ParentMac     string   `json:"parent-mac"`
	NodeRole      string   `json:"node-role"`
	HopCount      int      `json:"hop-count"`
	LinkSNR       int      `json:"link-snr"`
	ChildNodes    []string `json:"child-nodes"`
	LinkState     string   `json:"link-state"`
	BackhaulUsage int      `json:"backhaul-usage"`
}

type MeshStats struct {
	TotalNodes          int `json:"total-nodes"`
	RootNodes           int `json:"root-nodes"`
	MapNodes            int `json:"map-nodes"`
	ChildNodes          int `json:"child-nodes"`
	BackhaulUtilization int `json:"backhaul-utilization"`
}
