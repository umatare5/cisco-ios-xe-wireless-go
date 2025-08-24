package ap

import (
	"context"
	"errors"
	"fmt"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/helpers"
	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/ap"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/validation"
)

// Reload restarts an Access Point by MAC address causing temporary service interruption.
func (s Service) Reload(ctx context.Context, apMac string) error {
	if !validation.IsValidAPMacFormat(apMac) {
		return fmt.Errorf(ErrInvalidAPMacFormat, apMac)
	}

	resp, err := s.GetOperCapwapDataAll(ctx)
	if !validation.IsValid(resp, err) {
		if validation.HasError(err) {
			return fmt.Errorf(ErrFailedGetCAPWAPData, err)
		}
		return errors.New(ErrCAPWAPDataUnavailable)
	}

	apName, found := helpers.FindAPByMAC(resp, apMac)
	if !found {
		return fmt.Errorf(ErrAPNotFoundByMAC, apMac)
	}

	return s.reload(ctx, apName)
}

// reload is the internal helper function for AP reload operations
func (s Service) reload(ctx context.Context, apName string) error {
	requestBody := model.APReloadRPCPayload{
		Input: model.APReloadRPCInput{
			APName: apName,
		},
	}
	return core.PostRPCVoid(ctx, s.Client(), routes.APResetRPC, requestBody)
}
