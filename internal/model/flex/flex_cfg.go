// Package model provides data models for FlexConnect configuration data.
package model

// FlexCfg  represents the FlexConnect configuration data.
type FlexCfg struct {
	FlexCfgData FlexCfgData `json:"Cisco-IOS-XE-wireless-flex-cfg:flex-cfg-data"`
}

// FlexCfgFlexPolicyEntries  represents the flex policy entries.
type FlexCfgFlexPolicyEntries struct {
	FlexPolicyEntries FlexPolicyEntries `json:"Cisco-IOS-XE-wireless-flex-cfg:flex-policy-entries"`
}

type FlexCfgData struct {
	FlexPolicyEntries FlexPolicyEntries `json:"flex-policy-entries"`
}

type FlexPolicyEntries struct {
	FlexPolicyEntry []FlexPolicyEntry `json:"flex-policy-entry"`
}

type FlexPolicyEntry struct {
	PolicyName            string               `json:"policy-name"`
	NativeVlanID          int                  `json:"native-vlan-id,omitempty"`
	DefaultAP             bool                 `json:"default-ap,omitempty"`
	ACLMapMode            string               `json:"acl-map-mode,omitempty"`
	WebPolicyACLName      string               `json:"web-policy-acl-name,omitempty"`
	CentralAssociation    bool                 `json:"central-association,omitempty"`
	CentralAuthentication bool                 `json:"central-authentication,omitempty"`
	CentralDhcp           bool                 `json:"central-dhcp,omitempty"`
	CentralSwitching      bool                 `json:"central-switching,omitempty"`
	FlexAcls              []FlexACL            `json:"flex-acls,omitempty"`
	FlexVlanACLMappings   []FlexVlanACLMapping `json:"flex-vlan-acl-mappings,omitempty"`
	FlexLocalUsers        []FlexLocalUser      `json:"flex-local-users,omitempty"`
	RadiusServers         []FlexRadiusServer   `json:"radius-servers,omitempty"`
	PolicyACLName         string               `json:"policy-acl-name,omitempty"`
	UmbrellaProfiles      FlexUmbrellaProfiles `json:"umbrella-profiles,omitempty"`
	MdnsProfileName       string               `json:"mdns-profile-name,omitempty"`
	IPOverlapCfg          FlexIPOverlapConfig  `json:"ip-overlap-cfg,omitempty"`
	DhcpBcast             bool                 `json:"dhcp-bcast,omitempty"`
	PmkDistMethod         string               `json:"pmk-dist-method,omitempty"`
}

type FlexACL struct {
	ACLName      string `json:"acl-name"`
	ACLDirection string `json:"acl-direction"`
}

type FlexVlanACLMapping struct {
	VlanID      int    `json:"vlan-id"`
	InboundACL  string `json:"inbound-acl,omitempty"`
	OutboundACL string `json:"outbound-acl,omitempty"`
}

type FlexLocalUser struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	PasswordType string `json:"password-type"`
	UserRole     string `json:"user-role"`
}

type FlexRadiusServer struct {
	ServerName    string `json:"server-name"`
	ServerAddress string `json:"server-address"`
	SharedSecret  string `json:"shared-secret"`
}

type FlexUmbrellaProfiles struct {
	UmbrellaProfile []FlexUmbrellaProfile `json:"umbrella-profile"`
}

type FlexUmbrellaProfile struct {
	UmbrellaName string `json:"umbrella-name"`
	DeviceID     string `json:"device-id,omitempty"`
	APIKey       string `json:"api-key,omitempty"`
}

type FlexIPOverlapConfig struct {
	Enabled      bool   `json:"enabled,omitempty"`
	LocalSubnet  string `json:"local-subnet,omitempty"`
	LocalGateway string `json:"local-gateway,omitempty"`
}
