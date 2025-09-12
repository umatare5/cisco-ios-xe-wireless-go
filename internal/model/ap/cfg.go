package ap

// ApCfg represents the complete access point configuration.
type ApCfg struct {
	CiscoIOSXEWirelessApCfgApCfgData struct {
		LocationEntries          *LocationEntries          `json:"location-entries,omitempty"`             // AP location configurations (YANG: IOS-XE 17.12.1+)
		TagSourcePriorityConfigs *TagSourcePriorityConfigs `json:"tag-source-priority-configs,omitempty"`  // Tag source priority configurations
		ApFilterConfigs          *ApFilterConfigs          `json:"ap-filter-configs,omitempty"`            // AP filter configurations (YANG: IOS-XE 17.12.1+)
		ApRulePriorityConfigs    *ApRulePriorityConfigs    `json:"ap-rule-priority-cfg-entries,omitempty"` // AP rule priority configurations (YANG: IOS-XE 17.12.1+)
		ApTags                   *ApTags                   `json:"ap-tags,omitempty"`                      // AP tag configurations
	} `json:"Cisco-IOS-XE-wireless-ap-cfg:ap-cfg-data"`
}

// ApCfgLocationEntries represents the AP location entries wrapper.
type ApCfgLocationEntries struct {
	LocationEntries LocationEntries `json:"Cisco-IOS-XE-wireless-ap-cfg:location-entries"`
}

// ApCfgTagSourcePriorityConfigs represents the tag source priority configurations.
type ApCfgTagSourcePriorityConfigs struct {
	TagSourcePriorityConfigs TagSourcePriorityConfigs `json:"Cisco-IOS-XE-wireless-ap-cfg:tag-source-priority-configs"`
}

// ApCfgApFilterConfigs represents the AP filter configurations wrapper.
type ApCfgApFilterConfigs struct {
	ApFilterConfigs ApFilterConfigs `json:"Cisco-IOS-XE-wireless-ap-cfg:ap-filter-configs"`
}

// ApCfgApRulePriorityConfigs represents the AP rule priority configurations wrapper.
type ApCfgApRulePriorityConfigs struct {
	ApRulePriorityConfigs ApRulePriorityConfigs `json:"Cisco-IOS-XE-wireless-ap-cfg:ap-rule-priority-cfg-entries"`
}

// ApCfgApTag represents a single AP tag.
type ApCfgApTag struct {
	ApTag []ApTag `json:"Cisco-IOS-XE-wireless-ap-cfg:ap-tag"`
}

// ApCfgApTags represents the AP tags.
type ApCfgApTags struct {
	ApTags ApTags `json:"Cisco-IOS-XE-wireless-ap-cfg:ap-tags"`
}

// ApFilterConfigs represents the AP filter configurations.
type ApFilterConfigs struct {
	ApFilterConfig []ApFilterConfig `json:"ap-filter-config"`
}

// ApRulePriorityConfigs represents the AP rule priority configurations.
type ApRulePriorityConfigs struct {
	ApRulePriorityConfig []ApRulePriorityConfig `json:"ap-rule-priority-cfg-entry"`
}

// LocationEntries represents the AP location configuration entries.
type LocationEntries struct {
	LocationEntry []LocationEntry `json:"location-entry"`
}

// ApTags represents the collection of AP tags.
type ApTags struct {
	ApTag []ApTag `json:"ap-tag"`
}

// LocationEntry represents a single AP location configuration.
type LocationEntry struct {
	LocationName       string              `json:"location-name"`                 // AP location name (YANG: IOS-XE 17.12.1+)
	Description        string              `json:"description,omitempty"`         // Location description (YANG: IOS-XE 17.12.1+)
	TagInfo            *LocationTagInfo    `json:"tag-info,omitempty"`            // Tag information for location
	AssociatedAps      *AssociatedAps      `json:"associated-aps,omitempty"`      // Associated APs (YANG: IOS-XE 17.12.1+)
	LocationAttributes *LocationAttributes `json:"location-attributes,omitempty"` // Location attributes (YANG: IOS-XE 17.12.1+)
}

// LocationTagInfo represents tag information for AP location.
type LocationTagInfo struct {
	PolicyTag string `json:"policy-tag,omitempty"` // Policy tag for location
	SiteTag   string `json:"site-tag,omitempty"`   // Site tag for location
	RfTag     string `json:"rf-tag,omitempty"`     // RF tag for location
}

// AssociatedAps represents the container of associated APs.
type AssociatedAps struct {
	AssociatedAp []AssociatedAp `json:"associated-ap"`
}

// AssociatedAp represents a single associated AP.
type AssociatedAp struct {
	ApMac string `json:"ap-mac"` // AP MAC address
}

// LocationAttributes represents location attributes for APs (YANG: IOS-XE 17.18.1+).
type LocationAttributes struct {
	// Location attributes are present in 17.18.1+ but not defined in 17.12.1
}

// ApFilterConfig represents a single AP filter configuration.
type ApFilterConfig struct {
	FilterName     string        `json:"filter-name"`               // Filter name
	FilterString   string        `json:"filter-string,omitempty"`   // Regular expression string (YANG: IOS-XE 17.12.1+)
	FilterPriority *uint8        `json:"filter-priority,omitempty"` // Filter priority (0-127) (YANG: IOS-XE 17.12.1+)
	ApplyTagList   *ApplyTagList `json:"apply-tag-list,omitempty"`  // Tag list to apply (YANG: IOS-XE 17.12.1+)
	FilterType     string        `json:"filter-type,omitempty"`     // AP filter type (YANG: IOS-XE 17.12.1+)
	PrimingProfile string        `json:"priming-profile,omitempty"` // Applied AP priming profile name
}

// ApplyTagList represents the tag list for AP filter.
type ApplyTagList struct {
	PolicyTag string `json:"policy-tag,omitempty"` // Policy tag
	SiteTag   string `json:"site-tag,omitempty"`   // Site tag
	RfTag     string `json:"rf-tag,omitempty"`     // RF tag
}

// ApRulePriorityConfig represents a single AP rule priority configuration.
type ApRulePriorityConfig struct {
	Priority   uint32 `json:"priority"`    // Priority of filter rule (0-1023)
	FilterName string `json:"filter-name"` // Name of the filter rule
}

// TagSourcePriorityConfigs represents tag source priority configurations container.
type TagSourcePriorityConfigs struct {
	TagSourcePriorityConfig []TagSourcePriorityConfig `json:"tag-source-priority-config"`
}

// TagSourcePriorityConfig represents a single tag source priority configuration.
type TagSourcePriorityConfig struct {
	Priority uint8  `json:"priority"` // Priority level for tag source (0-4)
	TagSrc   string `json:"tag-src"`  // Tag source identifier
}

// ApTag represents AP tag assignment configuration.
type ApTag struct {
	ApMac          string `json:"ap-mac"`                    // Access point MAC address
	PolicyTag      string `json:"policy-tag,omitempty"`      // Policy tag assigned to the AP
	SiteTag        string `json:"site-tag,omitempty"`        // Site tag assigned to the AP
	RFTag          string `json:"rf-tag,omitempty"`          // RF tag assigned to the AP
	PrimingProfile string `json:"priming-profile,omitempty"` // Configuration of AP priming profile (YANG: IOS-XE 17.12.1+)
}
