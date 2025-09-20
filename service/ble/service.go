package ble

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/service"
)

// Service provides BLE (Bluetooth Low Energy) operations for Cisco IOS-XE Wireless LAN Controller.
type Service struct {
	service.BaseService
}

// NewService creates a new BLE service instance with the provided client.
func NewService(client *core.Client) Service {
	return Service{BaseService: service.NewBaseService(client)}
}

// GetOperational retrieves BLE operational data from the controller.
func (s Service) GetOperational(ctx context.Context) (*BLELtxOper, error) {
	return core.Get[BLELtxOper](ctx, s.Client(), routes.BLELtxOperPath)
}
