package mesh

import (
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/service"
)

// Service provides comprehensive Wireless Mesh Networking operations for Cisco IOS-XE Wireless LAN Controller.
//
// The Mesh service encompasses all aspects of wireless mesh networking management including:
//   - Mesh access point configuration
//   - Backhaul optimization and management
//   - Mesh topology monitoring
//   - Root access point selection
//   - Outdoor mesh network deployment
//
// All operations support monitoring and management of wireless mesh networking functionality,
// providing flexibility for different network management workflows and integration scenarios.
type Service struct {
	// BaseService provides common service infrastructure
	service.BaseService
}

// NewService creates a new Mesh service instance with the provided client.
//
// The service requires a valid core.Client configured with appropriate
// controller credentials and network settings. The client manages
// authentication, request/response handling, and error processing
// for all Mesh-related RESTCONF operations.
//
// Parameters:
//   - client: Configured core.Client for RESTCONF API communication
//
// Returns:
//   - Service: New Mesh service instance ready for operations
//
// Example:
//
//	client := core.NewClient("192.168.1.1", "your-token")
//	meshService := mesh.NewService(client)
func NewService(client *core.Client) Service {
	return Service{BaseService: service.NewBaseService(client)}
}
