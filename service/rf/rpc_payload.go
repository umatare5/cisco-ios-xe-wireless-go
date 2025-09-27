package rf

// CiscoIOSXEWirelessRFCfgRFTagPayload represents individual RF tag response structure.
type CiscoIOSXEWirelessRFCfgRFTagPayload struct {
	RFTags []RFTag `json:"Cisco-IOS-XE-wireless-rf-cfg:rf-tag"`
}

// CiscoIOSXEWirelessRFCfgRFTagsPayload represents request structure for rf-tags endpoint.
type CiscoIOSXEWirelessRFCfgRFTagsPayload struct {
	RFTag RFTag `json:"Cisco-IOS-XE-wireless-rf-cfg:rf-tag"`
}
