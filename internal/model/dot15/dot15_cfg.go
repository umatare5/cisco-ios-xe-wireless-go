// Package model provides data models for 802.15.4 configuration data.
package model

// Dot15Cfg  represents the structure for 802.15 configuration data.
type Dot15Cfg struct {
	Dot15CfgData Dot15CfgData `json:"Cisco-IOS-XE-wireless-dot15-cfg:dot15-cfg-data"`
}

// Dot15CfgDot15GlobalConfig  represents the structure for 802.15 global configuration.
type Dot15CfgDot15GlobalConfig struct {
	Dot15GlobalConfig *Dot15GlobalConfig `json:"Cisco-IOS-XE-wireless-dot15-cfg:dot15-global-config,omitempty"`
}

type Dot15CfgData struct {
	Dot15GlobalConfig *Dot15GlobalConfig `json:"dot15-global-config,omitempty"`
}

type Dot15GlobalConfig struct {
	GlobalRadioShut *bool `json:"global-radio-shut,omitempty"`
}
