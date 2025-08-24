package controller

import (
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/service"
)

// Service provides comprehensive Controller Management operations for Cisco IOS-XE Wireless LAN Controller.
//
// The Controller service encompasses all aspects of controller system management including:
//   - System reload and restart operations
//   - Controller administrative commands
//   - System-wide maintenance operations
//   - Emergency restart procedures
//   - Controller lifecycle management
//
// All operations support administrative control of the controller system,
// providing essential system management capabilities for network operations.
//
// WARNING: Controller management operations may cause service interruption.
// Use with caution in production environments.
type Service struct {
	// BaseService provides common service infrastructure
	service.BaseService
}

// NewService creates a new Controller service instance with the provided client.
//
// The service requires a valid core.Client configured with appropriate
// controller credentials and network settings. The client manages
// authentication, request/response handling, and error processing
// for all Controller management RESTCONF operations.
//
// Parameters:
//   - client: Configured core.Client for RESTCONF API communication
//
// Returns:
//   - Service: New Controller service instance ready for operations
//
// Example:
//
//	client := core.NewClient("192.168.1.1", "your-token")
//	controllerService := controller.NewService(client)
func NewService(client *core.Client) Service {
	return Service{BaseService: service.NewBaseService(client)}
}
