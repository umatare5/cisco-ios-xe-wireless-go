package model

// ApfCfgResponse represents the response structure for APF configuration data.
type ApfCfgResponse struct {
	ApfCfgData ApfCfgData `json:"Cisco-IOS-XE-wireless-apf-cfg:apf-cfg-data"`
}

// ApfCfgData contains APF configuration data
type ApfCfgData struct {
	Apf Apf `json:"apf"`
}

// Apf represents APF configuration settings
type Apf struct {
	NetworkName       string `json:"network-name"`
	ProbeLimit        int    `json:"probe-limit"`
	ProbeInterval     int    `json:"probe-interval"`
	VlanPersistent    bool   `json:"vlan-persistent"`
	TagPersistEnabled bool   `json:"tag-persist-enabled"`
}
