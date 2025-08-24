package rrm

import (
	"context"
	"fmt"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/rrm"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/validation"
)

// operOps provides high-level operational operations for rrm service
func (s Service) operOps() *core.OperationalOperations[model.RrmOper] {
	return core.NewOperationalOperations[model.RrmOper](s.Client(), routes.RRMOperBasePath)
}

// GetOper retrieves RRM operational data.
func (s Service) GetOper(ctx context.Context) (*model.RrmOper, error) {
	return s.operOps().GetAll(ctx)
}

// GetOperByDeviceID retrieves RRM operational data filtered by device ID.
func (s Service) GetOperByDeviceID(ctx context.Context, deviceID string) (*model.RrmOper, error) {
	if err := validation.ValidateNonEmptyString(deviceID, "deviceID"); err != nil {
		return nil, err
	}
	return s.operOps().GetByCompositeKey(ctx, "device", deviceID)
}

// GetOperByPhyType retrieves RRM operational data filtered by PHY type.
func (s Service) GetOperByPhyType(ctx context.Context, phyType string) (*model.RrmOper, error) {
	if err := validation.ValidateNonEmptyString(phyType, "phyType"); err != nil {
		return nil, err
	}
	return s.operOps().GetByCompositeKey(ctx, "phy-type", phyType)
}

// GetOperByWtpMacAndRadioSlot retrieves RRM operational data filtered by WTP MAC and radio slot ID.
func (s Service) GetOperByWtpMacAndRadioSlot(
	ctx context.Context,
	wtpMac string,
	radioSlotID int,
) (*model.RrmOper, error) {
	if err := validation.ValidateNonEmptyString(wtpMac, "wtpMac"); err != nil {
		return nil, err
	}
	if err := validation.ValidateSlotID(radioSlotID); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s=%s,%d", routes.RRMOperEndpoint, wtpMac, radioSlotID)
	return core.Get[model.RrmOper](ctx, s.Client(), url)
}

// GetGlobalOper retrieves RRM global operational data.
func (s Service) GetGlobalOper(ctx context.Context) (*model.RrmGlobalOper, error) {
	return core.Get[model.RrmGlobalOper](ctx, s.Client(), routes.RRMGlobalOperEndpoint)
}

// GetGlobalOperBy5GWtpMacAndRadioSlot returns RRM global operational data filtered by WTP MAC and radio slot ID for 5GHz.
func (s Service) GetGlobalOperBy5GWtpMacAndRadioSlot(
	ctx context.Context, wtpMac string, radioSlotID int,
) (*model.RrmGlobalOper, error) {
	if err := validation.ValidateNonEmptyString(wtpMac, "wtpMac"); err != nil {
		return nil, err
	}
	if err := validation.ValidateSlotID(radioSlotID); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s=%s,%d", routes.RRMOperEndpoint, wtpMac, radioSlotID)
	return core.Get[model.RrmGlobalOper](ctx, s.Client(), url)
}

// GetGlobalOperBy6GhzWtpMacAndRadioSlot returns 6 GHz RRM global operational data for a specific WTP MAC and radio slot from the Cisco Catalyst 9800 WLC.
func (s Service) GetGlobalOperBy6GhzWtpMacAndRadioSlot(
	ctx context.Context, wtpMac string, radioSlotID int,
) (*model.RrmGlobalOper, error) {
	if err := validation.ValidateNonEmptyString(wtpMac, "wtpMac"); err != nil {
		return nil, err
	}
	if err := validation.ValidateSlotID(radioSlotID); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s=%s,%d", routes.RRMOperEndpoint, wtpMac, radioSlotID)
	return core.Get[model.RrmGlobalOper](ctx, s.Client(), url)
}

// GetGlobalOperByApMac retrieves RRM global operational data filtered by AP MAC address.
func (s Service) GetGlobalOperByApMac(ctx context.Context, apMac string) (*model.RrmGlobalOper, error) {
	if err := validation.ValidateNonEmptyString(apMac, "apMac"); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s=%s", routes.RRMGlobalOperEndpoint, apMac)
	return core.Get[model.RrmGlobalOper](ctx, s.Client(), url)
}

// GetGlobalOperByBandID retrieves RRM global operational data filtered by band ID.
func (s Service) GetGlobalOperByBandID(ctx context.Context, bandID string) (*model.RrmGlobalOper, error) {
	if err := validation.ValidateNonEmptyString(bandID, "bandID"); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s=%s", routes.RRMGlobalOperEndpoint, bandID)
	return core.Get[model.RrmGlobalOper](ctx, s.Client(), url)
}

// GetGlobalOperByChannelPhyType retrieves RRM global operational data filtered by PHY type for channel parameters.
func (s Service) GetGlobalOperByChannelPhyType(
	ctx context.Context, phyType string,
) (*model.RrmGlobalOper, error) {
	if err := validation.ValidateNonEmptyString(phyType, "phyType"); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/phy-type=%s", routes.RRMOperEndpoint, phyType)
	return core.Get[model.RrmGlobalOper](ctx, s.Client(), url)
}

// GetGlobalOperByPhyType retrieves RRM global operational data filtered by PHY type.
func (s Service) GetGlobalOperByPhyType(ctx context.Context, phyType string) (*model.RrmGlobalOper, error) {
	if err := validation.ValidateNonEmptyString(phyType, "phyType"); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/phy-type=%s", routes.RRMOperEndpoint, phyType)
	return core.Get[model.RrmGlobalOper](ctx, s.Client(), url)
}

// GetGlobalOperByWtpMacAndRadioSlot retrieves RRM global operational data filtered by WTP MAC and radio slot ID for 2.4GHz.
func (s Service) GetGlobalOperByWtpMacAndRadioSlot(
	ctx context.Context, wtpMac string, radioSlotID int,
) (*model.RrmGlobalOper, error) {
	if err := validation.ValidateNonEmptyString(wtpMac, "wtpMac"); err != nil {
		return nil, err
	}
	if err := validation.ValidateSlotID(radioSlotID); err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s=%s,%d", routes.RRMOperEndpoint, wtpMac, radioSlotID)
	return core.Get[model.RrmGlobalOper](ctx, s.Client(), url)
}

// GetEmulOper retrieves RRM emulation operational data.
func (s Service) GetEmulOper(ctx context.Context) (*model.RrmEmulOper, error) {
	return core.Get[model.RrmEmulOper](ctx, s.Client(), routes.RRMEmulOperEndpoint)
}
