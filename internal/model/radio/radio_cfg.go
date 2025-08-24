// Package model provides data models for radio configuration data.
package model

// RadioCfg  represents the structure for Radio configuration data.
type RadioCfg struct {
	RadioCfgData RadioCfgData `json:"Cisco-IOS-XE-wireless-radio-cfg:radio-cfg-data"`
}

// RadioCfgRadioProfiles  represents the structure for radio profiles.
type RadioCfgRadioProfiles struct {
	RadioProfiles RadioProfiles `json:"Cisco-IOS-XE-wireless-radio-cfg:radio-profiles"`
}

type RadioCfgData struct {
	RadioProfiles RadioProfiles `json:"radio-profiles"`
}

type RadioProfiles struct {
	RadioProfile []RadioProfile `json:"radio-profile"`
}

type RadioProfile struct {
	Name         string `json:"name"`
	Desc         string `json:"desc,omitempty"`
	MeshBackhaul bool   `json:"mesh-backhaul"`
}
