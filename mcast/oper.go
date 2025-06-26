// Package mcast provides multicast operational data functionality for the Cisco Wireless Network Controller API.
package mcast

import (
	"context"

	wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
)

const (
	// McastOperBasePath defines the base path for multicast operational data endpoints.
	McastOperBasePath = "/restconf/data/Cisco-IOS-XE-wireless-mcast-oper:mcast-oper-data"
	// McastOperEndpoint defines the endpoint for multicast operational data.
	McastOperEndpoint = McastOperBasePath
	// FlexMediastreamClientSummaryEndpoint defines the endpoint for FlexConnect mediastream client summary data.
	FlexMediastreamClientSummaryEndpoint = McastOperBasePath + "/flex-mediastream-client-summary"
	// VlanL2MgidOpEndpoint defines the endpoint for VLAN Layer 2 multicast group ID operational data.
	VlanL2MgidOpEndpoint = McastOperBasePath + "/vlan-l2-mgid-op"
)

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

// GetMcastOper retrieves multicast operational data.
func GetMcastOper(client *wnc.Client, ctx context.Context) (*McastOperResponse, error) {
	var data McastOperResponse
	if err := client.SendAPIRequest(ctx, McastOperEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetMcastFlexMediastreamClientSummary retrieves FlexConnect mediastream client summary data.
func GetMcastFlexMediastreamClientSummary(client *wnc.Client, ctx context.Context) (*McastOperFlexMediastreamClientSummaryResponse, error) {
	var data McastOperFlexMediastreamClientSummaryResponse
	if err := client.SendAPIRequest(ctx, FlexMediastreamClientSummaryEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetMcastVlanL2MgidOp retrieves VLAN Layer 2 multicast group ID operational data.
func GetMcastVlanL2MgidOp(client *wnc.Client, ctx context.Context) (*McastOperVlanL2MgidOpResponse, error) {
	var data McastOperVlanL2MgidOpResponse
	if err := client.SendAPIRequest(ctx, VlanL2MgidOpEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}
