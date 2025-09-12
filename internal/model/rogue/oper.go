// Package rogue provides data models for rogue operational data.
package rogue

// RogueOper represents rogue operational data container from WNC 17.12.5.
type RogueOper struct {
	CiscoIOSXEWirelessRogueOperData struct {
		RogueStats      RogueStatsData      `json:"rogue-stats"`          // Rogue detection statistics
		RogueData       []RogueDataDetail   `json:"rogue-data"`           // Rogue access point details
		RogueClientData []RogueClientDetail `json:"rogue-client-data"`    // Rogue client details
		RLDPStats       *RLDPStats          `json:"rldp-stats,omitempty"` // RLDP statistics
	} `json:"Cisco-IOS-XE-wireless-rogue-oper:rogue-oper-data"` // Rogue operational data container // Rogue operational data container
}

// RogueData represents rogue access point detection data collection.
type RogueData struct {
	RogueData []RogueDataDetail `json:"Cisco-IOS-XE-wireless-rogue-oper:rogue-data"` // Rogue access point data
}

// RogueClientData represents rogue client detection data collection.
type RogueClientData struct {
	RogueClientData []RogueClientDetail `json:"Cisco-IOS-XE-wireless-rogue-oper:rogue-client-data"` // Rogue client data
}

// RogueStatsData represents rogue detection and classification statistics.
type RogueStatsData struct {
	RestartCount                int    `json:"restart-count"`                  // Rogue engine restart count
	PendingCount                int    `json:"pending-count"`                  // Pending rogue device count
	LradCount                   int    `json:"lrad-count"`                     // LRAD detection count
	OnMyNetworkCount            int    `json:"on-my-network-count"`            // Rogue devices on managed network
	AdhocCount                  int    `json:"adhoc-count"`                    // Ad-hoc network count
	UnknownCount                int    `json:"unknown-count"`                  // Unknown classification count
	UnclassifiedCount           int    `json:"unclassified-count"`             // Unclassified rogue count
	MaliciousCount              int    `json:"malicious-count"`                // Malicious rogue count
	FriendlyCount               int    `json:"friendly-count"`                 // Friendly rogue count
	CustomCount                 int    `json:"custom-count"`                   // Custom classification count
	NotAdhocCount               int    `json:"not-adhoc-count"`                // Non-adhoc rogue count
	TotalCount                  int    `json:"total-count"`                    // Total rogue device count
	ContainedCount              int    `json:"contained-count"`                // Contained rogue device count
	ContainedClientCount        int    `json:"contained-client-count"`         // Contained rogue client count
	ContainedPendingCount       int    `json:"contained-pending-count"`        // Contained pending count
	ContainedPendingClientCount int    `json:"contained-pending-client-count"` // Contained pending client count
	TotalClientCount            int    `json:"total-client-count"`             // Total rogue client count
	MaxCount                    int    `json:"max-count"`                      // Maximum rogue device limit
	MaxClientCount              int    `json:"max-client-count"`               // Maximum rogue client limit
	ReportCount                 string `json:"report-count"`                   // Rogue report count
	ClientReportCount           string `json:"client-report-count"`            // Rogue client report count
	RateReportCount             int    `json:"rate-report-count"`              // Report rate count
	RateClientReportCount       int    `json:"rate-client-report-count"`       // Client report rate count
	IappApPkt                   string `json:"iapp-ap-pkt"`                    // IAPP AP packet count
	IappClientPkt               string `json:"iapp-client-pkt"`                // IAPP client packet count
	RateIappApPkt               int    `json:"rate-iapp-ap-pkt"`               // IAPP AP packet rate
	RateIappClientPkt           int    `json:"rate-iapp-client-pkt"`           // IAPP client packet rate
	RldpCount                   string `json:"rldp-count"`                     // RLDP operation count
	AaaMsgRxCount               string `json:"aaa-msg-rx-count"`               // AAA message receive count
	AaaMsgTxCount               string `json:"aaa-msg-tx-count"`               // AAA message transmit count
	SnmpTrapsTxCount            string `json:"snmp-traps-tx-count"`            // SNMP trap transmit count
	LradOffCount                string `json:"lrad-off-count"`                 // LRAD offline count
	ApCreateCount               string `json:"ap-create-count"`                // AP creation count
	ApDeleteCount               string `json:"ap-delete-count"`                // AP deletion count
	ApRadioUpCount              string `json:"ap-radio-up-count"`              // AP radio up count
	ApRadioDownCount            string `json:"ap-radio-down-count"`            // AP radio down count
	ApNameChangeCount           int    `json:"ap-name-change-count"`           // AP name change count
	WncdIpcTxCount              string `json:"wncd-ipc-tx-count"`              // WNCD IPC transmit count
	WncdIpcRxCount              string `json:"wncd-ipc-rx-count"`              // WNCD IPC receive count
	WncmgrIpcRxCount            string `json:"wncmgr-ipc-rx-count"`            // WNC manager IPC receive count
	IosIpcTxCount               string `json:"ios-ipc-tx-count"`               // IOS IPC transmit count
	IosIpcRxCount               string `json:"ios-ipc-rx-count"`               // IOS IPC receive count
	NmspdIpcTxCount             string `json:"nmspd-ipc-tx-count"`             // NMSPD IPC transmit count
	NmspdIpcRxCount             string `json:"nmspd-ipc-rx-count"`             // NMSPD IPC receive count
	ContainMsgCount             string `json:"contain-msg-count"`              // Containment message count
	FsmErrors                   int    `json:"fsm-errors"`                     // Finite state machine error count
	TrapErrors                  int    `json:"trap-errors"`                    // SNMP trap error count

	EnqCount struct {
		Counters []struct {
			Value       string `json:"value"`       // Counter value
			Description string `json:"description"` // Counter description
		} `json:"counters"` // Counter entries
	} `json:"enq-count"` // Enqueue counter statistics

	SimilarApReportCount     string `json:"similar-ap-report-count"`     // Similar AP report count
	SimilarClientReportCount string `json:"similar-client-report-count"` // Similar client report count

	SnmpTrapsPerType struct {
		Counters []struct {
			Value       string `json:"value"`       // Trap count value
			Description string `json:"description"` // Trap type description
		} `json:"counters"` // Trap counter entries
	} `json:"snmp-traps-per-type"` // SNMP trap statistics by type

	IappReportMessagesLoadShedded int `json:"iapp-report-messages-load-shedded"` // IAPP report messages dropped due to load
	ManagedClientMessageCount     int `json:"managed-client-message-count"`      // Managed client message count
	ManagedClientJoinCount        int `json:"managed-client-join-count"`         // Managed client join count
	ManagedClientLeaveCount       int `json:"managed-client-leave-count"`        // Managed client leave count
	ManagedRogueClientCount       int `json:"managed-rogue-client-count"`        // Managed rogue client count

	ProcTime struct {
		Counters []struct {
			Value       string `json:"value"`       // Processing time value
			Description string `json:"description"` // Processing description
		} `json:"counters"` // Processing time entries
	} `json:"proc-time"` // Processing time statistics

	GlobalHistory struct {
		EventHistory []struct {
			Event      int    `json:"event"`       // Event identifier
			State      int    `json:"state"`       // State identifier
			Context    int    `json:"context"`     // Context identifier
			ContextStr string `json:"context-str"` // Context description
			CurrentRc  int    `json:"current-rc"`  // Current return code
			Count      int    `json:"count"`       // Event count
			Sticky     bool   `json:"sticky"`      // Sticky flag
			Timestamp  string `json:"timestamp"`   // Event timestamp
		} `json:"event-history"` // Event history entries
	} `json:"global-history"` // Global event history

	TblApfVapCacheReloadCount     int    `json:"tbl-apf-vap-cache-reload-count"`    // APF VAP cache reload count
	NewLradCount                  string `json:"new-lrad-count"`                    // New LRAD detection count
	LradPurgeCount                string `json:"lrad-purge-count"`                  // LRAD purge count
	RssiChangeCount               string `json:"rssi-change-count"`                 // RSSI change count
	FinalStateChangeCount         string `json:"final-state-change-count"`          // Final state change count
	ContainLevelChangeCount       string `json:"contain-level-change-count"`        // Containment level change count
	ClassChangeCount              string `json:"class-change-count"`                // Classification change count
	AdhocChangeCount              string `json:"adhoc-change-count"`                // Ad-hoc change count
	OnMyNetworkChangeCount        string `json:"on-my-network-change-count"`        // On-network change count
	NClientsChangedCount          string `json:"n-clients-changed-count"`           // Client count change
	ClientNewLradCount            string `json:"client-new-lrad-count"`             // Client new LRAD count
	ClientLradPurgeCount          string `json:"client-lrad-purge-count"`           // Client LRAD purge count
	ClientRssiChangeCount         string `json:"client-rssi-change-count"`          // Client RSSI change count
	ClientFinalStateChangeCount   string `json:"client-final-state-change-count"`   // Client state change count
	ClientContainLevelChangeCount string `json:"client-contain-level-change-count"` // Client containment change count
	ClientChannelChangeCount      string `json:"client-channel-change-count"`       // Client channel change count
	ClientIPChangeCount           string `json:"client-ip-change-count"`            // Client IP change count
	ClientRoamCount               string `json:"client-roam-count"`                 // Client roaming count

	RogueApReportsDroppedScale        string `json:"rogue-ap-reports-dropped-scale"`         // Rogue AP reports dropped scale
	RogueClientReportsDroppedScale    string `json:"rogue-client-reports-dropped-scale"`     // Rogue client reports dropped scale
	RogueClientReportsDroppedNoParent string `json:"rogue-client-reports-dropped-no-parent"` // Client reports dropped without parent

	RogueEnabled bool   `json:"rogue-enabled"`   // Rogue detection enabled flag
	MmIpcRxCount string `json:"mm-ipc-rx-count"` // MM IPC receive count

	RogueWsaEventsTriggeredCounter string `json:"rogue-wsa-events-triggered-counter"` // WSA events triggered counter
	RogueWsaEventsEnqueuedCounter  string `json:"rogue-wsa-events-enqueued-counter"`  // WSA events enqueued counter

	RogueWsaEventsTriggeredPerType struct {
		Counters []struct {
			Value       string `json:"value"`       // Event count value
			Description string `json:"description"` // Event type description
		} `json:"counters"` // WSA event counter entries
	} `json:"rogue-wsa-events-triggered-per-type"` // WSA events triggered by type

	RogueWsaEventsEnqueuedPerType struct {
		Counters []struct {
			Value       string `json:"value"`       // Enqueued count value
			Description string `json:"description"` // Event type description
		} `json:"counters"` // WSA enqueued counter entries
	} `json:"rogue-wsa-events-enqueued-per-type"` // WSA events enqueued by type

	BssidIpcCount        string `json:"bssid-ipc-count"`         // BSSID IPC count
	ApChannelChangeCount string `json:"ap-channel-change-count"` // AP channel change count
	BeaconDsAttackCount  string `json:"beacon-ds-attack-count"`  // Beacon DS attack count

	InternalCount int `json:"internal-count"` // Internal rogue count
	ExternalCount int `json:"external-count"` // External rogue count
	AlertCount    int `json:"alert-count"`    // Alert count
	ThreatCount   int `json:"threat-count"`   // Threat count

	RogueApReportFalseDrop string `json:"rogue-ap-report-false-drop"` // False drop report count

	AdhocUnknownCount      int `json:"adhoc-unknown-count"`      // Ad-hoc unknown count
	AdhocUnclassifiedCount int `json:"adhoc-unclassified-count"` // Ad-hoc unclassified count
	AdhocMaliciousCount    int `json:"adhoc-malicious-count"`    // Ad-hoc malicious count
	AdhocFriendlyCount     int `json:"adhoc-friendly-count"`     // Ad-hoc friendly count
	AdhocCustomCount       int `json:"adhoc-custom-count"`       // Ad-hoc custom count

	MaliciousInitCount    int `json:"malicious-init-count"`    // Malicious initialization count
	CustomInitCount       int `json:"custom-init-count"`       // Custom initialization count
	UnclassifiedInitCount int `json:"unclassified-init-count"` // Unclassified initialization count
	FriendlyInitCount     int `json:"friendly-init-count"`     // Friendly initialization count
	UnknownInitCount      int `json:"unknown-init-count"`      // Unknown initialization count
	InitCount             int `json:"init-count"`              // Total initialization count
	PostInitCount         int `json:"post-init-count"`         // Post initialization count
	MaxInitCount          int `json:"max-init-count"`          // Maximum initialization count

	TotalMaliciousCount    int `json:"total-malicious-count"`    // Total malicious count
	TotalCustomCount       int `json:"total-custom-count"`       // Total custom count
	TotalUnclassifiedCount int `json:"total-unclassified-count"` // Total unclassified count
	TotalFriendlyCount     int `json:"total-friendly-count"`     // Total friendly count
	TotalUnknownCount      int `json:"total-unknown-count"`      // Total unknown count
}

// RogueStats is a type alias for RogueStatsData to maintain backward compatibility with service layer.
type RogueStats = RogueStatsData

// RLDPStats represents Rogue Location Discovery Protocol statistics.
type RLDPStats struct {
	NumInProgress     int  `json:"num-in-progress"`     // Number of RLDP operations in progress
	NumRLDPStarted    int  `json:"num-rldp-started"`    // Number of RLDP operations started
	AuthTimeout       int  `json:"auth-timeout"`        // Authentication timeout count
	AssocTimeout      int  `json:"assoc-timeout"`       // Association timeout count
	DHCPTimeout       int  `json:"dhcp-timeout"`        // DHCP timeout count
	NotConnected      int  `json:"not-connected"`       // Not connected count
	Connected         int  `json:"connected"`           // Connected count
	RLDPSocketEnabled bool `json:"rldp-socket-enabled"` // RLDP socket enabled flag
}

// RogueDataDetail represents detailed information about detected rogue access points.
type RogueDataDetail struct {
	RogueAddress           string `json:"rogue-address"`              // Rogue device MAC address
	RogueClassType         string `json:"rogue-class-type"`           // Rogue classification type
	RogueMode              string `json:"rogue-mode"`                 // Rogue detection mode
	RogueContainmentLevel  int    `json:"rogue-containment-level"`    // Containment level setting
	ActualContainment      int    `json:"actual-containment"`         // Actual containment status
	ManualContained        bool   `json:"manual-contained"`           // Manual containment flag
	ClassOverrideSrc       string `json:"class-override-src"`         // Classification override source
	ContainmentType        string `json:"containment-type"`           // Type of containment applied
	Contained              bool   `json:"contained"`                  // Device contained flag
	SeverityScore          int    `json:"severity-score"`             // Security severity score
	ClassTypeCustomName    string `json:"class-type-custom-name"`     // Custom classification name
	RogueFirstTimestamp    string `json:"rogue-first-timestamp"`      // First detection timestamp
	RogueLastTimestamp     string `json:"rogue-last-timestamp"`       // Last detection timestamp
	RogueIsOnMyNetwork     bool   `json:"rogue-is-on-my-network"`     // Device is on managed network
	AdHoc                  bool   `json:"ad-hoc"`                     // Ad-hoc network flag
	AdHocBssid             string `json:"ad-hoc-bssid"`               // Ad-hoc BSSID identifier
	RogueRuleName          string `json:"rogue-rule-name"`            // Applied rogue rule name
	RogueRadioTypeLastSeen string `json:"rogue-radio-type-last-seen"` // Last seen radio type
	RldpRetries            int    `json:"rldp-retries"`               // RLDP retry count
	RogueClassTypeChange   string `json:"rogue-class-type-change"`    // Classification change history
	RogueStateChange       string `json:"rogue-state-change"`         // State change history
	RogueIfNum             int    `json:"rogue-if-num"`               // Interface number
	ManagedAp              bool   `json:"managed-ap"`                 // Managed AP flag
	AutocontainAdhocTrap   bool   `json:"autocontain-adhoc-trap"`     // Auto-contain ad-hoc trap
	AutocontainTrap        bool   `json:"autocontain-trap"`           // Auto-contain trap flag
	PotentialHoneypotTrap  bool   `json:"potential-honeypot-trap"`    // Potential honeypot trap

	History struct {
		EventHistory []struct {
			Event      int    `json:"event"`       // Event identifier
			State      int    `json:"state"`       // State identifier
			Context    int    `json:"context"`     // Context identifier
			ContextStr string `json:"context-str"` // Context description
			CurrentRc  int    `json:"current-rc"`  // Current return code
			Count      int    `json:"count"`       // Event occurrence count
			Sticky     bool   `json:"sticky"`      // Sticky event flag
			Timestamp  string `json:"timestamp"`   // Event timestamp
		} `json:"event-history"` // Event history entries
	} `json:"history"` // Rogue device event history

	RldpLastResult              string `json:"rldp-last-result"`                          // Last RLDP operation result
	RldpInProgress              bool   `json:"rldp-in-progress"`                          // RLDP operation in progress
	MaxDetectedRssi             int    `json:"max-detected-rssi"`                         // Maximum detected RSSI value
	SsidMaxRssi                 string `json:"ssid-max-rssi"`                             // SSID with maximum RSSI
	ApNameMaxRssi               string `json:"ap-name-max-rssi"`                          // AP name with maximum RSSI
	DetectingRadioType80211n25g []any  `json:"detecting-radio-type-80211n-24g,omitempty"` // 802.11n 2.4GHz radio detection
	DetectingRadioType80211g    []any  `json:"detecting-radio-type-80211g,omitempty"`     // 802.11g radio detection
	DetectingRadioType80211bg   []any  `json:"detecting-radio-type-802-11bg,omitempty"`   // 802.11b/g radio detection
	DRadioType80211ax24g        []any  `json:"d-radio-type-802-11ax24g,omitempty"`        // 802.11ax 2.4GHz radio detection
	DRadioType80211ax5g         []any  `json:"d-radio-type-802-11ax5g,omitempty"`         // 802.11ax 5GHz radio detection
	LradMacMaxRssi              string `json:"lrad-mac-max-rssi"`                         // LRAD MAC with maximum RSSI
	RogueRadioTypeMaxRssi       string `json:"rogue-radio-type-max-rssi"`                 // Radio type with maximum RSSI
	LastChannel                 int    `json:"last-channel"`                              // Last detected channel
	RadioTypeCount              []int  `json:"radio-type-count"`                          // Radio type detection count

	LastHeardLradKey struct {
		LradMacAddr string `json:"lrad-mac-addr"` // LRAD MAC address
		SlotID      int    `json:"slot-id"`       // Slot identifier
	} `json:"last-heard-lrad-key"` // Last heard LRAD key

	NLrads int `json:"n-lrads"` // Number of LRADs

	RogueLrad []struct {
		LradMacAddr string `json:"lrad-mac-addr"` // LRAD MAC address
		SlotID      int    `json:"slot-id"`       // LRAD slot identifier
		Ssid        string `json:"ssid"`          // Network SSID
		HiddenSsid  bool   `json:"hidden-ssid"`   // Hidden SSID flag
		Name        string `json:"name"`          // LRAD name
		Rssi        struct {
			Val int `json:"val"` // RSSI value
			Num int `json:"num"` // RSSI numerator
			Den int `json:"den"` // RSSI denominator
		} `json:"rssi"` // Received signal strength
		Snr struct {
			Val int `json:"val"` // SNR value
			Num int `json:"num"` // SNR numerator
			Den int `json:"den"` // SNR denominator
		} `json:"snr"` // Signal-to-noise ratio
		ShortPreamble           int    `json:"short-preamble"`            // Short preamble setting
		Channel                 int    `json:"channel"`                   // Operating channel
		Channels                []int  `json:"channels"`                  // Available channels
		Encrypted               int    `json:"encrypted"`                 // Encryption status
		WpaSupport              int    `json:"wpa-support"`               // WPA support indicator
		Dot11PhySupport         int    `json:"dot11-phy-support"`         // 802.11 PHY support
		LastHeard               string `json:"last-heard"`                // Last detection time
		ChanWidth               int    `json:"chan-width"`                // Channel width
		ChanWidthLabel          int    `json:"chan-width-label"`          // Channel width label
		ExtChan                 int    `json:"ext-chan"`                  // Extension channel
		BandID                  int    `json:"band-id"`                   // Band identifier
		NumSlots                int    `json:"num-slots"`                 // Number of slots
		ReportRadioType         int    `json:"report-radio-type"`         // Reported radio type
		ContainResult           string `json:"contain-result"`            // Containment result
		ContainSlotID           string `json:"contain-slot-id"`           // Containment slot ID
		ContainRadioType        int    `json:"contain-radio-type"`        // Containment radio type
		RadioType               string `json:"radio-type"`                // Radio type identifier
		ContainmentType         string `json:"containment-type"`          // Containment type
		ContainmentChannelCount int    `json:"containment-channel-count"` // Containment channel count
		RogueContainmentChans   string `json:"rogue-containment-chans"`   // Rogue containment channels
		AuthFailCount           int    `json:"auth-fail-count"`           // Authentication failure count
		MfpStatus               string `json:"mfp-status"`                // Management frame protection status
		ChannelFromDS           bool   `json:"channel-from-ds"`           // Channel from distribution system
		PhyApSlot               int    `json:"phy-ap-slot"`               // Physical AP slot
	} `json:"rogue-lrad"` // LRAD detection details

	NClients int `json:"n-clients"` // Number of rogue clients

	RogueClient *[]struct {
		RogueClientAddress string `json:"rogue-client-address"` // Rogue client MAC address
	} `json:"rogue-client,omitempty"` // Associated rogue clients

	RemoteOverride struct {
		RemoteOverrideClassType        string `json:"remote-override-class-type"`        // Remote override classification
		RemoteOverrideMode             string `json:"remote-override-mode"`              // Remote override mode
		RemoteOverrideContainmentLevel int    `json:"remote-override-containment-level"` // Remote override containment level
	} `json:"remote-override"` // Remote override settings

	LastHeardSsid     string `json:"last-heard-ssid"`      // Last heard SSID
	MfpRequired       bool   `json:"mfp-required"`         // Management frame protection required
	ChannelMaxRssi    int    `json:"channel-max-rssi"`     // Channel with maximum RSSI
	WpaSupportMaxRssi string `json:"wpa-support-max-rssi"` // WPA support with maximum RSSI
	EncryptedMaxRssi  string `json:"encrypted-max-rssi"`   // Encryption status with maximum RSSI
	SnrMaxRssi        int    `json:"snr-max-rssi"`         // SNR with maximum RSSI
	Properties        string `json:"properties"`           // Device properties

	BandData2dot4g struct {
		ChanWidth         int `json:"chan-width"`         // Channel width for 2.4GHz
		ActualContainment int `json:"actual-containment"` // Actual containment for 2.4GHz
	} `json:"band-data-2dot4g"` // 2.4GHz band data

	BandData5g struct {
		ChanWidth         int `json:"chan-width"`         // Channel width for 5GHz
		ActualContainment int `json:"actual-containment"` // Actual containment for 5GHz
	} `json:"band-data-5g"` // 5GHz band data

	BandData6g struct {
		ChanWidth         int `json:"chan-width"`         // Channel width for 6GHz
		ActualContainment int `json:"actual-containment"` // Actual containment for 6GHz
	} `json:"band-data-6g"` // 6GHz band data
}

// RogueClientDetail represents detailed information about detected rogue clients.
type RogueClientDetail struct {
	RogueClientAddress          string `json:"rogue-client-address"`           // Rogue client MAC address
	RogueClientBssid            string `json:"rogue-client-bssid"`             // Associated BSSID
	RogueClientGateway          string `json:"rogue-client-gateway"`           // Gateway address
	RogueClientState            string `json:"rogue-client-state"`             // Client state
	RogueClientContainmentLevel int    `json:"rogue-client-containment-level"` // Containment level
	ActualContainment           int    `json:"actual-containment"`             // Actual containment status
	ContainmentType             string `json:"containment-type"`               // Containment type
	RogueClientIfNum            int    `json:"rogue-client-if-num"`            // Interface number
	RogueClientIPv4Addr         string `json:"rogue-client-ipv4-addr"`         // IPv4 address
	RogueClientIPv6Addr         string `json:"rogue-client-ipv6-addr"`         // IPv6 address
	ManualContained             bool   `json:"manual-contained"`               // Manual containment flag
	Contained                   bool   `json:"contained"`                      // Client contained flag
	AaaCheck                    bool   `json:"aaa-check"`                      // AAA check performed
	CmxCheck                    bool   `json:"cmx-check"`                      // CMX check performed
	RogueClientFirstTimestamp   string `json:"rogue-client-first-timestamp"`   // First detection timestamp
	RogueClientLastTimestamp    string `json:"rogue-client-last-timestamp"`    // Last detection timestamp

	LastHeardLradKey struct {
		LradMacAddr string `json:"lrad-mac-addr"` // LRAD MAC address
		SlotID      int    `json:"slot-id"`       // Slot identifier
	} `json:"last-heard-lrad-key"` // Last heard LRAD key information

	History struct {
		EventHistory []struct {
			Event      int    `json:"event"`       // Event identifier
			State      int    `json:"state"`       // State identifier
			Context    int    `json:"context"`     // Context identifier
			ContextStr string `json:"context-str"` // Context description
			CurrentRc  int    `json:"current-rc"`  // Current return code
			Count      int    `json:"count"`       // Event occurrence count
			Sticky     bool   `json:"sticky"`      // Sticky event flag
			Timestamp  string `json:"timestamp"`   // Event timestamp
		} `json:"event-history"` // Event history entries
	} `json:"history"` // Rogue client event history

	ParentRogueDataAddress string `json:"parent-rogue-data-address"` // Parent rogue device address

	RogueClientLrad []struct {
		LradMacAddr string `json:"lrad-mac-addr"` // LRAD MAC address
		SlotID      int    `json:"slot-id"`       // Slot identifier
		LastHeard   string `json:"last-heard"`    // Last detection time
		Name        string `json:"name"`          // LRAD name
		Rssi        struct {
			Val int `json:"val"` // RSSI value
			Num int `json:"num"` // RSSI numerator
			Den int `json:"den"` // RSSI denominator
		} `json:"rssi"` // Received signal strength
		Snr struct {
			Val int `json:"val"` // SNR value
			Num int `json:"num"` // SNR numerator
			Den int `json:"den"` // SNR denominator
		} `json:"snr"` // Signal-to-noise ratio
		Channel          int    `json:"channel"`            // Operating channel
		BandID           int    `json:"band-id"`            // Band identifier
		NumSlots         int    `json:"num-slots"`          // Number of slots
		ContainSlotID    string `json:"contain-slot-id"`    // Containment slot ID
		ContainRadioType int    `json:"contain-radio-type"` // Containment radio type
		ContainResult    string `json:"contain-result"`     // Containment result
		PhyApSlot        int    `json:"phy-ap-slot"`        // Physical AP slot
	} `json:"rogue-client-lrad"` // LRAD detection details for client

	BandData2dot4g struct {
		NLrads            int `json:"n-lrads"`            // Number of LRADs for 2.4GHz
		ActualContainment int `json:"actual-containment"` // Actual containment for 2.4GHz
	} `json:"band-data-2dot4g"` // 2.4GHz band data for client

	BandData5g struct {
		NLrads            int `json:"n-lrads"`            // Number of LRADs for 5GHz
		ActualContainment int `json:"actual-containment"` // Actual containment for 5GHz
	} `json:"band-data-5g"` // 5GHz band data for client

	BandData6g struct {
		NLrads            int `json:"n-lrads"`            // Number of LRADs for 6GHz
		ActualContainment int `json:"actual-containment"` // Actual containment for 6GHz
	} `json:"band-data-6g"` // 6GHz band data for client
}
