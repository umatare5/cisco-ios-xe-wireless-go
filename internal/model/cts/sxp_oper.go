package cts

// CTSOper represents the structure for CTS SXP operational data.
type CTSOper struct {
	CiscoIOSXEWirelessCTSSxpOperData struct {
		FlexModeApSxpConnectionStatus []FlexModeApSxpConnectionStatus `json:"flex-mode-ap-sxp-connection-status,omitempty"` // FlexConnect AP SXP connection status list (YANG: IOS-XE 17.12.1)
	} `json:"Cisco-IOS-XE-wireless-cts-sxp-oper:cts-sxp-oper-data"` // CTS SXP operational data (YANG: IOS-XE 17.12.1)
}

// FlexModeApSxpConnectionStatus represents SXP connection status for FlexConnect AP.
type FlexModeApSxpConnectionStatus struct {
	WtpMAC       string  `json:"wtp-mac"`                 // SXP connection status for a specific AP MAC (YANG: IOS-XE 17.12.1)
	PeerIP       string  `json:"peer-ip"`                 // SXP connection peer IP (YANG: IOS-XE 17.12.1)
	ConnMode     string  `json:"conn-mode"`               // SXP connection mode (YANG: IOS-XE 17.12.1)
	SrcIP        *string `json:"src-ip,omitempty"`        // SXP connection source IP (YANG: IOS-XE 17.12.1)
	NegoVersion  *int    `json:"nego-version,omitempty"`  // SXP version running on the device (YANG: IOS-XE 17.12.1)
	ConnStatus   *int    `json:"conn-status,omitempty"`   // SXP connection status (YANG: IOS-XE 17.12.1)
	PasswordType *string `json:"password-type,omitempty"` // SXP connection password type (YANG: IOS-XE 17.12.1)
}
