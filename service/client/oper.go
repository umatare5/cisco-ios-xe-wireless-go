// Package client provides Cisco WNC client operational data models.
// Contains complete client operational data structure definitions for RESTCONF API.
package client

import "time"

// CiscoIOSXEWirelessClientOper represents the complete client operational data root structure.
type CiscoIOSXEWirelessClientOper struct {
	CiscoIOSXEWirelessClientOperData struct {
		CommonOperData    []CommonOperData    `json:"common-oper-data"`               // Common operational data for all wireless clients (Live: IOS-XE 17.12.6a)
		Dot11OperData     []Dot11OperData     `json:"dot11-oper-data"`                // IEEE 802.11 operational data for wireless clients (Live: IOS-XE 17.12.6a)
		MobilityOperData  []MobilityOperData  `json:"mobility-oper-data"`             // Mobility operational data for roaming clients (Live: IOS-XE 17.12.6a)
		MmIfClientStats   []MmIfClientStats   `json:"mm-if-client-stats,omitempty"`   // Mobility manager interface client statistics (YANG: IOS-XE 17.12.1)
		MmIfClientHistory []MmIfClientHistory `json:"mm-if-client-history,omitempty"` // Mobility manager interface client history (YANG: IOS-XE 17.12.1)
		TrafficStats      []TrafficStats      `json:"traffic-stats"`                  // Client traffic statistics and counters (Live: IOS-XE 17.12.6a)
		PolicyData        []PolicyData        `json:"policy-data"`                    // Client policy configuration and VLAN data (Live: IOS-XE 17.12.6a)
		SisfDBMac         []SisfDBMac         `json:"sisf-db-mac,omitempty"`          // SISF database MAC address bindings (YANG: IOS-XE 17.12.1)
		DcInfo            []DcInfo            `json:"dc-info,omitempty"`              // Device classification information (YANG: IOS-XE 17.12.1)
	} `json:"Cisco-IOS-XE-wireless-client-oper:client-oper-data"` // Client operational data (Live: IOS-XE 17.12.6a)
}

// CiscoIOSXEWirelessClientOperCommonOperData represents the common operational data.
type CiscoIOSXEWirelessClientOperCommonOperData struct {
	CommonOperData []CommonOperData `json:"Cisco-IOS-XE-wireless-client-oper:common-oper-data"`
}

// CiscoIOSXEWirelessClientOperDot11OperData represents the 802.11 operational data.
type CiscoIOSXEWirelessClientOperDot11OperData struct {
	Dot11OperData []Dot11OperData `json:"Cisco-IOS-XE-wireless-client-oper:dot11-oper-data"`
}

// CiscoIOSXEWirelessClientOperMobilityOperData represents the mobility operational data.
type CiscoIOSXEWirelessClientOperMobilityOperData struct {
	MobilityOperData []MobilityOperData `json:"Cisco-IOS-XE-wireless-client-oper:mobility-oper-data"`
}

// CiscoIOSXEWirelessClientOperMmIfClientStats represents the MM interface client statistics.
type CiscoIOSXEWirelessClientOperMmIfClientStats struct {
	MmIfClientStats []MmIfClientStats `json:"Cisco-IOS-XE-wireless-client-oper:mm-if-client-stats"`
}

// CiscoIOSXEWirelessClientOperMmIfClientHistory represents the MM interface client history.
type CiscoIOSXEWirelessClientOperMmIfClientHistory struct {
	MmIfClientHistory []MmIfClientHistory `json:"Cisco-IOS-XE-wireless-client-oper:mm-if-client-history"`
}

// CiscoIOSXEWirelessClientOperTrafficStatsData represents the traffic statistics data.
type CiscoIOSXEWirelessClientOperTrafficStatsData struct {
	TrafficStats []TrafficStats `json:"Cisco-IOS-XE-wireless-client-oper:traffic-stats"`
}

// CiscoIOSXEWirelessClientOperPolicyData represents the client policy data.
type CiscoIOSXEWirelessClientOperPolicyData struct {
	PolicyData []PolicyData `json:"Cisco-IOS-XE-wireless-client-oper:policy-data"`
}

// CiscoIOSXEWirelessClientOperSisfDBMac represents the SISF database MAC.
type CiscoIOSXEWirelessClientOperSisfDBMac struct {
	SisfDBMac []SisfDBMac `json:"Cisco-IOS-XE-wireless-client-oper:sisf-db-mac"`
}

// CiscoIOSXEWirelessClientOperDcInfo represents the discovery client information.
type CiscoIOSXEWirelessClientOperDcInfo struct {
	DcInfo []DcInfo `json:"Cisco-IOS-XE-wireless-client-oper:dc-info"`
}

// CommonOperData represents common client operational data.
type CommonOperData struct {
	ClientMAC                string           `json:"client-mac"`                  // MAC address used as network address for wireless stations (Live: IOS-XE 17.12.6a)
	ApName                   string           `json:"ap-name"`                     // Access Point name client is connected to (Live: IOS-XE 17.12.6a)
	MsApSlotID               int              `json:"ms-ap-slot-id"`               // Radio slot on AP client is connected to (Live: IOS-XE 17.12.6a)
	MsRadioType              string           `json:"ms-radio-type"`               // Wireless Radio type client connected with (Live: IOS-XE 17.12.6a)
	WlanID                   int              `json:"wlan-id"`                     // Unique Wireless LAN identifier client connected to (Live: IOS-XE 17.12.6a)
	ClientType               string           `json:"client-type"`                 // Wireless network type based on traffic switching mode (Live: IOS-XE 17.12.6a)
	CoState                  string           `json:"co-state"`                    // Last association phase client completed successfully (Live: IOS-XE 17.12.6a)
	AaaOverridePassphrase    bool             `json:"aaa-override-passphrase"`     // AAA override passphrase enabled status (Live: IOS-XE 17.12.6a)
	IsTviEnabled             bool             `json:"is-tvi-enabled"`              // Encrypted Traffic Analytics enablement for client (Live: IOS-XE 17.12.6a)
	WlanPolicy               ClientWlanPolicy `json:"wlan-policy"`                 // Client WLAN policy configuration (Live: IOS-XE 17.12.6a)
	Username                 string           `json:"username"`                    // Client username (Live: IOS-XE 17.12.6a)
	GuestLanClientInfo       ClientGuestInfo  `json:"guest-lan-client-info"`       // Guest LAN client information (Live: IOS-XE 17.12.6a)
	MethodID                 string           `json:"method-id"`                   // Method Identifier (Live: IOS-XE 17.12.6a)
	L3VlanOverrideReceived   bool             `json:"l3-vlan-override-received"`   // VLAN Override received after L3 authentication (Live: IOS-XE 17.12.6a)
	UpnID                    int              `json:"upn-id"`                      // User Defined Network Identity (Live: IOS-XE 17.12.6a)
	IsLocallyAdministeredMAC bool             `json:"is-locally-administered-mac"` // Client MAC address locally administered status (Live: IOS-XE 17.12.6a)
	IdleTimeout              int              `json:"idle-timeout"`                // Client idle timeout before deletion (Live: IOS-XE 17.12.6a)
	IdleTimestamp            time.Time        `json:"idle-timestamp"`              // Timestamp when client moved to idle state (Live: IOS-XE 17.12.6a)
	ClientDuid               string           `json:"client-duid"`                 // Client device user identity (Live: IOS-XE 17.12.6a)
	VrfName                  string           `json:"vrf-name"`                    // VRF Name (Live: IOS-XE 17.12.6a)

	// Wi-Fi 7 / 802.11be Support (YANG: IOS-XE 17.18.1)
	L3AccessEnabled bool `json:"l3-access-enabled"` // Client Layer 3 access enabled status (Live: IOS-XE 17.15.4b)
	MultiLinkClient bool `json:"multi-link-client"` // Multi Link client capability (Live: IOS-XE 17.15.4b)
}

// Dot11OperData represents 802.11 operational data.
type Dot11OperData struct {
	MsMACAddress        string    `json:"ms-mac-address"`         // Mac Address of the Client (Live: IOS-XE 17.12.6a)
	Dot11State          string    `json:"dot11-state"`            // DOT11 status for Client (Live: IOS-XE 17.12.6a)
	MsBssid             string    `json:"ms-bssid"`               // Basic Service Set Identifier client connected to (Live: IOS-XE 17.12.6a)
	ApMACAddress        string    `json:"ap-mac-address"`         // MAC Address of AP client has joined (Live: IOS-XE 17.12.6a)
	CurrentChannel      int       `json:"current-channel"`        // Current Channel client communicating on (Live: IOS-XE 17.12.6a)
	MsWlanID            int       `json:"ms-wlan-id"`             // Wireless LAN ID client connected to (Live: IOS-XE 17.12.6a)
	VapSsid             string    `json:"vap-ssid"`               // Service Set Identifier of Wireless LAN (Live: IOS-XE 17.12.6a)
	PolicyProfile       string    `json:"policy-profile"`         // Policy profile applied on WLAN (Live: IOS-XE 17.12.6a)
	MsApSlotID          int       `json:"ms-ap-slot-id"`          // Slot ID of AP radio client connected on (Live: IOS-XE 17.12.6a)
	RadioType           string    `json:"radio-type"`             // Type of Radio of AP client associated to (Live: IOS-XE 17.12.6a)
	MsAssociationID     int       `json:"ms-association-id"`      // Association ID of mobile station (Live: IOS-XE 17.12.6a)
	MsAuthAlgNum        string    `json:"ms-auth-alg-num"`        // Authentication algorithm (Live: IOS-XE 17.12.6a)
	MsReasonCode        string    `json:"ms-reason-code"`         // Reason code for deauth/disassoc frames (Live: IOS-XE 17.12.6a)
	MsAssocTime         time.Time `json:"ms-assoc-time"`          // Time association request received (Live: IOS-XE 17.12.6a)
	Is11GClient         bool      `json:"is-11g-client"`          // IEEE 802.11g protocol client indicator (Live: IOS-XE 17.12.6a)
	MsSupportedRatesStr string    `json:"ms-supported-rates-str"` // Supported radio rates by mobile station (Live: IOS-XE 17.12.6a)
	MsWifi              struct {
		WpaVersion           string `json:"wpa-version"`             // WPA version of the client (Live: IOS-XE 17.12.6a)
		CipherSuite          string `json:"cipher-suite"`            // IEEE 802.11i Cipher Suite type (Live: IOS-XE 17.12.6a)
		AuthKeyMgmt          string `json:"auth-key-mgmt"`           // IEEE 802.11i Auth Key Management (Live: IOS-XE 17.12.6a)
		GroupMgmtCipherSuite string `json:"group-mgmt-cipher-suite"` // IEEE 802.11i Group Management Cipher Suite (Live: IOS-XE 17.12.6a)
		GroupCipherSuite     string `json:"group-cipher-suite"`      // IEEE 802.11i Group Cipher Suite (Live: IOS-XE 17.12.6a)
		PweMode              string `json:"pwe-mode"`                // SAE Password Element Mode (Live: IOS-XE 17.12.6a)
	} `json:"ms-wifi"`
	MsWmeEnabled        bool   `json:"ms-wme-enabled"`         // Wireless Multimedia Extensions enabled indicator (Live: IOS-XE 17.12.6a)
	Dot11WEnabled       bool   `json:"dot11w-enabled"`         // IEEE 802.11w feature enabled indicator (Live: IOS-XE 17.12.6a)
	EwlcMsPhyType       string `json:"ewlc-ms-phy-type"`       // Radio PHY type client connected to (Live: IOS-XE 17.12.6a)
	EncryptionType      string `json:"encryption-type"`        // Encryption policy client uses with AP (Live: IOS-XE 17.12.6a)
	SecurityMode        string `json:"security-mode"`          // Security mode for client association (Live: IOS-XE 17.12.6a)
	ClientWepPolicyType string `json:"client-wep-policy-type"` // Client Wired Equivalent Privacy policy type (Live: IOS-XE 17.12.6a)
	BssTransCapable     bool   `json:"bss-trans-capable"`      // IEEE 802.11v capable indicator (Live: IOS-XE 17.12.6a)
	MsAppleCapable      bool   `json:"ms-apple-capable"`       // Client Fastlane Support indicator (Live: IOS-XE 17.12.6a)
	WlanProfile         string `json:"wlan-profile"`           // Profile applied on Wireless/Remote/Guest LAN (Live: IOS-XE 17.12.6a)
	DmsCapable          bool   `json:"dms-capable"`            // Directed Multicast Service capable indicator (Live: IOS-XE 17.12.6a)
	EogreClient         struct {
		IsEogre             bool   `json:"is-eogre"`              // Whether this is an EoGRE client (Live: IOS-XE 17.12.6a)
		PreviousMatchReason string `json:"previous-match-reason"` // Previous EoGRE client match process output (Live: IOS-XE 17.12.6a)
		MatchReason         string `json:"match-reason"`          // EoGRE client match process output (Live: IOS-XE 17.12.6a)
		IsAaaData           bool   `json:"is-aaa-data"`           // AAA override received for this client (Live: IOS-XE 17.12.6a)
		Realm               string `json:"realm"`                 // Client's realm matching EoGRE rule (Live: IOS-XE 17.12.6a)
		Vlan                int    `json:"vlan"`                  // VLAN tagging for EoGRE client traffic (Live: IOS-XE 17.12.6a)
		Domain              string `json:"domain"`                // EoGRE domain (Live: IOS-XE 17.12.6a)
		PlumbedGw           string `json:"plumbed-gw"`            // Tunnel Gateway Name for client traffic (Live: IOS-XE 17.12.6a)
		TunnelIfid          int    `json:"tunnel-ifid"`           // Tunnel Gateway datapath index (Live: IOS-XE 17.12.6a)
		IsCentralFwd        bool   `json:"is-central-fwd"`        // Client is centrally forwarded (Live: IOS-XE 17.12.6a)
	} `json:"eogre-client"`
	MsHs20Data struct {
		IsHs20                     bool      `json:"is-hs20"`                      // Hotspot 2.0 data received for this client (Live: IOS-XE 17.12.6a)
		Version                    string    `json:"version"`                      // Supported Hotspot release version (Live: IOS-XE 17.12.6a)
		ConsortiumOi               string    `json:"consortium-oi"`                // Roaming consortium membership OI (Live: IOS-XE 17.12.6a)
		PpsMoID                    int       `json:"pps-mo-id"`                    // Per provider subscription (Live: IOS-XE 17.12.6a)
		SwtTimer                   int       `json:"swt-timer"`                    // Session Warning Timer value (Live: IOS-XE 17.12.6a)
		SwtTimestamp               time.Time `json:"swt-timestamp"`                // SWT timestamp (Live: IOS-XE 17.12.6a)
		TermsConditionsURL         string    `json:"terms-conditions-url"`         // Terms and conditions URL (Live: IOS-XE 17.12.6a)
		SubscriptionRemediationURL string    `json:"subscription-remediation-url"` // Subscription remediation URL (Live: IOS-XE 17.12.6a)
		DeauthReasonURL            string    `json:"deauth-reason-url"`            // WNM deauthentication reason URL (Live: IOS-XE 17.12.6a)
	} `json:"ms-hs20-data"`
	QosmapCapable                   bool   `json:"qosmap-capable"`  // QoS map capability indicator (YANG: IOS-XE 17.12.1)
	RmCapabilities                  string `json:"rm-capabilities"` // Radio measurement capabilities (YANG: IOS-XE 17.12.1)
	Dot11KRmBeaconMeasReqParameters struct {
		Period              int       `json:"period"`                // Beacon measurement request period (YANG: IOS-XE 17.12.1)
		RepeatNum           int       `json:"repeat-num"`            // Measurement repeat number (YANG: IOS-XE 17.12.1)
		OperatingClass      int       `json:"operating-class"`       // IEEE 802.11 operating class (YANG: IOS-XE 17.12.1)
		ChannelNum          int       `json:"channel-num"`           // Channel number for measurement (YANG: IOS-XE 17.12.1)
		MeasMode            string    `json:"meas-mode"`             // Measurement mode active/passive (YANG: IOS-XE 17.12.1)
		CurrentBssid        bool      `json:"current-bssid"`         // Use current BSSID flag (YANG: IOS-XE 17.12.1)
		Bssid               string    `json:"bssid"`                 // Target BSSID for measurement (YANG: IOS-XE 17.12.1)
		CurrentSsid         bool      `json:"current-ssid"`          // Use current SSID flag (YANG: IOS-XE 17.12.1)
		Ssid                string    `json:"ssid"`                  // Target SSID for measurement (YANG: IOS-XE 17.12.1)
		DefaultRandInterval bool      `json:"default-rand-interval"` // Use default randomization interval (YANG: IOS-XE 17.12.1)
		RandInterval        int       `json:"rand-interval"`         // Randomization interval value (YANG: IOS-XE 17.12.1)
		DefaultMeasDuration bool      `json:"default-meas-duration"` // Use default measurement duration (YANG: IOS-XE 17.12.1)
		MeasDuration        int       `json:"meas-duration"`         // Measurement duration value (YANG: IOS-XE 17.12.1)
		DialogToken         int       `json:"dialog-token"`          // Dialog token for request tracking (YANG: IOS-XE 17.12.1)
		LastReqTrigger      string    `json:"last-req-trigger"`      // Last request trigger event (YANG: IOS-XE 17.12.1)
		LastReqTime         time.Time `json:"last-req-time"`         // Last request timestamp (YANG: IOS-XE 17.12.1)
		NextReqTime         time.Time `json:"next-req-time"`         // Next request timestamp (YANG: IOS-XE 17.12.1)
		LastReportTime      time.Time `json:"last-report-time"`      // Last report timestamp (YANG: IOS-XE 17.12.1)
	} `json:"dot11k-rm-beacon-meas-req-parameters"`
	CellularInfo struct {
		Capable     bool   `json:"capable"`      // Cellular capability indicator (YANG: IOS-XE 17.12.1)
		NetworkType string `json:"network-type"` // Cellular network type (YANG: IOS-XE 17.12.1)
		SignalScale string `json:"signal-scale"` // Signal scale measurement (YANG: IOS-XE 17.12.1)
		CellID      int    `json:"cell-id"`      // Cellular cell ID (YANG: IOS-XE 17.12.1)
	} `json:"cellular-info"`
	WifiDirectClientCapabilities struct {
		WifiDirectCapable bool `json:"wifi-direct-capable"` // Wi-Fi Direct capability indicator (YANG: IOS-XE 17.12.1)
	} `json:"wifi-direct-client-capabilities"`
	WtcSupport      bool `json:"wtc-support"`       // WTC support capability (YANG: IOS-XE 17.12.1)
	AbrSupport      bool `json:"abr-support"`       // ABR support capability (YANG: IOS-XE 17.12.1)
	WtcResp         bool `json:"wtc-resp"`          // WTC response status (YANG: IOS-XE 17.12.1)
	WtcRespCode     int  `json:"wtc-resp-code"`     // WTC response code (YANG: IOS-XE 17.12.1)
	Dot116GhzCap    bool `json:"dot11-6ghz-cap"`    // 802.11 6GHz capability (YANG: IOS-XE 17.12.1)
	LinkLocalEnable bool `json:"link-local-enable"` // Link local enablement (YANG: IOS-XE 17.12.1)

	// Wi-Fi 7 / 802.11be Support (YANG: IOS-XE 17.18.1)
	EhtCapable      bool          `json:"eht-capable,omitempty"`       // EHT capability indicator (YANG: IOS-XE 17.18.1)
	MultiLinkClient bool          `json:"multi-link-client,omitempty"` // Multi-link client capability (YANG: IOS-XE 17.18.1)
	MultilinkInfo   []LinkInfo    `json:"multilink-info,omitempty"`    // Multi-link information of the client (YANG: IOS-XE 17.18.1)
	KnownLinkInfo   []LinkInfoMin `json:"known-link-info,omitempty"`   // Known link information (YANG: IOS-XE 17.18.1)
	EmlrMode        string        `json:"emlr-mode,omitempty"`         // Enhanced multilink radio operation mode (YANG: IOS-XE 17.18.1)
	StrCapable      bool          `json:"str-capable,omitempty"`       // Simultaneous transmit and receive capability (YANG: IOS-XE 17.18.1)
}

// LinkInfo represents multi-link information for Wi-Fi 7 clients.
type LinkInfo struct {
	Band       string `json:"band"`         // Frequency band identifier (YANG: IOS-XE 17.18.1)
	BssMACAddr string `json:"bss-mac-addr"` // BSS MAC address (YANG: IOS-XE 17.18.1)
	SlotID     int    `json:"slot-id"`      // Slot identifier (YANG: IOS-XE 17.18.1)
	RadioType  string `json:"radio-type"`   // Radio type identifier (YANG: IOS-XE 17.18.1)
}

// LinkInfoMin represents minimal multi-link information for Wi-Fi 7 clients.
type LinkInfoMin struct {
	StaMAC string `json:"sta-mac"` // Station MAC address (YANG: IOS-XE 17.18.1)
	Band   string `json:"band"`    // Frequency band identifier (YANG: IOS-XE 17.18.1)
}

// MobilityOperData represents client mobility operational data.
type MobilityOperData struct {
	MsMACAddr           string    `json:"ms-mac-addr"`             // MAC address of wireless mobile station (Live: IOS-XE 17.12.6a)
	MmClientRole        string    `json:"mm-client-role"`          // Mobility role on Wireless LAN Controller (Live: IOS-XE 17.12.6a)
	MmClientRoamType    string    `json:"mm-client-roam-type"`     // Layer 2 or Layer 3 mobility roam type (Live: IOS-XE 17.12.6a)
	MmInstance          int       `json:"mm-instance"`             // Inter-controller roam count performed by client (Live: IOS-XE 17.12.6a)
	MmCompleteTimestamp time.Time `json:"mm-complete-timestamp"`   // Mobility discovery completion timestamp (Live: IOS-XE 17.12.6a)
	MmRemoteTunnelIP    string    `json:"mm-remote-tunnel-ip"`     // Primary IP of mobility peer for anchor/foreign client (Live: IOS-XE 17.12.6a)
	MmRemoteTunnelSecIP string    `json:"mm-remote-tunnel-sec-ip"` // Secondary IP of mobility peer (Live: IOS-XE 17.12.6a)
	MmRemotePlatformID  int       `json:"mm-remote-platform-id"`   // Platform ID of mobility peer (Live: IOS-XE 17.12.6a)
	MmRemoteTunnelID    int       `json:"mm-remote-tunnel-id"`     // Mobility peer tunnel identifier (Live: IOS-XE 17.12.6a)
	MmAnchorIP          string    `json:"mm-anchor-ip"`            // Anchor WLC address for foreign client (Live: IOS-XE 17.12.6a)

	// Wi-Fi 7 / MLO Support - Evidence-based classification
	Dot11RoamType string `json:"dot11-roam-type,omitempty"` // 802.11 roam type (YANG: IOS-XE 17.12.1)
	IsMloAssoc    string `json:"is-mlo-assoc,omitempty"`    // Multi-link operation association status (YANG: IOS-XE 17.18.1)
}

// MmIfClientStats represents mobility manager interface client statistics.
type MmIfClientStats struct {
	ClientMAC  string `json:"client-mac"` // Client MAC address (YANG: IOS-XE 17.12.1)
	MbltyStats struct {
		EventDataAllocs               int `json:"event-data-allocs"`                 // Event data allocations count (YANG: IOS-XE 17.12.1)
		EventDataFrees                int `json:"event-data-frees"`                  // Event data deallocations count (YANG: IOS-XE 17.12.1)
		MmifFsmInvalidEvents          int `json:"mmif-fsm-invalid-events"`           // MMIF finite state machine invalid events (YANG: IOS-XE 17.12.1)
		MmifScheduleErrors            int `json:"mmif-schedule-errors"`              // MMIF scheduling error count (YANG: IOS-XE 17.12.1)
		MmifFsmFailure                int `json:"mmif-fsm-failure"`                  // MMIF finite state machine failure count (YANG: IOS-XE 17.12.1)
		MmifIpcFailure                int `json:"mmif-ipc-failure"`                  // MMIF inter-process communication failures (YANG: IOS-XE 17.12.1)
		MmifDBFailure                 int `json:"mmif-db-failure"`                   // MMIF database operation failure count (YANG: IOS-XE 17.12.1)
		MmifInvalidParamsFailure      int `json:"mmif-invalid-params-failure"`       // MMIF invalid parameter failure count (YANG: IOS-XE 17.12.1)
		MmifMmMsgDecodeFailure        int `json:"mmif-mm-msg-decode-failure"`        // MMIF mobility message decode failures (YANG: IOS-XE 17.12.1)
		MmifUnknownFailure            int `json:"mmif-unknown-failure"`              // MMIF unknown failure count (YANG: IOS-XE 17.12.1)
		MmifClientHandoffFailure      int `json:"mmif-client-handoff-failure"`       // MMIF client handoff failure count (YANG: IOS-XE 17.12.1)
		MmifClientHandoffSuccess      int `json:"mmif-client-handoff-success"`       // MMIF client handoff success count (YANG: IOS-XE 17.12.1)
		MmifAnchorDeny                int `json:"mmif-anchor-deny"`                  // MMIF anchor denial count (YANG: IOS-XE 17.12.1)
		MmifRemoteDelete              int `json:"mmif-remote-delete"`                // MMIF remote deletion count (YANG: IOS-XE 17.12.1)
		MmifTunnelDownDelete          int `json:"mmif-tunnel-down-delete"`           // MMIF tunnel down deletion count (YANG: IOS-XE 17.12.1)
		MmifMbssidDownEvent           int `json:"mmif-mbssid-down-event"`            // MMIF multi-BSSID down event count (YANG: IOS-XE 17.12.1)
		IntraWncdRoamCount            int `json:"intra-wncd-roam-count"`             // Intra-WNC daemon roam count (YANG: IOS-XE 17.12.1)
		RemoteInterCtrlrRoams         int `json:"remote-inter-ctrlr-roams"`          // Remote inter-controller roam count (YANG: IOS-XE 17.12.1)
		RemoteWebauthPendRoams        int `json:"remote-webauth-pend-roams"`         // Remote web auth pending roam count (YANG: IOS-XE 17.12.1)
		AnchorRequestSent             int `json:"anchor-request-sent"`               // Anchor request messages sent count (YANG: IOS-XE 17.12.1)
		AnchorRequestGrantReceived    int `json:"anchor-request-grant-received"`     // Anchor request grant responses received (YANG: IOS-XE 17.12.1)
		AnchorRequestDenyReceived     int `json:"anchor-request-deny-received"`      // Anchor request deny responses received (YANG: IOS-XE 17.12.1)
		AnchorRequestReceived         int `json:"anchor-request-received"`           // Anchor request messages received count (YANG: IOS-XE 17.12.1)
		AnchorRequestGrantSent        int `json:"anchor-request-grant-sent"`         // Anchor request grant responses sent (YANG: IOS-XE 17.12.1)
		AnchorRequestDenySent         int `json:"anchor-request-deny-sent"`          // Anchor request deny responses sent (YANG: IOS-XE 17.12.1)
		HandoffReceivedOk             int `json:"handoff-received-ok"`               // Successful handoff messages received count (YANG: IOS-XE 17.12.1)
		HandoffReceivedGrpMismatch    int `json:"handoff-received-grp-mismatch"`     // Handoff received group mismatch count (YANG: IOS-XE 17.12.1)
		HandoffReceivedMsUnknown      int `json:"handoff-received-ms-unknown"`       // Handoff received unknown mobile station (YANG: IOS-XE 17.12.1)
		HandoffReceivedMsSsid         int `json:"handoff-received-ms-ssid"`          // Handoff received mobile station SSID (YANG: IOS-XE 17.12.1)
		HandoffReceivedDeny           int `json:"handoff-received-deny"`             // Handoff deny messages received count (YANG: IOS-XE 17.12.1)
		HandoffSentOk                 int `json:"handoff-sent-ok"`                   // Successful handoff messages sent count (YANG: IOS-XE 17.12.1)
		HandoffSentGrpMismatch        int `json:"handoff-sent-grp-mismatch"`         // Handoff sent group mismatch count (YANG: IOS-XE 17.12.1)
		HandoffSentMsUnknown          int `json:"handoff-sent-ms-unknown"`           // Handoff sent unknown mobile station (YANG: IOS-XE 17.12.1)
		HandoffSentMsSsid             int `json:"handoff-sent-ms-ssid"`              // Handoff sent mobile station SSID (YANG: IOS-XE 17.12.1)
		HandoffSentDeny               int `json:"handoff-sent-deny"`                 // Handoff deny messages sent count (YANG: IOS-XE 17.12.1)
		HandoffReceivedL3VlanOverride int `json:"handoff-received-l3-vlan-override"` // Handoff received L3 VLAN override count (YANG: IOS-XE 17.12.1)
		HandoffReceivedUnknownPeer    int `json:"handoff-received-unknown-peer"`     // Handoff received unknown peer count (YANG: IOS-XE 17.12.1)
		HandoffSentL3VlanOverride     int `json:"handoff-sent-l3-vlan-override"`     // Handoff sent L3 VLAN override count (YANG: IOS-XE 17.12.1)
	} `json:"mblty-stats"`
	IpcStats []struct {
		Type      int    `json:"type"`        // IPC message type identifier (YANG: IOS-XE 17.12.1)
		Allocs    int    `json:"allocs"`      // IPC message allocation count (YANG: IOS-XE 17.12.1)
		Frees     int    `json:"frees"`       // IPC message deallocation count (YANG: IOS-XE 17.12.1)
		Tx        int    `json:"tx"`          // IPC messages transmitted count (YANG: IOS-XE 17.12.1)
		Rx        int    `json:"rx"`          // IPC messages received count (YANG: IOS-XE 17.12.1)
		Forwarded int    `json:"forwarded"`   // IPC messages forwarded count (YANG: IOS-XE 17.12.1)
		TxErrors  int    `json:"tx-errors"`   // IPC transmission error count (YANG: IOS-XE 17.12.1)
		RxErrors  int    `json:"rx-errors"`   // IPC reception error count (YANG: IOS-XE 17.12.1)
		TxRetries int    `json:"tx-retries"`  // IPC transmission retry count (YANG: IOS-XE 17.12.1)
		Drops     int    `json:"drops"`       // IPC messages dropped count (YANG: IOS-XE 17.12.1)
		Built     int    `json:"built"`       // IPC messages built count (YANG: IOS-XE 17.12.1)
		Processed int    `json:"processed"`   // IPC messages processed count (YANG: IOS-XE 17.12.1)
		MmMsgType string `json:"mm-msg-type"` // Mobility manager message type (YANG: IOS-XE 17.12.1)
	} `json:"ipc-stats"`
}

// MmIfClientHistory represents mobility manager interface client history.
type MmIfClientHistory struct {
	ClientMAC       string `json:"client-mac"` // Client MAC address (YANG: IOS-XE 17.12.1)
	MobilityHistory struct {
		Entry []struct {
			InstanceID    int       `json:"instance-id"`     // Mobility instance identifier (YANG: IOS-XE 17.12.1)
			MsApSlotID    int       `json:"ms-ap-slot-id"`   // Mobile station AP slot ID (YANG: IOS-XE 17.12.1)
			MsAssocTime   time.Time `json:"ms-assoc-time"`   // Mobile station association timestamp (YANG: IOS-XE 17.12.1)
			Role          string    `json:"role"`            // Client mobility role (YANG: IOS-XE 17.12.1)
			Bssid         string    `json:"bssid"`           // Basic Service Set identifier (YANG: IOS-XE 17.12.1)
			ApName        string    `json:"ap-name"`         // Access point name (YANG: IOS-XE 17.12.1)
			RunLatency    int       `json:"run-latency"`     // Mobility operation latency (YANG: IOS-XE 17.12.1)
			Dot11RoamType string    `json:"dot11-roam-type"` // IEEE 802.11 roam type (YANG: IOS-XE 17.12.1)
		} `json:"entry"`
	} `json:"mobility-history"`
}

// TrafficStats represents client traffic statistics.
type TrafficStats struct {
	MsMACAddress             string    `json:"ms-mac-address"`              // MAC address used as network address for mobile station (Live: IOS-XE 17.12.6a)
	BytesRx                  string    `json:"bytes-rx"`                    // Bytes of wireless data traffic received (Live: IOS-XE 17.12.6a)
	BytesTx                  string    `json:"bytes-tx"`                    // Bytes of wireless data traffic transmitted (Live: IOS-XE 17.12.6a)
	PolicyErrs               string    `json:"policy-errs"`                 // Mobile station policy errors (Live: IOS-XE 17.12.6a)
	PktsRx                   string    `json:"pkts-rx"`                     // Packets of wireless data traffic received (Live: IOS-XE 17.12.6a)
	PktsTx                   string    `json:"pkts-tx"`                     // Packets of wireless data traffic transmitted (Live: IOS-XE 17.12.6a)
	DataRetries              string    `json:"data-retries"`                // Retries mobile station executed for data traffic (Live: IOS-XE 17.12.6a)
	RtsRetries               string    `json:"rts-retries"`                 // Frames received with retry bit set (Live: IOS-XE 17.12.6a)
	DuplicateRcv             string    `json:"duplicate-rcv"`               // Duplicate frames received (Live: IOS-XE 17.12.6a)
	DecryptFailed            string    `json:"decrypt-failed"`              // Decrypt failed packets (Live: IOS-XE 17.12.6a)
	MicMismatch              string    `json:"mic-mismatch"`                // Packets with Message Integrity Check mismatch (Live: IOS-XE 17.12.6a)
	MicMissing               string    `json:"mic-missing"`                 // Packets with Message Integrity Check missing (Live: IOS-XE 17.12.6a)
	MostRecentRSSI           int       `json:"most-recent-rssi"`            // Last updated Radio Signal Strength indicator (Live: IOS-XE 17.12.6a)
	MostRecentSNR            int       `json:"most-recent-snr"`             // Last updated Signal To Noise Ratio (Live: IOS-XE 17.12.6a)
	TxExcessiveRetries       string    `json:"tx-excessive-retries"`        // Mobile station excessive retries (Live: IOS-XE 17.12.6a)
	TxRetries                string    `json:"tx-retries"`                  // Frames transmitted with Retry bit set (Live: IOS-XE 17.12.6a)
	PowerSaveState           int       `json:"power-save-state"`            // Power save state (Live: IOS-XE 17.12.6a)
	CurrentRate              string    `json:"current-rate"`                // Current Rate (Live: IOS-XE 17.12.6a)
	Speed                    int       `json:"speed"`                       // Latest speed of connected client (Live: IOS-XE 17.12.6a)
	SpatialStream            int       `json:"spatial-stream"`              // Number of Spatial Streams supported (Live: IOS-XE 17.12.6a)
	ClientActive             bool      `json:"client-active"`               // Client status as active identification (Live: IOS-XE 17.12.6a)
	GlanStatsUpdateTimestamp time.Time `json:"glan-stats-update-timestamp"` // Guest-lan client statistics last update time (Live: IOS-XE 17.12.6a)
	GlanIdleUpdateTimestamp  time.Time `json:"glan-idle-update-timestamp"`  // Guest-lan client idle time last update (Live: IOS-XE 17.12.6a)
	RxGroupCounter           string    `json:"rx-group-counter"`            // Total broadcast and multicast frames sent (Live: IOS-XE 17.12.6a)
	TxTotalDrops             string    `json:"tx-total-drops"`              // Packets failed to transmit to client (Live: IOS-XE 17.12.6a)
}

// PolicyData represents client policy data.
type PolicyData struct {
	MAC         string `json:"mac"`           // Client MAC address (Live: IOS-XE 17.12.6a)
	ResVlanID   int    `json:"res-vlan-id"`   // Resolved VLAN ID (Live: IOS-XE 17.12.6a)
	ResVlanName string `json:"res-vlan-name"` // Resolved VLAN name (Live: IOS-XE 17.12.6a)
}

// SisfDBMac represents SISF database MAC information.
type SisfDBMac struct {
	MACAddr     string `json:"mac-addr"` // MAC address (YANG: IOS-XE 17.12.1)
	Ipv4Binding struct {
		IPKey struct {
			ZoneID int    `json:"zone-id"` // Zone identifier (YANG: IOS-XE 17.12.1)
			IPAddr string `json:"ip-addr"` // IPv4 address (YANG: IOS-XE 17.12.1)
		} `json:"ip-key"`
	} `json:"ipv4-binding"`
	Ipv6Binding []struct {
		Ipv6BindingIPKey struct {
			ZoneID int64  `json:"zone-id"` // Zone identifier (YANG: IOS-XE 17.12.1)
			IPAddr string `json:"ip-addr"` // IPv6 address (YANG: IOS-XE 17.12.1)
		} `json:"ip-key"`
	} `json:"ipv6-binding,omitempty"`
}

// DcInfo represents device classification information.
type DcInfo struct {
	ClientMAC        string    `json:"client-mac"`                   // Client MAC address (YANG: IOS-XE 17.12.1)
	DeviceType       string    `json:"device-type"`                  // Device type (YANG: IOS-XE 17.12.1)
	ProtocolMap      string    `json:"protocol-map"`                 // Protocol map (YANG: IOS-XE 17.12.1)
	ConfidenceLevel  int       `json:"confidence-level"`             // Confidence level (YANG: IOS-XE 17.12.1)
	ClassifiedTime   time.Time `json:"classified-time"`              // Classification time (YANG: IOS-XE 17.12.1)
	DayZeroDc        string    `json:"day-zero-dc"`                  // Day zero device classification (YANG: IOS-XE 17.12.1)
	SwVersionSrc     string    `json:"sw-version-src"`               // Software version source (YANG: IOS-XE 17.12.1)
	DeviceOs         string    `json:"device-os,omitempty"`          // Device operating system (YANG: IOS-XE 17.12.1)
	DeviceSubVersion string    `json:"device-sub-version,omitempty"` // Device sub-version (YANG: IOS-XE 17.12.1)
	DeviceOsSrc      string    `json:"device-os-src"`                // Device OS source (YANG: IOS-XE 17.12.1)
	DeviceName       string    `json:"device-name"`                  // Device name (YANG: IOS-XE 17.12.1)
	DeviceVendorSrc  string    `json:"device-vendor-src"`            // Device vendor source (YANG: IOS-XE 17.12.1)
	SalesCodeSrc     string    `json:"sales-code-src"`               // Sales code source (YANG: IOS-XE 17.12.1)
	DeviceSrc        string    `json:"device-src"`                   // Device source (YANG: IOS-XE 17.12.1)
	CountryNameSrc   string    `json:"country-name-src"`             // Country name source (YANG: IOS-XE 17.12.1)
	ModelNameSrc     string    `json:"model-name-src"`               // Model name source (YANG: IOS-XE 17.12.1)
	PowerTypeSrc     string    `json:"power-type-src"`               // Power type source (YANG: IOS-XE 17.12.1)
	HwModelSrc       string    `json:"hw-model-src"`                 // Hardware model source (YANG: IOS-XE 17.12.1)
	DeviceVendor     string    `json:"device-vendor,omitempty"`      // Device vendor (YANG: IOS-XE 17.12.1)
}

// ClientWlanPolicy represents client WLAN policy configuration.
type ClientWlanPolicy struct {
	CurrentSwitchingMode  string `json:"current-switching-mode"` // Current traffic switching mode for client (Live: IOS-XE 17.12.6a)
	WlanSwitchingMode     string `json:"wlan-switching-mode"`    // WLAN traffic switching mode based on properties (Live: IOS-XE 17.12.6a)
	CentralAuthentication string `json:"central-authentication"` // Central or local authentication mode (Live: IOS-XE 17.12.6a)
	CentralDHCP           bool   `json:"central-dhcp"`           // Central or local DHCP mode (Live: IOS-XE 17.12.6a)
	CentralAssocEnable    bool   `json:"central-assoc-enable"`   // Central or local association mode (Live: IOS-XE 17.12.6a)
	VlanCentralSwitching  bool   `json:"vlan-central-switching"` // VLAN-based central switching mode (Live: IOS-XE 17.12.6a)
	IsFabricClient        bool   `json:"is-fabric-client"`       // Client associated to Fabric WLAN (Live: IOS-XE 17.12.6a)
	IsGuestFabricClient   bool   `json:"is-guest-fabric-client"` // Client associated to Guest Fabric WLAN (Live: IOS-XE 17.12.6a)
	UpnBitFlag            string `json:"upn-bit-flag"`           // User Defined Network enabled status (Live: IOS-XE 17.12.6a)
}

// ClientGuestInfo represents guest LAN client information.
type ClientGuestInfo struct {
	WiredVlan       int `json:"wired-vlan"`        // VLAN client joins and is learned on foreign controller (Live: IOS-XE 17.12.6a)
	PhyIfid         int `json:"phy-ifid"`          // Physical interface ID client joins on foreign controller (Live: IOS-XE 17.12.6a)
	IdleTimeSeconds int `json:"idle-time-seconds"` // Idle time for guest-lan clients in seconds (Live: IOS-XE 17.12.6a)
}
