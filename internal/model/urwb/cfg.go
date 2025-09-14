package urwb

// UrwbCfg represents the complete URWB configuration from YANG 17.18.1.
type UrwbCfg struct {
	CiscoIOSXEWirelessUrwbCfgData struct {
		UrwbProfiles *UrwbProfiles `json:"urwb-profiles,omitempty"`
	} `json:"Cisco-IOS-XE-wireless-urwb-cfg:urwb-cfg-data"`
}

// UrwbCfgUrwbProfiles represents the URWB profiles container from YANG 17.18.1.
type UrwbCfgUrwbProfiles struct {
	UrwbProfiles *UrwbProfiles `json:"Cisco-IOS-XE-wireless-urwb-cfg:urwb-profiles,omitempty"`
}

// UrwbProfiles represents the URWB policy configuration profiles for APs from YANG 17.18.1.
type UrwbProfiles struct {
	UrwbProfile []UrwbProfile `json:"urwb-profile"`
}

// UrwbProfile represents URWB policy configuration for AP from YANG 17.18.1.
type UrwbProfile struct {
	ProfileName    string          `json:"profile-name"`
	Description    string          `json:"descp,omitempty"`
	Enabled        bool            `json:"enabled,omitempty"`
	Passphrase     string          `json:"passphr,omitempty"`
	Mob            *UrwbMob        `json:"mob,omitempty"`
	Mpls           *UrwbMpls       `json:"mpls,omitempty"`
	Mpo            *UrwbMpo        `json:"mpo,omitempty"`
	Multicast      bool            `json:"mcast,omitempty"`
	StrongNetkey   bool            `json:"strong-netkey,omitempty"`
	EthertypeLists *EthertypeLists `json:"ethertype-lists,omitempty"`
}

// UrwbMob represents URWB mobility configuration from YANG 17.18.1.
type UrwbMob struct {
	Role       string `json:"role,omitempty"`         // urwb-mob-mode enum
	ScanIdle   uint32 `json:"scan-idle,omitempty"`    // seconds (0-65535)
	ScanAfter  uint32 `json:"scan-after,omitempty"`   // milliseconds (0-65535)
	ScanRssiTh uint32 `json:"scan-rssi-th,omitempty"` // dBm RSSI threshold for active scanning (0-96)
	Warmup     uint32 `json:"warmup,omitempty"`       // milliseconds warm up time (0-300000)
	Timeout    uint16 `json:"timeout,omitempty"`      // milliseconds timeout for candidate infrastructure
	RssiDhi    uint16 `json:"rssi-dhi,omitempty"`     // dBm mobility handoff hysteresis high threshold (0-96)
	RssiDlo    uint16 `json:"rssi-dlo,omitempty"`     // mobility handoff hysteresis low threshold (0-96)
	RssiDth    uint16 `json:"rssi-dth,omitempty"`     // dBm low or high threshold value (0-96)
	RaRed      uint8  `json:"ra-red,omitempty"`       // level of redundancy for route advertisement (1-5)
	BcEth      string `json:"bc-eth,omitempty"`       // urwb-mob-bce enum - handoff when ethernet disconnected
	BcCrd      string `json:"bc-crd,omitempty"`       // urwb-mob-bcc enum - handoff when coordinator unreachable
}

// UrwbMpls represents URWB MPLS configuration from YANG 17.18.1.
type UrwbMpls struct {
	HaEnabled     bool   `json:"ha-en,omitempty"`      // high availability enabled
	HaTimeout     uint16 `json:"ha-timeout,omitempty"` // MPLS high availability timeout
	UnicastFlood  bool   `json:"uni-fl,omitempty"`     // MPLS unicast flood enabled
	UniFloodLimit bool   `json:"uni-fl-l,omitempty"`   // MPLS unicast flood limits
	Eth1          bool   `json:"eth-1,omitempty"`      // ethernet I frames forwarding
	EthFm         string `json:"eth-fm,omitempty"`     // urwb-efm enum - ethernet filter method
}

// UrwbMpo represents URWB MPO (Multi-Path Optimization) configuration from YANG 17.18.1.
type UrwbMpo struct {
	Status   string `json:"status,omitempty"`    // urwb-mpo-st enum - MPO status
	MaxLinks uint8  `json:"max-links,omitempty"` // maximum number of MPO links (1-4)
	MinRssi  int8   `json:"min-rssi,omitempty"`  // minimum RSSI to establish redundant links (0-96)
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

// UrwbRadio represents URWB radio configuration from YANG 17.18.1.
type UrwbRadio struct {
	RadioType      string            `json:"radio-type,omitempty"`
	Channel        int               `json:"channel,omitempty"`
	Power          int               `json:"power,omitempty"`
	ChannelWidth   string            `json:"channel-width,omitempty"`
	BeaconInterval int               `json:"beacon-interval,omitempty"`
	ChannelLists   *UrwbChannelLists `json:"channel-lists,omitempty"`
}

// UrwbChannelLists represents URWB channel list configuration from YANG 17.18.1.
type UrwbChannelLists struct {
	ChannelListEntry []UrwbChannelListEntry `json:"channel-list-entry"`
}

// UrwbChannelListEntry represents a single channel list entry from YANG 17.18.1.
type UrwbChannelListEntry struct {
	ChannelNumber int    `json:"channel-number"`
	DfsRequired   bool   `json:"dfs-required,omitempty"`
	MaxPower      int    `json:"max-power,omitempty"`
	CountryCode   string `json:"country-code,omitempty"`
}
