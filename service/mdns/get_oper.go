package mdns

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/mdns"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/validation"
)

// operOps provides high-level operational operations for mdns service
func (s Service) operOps() *core.OperationalOperations[model.MdnsOper] {
	return core.NewOperationalOperations[model.MdnsOper](s.Client(), routes.MDNSOperEndpoint)
}

// GetOper retrieves all mDNS operational data from the wireless controller.
// This function returns comprehensive mDNS operational information including
// global statistics, WLAN-specific statistics, and service discovery data.
//
// Parameters:
//   - ctx: Context for request timeout and cancellation control
//
// Returns:
//   - *model.MdnsOper: Complete mDNS operational data
//   - error: Error if the operation fails
//
// Example:
//
//	service := mdns.NewService(client)
//	data, err := service.GetOper(ctx)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("mDNS operational data: %+v\n", data)
func (s Service) GetOper(ctx context.Context) (*model.MdnsOper, error) {
	return s.operOps().GetAll(ctx)
}

// GetOperGlobalStats retrieves mDNS global statistics from the wireless controller.
// This function returns global mDNS statistics including packet counts, protocol metrics,
// and overall service discovery performance data.
//
// Parameters:
//   - ctx: Context for request timeout and cancellation control
//
// Returns:
//   - *model.MdnsGlobalStats: Global mDNS statistics data
//   - error: Error if the operation fails
//
// Example:
//
//	service := mdns.NewService(client)
//	stats, err := service.GetOperGlobalStats(ctx)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Global mDNS statistics: %+v\n", stats)
func (s Service) GetOperGlobalStats(ctx context.Context) (*model.MdnsGlobalStats, error) {
	if err := s.ValidateClient(); err != nil {
		return nil, err
	}
	return core.Get[model.MdnsGlobalStats](ctx, s.Client(), routes.MDNSGlobalStatsEndpoint)
}

// GetOperWlanStats retrieves mDNS WLAN statistics from the wireless controller.
// This function returns WLAN-specific mDNS statistics including per-WLAN packet counts,
// service discovery metrics, and WLAN-level mDNS performance data.
//
// Parameters:
//   - ctx: Context for request timeout and cancellation control
//
// Returns:
//   - *model.MdnsWlanStats: WLAN mDNS statistics data
//   - error: Error if the operation fails
//
// Example:
//
//	service := mdns.NewService(client)
//	stats, err := service.GetOperWlanStats(ctx)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("WLAN mDNS statistics: %+v\n", stats)
func (s Service) GetOperWlanStats(ctx context.Context) (*model.MdnsWlanStats, error) {
	if err := s.ValidateClient(); err != nil {
		return nil, err
	}
	return core.Get[model.MdnsWlanStats](ctx, s.Client(), routes.MDNSWlanStatsEndpoint)
}

// GetOperByWlanID retrieves mDNS operational data for a specific WLAN ID.
// This function provides targeted access to WLAN-specific mDNS statistics and
// service discovery data for a particular WLAN.
//
// Parameters:
//   - ctx: Context for request timeout and cancellation control
//   - wlanID: Specific WLAN ID to retrieve data for
//
// Returns:
//   - *model.MdnsWlanStats: mDNS data for the specified WLAN
//   - error: Error if the operation fails or wlanID is invalid
//
// Example:
//
//	service := mdns.NewService(client)
//	data, err := service.GetOperByWlanID(ctx, "1")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("WLAN mDNS data: %+v\n", data)
func (s Service) GetOperByWlanID(ctx context.Context, wlanID string) (*model.MdnsWlanStats, error) {
	if err := s.ValidateClient(); err != nil {
		return nil, err
	}
	if err := validation.ValidateWlanID(wlanID); err != nil {
		return nil, err
	}
	endpoint := s.Client().RestconfBuilder().BuildPathQueryURL(routes.MDNSOperEndpoint, "mdns-wlan-stats", wlanID)
	return core.Get[model.MdnsWlanStats](ctx, s.Client(), endpoint)
}
