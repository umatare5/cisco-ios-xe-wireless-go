package rf

import (
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/service"
)

// Service provides comprehensive Radio Frequency (RF) Management operations for Cisco IOS-XE Wireless LAN Controller.
//
// The RF service encompasses all aspects of radio frequency (rf) management management including:
//   - RF profile configuration and management
//   - Radio frequency parameter control
//   - Power and channel optimization settings
//   - RF policy enforcement and monitoring
//   - Global RF configuration management
//
// All operations support monitoring and management of radio frequency (rf) management functionality,
// providing flexibility for different network management workflows and integration scenarios.
type Service struct {
	// BaseService provides common service infrastructure
	service.BaseService
}

// NewService creates a new RF service instance with the provided client.
//
// The service requires a valid core.Client configured with appropriate
// controller credentials and network settings. The client manages
// authentication, request/response handling, and error processing
// for all RF-related RESTCONF operations.
//
// Parameters:
//   - client: Configured core.Client for RESTCONF API communication
//
// Returns:
//   - Service: New RF service instance ready for operations
//
// Example:
//
//	client := core.NewClient("192.168.1.1", "your-token")
//	rfService := rf.NewService(client)
func NewService(client *core.Client) Service {
	return Service{BaseService: service.NewBaseService(client)}
}
