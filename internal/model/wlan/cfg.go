// Package wlan provides data structures for WLAN configuration data.
package wlan

// WlanCfg represents the complete WLAN configuration.
type WlanCfg struct {
	WlanCfgData *WlanCfgData `json:"Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data"` // WLAN configuration data container
}

// WlanCfgWlanCfgEntries represents the WLAN configuration entries from WNC 17.12.5.
type WlanCfgWlanCfgEntries struct {
	WlanCfgEntries *WlanCfgEntries `json:"Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-entries,omitempty"` // WLAN configuration entries container
}

// WlanCfgPolicyListEntries represents the policy list entries from WNC 17.12.5.
type WlanCfgPolicyListEntries struct {
	PolicyListEntries *PolicyListEntries `json:"Cisco-IOS-XE-wireless-wlan-cfg:policy-list-entries,omitempty"` // Policy list entries container
}

// WlanCfgWirelessAaaPolicyConfigs represents the wireless AAA policy configurations from WNC 17.12.5.
type WlanCfgWirelessAaaPolicyConfigs struct {
	WirelessAaaPolicyConfigs *WirelessAaaPolicyConfigs `json:"Cisco-IOS-XE-wireless-wlan-cfg:wireless-aaa-policy-configs,omitempty"` // Wireless AAA policy configurations container
}

// WlanCfgData represents the master WLAN configuration data structure.
type WlanCfgData struct {
	WlanCfgEntries           *WlanCfgEntries           `json:"wlan-cfg-entries,omitempty"`            // WLAN configuration entries
	WlanPolicies             *WlanPolicies             `json:"wlan-policies,omitempty"`               // WLAN policy profiles
	PolicyListEntries        *PolicyListEntries        `json:"policy-list-entries,omitempty"`         // Policy tag mappings
	WirelessAaaPolicyConfigs *WirelessAaaPolicyConfigs `json:"wireless-aaa-policy-configs,omitempty"` // Wireless AAA policy configurations

	// Wi-Fi 7 / 802.11be Support (YANG: IOS-XE 17.18.1+)
	Dot11beProfiles *Dot11beProfiles `json:"dot11be-profiles,omitempty"` // Wi-Fi 7 profile configurations (YANG: IOS-XE 17.18.1+)
}

// WlanCfgEntries represents the WLAN configuration entries response.
type WlanCfgEntries struct {
	WlanCfgEntry []WlanCfgEntry `json:"wlan-cfg-entry"` // List of WLAN configuration entries
}

// WlanPolicies represents the WLAN policies configuration response.
type WlanPolicies struct {
	WlanPolicy []WlanPolicy `json:"wlan-policy"` // List of WLAN policy profiles
}

// PolicyListEntries represents the policy list entries response.
type PolicyListEntries struct {
	PolicyListEntry []PolicyListEntry `json:"policy-list-entry"` // List of policy tag entries
}

// WirelessAaaPolicyConfigs represents the wireless AAA policy configurations response.
type WirelessAaaPolicyConfigs struct {
	WirelessAaaPolicyConfig []WirelessAaaPolicyConfig `json:"wireless-aaa-policy-config"` // List of wireless AAA policy configurations
}

// WlanCfgEntry represents a single WLAN configuration entry.
type WlanCfgEntry struct {
	WlanID                 int                `json:"wlan-id"`                              // WLAN identifier number
	ProfileName            string             `json:"profile-name"`                         // WLAN profile name
	AuthKeyMgmtPsk         bool               `json:"auth-key-mgmt-psk,omitempty"`          // PSK authentication enabled
	AuthKeyMgmtDot1x       bool               `json:"auth-key-mgmt-dot1x,omitempty"`        // 802.1X authentication enabled
	AuthKeyMgmtDot1xSha256 bool               `json:"auth-key-mgmt-dot1x-sha256,omitempty"` // 802.1X SHA-256 authentication enabled
	PSK                    string             `json:"psk,omitempty"`                        // Pre-shared key
	PSKType                string             `json:"psk-type,omitempty"`                   // PSK encryption type
	FTMode                 string             `json:"ft-mode,omitempty"`                    // Fast transition mode
	PMFOptions             string             `json:"pmf-options,omitempty"`                // Protected Management Frame options
	WPA2Enabled            bool               `json:"wpa2-enabled,omitempty"`               // WPA2 protocol enabled
	WPA3Enabled            bool               `json:"wpa3-enabled,omitempty"`               // WPA3 protocol enabled
	LoadBalance            bool               `json:"load-balance,omitempty"`               // Load balancing enabled
	AuthenticationList     string             `json:"authentication-list,omitempty"`        // Authentication method list
	Wlan11kNeighList       bool               `json:"wlan-11k-neigh-list,omitempty"`        // 802.11k neighbor list enabled
	MulticastBufferValue   int                `json:"multicast-buffer-value,omitempty"`     // Multicast buffer size value
	ApfVapIDData           *ApfVapIDData      `json:"apf-vap-id-data,omitempty"`            // Virtual AP interface data
	ApfVap80211vData       *ApfVap80211vData  `json:"apf-vap-802-11v-data,omitempty"`       // 802.11v management data
	MdnsSDMode             string             `json:"mdns-sd-mode,omitempty"`               // mDNS service discovery mode
	WlanRadioPolicies      *WlanRadioPolicies `json:"wlan-radio-policies,omitempty"`        // Radio band policies
	ClientSteering         bool               `json:"client-steering,omitempty"`            // Band steering enabled
	WepKeyIndex            int                `json:"wep-key-index,omitempty"`              // WEP key index
}

// ApfVapIDData represents virtual AP interface identification data.
type ApfVapIDData struct {
	SSID       string `json:"ssid"`        // Service Set Identifier
	WlanStatus bool   `json:"wlan-status"` // WLAN administrative status
}

// ApfVap80211vData represents 802.11v wireless management configuration.
type ApfVap80211vData struct {
	Dot11vDms bool `json:"dot11v-dms"` // 802.11v Directed Multicast Service enabled
}

// WlanRadioPolicies represents WLAN radio band policy configuration.
type WlanRadioPolicies struct {
	WlanRadioPolicy []WlanRadioPolicy `json:"wlan-radio-policy"` // List of radio band policies
}

// WlanRadioPolicy represents individual radio band policy configuration.
type WlanRadioPolicy struct {
	Band string `json:"band"` // Radio band specification
}

// WlanPolicy represents WLAN policy profile configuration.
type WlanPolicy struct {
	PolicyProfileName       string                   `json:"policy-profile-name"`                   // Policy profile identifier name
	Description             string                   `json:"description,omitempty"`                 // Policy description
	Status                  bool                     `json:"status,omitempty"`                      // Policy enable status
	InterfaceName           string                   `json:"interface-name,omitempty"`              // Associated interface name
	WlanSwitchingPolicy     *WlanSwitchingPolicy     `json:"wlan-switching-policy,omitempty"`       // Switching policy configuration
	WlanTimeout             *WlanTimeout             `json:"wlan-timeout,omitempty"`                // Timeout policy configuration
	PerSsidQos              *PerSsidQos              `json:"per-ssid-qos,omitempty"`                // Quality of Service configuration
	DhcpParams              *DhcpParams              `json:"dhcp-params,omitempty"`                 // DHCP parameters configuration
	UmbrellaFlexParams      *UmbrellaFlexParams      `json:"umbrella-flex-params,omitempty"`        // Umbrella Flex parameters
	AtfPolicyMapEntries     *AtfPolicyMapEntries     `json:"atf-policy-map-entries,omitempty"`      // Airtime Fairness policy entries
	AvcIPv4FmIngressEntries *AvcIPv4FmIngressEntries `json:"avc-ipv4-fm-ingress-entries,omitempty"` // AVC IPv4 ingress flow monitor entries
	AvcIPv4FmEgressEntries  *AvcIPv4FmEgressEntries  `json:"avc-ipv4-fm-egress-entries,omitempty"`  // AVC IPv4 egress flow monitor entries
	AvcIPv6FmIngressEntries *AvcIPv6FmIngressEntries `json:"avc-ipv6-fm-ingress-entries,omitempty"` // AVC IPv6 ingress flow monitor entries
	AvcIPv6FmEgressEntries  *AvcIPv6FmEgressEntries  `json:"avc-ipv6-fm-egress-entries,omitempty"`  // AVC IPv6 egress flow monitor entries
}

// WlanSwitchingPolicy represents WLAN switching policy configuration.
type WlanSwitchingPolicy struct {
	CentralSwitching      bool `json:"central-switching,omitempty"`      // Central switching enabled
	CentralAuthentication bool `json:"central-authentication,omitempty"` // Central authentication enabled
	CentralDhcp           bool `json:"central-dhcp,omitempty"`           // Central DHCP enabled
	CentralAssocEnable    bool `json:"central-assoc-enable,omitempty"`   // Central association enabled
}

// WlanTimeout represents WLAN timeout configuration.
type WlanTimeout struct {
	SessionTimeout int `json:"session-timeout,omitempty"` // Session timeout in seconds
}

// PerSsidQos represents per-SSID QoS configuration.
type PerSsidQos struct {
	IngressServiceName string `json:"ingress-service-name,omitempty"` // Ingress QoS service name
	EgressServiceName  string `json:"egress-service-name,omitempty"`  // Egress QoS service name
}

// DhcpParams represents DHCP parameters configuration.
type DhcpParams struct {
	IsDhcpEnabled bool `json:"is-dhcp-enabled,omitempty"` // DHCP service enabled
}

// UmbrellaFlexParams represents Umbrella Flex parameters configuration.
type UmbrellaFlexParams struct {
	DhcpDNSOptionEnable bool `json:"dhcp-dns-option-enable,omitempty"` // DHCP DNS option enabled
}

// AtfPolicyMapEntries represents ATF (Airtime Fairness) policy map entries from live WNC 17.12.1.
type AtfPolicyMapEntries struct {
	AtfPolicyMapEntry []AtfPolicyMapEntry `json:"atf-policy-map-entry"` // List of ATF policy entries
}

// AtfPolicyMapEntry represents a single ATF policy map entry from live WNC 17.12.1.
type AtfPolicyMapEntry struct {
	BandID        int    `json:"band-id"`         // Radio band identifier
	AtfPolicyName string `json:"atf-policy-name"` // ATF policy name
}

// AvcIPv4FmIngressEntries represents AVC IPv4 flow monitor ingress entries from live WNC 17.12.1.
type AvcIPv4FmIngressEntries struct {
	AvcIPv4FmIngressEntry []AvcIPv4FmIngressEntry `json:"avc-ipv4-fm-ingress-entry"` // List of IPv4 ingress flow monitors
}

// AvcIPv4FmIngressEntry represents a single AVC IPv4 flow monitor ingress entry from live WNC 17.12.1.
type AvcIPv4FmIngressEntry struct {
	Name string `json:"name"` // Flow monitor name
}

// AvcIPv4FmEgressEntries represents AVC IPv4 flow monitor egress entries from live WNC 17.12.1.
type AvcIPv4FmEgressEntries struct {
	AvcIPv4FmEgressEntry []AvcIPv4FmEgressEntry `json:"avc-ipv4-fm-egress-entry"` // List of IPv4 egress flow monitors
}

// AvcIPv4FmEgressEntry represents a single AVC IPv4 flow monitor egress entry from live WNC 17.12.1.
type AvcIPv4FmEgressEntry struct {
	Name string `json:"name"` // Flow monitor name
}

// AvcIPv6FmIngressEntries represents AVC IPv6 flow monitor ingress entries from live WNC 17.12.1.
type AvcIPv6FmIngressEntries struct {
	AvcIPv6FmIngressEntry []AvcIPv6FmIngressEntry `json:"avc-ipv6-fm-ingress-entry"` // List of IPv6 ingress flow monitors
}

// AvcIPv6FmIngressEntry represents a single AVC IPv6 flow monitor ingress entry from live WNC 17.12.1.
type AvcIPv6FmIngressEntry struct {
	Name string `json:"name"` // Flow monitor name
}

// AvcIPv6FmEgressEntries represents AVC IPv6 flow monitor egress entries from live WNC 17.12.1.
type AvcIPv6FmEgressEntries struct {
	AvcIPv6FmEgressEntry []AvcIPv6FmEgressEntry `json:"avc-ipv6-fm-egress-entry"` // List of IPv6 egress flow monitors
}

// AvcIPv6FmEgressEntry represents a single AVC IPv6 flow monitor egress entry from live WNC 17.12.1.
type AvcIPv6FmEgressEntry struct {
	Name string `json:"name"` // Flow monitor name
}

// PolicyListEntry represents individual policy list entry.
type PolicyListEntry struct {
	TagName      string        `json:"tag-name,omitempty"`      // Policy tag name
	Description  string        `json:"description,omitempty"`   // Policy tag description
	WLANPolicies *WLANPolicies `json:"wlan-policies,omitempty"` // WLAN policy mappings
}

// WLANPolicies represents the container for WLAN policy mappings.
type WLANPolicies struct {
	WLANPolicy []WLANPolicyMap `json:"wlan-policy,omitempty"` // List of WLAN to policy mappings
}

// WLANPolicyMap represents a WLAN to policy profile mapping.
type WLANPolicyMap struct {
	WLANProfileName   string `json:"wlan-profile-name"`   // WLAN profile name
	PolicyProfileName string `json:"policy-profile-name"` // Policy profile name
}

// WirelessAaaPolicyConfig represents wireless AAA policy configuration.
type WirelessAaaPolicyConfig struct {
	PolicyName string `json:"policy-name"` // AAA policy name
}

// Dot11beProfiles represents Wi-Fi 7 / 802.11be profiles (YANG: IOS-XE 17.18.1+).
type Dot11beProfiles struct {
	Dot11beProfile []Dot11beProfile `json:"dot11be-profile"` // List of Wi-Fi 7 profiles (YANG: IOS-XE 17.18.1+)
}

// Dot11beProfile represents a single 802.11be profile (YANG: IOS-XE 17.18.1+).
type Dot11beProfile struct {
	ProfileName       string `json:"profile-name"`                  // Wi-Fi 7 profile identifier (YANG: IOS-XE 17.18.1+)
	Description       string `json:"description,omitempty"`         // Profile description (YANG: IOS-XE 17.18.1+)
	EhtOfdmaDownlink  bool   `json:"eht-ofdma-downlink,omitempty"`  // EHT OFDMA downlink enabled (YANG: IOS-XE 17.18.1+)
	EhtOfdmaUplink    bool   `json:"eht-ofdma-uplink,omitempty"`    // EHT OFDMA uplink enabled (YANG: IOS-XE 17.18.1+)
	EhtMumimoDownlink bool   `json:"eht-mumimo-downlink,omitempty"` // EHT MU-MIMO downlink enabled (YANG: IOS-XE 17.18.1+)
	EhtMumimoUplink   bool   `json:"eht-mumimo-uplink,omitempty"`   // EHT MU-MIMO uplink enabled (YANG: IOS-XE 17.18.1+)
	EhtOfdmaMultiRu   bool   `json:"eht-ofdma-multi-ru,omitempty"`  // EHT OFDMA multi-RU enabled (YANG: IOS-XE 17.18.1+)

	// Multi-link Operation (MLO) configurations (YANG: IOS-XE 17.18.1+)
	Mlo24Ghz   *MloGroup24Ghz   `json:"mlo-24ghz,omitempty"`    // MLO 2.4GHz link configuration (YANG: IOS-XE 17.18.1+)
	Mlo5Ghz    *MloGroup5Ghz    `json:"mlo-5ghz,omitempty"`     // MLO primary 5GHz link configuration (YANG: IOS-XE 17.18.1+)
	Mlo5GhzSec *MloGroup5GhzSec `json:"mlo-5ghz-sec,omitempty"` // MLO secondary 5GHz link configuration (YANG: IOS-XE 17.18.1+)
	Mlo6Ghz    *MloGroup6Ghz    `json:"mlo-6ghz,omitempty"`     // MLO 6GHz link configuration (YANG: IOS-XE 17.18.1+)
}

// MloGroup24Ghz represents Multi-link group configuration for 2.4GHz link (YANG: IOS-XE 17.18.1+).
type MloGroup24Ghz struct {
	GroupID int  `json:"group-id,omitempty"` // MLO group identifier (YANG: IOS-XE 17.18.1+)
	Enable  bool `json:"enable,omitempty"`   // MLO group enable status (YANG: IOS-XE 17.18.1+)
}

// MloGroup5Ghz represents Multi-link group configuration for primary 5GHz link (YANG: IOS-XE 17.18.1+).
type MloGroup5Ghz struct {
	GroupID int  `json:"group-id,omitempty"` // MLO group identifier (YANG: IOS-XE 17.18.1+)
	Enable  bool `json:"enable,omitempty"`   // MLO group enable status (YANG: IOS-XE 17.18.1+)
}

// MloGroup5GhzSec represents Multi-link group configuration for secondary 5GHz link (YANG: IOS-XE 17.18.1+).
type MloGroup5GhzSec struct {
	GroupID int  `json:"group-id,omitempty"` // MLO group identifier (YANG: IOS-XE 17.18.1+)
	Enable  bool `json:"enable,omitempty"`   // MLO group enable status (YANG: IOS-XE 17.18.1+)
}

// MloGroup6Ghz represents Multi-link group configuration for 6GHz link (YANG: IOS-XE 17.18.1+).
type MloGroup6Ghz struct {
	GroupID int  `json:"group-id,omitempty"` // MLO group identifier (YANG: IOS-XE 17.18.1+)
	Enable  bool `json:"enable,omitempty"`   // MLO group enable status (YANG: IOS-XE 17.18.1+)
}
