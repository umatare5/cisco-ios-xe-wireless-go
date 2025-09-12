package dot15

// Dot15Cfg represents the root structure for IEEE 802.15.4 configuration data.
type Dot15Cfg struct {
	Dot15CfgData Dot15CfgData `json:"Cisco-IOS-XE-wireless-dot15-cfg:dot15-cfg-data"`
}

// Dot15CfgDot15GlobalConfig represents the structure for 802.15 global configuration operations.
type Dot15CfgDot15GlobalConfig struct {
	Dot15GlobalConfig *Dot15GlobalConfig `json:"Cisco-IOS-XE-wireless-dot15-cfg:dot15-global-config,omitempty"`
}

// Dot15CfgData represents the IEEE 802.15.4 configuration data container.
type Dot15CfgData struct {
	Dot15GlobalConfig *Dot15GlobalConfig `json:"dot15-global-config,omitempty"`
}

// Dot15GlobalConfig represents the global IEEE 802.15.4 radio configuration.
type Dot15GlobalConfig struct {
	GlobalRadioShut *bool `json:"global-radio-shut,omitempty"` // Global IEEE 802.15.4 radio switch (YANG: IOS-XE 17.12.1+)
}
