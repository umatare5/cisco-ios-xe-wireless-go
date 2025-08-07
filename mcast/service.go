package mcast

import (
	"context"
	"net/http"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
	"github.com/umatare5/cisco-ios-xe-wireless-go/wnc"
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

// Service provides multicast operations.
type Service struct {
	c *wnc.Client
}

// NewService creates a new service instance.
func NewService(c *wnc.Client) Service {
	return Service{c: c}
}

// Oper returns multicast operational data.
func (s Service) Oper(ctx context.Context) (*model.McastOperResponse, error) {
	var out model.McastOperResponse
	return &out, s.c.Do(ctx, http.MethodGet, McastOperEndpoint, &out)
}

// FlexMediastreamClientSummary returns FlexConnect mediastream client summary data.
func (s Service) FlexMediastreamClientSummary(
	ctx context.Context,
) (*model.McastOperFlexMediastreamClientSummaryResponse, error) {
	var out model.McastOperFlexMediastreamClientSummaryResponse
	return &out, s.c.Do(ctx, http.MethodGet, FlexMediastreamClientSummaryEndpoint, &out)
}

// VlanL2MgidOp returns VLAN Layer 2 multicast group ID operational data.
func (s Service) VlanL2MgidOp(ctx context.Context) (*model.McastOperVlanL2MgidOpResponse, error) {
	var out model.McastOperVlanL2MgidOpResponse
	return &out, s.c.Do(ctx, http.MethodGet, VlanL2MgidOpEndpoint, &out)
}
