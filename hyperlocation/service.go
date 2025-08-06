// Package hyperlocation provides hyperlocation service for the Cisco Wireless Network Controller API.
// This package implements the Domain Service layer of the three-layer architecture.
package hyperlocation

import (
	"context"
	"net/http"

	wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
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
	err := s.c.CoreClient().Do(ctx, http.MethodGet, HyperlocationOperEndpoint, &out)
	return &out, err
}

// Profiles retrieves hyperlocation profiles.
func (s *Service) Profiles(ctx context.Context) (*model.HyperlocationProfilesResponse, error) {
	var out model.HyperlocationProfilesResponse
	err := s.c.CoreClient().Do(ctx, http.MethodGet, HyperlocationProfilesEndpoint, &out)
	return &out, err
}
