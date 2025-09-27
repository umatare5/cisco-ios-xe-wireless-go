package wlan

// CiscoIOSXEWirelessWlanCfg represents the complete WLAN configuration.
type CiscoIOSXEWirelessWlanCfg struct {
	CiscoIOSXEWirelessWlanCfgData struct {
		WlanCfgEntries           *WlanCfgEntries          `json:"wlan-cfg-entries"`            // WLAN configuration parameters (Live: IOS-XE 17.12.5)
		WlanPolicies             *WlanPolicies            `json:"wlan-policies"`               // WLAN policy configuration (Live: IOS-XE 17.12.5)
		PolicyListEntries        PolicyListEntries        `json:"policy-list-entries"`         // Policy list configuration (Live: IOS-XE 17.12.5)
		WirelessAaaPolicyConfigs WirelessAaaPolicyConfigs `json:"wireless-aaa-policy-configs"` // Wireless AAA policy configurations (Live: IOS-XE 17.12.5)
		Dot11beProfiles          *Dot11beProfiles         `json:"dot11be-profiles"`            // 802.11be profile parameters (Live: IOS-XE 17.15.4b)
	} `json:"Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data"` // WLAN configuration data container
}

// CiscoIOSXEWirelessWlanCfgWlanCfgEntries represents the WLAN configuration entries.
type CiscoIOSXEWirelessWlanCfgWlanCfgEntries struct {
	WlanCfgEntries *WlanCfgEntries `json:"Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-entries"`
}

// CiscoIOSXEWirelessWlanCfgPolicyListEntries represents the policy list entries.
type CiscoIOSXEWirelessWlanCfgPolicyListEntries struct {
	PolicyListEntries *PolicyListEntries `json:"Cisco-IOS-XE-wireless-wlan-cfg:policy-list-entries"`
}

// CiscoIOSXEWirelessWlanCfgWirelessAaaPolicyConfigs represents the wireless AAA policy configurations.
type CiscoIOSXEWirelessWlanCfgWirelessAaaPolicyConfigs struct {
	WirelessAaaPolicyConfigs *WirelessAaaPolicyConfigs `json:"Cisco-IOS-XE-wireless-wlan-cfg:wireless-aaa-policy-configs"`
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
	WlanID                 int                `json:"wlan-id"`                              // WLAN identifier (Live: IOS-XE 17.12.5)
	ProfileName            string             `json:"profile-name"`                         // WLAN profile name (Live: IOS-XE 17.12.5)
	AuthKeyMgmtPsk         bool               `json:"auth-key-mgmt-psk,omitempty"`          // Authentication key management PSK (Live: IOS-XE 17.12.5)
	AuthKeyMgmtDot1x       bool               `json:"auth-key-mgmt-dot1x,omitempty"`        // Authentication key management type 802.1x (Live: IOS-XE 17.12.5)
	AuthKeyMgmtDot1xSha256 bool               `json:"auth-key-mgmt-dot1x-sha256,omitempty"` // Authentication key management type 802.1x SHA256 (Live: IOS-XE 17.12.5)
	PSK                    string             `json:"psk,omitempty"`                        // Authentication pre-shared key (Live: IOS-XE 17.12.5)
	PSKType                string             `json:"psk-type,omitempty"`                   // Pre-shared key encryption type (Live: IOS-XE 17.12.5)
	FTMode                 string             `json:"ft-mode,omitempty"`                    // Configures Fast Transition Adaptive support (Live: IOS-XE 17.12.5)
	PMFOptions             string             `json:"pmf-options,omitempty"`                // Configures PMF as optional/required (Live: IOS-XE 17.12.5)
	WPA2Enabled            bool               `json:"wpa2-enabled,omitempty"`               // Configures WPA2 support (Live: IOS-XE 17.12.5)
	WPA3Enabled            bool               `json:"wpa3-enabled,omitempty"`               // Configures WPA3 support (Live: IOS-XE 17.12.5)
	LoadBalance            bool               `json:"load-balance,omitempty"`               // Allow/Disallow Load Balance on a WLAN (Live: IOS-XE 17.12.5)
	AuthenticationList     string             `json:"authentication-list,omitempty"`        // Enter the Authentication list name (Live: IOS-XE 17.12.5)
	Wlan11kNeighList       bool               `json:"wlan-11k-neigh-list,omitempty"`        // Indicates 11k Neighbor List enabled (YANG: IOS-XE 17.12.1)
	MulticastBufferValue   int                `json:"multicast-buffer-value,omitempty"`     // Configure Multicast Buffer Tuning (YANG: IOS-XE 17.12.1)
	APFVapIDData           *APFVapIDData      `json:"apf-vap-id-data,omitempty"`            // Virtual AP interface data (Live: IOS-XE 17.12.5)
	APFVap80211vData       *APFVap80211vData  `json:"apf-vap-802-11v-data,omitempty"`       // 802.11v wireless management configuration (Live: IOS-XE 17.12.5)
	MDNSSDMode             string             `json:"mdns-sd-mode,omitempty"`               // MDNS operational mode on WLAN (Live: IOS-XE 17.12.5)
	WlanRadioPolicies      *WlanRadioPolicies `json:"wlan-radio-policies,omitempty"`        // WLAN radio policy (Live: IOS-XE 17.12.5)
	ClientSteering         bool               `json:"client-steering,omitempty"`            // Enable/disable 6Ghz client steering (YANG: IOS-XE 17.12.1)
	WepKeyIndex            int                `json:"wep-key-index,omitempty"`              // WEP key index for Static WEP Authentication (Live: IOS-XE 17.12.5)
}

// APFVapIDData represents virtual AP interface identification data.
type APFVapIDData struct {
	SSID       string `json:"ssid"`        // Service Set Identifier (Live: IOS-XE 17.12.5)
	WlanStatus bool   `json:"wlan-status"` // WLAN administrative status (Live: IOS-XE 17.12.5)
}

// APFVap80211vData represents 802.11v wireless management configuration.
type APFVap80211vData struct {
	Dot11vDms bool `json:"dot11v-dms"` // 802.11v Directed Multicast Service enabled (Live: IOS-XE 17.12.5)
}

// WlanRadioPolicies represents WLAN radio band policy configuration.
type WlanRadioPolicies struct {
	WlanRadioPolicy []WlanRadioPolicy `json:"wlan-radio-policy"` // List of radio band policies
}

// WlanRadioPolicy represents individual radio band policy configuration.
type WlanRadioPolicy struct {
	Band string `json:"band"` // Radio band specification (Live: IOS-XE 17.12.5)
}

// WlanPolicy represents WLAN policy profile configuration.
type WlanPolicy struct {
	PolicyProfileName       string                   `json:"policy-profile-name"`                   // This object specifies policy profile instance (Live: IOS-XE 17.12.5)
	Description             string                   `json:"description,omitempty"`                 // Description associated to WLAN policy (Live: IOS-XE 17.12.5)
	Status                  bool                     `json:"status,omitempty"`                      // Whether policy profile is shutdown or active (Live: IOS-XE 17.12.5)
	InterfaceName           string                   `json:"interface-name,omitempty"`              // Interface attached to the wireless lan (Live: IOS-XE 17.12.5)
	WlanSwitchingPolicy     *WlanSwitchingPolicy     `json:"wlan-switching-policy,omitempty"`       // Switching policy configuration
	WlanTimeout             *WlanTimeout             `json:"wlan-timeout,omitempty"`                // Timeout policy configuration
	PerSsidQos              *PerSsidQos              `json:"per-ssid-qos,omitempty"`                // Quality of Service configuration
	DHCPParams              *DHCPParams              `json:"dhcp-params,omitempty"`                 // DHCP parameters configuration
	UmbrellaFlexParams      *UmbrellaFlexParams      `json:"umbrella-flex-params,omitempty"`        // Umbrella Flex parameters
	AtfPolicyMapEntries     *AtfPolicyMapEntries     `json:"atf-policy-map-entries,omitempty"`      // Airtime Fairness policy entries
	AvcIPv4FmIngressEntries *AvcIPv4FmIngressEntries `json:"avc-ipv4-fm-ingress-entries,omitempty"` // AVC IPv4 ingress flow monitor entries
	AvcIPv4FmEgressEntries  *AvcIPv4FmEgressEntries  `json:"avc-ipv4-fm-egress-entries,omitempty"`  // AVC IPv4 egress flow monitor entries
	AvcIPv6FmIngressEntries *AvcIPv6FmIngressEntries `json:"avc-ipv6-fm-ingress-entries,omitempty"` // AVC IPv6 ingress flow monitor entries
	AvcIPv6FmEgressEntries  *AvcIPv6FmEgressEntries  `json:"avc-ipv6-fm-egress-entries,omitempty"`  // AVC IPv6 egress flow monitor entries
}

// WlanSwitchingPolicy represents WLAN switching policy configuration.
type WlanSwitchingPolicy struct {
	CentralSwitching      bool `json:"central-switching,omitempty"`      // Central switching enabled (Live: IOS-XE 17.12.5)
	CentralAuthentication bool `json:"central-authentication,omitempty"` // Central authentication enabled (Live: IOS-XE 17.12.5)
	CentralDHCP           bool `json:"central-dhcp,omitempty"`           // Central dhcp for locally switched clients (Live: IOS-XE 17.12.5)
	CentralAssocEnable    bool `json:"central-assoc-enable,omitempty"`   // Central association enabled (Live: IOS-XE 17.12.5)
}

// WlanTimeout represents WLAN timeout configuration.
type WlanTimeout struct {
	SessionTimeout int `json:"session-timeout,omitempty"` // Session timeout in seconds (Live: IOS-XE 17.12.5)
}

// PerSsidQos represents per-SSID QoS configuration.
type PerSsidQos struct {
	IngressServiceName string `json:"ingress-service-name,omitempty"` // Ingress QoS service name (Live: IOS-XE 17.12.5)
	EgressServiceName  string `json:"egress-service-name,omitempty"`  // Egress QoS service name (Live: IOS-XE 17.12.5)
}

// DHCPParams represents DHCP parameters configuration.
type DHCPParams struct {
	IsDHCPEnabled bool `json:"is-dhcp-enabled,omitempty"` // Whether DHCP is enabled (Live: IOS-XE 17.12.5)
}

// UmbrellaFlexParams represents Umbrella Flex parameters configuration.
type UmbrellaFlexParams struct {
	DHCPDNSOptionEnable bool `json:"dhcp-dns-option-enable,omitempty"` // DHCP DNS option for Umbrella enabled (Live: IOS-XE 17.12.5)
}

// AtfPolicyMapEntries represents ATF (Airtime Fairness) policy map entries from live WNC 17.12.1.
type AtfPolicyMapEntries struct {
	AtfPolicyMapEntry []AtfPolicyMapEntry `json:"atf-policy-map-entry"` // List of ATF policy entries
}

// AtfPolicyMapEntry represents a single ATF policy map entry from live WNC 17.12.5.
type AtfPolicyMapEntry struct {
	BandID        int    `json:"band-id"`         // Radio band identifier (Live: IOS-XE 17.12.5)
	AtfPolicyName string `json:"atf-policy-name"` // ATF policy name (Live: IOS-XE 17.12.5)
}

// AvcIPv4FmIngressEntries represents AVC IPv4 flow monitor ingress entries from live WNC 17.12.5.
type AvcIPv4FmIngressEntries struct {
	AvcIPv4FmIngressEntry []AvcIPv4FmIngressEntry `json:"avc-ipv4-fm-ingress-entry"` // List of IPv4 ingress flow monitors (Live: IOS-XE 17.12.5)
}

// AvcIPv4FmIngressEntry represents a single AVC IPv4 flow monitor ingress entry from live WNC 17.12.5.
type AvcIPv4FmIngressEntry struct {
	Name string `json:"name"` // Flow monitor name (Live: IOS-XE 17.12.5)
}

// AvcIPv4FmEgressEntries represents AVC IPv4 flow monitor egress entries from live WNC 17.12.5.
type AvcIPv4FmEgressEntries struct {
	AvcIPv4FmEgressEntry []AvcIPv4FmEgressEntry `json:"avc-ipv4-fm-egress-entry"` // List of IPv4 egress flow monitors (Live: IOS-XE 17.12.5)
}

// AvcIPv4FmEgressEntry represents a single AVC IPv4 flow monitor egress entry from live WNC 17.12.5.
type AvcIPv4FmEgressEntry struct {
	Name string `json:"name"` // Flow monitor name (Live: IOS-XE 17.12.5)
}

// AvcIPv6FmIngressEntries represents AVC IPv6 flow monitor ingress entries from live WNC 17.12.5.
type AvcIPv6FmIngressEntries struct {
	AvcIPv6FmIngressEntry []AvcIPv6FmIngressEntry `json:"avc-ipv6-fm-ingress-entry"` // List of IPv6 ingress flow monitors (Live: IOS-XE 17.12.5)
}

// AvcIPv6FmIngressEntry represents a single AVC IPv6 flow monitor ingress entry from live WNC 17.12.5.
type AvcIPv6FmIngressEntry struct {
	Name string `json:"name"` // Flow monitor name (Live: IOS-XE 17.12.5)
}

// AvcIPv6FmEgressEntries represents AVC IPv6 flow monitor egress entries from live WNC 17.12.5.
type AvcIPv6FmEgressEntries struct {
	AvcIPv6FmEgressEntry []AvcIPv6FmEgressEntry `json:"avc-ipv6-fm-egress-entry"` // List of IPv6 egress flow monitors (Live: IOS-XE 17.12.5)
}

// AvcIPv6FmEgressEntry represents a single AVC IPv6 flow monitor egress entry from live WNC 17.12.5.
type AvcIPv6FmEgressEntry struct {
	Name string `json:"name"` // Flow monitor name (Live: IOS-XE 17.12.5)
}

// PolicyListEntry represents individual policy list entry.
type PolicyListEntry struct {
	TagName      string        `json:"tag-name,omitempty"`      // This object uniquely identifies the policy tag (Live: IOS-XE 17.12.5)
	Description  string        `json:"description,omitempty"`   // Description for the policy tag (Live: IOS-XE 17.12.5)
	WLANPolicies *WLANPolicies `json:"wlan-policies,omitempty"` // WLAN policy configuration (Live: IOS-XE 17.12.5)
}

// WLANPolicies represents the container for WLAN policy mappings.
type WLANPolicies struct {
	WLANPolicy []WLANPolicyMap `json:"wlan-policy,omitempty"` // List of WLAN to policy mappings
}

// WLANPolicyMap represents a WLAN to policy profile mapping.
type WLANPolicyMap struct {
	WLANProfileName   string `json:"wlan-profile-name"`   // Name of the WLAN profile (Live: IOS-XE 17.12.5)
	PolicyProfileName string `json:"policy-profile-name"` // Name of the policy profile (Live: IOS-XE 17.12.5)
}

// WirelessAaaPolicyConfig represents wireless AAA policy configuration.
type WirelessAaaPolicyConfig struct {
	PolicyName string `json:"policy-name"` // The wireless AAA policy name (Live: IOS-XE 17.12.5)
}

// Dot11beProfiles represents Wi-Fi 7 / 802.11be profiles (Live: IOS-XE 17.15.4b).
type Dot11beProfiles struct {
	Dot11beProfile []Dot11beProfile `json:"dot11be-profile"` // List of Wi-Fi 7 profiles (Live: IOS-XE 17.15.4b)
}

// Dot11beProfile represents a single 802.11be profile (Live: IOS-XE 17.15.4b).
type Dot11beProfile struct {
	ProfileName       string `json:"profile-name"`                  // 802.11be profile name (Live: IOS-XE 17.15.4b)
	Description       string `json:"description,omitempty"`         // 802.11be profile description (Live: IOS-XE 17.15.4b)
	EhtOfdmaDownlink  bool   `json:"eht-ofdma-downlink,omitempty"`  // 802.11be OFDMA downlink (YANG: IOS-XE 17.18.1)
	EhtOfdmaUplink    bool   `json:"eht-ofdma-uplink,omitempty"`    // 802.11be OFDMA uplink (YANG: IOS-XE 17.18.1)
	EhtMumimoDownlink bool   `json:"eht-mumimo-downlink,omitempty"` // 802.11be MU-MIMO downlink (YANG: IOS-XE 17.18.1)
	EhtMumimoUplink   bool   `json:"eht-mumimo-uplink,omitempty"`   // 802.11be MU-MIMO uplink (YANG: IOS-XE 17.18.1)
	EhtOfdmaMultiRu   bool   `json:"eht-ofdma-multi-ru,omitempty"`  // 802.11be OFDMA multiple resource unit (YANG: IOS-XE 17.18.1)

	// Multi-link Operation (MLO) configurations (YANG: IOS-XE 17.18.1)
	Mlo24Ghz   *MloGroup24Ghz   `json:"mlo-24ghz,omitempty"`    // Multi-link group config for 2.4GHz link (YANG: IOS-XE 17.18.1)
	Mlo5Ghz    *MloGroup5Ghz    `json:"mlo-5ghz,omitempty"`     // Multi-link group config for primary 5GHz (YANG: IOS-XE 17.18.1)
	Mlo5GhzSec *MloGroup5GhzSec `json:"mlo-5ghz-sec,omitempty"` // Multi-link group config for secondary 5GHz (YANG: IOS-XE 17.18.1)
	Mlo6Ghz    *MloGroup6Ghz    `json:"mlo-6ghz,omitempty"`     // Multi-link group config for 6GHz link (YANG: IOS-XE 17.18.1)
}

// MloGroup24Ghz represents Multi-link group configuration for 2.4GHz link (YANG: IOS-XE 17.18.1).
type MloGroup24Ghz struct {
	GroupID int  `json:"group-id,omitempty"` // MLO group identifier (YANG: IOS-XE 17.18.1)
	Enable  bool `json:"enable,omitempty"`   // MLO group enable status (YANG: IOS-XE 17.18.1)
}

// MloGroup5Ghz represents Multi-link group configuration for primary 5GHz link (YANG: IOS-XE 17.18.1).
type MloGroup5Ghz struct {
	GroupID int  `json:"group-id,omitempty"` // MLO group identifier (YANG: IOS-XE 17.18.1)
	Enable  bool `json:"enable,omitempty"`   // MLO group enable status (YANG: IOS-XE 17.18.1)
}

// MloGroup5GhzSec represents Multi-link group configuration for secondary 5GHz link (YANG: IOS-XE 17.18.1).
type MloGroup5GhzSec struct {
	GroupID int  `json:"group-id,omitempty"` // MLO group identifier (YANG: IOS-XE 17.18.1)
	Enable  bool `json:"enable,omitempty"`   // MLO group enable status (YANG: IOS-XE 17.18.1)
}

// MloGroup6Ghz represents Multi-link group configuration for 6GHz link (YANG: IOS-XE 17.18.1).
type MloGroup6Ghz struct {
	GroupID int  `json:"group-id,omitempty"` // MLO group identifier (YANG: IOS-XE 17.18.1)
	Enable  bool `json:"enable,omitempty"`   // MLO group enable status (YANG: IOS-XE 17.18.1)
}

// CiscoIOSXEWirelessWlanCfgWlanPolicies wraps the WlanPolicies structure of the WLAN configuration data.
type CiscoIOSXEWirelessWlanCfgWlanPolicies struct {
	CiscoIOSXEWirelessWlanCfgData struct {
		WlanPolicies *WlanPolicies `json:"wlan-policies,omitempty"`
	} `json:"Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data"`
}

// CiscoIOSXEWirelessWlanCfgDot11beProfiles wraps the Dot11beProfiles structure of the WLAN configuration data.
type CiscoIOSXEWirelessWlanCfgDot11beProfiles struct {
	CiscoIOSXEWirelessWlanCfgData struct {
		Dot11beProfiles *Dot11beProfiles `json:"dot11be-profiles,omitempty"`
	} `json:"Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data"`
}
