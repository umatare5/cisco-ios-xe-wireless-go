package location

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/location"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
)

// operOps provides high-level operational operations for location service
func (s Service) operOps() *core.OperationalOperations[model.LocationCfg] {
	return core.NewOperationalOperations[model.LocationCfg](s.Client(), routes.LocationOperEndpoint)
}

// GetOper retrieves all location operational data from the wireless controller.
// This function returns comprehensive location operational information including
// statistics and status information.
//
// Parameters:
//   - ctx: Context for request timeout and cancellation control
//
// Returns:
//   - *model.LocationCfg: Complete location operational data
//   - error: Error if the operation fails
//
// Example:
//
//	service := location.NewService(client)
//	data, err := service.GetOper(ctx)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Location operational data: %+v\n", data)
func (s Service) GetOper(ctx context.Context) (*model.LocationCfg, error) {
	return s.operOps().GetAll(ctx)
}

// GetOperStats retrieves location statistics operational data from the wireless controller.
// This function returns information about location service statistics and metrics.
//
// Parameters:
//   - ctx: Context for request timeout and cancellation control
//
// Returns:
//   - *model.LocationCfg: Location statistics operational data
//   - error: Error if the operation fails
//
// Example:
//
//	service := location.NewService(client)
//	stats, err := service.GetOperStats(ctx)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Location statistics: %+v\n", stats)
func (s Service) GetOperStats(ctx context.Context) (*model.LocationCfg, error) {
	return core.Get[model.LocationCfg](ctx, s.Client(), routes.LocationOperStatsEndpoint)
}
