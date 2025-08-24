package afc

import (
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/service"
)

// Service provides comprehensive AFC (Automated Frequency Coordination) operations for Cisco IOS-XE Wireless LAN Controller.
//
// The AFC service encompasses all aspects of 6 GHz spectrum coordination including:
//   - AFC operational data monitoring and retrieval
//   - Per-AP AFC response tracking and analysis
//   - AFC cloud service interaction and statistics
//   - Real-time spectrum coordination status monitoring
//   - Request/response correlation and performance metrics
//
// AFC operations are read-only as spectrum coordination is managed automatically
// by external AFC databases and controller coordination processes.
// All operations support MAC address-based filtering for detailed AP-specific analysis.
type Service struct {
	// BaseService provides common service infrastructure
	service.BaseService
}

// NewService creates a new AFC service instance with the provided client.
//
// The service requires a valid core.Client configured with appropriate
// controller credentials and network settings. The client manages
// authentication, request/response handling, and error processing
// for all AFC-related RESTCONF operations.
//
// Parameters:
//   - client: Configured core.Client for RESTCONF API communication
//
// Returns:
//   - Service: New AFC service instance ready for operations
//
// Example:
//
//	client := core.NewClient("192.168.1.1", "your-token")
//	afcService := afc.NewService(client)
func NewService(client *core.Client) Service {
	return Service{BaseService: service.NewBaseService(client)}
}
