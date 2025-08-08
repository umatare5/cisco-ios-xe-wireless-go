package hyperlocation

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/constants"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
)

const (
	// HyperlocationOperBasePath defines the base path for hyperlocation operational data endpoints.
	HyperlocationOperBasePath = constants.YANGModelPrefix + "hyperlocation-oper:hyperlocation-oper-data"
	// HyperlocationOperEndpoint defines the endpoint for hyperlocation operational data.
	HyperlocationOperEndpoint = HyperlocationOperBasePath
	// HyperlocationProfilesEndpoint defines the endpoint for hyperlocation profiles.
	HyperlocationProfilesEndpoint = HyperlocationOperBasePath + "/ewlc-hyperlocation-profile"
)

// Service provides Hyperlocation operations.
type Service struct {
	c *core.Client
}

// NewService creates a new service instance.
func NewService(c *core.Client) Service {
	return Service{c: c}
}

// GetOper returns hyperlocation operational data.
func (s Service) GetOper(ctx context.Context) (*model.HyperlocationOperResponse, error) {
	return core.Get[model.HyperlocationOperResponse](ctx, s.c, HyperlocationOperEndpoint)
}

// GetProfiles returns hyperlocation profiles.
func (s Service) GetProfiles(ctx context.Context) (*model.HyperlocationProfilesResponse, error) {
	return core.Get[model.HyperlocationProfilesResponse](ctx, s.c, HyperlocationProfilesEndpoint)
}
