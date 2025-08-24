package mobility

import (
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/service"
)

// Service provides comprehensive Client Mobility Management operations for Cisco IOS-XE Wireless LAN Controller.
//
// The Mobility service encompasses all aspects of client mobility management management including:
//   - Inter-controller mobility configuration
//   - Roaming optimization and fast transition
//   - Mobility group management
//   - Client load balancing across controllers
//   - Seamless handoff and session continuity
//
// All operations support monitoring and management of client mobility management functionality,
// providing flexibility for different network management workflows and integration scenarios.
type Service struct {
	// BaseService provides common service infrastructure
	service.BaseService
}

// NewService creates a new Mobility service instance with the provided client.
//
// The service requires a valid core.Client configured with appropriate
// controller credentials and network settings. The client manages
// authentication, request/response handling, and error processing
// for all Mobility-related RESTCONF operations.
//
// Parameters:
//   - client: Configured core.Client for RESTCONF API communication
//
// Returns:
//   - Service: New Mobility service instance ready for operations
//
// Example:
//
//	client := core.NewClient("192.168.1.1", "your-token")
//	mobilityService := mobility.NewService(client)
func NewService(client *core.Client) Service {
	return Service{BaseService: service.NewBaseService(client)}
}
