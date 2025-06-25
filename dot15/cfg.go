// Package dot15 provides 802.15 configuration functionality for the Cisco Wireless Network Controller API.
package dot15

import (
	"context"

	wnc "github.com/umatare5/cisco-xe-wireless-restconf-go"
)

const (
	// Dot15CfgBasePath defines the base path for 802.15 configuration endpoints
	Dot15CfgBasePath = "/restconf/data/Cisco-IOS-XE-wireless-dot15-cfg:dot15-cfg-data"
	// Dot15CfgEndpoint retrieves complete 802.15 configuration data
	Dot15CfgEndpoint = Dot15CfgBasePath
)

// Dot15CfgResponse represents the complete 802.15 configuration response
type Dot15CfgResponse struct {
	CiscoIOSXEWirelessDot15CfgDot15CfgData struct {
		Dot15GlobalConfig struct{} `json:"dot15-global-config"`
	} `json:"Cisco-IOS-XE-wireless-dot15-cfg:dot15-cfg-data"`
}

// Dot15GlobalConfigResponse represents the 802.15 global configuration response
type Dot15GlobalConfigResponse struct {
	Dot15GlobalConfig struct{} `json:"Cisco-IOS-XE-wireless-dot15-cfg:dot15-global-config"`
}

// GetDot15Cfg retrieves complete 802.15 configuration data.
func GetDot15Cfg(client *wnc.Client, ctx context.Context) (*Dot15CfgResponse, error) {
	var data Dot15CfgResponse
	if err := client.SendAPIRequest(ctx, Dot15CfgEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetDot15GlobalConfig retrieves 802.15 global configuration.
func GetDot15GlobalConfig(client *wnc.Client, ctx context.Context) (*Dot15GlobalConfigResponse, error) {
	var data Dot15GlobalConfigResponse
	if err := client.SendAPIRequest(ctx, Dot15CfgEndpoint+"/dot15-global-config", &data); err != nil {
		return nil, err
	}
	return &data, nil
}
