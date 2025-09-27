package nmsp

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
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

// GetOperational retrieves the complete NMSP operational data.
func (s Service) GetOperational(ctx context.Context) (*CiscoIOSXEWirelessNMSPOper, error) {
	return core.Get[CiscoIOSXEWirelessNMSPOper](ctx, s.Client(), routes.NMSPOperPath)
}

// ListClientRegistrations retrieves NMSP client registration data.
func (s Service) ListClientRegistrations(ctx context.Context) (*CiscoIOSXEWirelessNMSPClientRegistration, error) {
	return core.Get[CiscoIOSXEWirelessNMSPClientRegistration](ctx, s.Client(), routes.NMSPClientRegistrationPath)
}

// GetCMXConnectionInfo retrieves NMSP CMX connection information.
func (s Service) GetCMXConnectionInfo(ctx context.Context) (*CiscoIOSXEWirelessNMSPCmxConnection, error) {
	return core.Get[CiscoIOSXEWirelessNMSPCmxConnection](ctx, s.Client(), routes.NMSPCmxConnectionPath)
}

// GetCMXCloudInfo retrieves NMSP CMX cloud information.
func (s Service) GetCMXCloudInfo(ctx context.Context) (*CiscoIOSXEWirelessNMSPCmxCloudInfo, error) {
	return core.Get[CiscoIOSXEWirelessNMSPCmxCloudInfo](ctx, s.Client(), routes.NMSPCmxCloudInfoPath)
}
