package rf

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/rf"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/service"
)

// Service provides RF (Radio Frequency) management operations for Cisco IOS-XE Wireless LAN Controller.
type Service struct {
	service.BaseService
}

// NewService creates a new RF service instance with the provided client.
func NewService(client *core.Client) Service {
	return Service{BaseService: service.NewBaseService(client)}
}

// RFTag returns an RF tag service instance for RF tag management operations.
func (s Service) RFTag() *RFTagService {
	return NewRFTagService(s.Client())
}

// GetConfig retrieves RF configuration data including RF profiles and power settings.
func (s Service) GetConfig(ctx context.Context) (*model.RfCfg, error) {
	return core.Get[model.RfCfg](ctx, s.Client(), routes.RFCfgPath)
}
