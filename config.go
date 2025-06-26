package wnc

import (
	"log/slog"
	"time"
)

// Config represents the configuration for the WNC client.
// This struct contains all necessary settings for connecting to and interacting with
// the Cisco Wireless Network Controller API.
type Config struct {
	// Controller is the hostname or IP address of the WNC
	Controller string
	// AccessToken is the authentication token for API access
	AccessToken string
	// Timeout is the duration for API request timeouts (default: 15 seconds)
	Timeout time.Duration
	// InsecureSkipVerify skips TLS certificate verification when true
	// This should only be used for testing or with self-signed certificates
	InsecureSkipVerify bool
	// Logger is a custom logger instance (default: slog.Default())
	Logger *slog.Logger
}
