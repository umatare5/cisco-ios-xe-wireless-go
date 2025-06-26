// Package client provides client global operational data functionality for the Cisco Wireless Network Controller API.
package client

import (
	"context"

	wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
)

const (
	// ClientGlobalOperBasePath defines the base path for client global operational data endpoints.
	ClientGlobalOperBasePath = "/restconf/data/Cisco-IOS-XE-wireless-client-global-oper:client-global-oper-data"
	// ClientGlobalOperEndpoint defines the endpoint for client global operational data.
	ClientGlobalOperEndpoint = ClientGlobalOperBasePath
	// ClientLiveStatsEndpoint defines the endpoint for client live statistics.
	ClientLiveStatsEndpoint = ClientGlobalOperBasePath + "/client-live-stats"
	// ClientGlobalStatsDataEndpoint defines the endpoint for client global statistics data.
	ClientGlobalStatsDataEndpoint = ClientGlobalOperBasePath + "/client-global-stats-data"
	// ClientStatsEndpoint defines the endpoint for client statistics.
	ClientStatsEndpoint = ClientGlobalOperBasePath + "/client-stats"
	// ClientDot11StatsEndpoint defines the endpoint for client 802.11 statistics.
	ClientDot11StatsEndpoint = ClientGlobalOperBasePath + "/client-dot11-stats"
	// ClientLatencyStatsEndpoint defines the endpoint for client latency statistics.
	ClientLatencyStatsEndpoint = ClientGlobalOperBasePath + "/client-latency-stats"
	// SmWebauthStatsEndpoint defines the endpoint for session manager web authentication statistics.
	SmWebauthStatsEndpoint = ClientGlobalOperBasePath + "/sm-webauth-stats"
	// Dot1XGlobalStatsEndpoint defines the endpoint for 802.1X global statistics.
	Dot1XGlobalStatsEndpoint = ClientGlobalOperBasePath + "/dot1x-global-stats"
	// ClientExclusionStatsEndpoint defines the endpoint for client exclusion statistics.
	ClientExclusionStatsEndpoint = ClientGlobalOperBasePath + "/client-exclusion-stats"
	// SmDeviceCountEndpoint defines the endpoint for session manager device count data.
	SmDeviceCountEndpoint = ClientGlobalOperBasePath + "/sm-device-count"
	// TofStatsEndpoint defines the endpoint for time-of-flight statistics.
	TofStatsEndpoint = ClientGlobalOperBasePath + "/tof-stats"
)

// ClientGlobalOperResponse represents the response structure for client global operational data.
type ClientGlobalOperResponse struct {
	CiscoIOSXEWirelessClientGlobalOperClientGlobalOperData struct {
		ClientLiveStats       ClientLiveStats       `json:"client-live-stats"`
		ClientGlobalStatsData ClientGlobalStatsData `json:"client-global-stats-data"`
		ClientStats           ClientStats           `json:"client-stats"`
		ClientDot11Stats      ClientDot11Stats      `json:"client-dot11-stats"`
		ClientLatencyStats    ClientLatencyStats    `json:"client-latency-stats"`
		SmWebauthStats        SmWebauthStats        `json:"sm-webauth-stats"`
		Dot1XGlobalStats      Dot1XGlobalStats      `json:"dot1x-global-stats"`
		ClientExclusionStats  ClientExclusionStats  `json:"client-exclusion-stats"`
		SmDeviceCount         SmDeviceCount         `json:"sm-device-count"`
		TofStats              TofStats              `json:"tof-stats"`
	} `json:"Cisco-IOS-XE-wireless-client-global-oper:client-global-oper-data"`
}

// ClientLiveStatsResponse represents the response structure for client live statistics.
type ClientLiveStatsResponse struct {
	ClientLiveStats ClientLiveStats `json:"Cisco-IOS-XE-wireless-client-global-oper:client-live-stats"`
}

// ClientGlobalStatsDataResponse represents the response structure for client global statistics data.
type ClientGlobalStatsDataResponse struct {
	ClientGlobalStatsData ClientGlobalStatsData `json:"Cisco-IOS-XE-wireless-client-global-oper:client-global-stats-data"`
}

// ClientStatsResponse represents the response structure for client statistics.
type ClientStatsResponse struct {
	ClientStats ClientStats `json:"Cisco-IOS-XE-wireless-client-global-oper:client-stats"`
}

// ClientDot11StatsResponse represents the response structure for client 802.11 statistics.
type ClientDot11StatsResponse struct {
	ClientDot11Stats ClientDot11Stats `json:"Cisco-IOS-XE-wireless-client-global-oper:client-dot11-stats"`
}

// ClientLatencyStatsResponse represents the response structure for client latency statistics.
type ClientLatencyStatsResponse struct {
	ClientLatencyStats ClientLatencyStats `json:"Cisco-IOS-XE-wireless-client-global-oper:client-latency-stats"`
}

// SmWebauthStatsResponse represents the response structure for session manager web authentication statistics.
type SmWebauthStatsResponse struct {
	SmWebauthStats SmWebauthStats `json:"Cisco-IOS-XE-wireless-client-global-oper:sm-webauth-stats"`
}

// Dot1XGlobalStatsResponse represents the response structure for 802.1X global statistics.
type Dot1XGlobalStatsResponse struct {
	Dot1XGlobalStats Dot1XGlobalStats `json:"Cisco-IOS-XE-wireless-client-global-oper:dot1x-global-stats"`
}

// ClientExclusionStatsResponse represents the response structure for client exclusion statistics.
type ClientExclusionStatsResponse struct {
	ClientExclusionStats ClientExclusionStats `json:"Cisco-IOS-XE-wireless-client-global-oper:client-exclusion-stats"`
}

// SmDeviceCountResponse represents the response structure for session manager device count data.
type SmDeviceCountResponse struct {
	SmDeviceCount SmDeviceCount `json:"Cisco-IOS-XE-wireless-client-global-oper:sm-device-count"`
}

// TofStatsResponse represents the response structure for time-of-flight statistics.
type TofStatsResponse struct {
	TofStats TofStats `json:"Cisco-IOS-XE-wireless-client-global-oper:tof-stats"`
}

// ClientLiveStats represents live client statistics including counts of clients in various states.
type ClientLiveStats struct {
	AuthStateClients     int `json:"auth-state-clients"`
	MobilityStateClients int `json:"mobility-state-clients"`
	IplearnStateClients  int `json:"iplearn-state-clients"`
	WebauthStateClients  int `json:"webauth-state-clients"`
	RunStateClients      int `json:"run-state-clients"`
	DeleteStateClients   int `json:"delete-state-clients"`
	RandomMacClients     int `json:"random-mac-clients"`
}

// ClientGlobalStatsData represents global client statistics data including session statistics.
type ClientGlobalStatsData struct {
	SessionStats struct {
		MostMsEntries       int `json:"most-ms-entries"`
		CurMsEntries        int `json:"cur-ms-entries"`
		TotalMsEntries      int `json:"total-ms-entries"`
		TotalIpv4MsEntries  int `json:"total-ipv4-ms-entries"`
		CurWgbEntries       int `json:"cur-wgb-entries"`
		CurForeignEntries   int `json:"cur-foreign-entries"`
		CurAnchorEntries    int `json:"cur-anchor-entries"`
		CurLocalEntries     int `json:"cur-local-entries"`
		CurIdleStateEntries int `json:"cur-idle-state-entries"`
	} `json:"session-stats"`
}

// ClientStats represents client statistics including deletion reasons and various counters.
type ClientStats struct {
	CoClientDelReason struct {
		InvalidOperation        int `json:"invalid-operation"`
		InternalGenericErr      int `json:"internal-generic-err"`
		DeauthOrDisassocReq     int `json:"deauth-or-disassoc-req"`
		AuthFail                int `json:"auth-fail"`
		WebauthFail             int `json:"webauth-fail"`
		ApDelete                int `json:"ap-delete"`
		BssidDown               int `json:"bssid-down"`
		CapwapDown              int `json:"capwap-down"`
		ConnectTimeout          int `json:"connect-timeout"`
		MabFail                 int `json:"mab-fail"`
		DatapathFail            int `json:"datapath-fail"`
		WlanChange              int `json:"wlan-change"`
		VlanChange              int `json:"vlan-change"`
		AdminReset              int `json:"admin-reset"`
		QosFail                 int `json:"qos-fail"`
		KeyExchangeTimeout      int `json:"key-exchange-timeout"`
		GroupKeyUpdate          int `json:"group-key-update"`
		MaxSaQueries            int `json:"max-sa-queries"`
		RestartPurge            int `json:"restart-purge"`
		ClientBlockList         int `json:"client-block-list"`
		InterInstanceRoamSucc   int `json:"inter-instance-roam-succ"`
		IntraInstanceRoamFail   int `json:"intra-instance-roam-fail"`
		MobilityRoamSucc        int `json:"mobility-roam-succ"`
		MobilityRoamFail        int `json:"mobility-roam-fail"`
		ClientSessionTimeout    int `json:"client-session-timeout"`
		ClientIdleTimeout       int `json:"client-idle-timeout"`
		UserReq                 int `json:"user-req"`
		NasErr                  int `json:"nas-err"`
		AaaServiceUnavailable   int `json:"aaa-service-unavailable"`
		MobilityBssidDown       int `json:"mobility-bssid-down"`
		MobilityTunnelDown      int `json:"mobility-tunnel-down"`
		Dot11VSmartRoamFail     int `json:"dot11v-smart-roam-fail"`
		Dot11VTimerTimeout      int `json:"dot11v-timer-timeout"`
		Dot11VAssocFail         int `json:"dot11v-assoc-fail"`
		FtAuthResponseFail      int `json:"ft-auth-response-fail"`
		SaeAuthFail             int `json:"sae-auth-fail"`
		Dot11UnspecifiedFail    int `json:"dot11-unspecified-fail"`
		Dot11FailIgnoreReq      int `json:"dot11-fail-ignore-req"`
		Dot11CcxQosAddTsNoBw    int `json:"dot11-ccx-qos-add-ts-no-bw"`
		Dot11CapsUnsupported    int `json:"dot11-caps-unsupported"`
		Dot11AssocDenyUnspec    int `json:"dot11-assoc-deny-unspec"`
		Dot11MaxClient          int `json:"dot11-max-client"`
		Dot11DeniedRates        int `json:"dot11-denied-rates"`
		Dot11PoorChannel        int `json:"dot11-poor-channel"`
		Dot11InvalidQosParam    int `json:"dot11-invalid-qos-param"`
		Dot11InvalidIe          int `json:"dot11-invalid-ie"`
		Dot11GroupCipherInvalid int `json:"dot11-group-cipher-invalid"`
		Dot11UcastCipherInvalid int `json:"dot11-ucast-cipher-invalid"`
		Dot11AkmpInvalid        int `json:"dot11-akmp-invalid"`
		Dot11RsnVerNoSupported  int `json:"dot11-rsn-ver-no-supported"`
		Dot11InvalidRsnIe       int `json:"dot11-invalid-rsn-ie"`
		Dot11CipherSuiteReject  int `json:"dot11-cipher-suite-reject"`
		Dot11DenyHtCapab        int `json:"dot11-deny-ht-capab"`
		Dot11InvalidFtFrame     int `json:"dot11-invalid-ft-frame"`
		Dot11InvalidPmkid       int `json:"dot11-invalid-pmkid"`
		Dot11InvalidMdie        int `json:"dot11-invalid-mdie"`
		Dot11InvalidFtie        int `json:"dot11-invalid-ftie"`
		Dot11CcxQosPolicy       int `json:"dot11-ccx-qos-policy"`
		Dot11BandwidthFail      int `json:"dot11-bandwidth-fail"`
		Dot11CcxInvalidqosParam int `json:"dot11-ccx-invalidqos-param"`
		Dot11CcxNonOptChoice    int `json:"dot11-ccx-non-opt-choice"`
		WiredIappDisassoc       int `json:"wired-iapp-disassoc"`
		WiredWgbChange          int `json:"wired-wgb-change"`
		WiredVlanChange         int `json:"wired-vlan-change"`
		WiredWgbDelete          int `json:"wired-wgb-delete"`
		AvcReanchoring          int `json:"avc-reanchoring"`
		WgbClientDirectAssoc    int `json:"wgb-client-direct-assoc"`
		ApUpgrade               int `json:"ap-upgrade"`
		ClientDhcpFail          int `json:"client-dhcp-fail"`
		EapTimeoutFail          int `json:"eap-timeout-fail"`
		Client8021XFail         int `json:"client-8021x-fail"`
		ClientDeviceIdleTimeout int `json:"client-device-idle-timeout"`
		CaptivePortalSecFail    int `json:"captive-portal-sec-fail"`
		ClientDecryptFail       int `json:"client-decrypt-fail"`
		ClientIntfDisable       int `json:"client-intf-disable"`
		ClientDisassociated     int `json:"client-disassociated"`
		ClientMiscReason        int `json:"client-misc-reason"`
		ClientUnknownReason     int `json:"client-unknown-reason"`
		ClientPeerTriggered     int `json:"client-peer-triggered"`
		ClientBeaconLoss        int `json:"client-beacon-loss"`
		ClientEapIDTimeout      int `json:"client-eap-id-timeout"`
		ClientDot1XTimeout      int `json:"client-dot1x-timeout"`
		RecvEapKeyInvalidFrame  int `json:"recv-eap-key-invalid-frame"`
		RecvEapKeyInstallBit    int `json:"recv-eap-key-install-bit"`
		RecvEapKeyErrBit        int `json:"recv-eap-key-err-bit"`
		RecvEapKeyAckBit        int `json:"recv-eap-key-ack-bit"`
		RecvEapKeyInvalidKey    int `json:"recv-eap-key-invalid-key"`
		RecvEapKeySecureBit     int `json:"recv-eap-key-secure-bit"`
		RecvEapKeyDescVer       int `json:"recv-eap-key-desc-ver"`
		RecvEapkeyWrReplayCnt   int `json:"recv-eapkey-wr-replay-cnt"`
		RecvEapKeyNoMicBit      int `json:"recv-eap-key-no-mic-bit"`
		RecvEapKeyMicValidate   int `json:"recv-eap-key-mic-validate"`
		RecvEapKeyPtkCompute    int `json:"recv-eap-key-ptk-compute"`
		ClientCredentialFail    int `json:"client-credential-fail"`
		LostCarrier             int `json:"lost-carrier"`
		ReauthFail              int `json:"reauth-fail"`
		PortAdminDisable        int `json:"port-admin-disable"`
		SupplicantRestart       int `json:"supplicant-restart"`
		IPDownNoIP              int `json:"ip-down-no-ip"`
		AnchorThrottled         int `json:"anchor-throttled"`
		AnchorNoMemory          int `json:"anchor-no-memory"`
		AnchorInvalidMbssid     int `json:"anchor-invalid-mbssid"`
		AnchorCreateReqFail     int `json:"anchor-create-req-fail"`
		DbPopulateFail          int `json:"db-populate-fail"`
		DanglingCleanupTimer    int `json:"dangling-cleanup-timer"`
		ExcludeStaticConfig     int `json:"exclude-static-config"`
		ExcludeAssocFail        int `json:"exclude-assoc-fail"`
		ExcludeDot11AuthFail    int `json:"exclude-dot11-auth-fail"`
		ExcludeDot1XTimeout     int `json:"exclude-dot1x-timeout"`
		ExcludeDot1XAuthFail    int `json:"exclude-dot1x-auth-fail"`
		ExcludeWebAuthFail      int `json:"exclude-web-auth-fail"`
		ExcludePolicyBindFail   int `json:"exclude-policy-bind-fail"`
		ExcludeIPTheft          int `json:"exclude-ip-theft"`
		ExcludeMacTheft         int `json:"exclude-mac-theft"`
		ExcludeMacAndIPTheft    int `json:"exclude-mac-and-ip-theft"`
		ExcludeQosPolicyFail    int `json:"exclude-qos-policy-fail"`
		ExcludeQospolApsendFail int `json:"exclude-qospol-apsend-fail"`
		ExcludeQospolBindFail   int `json:"exclude-qospol-bind-fail"`
		ExcludeQospolUnbindFail int `json:"exclude-qospol-unbind-fail"`
		ExcludeStaIPAnchorFail  int `json:"exclude-sta-ip-anchor-fail"`
		ExcludeVlanFail         int `json:"exclude-vlan-fail"`
		ExcludeACLFail          int `json:"exclude-acl-fail"`
		ExcludePuntACLFail      int `json:"exclude-punt-acl-fail"`
		ExcludeAccountingFail   int `json:"exclude-accounting-fail"`
		ExcludeCtsFail          int `json:"exclude-cts-fail"`
		ExcludeFqdnNoDefFail    int `json:"exclude-fqdn-no-def-fail"`
		ExcludeFqdnPoauthFail   int `json:"exclude-fqdn-poauth-fail"`
		ExcludeFqdnPreauthFail  int `json:"exclude-fqdn-preauth-fail"`
		ExcludeFqdnzeroGidFail  int `json:"exclude-fqdnzero-gid-fail"`
		ExcludeMiscFail         int `json:"exclude-misc-fail"`
		ExcludeReauthFail       int `json:"exclude-reauth-fail"`
		ExcludeWrongPsk         int `json:"exclude-wrong-psk"`
		ExcludePolicyFail       int `json:"exclude-policy-fail"`
		ApIdleTimeout           int `json:"ap-idle-timeout"`
		ApClientACLMismatch     int `json:"ap-client-acl-mismatch"`
		ApAuthStop              int `json:"ap-auth-stop"`
		ApAssocExpiredAtAp      int `json:"ap-assoc-expired-at-ap"`
		FourwayHandshakeFail    int `json:"fourway-handshake-fail"`
		ApDhcpTimeout           int `json:"ap-dhcp-timeout"`
		ApReassocTimeout        int `json:"ap-reassoc-timeout"`
		ApSaQueryTimeout        int `json:"ap-sa-query-timeout"`
		ApIntraApRoam           int `json:"ap-intra-ap-roam"`
		ApChannelSwitchAtAp     int `json:"ap-channel-switch-at-ap"`
		ApDelBadAid             int `json:"ap-del-bad-aid"`
		ApDelReq                int `json:"ap-del-req"`
		ApDelIntfReset          int `json:"ap-del-intf-reset"`
		ApDelAllOnSlot          int `json:"ap-del-all-on-slot"`
		ApDelLinkChangeReaper   int `json:"ap-del-link-change-reaper"`
		ApDelSlotDisable        int `json:"ap-del-slot-disable"`
		ApDelMicFail            int `json:"ap-del-mic-fail"`
		ApDelVlanDel            int `json:"ap-del-vlan-del"`
		ApDelChannelChange      int `json:"ap-del-channel-change"`
		ApDelStopReassoc        int `json:"ap-del-stop-reassoc"`
		ApDelPakMaxRetry        int `json:"ap-del-pak-max-retry"`
		ApDelTxDeauth           int `json:"ap-del-tx-deauth"`
		ApSensorStaTimeout      int `json:"ap-sensor-sta-timeout"`
		ApDelAgeTimeout         int `json:"ap-del-age-timeout"`
		ApDelTxFailThold        int `json:"ap-del-tx-fail-thold"`
		ApUplinkRecvTimeout     int `json:"ap-uplink-recv-timeout"`
		ApSnsrscanNxtradio      int `json:"ap-snsrscan-nxtradio"`
		ApSnsrscanOtherbssid    int `json:"ap-snsrscan-otherbssid"`
		AaaServerUnavailable    int `json:"aaa-server-unavailable"`
		AaaNotReady             int `json:"aaa-not-ready"`
		NoDot1XAuthConfig       int `json:"no-dot1x-auth-config"`
		AbortRecv               int `json:"abort-recv"`
		AssocConnectTimeout     int `json:"assoc-connect-timeout"`
		MacauthConnectTimeout   int `json:"macauth-connect-timeout"`
		L2AuthConnectTimeout    int `json:"l2auth-connect-timeout"`
		WebauthConnectTimeout   int `json:"webauth-connect-timeout"`
		MobilityConnectTimeout  int `json:"mobility-connect-timeout"`
		StaticAnchorTimeout     int `json:"static-anchor-timeout"`
		SmSessionConnectTimeout int `json:"sm-session-connect-timeout"`
		IplearnConnectTimeout   int `json:"iplearn-connect-timeout"`
		DatapathIfidExists      int `json:"datapath-ifid-exists"`
		RadioDown               int `json:"radio-down"`
		FabricReject            int `json:"fabric-reject"`
		GuestLanInvalidMbssid   int `json:"guest-lan-invalid-mbssid"`
		GuestLanNoMemory        int `json:"guest-lan-no-memory"`
		GuestLanCreateReqFail   int `json:"guest-lan-create-req-fail"`
		TunnelEogreReset        int `json:"tunnel-eogre-reset"`
		TunnelEogreJoinFail     int `json:"tunnel-eogre-join-fail"`
		TunEogreReconcile       int `json:"tun-eogre-reconcile"`
		WiredIdleTimeout        int `json:"wired-idle-timeout"`
		IPUpdateTimeout         int `json:"ip-update-timeout"`
		RemoteMobilityDelete    int `json:"remote-mobility-delete"`
		SaeAuthInAssocedSt      int `json:"sae-auth-in-assoced-st"`
		NackIfidMismatch        int `json:"nack-ifid-mismatch"`
		TunnnelEogreInvalidVlan int `json:"tunnnel-eogre-invalid-vlan"`
		TunnnelEogreEmptyDomain int `json:"tunnnel-eogre-empty-domain"`
		TunnelEogreInvDomain    int `json:"tunnel-eogre-inv-domain"`
		TunnelEogreDomainShut   int `json:"tunnel-eogre-domain-shut"`
		TunnelEogreInvalidGway  int `json:"tunnel-eogre-invalid-gway"`
		TunnelEogreGwayDown     int `json:"tunnel-eogre-gway-down"`
		TunnelEogreflexNoActgw  int `json:"tunnel-eogreflex-no-actgw"`
		TunnelEogreRuleMatch    int `json:"tunnel-eogre-rule-match"`
		TunnelEogreAaaOverride  int `json:"tunnel-eogre-aaa-override"`
		TunnelEogreMspayload    int `json:"tunnel-eogre-mspayload"`
		TunnelEogreHandoffErr   int `json:"tunnel-eogre-handoff-err"`
		InvalidPmkLen           int `json:"invalid-pmk-len"`
		L3VlanOrideConnTimeout  int `json:"l3-vlan-oride-conn-timeout"`
		UserTriggerUnspecified  int `json:"user-trigger-unspecified"`
		UsrTriggerPwrofWifiof   int `json:"usr-trigger-pwrof-wifiof"`
		ConnectToOtherSsid      int `json:"connect-to-other-ssid"`
		UsrTriggerRemoveSsid    int `json:"usr-trigger-remove-ssid"`
		UsrTriggerAirplaneMode  int `json:"usr-trigger-airplane-mode"`
		L2ConnectionUnspecify   int `json:"l2connection-unspecify"`
		L2ConnectionAssoc       int `json:"l2connection-assoc"`
		FourwayHandshakeUnspec  int `json:"fourway-handshake-unspec"`
		EapKeyM1Fail            int `json:"eap-key-m1-fail"`
		EapKeyM3Fail            int `json:"eap-key-m3-fail"`
		EapkeyM3M4XchngTimeout  int `json:"eapkey-m3-m4-xchng-timeout"`
		DhcpFailUnspecified     int `json:"dhcp-fail-unspecified"`
		DhcpFailTimeout         int `json:"dhcp-fail-timeout"`
		DhcpFailTimeoutRoam     int `json:"dhcp-fail-timeout-roam"`
		DhcpFailLeaseExpired    int `json:"dhcp-fail-lease-expired"`
		DhcpFailNakInRenew      int `json:"dhcp-fail-nak-in-renew"`
		DhcpfailRenLeaseIP      int `json:"dhcpfail-ren-lease-ip"`
		DhcpFailIntErr          int `json:"dhcp-fail-int-err"`
		EapFailUnspecified      int `json:"eap-fail-unspecified"`
		EapFailCodeFail         int `json:"eap-fail-code-fail"`
		EapFailInfonotEnter     int `json:"eap-fail-infonot-enter"`
		EapFailMschapErr        int `json:"eap-fail-mschap-err"`
		EapFailSimAuthFail      int `json:"eap-fail-sim-auth-fail"`
		EapFailTimeout          int `json:"eap-fail-timeout"`
		EapFailTLSCertErr       int `json:"eap-fail-tls-cert-err"`
		NoInternetUnspecified   int `json:"no-internet-unspecified"`
		NoInternetDNS           int `json:"no-internet-dns"`
		NoInternetArp           int `json:"no-internet-arp"`
		ExcludePolTemplateFail  int `json:"exclude-pol-template-fail"`
		ApDelAuthExpired        int `json:"ap-del-auth-expired"`
		ApDelDisassoc           int `json:"ap-del-disassoc"`
		ApDelTxDisassoc         int `json:"ap-del-tx-disassoc"`
		FastroamMobilityFail    int `json:"fastroam-mobility-fail"`
		PolicyProfileDeny       int `json:"policy-profile-deny"`
		Dot11AidAllocReq        int `json:"dot11-aid-alloc-req"`
		ZoneChange              int `json:"zone-change"`
		WlanIDAttrMismatch      int `json:"wlan-id-attr-mismatch"`
		EpskAaaUnknownErr       int `json:"epsk-aaa-unknown-err"`
		EpskUnspecErr           int `json:"epsk-unspec-err"`
		EpskPskMismatch         int `json:"epsk-psk-mismatch"`
		EasyPskRadiusBusy       int `json:"easy-psk-radius-busy"`
		EasyPskLimitReached     int `json:"easy-psk-limit-reached"`
		EasyPskBad8021XFrame    int `json:"easy-psk-bad-8021x-frame"`
		EpskMissingParam        int `json:"epsk-missing-param"`
		ExcludeSuppNameFail     int `json:"exclude-supp-name-fail"`
		ExcludeUserNameFail     int `json:"exclude-user-name-fail"`
		ExcludeServiceSetidFail int `json:"exclude-service-setid-fail"`
		ExcludeAnchVlanidFail   int `json:"exclude-anch-vlanid-fail"`
		ExcludePskFail          int `json:"exclude-psk-fail"`
		ExcludePskModeFail      int `json:"exclude-psk-mode-fail"`
		ExcludeIntIntervalFail  int `json:"exclude-int-interval-fail"`
		RandomMac               int `json:"random-mac"`
		ApIplearnTimeout        int `json:"ap-iplearn-timeout"`
		ApFlexgroupChange       int `json:"ap-flexgroup-change"`
		ApEapolLogoff           int `json:"ap-eapol-logoff"`
		ApEapReqTimeout         int `json:"ap-eap-req-timeout"`
		ApFourwayHandshakeFail  int `json:"ap-fourway-handshake-fail"`
		ApMicValidation         int `json:"ap-mic-validation"`
		ApWrongReplayCounter    int `json:"ap-wrong-replay-counter"`
		ApTunnelDown            int `json:"ap-tunnel-down"`
		InterApRoam             int `json:"inter-ap-roam"`
		ApUnknownClient         int `json:"ap-unknown-client"`
		ApReauthTimeout         int `json:"ap-reauth-timeout"`
		ApContIdleTimeout       int `json:"ap-cont-idle-timeout"`
		ApRldpCleanup           int `json:"ap-rldp-cleanup"`
		ApIntraSwitchRoam       int `json:"ap-intra-switch-roam"`
		ApPemCleanup            int `json:"ap-pem-cleanup"`
		ApRlanCentralSwitch     int `json:"ap-rlan-central-switch"`
		ApRlanDpAddFail         int `json:"ap-rlan-dp-add-fail"`
		ApRlanDelete            int `json:"ap-rlan-delete"`
		ApRlanInactiveTimeout   int `json:"ap-rlan-inactive-timeout"`
		ApRlanMabFail           int `json:"ap-rlan-mab-fail"`
		ApNoMemory              int `json:"ap-no-memory"`
		ApBssidMismatch         int `json:"ap-bssid-mismatch"`
		ApDeleteNoACL           int `json:"ap-delete-no-acl"`
		ApDelNoParentWgb        int `json:"ap-del-no-parent-wgb"`
		ApKeyPlumbFail          int `json:"ap-key-plumb-fail"`
		ApMeshKeyplumbFail      int `json:"ap-mesh-keyplumb-fail"`
		ApDatapathAddFail       int `json:"ap-datapath-add-fail"`
		ApAuthRespReject        int `json:"ap-auth-resp-reject"`
		ApAuthRespSendFail      int `json:"ap-auth-resp-send-fail"`
		ApAssocRespSendFail     int `json:"ap-assoc-resp-send-fail"`
		ApAssocRespFailStat     int `json:"ap-assoc-resp-fail-stat"`
		ApWebauthTimerExp       int `json:"ap-webauth-timer-exp"`
		ApDot1XTimerExp         int `json:"ap-dot1x-timer-exp"`
		ApDeauthDisassocFail    int `json:"ap-deauth-disassoc-fail"`
		ApDvrEvtClass3Recv      int `json:"ap-dvr-evt-class3-recv"`
		ApDvrEvtPspUnauth       int `json:"ap-dvr-evt-psp-unauth"`
		ApDvrEvtIoctlErr        int `json:"ap-dvr-evt-ioctl-err"`
		ApFlexFtFail            int `json:"ap-flex-ft-fail"`
		ApDvrAddFail            int `json:"ap-dvr-add-fail"`
		ApDvrClientNotFound     int `json:"ap-dvr-client-not-found"`
		ApDvrMgmtpktAllocFail   int `json:"ap-dvr-mgmtpkt-alloc-fail"`
		ApDvrInvalidCipher      int `json:"ap-dvr-invalid-cipher"`
		ApDvrInvalidAid         int `json:"ap-dvr-invalid-aid"`
		ApDvrInvalidKey         int `json:"ap-dvr-invalid-key"`
		ApDvrFwKeySetFail       int `json:"ap-dvr-fw-key-set-fail"`
		ApDvrInvalidHtvhtRates  int `json:"ap-dvr-invalid-htvht-rates"`
		ApDvrInvalidLegRates    int `json:"ap-dvr-invalid-leg-rates"`
		ApDvrNoOlapLegRates     int `json:"ap-dvr-no-olap-leg-rates"`
		ApDvrMaxVhtStreams      int `json:"ap-dvr-max-vht-streams"`
		ApDriverAidInuse        int `json:"ap-driver-aid-inuse"`
		ApDvrAssocTooMany       int `json:"ap-dvr-assoc-too-many"`
		ApDvrCipherAttachFail   int `json:"ap-dvr-cipher-attach-fail"`
		ApDvrAlgoMismatch       int `json:"ap-dvr-algo-mismatch"`
		ApDvrInvalidKeylen      int `json:"ap-dvr-invalid-keylen"`
		ApDvrInvalidKeyIndex    int `json:"ap-dvr-invalid-key-index"`
		ApDvrAssocScbNoauth     int `json:"ap-dvr-assoc-scb-noauth"`
		ApDvrStadbInitFail      int `json:"ap-dvr-stadb-init-fail"`
		ApDvrAddEntryFail       int `json:"ap-dvr-add-entry-fail"`
		ApDvrAddEntFwFail       int `json:"ap-dvr-add-ent-fw-fail"`
		ApReasonInternalFail    int `json:"ap-reason-internal-fail"`
		LinkLocalbridgeVlanFail int `json:"link-localbridge-vlan-fail"`
		MaxStaOnAp              int `json:"max-sta-on-ap"`
		MaxStaOnVap             int `json:"max-sta-on-vap"`
		MaxStaOnBssid           int `json:"max-sta-on-bssid"`
		MaxStaOnRadio           int `json:"max-sta-on-radio"`
		ApMaxStaOnAp            int `json:"ap-max-sta-on-ap"`
		ApMaxStaOnBssid         int `json:"ap-max-sta-on-bssid"`
		ApMaxStaOnRadio         int `json:"ap-max-sta-on-radio"`
		StaTriggeredPmkTimeout  int `json:"sta-triggered-pmk-timeout"`
		ExcessArpRate           int `json:"excess-arp-rate"`
		Dot11UnspecQosReason    int `json:"dot11-unspec-qos-reason"`
		ExcessNdpRate           int `json:"excess-ndp-rate"`
		VrfVlanMismatch         int `json:"vrf-vlan-mismatch"`
		DpathEncodeFailed       int `json:"dpath-encode-failed"`
	} `json:"co-client-del-reason"`
}

// ClientDot11Stats represents 802.11 client statistics including association, DMS, and roaming statistics.
type ClientDot11Stats struct {
	ClientAssocStats struct {
		TotalAttempt    int `json:"total-attempt"`
		TotalFail       int `json:"total-fail"`
		TotalRespAccept int `json:"total-resp-accept"`
		TotalRespReject int `json:"total-resp-reject"`
		TotalRespError  int `json:"total-resp-error"`
	} `json:"client-assoc-stats"`
	ClientDmsStats struct {
		ActionFrameReq  int `json:"action-frame-req"`
		ActionFrameResp int `json:"action-frame-resp"`
		ReassocReq      int `json:"reassoc-req"`
	} `json:"client-dms-stats"`
	ClientRoamingStats struct {
		TotalRoam            int `json:"total-roam"`
		CckmRoam             int `json:"cckm-roam"`
		Dot11RRoam           int `json:"dot11r-roam"`
		Dot11IFastRoam       int `json:"dot11i-fast-roam"`
		Dot11ISlowRoam       int `json:"dot11i-slow-roam"`
		RoamFail             int `json:"roam-fail"`
		ApAuthRoams          int `json:"ap-auth-roams"`
		ApAuthDot11IFastRoam int `json:"ap-auth-dot11i-fast-roam"`
		ApAuthDot11ISlowRoam int `json:"ap-auth-dot11i-slow-roam"`
		Flex11RRoam          int `json:"flex-11r-roam"`
		Dot11RSlowRoam       int `json:"dot11r-slow-roam"`
		Pmkr0NameMismatch    int `json:"pmkr0-name-mismatch"`
		Pmkr1NameMismatch    int `json:"pmkr1-name-mismatch"`
	} `json:"client-roaming-stats"`
	LoadBalanceStats struct {
		TotalDenied      int `json:"total-denied"`
		TotalDeniedSent  int `json:"total-denied-sent"`
		TotalExMaxDenial int `json:"total-ex-max-denial"`
		Cand5G           int `json:"cand-5g"`
		Cand24G          int `json:"cand-24g"`
	} `json:"load-balance-stats"`
	OtherRoamAttempts   int `json:"other-roam-attempts"`
	AidAllocationFail   int `json:"aid-allocation-fail"`
	AidFreeFail         int `json:"aid-free-fail"`
	ClientDot11RFtStats struct {
		TotalAuthReqRx         int `json:"total-auth-req-rx"`
		TotalAuthRespSuccess   int `json:"total-auth-resp-success"`
		TotalAuthRespFail      int `json:"total-auth-resp-fail"`
		TotalActionReqRx       int `json:"total-action-req-rx"`
		TotalActionRespSuccess int `json:"total-action-resp-success"`
		TotalActionRespFail    int `json:"total-action-resp-fail"`
	} `json:"client-dot11r-ft-stats"`
	NumClientsOn24GhzRadio int `json:"num-clients-on-24ghz-radio"`
	NumClientsOn5GhzRadio  int `json:"num-clients-on-5ghz-radio"`
	Wpa3SaeStats           struct {
		Attempts                int `json:"attempts"`
		ProtocolIncomplete      int `json:"protocol-incomplete"`
		CommitRx                int `json:"commit-rx"`
		CommitRejected          int `json:"commit-rejected"`
		CommitUnsupportedGroup  int `json:"commit-unsupported-group"`
		CommitSent              int `json:"commit-sent"`
		ConfirmRx               int `json:"confirm-rx"`
		ConfirmRejected         int `json:"confirm-rejected"`
		ConfirmFieldMismatch    int `json:"confirm-field-mismatch"`
		ConfirmMsgInvalidLength int `json:"confirm-msg-invalid-length"`
		ConfirmSent             int `json:"confirm-sent"`
		OpenSessions            int `json:"open-sessions"`
		Accepted                int `json:"accepted"`
		H2ECommitRx             int `json:"h2e-commit-rx"`
		HnpCommitRx             int `json:"hnp-commit-rx"`
		H2ECommitPweMismatch    int `json:"h2e-commit-pwe-mismatch"`
		HnpCommitPweMismatch    int `json:"hnp-commit-pwe-mismatch"`
	} `json:"wpa3-sae-stats"`
	MdidMismatch           int `json:"mdid-mismatch"`
	WifiDirectAssocFail    int `json:"wifi-direct-assoc-fail"`
	WifiDirectAssocSuccess int `json:"wifi-direct-assoc-success"`
	Num6GhzClients         int `json:"num-6ghz-clients"`
}

// ClientLatencyStats represents client latency statistics including state durations and timing information.
type ClientLatencyStats struct {
	ClientStatesStats struct {
		ClientStateStatsValue []struct {
			AvgClientStateDuration string `json:"avg-client-state-duration"`
			TotalSessions          string `json:"total-sessions"`
		} `json:"client-state-stats-value"`
	} `json:"client-states-stats"`
	AvgRunStateLatency struct {
		RunState              int `json:"run-state"`
		RunStateSansUserDelay int `json:"run-state-sans-user-delay"`
	} `json:"avg-run-state-latency"`
}

// HTTPStats represents HTTP-related statistics including request and event counts.
type HTTPStats struct {
	HTTPRequestCount   int `json:"http-request-count"`
	ReadEventCount     int `json:"read-event-count"`
	ReadCompleteCount  int `json:"read-complete-count"`
	WriteEventCount    int `json:"write-event-count"`
	WriteCompleteCount int `json:"write-complete-count"`
	AaaMessageCount    int `json:"aaa-message-count"`
	SslEventStats      struct {
		EventReadBlockCount  int `json:"event-read-block-count"`
		EventWriteBlockCount int `json:"event-write-block-count"`
		EventOkCount         int `json:"event-ok-count"`
	} `json:"ssl-event-stats"`
	HTTPNewReqStats struct {
		NewReqNoSessionError int `json:"new-req-no-session-error"`
		NewReqCtxExistsError int `json:"new-req-ctx-exists-error"`
	} `json:"http-new-req-stats"`
	NumOfSocketOpened int `json:"num-of-socket-opened"`
	NumOfSocketClosed int `json:"num-of-socket-closed"`
}

// SmWebauthStats represents session manager web authentication statistics including HTTP and I/O metrics.
type SmWebauthStats struct {
	HTTPStats  HTTPStats `json:"http-stats"`
	IomReading struct {
		Total   int `json:"total"`
		Max     int `json:"max"`
		Min     int `json:"min"`
		Samples int `json:"samples"`
	} `json:"iom-reading"`
	MethodReading struct {
		Total   int `json:"total"`
		Max     int `json:"max"`
		Min     int `json:"min"`
		Samples int `json:"samples"`
	} `json:"method-reading"`
	IomWriting struct {
		Total   int `json:"total"`
		Max     int `json:"max"`
		Min     int `json:"min"`
		Samples int `json:"samples"`
	} `json:"iom-writing"`
	MethodWriting struct {
		Total   int `json:"total"`
		Max     int `json:"max"`
		Min     int `json:"min"`
		Samples int `json:"samples"`
	} `json:"method-writing"`
	IomAaa struct {
		Total   int `json:"total"`
		Max     int `json:"max"`
		Min     int `json:"min"`
		Samples int `json:"samples"`
	} `json:"iom-aaa"`
	MethodAaa struct {
		Total   int `json:"total"`
		Max     int `json:"max"`
		Min     int `json:"min"`
		Samples int `json:"samples"`
	} `json:"method-aaa"`
	NumOfSleepingClients int `json:"num-of-sleeping-clients"`
	SessionCount         int `json:"session-count"`
	HalfOpenCount        int `json:"half-open-count"`
	BackpressureCounters struct {
		SslHandshakePending    int `json:"ssl-handshake-pending"`
		HTTPSNewRequestPending int `json:"https-new-request-pending"`
		AaaReplyPending        int `json:"aaa-reply-pending"`
	} `json:"backpressure-counters"`
}

// Dot1XGlobalStats represents global 802.1X statistics including EAPOL message counters.
type Dot1XGlobalStats struct {
	EapolRx            int `json:"eapol-rx"`
	EapolRxStart       int `json:"eapol-rx-start"`
	EapolRxLogoff      int `json:"eapol-rx-logoff"`
	EapolRxResp        int `json:"eapol-rx-resp"`
	EapolRxRespID      int `json:"eapol-rx-resp-id"`
	EapolRxReq         int `json:"eapol-rx-req"`
	EapolRxInvalid     int `json:"eapol-rx-invalid"`
	EapolRxLenError    int `json:"eapol-rx-len-error"`
	EapolTx            int `json:"eapol-tx"`
	EapolTxStart       int `json:"eapol-tx-start"`
	EapolTxLogoff      int `json:"eapol-tx-logoff"`
	EapolTxResp        int `json:"eapol-tx-resp"`
	EapolTxReq         int `json:"eapol-tx-req"`
	EapolRetxReq       int `json:"eapol-retx-req"`
	EapolRetxReqFail   int `json:"eapol-retx-req-fail"`
	EapolTxReqID       int `json:"eapol-tx-req-id"`
	EapolRetxReqID     int `json:"eapol-retx-req-id"`
	EapolRetxReqIDFail int `json:"eapol-retx-req-id-fail"`
}

// ClientExclusionStats represents client exclusion statistics including excluded and disabled client counts.
type ClientExclusionStats struct {
	ExcludedClients int `json:"excluded-clients"`
	DisabledClients int `json:"disabled-clients"`
}

// SmDeviceCount represents session manager device count data by device type.
type SmDeviceCount struct {
	SmDeviceList []struct {
		DeviceType  string `json:"device-type"`
		DeviceCount int    `json:"device-count"`
	} `json:"sm-device-list"`
}

// TofStats represents time-of-flight statistics containing ToF tags.
type TofStats struct {
	TofTag []string `json:"tof-tag"`
}

// GetClientGlobalOper retrieves client global operational data.
func GetClientGlobalOper(client *wnc.Client, ctx context.Context) (*ClientGlobalOperResponse, error) {
	var data ClientGlobalOperResponse
	if err := client.SendAPIRequest(ctx, ClientGlobalOperEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetClientLiveStats retrieves client live statistics.
func GetClientLiveStats(client *wnc.Client, ctx context.Context) (*ClientLiveStatsResponse, error) {
	var data ClientLiveStatsResponse
	if err := client.SendAPIRequest(ctx, ClientLiveStatsEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetClientGlobalStatsData retrieves client global statistics data.
func GetClientGlobalStatsData(client *wnc.Client, ctx context.Context) (*ClientGlobalStatsDataResponse, error) {
	var data ClientGlobalStatsDataResponse
	if err := client.SendAPIRequest(ctx, ClientGlobalStatsDataEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetClientStats retrieves client statistics.
func GetClientStats(client *wnc.Client, ctx context.Context) (*ClientStatsResponse, error) {
	var data ClientStatsResponse
	if err := client.SendAPIRequest(ctx, ClientStatsEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetClientDot11Stats retrieves client 802.11 statistics.
func GetClientDot11Stats(client *wnc.Client, ctx context.Context) (*ClientDot11StatsResponse, error) {
	var data ClientDot11StatsResponse
	if err := client.SendAPIRequest(ctx, ClientDot11StatsEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetClientLatencyStats retrieves client latency statistics.
func GetClientLatencyStats(client *wnc.Client, ctx context.Context) (*ClientLatencyStatsResponse, error) {
	var data ClientLatencyStatsResponse
	if err := client.SendAPIRequest(ctx, ClientLatencyStatsEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetClientSmWebauthStats retrieves session manager web authentication statistics.
func GetClientSmWebauthStats(client *wnc.Client, ctx context.Context) (*SmWebauthStatsResponse, error) {
	var data SmWebauthStatsResponse
	if err := client.SendAPIRequest(ctx, SmWebauthStatsEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetClientDot1XGlobalStats retrieves 802.1X global statistics.
func GetClientDot1XGlobalStats(client *wnc.Client, ctx context.Context) (*Dot1XGlobalStatsResponse, error) {
	var data Dot1XGlobalStatsResponse
	if err := client.SendAPIRequest(ctx, Dot1XGlobalStatsEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetClientExclusionStats retrieves client exclusion statistics.
func GetClientExclusionStats(client *wnc.Client, ctx context.Context) (*ClientExclusionStatsResponse, error) {
	var data ClientExclusionStatsResponse
	if err := client.SendAPIRequest(ctx, ClientExclusionStatsEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetClientSmDeviceCount retrieves session manager device count data.
func GetClientSmDeviceCount(client *wnc.Client, ctx context.Context) (*SmDeviceCountResponse, error) {
	var data SmDeviceCountResponse
	if err := client.SendAPIRequest(ctx, SmDeviceCountEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetClientTofStats retrieves time-of-flight statistics.
func GetClientTofStats(client *wnc.Client, ctx context.Context) (*TofStatsResponse, error) {
	var data TofStatsResponse
	if err := client.SendAPIRequest(ctx, TofStatsEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}
