package dot15

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/service"
)

// Service provides IEEE 802.15 Bluetooth configuration operations for Cisco IOS-XE Wireless LAN Controller.
type Service struct {
	service.BaseService
}

// NewService creates a new 802.15 service instance with the provided client.
func NewService(client *core.Client) Service {
	return Service{BaseService: service.NewBaseService(client)}
}

// GetConfig retrieves dot15.4 configuration data from the controller.
func (s Service) GetConfig(ctx context.Context) (*Dot15Cfg, error) {
	return core.Get[Dot15Cfg](ctx, s.Client(), routes.Dot15CfgPath)
}

// ListDot15GlobalConfigs retrieves 802.15 global configuration.
func (s Service) ListDot15GlobalConfigs(ctx context.Context) (*Dot15CfgDot15GlobalConfig, error) {
	return core.Get[Dot15CfgDot15GlobalConfig](ctx, s.Client(), routes.Dot15GlobalConfigPath)
}
