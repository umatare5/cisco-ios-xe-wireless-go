package nmsp

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/constants"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
)

const (
	// NmspOperBasePath defines the base path for NMSP operational data endpoints.
	NmspOperBasePath = constants.YANGModelPrefix + "nmsp-oper:nmsp-oper-data"
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

// GetOper returns NMSP operational data.
func (s Service) GetOper(ctx context.Context) (*model.NmspOperResponse, error) {
	return core.Get[model.NmspOperResponse](ctx, s.c, NmspOperEndpoint)
}

// GetClientRegistration returns NMSP client registration data.
func (s Service) GetClientRegistration(ctx context.Context) (*model.NmspClientRegistrationResponse, error) {
	return core.Get[model.NmspClientRegistrationResponse](ctx, s.c, ClientRegistrationEndpoint)
}

// GetCmxConnection returns NMSP CMX connection data.
func (s Service) GetCmxConnection(ctx context.Context) (*model.NmspCmxConnectionResponse, error) {
	return core.Get[model.NmspCmxConnectionResponse](ctx, s.c, CmxConnectionEndpoint)
}

// GetCmxCloudInfo returns NMSP CMX cloud information.
func (s Service) GetCmxCloudInfo(ctx context.Context) (*model.NmspCmxCloudInfoResponse, error) {
	return core.Get[model.NmspCmxCloudInfoResponse](ctx, s.c, CmxCloudInfoEndpoint)
}
