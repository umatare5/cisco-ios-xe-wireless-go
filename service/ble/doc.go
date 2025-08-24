// Package ble provides Bluetooth Low Energy (BLE) functionality for the Cisco IOS-XE Wireless Network Controller API.
//
// This package allows you to monitor and retrieve BLE Location Transmit (LTX) data from a Cisco Catalyst 9800 Wireless LAN Controller.
// It provides methods for retrieving operational data and managing BLE beacon monitoring, antenna management, and proximity analytics.
//
// # Main Features
//
// - BLE LTX Operational Data: GetOper()
// - AP-specific BLE Data: GetOperByApMac()
// - Antenna-specific Data: GetOperByApMacSlotAntenna()
// - BLE beacon monitoring
// - Location transmit analytics
// - Real-time proximity analysis
//
// # Usage Example
//
//	// Create a client and access BLE service
//	client := wnc.NewClient("controller.example.com", "your-token")
//	bleService := client.BLE()
//
//	// Get all BLE LTX operational data
//	oper, err := bleService.GetOper(context.Background())
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// Get BLE data for specific AP
//	apData, err := bleService.GetOperByApMac(context.Background(), "28:ac:9e:bb:3c:80")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// Get antenna data for specific AP, slot, and antenna
//	antennaData, err := bleService.GetOperByApMacSlotAntenna(context.Background(), "28:ac:9e:bb:3c:80", 0, 1)
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
package ble
