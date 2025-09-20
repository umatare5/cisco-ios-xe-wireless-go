package urwb

// URWBCfg represents the complete URWB configuration from YANG 17.18.1.
type URWBCfg struct {
	CiscoIOSXEWirelessURWBCfgData struct {
		URWBProfiles *URWBProfiles `json:"urwb-profiles,omitempty"`
	} `json:"Cisco-IOS-XE-wireless-urwb-cfg:urwb-cfg-data"`
}

// URWBCfgURWBProfiles represents the URWB profiles container from YANG 17.18.1.
type URWBCfgURWBProfiles struct {
	URWBProfiles *URWBProfiles `json:"Cisco-IOS-XE-wireless-urwb-cfg:urwb-profiles,omitempty"`
}

// URWBProfiles represents the URWB policy configuration profiles for APs from YANG 17.18.1.
type URWBProfiles struct {
	URWBProfile []URWBProfile `json:"urwb-profile"`
}

// URWBProfile represents URWB policy configuration for AP from YANG 17.18.1.
type URWBProfile struct {
	ProfileName    string          `json:"profile-name"`
	Description    string          `json:"descp,omitempty"`
	Enabled        bool            `json:"enabled,omitempty"`
	Passphrase     string          `json:"passphr,omitempty"`
	Mob            *URWBMob        `json:"mob,omitempty"`
	Mpls           *URWBMpls       `json:"mpls,omitempty"`
	Mpo            *URWBMpo        `json:"mpo,omitempty"`
	Multicast      bool            `json:"mcast,omitempty"`
	StrongNetkey   bool            `json:"strong-netkey,omitempty"`
	EthertypeLists *EthertypeLists `json:"ethertype-lists,omitempty"`
}

// URWBMob represents URWB mobility configuration from YANG 17.18.1.
type URWBMob struct {
	Role       string `json:"role,omitempty"`         // urwb-mob-mode enum
	ScanIdle   uint32 `json:"scan-idle,omitempty"`    // seconds (0-65535)
	ScanAfter  uint32 `json:"scan-after,omitempty"`   // milliseconds (0-65535)
	ScanRSSITh uint32 `json:"scan-rssi-th,omitempty"` // dBm RSSI threshold for active scanning (0-96)
	Warmup     uint32 `json:"warmup,omitempty"`       // milliseconds warm up time (0-300000)
	Timeout    uint16 `json:"timeout,omitempty"`      // milliseconds timeout for candidate infrastructure
	RSSIDhi    uint16 `json:"rssi-dhi,omitempty"`     // dBm mobility handoff hysteresis high threshold (0-96)
	RSSIDlo    uint16 `json:"rssi-dlo,omitempty"`     // mobility handoff hysteresis low threshold (0-96)
	RSSIDth    uint16 `json:"rssi-dth,omitempty"`     // dBm low or high threshold value (0-96)
	RaRed      uint8  `json:"ra-red,omitempty"`       // level of redundancy for route advertisement (1-5)
	BcEth      string `json:"bc-eth,omitempty"`       // urwb-mob-bce enum - handoff when ethernet disconnected
	BcCrd      string `json:"bc-crd,omitempty"`       // urwb-mob-bcc enum - handoff when coordinator unreachable
}

// URWBMpls represents URWB MPLS configuration from YANG 17.18.1.
type URWBMpls struct {
	HaEnabled     bool   `json:"ha-en,omitempty"`      // high availability enabled
	HaTimeout     uint16 `json:"ha-timeout,omitempty"` // MPLS high availability timeout
	UnicastFlood  bool   `json:"uni-fl,omitempty"`     // MPLS unicast flood enabled
	UniFloodLimit bool   `json:"uni-fl-l,omitempty"`   // MPLS unicast flood limits
	Eth1          bool   `json:"eth-1,omitempty"`      // ethernet I frames forwarding
	EthFm         string `json:"eth-fm,omitempty"`     // urwb-efm enum - ethernet filter method
}

// URWBMpo represents URWB MPO (Multi-Path Optimization) configuration from YANG 17.18.1.
type URWBMpo struct {
	Status   string `json:"status,omitempty"`    // urwb-mpo-st enum - MPO status
	MaxLinks uint8  `json:"max-links,omitempty"` // maximum number of MPO links (1-4)
	MinRSSI  int8   `json:"min-rssi,omitempty"`  // minimum RSSI to establish redundant links (0-96)
	ClassCs  uint8  `json:"class-cs,omitempty"`  // class-of-service for MPO redundancy (0-7)
	Tlmtry   bool   `json:"tlmtry,omitempty"`    // MPO telemetry enabled
}

// EthertypeLists represents URWB ether types list configuration from YANG 17.18.1.
type EthertypeLists struct {
	EthertypeList []EthertypeList `json:"ethertype-list"`
}

// EthertypeList represents a single ether type configuration from YANG 17.18.1.
type EthertypeList struct {
	Ethertype int `json:"ethertype"` // ether type value
}

// URWBRadio represents URWB radio configuration from YANG 17.18.1.
type URWBRadio struct {
	RadioType      string            `json:"radio-type,omitempty"`
	Channel        int               `json:"channel,omitempty"`
	Power          int               `json:"power,omitempty"`
	ChannelWidth   string            `json:"channel-width,omitempty"`
	BeaconInterval int               `json:"beacon-interval,omitempty"`
	ChannelLists   *URWBChannelLists `json:"channel-lists,omitempty"`
}

// URWBChannelLists represents URWB channel list configuration from YANG 17.18.1.
type URWBChannelLists struct {
	ChannelListEntry []URWBChannelListEntry `json:"channel-list-entry"`
}

// URWBChannelListEntry represents a single channel list entry from YANG 17.18.1.
type URWBChannelListEntry struct {
	ChannelNumber int    `json:"channel-number"`
	DFSRequired   bool   `json:"dfs-required,omitempty"`
	MaxPower      int    `json:"max-power,omitempty"`
	CountryCode   string `json:"country-code,omitempty"`
}
