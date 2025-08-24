package client

import (
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/service"
)

// Service provides comprehensive wireless client operations for Cisco IOS-XE Wireless LAN Controller.
//
// The Client service encompasses all aspects of wireless client management including:
//   - Operational data retrieval and monitoring
//   - Client connection and session management
//   - Traffic statistics and analytics
//   - Mobility management and roaming
//   - Policy enforcement and compliance
//   - Real-time client status and performance monitoring
//
// All operations support monitoring of wireless clients and their connectivity patterns,
// providing flexibility for different client management workflows and integration scenarios.
type Service struct {
	// BaseService provides common service infrastructure
	service.BaseService
}

// NewService creates a new Client service instance with the provided client.
//
// The service requires a valid core.Client configured with appropriate
// controller credentials and network settings. The client manages
// authentication, request/response handling, and error processing
// for all Client-related RESTCONF operations.
//
// Parameters:
//   - client: Configured core.Client for RESTCONF API communication
//
// Returns:
//   - Service: New Client service instance ready for operations
//
// Example:
//
//	client := core.NewClient("192.168.1.1", "your-token")
//	clientService := client.Client()
func NewService(client *core.Client) Service {
	return Service{BaseService: service.NewBaseService(client)}
}
