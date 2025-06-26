// Package location provides location configuration functionality for the Cisco Wireless Network Controller API.
package location

import (
	"context"

	wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
)

const (
	// LocationCfgBasePath defines the base path for location configuration endpoints
	LocationCfgBasePath = "/restconf/data/Cisco-IOS-XE-wireless-location-cfg:location-cfg-data"
	// LocationCfgEndpoint retrieves complete location configuration data
	LocationCfgEndpoint = LocationCfgBasePath
	// LocationCfgNmspConfigEndpoint retrieves NMSP configuration
	LocationCfgNmspConfigEndpoint = LocationCfgBasePath + "/nmsp-config"
)

// LocationCfgResponse represents the complete location configuration response
type LocationCfgResponse struct {
	CiscoIOSXEWirelessLocationCfgLocationCfgData struct {
		NmspConfig struct{} `json:"nmsp-config"`
	} `json:"Cisco-IOS-XE-wireless-location-cfg:location-cfg-data"`
}

// LocationNmspConfigResponse represents the NMSP configuration response
type LocationNmspConfigResponse struct {
	NmspConfig struct{} `json:"Cisco-IOS-XE-wireless-location-cfg:nmsp-config"`
}

// GetLocationCfg retrieves complete location configuration data.
func GetLocationCfg(client *wnc.Client, ctx context.Context) (*LocationCfgResponse, error) {
	var data LocationCfgResponse
	if err := client.SendAPIRequest(ctx, LocationCfgEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetLocationNmspConfig retrieves NMSP configuration data.
func GetLocationNmspConfig(client *wnc.Client, ctx context.Context) (*LocationNmspConfigResponse, error) {
	var data LocationNmspConfigResponse
	if err := client.SendAPIRequest(ctx, LocationCfgNmspConfigEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}
