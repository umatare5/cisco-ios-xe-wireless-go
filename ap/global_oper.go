// Package ap provides access point global operational data functionality for the Cisco Wireless Network Controller API.
package ap

import (
	"context"

	"time"

	wnc "github.com/umatare5/cisco-xe-wireless-restconf-go"
)

const (
	// ApGlobalOperBasePath defines the base path for AP global operational data endpoints.
	ApGlobalOperBasePath = "/restconf/data/Cisco-IOS-XE-wireless-ap-global-oper:ap-global-oper-data"
	// ApGlobalOperEndpoint defines the endpoint for AP global operational data.
	ApGlobalOperEndpoint = ApGlobalOperBasePath
	// ApHistoryEndpoint defines the endpoint for AP history data.
	ApHistoryEndpoint = ApGlobalOperBasePath + "/ap-history"
	// EwlcApStatsEndpoint defines the endpoint for EWLC AP statistics.
	EwlcApStatsEndpoint = ApGlobalOperBasePath + "/ewlc-ap-stats"
	// ApImgPredownloadStatsEndpoint defines the endpoint for AP image predownload statistics.
	ApImgPredownloadStatsEndpoint = ApGlobalOperBasePath + "/ap-img-predownload-stats"
	// ApJoinStatsEndpoint defines the endpoint for AP join statistics.
	ApJoinStatsEndpoint = ApGlobalOperBasePath + "/ap-join-stats"
	// WlanClientStatsEndpoint defines the endpoint for WLAN client statistics.
	WlanClientStatsEndpoint = ApGlobalOperBasePath + "/wlan-client-stats"
	// EmltdJoinCountStatEndpoint defines the endpoint for EMLTD join count statistics.
	EmltdJoinCountStatEndpoint = ApGlobalOperBasePath + "/emltd-join-count-stat"
)

// ApGlobalOperResponse represents the response structure for AP global operational data.
type ApGlobalOperResponse struct {
	ApHistory             []ApHistory           `json:"ap-history"`
	EwlcApStats           EwlcApStats           `json:"ewlc-ap-stats"`
	ApImgPredownloadStats ApImgPredownloadStats `json:"ap-img-predownload-stats"`
	ApJoinStats           []ApJoinStats         `json:"ap-join-stats"`
	WlanClientStats       []WlanClientStats     `json:"wlan-client-stats"`
	EmltdJoinCountStat    EmltdJoinCountStat    `json:"emltd-join-count-stat"`
}

// ApGlobalOperApHistoryResponse represents the response structure for AP history data.
type ApGlobalOperApHistoryResponse struct {
	ApHistory []ApHistory `json:"Cisco-IOS-XE-wireless-ap-global-oper:ap-history"`
}

// ApGlobalOperEwlcApStatsResponse represents the response structure for EWLC AP statistics.
type ApGlobalOperEwlcApStatsResponse struct {
	EwlcApStats EwlcApStats `json:"Cisco-IOS-XE-wireless-ap-global-oper:ewlc-ap-stats"`
}

// ApGlobalOperApImgPredownloadStatsResponse represents the response structure for AP image predownload statistics.
type ApGlobalOperApImgPredownloadStatsResponse struct {
	ApImgPredownloadStats ApImgPredownloadStats `json:"Cisco-IOS-XE-wireless-ap-global-oper:ap-img-predownload-stats"`
}

// ApGlobalOperApJoinStatsResponse represents the response structure for AP join statistics.
type ApGlobalOperApJoinStatsResponse struct {
	ApJoinStats []ApJoinStats `json:"Cisco-IOS-XE-wireless-ap-global-oper:ap-join-stats"`
}

// ApGlobalOperWlanClientStatsResponse represents the response structure for WLAN client statistics.
type ApGlobalOperWlanClientStatsResponse struct {
	WlanClientStats []WlanClientStats `json:"Cisco-IOS-XE-wireless-ap-global-oper:wlan-client-stats"`
}

// ApGlobalOperEmltdJoinCountStatResponse represents the response structure for EMLTD join count statistics.
type ApGlobalOperEmltdJoinCountStatResponse struct {
	EmltdJoinCountStat EmltdJoinCountStat `json:"Cisco-IOS-XE-wireless-ap-global-oper:emltd-join-count-stat"`
}

type ApHistory struct {
	EthernetMac    string              `json:"ethernet-mac"`
	ApName         string              `json:"ap-name"`
	WtpMac         string              `json:"wtp-mac"`
	EwlcApStatePtr []EwlcApStateRecord `json:"ewlc-ap-state-ptr"`
}

type EwlcApStateRecord struct {
	IsApJoined              bool      `json:"is-ap-joined"`
	Timestamp               time.Time `json:"timestamp"`
	LastDisconnectTimestamp time.Time `json:"last-disconnect-timestamp"`
	Disconnects             int       `json:"disconnects"`
	ApDisconnectReasonStr   string    `json:"ap-disconnect-reason-str"`
}

type EwlcApStats struct {
	Stats80211ARad         RadioStats `json:"stats-80211-a-rad"`
	Stats80211BgRad        RadioStats `json:"stats-80211-bg-rad"`
	Stats80211XorRad       RadioStats `json:"stats-80211-xor-rad"`
	Stats80211RxOnlyRad    RadioStats `json:"stats-80211-rx-only-rad"`
	Stats80211AllRad       RadioStats `json:"stats-80211-all-rad"`
	Stats80211BgClntSrvg   RadioStats `json:"stats-80211bg-clnt-srvg"`
	Stats80211AClntSrvg    RadioStats `json:"stats-80211a-clnt-srvg"`
	StatsRadMonMode        RadioStats `json:"stats-rad-mon-mode"`
	StatsMisconfiguredAps  int        `json:"stats-misconfigured-aps"`
	Stats802116GhzRadios   RadioStats `json:"stats-80211-6ghz-radios"`
	Stats802116GhzClntSrvg RadioStats `json:"stats-80211-6ghz-clnt-srvg"`
	DualBandRadMonMode     RadioStats `json:"dual-band-rad-mon-mode"`
	Stats80211BgRadMonMode RadioStats `json:"stats-80211bg-rad-mon-mode"`
	Stats80211ARadMonMode  RadioStats `json:"stats-80211a-rad-mon-mode"`
	RadMonMode802116Ghz    RadioStats `json:"rad-mon-mode-80211-6ghz"`
	StatsDtlsLscFbkAps     int        `json:"stats-dtls-lsc-fbk-aps"`
	TotalHighCPUReload     int        `json:"total-high-cpu-reload"`
	TotalHighMemReload     int        `json:"total-high-mem-reload"`
	TotalRadioStuckReset   int        `json:"total-radio-stuck-reset"`
	DualBandRadSnfrMode    RadioStats `json:"dual-band-rad-snfr-mode"`
	RadioSnfrMode80211Bg   RadioStats `json:"radio-snfr-mode-80211bg"`
	RadioSnfrMode80211A    RadioStats `json:"radio-snfr-mode-80211a"`
	RadioSnfrMode802116Ghz RadioStats `json:"radio-snfr-mode-80211-6ghz"`
	RadioSnfrMode          RadioStats `json:"radio-snfr-mode"`
	Total80211Xor56GhzRad  RadioStats `json:"total-80211-xor-5-6ghz-rad"`
}

type RadioStats struct {
	TotalRadios int `json:"total-radios"`
	RadiosUp    int `json:"radios-up"`
	RadiosDown  int `json:"radios-down"`
}

type ApImgPredownloadStats struct {
	PredownloadStats struct {
		NumInitiated            int  `json:"num-initiated"`
		NumInProgress           int  `json:"num-in-progress"`
		NumComplete             int  `json:"num-complete"`
		NumUnsupported          int  `json:"num-unsupported"`
		NumFailed               int  `json:"num-failed"`
		IsPredownloadInProgress bool `json:"is-predownload-in-progress"`
		NumTotal                int  `json:"num-total"`
	} `json:"predownload-stats"`
	DownloadsInProgress int `json:"downloads-in-progress"`
	DownloadsComplete   int `json:"downloads-complete"`
	WlcPredownloadStats struct {
		NumInitiated            int  `json:"num-initiated"`
		NumInProgress           int  `json:"num-in-progress"`
		NumComplete             int  `json:"num-complete"`
		NumUnsupported          int  `json:"num-unsupported"`
		NumFailed               int  `json:"num-failed"`
		IsPredownloadInProgress bool `json:"is-predownload-in-progress"`
		NumTotal                int  `json:"num-total"`
	} `json:"wlc-predownload-stats"`
}

type ApJoinStats struct {
	WtpMac             string          `json:"wtp-mac"`
	ApJoinInfo         ApJoinInfo      `json:"ap-join-info"`
	ApDiscoveryInfo    ApDiscoveryInfo `json:"ap-discovery-info"`
	DtlsSessInfo       DtlsSessInfo    `json:"dtls-sess-info"`
	ApDisconnectReason string          `json:"ap-disconnect-reason"`
	RebootReason       string          `json:"reboot-reason"`
	DisconnectReason   string          `json:"disconnect-reason"`
}

type ApJoinInfo struct {
	ApIPAddr              string    `json:"ap-ip-addr"`
	ApEthernetMac         string    `json:"ap-ethernet-mac"`
	ApName                string    `json:"ap-name"`
	IsJoined              bool      `json:"is-joined"`
	NumJoinReqRecvd       int       `json:"num-join-req-recvd"`
	NumConfigReqRecvd     int       `json:"num-config-req-recvd"`
	LastJoinFailureType   string    `json:"last-join-failure-type"`
	LastConfigFailureType string    `json:"last-config-failure-type"`
	LastErrorType         string    `json:"last-error-type"`
	LastErrorTime         time.Time `json:"last-error-time"`
	LastMsgDecrFailReason string    `json:"last-msg-decr-fail-reason"`
	NumSuccJoinRespSent   int       `json:"num-succ-join-resp-sent"`
	NumUnsuccJoinReqProcn int       `json:"num-unsucc-join-req-procn"`
	NumSuccConfRespSent   int       `json:"num-succ-conf-resp-sent"`
	NumUnsuccConfReqProcn int       `json:"num-unsucc-conf-req-procn"`
	LastSuccJoinAtmptTime time.Time `json:"last-succ-join-atmpt-time"`
	LastFailJoinAtmptTime time.Time `json:"last-fail-join-atmpt-time"`
	LastSuccConfAtmptTime time.Time `json:"last-succ-conf-atmpt-time"`
	LastFailConfAtmptTime time.Time `json:"last-fail-conf-atmpt-time"`
}

type ApDiscoveryInfo struct {
	WtpMac               string    `json:"wtp-mac"`
	EthernetMac          string    `json:"ethernet-mac"`
	ApIPAddress          string    `json:"ap-ip-address"`
	NumDiscoveryReqRecvd int       `json:"num-discovery-req-recvd"`
	NumSuccDiscRespSent  int       `json:"num-succ-disc-resp-sent"`
	NumErrDiscReq        int       `json:"num-err-disc-req"`
	LastDiscFailureType  string    `json:"last-disc-failure-type"`
	LastSuccessDiscTime  time.Time `json:"last-success-disc-time"`
	LastFailedDiscTime   time.Time `json:"last-failed-disc-time"`
}

type DtlsSessInfo struct {
	MacAddr               string    `json:"mac-addr"`
	DataDtlsSetupReq      int       `json:"data-dtls-setup-req"`
	DataDtlsSuccess       int       `json:"data-dtls-success"`
	DataDtlsFailure       int       `json:"data-dtls-failure"`
	CtrlDtlsSetupReq      int       `json:"ctrl-dtls-setup-req"`
	CtrlDtlsSuccess       int       `json:"ctrl-dtls-success"`
	CtrlDtlsFailure       int       `json:"ctrl-dtls-failure"`
	DataDtlsFailureType   string    `json:"data-dtls-failure-type"`
	CtrlDtlsFailureType   string    `json:"ctrl-dtls-failure-type"`
	CtrlDtlsDecryptErr    int       `json:"ctrl-dtls-decrypt-err"`
	CtrlDtlsAntiReplayErr int       `json:"ctrl-dtls-anti-replay-err"`
	DataDtlsDecryptErr    int       `json:"data-dtls-decrypt-err"`
	DataDtlsAntiReplayErr int       `json:"data-dtls-anti-replay-err"`
	DataDtlsFailureTime   time.Time `json:"data-dtls-failure-time"`
	DataDtlsSuccessTime   time.Time `json:"data-dtls-success-time"`
	CtrlDtlsFailureTime   time.Time `json:"ctrl-dtls-failure-time"`
	CtrlDtlsSuccessTime   time.Time `json:"ctrl-dtls-success-time"`
}

type WlanClientStats struct {
	WlanID                  int    `json:"wlan-id"`
	WlanProfileName         string `json:"wlan-profile-name"`
	DataUsage               string `json:"data-usage"`
	TotalRandomMacClients   int    `json:"total-random-mac-clients"`
	ClientCurrStateL2Auth   int    `json:"client-curr-state-l2auth"`
	ClientCurrStateMobility int    `json:"client-curr-state-mobility"`
	ClientCurrStateIplearn  int    `json:"client-curr-state-iplearn"`
	CurrStateWebauthPending int    `json:"curr-state-webauth-pending"`
	ClientCurrStateRun      int    `json:"client-curr-state-run"`
}

type EmltdJoinCountStat struct {
	JoinedApsCount int `json:"joined-aps-count"`
}

// GetApGlobalOper retrieves the entire AP global operational data.
func GetApGlobalOper(client *wnc.Client, ctx context.Context) (*ApGlobalOperResponse, error) {
	var data ApGlobalOperResponse
	if err := client.SendAPIRequest(ctx, ApGlobalOperEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetApHistory retrieves AP history data.
func GetApHistory(client *wnc.Client, ctx context.Context) (*ApGlobalOperApHistoryResponse, error) {
	var data ApGlobalOperApHistoryResponse
	if err := client.SendAPIRequest(ctx, ApHistoryEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetApEwlcApStats retrieves EWLC AP statistics.
func GetApEwlcApStats(client *wnc.Client, ctx context.Context) (*ApGlobalOperEwlcApStatsResponse, error) {
	var data ApGlobalOperEwlcApStatsResponse
	if err := client.SendAPIRequest(ctx, EwlcApStatsEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetApImgPredownloadStats retrieves AP image predownload statistics.
func GetApImgPredownloadStats(client *wnc.Client, ctx context.Context) (*ApGlobalOperApImgPredownloadStatsResponse, error) {
	var data ApGlobalOperApImgPredownloadStatsResponse
	if err := client.SendAPIRequest(ctx, ApImgPredownloadStatsEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetApJoinStats retrieves AP join statistics.
func GetApJoinStats(client *wnc.Client, ctx context.Context) (*ApGlobalOperApJoinStatsResponse, error) {
	var data ApGlobalOperApJoinStatsResponse
	if err := client.SendAPIRequest(ctx, ApJoinStatsEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetApWlanClientStats retrieves WLAN client statistics.
func GetApWlanClientStats(client *wnc.Client, ctx context.Context) (*ApGlobalOperWlanClientStatsResponse, error) {
	var data ApGlobalOperWlanClientStatsResponse
	if err := client.SendAPIRequest(ctx, WlanClientStatsEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetApEmltdJoinCountStat retrieves EMLTD join count statistics.
func GetApEmltdJoinCountStat(client *wnc.Client, ctx context.Context) (*ApGlobalOperEmltdJoinCountStatResponse, error) {
	var data ApGlobalOperEmltdJoinCountStatResponse
	if err := client.SendAPIRequest(ctx, EmltdJoinCountStatEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}
