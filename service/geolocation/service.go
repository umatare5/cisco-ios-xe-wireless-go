package geolocation

import (
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/service"
)

// Service provides comprehensive Geolocation Tracking Service operations for Cisco IOS-XE Wireless LAN Controller.
//
// The Geolocation service encompasses all aspects of geolocation tracking service management including:
//   - Location tracking and positioning data
//   - Client device position monitoring
//   - Geofencing and location analytics
//   - Real-time positioning accuracy metrics
//   - Location-based service integration
//
// All operations support monitoring and management of geolocation tracking service functionality,
// providing flexibility for different network management workflows and integration scenarios.
type Service struct {
	// BaseService provides common service infrastructure
	service.BaseService
}

// NewService creates a new Geolocation service instance with the provided client.
//
// The service requires a valid core.Client configured with appropriate
// controller credentials and network settings. The client manages
// authentication, request/response handling, and error processing
// for all Geolocation-related RESTCONF operations.
//
// Parameters:
//   - client: Configured core.Client for RESTCONF API communication
//
// Returns:
//   - Service: New Geolocation service instance ready for operations
//
// Example:
//
//	client := core.NewClient("192.168.1.1", "your-token")
//	geolocationService := geolocation.NewService(client)
func NewService(client *core.Client) Service {
	return Service{BaseService: service.NewBaseService(client)}
}
