package ap

import (
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/service"
)

// Service provides comprehensive access point operations for Cisco IOS-XE Wireless LAN Controller.
//
// The AP service encompasses all aspects of access point management including:
//   - Configuration retrieval and monitoring
//   - Operational data collection and analysis
//   - Administrative state control (enable/disable)
//   - Radio interface management per slot
//   - Tag-based organization and policy assignment
//   - Maintenance operations (reload, diagnostics)
//   - Real-time status and performance monitoring
//
// All operations support both MAC address and AP name identification methods,
// providing flexibility for different network management workflows and integration scenarios.
type Service struct {
	// BaseService provides common service infrastructure
	service.BaseService
}

// NewService creates a new AP service instance with the provided client.
//
// The service requires a valid core.Client configured with appropriate
// controller credentials and network settings. The client manages
// authentication, request/response handling, and error processing
// for all AP-related RESTCONF operations.
//
// Parameters:
//   - client: Configured core.Client for RESTCONF API communication
//
// Returns:
//   - Service: New AP service instance ready for operations
//
// Example:
//
//	client := core.NewClient("192.168.1.1", "your-token")
//	apService := ap.NewService(client)
func NewService(client *core.Client) Service {
	return Service{BaseService: service.NewBaseService(client)}
}
