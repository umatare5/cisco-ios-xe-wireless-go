package lisp

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/lisp"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/service"
)

// Service provides LISP (Locator/ID Separation Protocol) operations for Cisco IOS-XE Wireless LAN Controller.
type Service struct {
	service.BaseService
}

// NewService creates a new LISP service instance with the provided client.
func NewService(client *core.Client) Service {
	return Service{BaseService: service.NewBaseService(client)}
}

// GetOperational retrieves LISP agent operational data from the controller.
func (s Service) GetOperational(ctx context.Context) (*model.LispAgentOper, error) {
	return core.Get[model.LispAgentOper](ctx, s.Client(), routes.LispOperPath)
}

// GetMemoryStats retrieves LISP agent memory statistics from the wireless controller.
func (s Service) GetMemoryStats(ctx context.Context) (*model.LispAgentOper, error) {
	return core.Get[model.LispAgentOper](ctx, s.Client(), routes.LispMemoryStatsPath)
}

// GetCapabilities retrieves LISP WLC capabilities from the wireless controller.
func (s Service) GetCapabilities(ctx context.Context) (*model.LispAgentOper, error) {
	return core.Get[model.LispAgentOper](ctx, s.Client(), routes.LispCapabilitiesPath)
}
