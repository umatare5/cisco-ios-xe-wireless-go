// Package ap provides access point functionality for the Cisco IOS-XE Wireless Network Controller API.
//
// This package allows you to configure and monitor access points connected to a Cisco Catalyst 9800 Wireless LAN Controller.
// It provides methods for retrieving operational data, configuring access point settings, and managing access point administrative states.
//
// # Main Features
//
// - Access Point Configuration: GetCfg(), GetApTagsCfg(), GetApTagsByMAC(), GetTagSourcePriorityCfg()
// - Operational Data: GetOper(), GetOperCapwapData(), GetOperRadioStatus()
// - Operational Filtering: GetOperCapwapDataByWtpMac(), GetOperNameMacMapByWtpName(), GetOperRadioStatusByWtpMacSlot()
// - Global Operational Data: GetGlobalOper(), GetGlobalOperEwlcApStats()
// - Administrative Control: SetAPAdminState(), SetAPSlotAdminState(), RunAPReload()
// - Tag Management: SetAPTags()
//
// # Usage Example
//
//	// Create a client and access AP service
//	client := wnc.NewClient("controller.example.com", "your-token")
//	apService := client.AP()
//
//	// Get all access point configurations
//	cfg, err := apService.GetCfg(context.Background())
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// Get operational data
//	oper, err := apService.GetOper(context.Background())
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// Get specific AP tags by MAC address
//	tags, err := apService.GetApTagsByMAC(context.Background(), "28:ac:9e:bb:3c:80")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// Get CAPWAP data for specific AP
//	capwap, err := apService.GetOperCapwapDataByWtpMac(context.Background(), "28:ac:9e:bb:3c:80")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// Get radio status for specific AP and slot
//	radio, err := apService.GetOperRadioStatusByWtpMacSlot(context.Background(), "28:ac:9e:bb:3c:80", 0)
//	if err != nil {
//		log.Fatal(err)
//	}
//
// # KNOWN LIMITATIONS
//
// Based on testing with actual Cisco Catalyst 9800 controllers, the following YANG model endpoints
// are not implemented in the RESTCONF API and will return 404 errors:
//
// - live-data and all its filtering operations (by WTP name, MAC, IP, etc.)
// - name-mac-map endpoint
// - ap-join-stats-data and all its filtering operations
// - history endpoint
// - ap-cfg-entries endpoint (originally planned for configuration filtering)
//
// These limitations affect the following methods that have been removed:
// - GetOperByWtpName, GetOperByWtpMac, GetOperByApIPAddr, etc.
// - GetGlobalOperByWtpMac, GetGlobalOperByWtpName, etc.
// - GetCfgByMAC, GetCfgByPolicyTag, GetCfgBySiteTag, GetCfgByRfTag, etc.
//
// The available functionality includes:
// - Basic configuration and operational data retrieval
// - CAPWAP data access
// - Radio operational status
// - EWLC statistics
// - AP tag configuration by MAC address (working filter)
// - Administrative state management
// - AP reload operations
//
// # Error Handling
//
// Methods may return HTTP 404 errors for unavailable endpoints. This is expected behavior
// for endpoints not implemented in the target controller deployment.
//
// # Requirements
//
// - Cisco Catalyst 9800 Wireless LAN Controller
// - IOS-XE 17.12 or later
// - RESTCONF API access enabled
// - Valid authentication credentials
package ap
