package apf

// ApfCfg represents APF configuration data container.
type ApfCfg struct {
	ApfCfgData ApfCfgData `json:"Cisco-IOS-XE-wireless-apf-cfg:apf-cfg-data"`
}

// ApfCfgApf represents APF configuration wrapper.
type ApfCfgApf struct {
	Apf Apf `json:"Cisco-IOS-XE-wireless-apf-cfg:apf"`
}

// ApfCfgData represents APF configuration data structure.
type ApfCfgData struct {
	Apf Apf `json:"apf"`
}

// Apf represents Access Point Filter configuration parameters.
type Apf struct {
	NetworkName       string `json:"network-name"`        // Wireless network name identifier
	ProbeLimit        int    `json:"probe-limit"`         // Maximum probe request limit
	ProbeInterval     int    `json:"probe-interval"`      // Probe request interval in milliseconds
	VlanPersistent    bool   `json:"vlan-persistent"`     // VLAN persistence across roaming
	TagPersistEnabled bool   `json:"tag-persist-enabled"` // Tag persistence enable status
}
