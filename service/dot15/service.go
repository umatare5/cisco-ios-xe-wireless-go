package dot15

import (
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/service"
)

// Service provides comprehensive IEEE 802.15 Bluetooth Configuration operations for Cisco IOS-XE Wireless LAN Controller.
//
// The 802.15 service encompasses all aspects of ieee 802.15 bluetooth configuration management including:
//   - 802.15 Bluetooth protocol configuration
//   - Personal area network (PAN) settings
//   - Bluetooth coexistence management
//   - Low energy protocol optimization
//   - Device pairing and security policies
//
// All operations support monitoring and management of ieee 802.15 bluetooth configuration functionality,
// providing flexibility for different network management workflows and integration scenarios.
type Service struct {
	// BaseService provides common service infrastructure
	service.BaseService
}

// NewService creates a new 802.15 service instance with the provided client.
//
// The service requires a valid core.Client configured with appropriate
// controller credentials and network settings. The client manages
// authentication, request/response handling, and error processing
// for all 802.15-related RESTCONF operations.
//
// Parameters:
//   - client: Configured core.Client for RESTCONF API communication
//
// Returns:
//   - Service: New 802.15 service instance ready for operations
//
// Example:
//
//	client := core.NewClient("192.168.1.1", "your-token")
//	dot15Service := dot15.NewService(client)
func NewService(client *core.Client) Service {
	return Service{BaseService: service.NewBaseService(client)}
}
