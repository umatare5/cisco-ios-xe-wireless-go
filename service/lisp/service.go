package lisp

import (
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/service"
)

// Service provides comprehensive Locator/ID Separation Protocol (LISP) operations for Cisco IOS-XE Wireless LAN Controller.
//
// The LISP service encompasses all aspects of locator/id separation protocol (lisp) management including:
//   - LISP mobility and location services
//   - Endpoint identifier (EID) management
//   - Routing locator (RLOC) configuration
//   - Map server and resolver settings
//   - Network virtualization and scalability
//
// All operations support monitoring and management of locator/id separation protocol (lisp) functionality,
// providing flexibility for different network management workflows and integration scenarios.
type Service struct {
	// BaseService provides common service infrastructure
	service.BaseService
}

// NewService creates a new LISP service instance with the provided client.
//
// The service requires a valid core.Client configured with appropriate
// controller credentials and network settings. The client manages
// authentication, request/response handling, and error processing
// for all LISP-related RESTCONF operations.
//
// Parameters:
//   - client: Configured core.Client for RESTCONF API communication
//
// Returns:
//   - Service: New LISP service instance ready for operations
//
// Example:
//
//	client := core.NewClient("192.168.1.1", "your-token")
//	lispService := lisp.NewService(client)
func NewService(client *core.Client) Service {
	return Service{BaseService: service.NewBaseService(client)}
}
