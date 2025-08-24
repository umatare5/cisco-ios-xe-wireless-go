package mdns

import (
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/service"
)

// Service provides comprehensive Multicast DNS (mDNS) operations for Cisco IOS-XE Wireless LAN Controller.
//
// The MDNS service encompasses all aspects of Multicast DNS (mDNS) management including:
//   - Service discovery monitoring and analysis
//   - Bonjour protocol operational data
//   - Zero-configuration networking statistics
//   - WLAN-specific mDNS performance metrics
//   - Multicast DNS traffic analysis
//
// All operations support monitoring and management of multicast dns (mdns) functionality,
// providing flexibility for different network management workflows and integration scenarios.
type Service struct {
	// BaseService provides common service infrastructure
	service.BaseService
}

// NewService creates a new MDNS service instance with the provided client.
//
// The service requires a valid core.Client configured with appropriate
// controller credentials and network settings. The client manages
// authentication, request/response handling, and error processing
// for all MDNS-related RESTCONF operations.
//
// Parameters:
//   - client: Configured core.Client for RESTCONF API communication
//
// Returns:
//   - Service: New MDNS service instance ready for operations
//
// Example:
//
//	client := core.NewClient("192.168.1.1", "your-token")
//	mdnsService := mdns.NewService(client)
func NewService(client *core.Client) Service {
	return Service{BaseService: service.NewBaseService(client)}
}
