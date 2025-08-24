package site

import (
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/service"
)

// Service provides comprehensive site configuration operations for Cisco IOS-XE Wireless LAN Controller.
//
// The Site service encompasses all aspects of site management including:
//   - Site configuration retrieval and monitoring
//   - AP configuration profile management
//   - Site tag configuration and assignment
//   - AP packet capture and trace profile management
//   - Site-specific operational data collection
//
// All operations support various filtering methods including by profile name,
// site tag name, and other site-specific identifiers for flexible network
// management workflows and integration scenarios.
type Service struct {
	// BaseService provides common service infrastructure
	service.BaseService
}

// NewService creates a new Site service instance with the provided client.
//
// The service requires a valid core.Client configured with appropriate
// controller credentials and network settings. The client manages
// authentication, request/response handling, and error processing
// for all site-related RESTCONF operations.
//
// Parameters:
//   - client: Configured core.Client for RESTCONF API communication
//
// Returns:
//   - Service: New Site service instance ready for operations
//
// Example:
//
//	client := core.NewClient("192.168.1.1", "your-token")
//	siteService := site.NewService(client)
func NewService(client *core.Client) Service {
	return Service{BaseService: service.NewBaseService(client)}
}
