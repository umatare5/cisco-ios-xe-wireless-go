package dot11

import (
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/service"
)

// Service provides comprehensive IEEE 802.11 Wireless Configuration operations for Cisco IOS-XE Wireless LAN Controller.
//
// The 802.11 service encompasses all aspects of ieee 802.11 wireless configuration management including:
//   - 802.11 protocol configuration management
//   - Wireless standard parameters control
//   - Protocol compliance and optimization
//   - Radio transmission settings
//   - Standards-based configuration enforcement
//
// All operations support monitoring and management of ieee 802.11 wireless configuration functionality,
// providing flexibility for different network management workflows and integration scenarios.
type Service struct {
	// BaseService provides common service infrastructure
	service.BaseService
}

// NewService creates a new 802.11 service instance with the provided client.
//
// The service requires a valid core.Client configured with appropriate
// controller credentials and network settings. The client manages
// authentication, request/response handling, and error processing
// for all 802.11-related RESTCONF operations.
//
// Parameters:
//   - client: Configured core.Client for RESTCONF API communication
//
// Returns:
//   - Service: New 802.11 service instance ready for operations
//
// Example:
//
//	client := core.NewClient("192.168.1.1", "your-token")
//	dot11Service := dot11.NewService(client)
func NewService(client *core.Client) Service {
	return Service{BaseService: service.NewBaseService(client)}
}
