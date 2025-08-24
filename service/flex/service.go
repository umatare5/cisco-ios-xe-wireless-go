package flex

import (
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/service"
)

// Service provides comprehensive Cisco FlexConnect operations for Cisco IOS-XE Wireless LAN Controller.
//
// The FlexConnect service encompasses all aspects of cisco flexconnect management including:
//   - FlexConnect group and AP management
//   - Local switching and VLAN configuration
//   - Central and local authentication modes
//   - Backup RADIUS and local user database
//   - Branch office wireless optimization
//
// All operations support monitoring and management of cisco flexconnect functionality,
// providing flexibility for different network management workflows and integration scenarios.
type Service struct {
	// BaseService provides common service infrastructure
	service.BaseService
}

// NewService creates a new FlexConnect service instance with the provided client.
//
// The service requires a valid core.Client configured with appropriate
// controller credentials and network settings. The client manages
// authentication, request/response handling, and error processing
// for all FlexConnect-related RESTCONF operations.
//
// Parameters:
//   - client: Configured core.Client for RESTCONF API communication
//
// Returns:
//   - Service: New FlexConnect service instance ready for operations
//
// Example:
//
//	client := core.NewClient("192.168.1.1", "your-token")
//	flexService := flex.NewService(client)
func NewService(client *core.Client) Service {
	return Service{BaseService: service.NewBaseService(client)}
}
