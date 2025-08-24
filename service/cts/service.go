package cts

import (
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/service"
)

// Service provides comprehensive Cisco TrustSec (CTS) operations for Cisco IOS-XE Wireless LAN Controller.
//
// The CTS service encompasses all aspects of cisco trustsec (cts) management including:
//   - Security group tag (SGT) management
//   - Network segmentation and access control
//   - Security policy enforcement
//   - TrustSec configuration and monitoring
//   - Identity-based network security
//
// All operations support monitoring and management of cisco trustsec (cts) functionality,
// providing flexibility for different network management workflows and integration scenarios.
type Service struct {
	// BaseService provides common service infrastructure
	service.BaseService
}

// NewService creates a new CTS service instance with the provided client.
//
// The service requires a valid core.Client configured with appropriate
// controller credentials and network settings. The client manages
// authentication, request/response handling, and error processing
// for all CTS-related RESTCONF operations.
//
// Parameters:
//   - client: Configured core.Client for RESTCONF API communication
//
// Returns:
//   - Service: New CTS service instance ready for operations
//
// Example:
//
//	client := core.NewClient("192.168.1.1", "your-token")
//	ctsService := cts.NewService(client)
func NewService(client *core.Client) Service {
	return Service{BaseService: service.NewBaseService(client)}
}
