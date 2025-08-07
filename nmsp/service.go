package nmsp

import (
	"context"
	"net/http"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
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

// Service provides NMSP operations.
type Service struct {
	c *core.Client
}

// NewService creates a new service instance.
func NewService(c *core.Client) Service {
	return Service{c: c}
}

// Oper returns NMSP operational data.
func (s Service) Oper(ctx context.Context) (*model.NmspOperResponse, error) {
	var out model.NmspOperResponse
	return &out, s.c.Do(ctx, http.MethodGet, NmspOperEndpoint, &out)
}

// ClientRegistration returns NMSP client registration data.
func (s Service) ClientRegistration(ctx context.Context) (*model.NmspClientRegistrationResponse, error) {
	var out model.NmspClientRegistrationResponse
	return &out, s.c.Do(ctx, http.MethodGet, ClientRegistrationEndpoint, &out)
}

// CmxConnection returns NMSP CMX connection data.
func (s Service) CmxConnection(ctx context.Context) (*model.NmspCmxConnectionResponse, error) {
	var out model.NmspCmxConnectionResponse
	return &out, s.c.Do(ctx, http.MethodGet, CmxConnectionEndpoint, &out)
}

// CmxCloudInfo returns NMSP CMX cloud information.
func (s Service) CmxCloudInfo(ctx context.Context) (*model.NmspCmxCloudInfoResponse, error) {
	var out model.NmspCmxCloudInfoResponse
	return &out, s.c.Do(ctx, http.MethodGet, CmxCloudInfoEndpoint, &out)
}
