package model

// MeshCfgResponse represents the response structure for Mesh configuration data.
type MeshCfgResponse struct {
	MeshCfgData MeshCfgData `json:"Cisco-IOS-XE-wireless-mesh-cfg:mesh-cfg-data"`
}

// MeshCfgData contains Mesh configuration data
type MeshCfgData struct {
	MeshProfiles []MeshProfile `json:"mesh-profiles"`
}

// MeshProfile represents mesh profile configuration
type MeshProfile struct {
	ProfileName           string `json:"profile-name"`
	SecurityMode          string `json:"security-mode"`
	BackhaulDataRate      string `json:"backhaul-data-rate"`
	BridgeGroupName       string `json:"bridge-group-name"`
	StrictSecurity        bool   `json:"strict-security"`
	BackgroundScanning    bool   `json:"background-scanning"`
	ParentFallbackTimeout int    `json:"parent-fallback-timeout"`
}
