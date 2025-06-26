// Package mesh provides mesh networking configuration functionality for the Cisco Wireless Network Controller API.
package mesh

import (
	"context"

	wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
)

const (
	// MeshCfgBasePath defines the base path for mesh configuration endpoints
	MeshCfgBasePath = "/restconf/data/Cisco-IOS-XE-wireless-mesh-cfg:mesh-cfg-data"
	// MeshCfgEndpoint retrieves complete mesh configuration data
	MeshCfgEndpoint = MeshCfgBasePath
	// MeshMeshEndpoint retrieves mesh global configuration
	MeshMeshEndpoint = MeshCfgBasePath + "/mesh"
	// MeshProfilesEndpoint retrieves mesh profiles
	MeshProfilesEndpoint = MeshCfgBasePath + "/mesh-profiles"
)

// MeshCfgResponse represents the complete mesh configuration response
type MeshCfgResponse struct {
	CiscoIOSXEWirelessMeshCfgMeshCfgData struct {
		Mesh         struct{}     `json:"mesh"`
		MeshProfiles MeshProfiles `json:"mesh-profiles"`
	} `json:"Cisco-IOS-XE-wireless-mesh-cfg:mesh-cfg-data"`
}

// MeshResponse represents the mesh global configuration response
type MeshResponse struct {
	Mesh struct{} `json:"Cisco-IOS-XE-wireless-mesh-cfg:mesh"`
}

// MeshProfilesResponse represents the mesh profiles configuration response
type MeshProfilesResponse struct {
	MeshProfiles MeshProfiles `json:"Cisco-IOS-XE-wireless-mesh-cfg:mesh-profiles"`
}

// MeshProfiles contains mesh profile configuration entries
type MeshProfiles struct {
	MeshProfile []MeshProfile `json:"mesh-profile"`
}

// MeshProfile represents a mesh profile configuration
type MeshProfile struct {
	ProfileName string `json:"profile-name"`          // Mesh profile name
	Description string `json:"description,omitempty"` // Optional profile description
}

// GetMeshCfg retrieves complete mesh configuration data.
// Returns mesh configuration including global settings and profiles.
func GetMeshCfg(client *wnc.Client, ctx context.Context) (*MeshCfgResponse, error) {
	var data MeshCfgResponse
	err := client.SendAPIRequest(ctx, MeshCfgEndpoint, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// GetMesh retrieves mesh global configuration data.
// Returns global mesh configuration settings.
func GetMesh(client *wnc.Client, ctx context.Context) (*MeshResponse, error) {
	var data MeshResponse
	err := client.SendAPIRequest(ctx, MeshMeshEndpoint, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// GetMeshProfiles retrieves mesh profile configurations.
// Returns mesh profile settings and descriptions.
func GetMeshProfiles(client *wnc.Client, ctx context.Context) (*MeshProfilesResponse, error) {
	var data MeshProfilesResponse
	err := client.SendAPIRequest(ctx, MeshProfilesEndpoint, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
