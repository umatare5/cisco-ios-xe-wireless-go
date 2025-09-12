package cts

// CtsCfg represents the CTS SXP configuration data.
type CtsCfg struct {
	CtsSxpCfgData CtsSxpCfgData `json:"Cisco-IOS-XE-wireless-cts-sxp-cfg:cts-sxp-cfg-data"`
}

// CtsCfgFilter represents the structure for filtered CTS SXP configuration data.
type CtsCfgFilter struct {
	CtsSxpConfig []CtsSxpConfig `json:"Cisco-IOS-XE-wireless-cts-sxp-cfg:cts-sxp-config"`
}

// CtsSxpCfgData represents CTS SXP configuration data container.
type CtsSxpCfgData struct {
	CtsSxpConfiguration CtsSxpConfiguration `json:"cts-sxp-configuration"`
}

// CtsSxpConfiguration represents CTS SXP configuration settings.
type CtsSxpConfiguration struct {
	CtsSxpConfig []CtsSxpConfig `json:"cts-sxp-config"`
}

// CtsSxpConfig represents individual CTS SXP profile configuration.
type CtsSxpConfig struct {
	SxpProfileName          string          `json:"sxp-profile-name"`                    // CTS SXP profile name (Live WNC)
	Enable                  *bool           `json:"enable,omitempty"`                    // Enable CTS SXP configuration (YANG: IOS-XE 17.12.1+)
	ListenerMinimumHoldtime *int            `json:"listener-minimum-holdtime,omitempty"` // Listener minimum hold time (YANG: IOS-XE 17.12.1+)
	ListenerMaximumHoldtime *int            `json:"listener-maximum-holdtime,omitempty"` // Listener maximum hold time (YANG: IOS-XE 17.12.1+)
	DefaultPassword         *string         `json:"default-password,omitempty"`          // CTS SXP default password (YANG: IOS-XE 17.12.1+)
	SpeakerHoldtime         *int            `json:"speaker-holdtime,omitempty"`          // Speaker hold time (YANG: IOS-XE 17.12.1+)
	ReconcilePeriod         *int            `json:"reconcile-period,omitempty"`          // SXP reconcile period (YANG: IOS-XE 17.12.1+)
	RetryPeriod             *int            `json:"retry-period,omitempty"`              // SXP retry period (YANG: IOS-XE 17.12.1+)
	SxpConnections          *SxpConnections `json:"sxp-connections,omitempty"`           // SXP connection configuration container (YANG: IOS-XE 17.12.1+)
}

// SxpConnections represents SXP connection configuration container.
type SxpConnections struct {
	SxpConnectionConfig []SxpConnectionConfig `json:"sxp-connection-config"`
}

// SxpConnectionConfig represents individual SXP connection configuration.
type SxpConnectionConfig struct {
	PeerIPAddress  string `json:"peer-ip-address"` // Peer IP address for SXP connection (YANG: IOS-XE 17.12.1+)
	ConnectionMode string `json:"connection-mode"` // SXP connection mode (YANG: IOS-XE 17.12.1+)
	PasswordType   string `json:"password-type"`   // SXP connection password type (YANG: IOS-XE 17.12.1+)
}
