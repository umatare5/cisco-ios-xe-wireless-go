// Package awips provides Automated Wireless Intrusion Prevention System (AWIPS) functionality for the Cisco IOS-XE Wireless Network Controller API.
//
// This package allows you to monitor and retrieve AWIPS data from a Cisco Catalyst 9800 Wireless LAN Controller.
// It provides methods for retrieving operational data and managing wireless security monitoring and intrusion detection.
//
// # Main Features
//
// - AWIPS Operational Data: GetOper()
// - AP-specific AWIPS Data: GetOperByApMac()
// - Download Status Monitoring: GetOperByApMacDownloadStatus()
// - Wireless security threat monitoring
// - Intrusion detection statistics
// - Real-time threat analysis
//
// # Usage Example
//
//	// Create a client and access AWIPS service
//	client := wnc.NewClient("controller.example.com", "your-token")
//	awipsService := client.AWIPS()
//
//	// Get all AWIPS operational data
//	oper, err := awipsService.GetOper(context.Background())
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// Get AWIPS data for specific AP
//	apData, err := awipsService.GetOperByApMac(context.Background(), "28:ac:9e:bb:3c:80")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// Get download status for specific AP
//	status, err := awipsService.GetOperByApMacDownloadStatus(context.Background(), "28:ac:9e:bb:3c:80")
//	if err != nil {
//		log.Fatal(err)
//	}
//
// # Requirements
//
// - Cisco Catalyst 9800 Wireless LAN Controller
// - IOS-XE 17.12 or later
// - RESTCONF API access enabled
// - Valid authentication credentials
package awips
