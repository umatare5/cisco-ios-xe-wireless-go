// Package model contains generated response structures for the Cisco WNC API.
// This file contains multicast operational data structures.
package model

// McastOperResponse represents the response structure for multicast operational data.
type McastOperResponse struct {
	CiscoIOSXEWirelessMcastOperMcastOperData struct {
		FlexMediastreamClientSummary []FlexMediastreamClientSummary `json:"flex-mediastream-client-summary"`
		VlanL2MgidOp                 []VlanL2MgidOp                 `json:"vlan-l2-mgid-op"`
	} `json:"Cisco-IOS-XE-wireless-mcast-oper:mcast-oper-data"`
}

// McastOperFlexMediastreamClientSummaryResponse represents the response structure for FlexConnect mediastream client summary data.
type McastOperFlexMediastreamClientSummaryResponse struct {
	FlexMediastreamClientSummary []FlexMediastreamClientSummary `json:"Cisco-IOS-XE-wireless-mcast-oper:flex-mediastream-client-summary"`
}

// McastOperVlanL2MgidOpResponse represents the response structure for VLAN Layer 2 multicast group ID operational data.
type McastOperVlanL2MgidOpResponse struct {
	VlanL2MgidOp []VlanL2MgidOp `json:"Cisco-IOS-XE-wireless-mcast-oper:vlan-l2-mgid-op"`
}

// FlexMediastreamClientSummary represents a FlexConnect mediastream client summary entry.
type FlexMediastreamClientSummary struct {
	ClientMac            string                 `json:"client-mac"`
	VlanID               int                    `json:"vlan-id"`
	FlexMcastClientGroup []FlexMcastClientGroup `json:"flex-mcast-client-group"`
}

// FlexMcastClientGroup represents a FlexConnect multicast client group entry.
type FlexMcastClientGroup struct {
	McastIP    string `json:"mcast-ip"`
	StreamName string `json:"stream-name"`
	ApMac      string `json:"ap-mac"`
	IsDirect   bool   `json:"is-direct"`
}

// VlanL2MgidOp represents VLAN Layer 2 multicast group ID operational data.
type VlanL2MgidOp struct {
	VlanIndex               int  `json:"vlan-index"`
	IsNonipMulticastEnabled bool `json:"is-nonip-multicast-enabled"`
	IsBroadcastEnable       bool `json:"is-broadcast-enable"`
}
