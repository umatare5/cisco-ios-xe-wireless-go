package model

// ApCfgApTagsParams  represents the tag configuration parameters for an Access Point.
// Tags are used to group APs and apply consistent configurations.
type ApCfgApTagsParams struct {
	SiteTag   string `json:"site-tag,omitempty"`
	PolicyTag string `json:"policy-tag,omitempty"`
	RFTag     string `json:"rf-tag,omitempty"`
}

type APTags = ApCfgApTagsParams

// ApCfg  represents the complete access point configuration
type ApCfg struct {
	CiscoIOSXEWirelessApCfgApCfgData struct {
		TagSourcePriorityConfigs TagSourcePriorityConfigs `json:"tag-source-priority-configs"`
		ApTags                   ApTags                   `json:"ap-tags"`
	} `json:"Cisco-IOS-XE-wireless-ap-cfg:ap-cfg-data"`
}

// ApCfgTagSourcePriorityConfigs  represents the tag source priority configurations.
type ApCfgTagSourcePriorityConfigs struct {
	TagSourcePriorityConfigs TagSourcePriorityConfigs `json:"Cisco-IOS-XE-wireless-ap-cfg:tag-source-priority-configs"`
}

// ApCfgApTags  represents the AP tags.
type ApCfgApTags struct {
	ApTags ApTags `json:"Cisco-IOS-XE-wireless-ap-cfg:ap-tags"`
}

// ApCfgApTag  represents a single AP tag.
type ApCfgApTag struct {
	ApTag []ApTag `json:"Cisco-IOS-XE-wireless-ap-cfg:ap-tag"`
}

type TagSourcePriorityConfigs struct {
	TagSourcePriorityConfig []struct {
		Priority int    `json:"priority"` // Priority level for tag source
		TagSrc   string `json:"tag-src"`  // Tag source identifier
	} `json:"tag-source-priority-config"`
}

type ApTags struct {
	ApTag []ApTag `json:"ap-tag"`
}

type ApTag struct {
	ApMac       string        `json:"ap-mac"`                 // Access point MAC address
	PolicyTag   string        `json:"policy-tag"`             // Policy tag assigned to the AP
	SiteTag     string        `json:"site-tag"`               // Site tag assigned to the AP
	RfTag       string        `json:"rf-tag,omitempty"`       // RF tag assigned to the AP (optional)
	RadioParams []RadioParams `json:"radio-params,omitempty"` // Radio parameters for the AP (optional)
}

type RadioParams struct {
	RadioSlotID int  `json:"radio-slot-id"` // Radio slot identifier (0=2.4GHz, 1=5GHz)
	AdminState  bool `json:"admin-state"`   // Administrative state (true=enabled, false=disabled)
}
