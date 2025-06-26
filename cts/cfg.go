// Package cts provides Cisco TrustSec configuration functionality for the Cisco Wireless Network Controller API.
package cts

import (
	"context"

	wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
)

const (
	// CtsSxpCfgBasePath defines the base path for CTS SXP configuration endpoints
	CtsSxpCfgBasePath = "/restconf/data/Cisco-IOS-XE-wireless-cts-sxp-cfg:cts-sxp-cfg-data"
	// CtsSxpCfgEndpoint retrieves complete CTS SXP configuration data
	CtsSxpCfgEndpoint = CtsSxpCfgBasePath
	// CtsSxpConfigurationEndpoint retrieves CTS SXP configuration
	CtsSxpConfigurationEndpoint = CtsSxpCfgBasePath + "/cts-sxp-configuration"
)

// CtsSxpCfgResponse represents the complete CTS SXP configuration response
type CtsSxpCfgResponse struct {
	CiscoIOSXEWirelessCtsSxpCfgCtsSxpCfgData struct {
		CtsSxpConfiguration CtsSxpConfiguration `json:"cts-sxp-configuration"`
	} `json:"Cisco-IOS-XE-wireless-cts-sxp-cfg:cts-sxp-cfg-data"`
}

// CtsSxpConfigurationResponse represents the CTS SXP configuration response
type CtsSxpConfigurationResponse struct {
	CtsSxpConfiguration CtsSxpConfiguration `json:"Cisco-IOS-XE-wireless-cts-sxp-cfg:cts-sxp-configuration"`
}

// CtsSxpConfiguration contains CTS SXP configuration entries
type CtsSxpConfiguration struct {
	CtsSxpConfig []CtsSxpConfig `json:"cts-sxp-config"`
}

// CtsSxpConfig represents a CTS SXP configuration entry
type CtsSxpConfig struct {
	SxpProfileName string `json:"sxp-profile-name"` // SXP profile name
}

// GetCtsSxpCfg retrieves complete CTS SXP configuration data.
func GetCtsSxpCfg(client *wnc.Client, ctx context.Context) (*CtsSxpCfgResponse, error) {
	var data CtsSxpCfgResponse
	if err := client.SendAPIRequest(ctx, CtsSxpCfgEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetCtsSxpConfiguration retrieves CTS SXP configuration entries.
func GetCtsSxpConfiguration(client *wnc.Client, ctx context.Context) (*CtsSxpConfigurationResponse, error) {
	var data CtsSxpConfigurationResponse
	if err := client.SendAPIRequest(ctx, CtsSxpConfigurationEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}
