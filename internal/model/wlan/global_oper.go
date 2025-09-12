package wlan

// WlanGlobalOper represents WLAN global operational data from live WNC 17.12.1.
type WlanGlobalOper struct {
	WlanGlobalOperData struct {
		WlanInfo []WlanInfo `json:"wlan-info"`
	} `json:"Cisco-IOS-XE-wireless-wlan-global-oper:wlan-global-oper-data"`
}

// WlanInfo represents WLAN information from live WNC 17.12.1.
type WlanInfo struct {
	WlanProfile            string `json:"wlan-profile"`               // WLAN profile name
	CurrClientsCount       int    `json:"curr-clients-count"`         // Current connected clients count
	PerWlanMaxClientSyslog bool   `json:"per-wlan-max-client-syslog"` // Per-WLAN max client syslog enabled
}
