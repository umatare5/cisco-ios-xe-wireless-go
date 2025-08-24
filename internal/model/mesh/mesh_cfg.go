// Package model provides data models for mesh configuration data.
package model

// MeshCfg  represents the structure for Mesh configuration data.
type MeshCfg struct {
	MeshCfgData MeshCfgData `json:"Cisco-IOS-XE-wireless-mesh-cfg:mesh-cfg-data"`
}

// MeshCfgMeshProfiles  represents the structure for mesh profiles.
type MeshCfgMeshProfiles struct {
	MeshProfiles []MeshProfile `json:"Cisco-IOS-XE-wireless-mesh-cfg:mesh-profiles"`
}

type MeshCfgData struct {
	MeshProfiles []MeshProfile `json:"mesh-profiles"`
}

type MeshProfile struct {
	ProfileName           string `json:"profile-name"`
	SecurityMode          string `json:"security-mode"`
	BackhaulDataRate      string `json:"backhaul-data-rate"`
	BridgeGroupName       string `json:"bridge-group-name"`
	StrictSecurity        bool   `json:"strict-security"`
	BackgroundScanning    bool   `json:"background-scanning"`
	ParentFallbackTimeout int    `json:"parent-fallback-timeout"`
}
