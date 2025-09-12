// Package client provides data models for client operational data.
package client

import "time"

// ClientOper represents the complete client operational data from WNC 17.12.5.
type ClientOper struct {
	CiscoIOSXEWirelessClientOperClientOperData struct {
		CommonOperData    []CommonOperData    `json:"common-oper-data"`
		Dot11OperData     []Dot11OperData     `json:"dot11-oper-data"`
		MobilityOperData  []MobilityOperData  `json:"mobility-oper-data"`
		MmIfClientStats   []MmIfClientStats   `json:"mm-if-client-stats"`
		MmIfClientHistory []MmIfClientHistory `json:"mm-if-client-history"`
		TrafficStats      []TrafficStats      `json:"traffic-stats"`
		PolicyData        []PolicyData        `json:"policy-data"`
		SisfDBMac         []SisfDBMac         `json:"sisf-db-mac"`
		DcInfo            []DcInfo            `json:"dc-info"`
	} `json:"Cisco-IOS-XE-wireless-client-oper:client-oper-data"`
}

// ClientOperCommonOperData represents the common operational data from WNC 17.12.5.
type ClientOperCommonOperData struct {
	CommonOperData []CommonOperData `json:"Cisco-IOS-XE-wireless-client-oper:common-oper-data"`
}

// ClientOperDot11OperData represents the 802.11 operational data from WNC 17.12.5.
type ClientOperDot11OperData struct {
	Dot11OperData []Dot11OperData `json:"Cisco-IOS-XE-wireless-client-oper:dot11-oper-data"`
}

// ClientOperMobilityOperData represents the mobility operational data from WNC 17.12.5.
type ClientOperMobilityOperData struct {
	MobilityOperData []MobilityOperData `json:"Cisco-IOS-XE-wireless-client-oper:mobility-oper-data"`
}

// ClientOperMmIfClientStats represents the MM interface client statistics.
type ClientOperMmIfClientStats struct {
	MmIfClientStats []MmIfClientStats `json:"Cisco-IOS-XE-wireless-client-oper:mm-if-client-stats"`
}

// ClientOperMmIfClientHistory represents the MM interface client history.
type ClientOperMmIfClientHistory struct {
	MmIfClientHistory []MmIfClientHistory `json:"Cisco-IOS-XE-wireless-client-oper:mm-if-client-history"`
}

// ClientOperTrafficStats represents the client traffic statistics.
type ClientOperTrafficStats struct {
	TrafficStats []TrafficStats `json:"Cisco-IOS-XE-wireless-client-oper:traffic-stats"`
}

// ClientOperPolicyData represents the client policy data.
type ClientOperPolicyData struct {
	PolicyData []PolicyData `json:"Cisco-IOS-XE-wireless-client-oper:policy-data"`
}

// ClientOperSisfDBMac represents the SISF database MAC.
type ClientOperSisfDBMac struct {
	SisfDBMac []SisfDBMac `json:"Cisco-IOS-XE-wireless-client-oper:sisf-db-mac"`
}

// ClientOperDcInfo represents the discovery client information.
type ClientOperDcInfo struct {
	DcInfo []DcInfo `json:"Cisco-IOS-XE-wireless-client-oper:dc-info"`
}

// CommonOperData represents common client operational data.
type CommonOperData struct {
	ClientMac             string `json:"client-mac"`              // Client MAC address
	ApName                string `json:"ap-name"`                 // Access point name
	MsApSlotID            int    `json:"ms-ap-slot-id"`           // AP slot identifier
	MsRadioType           string `json:"ms-radio-type"`           // Radio type identifier
	WlanID                int    `json:"wlan-id"`                 // WLAN identifier
	ClientType            string `json:"client-type"`             // Client device type
	CoState               string `json:"co-state"`                // Client operational state
	AaaOverridePassphrase bool   `json:"aaa-override-passphrase"` // AAA passphrase override flag
	IsTviEnabled          bool   `json:"is-tvi-enabled"`          // TVI enablement status
	WlanPolicy            struct {
		CurrentSwitchingMode  string `json:"current-switching-mode"` // Current data switching mode
		WlanSwitchingMode     string `json:"wlan-switching-mode"`    // WLAN switching mode configuration
		CentralAuthentication string `json:"central-authentication"` // Central authentication setting
		CentralDhcp           bool   `json:"central-dhcp"`           // Central DHCP enablement
		CentralAssocEnable    bool   `json:"central-assoc-enable"`   // Central association enablement
		VlanCentralSwitching  bool   `json:"vlan-central-switching"` // VLAN central switching flag
		IsFabricClient        bool   `json:"is-fabric-client"`       // Fabric client status
		IsGuestFabricClient   bool   `json:"is-guest-fabric-client"` // Guest fabric client status
		UpnBitFlag            string `json:"upn-bit-flag"`           // UPN bit flag setting
	} `json:"wlan-policy"`
	Username           string `json:"username"` // Client username
	GuestLanClientInfo struct {
		WiredVlan       int `json:"wired-vlan"`        // Wired VLAN identifier
		PhyIfid         int `json:"phy-ifid"`          // Physical interface identifier
		IdleTimeSeconds int `json:"idle-time-seconds"` // Idle time in seconds
	} `json:"guest-lan-client-info"`
	MethodID                 string    `json:"method-id"`                   // Authentication method identifier
	L3VlanOverrideReceived   bool      `json:"l3-vlan-override-received"`   // L3 VLAN override received flag
	UpnID                    int       `json:"upn-id"`                      // UPN identifier
	IsLocallyAdministeredMac bool      `json:"is-locally-administered-mac"` // Locally administered MAC flag
	IdleTimeout              int       `json:"idle-timeout"`                // Idle timeout value
	IdleTimestamp            time.Time `json:"idle-timestamp"`              // Idle timestamp
	ClientDuid               string    `json:"client-duid"`                 // Client DHCP unique identifier
	VrfName                  string    `json:"vrf-name"`                    // VRF name

	// Wi-Fi 7 / 802.11be Support (YANG: IOS-XE 17.18.1+)
	L3AccessEnabled bool `json:"l3-access-enabled,omitempty"` // L3 access enablement (YANG: IOS-XE 17.18.1+)
	MultiLinkClient bool `json:"multi-link-client,omitempty"` // Multi-link client capability (YANG: IOS-XE 17.18.1+)
}

// Dot11OperData represents 802.11 operational data.
type Dot11OperData struct {
	MsMacAddress        string    `json:"ms-mac-address"`         // Mobile station MAC address
	Dot11State          string    `json:"dot11-state"`            // 802.11 connection state
	MsBssid             string    `json:"ms-bssid"`               // Mobile station BSSID
	ApMacAddress        string    `json:"ap-mac-address"`         // Access point MAC address
	CurrentChannel      int       `json:"current-channel"`        // Current radio channel
	MsWlanID            int       `json:"ms-wlan-id"`             // Mobile station WLAN ID
	VapSsid             string    `json:"vap-ssid"`               // Virtual AP SSID
	PolicyProfile       string    `json:"policy-profile"`         // Policy profile name
	MsApSlotID          int       `json:"ms-ap-slot-id"`          // Mobile station AP slot ID
	RadioType           string    `json:"radio-type"`             // Radio type identifier
	MsAssociationID     int       `json:"ms-association-id"`      // Mobile station association ID
	MsAuthAlgNum        string    `json:"ms-auth-alg-num"`        // Authentication algorithm number
	MsReasonCode        string    `json:"ms-reason-code"`         // Mobile station reason code
	MsAssocTime         time.Time `json:"ms-assoc-time"`          // Mobile station association time
	Is11GClient         bool      `json:"is-11g-client"`          // 802.11g client capability
	MsSupportedRatesStr string    `json:"ms-supported-rates-str"` // Mobile station supported rates
	MsWifi              struct {
		WpaVersion           string `json:"wpa-version"`             // WPA version information
		CipherSuite          string `json:"cipher-suite"`            // Cipher suite configuration
		AuthKeyMgmt          string `json:"auth-key-mgmt"`           // Authentication key management
		GroupMgmtCipherSuite string `json:"group-mgmt-cipher-suite"` // Group management cipher suite
		GroupCipherSuite     string `json:"group-cipher-suite"`      // Group cipher suite
		PweMode              string `json:"pwe-mode"`                // PWE mode setting
	} `json:"ms-wifi"`
	MsWmeEnabled        bool   `json:"ms-wme-enabled"`         // Mobile station WME enablement
	Dot11WEnabled       bool   `json:"dot11w-enabled"`         // 802.11w enablement status
	EwlcMsPhyType       string `json:"ewlc-ms-phy-type"`       // EWLC mobile station PHY type
	EncryptionType      string `json:"encryption-type"`        // Encryption type
	SecurityMode        string `json:"security-mode"`          // Security mode
	ClientWepPolicyType string `json:"client-wep-policy-type"` // Client WEP policy type
	BssTransCapable     bool   `json:"bss-trans-capable"`      // BSS transition capability
	MsAppleCapable      bool   `json:"ms-apple-capable"`       // Mobile station Apple capability
	WlanProfile         string `json:"wlan-profile"`           // WLAN profile name
	DmsCapable          bool   `json:"dms-capable"`            // DMS capability
	EogreClient         struct {
		IsEogre             bool   `json:"is-eogre"`              // EoGRE client status
		PreviousMatchReason string `json:"previous-match-reason"` // Previous match reason
		MatchReason         string `json:"match-reason"`          // Current match reason
		IsAaaData           bool   `json:"is-aaa-data"`           // AAA data availability
		Realm               string `json:"realm"`                 // Authentication realm
		Vlan                int    `json:"vlan"`                  // VLAN identifier
		Domain              string `json:"domain"`                // Domain name
		PlumbedGw           string `json:"plumbed-gw"`            // Plumbed gateway
		TunnelIfid          int    `json:"tunnel-ifid"`           // Tunnel interface ID
		IsCentralFwd        bool   `json:"is-central-fwd"`        // Central forwarding flag
	} `json:"eogre-client"`
	MsHs20Data struct {
		IsHs20                     bool      `json:"is-hs20"`                      // Hotspot 2.0 capability
		Version                    string    `json:"version"`                      // Hotspot 2.0 version
		ConsortiumOi               string    `json:"consortium-oi"`                // Consortium organization identifier
		PpsMoID                    int       `json:"pps-mo-id"`                    // PPS management object ID
		SwtTimer                   int       `json:"swt-timer"`                    // Session wait timer
		SwtTimestamp               time.Time `json:"swt-timestamp"`                // Session wait timestamp
		TermsConditionsURL         string    `json:"terms-conditions-url"`         // Terms and conditions URL
		SubscriptionRemediationURL string    `json:"subscription-remediation-url"` // Subscription remediation URL
		DeauthReasonURL            string    `json:"deauth-reason-url"`            // Deauthentication reason URL
	} `json:"ms-hs20-data"`
	QosmapCapable                   bool   `json:"qosmap-capable"`  // QoS map capability
	RmCapabilities                  string `json:"rm-capabilities"` // Radio measurement capabilities
	Dot11KRmBeaconMeasReqParameters struct {
		Period              int       `json:"period"`                // Measurement period
		RepeatNum           int       `json:"repeat-num"`            // Repeat number
		OperatingClass      int       `json:"operating-class"`       // Operating class
		ChannelNum          int       `json:"channel-num"`           // Channel number
		MeasMode            string    `json:"meas-mode"`             // Measurement mode
		CurrentBssid        bool      `json:"current-bssid"`         // Current BSSID flag
		Bssid               string    `json:"bssid"`                 // BSSID value
		CurrentSsid         bool      `json:"current-ssid"`          // Current SSID flag
		Ssid                string    `json:"ssid"`                  // SSID value
		DefaultRandInterval bool      `json:"default-rand-interval"` // Default random interval flag
		RandInterval        int       `json:"rand-interval"`         // Random interval
		DefaultMeasDuration bool      `json:"default-meas-duration"` // Default measurement duration flag
		MeasDuration        int       `json:"meas-duration"`         // Measurement duration
		DialogToken         int       `json:"dialog-token"`          // Dialog token
		LastReqTrigger      string    `json:"last-req-trigger"`      // Last request trigger
		LastReqTime         time.Time `json:"last-req-time"`         // Last request time
		NextReqTime         time.Time `json:"next-req-time"`         // Next request time
		LastReportTime      time.Time `json:"last-report-time"`      // Last report time
	} `json:"dot11k-rm-beacon-meas-req-parameters"`
	CellularInfo struct {
		Capable     bool   `json:"capable"`      // Cellular capability
		NetworkType string `json:"network-type"` // Cellular network type
		SignalScale string `json:"signal-scale"` // Signal scale measurement
		CellID      int    `json:"cell-id"`      // Cellular cell ID
	} `json:"cellular-info"`
	WifiDirectClientCapabilities struct {
		WifiDirectCapable bool `json:"wifi-direct-capable"` // Wi-Fi Direct capability
	} `json:"wifi-direct-client-capabilities"`
	WtcSupport      bool `json:"wtc-support"`       // WTC support capability
	AbrSupport      bool `json:"abr-support"`       // ABR support capability
	WtcResp         bool `json:"wtc-resp"`          // WTC response status
	WtcRespCode     int  `json:"wtc-resp-code"`     // WTC response code
	Dot116GhzCap    bool `json:"dot11-6ghz-cap"`    // 802.11 6GHz capability
	LinkLocalEnable bool `json:"link-local-enable"` // Link local enablement

	// Wi-Fi 7 / 802.11be Support (YANG: IOS-XE 17.18.1+)
	EhtCapable      bool          `json:"eht-capable,omitempty"`       // EHT capability (YANG: IOS-XE 17.18.1+)
	MultiLinkClient bool          `json:"multi-link-client,omitempty"` // Multi-link client capability (YANG: IOS-XE 17.18.1+)
	MultilinkInfo   []LinkInfo    `json:"multilink-info,omitempty"`    // Multi-link information (YANG: IOS-XE 17.18.1+)
	KnownLinkInfo   []LinkInfoMin `json:"known-link-info,omitempty"`   // Known link information (YANG: IOS-XE 17.18.1+)
	EmlrMode        string        `json:"emlr-mode,omitempty"`         // EMLR mode (YANG: IOS-XE 17.18.1+)
	StrCapable      bool          `json:"str-capable,omitempty"`       // STR capability (YANG: IOS-XE 17.18.1+)
}

// LinkInfo represents multi-link information for Wi-Fi 7 clients.
type LinkInfo struct {
	Band       string `json:"band"`         // Frequency band (YANG: IOS-XE 17.18.1+)
	BssMacAddr string `json:"bss-mac-addr"` // BSS MAC address (YANG: IOS-XE 17.18.1+)
	SlotID     int    `json:"slot-id"`      // Slot identifier (YANG: IOS-XE 17.18.1+)
	RadioType  string `json:"radio-type"`   // Radio type (YANG: IOS-XE 17.18.1+)
}

// LinkInfoMin represents minimal multi-link information for Wi-Fi 7 clients.
type LinkInfoMin struct {
	StaMac string `json:"sta-mac"` // Station MAC address (YANG: IOS-XE 17.18.1+)
	Band   string `json:"band"`    // Frequency band (YANG: IOS-XE 17.18.1+)
}

// MobilityOperData represents client mobility operational data.
type MobilityOperData struct {
	MsMacAddr           string    `json:"ms-mac-addr"`             // Mobile station MAC address
	MmClientRole        string    `json:"mm-client-role"`          // Mobility manager client role
	MmClientRoamType    string    `json:"mm-client-roam-type"`     // Mobility manager client roam type
	MmInstance          int       `json:"mm-instance"`             // Mobility manager instance
	MmCompleteTimestamp time.Time `json:"mm-complete-timestamp"`   // Mobility manager completion timestamp
	MmRemoteTunnelIP    string    `json:"mm-remote-tunnel-ip"`     // Mobility manager remote tunnel IP
	MmRemoteTunnelSecIP string    `json:"mm-remote-tunnel-sec-ip"` // Mobility manager remote tunnel secondary IP
	MmRemotePlatformID  int       `json:"mm-remote-platform-id"`   // Mobility manager remote platform ID
	MmRemoteTunnelID    int       `json:"mm-remote-tunnel-id"`     // Mobility manager remote tunnel ID
	MmAnchorIP          string    `json:"mm-anchor-ip"`            // Mobility manager anchor IP

	// Wi-Fi 7 / MLO Support - Evidence-based classification
	Dot11RoamType string `json:"dot11-roam-type,omitempty"` // 802.11 roam type
	IsMloAssoc    string `json:"is-mlo-assoc,omitempty"`    // Multi-link operation association status (YANG: IOS-XE 17.18.1+)
}

// MmIfClientStats represents mobility manager interface client statistics.
type MmIfClientStats struct {
	ClientMac  string `json:"client-mac"` // Client MAC address
	MbltyStats struct {
		EventDataAllocs               int `json:"event-data-allocs"`                 // Event data allocations
		EventDataFrees                int `json:"event-data-frees"`                  // Event data frees
		MmifFsmInvalidEvents          int `json:"mmif-fsm-invalid-events"`           // MMIF FSM invalid events
		MmifScheduleErrors            int `json:"mmif-schedule-errors"`              // MMIF schedule errors
		MmifFsmFailure                int `json:"mmif-fsm-failure"`                  // MMIF FSM failures
		MmifIpcFailure                int `json:"mmif-ipc-failure"`                  // MMIF IPC failures
		MmifDBFailure                 int `json:"mmif-db-failure"`                   // MMIF database failures
		MmifInvalidParamsFailure      int `json:"mmif-invalid-params-failure"`       // MMIF invalid parameters failures
		MmifMmMsgDecodeFailure        int `json:"mmif-mm-msg-decode-failure"`        // MMIF MM message decode failures
		MmifUnknownFailure            int `json:"mmif-unknown-failure"`              // MMIF unknown failures
		MmifClientHandoffFailure      int `json:"mmif-client-handoff-failure"`       // MMIF client handoff failures
		MmifClientHandoffSuccess      int `json:"mmif-client-handoff-success"`       // MMIF client handoff successes
		MmifAnchorDeny                int `json:"mmif-anchor-deny"`                  // MMIF anchor denials
		MmifRemoteDelete              int `json:"mmif-remote-delete"`                // MMIF remote deletions
		MmifTunnelDownDelete          int `json:"mmif-tunnel-down-delete"`           // MMIF tunnel down deletions
		MmifMbssidDownEvent           int `json:"mmif-mbssid-down-event"`            // MMIF MBSSID down events
		IntraWncdRoamCount            int `json:"intra-wncd-roam-count"`             // Intra-WNCD roam count
		RemoteInterCtrlrRoams         int `json:"remote-inter-ctrlr-roams"`          // Remote inter-controller roams
		RemoteWebauthPendRoams        int `json:"remote-webauth-pend-roams"`         // Remote webauth pending roams
		AnchorRequestSent             int `json:"anchor-request-sent"`               // Anchor requests sent
		AnchorRequestGrantReceived    int `json:"anchor-request-grant-received"`     // Anchor request grants received
		AnchorRequestDenyReceived     int `json:"anchor-request-deny-received"`      // Anchor request denials received
		AnchorRequestReceived         int `json:"anchor-request-received"`           // Anchor requests received
		AnchorRequestGrantSent        int `json:"anchor-request-grant-sent"`         // Anchor request grants sent
		AnchorRequestDenySent         int `json:"anchor-request-deny-sent"`          // Anchor request denials sent
		HandoffReceivedOk             int `json:"handoff-received-ok"`               // Handoff received OK
		HandoffReceivedGrpMismatch    int `json:"handoff-received-grp-mismatch"`     // Handoff received group mismatch
		HandoffReceivedMsUnknown      int `json:"handoff-received-ms-unknown"`       // Handoff received MS unknown
		HandoffReceivedMsSsid         int `json:"handoff-received-ms-ssid"`          // Handoff received MS SSID
		HandoffReceivedDeny           int `json:"handoff-received-deny"`             // Handoff received denials
		HandoffSentOk                 int `json:"handoff-sent-ok"`                   // Handoff sent OK
		HandoffSentGrpMismatch        int `json:"handoff-sent-grp-mismatch"`         // Handoff sent group mismatch
		HandoffSentMsUnknown          int `json:"handoff-sent-ms-unknown"`           // Handoff sent MS unknown
		HandoffSentMsSsid             int `json:"handoff-sent-ms-ssid"`              // Handoff sent MS SSID
		HandoffSentDeny               int `json:"handoff-sent-deny"`                 // Handoff sent denials
		HandoffReceivedL3VlanOverride int `json:"handoff-received-l3-vlan-override"` // Handoff received L3 VLAN override
		HandoffReceivedUnknownPeer    int `json:"handoff-received-unknown-peer"`     // Handoff received unknown peer
		HandoffSentL3VlanOverride     int `json:"handoff-sent-l3-vlan-override"`     // Handoff sent L3 VLAN override
	} `json:"mblty-stats"`
	IpcStats []struct {
		Type      int    `json:"type"`        // IPC message type
		Allocs    int    `json:"allocs"`      // Allocations
		Frees     int    `json:"frees"`       // Frees
		Tx        int    `json:"tx"`          // Transmitted messages
		Rx        int    `json:"rx"`          // Received messages
		Forwarded int    `json:"forwarded"`   // Forwarded messages
		TxErrors  int    `json:"tx-errors"`   // Transmission errors
		RxErrors  int    `json:"rx-errors"`   // Reception errors
		TxRetries int    `json:"tx-retries"`  // Transmission retries
		Drops     int    `json:"drops"`       // Dropped messages
		Built     int    `json:"built"`       // Built messages
		Processed int    `json:"processed"`   // Processed messages
		MmMsgType string `json:"mm-msg-type"` // MM message type
	} `json:"ipc-stats"`
}

// MmIfClientHistory represents mobility manager interface client history.
type MmIfClientHistory struct {
	ClientMac       string `json:"client-mac"` // Client MAC address
	MobilityHistory struct {
		Entry []struct {
			InstanceID    int       `json:"instance-id"`     // Instance identifier
			MsApSlotID    int       `json:"ms-ap-slot-id"`   // Mobile station AP slot ID
			MsAssocTime   time.Time `json:"ms-assoc-time"`   // Mobile station association time
			Role          string    `json:"role"`            // Client role
			Bssid         string    `json:"bssid"`           // BSSID
			ApName        string    `json:"ap-name"`         // Access point name
			RunLatency    int       `json:"run-latency"`     // Run latency
			Dot11RoamType string    `json:"dot11-roam-type"` // 802.11 roam type
		} `json:"entry"`
	} `json:"mobility-history"`
}

// TrafficStats represents client traffic statistics.
type TrafficStats struct {
	MsMacAddress             string    `json:"ms-mac-address"`              // Mobile station MAC address
	BytesRx                  string    `json:"bytes-rx"`                    // Bytes received
	BytesTx                  string    `json:"bytes-tx"`                    // Bytes transmitted
	PolicyErrs               string    `json:"policy-errs"`                 // Policy errors
	PktsRx                   string    `json:"pkts-rx"`                     // Packets received
	PktsTx                   string    `json:"pkts-tx"`                     // Packets transmitted
	DataRetries              string    `json:"data-retries"`                // Data retries
	RtsRetries               string    `json:"rts-retries"`                 // RTS retries
	DuplicateRcv             string    `json:"duplicate-rcv"`               // Duplicate receives
	DecryptFailed            string    `json:"decrypt-failed"`              // Decryption failures
	MicMismatch              string    `json:"mic-mismatch"`                // MIC mismatches
	MicMissing               string    `json:"mic-missing"`                 // MIC missing
	MostRecentRssi           int       `json:"most-recent-rssi"`            // Most recent RSSI
	MostRecentSnr            int       `json:"most-recent-snr"`             // Most recent SNR
	TxExcessiveRetries       string    `json:"tx-excessive-retries"`        // Transmission excessive retries
	TxRetries                string    `json:"tx-retries"`                  // Transmission retries
	PowerSaveState           int       `json:"power-save-state"`            // Power save state
	CurrentRate              string    `json:"current-rate"`                // Current data rate
	Speed                    int       `json:"speed"`                       // Connection speed
	SpatialStream            int       `json:"spatial-stream"`              // Spatial stream count
	ClientActive             bool      `json:"client-active"`               // Client active status
	GlanStatsUpdateTimestamp time.Time `json:"glan-stats-update-timestamp"` // GLAN stats update timestamp
	GlanIdleUpdateTimestamp  time.Time `json:"glan-idle-update-timestamp"`  // GLAN idle update timestamp
	RxGroupCounter           string    `json:"rx-group-counter"`            // RX group counter
	TxTotalDrops             string    `json:"tx-total-drops"`              // Total transmission drops
}

// PolicyData represents client policy data.
type PolicyData struct {
	Mac         string `json:"mac"`           // Client MAC address
	ResVlanID   int    `json:"res-vlan-id"`   // Resolved VLAN ID
	ResVlanName string `json:"res-vlan-name"` // Resolved VLAN name
}

// SisfDBMac represents SISF database MAC information.
type SisfDBMac struct {
	MacAddr     string `json:"mac-addr"` // MAC address
	Ipv4Binding struct {
		IPKey struct {
			ZoneID int    `json:"zone-id"` // Zone identifier
			IPAddr string `json:"ip-addr"` // IPv4 address
		} `json:"ip-key"`
	} `json:"ipv4-binding"`
	Ipv6Binding []struct {
		Ipv6BindingIPKey struct {
			ZoneID int64  `json:"zone-id"` // Zone identifier
			IPAddr string `json:"ip-addr"` // IPv6 address
		} `json:"ip-key"`
	} `json:"ipv6-binding,omitempty"`
}

// DcInfo represents device classification information.
type DcInfo struct {
	ClientMac        string    `json:"client-mac"`                   // Client MAC address
	DeviceType       string    `json:"device-type"`                  // Device type
	ProtocolMap      string    `json:"protocol-map"`                 // Protocol map
	ConfidenceLevel  int       `json:"confidence-level"`             // Confidence level
	ClassifiedTime   time.Time `json:"classified-time"`              // Classification time
	DayZeroDc        string    `json:"day-zero-dc"`                  // Day zero device classification
	SwVersionSrc     string    `json:"sw-version-src"`               // Software version source
	DeviceOs         string    `json:"device-os,omitempty"`          // Device operating system
	DeviceSubVersion string    `json:"device-sub-version,omitempty"` // Device sub-version
	DeviceOsSrc      string    `json:"device-os-src"`                // Device OS source
	DeviceName       string    `json:"device-name"`                  // Device name
	DeviceVendorSrc  string    `json:"device-vendor-src"`            // Device vendor source
	SalesCodeSrc     string    `json:"sales-code-src"`               // Sales code source
	DeviceSrc        string    `json:"device-src"`                   // Device source
	CountryNameSrc   string    `json:"country-name-src"`             // Country name source
	ModelNameSrc     string    `json:"model-name-src"`               // Model name source
	PowerTypeSrc     string    `json:"power-type-src"`               // Power type source
	HwModelSrc       string    `json:"hw-model-src"`                 // Hardware model source
	DeviceVendor     string    `json:"device-vendor,omitempty"`      // Device vendor
}
