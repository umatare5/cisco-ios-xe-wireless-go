// Package model provides data models for RF configuration data.
package model

// RfCfg  represents the structure for RF configuration data
type RfCfg struct {
	RfCfgData RfCfgData `json:"Cisco-IOS-XE-wireless-rf-cfg:rf-cfg-data"`
}

// RfCfgData  represents the complete RF configuration data structure
type RfCfgData struct {
	MultiBssidProfiles      *MultiBssidProfiles      `json:"multi-bssid-profiles,omitempty"`
	AtfPolicies             *AtfPolicies             `json:"atf-policies,omitempty"`
	RfTags                  *RfTags                  `json:"rf-tags,omitempty"`
	RfProfiles              *RfProfiles              `json:"rf-profiles,omitempty"`
	RfProfileDefaultEntries *RfProfileDefaultEntries `json:"rf-profile-default-entries,omitempty"`
}

// MultiBssidProfiles  represents the Multi-BSSID profiles response
type MultiBssidProfiles struct {
	MultiBssidProfileList []MultiBssidProfile `json:"multi-bssid-profile"`
}

// MultiBssidProfile  represents a Multi-BSSID profile configuration
type MultiBssidProfile struct {
	ProfileName string `json:"profile-name"`
	Description string `json:"description,omitempty"`
}

// AtfPolicies  represents the ATF policies response
type AtfPolicies struct {
	AtfPolicyList []AtfPolicyDetail `json:"atf-policy"`
}

// AtfPolicy represents a single ATF policy structure.
type AtfPolicy struct {
	AtfPolicyList []AtfPolicyDetail `json:"atf-policy"`
}

// AtfPolicyDetail  represents an ATF policy configuration
type AtfPolicyDetail struct {
	PolicyID      int    `json:"policy-id"`
	AtfPolicyName string `json:"atfpolicy-name"`
	PolicyWeight  int    `json:"policy-weight,omitempty"`
	ClientSharing bool   `json:"client-sharing,omitempty"`
}

// RfTags  represents the RF tags response
type RfTags struct {
	RfTagList []RfTag `json:"rf-tag"`
}

// RfTag  represents an individual RF tag configuration
type RfTag struct {
	TagName             string              `json:"tag-name"`
	Description         string              `json:"description,omitempty"`
	Dot11ARfProfileName string              `json:"dot11a-rf-profile-name,omitempty"`
	Dot11BRfProfileName string              `json:"dot11b-rf-profile-name,omitempty"`
	Dot116GhzRfProfName string              `json:"dot11-6ghz-rf-prof-name,omitempty"`
	RfTagRadioProfiles  *RfTagRadioProfiles `json:"rf-tag-radio-profiles,omitempty"`
}

// RfTagRadioProfiles  represents RF tag radio profiles
type RfTagRadioProfiles struct {
	RfTagRadioProfile []RfTagRadioProfile `json:"rf-tag-radio-profile"`
}

// RfTagRadioProfile  represents an RF tag radio profile
type RfTagRadioProfile struct {
	SlotID string `json:"slot-id"`
	BandID string `json:"band-id"`
}

// RfProfiles  represents the RF profiles response
type RfProfiles struct {
	RfProfileList []RfProfileDetail `json:"Cisco-IOS-XE-wireless-rf-cfg:rf-profiles"`
}

// RfProfile represents a single RF profile structure.
type RfProfile struct {
	RfProfileList []RfProfileDetail `json:"Cisco-IOS-XE-wireless-rf-cfg:rf-profile"`
}

// RfProfileDetail  represents an RF profile configuration
type RfProfileDetail struct {
	Name                             string                `json:"name"`
	Description                      string                `json:"description,omitempty"`
	TxPowerMin                       int                   `json:"tx-power-min,omitempty"`
	TxPowerV1Threshold               int                   `json:"tx-power-v1-threshold,omitempty"`
	Status                           bool                  `json:"status,omitempty"`
	Band                             string                `json:"band,omitempty"`
	DataRate1M                       string                `json:"data-rate-1m,omitempty"`
	DataRate2M                       string                `json:"data-rate-2m,omitempty"`
	DataRate5_5M                     string                `json:"data-rate-5-5m,omitempty"`
	DataRate11M                      string                `json:"data-rate-11m,omitempty"`
	DataRate6M                       string                `json:"data-rate-6m,omitempty"`
	DataRate9M                       string                `json:"data-rate-9m,omitempty"`
	DataRate12M                      string                `json:"data-rate-12m,omitempty"`
	DataRate18M                      string                `json:"data-rate-18m,omitempty"`
	DataRate24M                      string                `json:"data-rate-24m,omitempty"`
	CoverageDataPacketRssiThreshold  int                   `json:"coverage-data-packet-rssi-threshold,omitempty"`
	CoverageVoicePacketRssiThreshold int                   `json:"coverage-voice-packet-rssi-threshold,omitempty"`
	LoadBalancingWindow              int                   `json:"load-balancing-window,omitempty"`
	LoadBalancingDenialCount         int                   `json:"load-balancing-denial-count,omitempty"`
	AtfOperMode                      string                `json:"atf-oper-mode,omitempty"`
	AtfOptimization                  string                `json:"atf-optimization,omitempty"`
	RfMcsEntries                     *RfMcsEntries         `json:"rf-mcs-entries,omitempty"`
	RfdcaRemovedChannels             *RfdcaRemovedChannels `json:"rfdca-removed-channels,omitempty"`
	ChannelWidthMax                  string                `json:"channel-width-max,omitempty"`
	MinNumClients                    int                   `json:"min-num-clients,omitempty"`
	RxSenSopThreshold                string                `json:"rx-sen-sop-threshold,omitempty"`
	BandSelectProbeResponse          bool                  `json:"band-select-probe-response,omitempty"`
}

// RfMcsEntries  represents RF MCS entries
type RfMcsEntries struct {
	RfMcsEntry []RfMcsEntry `json:"rf-mcs-entry"`
}

// RfMcsEntry  represents an RF MCS entry
type RfMcsEntry struct {
	RfIndex           int   `json:"rf-index"`
	Rf80211NMcsEnable *bool `json:"rf-80211n-mcs-enable,omitempty"`
}

// RfdcaRemovedChannels  represents DCA removed channels
type RfdcaRemovedChannels struct {
	RfdcaRemovedChannel []RfdcaRemovedChannel `json:"rfdca-removed-channel"`
}

// RfdcaRemovedChannel  represents a DCA removed channel
type RfdcaRemovedChannel struct {
	Channel int `json:"channel"`
}

// RfProfileDefaultEntries  represents the RF profile default entries response
type RfProfileDefaultEntries struct {
	RfProfileDefaultEntryList []RfProfileDefaultEntry `json:"Cisco-IOS-XE-wireless-rf-cfg:rf-profile-default-entries"`
}

// RfProfileDefaultEntry  represents an RF profile default entry
type RfProfileDefaultEntry struct {
	Band                     string               `json:"band"`
	Name                     string               `json:"name"`
	Description              string               `json:"description,omitempty"`
	DataRate1M               string               `json:"data-rate-1m,omitempty"`
	DataRate2M               string               `json:"data-rate-2m,omitempty"`
	DataRate5_5M             string               `json:"data-rate-5-5m,omitempty"`
	DataRate11M              string               `json:"data-rate-11m,omitempty"`
	DataRate6M               string               `json:"data-rate-6m,omitempty"`
	DataRate9M               string               `json:"data-rate-9m,omitempty"`
	DataRate12M              string               `json:"data-rate-12m,omitempty"`
	DataRate18M              string               `json:"data-rate-18m,omitempty"`
	DataRate24M              string               `json:"data-rate-24m,omitempty"`
	BandSelectProbeResponse  bool                 `json:"band-select-probe-response,omitempty"`
	LoadBalancingWindow      int                  `json:"load-balancing-window,omitempty"`
	LoadBalancingDenialCount int                  `json:"load-balancing-denial-count,omitempty"`
	AtfOperMode              string               `json:"atf-oper-mode,omitempty"`
	AtfOptimization          string               `json:"atf-optimization,omitempty"`
	RfMcsDefaultEntries      *RfMcsDefaultEntries `json:"rf-mcs-default-entries,omitempty"`
}

// RfMcsDefaultEntries  represents RF MCS default entries
type RfMcsDefaultEntries struct {
	RfMcsDefaultEntry []RfMcsDefaultEntry `json:"rf-mcs-default-entry"`
}

// RfMcsDefaultEntry  represents an RF MCS default entry
type RfMcsDefaultEntry struct {
	RfIndex     int `json:"rf-index"`
	McsDataRate int `json:"mcs-data-rate"`
}
