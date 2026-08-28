package wlan

import "log/slog"

// CiscoIOSXEWirelessWlanCfg represents the complete WLAN configuration.
type CiscoIOSXEWirelessWlanCfg struct {
	CiscoIOSXEWirelessWlanCfgData *CiscoIOSXEWirelessWlanCfgData `json:"Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data"` // WLAN configuration data container
}

// CiscoIOSXEWirelessWlanCfgData represents WLAN configuration data container.
type CiscoIOSXEWirelessWlanCfgData struct {
	WlanCfgEntries           *WlanCfgEntries           `json:"wlan-cfg-entries"`                      // WLAN configuration parameters (Live: IOS-XE 17.12.6a)
	WlanPolicies             *WlanPolicies             `json:"wlan-policies"`                         // WLAN policy configuration (Live: IOS-XE 17.12.6a)
	PolicyListEntries        *PolicyListEntries        `json:"policy-list-entries,omitempty"`         // Policy list configuration (Live: IOS-XE 17.12.6a)
	WirelessAaaPolicyConfigs *WirelessAaaPolicyConfigs `json:"wireless-aaa-policy-configs,omitempty"` // Wireless AAA policy configurations (Live: IOS-XE 17.12.6a)
	Dot11beProfiles          *Dot11beProfiles          `json:"dot11be-profiles"`                      // 802.11be profile parameters (Live: IOS-XE 17.15.4b)
}

// CiscoIOSXEWirelessWlanCfgWlanCfgEntries represents the WLAN configuration entries.
type CiscoIOSXEWirelessWlanCfgWlanCfgEntries struct {
	WlanCfgEntries *WlanCfgEntries `json:"Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-entries"`
}

// CiscoIOSXEWirelessWlanCfgPolicyListEntries represents the policy list entries.
type CiscoIOSXEWirelessWlanCfgPolicyListEntries struct {
	PolicyListEntries *PolicyListEntries `json:"Cisco-IOS-XE-wireless-wlan-cfg:policy-list-entries"`
}

// CiscoIOSXEWirelessWlanCfgPolicyListEntry represents the response to a keyed read of one policy
// list entry. A keyed list read answers with the singular node name and a list of one record,
// measured on a controller, which is why the sole field is a slice.
type CiscoIOSXEWirelessWlanCfgPolicyListEntry struct {
	PolicyListEntry []PolicyListEntry `json:"Cisco-IOS-XE-wireless-wlan-cfg:policy-list-entry"`
}

// CiscoIOSXEWirelessWlanCfgWirelessAaaPolicyConfigs represents the wireless AAA policy configurations.
type CiscoIOSXEWirelessWlanCfgWirelessAaaPolicyConfigs struct {
	WirelessAaaPolicyConfigs *WirelessAaaPolicyConfigs `json:"Cisco-IOS-XE-wireless-wlan-cfg:wireless-aaa-policy-configs"`
}

// CiscoIOSXEWirelessWlanCfgWlanPolicies wraps the WlanPolicies structure of the WLAN configuration data.
type CiscoIOSXEWirelessWlanCfgWlanPolicies struct {
	WlanPolicies *WlanPolicies `json:"Cisco-IOS-XE-wireless-wlan-cfg:wlan-policies"`
}

// CiscoIOSXEWirelessWlanCfgDot11beProfiles wraps the Dot11beProfiles structure of the WLAN configuration data.
type CiscoIOSXEWirelessWlanCfgDot11beProfiles struct {
	Dot11beProfiles *Dot11beProfiles `json:"Cisco-IOS-XE-wireless-wlan-cfg:dot11be-profiles"`
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

// Secret holds key material. Its String and LogValue redact, so a formatted struct and a slog
// record carry no key; call Reveal where the value is needed.
//
// MarshalJSON is deliberately not implemented: encoding/json is how this type reaches a write
// payload. So json.Marshal, encoding/json/v2 and the %#v verb still render the key, and they are
// the three paths a caller has to keep out of a log.
type Secret string

// String returns the redaction, which is what keeps the key out of every fmt verb but %#v.
func (s Secret) String() string { return redacted }

// LogValue returns the redaction. String is not enough on its own: slog.JSONHandler renders a
// KindAny value through json.Marshal, which honors json.Marshaler and ignores fmt.Stringer.
func (s Secret) LogValue() slog.Value { return slog.StringValue(redacted) }

// Reveal returns the key itself, for the one caller that needs it.
func (s Secret) Reveal() string { return string(s) }

// redacted is what String and LogValue render in place of key material.
const redacted = "[REDACTED]"

// WlanCfgEntry represents a single WLAN configuration entry.
//
// A security leaf the controller omits holds its schema default, so the two Go shapes below are
// not interchangeable: a leaf whose default is true is a pointer and says so in its own comment,
// because decoding that absence as false inverts the reading, while a default-false leaf is a
// value bool because false is what its absence means. An enumeration leaf is a plain string, so an
// omitted one holds the default spelling its schema declares rather than "" — dot11-auth-type names
// its own in the comment below. The four leaves marked 17.15.1 are not declared at all on 17.12,
// where their absence is the release rather than a reading — ask the controller which modules it
// serves if that distinction matters.
//
// LogValue reduces the entry to the two leaves that identify it, because slog resolves LogValuer
// on the Attr's own value and never on a field inside it: without it, slog.JSONHandler renders the
// whole entry through json.Marshal and the key goes to the log.
type WlanCfgEntry struct {
	WlanID                 int                `json:"wlan-id"`                              // WLAN identifier (Live: IOS-XE 17.12.6a)
	ProfileName            string             `json:"profile-name"`                         // WLAN profile name (Live: IOS-XE 17.12.6a)
	AuthKeyMgmtPsk         bool               `json:"auth-key-mgmt-psk,omitempty"`          // Authentication key management PSK (Live: IOS-XE 17.12.6a)
	AuthKeyMgmtPskSha256   bool               `json:"auth-key-mgmt-psk-sha256,omitempty"`   // Authentication key management PSK SHA256 (YANG: IOS-XE 17.12.1)
	AuthKeyMgmtDot1x       *bool              `json:"auth-key-mgmt-dot1x,omitempty"`        // Authentication key management type 802.1x; schema default true, so absence is not false (Live: IOS-XE 17.12.6a)
	AuthKeyMgmtDot1xSha256 bool               `json:"auth-key-mgmt-dot1x-sha256,omitempty"` // Authentication key management type 802.1x SHA256 (Live: IOS-XE 17.12.6a)
	AuthKeyMgmtCckm        bool               `json:"auth-key-mgmt-cckm,omitempty"`         // Authentication key management CCKM; the schema marks the leaf deprecated on 17.12 and 17.15 and not on 17.18 (YANG: IOS-XE 17.12.1)
	AuthKeyMgmtSae         bool               `json:"auth-key-mgmt-sae,omitempty"`          // Authentication key management SAE (YANG: IOS-XE 17.12.1)
	AuthKeyMgmtFtPsk       bool               `json:"auth-key-mgmt-ft-psk,omitempty"`       // Authentication key management fast-transition PSK (YANG: IOS-XE 17.12.1)
	AuthKeyMgmtFtDot1x     bool               `json:"auth-key-mgmt-ft-dot1x,omitempty"`     // Authentication key management fast-transition 802.1x (YANG: IOS-XE 17.12.1)
	AuthKeyMgmtFtSae       bool               `json:"auth-key-mgmt-ft-sae,omitempty"`       // Authentication key management fast-transition SAE (YANG: IOS-XE 17.12.1)
	AuthKeyMgmtSuiteB      bool               `json:"auth-key-mgmt-suite-b,omitempty"`      // Authentication key management Suite-B (YANG: IOS-XE 17.15.1)
	AuthKeyMgmtSuiteB192   bool               `json:"auth-key-mgmt-suite-b-192,omitempty"`  // Authentication key management Suite-B 192-bit (YANG: IOS-XE 17.15.1)
	AkmOwe                 bool               `json:"akm-owe,omitempty"`                    // Authentication key management Opportunistic Wireless Encryption (YANG: IOS-XE 17.12.1)
	AkmSaeExtKey           bool               `json:"akm-sae-ext-key,omitempty"`            // Authentication key management SAE with extended key (YANG: IOS-XE 17.15.1)
	AkmFtSaeExtKey         bool               `json:"akm-ft-sae-ext-key,omitempty"`         // Authentication key management fast-transition SAE with extended key (YANG: IOS-XE 17.15.1)
	PSK                    Secret             `json:"psk,omitempty"`                        // Authentication pre-shared key — secret, never log (Live: IOS-XE 17.12.6a)
	PSKType                string             `json:"psk-type,omitempty"`                   // Pre-shared key encryption type (Live: IOS-XE 17.12.6a)
	FTMode                 string             `json:"ft-mode,omitempty"`                    // Configures Fast Transition Adaptive support (Live: IOS-XE 17.12.6a)
	PMFOptions             string             `json:"pmf-options,omitempty"`                // Configures PMF as optional/required (Live: IOS-XE 17.12.6a)
	SecurityWPA            *bool              `json:"security-wpa,omitempty"`               // WPA security enabled on the WLAN; schema default true, so absence is not false (YANG: IOS-XE 17.12.1)
	WPA1Enabled            bool               `json:"wpa1-enabled,omitempty"`               // Configures WPA1 support (YANG: IOS-XE 17.12.1)
	WPA2Enabled            *bool              `json:"wpa2-enabled,omitempty"`               // Configures WPA2 support; schema default true, so absence is not false (Live: IOS-XE 17.12.6a)
	WPA2AES                *bool              `json:"wpa2-aes,omitempty"`                   // WPA2 AES encryption; schema default true, so absence is not false (YANG: IOS-XE 17.12.1)
	WPA3Enabled            bool               `json:"wpa3-enabled,omitempty"`               // Configures WPA3 support (Live: IOS-XE 17.12.6a)
	WEPEnabled             bool               `json:"wep-enabled,omitempty"`                // Static WEP enabled on the WLAN (YANG: IOS-XE 17.12.1)
	OSEN                   bool               `json:"osen,omitempty"`                       // Hotspot 2.0 OSEN enabled (YANG: IOS-XE 17.12.1)
	OKC                    *bool              `json:"okc,omitempty"`                        // Opportunistic key caching; schema default true, so absence is not false (YANG: IOS-XE 17.12.1)
	Dot11AuthType          string             `json:"dot11-auth-type,omitempty"`            // 802.11 authentication type; an omitted leaf means the default spelling apf-vap-80211-auth-open, not "" (YANG: IOS-XE 17.12.1)
	LoadBalance            bool               `json:"load-balance,omitempty"`               // Allow/Disallow Load Balance on a WLAN (Live: IOS-XE 17.12.6a)
	AuthenticationList     string             `json:"authentication-list,omitempty"`        // Enter the Authentication list name (Live: IOS-XE 17.12.6a)
	Wlan11kNeighList       *bool              `json:"wlan-11k-neigh-list,omitempty"`        // Indicates 11k Neighbor List enabled; schema default true, so absence is not false (YANG: IOS-XE 17.12.1)
	MulticastBufferValue   int                `json:"multicast-buffer-value,omitempty"`     // Configure Multicast Buffer Tuning (YANG: IOS-XE 17.12.1)
	APFVapIDData           *APFVapIDData      `json:"apf-vap-id-data,omitempty"`            // Virtual AP interface data (Live: IOS-XE 17.12.6a)
	APFVap80211vData       *APFVap80211vData  `json:"apf-vap-802-11v-data,omitempty"`       // 802.11v wireless management configuration (Live: IOS-XE 17.12.6a)
	MDNSSDMode             string             `json:"mdns-sd-mode,omitempty"`               // MDNS operational mode on WLAN (Live: IOS-XE 17.12.6a)
	WlanRadioPolicies      *WlanRadioPolicies `json:"wlan-radio-policies,omitempty"`        // WLAN radio policy (Live: IOS-XE 17.12.6a)
	ClientSteering         bool               `json:"client-steering,omitempty"`            // Enable/disable 6Ghz client steering (YANG: IOS-XE 17.12.1)
	WepKeyIndex            int                `json:"wep-key-index,omitempty"`              // WEP key index for Static WEP Authentication (Live: IOS-XE 17.12.6a)
}

// LogValue returns the leaves that identify this WLAN. The entry is not rendered whole: a slog
// record built from it would otherwise carry every configuration leaf, the pre-shared key
// included, whenever the handler is a JSONHandler.
func (e WlanCfgEntry) LogValue() slog.Value {
	return slog.GroupValue(
		slog.Int("wlan-id", e.WlanID),
		slog.String("profile-name", e.ProfileName),
	)
}

// APFVapIDData represents virtual AP interface identification data.
type APFVapIDData struct {
	SSID           string `json:"ssid"`                       // Service Set Identifier (Live: IOS-XE 17.12.6a)
	WlanStatus     *bool  `json:"wlan-status,omitempty"`      // WLAN administrative status (Live: IOS-XE 17.12.6a)
	BroadcastSSID  *bool  `json:"broadcast-ssid,omitempty"`   // SSID broadcast in the beacon; schema default true, so absence is not false (YANG: IOS-XE 17.12.1)
	P2PBlockAction string `json:"p2p-block-action,omitempty"` // Peer-to-peer blocking action; an omitted leaf means the default spelling p2p-blocking-action-none, not "" (YANG: IOS-XE 17.12.1)
}

// APFVap80211vData represents 802.11v wireless management configuration.
type APFVap80211vData struct {
	Dot11vDms *bool `json:"dot11v-dms,omitempty"` // 802.11v Directed Multicast Service enabled; schema default true, so absence is not false (Live: IOS-XE 17.12.6a)
}

// WlanRadioPolicies represents WLAN radio band policy configuration.
type WlanRadioPolicies struct {
	WlanRadioPolicy []WlanRadioPolicy `json:"wlan-radio-policy"` // List of radio band policies
}

// WlanRadioPolicy represents individual radio band policy configuration.
type WlanRadioPolicy struct {
	Band string `json:"band"` // Radio band specification (Live: IOS-XE 17.12.6a)
}

// WlanPolicy represents WLAN policy profile configuration.
type WlanPolicy struct {
	PolicyProfileName       string                   `json:"policy-profile-name"`                   // This object specifies policy profile instance (Live: IOS-XE 17.12.6a)
	Description             string                   `json:"description,omitempty"`                 // Description associated to WLAN policy (Live: IOS-XE 17.12.6a)
	Status                  bool                     `json:"status,omitempty"`                      // Whether policy profile is shutdown or active (Live: IOS-XE 17.12.6a)
	InterfaceName           string                   `json:"interface-name,omitempty"`              // Interface attached to the wireless lan (Live: IOS-XE 17.12.6a)
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
	CentralSwitching      *bool `json:"central-switching,omitempty"`      // Central switching enabled; schema default true, so absence is not false (Live: IOS-XE 17.12.6a)
	CentralAuthentication *bool `json:"central-authentication,omitempty"` // Central authentication enabled; schema default true, so absence is not false (Live: IOS-XE 17.12.6a)
	CentralDHCP           *bool `json:"central-dhcp,omitempty"`           // Central dhcp for locally switched clients; schema default true, so absence is not false (Live: IOS-XE 17.12.6a)
	CentralAssocEnable    *bool `json:"central-assoc-enable,omitempty"`   // Central association enabled; schema default true, so absence is not false (Live: IOS-XE 17.12.6a)
}

// WlanTimeout represents WLAN timeout configuration.
type WlanTimeout struct {
	SessionTimeout *int `json:"session-timeout,omitempty"` // Session timeout in seconds (Live: IOS-XE 17.12.6a)
	IdleTimeout    *int `json:"idle-timeout,omitempty"`    // Client idle timeout (Live: IOS-XE 17.12.7a)
	IdleThreshold  *int `json:"idle-threshold,omitempty"`  // Client idle traffic threshold (Live: IOS-XE 17.12.7a)
}

// PerSsidQos represents per-SSID QoS configuration.
type PerSsidQos struct {
	IngressServiceName string `json:"ingress-service-name,omitempty"` // Ingress QoS service name (Live: IOS-XE 17.12.6a)
	EgressServiceName  string `json:"egress-service-name,omitempty"`  // Egress QoS service name (Live: IOS-XE 17.12.6a)
}

// DHCPParams represents DHCP parameters configuration.
type DHCPParams struct {
	IsDHCPEnabled bool `json:"is-dhcp-enabled,omitempty"` // Whether DHCP is enabled (Live: IOS-XE 17.12.6a)
}

// UmbrellaFlexParams represents Umbrella Flex parameters configuration.
type UmbrellaFlexParams struct {
	DHCPDNSOptionEnable *bool `json:"dhcp-dns-option-enable,omitempty"` // DHCP DNS option for Umbrella enabled; schema default true, so absence is not false (Live: IOS-XE 17.12.6a)
	ModeForce           bool  `json:"mode-force,omitempty"`             // Force Umbrella mode on the profile (Live: IOS-XE 17.12.7a)
}

// AtfPolicyMapEntries represents ATF (Airtime Fairness) policy map entries from live WNC 17.12.1.
type AtfPolicyMapEntries struct {
	AtfPolicyMapEntry []AtfPolicyMapEntry `json:"atf-policy-map-entry"` // List of ATF policy entries
}

// AtfPolicyMapEntry represents a single ATF policy map entry.
type AtfPolicyMapEntry struct {
	BandID        int    `json:"band-id"`         // Radio band identifier (Live: IOS-XE 17.12.6a)
	AtfPolicyName string `json:"atf-policy-name"` // ATF policy name (Live: IOS-XE 17.12.6a)
}

// AvcIPv4FmIngressEntries represents AVC IPv4 flow monitor ingress entries.
type AvcIPv4FmIngressEntries struct {
	AvcIPv4FmIngressEntry []AvcIPv4FmIngressEntry `json:"avc-ipv4-fm-ingress-entry"` // List of IPv4 ingress flow monitors (Live: IOS-XE 17.12.6a)
}

// AvcIPv4FmIngressEntry represents a single AVC IPv4 flow monitor ingress entry.
type AvcIPv4FmIngressEntry struct {
	Name string `json:"name"` // Flow monitor name (Live: IOS-XE 17.12.6a)
}

// AvcIPv4FmEgressEntries represents AVC IPv4 flow monitor egress entries.
type AvcIPv4FmEgressEntries struct {
	AvcIPv4FmEgressEntry []AvcIPv4FmEgressEntry `json:"avc-ipv4-fm-egress-entry"` // List of IPv4 egress flow monitors (Live: IOS-XE 17.12.6a)
}

// AvcIPv4FmEgressEntry represents a single AVC IPv4 flow monitor egress entry.
type AvcIPv4FmEgressEntry struct {
	Name string `json:"name"` // Flow monitor name (Live: IOS-XE 17.12.6a)
}

// AvcIPv6FmIngressEntries represents AVC IPv6 flow monitor ingress entries.
type AvcIPv6FmIngressEntries struct {
	AvcIPv6FmIngressEntry []AvcIPv6FmIngressEntry `json:"avc-ipv6-fm-ingress-entry"` // List of IPv6 ingress flow monitors (Live: IOS-XE 17.12.6a)
}

// AvcIPv6FmIngressEntry represents a single AVC IPv6 flow monitor ingress entry.
type AvcIPv6FmIngressEntry struct {
	Name string `json:"name"` // Flow monitor name (Live: IOS-XE 17.12.6a)
}

// AvcIPv6FmEgressEntries represents AVC IPv6 flow monitor egress entries.
type AvcIPv6FmEgressEntries struct {
	AvcIPv6FmEgressEntry []AvcIPv6FmEgressEntry `json:"avc-ipv6-fm-egress-entry"` // List of IPv6 egress flow monitors (Live: IOS-XE 17.12.6a)
}

// AvcIPv6FmEgressEntry represents a single AVC IPv6 flow monitor egress entry.
type AvcIPv6FmEgressEntry struct {
	Name string `json:"name"` // Flow monitor name (Live: IOS-XE 17.12.6a)
}

// PolicyListEntry represents individual policy list entry.
type PolicyListEntry struct {
	TagName      string        `json:"tag-name,omitempty"`      // This object uniquely identifies the policy tag (Live: IOS-XE 17.12.6a)
	Description  *string       `json:"description,omitempty"`   // Description for the policy tag (Live: IOS-XE 17.12.6a)
	WLANPolicies *WLANPolicies `json:"wlan-policies,omitempty"` // WLAN policy configuration (Live: IOS-XE 17.12.6a)
}

// WLANPolicies represents the container for WLAN policy mappings.
type WLANPolicies struct {
	WLANPolicy []WLANPolicyMap `json:"wlan-policy,omitempty"` // List of WLAN to policy mappings
}

// WLANPolicyMap represents a WLAN to policy profile mapping.
type WLANPolicyMap struct {
	WLANProfileName   string `json:"wlan-profile-name"`   // Name of the WLAN profile (Live: IOS-XE 17.12.6a)
	PolicyProfileName string `json:"policy-profile-name"` // Name of the policy profile (Live: IOS-XE 17.12.6a)
}

// WirelessAaaPolicyConfig represents wireless AAA policy configuration.
type WirelessAaaPolicyConfig struct {
	PolicyName string `json:"policy-name"` // The wireless AAA policy name (Live: IOS-XE 17.12.6a)
}

// Dot11beProfiles represents Wi-Fi 7 / 802.11be profiles (Live: IOS-XE 17.15.4b).
type Dot11beProfiles struct {
	Dot11beProfile []Dot11beProfile `json:"dot11be-profile"` // List of Wi-Fi 7 profiles (Live: IOS-XE 17.15.4b)
}

// Dot11beProfile represents a single 802.11be profile (Live: IOS-XE 17.15.4b).
type Dot11beProfile struct {
	ProfileName       string `json:"profile-name"`                  // 802.11be profile name (Live: IOS-XE 17.15.4b)
	Description       string `json:"description,omitempty"`         // 802.11be profile description (Live: IOS-XE 17.15.4b)
	EhtOfdmaDownlink  *bool  `json:"eht-ofdma-downlink,omitempty"`  // 802.11be OFDMA downlink; schema default true, so absence is not false (YANG: IOS-XE 17.18.1)
	EhtOfdmaUplink    *bool  `json:"eht-ofdma-uplink,omitempty"`    // 802.11be OFDMA uplink; schema default true, so absence is not false (YANG: IOS-XE 17.18.1)
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
//
// The container holds one leaf of the same name, carrying the enumeration spelling
// "24ghz-enable" or "24ghz-disable". The group identifier and enable flag this type
// declared before are not nodes the controller sends.
type MloGroup24Ghz struct {
	Mlo24Ghz *string `json:"mlo-24ghz,omitempty"` // 802.11be MLO state for the 2.4GHz link (YANG: IOS-XE 17.18.1)
}

// MloGroup5Ghz represents Multi-link group configuration for primary 5GHz link (YANG: IOS-XE 17.18.1).
type MloGroup5Ghz struct {
	Mlo5Ghz *string `json:"mlo-5ghz,omitempty"` // 802.11be MLO state for the primary 5GHz link (YANG: IOS-XE 17.18.1)
}

// MloGroup5GhzSec represents Multi-link group configuration for secondary 5GHz link (YANG: IOS-XE 17.18.1).
type MloGroup5GhzSec struct {
	Mlo5GhzSec *string `json:"mlo-5ghz-sec,omitempty"` // 802.11be MLO state for the secondary 5GHz link (YANG: IOS-XE 17.18.1)
}

// MloGroup6Ghz represents Multi-link group configuration for 6GHz link (YANG: IOS-XE 17.18.1).
type MloGroup6Ghz struct {
	Mlo6Ghz *string `json:"mlo-6ghz,omitempty"` // 802.11be MLO state for the 6GHz link (YANG: IOS-XE 17.18.1)
}
