package ble

import (
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/service"
)

// Service provides comprehensive Bluetooth Low Energy (BLE) operations for Cisco IOS-XE Wireless LAN Controller.
//
// The BLE service encompasses all aspects of BLE Locate and Track (LTX) management including:
//   - Operational data retrieval and monitoring
//   - Location tracking and analytics
//   - BLE beacon management
//   - Real-time tracking status and performance monitoring
//
// All operations support monitoring of BLE devices and location tracking services,
// providing flexibility for different IoT and location-based service workflows and integration scenarios.
type Service struct {
	// BaseService provides common service infrastructure
	service.BaseService
}

// NewService creates a new BLE service instance with the provided client.
//
// The service requires a valid core.Client configured with appropriate
// controller credentials and network settings. The client manages
// authentication, request/response handling, and error processing
// for all BLE-related RESTCONF operations.
//
// Parameters:
//   - client: Configured core.Client for RESTCONF API communication
//
// Returns:
//   - Service: New BLE service instance ready for operations
//
// Example:
//
//	client := core.NewClient("192.168.1.1", "your-token")
//	bleService := ble.NewService(client)
func NewService(client *core.Client) Service {
	return Service{BaseService: service.NewBaseService(client)}
}
