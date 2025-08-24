package apf

import (
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/service"
)

// Service provides comprehensive Application Policy Framework (APF) operations for Cisco IOS-XE Wireless LAN Controller.
//
// The APF service encompasses all aspects of application policy framework management including:
//   - Configuration retrieval and monitoring
//   - Application classification management
//   - Quality of Service (QoS) policy configuration
//   - Traffic handling and classification rules
//   - Real-time policy status and performance monitoring
//
// All operations support configuration management for application-aware wireless networking,
// providing flexibility for different network policy workflows and integration scenarios.
type Service struct {
	// BaseService provides common service infrastructure
	service.BaseService
}

// NewService creates a new APF service instance with the provided client.
//
// The service requires a valid core.Client configured with appropriate
// controller credentials and network settings. The client manages
// authentication, request/response handling, and error processing
// for all APF-related RESTCONF operations.
//
// Parameters:
//   - client: Configured core.Client for RESTCONF API communication
//
// Returns:
//   - Service: New APF service instance ready for operations
//
// Example:
//
//	client := core.NewClient("192.168.1.1", "your-token")
//	apfService := apf.NewService(client)
func NewService(client *core.Client) Service {
	return Service{BaseService: service.NewBaseService(client)}
}
