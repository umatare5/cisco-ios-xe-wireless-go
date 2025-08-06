package nmsp

import (
	"context"
	"net/http"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
	"github.com/umatare5/cisco-ios-xe-wireless-go/wnc"
)

const (
	// NmspOperBasePath defines the base path for NMSP operational data endpoints.
	NmspOperBasePath = "Cisco-IOS-XE-wireless-nmsp-oper:nmsp-oper-data"
	// NmspOperEndpoint defines the endpoint for NMSP operational data.
	NmspOperEndpoint = NmspOperBasePath
	// ClientRegistrationEndpoint defines the endpoint for client registration data.
	ClientRegistrationEndpoint = NmspOperBasePath + "/client-registration"
	// CmxConnectionEndpoint defines the endpoint for CMX connection data.
	CmxConnectionEndpoint = NmspOperBasePath + "/cmx-connection"
	// CmxCloudInfoEndpoint defines the endpoint for CMX cloud information.
	CmxCloudInfoEndpoint = NmspOperBasePath + "/cmx-cloud-info"
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
	err := s.c.Do(ctx, http.MethodGet, NmspOperEndpoint, &out)
	return &out, err
}

// ClientRegistration retrieves NMSP client registration data.
func (s *Service) ClientRegistration(ctx context.Context) (*model.NmspClientRegistrationResponse, error) {
	var out model.NmspClientRegistrationResponse
	err := s.c.Do(ctx, http.MethodGet, ClientRegistrationEndpoint, &out)
	return &out, err
}

// CmxConnection retrieves NMSP CMX connection data.
func (s *Service) CmxConnection(ctx context.Context) (*model.NmspCmxConnectionResponse, error) {
	var out model.NmspCmxConnectionResponse
	err := s.c.Do(ctx, http.MethodGet, CmxConnectionEndpoint, &out)
	return &out, err
}

// CmxCloudInfo retrieves NMSP CMX cloud information.
func (s *Service) CmxCloudInfo(ctx context.Context) (*model.NmspCmxCloudInfoResponse, error) {
	var out model.NmspCmxCloudInfoResponse
	err := s.c.Do(ctx, http.MethodGet, CmxCloudInfoEndpoint, &out)
	return &out, err
}
