// Package net provides network and HTTP-related constants used across the Cisco IOS-XE Wireless Go SDK.
package net

import "time"

// Timeout duration constants (as time.Duration)
const (
	// QuickTimeout for fast operations (TLS handshake, response headers)
	QuickTimeout = 5 * time.Second

	// DefaultHTTPTimeoutDuration for client creation
	DefaultHTTPTimeoutDuration = 30 * time.Second

	// StandardTimeout for normal operations
	StandardTimeout = 60 * time.Second

	// ExtendedTimeout for longer operations (idle connections)
	ExtendedTimeout = 90 * time.Second
)
