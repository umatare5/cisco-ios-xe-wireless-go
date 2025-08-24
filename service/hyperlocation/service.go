package hyperlocation

import (
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/service"
)

// Service provides comprehensive Cisco Hyperlocation operations for Cisco IOS-XE Wireless LAN Controller.
//
// The Hyperlocation service encompasses all aspects of cisco hyperlocation management including:
//   - High-precision indoor positioning
//   - Real-time location analytics
//   - Asset tracking and monitoring
//   - Location-based automation triggers
//   - Advanced positioning algorithms
//
// All operations support monitoring and management of cisco hyperlocation functionality,
// providing flexibility for different network management workflows and integration scenarios.
type Service struct {
	// BaseService provides common service infrastructure
	service.BaseService
}

// NewService creates a new Hyperlocation service instance with the provided client.
//
// The service requires a valid core.Client configured with appropriate
// controller credentials and network settings. The client manages
// authentication, request/response handling, and error processing
// for all Hyperlocation-related RESTCONF operations.
//
// Parameters:
//   - client: Configured core.Client for RESTCONF API communication
//
// Returns:
//   - Service: New Hyperlocation service instance ready for operations
//
// Example:
//
//	client := core.NewClient("192.168.1.1", "your-token")
//	hyperlocationService := hyperlocation.NewService(client)
func NewService(client *core.Client) Service {
	return Service{BaseService: service.NewBaseService(client)}
}
