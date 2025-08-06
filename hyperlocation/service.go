package hyperlocation

import (
	"context"
	"net/http"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
	"github.com/umatare5/cisco-ios-xe-wireless-go/wnc"
)

const (
	// HyperlocationOperBasePath defines the base path for hyperlocation operational data endpoints.
	HyperlocationOperBasePath = "Cisco-IOS-XE-wireless-hyperlocation-oper:hyperlocation-oper-data"
	// HyperlocationOperEndpoint defines the endpoint for hyperlocation operational data.
	HyperlocationOperEndpoint = HyperlocationOperBasePath
	// HyperlocationProfilesEndpoint defines the endpoint for hyperlocation profiles.
	HyperlocationProfilesEndpoint = HyperlocationOperBasePath + "/ewlc-hyperlocation-profile"
)

// Service provides hyperlocation operations.
type Service struct {
	c *wnc.Client
}

// NewService creates a new hyperlocation service instance.
func NewService(c *wnc.Client) *Service {
	return &Service{c: c}
}

// Oper retrieves hyperlocation operational data.
func (s *Service) Oper(ctx context.Context) (*model.HyperlocationOperResponse, error) {
	var out model.HyperlocationOperResponse
	err := s.c.Do(ctx, http.MethodGet, HyperlocationOperEndpoint, &out)
	return &out, err
}

// Profiles retrieves hyperlocation profiles.
func (s *Service) Profiles(ctx context.Context) (*model.HyperlocationProfilesResponse, error) {
	var out model.HyperlocationProfilesResponse
	err := s.c.Do(ctx, http.MethodGet, HyperlocationProfilesEndpoint, &out)
	return &out, err
}
