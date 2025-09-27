package wlan

// WlanGlobalOper represents WLAN global operational data.
type WlanGlobalOper struct {
	CiscoIOSXEWirelessWlanGlobalOperData struct {
		WlanInfo []WlanInfo `json:"wlan-info"` // WLAN client statistics (Live: IOS-XE 17.12.6a)
	} `json:"Cisco-IOS-XE-wireless-wlan-global-oper:wlan-global-oper-data"` // Root container for WLAN operational data (Live: IOS-XE 17.12.6a)
}

// WlanInfo represents WLAN information.
type WlanInfo struct {
	WlanProfile            string `json:"wlan-profile"`               // WLAN profile name (Live: IOS-XE 17.12.6a)
	CurrClientsCount       int    `json:"curr-clients-count"`         // Number of active clients for this WLAN (Live: IOS-XE 17.12.6a)
	PerWlanMaxClientSyslog bool   `json:"per-wlan-max-client-syslog"` // Syslog message enabled when max clients reached (Live: IOS-XE 17.12.6a)
}

// CiscoIOSXEWirelessWlanGlobalOperWlanInfo wraps the WlanInfo structure of the WLAN global operational data.
type CiscoIOSXEWirelessWlanGlobalOperWlanInfo struct {
	CiscoIOSXEWirelessWlanGlobalOperData struct {
		WlanInfo []WlanInfo `json:"wlan-info,omitempty"`
	} `json:"Cisco-IOS-XE-wireless-wlan-global-oper:wlan-global-oper-data"`
}
