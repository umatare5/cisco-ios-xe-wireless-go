package hyperlocation

import (
	"context"
	"net/http"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
)

const (
	// HyperlocationOperBasePath defines the base path for hyperlocation operational data endpoints.
	HyperlocationOperBasePath = "Cisco-IOS-XE-wireless-hyperlocation-oper:hyperlocation-oper-data"
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

// Oper returns hyperlocation operational data.
func (s Service) Oper(ctx context.Context) (*model.HyperlocationOperResponse, error) {
	var out model.HyperlocationOperResponse
	return &out, s.c.Do(ctx, http.MethodGet, HyperlocationOperEndpoint, &out)
}

// Profiles returns hyperlocation profiles.
func (s Service) Profiles(ctx context.Context) (*model.HyperlocationProfilesResponse, error) {
	var out model.HyperlocationProfilesResponse
	return &out, s.c.Do(ctx, http.MethodGet, HyperlocationProfilesEndpoint, &out)
}
