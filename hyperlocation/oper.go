// Package hyperlocation provides hyperlocation operational data functionality for the Cisco Wireless Network Controller API.
package hyperlocation

import (
	"context"

	wnc "github.com/umatare5/cisco-xe-wireless-restconf-go"
)

const (
	// HyperlocationOperBasePath defines the base path for hyperlocation operational data endpoints.
	HyperlocationOperBasePath = "/restconf/data/Cisco-IOS-XE-wireless-hyperlocation-oper:hyperlocation-oper-data"
	// HyperlocationOperEndpoint defines the endpoint for hyperlocation operational data.
	HyperlocationOperEndpoint = HyperlocationOperBasePath
	// HyperlocationProfilesEndpoint defines the endpoint for hyperlocation profiles.
	HyperlocationProfilesEndpoint = HyperlocationOperBasePath + "/ewlc-hyperlocation-profile"
)

// HyperlocationOperResponse represents the response structure for hyperlocation operational data.
type HyperlocationOperResponse struct {
	CiscoIOSXEWirelessHyperlocationOperHyperlocationOperData struct {
		EwlcHyperlocationProfile []EwlcHyperlocationProfile `json:"ewlc-hyperlocation-profile"`
	} `json:"Cisco-IOS-XE-wireless-hyperlocation-oper:hyperlocation-oper-data"`
}

// HyperlocationProfilesResponse represents the response structure for hyperlocation profiles.
type HyperlocationProfilesResponse struct {
	EwlcHyperlocationProfile []EwlcHyperlocationProfile `json:"Cisco-IOS-XE-wireless-hyperlocation-oper:ewlc-hyperlocation-profile"`
}

// EwlcHyperlocationProfile represents an EWLC hyperlocation profile configuration.
type EwlcHyperlocationProfile struct {
	Name              string `json:"name"`
	HyperlocationData struct {
		HyperlocationEnable       bool `json:"hyperlocation-enable"`
		PakRssiThresholdDetection int  `json:"pak-rssi-threshold-detection"`
		PakRssiThresholdTrigger   int  `json:"pak-rssi-threshold-trigger"`
		PakRssiThresholdReset     int  `json:"pak-rssi-threshold-reset"`
	} `json:"hyperlocation-data"`
	NtpServer  string `json:"ntp-server"`
	Status     bool   `json:"status"`
	ReasonDown string `json:"reason-down"`
}

// GetHyperlocationOper retrieves hyperlocation operational data.
func GetHyperlocationOper(client *wnc.Client, ctx context.Context) (*HyperlocationOperResponse, error) {
	var data HyperlocationOperResponse
	if err := client.SendAPIRequest(ctx, HyperlocationOperEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetHyperlocationProfiles retrieves hyperlocation profiles.
func GetHyperlocationProfiles(client *wnc.Client, ctx context.Context) (*HyperlocationProfilesResponse, error) {
	var data HyperlocationProfilesResponse
	if err := client.SendAPIRequest(ctx, HyperlocationProfilesEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}
