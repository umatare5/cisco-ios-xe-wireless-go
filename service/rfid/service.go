package rfid

import (
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/service"
)

// Service provides comprehensive Radio Frequency Identification (RFID) operations for Cisco IOS-XE Wireless LAN Controller.
//
// The RFID service encompasses all aspects of radio frequency identification (rfid) management including:
//   - RFID tag tracking and monitoring
//   - Asset management and inventory
//   - RFID reader configuration
//   - Tag detection and location services
//   - Real-time asset visibility
//
// All operations support monitoring and management of radio frequency identification (rfid) functionality,
// providing flexibility for different network management workflows and integration scenarios.
type Service struct {
	// BaseService provides common service infrastructure
	service.BaseService
}

// NewService creates a new RFID service instance with the provided client.
//
// The service requires a valid core.Client configured with appropriate
// controller credentials and network settings. The client manages
// authentication, request/response handling, and error processing
// for all RFID-related RESTCONF operations.
//
// Parameters:
//   - client: Configured core.Client for RESTCONF API communication
//
// Returns:
//   - Service: New RFID service instance ready for operations
//
// Example:
//
//	client := core.NewClient("192.168.1.1", "your-token")
//	rfidService := rfid.NewService(client)
func NewService(client *core.Client) Service {
	return Service{BaseService: service.NewBaseService(client)}
}
