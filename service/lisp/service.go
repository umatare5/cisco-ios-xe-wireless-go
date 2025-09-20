package lisp

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
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
func (s Service) GetOperational(ctx context.Context) (*LISPAgentOper, error) {
	return core.Get[LISPAgentOper](ctx, s.Client(), routes.LISPOperPath)
}

// GetMemoryStats retrieves LISP agent memory statistics from the wireless controller.
func (s Service) GetMemoryStats(ctx context.Context) (*LISPAgentMemoryStats, error) {
	return core.Get[LISPAgentMemoryStats](ctx, s.Client(), routes.LISPMemoryStatsPath)
}

// GetCapabilities retrieves LISP WLC capabilities from the wireless controller.
func (s Service) GetCapabilities(ctx context.Context) (*LISPWLCCapabilities, error) {
	return core.Get[LISPWLCCapabilities](ctx, s.Client(), routes.LISPCapabilitiesPath)
}
