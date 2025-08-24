// Package controller provides wireless controller management functionality for the Cisco IOS-XE Wireless Network Controller API.
//
// This package allows you to perform administrative operations on a Cisco Catalyst 9800 Wireless LAN Controller.
// It provides essential system management capabilities including controller restart operations and system-wide maintenance commands.
//
// # Main Features
//
// - Controller Restart Operations: Reload(), ReloadWithReason()
// - System Administrative Commands
// - Maintenance Operation Management
// - Emergency System Control
//
// # Usage Example
//
//	// Create a client and access Controller service
//	client := wnc.NewClient("controller.example.com", "your-token")
//	controllerService := client.Controller()
//
//	// Restart controller with reason and force flag
//	err := controllerService.Reload(context.Background(), "Scheduled maintenance", false)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// Restart controller with reason only
//	err = controllerService.ReloadWithReason(context.Background(), "Memory leak detected")
//	if err != nil {
//		log.Fatal(err)
//	}
//
// # CRITICAL WARNING
//
// ⚠️  **DANGER: Controller operations will cause complete service interruption!**
//
// All controller management operations in this package are **DESTRUCTIVE** and will cause:
// - Complete wireless network service outage
// - All connected access points to go offline
// - All wireless clients to lose connectivity
// - Controller to become temporarily unavailable
//
// These operations should ONLY be performed:
// - During scheduled maintenance windows
// - In emergency situations requiring immediate controller restart
// - With proper coordination and advance notice to network users
//
// # Production Usage Guidelines
//
// 1. **Schedule Maintenance Windows**: Always perform these operations during planned maintenance
// 2. **Notify Stakeholders**: Inform all relevant parties before controller restart
// 3. **Verify Backups**: Ensure configuration backups are current before restart
// 4. **Test in Lab**: Validate restart procedures in non-production environment first
// 5. **Monitor Recovery**: Stay available to monitor controller restart and recovery
//
// # Integration Testing
//
// Controller restart operations are intentionally excluded from automated integration tests
// due to their destructive nature. To manually test controller functionality:
//
//  1. Set WNC_TEST_HOST, WNC_TEST_USERNAME, WNC_TEST_PASSWORD environment variables
//  2. Run: go test -run TestWNCReloadIntegration -tags integration
//  3. Be prepared for complete service outage during reload
//
// # Error Handling
//
// Methods return errors for:
// - Invalid authentication credentials
// - Network connectivity issues
// - Empty or invalid reload reasons
// - Controller state preventing restart
//
// Always check for errors and implement proper error handling when using these operations.
//
// # Requirements
//
// - Cisco Catalyst 9800 Wireless LAN Controller
// - IOS-XE 17.12 or later
// - RESTCONF API access enabled
// - Administrative privileges required
// - Valid authentication credentials
package controller
