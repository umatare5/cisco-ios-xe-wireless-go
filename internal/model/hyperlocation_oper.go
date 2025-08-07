// Package model contains generated response structures for the Cisco WNC API.
// This package is part of the three-layer architecture providing Generated Type separation.
package model

// Hyperlocation Operational Response Types

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

// Hyperlocation Supporting Types

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
