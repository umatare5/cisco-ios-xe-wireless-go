package ap

import (
	"context"
	"strconv"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	ierrors "github.com/umatare5/cisco-ios-xe-wireless-go/internal/errors"
	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/ap"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/validation"
)

// EnableRadio enables a radio on an Access Point using MAC address.
func (s Service) EnableRadio(ctx context.Context, apMac string, radioBand core.RadioBand) error {
	if err := validation.ValidateAPMac(apMac); err != nil {
		return err
	}
	return s.updateRadioState(ctx, apMac, &radioBand, true)
}

// DisableRadio disables a radio on an Access Point using MAC address.
func (s Service) DisableRadio(ctx context.Context, apMac string, radioBand core.RadioBand) error {
	if err := validation.ValidateAPMac(apMac); err != nil {
		return err
	}
	return s.updateRadioState(ctx, apMac, &radioBand, false)
}

// updateRadioState handles radio-level state changes
func (s Service) updateRadioState(ctx context.Context, apMac string, radioBand *core.RadioBand, enabled bool) error {
	if validation.IsNil(radioBand) {
		return ierrors.RequiredParameterError("radio band")
	}

	radioBandInfo, err := core.GetRadioBandInfo(int(*radioBand))
	if err != nil {
		return err
	}

	payload := model.APSlotConfigRPCPayload{
		Input: model.APSlotConfigRPCInput{
			Mode:    core.GetAdminStateMode(enabled),
			SlotID:  int(radioBandInfo.SlotID),
			Band:    strconv.Itoa(int(radioBandInfo.Band)),
			MACAddr: validation.NormalizeAPMac(apMac),
		},
	}

	if err := core.PostRPCVoid(ctx, s.Client(), routes.SetAPSlotAdminStateRPC, payload); err != nil {
		return ierrors.ServiceOperationError("set", "AP radio", "state", err)
	}
	return nil
}
