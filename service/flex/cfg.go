package flex

// CiscoIOSXEWirelessFlexCfg represents FlexConnect configuration data container.
type CiscoIOSXEWirelessFlexCfg struct {
	CiscoIOSXEWirelessFlexCfgData struct {
		FlexPolicyEntries FlexPolicyEntries `json:"flex-policy-entries"` // FlexConnect policy entries configuration (Live: IOS-XE 17.12.6a)
	} `json:"Cisco-IOS-XE-wireless-flex-cfg:flex-cfg-data"` // FlexConnect configuration data (Live: IOS-XE 17.12.6a)
}

// CiscoIOSXEWirelessFlexCfgFlexPolicyEntries represents FlexConnect policy entries container.
type CiscoIOSXEWirelessFlexCfgFlexPolicyEntries struct {
	FlexPolicyEntries FlexPolicyEntries `json:"Cisco-IOS-XE-wireless-flex-cfg:flex-policy-entries"`
}

// FlexPolicyEntries represents FlexConnect policy entries collection.
type FlexPolicyEntries struct {
	FlexPolicyEntry []FlexPolicyEntry `json:"flex-policy-entry"` // List of FlexConnect policy configurations (Live: IOS-XE 17.12.6a)
}

// FlexPolicyEntry represents individual FlexConnect policy configuration.
type FlexPolicyEntry struct {
	PolicyName                string               `json:"policy-name"`                             // Name of the flex profile (Live: IOS-XE 17.12.6a)
	Description               string               `json:"description,omitempty"`                   // Description for the flex profile (Live: IOS-XE 17.12.6a)
	IfNameVlanIDs             *FlexIfNameVlanIDs   `json:"if-name-vlan-ids,omitempty"`              // Interface name VLAN IDs container (Live: IOS-XE 17.12.6a)
	EAPFastProfileName        string               `json:"eap-fast-profile-name,omitempty"`         // EAP fast profile for local authentication (YANG: IOS-XE 17.12.1)
	RADIUSServerGroupName     string               `json:"radius-server-group-name,omitempty"`      // Radius server group for authentication (YANG: IOS-XE 17.12.1)
	FallbackRadioShut         *bool                `json:"fallback-radio-shut,omitempty"`           // Fallback Radio Shut feature enable flag (YANG: IOS-XE 17.12.1)
	ARPCaching                *bool                `json:"arp-caching,omitempty"`                   // ARP cache feature enable flag (YANG: IOS-XE 17.12.1)
	CTSInlineTagging          *bool                `json:"cts-inline-tagging,omitempty"`            // CTS inline tagging feature enable flag (YANG: IOS-XE 17.12.1)
	CTSRolebasedEnforce       *bool                `json:"cts-rolebased-enforce,omitempty"`         // CTS rolebased enforcement enable flag (YANG: IOS-XE 17.12.1)
	CTSProfileName            string               `json:"cts-profile-name,omitempty"`              // CTS SXP profile name (YANG: IOS-XE 17.12.1)
	JoinMinLatency            *bool                `json:"join-min-latency,omitempty"`              // AP joins controller with smallest latency (YANG: IOS-XE 17.12.1)
	RADIUSEnable              *bool                `json:"radius-enable,omitempty"`                 // Enable or disable RADIUS (YANG: IOS-XE 17.12.1)
	VlanEnable                *bool                `json:"vlan-enable,omitempty"`                   // Availability of Native VLAN on REAP (YANG: IOS-XE 17.12.1)
	IsHomeAPEnable            *bool                `json:"is-home-ap-enable,omitempty"`             // APs connected to profile are Home APs (YANG: IOS-XE 17.12.1)
	IsRadioBackhaul           *bool                `json:"is-radio-backhaul,omitempty"`             // Enable/disable WLAN on backhaul radio (YANG: IOS-XE 17.12.1)
	IsResilientMode           *bool                `json:"is-resilient-mode,omitempty"`             // Enable/disable standalone mode on REAP AP (YANG: IOS-XE 17.12.1)
	EfficientAPUpgradeEnable  *bool                `json:"efficient-ap-upgrade-enable,omitempty"`   // Efficient AP image upgrade enable flag (YANG: IOS-XE 17.12.1)
	HTTPProxyIP               string               `json:"http-proxy-ip,omitempty"`                 // HTTP proxy IP address (YANG: IOS-XE 17.12.1)
	HTTPProxyPort             int                  `json:"http-proxy-port,omitempty"`               // HTTP proxy port (YANG: IOS-XE 17.12.1)
	NativeVlanID              int                  `json:"native-vlan-id,omitempty"`                // Native VLAN ID for particular AP (YANG: IOS-XE 17.12.1)
	SlaveMaxRetryCount        int                  `json:"slave-max-retry-count,omitempty"`         // Max retries for slave download from master (YANG: IOS-XE 17.12.1)
	AcctRADIUSServerGroupName string               `json:"acct-radius-server-group-name,omitempty"` // Radius server group for accounting (YANG: IOS-XE 17.12.1)
	IsLocalRoamingEnable      *bool                `json:"is-local-roaming-enable,omitempty"`       // Distributed client data caching for roaming (YANG: IOS-XE 17.12.1)
	PolicyACLs                []FlexPolicyACL      `json:"policy-acls,omitempty"`                   // Policy ACL configurations (YANG: IOS-XE 17.12.1)
	VlanACLs                  []FlexVlanACL        `json:"vlan-acls,omitempty"`                     // VLAN ACL mappings (obsolete) (YANG: IOS-XE 17.12.1)
	LocalAuthUsers            []FlexLocalAuthUser  `json:"local-auth-users,omitempty"`              // Local authenticated user configurations (YANG: IOS-XE 17.12.1)
	UmbrellaProfiles          FlexUmbrellaProfiles `json:"umbrella-profiles,omitempty"`             // Umbrella profile configurations (YANG: IOS-XE 17.12.1)
	MDNSProfileName           string               `json:"mdns-profile-name,omitempty"`             // mDNS flex profile name (YANG: IOS-XE 17.12.1)
	IPOverlapCfg              FlexIPOverlapConfig  `json:"ip-overlap-cfg,omitempty"`                // IP overlap configuration (YANG: IOS-XE 17.12.1)
	DHCPBcast                 bool                 `json:"dhcp-bcast,omitempty"`                    // DHCP broadcast for locally switched clients (YANG: IOS-XE 17.12.1)
	PmkDistMethod             string               `json:"pmk-dist-method,omitempty"`               // PMK distribution with APs (YANG: IOS-XE 17.12.1)
}

// FlexPolicyACL represents FlexConnect policy ACL configuration.
type FlexPolicyACL struct {
	ACLName       string `json:"acl-name"`                     // Webpolicy ACL name to map to REAP group (YANG: IOS-XE 17.12.1)
	IsCWA         *bool  `json:"is-cwa,omitempty"`             // Enable/disable central webauth for this ACL (YANG: IOS-XE 17.12.1)
	URLFilterName string `json:"urlfilterlist-name,omitempty"` // IPv4/IPv6 ACL name to url filter list mapping (YANG: IOS-XE 17.12.1)
}

// FlexVlanACL represents FlexConnect VLAN ACL mapping configuration (obsolete).
type FlexVlanACL struct {
	VlanID     int    `json:"vlan-id"`                    // VLAN ID to map to ACL for REAP group (YANG: IOS-XE 17.12.1)
	IngressACL string `json:"ingress-acl-name,omitempty"` // Ingress Access Control List name (YANG: IOS-XE 17.12.1)
	EgressACL  string `json:"egress-acl-name,omitempty"`  // Egress Access Control List name (YANG: IOS-XE 17.12.1)
}

// FlexLocalAuthUser represents FlexConnect local authenticated user configuration.
type FlexLocalAuthUser struct {
	UserName     string `json:"user-name"`     // Username for client authentication (YANG: IOS-XE 17.12.1)
	Password     string `json:"password"`      // Password for given username (YANG: IOS-XE 17.12.1)
	PasswordType string `json:"password-type"` // Password encryption type for AP auth (YANG: IOS-XE 17.12.1)
}

// FlexUmbrellaProfiles represents FlexConnect Umbrella profiles container.
type FlexUmbrellaProfiles struct {
	UmbrellaProfile []FlexUmbrellaProfile `json:"umbrella-profile"` // List of umbrella profiles (YANG: IOS-XE 17.12.1)
}

// FlexUmbrellaProfile represents individual FlexConnect Umbrella profile configuration.
type FlexUmbrellaProfile struct {
	UmbrellaName string `json:"umbrella-name"` // Umbrella profile name (YANG: IOS-XE 17.12.1)
}

// FlexIPOverlapConfig represents FlexConnect IP overlap configuration.
type FlexIPOverlapConfig struct {
	Enabled bool `json:"enabled,omitempty"` // Enable IP overlap support for site (YANG: IOS-XE 17.12.1)
}

// FlexIfNameVlanIDs represents FlexConnect interface name VLAN IDs container.
type FlexIfNameVlanIDs struct {
	IfNameVlanID []FlexIfNameVlanID `json:"if-name-vlan-id"` // Interface name VLAN ID list (Live: IOS-XE 17.12.6a)
}

// FlexIfNameVlanID represents individual FlexConnect interface name VLAN ID mapping.
type FlexIfNameVlanID struct {
	InterfaceName string `json:"interface-name"` // VLAN name for vlan-acl mapping (Live: IOS-XE 17.12.6a)
	VlanID        int    `json:"vlan-id"`        // VLAN ID mapped to ACL for AP by VLAN name (Live: IOS-XE 17.12.6a)
}
