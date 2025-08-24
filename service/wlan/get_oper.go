package wlan

import (
	"context"
	"fmt"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/wlan"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
)

// operOps provides high-level operational operations for wlan service
func (s Service) operOps() *core.OperationalOperations[model.WlanGlobalOper] {
	return core.NewOperationalOperations[model.WlanGlobalOper](s.Client(), routes.WlanGlobalOperDataEndpoint)
}

// GetOper retrieves WLAN operational data from the controller.
//
// This method provides access to real-time operational information about
// WLAN configurations and their current status. Note that some controllers
// may not support all operational data endpoints.
//
// Parameters:
//   - ctx: Context for the request, allowing for timeout and cancellation control
//
// Returns:
//   - *model.WlanGlobalOper: WLAN operational data structure
//   - error: Any error encountered during the operation, including 404 if not supported
//
// Example:
//
//	wlanService := wlan.NewService(client)
//	oper, err := wlanService.GetOper(ctx)
//	if err != nil {
//		// Note: Some controllers may not support operational data
//		return err
//	}
func (s Service) GetOper(ctx context.Context) (*model.WlanGlobalOper, error) {
	return s.operOps().GetAll(ctx)
}

// GetOperByWlanProfile retrieves operational data for a specified WLAN profile.
//
// This method allows retrieval of operational data for a specific WLAN
// identified by its profile name. This provides targeted access to
// operational metrics for individual WLANs.
//
// Parameters:
//   - ctx: Context for the request, allowing for timeout and cancellation control
//   - wlanProfile: Name of the WLAN profile to retrieve operational data for
//
// Returns:
//   - *model.WlanGlobalOper: WLAN operational data for the specified profile
//   - error: Any error encountered during the operation, including 404 if not found
//
// Example:
//
//	oper, err := wlanService.GetOperByWlanProfile(ctx, "corporate-wlan")
//	if err != nil {
//		return err
//	}
func (s Service) GetOperByWlanProfile(
	ctx context.Context,
	wlanProfile string,
) (*model.WlanGlobalOper, error) {
	url := fmt.Sprintf("%s/wlan-profile=%s", routes.WlanProfilesEndpoint, wlanProfile)
	return core.Get[model.WlanGlobalOper](ctx, s.Client(), url)
}
