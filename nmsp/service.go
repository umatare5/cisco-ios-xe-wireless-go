// Package nmsp provides Network Mobility Services Protocol service for the Cisco Wireless Network Controller API.
// This package implements the Domain Service layer of the three-layer architecture.
package nmsp

import (
	"context"
	"net/http"

	wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
)

// Service provides NMSP (Network Mobility Services Protocol) operations.
type Service struct {
	c *wnc.Client
}

// NewService creates a new NMSP service instance.
func NewService(c *wnc.Client) *Service {
	return &Service{c: c}
}

// Oper retrieves NMSP operational data.
func (s *Service) Oper(ctx context.Context) (*model.NmspOperResponse, error) {
	var out model.NmspOperResponse
	err := s.c.CoreClient().Do(ctx, http.MethodGet, NmspOperEndpoint, &out)
	return &out, err
}

// ClientRegistration retrieves NMSP client registration data.
func (s *Service) ClientRegistration(ctx context.Context) (*model.NmspClientRegistrationResponse, error) {
	var out model.NmspClientRegistrationResponse
	err := s.c.CoreClient().Do(ctx, http.MethodGet, ClientRegistrationEndpoint, &out)
	return &out, err
}

// CmxConnection retrieves NMSP CMX connection data.
func (s *Service) CmxConnection(ctx context.Context) (*model.NmspCmxConnectionResponse, error) {
	var out model.NmspCmxConnectionResponse
	err := s.c.CoreClient().Do(ctx, http.MethodGet, CmxConnectionEndpoint, &out)
	return &out, err
}

// CmxCloudInfo retrieves NMSP CMX cloud information.
func (s *Service) CmxCloudInfo(ctx context.Context) (*model.NmspCmxCloudInfoResponse, error) {
	var out model.NmspCmxCloudInfoResponse
	err := s.c.CoreClient().Do(ctx, http.MethodGet, CmxCloudInfoEndpoint, &out)
	return &out, err
}
