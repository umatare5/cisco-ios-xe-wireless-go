package cts

// CiscoIOSXEWirelessCTSCfg represents the CTS SXP configuration data.
type CiscoIOSXEWirelessCTSCfg struct {
	CiscoIOSXEWirelessCTSCfgData struct {
		CTSSxpConfiguration CTSSxpConfiguration `json:"cts-sxp-configuration"` // CTS SXP configuration container (Live: IOS-XE 17.12.6a)
	} `json:"Cisco-IOS-XE-wireless-cts-sxp-cfg:cts-sxp-cfg-data"` // CTS SXP configuration data (Live: IOS-XE 17.12.6a)
}

// CTSSxpConfiguration represents the CTS SXP configuration container.
type CTSSxpConfiguration struct {
	CTSSxpConfig []CTSSxpConfig `json:"cts-sxp-config"` // CTS SXP configuration list (Live: IOS-XE 17.12.6a)
}

// CiscoIOSXEWirelessCTSCfgCTSSxpConfig represents CTS SXP configuration wrapper.
type CiscoIOSXEWirelessCTSCfgCTSSxpConfig struct {
	CTSSxpConfig []CTSSxpConfig `json:"Cisco-IOS-XE-wireless-cts-sxp-cfg:cts-sxp-config"`
}

// CTSSxpConfig represents individual CTS SXP profile configuration.
type CTSSxpConfig struct {
	SxpProfileName          string          `json:"sxp-profile-name"`                    // CTS SXP profile name (Live: IOS-XE 17.12.6a)
	Enable                  *bool           `json:"enable,omitempty"`                    // Enable CTS SXP configuration (YANG: IOS-XE 17.12.1)
	ListenerMinimumHoldtime *int            `json:"listener-minimum-holdtime,omitempty"` // CTS SXP listener hold time minimum (YANG: IOS-XE 17.12.1)
	ListenerMaximumHoldtime *int            `json:"listener-maximum-holdtime,omitempty"` // CTS SXP listener hold time maximum (YANG: IOS-XE 17.12.1)
	DefaultPassword         *string         `json:"default-password,omitempty"`          // CTS SXP default password (YANG: IOS-XE 17.12.1)
	SpeakerHoldtime         *int            `json:"speaker-holdtime,omitempty"`          // CTS SXP speaker hold time (YANG: IOS-XE 17.12.1)
	ReconcilePeriod         *int            `json:"reconcile-period,omitempty"`          // CTS SXP reconcile period (YANG: IOS-XE 17.12.1)
	RetryPeriod             *int            `json:"retry-period,omitempty"`              // CTS SXP retry period (YANG: IOS-XE 17.12.1)
	SxpConnections          *SxpConnections `json:"sxp-connections,omitempty"`           // SXP connection information container (YANG: IOS-XE 17.12.1)
}

// SxpConnections represents SXP connection configuration container.
type SxpConnections struct {
	SxpConnectionConfig []SxpConnectionConfig `json:"sxp-connection-config"` // SXP connections information list (YANG: IOS-XE 17.12.1)
}

// SxpConnectionConfig represents individual SXP connection configuration.
type SxpConnectionConfig struct {
	PeerIPAddress  string `json:"peer-ip-address"` // Peer IP address for SXP connection (YANG: IOS-XE 17.12.1)
	ConnectionMode string `json:"connection-mode"` // SXP connection mode (YANG: IOS-XE 17.12.1)
	PasswordType   string `json:"password-type"`   // SXP connection password type (YANG: IOS-XE 17.12.1)
}
