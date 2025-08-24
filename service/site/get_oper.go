package site

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/site"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
)

// operOps provides high-level operational operations for site service
func (s Service) operOps() *core.OperationalOperations[model.SiteOper] {
	return core.NewOperationalOperations[model.SiteOper](s.Client(), routes.SiteOperBasePath)
}

// GetOper retrieves complete site operational data from the wireless controller.
//
// This method returns all site operational information including current site
// status, active configurations, and runtime statistics. The response contains
// the full site operational data hierarchy as defined in the Cisco-IOS-XE-wireless-site-oper
// YANG model.
//
// Note: Site operational data may not be available on all controller models
// or firmware versions. Check the response carefully for data availability.
//
// Parameters:
//   - ctx: Context for the request, allowing for timeout and cancellation control
//
// Returns:
//   - *model.SiteOper: Complete site operational data structure
//   - error: Any error encountered during the operation
//
// Example:
//
//	oper, err := siteService.GetOper(ctx)
//	if err != nil {
//		log.Printf("Failed to get site operational data: %v", err)
//		return err
//	}
//
//	// Access site operational information
//	if oper.SiteOperData != nil {
//		fmt.Printf("Site operational data retrieved successfully\n")
//		// Process operational data as needed
//	} else {
//		fmt.Printf("Site operational data not available on this controller\n")
//	}
func (s Service) GetOper(ctx context.Context) (*model.SiteOper, error) {
	return s.operOps().GetAll(ctx)
}
