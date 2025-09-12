package hyperlocation

// HyperlocationOper represents the hyperlocation operational data.
type HyperlocationOper struct {
	CiscoIOSXEWirelessHyperlocationOperHyperlocationOperData *HyperlocationOperData `json:"Cisco-IOS-XE-wireless-hyperlocation-oper:hyperlocation-oper-data,omitempty"`
}

// HyperlocationOperData represents the hyperlocation operational data container.
type HyperlocationOperData struct {
	EwlcHyperlocationProfile []EwlcHyperlocationProfile `json:"ewlc-hyperlocation-profile,omitempty"`
}

// HyperlocationOperEwlcHyperlocationProfile represents the EWLC hyperlocation profile endpoint response.
type HyperlocationOperEwlcHyperlocationProfile struct {
	EwlcHyperlocationProfile []EwlcHyperlocationProfile `json:"Cisco-IOS-XE-wireless-hyperlocation-oper:ewlc-hyperlocation-profile,omitempty"`
}

// HyperlocationProfiles represents the hyperlocation profiles collection.
type HyperlocationProfiles struct {
	EwlcHyperlocationProfile []EwlcHyperlocationProfile `json:"Cisco-IOS-XE-wireless-hyperlocation-oper:ewlc-hyperlocation-profile,omitempty"`
}

// EwlcHyperlocationProfile represents a single hyperlocation profile.
type EwlcHyperlocationProfile struct {
	Name              string             `json:"name"`                         // Hyperlocation profile name
	HyperlocationData *HyperlocationData `json:"hyperlocation-data,omitempty"` // Hyperlocation configuration container
	NtpServer         string             `json:"ntp-server"`                   // Configured NTP server address
	Status            bool               `json:"status"`                       // Operational status flag
	ReasonDown        string             `json:"reason-down"`                  // Reason for operational status being down
	OperNtpServer     *string            `json:"oper-ntp-server,omitempty"`    // Operational NTP server address (YANG: IOS-XE 17.12.1+)
}

// HyperlocationData represents the hyperlocation configuration data.
type HyperlocationData struct {
	HyperlocationEnable       bool `json:"hyperlocation-enable"`         // Hyperlocation enable status
	PakRssiThresholdDetection int  `json:"pak-rssi-threshold-detection"` // RSSI threshold for detection
	PakRssiThresholdTrigger   int  `json:"pak-rssi-threshold-trigger"`   // RSSI threshold trigger value
	PakRssiThresholdReset     int  `json:"pak-rssi-threshold-reset"`     // RSSI threshold reset value
}
