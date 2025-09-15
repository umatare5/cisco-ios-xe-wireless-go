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

// GetOperational retrieves LISP operational data from the wireless controller.
func (s Service) GetOperational(ctx context.Context) (*model.LISPAgentOper, error) {
	return core.Get[model.LISPAgentOper](ctx, s.Client(), routes.LISPOperPath)
}

// GetMemoryStats retrieves LISP agent memory statistics from the wireless controller.
func (s Service) GetMemoryStats(ctx context.Context) (*model.LISPAgentMemoryStats, error) {
	return core.Get[model.LISPAgentMemoryStats](ctx, s.Client(), routes.LISPMemoryStatsPath)
}

// GetCapabilities retrieves LISP WLC capabilities from the wireless controller.
func (s Service) GetCapabilities(ctx context.Context) (*model.LISPWLCCapabilities, error) {
	return core.Get[model.LISPWLCCapabilities](ctx, s.Client(), routes.LISPCapabilitiesPath)
}
