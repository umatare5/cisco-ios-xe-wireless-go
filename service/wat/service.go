package wat

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/service"
)

// Service provides Wireless Application Templates (WAT) operations for Cisco IOS-XE Wireless LAN Controller.
type Service struct {
	service.BaseService
}

// NewService creates a new WAT service instance with the provided client.
// EXPERIMENTAL: Requires IOS-XE 17.18.1+.
func NewService(client *core.Client) Service {
	return Service{BaseService: service.NewBaseService(client)}
}

// GetConfig retrieves the complete WAT configuration from the controller.
// EXPERIMENTAL: Requires IOS-XE 17.18.1+.
func (s Service) GetConfig(ctx context.Context) (*WATCfg, error) {
	return core.Get[WATCfg](ctx, s.Client(), routes.WATCfgPath)
}
