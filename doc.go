// Package wnc provides a unified Go client library for the Cisco Catalyst C9800 Wireless LAN Controller RESTCONF API.
//
// This library enables developers to communicate with Cisco Catalyst 9800 controllers in an idiomatic,
// robust, simple, and maintainable way using Go.
//
// # Architecture
//
// The library follows a three-layer architecture:
//
//   - Core Layer (internal/core): Handles HTTP communication, authentication, and low-level RESTCONF operations
//   - Domain Service Layer (service packages): Provides domain-specific APIs for each wireless controller feature
//   - Generated Type Layer (internal/model): Contains strongly-typed data structures for all API responses
//
// # Usage
//
// The unified client provides single-import access to all wireless controller functionality:
//
//	import wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
//
//	client, err := wnc.NewClient("controller.example.com", "your-access-token",
//		wnc.WithTimeout(30*time.Second),
//		wnc.WithInsecureSkipVerify(true))
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// Access different domain services
//	ctx := context.Background()
//	afcData, err := client.AFC().GetOper(ctx)
//	generalData, err := client.General().GetOper(ctx)
//	apData, err := client.AP().GetOper(ctx)
//
// # Domain Services
//
// The client provides access to the following domain services:
//
//   - AFC: Automated Frequency Coordination operations
//   - AP: Access Point management and operational data
//   - APF: Application Policy Framework configuration
//   - AWIPS: Advanced Weather Interactive Processing System
//   - BLE: Bluetooth Low Energy operations
//   - Client: Wireless client management and statistics
//   - CTS: Cisco TrustSec configuration
//   - Dot11: 802.11 wireless standard configuration
//   - Dot15: 802.15 standard configuration
//   - Fabric: SD-Access Fabric operations
//   - Flex: FlexConnect configuration and operations
//   - General: General controller configuration and status
//   - Geolocation: Location services and positioning
//   - Hyperlocation: High-precision location services
//   - LISP: Locator/Identifier Separation Protocol
//   - Location: Location services configuration
//   - Mcast: Multicast configuration and operations
//   - Mdns: Multicast DNS services
//   - Mesh: Wireless mesh networking
//   - Mobility: Client mobility management
//   - NMSP: Network Mobility Services Protocol
//   - Radio: Radio management and configuration
//   - RF: Radio Frequency management
//   - RFID: RFID tracking and management
//   - Rogue: Rogue access point detection
//   - RRM: Radio Resource Management
//   - Site: Site-based configuration
//   - WLAN: Wireless LAN configuration
//
// # Authentication
//
// The client uses Basic Authentication with a base64-encoded token. You can generate the token using:
//
//	echo -n "username:password" | base64
//
// # Error Handling
//
// The library provides structured error handling with specific error types:
//
//	if err != nil {
//		var apiErr *wnc.APIError
//		if errors.As(err, &apiErr) {
//			fmt.Printf("API Error (HTTP %d): %s\n", apiErr.StatusCode, apiErr.Message)
//		}
//	}
//
// # Compatibility
//
// This library is designed for Cisco IOS-XE 17.12.x running on Cisco Catalyst 9800 series controllers.
// It uses YANG models from the cisco-xe/17121 specification.
package wnc
