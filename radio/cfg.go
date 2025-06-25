// Package radio provides radio configuration functionality for the Cisco Wireless Network Controller API.
package radio

import (
	"context"

	wnc "github.com/umatare5/cisco-xe-wireless-restconf-go"
)

const (
	// RadioCfgBasePath defines the base path for radio configuration endpoints
	RadioCfgBasePath = "/restconf/data/Cisco-IOS-XE-wireless-radio-cfg:radio-cfg-data"
	// RadioCfgEndpoint retrieves complete radio configuration data
	RadioCfgEndpoint = RadioCfgBasePath
	// RadioProfilesEndpoint retrieves radio profiles configuration
	RadioProfilesEndpoint = RadioCfgBasePath + "/radio-profiles"
)

// RadioCfgResponse represents the complete radio configuration response
type RadioCfgResponse struct {
	CiscoIOSXEWirelessRadioCfgData struct {
		RadioProfiles RadioProfiles `json:"radio-profiles"`
	} `json:"Cisco-IOS-XE-wireless-radio-cfg:radio-cfg-data"`
}

// RadioProfilesResponse represents the radio profiles configuration response
type RadioProfilesResponse struct {
	RadioProfiles RadioProfiles `json:"Cisco-IOS-XE-wireless-radio-cfg:radio-profiles"`
}

// RadioProfiles contains radio profile configuration entries
type RadioProfiles struct {
	RadioProfile []RadioProfile `json:"radio-profile"`
}

// RadioProfile represents a radio profile configuration
type RadioProfile struct {
	Name         string `json:"name"`           // Radio profile name
	Desc         string `json:"desc,omitempty"` // Optional profile description
	MeshBackhaul bool   `json:"mesh-backhaul"`  // Enable mesh backhaul
}

// GetRadioCfg retrieves complete radio configuration data.
func GetRadioCfg(client *wnc.Client, ctx context.Context) (*RadioCfgResponse, error) {
	var data RadioCfgResponse
	if err := client.SendAPIRequest(ctx, RadioCfgEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetRadioProfiles retrieves radio profile configurations.
func GetRadioProfiles(client *wnc.Client, ctx context.Context) (*RadioProfilesResponse, error) {
	var data RadioProfilesResponse
	if err := client.SendAPIRequest(ctx, RadioProfilesEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}
