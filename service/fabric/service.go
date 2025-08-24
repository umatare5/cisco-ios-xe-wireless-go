package fabric

import (
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/service"
)

// Service provides comprehensive Cisco SD-Access Fabric operations for Cisco IOS-XE Wireless LAN Controller.
//
// The Fabric service encompasses all aspects of cisco sd-access fabric management including:
//   - Software-defined access fabric configuration
//   - Overlay network management
//   - Border node and control plane settings
//   - Fabric-enabled SSID management
//   - Integration with DNA Center policies
//
// All operations support monitoring and management of cisco sd-access fabric functionality,
// providing flexibility for different network management workflows and integration scenarios.
type Service struct {
	// BaseService provides common service infrastructure
	service.BaseService
}

// NewService creates a new Fabric service instance with the provided client.
//
// The service requires a valid core.Client configured with appropriate
// controller credentials and network settings. The client manages
// authentication, request/response handling, and error processing
// for all Fabric-related RESTCONF operations.
//
// Parameters:
//   - client: Configured core.Client for RESTCONF API communication
//
// Returns:
//   - Service: New Fabric service instance ready for operations
//
// Example:
//
//	client := core.NewClient("192.168.1.1", "your-token")
//	fabricService := fabric.NewService(client)
func NewService(client *core.Client) Service {
	return Service{BaseService: service.NewBaseService(client)}
}
