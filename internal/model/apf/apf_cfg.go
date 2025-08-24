// Package model provides data models for AP Filter configuration data.
package model

// ApfCfg  represents the APF configuration data.
type ApfCfg struct {
	ApfCfgData ApfCfgData `json:"Cisco-IOS-XE-wireless-apf-cfg:apf-cfg-data"`
}

// ApfCfgApf  represents the APF configuration.
type ApfCfgApf struct {
	Apf Apf `json:"Cisco-IOS-XE-wireless-apf-cfg:apf"`
}

type ApfCfgData struct {
	Apf Apf `json:"apf"`
}

type Apf struct {
	NetworkName       string `json:"network-name"`
	ProbeLimit        int    `json:"probe-limit"`
	ProbeInterval     int    `json:"probe-interval"`
	VlanPersistent    bool   `json:"vlan-persistent"`
	TagPersistEnabled bool   `json:"tag-persist-enabled"`
}
