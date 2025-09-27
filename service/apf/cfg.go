package apf

// CiscoIOSXEWirelessAPFCfg represents APF configuration data container.
type CiscoIOSXEWirelessAPFCfg struct {
	CiscoIOSXEWirelessAPFCfgData struct {
		APF APF `json:"apf"` // APF configuration parameters (Live: IOS-XE 17.12.5)
	} `json:"Cisco-IOS-XE-wireless-apf-cfg:apf-cfg-data"` // APF configuration data (Live: IOS-XE 17.12.5)
}

// CiscoIOSXEWirelessAPFCfgAPF represents APF configuration wrapper.
type CiscoIOSXEWirelessAPFCfgAPF struct {
	APF APF `json:"Cisco-IOS-XE-wireless-apf-cfg:apf"`
}

// APF represents Access Point Filter configuration parameters.
type APF struct {
	NetworkName       string `json:"network-name"`        // RF network name identifier (Live: IOS-XE 17.12.5)
	ProbeLimit        int    `json:"probe-limit"`         // Maximum probe request limit (Live: IOS-XE 17.12.5)
	ProbeInterval     int    `json:"probe-interval"`      // Probe request interval in milliseconds (Live: IOS-XE 17.12.5)
	VlanPersistent    bool   `json:"vlan-persistent"`     // VLAN persistence across roaming (Live: IOS-XE 17.12.5)
	TagPersistEnabled bool   `json:"tag-persist-enabled"` // Tag persistence enable status (Live: IOS-XE 17.12.5)
}
