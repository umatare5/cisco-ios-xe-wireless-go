// Package mobility provides data models for mobility operational data.
package mobility

// MobilityOper represents the root mobility operational data container.
type MobilityOper struct {
	MobilityOperData MobilitySystemOperData `json:"Cisco-IOS-XE-wireless-mobility-oper:mobility-oper-data"` // Mobility operational data
}

// MobilityOperApCache represents the AP cache data from WNC 17.12.5.
type MobilityOperApCache struct {
	ApCache []ApCache `json:"Cisco-IOS-XE-wireless-mobility-oper:ap-cache"`
}

// MobilityOperApPeerList represents the AP peer list data from WNC 17.12.5.
type MobilityOperApPeerList struct {
	ApPeerList []ApPeerList `json:"Cisco-IOS-XE-wireless-mobility-oper:ap-peer-list"`
}

// MobilityOperMmGlobalData represents the structure for MM global data from WNC 17.12.5.
type MobilityOperMmGlobalData struct {
	MmGlobalData MmGlobalData `json:"Cisco-IOS-XE-wireless-mobility-oper:mm-global-data"`
}

// MobilityOperMmIfGlobalMsgStats represents the structure for MM interface global message statistics from WNC 17.12.5.
type MobilityOperMmIfGlobalMsgStats struct {
	MmIfGlobalMsgStats MmIfGlobalMsgStats `json:"Cisco-IOS-XE-wireless-mobility-oper:mm-if-global-msg-stats"`
}

// MobilityOperMmIfGlobalStats represents the structure for MM interface global statistics from WNC 17.12.5.
type MobilityOperMmIfGlobalStats struct {
	MmIfGlobalStats MmIfGlobalStats `json:"Cisco-IOS-XE-wireless-mobility-oper:mm-if-global-stats"`
}

// MobilityOperMobilityClientData represents the structure for mobility client data.
type MobilityOperMobilityClientData struct {
	MobilityClientData []MobilityClientData `json:"Cisco-IOS-XE-wireless-mobility-oper:mobility-client-data"`
}

// MobilityOperMobilityClientStats represents the structure for mobility client statistics.
type MobilityOperMobilityClientStats struct {
	MobilityClientStats MobilityClientStats `json:"Cisco-IOS-XE-wireless-mobility-oper:mobility-client-stats"`
}

// MobilityOperMobilityGlobalDtlsStats represents the structure for mobility global DTLS statistics.
type MobilityOperMobilityGlobalDtlsStats struct {
	MobilityGlobalDtlsStats MobilityGlobalDtlsStats `json:"Cisco-IOS-XE-wireless-mobility-oper:mobility-global-dtls-stats"`
}

// MobilityOperMobilityGlobalMsgStats represents the structure for mobility global message statistics.
type MobilityOperMobilityGlobalMsgStats struct {
	MobilityGlobalMsgStats MobilityGlobalMsgStats `json:"Cisco-IOS-XE-wireless-mobility-oper:mobility-global-msg-stats"`
}

// MobilityOperMobilityGlobalStats represents the structure for mobility global statistics.
type MobilityOperMobilityGlobalStats struct {
	MobilityGlobalStats MobilityGlobalStats `json:"Cisco-IOS-XE-wireless-mobility-oper:mobility-global-stats"`
}

// MobilityOperWlanClientLimit represents the structure for WLAN client limit data.
type MobilityOperWlanClientLimit struct {
	WlanClientLimit []WlanClientLimit `json:"Cisco-IOS-XE-wireless-mobility-oper:wlan-client-limit"`
}

// MobilitySystemOperData represents mobility system operational data container.
type MobilitySystemOperData struct {
	ApCache                 []ApCache               `json:"ap-cache"`
	ApPeerList              []ApPeerList            `json:"ap-peer-list"`
	MmGlobalData            MmGlobalData            `json:"mm-global-data"`
	MmIfGlobalMsgStats      MmIfGlobalMsgStats      `json:"mm-if-global-msg-stats"`
	MmIfGlobalStats         MmIfGlobalStats         `json:"mm-if-global-stats"`
	MobilityClientData      []MobilityClientData    `json:"mobility-client-data"`
	MobilityClientStats     []MobilityClientStats   `json:"mobility-client-stats"`
	MobilityGlobalDtlsStats MobilityGlobalDtlsStats `json:"mobility-global-dtls-stats"`
	MobilityGlobalMsgStats  MobilityGlobalMsgStats  `json:"mobility-global-msg-stats"`
	MobilityGlobalStats     MobilityGlobalStats     `json:"mobility-global-stats"`
	WlanClientLimit         []WlanClientLimit       `json:"wlan-client-limit"`
}

// ApCache represents access point cache information.
type ApCache struct {
	MacAddr      string `json:"mac-addr"`      // Access point MAC address
	Name         string `json:"name"`          // Access point hostname
	ControllerIP string `json:"controller-ip"` // Controller IP address
	Source       string `json:"source"`        // Cache entry source
}

// ApPeerList represents access point peer list information.
type ApPeerList struct {
	PeerIP  string `json:"peer-ip"`  // Peer controller IP address
	ApCount int    `json:"ap-count"` // Number of access points
	Source  string `json:"source"`   // Peer entry source
}

// MmGlobalData represents mobility manager global configuration data.
type MmGlobalData struct {
	MmMacAddr string `json:"mm-mac-addr"` // Mobility manager MAC address
}

// MmIfGlobalMsgStats represents mobility manager interface global message statistics.
type MmIfGlobalMsgStats struct {
	IpcStats []IpcStats `json:"ipc-stats"`
}

// IpcStats represents inter-process communication statistics.
type IpcStats struct {
	Type      int    `json:"type"`        // Message type identifier
	Allocs    int    `json:"allocs"`      // Memory allocation count
	Frees     int    `json:"frees"`       // Memory free count
	TX        int    `json:"tx"`          // Transmitted message count
	RX        int    `json:"rx"`          // Received message count
	Forwarded int    `json:"forwarded"`   // Forwarded message count
	TXErrors  int    `json:"tx-errors"`   // Transmission error count
	RXErrors  int    `json:"rx-errors"`   // Reception error count
	TXRetries int    `json:"tx-retries"`  // Transmission retry count
	Drops     int    `json:"drops"`       // Dropped message count
	Built     int    `json:"built"`       // Built message count
	Processed int    `json:"processed"`   // Processed message count
	MmMsgType string `json:"mm-msg-type"` // Mobility message type
}

// MmIfGlobalStats represents mobility manager interface global statistics.
type MmIfGlobalStats struct {
	MbltyStats      MbltyStats      `json:"mblty-stats"`
	MbltyDomainInfo MbltyDomainInfo `json:"mblty-domain-info"`
}

// MbltyDomainInfo represents mobility domain information.
type MbltyDomainInfo struct {
	MobilityDomainID int `json:"mobility-domain-id"` // Mobility domain identifier
}

// MbltyStats represents mobility statistics.
type MbltyStats struct {
	EventDataAllocs            int `json:"event-data-allocs"`             // Event data allocation count
	EventDataFrees             int `json:"event-data-frees"`              // Event data free count
	MmifFsmInvalidEvents       int `json:"mmif-fsm-invalid-events"`       // FSM invalid event count
	MmifScheduleErrors         int `json:"mmif-schedule-errors"`          // Schedule error count
	MmifFsmFailure             int `json:"mmif-fsm-failure"`              // FSM failure count
	MmifIpcFailure             int `json:"mmif-ipc-failure"`              // IPC failure count
	MmifDBFailure              int `json:"mmif-db-failure"`               // Database failure count
	MmifInvalidParamsFailure   int `json:"mmif-invalid-params-failure"`   // Invalid parameter failure count
	MmifMmMsgDecodeFailure     int `json:"mmif-mm-msg-decode-failure"`    // Message decode failure count
	MmifUnknownFailure         int `json:"mmif-unknown-failure"`          // Unknown failure count
	MmifClientHandoffFailure   int `json:"mmif-client-handoff-failure"`   // Client handoff failure count
	MmifClientHandoffSuccess   int `json:"mmif-client-handoff-success"`   // Client handoff success count
	MmifAnchorDeny             int `json:"mmif-anchor-deny"`              // Anchor deny count
	MmifRemoteDelete           int `json:"mmif-remote-delete"`            // Remote delete count
	MmifTunnelDownDelete       int `json:"mmif-tunnel-down-delete"`       // Tunnel down delete count
	MmifMbssidDownEvent        int `json:"mmif-mbssid-down-event"`        // MBSSID down event count
	IntraWncdRoamCount         int `json:"intra-wncd-roam-count"`         // Intra-WNC roam count
	RemoteInterCtrlrRoams      int `json:"remote-inter-ctrlr-roams"`      // Remote inter-controller roam count
	RemoteWebauthPendRoams     int `json:"remote-webauth-pend-roams"`     // Remote webauth pending roam count
	AnchorRequestSent          int `json:"anchor-request-sent"`           // Anchor request sent count
	AnchorRequestGrantReceived int `json:"anchor-request-grant-received"` // Anchor request grant received count
	// Live WNC fields not in original struct
	AnchorRequestDenyReceived     int `json:"anchor-request-deny-received"`      // Anchor request deny received count
	AnchorRequestDenySent         int `json:"anchor-request-deny-sent"`          // Anchor request deny sent count
	AnchorRequestGrantSent        int `json:"anchor-request-grant-sent"`         // Anchor request grant sent count
	AnchorRequestReceived         int `json:"anchor-request-received"`           // Anchor request received count
	HandoffReceivedDeny           int `json:"handoff-received-deny"`             // Handoff received deny count
	HandoffReceivedGrpMismatch    int `json:"handoff-received-grp-mismatch"`     // Handoff received group mismatch count
	HandoffReceivedL3VlanOverride int `json:"handoff-received-l3-vlan-override"` // Handoff received L3 VLAN override count
	HandoffReceivedMsSsid         int `json:"handoff-received-ms-ssid"`          // Handoff received MS SSID count
	HandoffReceivedMsUnknown      int `json:"handoff-received-ms-unknown"`       // Handoff received MS unknown count
	HandoffReceivedOk             int `json:"handoff-received-ok"`               // Handoff received OK count
	HandoffReceivedUnknownPeer    int `json:"handoff-received-unknown-peer"`     // Handoff received unknown peer count
	HandoffSentDeny               int `json:"handoff-sent-deny"`                 // Handoff sent deny count
	HandoffSentGrpMismatch        int `json:"handoff-sent-grp-mismatch"`         // Handoff sent group mismatch count
	HandoffSentL3VlanOverride     int `json:"handoff-sent-l3-vlan-override"`     // Handoff sent L3 VLAN override count
	HandoffSentMsSsid             int `json:"handoff-sent-ms-ssid"`              // Handoff sent MS SSID count
	HandoffSentMsUnknown          int `json:"handoff-sent-ms-unknown"`           // Handoff sent MS unknown count
	HandoffSentOk                 int `json:"handoff-sent-ok"`                   // Handoff sent OK count
}

// MobilityClientData represents mobility client data.
type MobilityClientData struct {
	ClientMac           string `json:"client-mac"`             // Client MAC address
	ClientIfid          int64  `json:"client-ifid"`            // Client interface identifier
	OwnerInstance       int    `json:"owner-instance"`         // Client owner instance
	AssocTimeStamp      string `json:"assoc-time-stamp"`       // Client association timestamp
	MobilityState       string `json:"mobility-state"`         // Mobility state description
	ClientRole          string `json:"client-role"`            // Client role type
	ClientType          string `json:"client-type"`            // Client connection type
	ClientMode          string `json:"client-mode"`            // Client operational mode
	ClientRoamType      string `json:"client-roam-type"`       // Client roaming type
	PeerIP              string `json:"peer-ip"`                // Peer controller IP address
	EntryLastUpdateTime string `json:"entry-last-update-time"` // Entry last update timestamp
	ClientCreateTime    string `json:"client-create-time"`     // Client creation timestamp
	WlanProfile         string `json:"wlan-profile"`           // WLAN profile name
}

// MobilityClientStats represents mobility client statistics from WNC 17.12.5.
type MobilityClientStats struct {
	MmMbltyStats MmMbltyStats `json:"mm-mblty-stats"`
	IpcStats     []IpcStats   `json:"ipc-stats"`
	DgramStats   []IpcStats   `json:"dgram-stats"`
}

// DgramStats represents datagram statistics container.
type DgramStats struct {
	DgramStatsDbgCounters DgramStatsDbgCounters `json:"dgram-stats-dbg-counters"`
}

// DgramStatsDbgCounters represents datagram debug counters.
type DgramStatsDbgCounters struct {
	DgramTxPkts          int `json:"dgram-tx-pkts"`           // Datagram transmitted packets count
	DgramRxPkts          int `json:"dgram-rx-pkts"`           // Datagram received packets count
	DgramDiscards        int `json:"dgram-discards"`          // Datagram discarded packets count
	DgramSwitchTxTimeout int `json:"dgram-switch-tx-timeout"` // Datagram switch transmission timeout count
}

// MobilityGlobalDtlsStats represents mobility global DTLS statistics.
type MobilityGlobalDtlsStats struct {
	EventStats []DtlsEventStats `json:"event-stats"` // DTLS event statistics
	MsgStats   []DtlsMsgStats   `json:"msg-stats"`   // DTLS message statistics
}

// DtlsEventStats represents DTLS event statistics.
type DtlsEventStats struct {
	ConnectStart       int    `json:"connect-start"`       // Connection start count
	ConnectEstablished int    `json:"connect-established"` // Connection established count
	Close              int    `json:"close"`               // Connection close count
	KeyPlumbStart      int    `json:"key-plumb-start"`     // Key plumb start count
	KeyPlumbAcked      int    `json:"key-plumb-acked"`     // Key plumb acknowledged count
	TunnelType         string `json:"tunnel-type"`         // Tunnel type identifier
}

// DtlsMsgStats represents DTLS message statistics.
type DtlsMsgStats struct {
	HandshakeMsgTX int    `json:"handshake-msg-tx"` // Handshake message transmission count
	HandshakeMsgRX int    `json:"handshake-msg-rx"` // Handshake message reception count
	EncryptedMsgTX int    `json:"encrypted-msg-tx"` // Encrypted message transmission count
	EncryptedMsgRX int    `json:"encrypted-msg-rx"` // Encrypted message reception count
	TunnelType     string `json:"tunnel-type"`      // Tunnel type identifier
}

// MobilityGlobalMsgStats represents mobility global message statistics container.
type MobilityGlobalMsgStats struct {
	MsgStats MsgStats `json:"msg-stats"`
}

// MsgStats represents message statistics.
type MsgStats struct {
	MobilityAnnounceSent           int `json:"mobility-announce-sent"`            // Mobility announce messages sent count
	MobilityAnnounceReceived       int `json:"mobility-announce-received"`        // Mobility announce messages received count
	MobilityDenylistAddSent        int `json:"mobility-denylist-add-sent"`        // Mobility deny list add messages sent count
	MobilityDenylistAddReceived    int `json:"mobility-denylist-add-received"`    // Mobility deny list add messages received count
	MobilityDenylistDelSent        int `json:"mobility-denylist-del-sent"`        // Mobility deny list delete messages sent count
	MobilityDenylistDelReceived    int `json:"mobility-denylist-del-received"`    // Mobility deny list delete messages received count
	MobilityHandoffRequestSent     int `json:"mobility-handoff-request-sent"`     // Mobility handoff request messages sent count
	MobilityHandoffRequestReceived int `json:"mobility-handoff-request-received"` // Mobility handoff request messages received count
	MobilityHandoffReplySent       int `json:"mobility-handoff-reply-sent"`       // Mobility handoff reply messages sent count
	MobilityHandoffReplyReceived   int `json:"mobility-handoff-reply-received"`   // Mobility handoff reply messages received count
	MobilityHandoffEndSent         int `json:"mobility-handoff-end-sent"`         // Mobility handoff end messages sent count
	MobilityHandoffEndReceived     int `json:"mobility-handoff-end-received"`     // Mobility handoff end messages received count
	MobilityRevokeSent             int `json:"mobility-revoke-sent"`              // Mobility revoke messages sent count
	MobilityRevokeReceived         int `json:"mobility-revoke-received"`          // Mobility revoke messages received count
	MobilityRevokeAckSent          int `json:"mobility-revoke-ack-sent"`          // Mobility revoke acknowledgment messages sent count
	MobilityRevokeAckReceived      int `json:"mobility-revoke-ack-received"`      // Mobility revoke acknowledgment messages received count
	MobilityDirectiveAddSent       int `json:"mobility-directive-add-sent"`       // Mobility directive add messages sent count
	MobilityDirectiveAddReceived   int `json:"mobility-directive-add-received"`   // Mobility directive add messages received count
	MobilityDirectiveDelSent       int `json:"mobility-directive-del-sent"`       // Mobility directive delete messages sent count
	MobilityDirectiveDelReceived   int `json:"mobility-directive-del-received"`   // Mobility directive delete messages received count
	MobilityWlanStatusSent         int `json:"mobility-wlan-status-sent"`         // Mobility WLAN status messages sent count
	MobilityWlanStatusReceived     int `json:"mobility-wlan-status-received"`     // Mobility WLAN status messages received count
}

// MobilityGlobalStats represents mobility global statistics.
type MobilityGlobalStats struct {
	MmMbltyStats         MmMbltyStats `json:"mm-mblty-stats"`          // Mobility manager statistics
	NumOfSleepingClients int          `json:"num-of-sleeping-clients"` // Number of sleeping clients
}

// MmMbltyStats represents mobility manager statistics.
type MmMbltyStats struct {
	EventDataAllocs                     int `json:"event-data-allocs"`                        // Event data allocation count
	EventDataFrees                      int `json:"event-data-frees"`                         // Event data free count
	FsmSetAllocs                        int `json:"fsm-set-allocs"`                           // FSM set allocation count
	FsmSetFrees                         int `json:"fsm-set-frees"`                            // FSM set free count
	TimerAllocs                         int `json:"timer-allocs"`                             // Timer allocation count
	TimerFrees                          int `json:"timer-frees"`                              // Timer free count
	TimerStarts                         int `json:"timer-starts"`                             // Timer start count
	TimerStops                          int `json:"timer-stops"`                              // Timer stop count
	McfsmInvalidEvents                  int `json:"mcfsm-invalid-events"`                     // MC FSM invalid event count
	McfsmInternalError                  int `json:"mcfsm-internal-error"`                     // MC FSM internal error count
	JoinedAsLocal                       int `json:"joined-as-local"`                          // Client joined as local count
	JoinedAsForeign                     int `json:"joined-as-foreign"`                        // Client joined as foreign count
	JoinedAsExportForeign               int `json:"joined-as-export-foreign"`                 // Client joined as export foreign count
	JoinedAsExportAnchor                int `json:"joined-as-export-anchor"`                  // Client joined as export anchor count
	LocalToAnchor                       int `json:"local-to-anchor"`                          // Local to anchor transition count
	AnchorToLocal                       int `json:"anchor-to-local"`                          // Anchor to local transition count
	LocalDelete                         int `json:"local-delete"`                             // Local client delete count
	RemoteDelete                        int `json:"remote-delete"`                            // Remote client delete count
	McfsmDeleteInternalError            int `json:"mcfsm-delete-internal-error"`              // MC FSM delete internal error count
	McfsmRoamInternalError              int `json:"mcfsm-roam-internal-error"`                // MC FSM roam internal error count
	L2RoamCount                         int `json:"l2-roam-count"`                            // Layer 2 roam count
	L3RoamCount                         int `json:"l3-roam-count"`                            // Layer 3 roam count
	FlexClientRoamingCount              int `json:"flex-client-roaming-count"`                // FlexConnect client roaming count
	InterWncdRoamCount                  int `json:"inter-wncd-roam-count"`                    // Inter-WNC roam count
	ExpAncReqSent                       int `json:"exp-anc-req-sent"`                         // Export anchor request sent count
	ExpAncReqReceived                   int `json:"exp-anc-req-received"`                     // Export anchor request received count
	ExpAncRespOkSent                    int `json:"exp-anc-resp-ok-sent"`                     // Export anchor response OK sent count
	ExpAncRespGenericDenySent           int `json:"exp-anc-resp-generic-deny-sent"`           // Export anchor response generic deny sent count
	ExpAncRespClientBlacklistedSent     int `json:"exp-anc-resp-client-blacklisted-sent"`     // Export anchor response client blacklisted sent count
	ExpAncRespLimitReachedSent          int `json:"exp-anc-resp-limit-reached-sent"`          // Export anchor response limit reached sent count
	ExpAncRespProfileMismatchSent       int `json:"exp-anc-resp-profile-mismatch-sent"`       // Export anchor response profile mismatch sent count
	ExpAncRespOkReceived                int `json:"exp-anc-resp-ok-received"`                 // Export anchor response OK received count
	ExpAncRespGenericDenyReceived       int `json:"exp-anc-resp-generic-deny-received"`       // Export anchor response generic deny received count
	ExpAncRespClientBlacklistedReceived int `json:"exp-anc-resp-client-blacklisted-received"` // Export anchor response client blacklisted received count
	ExpAncRespLimitReachedReceived      int `json:"exp-anc-resp-limit-reached-received"`      // Export anchor response limit reached received count
	ExpAncRespProfileMismatchReceived   int `json:"exp-anc-resp-profile-mismatch-received"`   // Export anchor response profile mismatch received count
	ExpAncRespUnknownReceived           int `json:"exp-anc-resp-unknown-received"`            // Export anchor response unknown received count
	HandoffSentMsBlacklist              int `json:"handoff-sent-ms-blacklist"`                // Handoff sent MS blacklist count
	HandoffReceivedMsBlacklist          int `json:"handoff-received-ms-blacklist"`            // Handoff received MS blacklist count
}

// WlanClientLimit represents WLAN client limit information.
type WlanClientLimit struct {
	WlanProfile      string `json:"wlan-profile"`       // WLAN profile name
	CurrClientsCount int    `json:"curr-clients-count"` // Current client count
}
