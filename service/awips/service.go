package awips

import (
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/service"
)

// Service provides comprehensive Automated Wireless IPS (AWIPS) operations for Cisco IOS-XE Wireless LAN Controller.
//
// The AWIPS service encompasses all aspects of automated wireless intrusion prevention system management including:
//   - Operational data retrieval and monitoring
//   - Threat detection and analysis
//   - Security policy enforcement
//   - Real-time threat status and performance monitoring
//
// All operations support monitoring of wireless security threats and intrusion detection,
// providing flexibility for different security management workflows and integration scenarios.
type Service struct {
	// BaseService provides common service infrastructure
	service.BaseService
}

// NewService creates a new AWIPS service instance with the provided client.
//
// The service requires a valid core.Client configured with appropriate
// controller credentials and network settings. The client manages
// authentication, request/response handling, and error processing
// for all AWIPS-related RESTCONF operations.
//
// Parameters:
//   - client: Configured core.Client for RESTCONF API communication
//
// Returns:
//   - Service: New AWIPS service instance ready for operations
//
// Example:
//
//	client := core.NewClient("192.168.1.1", "your-token")
//	awipsService := awips.NewService(client)
func NewService(client *core.Client) Service {
	return Service{BaseService: service.NewBaseService(client)}
}
