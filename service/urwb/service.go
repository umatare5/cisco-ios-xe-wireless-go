package urwb

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/urwb"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/service"
)

// Service provides access to URWB (Ultra Reliable Wireless Backhaul) operations.
type Service struct {
	service.BaseService
}

// NewService creates a new URWB service instance with the provided client.
// EXPERIMENTAL: Requires IOS-XE 17.18.1+.
func NewService(client *core.Client) Service {
	return Service{BaseService: service.NewBaseService(client)}
}

// GetConfig retrieves the complete URWB configuration from the controller.
// EXPERIMENTAL: Requires IOS-XE 17.18.1+.
func (s Service) GetConfig(ctx context.Context) (*model.URWBCfg, error) {
	return core.Get[model.URWBCfg](ctx, s.Client(), routes.URWBCfgPath)
}

// GetURWBNetOperational retrieves URWB network operational data from the controller.
// EXPERIMENTAL: Requires IOS-XE 17.18.1+.
func (s Service) GetURWBNetOperational(ctx context.Context) (*model.URWBnetOper, error) {
	return core.Get[model.URWBnetOper](ctx, s.Client(), routes.URWBNetOperPath)
}
