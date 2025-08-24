package rogue

import (
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/service"
)

// Service provides comprehensive Rogue Detection and Mitigation operations for Cisco IOS-XE Wireless LAN Controller.
//
// The Rogue service encompasses all aspects of rogue detection and mitigation management including:
//   - Rogue access point detection
//   - Security threat identification
//   - Automatic containment policies
//   - Rogue classification and analysis
//   - Wireless security monitoring
//
// All operations support monitoring and management of rogue detection and mitigation functionality,
// providing flexibility for different network management workflows and integration scenarios.
type Service struct {
	// BaseService provides common service infrastructure
	service.BaseService
}

// NewService creates a new Rogue service instance with the provided client.
//
// The service requires a valid core.Client configured with appropriate
// controller credentials and network settings. The client manages
// authentication, request/response handling, and error processing
// for all Rogue-related RESTCONF operations.
//
// Parameters:
//   - client: Configured core.Client for RESTCONF API communication
//
// Returns:
//   - Service: New Rogue service instance ready for operations
//
// Example:
//
//	client := core.NewClient("192.168.1.1", "your-token")
//	rogueService := rogue.NewService(client)
func NewService(client *core.Client) Service {
	return Service{BaseService: service.NewBaseService(client)}
}
