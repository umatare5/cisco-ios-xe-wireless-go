// Package mobility provides mobility operational data functionality for the Cisco Wireless Network Controller API.
package mobility

import (
	"context"

	wnc "github.com/umatare5/cisco-xe-wireless-restconf-go"
)

const (
	// MobilityOperBasePath defines the base path for mobility operational data endpoints.
	MobilityOperBasePath = "/restconf/data/Cisco-IOS-XE-wireless-mobility-oper:mobility-oper-data"
	// MobilityOperEndpoint defines the endpoint for mobility operational data.
	MobilityOperEndpoint = MobilityOperBasePath
	// MmIfGlobalStatsEndpoint defines the endpoint for mobility manager interface global statistics.
	MmIfGlobalStatsEndpoint = MobilityOperBasePath + "/mm-if-global-stats"
	// MmIfGlobalMsgStatsEndpoint defines the endpoint for mobility manager interface global message statistics.
	MmIfGlobalMsgStatsEndpoint = MobilityOperBasePath + "/mm-if-global-msg-stats"
	// MobilityGlobalStatsEndpoint defines the endpoint for mobility global statistics.
	MobilityGlobalStatsEndpoint = MobilityOperBasePath + "/mobility-global-stats"
	// MmGlobalDataEndpoint defines the endpoint for mobility manager global data.
	MmGlobalDataEndpoint = MobilityOperBasePath + "/mm-global-data"
	// MobilityGlobalMsgStatsEndpoint defines the endpoint for mobility global message statistics.
	MobilityGlobalMsgStatsEndpoint = MobilityOperBasePath + "/mobility-global-msg-stats"
	// MobilityClientDataEndpoint defines the endpoint for mobility client data.
	MobilityClientDataEndpoint = MobilityOperBasePath + "/mobility-client-data"
	// ApCacheEndpoint defines the endpoint for AP cache data.
	ApCacheEndpoint = MobilityOperBasePath + "/ap-cache"
	// ApPeerListEndpoint defines the endpoint for AP peer list data.
	ApPeerListEndpoint = MobilityOperBasePath + "/ap-peer-list"
	// MobilityClientStatsEndpoint defines the endpoint for mobility client statistics.
	MobilityClientStatsEndpoint = MobilityOperBasePath + "/mobility-client-stats"
	// WlanClientLimitEndpoint defines the endpoint for WLAN client limit data.
	WlanClientLimitEndpoint = MobilityOperBasePath + "/wlan-client-limit"
	// MobilityGlobalDTLSStatsEndpoint defines the endpoint for mobility global DTLS statistics.
	MobilityGlobalDTLSStatsEndpoint = MobilityOperBasePath + "/mobility-global-dtls-stats"
)

// MobilityOperResponse represents the response structure for mobility operational data.
type MobilityOperResponse struct {
	CiscoIOSXEWirelessMobilityOperMobilityOperData struct {
		MmIfGlobalStats         MmIfGlobalStats         `json:"mm-if-global-stats"`
		MmIfGlobalMsgStats      MmIfGlobalMsgStats      `json:"mm-if-global-msg-stats"`
		MobilityGlobalStats     MobilityGlobalStats     `json:"mobility-global-stats"`
		MmGlobalData            MmGlobalData            `json:"mm-global-data"`
		MobilityGlobalMsgStats  MobilityGlobalMsgStats  `json:"mobility-global-msg-stats"`
		MobilityClientData      []MobilityClient        `json:"mobility-client-data"`
		ApCache                 []ApCache               `json:"ap-cache"`
		ApPeerList              []ApPeer                `json:"ap-peer-list"`
		MobilityClientStats     []MobilityClientStat    `json:"mobility-client-stats"`
		WlanClientLimit         []WlanClientLimit       `json:"wlan-client-limit"`
		MobilityGlobalDTLSStats MobilityGlobalDTLSStats `json:"mobility-global-dtls-stats"`
	} `json:"Cisco-IOS-XE-wireless-mobility-oper:mobility-oper-data"`
}

// MmIfGlobalStatsResponse represents the response structure for mobility manager interface global statistics.
type MmIfGlobalStatsResponse struct {
	MmIfGlobalStats MmIfGlobalStats `json:"Cisco-IOS-XE-wireless-mobility-oper:mm-if-global-stats"`
}

// MmIfGlobalMsgStatsResponse represents the response structure for mobility manager interface global message statistics.
type MmIfGlobalMsgStatsResponse struct {
	MmIfGlobalMsgStats MmIfGlobalMsgStats `json:"Cisco-IOS-XE-wireless-mobility-oper:mm-if-global-msg-stats"`
}

// MobilityGlobalStatsResponse represents the response structure for mobility global statistics.
type MobilityGlobalStatsResponse struct {
	MobilityGlobalStats MobilityGlobalStats `json:"Cisco-IOS-XE-wireless-mobility-oper:mobility-global-stats"`
}

// MmGlobalDataResponse represents the response structure for mobility manager global data.
type MmGlobalDataResponse struct {
	MmGlobalData MmGlobalData `json:"Cisco-IOS-XE-wireless-mobility-oper:mm-global-data"`
}

type MobilityGlobalMsgStatsResponse struct {
	MobilityGlobalMsgStats MobilityGlobalMsgStats `json:"Cisco-IOS-XE-wireless-mobility-oper:mobility-global-msg-stats"`
}

type MobilityClientDataResponse struct {
	MobilityClientData []MobilityClient `json:"Cisco-IOS-XE-wireless-mobility-oper:mobility-client-data"`
}

type ApCacheResponse struct {
	ApCache []ApCache `json:"Cisco-IOS-XE-wireless-mobility-oper:ap-cache"`
}

type ApPeerListResponse struct {
	ApPeerList []ApPeer `json:"Cisco-IOS-XE-wireless-mobility-oper:ap-peer-list"`
}

type MobilityClientStatsResponse struct {
	MobilityClientStats []MobilityClientStat `json:"Cisco-IOS-XE-wireless-mobility-oper:mobility-client-stats"`
}

type WlanClientLimitResponse struct {
	WlanClientLimit []WlanClientLimit `json:"Cisco-IOS-XE-wireless-mobility-oper:wlan-client-limit"`
}

type MobilityGlobalDTLSStatsResponse struct {
	MobilityGlobalDTLSStats MobilityGlobalDTLSStats `json:"Cisco-IOS-XE-wireless-mobility-oper:mobility-global-dtls-stats"`
}

type MmIfGlobalStats struct {
	MbltyStats struct {
		EventDataAllocs               int `json:"event-data-allocs"`
		EventDataFrees                int `json:"event-data-frees"`
		MmifFsmInvalidEvents          int `json:"mmif-fsm-invalid-events"`
		MmifScheduleErrors            int `json:"mmif-schedule-errors"`
		MmifFsmFailure                int `json:"mmif-fsm-failure"`
		MmifIpcFailure                int `json:"mmif-ipc-failure"`
		MmifDbFailure                 int `json:"mmif-db-failure"`
		MmifInvalidParamsFailure      int `json:"mmif-invalid-params-failure"`
		MmifMmMsgDecodeFailure        int `json:"mmif-mm-msg-decode-failure"`
		MmifUnknownFailure            int `json:"mmif-unknown-failure"`
		MmifClientHandoffFailure      int `json:"mmif-client-handoff-failure"`
		MmifClientHandoffSuccess      int `json:"mmif-client-handoff-success"`
		MmifAnchorDeny                int `json:"mmif-anchor-deny"`
		MmifRemoteDelete              int `json:"mmif-remote-delete"`
		MmifTunnelDownDelete          int `json:"mmif-tunnel-down-delete"`
		MmifMbssidDownEvent           int `json:"mmif-mbssid-down-event"`
		IntraWncdRoamCount            int `json:"intra-wncd-roam-count"`
		RemoteInterCtrlrRoams         int `json:"remote-inter-ctrlr-roams"`
		RemoteWebauthPendRoams        int `json:"remote-webauth-pend-roams"`
		AnchorRequestSent             int `json:"anchor-request-sent"`
		AnchorRequestGrantReceived    int `json:"anchor-request-grant-received"`
		AnchorRequestDenyReceived     int `json:"anchor-request-deny-received"`
		AnchorRequestReceived         int `json:"anchor-request-received"`
		AnchorRequestGrantSent        int `json:"anchor-request-grant-sent"`
		AnchorRequestDenySent         int `json:"anchor-request-deny-sent"`
		HandoffReceivedOk             int `json:"handoff-received-ok"`
		HandoffReceivedGrpMismatch    int `json:"handoff-received-grp-mismatch"`
		HandoffReceivedMsUnknown      int `json:"handoff-received-ms-unknown"`
		HandoffReceivedMsSsid         int `json:"handoff-received-ms-ssid"`
		HandoffReceivedDeny           int `json:"handoff-received-deny"`
		HandoffSentOk                 int `json:"handoff-sent-ok"`
		HandoffSentGrpMismatch        int `json:"handoff-sent-grp-mismatch"`
		HandoffSentMsUnknown          int `json:"handoff-sent-ms-unknown"`
		HandoffSentMsSsid             int `json:"handoff-sent-ms-ssid"`
		HandoffSentDeny               int `json:"handoff-sent-deny"`
		HandoffReceivedL3VlanOverride int `json:"handoff-received-l3-vlan-override"`
		HandoffReceivedUnknownPeer    int `json:"handoff-received-unknown-peer"`
		HandoffSentL3VlanOverride     int `json:"handoff-sent-l3-vlan-override"`
	} `json:"mblty-stats"`
	MbltyDomainInfo struct {
		MobilityDomainID int `json:"mobility-domain-id"`
	} `json:"mblty-domain-info"`
}

type MmIfGlobalMsgStats struct {
	IpcStats []IpcStats `json:"ipc-stats"`
}

type IpcStats struct {
	Type      int    `json:"type"`
	Allocs    int    `json:"allocs"`
	Frees     int    `json:"frees"`
	Tx        int    `json:"tx"`
	Rx        int    `json:"rx"`
	Forwarded int    `json:"forwarded"`
	TxErrors  int    `json:"tx-errors"`
	RxErrors  int    `json:"rx-errors"`
	TxRetries int    `json:"tx-retries"`
	Drops     int    `json:"drops"`
	Built     int    `json:"built"`
	Processed int    `json:"processed"`
	MmMsgType string `json:"mm-msg-type"`
}

type MobilityGlobalStats struct {
	MmMbltyStats struct {
		EventDataAllocs                     int `json:"event-data-allocs"`
		EventDataFrees                      int `json:"event-data-frees"`
		FsmSetAllocs                        int `json:"fsm-set-allocs"`
		FsmSetFrees                         int `json:"fsm-set-frees"`
		TimerAllocs                         int `json:"timer-allocs"`
		TimerFrees                          int `json:"timer-frees"`
		TimerStarts                         int `json:"timer-starts"`
		TimerStops                          int `json:"timer-stops"`
		McfsmInvalidEvents                  int `json:"mcfsm-invalid-events"`
		McfsmInternalError                  int `json:"mcfsm-internal-error"`
		JoinedAsLocal                       int `json:"joined-as-local"`
		JoinedAsForeign                     int `json:"joined-as-foreign"`
		JoinedAsExportForeign               int `json:"joined-as-export-foreign"`
		JoinedAsExportAnchor                int `json:"joined-as-export-anchor"`
		LocalToAnchor                       int `json:"local-to-anchor"`
		AnchorToLocal                       int `json:"anchor-to-local"`
		LocalDelete                         int `json:"local-delete"`
		RemoteDelete                        int `json:"remote-delete"`
		McfsmDeleteInternalError            int `json:"mcfsm-delete-internal-error"`
		McfsmRoamInternalError              int `json:"mcfsm-roam-internal-error"`
		L2RoamCount                         int `json:"l2-roam-count"`
		L3RoamCount                         int `json:"l3-roam-count"`
		FlexClientRoamingCount              int `json:"flex-client-roaming-count"`
		InterWncdRoamCount                  int `json:"inter-wncd-roam-count"`
		ExpAncReqSent                       int `json:"exp-anc-req-sent"`
		ExpAncReqReceived                   int `json:"exp-anc-req-received"`
		ExpAncRespOkSent                    int `json:"exp-anc-resp-ok-sent"`
		ExpAncRespGenericDenySent           int `json:"exp-anc-resp-generic-deny-sent"`
		ExpAncRespClientBlacklistedSent     int `json:"exp-anc-resp-client-blacklisted-sent"`
		ExpAncRespLimitReachedSent          int `json:"exp-anc-resp-limit-reached-sent"`
		ExpAncRespProfileMismatchSent       int `json:"exp-anc-resp-profile-mismatch-sent"`
		ExpAncRespOkReceived                int `json:"exp-anc-resp-ok-received"`
		ExpAncRespGenericDenyReceived       int `json:"exp-anc-resp-generic-deny-received"`
		ExpAncRespClientBlacklistedReceived int `json:"exp-anc-resp-client-blacklisted-received"`
		ExpAncRespLimitReachedReceived      int `json:"exp-anc-resp-limit-reached-received"`
		ExpAncRespProfileMismatchReceived   int `json:"exp-anc-resp-profile-mismatch-received"`
		ExpAncRespUnknownReceived           int `json:"exp-anc-resp-unknown-received"`
		HandoffSentMsBlacklist              int `json:"handoff-sent-ms-blacklist"`
		HandoffReceivedMsBlacklist          int `json:"handoff-received-ms-blacklist"`
	} `json:"mm-mblty-stats"`
	NumOfSleepingClients int `json:"num-of-sleeping-clients"`
}

type MmGlobalData struct {
	MmMacAddr string `json:"mm-mac-addr"`
}

type MobilityGlobalMsgStats struct {
	IpcStats       []IpcStats       `json:"ipc-stats"`
	DgramStats     []DgramStats     `json:"dgram-stats"`
	DgramDataStats []DgramDataStats `json:"dgram-data-stats"`
}

type DgramStats struct {
	Type      int    `json:"type"`
	Allocs    int    `json:"allocs"`
	Frees     int    `json:"frees"`
	Tx        int    `json:"tx"`
	Rx        int    `json:"rx"`
	Forwarded int    `json:"forwarded"`
	TxErrors  int    `json:"tx-errors"`
	RxErrors  int    `json:"rx-errors"`
	TxRetries int    `json:"tx-retries"`
	Drops     int    `json:"drops"`
	Built     int    `json:"built"`
	Processed int    `json:"processed"`
	MmMsgType string `json:"mm-msg-type"`
}

type DgramDataStats struct {
	Type      int    `json:"type"`
	Allocs    int    `json:"allocs"`
	Frees     int    `json:"frees"`
	Tx        int    `json:"tx"`
	Rx        int    `json:"rx"`
	Forwarded int    `json:"forwarded"`
	TxErrors  int    `json:"tx-errors"`
	RxErrors  int    `json:"rx-errors"`
	TxRetries int    `json:"tx-retries"`
	Drops     int    `json:"drops"`
	Built     int    `json:"built"`
	Processed int    `json:"processed"`
	MmMsgType string `json:"mm-msg-type"`
}

type MobilityClient struct {
	ClientMac           string `json:"client-mac"`
	ClientIfID          int64  `json:"client-ifid"`
	OwnerInstance       int    `json:"owner-instance"`
	AssocTimeStamp      string `json:"assoc-time-stamp"`
	MobilityState       string `json:"mobility-state"`
	ClientRole          string `json:"client-role"`
	ClientType          string `json:"client-type"`
	ClientMode          string `json:"client-mode"`
	ClientRoamType      string `json:"client-roam-type"`
	PeerIP              string `json:"peer-ip"`
	EntryLastUpdateTime string `json:"entry-last-update-time"`
	ClientCreateTime    string `json:"client-create-time"`
	WlanProfile         string `json:"wlan-profile"`
}

type ApCache struct {
	MacAddr      string `json:"mac-addr"`
	Name         string `json:"name"`
	ControllerIP string `json:"controller-ip"`
	Source       string `json:"source"`
}

type ApPeer struct {
	PeerIP  string `json:"peer-ip"`
	ApCount int    `json:"ap-count"`
	Source  string `json:"source"`
}

type MobilityClientStat struct {
	ClientMac    string `json:"client-mac"`
	MmMbltyStats struct {
		EventDataAllocs                     int `json:"event-data-allocs"`
		EventDataFrees                      int `json:"event-data-frees"`
		FsmSetAllocs                        int `json:"fsm-set-allocs"`
		FsmSetFrees                         int `json:"fsm-set-frees"`
		TimerAllocs                         int `json:"timer-allocs"`
		TimerFrees                          int `json:"timer-frees"`
		TimerStarts                         int `json:"timer-starts"`
		TimerStops                          int `json:"timer-stops"`
		McfsmInvalidEvents                  int `json:"mcfsm-invalid-events"`
		McfsmInternalError                  int `json:"mcfsm-internal-error"`
		JoinedAsLocal                       int `json:"joined-as-local"`
		JoinedAsForeign                     int `json:"joined-as-foreign"`
		JoinedAsExportForeign               int `json:"joined-as-export-foreign"`
		JoinedAsExportAnchor                int `json:"joined-as-export-anchor"`
		LocalToAnchor                       int `json:"local-to-anchor"`
		AnchorToLocal                       int `json:"anchor-to-local"`
		LocalDelete                         int `json:"local-delete"`
		RemoteDelete                        int `json:"remote-delete"`
		McfsmDeleteInternalError            int `json:"mcfsm-delete-internal-error"`
		McfsmRoamInternalError              int `json:"mcfsm-roam-internal-error"`
		L2RoamCount                         int `json:"l2-roam-count"`
		L3RoamCount                         int `json:"l3-roam-count"`
		FlexClientRoamingCount              int `json:"flex-client-roaming-count"`
		InterWncdRoamCount                  int `json:"inter-wncd-roam-count"`
		ExpAncReqSent                       int `json:"exp-anc-req-sent"`
		ExpAncReqReceived                   int `json:"exp-anc-req-received"`
		ExpAncRespOkSent                    int `json:"exp-anc-resp-ok-sent"`
		ExpAncRespGenericDenySent           int `json:"exp-anc-resp-generic-deny-sent"`
		ExpAncRespClientBlacklistedSent     int `json:"exp-anc-resp-client-blacklisted-sent"`
		ExpAncRespLimitReachedSent          int `json:"exp-anc-resp-limit-reached-sent"`
		ExpAncRespProfileMismatchSent       int `json:"exp-anc-resp-profile-mismatch-sent"`
		ExpAncRespOkReceived                int `json:"exp-anc-resp-ok-received"`
		ExpAncRespGenericDenyReceived       int `json:"exp-anc-resp-generic-deny-received"`
		ExpAncRespClientBlacklistedReceived int `json:"exp-anc-resp-client-blacklisted-received"`
		ExpAncRespLimitReachedReceived      int `json:"exp-anc-resp-limit-reached-received"`
		ExpAncRespProfileMismatchReceived   int `json:"exp-anc-resp-profile-mismatch-received"`
		ExpAncRespUnknownReceived           int `json:"exp-anc-resp-unknown-received"`
		HandoffSentMsBlacklist              int `json:"handoff-sent-ms-blacklist"`
		HandoffReceivedMsBlacklist          int `json:"handoff-received-ms-blacklist"`
	} `json:"mm-mblty-stats"`
	IpcStats []struct {
		Type      int    `json:"type"`
		Allocs    int    `json:"allocs"`
		Frees     int    `json:"frees"`
		Tx        int    `json:"tx"`
		Rx        int    `json:"rx"`
		Forwarded int    `json:"forwarded"`
		TxErrors  int    `json:"tx-errors"`
		RxErrors  int    `json:"rx-errors"`
		TxRetries int    `json:"tx-retries"`
		Drops     int    `json:"drops"`
		Built     int    `json:"built"`
		Processed int    `json:"processed"`
		MmMsgType string `json:"mm-msg-type"`
	} `json:"ipc-stats"`
}

type WlanClientLimit struct {
	WlanProfile      string `json:"wlan-profile"`
	CurrClientsCount int    `json:"curr-clients-count"`
}

type MobilityGlobalDTLSStats struct {
	EventStats []DTLSEventStat `json:"event-stats"`
	MsgStats   []DTLSMsgStat   `json:"msg-stats"`
}

type DTLSEventStat struct {
	ConnectStart       int    `json:"connect-start"`
	ConnectEstablished int    `json:"connect-established"`
	Close              int    `json:"close"`
	KeyPlumbStart      int    `json:"key-plumb-start"`
	KeyPlumbAcked      int    `json:"key-plumb-acked"`
	TunnelType         string `json:"tunnel-type"`
}

type DTLSMsgStat struct {
	HandshakeMsgTx int    `json:"handshake-msg-tx"`
	HandshakeMsgRx int    `json:"handshake-msg-rx"`
	EncryptedMsgTx int    `json:"encrypted-msg-tx"`
	EncryptedMsgRx int    `json:"encrypted-msg-rx"`
	TunnelType     string `json:"tunnel-type"`
}

// GetMobilityOper retrieves mobility operational data.
func GetMobilityOper(client *wnc.Client, ctx context.Context) (*MobilityOperResponse, error) {
	var data MobilityOperResponse
	if err := client.SendAPIRequest(ctx, MobilityOperEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetMobilityMmIfGlobalStats retrieves mobility manager interface global statistics.
func GetMobilityMmIfGlobalStats(client *wnc.Client, ctx context.Context) (*MmIfGlobalStatsResponse, error) {
	var data MmIfGlobalStatsResponse
	if err := client.SendAPIRequest(ctx, MmIfGlobalStatsEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetMobilityMmIfGlobalMsgStats retrieves mobility manager interface global message statistics.
func GetMobilityMmIfGlobalMsgStats(client *wnc.Client, ctx context.Context) (*MmIfGlobalMsgStatsResponse, error) {
	var data MmIfGlobalMsgStatsResponse
	if err := client.SendAPIRequest(ctx, MmIfGlobalMsgStatsEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetMobilityGlobalStats retrieves mobility global statistics.
func GetMobilityGlobalStats(client *wnc.Client, ctx context.Context) (*MobilityGlobalStatsResponse, error) {
	var data MobilityGlobalStatsResponse
	if err := client.SendAPIRequest(ctx, MobilityGlobalStatsEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetMobilityMmGlobalData retrieves mobility manager global data.
func GetMobilityMmGlobalData(client *wnc.Client, ctx context.Context) (*MmGlobalDataResponse, error) {
	var data MmGlobalDataResponse
	if err := client.SendAPIRequest(ctx, MmGlobalDataEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetMobilityGlobalMsgStats retrieves mobility global message statistics.
func GetMobilityGlobalMsgStats(client *wnc.Client, ctx context.Context) (*MobilityGlobalMsgStatsResponse, error) {
	var data MobilityGlobalMsgStatsResponse
	if err := client.SendAPIRequest(ctx, MobilityGlobalMsgStatsEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetMobilityClientData retrieves mobility client data.
func GetMobilityClientData(client *wnc.Client, ctx context.Context) (*MobilityClientDataResponse, error) {
	var data MobilityClientDataResponse
	if err := client.SendAPIRequest(ctx, MobilityClientDataEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetMobilityApCache retrieves AP cache data.
func GetMobilityApCache(client *wnc.Client, ctx context.Context) (*ApCacheResponse, error) {
	var data ApCacheResponse
	if err := client.SendAPIRequest(ctx, ApCacheEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetMobilityApPeerList retrieves AP peer list data.
func GetMobilityApPeerList(client *wnc.Client, ctx context.Context) (*ApPeerListResponse, error) {
	var data ApPeerListResponse
	if err := client.SendAPIRequest(ctx, ApPeerListEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetMobilityClientStats retrieves mobility client statistics.
func GetMobilityClientStats(client *wnc.Client, ctx context.Context) (*MobilityClientStatsResponse, error) {
	var data MobilityClientStatsResponse
	if err := client.SendAPIRequest(ctx, MobilityClientStatsEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetMobilityWlanClientLimit retrieves WLAN client limit data.
func GetMobilityWlanClientLimit(client *wnc.Client, ctx context.Context) (*WlanClientLimitResponse, error) {
	var data WlanClientLimitResponse
	if err := client.SendAPIRequest(ctx, WlanClientLimitEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetMobilityGlobalDTLSStats retrieves mobility global DTLS statistics.
func GetMobilityGlobalDTLSStats(client *wnc.Client, ctx context.Context) (*MobilityGlobalDTLSStatsResponse, error) {
	var data MobilityGlobalDTLSStatsResponse
	if err := client.SendAPIRequest(ctx, MobilityGlobalDTLSStatsEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}
