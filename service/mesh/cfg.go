package mesh

// MeshCfg represents mesh configuration data.
type MeshCfg struct {
	MeshCfgData struct {
		Mesh         *Mesh        `json:"mesh,omitempty"` // Global mesh configuration (YANG: IOS-XE 17.12.1)
		MeshProfiles MeshProfiles `json:"mesh-profiles"`  // Mesh profiles configuration (Live: IOS-XE 17.12.6a)
	} `json:"Cisco-IOS-XE-wireless-mesh-cfg:mesh-cfg-data"` // Mesh configuration data (YANG: IOS-XE 17.12.1)
}

// CiscoIOSXEWirelessMeshCfgMeshProfiles represents mesh profiles container structure.
type CiscoIOSXEWirelessMeshCfgMeshProfiles struct {
	MeshProfiles []MeshProfile `json:"Cisco-IOS-XE-wireless-mesh-cfg:mesh-profiles"`
}

// MeshProfiles represents mesh profiles container.
type MeshProfiles struct {
	MeshProfile []MeshProfile `json:"mesh-profile"`
}

// Mesh represents global mesh configuration.
type Mesh struct {
	AssocCountAlarmThresh        *int  `json:"assoc-count-alarm-thresh,omitempty"`         // Association count alarm threshold (YANG: IOS-XE 17.12.1)
	HighSNRAlarmThresh           *int  `json:"high-snr-alarm-thresh,omitempty"`            // High SNR alarm threshold (YANG: IOS-XE 17.12.1)
	LowSNRAlarmThresh            *int  `json:"low-snr-alarm-thresh,omitempty"`             // Low SNR alarm threshold (YANG: IOS-XE 17.12.1)
	MaxMAPChildrenAlarmThresh    *int  `json:"max-map-children-alarm-thresh,omitempty"`    // Maximum MAP children alarm threshold (YANG: IOS-XE 17.12.1)
	MaxRAPChildrenAlarmThresh    *int  `json:"max-rap-children-alarm-thresh,omitempty"`    // Maximum RAP children alarm threshold (YANG: IOS-XE 17.12.1)
	MaxHopAlarmThresh            *int  `json:"max-hop-alarm-thresh,omitempty"`             // Maximum hop alarm threshold (YANG: IOS-XE 17.12.1)
	ParentChangeCountAlarmThresh *int  `json:"parent-change-count-alarm-thresh,omitempty"` // Parent change count alarm threshold (YANG: IOS-XE 17.12.1)
	BackhaulRRMEnabled           *bool `json:"bhaul-rrm-enabled,omitempty"`                // Backhaul RRM enabled flag (YANG: IOS-XE 17.12.1)
	CACEnabled                   *bool `json:"cac-enabled,omitempty"`                      // Call Admission Control enabled flag (YANG: IOS-XE 17.12.1)
	RAPChannelSyncEnabled        *bool `json:"rap-channel-sync-enabled,omitempty"`         // RAP channel sync enabled flag (YANG: IOS-XE 17.12.1)
}

// MeshProfile represents individual mesh profile configuration.
type MeshProfile struct {
	ProfileName           string                `json:"profile-name"`                       // Mesh profile name identifier (Live: IOS-XE 17.12.6a)
	Description           string                `json:"description"`                        // Mesh profile description (Live: IOS-XE 17.12.6a)
	AMSDUEnabled          *bool                 `json:"amsdu-enabled,omitempty"`            // AMSDU aggregation enabled flag (YANG: IOS-XE 17.12.1)
	BackgroundScanEnabled *bool                 `json:"bg-scan-enabled,omitempty"`          // Background scanning enabled flag (YANG: IOS-XE 17.12.1)
	CCNMode               *bool                 `json:"ccn-mode,omitempty"`                 // Cisco Compatible eXtensions mode (YANG: IOS-XE 17.12.1)
	BackhaulClientAccess  *bool                 `json:"bhaul-client-access,omitempty"`      // Backhaul client access enabled flag (YANG: IOS-XE 17.12.1)
	EthVLANTransparent    *bool                 `json:"eth-vlan-transparent,omitempty"`     // Ethernet VLAN transparent mode (YANG: IOS-XE 17.12.1)
	SecurityMode          *string               `json:"security-mode,omitempty"`            // Security mode configuration (YANG: IOS-XE 17.12.1)
	BridgeGroupName       *string               `json:"bridgegroupname,omitempty"`          // Bridge group name assignment (YANG: IOS-XE 17.12.1)
	BGNStrictMatchEnabled *bool                 `json:"bgn-strict-match-enabled,omitempty"` // Bridge group name strict matching (YANG: IOS-XE 17.12.1)
	BackhaulTxRateDot11BG *MeshBackhaulDataRate `json:"bhaul-tx-rate-dot11bg,omitempty"`    // Backhaul data rate for 2.4GHz radio (YANG: IOS-XE 17.12.1)
	BackhaulTxRateDot11A  *MeshBackhaulDataRate `json:"bhaul-tx-rate-dot11a,omitempty"`     // Backhaul data rate for 5GHz radio (YANG: IOS-XE 17.12.1)
}

// MeshBackhaulDataRate represents mesh backhaul data rate configuration.
type MeshBackhaulDataRate struct {
	Type                   *string `json:"type,omitempty"`                      // Backhaul data rate type configuration (YANG: IOS-XE 17.12.1)
	Rate                   *string `json:"rate,omitempty"`                      // 802.11abg transmission rate setting (YANG: IOS-XE 17.12.1)
	Dot11NMCSIndex         *uint8  `json:"dot11n-mcs-index,omitempty"`          // 802.11n MCS index configuration (YANG: IOS-XE 17.12.1)
	Dot11ACMCSIndex        *uint8  `json:"dot11ac-mcs-index,omitempty"`         // 802.11ac MCS index configuration (YANG: IOS-XE 17.12.1)
	SpatialStream          *uint8  `json:"spatial-stream,omitempty"`            // 802.11ac spatial stream count (YANG: IOS-XE 17.12.1)
	Dot11AXMCSIndex        *uint8  `json:"dot11ax-mcs-index,omitempty"`         // 802.11ax MCS index configuration (YANG: IOS-XE 17.12.1)
	Dot11AXSpatialStreamBG *uint8  `json:"dot11ax-spatial-stream-bg,omitempty"` // 802.11ax 2.4GHz spatial stream count (YANG: IOS-XE 17.12.1)
	Dot11AXSpatialStreamA  *uint8  `json:"dot11ax-spatial-stream-a,omitempty"`  // 802.11ax 5GHz spatial stream count (YANG: IOS-XE 17.12.1)
}
