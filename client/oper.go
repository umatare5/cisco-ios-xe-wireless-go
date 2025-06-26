// Package client provides client operational data functionality for the Cisco Wireless Network Controller API.
package client

import (
	"context"

	"time"

	wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
)

// Client operational data API endpoints
const (
	// ClientOperBasePath is the base path for client operational data endpoints
	ClientOperBasePath = "/restconf/data/Cisco-IOS-XE-wireless-client-oper:client-oper-data"
	// ClientOperEndpoint retrieves complete client operational data
	ClientOperEndpoint = ClientOperBasePath
	// CommonOperDataEndpoint retrieves common operational data for clients
	CommonOperDataEndpoint = ClientOperBasePath + "/common-oper-data"
	// Dot11OperDataEndpoint retrieves 802.11 operational data for clients
	Dot11OperDataEndpoint = ClientOperBasePath + "/dot11-oper-data"
	// MobilityOperDataEndpoint retrieves mobility operational data for clients
	MobilityOperDataEndpoint = ClientOperBasePath + "/mobility-oper-data"
	// MmIfClientStatsEndpoint retrieves mobility manager interface client statistics
	MmIfClientStatsEndpoint = ClientOperBasePath + "/mm-if-client-stats"
	// MmIfClientHistoryEndpoint retrieves mobility manager interface client history
	MmIfClientHistoryEndpoint = ClientOperBasePath + "/mm-if-client-history"
	// TrafficStatsEndpoint retrieves client traffic statistics
	TrafficStatsEndpoint = ClientOperBasePath + "/traffic-stats"
	// PolicyDataEndpoint retrieves client policy data
	PolicyDataEndpoint = ClientOperBasePath + "/policy-data"
	// SisfDbMacEndpoint retrieves SISF database MAC information
	SisfDbMacEndpoint = ClientOperBasePath + "/sisf-db-mac"
	// DcInfoEndpoint retrieves discovery client information
	DcInfoEndpoint = ClientOperBasePath + "/dc-info"
)

// ClientOperResponse represents the complete client operational data response
type ClientOperResponse struct {
	CiscoIOSXEWirelessClientOperClientOperData struct {
		CommonOperData    []CommonOperData    `json:"common-oper-data"`
		Dot11OperData     []Dot11OperData     `json:"dot11-oper-data"`
		MobilityOperData  []MobilityOperData  `json:"mobility-oper-data"`
		MmIfClientStats   []MmIfClientStats   `json:"mm-if-client-stats"`
		MmIfClientHistory []MmIfClientHistory `json:"mm-if-client-history"`
		TrafficStats      []TrafficStats      `json:"traffic-stats"`
		PolicyData        []PolicyData        `json:"policy-data"`
		SisfDbMac         []SisfDbMac         `json:"sisf-db-mac"`
		DcInfo            []DcInfo            `json:"dc-info"`
	} `json:"Cisco-IOS-XE-wireless-client-oper:client-oper-data"`
}

// ClientOperCommonOperDataResponse represents the common operational data response
type ClientOperCommonOperDataResponse struct {
	CommonOperData []CommonOperData `json:"Cisco-IOS-XE-wireless-client-oper:common-oper-data"`
}

// ClientOperDot11OperDataResponse represents the 802.11 operational data response
type ClientOperDot11OperDataResponse struct {
	Dot11OperData []Dot11OperData `json:"Cisco-IOS-XE-wireless-client-oper:dot11-oper-data"`
}

// ClientOperMobilityOperDataResponse represents the mobility operational data response
type ClientOperMobilityOperDataResponse struct {
	MobilityOperData []MobilityOperData `json:"Cisco-IOS-XE-wireless-client-oper:mobility-oper-data"`
}

// ClientOperMmIfClientStatsResponse represents the MM interface client statistics response
type ClientOperMmIfClientStatsResponse struct {
	MmIfClientStats []MmIfClientStats `json:"Cisco-IOS-XE-wireless-client-oper:mm-if-client-stats"`
}

// ClientOperMmIfClientHistoryResponse represents the MM interface client history response
type ClientOperMmIfClientHistoryResponse struct {
	MmIfClientHistory []MmIfClientHistory `json:"Cisco-IOS-XE-wireless-client-oper:mm-if-client-history"`
}

// ClientOperTrafficStatsResponse represents the client traffic statistics response
type ClientOperTrafficStatsResponse struct {
	TrafficStats []TrafficStats `json:"Cisco-IOS-XE-wireless-client-oper:traffic-stats"`
}

// ClientOperPolicyDataResponse represents the client policy data response
type ClientOperPolicyDataResponse struct {
	PolicyData []PolicyData `json:"Cisco-IOS-XE-wireless-client-oper:policy-data"`
}

// ClientOperSisfDbMacResponse represents the SISF database MAC response
type ClientOperSisfDbMacResponse struct {
	SisfDbMac []SisfDbMac `json:"Cisco-IOS-XE-wireless-client-oper:sisf-db-mac"`
}

// ClientOperDcInfoResponse represents the discovery client information response
type ClientOperDcInfoResponse struct {
	DcInfo []DcInfo `json:"Cisco-IOS-XE-wireless-client-oper:dc-info"`
}

type CommonOperData struct {
	ClientMac             string `json:"client-mac"`
	ApName                string `json:"ap-name"`
	MsApSlotID            int    `json:"ms-ap-slot-id"`
	MsRadioType           string `json:"ms-radio-type"`
	WlanID                int    `json:"wlan-id"`
	ClientType            string `json:"client-type"`
	CoState               string `json:"co-state"`
	AaaOverridePassphrase bool   `json:"aaa-override-passphrase"`
	IsTviEnabled          bool   `json:"is-tvi-enabled"`
	WlanPolicy            struct {
		CurrentSwitchingMode  string `json:"current-switching-mode"`
		WlanSwitchingMode     string `json:"wlan-switching-mode"`
		CentralAuthentication string `json:"central-authentication"`
		CentralDhcp           bool   `json:"central-dhcp"`
		CentralAssocEnable    bool   `json:"central-assoc-enable"`
		VlanCentralSwitching  bool   `json:"vlan-central-switching"`
		IsFabricClient        bool   `json:"is-fabric-client"`
		IsGuestFabricClient   bool   `json:"is-guest-fabric-client"`
		UpnBitFlag            string `json:"upn-bit-flag"`
	} `json:"wlan-policy"`
	Username           string `json:"username"`
	GuestLanClientInfo struct {
		WiredVlan       int `json:"wired-vlan"`
		PhyIfid         int `json:"phy-ifid"`
		IdleTimeSeconds int `json:"idle-time-seconds"`
	} `json:"guest-lan-client-info"`
	MethodID                 string    `json:"method-id"`
	L3VlanOverrideReceived   bool      `json:"l3-vlan-override-received"`
	UpnID                    int       `json:"upn-id"`
	IsLocallyAdministeredMac bool      `json:"is-locally-administered-mac"`
	IdleTimeout              int       `json:"idle-timeout"`
	IdleTimestamp            time.Time `json:"idle-timestamp"`
	ClientDuid               string    `json:"client-duid"`
	VrfName                  string    `json:"vrf-name"`
}

type Dot11OperData struct {
	MsMacAddress        string    `json:"ms-mac-address"`
	Dot11State          string    `json:"dot11-state"`
	MsBssid             string    `json:"ms-bssid"`
	ApMacAddress        string    `json:"ap-mac-address"`
	CurrentChannel      int       `json:"current-channel"`
	MsWlanID            int       `json:"ms-wlan-id"`
	VapSsid             string    `json:"vap-ssid"`
	PolicyProfile       string    `json:"policy-profile"`
	MsApSlotID          int       `json:"ms-ap-slot-id"`
	RadioType           string    `json:"radio-type"`
	MsAssociationID     int       `json:"ms-association-id"`
	MsAuthAlgNum        string    `json:"ms-auth-alg-num"`
	MsReasonCode        string    `json:"ms-reason-code"`
	MsAssocTime         time.Time `json:"ms-assoc-time"`
	Is11GClient         bool      `json:"is-11g-client"`
	MsSupportedRatesStr string    `json:"ms-supported-rates-str"`
	MsWifi              struct {
		WpaVersion           string `json:"wpa-version"`
		CipherSuite          string `json:"cipher-suite"`
		AuthKeyMgmt          string `json:"auth-key-mgmt"`
		GroupMgmtCipherSuite string `json:"group-mgmt-cipher-suite"`
		GroupCipherSuite     string `json:"group-cipher-suite"`
		PweMode              string `json:"pwe-mode"`
	} `json:"ms-wifi"`
	MsWmeEnabled        bool   `json:"ms-wme-enabled"`
	Dot11WEnabled       bool   `json:"dot11w-enabled"`
	EwlcMsPhyType       string `json:"ewlc-ms-phy-type"`
	EncryptionType      string `json:"encryption-type"`
	SecurityMode        string `json:"security-mode"`
	ClientWepPolicyType string `json:"client-wep-policy-type"`
	BssTransCapable     bool   `json:"bss-trans-capable"`
	MsAppleCapable      bool   `json:"ms-apple-capable"`
	WlanProfile         string `json:"wlan-profile"`
	DmsCapable          bool   `json:"dms-capable"`
	EogreClient         struct {
		IsEogre             bool   `json:"is-eogre"`
		PreviousMatchReason string `json:"previous-match-reason"`
		MatchReason         string `json:"match-reason"`
		IsAaaData           bool   `json:"is-aaa-data"`
		Realm               string `json:"realm"`
		Vlan                int    `json:"vlan"`
		Domain              string `json:"domain"`
		PlumbedGw           string `json:"plumbed-gw"`
		TunnelIfid          int    `json:"tunnel-ifid"`
		IsCentralFwd        bool   `json:"is-central-fwd"`
	} `json:"eogre-client"`
	MsHs20Data struct {
		IsHs20                     bool      `json:"is-hs20"`
		Version                    string    `json:"version"`
		ConsortiumOi               string    `json:"consortium-oi"`
		PpsMoID                    int       `json:"pps-mo-id"`
		SwtTimer                   int       `json:"swt-timer"`
		SwtTimestamp               time.Time `json:"swt-timestamp"`
		TermsConditionsURL         string    `json:"terms-conditions-url"`
		SubscriptionRemediationURL string    `json:"subscription-remediation-url"`
		DeauthReasonURL            string    `json:"deauth-reason-url"`
	} `json:"ms-hs20-data"`
	QosmapCapable                   bool   `json:"qosmap-capable"`
	RmCapabilities                  string `json:"rm-capabilities"`
	Dot11KRmBeaconMeasReqParameters struct {
		Period              int       `json:"period"`
		RepeatNum           int       `json:"repeat-num"`
		OperatingClass      int       `json:"operating-class"`
		ChannelNum          int       `json:"channel-num"`
		MeasMode            string    `json:"meas-mode"`
		CurrentBssid        bool      `json:"current-bssid"`
		Bssid               string    `json:"bssid"`
		CurrentSsid         bool      `json:"current-ssid"`
		Ssid                string    `json:"ssid"`
		DefaultRandInterval bool      `json:"default-rand-interval"`
		RandInterval        int       `json:"rand-interval"`
		DefaultMeasDuration bool      `json:"default-meas-duration"`
		MeasDuration        int       `json:"meas-duration"`
		DialogToken         int       `json:"dialog-token"`
		LastReqTrigger      string    `json:"last-req-trigger"`
		LastReqTime         time.Time `json:"last-req-time"`
		NextReqTime         time.Time `json:"next-req-time"`
		LastReportTime      time.Time `json:"last-report-time"`
	} `json:"dot11k-rm-beacon-meas-req-parameters"`
	CellularInfo struct {
		Capable     bool   `json:"capable"`
		NetworkType string `json:"network-type"`
		SignalScale string `json:"signal-scale"`
		CellID      int    `json:"cell-id"`
	} `json:"cellular-info"`
	WifiDirectClientCapabilities struct {
		WifiDirectCapable bool `json:"wifi-direct-capable"`
	} `json:"wifi-direct-client-capabilities"`
	WtcSupport      bool `json:"wtc-support"`
	AbrSupport      bool `json:"abr-support"`
	WtcResp         bool `json:"wtc-resp"`
	WtcRespCode     int  `json:"wtc-resp-code"`
	Dot116GhzCap    bool `json:"dot11-6ghz-cap"`
	LinkLocalEnable bool `json:"link-local-enable"`
}

type MobilityOperData struct {
	MsMacAddr           string    `json:"ms-mac-addr"`
	MmClientRole        string    `json:"mm-client-role"`
	MmClientRoamType    string    `json:"mm-client-roam-type"`
	MmInstance          int       `json:"mm-instance"`
	MmCompleteTimestamp time.Time `json:"mm-complete-timestamp"`
	MmRemoteTunnelIP    string    `json:"mm-remote-tunnel-ip"`
	MmRemoteTunnelSecIP string    `json:"mm-remote-tunnel-sec-ip"`
	MmRemotePlatformID  int       `json:"mm-remote-platform-id"`
	MmRemoteTunnelID    int       `json:"mm-remote-tunnel-id"`
	MmAnchorIP          string    `json:"mm-anchor-ip"`
}

type MmIfClientStats struct {
	ClientMac  string `json:"client-mac"`
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

type MmIfClientHistory struct {
	ClientMac       string `json:"client-mac"`
	MobilityHistory struct {
		Entry []struct {
			InstanceID    int       `json:"instance-id"`
			MsApSlotID    int       `json:"ms-ap-slot-id"`
			MsAssocTime   time.Time `json:"ms-assoc-time"`
			Role          string    `json:"role"`
			Bssid         string    `json:"bssid"`
			ApName        string    `json:"ap-name"`
			RunLatency    int       `json:"run-latency"`
			Dot11RoamType string    `json:"dot11-roam-type"`
		} `json:"entry"`
	} `json:"mobility-history"`
}

type TrafficStats struct {
	MsMacAddress             string    `json:"ms-mac-address"`
	BytesRx                  string    `json:"bytes-rx"`
	BytesTx                  string    `json:"bytes-tx"`
	PolicyErrs               string    `json:"policy-errs"`
	PktsRx                   string    `json:"pkts-rx"`
	PktsTx                   string    `json:"pkts-tx"`
	DataRetries              string    `json:"data-retries"`
	RtsRetries               string    `json:"rts-retries"`
	DuplicateRcv             string    `json:"duplicate-rcv"`
	DecryptFailed            string    `json:"decrypt-failed"`
	MicMismatch              string    `json:"mic-mismatch"`
	MicMissing               string    `json:"mic-missing"`
	MostRecentRssi           int       `json:"most-recent-rssi"`
	MostRecentSnr            int       `json:"most-recent-snr"`
	TxExcessiveRetries       string    `json:"tx-excessive-retries"`
	TxRetries                string    `json:"tx-retries"`
	PowerSaveState           int       `json:"power-save-state"`
	CurrentRate              string    `json:"current-rate"`
	Speed                    int       `json:"speed"`
	SpatialStream            int       `json:"spatial-stream"`
	ClientActive             bool      `json:"client-active"`
	GlanStatsUpdateTimestamp time.Time `json:"glan-stats-update-timestamp"`
	GlanIdleUpdateTimestamp  time.Time `json:"glan-idle-update-timestamp"`
	RxGroupCounter           string    `json:"rx-group-counter"`
	TxTotalDrops             string    `json:"tx-total-drops"`
}

type PolicyData struct {
	Mac         string `json:"mac"`
	ResVlanID   int    `json:"res-vlan-id"`
	ResVlanName string `json:"res-vlan-name"`
}

type SisfDbMac struct {
	MacAddr     string `json:"mac-addr"`
	Ipv4Binding struct {
		IPKey struct {
			ZoneID int    `json:"zone-id"`
			IPAddr string `json:"ip-addr"`
		} `json:"ip-key"`
	} `json:"ipv4-binding"`
	Ipv6Binding []struct {
		Ipv6BindingIPKey struct {
			ZoneID int64  `json:"zone-id"`
			IPAddr string `json:"ip-addr"`
		} `json:"ip-key"`
	} `json:"ipv6-binding,omitempty"`
}

type DcInfo struct {
	ClientMac        string    `json:"client-mac"`
	DeviceType       string    `json:"device-type"`
	ProtocolMap      string    `json:"protocol-map"`
	ConfidenceLevel  int       `json:"confidence-level"`
	ClassifiedTime   time.Time `json:"classified-time"`
	DayZeroDc        string    `json:"day-zero-dc"`
	SwVersionSrc     string    `json:"sw-version-src"`
	DeviceOs         string    `json:"device-os,omitempty"`
	DeviceSubVersion string    `json:"device-sub-version,omitempty"`
	DeviceOsSrc      string    `json:"device-os-src"`
	DeviceName       string    `json:"device-name"`
	DeviceVendorSrc  string    `json:"device-vendor-src"`
	SalesCodeSrc     string    `json:"sales-code-src"`
	DeviceSrc        string    `json:"device-src"`
	CountryNameSrc   string    `json:"country-name-src"`
	ModelNameSrc     string    `json:"model-name-src"`
	PowerTypeSrc     string    `json:"power-type-src"`
	HwModelSrc       string    `json:"hw-model-src"`
	DeviceVendor     string    `json:"device-vendor,omitempty"`
}

// GetClientOper retrieves complete client operational data.
func GetClientOper(client *wnc.Client, ctx context.Context) (*ClientOperResponse, error) {
	var data ClientOperResponse
	if err := client.SendAPIRequest(ctx, ClientOperEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetClientOperCommonOperData retrieves common operational data for wireless clients.
func GetClientOperCommonOperData(client *wnc.Client, ctx context.Context) (*ClientOperCommonOperDataResponse, error) {
	var data ClientOperCommonOperDataResponse
	if err := client.SendAPIRequest(ctx, CommonOperDataEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetClientOperDot11OperData retrieves 802.11 operational data for wireless clients.
func GetClientOperDot11OperData(client *wnc.Client, ctx context.Context) (*ClientOperDot11OperDataResponse, error) {
	var data ClientOperDot11OperDataResponse
	if err := client.SendAPIRequest(ctx, Dot11OperDataEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetClientOperMobilityOperData retrieves mobility operational data for wireless clients.
func GetClientOperMobilityOperData(client *wnc.Client, ctx context.Context) (*ClientOperMobilityOperDataResponse, error) {
	var data ClientOperMobilityOperDataResponse
	if err := client.SendAPIRequest(ctx, MobilityOperDataEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetClientOperMmIfClientStats retrieves mobility manager interface client statistics.
func GetClientOperMmIfClientStats(client *wnc.Client, ctx context.Context) (*ClientOperMmIfClientStatsResponse, error) {
	var data ClientOperMmIfClientStatsResponse
	if err := client.SendAPIRequest(ctx, MmIfClientStatsEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetClientOperMmIfClientHistory retrieves mobility manager interface client history.
func GetClientOperMmIfClientHistory(client *wnc.Client, ctx context.Context) (*ClientOperMmIfClientHistoryResponse, error) {
	var data ClientOperMmIfClientHistoryResponse
	if err := client.SendAPIRequest(ctx, MmIfClientHistoryEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetClientOperTrafficStats retrieves traffic statistics for wireless clients.
func GetClientOperTrafficStats(client *wnc.Client, ctx context.Context) (*ClientOperTrafficStatsResponse, error) {
	var data ClientOperTrafficStatsResponse
	if err := client.SendAPIRequest(ctx, TrafficStatsEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetClientOperPolicyData retrieves policy data for wireless clients.
func GetClientOperPolicyData(client *wnc.Client, ctx context.Context) (*ClientOperPolicyDataResponse, error) {
	var data ClientOperPolicyDataResponse
	if err := client.SendAPIRequest(ctx, PolicyDataEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetClientOperSisfDbMac retrieves SISF database MAC information.
func GetClientOperSisfDbMac(client *wnc.Client, ctx context.Context) (*ClientOperSisfDbMacResponse, error) {
	var data ClientOperSisfDbMacResponse
	if err := client.SendAPIRequest(ctx, SisfDbMacEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetClientOperDcInfo retrieves discovery client information.
func GetClientOperDcInfo(client *wnc.Client, ctx context.Context) (*ClientOperDcInfoResponse, error) {
	var data ClientOperDcInfoResponse
	if err := client.SendAPIRequest(ctx, DcInfoEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}
