package spaces

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/service"
)

// Service provides Cisco Spaces integration operations for Cisco IOS-XE Wireless LAN Controller.
type Service struct {
	service.BaseService
}

// NewService creates a new Spaces service instance with the provided client.
// EXPERIMENTAL: Requires IOS-XE 17.18.1+.
func NewService(client *core.Client) Service {
	return Service{BaseService: service.NewBaseService(client)}
}

// GetOperational retrieves all Cisco Spaces operational data from the controller.
// EXPERIMENTAL: Requires IOS-XE 17.18.1+.
func (s Service) GetOperational(ctx context.Context) (*CiscoSpacesOper, error) {
	return core.Get[CiscoSpacesOper](ctx, s.Client(), routes.SpacesOperPath)
}

// GetConnectionDetails retrieves detailed connection information for Cisco Spaces integration.
// EXPERIMENTAL: Requires IOS-XE 17.18.1+.
func (s Service) GetConnectionDetails(ctx context.Context) (*SpacesConnectionDetail, error) {
	return core.Get[SpacesConnectionDetail](ctx, s.Client(), routes.SpacesConnectionDetailPath)
}
