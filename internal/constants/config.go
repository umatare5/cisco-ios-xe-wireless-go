// Package constants provides shared constants used across the Cisco IOS-XE Wireless Go SDK.
package constants

// Test configuration constants
const (
	// DefaultTestMethods represents the standard number of methods for most service tests
	DefaultTestMethods = 10

	// StandardTestPhases represents the standard number of test phases (Unit, Mock, Error, Integration)
	StandardTestPhases = 4

	// DefaultTestGoroutines represents the default number of goroutines for concurrent testing
	DefaultTestGoroutines = 10

	// RogueServiceMethods represents the number of methods in the Rogue service
	RogueServiceMethods = 5

	// SingleMethodServices represents the number of methods for simple services
	SingleMethodServices = 1

	// WLANServiceMethods represents the number of methods in the WLAN service
	WLANServiceMethods = 6 // DefaultTestMethods - 4 unused methods
)

// HTTP Configuration
const (
	// DefaultHTTPTimeout defines the default HTTP timeout in seconds
	DefaultHTTPTimeout = 30

	// DefaultMaxRetries defines the default number of retry attempts
	DefaultMaxRetries = 3
)
