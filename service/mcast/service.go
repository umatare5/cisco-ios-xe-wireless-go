package mcast

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
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
func (s Service) GetOperational(ctx context.Context) (*McastOper, error) {
	return core.Get[McastOper](ctx, s.Client(), routes.McastOperPath)
}

// GetFlexConnectMediastreamClientSummary retrieves FlexConnect mediastream client summary data from the wireless controller.
// This function returns information about FlexConnect mediastream clients and their status.
//
// Parameters:
//   - ctx: Context for request timeout and cancellation control
//
// Returns:
//   - *McastOperFlexMediastreamClientSummary: FlexConnect mediastream data
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
) (*McastOperFlexMediastreamClientSummary, error) {
	return core.Get[McastOperFlexMediastreamClientSummary](
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
//   - *McastOperVlanL2MgidOp: VLAN Layer 2 multicast group data
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
func (s Service) ListVLANL2MGIDs(ctx context.Context) (*McastOperVlanL2MgidOp, error) {
	return core.Get[McastOperVlanL2MgidOp](ctx, s.Client(), routes.McastVlanL2MgidPath)
}

// GetFabricMediastreamClientSummary retrieves fabric mediastream client summary data from the wireless controller.
// This function returns information about fabric mediastream clients and their status.
//
// Parameters:
//   - ctx: Context for request timeout and cancellation control
//
// Returns:
//   - *McastOperFabricMediaStreamClientSummary: Fabric mediastream data
//   - error: Error if the operation fails
func (s Service) GetFabricMediastreamClientSummary(
	ctx context.Context,
) (*McastOperFabricMediaStreamClientSummary, error) {
	return core.Get[McastOperFabricMediaStreamClientSummary](
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
//   - *McastOperMcastMgidInfo: Multicast MGID information
//   - error: Error if the operation fails
func (s Service) GetMcastMgidInfo(ctx context.Context) (*McastOperMcastMgidInfo, error) {
	return core.Get[McastOperMcastMgidInfo](ctx, s.Client(), routes.McastMgidInfoPath)
}

// GetMulticastOperData retrieves multicast operational data from the wireless controller.
// This function returns detailed operational information about multicast operations.
//
// Parameters:
//   - ctx: Context for request timeout and cancellation control
//
// Returns:
//   - *McastOperMulticastOperData: Multicast operational data
//   - error: Error if the operation fails
func (s Service) GetMulticastOperData(ctx context.Context) (*McastOperMulticastOperData, error) {
	return core.Get[McastOperMulticastOperData](ctx, s.Client(), routes.McastMulticastOperDataPath)
}
