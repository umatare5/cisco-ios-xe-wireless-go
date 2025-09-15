package rf

// RFCfg represents RF configuration data response structure.
type RFCfg struct {
	RFCfgData struct {
		MultiBssidProfiles      *MultiBssidProfiles      `json:"multi-bssid-profiles,omitempty"`       // 802.11ax Multi BSSID profile configuration (Live: IOS-XE 17.12.5)
		AtfPolicies             *AtfPolicies             `json:"atf-policies,omitempty"`               // Air Time Fairness policy configurations (Live: IOS-XE 17.12.5)
		RFTags                  *RFTags                  `json:"rf-tags,omitempty"`                    // RF tag configuration data (Live: IOS-XE 17.12.5)
		RFProfiles              *RFProfiles              `json:"rf-profiles,omitempty"`                // RF profile configuration data (Live: IOS-XE 17.12.5)
		RFProfileDefaultEntries *RFProfileDefaultEntries `json:"rf-profile-default-entries,omitempty"` // Default RF profile entries (YANG: IOS-XE 17.12.1)
	} `json:"Cisco-IOS-XE-wireless-rf-cfg:rf-cfg-data"` // RF configuration data (Live: IOS-XE 17.12.5)
}

// RFCfgRFTags represents RF tags list response structure.
type RFCfgRFTags struct {
	RFTags RFTags `json:"Cisco-IOS-XE-wireless-rf-cfg:rf-tags"`
}

// RFCfgRFTag represents individual RF tag response structure.
type RFCfgRFTag struct {
	RFTagList []RFTag `json:"Cisco-IOS-XE-wireless-rf-cfg:rf-tag"`
}

// MultiBssidProfiles represents Multi-BSSID profiles collection.
type MultiBssidProfiles struct {
	MultiBssidProfileList []MultiBssidProfile `json:"multi-bssid-profile"` // Multi BSSID profile list (Live: IOS-XE 17.12.5)
}

// AtfPolicies represents ATF policies collection.
type AtfPolicies struct {
	AtfPolicyList []AtfPolicyDetail `json:"atf-policy"` // Air Time Fairness policy list (Live: IOS-XE 17.12.5)
}

// RFTags represents RF tags collection.
type RFTags struct {
	RFTagList []RFTag `json:"rf-tag"` // RF tag configuration list (Live: IOS-XE 17.12.5)
}

// RFProfiles represents RF profiles collection.
type RFProfiles struct {
	RFProfileList []RFProfileDetail `json:"rf-profile"` // Radio-Frequency profile list (Live: IOS-XE 17.12.5)
}

// RFProfileDefaultEntries represents RF profile default entries collection.
type RFProfileDefaultEntries struct {
	RFProfileDefaultEntryList []RFProfileDefaultEntry `json:"rf-profile-default-entry"` // Default RF profile entry list (Live: IOS-XE 17.12.5)
}

// MultiBssidProfile represents Multi-BSSID profile configuration.
type MultiBssidProfile struct {
	ProfileName string `json:"profile-name"`          // 802.11ax Multi BSSID profile name (Live: IOS-XE 17.12.5)
	Description string `json:"description,omitempty"` // Brief Multi BSSID profile description (Live: IOS-XE 17.12.5)
}

// AtfPolicy represents ATF policy structure container.
type AtfPolicy struct {
	AtfPolicyList []AtfPolicyDetail `json:"atf-policy"` // Air Time Fairness policy objects list (Live: IOS-XE 17.12.5)
}

// AtfPolicyDetail represents ATF policy configuration details.
type AtfPolicyDetail struct {
	PolicyID      int    `json:"policy-id"`                // Unique ATF policy identifier (Live: IOS-XE 17.12.5)
	AtfPolicyName string `json:"atfpolicy-name"`           // ATF policy name (Live: IOS-XE 17.12.5)
	PolicyWeight  int    `json:"policy-weight,omitempty"`  // ATF policy weight percentage (Live: IOS-XE 17.12.5)
	ClientSharing bool   `json:"client-sharing,omitempty"` // ATF client fair sharing flag (Live: IOS-XE 17.12.5)
}

// RFTag represents RF tag configuration.
type RFTag struct {
	TagName             string              `json:"tag-name"`                          // RF tag name identifier (Live: IOS-XE 17.12.5)
	Description         string              `json:"description,omitempty"`             // RF tag description (Live: IOS-XE 17.12.5)
	Dot11ARfProfileName string              `json:"dot11a-rf-profile-name,omitempty"`  // 802.11a RF profile name (Live: IOS-XE 17.12.5)
	Dot11BRfProfileName string              `json:"dot11b-rf-profile-name,omitempty"`  // 802.11b RF profile name (Live: IOS-XE 17.12.5)
	Dot116GhzRFProfName string              `json:"dot11-6ghz-rf-prof-name,omitempty"` // 802.11 6GHz RF profile name (Live: IOS-XE 17.12.5)
	RFTagRadioProfiles  *RFTagRadioProfiles `json:"rf-tag-radio-profiles,omitempty"`   // RF tag radio profiles data (Live: IOS-XE 17.12.5)
}

// RFTagRadioProfiles represents RF tag radio profiles collection.
type RFTagRadioProfiles struct {
	RFTagRadioProfile []RFTagRadioProfile `json:"rf-tag-radio-profile"` // Slot specific radio profile list (Live: IOS-XE 17.12.5)
}

// RFTagRadioProfile represents RF tag radio profile configuration.
type RFTagRadioProfile struct {
	SlotID string `json:"slot-id"` // Radio slot identifier (Live: IOS-XE 17.12.5)
	BandID string `json:"band-id"` // Radio band identifier (Live: IOS-XE 17.12.5)
}

// RFProfileDetail represents RF profile configuration details.
type RFProfileDetail struct {
	Name                             string                `json:"name"`                                           // RF profile name identifier (Live: IOS-XE 17.12.5)
	Description                      string                `json:"description,omitempty"`                          // RF profile description (YANG: IOS-XE 17.12.1)
	TxPowerMin                       int                   `json:"tx-power-min,omitempty"`                         // Minimum transmit power in dBm (Live: IOS-XE 17.12.5)
	TxPowerMax                       int                   `json:"tx-power-max,omitempty"`                         // Maximum transmit power in dBm (YANG: IOS-XE 17.12.1)
	TxPowerV1Threshold               int                   `json:"tx-power-v1-threshold,omitempty"`                // TPC version 1 threshold in dBm (Live: IOS-XE 17.12.5)
	TxPowerV2Threshold               int                   `json:"tx-power-v2-threshold,omitempty"`                // TPC version 2 threshold in dBm (YANG: IOS-XE 17.12.1)
	Status                           bool                  `json:"status,omitempty"`                               // RF profile operational state (Live: IOS-XE 17.12.5)
	Band                             string                `json:"band,omitempty"`                                 // Radio frequency band type (Live: IOS-XE 17.12.5)
	DataRate1M                       string                `json:"data-rate-1m,omitempty"`                         // 1 Mbps data rate state (YANG: IOS-XE 17.12.1)
	DataRate2M                       string                `json:"data-rate-2m,omitempty"`                         // 2 Mbps data rate state (YANG: IOS-XE 17.12.1)
	DataRate5_5M                     string                `json:"data-rate-5-5m,omitempty"`                       // 5.5 Mbps data rate state (YANG: IOS-XE 17.12.1)
	DataRate11M                      string                `json:"data-rate-11m,omitempty"`                        // 11 Mbps data rate state (YANG: IOS-XE 17.12.1)
	DataRate6M                       string                `json:"data-rate-6m,omitempty"`                         // 6 Mbps data rate state (Live: IOS-XE 17.12.5)
	DataRate9M                       string                `json:"data-rate-9m,omitempty"`                         // 9 Mbps data rate state (Live: IOS-XE 17.12.5)
	DataRate12M                      string                `json:"data-rate-12m,omitempty"`                        // 12 Mbps data rate state (Live: IOS-XE 17.12.5)
	DataRate18M                      string                `json:"data-rate-18m,omitempty"`                        // 18 Mbps data rate state (Live: IOS-XE 17.12.5)
	DataRate24M                      string                `json:"data-rate-24m,omitempty"`                        // 24 Mbps data rate state (Live: IOS-XE 17.12.5)
	DataRate36M                      string                `json:"data-rate-36m,omitempty"`                        // 36 Mbps data rate state (YANG: IOS-XE 17.12.1)
	DataRate48M                      string                `json:"data-rate-48m,omitempty"`                        // 48 Mbps data rate state (YANG: IOS-XE 17.12.1)
	CoverageDataPacketRSSIThreshold  int                   `json:"coverage-data-packet-rssi-threshold,omitempty"`  // Data packet RSSI threshold dBm (YANG: IOS-XE 17.12.1)
	CoverageVoicePacketRSSIThreshold int                   `json:"coverage-voice-packet-rssi-threshold,omitempty"` // Voice packet RSSI threshold dBm (YANG: IOS-XE 17.12.1)
	LoadBalancingWindow              int                   `json:"load-balancing-window,omitempty"`                // Load balancing window seconds (Live: IOS-XE 17.12.5)
	LoadBalancingDenialCount         int                   `json:"load-balancing-denial-count,omitempty"`          // Load balancing denial count (Live: IOS-XE 17.12.5)
	AtfOperMode                      string                `json:"atf-oper-mode,omitempty"`                        // ATF operation mode (YANG: IOS-XE 17.12.1)
	AtfOptimization                  string                `json:"atf-optimization,omitempty"`                     // ATF optimization policy (YANG: IOS-XE 17.12.1)
	RFMcsEntries                     *RFMcsEntries         `json:"rf-mcs-entries,omitempty"`                       // RF MCS entries data (YANG: IOS-XE 17.12.1)
	RfdcaRemovedChannels             *RfdcaRemovedChannels `json:"rfdca-removed-channels,omitempty"`               // RF DCA removed channels data (Live: IOS-XE 17.12.5)
	ChannelWidthMax                  string                `json:"channel-width-max,omitempty"`                    // Maximum channel width cap (Live: IOS-XE 17.12.5)
	MinNumClients                    int                   `json:"min-num-clients,omitempty"`                      // Minimum client exception level (YANG: IOS-XE 17.12.1)
	RxSenSopThreshold                string                `json:"rx-sen-sop-threshold,omitempty"`                 // RX SOP sensitivity threshold (YANG: IOS-XE 17.12.1)
	BandSelectProbeResponse          bool                  `json:"band-select-probe-response,omitempty"`           // Band select probe response flag (YANG: IOS-XE 17.12.1)
}

// RFMcsEntries represents RF MCS entries collection.
type RFMcsEntries struct {
	RFMcsEntry []RFMcsEntry `json:"rf-mcs-entry"` // RF MCS entry configuration list (YANG: IOS-XE 17.12.1)
}

// RFMcsEntry represents RF MCS entry configuration.
type RFMcsEntry struct {
	RFIndex           int   `json:"rf-index"`                       // RF MCS index identifier (YANG: IOS-XE 17.12.1)
	RF80211NMcsEnable *bool `json:"rf-80211n-mcs-enable,omitempty"` // 802.11n MCS enable flag (YANG: IOS-XE 17.12.1)
}

// RfdcaRemovedChannels represents DCA removed channels collection.
type RfdcaRemovedChannels struct {
	RfdcaRemovedChannel []RfdcaRemovedChannel `json:"rfdca-removed-channel"` // DCA removed channel list (Live: IOS-XE 17.12.5)
}

// RfdcaRemovedChannel represents DCA removed channel configuration.
type RfdcaRemovedChannel struct {
	Channel int `json:"channel"` // Channel number to remove from DCA (Live: IOS-XE 17.12.5)
}

// RFProfileDefaultEntry represents RF profile default entry configuration.
type RFProfileDefaultEntry struct {
	Band                     string               `json:"band"`                                  // Radio frequency band type (YANG: IOS-XE 17.12.1)
	Name                     string               `json:"name"`                                  // RF profile default entry name (YANG: IOS-XE 17.12.1)
	Description              string               `json:"description,omitempty"`                 // RF profile default description (YANG: IOS-XE 17.12.1)
	DataRate1M               string               `json:"data-rate-1m,omitempty"`                // 1 Mbps data rate state (YANG: IOS-XE 17.12.1)
	DataRate2M               string               `json:"data-rate-2m,omitempty"`                // 2 Mbps data rate state (YANG: IOS-XE 17.12.1)
	DataRate5_5M             string               `json:"data-rate-5-5m,omitempty"`              // 5.5 Mbps data rate state (YANG: IOS-XE 17.12.1)
	DataRate11M              string               `json:"data-rate-11m,omitempty"`               // 11 Mbps data rate state (YANG: IOS-XE 17.12.1)
	DataRate6M               string               `json:"data-rate-6m,omitempty"`                // 6 Mbps data rate state (YANG: IOS-XE 17.12.1)
	DataRate9M               string               `json:"data-rate-9m,omitempty"`                // 9 Mbps data rate state (YANG: IOS-XE 17.12.1)
	DataRate12M              string               `json:"data-rate-12m,omitempty"`               // 12 Mbps data rate state (YANG: IOS-XE 17.12.1)
	DataRate18M              string               `json:"data-rate-18m,omitempty"`               // 18 Mbps data rate state (YANG: IOS-XE 17.12.1)
	DataRate24M              string               `json:"data-rate-24m,omitempty"`               // 24 Mbps data rate state (YANG: IOS-XE 17.12.1)
	BandSelectProbeResponse  bool                 `json:"band-select-probe-response,omitempty"`  // Band select probe response flag (YANG: IOS-XE 17.12.1)
	LoadBalancingWindow      int                  `json:"load-balancing-window,omitempty"`       // Load balancing window seconds (YANG: IOS-XE 17.12.1)
	LoadBalancingDenialCount int                  `json:"load-balancing-denial-count,omitempty"` // Load balancing denial count (YANG: IOS-XE 17.12.1)
	AtfOperMode              string               `json:"atf-oper-mode,omitempty"`               // ATF operation mode (YANG: IOS-XE 17.12.1)
	AtfOptimization          string               `json:"atf-optimization,omitempty"`            // ATF optimization policy (YANG: IOS-XE 17.12.1)
	RFMcsDefaultEntries      *RFMcsDefaultEntries `json:"rf-mcs-default-entries,omitempty"`      // RF MCS default entries data (YANG: IOS-XE 17.12.1)
}

// RFMcsDefaultEntries represents RF MCS default entries collection.
type RFMcsDefaultEntries struct {
	RFMcsDefaultEntry []RFMcsDefaultEntry `json:"rf-mcs-default-entry"` // RF MCS default entry list (YANG: IOS-XE 17.12.1)
}

// RFMcsDefaultEntry represents RF MCS default entry configuration.
type RFMcsDefaultEntry struct {
	RFIndex     int `json:"rf-index"`      // RF MCS index identifier (YANG: IOS-XE 17.12.1)
	McsDataRate int `json:"mcs-data-rate"` // MCS data rate value (YANG: IOS-XE 17.12.1)
}
