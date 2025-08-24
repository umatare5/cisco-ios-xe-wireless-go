package nmsp

import (
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/service"
)

// Service provides comprehensive Network Mobility Services Protocol (NMSP) operations for Cisco IOS-XE Wireless LAN Controller.
//
// The NMSP service encompasses all aspects of network mobility services protocol (nmsp) management including:
//   - Location services protocol management
//   - Third-party application integration
//   - Real-time location streaming
//   - NMSP client configuration and monitoring
//   - Location data export and analytics
//
// All operations support monitoring and management of network mobility services protocol (nmsp) functionality,
// providing flexibility for different network management workflows and integration scenarios.
type Service struct {
	// BaseService provides common service infrastructure
	service.BaseService
}

// NewService creates a new NMSP service instance with the provided client.
//
// The service requires a valid core.Client configured with appropriate
// controller credentials and network settings. The client manages
// authentication, request/response handling, and error processing
// for all NMSP-related RESTCONF operations.
//
// Parameters:
//   - client: Configured core.Client for RESTCONF API communication
//
// Returns:
//   - Service: New NMSP service instance ready for operations
//
// Example:
//
//	client := core.NewClient("192.168.1.1", "your-token")
//	nmspService := nmsp.NewService(client)
func NewService(client *core.Client) Service {
	return Service{BaseService: service.NewBaseService(client)}
}
