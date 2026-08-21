package fabric

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/service"
)

// Service provides SD-Access Fabric operations for Cisco IOS-XE Wireless LAN Controller.
type Service struct {
	service.BaseService
}

// NewService creates a new Fabric service instance with the provided client.
func NewService(client *core.Client) Service {
	return Service{BaseService: service.NewBaseService(client)}
}

// GetConfig retrieves fabric configuration data from the controller.
func (s Service) GetConfig(ctx context.Context, opts ...core.GetOption) (*CiscoIOSXEWirelessFabricCfg, error) {
	return core.Get[CiscoIOSXEWirelessFabricCfg](ctx, s.Client(), routes.FabricCfgPath, opts...)
}

// ListCfgFabric retrieves fabric configuration wrapper data.
func (s Service) ListCfgFabric(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessFabricCfgFabric, error) {
	return core.Get[CiscoIOSXEWirelessFabricCfgFabric](ctx, s.Client(), routes.FabricPath, opts...)
}

// ListCfgFabricProfiles retrieves fabric profiles wrapper data.
func (s Service) ListCfgFabricProfiles(
	ctx context.Context,
	opts ...core.GetOption,
) (*CiscoIOSXEWirelessFabricCfgFabricProfiles, error) {
	return core.Get[CiscoIOSXEWirelessFabricCfgFabricProfiles](ctx, s.Client(), routes.FabricProfilesPath, opts...)
}

// ListCfgFabricControlplaneNames retrieves fabric controlplane names wrapper data.
func (s Service) ListCfgFabricControlplaneNames(
	ctx context.Context, opts ...core.GetOption,
) (*CiscoIOSXEWirelessFabricCfgFabricControlplaneNames, error) {
	return core.Get[CiscoIOSXEWirelessFabricCfgFabricControlplaneNames](
		ctx,
		s.Client(),
		routes.FabricControlplaneNamesPath, opts...,
	)
}
