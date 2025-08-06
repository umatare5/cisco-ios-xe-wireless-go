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

// Service provides multicast operational data access.
type Service struct {
	c *wnc.Client
}

// NewService creates a new mcast service instance.
func NewService(c *wnc.Client) Service {
	return Service{c: c}
}

// Oper retrieves multicast operational data.
func (s Service) Oper(ctx context.Context) (*model.McastOperResponse, error) {
	var out model.McastOperResponse
	err := s.c.Do(ctx, http.MethodGet, McastOperEndpoint, &out)
	return &out, err
}

// FlexMediastreamClientSummary retrieves FlexConnect mediastream client summary data.
func (s Service) FlexMediastreamClientSummary(ctx context.Context) (*model.McastOperFlexMediastreamClientSummaryResponse, error) {
	var out model.McastOperFlexMediastreamClientSummaryResponse
	err := s.c.Do(ctx, http.MethodGet, FlexMediastreamClientSummaryEndpoint, &out)
	return &out, err
}

// VlanL2MgidOp retrieves VLAN Layer 2 multicast group ID operational data.
func (s Service) VlanL2MgidOp(ctx context.Context) (*model.McastOperVlanL2MgidOpResponse, error) {
	var out model.McastOperVlanL2MgidOpResponse
	err := s.c.Do(ctx, http.MethodGet, VlanL2MgidOpEndpoint, &out)
	return &out, err
}
