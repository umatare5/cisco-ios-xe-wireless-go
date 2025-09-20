package dot15

// Dot15Cfg represents the root structure for IEEE 802.15.4 configuration data.
type Dot15Cfg struct {
	CiscoIOSXEWirelessDot15CfgData struct {
		Dot15GlobalConfig *Dot15GlobalConfig `json:"dot15-global-config"` // 802.15 global configuration (Live: IOS-XE 17.12.5)
	} `json:"Cisco-IOS-XE-wireless-dot15-cfg:dot15-cfg-data"` // IEEE 802.15 configuration data (Live: IOS-XE 17.12.5)
}

// Dot15CfgDot15GlobalConfig represents the structure for 802.15 global configuration operations.
type Dot15CfgDot15GlobalConfig struct {
	Dot15GlobalConfig *Dot15GlobalConfig `json:"Cisco-IOS-XE-wireless-dot15-cfg:dot15-global-config,omitempty"`
}

// Dot15GlobalConfig represents the global IEEE 802.15.4 radio configuration.
type Dot15GlobalConfig struct {
	GlobalRadioShut *bool `json:"global-radio-shut,omitempty"` // Global 802.15 radio switch (YANG: IOS-XE 17.12.1)
}
