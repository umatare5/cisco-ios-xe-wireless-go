package ap

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	ierrors "github.com/umatare5/cisco-ios-xe-wireless-go/internal/errors"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/validation"
)

// EnableAP enables the administrative state of an access point
func (s *Service) EnableAP(ctx context.Context, mac string) error {
	return s.updateAPState(ctx, mac, "admin-state-enabled")
}

// DisableAP disables the administrative state of an access point
func (s *Service) DisableAP(ctx context.Context, mac string) error {
	return s.updateAPState(ctx, mac, "admin-state-disabled")
}

// updateAPState handles AP admin state changes with mac and mode parameters
func (s *Service) updateAPState(ctx context.Context, mac, mode string) error {
	if err := s.ValidateClient(); err != nil {
		return err
	}

	normalizedMAC := validation.NormalizeAPMac(mac)

	payload := map[string]any{
		"input": map[string]any{
			"mac-addr": normalizedMAC,
			"mode":     mode,
		},
	}

	if err := core.PostRPCVoid(ctx, s.Client(), routes.SetAPAdminStateRPC, payload); err != nil {
		return ierrors.ServiceOperationError("update", "AP", "admin state", err)
	}

	return nil
}
