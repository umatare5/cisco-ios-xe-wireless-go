// Package model contains generated response structures for the Cisco WNC API.
// This file contains rogue access point detection operational data structures.
package model

// RogueOperResponse represents the response structure for rogue operational data.
type RogueOperResponse struct {
	CiscoIOSXEWirelessRogueOperData struct {
		RogueStats      RogueStats        `json:"rogue-stats"`
		RogueData       []RogueData       `json:"rogue-data"`
		RogueClientData []RogueClientData `json:"rogue-client-data"`
		RldpStats       RldpStats         `json:"rldp-stats"`
	} `json:"Cisco-IOS-XE-wireless-rogue-oper:rogue-oper-data"`
}

// RogueStatsResponse represents the response structure for rogue statistics.
type RogueStatsResponse struct {
	RogueStats RogueStats `json:"Cisco-IOS-XE-wireless-rogue-oper:rogue-stats"`
}

// RogueDataResponse represents the response structure for rogue data.
type RogueDataResponse struct {
	RogueData []RogueData `json:"Cisco-IOS-XE-wireless-rogue-oper:rogue-data"`
}

// RogueClientDataResponse represents the response structure for rogue client data.
type RogueClientDataResponse struct {
	RogueClientData []RogueClientData `json:"Cisco-IOS-XE-wireless-rogue-oper:rogue-client-data"`
}

// RldpStatsResponse represents the response structure for RLDP statistics.
type RldpStatsResponse struct {
	RldpStats RldpStats `json:"Cisco-IOS-XE-wireless-rogue-oper:rldp-stats"`
}

// RogueStats represents rogue statistics including counts of various rogue types.
type RogueStats struct {
	RestartCount                int    `json:"restart-count"`
	PendingCount                int    `json:"pending-count"`
	LradCount                   int    `json:"lrad-count"`
	OnMyNetworkCount            int    `json:"on-my-network-count"`
	AdhocCount                  int    `json:"adhoc-count"`
	UnknownCount                int    `json:"unknown-count"`
	UnclassifiedCount           int    `json:"unclassified-count"`
	MaliciousCount              int    `json:"malicious-count"`
	FriendlyCount               int    `json:"friendly-count"`
	CustomCount                 int    `json:"custom-count"`
	NotAdhocCount               int    `json:"not-adhoc-count"`
	TotalCount                  int    `json:"total-count"`
	ContainedCount              int    `json:"contained-count"`
	ContainedClientCount        int    `json:"contained-client-count"`
	ContainedPendingCount       int    `json:"contained-pending-count"`
	ContainedPendingClientCount int    `json:"contained-pending-client-count"`
	TotalClientCount            int    `json:"total-client-count"`
	MaxCount                    int    `json:"max-count"`
	MaxClientCount              int    `json:"max-client-count"`
	ReportCount                 string `json:"report-count"`
	ClientReportCount           string `json:"client-report-count"`
	RateReportCount             int    `json:"rate-report-count"`
	RateClientReportCount       int    `json:"rate-client-report-count"`
	IappApPkt                   string `json:"iapp-ap-pkt"`
	IappClientPkt               string `json:"iapp-client-pkt"`
	RateIappApPkt               int    `json:"rate-iapp-ap-pkt"`
	RateIappClientPkt           int    `json:"rate-iapp-client-pkt"`
	RldpCount                   string `json:"rldp-count"`
	AaaMsgRxCount               string `json:"aaa-msg-rx-count"`
	AaaMsgTxCount               string `json:"aaa-msg-tx-count"`
	SnmpTrapsTxCount            string `json:"snmp-traps-tx-count"`
	LradOffCount                string `json:"lrad-off-count"`
	ApCreateCount               string `json:"ap-create-count"`
	ApDeleteCount               string `json:"ap-delete-count"`
	ApRadioUpCount              string `json:"ap-radio-up-count"`
	ApRadioDownCount            string `json:"ap-radio-down-count"`
	ApNameChangeCount           int    `json:"ap-name-change-count"`
	WncdIpcTxCount              string `json:"wncd-ipc-tx-count"`
	WncdIpcRxCount              string `json:"wncd-ipc-rx-count"`
	WncmgrIpcRxCount            string `json:"wncmgr-ipc-rx-count"`
	IosIpcTxCount               string `json:"ios-ipc-tx-count"`
	IosIpcRxCount               string `json:"ios-ipc-rx-count"`
	NmspdIpcTxCount             string `json:"nmspd-ipc-tx-count"`
	NmspdIpcRxCount             string `json:"nmspd-ipc-rx-count"`
	ContainMsgCount             string `json:"contain-msg-count"`
	FsmErrors                   int    `json:"fsm-errors"`
	TrapErrors                  int    `json:"trap-errors"`

	EnqCount struct {
		Counters []struct {
			Value       string `json:"value"`
			Description string `json:"description"`
		} `json:"counters"`
	} `json:"enq-count"`

	SimilarApReportCount     string `json:"similar-ap-report-count"`
	SimilarClientReportCount string `json:"similar-client-report-count"`

	SnmpTrapsPerType struct {
		Counters []struct {
			Value       string `json:"value"`
			Description string `json:"description"`
		} `json:"counters"`
	} `json:"snmp-traps-per-type"`

	IappReportMessagesLoadShedded int `json:"iapp-report-messages-load-shedded"`
	ManagedClientMessageCount     int `json:"managed-client-message-count"`
	ManagedClientJoinCount        int `json:"managed-client-join-count"`
	ManagedClientLeaveCount       int `json:"managed-client-leave-count"`
	ManagedRogueClientCount       int `json:"managed-rogue-client-count"`

	ProcTime struct {
		Counters []struct {
			Value       string `json:"value"`
			Description string `json:"description"`
		} `json:"counters"`
	} `json:"proc-time"`

	GlobalHistory struct {
		EventHistory []struct {
			Event      int    `json:"event"`
			State      int    `json:"state"`
			Context    int    `json:"context"`
			ContextStr string `json:"context-str"`
			CurrentRc  int    `json:"current-rc"`
			Count      int    `json:"count"`
			Sticky     bool   `json:"sticky"`
			Timestamp  string `json:"timestamp"`
		} `json:"event-history"`
	} `json:"global-history"`

	TblApfVapCacheReloadCount     int    `json:"tbl-apf-vap-cache-reload-count"`
	NewLradCount                  string `json:"new-lrad-count"`
	LradPurgeCount                string `json:"lrad-purge-count"`
	RssiChangeCount               string `json:"rssi-change-count"`
	FinalStateChangeCount         string `json:"final-state-change-count"`
	ContainLevelChangeCount       string `json:"contain-level-change-count"`
	ClassChangeCount              string `json:"class-change-count"`
	AdhocChangeCount              string `json:"adhoc-change-count"`
	OnMyNetworkChangeCount        string `json:"on-my-network-change-count"`
	NClientsChangedCount          string `json:"n-clients-changed-count"`
	ClientNewLradCount            string `json:"client-new-lrad-count"`
	ClientLradPurgeCount          string `json:"client-lrad-purge-count"`
	ClientRssiChangeCount         string `json:"client-rssi-change-count"`
	ClientFinalStateChangeCount   string `json:"client-final-state-change-count"`
	ClientContainLevelChangeCount string `json:"client-contain-level-change-count"`
	ClientChannelChangeCount      string `json:"client-channel-change-count"`
	ClientIPChangeCount           string `json:"client-ip-change-count"`
	ClientRoamCount               string `json:"client-roam-count"`

	RogueApReportsDroppedScale        string `json:"rogue-ap-reports-dropped-scale"`
	RogueClientReportsDroppedScale    string `json:"rogue-client-reports-dropped-scale"`
	RogueClientReportsDroppedNoParent string `json:"rogue-client-reports-dropped-no-parent"`

	RogueEnabled bool   `json:"rogue-enabled"`
	MmIpcRxCount string `json:"mm-ipc-rx-count"`

	RogueWsaEventsTriggeredCounter string `json:"rogue-wsa-events-triggered-counter"`
	RogueWsaEventsEnqueuedCounter  string `json:"rogue-wsa-events-enqueued-counter"`

	RogueWsaEventsTriggeredPerType struct {
		Counters []struct {
			Value       string `json:"value"`
			Description string `json:"description"`
		} `json:"counters"`
	} `json:"rogue-wsa-events-triggered-per-type"`

	RogueWsaEventsEnqueuedPerType struct {
		Counters []struct {
			Value       string `json:"value"`
			Description string `json:"description"`
		} `json:"counters"`
	} `json:"rogue-wsa-events-enqueued-per-type"`

	BssidIpcCount        string `json:"bssid-ipc-count"`
	ApChannelChangeCount string `json:"ap-channel-change-count"`
	BeaconDsAttackCount  string `json:"beacon-ds-attack-count"`

	InternalCount int `json:"internal-count"`
	ExternalCount int `json:"external-count"`
	AlertCount    int `json:"alert-count"`
	ThreatCount   int `json:"threat-count"`

	RogueApReportFalseDrop string `json:"rogue-ap-report-false-drop"`

	AdhocUnknownCount      int `json:"adhoc-unknown-count"`
	AdhocUnclassifiedCount int `json:"adhoc-unclassified-count"`
	AdhocMaliciousCount    int `json:"adhoc-malicious-count"`
	AdhocFriendlyCount     int `json:"adhoc-friendly-count"`
	AdhocCustomCount       int `json:"adhoc-custom-count"`

	MaliciousInitCount    int `json:"malicious-init-count"`
	CustomInitCount       int `json:"custom-init-count"`
	UnclassifiedInitCount int `json:"unclassified-init-count"`
	FriendlyInitCount     int `json:"friendly-init-count"`
	UnknownInitCount      int `json:"unknown-init-count"`
	InitCount             int `json:"init-count"`
	PostInitCount         int `json:"post-init-count"`
	MaxInitCount          int `json:"max-init-count"`

	TotalMaliciousCount    int `json:"total-malicious-count"`
	TotalCustomCount       int `json:"total-custom-count"`
	TotalUnclassifiedCount int `json:"total-unclassified-count"`
	TotalFriendlyCount     int `json:"total-friendly-count"`
	TotalUnknownCount      int `json:"total-unknown-count"`
}

// RogueData represents detailed information about a detected rogue device.
type RogueData struct {
	RogueAddress           string `json:"rogue-address"`
	RogueClassType         string `json:"rogue-class-type"`
	RogueMode              string `json:"rogue-mode"`
	RogueContainmentLevel  int    `json:"rogue-containment-level"`
	ActualContainment      int    `json:"actual-containment"`
	ManualContained        bool   `json:"manual-contained"`
	ClassOverrideSrc       string `json:"class-override-src"`
	ContainmentType        string `json:"containment-type"`
	Contained              bool   `json:"contained"`
	SeverityScore          int    `json:"severity-score"`
	ClassTypeCustomName    string `json:"class-type-custom-name"`
	RogueFirstTimestamp    string `json:"rogue-first-timestamp"`
	RogueLastTimestamp     string `json:"rogue-last-timestamp"`
	RogueIsOnMyNetwork     bool   `json:"rogue-is-on-my-network"`
	AdHoc                  bool   `json:"ad-hoc"`
	AdHocBssid             string `json:"ad-hoc-bssid"`
	RogueRuleName          string `json:"rogue-rule-name"`
	RogueRadioTypeLastSeen string `json:"rogue-radio-type-last-seen"`
	RldpRetries            int    `json:"rldp-retries"`
	RogueClassTypeChange   string `json:"rogue-class-type-change"`
	RogueStateChange       string `json:"rogue-state-change"`
	RogueIfNum             int    `json:"rogue-if-num"`
	ManagedAp              bool   `json:"managed-ap"`
	AutocontainAdhocTrap   bool   `json:"autocontain-adhoc-trap"`
	AutocontainTrap        bool   `json:"autocontain-trap"`
	PotentialHoneypotTrap  bool   `json:"potential-honeypot-trap"`

	History struct {
		EventHistory []struct {
			Event      int    `json:"event"`
			State      int    `json:"state"`
			Context    int    `json:"context"`
			ContextStr string `json:"context-str"`
			CurrentRc  int    `json:"current-rc"`
			Count      int    `json:"count"`
			Sticky     bool   `json:"sticky"`
			Timestamp  string `json:"timestamp"`
		} `json:"event-history"`
	} `json:"history"`

	RldpLastResult              string `json:"rldp-last-result"`
	RldpInProgress              bool   `json:"rldp-in-progress"`
	MaxDetectedRssi             int    `json:"max-detected-rssi"`
	SsidMaxRssi                 string `json:"ssid-max-rssi"`
	ApNameMaxRssi               string `json:"ap-name-max-rssi"`
	DetectingRadioType80211n24g []any  `json:"detecting-radio-type-80211n-24g,omitempty"`
	LradMacMaxRssi              string `json:"lrad-mac-max-rssi"`
	RogueRadioTypeMaxRssi       string `json:"rogue-radio-type-max-rssi"`
	LastChannel                 int    `json:"last-channel"`
	RadioTypeCount              []int  `json:"radio-type-count"`

	LastHeardLradKey struct {
		LradMacAddr string `json:"lrad-mac-addr"`
		SlotID      int    `json:"slot-id"`
	} `json:"last-heard-lrad-key"`

	NLrads int `json:"n-lrads"`

	RogueLrad []struct {
		LradMacAddr string `json:"lrad-mac-addr"`
		SlotID      int    `json:"slot-id"`
		Ssid        string `json:"ssid"`
		HiddenSsid  bool   `json:"hidden-ssid"`
		Name        string `json:"name"`
		Rssi        struct {
			Val int `json:"val"`
			Num int `json:"num"`
			Den int `json:"den"`
		} `json:"rssi"`
		Snr struct {
			Val int `json:"val"`
			Num int `json:"num"`
			Den int `json:"den"`
		} `json:"snr"`
		ShortPreamble           int    `json:"short-preamble"`
		Channel                 int    `json:"channel"`
		Channels                []int  `json:"channels"`
		Encrypted               int    `json:"encrypted"`
		WpaSupport              int    `json:"wpa-support"`
		Dot11PhySupport         int    `json:"dot11-phy-support"`
		LastHeard               string `json:"last-heard"`
		ChanWidth               int    `json:"chan-width"`
		ChanWidthLabel          int    `json:"chan-width-label"`
		ExtChan                 int    `json:"ext-chan"`
		BandID                  int    `json:"band-id"`
		NumSlots                int    `json:"num-slots"`
		ReportRadioType         int    `json:"report-radio-type"`
		ContainResult           string `json:"contain-result"`
		ContainSlotID           string `json:"contain-slot-id"`
		ContainRadioType        int    `json:"contain-radio-type"`
		RadioType               string `json:"radio-type"`
		ContainmentType         string `json:"containment-type"`
		ContainmentChannelCount int    `json:"containment-channel-count"`
		RogueContainmentChans   string `json:"rogue-containment-chans"`
		AuthFailCount           int    `json:"auth-fail-count"`
		MfpStatus               string `json:"mfp-status"`
		ChannelFromDS           bool   `json:"channel-from-ds"`
		PhyApSlot               int    `json:"phy-ap-slot"`
	} `json:"rogue-lrad"`

	NClients int `json:"n-clients"`

	RemoteOverride struct {
		RemoteOverrideClassType        string `json:"remote-override-class-type"`
		RemoteOverrideMode             string `json:"remote-override-mode"`
		RemoteOverrideContainmentLevel int    `json:"remote-override-containment-level"`
	} `json:"remote-override"`

	LastHeardSsid     string `json:"last-heard-ssid"`
	MfpRequired       bool   `json:"mfp-required"`
	ChannelMaxRssi    int    `json:"channel-max-rssi"`
	WpaSupportMaxRssi string `json:"wpa-support-max-rssi"`
	EncryptedMaxRssi  string `json:"encrypted-max-rssi"`
	SnrMaxRssi        int    `json:"snr-max-rssi"`
	Properties        string `json:"properties"`

	BandData2dot4g struct {
		ChanWidth         int `json:"chan-width"`
		ActualContainment int `json:"actual-containment"`
	} `json:"band-data-2dot4g"`

	BandData5g struct {
		ChanWidth         int `json:"chan-width"`
		ActualContainment int `json:"actual-containment"`
	} `json:"band-data-5g"`

	BandData6g struct {
		ChanWidth         int `json:"chan-width"`
		ActualContainment int `json:"actual-containment"`
	} `json:"band-data-6g"`
}

// RogueClientData represents detailed information about a rogue client device.
type RogueClientData struct {
	RogueClientAddress          string `json:"rogue-client-address"`
	RogueClientBssid            string `json:"rogue-client-bssid"`
	RogueClientGateway          string `json:"rogue-client-gateway"`
	RogueClientState            string `json:"rogue-client-state"`
	RogueClientContainmentLevel int    `json:"rogue-client-containment-level"`
	ActualContainment           int    `json:"actual-containment"`
	ContainmentType             string `json:"containment-type"`
	RogueClientIfNum            int    `json:"rogue-client-if-num"`
	RogueClientIPv4Addr         string `json:"rogue-client-ipv4-addr"`
	RogueClientIPv6Addr         string `json:"rogue-client-ipv6-addr"`
	ManualContained             bool   `json:"manual-contained"`
	Contained                   bool   `json:"contained"`
	AaaCheck                    bool   `json:"aaa-check"`
	CmxCheck                    bool   `json:"cmx-check"`
	RogueClientFirstTimestamp   string `json:"rogue-client-first-timestamp"`
	RogueClientLastTimestamp    string `json:"rogue-client-last-timestamp"`

	LastHeardLradKey struct {
		LradMacAddr string `json:"lrad-mac-addr"`
		SlotID      int    `json:"slot-id"`
	} `json:"last-heard-lrad-key"`

	History struct {
		EventHistory []struct {
			Event      int    `json:"event"`
			State      int    `json:"state"`
			Context    int    `json:"context"`
			ContextStr string `json:"context-str"`
			CurrentRc  int    `json:"current-rc"`
			Count      int    `json:"count"`
			Sticky     bool   `json:"sticky"`
			Timestamp  string `json:"timestamp"`
		} `json:"event-history"`
	} `json:"history"`

	ParentRogueDataAddress string `json:"parent-rogue-data-address"`

	RogueClientLrad []struct {
		LradMacAddr string `json:"lrad-mac-addr"`
		SlotID      int    `json:"slot-id"`
		LastHeard   string `json:"last-heard"`
		Name        string `json:"name"`
		Rssi        struct {
			Val int `json:"val"`
			Num int `json:"num"`
			Den int `json:"den"`
		} `json:"rssi"`
		Snr struct {
			Val int `json:"val"`
			Num int `json:"num"`
			Den int `json:"den"`
		} `json:"snr"`
		Channel          int    `json:"channel"`
		BandID           int    `json:"band-id"`
		NumSlots         int    `json:"num-slots"`
		ContainSlotID    string `json:"contain-slot-id"`
		ContainRadioType int    `json:"contain-radio-type"`
		ContainResult    string `json:"contain-result"`
		PhyApSlot        int    `json:"phy-ap-slot"`
	} `json:"rogue-client-lrad"`

	BandData2dot4g struct {
		NLrads            int `json:"n-lrads"`
		ActualContainment int `json:"actual-containment"`
	} `json:"band-data-2dot4g"`

	BandData5g struct {
		NLrads            int `json:"n-lrads"`
		ActualContainment int `json:"actual-containment"`
	} `json:"band-data-5g"`

	BandData6g struct {
		NLrads            int `json:"n-lrads"`
		ActualContainment int `json:"actual-containment"`
	} `json:"band-data-6g"`
}

// RldpStats represents RLDP (Rogue Location Discovery Protocol) statistics.
type RldpStats struct {
	NumInProgress     int  `json:"num-in-progress"`
	NumRldpStarted    int  `json:"num-rldp-started"`
	AuthTimeout       int  `json:"auth-timeout"`
	AssocTimeout      int  `json:"assoc-timeout"`
	DhcpTimeout       int  `json:"dhcp-timeout"`
	NotConnected      int  `json:"not-connected"`
	Connected         int  `json:"connected"`
	RldpSocketEnabled bool `json:"rldp-socket-enabled"`
}
