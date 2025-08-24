// Package afc provides Automated Frequency Coordination (AFC) functionality for the Cisco IOS-XE Wireless Network Controller API.
//
// This package allows you to monitor AFC (Automated Frequency Coordination) operations for 6 GHz spectrum management
// on a Cisco Catalyst 9800 Wireless LAN Controller. AFC provides dynamic spectrum coordination with external AFC databases
// to ensure optimal spectrum utilization and interference avoidance in the 6 GHz band.
//
// # Main Features
//
// - AFC Operational Data: GetOper(), GetOperAPResp()
// - AFC Cloud Operations: GetOperCloudOper(), GetOperCloudStats()
// - Operational Filtering: GetOperByApMac(), GetOperByApMacAndSlot(), GetOperByApMacAndRequestID()
// - Specialized Access: GetOperBySlot(), GetOperByRequestID()
//
// # Usage Example
//
//	// Create a client and access AFC service
//	client := wnc.NewClient("controller.example.com", "your-token")
//	afcService := client.AFC()
//
//	// Get all AFC operational data
//	oper, err := afcService.GetOper(context.Background())
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// Get AFC response data for specific AP
//	apResp, err := afcService.GetOperByApMac(context.Background(), "28:ac:9e:bb:3c:80")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// Get AFC cloud operational data
//	cloudOper, err := afcService.GetOperCloudOper(context.Background())
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// Get AFC cloud statistics
//	cloudStats, err := afcService.GetOperCloudStats(context.Background())
//	if err != nil {
//		log.Fatal(err)
//	}
//
// # KNOWN LIMITATIONS
//
// AFC operations are read-only as they are managed by external AFC systems and the controller's
// automatic coordination processes. Configuration of AFC parameters is handled through:
//
// - External AFC database registration
// - Controller system-level AFC configuration
// - 6 GHz radio profile settings
//
// The available functionality includes:
// - AFC operational status monitoring
// - Per-AP AFC response tracking
// - AFC cloud service statistics
// - Request/response correlation data
// - Spectrum coordination status
//
// Set/Create/Update/Delete operations are not available for AFC service as it operates
// as a monitoring and coordination service with external AFC providers.
//
// # Error Handling
//
// Methods may return HTTP 404 errors for unavailable AFC data when:
// - No AFC-capable APs are present
// - AFC service is not configured
// - 6 GHz operation is not enabled
//
// # Requirements
//
// - Cisco Catalyst 9800 Wireless LAN Controller
// - IOS-XE 17.12 or later
// - RESTCONF API access enabled
// - 6 GHz AFC service configuration
// - Valid authentication credentials
package afc
