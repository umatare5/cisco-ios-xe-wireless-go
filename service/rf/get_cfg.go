package rf

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/rf"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
)

// GetCfg retrieves RF configuration data including RF profiles, power settings, and channel configurations.
//
// This method returns comprehensive configuration information including all
// RF settings, policies, and operational parameters. The response
// contains the complete configuration hierarchy as defined in the corresponding
// YANG model.
//
// Parameters:
//   - ctx: Context for the request, allowing for timeout and cancellation control
//
// Returns:
//   - *model.RfCfg: Complete RF configuration data structure
//   - error: Any error encountered during the operation
//
// Example:
//
//	cfg, err := rfService.GetCfg(ctx)
//	if err != nil {
//		log.Printf("Failed to get RF configuration: %v", err)
//		return err
//	}
//
//	// Process configuration data
//	fmt.Printf("Retrieved RF configuration successfully\n")
func (s Service) GetCfg(ctx context.Context) (*model.RfCfg, error) {
	if err := s.ValidateClient(); err != nil {
		return nil, err
	}
	return core.Get[model.RfCfg](ctx, s.Client(), routes.RFCfgBasePath)
}
