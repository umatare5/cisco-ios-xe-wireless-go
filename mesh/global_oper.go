// Package mesh provides mesh networking global operational data functionality for the Cisco Wireless Network Controller API.
package mesh

import (
	"context"

	wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
)

const (
	// MeshGlobalOperBasePath defines the base path for mesh global operational data endpoints.
	MeshGlobalOperBasePath = "/restconf/data/Cisco-IOS-XE-wireless-mesh-global-oper:mesh-global-oper-data"
	// MeshGlobalOperEndpoint defines the endpoint for mesh global operational data.
	MeshGlobalOperEndpoint = MeshGlobalOperBasePath
	// MeshGlobalStatsEndpoint defines the endpoint for mesh global statistics.
	MeshGlobalStatsEndpoint = MeshGlobalOperBasePath + "/mesh-global-stats"
	// MeshApTreeDataEndpoint defines the endpoint for mesh AP tree data.
	MeshApTreeDataEndpoint = MeshGlobalOperBasePath + "/mesh-ap-tree-data"
)

// MeshGlobalOperResponse represents the response structure for mesh global operational data.
type MeshGlobalOperResponse struct {
	CiscoIOSXEWirelessMeshGlobalOperMeshGlobalOperData struct {
		MeshGlobalStats MeshGlobalStats  `json:"mesh-global-stats"`
		MeshApTreeData  []MeshApTreeData `json:"mesh-ap-tree-data"`
	} `json:"Cisco-IOS-XE-wireless-mesh-global-oper:mesh-global-oper-data"`
}

// MeshGlobalStatsResponse represents the response structure for mesh global statistics.
type MeshGlobalStatsResponse struct {
	MeshGlobalStats MeshGlobalStats `json:"Cisco-IOS-XE-wireless-mesh-global-oper:mesh-global-stats"`
}

// MeshApTreeDataResponse represents the response structure for mesh AP tree data.
type MeshApTreeDataResponse struct {
	MeshApTreeData []MeshApTreeData `json:"Cisco-IOS-XE-wireless-mesh-global-oper:mesh-ap-tree-data"`
}

// MeshGlobalStats represents mesh global statistics including counts of various AP types.
type MeshGlobalStats struct {
	NumOfBridgeAps      int `json:"num-of-bridge-aps"`
	NumOfMaps           int `json:"num-of-maps"`
	NumOfRaps           int `json:"num-of-raps"`
	NumOfFlexBridgeAps  int `json:"num-of-flex-bridge-aps"`
	NumOfFlexBridgeMaps int `json:"num-of-flex-bridge-maps"`
	NumOfFlexBridgeRaps int `json:"num-of-flex-bridge-raps"`
}

// MeshApTreeData represents mesh AP tree data including sector information and AP counts.
type MeshApTreeData struct {
	SectorNumber int    `json:"sector-number"`
	WtpMac       string `json:"wtp-mac"`
	MeshApCount  int    `json:"mesh-ap-count"`
	RapCount     int    `json:"rap-count"`
	MapCount     int    `json:"map-count"`
}

// GetMeshGlobalOper retrieves mesh global operational data.
func GetMeshGlobalOper(client *wnc.Client, ctx context.Context) (*MeshGlobalOperResponse, error) {
	var data MeshGlobalOperResponse
	if err := client.SendAPIRequest(ctx, MeshGlobalOperEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetMeshGlobalStats retrieves mesh global statistics.
func GetMeshGlobalStats(client *wnc.Client, ctx context.Context) (*MeshGlobalStatsResponse, error) {
	var data MeshGlobalStatsResponse
	if err := client.SendAPIRequest(ctx, MeshGlobalStatsEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetMeshApTreeData retrieves mesh AP tree data.
func GetMeshApTreeData(client *wnc.Client, ctx context.Context) (*MeshApTreeDataResponse, error) {
	var data MeshApTreeDataResponse
	if err := client.SendAPIRequest(ctx, MeshApTreeDataEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}
