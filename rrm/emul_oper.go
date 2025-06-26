// Package rrm provides Radio Resource Management emulation operational data functionality for the Cisco Wireless Network Controller API.
package rrm

import (
	"context"

	wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
)

const (
	// RrmEmulOperBasePath defines the base path for RRM emulation operational data endpoints.
	RrmEmulOperBasePath = "/restconf/data/Cisco-IOS-XE-wireless-rrm-emul-oper:rrm-emul-oper-data"
	// RrmEmulOperEndpoint defines the endpoint for RRM emulation operational data.
	RrmEmulOperEndpoint = RrmEmulOperBasePath
	// RrmEmulOperRrmFraStatsEndpoint defines the endpoint for RRM emulation FRA statistics.
	RrmEmulOperRrmFraStatsEndpoint = RrmEmulOperBasePath + "/rrm-fra-stats"
)

// RrmEmulOperResponse represents the response structure for RRM emulation operational data.
type RrmEmulOperResponse struct {
	CiscoIOSXEWirelessRrmEmulOperData struct {
		RrmFraStats RrmEmulOperRrmFraStats `json:"rrm-fra-stats"`
	} `json:"Cisco-IOS-XE-wireless-rrm-emul-oper:rrm-emul-oper-data"`
}

// RrmEmulOperRrmFraStatsResponse represents the response structure for RRM emulation FRA statistics.
type RrmEmulOperRrmFraStatsResponse struct {
	RrmFraStats RrmEmulOperRrmFraStats `json:"Cisco-IOS-XE-wireless-rrm-emul-oper:rrm-fra-stats"`
}

// RrmEmulOperRrmFraStats represents RRM emulation FRA statistics including band switching counts.
type RrmEmulOperRrmFraStats struct {
	DualBandMonitorTo24ghz int `json:"dual-band-monitor-to-24ghz"`
	DualBandMonitorTo5ghz  int `json:"dual-band-monitor-to-5ghz"`
	DualBand24ghzTo5ghz    int `json:"dual-band-24ghz-to-5ghz"`
	DualBand24ghzToMonitor int `json:"dual-band-24ghz-to-monitor"`
	DualBand5ghzTo24ghz    int `json:"dual-band-5ghz-to-24ghz"`
	DualBand5ghzToMonitor  int `json:"dual-band-5ghz-to-monitor"`
	SecRadioMonitorTo5ghz  int `json:"sec-radio-monitor-to-5ghz"`
	SecRadio5ghzToMonitor  int `json:"sec-radio-5ghz-to-monitor"`
	DualBand6ghzTo5ghz     int `json:"dual-band-6ghz-to-5ghz"`
	DualBand5ghzTo6ghz     int `json:"dual-band-5ghz-to-6ghz"`
}

// GetRrmEmulOper retrieves RRM emulation operational data.
func GetRrmEmulOper(client *wnc.Client, ctx context.Context) (*RrmEmulOperResponse, error) {
	var data RrmEmulOperResponse
	if err := client.SendAPIRequest(ctx, RrmEmulOperEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetRrmEmulRrmFraStats retrieves RRM emulation FRA statistics.
func GetRrmEmulRrmFraStats(client *wnc.Client, ctx context.Context) (*RrmEmulOperRrmFraStatsResponse, error) {
	var data RrmEmulOperRrmFraStatsResponse
	if err := client.SendAPIRequest(ctx, RrmEmulOperRrmFraStatsEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}
