// Package model provides data models for mobility operational data.
package model

// MobilityOper  represents the Mobility operational data.
type MobilityOper struct {
	MobilityOperData MobilitySystemOperData `json:"Cisco-IOS-XE-wireless-mobility-oper:mobility-oper-data"`
}

// MobilityOperApCache  represents the AP cache data.
type MobilityOperApCache struct {
	ApCache []ApCache `json:"Cisco-IOS-XE-wireless-mobility-oper:ap-cache"`
}

// MobilityOperApPeerList  represents the AP peer list data.
type MobilityOperApPeerList struct {
	ApPeerList []ApPeerList `json:"Cisco-IOS-XE-wireless-mobility-oper:ap-peer-list"`
}

// MobilityOperMmGlobalData  represents the structure for MM global data.
type MobilityOperMmGlobalData struct {
	MmGlobalData MmGlobalData `json:"Cisco-IOS-XE-wireless-mobility-oper:mm-global-data"`
}

// MobilityOperMmIfGlobalMsgStats  represents the structure for MM interface global message statistics.
type MobilityOperMmIfGlobalMsgStats struct {
	MmIfGlobalMsgStats MmIfGlobalMsgStats `json:"Cisco-IOS-XE-wireless-mobility-oper:mm-if-global-msg-stats"`
}

// MobilityOperMmIfGlobalStats  represents the structure for MM interface global statistics.
type MobilityOperMmIfGlobalStats struct {
	MmIfGlobalStats MmIfGlobalStats `json:"Cisco-IOS-XE-wireless-mobility-oper:mm-if-global-stats"`
}

// MobilityOperMobilityClientData  represents the structure for mobility client data.
type MobilityOperMobilityClientData struct {
	MobilityClientData []MobilityClientData `json:"Cisco-IOS-XE-wireless-mobility-oper:mobility-client-data"`
}

// MobilityOperMobilityClientStats  represents the structure for mobility client statistics.
type MobilityOperMobilityClientStats struct {
	MobilityClientStats MobilityClientStats `json:"Cisco-IOS-XE-wireless-mobility-oper:mobility-client-stats"`
}

// MobilityOperMobilityGlobalDtlsStats  represents the structure for mobility global DTLS statistics.
type MobilityOperMobilityGlobalDtlsStats struct {
	MobilityGlobalDtlsStats MobilityGlobalDtlsStats `json:"Cisco-IOS-XE-wireless-mobility-oper:mobility-global-dtls-stats"`
}

// MobilityOperMobilityGlobalMsgStats  represents the structure for mobility global message statistics.
type MobilityOperMobilityGlobalMsgStats struct {
	MobilityGlobalMsgStats MobilityGlobalMsgStats `json:"Cisco-IOS-XE-wireless-mobility-oper:mobility-global-msg-stats"`
}

// MobilityOperMobilityGlobalStats  represents the structure for mobility global statistics.
type MobilityOperMobilityGlobalStats struct {
	MobilityGlobalStats MobilityGlobalStats `json:"Cisco-IOS-XE-wireless-mobility-oper:mobility-global-stats"`
}

// MobilityOperWlanClientLimit  represents the structure for WLAN client limit data.
type MobilityOperWlanClientLimit struct {
	WlanClientLimit []WlanClientLimit `json:"Cisco-IOS-XE-wireless-mobility-oper:wlan-client-limit"`
}

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

type ApCache struct {
	// Define based on actual structure when available
}

type ApPeerList struct {
	// Define based on actual structure when available
}

type MmGlobalData struct {
	// Define based on actual structure when available
}

type MmIfGlobalMsgStats struct {
	// Define based on actual structure when available
}

type MmIfGlobalStats struct {
	MbltyStats MbltyStats `json:"mblty-stats"`
}

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

type MobilityClientData struct {
	// Define based on actual structure when available
}

type MobilityClientStats struct {
	// Define based on actual structure when available
}

type MobilityGlobalDtlsStats struct {
	// Define based on actual structure when available
}

type MobilityGlobalMsgStats struct {
	// Define based on actual structure when available
}

type MobilityGlobalStats struct {
	// Define based on actual structure when available
}

type WlanClientLimit struct {
	// Define based on actual structure when available
}
