package rogue

// CiscoIOSXEWirelessRogueOper represents rogue operational data container.
type CiscoIOSXEWirelessRogueOper struct {
	CiscoIOSXEWirelessRogueOperData struct {
		RogueStats      RogueStatsData      `json:"rogue-stats"`          // Rogue detection statistics (Live: IOS-XE 17.12.6a)
		RogueData       []RogueDataDetail   `json:"rogue-data"`           // Rogue access point details (Live: IOS-XE 17.12.6a)
		RogueClientData []RogueClientDetail `json:"rogue-client-data"`    // Rogue client details (Live: IOS-XE 17.12.6a)
		RLDPStats       *RLDPStats          `json:"rldp-stats,omitempty"` // RLDP statistics (Live: IOS-XE 17.12.6a)
	} `json:"Cisco-IOS-XE-wireless-rogue-oper:rogue-oper-data"` // Rogue operational data container (Live: IOS-XE 17.12.6a)
}

// CiscoIOSXEWirelessRogueData represents rogue access point detection data collection.
type CiscoIOSXEWirelessRogueData struct {
	RogueData []RogueDataDetail `json:"Cisco-IOS-XE-wireless-rogue-oper:rogue-data"` // Rogue access point data (Live: IOS-XE 17.12.6a)
}

// CiscoIOSXEWirelessRogueClientData represents rogue client detection data collection.
type CiscoIOSXEWirelessRogueClientData struct {
	RogueClientData []RogueClientDetail `json:"Cisco-IOS-XE-wireless-rogue-oper:rogue-client-data"` // Rogue client data (Live: IOS-XE 17.12.6a)
}

// RogueStatsData represents rogue detection and classification statistics.
type RogueStatsData struct {
	RestartCount                int    `json:"restart-count"`                  // Number of process restarts (Live: IOS-XE 17.12.6a)
	PendingCount                int    `json:"pending-count"`                  // Number of rogue AP in pending state (Live: IOS-XE 17.12.6a)
	LradCount                   int    `json:"lrad-count"`                     // Number of rogue AP in LRAD state (Live: IOS-XE 17.12.6a)
	OnMyNetworkCount            int    `json:"on-my-network-count"`            // Number of rogue AP in my network (Live: IOS-XE 17.12.6a)
	AdhocCount                  int    `json:"adhoc-count"`                    // Number of ad-hoc rogue APs (Live: IOS-XE 17.12.6a)
	UnknownCount                int    `json:"unknown-count"`                  // Number of unknown rogue APs (Live: IOS-XE 17.12.6a)
	UnclassifiedCount           int    `json:"unclassified-count"`             // Number of unclassified rogue APs (Live: IOS-XE 17.12.6a)
	MaliciousCount              int    `json:"malicious-count"`                // Number of malicious rogue APs (Live: IOS-XE 17.12.6a)
	FriendlyCount               int    `json:"friendly-count"`                 // Number of friendly rogue APs (Live: IOS-XE 17.12.6a)
	CustomCount                 int    `json:"custom-count"`                   // Number of custom rogue APs (Live: IOS-XE 17.12.6a)
	NotAdhocCount               int    `json:"not-adhoc-count"`                // Number of rogue APs (not adhoc) (Live: IOS-XE 17.12.6a)
	TotalCount                  int    `json:"total-count"`                    // Number of rogue APs in total (Live: IOS-XE 17.12.6a)
	ContainedCount              int    `json:"contained-count"`                // Number of contained rogue APs (Live: IOS-XE 17.12.6a)
	ContainedClientCount        int    `json:"contained-client-count"`         // Number of contained rogue clients (Live: IOS-XE 17.12.6a)
	ContainedPendingCount       int    `json:"contained-pending-count"`        // Number of containment-pending rogue APs (Live: IOS-XE 17.12.6a)
	ContainedPendingClientCount int    `json:"contained-pending-client-count"` // Number of containment-pending rogue clients (Live: IOS-XE 17.12.6a)
	TotalClientCount            int    `json:"total-client-count"`             // Number of rogue clients in total (Live: IOS-XE 17.12.6a)
	MaxCount                    int    `json:"max-count"`                      // Number of rogue APs that system can support (Live: IOS-XE 17.12.6a)
	MaxClientCount              int    `json:"max-client-count"`               // Number of rogue clients that system can support (Live: IOS-XE 17.12.6a)
	ReportCount                 string `json:"report-count"`                   // Number of IAPP AP reports received (Live: IOS-XE 17.12.6a)
	ClientReportCount           string `json:"client-report-count"`            // Number of IAPP Client reports received (Live: IOS-XE 17.12.6a)
	RateReportCount             int    `json:"rate-report-count"`              // Number of IAPP AP reports received in last minute (Live: IOS-XE 17.12.6a)
	RateClientReportCount       int    `json:"rate-client-report-count"`       // Number of IAPP Client reports received in last minute (Live: IOS-XE 17.12.6a)
	IappApPkt                   string `json:"iapp-ap-pkt"`                    // Number of IAPP AP packets received (Live: IOS-XE 17.12.6a)
	IappClientPkt               string `json:"iapp-client-pkt"`                // Number of IAPP Client packets received (Live: IOS-XE 17.12.6a)
	RateIappApPkt               int    `json:"rate-iapp-ap-pkt"`               // Number of IAPP AP packets received in last minute (Live: IOS-XE 17.12.6a)
	RateIappClientPkt           int    `json:"rate-iapp-client-pkt"`           // Number of IAPP Client packets received in last minute (Live: IOS-XE 17.12.6a)
	RldpCount                   string `json:"rldp-count"`                     // Number of RLDP procedure started (Live: IOS-XE 17.12.6a)
	AaaMsgRxCount               string `json:"aaa-msg-rx-count"`               // Number of AAA messages received (Live: IOS-XE 17.12.6a)
	AaaMsgTxCount               string `json:"aaa-msg-tx-count"`               // Number of AAA messages sent (Live: IOS-XE 17.12.6a)
	SnmpTrapsTxCount            string `json:"snmp-traps-tx-count"`            // Number of SNMP traps sent (Live: IOS-XE 17.12.6a)
	LradOffCount                string `json:"lrad-off-count"`                 // Number of LRAD off events (Live: IOS-XE 17.12.6a)
	ApCreateCount               string `json:"ap-create-count"`                // Number of AP create events (Live: IOS-XE 17.12.6a)
	ApDeleteCount               string `json:"ap-delete-count"`                // Number of AP delete events (Live: IOS-XE 17.12.6a)
	ApRadioUpCount              string `json:"ap-radio-up-count"`              // Number of AP Radio Up events (Live: IOS-XE 17.12.6a)
	ApRadioDownCount            string `json:"ap-radio-down-count"`            // Number of AP Radio Down events (Live: IOS-XE 17.12.6a)
	ApNameChangeCount           int    `json:"ap-name-change-count"`           // Number of AP Name Change events (Live: IOS-XE 17.12.6a)
	WncdIpcTxCount              string `json:"wncd-ipc-tx-count"`              // Number of IPCs to WNCDs sent (Live: IOS-XE 17.12.6a)
	WncdIpcRxCount              string `json:"wncd-ipc-rx-count"`              // Number of IPCs from WNCDs received (Live: IOS-XE 17.12.6a)
	WncmgrIpcRxCount            string `json:"wncmgr-ipc-rx-count"`            // Number of IPCs from WNCMGR received (Live: IOS-XE 17.12.6a)
	IosIpcTxCount               string `json:"ios-ipc-tx-count"`               // Number of IPCs to IOS sent (Live: IOS-XE 17.12.6a)
	IosIpcRxCount               string `json:"ios-ipc-rx-count"`               // Number of IPCs from IOS received (Live: IOS-XE 17.12.6a)
	NmspdIpcTxCount             string `json:"nmspd-ipc-tx-count"`             // Number of IPCs to NMSPD sent (Live: IOS-XE 17.12.6a)
	NmspdIpcRxCount             string `json:"nmspd-ipc-rx-count"`             // Number of IPCs from NMSPD received (Live: IOS-XE 17.12.6a)
	ContainMsgCount             string `json:"contain-msg-count"`              // Number of Containment msgs sent to APs (Live: IOS-XE 17.12.6a)
	FsmErrors                   int    `json:"fsm-errors"`                     // Number of FSM errors (Live: IOS-XE 17.12.6a)
	TrapErrors                  int    `json:"trap-errors"`                    // Number of TRAP errors (Live: IOS-XE 17.12.6a)

	EnqCount struct {
		Counters []struct {
			Value       string `json:"value"`       // Counter value (Live: IOS-XE 17.12.6a)
			Description string `json:"description"` // Counter description (Live: IOS-XE 17.12.6a)
		} `json:"counters"` // Counter entries (Live: IOS-XE 17.12.6a)
	} `json:"enq-count"` // Number of object enqueues (Live: IOS-XE 17.12.6a)

	SimilarApReportCount     string `json:"similar-ap-report-count"`     // Number of very-similar AP reports (Live: IOS-XE 17.12.6a)
	SimilarClientReportCount string `json:"similar-client-report-count"` // Number of very-similar client reports (Live: IOS-XE 17.12.6a)

	SnmpTrapsPerType struct {
		Counters []struct {
			Value       string `json:"value"`       // Trap count value (Live: IOS-XE 17.12.6a)
			Description string `json:"description"` // Trap type description (Live: IOS-XE 17.12.6a)
		} `json:"counters"` // Trap counter entries (Live: IOS-XE 17.12.6a)
	} `json:"snmp-traps-per-type"` // per-trap-type counter (Live: IOS-XE 17.12.6a)

	IappReportMessagesLoadShedded int `json:"iapp-report-messages-load-shedded"` // Number of IAPP report messages not processed due to rx IPC too-high occupancy (Live: IOS-XE 17.12.6a)
	ManagedClientMessageCount     int `json:"managed-client-message-count"`      // Number of client join/leave messages (Live: IOS-XE 17.12.6a)
	ManagedClientJoinCount        int `json:"managed-client-join-count"`         // number of client join events (Live: IOS-XE 17.12.6a)
	ManagedClientLeaveCount       int `json:"managed-client-leave-count"`        // Number of client leaving the run state events (Live: IOS-XE 17.12.6a)
	ManagedRogueClientCount       int `json:"managed-rogue-client-count"`        // Number of managed clients matching rogue clients (Live: IOS-XE 17.12.6a)

	ProcTime struct {
		Counters []struct {
			Value       string `json:"value"`       // Processing time value (Live: IOS-XE 17.12.6a)
			Description string `json:"description"` // Processing description (Live: IOS-XE 17.12.6a)
		} `json:"counters"` // Processing time entries (Live: IOS-XE 17.12.6a)
	} `json:"proc-time"` // per-processing-type average processing time (Live: IOS-XE 17.12.6a)

	GlobalHistory struct {
		EventHistory []struct {
			Event      int    `json:"event"`       // Event identifier (Live: IOS-XE 17.12.6a)
			State      int    `json:"state"`       // State identifier (Live: IOS-XE 17.12.6a)
			Context    int    `json:"context"`     // Context identifier (Live: IOS-XE 17.12.6a)
			ContextStr string `json:"context-str"` // Context description (Live: IOS-XE 17.12.6a)
			CurrentRc  int    `json:"current-rc"`  // Current return code (Live: IOS-XE 17.12.6a)
			Count      int    `json:"count"`       // Event count (Live: IOS-XE 17.12.6a)
			Sticky     bool   `json:"sticky"`      // Sticky flag (Live: IOS-XE 17.12.6a)
			Timestamp  string `json:"timestamp"`   // Event timestamp (Live: IOS-XE 17.12.6a)
		} `json:"event-history"` // Event history entries (Live: IOS-XE 17.12.6a)
	} `json:"global-history"` // Global Event History (Live: IOS-XE 17.12.6a)

	TblAPFVapCacheReloadCount     int    `json:"tbl-apf-vap-cache-reload-count"`    // Count of APF VAP SSID cache reloads (Live: IOS-XE 17.12.6a)
	NewLradCount                  string `json:"new-lrad-count"`                    // Number of times a new LRAD has been added (Live: IOS-XE 17.12.6a)
	LradPurgeCount                string `json:"lrad-purge-count"`                  // Number of LRAD purge events (Live: IOS-XE 17.12.6a)
	RSSIChangeCount               string `json:"rssi-change-count"`                 // Number of RSSI change events (Live: IOS-XE 17.12.6a)
	FinalStateChangeCount         string `json:"final-state-change-count"`          // Number of times the final state has changed (Live: IOS-XE 17.12.6a)
	ContainLevelChangeCount       string `json:"contain-level-change-count"`        // Number of times the containment level has changed (Live: IOS-XE 17.12.6a)
	ClassChangeCount              string `json:"class-change-count"`                // Number of Classification Type changes (Live: IOS-XE 17.12.6a)
	AdhocChangeCount              string `json:"adhoc-change-count"`                // Number of times adhoc status changed (Live: IOS-XE 17.12.6a)
	OnMyNetworkChangeCount        string `json:"on-my-network-change-count"`        // Number of times on-my-network status changed (Live: IOS-XE 17.12.6a)
	NClientsChangedCount          string `json:"n-clients-changed-count"`           // Number of times the client-number has changed (Live: IOS-XE 17.12.6a)
	ClientNewLradCount            string `json:"client-new-lrad-count"`             // Number of times a new client LRAD has been added (Live: IOS-XE 17.12.6a)
	ClientLradPurgeCount          string `json:"client-lrad-purge-count"`           // Number of client LRAD purge events (Live: IOS-XE 17.12.6a)
	ClientRSSIChangeCount         string `json:"client-rssi-change-count"`          // Number of client RSSI change events (Live: IOS-XE 17.12.6a)
	ClientFinalStateChangeCount   string `json:"client-final-state-change-count"`   // Number of times the final client state has changed (Live: IOS-XE 17.12.6a)
	ClientContainLevelChangeCount string `json:"client-contain-level-change-count"` // Number of times the client containment level has changed (Live: IOS-XE 17.12.6a)
	ClientChannelChangeCount      string `json:"client-channel-change-count"`       // Number of channel change events (Live: IOS-XE 17.12.6a)
	ClientIPChangeCount           string `json:"client-ip-change-count"`            // Number of IP-change events (Live: IOS-XE 17.12.6a)
	ClientRoamCount               string `json:"client-roam-count"`                 // Number of rogue-client-roam events (Live: IOS-XE 17.12.6a)

	RogueApReportsDroppedScale        string `json:"rogue-ap-reports-dropped-scale"`         // Number of rogue AP reports dropped due to max. scale reached (Live: IOS-XE 17.12.6a)
	RogueClientReportsDroppedScale    string `json:"rogue-client-reports-dropped-scale"`     // Number of rogue client reports dropped due to max. scale reached (Live: IOS-XE 17.12.6a)
	RogueClientReportsDroppedNoParent string `json:"rogue-client-reports-dropped-no-parent"` // Number of rogue client reports dropped due to missing parent rogue AP (Live: IOS-XE 17.12.6a)

	RogueEnabled bool   `json:"rogue-enabled"`   // Rogue socket on port 5247 is enabled (Live: IOS-XE 17.12.6a)
	MmIpcRxCount string `json:"mm-ipc-rx-count"` // Number of IPCs from Mobilityd received (Live: IOS-XE 17.12.6a)

	RogueWsaEventsTriggeredCounter string `json:"rogue-wsa-events-triggered-counter"` // Number of Rogue WSA events triggered (Live: IOS-XE 17.12.6a)
	RogueWsaEventsEnqueuedCounter  string `json:"rogue-wsa-events-enqueued-counter"`  // Number of Rogue WSA events enqueued (Live: IOS-XE 17.12.6a)

	RogueWsaEventsTriggeredPerType struct {
		Counters []struct {
			Value       string `json:"value"`       // Event count value
			Description string `json:"description"` // Event type description
		} `json:"counters"` // WSA event counter entries
	} `json:"rogue-wsa-events-triggered-per-type"` // Number of Rogue WSA events triggered per-type (Live: IOS-XE 17.12.6a)

	RogueWsaEventsEnqueuedPerType struct {
		Counters []struct {
			Value       string `json:"value"`       // Enqueued count value
			Description string `json:"description"` // Event type description
		} `json:"counters"` // WSA enqueued counter entries
	} `json:"rogue-wsa-events-enqueued-per-type"` // Number of Rogue WSA events enqueued per-type (Live: IOS-XE 17.12.6a)

	BssidIpcCount        string `json:"bssid-ipc-count"`         // Number of BSSID cache update events (Live: IOS-XE 17.12.6a)
	ApChannelChangeCount string `json:"ap-channel-change-count"` // Number of AP channel change events (Live: IOS-XE 17.12.6a)
	BeaconDsAttackCount  string `json:"beacon-ds-attack-count"`  // Number of Beacon DS attacks detected (Live: IOS-XE 17.12.6a)

	InternalCount int `json:"internal-count"` // Number of rogue APs in internal state (Live: IOS-XE 17.12.6a)
	ExternalCount int `json:"external-count"` // Number of rogue APs in external state (Live: IOS-XE 17.12.6a)
	AlertCount    int `json:"alert-count"`    // Number of rogue APs in alert state (Live: IOS-XE 17.12.6a)
	ThreatCount   int `json:"threat-count"`   // Number of rogue APs in threat state (Live: IOS-XE 17.12.6a)

	RogueApReportFalseDrop string `json:"rogue-ap-report-false-drop"` // Number of rogue AP reports dropped because they might be falsely reported neighbor APs (Live: IOS-XE 17.12.6a)

	AdhocUnknownCount      int `json:"adhoc-unknown-count"`      // Number of unknown ad-hoc rogue APs (Live: IOS-XE 17.12.6a)
	AdhocUnclassifiedCount int `json:"adhoc-unclassified-count"` // Number of unclassified ad-hoc rogue APs (Live: IOS-XE 17.12.6a)
	AdhocMaliciousCount    int `json:"adhoc-malicious-count"`    // Number of malicious ad-hoc rogue APs (Live: IOS-XE 17.12.6a)
	AdhocFriendlyCount     int `json:"adhoc-friendly-count"`     // Number of friendly ad-hoc rogue APs (Live: IOS-XE 17.12.6a)
	AdhocCustomCount       int `json:"adhoc-custom-count"`       // Number of custom ad-hoc rogue APs (Live: IOS-XE 17.12.6a)

	MaliciousInitCount    int `json:"malicious-init-count"`    // Number of malicious rogue APs in init state (Live: IOS-XE 17.12.6a)
	CustomInitCount       int `json:"custom-init-count"`       // Number of custom rogue APs in init state (Live: IOS-XE 17.12.6a)
	UnclassifiedInitCount int `json:"unclassified-init-count"` // Number of unclassified rogue APs in init state (Live: IOS-XE 17.12.6a)
	FriendlyInitCount     int `json:"friendly-init-count"`     // Number of friendly rogue APs in init state (Live: IOS-XE 17.12.6a)
	UnknownInitCount      int `json:"unknown-init-count"`      // Number of unknown rogue APs in init state (Live: IOS-XE 17.12.6a)
	InitCount             int `json:"init-count"`              // Total number of rogue APs in init state (Live: IOS-XE 17.12.6a)
	PostInitCount         int `json:"post-init-count"`         // Total number of rogue APs in post init state (Live: IOS-XE 17.12.6a)
	MaxInitCount          int `json:"max-init-count"`          // Maximum number of rogue APs in init state (Live: IOS-XE 17.12.6a)

	TotalMaliciousCount    int `json:"total-malicious-count"`    // Total number of malicious classification rogues (Live: IOS-XE 17.12.6a)
	TotalCustomCount       int `json:"total-custom-count"`       // Total number of custom classification rogues (Live: IOS-XE 17.12.6a)
	TotalUnclassifiedCount int `json:"total-unclassified-count"` // Total number of unclassified classification rogues (Live: IOS-XE 17.12.6a)
	TotalFriendlyCount     int `json:"total-friendly-count"`     // Total number of friendly classification rogues (Live: IOS-XE 17.12.6a)
	TotalUnknownCount      int `json:"total-unknown-count"`      // Total number of unknown classification rogues (Live: IOS-XE 17.12.6a)

	// Fields added in IOS-XE 17.18.1
	RogueApMldLinkCount     int    `json:"rogue-ap-mld-link-count"`   // Total number of Rogue AP backward compatible MLD-Link records (Live: IOS-XE 17.15.4b)
	RogueClientMldLinkCnt   int    `json:"rogue-client-mld-link-cnt"` // Total number of Rogue Client backward compatible MLD-Link records (Live: IOS-XE 17.15.4b)
	ApDropMldMismatch       string `json:"ap-drop-mld-mismatch"`      // Total number of rogue AP reports dropped due to MLD / Non-MLD type mismatch (Live: IOS-XE 17.15.4b)
	ClientDropMldMismatch   string `json:"client-drop-mld-mismatch"`  // Total number of rogue AP reports dropped due to MLD / Non-MLD type mismatch (Live: IOS-XE 17.15.4b)
	IappUnconnectedClient   uint64 `json:"iapp-unconnected-client"`   // Number of IAPP Unconnected Client packets received (Live: IOS-XE 17.15.4b)
	UnconnectedClientReport uint64 `json:"unconnected-client-report"` // Number of IAPP Client reports received (Live: IOS-XE 17.15.4b)
	UnconnectedClientCount  uint64 `json:"unconnected-client-count"`  // Number of unconnected client in total (Live: IOS-XE 17.15.4b)
	UnconnectedReportsDrop  uint64 `json:"unconnected-reports-drop"`  // Number of unconnected clients dropped due to max. scale reached (Live: IOS-XE 17.15.4b)
	ApDropURWBLink          uint64 `json:"ap-drop-urwb-link"`         // Total number of rogue AP reports dropped due to URWB link address reported as rogue AP (Live: IOS-XE 17.15.4b)
}

// RLDPStats represents Rogue Location Discovery Protocol statistics.
type RLDPStats struct {
	NumInProgress     int  `json:"num-in-progress"`     // Number of RLDP operations in progress (Live: IOS-XE 17.12.6a)
	NumRLDPStarted    int  `json:"num-rldp-started"`    // Number of RLDP operations started (Live: IOS-XE 17.12.6a)
	AuthTimeout       int  `json:"auth-timeout"`        // Authentication timeout count (Live: IOS-XE 17.12.6a)
	AssocTimeout      int  `json:"assoc-timeout"`       // Association timeout count (Live: IOS-XE 17.12.6a)
	DHCPTimeout       int  `json:"dhcp-timeout"`        // DHCP timeout count (Live: IOS-XE 17.12.6a)
	NotConnected      int  `json:"not-connected"`       // Not connected count (Live: IOS-XE 17.12.6a)
	Connected         int  `json:"connected"`           // Connected count (Live: IOS-XE 17.12.6a)
	RLDPSocketEnabled bool `json:"rldp-socket-enabled"` // RLDP socket enabled flag (Live: IOS-XE 17.12.6a)
}

// RogueDataDetail represents detailed information about detected rogue access points.
type RogueDataDetail struct {
	RogueAddress           string `json:"rogue-address"`              // MAC Address of a rogue AP (Live: IOS-XE 17.12.6a)
	RogueClassType         string `json:"rogue-class-type"`           // Type of a rogue AP (Live: IOS-XE 17.12.6a)
	RogueMode              string `json:"rogue-mode"`                 // State in which the rogue AP is (Live: IOS-XE 17.12.6a)
	RogueContainmentLevel  int    `json:"rogue-containment-level"`    // Containment Level (Live: IOS-XE 17.12.6a)
	ActualContainment      int    `json:"actual-containment"`         // Number of Containing APs (Live: IOS-XE 17.12.6a)
	ManualContained        bool   `json:"manual-contained"`           // Manually Contained (Live: IOS-XE 17.12.6a)
	ClassOverrideSrc       string `json:"class-override-src"`         // Source of classification/containment override (Live: IOS-XE 17.12.6a)
	ContainmentType        string `json:"containment-type"`           // Containment mode applied to this rogue AP (Live: IOS-XE 17.12.6a)
	Contained              bool   `json:"contained"`                  // Contained (Live: IOS-XE 17.12.6a)
	SeverityScore          int    `json:"severity-score"`             // Custom classification severity score (Live: IOS-XE 17.12.6a)
	ClassTypeCustomName    string `json:"class-type-custom-name"`     // Custom rule (Live: IOS-XE 17.12.6a)
	RogueFirstTimestamp    string `json:"rogue-first-timestamp"`      // Time when this Rogue was First Detected (Live: IOS-XE 17.12.6a)
	RogueLastTimestamp     string `json:"rogue-last-timestamp"`       // Time when this Rogue was Last Detected (Live: IOS-XE 17.12.6a)
	RogueIsOnMyNetwork     bool   `json:"rogue-is-on-my-network"`     // Specifies if the Rogue is on Wired Network (Live: IOS-XE 17.12.6a)
	AdHoc                  bool   `json:"ad-hoc"`                     // Specifies if the Rogue is ad-hoc type or AP (Live: IOS-XE 17.12.6a)
	AdHocBssid             string `json:"ad-hoc-bssid"`               // BSSID for Ad-Hoc Rogue (Live: IOS-XE 17.12.6a)
	RogueRuleName          string `json:"rogue-rule-name"`            // Rule Name (Live: IOS-XE 17.12.6a)
	RogueRadioTypeLastSeen string `json:"rogue-radio-type-last-seen"` // Last Seen Radio Type (Live: IOS-XE 17.12.6a)
	RldpRetries            int    `json:"rldp-retries"`               // RLDP attempts (Live: IOS-XE 17.12.6a)
	RogueClassTypeChange   string `json:"rogue-class-type-change"`    // Classification Type Change (Live: IOS-XE 17.12.6a)
	RogueStateChange       string `json:"rogue-state-change"`         // State Type Change (Live: IOS-XE 17.12.6a)
	RogueIfNum             int    `json:"rogue-if-num"`               // Interface Number (Live: IOS-XE 17.12.6a)
	ManagedAp              bool   `json:"managed-ap"`                 // Managed AP locally or via AP list with same MAC (Live: IOS-XE 17.12.6a)
	AutocontainAdhocTrap   bool   `json:"autocontain-adhoc-trap"`     // Trap for AdHoc Auto-Containment sent (Live: IOS-XE 17.12.6a)
	AutocontainTrap        bool   `json:"autocontain-trap"`           // Trap for Auto-Containment sent (Live: IOS-XE 17.12.6a)
	PotentialHoneypotTrap  bool   `json:"potential-honeypot-trap"`    // Trap for Potential Honeypot sent (Live: IOS-XE 17.12.6a)

	History struct {
		EventHistory []struct {
			Event      int    `json:"event"`       // Event identifier (Live: IOS-XE 17.12.6a)
			State      int    `json:"state"`       // State identifier (Live: IOS-XE 17.12.6a)
			Context    int    `json:"context"`     // Context identifier (Live: IOS-XE 17.12.6a)
			ContextStr string `json:"context-str"` // Context description (Live: IOS-XE 17.12.6a)
			CurrentRc  int    `json:"current-rc"`  // Current return code (Live: IOS-XE 17.12.6a)
			Count      int    `json:"count"`       // Event occurrence count (Live: IOS-XE 17.12.6a)
			Sticky     bool   `json:"sticky"`      // Sticky event flag (Live: IOS-XE 17.12.6a)
			Timestamp  string `json:"timestamp"`   // Event timestamp (Live: IOS-XE 17.12.6a)
		} `json:"event-history"` // Event history entries (Live: IOS-XE 17.12.6a)
	} `json:"history"` // Event History (Live: IOS-XE 17.12.6a)

	RldpLastResult              string `json:"rldp-last-result"`                          // RLDP Last Result (Live: IOS-XE 17.12.6a)
	RldpInProgress              bool   `json:"rldp-in-progress"`                          // RLDP in progress (Live: IOS-XE 17.12.6a)
	MaxDetectedRSSI             int    `json:"max-detected-rssi"`                         // Max RSSI value of all detecting APs (Live: IOS-XE 17.12.6a)
	SsidMaxRSSI                 string `json:"ssid-max-rssi"`                             // SSID of rogue detected by AP with max RSSI (Live: IOS-XE 17.12.6a)
	ApNameMaxRSSI               string `json:"ap-name-max-rssi"`                          // AP name of detecting AP with max RSSI (Live: IOS-XE 17.12.6a)
	DetectingRadioType80211n25g []any  `json:"detecting-radio-type-80211n-24g,omitempty"` // Radio type detecting APs. 802.11n 2.4GHz (Live: IOS-XE 17.12.6a)
	DetectingRadioType80211g    []any  `json:"detecting-radio-type-80211g,omitempty"`     // Radio type detecting APs. 802.11g (Live: IOS-XE 17.12.6a)
	DetectingRadioType80211bg   []any  `json:"detecting-radio-type-802-11bg,omitempty"`   // Radio type detecting APs. 802.11bg (Live: IOS-XE 17.12.6a)
	DRadioType80211ax24g        []any  `json:"d-radio-type-802-11ax24g,omitempty"`        // Radio type detecting APs. 802.11ax 2.4GHz (Live: IOS-XE 17.12.6a)
	DRadioType80211ax5g         []any  `json:"d-radio-type-802-11ax5g,omitempty"`         // Radio type detecting APs. 802.11ax 5GHz (Live: IOS-XE 17.12.6a)
	LradMACMaxRSSI              string `json:"lrad-mac-max-rssi"`                         // MAC Address of detecting AP with max RSSI (Live: IOS-XE 17.12.6a)
	RogueRadioTypeMaxRSSI       string `json:"rogue-radio-type-max-rssi"`                 // Radio type of detecting AP with max RSSI (Live: IOS-XE 17.12.6a)
	LastChannel                 int    `json:"last-channel"`                              // Channel number of last detecting APs (Live: IOS-XE 17.12.6a)
	RadioTypeCount              []int  `json:"radio-type-count"`                          // Number of radio type count (Live: IOS-XE 17.12.6a)

	LastHeardLradKey struct {
		LradMACAddr string `json:"lrad-mac-addr"` // MAC Address of AP interface that detected (Live: IOS-XE 17.12.6a)
		SlotID      int    `json:"slot-id"`       // Slot identifier (Live: IOS-XE 17.12.6a)
	} `json:"last-heard-lrad-key"` // Last Local Radio Key (Live: IOS-XE 17.12.6a)

	NLrads int `json:"n-lrads"` // Total number of APs that detected this rogue (Live: IOS-XE 17.12.6a)

	RogueLrad []struct {
		LradMACAddr string `json:"lrad-mac-addr"` // MAC Address of AP interface that detected (Live: IOS-XE 17.12.6a)
		SlotID      int    `json:"slot-id"`       // Slot identifier (Live: IOS-XE 17.12.6a)
		Ssid        string `json:"ssid"`          // SSID Advertised by Rogue Station (Live: IOS-XE 17.12.6a)
		HiddenSsid  bool   `json:"hidden-ssid"`   // Hidden ssid indication on detecting AP (Live: IOS-XE 17.12.6a)
		Name        string `json:"name"`          // Name of Detecting AP Interface (Live: IOS-XE 17.12.6a)
		RSSI        struct {
			Val int `json:"val"` // Value (Live: IOS-XE 17.12.6a)
			Num int `json:"num"` // Numerator (Live: IOS-XE 17.12.6a)
			Den int `json:"den"` // Denominator (Live: IOS-XE 17.12.6a)
		} `json:"rssi"` // Rogue RSSI as seen by Detecting AP Interface (Live: IOS-XE 17.12.6a)
		SNR struct {
			Val int `json:"val"` // Value (Live: IOS-XE 17.12.6a)
			Num int `json:"num"` // Numerator (Live: IOS-XE 17.12.6a)
			Den int `json:"den"` // Denominator (Live: IOS-XE 17.12.6a)
		} `json:"snr"` // SNR seen by Detecting AP Interface from Rogue (Live: IOS-XE 17.12.6a)
		ShortPreamble           int    `json:"short-preamble"`            // Preamble on this detecting AP (Live: IOS-XE 17.12.6a)
		Channel                 int    `json:"channel"`                   // Advertised Channel Number of Detecting AP Interface (Live: IOS-XE 17.12.6a)
		Channels                []int  `json:"channels"`                  // Advertised channels (Live: IOS-XE 17.12.6a)
		Encrypted               int    `json:"encrypted"`                 // Encryption mode on this detecting AP (Live: IOS-XE 17.12.6a)
		WpaSupport              int    `json:"wpa-support"`               // WPA mode on this detecting AP (Live: IOS-XE 17.12.6a)
		Dot11PhySupport         int    `json:"dot11-phy-support"`         // Rogue Radio Type (Live: IOS-XE 17.12.6a)
		LastHeard               string `json:"last-heard"`                // No. of seconds ago when Rogue was last heard (Live: IOS-XE 17.12.6a)
		ChanWidth               int    `json:"chan-width"`                // Channel Width (Live: IOS-XE 17.12.6a)
		ChanWidthLabel          int    `json:"chan-width-label"`          // Channel Width Label (Live: IOS-XE 17.12.6a)
		ExtChan                 int    `json:"ext-chan"`                  // Extension Channel (Live: IOS-XE 17.12.6a)
		BandID                  int    `json:"band-id"`                   // Band Identifier (Live: IOS-XE 17.12.6a)
		NumSlots                int    `json:"num-slots"`                 // Number of slots (Live: IOS-XE 17.12.6a)
		ReportRadioType         int    `json:"report-radio-type"`         // Reported radio type (Live: IOS-XE 17.12.6a)
		ContainResult           string `json:"contain-result"`            // Containment result (Live: IOS-XE 17.12.6a)
		ContainSlotID           string `json:"contain-slot-id"`           // Containment slot ID (Live: IOS-XE 17.12.6a)
		ContainRadioType        int    `json:"contain-radio-type"`        // Containment radio type (Live: IOS-XE 17.12.6a)
		RadioType               string `json:"radio-type"`                // Radio type identifier (Live: IOS-XE 17.12.6a)
		ContainmentType         string `json:"containment-type"`          // Containment type (Live: IOS-XE 17.12.6a)
		ContainmentChannelCount int    `json:"containment-channel-count"` // Containment channel count (Live: IOS-XE 17.12.6a)
		RogueContainmentChans   string `json:"rogue-containment-chans"`   // Rogue containment channels (Live: IOS-XE 17.12.6a)
		AuthFailCount           int    `json:"auth-fail-count"`           // Authentication failure count (Live: IOS-XE 17.12.6a)
		MfpStatus               string `json:"mfp-status"`                // Management frame protection status (Live: IOS-XE 17.12.6a)
		ChannelFromDS           bool   `json:"channel-from-ds"`           // Channel from distribution system (Live: IOS-XE 17.12.6a)
		PhyApSlot               int    `json:"phy-ap-slot"`               // Physical AP slot (Live: IOS-XE 17.12.6a)
	} `json:"rogue-lrad"` // Local Radio that detected this rogue (Live: IOS-XE 17.12.6a)

	NClients int `json:"n-clients"` // Total number of Clients detected on this rogue (Live: IOS-XE 17.12.6a)

	RogueClient *[]struct {
		RogueClientAddress string `json:"rogue-client-address"` // Rogue Client Address (Live: IOS-XE 17.12.6a)
	} `json:"rogue-client,omitempty"` // Rogue Client Address (Live: IOS-XE 17.12.6a)

	RemoteOverride struct {
		RemoteOverrideClassType        string `json:"remote-override-class-type"`        // Remote override classification (Live: IOS-XE 17.12.6a)
		RemoteOverrideMode             string `json:"remote-override-mode"`              // Remote override mode (Live: IOS-XE 17.12.6a)
		RemoteOverrideContainmentLevel int    `json:"remote-override-containment-level"` // Remote override containment level (Live: IOS-XE 17.12.6a)
	} `json:"remote-override"` // Remote containment override (Live: IOS-XE 17.12.6a)

	LastHeardSsid     string `json:"last-heard-ssid"`      // Last detected SSID advertised by Rogue station (Live: IOS-XE 17.12.6a)
	MfpRequired       bool   `json:"mfp-required"`         // This rogue requires 802.11w PMF (Live: IOS-XE 17.12.6a)
	ChannelMaxRSSI    int    `json:"channel-max-rssi"`     // Channel reported by the maximum RSSI LRAD (Live: IOS-XE 17.12.6a)
	WpaSupportMaxRSSI string `json:"wpa-support-max-rssi"` // WPA encryption reported by maximum RSSI LRAD (Live: IOS-XE 17.12.6a)
	EncryptedMaxRSSI  string `json:"encrypted-max-rssi"`   // Encryption reported by the maximum RSSI LRAD (Live: IOS-XE 17.12.6a)
	SNRMaxRSSI        int    `json:"snr-max-rssi"`         // Signal-to-noise ratio of maximum RSSI LRAD (Live: IOS-XE 17.12.6a)
	Properties        string `json:"properties"`           // Rogue AP properties (Live: IOS-XE 17.12.6a)

	BandData2dot4g struct {
		ChanWidth         int `json:"chan-width"`         // Channel Width (Live: IOS-XE 17.12.6a)
		ActualContainment int `json:"actual-containment"` // Actual containment (Live: IOS-XE 17.12.6a)
	} `json:"band-data-2dot4g"` // Rogue AP 2.4GHz band-specific data (Live: IOS-XE 17.12.6a)

	BandData5g struct {
		ChanWidth         int `json:"chan-width"`         // Channel Width (Live: IOS-XE 17.12.6a)
		ActualContainment int `json:"actual-containment"` // Actual containment (Live: IOS-XE 17.12.6a)
	} `json:"band-data-5g"` // Rogue AP 5GHz band-specific data (Live: IOS-XE 17.12.6a)

	BandData6g struct {
		ChanWidth         int `json:"chan-width"`         // Channel Width (Live: IOS-XE 17.12.6a)
		ActualContainment int `json:"actual-containment"` // Actual containment (Live: IOS-XE 17.12.6a)
	} `json:"band-data-6g"` // Rogue AP 6GHz band-specific data (Live: IOS-XE 17.12.6a)
}

// RogueClientDetail represents detailed information about detected rogue clients.
type RogueClientDetail struct {
	RogueClientAddress          string `json:"rogue-client-address"`           // Mac Address of Rogue Station (YANG: IOS-XE 17.12.1)
	RogueClientBssid            string `json:"rogue-client-bssid"`             // Rogue BSSID (YANG: IOS-XE 17.12.1)
	RogueClientGateway          string `json:"rogue-client-gateway"`           // MAC Address of the rogue AP Client gateway (YANG: IOS-XE 17.12.1)
	RogueClientState            string `json:"rogue-client-state"`             // State in which the Rogue AP is (YANG: IOS-XE 17.12.1)
	RogueClientContainmentLevel int    `json:"rogue-client-containment-level"` // Level of containment if the state is contained (YANG: IOS-XE 17.12.1)
	ActualContainment           int    `json:"actual-containment"`             // Number of Containing APs (YANG: IOS-XE 17.12.1)
	ContainmentType             string `json:"containment-type"`               // Containment mode applied to this rogue client (YANG: IOS-XE 17.12.1)
	RogueClientIfNum            int    `json:"rogue-client-if-num"`            // Rouge Client interface number (YANG: IOS-XE 17.12.1)
	RogueClientIPv4Addr         string `json:"rogue-client-ipv4-addr"`         // Rogue Client IPv4 address (YANG: IOS-XE 17.12.1)
	RogueClientIPv6Addr         string `json:"rogue-client-ipv6-addr"`         // Rogue Client IPv6 address (YANG: IOS-XE 17.12.1)
	ManualContained             bool   `json:"manual-contained"`               // Manually Contained (YANG: IOS-XE 17.12.1)
	Contained                   bool   `json:"contained"`                      // Contained (YANG: IOS-XE 17.12.1)
	AaaCheck                    bool   `json:"aaa-check"`                      // AAA Validation Status of this rogue client (YANG: IOS-XE 17.12.1)
	CmxCheck                    bool   `json:"cmx-check"`                      // CMX Validation Status of this rogue client (YANG: IOS-XE 17.12.1)
	RogueClientFirstTimestamp   string `json:"rogue-client-first-timestamp"`   // Time Stamp when this Rogue was First Detected (YANG: IOS-XE 17.12.1)
	RogueClientLastTimestamp    string `json:"rogue-client-last-timestamp"`    // Time Stamp when this Rogue was Last Detected (YANG: IOS-XE 17.12.1)

	LastHeardLradKey struct {
		LradMACAddr string `json:"lrad-mac-addr"` // MAC Address of AP interface that detected (YANG: IOS-XE 17.12.1)
		SlotID      int    `json:"slot-id"`       // Slot identifier (YANG: IOS-XE 17.12.1)
	} `json:"last-heard-lrad-key"` // Last Local Radio Key (YANG: IOS-XE 17.12.1)

	History struct {
		EventHistory []struct {
			Event      int    `json:"event"`       // Event identifier (Live: IOS-XE 17.12.6a)
			State      int    `json:"state"`       // State identifier (Live: IOS-XE 17.12.6a)
			Context    int    `json:"context"`     // Context identifier (Live: IOS-XE 17.12.6a)
			ContextStr string `json:"context-str"` // Context description (Live: IOS-XE 17.12.6a)
			CurrentRc  int    `json:"current-rc"`  // Current return code (Live: IOS-XE 17.12.6a)
			Count      int    `json:"count"`       // Event occurrence count (Live: IOS-XE 17.12.6a)
			Sticky     bool   `json:"sticky"`      // Sticky event flag (Live: IOS-XE 17.12.6a)
			Timestamp  string `json:"timestamp"`   // Event timestamp (Live: IOS-XE 17.12.6a)
		} `json:"event-history"` // Event history entries (Live: IOS-XE 17.12.6a)
	} `json:"history"` // Event history for this rogue client (YANG: IOS-XE 17.12.1)

	ParentRogueDataAddress string `json:"parent-rogue-data-address"` // MAC Address of the rogue AP this client is connected to (YANG: IOS-XE 17.12.1)

	RogueClientLrad []struct {
		LradMACAddr string `json:"lrad-mac-addr"` // MAC Address of AP interface that detected (YANG: IOS-XE 17.12.1)
		SlotID      int    `json:"slot-id"`       // Slot identifier (YANG: IOS-XE 17.12.1)
		LastHeard   string `json:"last-heard"`    // No of seconds ago when this Rogue was last heard by this AP (YANG: IOS-XE 17.12.1)
		Name        string `json:"name"`          // Name of Airespace AP Interface that detected the Rogue (YANG: IOS-XE 17.12.1)
		RSSI        struct {
			Val int `json:"val"` // Value (Live: IOS-XE 17.12.6a)
			Num int `json:"num"` // Numerator (Live: IOS-XE 17.12.6a)
			Den int `json:"den"` // Denominator (Live: IOS-XE 17.12.6a)
		} `json:"rssi"` // RSSI seen by Airespace AP Interface from the Rogue (YANG: IOS-XE 17.12.1)
		SNR struct {
			Val int `json:"val"` // Value (Live: IOS-XE 17.12.6a)
			Num int `json:"num"` // Numerator (Live: IOS-XE 17.12.6a)
			Den int `json:"den"` // Denominator (Live: IOS-XE 17.12.6a)
		} `json:"snr"` // SNR seen by Airespace AP Interface from Rogue (YANG: IOS-XE 17.12.1)
		Channel          int    `json:"channel"`            // The advertised Channel Number of that the Airespace AP Interface picked up from the Rogue (YANG: IOS-XE 17.12.1)
		BandID           int    `json:"band-id"`            // Band Identifier (YANG: IOS-XE 17.12.1)
		NumSlots         int    `json:"num-slots"`          // Number of slots in this AP (YANG: IOS-XE 17.12.1)
		ContainSlotID    string `json:"contain-slot-id"`    // Slot performing containment (YANG: IOS-XE 17.12.1)
		ContainRadioType int    `json:"contain-radio-type"` // Radio type for containment (YANG: IOS-XE 17.12.1)
		ContainResult    string `json:"contain-result"`     // Last containment result (YANG: IOS-XE 17.12.1)
		PhyApSlot        int    `json:"phy-ap-slot"`        // Manageability AP slot of the reporting LRAD (YANG: IOS-XE 17.12.1)
	} `json:"rogue-client-lrad"` // Local Radio that detected this rogue (YANG: IOS-XE 17.12.1)

	BandData2dot4g struct {
		NLrads            int `json:"n-lrads"`            // Number of LRADs for 2.4GHz (Live: IOS-XE 17.12.6a)
		ActualContainment int `json:"actual-containment"` // Actual containment for 2.4GHz (Live: IOS-XE 17.12.6a)
	} `json:"band-data-2dot4g"` // 2.4GHz band data for client (Live: IOS-XE 17.12.6a)

	BandData5g struct {
		NLrads            int `json:"n-lrads"`            // Number of LRADs for 5GHz (Live: IOS-XE 17.12.6a)
		ActualContainment int `json:"actual-containment"` // Actual containment for 5GHz (Live: IOS-XE 17.12.6a)
	} `json:"band-data-5g"` // 5GHz band data for client (Live: IOS-XE 17.12.6a)

	BandData6g struct {
		NLrads            int `json:"n-lrads"`            // Number of LRADs for 6GHz (Live: IOS-XE 17.12.6a)
		ActualContainment int `json:"actual-containment"` // Actual containment for 6GHz (Live: IOS-XE 17.12.6a)
	} `json:"band-data-6g"` // 6GHz band data for client (Live: IOS-XE 17.12.6a)
}
