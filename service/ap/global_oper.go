package ap

import "time"

// CiscoIOSXEWirelessAPGlobalOper represents the top-level wrapper for AP global operational data.
type CiscoIOSXEWirelessAPGlobalOper struct {
	ApGlobalOperData struct {
		ApHistory             []ApHistory           `json:"ap-history"`                  // AP history data for state tracking (Live: IOS-XE 17.12.6a)
		EwlcApStats           EwlcApStats           `json:"ewlc-ap-stats"`               // AP radio statistics (Live: IOS-XE 17.12.6a)
		ApImgPredownloadStats ApImgPredownloadStats `json:"ap-img-predownload-stats"`    // AP image predownload stats (Live: IOS-XE 17.12.6a)
		ApLocationStats       []ApLocationStats     `json:"ap-location-stats,omitempty"` // AP location statistics (YANG: IOS-XE 17.12.1)
		ApJoinStats           []ApJoinStats         `json:"ap-join-stats"`               // AP join statistics (Live: IOS-XE 17.12.6a)
		WlanClientStats       []WlanClientStats     `json:"wlan-client-stats"`           // WLAN client stats (Live: IOS-XE 17.12.6a)
		EmltdJoinCountStat    EmltdJoinCountStat    `json:"emltd-join-count-stat"`       // AP join count statistics (Live: IOS-XE 17.12.6a)
	} `json:"Cisco-IOS-XE-wireless-ap-global-oper:ap-global-oper-data"`
}

// CiscoIOSXEWirelessApGlobalOperApHistory represents the structure for AP history data.
type CiscoIOSXEWirelessApGlobalOperApHistory struct {
	ApHistory []ApHistory `json:"Cisco-IOS-XE-wireless-ap-global-oper:ap-history"`
}

// CiscoIOSXEWirelessApGlobalOperEwlcApStats represents the structure for EWLC AP statistics.
type CiscoIOSXEWirelessApGlobalOperEwlcApStats struct {
	EwlcApStats EwlcApStats `json:"Cisco-IOS-XE-wireless-ap-global-oper:ewlc-ap-stats"`
}

// CiscoIOSXEWirelessApGlobalOperApImgPredownloadStats represents the structure for AP image predownload statistics.
type CiscoIOSXEWirelessApGlobalOperApImgPredownloadStats struct {
	ApImgPredownloadStats ApImgPredownloadStats `json:"Cisco-IOS-XE-wireless-ap-global-oper:ap-img-predownload-stats"`
}

// CiscoIOSXEWirelessApGlobalOperApJoinStats represents the structure for AP join statistics.
type CiscoIOSXEWirelessApGlobalOperApJoinStats struct {
	ApJoinStats []ApJoinStats `json:"Cisco-IOS-XE-wireless-ap-global-oper:ap-join-stats"`
}

// CiscoIOSXEWirelessApGlobalOperWlanClientStats represents the structure for WLAN client statistics.
type CiscoIOSXEWirelessApGlobalOperWlanClientStats struct {
	WlanClientStats []WlanClientStats `json:"Cisco-IOS-XE-wireless-ap-global-oper:wlan-client-stats"`
}

// CiscoIOSXEWirelessApGlobalOperApLocationStats represents the structure for AP location statistics.
type CiscoIOSXEWirelessApGlobalOperApLocationStats struct {
	ApLocationStats []ApLocationStats `json:"Cisco-IOS-XE-wireless-ap-global-oper:ap-location-stats"`
}

// CiscoIOSXEWirelessApGlobalOperEmltdJoinCountStat represents the structure for EMLTD join count statistics.
type CiscoIOSXEWirelessApGlobalOperEmltdJoinCountStat struct {
	EmltdJoinCountStat EmltdJoinCountStat `json:"Cisco-IOS-XE-wireless-ap-global-oper:emltd-join-count-stat"`
}

// ApLocationStats represents AP location statistics data.
type ApLocationStats struct {
	Location      string `json:"location"`       // AP location name (YANG: IOS-XE 17.12.1)
	ClientsJoined uint64 `json:"clients-joined"` // Number of clients joined (YANG: IOS-XE 17.12.1)
	ClientsOn11a  uint64 `json:"clients-on-11a"` // Number of clients joined on 11a (YANG: IOS-XE 17.12.1)
	ClientsOn11b  uint64 `json:"clients-on-11b"` // Number of clients joined on 11b (YANG: IOS-XE 17.12.1)
	ApsJoined     uint64 `json:"aps-joined"`     // Number of APs joined (YANG: IOS-XE 17.12.1)
}

// ApHistory represents AP historical state information.
type ApHistory struct {
	EthernetMAC    string              `json:"ethernet-mac"`      // Ethernet MAC address (YANG: IOS-XE 17.12.1)
	ApName         string              `json:"ap-name"`           // AP name (YANG: IOS-XE 17.12.1)
	WtpMAC         string              `json:"wtp-mac"`           // AP WTP mac (YANG: IOS-XE 17.12.1)
	EwlcApStatePtr []EwlcApStateRecord `json:"ewlc-ap-state-ptr"` // AP state (YANG: IOS-XE 17.12.1)
}

// EwlcApStateRecord represents EWLC AP state record information.
type EwlcApStateRecord struct {
	IsApJoined              bool      `json:"is-ap-joined"`              // AP joined or not (YANG: IOS-XE 17.12.1)
	Timestamp               time.Time `json:"timestamp"`                 // AP Joined or first disjoined timestamp (YANG: IOS-XE 17.12.1)
	LastDisconnectTimestamp time.Time `json:"last-disconnect-timestamp"` // Last disconnect timestamp (YANG: IOS-XE 17.12.1)
	Disconnects             int       `json:"disconnects"`               // Number of times AP disconnected (YANG: IOS-XE 17.12.1)
	ApDisconnectReasonStr   string    `json:"ap-disconnect-reason-str"`  // AP disconnect string (YANG: IOS-XE 17.12.1)
}

// EwlcApStats represents EWLC AP statistics data.
type EwlcApStats struct {
	Stats80211ARad         RadioStats `json:"stats-80211-a-rad"`          // 802.11 5 GHz radio stats (Live: IOS-XE 17.12.6a)
	Stats80211BgRad        RadioStats `json:"stats-80211-bg-rad"`         // 802.11 2.4 GHz radio stats (Live: IOS-XE 17.12.6a)
	Stats80211XorRad       RadioStats `json:"stats-80211-xor-rad"`        // 802.11 dual band radio stats (Live: IOS-XE 17.12.6a)
	Stats80211RxOnlyRad    RadioStats `json:"stats-80211-rx-only-rad"`    // 802.11 RX radio stats (Live: IOS-XE 17.12.6a)
	Stats80211AllRad       RadioStats `json:"stats-80211-all-rad"`        // All radio stats (Live: IOS-XE 17.12.6a)
	Stats80211BgClntSrvg   RadioStats `json:"stats-80211bg-clnt-srvg"`    // 802.11 2.4 GHz client serving radio stats (Live: IOS-XE 17.12.6a)
	Stats80211AClntSrvg    RadioStats `json:"stats-80211a-clnt-srvg"`     // 802.11 5 GHz client serving radio stats (Live: IOS-XE 17.12.6a)
	StatsRadMonMode        RadioStats `json:"stats-rad-mon-mode"`         // Monitor radio stats (Live: IOS-XE 17.12.6a)
	StatsMisconfiguredAps  int        `json:"stats-misconfigured-aps"`    // Total number of misconfigured APs (Live: IOS-XE 17.12.6a)
	Stats802116GhzRadios   RadioStats `json:"stats-80211-6ghz-radios"`    // 802.11 6 GHz radio stats (Live: IOS-XE 17.12.6a)
	Stats802116GhzClntSrvg RadioStats `json:"stats-80211-6ghz-clnt-srvg"` // 802.11 6 GHz client serving radio stats (Live: IOS-XE 17.12.6a)
	DualBandRadMonMode     RadioStats `json:"dual-band-rad-mon-mode"`     // 802.11 dual band monitor radio stats (Live: IOS-XE 17.12.6a)
	Stats80211BgRadMonMode RadioStats `json:"stats-80211bg-rad-mon-mode"` // 802.11 2.4 GHz monitor radio stats (Live: IOS-XE 17.12.6a)
	Stats80211ARadMonMode  RadioStats `json:"stats-80211a-rad-mon-mode"`  // 802.11 5 GHz monitor radio stats (Live: IOS-XE 17.12.6a)
	RadMonMode802116Ghz    RadioStats `json:"rad-mon-mode-80211-6ghz"`    // 802.11 6 GHz monitor radio stats (Live: IOS-XE 17.12.6a)
	StatsDTLSLscFbkAps     int        `json:"stats-dtls-lsc-fbk-aps"`     // Total number of DTLS LSC fallback APs (Live: IOS-XE 17.12.6a)
	TotalHighCPUReload     int        `json:"total-high-cpu-reload"`      // Total number of AP reloads due to high CPU (Live: IOS-XE 17.12.6a)
	TotalHighMemReload     int        `json:"total-high-mem-reload"`      // Total number of AP reloads due to high memory (Live: IOS-XE 17.12.6a)
	TotalRadioStuckReset   int        `json:"total-radio-stuck-reset"`    // Total number of radio stuck resets (Live: IOS-XE 17.12.6a)
	DualBandRadSnfrMode    RadioStats `json:"dual-band-rad-snfr-mode"`    // 802.11 dual band sniffer radio stats (Live: IOS-XE 17.12.6a)
	RadioSnfrMode80211Bg   RadioStats `json:"radio-snfr-mode-80211bg"`    // 802.11 2.4 GHz sniffer radio stats (Live: IOS-XE 17.12.6a)
	RadioSnfrMode80211A    RadioStats `json:"radio-snfr-mode-80211a"`     // 802.11 5 GHz sniffer radio stats (Live: IOS-XE 17.12.6a)
	RadioSnfrMode802116Ghz RadioStats `json:"radio-snfr-mode-80211-6ghz"` // 802.11 6 GHz sniffer radio stats (Live: IOS-XE 17.12.6a)
	RadioSnfrMode          RadioStats `json:"radio-snfr-mode"`            // All sniffer radio stats (Live: IOS-XE 17.12.6a)
	Total80211Xor56GhzRad  RadioStats `json:"total-80211-xor-5-6ghz-rad"` // 802.11 dual band 5/6 GHz radio stats (Live: IOS-XE 17.12.6a)
}

// RadioStats represents radio statistics information.
type RadioStats struct {
	TotalRadios int `json:"total-radios"` // Total number of radios (Live: IOS-XE 17.12.6a)
	RadiosUp    int `json:"radios-up"`    // Total number of radios up (Live: IOS-XE 17.12.6a)
	RadiosDown  int `json:"radios-down"`  // Total number of radios down (Live: IOS-XE 17.12.6a)
}

// ApImgPredownloadStats represents AP image predownload statistics.
type ApImgPredownloadStats struct {
	PredownloadStats struct {
		NumInitiated            int  `json:"num-initiated"`              // Total AP predownload initiated for proactive distribution (Live: IOS-XE 17.12.6a)
		NumInProgress           int  `json:"num-in-progress"`            // Total AP predownload in-progress for bandwidth monitoring (Live: IOS-XE 17.12.6a)
		NumComplete             int  `json:"num-complete"`               // Total AP predownload completed for deployment tracking (Live: IOS-XE 17.12.6a)
		NumUnsupported          int  `json:"num-unsupported"`            // Total AP predownload not supported for compatibility check (Live: IOS-XE 17.12.6a)
		NumFailed               int  `json:"num-failed"`                 // Total AP predownload failed for troubleshooting analysis (Live: IOS-XE 17.12.6a)
		IsPredownloadInProgress bool `json:"is-predownload-in-progress"` // Status of AP image predownload process for monitoring (Live: IOS-XE 17.12.6a)
		NumTotal                int  `json:"num-total"`                  // Total AP connected for comprehensive coverage tracking (Live: IOS-XE 17.12.6a)
	} `json:"predownload-stats"` // AP predownload statistics for firmware management (Live: IOS-XE 17.12.6a)
	DownloadsInProgress int `json:"downloads-in-progress"` // Total APs download in-progress for bandwidth monitoring (Live: IOS-XE 17.12.6a)
	DownloadsComplete   int `json:"downloads-complete"`    // Total APs download completed for deployment tracking (Live: IOS-XE 17.12.6a)
	WlcPredownloadStats struct {
		NumInitiated            int  `json:"num-initiated"`              // Total AP predownload initiated for controller orchestration (Live: IOS-XE 17.12.6a)
		NumInProgress           int  `json:"num-in-progress"`            // Total AP predownload in-progress for resource monitoring (Live: IOS-XE 17.12.6a)
		NumComplete             int  `json:"num-complete"`               // Total AP predownload completed for deployment success (Live: IOS-XE 17.12.6a)
		NumUnsupported          int  `json:"num-unsupported"`            // Total AP predownload not supported for compatibility check (Live: IOS-XE 17.12.6a)
		NumFailed               int  `json:"num-failed"`                 // Total AP predownload failed for troubleshooting recovery (Live: IOS-XE 17.12.6a)
		IsPredownloadInProgress bool `json:"is-predownload-in-progress"` // Status of AP image predownload for coordination monitoring (Live: IOS-XE 17.12.6a)
		NumTotal                int  `json:"num-total"`                  // Total AP connected for controller-wide coordination (Live: IOS-XE 17.12.6a)
	} `json:"wlc-predownload-stats"` // Wireless controller predownload statistics for firmware management (Live: IOS-XE 17.12.6a)
}

// ApJoinStats represents AP join statistics data.
type ApJoinStats struct {
	WtpMAC             string          `json:"wtp-mac"`              // AP radio MAC address for CAPWAP session identification (Live: IOS-XE 17.12.6a)
	ApJoinInfo         ApJoinInfo      `json:"ap-join-info"`         // AP join information for connection status tracking (Live: IOS-XE 17.12.6a)
	ApDiscoveryInfo    ApDiscoveryInfo `json:"ap-discovery-info"`    // AP discovery information for controller identification (Live: IOS-XE 17.12.6a)
	DTLSSessInfo       DTLSSessInfo    `json:"dtls-sess-info"`       // Data and Control DTLS phase statistics for secure tunnel (Live: IOS-XE 17.12.6a)
	ApDisconnectReason string          `json:"ap-disconnect-reason"` // Last disconnect reason of AP for troubleshooting (Live: IOS-XE 17.12.6a)
	RebootReason       string          `json:"reboot-reason"`        // Reboot reason from AP for system stability analysis (Live: IOS-XE 17.12.6a)
	DisconnectReason   string          `json:"disconnect-reason"`    // Disconnect reason from AP for failure diagnosis (Live: IOS-XE 17.12.6a)
}

// ApJoinInfo represents AP join process information.
type ApJoinInfo struct {
	ApIPAddr              string    `json:"ap-ip-addr"`                // IP address of the AP for network connectivity (Live: IOS-XE 17.12.6a)
	ApEthernetMAC         string    `json:"ap-ethernet-mac"`           // AP ethernet MAC address for identification (Live: IOS-XE 17.12.6a)
	ApName                string    `json:"ap-name"`                   // Name of the AP for administrative identification (Live: IOS-XE 17.12.6a)
	IsJoined              bool      `json:"is-joined"`                 // AP join status flag indicating controller association (Live: IOS-XE 17.12.6a)
	NumJoinReqRecvd       int       `json:"num-join-req-recvd"`        // Total number of join requests received from AP (Live: IOS-XE 17.12.6a)
	NumConfigReqRecvd     int       `json:"num-config-req-recvd"`      // Total number of configuration requests received (Live: IOS-XE 17.12.6a)
	LastJoinFailureType   string    `json:"last-join-failure-type"`    // Last AP join failure reason for troubleshooting (Live: IOS-XE 17.12.6a)
	LastConfigFailureType string    `json:"last-config-failure-type"`  // Last AP config failure reason for diagnosis (Live: IOS-XE 17.12.6a)
	LastErrorType         string    `json:"last-error-type"`           // Last failure phase of AP connection for analysis (Live: IOS-XE 17.12.6a)
	LastErrorTime         time.Time `json:"last-error-time"`           // Time at which the last join error occurred (Live: IOS-XE 17.12.6a)
	LastMsgDecrFailReason string    `json:"last-msg-decr-fail-reason"` // Reason for last message decryption failure (Live: IOS-XE 17.12.6a)
	NumSuccJoinRespSent   int       `json:"num-succ-join-resp-sent"`   // Total number of successful join response sent (Live: IOS-XE 17.12.6a)
	NumUnsuccJoinReqProcn int       `json:"num-unsucc-join-req-procn"` // Total number of unsuccessful join request processed (Live: IOS-XE 17.12.6a)
	NumSuccConfRespSent   int       `json:"num-succ-conf-resp-sent"`   // Total number of successful configure response sent (Live: IOS-XE 17.12.6a)
	NumUnsuccConfReqProcn int       `json:"num-unsucc-conf-req-procn"` // Total number of unsuccessful config request processed (Live: IOS-XE 17.12.6a)
	LastSuccJoinAtmptTime time.Time `json:"last-succ-join-atmpt-time"` // Last successful join attempt time for baseline (Live: IOS-XE 17.12.6a)
	LastFailJoinAtmptTime time.Time `json:"last-fail-join-atmpt-time"` // Last join failure time for pattern analysis (Live: IOS-XE 17.12.6a)
	LastSuccConfAtmptTime time.Time `json:"last-succ-conf-atmpt-time"` // Last successful config attempt time for analysis (Live: IOS-XE 17.12.6a)
	LastFailConfAtmptTime time.Time `json:"last-fail-conf-atmpt-time"` // Last failed config attempt time for troubleshooting (Live: IOS-XE 17.12.6a)
}

// ApDiscoveryInfo represents AP discovery process information.
type ApDiscoveryInfo struct {
	WtpMAC               string    `json:"wtp-mac"`                 // AP radio MAC address for CAPWAP discovery session (Live: IOS-XE 17.12.6a)
	EthernetMAC          string    `json:"ethernet-mac"`            // AP ethernet MAC address for network layer discovery (Live: IOS-XE 17.12.6a)
	ApIPAddress          string    `json:"ap-ip-address"`           // AP IP address used during CAPWAP discovery process (Live: IOS-XE 17.12.6a)
	NumDiscoveryReqRecvd int       `json:"num-discovery-req-recvd"` // Total number of discovery request received (Live: IOS-XE 17.12.6a)
	NumSuccDiscRespSent  int       `json:"num-succ-disc-resp-sent"` // Total number of successful discovery response sent (Live: IOS-XE 17.12.6a)
	NumErrDiscReq        int       `json:"num-err-disc-req"`        // Total number of errored discovery requests (Live: IOS-XE 17.12.6a)
	LastDiscFailureType  string    `json:"last-disc-failure-type"`  // Last discovery failure type for troubleshooting (Live: IOS-XE 17.12.6a)
	LastSuccessDiscTime  time.Time `json:"last-success-disc-time"`  // Last successful discovery attempt time for baseline (Live: IOS-XE 17.12.6a)
	LastFailedDiscTime   time.Time `json:"last-failed-disc-time"`   // Last failed discovery attempt time for tracking (Live: IOS-XE 17.12.6a)
}

// DTLSSessInfo represents DTLS session information.
type DTLSSessInfo struct {
	MACAddr               string    `json:"mac-addr"`                  // AP MAC address for DTLS session correlation (Live: IOS-XE 17.12.6a)
	DataDTLSSetupReq      int       `json:"data-dtls-setup-req"`       // DTLS session requests received for data channel (Live: IOS-XE 17.12.6a)
	DataDTLSSuccess       int       `json:"data-dtls-success"`         // Established DTLS session for data channel (Live: IOS-XE 17.12.6a)
	DataDTLSFailure       int       `json:"data-dtls-failure"`         // Unsuccessful DTLS session for data channel (Live: IOS-XE 17.12.6a)
	CtrlDTLSSetupReq      int       `json:"ctrl-dtls-setup-req"`       // DTLS session requests received for control channel (Live: IOS-XE 17.12.6a)
	CtrlDTLSSuccess       int       `json:"ctrl-dtls-success"`         // Established Dtls session for control channel (Live: IOS-XE 17.12.6a)
	CtrlDTLSFailure       int       `json:"ctrl-dtls-failure"`         // Unsuccessful Dtls session for control channel (Live: IOS-XE 17.12.6a)
	DataDTLSFailureType   string    `json:"data-dtls-failure-type"`    // Reason for last unsuccessful DTLS session on data (Live: IOS-XE 17.12.6a)
	CtrlDTLSFailureType   string    `json:"ctrl-dtls-failure-type"`    // Reason for last unsuccessful DTLS session on control (Live: IOS-XE 17.12.6a)
	CtrlDTLSDecryptErr    int       `json:"ctrl-dtls-decrypt-err"`     // SSL decrypt errors for control channel (Live: IOS-XE 17.12.6a)
	CtrlDTLSAntiReplayErr int       `json:"ctrl-dtls-anti-replay-err"` // SSL anti replay errors for control channel (Live: IOS-XE 17.12.6a)
	DataDTLSDecryptErr    int       `json:"data-dtls-decrypt-err"`     // SSL decrypt errors for data channel (Live: IOS-XE 17.12.6a)
	DataDTLSAntiReplayErr int       `json:"data-dtls-anti-replay-err"` // SSL anti replay errors for data channel (Live: IOS-XE 17.12.6a)
	DataDTLSFailureTime   time.Time `json:"data-dtls-failure-time"`    // Last unsuccessful data dtls session time (Live: IOS-XE 17.12.6a)
	DataDTLSSuccessTime   time.Time `json:"data-dtls-success-time"`    // Last successful data dtls session time (Live: IOS-XE 17.12.6a)
	CtrlDTLSFailureTime   time.Time `json:"ctrl-dtls-failure-time"`    // Last unsuccessful control dtls session time (Live: IOS-XE 17.12.6a)
	CtrlDTLSSuccessTime   time.Time `json:"ctrl-dtls-success-time"`    // Last successful control dtls session time (Live: IOS-XE 17.12.6a)
}

// WlanClientStats represents WLAN client statistics data.
type WlanClientStats struct {
	WlanID                  int    `json:"wlan-id"`                    // WLAN identifier for service-specific client tracking (Live: IOS-XE 17.12.6a)
	WlanProfileName         string `json:"wlan-profile-name"`          // WLAN profile name for administrative identification (Live: IOS-XE 17.12.6a)
	DataUsage               string `json:"data-usage"`                 // Data usage statistics for bandwidth monitoring (Live: IOS-XE 17.12.6a)
	TotalRandomMACClients   int    `json:"total-random-mac-clients"`   // Total random MAC clients for privacy protection tracking (Live: IOS-XE 17.12.6a)
	ClientCurrStateL2Auth   int    `json:"client-curr-state-l2auth"`   // Total number of clients in L2-authenticating state (Live: IOS-XE 17.12.6a)
	ClientCurrStateMobility int    `json:"client-curr-state-mobility"` // Total number of clients in mobility state (Live: IOS-XE 17.12.6a)
	ClientCurrStateIplearn  int    `json:"client-curr-state-iplearn"`  // Total number of clients in iplearn state (Live: IOS-XE 17.12.6a)
	CurrStateWebauthPending int    `json:"curr-state-webauth-pending"` // Total number of clients in webauth pending state (Live: IOS-XE 17.12.6a)
	ClientCurrStateRun      int    `json:"client-curr-state-run"`      // Total number of clients in run state (Live: IOS-XE 17.12.6a)
}

// EmltdJoinCountStat represents emulated join count statistics.
type EmltdJoinCountStat struct {
	JoinedApsCount int `json:"joined-aps-count"` // Number of APs joined on wireless LAN controller (Live: IOS-XE 17.12.6a)
}
