package rrm

import (
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/service"
)

// Service provides comprehensive Radio Resource Management (RRM) operations for Cisco IOS-XE Wireless LAN Controller.
//
// The RRM service encompasses all aspects of radio resource management (rrm) management including:
//   - Dynamic channel assignment and optimization
//   - Power control and interference mitigation
//   - Coverage hole detection and remediation
//   - Load balancing and capacity planning
//   - Real-time RF environment monitoring
//
// All operations support monitoring and management of radio resource management (rrm) functionality,
// providing flexibility for different network management workflows and integration scenarios.
type Service struct {
	// BaseService provides common service infrastructure
	service.BaseService
}

// NewService creates a new RRM service instance with the provided client.
//
// The service requires a valid core.Client configured with appropriate
// controller credentials and network settings. The client manages
// authentication, request/response handling, and error processing
// for all RRM-related RESTCONF operations.
//
// Parameters:
//   - client: Configured core.Client for RESTCONF API communication
//
// Returns:
//   - Service: New RRM service instance ready for operations
//
// Example:
//
//	client := core.NewClient("192.168.1.1", "your-token")
//	rrmService := rrm.NewService(client)
func NewService(client *core.Client) Service {
	return Service{BaseService: service.NewBaseService(client)}
}
