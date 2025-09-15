// Package service provides common service infrastructure for Cisco IOS-XE Wireless Network Controller API services.
package service

import (
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
)

// BaseService provides common service infrastructure that all domain services can embed.
//
// This structure standardizes service creation patterns and provides consistent
// error handling and validation across all service implementations.
type BaseService struct {
	// client provides the underlying HTTP client for RESTCONF API communication
	client *core.Client
}

// NewBaseService creates a new base service instance with the provided client.
//
// This constructor provides standardized validation and setup for all services
// that embed BaseService, ensuring consistent initialization patterns.
//
// Parameters:
//   - client: Configured core.Client for RESTCONF API communication
//
// Returns:
//   - BaseService: New base service instance ready for use by embedding services
func NewBaseService(client *core.Client) BaseService {
	return BaseService{client: client}
}

// Client returns the underlying HTTP client for advanced use cases.
//
// This method provides access to the configured client instance for services
// that need direct client access or custom operations beyond the standard patterns.
//
// Note: This method will return nil if the BaseService was not properly initialized
// with a valid client. Services should handle nil return values appropriately.
//
// Returns:
//   - *core.Client: The configured HTTP client instance, or nil if not initialized
func (b BaseService) Client() *core.Client {
	return b.client
}
