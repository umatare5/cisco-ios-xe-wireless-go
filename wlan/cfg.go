// Package wlan provides WLAN configuration functionality for the Cisco Wireless Network Controller API.
package wlan

import (
	"context"

	wnc "github.com/umatare5/cisco-xe-wireless-restconf-go"
)

const (
	// WlanCfgBasePath defines the base path for WLAN configuration endpoints
	WlanCfgBasePath = "/restconf/data/Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data"
	// WlanCfgEndpoint retrieves complete WLAN configuration data
	WlanCfgEndpoint = WlanCfgBasePath
	// WlanCfgEntriesEndpoint retrieves WLAN configuration entries
	WlanCfgEntriesEndpoint = WlanCfgBasePath + "/wlan-cfg-entries"
	// WlanPoliciesEndpoint retrieves WLAN policies
	WlanPoliciesEndpoint = WlanCfgBasePath + "/wlan-policies"
	// PolicyListEntriesEndpoint retrieves policy list entries
	PolicyListEntriesEndpoint = WlanCfgBasePath + "/policy-list-entries"
	// WirelessAaaPolicyConfigsEndpoint retrieves wireless AAA policy configurations
	WirelessAaaPolicyConfigsEndpoint = WlanCfgBasePath + "/wireless-aaa-policy-configs"
)

// WlanCfgResponse represents the complete WLAN configuration response
type WlanCfgResponse struct {
	CiscoIOSXEWirelessWlanCfgWlanCfgData struct {
		WlanCfgEntries           WlanCfgEntries           `json:"wlan-cfg-entries"`
		WlanPolicies             WlanPolicies             `json:"wlan-policies"`
		PolicyListEntries        PolicyListEntries        `json:"policy-list-entries"`
		WirelessAaaPolicyConfigs WirelessAaaPolicyConfigs `json:"wireless-aaa-policy-configs"`
	} `json:"Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data"`
}

// WlanCfgEntriesResponse represents the WLAN configuration entries response
type WlanCfgEntriesResponse struct {
	WlanCfgEntries WlanCfgEntries `json:"Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-entries"`
}

// WlanPoliciesResponse represents the WLAN policies configuration response
type WlanPoliciesResponse struct {
	WlanPolicies WlanPolicies `json:"Cisco-IOS-XE-wireless-wlan-cfg:wlan-policies"`
}

// PolicyListEntriesResponse represents the policy list entries response
type PolicyListEntriesResponse struct {
	PolicyListEntries PolicyListEntries `json:"Cisco-IOS-XE-wireless-wlan-cfg:policy-list-entries"`
}

// WirelessAaaPolicyConfigsResponse represents the wireless AAA policy configurations response
type WirelessAaaPolicyConfigsResponse struct {
	WirelessAaaPolicyConfigs WirelessAaaPolicyConfigs `json:"Cisco-IOS-XE-wireless-wlan-cfg:wireless-aaa-policy-configs"`
}

// WlanCfgEntries contains WLAN configuration entries
type WlanCfgEntries struct {
	WlanCfgEntry []WlanCfgEntry `json:"wlan-cfg-entry"`
}

// WlanCfgEntry represents a WLAN configuration entry with security and authentication settings
type WlanCfgEntry struct {
	ProfileName               string `json:"profile-name"`            // WLAN profile name
	WlanID                    int    `json:"wlan-id"`                 // WLAN identifier
	WepKeyIndex               int    `json:"wep-key-index,omitempty"` // WEP key index (deprecated)
	AuthKeyMgmtPsk            bool   `json:"auth-key-mgmt-psk"`       // PSK authentication enabled
	AuthKeyMgmtDot1x          bool   `json:"auth-key-mgmt-dot1x"`     // 802.1X authentication enabled
	Psk                       string `json:"psk"`                     // Pre-shared key
	PskType                   string `json:"psk-type"`                // PSK type (ascii/hex)
	QosWmmStatus              string `json:"qos-wmm-status,omitempty"`
	FtMode                    string `json:"ft-mode"`
	PmfOptions                string `json:"pmf-options,omitempty"`
	BandSteeringAllowed       bool   `json:"band-steering-allowed,omitempty"`
	MuMimo                    bool   `json:"mu-mimo,omitempty"`
	MulticastMcDirect         bool   `json:"multicast-mc-direct,omitempty"`
	MaxClientsAllowed         int    `json:"max-clients-allowed,omitempty"`
	MaxClientsPerApPerWlan    int    `json:"max-clients-per-ap-per-wlan,omitempty"`
	MaxClientsPerRadioPerWlan int    `json:"max-clients-per-radio-per-wlan,omitempty"`
	IPSourceGuardEnabled      bool   `json:"ip-source-guard-enabled,omitempty"`
	LoadBalance               bool   `json:"load-balance,omitempty"`
	Wlan11kNeighList          bool   `json:"wlan-11k-neigh-list,omitempty"`
	Wlan11kAssistedRoaming    bool   `json:"wlan-11k-assisted-roaming,omitempty"`
	Wlan11kDualBandNeighList  bool   `json:"wlan-11k-dual-band-neigh-list,omitempty"`
	MulticastBufferValue      int    `json:"multicast-buffer-value,omitempty"`

	ApfVapIDData struct {
		BroadcastSsid  bool   `json:"broadcast-ssid,omitempty"`
		P2PBlockAction string `json:"p2p-block-action,omitempty"`
		CcxAironetIe   bool   `json:"ccx-aironet-ie,omitempty"`
		Dot11aDtim     int    `json:"dot11a-dtim,omitempty"`
		Dot11bDtim     int    `json:"dot11b-dtim,omitempty"`
		Chd            bool   `json:"chd,omitempty"`
		SSID           string `json:"ssid"`
		WlanStatus     bool   `json:"wlan-status"`
	} `json:"apf-vap-id-data"`

	ApfVap80211vData struct {
		Dot11vDms                 bool `json:"dot11v-dms,omitempty"`
		Dot11vBssMaxIdleProtected bool `json:"dot11v-bss-max-idle-protected,omitempty"`
		Dot11vBssTransition       bool `json:"dot11v-bss-transition,omitempty"`
		Dot11vDualList            bool `json:"dot11v-dual-list,omitempty"`
	} `json:"apf-vap-80211v-data"`
	VapDot11axCfg struct {
		HeOfdmaDownlink  bool `json:"he-ofdma-downlink,omitempty"`
		HeOfdmaUplink    bool `json:"he-ofdma-uplink,omitempty"`
		HeMuMimoDownlink bool `json:"he-mumimo-downlink,omitempty"`
		HeMuMimoUplink   bool `json:"he-mumimo-uplink,omitempty"`
		HeTwtEnable      bool `json:"he-twt-enable,omitempty"`
	} `json:"vap-dot11ax-cfg"`
	MdnsSdMode      string `json:"mdns-sd-mode"`
	Wpa3Enabled     bool   `json:"wpa3-enabled,omitempty"`
	AuthKeyMgmtSae  bool   `json:"auth-key-mgmt-sae,omitempty"`
	DeviceAnalytics struct {
		DaReport bool `json:"da-report,omitempty"`
	} `json:"device-analytics"`
	Dot11kRmBeaconMeasReq struct {
		OnAssoc bool `json:"on-assoc,omitempty"`
		OnRoam  bool `json:"on-roam,omitempty"`
	} `json:"dot11k-rm-beacon-meas-req"`
	LaaParams struct {
		LaaClientDenial bool `json:"laa-client-denial,omitempty"`
	} `json:"laa-params"`
	WlanRadioPolicies struct {
		WlanRadioPolicy []struct {
			Band string `json:"band"`
		} `json:"wlan-radio-policy"`
	} `json:"wlan-radio-policies"`
	ClientSteering bool `json:"client-steering,omitempty"`
	LatencyMa      bool `json:"latency-ma,omitempty"`
}

type WlanPolicies struct {
	WlanPolicy []WlanPolicy `json:"wlan-policy"`
}

type WlanPolicy struct {
	PolicyProfileName   string `json:"policy-profile-name"`
	Description         string `json:"description,omitempty"`
	Status              bool   `json:"status"`
	InterfaceName       string `json:"interface-name,omitempty"`
	WlanSwitchingPolicy struct {
		CentralSwitching      bool `json:"central-switching"`
		CentralAuthentication bool `json:"central-authentication,omitempty"`
		CentralDhcp           bool `json:"central-dhcp"`
		CentralAssocEnable    bool `json:"central-assoc-enable,omitempty"`
	} `json:"wlan-switching-policy"`
	WlanTimeout struct {
		SessionTimeout int `json:"session-timeout"`
	} `json:"wlan-timeout,omitempty"`
	PerSsidQos struct {
		IngressServiceName string `json:"ingress-service-name"`
		EgressServiceName  string `json:"egress-service-name"`
	} `json:"per-ssid-qos,omitempty"`
	DhcpParams struct {
		IsDhcpEnabled bool `json:"is-dhcp-enabled"`
	} `json:"dhcp-params"`
	AtfPolicyMapEntries struct {
		Entries []struct {
			BandID        int    `json:"band-id,omitempty"`
			AtfPolicyName string `json:"atf-policy-name,omitempty"`
		} `json:"atf-policy-map-entry,omitempty"`
	} `json:"atf-policy-map-entries"`
	AvcIpv4FmIngressEntries struct {
		Entries []struct {
			Name string `json:"name"`
		} `json:"avc-ipv4-fm-ingress-entry"`
	} `json:"avc-ipv4-fm-ingress-entries,omitempty"`
	AvcIpv4FmEgressEntries struct {
		Entries []struct {
			Name string `json:"name"`
		} `json:"avc-ipv4-fm-egress-entry"`
	} `json:"avc-ipv4-fm-egress-entries,omitempty"`
	AvcIpv6FmIngressEntries struct {
		Entries []struct {
			Name string `json:"name"`
		} `json:"avc-ipv6-fm-ingress-entry"`
	} `json:"avc-ipv6-fm-ingress-entries,omitempty"`
	AvcIpv6FmEgressEntries struct {
		Entries []struct {
			Name string `json:"name"`
		} `json:"avc-ipv6-fm-egress-entry"`
	} `json:"avc-ipv6-fm-egress-entries,omitempty"`
	UmbrellaFlexParams struct {
		DhcpDnsOptionEnable bool `json:"dhcp-dns-option-enable"`
	} `json:"umbrella-flex-params,omitempty"`
	WlanLocalProfiling struct {
		HttpTlvCaching bool `json:"http-tlv-caching"`
		DhcpTlvCaching bool `json:"dhcp-tlv-caching"`
	} `json:"wlan-local-profiling,omitempty"`
}

type PolicyListEntries struct {
	PolicyListEntry []PolicyListEntry `json:"policy-list-entry"`
}

type PolicyListEntry struct {
	TagName      string `json:"tag-name"`
	Description  string `json:"description,omitempty"`
	WlanPolicies struct {
		WlanPolicy []struct {
			WlanProfileName   string `json:"wlan-profile-name"`
			PolicyProfileName string `json:"policy-profile-name"`
		} `json:"wlan-policy"`
	} `json:"wlan-policies,omitempty"`
}

type WirelessAaaPolicyConfigs struct {
	WirelessAaaPolicyConfig []WirelessAaaPolicyConfig `json:"wireless-aaa-policy-config"`
}

type WirelessAaaPolicyConfig struct {
	PolicyName string `json:"policy-name"`
}

// GetWlanCfg retrieves WLAN configuration with context support
func GetWlanCfg(client *wnc.Client, ctx context.Context) (*WlanCfgResponse, error) {
	var data WlanCfgResponse
	if err := client.SendAPIRequest(ctx, WlanCfgEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

func GetWlanCfgEntries(client *wnc.Client, ctx context.Context) (*WlanCfgEntriesResponse, error) {
	var data WlanCfgEntriesResponse
	if err := client.SendAPIRequest(ctx, WlanCfgEntriesEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

func GetWlanPolicies(client *wnc.Client, ctx context.Context) (*WlanPoliciesResponse, error) {
	var data WlanPoliciesResponse
	if err := client.SendAPIRequest(ctx, WlanPoliciesEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

func GetPolicyListEntries(client *wnc.Client, ctx context.Context) (*PolicyListEntriesResponse, error) {
	var data PolicyListEntriesResponse
	if err := client.SendAPIRequest(ctx, PolicyListEntriesEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

func GetWirelessAaaPolicyConfigs(client *wnc.Client, ctx context.Context) (*WirelessAaaPolicyConfigsResponse, error) {
	var data WirelessAaaPolicyConfigsResponse
	if err := client.SendAPIRequest(ctx, WirelessAaaPolicyConfigsEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}
