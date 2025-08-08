package mcast

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/constants"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
)

const (
	// McastOperBasePath defines the base path for multicast operational data endpoints.
	McastOperBasePath = constants.YANGModelPrefix + "mcast-oper:mcast-oper-data"
	// McastOperEndpoint defines the endpoint for multicast operational data.
	McastOperEndpoint = McastOperBasePath
	// FlexMediastreamClientSummaryEndpoint defines the endpoint for FlexConnect mediastream client summary data.
	FlexMediastreamClientSummaryEndpoint = McastOperBasePath + "/flex-mediastream-client-summary"
	// VlanL2MgidOpEndpoint defines the endpoint for VLAN Layer 2 multicast group ID operational data.
	VlanL2MgidOpEndpoint = McastOperBasePath + "/vlan-l2-mgid-op"
)

// Service provides Multicast operations.
type Service struct {
	c *core.Client
}

// NewService creates a new service instance.
func NewService(c *core.Client) Service {
	return Service{c: c}
}

// Oper returns multicast operational data.
func (s Service) Oper(ctx context.Context) (*model.McastOperResponse, error) {
	return core.Get[model.McastOperResponse](ctx, s.c, McastOperEndpoint)
}

// FlexMediastreamClientSummary returns FlexConnect mediastream client summary data.
func (s Service) FlexMediastreamClientSummary(
	ctx context.Context,
) (*model.McastOperFlexMediastreamClientSummaryResponse, error) {
	return core.Get[model.McastOperFlexMediastreamClientSummaryResponse](
		ctx, s.c, FlexMediastreamClientSummaryEndpoint,
	)
}

// VlanL2MgidOp returns VLAN Layer 2 multicast group ID operational data.
func (s Service) VlanL2MgidOp(ctx context.Context) (*model.McastOperVlanL2MgidOpResponse, error) {
	return core.Get[model.McastOperVlanL2MgidOpResponse](ctx, s.c, VlanL2MgidOpEndpoint)
}
