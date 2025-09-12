package awips

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/awips"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/service"
)

// Service provides AWIPS (Automated Wireless IPS) operations for Cisco IOS-XE Wireless LAN Controller.
type Service struct {
	service.BaseService
}

// NewService creates a new AWIPS service instance with the provided client.
func NewService(client *core.Client) Service {
	return Service{BaseService: service.NewBaseService(client)}
}

// GetOperational retrieves the complete AWIPS operational data.
func (s Service) GetOperational(ctx context.Context) (*model.AwipsOper, error) {
	return core.Get[model.AwipsOper](ctx, s.Client(), routes.AWIPSOperPath)
}
