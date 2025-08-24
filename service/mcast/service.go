package mcast

import (
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/service"
)

// Service provides comprehensive Multicast Management operations for Cisco IOS-XE Wireless LAN Controller.
//
// The Multicast service encompasses all aspects of multicast management management including:
//   - Multicast routing and forwarding
//   - IGMP snooping and optimization
//   - Multicast VLAN management
//   - Video streaming optimization
//   - Bandwidth efficiency controls
//
// All operations support monitoring and management of multicast management functionality,
// providing flexibility for different network management workflows and integration scenarios.
type Service struct {
	// BaseService provides common service infrastructure
	service.BaseService
}

// NewService creates a new Multicast service instance with the provided client.
//
// The service requires a valid core.Client configured with appropriate
// controller credentials and network settings. The client manages
// authentication, request/response handling, and error processing
// for all Multicast-related RESTCONF operations.
//
// Parameters:
//   - client: Configured core.Client for RESTCONF API communication
//
// Returns:
//   - Service: New Multicast service instance ready for operations
//
// Example:
//
//	client := core.NewClient("192.168.1.1", "your-token")
//	mcastService := mcast.NewService(client)
func NewService(client *core.Client) Service {
	return Service{BaseService: service.NewBaseService(client)}
}
