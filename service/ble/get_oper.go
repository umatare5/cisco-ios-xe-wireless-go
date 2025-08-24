// Package ble provides Bluetooth Low Energy operational data operations for the Cisco IOS-XE Wireless Network Controller API.
package ble

import (
	"context"
	"strconv"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/ble"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
)

// operOps provides high-level operational operations for ble service
func (s Service) operOps() *core.OperationalOperations[model.BleLtxOper] {
	return core.NewOperationalOperations[model.BleLtxOper](s.Client(), routes.BLELtxOperBasePath)
}

// GetOper retrieves the complete BLE LTX operational data
func (s Service) GetOper(ctx context.Context) (*model.BleLtxOper, error) {
	return s.operOps().GetAll(ctx)
}

// GetOperByApMac retrieves BLE LTX operational data filtered by AP MAC address
func (s Service) GetOperByApMac(ctx context.Context, apMacAddr string) (*model.BleLtxOper, error) {
	return s.operOps().GetByMAC(ctx, "ble-ltx-ap", apMacAddr)
}

// GetOperByApMacSlotAntenna retrieves BLE LTX antenna data filtered by AP MAC address, slot ID, and antenna ID
func (s Service) GetOperByApMacSlotAntenna(
	ctx context.Context, apMacAddr string, slotID, antennaID int,
) (*model.BleLtxOper, error) {
	keys := []string{apMacAddr, strconv.Itoa(slotID), strconv.Itoa(antennaID)}
	return s.operOps().GetByCompositeKey(ctx, "ble-ltx-ap-antenna", keys...)
}
