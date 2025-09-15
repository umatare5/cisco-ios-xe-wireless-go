// Package mobility provides data models for mobility operational data.
package mobility

// MobilityOper represents the root mobility operational data container.
type MobilityOper struct {
	MobilityOperData struct {
		ApCache                 []ApCache               `json:"ap-cache"`                   // AP cache info by mobility controller (YANG: IOS-XE 17.12.1)
		ApPeerList              []ApPeerList            `json:"ap-peer-list"`               // AP count reported by peer controllers (YANG: IOS-XE 17.12.1)
		MmGlobalData            MmGlobalData            `json:"mm-global-data"`             // Container for global mobility data (YANG: IOS-XE 17.12.1)
		MmIfGlobalMsgStats      MmIfGlobalMsgStats      `json:"mm-if-global-msg-stats"`     // Global mobility interface message stats (YANG: IOS-XE 17.12.1)
		MmIfGlobalStats         MmIfGlobalStats         `json:"mm-if-global-stats"`         // Global mobility interface event stats (Live: IOS-XE 17.12.5)
		MobilityClientData      []MobilityClientData    `json:"mobility-client-data"`       // 802.11 LWAPP Mobility Clients info (YANG: IOS-XE 17.12.1)
		MobilityClientStats     []MobilityClientStats   `json:"mobility-client-stats"`      // Client mobility event and message stats (YANG: IOS-XE 17.12.1)
		MobilityGlobalDTLSStats MobilityGlobalDTLSStats `json:"mobility-global-dtls-stats"` // Global mobility DTLS tunnel stats (YANG: IOS-XE 17.12.1)
		MobilityGlobalMsgStats  MobilityGlobalMsgStats  `json:"mobility-global-msg-stats"`  // Global mobility message exchange stats (YANG: IOS-XE 17.12.1)
		MobilityGlobalStats     MobilityGlobalStats     `json:"mobility-global-stats"`      // Global mobility event and tunnel stats (YANG: IOS-XE 17.12.1)
		WlanClientLimit         []WlanClientLimit       `json:"wlan-client-limit"`          // WLAN client limit configuration data (YANG: IOS-XE 17.12.1)
	} `json:"Cisco-IOS-XE-wireless-mobility-oper:mobility-oper-data"` // Mobility operational data (YANG: IOS-XE 17.12.1)
}

// MobilityOperApCache represents the AP cache data.
type MobilityOperApCache struct {
	ApCache []ApCache `json:"Cisco-IOS-XE-wireless-mobility-oper:ap-cache"`
}

// MobilityOperApPeerList represents the AP peer list data.
type MobilityOperApPeerList struct {
	ApPeerList []ApPeerList `json:"Cisco-IOS-XE-wireless-mobility-oper:ap-peer-list"`
}

// MobilityOperMmGlobalData represents the structure for MM global data.
type MobilityOperMmGlobalData struct {
	MmGlobalData MmGlobalData `json:"Cisco-IOS-XE-wireless-mobility-oper:mm-global-data"`
}

// MobilityOperMmIfGlobalMsgStats represents the structure for MM interface global message statistics.
type MobilityOperMmIfGlobalMsgStats struct {
	MmIfGlobalMsgStats MmIfGlobalMsgStats `json:"Cisco-IOS-XE-wireless-mobility-oper:mm-if-global-msg-stats"`
}

// MobilityOperMmIfGlobalStats represents the structure for MM interface global statistics.
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

// MobilityOperMobilityGlobalDTLSStats represents the structure for mobility global DTLS statistics.
type MobilityOperMobilityGlobalDTLSStats struct {
	MobilityGlobalDTLSStats MobilityGlobalDTLSStats `json:"Cisco-IOS-XE-wireless-mobility-oper:mobility-global-dtls-stats"`
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

// ApCache represents access point cache information from mobility controllers.
type ApCache struct {
	ControllerIP string `json:"controller-ip,omitempty"` // Reporting device's IP address (YANG: IOS-XE 17.12.1)
}

// ApPeerList represents access point peer list information.
type ApPeerList struct {
	PeerIP  string `json:"peer-ip"`  // Reporting device's IP address (YANG: IOS-XE 17.12.1)
	ApCount int    `json:"ap-count"` // Total number of APs reported by this device (YANG: IOS-XE 17.12.1)
}

// MmGlobalData represents mobility manager global configuration data.
type MmGlobalData struct {
	MmMACAddr string `json:"mm-mac-addr"` // MAC address being used by mobility module (Live: IOS-XE 17.12.5)
}

// MmIfGlobalMsgStats represents mobility manager interface global message statistics.
type MmIfGlobalMsgStats struct {
	IpcStats []IpcStats `json:"ipc-stats"` // Inter-process communication stats (YANG: IOS-XE 17.12.1)
}

// IpcStats represents inter-process communication statistics.
type IpcStats struct {
	Type      int    `json:"type"`        // CAPWAP messages type for mobility client (Live: IOS-XE 17.12.5)
	Allocs    int    `json:"allocs"`      // Number of CAPWAP messages allocated for mobility client (Live: IOS-XE 17.12.5)
	Frees     int    `json:"frees"`       // Number of CAPWAP messages freed for mobility client (Live: IOS-XE 17.12.5)
	TX        int    `json:"tx"`          // Number of CAPWAP messages transmitted for mobility client (Live: IOS-XE 17.12.5)
	RX        int    `json:"rx"`          // Number of CAPWAP messages received for mobility client (Live: IOS-XE 17.12.5)
	Forwarded int    `json:"forwarded"`   // Number of CAPWAP messages forwarded for mobility client (Live: IOS-XE 17.12.5)
	TXErrors  int    `json:"tx-errors"`   // Number of CAPWAP message transmit errors for mobility client (Live: IOS-XE 17.12.5)
	RXErrors  int    `json:"rx-errors"`   // Number of CAPWAP message receive errors for mobility client (Live: IOS-XE 17.12.5)
	TXRetries int    `json:"tx-retries"`  // Number of retries for CAPWAP message transmit error for mobility client (Live: IOS-XE 17.12.5)
	Drops     int    `json:"drops"`       // Number of dropped CAPWAP messages for mobility client (Live: IOS-XE 17.12.5)
	Built     int    `json:"built"`       // Number of CAPWAP messages built for mobility client (Live: IOS-XE 17.12.5)
	Processed int    `json:"processed"`   // Number of processed CAPWAP messages for mobility client (Live: IOS-XE 17.12.5)
	MmMsgType string `json:"mm-msg-type"` // CAPWAP mobility message type (Live: IOS-XE 17.12.5)
}

// MmIfGlobalStats represents mobility manager interface global statistics.
type MmIfGlobalStats struct {
	MbltyStats      MbltyStats      `json:"mblty-stats"`       // Mobility statistics data (Live: IOS-XE 17.12.5)
	MbltyDomainInfo MbltyDomainInfo `json:"mblty-domain-info"` // Mobility domain information (Live: IOS-XE 17.12.5)
}

// MbltyDomainInfo represents mobility domain information.
type MbltyDomainInfo struct {
	MobilityDomainID int `json:"mobility-domain-id"` // Mobility domain identifier (Live: IOS-XE 17.12.5)
}

// MbltyStats represents mobility statistics.
type MbltyStats struct {
	EventDataAllocs               int `json:"event-data-allocs"`                 // Total number of mobility interface event data allocations (Live: IOS-XE 17.12.5)
	EventDataFrees                int `json:"event-data-frees"`                  // Total number of mobility interface event data frees (Live: IOS-XE 17.12.5)
	MmifFsmInvalidEvents          int `json:"mmif-fsm-invalid-events"`           // Total number of invalid events received by mobility interface (Live: IOS-XE 17.12.5)
	MmifScheduleErrors            int `json:"mmif-schedule-errors"`              // Total number of mobility interface event scheduling errors (Live: IOS-XE 17.12.5)
	MmifFsmFailure                int `json:"mmif-fsm-failure"`                  // FSM failure count (YANG: IOS-XE 17.12.1)
	MmifIpcFailure                int `json:"mmif-ipc-failure"`                  // Total number of mobility interface event processing errors due to IPC messaging failure (Live: IOS-XE 17.12.5)
	MmifDBFailure                 int `json:"mmif-db-failure"`                   // Total number of mobility interface event processing errors due to database operation failure (Live: IOS-XE 17.12.5)
	MmifInvalidParamsFailure      int `json:"mmif-invalid-params-failure"`       // Invalid parameter failure count (YANG: IOS-XE 17.12.1)
	MmifMmMsgDecodeFailure        int `json:"mmif-mm-msg-decode-failure"`        // Message decode failure count (YANG: IOS-XE 17.12.1)
	MmifUnknownFailure            int `json:"mmif-unknown-failure"`              // Unknown failure count (YANG: IOS-XE 17.12.1)
	MmifClientHandoffFailure      int `json:"mmif-client-handoff-failure"`       // Total number of client handoff failures for mobile stations on wireless LAN controller (Live: IOS-XE 17.12.5)
	MmifClientHandoffSuccess      int `json:"mmif-client-handoff-success"`       // Total number of client handoff successes for mobile stations on wireless LAN controller (Live: IOS-XE 17.12.5)
	MmifAnchorDeny                int `json:"mmif-anchor-deny"`                  // Anchor deny count (YANG: IOS-XE 17.12.1)
	MmifRemoteDelete              int `json:"mmif-remote-delete"`                // Remote delete count (YANG: IOS-XE 17.12.1)
	MmifTunnelDownDelete          int `json:"mmif-tunnel-down-delete"`           // Tunnel down delete count (YANG: IOS-XE 17.12.1)
	MmifMbssidDownEvent           int `json:"mmif-mbssid-down-event"`            // MBSSID down event count (YANG: IOS-XE 17.12.1)
	IntraWncdRoamCount            int `json:"intra-wncd-roam-count"`             // Total number of intra-process roams within wireless LAN controller (Live: IOS-XE 17.12.5)
	RemoteInterCtrlrRoams         int `json:"remote-inter-ctrlr-roams"`          // Total number of inter-controller roams performed on peer controllers by anchored clients (Live: IOS-XE 17.12.5)
	RemoteWebauthPendRoams        int `json:"remote-webauth-pend-roams"`         // Remote webauth pending roam count (YANG: IOS-XE 17.12.1)
	AnchorRequestSent             int `json:"anchor-request-sent"`               // Total number of anchor requests sent for mobile stations on wireless LAN controller (Live: IOS-XE 17.12.5)
	AnchorRequestGrantReceived    int `json:"anchor-request-grant-received"`     // Total number of anchor request grants received for mobile stations on wireless LAN controller (Live: IOS-XE 17.12.5)
	AnchorRequestDenyReceived     int `json:"anchor-request-deny-received"`      // Total number of anchor request denies received for mobile stations on wireless LAN controller (Live: IOS-XE 17.12.5)
	AnchorRequestDenySent         int `json:"anchor-request-deny-sent"`          // Anchor request deny sent count (Live: IOS-XE 17.12.5)
	AnchorRequestGrantSent        int `json:"anchor-request-grant-sent"`         // Anchor request grant sent count (Live: IOS-XE 17.12.5)
	AnchorRequestReceived         int `json:"anchor-request-received"`           // Anchor request received count (Live: IOS-XE 17.12.5)
	HandoffReceivedDeny           int `json:"handoff-received-deny"`             // Handoff received deny count (Live: IOS-XE 17.12.5)
	HandoffReceivedGrpMismatch    int `json:"handoff-received-grp-mismatch"`     // Handoff received group mismatch count (Live: IOS-XE 17.12.5)
	HandoffReceivedL3VlanOverride int `json:"handoff-received-l3-vlan-override"` // Handoff received L3 VLAN override count (Live: IOS-XE 17.12.5)
	HandoffReceivedMsSsid         int `json:"handoff-received-ms-ssid"`          // Handoff received MS SSID count (Live: IOS-XE 17.12.5)
	HandoffReceivedMsUnknown      int `json:"handoff-received-ms-unknown"`       // Handoff received MS unknown count (Live: IOS-XE 17.12.5)
	HandoffReceivedOk             int `json:"handoff-received-ok"`               // Total number of handoff status success received for mobile stations on wireless LAN controller (Live: IOS-XE 17.12.5)
	HandoffReceivedUnknownPeer    int `json:"handoff-received-unknown-peer"`     // Handoff received unknown peer count (Live: IOS-XE 17.12.5)
	HandoffSentDeny               int `json:"handoff-sent-deny"`                 // Handoff sent deny count (Live: IOS-XE 17.12.5)
	HandoffSentGrpMismatch        int `json:"handoff-sent-grp-mismatch"`         // Handoff sent group mismatch count (Live: IOS-XE 17.12.5)
	HandoffSentL3VlanOverride     int `json:"handoff-sent-l3-vlan-override"`     // Handoff sent L3 VLAN override count (Live: IOS-XE 17.12.5)
	HandoffSentMsSsid             int `json:"handoff-sent-ms-ssid"`              // Handoff sent MS SSID count (Live: IOS-XE 17.12.5)
	HandoffSentMsUnknown          int `json:"handoff-sent-ms-unknown"`           // Handoff sent MS unknown count (Live: IOS-XE 17.12.5)
	HandoffSentOk                 int `json:"handoff-sent-ok"`                   // Total number of handoff status OK sent for mobile stations on wireless LAN controller (Live: IOS-XE 17.12.5)
}

// MobilityClientData represents mobility client data.
type MobilityClientData struct {
	ClientMAC           string `json:"client-mac"`             // Client MAC address (Live: IOS-XE 17.12.5)
	ClientIfid          int64  `json:"client-ifid"`            // Client IFID (Live: IOS-XE 17.12.5)
	OwnerInstance       int    `json:"owner-instance"`         // Owner instance for mobility client (Live: IOS-XE 17.12.5)
	AssocTimeStamp      string `json:"assoc-time-stamp"`       // Client association timestamp (Live: IOS-XE 17.12.5)
	MobilityState       string `json:"mobility-state"`         // Mobility state (Live: IOS-XE 17.12.5)
	ClientRole          string `json:"client-role"`            // Client role (Live: IOS-XE 17.12.5)
	ClientType          string `json:"client-type"`            // Client type (Live: IOS-XE 17.12.5)
	ClientMode          string `json:"client-mode"`            // Mobility client mode (Live: IOS-XE 17.12.5)
	ClientRoamType      string `json:"client-roam-type"`       // Mobility client roam type (Live: IOS-XE 17.12.5)
	PeerIP              string `json:"peer-ip"`                // Mobility peer IP for anchor or foreign client (Live: IOS-XE 17.12.5)
	EntryLastUpdateTime string `json:"entry-last-update-time"` // Entry last update timestamp (Live: IOS-XE 17.12.5)
	ClientCreateTime    string `json:"client-create-time"`     // Client creation timestamp (Live: IOS-XE 17.12.5)
	WlanProfile         string `json:"wlan-profile"`           // Mobility client wlan profile name (Live: IOS-XE 17.12.5)
}

// MobilityClientStats represents mobility client statistics.
type MobilityClientStats struct {
	MmMbltyStats MmMbltyStats `json:"mm-mblty-stats"` // Mobility statistics data (YANG: IOS-XE 17.12.1)
	IpcStats     []IpcStats   `json:"ipc-stats"`      // Inter-process communication stats (YANG: IOS-XE 17.12.1)
	DgramStats   []IpcStats   `json:"dgram-stats"`    // Datagram statistics (YANG: IOS-XE 17.12.1)
}

// DgramStats represents datagram statistics container.
type DgramStats struct {
	DgramStatsDbgCounters DgramStatsDbgCounters `json:"dgram-stats-dbg-counters"` // Datagram debug counters (YANG: IOS-XE 17.12.1)
}

// DgramStatsDbgCounters represents datagram debug counters.
type DgramStatsDbgCounters struct {
	DgramTxPkts          int `json:"dgram-tx-pkts"`           // Datagram transmitted packets count (YANG: IOS-XE 17.12.1)
	DgramRxPkts          int `json:"dgram-rx-pkts"`           // Datagram received packets count (YANG: IOS-XE 17.12.1)
	DgramDiscards        int `json:"dgram-discards"`          // Datagram discarded packets count (YANG: IOS-XE 17.12.1)
	DgramSwitchTxTimeout int `json:"dgram-switch-tx-timeout"` // Datagram switch transmission timeout count (YANG: IOS-XE 17.12.1)
}

// MobilityGlobalDTLSStats represents mobility global DTLS statistics.
type MobilityGlobalDTLSStats struct {
	EventStats []DTLSEventStats `json:"event-stats"` // DTLS event statistics (YANG: IOS-XE 17.12.1)
	MsgStats   []DTLSMsgStats   `json:"msg-stats"`   // DTLS message statistics (YANG: IOS-XE 17.12.1)
}

// DTLSEventStats represents DTLS event statistics.
type DTLSEventStats struct {
	ConnectStart       int    `json:"connect-start"`       // Connections attempted (YANG: IOS-XE 17.12.1)
	ConnectEstablished int    `json:"connect-established"` // Connections established (YANG: IOS-XE 17.12.1)
	Close              int    `json:"close"`               // Connections closed (YANG: IOS-XE 17.12.1)
	KeyPlumbStart      int    `json:"key-plumb-start"`     // Data plane key plumb requests (YANG: IOS-XE 17.12.1)
	KeyPlumbAcked      int    `json:"key-plumb-acked"`     // Data plane key plumb acknowledgements (YANG: IOS-XE 17.12.1)
	TunnelType         string `json:"tunnel-type"`         // CAPWAP mobility tunnel type (YANG: IOS-XE 17.12.1)
}

// DTLSMsgStats represents DTLS message statistics.
type DTLSMsgStats struct {
	HandshakeMsgTX int    `json:"handshake-msg-tx"` // Handshake messages sent (YANG: IOS-XE 17.12.1)
	HandshakeMsgRX int    `json:"handshake-msg-rx"` // Handshake messages received (YANG: IOS-XE 17.12.1)
	EncryptedMsgTX int    `json:"encrypted-msg-tx"` // Encrypted messages sent (YANG: IOS-XE 17.12.1)
	EncryptedMsgRX int    `json:"encrypted-msg-rx"` // Encrypted messages received (YANG: IOS-XE 17.12.1)
	TunnelType     string `json:"tunnel-type"`      // CAPWAP mobility tunnel type (YANG: IOS-XE 17.12.1)
}

// MobilityGlobalMsgStats represents mobility global message statistics container.
type MobilityGlobalMsgStats struct {
	MsgStats MsgStats `json:"msg-stats"` // Message statistics data (YANG: IOS-XE 17.12.1)
}

// MsgStats represents message statistics.
type MsgStats struct {
	MobilityAnnounceSent           int `json:"mobility-announce-sent"`            // Mobility announce messages sent count (YANG: IOS-XE 17.12.1)
	MobilityAnnounceReceived       int `json:"mobility-announce-received"`        // Mobility announce messages received count (YANG: IOS-XE 17.12.1)
	MobilityDenylistAddSent        int `json:"mobility-denylist-add-sent"`        // Mobility deny list add messages sent count (YANG: IOS-XE 17.12.1)
	MobilityDenylistAddReceived    int `json:"mobility-denylist-add-received"`    // Mobility deny list add messages received count (YANG: IOS-XE 17.12.1)
	MobilityDenylistDelSent        int `json:"mobility-denylist-del-sent"`        // Mobility deny list delete messages sent count (YANG: IOS-XE 17.12.1)
	MobilityDenylistDelReceived    int `json:"mobility-denylist-del-received"`    // Mobility deny list delete messages received count (YANG: IOS-XE 17.12.1)
	MobilityHandoffRequestSent     int `json:"mobility-handoff-request-sent"`     // Mobility handoff request messages sent count (YANG: IOS-XE 17.12.1)
	MobilityHandoffRequestReceived int `json:"mobility-handoff-request-received"` // Mobility handoff request messages received count (YANG: IOS-XE 17.12.1)
	MobilityHandoffReplySent       int `json:"mobility-handoff-reply-sent"`       // Mobility handoff reply messages sent count (YANG: IOS-XE 17.12.1)
	MobilityHandoffReplyReceived   int `json:"mobility-handoff-reply-received"`   // Mobility handoff reply messages received count (YANG: IOS-XE 17.12.1)
	MobilityHandoffEndSent         int `json:"mobility-handoff-end-sent"`         // Mobility handoff end messages sent count (YANG: IOS-XE 17.12.1)
	MobilityHandoffEndReceived     int `json:"mobility-handoff-end-received"`     // Mobility handoff end messages received count (YANG: IOS-XE 17.12.1)
	MobilityRevokeSent             int `json:"mobility-revoke-sent"`              // Mobility revoke messages sent count (YANG: IOS-XE 17.12.1)
	MobilityRevokeReceived         int `json:"mobility-revoke-received"`          // Mobility revoke messages received count (YANG: IOS-XE 17.12.1)
	MobilityRevokeAckSent          int `json:"mobility-revoke-ack-sent"`          // Mobility revoke acknowledgment messages sent count (YANG: IOS-XE 17.12.1)
	MobilityRevokeAckReceived      int `json:"mobility-revoke-ack-received"`      // Mobility revoke acknowledgment messages received count (YANG: IOS-XE 17.12.1)
	MobilityDirectiveAddSent       int `json:"mobility-directive-add-sent"`       // Mobility directive add messages sent count (YANG: IOS-XE 17.12.1)
	MobilityDirectiveAddReceived   int `json:"mobility-directive-add-received"`   // Mobility directive add messages received count (YANG: IOS-XE 17.12.1)
	MobilityDirectiveDelSent       int `json:"mobility-directive-del-sent"`       // Mobility directive delete messages sent count (YANG: IOS-XE 17.12.1)
	MobilityDirectiveDelReceived   int `json:"mobility-directive-del-received"`   // Mobility directive delete messages received count (YANG: IOS-XE 17.12.1)
	MobilityWlanStatusSent         int `json:"mobility-wlan-status-sent"`         // Mobility WLAN status messages sent count (YANG: IOS-XE 17.12.1)
	MobilityWlanStatusReceived     int `json:"mobility-wlan-status-received"`     // Mobility WLAN status messages received count (YANG: IOS-XE 17.12.1)
}

// MobilityGlobalStats represents mobility global statistics.
type MobilityGlobalStats struct {
	MmMbltyStats         MmMbltyStats `json:"mm-mblty-stats"`          // Mobility manager statistics (Live: IOS-XE 17.12.5)
	NumOfSleepingClients int          `json:"num-of-sleeping-clients"` // Number of sleeping clients (Live: IOS-XE 17.12.5)
}

// MmMbltyStats represents mobility manager statistics.
type MmMbltyStats struct {
	EventDataAllocs                     int `json:"event-data-allocs"`                        // Event data allocation count (YANG: IOS-XE 17.12.1)
	EventDataFrees                      int `json:"event-data-frees"`                         // Event data free count (YANG: IOS-XE 17.12.1)
	FsmSetAllocs                        int `json:"fsm-set-allocs"`                           // FSM set allocation count (YANG: IOS-XE 17.12.1)
	FsmSetFrees                         int `json:"fsm-set-frees"`                            // FSM set free count (YANG: IOS-XE 17.12.1)
	TimerAllocs                         int `json:"timer-allocs"`                             // Timer allocation count (YANG: IOS-XE 17.12.1)
	TimerFrees                          int `json:"timer-frees"`                              // Timer free count (YANG: IOS-XE 17.12.1)
	TimerStarts                         int `json:"timer-starts"`                             // Timer start count (YANG: IOS-XE 17.12.1)
	TimerStops                          int `json:"timer-stops"`                              // Timer stop count (YANG: IOS-XE 17.12.1)
	McfsmInvalidEvents                  int `json:"mcfsm-invalid-events"`                     // MC FSM invalid event count (YANG: IOS-XE 17.12.1)
	McfsmInternalError                  int `json:"mcfsm-internal-error"`                     // MC FSM internal error count (YANG: IOS-XE 17.12.1)
	JoinedAsLocal                       int `json:"joined-as-local"`                          // Client joined as local count (YANG: IOS-XE 17.12.1)
	JoinedAsForeign                     int `json:"joined-as-foreign"`                        // Client joined as foreign count (YANG: IOS-XE 17.12.1)
	JoinedAsExportForeign               int `json:"joined-as-export-foreign"`                 // Client joined as export foreign count (YANG: IOS-XE 17.12.1)
	JoinedAsExportAnchor                int `json:"joined-as-export-anchor"`                  // Client joined as export anchor count (YANG: IOS-XE 17.12.1)
	LocalToAnchor                       int `json:"local-to-anchor"`                          // Local to anchor transition count (YANG: IOS-XE 17.12.1)
	AnchorToLocal                       int `json:"anchor-to-local"`                          // Anchor to local transition count (YANG: IOS-XE 17.12.1)
	LocalDelete                         int `json:"local-delete"`                             // Local client delete count (YANG: IOS-XE 17.12.1)
	RemoteDelete                        int `json:"remote-delete"`                            // Remote client delete count (YANG: IOS-XE 17.12.1)
	McfsmDeleteInternalError            int `json:"mcfsm-delete-internal-error"`              // MC FSM delete internal error count (YANG: IOS-XE 17.12.1)
	McfsmRoamInternalError              int `json:"mcfsm-roam-internal-error"`                // MC FSM roam internal error count (YANG: IOS-XE 17.12.1)
	L2RoamCount                         int `json:"l2-roam-count"`                            // Layer 2 roam count (YANG: IOS-XE 17.12.1)
	L3RoamCount                         int `json:"l3-roam-count"`                            // Layer 3 roam count (YANG: IOS-XE 17.12.1)
	FlexClientRoamingCount              int `json:"flex-client-roaming-count"`                // FlexConnect client roaming count (YANG: IOS-XE 17.12.1)
	InterWncdRoamCount                  int `json:"inter-wncd-roam-count"`                    // Inter-WNC roam count (YANG: IOS-XE 17.12.1)
	ExpAncReqSent                       int `json:"exp-anc-req-sent"`                         // Export anchor request sent count (YANG: IOS-XE 17.12.1)
	ExpAncReqReceived                   int `json:"exp-anc-req-received"`                     // Export anchor request received count (YANG: IOS-XE 17.12.1)
	ExpAncRespOkSent                    int `json:"exp-anc-resp-ok-sent"`                     // Export anchor response OK sent count (YANG: IOS-XE 17.12.1)
	ExpAncRespGenericDenySent           int `json:"exp-anc-resp-generic-deny-sent"`           // Export anchor response generic deny sent count (YANG: IOS-XE 17.12.1)
	ExpAncRespClientBlacklistedSent     int `json:"exp-anc-resp-client-blacklisted-sent"`     // Export anchor response client blacklisted sent count (YANG: IOS-XE 17.12.1)
	ExpAncRespLimitReachedSent          int `json:"exp-anc-resp-limit-reached-sent"`          // Export anchor response limit reached sent count (YANG: IOS-XE 17.12.1)
	ExpAncRespProfileMismatchSent       int `json:"exp-anc-resp-profile-mismatch-sent"`       // Export anchor response profile mismatch sent count (YANG: IOS-XE 17.12.1)
	ExpAncRespOkReceived                int `json:"exp-anc-resp-ok-received"`                 // Export anchor response OK received count (YANG: IOS-XE 17.12.1)
	ExpAncRespGenericDenyReceived       int `json:"exp-anc-resp-generic-deny-received"`       // Export anchor response generic deny received count (YANG: IOS-XE 17.12.1)
	ExpAncRespClientBlacklistedReceived int `json:"exp-anc-resp-client-blacklisted-received"` // Export anchor response client blacklisted received count (YANG: IOS-XE 17.12.1)
	ExpAncRespLimitReachedReceived      int `json:"exp-anc-resp-limit-reached-received"`      // Export anchor response limit reached received count (YANG: IOS-XE 17.12.1)
	ExpAncRespProfileMismatchReceived   int `json:"exp-anc-resp-profile-mismatch-received"`   // Export anchor response profile mismatch received count (YANG: IOS-XE 17.12.1)
	ExpAncRespUnknownReceived           int `json:"exp-anc-resp-unknown-received"`            // Export anchor response unknown received count (YANG: IOS-XE 17.12.1)
	HandoffSentMsBlacklist              int `json:"handoff-sent-ms-blacklist"`                // Handoff sent MS blacklist count (YANG: IOS-XE 17.12.1)
	HandoffReceivedMsBlacklist          int `json:"handoff-received-ms-blacklist"`            // Handoff received MS blacklist count (YANG: IOS-XE 17.12.1)
}

// WlanClientLimit represents WLAN client limit information.
type WlanClientLimit struct {
	WlanProfile      string `json:"wlan-profile"`       // Name of the wlan profile (Live: IOS-XE 17.12.5)
	CurrClientsCount int    `json:"curr-clients-count"` // Current client count (YANG: IOS-XE 17.12.1)
}
