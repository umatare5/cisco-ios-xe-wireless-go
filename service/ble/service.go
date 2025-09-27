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
func (s Service) GetOperational(ctx context.Context) (*CiscoIOSXEWirelessBLELtxOper, error) {
	return core.Get[CiscoIOSXEWirelessBLELtxOper](ctx, s.Client(), routes.BLELtxOperPath)
}

// ListBLELtxAp retrieves BLE LTX AP operational data.
func (s Service) ListBLELtxAp(ctx context.Context) (*CiscoIOSXEWirelessBLELtxOperBLELtxAp, error) {
	return core.Get[CiscoIOSXEWirelessBLELtxOperBLELtxAp](ctx, s.Client(), routes.BLELtxApPath)
}

// ListBLELtxApAntenna retrieves BLE LTX AP antenna operational data.
func (s Service) ListBLELtxApAntenna(ctx context.Context) (*CiscoIOSXEWirelessBLELtxOperBLELtxApAntenna, error) {
	return core.Get[CiscoIOSXEWirelessBLELtxOperBLELtxApAntenna](ctx, s.Client(), routes.BLELtxApAntennaPath)
}

// GetMgmtOperational retrieves BLE management operational data from the controller.
func (s Service) GetMgmtOperational(ctx context.Context) (*CiscoIOSXEWirelessBLEMgmtOper, error) {
	return core.Get[CiscoIOSXEWirelessBLEMgmtOper](ctx, s.Client(), routes.BLEMgmtOperPath)
}

// ListBLEMgmtAp retrieves BLE management AP operational data.
func (s Service) ListBLEMgmtAp(ctx context.Context) (*CiscoIOSXEWirelessBLEMgmtOperBLEMgmtAp, error) {
	return core.Get[CiscoIOSXEWirelessBLEMgmtOperBLEMgmtAp](ctx, s.Client(), routes.BLEMgmtApPath)
}

// ListBLEMgmtCmx retrieves BLE management CMX operational data.
func (s Service) ListBLEMgmtCmx(ctx context.Context) (*CiscoIOSXEWirelessBLEMgmtOperBLEMgmtCmx, error) {
	return core.Get[CiscoIOSXEWirelessBLEMgmtOperBLEMgmtCmx](ctx, s.Client(), routes.BLEMgmtCmxPath)
}
