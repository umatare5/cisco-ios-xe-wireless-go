// Package apf provides Access Point Filter configuration functionality for the Cisco Wireless Network Controller API.
package apf

import (
	"context"

	wnc "github.com/umatare5/cisco-xe-wireless-restconf-go"
)

const (
	// ApfCfgBasePath defines the base path for APF configuration endpoints
	ApfCfgBasePath = "/restconf/data/Cisco-IOS-XE-wireless-apf-cfg:apf-cfg-data"
	// ApfCfgEndpoint retrieves complete APF configuration data
	ApfCfgEndpoint = ApfCfgBasePath
	// ApfEndpoint retrieves APF specific configuration
	ApfEndpoint = ApfCfgBasePath + "/apf"
)

// ApfCfgResponse represents the complete APF configuration response
type ApfCfgResponse struct {
	CiscoIOSXEWirelessApfCfgApfCfgData struct {
		Apf Apf `json:"apf"`
	} `json:"Cisco-IOS-XE-wireless-apf-cfg:apf-cfg-data"`
}

// ApfCfgApfResponse represents the APF configuration response
type ApfCfgApfResponse struct {
	Apf Apf `json:"Cisco-IOS-XE-wireless-apf-cfg:apf"`
}

// Apf contains Access Point Framework configuration settings.
// The Access Point Framework provides management capabilities for wireless access points
// including system management options and network identification settings.
type Apf struct {
	SystemMgmtViaWireless bool   `json:"system-mgmt-via-wireless"` // Enable system management via wireless interface
	NetworkName           string `json:"network-name"`             // Network name identifier for APF
}

// GetApfCfg retrieves complete APF configuration data.
func GetApfCfg(client *wnc.Client, ctx context.Context) (*ApfCfgResponse, error) {
	var data ApfCfgResponse
	if err := client.SendAPIRequest(ctx, ApfCfgEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetApf retrieves APF specific configuration data.
func GetApf(client *wnc.Client, ctx context.Context) (*ApfCfgApfResponse, error) {
	var data ApfCfgApfResponse
	if err := client.SendAPIRequest(ctx, ApfEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}
