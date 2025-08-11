package model

// FlexCfgResponse represents the response structure for FlexConnect configuration data.
type FlexCfgResponse struct {
	FlexCfgData FlexCfgData `json:"Cisco-IOS-XE-wireless-flex-cfg:flex-cfg-data"`
}

// FlexCfgData contains FlexConnect configuration data
type FlexCfgData struct {
	FlexGroups []FlexGroup `json:"flex-groups"`
}

// FlexGroup represents a FlexConnect group configuration
type FlexGroup struct {
	GroupName           string     `json:"group-name"`
	Vlans               []FlexVlan `json:"vlans"`
	CentralAuth         bool       `json:"central-auth"`
	CentralDhcp         bool       `json:"central-dhcp"`
	BackupRadiusServers []string   `json:"backup-radius-servers"`
	LocalSplitTunneling bool       `json:"local-split-tunneling"`
}

// FlexVlan represents VLAN configuration in FlexConnect
type FlexVlan struct {
	VlanID     int    `json:"vlan-id"`
	VlanName   string `json:"vlan-name"`
	NativeVlan bool   `json:"native-vlan"`
}
