// Package mcast provides data models for multicast operational data.
package mcast

// McastOper represents multicast operational data from controller.
type McastOper struct {
	CiscoIOSXEWirelessMcastOperMcastOperData struct {
		FlexMediastreamClientSummary   []FlexMediastreamClientSummary   `json:"flex-mediastream-client-summary,omitempty"`
		VlanL2MgidOp                   []VlanL2MgidOp                   `json:"vlan-l2-mgid-op,omitempty"`
		FabricMediaStreamClientSummary []FabricMediaStreamClientSummary `json:"fabric-media-stream-client-summary,omitempty"`
		McastMgidInfo                  []McastMgidInfo                  `json:"mcast-mgid-info,omitempty"`
		MulticastOperData              []MulticastOperData              `json:"multicast-oper-data,omitempty"`
		RrcHistoryClientRecordData     []RrcHistoryClientRecordData     `json:"rrc-history-client-record-data,omitempty"`
		RrcSrRadioRecord               []RrcSrRadioRecord               `json:"rrc-sr-radio-record,omitempty"`
		RrcStreamRecord                []RrcStreamRecord                `json:"rrc-stream-record,omitempty"`
		RrcStreamAdmitRecord           []RrcStreamAdmitRecord           `json:"rrc-stream-admit-record,omitempty"`
		RrcStreamDenyRecord            []RrcStreamDenyRecord            `json:"rrc-stream-deny-record,omitempty"`
	} `json:"Cisco-IOS-XE-wireless-mcast-oper:mcast-oper-data"`
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
	ClientMac            string                 `json:"client-mac"`              // Client MAC address
	VlanID               int                    `json:"vlan-id"`                 // VLAN identifier
	FlexMcastClientGroup []FlexMcastClientGroup `json:"flex-mcast-client-group"` // FlexConnect multicast client groups
}

// FlexMcastClientGroup represents FlexConnect multicast client group configuration.
type FlexMcastClientGroup struct {
	McastIP    string `json:"mcast-ip"`    // Multicast IP address
	StreamName string `json:"stream-name"` // Stream name identifier
	ApMac      string `json:"ap-mac"`      // Access point MAC address
	IsDirect   bool   `json:"is-direct"`   // Direct connection flag
}

// VlanL2MgidOp represents VLAN Layer 2 multicast group ID operational data.
type VlanL2MgidOp struct {
	VlanIndex               int  `json:"vlan-index"`                 // VLAN index number
	IsNonipMulticastEnabled bool `json:"is-nonip-multicast-enabled"` // Non-IP multicast enable status
	IsBroadcastEnable       bool `json:"is-broadcast-enable"`        // Broadcast enable status
}

// FabricMediaStreamClientSummary represents fabric mediastream client summary.
type FabricMediaStreamClientSummary struct {
	ClientMac              string                   `json:"client-mac"`                          // Client MAC address (YANG: IOS-XE 17.12.1+)
	VlanID                 int                      `json:"vlan-id"`                             // VLAN identifier (YANG: IOS-XE 17.12.1+)
	FabricMcastClientGroup []FabricMcastClientGroup `json:"fabric-mcast-client-group,omitempty"` // Fabric multicast client groups (YANG: IOS-XE 17.12.1+)
}

// FabricMcastClientGroup represents fabric multicast client group.
type FabricMcastClientGroup struct {
	McastIP    string `json:"mcast-ip"`    // Multicast IP address (YANG: IOS-XE 17.12.1+)
	StreamName string `json:"stream-name"` // Stream name identifier (YANG: IOS-XE 17.12.1+)
	ApMac      string `json:"ap-mac"`      // Access point MAC address (YANG: IOS-XE 17.12.1+)
	IsDirect   bool   `json:"is-direct"`   // Direct connection flag (YANG: IOS-XE 17.12.1+)
}

// McastMgidInfo represents multicast MGID information.
type McastMgidInfo struct {
	Mgid                   int                    `json:"mgid"`                             // Multicast group ID (YANG: IOS-XE 17.12.1+)
	Vlan                   int                    `json:"vlan"`                             // VLAN identifier (YANG: IOS-XE 17.12.1+)
	McOnlyNumClients       int                    `json:"mc-only-num-clients"`              // Multicast-only client count (YANG: IOS-XE 17.12.1+)
	Mc2ucNumClients        int                    `json:"mc2uc-num-clients"`                // Multicast-to-unicast client count (YANG: IOS-XE 17.12.1+)
	Mc2ucNumDenyClients    int                    `json:"mc2uc-num-deny-clients"`           // Multicast-to-unicast denied client count (YANG: IOS-XE 17.12.1+)
	Mc2ucNumPendingClients int                    `json:"mc2uc-num-pending-clients"`        // Multicast-to-unicast pending client count (YANG: IOS-XE 17.12.1+)
	Group                  string                 `json:"group"`                            // Multicast group address (YANG: IOS-XE 17.12.1+)
	McastMgidClientList    []McastMgidClientEntry `json:"mcast-mgid-client-list,omitempty"` // Multicast MGID client list (YANG: IOS-XE 17.12.1+)
}

// McastMgidClientEntry represents multicast MGID client entry.
type McastMgidClientEntry struct {
	ClientMacAddr string `json:"client-mac-addr"` // Client MAC address (YANG: IOS-XE 17.12.1+)
	ClientIPAddr  string `json:"client-ip-addr"`  // Client IP address (YANG: IOS-XE 17.12.1+)
	ClientStatus  string `json:"client-status"`   // Client status (YANG: IOS-XE 17.12.1+)
}

// MulticastOperData represents multicast operational data.
type MulticastOperData struct {
	MsMacAddress string             `json:"ms-mac-address"`          // Mobility switch MAC address (YANG: IOS-XE 17.12.1+)
	NumEntries   int                `json:"num-entries"`             // Number of entries (YANG: IOS-XE 17.12.1+)
	Entry        []McastClientEntry `json:"entry,omitempty"`         // Multicast client entries (YANG: IOS-XE 17.12.1+)
	ClientIPv6   string             `json:"client-ipv6,omitempty"`   // Client IPv6 address (YANG: IOS-XE 17.12.1+)
	CapwapIifID  int                `json:"capwap-iif-id,omitempty"` // CAPWAP interface ID (YANG: IOS-XE 17.12.1+)
	ClientIP     string             `json:"client-ip,omitempty"`     // Client IP address (YANG: IOS-XE 17.12.1+)
}

// McastClientEntry represents multicast client entry.
type McastClientEntry struct {
	Vlan         int    `json:"vlan"`          // VLAN identifier (YANG: IOS-XE 17.12.1+)
	Mgid         int    `json:"mgid"`          // Multicast group ID (YANG: IOS-XE 17.12.1+)
	Group        string `json:"group"`         // Multicast group address (YANG: IOS-XE 17.12.1+)
	ClientStatus string `json:"client-status"` // Client status (YANG: IOS-XE 17.12.1+)
	Qos          string `json:"qos"`           // Quality of service (YANG: IOS-XE 17.12.1+)
	Used         bool   `json:"used"`          // Entry usage flag (YANG: IOS-XE 17.12.1+)
}

// RrcHistoryClientRecordData represents RRC history client record data.
type RrcHistoryClientRecordData struct {
	UserTimeStamp        string `json:"user-time-stamp"`         // User timestamp (YANG: IOS-XE 17.12.1+)
	ClientMac            string `json:"client-mac"`              // Client MAC address (YANG: IOS-XE 17.12.1+)
	Qos                  int    `json:"qos"`                     // Quality of service (YANG: IOS-XE 17.12.1+)
	Decision             string `json:"decision"`                // RRC decision (YANG: IOS-XE 17.12.1+)
	ReasonCode           int    `json:"reason-code"`             // Reason code (YANG: IOS-XE 17.12.1+)
	ApMac                string `json:"ap-mac"`                  // Access point MAC address (YANG: IOS-XE 17.12.1+)
	VapID                int    `json:"vap-id"`                  // Virtual AP ID (YANG: IOS-XE 17.12.1+)
	SlotID               int    `json:"slot-id"`                 // Slot ID (YANG: IOS-XE 17.12.1+)
	CacEnabled           int    `json:"cac-enabled"`             // CAC enabled flag (YANG: IOS-XE 17.12.1+)
	StreamName           string `json:"stream-name"`             // Stream name (YANG: IOS-XE 17.12.1+)
	DstIPAddress         string `json:"dst-ip-address"`          // Destination IP address (YANG: IOS-XE 17.12.1+)
	CfgStreamBw          int    `json:"cfg-stream-bw"`           // Configured stream bandwidth (YANG: IOS-XE 17.12.1+)
	CurrentRate          int    `json:"current-rate"`            // Current rate (YANG: IOS-XE 17.12.1+)
	VideoPktSize         int    `json:"video-pkt-size"`          // Video packet size (YANG: IOS-XE 17.12.1+)
	CurrVideoUtil        int    `json:"curr-video-util"`         // Current video utilization (YANG: IOS-XE 17.12.1+)
	CurrVoiceUtil        int    `json:"curr-voice-util"`         // Current voice utilization (YANG: IOS-XE 17.12.1+)
	CurrChannelUtil      int    `json:"curr-channel-util"`       // Current channel utilization (YANG: IOS-XE 17.12.1+)
	CurrQueueUtil        int    `json:"curr-queue-util"`         // Current queue utilization (YANG: IOS-XE 17.12.1+)
	CurrVideoPps         int    `json:"curr-video-pps"`          // Current video packets per second (YANG: IOS-XE 17.12.1+)
	VideoDelayHistSevere int    `json:"video-delay-hist-severe"` // Video delay histogram severe (YANG: IOS-XE 17.12.1+)
	VideoPktLossDiscard  int    `json:"video-pkt-loss-discard"`  // Video packet loss discard (YANG: IOS-XE 17.12.1+)
	VideoPktLossFail     int    `json:"video-pkt-loss-fail"`     // Video packet loss fail (YANG: IOS-XE 17.12.1+)
	NumVideoStreams      int    `json:"num-video-streams"`       // Number of video streams (YANG: IOS-XE 17.12.1+)
}

// RrcSrRadioRecord represents RRC stream radio record.
type RrcSrRadioRecord struct {
	ApMac               string          `json:"ap-mac"`                        // Access point MAC address (YANG: IOS-XE 17.12.1+)
	SlotID              int             `json:"slot-id"`                       // Slot ID (YANG: IOS-XE 17.12.1+)
	RadioType           int             `json:"radio-type"`                    // Radio type (YANG: IOS-XE 17.12.1+)
	DuplicatedBandWidth int             `json:"duplicated-band-width"`         // Duplicated bandwidth (YANG: IOS-XE 17.12.1+)
	LastReRrc           string          `json:"last-re-rrc"`                   // Last re-RRC timestamp (YANG: IOS-XE 17.12.1+)
	NumberOfAdmitted    int             `json:"number-of-admitted"`            // Number of admitted streams (YANG: IOS-XE 17.12.1+)
	RrcGroupsInRadio    []RrcGroupsInfo `json:"rrc-groups-in-radio,omitempty"` // RRC groups in radio (YANG: IOS-XE 17.12.1+)
}

// RrcGroupsInfo represents RRC groups information.
type RrcGroupsInfo struct {
	DestIP      string `json:"dest-ip"`       // Destination IP address (YANG: IOS-XE 17.12.1+)
	NoOfStreams int    `json:"no-of-streams"` // Number of streams (YANG: IOS-XE 17.12.1+)
}

// RrcStreamRecord represents RRC stream record.
type RrcStreamRecord struct {
	StreamNameStr  string             `json:"stream-name-str"`            // Stream name string (YANG: IOS-XE 17.12.1+)
	GroupIP        string             `json:"group-ip"`                   // Group IP address (YANG: IOS-XE 17.12.1+)
	ClMac          string             `json:"cl-mac"`                     // Client MAC address (YANG: IOS-XE 17.12.1+)
	ClientMac      string             `json:"client-mac"`                 // Client MAC address (YANG: IOS-XE 17.12.1+)
	DestIP         string             `json:"dest-ip"`                    // Destination IP address (YANG: IOS-XE 17.12.1+)
	VapID          int                `json:"vap-id"`                     // Virtual AP ID (YANG: IOS-XE 17.12.1+)
	Vlan           int                `json:"vlan"`                       // VLAN identifier (YANG: IOS-XE 17.12.1+)
	WlanID         int                `json:"wlan-id"`                    // WLAN ID (YANG: IOS-XE 17.12.1+)
	Mgid           int                `json:"mgid"`                       // Multicast group ID (YANG: IOS-XE 17.12.1+)
	Priority       int                `json:"priority"`                   // Priority level (YANG: IOS-XE 17.12.1+)
	RerrcEnable    bool               `json:"rerrc-enable"`               // Re-RRC enable flag (YANG: IOS-XE 17.12.1+)
	RerrcDrop      bool               `json:"rerrc-drop"`                 // Re-RRC drop flag (YANG: IOS-XE 17.12.1+)
	Decision       string             `json:"decision"`                   // RRC decision (YANG: IOS-XE 17.12.1+)
	Qos            string             `json:"qos"`                        // Quality of service (YANG: IOS-XE 17.12.1+)
	KbpsBandwidth  int                `json:"kbps-bandwidth"`             // Bandwidth in kbps (YANG: IOS-XE 17.12.1+)
	Radio          string             `json:"radio"`                      // Radio identifier (YANG: IOS-XE 17.12.1+)
	StreamName     string             `json:"stream-name"`                // Stream name (YANG: IOS-XE 17.12.1+)
	ApName         string             `json:"ap-name"`                    // Access point name (YANG: IOS-XE 17.12.1+)
	StartTime      string             `json:"start-time"`                 // Start time (YANG: IOS-XE 17.12.1+)
	LastUpdated    string             `json:"last-updated"`               // Last updated timestamp (YANG: IOS-XE 17.12.1+)
	RrcRadioRecord *RrcRadioRecordKey `json:"rrc-radio-record,omitempty"` // RRC radio record (YANG: IOS-XE 17.12.1+)
}

// RrcRadioRecordKey represents RRC radio record key.
type RrcRadioRecordKey struct {
	ApMac  string `json:"ap-mac"`  // Access point MAC address (YANG: IOS-XE 17.12.1+)
	SlotID int    `json:"slot-id"` // Slot ID (YANG: IOS-XE 17.12.1+)
}

// RrcStreamAdmitRecord represents RRC stream admit record.
type RrcStreamAdmitRecord struct {
	LastUpdated string `json:"last-updated"` // Last updated timestamp (YANG: IOS-XE 17.12.1+)
	ClientMac   string `json:"client-mac"`   // Client MAC address (YANG: IOS-XE 17.12.1+)
	DestIP      string `json:"dest-ip"`      // Destination IP address (YANG: IOS-XE 17.12.1+)
}

// RrcStreamDenyRecord represents RRC stream deny record.
type RrcStreamDenyRecord struct {
	LastUpdated string `json:"last-updated"` // Last updated timestamp (YANG: IOS-XE 17.12.1+)
	ClientMac   string `json:"client-mac"`   // Client MAC address (YANG: IOS-XE 17.12.1+)
	DestIP      string `json:"dest-ip"`      // Destination IP address (YANG: IOS-XE 17.12.1+)
}
