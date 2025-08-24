// Package model provides data models for CTS configuration data.
package model

// CtsCfg  represents the structure for CTS SXP configuration data.
type CtsCfg struct {
	CtsSxpCfgData CtsSxpCfgData `json:"Cisco-IOS-XE-wireless-cts-sxp-cfg:cts-sxp-cfg-data"`
}

// CtsCfgFilter  represents the structure for filtered CTS SXP configuration data.
type CtsCfgFilter struct {
	CtsSxpConfig []CtsSxpConfig `json:"Cisco-IOS-XE-wireless-cts-sxp-cfg:cts-sxp-config"`
}

type CtsSxpCfgData struct {
	CtsSxpConfiguration CtsSxpConfiguration `json:"cts-sxp-configuration"`
}

type CtsSxpConfiguration struct {
	CtsSxpConfig []CtsSxpConfig `json:"cts-sxp-config"`
}

type CtsSxpConfig struct {
	SxpProfileName          string          `json:"sxp-profile-name"`
	Enable                  *bool           `json:"enable,omitempty"`
	SxpDefaultPassword      *string         `json:"sxp-default-password,omitempty"`
	SxpDefaultSourceIP      *string         `json:"sxp-default-source-ip,omitempty"`
	SpeakerMinimumHoldtime  *int            `json:"speaker-minimum-holdtime,omitempty"`
	SpeakerMaximumHoldtime  *int            `json:"speaker-maximum-holdtime,omitempty"`
	ListenerMinimumHoldtime *int            `json:"listener-minimum-holdtime,omitempty"`
	ListenerMaximumHoldtime *int            `json:"listener-maximum-holdtime,omitempty"`
	SxpConnections          *SxpConnections `json:"sxp-connections,omitempty"`
}

type SxpConnections struct {
	SxpConnectionConfig []SxpConnectionConfig `json:"sxp-connection-config"`
}

type SxpConnectionConfig struct {
	PeerIPAddress  string `json:"peer-ip-address"`
	ConnectionMode string `json:"connection-mode"`
	PasswordType   string `json:"password-type"`
}
