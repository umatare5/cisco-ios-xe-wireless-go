// Package wlan provides WLAN domain services feless Network Controller API.
package wlan

import (
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/service"
)

// Service provides comprehensive WLAN operations for Cisco IOS-XE Wireless LAN Controller.
//
// The WLAN service encompasses all aspects of wireless LAN management including:
//   - WLAN configuration retrieval and monitoring
//   - Operational data collection and analysis
//   - WLAN enable/disable control and status management
//   - Policy management and association
//   - Authentication and encryption configuration
//   - Real-time WLAN status and performance monitoring
//   - Policy tag operations through integrated sub-services
//
// All operations support both WLAN ID and profile name identification methods,
// providing flexibility for different network management workflows and integration scenarios.
type Service struct {
	// BaseService provides common service infrastructure
	service.BaseService
}

// NewService creates a new WLAN service instance with the provided client.
//
// The service requires a valid core.Client configured with appropriate
// controller credentials and network settings. The client manages
// authentication, request/response handling, and error processing
// for all WLAN-related RESTCONF operations.
//
// Parameters:
//   - client: Configured core.Client for RESTCONF API communication
//
// Returns:
//   - Service: New WLAN service instance ready for operations
//
// Example:
//
//	client := core.NewClient("192.168.1.1", "your-token")
//	wlanService := wlan.NewService(client)
func NewService(client *core.Client) Service {
	return Service{BaseService: service.NewBaseService(client)}
}

// PolicyTag returns a PolicyTagService for policy tag operations
func (s Service) PolicyTag() *PolicyTagService {
	return NewPolicyTagService(s.Client())
}
