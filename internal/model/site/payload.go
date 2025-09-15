package site

// SiteTagConfigsPayload represents request structure for site-tag-configs endpoint.
type SiteTagConfigsPayload struct {
	SiteListEntry SiteListEntry `json:"Cisco-IOS-XE-wireless-site-cfg:site-tag-config"`
}
