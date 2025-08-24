package lisp

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/lisp"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
)

// operOps provides high-level operational operations for lisp service
func (s Service) operOps() *core.OperationalOperations[model.LispOper] {
	return core.NewOperationalOperations[model.LispOper](s.Client(), routes.LispOperBasePath)
}

// GetOper retrieves all LISP agent operational data from the wireless controller.
// This function returns comprehensive LISP agent operational information including
// memory statistics, WLC capabilities, and AP type capabilities.
//
// Parameters:
//   - ctx: Context for request timeout and cancellation control
//
// Returns:
//   - *model.LispOper: Complete LISP agent operational data
//   - error: Error if the operation fails
//
// Example:
//
//	service := lisp.NewService(client)
//	data, err := service.GetOper(ctx)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("LISP agent operational data: %+v\n", data)
func (s Service) GetOper(ctx context.Context) (*model.LispOper, error) {
	return s.operOps().GetAll(ctx)
}

// GetOperMemoryStats retrieves LISP agent memory statistics from the wireless controller.
// This function returns detailed memory allocation and deallocation statistics
// for various LISP agent components.
//
// Parameters:
//   - ctx: Context for request timeout and cancellation control
//
// Returns:
//   - *model.LispOper: LISP agent memory statistics
//   - error: Error if the operation fails
//
// Example:
//
//	service := lisp.NewService(client)
//	stats, err := service.GetOperMemoryStats(ctx)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("LISP memory stats: %+v\n", stats)
func (s Service) GetOperMemoryStats(ctx context.Context) (*model.LispOper, error) {
	ops := core.NewOperationalOperations[model.LispOper](s.Client(), routes.LispOperMemoryStatsEndpoint)
	return ops.GetAll(ctx)
}

// GetOperCapabilities retrieves LISP WLC capabilities from the wireless controller.
// This function returns information about LISP fabric capabilities and AP type support.
//
// Parameters:
//   - ctx: Context for request timeout and cancellation control
//
// Returns:
//   - *model.LispOper: LISP WLC and AP capabilities
//   - error: Error if the operation fails
//
// Example:
//
//	service := lisp.NewService(client)
//	capabilities, err := service.GetOperCapabilities(ctx)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("LISP capabilities: %+v\n", capabilities)
func (s Service) GetOperCapabilities(ctx context.Context) (*model.LispOper, error) {
	ops := core.NewOperationalOperations[model.LispOper](s.Client(), routes.LispOperCapabilitiesEndpoint)
	return ops.GetAll(ctx)
}
