// Package wnc provides the core client for the Cisco Wireless Network Controller API.
// This package implements the three-layer architecture with Core, Domain Service, and Generated Type separation.
package wnc

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"time"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/httpx"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf"
)

// Default timeout constant
const (
	// DefaultTimeout is the default timeout for API requests
	DefaultTimeout = 60 * time.Second
)

// HTTPError represents an HTTP error response from the API
type HTTPError struct {
	Status int    // HTTP status code
	Body   []byte // Response body
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("HTTP %d: %s", e.Status, string(e.Body))
}

// Client represents the core WNC API client with connection pooling and structured logging
type Client struct {
	httpClient *http.Client      // Reused HTTP client with connection pool
	rest       *restconf.Builder // RESTCONF URL builder
	logger     *slog.Logger      // Structured logger
	token      string            // Access token for authorization
}

// Option represents a functional option for configuring the Client
type Option func(*Client) error

// WithTimeout sets the timeout duration for HTTP requests
func WithTimeout(timeout time.Duration) Option {
	return func(c *Client) error {
		if timeout <= 0 {
			return fmt.Errorf("timeout must be positive, got %v", timeout)
		}
		c.httpClient.Timeout = timeout
		return nil
	}
}

// WithInsecureSkipVerify configures TLS certificate verification
func WithInsecureSkipVerify(skip bool) Option {
	return func(c *Client) error {
		// Create new transport with updated TLS settings
		transport := httpx.NewTransport(skip)
		c.httpClient.Transport = transport
		return nil
	}
}

// WithLogger sets a custom logger for the client
func WithLogger(logger *slog.Logger) Option {
	return func(c *Client) error {
		if logger == nil {
			return fmt.Errorf("logger cannot be nil")
		}
		c.logger = logger
		return nil
	}
}

// WithUserAgent sets a custom User-Agent header
func WithUserAgent(userAgent string) Option {
	return func(c *Client) error {
		// This will be handled in the headers when making requests
		// For now, we store it in the client context (not implemented yet)
		return nil
	}
}

// New creates a new WNC client with the specified host, token, and options
func New(host, token string, opts ...Option) (*Client, error) {
	// Validate inputs using existing validation functions
	if !isValidController(host) {
		return nil, fmt.Errorf("controller address is required")
	}
	if !isValidAccessToken(token) {
		return nil, fmt.Errorf("access token is required")
	}

	// Create HTTP transport with default settings
	transport := httpx.NewTransport(false) // Default to secure

	// Create HTTP client with transport
	httpClient := &http.Client{
		Transport: transport,
		Timeout:   DefaultTimeout,
	}

	// Create RESTCONF URL builder
	restBuilder := restconf.NewBuilder(restconf.DefaultProtocol, host)

	// Create client with defaults
	client := &Client{
		httpClient: httpClient,
		rest:       restBuilder,
		logger:     slog.Default(),
		token:      token,
	}

	// Apply options
	for _, opt := range opts {
		if err := opt(client); err != nil {
			return nil, fmt.Errorf("failed to apply option: %w", err)
		}
	}

	return client, nil
}

// Do performs a generic HTTP request to the specified path and unmarshals the response into out
func (c *Client) Do(ctx context.Context, method, path string, out any) error {
	if ctx == nil {
		return fmt.Errorf("context cannot be nil")
	}
	if out == nil {
		return fmt.Errorf("output parameter cannot be nil")
	}

	// Build full URL using RESTCONF builder
	url := c.rest.BuildRESTCONFURL(path)

	// Create HTTP request
	req, err := http.NewRequestWithContext(ctx, method, url, nil)
	if err != nil {
		c.logger.Error("Failed to create HTTP request", "error", err, "url", url)
		return fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers using httpx helper
	req.Header = httpx.DefaultHeaders(c.token)

	c.logger.Debug("Sending API request", "method", method, "url", url)

	// Execute request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		c.logger.Error("HTTP request failed", "error", err, "url", url)
		return fmt.Errorf("request failed: %w", err)
	}
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			c.logger.Error("Failed to close response body", "error", closeErr)
		}
	}()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.logger.Error("Failed to read response body", "error", err)
		return fmt.Errorf("failed to read response: %w", err)
	}

	c.logger.Debug("Received API response", "status", resp.StatusCode, "content_length", len(body))

	// Check for HTTP errors
	if resp.StatusCode >= 400 {
		c.logger.Error("HTTP error response", "status", resp.StatusCode, "body", string(body))
		return &HTTPError{
			Status: resp.StatusCode,
			Body:   body,
		}
	}

	// Unmarshal JSON response
	if err := json.Unmarshal(body, out); err != nil {
		c.logger.Error("Failed to unmarshal JSON response", "error", err, "body_length", len(body))
		return fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.Debug("Successfully processed API response", "path", path)
	return nil
}

// Domain service interfaces - these will be implemented by the respective packages

// AFCService defines the interface for Automated Frequency Coordination operations
type AFCService interface {
	// Methods will be added as the afc package is refactored
}

// APService defines the interface for Access Point operations
type APService interface {
	// Methods will be added as the ap package is refactored
}

// ClientService defines the interface for wireless client operations
type ClientService interface {
	// Methods will be added as the client package is refactored
}

// GeneralService defines the interface for general controller operations
type GeneralService interface {
	// Methods will be added as the general package is refactored
}

// RRMService defines the interface for Radio Resource Management operations
type RRMService interface {
	// Methods will be added as the rrm package is refactored
}

// WLANService defines the interface for WLAN operations
type WLANService interface {
	// Methods will be added as the wlan package is refactored
}

// Domain service accessors - these create service instances that use the client's Do() method

// AFC returns an AFC service instance
func (c *Client) AFC() AFCService {
	// This will be implemented when AFC package is refactored to use the new client
	return nil // Placeholder
}

// AP returns an Access Point service instance
func (c *Client) AP() APService {
	// This will be implemented when AP package is refactored to use the new client
	return nil // Placeholder
}

// Client returns a wireless client service instance
func (c *Client) Client() ClientService {
	// This will be implemented when client package is refactored to use the new client
	return nil // Placeholder
}

// General returns a general controller service instance
func (c *Client) General() GeneralService {
	// This will be implemented when general package is refactored to use the new client
	return nil // Placeholder
}

// RRM returns a Radio Resource Management service instance
func (c *Client) RRM() RRMService {
	// This will be implemented when RRM package is refactored to use the new client
	return nil // Placeholder
}

// WLAN returns a WLAN service instance
func (c *Client) WLAN() WLANService {
	// This will be implemented when WLAN package is refactored to use the new client
	return nil // Placeholder
}

// Helper functions - these need to be imported or defined

// isValidController checks if controller address is valid
func isValidController(controller string) bool {
	return controller != ""
}

// isValidAccessToken checks if access token is valid
func isValidAccessToken(accessToken string) bool {
	return accessToken != ""
}

// isPositiveTimeout checks if timeout is greater than zero
func isPositiveTimeout(timeout time.Duration) bool {
	return timeout > 0
}
