// Package model provides data models for multicast operational data.
package model

// McastOper  represents the structure for multicast operational data.
type McastOper struct {
	CiscoIOSXEWirelessMcastOperMcastOperData struct {
		FlexMediastreamClientSummary []FlexMediastreamClientSummary `json:"flex-mediastream-client-summary"`
		VlanL2MgidOp                 []VlanL2MgidOp                 `json:"vlan-l2-mgid-op"`
	} `json:"Cisco-IOS-XE-wireless-mcast-oper:mcast-oper-data"`
}

// McastOperFlexMediastreamClientSummary  represents the structure for FlexConnect mediastream client summary data.
type McastOperFlexMediastreamClientSummary struct {
	FlexMediastreamClientSummary []FlexMediastreamClientSummary `json:"Cisco-IOS-XE-wireless-mcast-oper:flex-mediastream-client-summary"`
}

// McastOperVlanL2MgidOp  represents the structure for VLAN Layer 2 multicast group ID operational data.
type McastOperVlanL2MgidOp struct {
	VlanL2MgidOp []VlanL2MgidOp `json:"Cisco-IOS-XE-wireless-mcast-oper:vlan-l2-mgid-op"`
}

type FlexMediastreamClientSummary struct {
	ClientMac            string                 `json:"client-mac"`
	VlanID               int                    `json:"vlan-id"`
	FlexMcastClientGroup []FlexMcastClientGroup `json:"flex-mcast-client-group"`
}

type FlexMcastClientGroup struct {
	McastIP    string `json:"mcast-ip"`
	StreamName string `json:"stream-name"`
	ApMac      string `json:"ap-mac"`
	IsDirect   bool   `json:"is-direct"`
}

type VlanL2MgidOp struct {
	VlanIndex               int  `json:"vlan-index"`
	IsNonipMulticastEnabled bool `json:"is-nonip-multicast-enabled"`
	IsBroadcastEnable       bool `json:"is-broadcast-enable"`
}
