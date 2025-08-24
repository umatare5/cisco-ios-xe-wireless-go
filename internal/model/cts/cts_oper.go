// Package model provides data models for CTS operational data.
package model

// CtsOper  represents the structure for CTS SXP operational data.
type CtsOper struct {
	CtsSxpOperData CtsSxpOperData `json:"Cisco-IOS-XE-wireless-cts-sxp-oper:cts-sxp-oper-data"`
}

type CtsSxpOperData struct {
	FlexModeApSxpConnectionStatus []FlexModeApSxpConnectionStatus `json:"flex-mode-ap-sxp-connection-status"`
}

type FlexModeApSxpConnectionStatus struct {
	WtpMAC    string  `json:"wtp-mac"`
	PeerIP    string  `json:"peer-ip"`
	ConnMode  string  `json:"conn-mode"`
	ConnState *string `json:"conn-state,omitempty"`
	Duration  *string `json:"duration,omitempty"`
	LocalMode *string `json:"local-mode,omitempty"`
	PeerMode  *string `json:"peer-mode,omitempty"`
}
