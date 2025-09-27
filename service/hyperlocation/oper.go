package hyperlocation

// CiscoIOSXEWirelessHyperlocationOper represents the hyperlocation operational data.
type CiscoIOSXEWirelessHyperlocationOper struct {
	CiscoIOSXEWirelessHyperlocationOperData struct {
		EwlcHyperlocationProfile []EwlcHyperlocationProfile `json:"ewlc-hyperlocation-profile"` // Hyperlocation AP profile data (Live: IOS-XE 17.12.5)
	} `json:"Cisco-IOS-XE-wireless-hyperlocation-oper:hyperlocation-oper-data"` // Hyperlocation operational data (Live: IOS-XE 17.12.5)
}

// CiscoIOSXEWirelessHyperlocationProfiles represents the hyperlocation profiles collection.
type CiscoIOSXEWirelessHyperlocationProfiles struct {
	EwlcHyperlocationProfile []EwlcHyperlocationProfile `json:"Cisco-IOS-XE-wireless-hyperlocation-oper:ewlc-hyperlocation-profile"`
}

// EwlcHyperlocationProfile represents a single hyperlocation profile.
type EwlcHyperlocationProfile struct {
	Name              string             `json:"name"`                         // AP profile name (Live: IOS-XE 17.12.5)
	HyperlocationData *HyperlocationData `json:"hyperlocation-data,omitempty"` // Cisco AP hyperlocation details (Live: IOS-XE 17.12.5)
	NtpServer         string             `json:"ntp-server"`                   // Configured hyperlocation NTP server (Live: IOS-XE 17.12.5)
	Status            bool               `json:"status"`                       // Operational status (Live: IOS-XE 17.12.5)
	ReasonDown        string             `json:"reason-down"`                  // Reason for operational status being down (Live: IOS-XE 17.12.5)
	OperNtpServer     *string            `json:"oper-ntp-server,omitempty"`    // Operational NTP server (YANG: IOS-XE 17.18.1)
}

// HyperlocationData represents the hyperlocation configuration data.
type HyperlocationData struct {
	HyperlocationEnable       bool `json:"hyperlocation-enable"`         // Enable hyperlocation (Live: IOS-XE 17.12.5)
	PakRSSIThresholdDetection int  `json:"pak-rssi-threshold-detection"` // Pak rssi threshold detection (Live: IOS-XE 17.12.5)
	PakRSSIThresholdTrigger   int  `json:"pak-rssi-threshold-trigger"`   // Pak rssi threshold trigger (Live: IOS-XE 17.12.5)
	PakRSSIThresholdReset     int  `json:"pak-rssi-threshold-reset"`     // Pak rssi threshold reset (Live: IOS-XE 17.12.5)
}
