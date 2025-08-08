// Package constants provides shared constants used across the Cisco IOS-XE Wireless Go SDK.
package constants

import (
	"time"
)

// YANG Model prefixes
const (
	// YANGModelPrefix is the standard prefix for wireless-related YANG models
	YANGModelPrefix = "Cisco-IOS-XE-wireless-"

	// YANGModelPrefixAccess is the prefix for access-related YANG models
	YANGModelPrefixAccess = "Cisco-IOS-XE-access-"

	// YANGModelPrefixSite is the prefix for site-related YANG models
	YANGModelPrefixSite = "Cisco-IOS-XE-site-"
)

// YANG Module Suffixes
const (
	// CfgSuffix represents configuration module suffix
	CfgSuffix = "-cfg"

	// OperSuffix represents operational module suffix
	OperSuffix = "-oper"

	// CfgDataSuffix represents configuration data container suffix
	CfgDataSuffix = "-cfg-data"

	// OperDataSuffix represents operational data container suffix
	OperDataSuffix = "-oper-data"
)

// YANG Module Names (kept minimal for tests)
const (
	YangModuleAFC           = "afc"
	YangModuleAP            = "ap"
	YangModuleAPF           = "apf"
	YangModuleAWIPS         = "awips"
	YangModuleBLE           = "ble"
	YangModuleClient        = "client"
	YangModuleCTS           = "cts"
	YangModuleDot11         = "dot11"
	YangModuleDot15         = "dot15"
	YangModuleFabric        = "fabric"
	YangModuleFlex          = "flex"
	YangModuleGeneral       = "general"
	YangModuleGeolocation   = "geolocation"
	YangModuleHyperlocation = "hyperlocation"
	YangModuleLISP          = "lisp"
	YangModuleLocation      = "location"
	YangModuleMcast         = "mcast"
	YangModuleMDNS          = "mdns"
	YangModuleMesh          = "mesh"
	YangModuleMobility      = "mobility"
	YangModuleNMSP          = "nmsp"
	YangModuleRadio         = "radio"
	YangModuleRF            = "rf"
	YangModuleRFID          = "rfid"
	YangModuleRogue         = "rogue"
	YangModuleRRM           = "rrm"
	YangModuleSite          = "site"
	YangModuleWLAN          = "wlan"
)

// AP Service specific YANG models
const (
	APCfgModel        = YANGModelPrefix + "ap-cfg"
	APOperModel       = YANGModelPrefix + "access-point-oper"
	APGlobalOperModel = YANGModelPrefix + "ap-global-oper"
)

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
	WLANServiceMethods = 6
)

// HTTP Configuration
const (
	// DefaultHTTPTimeout defines the default HTTP timeout in seconds
	DefaultHTTPTimeout = 30

	// DefaultMaxRetries defines the default number of retry attempts
	DefaultMaxRetries = 3
)

// Network and protocol constants
const (
	// NetworkTimeoutSeconds defines timeout in seconds for backward compatibility
	NetworkTimeoutSeconds = 60

	// HTTPSScheme defines the HTTPS URL scheme
	HTTPSScheme = "https"

	// HTTPScheme defines the HTTP URL scheme
	HTTPScheme = "http"

	// URLSchemeSeparator defines the scheme separator in URLs
	URLSchemeSeparator = "://"
)

// Timeout duration constants
const (
	// QuickTimeoutSeconds for fast operations
	QuickTimeoutSeconds = 5

	// StandardTimeoutSeconds for normal operations
	StandardTimeoutSeconds = NetworkTimeoutSeconds

	// ExtendedTimeoutSeconds for longer operations
	ExtendedTimeoutSeconds = 90

	// ComprehensiveTimeoutSeconds for test suites
	ComprehensiveTimeoutSeconds = 150

	// MicroTimeoutMicroseconds for immediate cancellation tests
	MicroTimeoutMicroseconds = 1
)

// Timeout variation constants
const (
	// QuickTimeout for fast operations
	QuickTimeout = QuickTimeoutSeconds * time.Second

	// StandardTimeout for normal operations
	StandardTimeout = StandardTimeoutSeconds * time.Second

	// ExtendedTimeout for longer operations
	ExtendedTimeout = ExtendedTimeoutSeconds * time.Second

	// ComprehensiveTimeout for test suites
	ComprehensiveTimeout = ComprehensiveTimeoutSeconds * time.Second

	// MicroTimeout for immediate cancellation tests
	MicroTimeout = MicroTimeoutMicroseconds * time.Microsecond
)

// Environment variable names
const (
	// EnvVarController is the environment variable name for controller address
	EnvVarController = "WNC_CONTROLLER"

	// EnvVarAccessToken is the environment variable name for access token
	EnvVarAccessToken = "WNC_ACCESS_TOKEN"
)

// Default values
// (Intentionally no implicit default controller; users must supply one explicitly.)

// Documentation and example constants
const (
	// ExampleControllerIPAddress is used in documentation examples
	ExampleControllerIPAddress = "192.168.1.100"

	// ExampleControllerHostname is used in documentation examples
	ExampleControllerHostname = "core.example.local"

	// ExampleAccessToken is used in documentation examples
	ExampleAccessToken = "your-token"

	// ExampleTimeoutSeconds is used in documentation examples
	ExampleTimeoutSeconds = 20

	// ExampleTestHostname is used in test examples
	ExampleTestHostname = "test.local"
)

// Test constants
const (
	// TestAccessTokenValue is a base64 encoded test token for "test:test"
	TestAccessTokenValue = "dGVzdDp0ZXN0" //nolint:gosec // Test credential, not production

	// TestTimestamp defines a standard test timestamp
	TestTimestamp = "2024-01-01T00:00:00.000Z"

	// TestAPName defines a standard test access point name
	TestAPName = "test-ap-01"
)

// BuildYangModulePath constructs a YANG module path using the common pattern
func BuildYangModulePath(module, moduleType string) string {
	return YANGModelPrefix + module + "-" + moduleType + ":" + module + "-" + moduleType + "-data"
}

// BuildWirelessYangModule constructs a wireless YANG module name
func BuildWirelessYangModule(module, moduleType string) string {
	return YANGModelPrefix + module + "-" + moduleType
}

// BuildYangEndpoint constructs a YANG endpoint path
func BuildYangEndpoint(module, moduleType, endpoint string) string {
	basePath := BuildYangModulePath(module, moduleType)
	if endpoint == "" {
		return basePath
	}
	return basePath + "/" + endpoint
}

// BuildAPCfgPath builds AP configuration paths
func BuildAPCfgPath(endpoint string) string {
	if endpoint == "" {
		return APCfgModel + ":" + YangModuleAP + CfgDataSuffix
	}
	return APCfgModel + ":" + YangModuleAP + CfgDataSuffix + "/" + endpoint
}

// BuildAPOperPath builds AP operational paths
func BuildAPOperPath(endpoint string) string {
	basePath := "/" + APOperModel + ":" + "access-point" + OperDataSuffix
	if endpoint == "" {
		return basePath
	}
	return basePath + "/" + endpoint
}

// BuildAPGlobalOperPath builds AP global operational paths
func BuildAPGlobalOperPath(endpoint string) string {
	basePath := APGlobalOperModel + ":" + YangModuleAP + "-global" + OperDataSuffix
	if endpoint == "" {
		return basePath
	}
	return basePath + "/" + endpoint
}
