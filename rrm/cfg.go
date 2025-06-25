// Package rrm provides Radio Resource Management configuration functionality for the Cisco Wireless Network Controller API.
package rrm

import (
	"context"

	wnc "github.com/umatare5/cisco-xe-wireless-restconf-go"
)

const (
	// RrmCfgBasePath defines the base path for RRM configuration endpoints
	RrmCfgBasePath = "/restconf/data/Cisco-IOS-XE-wireless-rrm-cfg:rrm-cfg-data"
	// RrmCfgEndpoint retrieves complete RRM configuration data
	RrmCfgEndpoint = RrmCfgBasePath
	// RrmCfgRrmsEndpoint retrieves RRM configuration entries
	RrmCfgRrmsEndpoint = RrmCfgBasePath + "/rrms"
	// RrmCfgRrmMgrCfgEntriesEndpoint retrieves RRM manager configuration entries
	RrmCfgRrmMgrCfgEntriesEndpoint = RrmCfgBasePath + "/rrm-mgr-cfg-entries"
)

// RrmCfgResponse represents the complete RRM configuration response
type RrmCfgResponse struct {
	CiscoIOSXEWirelessRrmCfgRrmCfgData struct {
		Rrms             Rrms             `json:"rrms"`
		RrmMgrCfgEntries RrmMgrCfgEntries `json:"rrm-mgr-cfg-entries"`
	} `json:"Cisco-IOS-XE-wireless-rrm-cfg:rrm-cfg-data"`
}

// RrmRrmsResponse represents the RRM configuration entries response
type RrmRrmsResponse struct {
	Rrms Rrms `json:"Cisco-IOS-XE-wireless-rrm-cfg:rrms"`
}

// RrmMgrCfgEntriesResponse represents the RRM manager configuration entries response
type RrmMgrCfgEntriesResponse struct {
	RrmMgrCfgEntries RrmMgrCfgEntries `json:"Cisco-IOS-XE-wireless-rrm-cfg:rrm-mgr-cfg-entries"`
}

// Rrms contains Radio Resource Management configuration entries
type Rrms struct {
	Rrm []Rrm `json:"rrm"`
}

// Rrm represents RRM configuration for a specific radio band
type Rrm struct {
	Band string `json:"band"` // Radio band (e.g., 2.4GHz, 5GHz)
	Rrm  *struct {
		MeasurementInterval int `json:"measurement-interval"` // Measurement interval in seconds
	} `json:"rrm,omitempty"`
}

// RrmMgrCfgEntries contains RRM manager configuration entries
type RrmMgrCfgEntries struct {
	RrmMgrCfgEntry []RrmMgrCfgEntry `json:"rrm-mgr-cfg-entry"`
}

// RrmMgrCfgEntry represents RRM manager configuration for a specific band
type RrmMgrCfgEntry struct {
	Band string `json:"band"` // Radio band identifier
}

// GetRrmCfg retrieves complete RRM configuration data.
// Returns RRM configuration including band settings and manager entries.
func GetRrmCfg(client *wnc.Client, ctx context.Context) (*RrmCfgResponse, error) {
	var data RrmCfgResponse
	err := client.SendAPIRequest(ctx, RrmCfgEndpoint, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// GetRrmRrms retrieves RRM configuration entries.
// Returns RRM configuration settings for different radio bands.
func GetRrmRrms(client *wnc.Client, ctx context.Context) (*RrmRrmsResponse, error) {
	var data RrmRrmsResponse
	err := client.SendAPIRequest(ctx, RrmCfgRrmsEndpoint, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// GetRrmMgrCfgEntries retrieves RRM manager configuration entries.
// Returns RRM manager configuration for different radio bands.
func GetRrmMgrCfgEntries(client *wnc.Client, ctx context.Context) (*RrmMgrCfgEntriesResponse, error) {
	var data RrmMgrCfgEntriesResponse
	err := client.SendAPIRequest(ctx, RrmCfgRrmMgrCfgEntriesEndpoint, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
