package wnc

import (
	"time"
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

// HTTP and API related constants
const (
	// DefaultTimeout is the default timeout for API requests
	DefaultTimeout = NetworkTimeoutSeconds * time.Second
)

// Timeout duration constants in seconds for readability
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

	// StandardTimeout for normal operations (same as DefaultTimeout for compatibility)
	StandardTimeout = DefaultTimeout

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
const (
	// DefaultController is the default controller hostname
	DefaultController = "wnc1.example.internal"
)

// Documentation and example constants
const (
	// ExampleControllerIPAddress is used in documentation examples
	ExampleControllerIPAddress = "192.168.1.100"

	// ExampleControllerHostname is used in documentation examples
	ExampleControllerHostname = "wnc1.example.internal"

	// ExampleAccessToken is used in documentation examples. A real token is
	// base64("username:password") for HTTP Basic; see README for how to build one.
	ExampleAccessToken = "test-token-123"

	// ExampleTimeoutSeconds is used in documentation examples
	ExampleTimeoutSeconds = 20

	// ExampleTestHostname is used in test examples
	ExampleTestHostname = "wnc1.example.internal"
)

// Test constants
const (
	// TestAccessTokenValue is the access token used by every unit test fixture.
	// The controller expects base64("user:password") for HTTP Basic auth, but no
	// unit test decodes or transmits this value, so a readable placeholder is used.
	TestAccessTokenValue = "test-token-123"

	// TestTimestamp defines a standard test timestamp
	TestTimestamp = "2024-01-01T00:00:00.000Z"

	// TestAPName defines a standard test access point name.
	// Fixture access points are named TEST-APnn; the nn suffix matches the last
	// octet of the AP radio MAC address used in the same fixture.
	TestAPName = "TEST-AP01"
)
