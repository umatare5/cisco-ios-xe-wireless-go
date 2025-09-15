package nmsp

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/nmsp"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/service"
)

// Service provides NMSP (Network Mobility Services Protocol) operations for Cisco IOS-XE Wireless LAN Controller.
type Service struct {
	service.BaseService
}

// NewService creates a new NMSP service instance with the provided client.
func NewService(client *core.Client) Service {
	return Service{BaseService: service.NewBaseService(client)}
}

// GetOperational retrieves NMSP operational data from the controller.
func (s Service) GetOperational(ctx context.Context) (*model.NMSPOper, error) {
	return core.Get[model.NMSPOper](ctx, s.Client(), routes.NMSPOperPath)
}

// ListClientRegistrations retrieves NMSP client registration data.
func (s Service) ListClientRegistrations(ctx context.Context) (*model.NMSPClientRegistration, error) {
	return core.Get[model.NMSPClientRegistration](ctx, s.Client(), routes.NMSPClientRegistrationPath)
}

// GetCMXConnectionInfo retrieves NMSP CMX connection information.
func (s Service) GetCMXConnectionInfo(ctx context.Context) (*model.NMSPCmxConnection, error) {
	return core.Get[model.NMSPCmxConnection](ctx, s.Client(), routes.NMSPCmxConnectionPath)
}

// GetCMXCloudInfo retrieves NMSP CMX cloud information.
func (s Service) GetCMXCloudInfo(ctx context.Context) (*model.NMSPCmxCloudInfo, error) {
	return core.Get[model.NMSPCmxCloudInfo](ctx, s.Client(), routes.NMSPCmxCloudInfoPath)
}
