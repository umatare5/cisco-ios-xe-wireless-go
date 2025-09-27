package flex

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/service"
)

// Service provides FlexConnect operations for Cisco IOS-XE Wireless LAN Controller.
type Service struct {
	service.BaseService
}

// NewService creates a new FlexConnect service instance with the provided client.
func NewService(client *core.Client) Service {
	return Service{BaseService: service.NewBaseService(client)}
}

// GetConfig retrieves FlexConnect configuration data.
func (s Service) GetConfig(ctx context.Context) (*CiscoIOSXEWirelessFlexCfg, error) {
	return core.Get[CiscoIOSXEWirelessFlexCfg](ctx, s.Client(), routes.FlexCfgPath)
}

// ListFlexPolicyEntries retrieves FlexConnect policy entries.
func (s Service) ListFlexPolicyEntries(ctx context.Context) (*CiscoIOSXEWirelessFlexCfgFlexPolicyEntries, error) {
	return core.Get[CiscoIOSXEWirelessFlexCfgFlexPolicyEntries](ctx, s.Client(), routes.FlexPolicyEntriesPath)
}
