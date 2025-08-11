package model

// MobilityOperResponse represents the response structure for Mobility operational data.
type MobilityOperResponse struct {
	MobilityOperData MobilitySystemOperData `json:"Cisco-IOS-XE-wireless-mobility-oper:mobility-oper-data"`
}

// MobilitySystemOperData contains Mobility operational data
type MobilitySystemOperData struct {
	ApCache                 []ApCache               `json:"ap-cache"`
	ApPeerList              []ApPeerList            `json:"ap-peer-list"`
	MmGlobalData            MmGlobalData            `json:"mm-global-data"`
	MmIfGlobalMsgStats      MmIfGlobalMsgStats      `json:"mm-if-global-msg-stats"`
	MmIfGlobalStats         MmIfGlobalStats         `json:"mm-if-global-stats"`
	MobilityClientData      []MobilityClientData    `json:"mobility-client-data"`
	MobilityClientStats     MobilityClientStats     `json:"mobility-client-stats"`
	MobilityGlobalDtlsStats MobilityGlobalDtlsStats `json:"mobility-global-dtls-stats"`
	MobilityGlobalMsgStats  MobilityGlobalMsgStats  `json:"mobility-global-msg-stats"`
	MobilityGlobalStats     MobilityGlobalStats     `json:"mobility-global-stats"`
	WlanClientLimit         []WlanClientLimit       `json:"wlan-client-limit"`
}

// ApCache represents AP cache information
type ApCache struct {
	// Define based on actual structure when available
}

// ApPeerList represents AP peer list information
type ApPeerList struct {
	// Define based on actual structure when available
}

// MmGlobalData represents MM global data
type MmGlobalData struct {
	// Define based on actual structure when available
}

// MmIfGlobalMsgStats represents MM interface global message statistics
type MmIfGlobalMsgStats struct {
	// Define based on actual structure when available
}

// MmIfGlobalStats represents MM interface global statistics
type MmIfGlobalStats struct {
	MbltyStats MbltyStats `json:"mblty-stats"`
}

// MbltyStats represents mobility statistics
type MbltyStats struct {
	EventDataAllocs          int `json:"event-data-allocs"`
	EventDataFrees           int `json:"event-data-frees"`
	MmifFsmInvalidEvents     int `json:"mmif-fsm-invalid-events"`
	MmifScheduleErrors       int `json:"mmif-schedule-errors"`
	MmifFsmFailure           int `json:"mmif-fsm-failure"`
	MmifIpcFailure           int `json:"mmif-ipc-failure"`
	MmifDBFailure            int `json:"mmif-db-failure"`
	MmifInvalidParamsFailure int `json:"mmif-invalid-params-failure"`
	MmifMmMsgDecodeFailure   int `json:"mmif-mm-msg-decode-failure"`
	MmifUnknownFailure       int `json:"mmif-unknown-failure"`
	MmifClientHandoffFailure int `json:"mmif-client-handoff-failure"`
	MmifClientHandoffSuccess int `json:"mmif-client-handoff-success"`
	MmifAnchorDeny           int `json:"mmif-anchor-deny"`
	MmifRemoteDelete         int `json:"mmif-remote-delete"`
	MmifTunnelDownDelete     int `json:"mmif-tunnel-down-delete"`
	MmifMbssidDownEvent      int `json:"mmif-mbssid-down-event"`
}

// MobilityClientData represents mobility client data
type MobilityClientData struct {
	// Define based on actual structure when available
}

// MobilityClientStats represents mobility client statistics
type MobilityClientStats struct {
	// Define based on actual structure when available
}

// MobilityGlobalDtlsStats represents mobility global DTLS statistics
type MobilityGlobalDtlsStats struct {
	// Define based on actual structure when available
}

// MobilityGlobalMsgStats represents mobility global message statistics
type MobilityGlobalMsgStats struct {
	// Define based on actual structure when available
}

// MobilityGlobalStats represents mobility global statistics
type MobilityGlobalStats struct {
	// Define based on actual structure when available
}

// WlanClientLimit represents WLAN client limit information
type WlanClientLimit struct {
	// Define based on actual structure when available
}
