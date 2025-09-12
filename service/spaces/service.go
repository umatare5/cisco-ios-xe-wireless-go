package spaces

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/spaces"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/service"
)

// Service provides Cisco Spaces integration operations for Cisco IOS-XE Wireless LAN Controller.
type Service struct {
	service.BaseService
}

// NewService creates a new Spaces service instance with the provided client.
func NewService(client *core.Client) Service {
	return Service{BaseService: service.NewBaseService(client)}
}

// GetOperational retrieves all Cisco Spaces operational data from the controller.
// EXPERIMENTAL: Requires IOS-XE 17.18.1+.
func (s Service) GetOperational(ctx context.Context) (*model.CiscoSpacesOper, error) {
	return core.Get[model.CiscoSpacesOper](ctx, s.Client(), routes.SpacesOperPath)
}

// GetConnectionDetails retrieves detailed connection information for Cisco Spaces integration.
// EXPERIMENTAL: Requires IOS-XE 17.18.1+.
func (s Service) GetConnectionDetails(ctx context.Context) (*model.SpacesConnectionDetail, error) {
	return core.Get[model.SpacesConnectionDetail](ctx, s.Client(), routes.SpacesConnectionDetailPath)
}

// GetTenantInfo retrieves tenant information for Cisco Spaces multi-tenant environments.
// EXPERIMENTAL: Requires IOS-XE 17.18.1+.
func (s Service) GetTenantInfo(ctx context.Context) (*model.SpacesTenant, error) {
	// Get connection details first, then extract tenant info
	connDetails, err := s.GetConnectionDetails(ctx)
	if err != nil {
		return nil, err
	}

	return connDetails.Tenant, nil
}

// GetConnectionStats retrieves statistics and performance metrics for Cisco Spaces integration.
// EXPERIMENTAL: Requires IOS-XE 17.18.1+.
func (s Service) GetConnectionStats(ctx context.Context) (*model.SpacesConnectionStats, error) {
	// Get connection details first, then extract stats
	connDetails, err := s.GetConnectionDetails(ctx)
	if err != nil {
		return nil, err
	}

	return connDetails.Stats, nil
}
