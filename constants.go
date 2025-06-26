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
	ExampleControllerHostname = "wnc.example.local"

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
	TestAccessTokenValue = "dGVzdDp0ZXN0"

	// TestTimestamp defines a standard test timestamp
	TestTimestamp = "2024-01-01T00:00:00.000Z"

	// TestAPName defines a standard test access point name
	TestAPName = "test-ap-01"
)
