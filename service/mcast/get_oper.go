package mcast

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/mcast"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/validation"
)

// operOps provides high-level operational operations for mcast service
func (s Service) operOps() *core.OperationalOperations[model.McastOper] {
	return core.NewOperationalOperations[model.McastOper](s.Client(), routes.McastOperBasePath)
}

// GetOper retrieves all multicast operational data from the wireless controller.
// This function returns comprehensive multicast operational information including
// FlexConnect mediastream clients, VLAN Layer 2 multicast groups, and statistics.
//
// Parameters:
//   - ctx: Context for request timeout and cancellation control
//
// Returns:
//   - *model.McastOper: Complete multicast operational data
//   - error: Error if the operation fails
//
// Example:
//
//	service := mcast.NewService(client)
//	data, err := service.GetOper(ctx)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Multicast operational data: %+v\n", data)
func (s Service) GetOper(ctx context.Context) (*model.McastOper, error) {
	if err := s.ValidateClient(); err != nil {
		return nil, err
	}
	return s.operOps().GetAll(ctx)
}

// GetOperFlexMediastream retrieves FlexConnect mediastream client summary data from the wireless controller.
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
//	data, err := service.GetOperFlexMediastream(ctx)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("FlexConnect mediastream data: %+v\n", data)
func (s Service) GetOperFlexMediastream(
	ctx context.Context,
) (*model.McastOperFlexMediastreamClientSummary, error) {
	if err := s.ValidateClient(); err != nil {
		return nil, err
	}
	return core.Get[model.McastOperFlexMediastreamClientSummary](
		ctx, s.Client(), routes.McastOperFlexMediastreamEndpoint,
	)
}

// GetOperVlanL2Mgid retrieves VLAN Layer 2 multicast group ID operational data from the wireless controller.
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
//	data, err := service.GetOperVlanL2Mgid(ctx)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("VLAN L2 multicast group data: %+v\n", data)
func (s Service) GetOperVlanL2Mgid(ctx context.Context) (*model.McastOperVlanL2MgidOp, error) {
	if err := s.ValidateClient(); err != nil {
		return nil, err
	}
	return core.Get[model.McastOperVlanL2MgidOp](ctx, s.Client(), routes.McastOperVlanL2MgidEndpoint)
}

// GetOperByClientMAC retrieves multicast operational data for a specific client MAC address.
// This function provides targeted access to FlexConnect mediastream data for a particular client.
//
// Parameters:
//   - ctx: Context for request timeout and cancellation control
//   - clientMAC: Specific client MAC address to retrieve data for
//
// Returns:
//   - *model.McastOperFlexMediastreamClientSummary: Multicast data for the specified client
//   - error: Error if the operation fails or clientMAC is invalid
//
// Example:
//
//	service := mcast.NewService(client)
//	data, err := service.GetOperByClientMAC(ctx, "aa:bb:cc:dd:ee:ff")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("Client multicast data: %+v\n", data)
func (s Service) GetOperByClientMAC(
	ctx context.Context,
	clientMAC string,
) (*model.McastOperFlexMediastreamClientSummary, error) {
	if err := s.ValidateClient(); err != nil {
		return nil, err
	}
	if err := validation.ValidateAPMac(clientMAC); err != nil {
		return nil, err
	}
	endpoint := s.Client().RestconfBuilder().BuildPathQueryURL(
		routes.McastOperEndpoint, "flex-mediastream-client-summary", clientMAC)
	return core.Get[model.McastOperFlexMediastreamClientSummary](ctx, s.Client(), endpoint)
}

// GetOperByVlanIndex retrieves multicast operational data for a specific VLAN index.
// This function provides targeted access to VLAN Layer 2 multicast group data for a particular VLAN.
//
// Parameters:
//   - ctx: Context for request timeout and cancellation control
//   - vlanIndex: Specific VLAN index to retrieve data for
//
// Returns:
//   - *model.McastOperVlanL2MgidOp: Multicast data for the specified VLAN
//   - error: Error if the operation fails or vlanIndex is invalid
//
// Example:
//
//	service := mcast.NewService(client)
//	data, err := service.GetOperByVlanIndex(ctx, "100")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Printf("VLAN multicast data: %+v\n", data)
func (s Service) GetOperByVlanIndex(
	ctx context.Context,
	vlanIndex string,
) (*model.McastOperVlanL2MgidOp, error) {
	if err := s.ValidateClient(); err != nil {
		return nil, err
	}
	if err := validation.ValidateNonEmptyString(vlanIndex, "VLAN index"); err != nil {
		return nil, err
	}
	endpoint := s.Client().RestconfBuilder().BuildPathQueryURL(routes.McastOperEndpoint, "vlan-l2-mgid-op", vlanIndex)
	return core.Get[model.McastOperVlanL2MgidOp](ctx, s.Client(), endpoint)
}
