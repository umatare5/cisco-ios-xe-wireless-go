package radio

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/service"
)

// Service provides radio configuration management operations for Cisco IOS-XE Wireless LAN Controller.
type Service struct {
	service.BaseService
}

// NewService creates a new Radio service instance with the provided client.
func NewService(client *core.Client) Service {
	return Service{BaseService: service.NewBaseService(client)}
}

// GetConfig retrieves radio configuration data from the controller.
func (s Service) GetConfig(ctx context.Context) (*RadioCfg, error) {
	return core.Get[RadioCfg](ctx, s.Client(), routes.RadioCfgPath)
}

// ListProfileConfigs retrieves radio profiles configuration data.
func (s Service) ListProfileConfigs(ctx context.Context) (*RadioProfiles, error) {
	return core.Get[RadioProfiles](ctx, s.Client(), routes.RadioCfgPath+"/radio-profiles")
}
