package mcast

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/mcast"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/service"
)

// Service provides multicast management operations for Cisco IOS-XE Wireless LAN Controller.
type Service struct {
	service.BaseService
}

// NewService creates a new Multicast service instance with the provided client.
func NewService(client *core.Client) Service {
	return Service{BaseService: service.NewBaseService(client)}
}

// GetOperational retrieves multicast operational data from the controller.
func (s Service) GetOperational(ctx context.Context) (*model.McastOper, error) {
	return core.Get[model.McastOper](ctx, s.Client(), routes.McastOperPath)
}

// GetFlexConnectMediastreamClientSummary retrieves FlexConnect mediastream client summary data from the wireless controller.
// This function returns information about FlexConnect mediastream clients and their status.
//
// Parameters:
//   - ctx: Context for request timeout and cancellation control
//
// Returns:
//   - *model.McastOperFlexMediastreamClientSummary: FlexConnect mediastream data
//   - error: Error if the operation fails
//
// Example:
//
//	service := mcast.NewService(client)
//	data, err := service.GetFlexConnectMediastreamClientSummary(ctx)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("FlexConnect mediastream data: %+v\n", data)
func (s Service) GetFlexConnectMediastreamClientSummary(
	ctx context.Context,
) (*model.McastOperFlexMediastreamClientSummary, error) {
	return core.Get[model.McastOperFlexMediastreamClientSummary](
		ctx, s.Client(), routes.McastFlexMediastreamPath,
	)
}

// ListVLANL2MGIDs retrieves VLAN Layer 2 multicast group ID operational data from the wireless controller.
// This function returns information about VLAN Layer 2 multicast groups and their configuration.
//
// Parameters:
//   - ctx: Context for request timeout and cancellation control
//
// Returns:
//   - *model.McastOperVlanL2MgidOp: VLAN Layer 2 multicast group data
//   - error: Error if the operation fails
//
// Example:
//
//	service := mcast.NewService(client)
//	data, err := service.ListVLANL2MGIDs(ctx)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("VLAN L2 multicast group data: %+v\n", data)
func (s Service) ListVLANL2MGIDs(ctx context.Context) (*model.McastOperVlanL2MgidOp, error) {
	return core.Get[model.McastOperVlanL2MgidOp](ctx, s.Client(), routes.McastVlanL2MgidPath)
}

// GetFabricMediastreamClientSummary retrieves fabric mediastream client summary data from the wireless controller.
// This function returns information about fabric mediastream clients and their status.
//
// Parameters:
//   - ctx: Context for request timeout and cancellation control
//
// Returns:
//   - *model.McastOperFabricMediaStreamClientSummary: Fabric mediastream data
//   - error: Error if the operation fails
func (s Service) GetFabricMediastreamClientSummary(
	ctx context.Context,
) (*model.McastOperFabricMediaStreamClientSummary, error) {
	return core.Get[model.McastOperFabricMediaStreamClientSummary](
		ctx, s.Client(), routes.McastFabricMediastreamPath,
	)
}

// GetMcastMgidInfo retrieves multicast MGID information from the wireless controller.
// This function returns information about multicast group IDs and their configuration.
//
// Parameters:
//   - ctx: Context for request timeout and cancellation control
//
// Returns:
//   - *model.McastOperMcastMgidInfo: Multicast MGID information
//   - error: Error if the operation fails
func (s Service) GetMcastMgidInfo(ctx context.Context) (*model.McastOperMcastMgidInfo, error) {
	return core.Get[model.McastOperMcastMgidInfo](ctx, s.Client(), routes.McastMgidInfoPath)
}

// GetMulticastOperData retrieves multicast operational data from the wireless controller.
// This function returns detailed operational information about multicast operations.
//
// Parameters:
//   - ctx: Context for request timeout and cancellation control
//
// Returns:
//   - *model.McastOperMulticastOperData: Multicast operational data
//   - error: Error if the operation fails
func (s Service) GetMulticastOperData(ctx context.Context) (*model.McastOperMulticastOperData, error) {
	return core.Get[model.McastOperMulticastOperData](ctx, s.Client(), routes.McastMulticastOperDataPath)
}
