package model

// MeshOperResponse represents the response structure for Mesh operational data.
type MeshOperResponse struct {
	MeshOperData MeshOperData `json:"Cisco-IOS-XE-wireless-mesh-oper:mesh-oper-data"`
}

// MeshOperData contains Mesh operational data
type MeshOperData struct {
	MeshNodes []MeshNode `json:"mesh-nodes"`
	MeshStats MeshStats  `json:"mesh-stats"`
}

// MeshNode represents mesh node information
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

// MeshStats represents mesh statistics
type MeshStats struct {
	TotalNodes          int `json:"total-nodes"`
	RootNodes           int `json:"root-nodes"`
	MapNodes            int `json:"map-nodes"`
	ChildNodes          int `json:"child-nodes"`
	BackhaulUtilization int `json:"backhaul-utilization"`
}
