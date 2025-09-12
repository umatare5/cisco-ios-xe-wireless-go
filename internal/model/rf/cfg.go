// Package rf provides data models for RF configuration data.
package rf

// RfCfg represents RF configuration data response structure.
type RfCfg struct {
	RfCfgData RfCfgData `json:"Cisco-IOS-XE-wireless-rf-cfg:rf-cfg-data"`
}

// RfTagsPayload represents RF tag configuration request structure.
type RfTagsPayload struct {
	RfTag RfTag `json:"Cisco-IOS-XE-wireless-rf-cfg:rf-tag"`
}

// RfTagsResponse represents RF tags list response structure.
type RfTagsResponse struct {
	RfTags RfTags `json:"Cisco-IOS-XE-wireless-rf-cfg:rf-tags"`
}

// RfTagResponse represents individual RF tag response structure.
type RfTagResponse struct {
	RfTagList []RfTag `json:"Cisco-IOS-XE-wireless-rf-cfg:rf-tag"`
}

// RfCfgData represents complete RF configuration data container.
type RfCfgData struct {
	MultiBssidProfiles      *MultiBssidProfiles      `json:"multi-bssid-profiles,omitempty"`       // Multi-BSSID profiles configuration
	AtfPolicies             *AtfPolicies             `json:"atf-policies,omitempty"`               // Airtime Fairness policies configuration
	RfTags                  *RfTags                  `json:"rf-tags,omitempty"`                    // RF tag configurations
	RfProfiles              *RfProfiles              `json:"rf-profiles,omitempty"`                // RF profile configurations
	RfProfileDefaultEntries *RfProfileDefaultEntries `json:"rf-profile-default-entries,omitempty"` // RF profile default entries
}

// MultiBssidProfiles represents Multi-BSSID profiles collection.
type MultiBssidProfiles struct {
	MultiBssidProfileList []MultiBssidProfile `json:"multi-bssid-profile"`
}

// MultiBssidProfile represents Multi-BSSID profile configuration.
type MultiBssidProfile struct {
	ProfileName string `json:"profile-name"`          // Multi-BSSID profile name identifier
	Description string `json:"description,omitempty"` // Multi-BSSID profile description
}

// AtfPolicies represents ATF policies collection.
type AtfPolicies struct {
	AtfPolicyList []AtfPolicyDetail `json:"atf-policy"`
}

// AtfPolicy represents ATF policy structure container.
type AtfPolicy struct {
	AtfPolicyList []AtfPolicyDetail `json:"atf-policy"`
}

// AtfPolicyDetail represents ATF policy configuration details.
type AtfPolicyDetail struct {
	PolicyID      int    `json:"policy-id"`                // ATF policy identifier
	AtfPolicyName string `json:"atfpolicy-name"`           // ATF policy name
	PolicyWeight  int    `json:"policy-weight,omitempty"`  // ATF policy weight percentage
	ClientSharing bool   `json:"client-sharing,omitempty"` // ATF client sharing enable flag
}

// RfTags represents RF tags collection.
type RfTags struct {
	RfTagList []RfTag `json:"rf-tag"`
}

// RfTag represents RF tag configuration.
type RfTag struct {
	TagName             string              `json:"tag-name"`                          // RF tag name identifier
	Description         string              `json:"description,omitempty"`             // RF tag description
	Dot11ARfProfileName string              `json:"dot11a-rf-profile-name,omitempty"`  // 802.11a RF profile name
	Dot11BRfProfileName string              `json:"dot11b-rf-profile-name,omitempty"`  // 802.11b RF profile name
	Dot116GhzRfProfName string              `json:"dot11-6ghz-rf-prof-name,omitempty"` // 802.11 6GHz RF profile name
	RfTagRadioProfiles  *RfTagRadioProfiles `json:"rf-tag-radio-profiles,omitempty"`   // RF tag radio profiles configuration
}

// RfTagRadioProfiles represents RF tag radio profiles collection.
type RfTagRadioProfiles struct {
	RfTagRadioProfile []RfTagRadioProfile `json:"rf-tag-radio-profile"`
}

// RfTagRadioProfile represents RF tag radio profile configuration.
type RfTagRadioProfile struct {
	SlotID string `json:"slot-id"` // Radio slot identifier
	BandID string `json:"band-id"` // Radio band identifier
}

// RfProfiles represents RF profiles collection.
type RfProfiles struct {
	RfProfileList []RfProfileDetail `json:"rf-profile"`
}

// RfProfile represents RF profile structure container.
type RfProfile struct {
	RfProfileList []RfProfileDetail `json:"rf-profile"`
}

// RfProfileDetail represents RF profile configuration details.
type RfProfileDetail struct {
	Name                             string                `json:"name"`                                           // RF profile name identifier
	Description                      string                `json:"description,omitempty"`                          // RF profile description
	TxPowerMin                       int                   `json:"tx-power-min,omitempty"`                         // Minimum transmit power in dBm
	TxPowerMax                       int                   `json:"tx-power-max,omitempty"`                         // Maximum transmit power in dBm
	TxPowerV1Threshold               int                   `json:"tx-power-v1-threshold,omitempty"`                // Transmit power version 1 threshold
	TxPowerV2Threshold               int                   `json:"tx-power-v2-threshold,omitempty"`                // Transmit power version 2 threshold
	Status                           bool                  `json:"status,omitempty"`                               // RF profile enable status
	Band                             string                `json:"band,omitempty"`                                 // Radio frequency band
	DataRate1M                       string                `json:"data-rate-1m,omitempty"`                         // 1 Mbps data rate setting
	DataRate2M                       string                `json:"data-rate-2m,omitempty"`                         // 2 Mbps data rate setting
	DataRate5_5M                     string                `json:"data-rate-5-5m,omitempty"`                       // 5.5 Mbps data rate setting
	DataRate11M                      string                `json:"data-rate-11m,omitempty"`                        // 11 Mbps data rate setting
	DataRate6M                       string                `json:"data-rate-6m,omitempty"`                         // 6 Mbps data rate setting
	DataRate9M                       string                `json:"data-rate-9m,omitempty"`                         // 9 Mbps data rate setting
	DataRate12M                      string                `json:"data-rate-12m,omitempty"`                        // 12 Mbps data rate setting
	DataRate18M                      string                `json:"data-rate-18m,omitempty"`                        // 18 Mbps data rate setting
	DataRate24M                      string                `json:"data-rate-24m,omitempty"`                        // 24 Mbps data rate setting
	DataRate36M                      string                `json:"data-rate-36m,omitempty"`                        // 36 Mbps data rate setting
	DataRate48M                      string                `json:"data-rate-48m,omitempty"`                        // 48 Mbps data rate setting
	CoverageDataPacketRssiThreshold  int                   `json:"coverage-data-packet-rssi-threshold,omitempty"`  // Coverage data packet RSSI threshold in dBm
	CoverageVoicePacketRssiThreshold int                   `json:"coverage-voice-packet-rssi-threshold,omitempty"` // Coverage voice packet RSSI threshold in dBm
	LoadBalancingWindow              int                   `json:"load-balancing-window,omitempty"`                // Load balancing window in seconds
	LoadBalancingDenialCount         int                   `json:"load-balancing-denial-count,omitempty"`          // Load balancing denial count threshold
	AtfOperMode                      string                `json:"atf-oper-mode,omitempty"`                        // Airtime Fairness operation mode
	AtfOptimization                  string                `json:"atf-optimization,omitempty"`                     // Airtime Fairness optimization setting
	RfMcsEntries                     *RfMcsEntries         `json:"rf-mcs-entries,omitempty"`                       // RF MCS entries configuration
	RfdcaRemovedChannels             *RfdcaRemovedChannels `json:"rfdca-removed-channels,omitempty"`               // RF DCA removed channels configuration
	ChannelWidthMax                  string                `json:"channel-width-max,omitempty"`                    // Maximum channel width setting
	MinNumClients                    int                   `json:"min-num-clients,omitempty"`                      // Minimum number of clients threshold
	RxSenSopThreshold                string                `json:"rx-sen-sop-threshold,omitempty"`                 // Receive sensitivity SOP threshold
	BandSelectProbeResponse          bool                  `json:"band-select-probe-response,omitempty"`           // Band select probe response enable flag
}

// RfMcsEntries represents RF MCS entries collection.
type RfMcsEntries struct {
	RfMcsEntry []RfMcsEntry `json:"rf-mcs-entry"`
}

// RfMcsEntry represents RF MCS entry configuration.
type RfMcsEntry struct {
	RfIndex           int   `json:"rf-index"`                       // RF MCS index identifier
	Rf80211NMcsEnable *bool `json:"rf-80211n-mcs-enable,omitempty"` // 802.11n MCS enable flag
}

// RfdcaRemovedChannels represents DCA removed channels collection.
type RfdcaRemovedChannels struct {
	RfdcaRemovedChannel []RfdcaRemovedChannel `json:"rfdca-removed-channel"`
}

// RfdcaRemovedChannel represents DCA removed channel configuration.
type RfdcaRemovedChannel struct {
	Channel int `json:"channel"` // Channel number to remove from DCA
}

// RfProfileDefaultEntries represents RF profile default entries collection.
type RfProfileDefaultEntries struct {
	RfProfileDefaultEntryList []RfProfileDefaultEntry `json:"rf-profile-default-entry"`
}

// RfProfileDefaultEntry represents RF profile default entry configuration.
type RfProfileDefaultEntry struct {
	Band                     string               `json:"band"`                                  // Radio frequency band
	Name                     string               `json:"name"`                                  // RF profile default entry name
	Description              string               `json:"description,omitempty"`                 // RF profile default entry description
	DataRate1M               string               `json:"data-rate-1m,omitempty"`                // 1 Mbps data rate setting
	DataRate2M               string               `json:"data-rate-2m,omitempty"`                // 2 Mbps data rate setting
	DataRate5_5M             string               `json:"data-rate-5-5m,omitempty"`              // 5.5 Mbps data rate setting
	DataRate11M              string               `json:"data-rate-11m,omitempty"`               // 11 Mbps data rate setting
	DataRate6M               string               `json:"data-rate-6m,omitempty"`                // 6 Mbps data rate setting
	DataRate9M               string               `json:"data-rate-9m,omitempty"`                // 9 Mbps data rate setting
	DataRate12M              string               `json:"data-rate-12m,omitempty"`               // 12 Mbps data rate setting
	DataRate18M              string               `json:"data-rate-18m,omitempty"`               // 18 Mbps data rate setting
	DataRate24M              string               `json:"data-rate-24m,omitempty"`               // 24 Mbps data rate setting
	BandSelectProbeResponse  bool                 `json:"band-select-probe-response,omitempty"`  // Band select probe response enable flag
	LoadBalancingWindow      int                  `json:"load-balancing-window,omitempty"`       // Load balancing window in seconds
	LoadBalancingDenialCount int                  `json:"load-balancing-denial-count,omitempty"` // Load balancing denial count threshold
	AtfOperMode              string               `json:"atf-oper-mode,omitempty"`               // Airtime Fairness operation mode
	AtfOptimization          string               `json:"atf-optimization,omitempty"`            // Airtime Fairness optimization setting
	RfMcsDefaultEntries      *RfMcsDefaultEntries `json:"rf-mcs-default-entries,omitempty"`      // RF MCS default entries configuration
}

// RfMcsDefaultEntries represents RF MCS default entries collection.
type RfMcsDefaultEntries struct {
	RfMcsDefaultEntry []RfMcsDefaultEntry `json:"rf-mcs-default-entry"`
}

// RfMcsDefaultEntry represents RF MCS default entry configuration.
type RfMcsDefaultEntry struct {
	RfIndex     int `json:"rf-index"`      // RF MCS index identifier
	McsDataRate int `json:"mcs-data-rate"` // MCS data rate value
}
