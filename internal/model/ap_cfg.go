// Package model contains generated response structures for the Cisco WNC API.
package model

// ApCfgResponse represents the complete access point configuration response
type ApCfgResponse struct {
	CiscoIOSXEWirelessApCfgApCfgData struct {
		TagSourcePriorityConfigs TagSourcePriorityConfigs `json:"tag-source-priority-configs"`
		ApTags                   ApTags                   `json:"ap-tags"`
	} `json:"Cisco-IOS-XE-wireless-ap-cfg:ap-cfg-data"`
}

// ApCfgTagSourcePriorityConfigsResponse represents tag source priority configurations response
type ApCfgTagSourcePriorityConfigsResponse struct {
	TagSourcePriorityConfigs TagSourcePriorityConfigs `json:"Cisco-IOS-XE-wireless-ap-cfg:tag-source-priority-configs"`
}

// ApCfgApTagsResponse represents access point tags configuration response
type ApCfgApTagsResponse struct {
	ApTags ApTags `json:"Cisco-IOS-XE-wireless-ap-cfg:ap-tags"`
}

// TagSourcePriorityConfigs contains tag source priority configuration settings
type TagSourcePriorityConfigs struct {
	TagSourcePriorityConfig []struct {
		Priority int    `json:"priority"` // Priority level for tag source
		TagSrc   string `json:"tag-src"`  // Tag source identifier
	} `json:"tag-source-priority-config"`
}

// ApTags contains access point tag configuration data
type ApTags struct {
	ApTag []ApTag `json:"ap-tag"`
}

// ApTag represents tag assignments for a specific access point
type ApTag struct {
	ApMac     string `json:"ap-mac"`           // Access point MAC address
	PolicyTag string `json:"policy-tag"`       // Policy tag assigned to the AP
	SiteTag   string `json:"site-tag"`         // Site tag assigned to the AP
	RfTag     string `json:"rf-tag,omitempty"` // RF tag assigned to the AP (optional)
}
