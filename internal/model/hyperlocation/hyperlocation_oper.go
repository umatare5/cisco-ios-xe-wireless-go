// Package model provides type definitions for Cisco IOS-XE wireless controller operations.
package model

// Hyperlocation Operational Response Types

// HyperlocationOper  represents the hyperlocation operational data.
type HyperlocationOper struct {
	CiscoIOSXEWirelessHyperlocationOperHyperlocationOperData struct {
		EwlcHyperlocationProfile []EwlcHyperlocationProfile `json:"ewlc-hyperlocation-profile"`
	} `json:"Cisco-IOS-XE-wireless-hyperlocation-oper:hyperlocation-oper-data"`
}

// HyperlocationOperEwlcHyperlocationProfile  represents the EWLC hyperlocation profile.
type HyperlocationOperEwlcHyperlocationProfile struct {
	EwlcHyperlocationProfile []EwlcHyperlocationProfile `json:"Cisco-IOS-XE-wireless-hyperlocation-oper:ewlc-hyperlocation-profile"`
}

// HyperlocationProfiles  represents the hyperlocation profiles.
type HyperlocationProfiles struct {
	EwlcHyperlocationProfile []EwlcHyperlocationProfile `json:"Cisco-IOS-XE-wireless-hyperlocation-oper:ewlc-hyperlocation-profile"`
}

// Hyperlocation Supporting Types

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
