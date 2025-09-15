package mcast

// McastOper represents multicast operational data from controller.
type McastOper struct {
	CiscoIOSXEWirelessMcastOperMcastOperData struct {
		FlexMediastreamClientSummary   []FlexMediastreamClientSummary   `json:"flex-mediastream-client-summary,omitempty"`    // FlexConnect mediastream client summary data (Live: IOS-XE 17.12.5)
		VlanL2MgidOp                   []VlanL2MgidOp                   `json:"vlan-l2-mgid-op,omitempty"`                    // VLAN Layer 2 multicast group ID operational data (Live: IOS-XE 17.12.5)
		FabricMediaStreamClientSummary []FabricMediaStreamClientSummary `json:"fabric-media-stream-client-summary,omitempty"` // Fabric mediastream client summary data (Live: IOS-XE 17.12.5)
		McastMgidInfo                  []McastMgidInfo                  `json:"mcast-mgid-info,omitempty"`                    // Multicast MGID information (YANG: IOS-XE 17.12.1)
		MulticastOperData              []MulticastOperData              `json:"multicast-oper-data,omitempty"`                // Multicast operational data (YANG: IOS-XE 17.12.1)
		RrcHistoryClientRecordData     []RrcHistoryClientRecordData     `json:"rrc-history-client-record-data,omitempty"`     // RRC history client record data (YANG: IOS-XE 17.12.1)
		RrcSrRadioRecord               []RrcSrRadioRecord               `json:"rrc-sr-radio-record,omitempty"`                // RRC stream radio record data (YANG: IOS-XE 17.12.1)
		RrcStreamRecord                []RrcStreamRecord                `json:"rrc-stream-record,omitempty"`                  // RRC stream record data (YANG: IOS-XE 17.12.1)
		RrcStreamAdmitRecord           []RrcStreamAdmitRecord           `json:"rrc-stream-admit-record,omitempty"`            // RRC stream admit record data (YANG: IOS-XE 17.12.1)
		RrcStreamDenyRecord            []RrcStreamDenyRecord            `json:"rrc-stream-deny-record,omitempty"`             // RRC stream deny record data (YANG: IOS-XE 17.12.1)
	} `json:"Cisco-IOS-XE-wireless-mcast-oper:mcast-oper-data"` // Multicast operational data container (Live: IOS-XE 17.12.5)
}

// McastOperFlexMediastreamClientSummary represents the structure for FlexConnect mediastream client summary data.
type McastOperFlexMediastreamClientSummary struct {
	FlexMediastreamClientSummary []FlexMediastreamClientSummary `json:"Cisco-IOS-XE-wireless-mcast-oper:flex-mediastream-client-summary"`
}

// McastOperVlanL2MgidOp represents the structure for VLAN Layer 2 multicast group ID operational data.
type McastOperVlanL2MgidOp struct {
	VlanL2MgidOp []VlanL2MgidOp `json:"Cisco-IOS-XE-wireless-mcast-oper:vlan-l2-mgid-op"`
}

// McastOperFabricMediaStreamClientSummary represents fabric mediastream client summary wrapper.
type McastOperFabricMediaStreamClientSummary struct {
	FabricMediaStreamClientSummary []FabricMediaStreamClientSummary `json:"Cisco-IOS-XE-wireless-mcast-oper:fabric-media-stream-client-summary"`
}

// McastOperMcastMgidInfo represents multicast MGID info wrapper.
type McastOperMcastMgidInfo struct {
	McastMgidInfo []McastMgidInfo `json:"Cisco-IOS-XE-wireless-mcast-oper:mcast-mgid-info"`
}

// McastOperMulticastOperData represents multicast operational data wrapper.
type McastOperMulticastOperData struct {
	MulticastOperData []MulticastOperData `json:"Cisco-IOS-XE-wireless-mcast-oper:multicast-oper-data"`
}

// FlexMediastreamClientSummary represents FlexConnect mediastream client information.
type FlexMediastreamClientSummary struct {
	ClientMAC            string                 `json:"client-mac"`              // Multicast flex client MAC address (Live: IOS-XE 17.12.5)
	VlanID               int                    `json:"vlan-id"`                 // Multicast client's VLAN (Live: IOS-XE 17.12.5)
	FlexMcastClientGroup []FlexMcastClientGroup `json:"flex-mcast-client-group"` // Flex multicast client group (Live: IOS-XE 17.12.5)
}

// FlexMcastClientGroup represents FlexConnect multicast client group configuration.
type FlexMcastClientGroup struct {
	McastIP    string `json:"mcast-ip"`    // Multicast group IP (Live: IOS-XE 17.12.5)
	StreamName string `json:"stream-name"` // Stream name associated with multicast group (Live: IOS-XE 17.12.5)
	ApMAC      string `json:"ap-mac"`      // AP MAC on which Multicast group is reported (Live: IOS-XE 17.12.5)
	IsDirect   bool   `json:"is-direct"`   // Stream is Multicast-Direct or Multicast (Live: IOS-XE 17.12.5)
}

// VlanL2MgidOp represents VLAN Layer 2 multicast group ID operational data.
type VlanL2MgidOp struct {
	VlanIndex               int  `json:"vlan-index"`                 // Multicast VLAN index (Live: IOS-XE 17.12.5)
	IsNonipMulticastEnabled bool `json:"is-nonip-multicast-enabled"` // Is non IP multicast enabled (Live: IOS-XE 17.12.5)
	IsBroadcastEnable       bool `json:"is-broadcast-enable"`        // Is broadcast enabled (Live: IOS-XE 17.12.5)
}

// FabricMediaStreamClientSummary represents fabric mediastream client summary.
type FabricMediaStreamClientSummary struct {
	ClientMAC              string                   `json:"client-mac"`                          // Multicast fabric client MAC address (Live: IOS-XE 17.12.5)
	VlanID                 int                      `json:"vlan-id"`                             // Multicast client's VLAN (Live: IOS-XE 17.12.5)
	FabricMcastClientGroup []FabricMcastClientGroup `json:"fabric-mcast-client-group,omitempty"` // Fabric multicast client group (Live: IOS-XE 17.12.5)
}

// FabricMcastClientGroup represents fabric multicast client group.
type FabricMcastClientGroup struct {
	McastIP    string `json:"mcast-ip"`    // Multicast IP (Live: IOS-XE 17.12.5)
	StreamName string `json:"stream-name"` // Stream name (Live: IOS-XE 17.12.5)
	ApMAC      string `json:"ap-mac"`      // AP MAC on which Multicast group is reported (Live: IOS-XE 17.12.5)
	IsDirect   bool   `json:"is-direct"`   // Stream is Multicast-Direct or Multicast (Live: IOS-XE 17.12.5)
}

// McastMgidInfo represents multicast MGID information.
type McastMgidInfo struct {
	Mgid                   int                    `json:"mgid"`                             // MGID for Multicast group (YANG: IOS-XE 17.12.1)
	Vlan                   int                    `json:"vlan"`                             // VLAN used for MGID (YANG: IOS-XE 17.12.1)
	McOnlyNumClients       int                    `json:"mc-only-num-clients"`              // Number of MC-only clients for MGID (YANG: IOS-XE 17.12.1)
	Mc2ucNumClients        int                    `json:"mc2uc-num-clients"`                // Number of MC-UC clients for MGID (YANG: IOS-XE 17.12.1)
	Mc2ucNumDenyClients    int                    `json:"mc2uc-num-deny-clients"`           // Number of MC-UC deny clients for MGID (YANG: IOS-XE 17.12.1)
	Mc2ucNumPendingClients int                    `json:"mc2uc-num-pending-clients"`        // Number of MC-UC pending client for MGID (YANG: IOS-XE 17.12.1)
	Group                  string                 `json:"group"`                            // Multicast group IP address (YANG: IOS-XE 17.12.1)
	McastMgidClientList    []McastMgidClientEntry `json:"mcast-mgid-client-list,omitempty"` // Multicast MGID client list (YANG: IOS-XE 17.12.1)
}

// McastMgidClientEntry represents multicast MGID client entry.
type McastMgidClientEntry struct {
	ClientMACAddr string `json:"client-mac-addr"` // client MAC address (YANG: IOS-XE 17.12.1)
	ClientIPAddr  string `json:"client-ip-addr"`  // Client IPv4 address (YANG: IOS-XE 17.12.1)
	ClientStatus  string `json:"client-status"`   // Client's multicast status (YANG: IOS-XE 17.12.1)
}

// MulticastOperData represents multicast operational data.
type MulticastOperData struct {
	MsMACAddress string             `json:"ms-mac-address"`          // Client MAC address (YANG: IOS-XE 17.12.1)
	NumEntries   int                `json:"num-entries"`             // Number of currently filled entries (YANG: IOS-XE 17.12.1)
	Entry        []McastClientEntry `json:"entry,omitempty"`         // Multicast client entry (YANG: IOS-XE 17.12.1)
	ClientIPv6   string             `json:"client-ipv6,omitempty"`   // Multicast client IPv6 address (YANG: IOS-XE 17.12.1)
	CAPWAPIifID  int                `json:"capwap-iif-id,omitempty"` // Client capwap IIF ID (YANG: IOS-XE 17.12.1)
	ClientIP     string             `json:"client-ip,omitempty"`     // Client IPv4 address (YANG: IOS-XE 17.12.1)
}

// McastClientEntry represents multicast client entry.
type McastClientEntry struct {
	Vlan         int    `json:"vlan"`          // Client's vlan for Multicast (YANG: IOS-XE 17.12.1)
	Mgid         int    `json:"mgid"`          // Mgid for Multicast group (YANG: IOS-XE 17.12.1)
	Group        string `json:"group"`         // Multicast group IP address (YANG: IOS-XE 17.12.1)
	ClientStatus string `json:"client-status"` // Client status of Multicast group (YANG: IOS-XE 17.12.1)
	Qos          string `json:"qos"`           // QOS value for Multicast group (YANG: IOS-XE 17.12.1)
	Used         bool   `json:"used"`          // Multicast client entry is in use or not (YANG: IOS-XE 17.12.1)
}

// RrcHistoryClientRecordData represents RRC history client record data.
type RrcHistoryClientRecordData struct {
	UserTimeStamp        string `json:"user-time-stamp"`         // RRC history client user timestamp (YANG: IOS-XE 17.12.1)
	ClientMAC            string `json:"client-mac"`              // RRC history client MAC address (YANG: IOS-XE 17.12.1)
	Qos                  int    `json:"qos"`                     // QOS for the streaming (YANG: IOS-XE 17.12.1)
	Decision             string `json:"decision"`                // RRC history Video-streaming decision (YANG: IOS-XE 17.12.1)
	ReasonCode           int    `json:"reason-code"`             // RRC decision's reason code (YANG: IOS-XE 17.12.1)
	ApMAC                string `json:"ap-mac"`                  // AP's MAC address to which client was connected (YANG: IOS-XE 17.12.1)
	VapID                int    `json:"vap-id"`                  // Client's VAP ID (YANG: IOS-XE 17.12.1)
	SlotID               int    `json:"slot-id"`                 // AP's Slot ID to which client was connected (YANG: IOS-XE 17.12.1)
	CacEnabled           int    `json:"cac-enabled"`             // RRC CAC is enabled or not (YANG: IOS-XE 17.12.1)
	StreamName           string `json:"stream-name"`             // Stream name associated with multicast group (YANG: IOS-XE 17.12.1)
	DstIPAddress         string `json:"dst-ip-address"`          // Multicast group destination IP address (YANG: IOS-XE 17.12.1)
	CfgStreamBw          int    `json:"cfg-stream-bw"`           // Configured max video bandwidth for new stream (YANG: IOS-XE 17.12.1)
	CurrentRate          int    `json:"current-rate"`            // Current data rate of client requesting stream (YANG: IOS-XE 17.12.1)
	VideoPktSize         int    `json:"video-pkt-size"`          // Packet size for the new video stream (YANG: IOS-XE 17.12.1)
	CurrVideoUtil        int    `json:"curr-video-util"`         // Current video utilization of AP radio (YANG: IOS-XE 17.12.1)
	CurrVoiceUtil        int    `json:"curr-voice-util"`         // Current voice utilization of AP radio (YANG: IOS-XE 17.12.1)
	CurrChannelUtil      int    `json:"curr-channel-util"`       // Current channel utilization of AP radio (YANG: IOS-XE 17.12.1)
	CurrQueueUtil        int    `json:"curr-queue-util"`         // Current queue utilization of AP radio (YANG: IOS-XE 17.12.1)
	CurrVideoPps         int    `json:"curr-video-pps"`          // Current video rate in packets per second (YANG: IOS-XE 17.12.1)
	VideoDelayHistSevere int    `json:"video-delay-hist-severe"` // Number of video packets with severe delay (YANG: IOS-XE 17.12.1)
	VideoPktLossDiscard  int    `json:"video-pkt-loss-discard"`  // Video packet loss discarded by AP (YANG: IOS-XE 17.12.1)
	VideoPktLossFail     int    `json:"video-pkt-loss-fail"`     // Video packet loss fail (YANG: IOS-XE 17.12.1)
	NumVideoStreams      int    `json:"num-video-streams"`       // Number of video streams (YANG: IOS-XE 17.12.1)
}

// RrcSrRadioRecord represents RRC stream radio record.
type RrcSrRadioRecord struct {
	ApMAC               string          `json:"ap-mac"`                        // AP MAC on which Multicast group is reported (YANG: IOS-XE 17.12.1)
	SlotID              int             `json:"slot-id"`                       // Radio Slot ID (YANG: IOS-XE 17.12.1)
	RadioType           int             `json:"radio-type"`                    // Radio type of the stream (YANG: IOS-XE 17.12.1)
	DuplicatedBandWidth int             `json:"duplicated-band-width"`         // Duplicated bandwidth on this radio (YANG: IOS-XE 17.12.1)
	LastReRrc           string          `json:"last-re-rrc"`                   // Last Re-RRC timestamp for stream (YANG: IOS-XE 17.12.1)
	NumberOfAdmitted    int             `json:"number-of-admitted"`            // Number of admitted streams (YANG: IOS-XE 17.12.1)
	RrcGroupsInRadio    []RrcGroupsInfo `json:"rrc-groups-in-radio,omitempty"` // RRC stream groups in radio (YANG: IOS-XE 17.12.1)
}

// RrcGroupsInfo represents RRC groups information.
type RrcGroupsInfo struct {
	DestIP      string `json:"dest-ip"`       // Multicast Group IP address (YANG: IOS-XE 17.12.1)
	NoOfStreams int    `json:"no-of-streams"` // Number of streams for group (YANG: IOS-XE 17.12.1)
}

// RrcStreamRecord represents RRC stream record.
type RrcStreamRecord struct {
	StreamNameStr  string             `json:"stream-name-str"`            // Name of the Media-stream (YANG: IOS-XE 17.12.1)
	GroupIP        string             `json:"group-ip"`                   // Multicast group IP address for this stream (YANG: IOS-XE 17.12.1)
	ClMAC          string             `json:"cl-mac"`                     // Media-stream client MAC address (YANG: IOS-XE 17.12.1)
	ClientMAC      string             `json:"client-mac"`                 // Client MAC address (YANG: IOS-XE 17.12.1)
	DestIP         string             `json:"dest-ip"`                    // Multicast group destination IP address (YANG: IOS-XE 17.12.1)
	VapID          int                `json:"vap-id"`                     // VAP ID associated with stream (YANG: IOS-XE 17.12.1)
	Vlan           int                `json:"vlan"`                       // VLAN ID associated with stream (YANG: IOS-XE 17.12.1)
	WlanID         int                `json:"wlan-id"`                    // Client's WLAN ID associated with this stream (YANG: IOS-XE 17.12.1)
	Mgid           int                `json:"mgid"`                       // MGID assigned to stream (YANG: IOS-XE 17.12.1)
	Priority       int                `json:"priority"`                   // Stream's priority (YANG: IOS-XE 17.12.1)
	RerrcEnable    bool               `json:"rerrc-enable"`               // Is Re-RRC enable on this stream (YANG: IOS-XE 17.12.1)
	RerrcDrop      bool               `json:"rerrc-drop"`                 // Re-RRC drop decision for violation (YANG: IOS-XE 17.12.1)
	Decision       string             `json:"decision"`                   // RRC decision for stream (YANG: IOS-XE 17.12.1)
	Qos            string             `json:"qos"`                        // Stream's QOS (YANG: IOS-XE 17.12.1)
	KbpsBandwidth  int                `json:"kbps-bandwidth"`             // Kbps bandwidth of stream (YANG: IOS-XE 17.12.1)
	Radio          string             `json:"radio"`                      // Radio on which Multicast group is reported (YANG: IOS-XE 17.12.1)
	StreamName     string             `json:"stream-name"`                // Name of the stream (YANG: IOS-XE 17.12.1)
	ApName         string             `json:"ap-name"`                    // AP name on which Multicast group is reported (YANG: IOS-XE 17.12.1)
	StartTime      string             `json:"start-time"`                 // Stream start time (YANG: IOS-XE 17.12.1)
	LastUpdated    string             `json:"last-updated"`               // Latest timestamp when Stream is updated by RRC (YANG: IOS-XE 17.12.1)
	RrcRadioRecord *RrcRadioRecordKey `json:"rrc-radio-record,omitempty"` // RRC stream radio record (YANG: IOS-XE 17.12.1)
}

// RrcRadioRecordKey represents RRC radio record key.
type RrcRadioRecordKey struct {
	ApMAC  string `json:"ap-mac"`  // AP MAC on which Multicast group is reported (YANG: IOS-XE 17.12.1)
	SlotID int    `json:"slot-id"` // Radio Slot ID (YANG: IOS-XE 17.12.1)
}

// RrcStreamAdmitRecord represents RRC stream admit record.
type RrcStreamAdmitRecord struct {
	LastUpdated string `json:"last-updated"` // Stream last updated by RRC (YANG: IOS-XE 17.12.1)
	ClientMAC   string `json:"client-mac"`   // Client MAC address (YANG: IOS-XE 17.12.1)
	DestIP      string `json:"dest-ip"`      // Multicast group destination IP address (YANG: IOS-XE 17.12.1)
}

// RrcStreamDenyRecord represents RRC stream deny record.
type RrcStreamDenyRecord struct {
	LastUpdated string `json:"last-updated"` // Stream last updated by RRC (YANG: IOS-XE 17.12.1)
	ClientMAC   string `json:"client-mac"`   // Client MAC address (YANG: IOS-XE 17.12.1)
	DestIP      string `json:"dest-ip"`      // Multicast group destination IP address (YANG: IOS-XE 17.12.1)
}
