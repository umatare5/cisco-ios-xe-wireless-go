// Package wnc provides validation functions and configuration options for the Cisco Wireless Network Controller API client.
package wnc

import (
	"log/slog"
	"time"
)

// ClientOption represents an option for configuring the WNC client.
// This provides a functional options pattern for backwards compatibility.
type ClientOption func(*Client)

// WithTimeout sets the timeout duration for API requests.
func WithTimeout(timeout time.Duration) ClientOption {
	return func(c *Client) {
		c.timeout = timeout
	}
}

// WithInsecureSkipVerify skips TLS certificate verification.
// This should only be used for testing or when connecting to controllers with self-signed certificates.
func WithInsecureSkipVerify(skip bool) ClientOption {
	return func(c *Client) {
		c.insecureSkipVerify = skip
	}
}

// WithLogger sets a custom logger for the client.
func WithLogger(logger *slog.Logger) ClientOption {
	return func(c *Client) {
		c.logger = logger
	}
}
