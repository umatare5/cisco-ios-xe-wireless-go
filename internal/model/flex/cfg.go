// Package flex provides data structures for Cisco IOS-XE wireless FlexConnect
package flex

// FlexCfg represents FlexConnect configuration data container.
type FlexCfg struct {
	FlexCfgData FlexCfgData `json:"Cisco-IOS-XE-wireless-flex-cfg:flex-cfg-data"`
}

// FlexCfgFlexPolicyEntries represents FlexConnect policy entries container.
type FlexCfgFlexPolicyEntries struct {
	FlexPolicyEntries FlexPolicyEntries `json:"Cisco-IOS-XE-wireless-flex-cfg:flex-policy-entries"`
}

// FlexCfgData represents FlexConnect configuration data container.
type FlexCfgData struct {
	FlexPolicyEntries FlexPolicyEntries `json:"flex-policy-entries"`
}

// FlexPolicyEntries represents FlexConnect policy entries collection.
type FlexPolicyEntries struct {
	FlexPolicyEntry []FlexPolicyEntry `json:"flex-policy-entry"`
}

// FlexPolicyEntry represents individual FlexConnect policy configuration.
type FlexPolicyEntry struct {
	PolicyName               string               `json:"policy-name"`                           // FlexConnect policy identifier
	Description              string               `json:"description,omitempty"`                 // Policy description
	IfNameVlanIDs            *FlexIfNameVlanIDs   `json:"if-name-vlan-ids,omitempty"`            // Interface name VLAN ID mappings
	EAPFastProfileName       string               `json:"eap-fast-profile-name,omitempty"`       // EAP-FAST profile name (YANG: IOS-XE 17.12.1+)
	RadiusServerGroupName    string               `json:"radius-server-group-name,omitempty"`    // RADIUS server group name (YANG: IOS-XE 17.12.1+)
	FallbackRadioShut        *bool                `json:"fallback-radio-shut,omitempty"`         // Fallback radio shutdown setting (YANG: IOS-XE 17.12.1+)
	ARPCaching               *bool                `json:"arp-caching,omitempty"`                 // ARP caching enable flag (YANG: IOS-XE 17.12.1+)
	CTSInlineTagging         *bool                `json:"cts-inline-tagging,omitempty"`          // CTS inline tagging setting (YANG: IOS-XE 17.12.1+)
	CTSRolebasedEnforce      *bool                `json:"cts-rolebased-enforce,omitempty"`       // CTS role-based enforcement setting (YANG: IOS-XE 17.12.1+)
	CTSProfileName           string               `json:"cts-profile-name,omitempty"`            // CTS profile name (YANG: IOS-XE 17.12.1+)
	JoinMinLatency           *bool                `json:"join-min-latency,omitempty"`            // Join minimum latency setting (YANG: IOS-XE 17.12.1+)
	RadiusEnable             *bool                `json:"radius-enable,omitempty"`               // RADIUS enable flag (YANG: IOS-XE 17.12.1+)
	VlanEnable               *bool                `json:"vlan-enable,omitempty"`                 // VLAN enable flag (YANG: IOS-XE 17.12.1+)
	IsHomeAPEnable           *bool                `json:"is-home-ap-enable,omitempty"`           // Home AP enable flag (YANG: IOS-XE 17.12.1+)
	IsRadioBackhaul          *bool                `json:"is-radio-backhaul,omitempty"`           // Radio backhaul enable flag (YANG: IOS-XE 17.12.1+)
	IsResilientMode          *bool                `json:"is-resilient-mode,omitempty"`           // Resilient mode enable flag (YANG: IOS-XE 17.12.1+)
	EfficientAPUpgradeEnable *bool                `json:"efficient-ap-upgrade-enable,omitempty"` // Efficient AP upgrade enable flag (YANG: IOS-XE 17.12.1+)
	HTTPProxyIP              string               `json:"http-proxy-ip,omitempty"`               // HTTP proxy IP address (YANG: IOS-XE 17.12.1+)
	NativeVlanID             int                  `json:"native-vlan-id,omitempty"`              // Native VLAN ID (YANG: IOS-XE 17.12.1+)
	DefaultAP                bool                 `json:"default-ap,omitempty"`                  // Default AP policy flag (YANG: IOS-XE 17.12.1+)
	ACLMapMode               string               `json:"acl-map-mode,omitempty"`                // ACL mapping mode (YANG: IOS-XE 17.12.1+)
	WebPolicyACLName         string               `json:"web-policy-acl-name,omitempty"`         // Web policy ACL name (YANG: IOS-XE 17.12.1+)
	CentralAssociation       bool                 `json:"central-association,omitempty"`         // Central association enable flag (YANG: IOS-XE 17.12.1+)
	CentralAuthentication    bool                 `json:"central-authentication,omitempty"`      // Central authentication enable flag (YANG: IOS-XE 17.12.1+)
	CentralDhcp              bool                 `json:"central-dhcp,omitempty"`                // Central DHCP enable flag (YANG: IOS-XE 17.12.1+)
	CentralSwitching         bool                 `json:"central-switching,omitempty"`           // Central switching enable flag (YANG: IOS-XE 17.12.1+)
	PolicyACLs               []FlexPolicyACL      `json:"policy-acls,omitempty"`                 // Policy ACL configurations (YANG: IOS-XE 17.12.1+)
	VlanACLs                 []FlexVlanACL        `json:"vlan-acls,omitempty"`                   // VLAN ACL mappings (obsolete) (YANG: IOS-XE 17.12.1+)
	LocalAuthUsers           []FlexLocalAuthUser  `json:"local-auth-users,omitempty"`            // Local authenticated user configurations (YANG: IOS-XE 17.12.1+)
	PolicyACLName            string               `json:"policy-acl-name,omitempty"`             // Policy ACL name (YANG: IOS-XE 17.12.1+)
	UmbrellaProfiles         FlexUmbrellaProfiles `json:"umbrella-profiles,omitempty"`           // Umbrella profile configurations (YANG: IOS-XE 17.12.1+)
	MdnsProfileName          string               `json:"mdns-profile-name,omitempty"`           // mDNS profile name (YANG: IOS-XE 17.12.1+)
	IPOverlapCfg             FlexIPOverlapConfig  `json:"ip-overlap-cfg,omitempty"`              // IP overlap configuration (YANG: IOS-XE 17.12.1+)
	DhcpBcast                bool                 `json:"dhcp-bcast,omitempty"`                  // DHCP broadcast enable flag (YANG: IOS-XE 17.12.1+)
	PmkDistMethod            string               `json:"pmk-dist-method,omitempty"`             // PMK distribution method (YANG: IOS-XE 17.12.1+)
}

// FlexPolicyACL represents FlexConnect policy ACL configuration.
type FlexPolicyACL struct {
	ACLName       string `json:"acl-name"`                     // Policy ACL name (YANG: IOS-XE 17.12.1+)
	IsCWA         *bool  `json:"is-cwa,omitempty"`             // Central webauth enable flag (YANG: IOS-XE 17.12.1+)
	URLFilterName string `json:"urlfilterlist-name,omitempty"` // URL filter list name (YANG: IOS-XE 17.12.1+)
}

// FlexVlanACL represents FlexConnect VLAN ACL mapping configuration (obsolete).
type FlexVlanACL struct {
	VlanID     int    `json:"vlan-id"`                    // VLAN identifier (YANG: IOS-XE 17.12.1+)
	IngressACL string `json:"ingress-acl-name,omitempty"` // Ingress ACL name (YANG: IOS-XE 17.12.1+)
	EgressACL  string `json:"egress-acl-name,omitempty"`  // Egress ACL name (YANG: IOS-XE 17.12.1+)
}

// FlexLocalAuthUser represents FlexConnect local authenticated user configuration.
type FlexLocalAuthUser struct {
	UserName     string `json:"user-name"`     // Local username (YANG: IOS-XE 17.12.1+)
	Password     string `json:"password"`      // User password (YANG: IOS-XE 17.12.1+)
	PasswordType string `json:"password-type"` // Password type (YANG: IOS-XE 17.12.1+)
}

// FlexUmbrellaProfiles represents FlexConnect Umbrella profiles container.
type FlexUmbrellaProfiles struct {
	UmbrellaProfile []FlexUmbrellaProfile `json:"umbrella-profile"` // Umbrella profile entries (YANG: IOS-XE 17.12.1+)
}

// FlexUmbrellaProfile represents individual FlexConnect Umbrella profile configuration.
type FlexUmbrellaProfile struct {
	UmbrellaName string `json:"umbrella-name"` // Umbrella profile name (YANG: IOS-XE 17.12.1+)
}

// FlexIPOverlapConfig represents FlexConnect IP overlap configuration.
type FlexIPOverlapConfig struct {
	Enabled bool `json:"enabled,omitempty"` // IP overlap enable flag (YANG: IOS-XE 17.12.1+)
}

// FlexIfNameVlanIDs represents FlexConnect interface name VLAN IDs container.
type FlexIfNameVlanIDs struct {
	IfNameVlanID []FlexIfNameVlanID `json:"if-name-vlan-id"` // Interface name VLAN ID mappings
}

// FlexIfNameVlanID represents individual FlexConnect interface name VLAN ID mapping.
type FlexIfNameVlanID struct {
	InterfaceName string `json:"interface-name"` // Interface name identifier
	VlanID        int    `json:"vlan-id"`        // VLAN identifier
}
