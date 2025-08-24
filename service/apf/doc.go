// Package apf provides Application Policy Framework (APF) functionality for the Cisco IOS-XE Wireless Network Controller API.
//
// This package allows you to configure and monitor Application Policy Framework settings on a Cisco Catalyst 9800 Wireless LAN Controller.
// It provides methods for retrieving APF configuration data and managing application-aware wireless networking policies.
//
// # Main Features
//
// - APF Configuration: GetCfg()
// - Application classification management
// - Quality of Service (QoS) policy configuration
// - Traffic handling and classification rules
//
// # Usage Example
//
//	// Create a client and access APF service
//	client := wnc.NewClient("controller.example.com", "your-token")
//	apfService := client.APF()
//
//	// Get APF configuration
//	cfg, err := apfService.GetCfg(context.Background())
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
package apf
