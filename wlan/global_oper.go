// Package wlan provides WLAN global operational data functionality for the Cisco Wireless Network Controller API.
package wlan

import (
	"context"

	wnc "github.com/umatare5/cisco-xe-wireless-restconf-go"
)

const (
	// WlanGlobalOperBasePath defines the base path for WLAN global operational data endpoints.
	WlanGlobalOperBasePath = "/restconf/data/Cisco-IOS-XE-wireless-wlan-global-oper:wlan-global-oper-data"
	// WlanGlobalOperDataEndpoint defines the endpoint for WLAN global operational data.
	WlanGlobalOperDataEndpoint = WlanGlobalOperBasePath
	// WlanGlobalOperWlanInfoEndpoint defines the endpoint for WLAN information.
	WlanGlobalOperWlanInfoEndpoint = WlanGlobalOperBasePath + "/wlan-info"
)

// WlanGlobalOperResponse represents the response structure for WLAN global operational data.
type WlanGlobalOperResponse struct {
	CiscoIOSXEWirelessWlanGlobalOperData struct {
		WlanInfo []WlanInfo `json:"wlan-info"`
	} `json:"Cisco-IOS-XE-wireless-wlan-global-oper:wlan-global-oper-data"`
}

// WlanGlobalOperWlanInfoResponse represents the response structure for WLAN information.
type WlanGlobalOperWlanInfoResponse struct {
	WlanInfo []WlanInfo `json:"Cisco-IOS-XE-wireless-wlan-global-oper:wlan-info"`
}

// WlanInfo represents WLAN information including profile name and client counts.
type WlanInfo struct {
	WlanProfile            string `json:"wlan-profile"`
	CurrClientsCount       int    `json:"curr-clients-count"`
	PerWlanMaxClientSyslog bool   `json:"per-wlan-max-client-syslog"`
}

// GetWlanGlobalOper retrieves WLAN global operational data.
func GetWlanGlobalOper(client *wnc.Client, ctx context.Context) (*WlanGlobalOperResponse, error) {
	var data WlanGlobalOperResponse
	if err := client.SendAPIRequest(ctx, WlanGlobalOperDataEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetWlanGlobalOperWlanInfo retrieves WLAN information.
func GetWlanGlobalOperWlanInfo(client *wnc.Client, ctx context.Context) (*WlanGlobalOperWlanInfoResponse, error) {
	var data WlanGlobalOperWlanInfoResponse
	if err := client.SendAPIRequest(ctx, WlanGlobalOperWlanInfoEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}
