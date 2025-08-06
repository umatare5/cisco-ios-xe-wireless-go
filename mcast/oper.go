// Package mcast provides multicast operational data functionality for the Cisco Wireless Network Controller API.
package mcast

import (
	"context"
	"errors"

	wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
)

const (
	// McastOperBasePath defines the base path for multicast operational data endpoints.
	McastOperBasePath = "Cisco-IOS-XE-wireless-mcast-oper:mcast-oper-data"
	// McastOperEndpoint defines the endpoint for multicast operational data.
	McastOperEndpoint = McastOperBasePath
	// FlexMediastreamClientSummaryEndpoint defines the endpoint for FlexConnect mediastream client summary data.
	FlexMediastreamClientSummaryEndpoint = McastOperBasePath + "/flex-mediastream-client-summary"
	// VlanL2MgidOpEndpoint defines the endpoint for VLAN Layer 2 multicast group ID operational data.
	VlanL2MgidOpEndpoint = McastOperBasePath + "/vlan-l2-mgid-op"
)

// Type aliases for backward compatibility.
type (
	// Deprecated: Use model.McastOperResponse instead.
	McastOperResponse = model.McastOperResponse
	// Deprecated: Use model.McastOperFlexMediastreamClientSummaryResponse instead.
	McastOperFlexMediastreamClientSummaryResponse = model.McastOperFlexMediastreamClientSummaryResponse
	// Deprecated: Use model.McastOperVlanL2MgidOpResponse instead.
	McastOperVlanL2MgidOpResponse = model.McastOperVlanL2MgidOpResponse
	// Deprecated: Use model.FlexMediastreamClientSummary instead.
	FlexMediastreamClientSummary = model.FlexMediastreamClientSummary
	// Deprecated: Use model.FlexMcastClientGroup instead.
	FlexMcastClientGroup = model.FlexMcastClientGroup
	// Deprecated: Use model.VlanL2MgidOp instead.
	VlanL2MgidOp = model.VlanL2MgidOp
)

// Deprecated: Use client.Mcast().Oper(ctx) instead.
// GetMcastOper retrieves multicast operational data.
func GetMcastOper(client *wnc.Client, ctx context.Context) (*McastOperResponse, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	service := NewService(client.CoreClient())
	return service.Oper(ctx)
}

// Deprecated: Use client.Mcast().FlexMediastreamClientSummary(ctx) instead.
// GetMcastFlexMediastreamClientSummary retrieves FlexConnect mediastream client summary data.
func GetMcastFlexMediastreamClientSummary(client *wnc.Client, ctx context.Context) (*McastOperFlexMediastreamClientSummaryResponse, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	service := NewService(client.CoreClient())
	return service.FlexMediastreamClientSummary(ctx)
}

// Deprecated: Use client.Mcast().VlanL2MgidOp(ctx) instead.
// GetMcastVlanL2MgidOp retrieves VLAN Layer 2 multicast group ID operational data.
func GetMcastVlanL2MgidOp(client *wnc.Client, ctx context.Context) (*McastOperVlanL2MgidOpResponse, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	service := NewService(client.CoreClient())
	return service.VlanL2MgidOp(ctx)
}
