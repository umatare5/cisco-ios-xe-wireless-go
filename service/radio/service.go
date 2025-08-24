package radio

import (
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/service"
)

// Service provides comprehensive Radio Configuration Management operations for Cisco IOS-XE Wireless LAN Controller.
//
// The Radio service encompasses all aspects of radio configuration management management including:
//   - Radio interface configuration and control
//   - Antenna settings and RF optimization
//   - Channel and power management
//   - Radio profile assignment
//   - Per-radio performance monitoring
//
// All operations support monitoring and management of radio configuration management functionality,
// providing flexibility for different network management workflows and integration scenarios.
type Service struct {
	// BaseService provides common service infrastructure
	service.BaseService
}

// NewService creates a new Radio service instance with the provided client.
//
// The service requires a valid core.Client configured with appropriate
// controller credentials and network settings. The client manages
// authentication, request/response handling, and error processing
// for all Radio-related RESTCONF operations.
//
// Parameters:
//   - client: Configured core.Client for RESTCONF API communication
//
// Returns:
//   - Service: New Radio service instance ready for operations
//
// Example:
//
//	client := core.NewClient("192.168.1.1", "your-token")
//	radioService := radio.NewService(client)
func NewService(client *core.Client) Service {
	return Service{BaseService: service.NewBaseService(client)}
}
