package dot11

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/service"
)

// Service provides IEEE 802.11 wireless configuration operations for Cisco IOS-XE Wireless LAN Controller.
type Service struct {
	service.BaseService
}

// NewService creates a new 802.11 service instance with the provided client.
func NewService(client *core.Client) Service {
	return Service{BaseService: service.NewBaseService(client)}
}

// GetConfig retrieves dot11 configuration data from the controller.
func (s Service) GetConfig(ctx context.Context) (*Dot11Cfg, error) {
	return core.Get[Dot11Cfg](ctx, s.Client(), routes.Dot11CfgPath)
}
