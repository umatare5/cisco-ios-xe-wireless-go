package ap

import "time"

// ApGlobalOper represents the structure for AP global operational data from WNC 17.12.5.
type ApGlobalOper struct {
	ApHistory             []ApHistory           `json:"ap-history"` // AP history data
	EwlcApStats           EwlcApStats           `json:"ewlc-ap-stats"`
	ApImgPredownloadStats ApImgPredownloadStats `json:"ap-img-predownload-stats"`
	ApLocationStats       []ApLocationStats     `json:"ap-location-stats"` // AP location statistics (YANG: IOS-XE 17.12.1+)
	ApJoinStats           []ApJoinStats         `json:"ap-join-stats"`
	WlanClientStats       []WlanClientStats     `json:"wlan-client-stats"`
	EmltdJoinCountStat    EmltdJoinCountStat    `json:"emltd-join-count-stat"`
}

// ApGlobalOperApHistory represents the structure for AP history data.
type ApGlobalOperApHistory struct {
	ApHistory []ApHistory `json:"Cisco-IOS-XE-wireless-ap-global-oper:ap-history"`
}

// ApGlobalOperEwlcApStats represents the structure for EWLC AP statistics.
type ApGlobalOperEwlcApStats struct {
	EwlcApStats EwlcApStats `json:"Cisco-IOS-XE-wireless-ap-global-oper:ewlc-ap-stats"`
}

// ApGlobalOperApImgPredownloadStats represents the structure for AP image predownload statistics.
type ApGlobalOperApImgPredownloadStats struct {
	ApImgPredownloadStats ApImgPredownloadStats `json:"Cisco-IOS-XE-wireless-ap-global-oper:ap-img-predownload-stats"`
}

// ApGlobalOperApJoinStats represents the structure for AP join statistics.
type ApGlobalOperApJoinStats struct {
	ApJoinStats []ApJoinStats `json:"Cisco-IOS-XE-wireless-ap-global-oper:ap-join-stats"`
}

// ApGlobalOperWlanClientStats represents the structure for WLAN client statistics.
type ApGlobalOperWlanClientStats struct {
	WlanClientStats []WlanClientStats `json:"Cisco-IOS-XE-wireless-ap-global-oper:wlan-client-stats"`
}

// ApGlobalOperApLocationStats represents the structure for AP location statistics.
type ApGlobalOperApLocationStats struct {
	ApLocationStats []ApLocationStats `json:"Cisco-IOS-XE-wireless-ap-global-oper:ap-location-stats"`
}

// ApGlobalOperEmltdJoinCountStat represents the structure for EMLTD join count statistics.
type ApGlobalOperEmltdJoinCountStat struct {
	EmltdJoinCountStat EmltdJoinCountStat `json:"Cisco-IOS-XE-wireless-ap-global-oper:emltd-join-count-stat"`
}

// ApGlobalOperData represents the top-level wrapper for AP global operational data.
type ApGlobalOperData struct {
	ApGlobalOper ApGlobalOper `json:"Cisco-IOS-XE-wireless-ap-global-oper:ap-global-oper-data"`
}

// ApLocationStats represents AP location statistics data.
type ApLocationStats struct {
	Location      string `json:"location"`       // AP location name
	ClientsJoined uint64 `json:"clients-joined"` // Number of clients joined (YANG: IOS-XE 17.12.1+)
	ClientsOn11a  uint64 `json:"clients-on-11a"` // Number of clients joined on 11a (YANG: IOS-XE 17.12.1+)
	ClientsOn11b  uint64 `json:"clients-on-11b"` // Number of clients joined on 11b (YANG: IOS-XE 17.12.1+)
	ApsJoined     uint64 `json:"aps-joined"`     // Number of APs joined (YANG: IOS-XE 17.12.1+)
}

// ApHistory represents AP historical state information.
type ApHistory struct {
	EthernetMac    string              `json:"ethernet-mac"`      // AP ethernet MAC address
	ApName         string              `json:"ap-name"`           // Access point name
	WtpMac         string              `json:"wtp-mac"`           // WTP MAC address
	EwlcApStatePtr []EwlcApStateRecord `json:"ewlc-ap-state-ptr"` // EWLC AP state records
}

// EwlcApStateRecord represents EWLC AP state record information.
type EwlcApStateRecord struct {
	IsApJoined              bool      `json:"is-ap-joined"`              // AP join status flag
	Timestamp               time.Time `json:"timestamp"`                 // State record timestamp
	LastDisconnectTimestamp time.Time `json:"last-disconnect-timestamp"` // Last disconnect time
	Disconnects             int       `json:"disconnects"`               // Number of disconnections
	ApDisconnectReasonStr   string    `json:"ap-disconnect-reason-str"`  // Disconnect reason description
}

// EwlcApStats represents EWLC AP statistics data.
type EwlcApStats struct {
	Stats80211ARad         RadioStats `json:"stats-80211-a-rad"`          // 802.11a radio statistics
	Stats80211BgRad        RadioStats `json:"stats-80211-bg-rad"`         // 802.11bg radio statistics
	Stats80211XorRad       RadioStats `json:"stats-80211-xor-rad"`        // 802.11 XOR radio statistics
	Stats80211RxOnlyRad    RadioStats `json:"stats-80211-rx-only-rad"`    // 802.11 RX-only radio statistics
	Stats80211AllRad       RadioStats `json:"stats-80211-all-rad"`        // 802.11 all radio statistics
	Stats80211BgClntSrvg   RadioStats `json:"stats-80211bg-clnt-srvg"`    // 802.11bg client serving statistics
	Stats80211AClntSrvg    RadioStats `json:"stats-80211a-clnt-srvg"`     // 802.11a client serving statistics
	StatsRadMonMode        RadioStats `json:"stats-rad-mon-mode"`         // Radio monitor mode statistics
	StatsMisconfiguredAps  int        `json:"stats-misconfigured-aps"`    // Number of misconfigured APs
	Stats802116GhzRadios   RadioStats `json:"stats-80211-6ghz-radios"`    // 802.11 6GHz radio statistics
	Stats802116GhzClntSrvg RadioStats `json:"stats-80211-6ghz-clnt-srvg"` // 802.11 6GHz client serving statistics
	DualBandRadMonMode     RadioStats `json:"dual-band-rad-mon-mode"`     // Dual-band radio monitor mode statistics
	Stats80211BgRadMonMode RadioStats `json:"stats-80211bg-rad-mon-mode"` // 802.11bg radio monitor mode statistics
	Stats80211ARadMonMode  RadioStats `json:"stats-80211a-rad-mon-mode"`  // 802.11a radio monitor mode statistics
	RadMonMode802116Ghz    RadioStats `json:"rad-mon-mode-80211-6ghz"`    // 6GHz radio monitor mode statistics
	StatsDtlsLscFbkAps     int        `json:"stats-dtls-lsc-fbk-aps"`     // DTLS LSC fallback APs count
	TotalHighCPUReload     int        `json:"total-high-cpu-reload"`      // High CPU reload events count
	TotalHighMemReload     int        `json:"total-high-mem-reload"`      // High memory reload events count
	TotalRadioStuckReset   int        `json:"total-radio-stuck-reset"`    // Radio stuck reset events count
	DualBandRadSnfrMode    RadioStats `json:"dual-band-rad-snfr-mode"`    // Dual-band radio sniffer mode statistics
	RadioSnfrMode80211Bg   RadioStats `json:"radio-snfr-mode-80211bg"`    // 802.11bg radio sniffer mode statistics
	RadioSnfrMode80211A    RadioStats `json:"radio-snfr-mode-80211a"`     // 802.11a radio sniffer mode statistics
	RadioSnfrMode802116Ghz RadioStats `json:"radio-snfr-mode-80211-6ghz"` // 6GHz radio sniffer mode statistics
	RadioSnfrMode          RadioStats `json:"radio-snfr-mode"`            // Radio sniffer mode statistics
	Total80211Xor56GhzRad  RadioStats `json:"total-80211-xor-5-6ghz-rad"` // 802.11 XOR 5/6GHz radio statistics
}

// RadioStats represents radio statistics information.
type RadioStats struct {
	TotalRadios int `json:"total-radios"` // Total number of radios
	RadiosUp    int `json:"radios-up"`    // Number of radios in up state
	RadiosDown  int `json:"radios-down"`  // Number of radios in down state
}

// ApImgPredownloadStats represents AP image predownload statistics.
type ApImgPredownloadStats struct {
	PredownloadStats struct {
		NumInitiated            int  `json:"num-initiated"`              // Number of initiated predownloads
		NumInProgress           int  `json:"num-in-progress"`            // Number of predownloads in progress
		NumComplete             int  `json:"num-complete"`               // Number of completed predownloads
		NumUnsupported          int  `json:"num-unsupported"`            // Number of unsupported predownloads
		NumFailed               int  `json:"num-failed"`                 // Number of failed predownloads
		IsPredownloadInProgress bool `json:"is-predownload-in-progress"` // Predownload progress status
		NumTotal                int  `json:"num-total"`                  // Total number of predownloads
	} `json:"predownload-stats"` // AP predownload statistics
	DownloadsInProgress int `json:"downloads-in-progress"` // Current downloads in progress
	DownloadsComplete   int `json:"downloads-complete"`    // Completed downloads count
	WlcPredownloadStats struct {
		NumInitiated            int  `json:"num-initiated"`              // WLC initiated predownloads
		NumInProgress           int  `json:"num-in-progress"`            // WLC predownloads in progress
		NumComplete             int  `json:"num-complete"`               // WLC completed predownloads
		NumUnsupported          int  `json:"num-unsupported"`            // WLC unsupported predownloads
		NumFailed               int  `json:"num-failed"`                 // WLC failed predownloads
		IsPredownloadInProgress bool `json:"is-predownload-in-progress"` // WLC predownload progress status
		NumTotal                int  `json:"num-total"`                  // WLC total predownloads
	} `json:"wlc-predownload-stats"` // WLC predownload statistics
}

// ApJoinStats represents AP join statistics data.
type ApJoinStats struct {
	WtpMac             string          `json:"wtp-mac"`              // WTP MAC address
	ApJoinInfo         ApJoinInfo      `json:"ap-join-info"`         // AP join process information
	ApDiscoveryInfo    ApDiscoveryInfo `json:"ap-discovery-info"`    // AP discovery process information
	DtlsSessInfo       DtlsSessInfo    `json:"dtls-sess-info"`       // DTLS session information
	ApDisconnectReason string          `json:"ap-disconnect-reason"` // AP disconnect reason
	RebootReason       string          `json:"reboot-reason"`        // AP reboot reason
	DisconnectReason   string          `json:"disconnect-reason"`    // Disconnect reason details
}

// ApJoinInfo represents AP join process information.
type ApJoinInfo struct {
	ApIPAddr              string    `json:"ap-ip-addr"`                // AP IP address
	ApEthernetMac         string    `json:"ap-ethernet-mac"`           // AP ethernet MAC address
	ApName                string    `json:"ap-name"`                   // Access point name
	IsJoined              bool      `json:"is-joined"`                 // AP join status flag
	NumJoinReqRecvd       int       `json:"num-join-req-recvd"`        // Number of join requests received
	NumConfigReqRecvd     int       `json:"num-config-req-recvd"`      // Number of config requests received
	LastJoinFailureType   string    `json:"last-join-failure-type"`    // Last join failure type
	LastConfigFailureType string    `json:"last-config-failure-type"`  // Last config failure type
	LastErrorType         string    `json:"last-error-type"`           // Last error type
	LastErrorTime         time.Time `json:"last-error-time"`           // Last error timestamp
	LastMsgDecrFailReason string    `json:"last-msg-decr-fail-reason"` // Last message decryption failure reason
	NumSuccJoinRespSent   int       `json:"num-succ-join-resp-sent"`   // Number of successful join responses sent
	NumUnsuccJoinReqProcn int       `json:"num-unsucc-join-req-procn"` // Number of unsuccessful join request processing
	NumSuccConfRespSent   int       `json:"num-succ-conf-resp-sent"`   // Number of successful config responses sent
	NumUnsuccConfReqProcn int       `json:"num-unsucc-conf-req-procn"` // Number of unsuccessful config request processing
	LastSuccJoinAtmptTime time.Time `json:"last-succ-join-atmpt-time"` // Last successful join attempt time
	LastFailJoinAtmptTime time.Time `json:"last-fail-join-atmpt-time"` // Last failed join attempt time
	LastSuccConfAtmptTime time.Time `json:"last-succ-conf-atmpt-time"` // Last successful config attempt time
	LastFailConfAtmptTime time.Time `json:"last-fail-conf-atmpt-time"` // Last failed config attempt time
}

// ApDiscoveryInfo represents AP discovery process information.
type ApDiscoveryInfo struct {
	WtpMac               string    `json:"wtp-mac"`                 // WTP MAC address
	EthernetMac          string    `json:"ethernet-mac"`            // Ethernet MAC address
	ApIPAddress          string    `json:"ap-ip-address"`           // AP IP address
	NumDiscoveryReqRecvd int       `json:"num-discovery-req-recvd"` // Number of discovery requests received
	NumSuccDiscRespSent  int       `json:"num-succ-disc-resp-sent"` // Number of successful discovery responses sent
	NumErrDiscReq        int       `json:"num-err-disc-req"`        // Number of error discovery requests
	LastDiscFailureType  string    `json:"last-disc-failure-type"`  // Last discovery failure type
	LastSuccessDiscTime  time.Time `json:"last-success-disc-time"`  // Last successful discovery time
	LastFailedDiscTime   time.Time `json:"last-failed-disc-time"`   // Last failed discovery time
}

// DtlsSessInfo represents DTLS session information.
type DtlsSessInfo struct {
	MacAddr               string    `json:"mac-addr"`                  // MAC address
	DataDtlsSetupReq      int       `json:"data-dtls-setup-req"`       // Data DTLS setup requests
	DataDtlsSuccess       int       `json:"data-dtls-success"`         // Data DTLS successful connections
	DataDtlsFailure       int       `json:"data-dtls-failure"`         // Data DTLS failed connections
	CtrlDtlsSetupReq      int       `json:"ctrl-dtls-setup-req"`       // Control DTLS setup requests
	CtrlDtlsSuccess       int       `json:"ctrl-dtls-success"`         // Control DTLS successful connections
	CtrlDtlsFailure       int       `json:"ctrl-dtls-failure"`         // Control DTLS failed connections
	DataDtlsFailureType   string    `json:"data-dtls-failure-type"`    // Data DTLS failure type
	CtrlDtlsFailureType   string    `json:"ctrl-dtls-failure-type"`    // Control DTLS failure type
	CtrlDtlsDecryptErr    int       `json:"ctrl-dtls-decrypt-err"`     // Control DTLS decryption errors
	CtrlDtlsAntiReplayErr int       `json:"ctrl-dtls-anti-replay-err"` // Control DTLS anti-replay errors
	DataDtlsDecryptErr    int       `json:"data-dtls-decrypt-err"`     // Data DTLS decryption errors
	DataDtlsAntiReplayErr int       `json:"data-dtls-anti-replay-err"` // Data DTLS anti-replay errors
	DataDtlsFailureTime   time.Time `json:"data-dtls-failure-time"`    // Data DTLS failure timestamp
	DataDtlsSuccessTime   time.Time `json:"data-dtls-success-time"`    // Data DTLS success timestamp
	CtrlDtlsFailureTime   time.Time `json:"ctrl-dtls-failure-time"`    // Control DTLS failure timestamp
	CtrlDtlsSuccessTime   time.Time `json:"ctrl-dtls-success-time"`    // Control DTLS success timestamp
}

// WlanClientStats represents WLAN client statistics data.
type WlanClientStats struct {
	WlanID                  int    `json:"wlan-id"`                    // WLAN identifier
	WlanProfileName         string `json:"wlan-profile-name"`          // WLAN profile name
	DataUsage               string `json:"data-usage"`                 // Data usage in bytes
	TotalRandomMacClients   int    `json:"total-random-mac-clients"`   // Total random MAC clients
	ClientCurrStateL2Auth   int    `json:"client-curr-state-l2auth"`   // Clients in L2 authentication state
	ClientCurrStateMobility int    `json:"client-curr-state-mobility"` // Clients in mobility state
	ClientCurrStateIplearn  int    `json:"client-curr-state-iplearn"`  // Clients in IP learning state
	CurrStateWebauthPending int    `json:"curr-state-webauth-pending"` // Clients in web auth pending state
	ClientCurrStateRun      int    `json:"client-curr-state-run"`      // Clients in running state
}

// EmltdJoinCountStat represents emulated join count statistics.
type EmltdJoinCountStat struct {
	JoinedApsCount int `json:"joined-aps-count"` // Number of joined APs in emulated mode
}
