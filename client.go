// Package wnc provides the components for interacting with the Cisco Wireless Network Controller API.
package wnc

import (
	"context"
	"log/slog"
	"net/http"

	wnccore "github.com/umatare5/cisco-ios-xe-wireless-go/wnc"
)

// HTTP method constants
const (
	// HTTPMethodGet defines the GET HTTP method
	HTTPMethodGet = http.MethodGet

	// HTTPMethodPost defines the POST HTTP method
	HTTPMethodPost = http.MethodPost

	// HTTPMethodPut defines the PUT HTTP method
	HTTPMethodPut = http.MethodPut

	// HTTPMethodDelete defines the DELETE HTTP method
	HTTPMethodDelete = http.MethodDelete
)

// NewClient creates a new WNC client using a configuration struct.
// This is the primary constructor that follows the architectural guidelines.
//
// Deprecated: Use wnc.New() from the wnc package instead.
//
// Example usage:
//
//	config := wnc.Config{
//		Controller:         "controller.example.com",
//		AccessToken:        "your-access-token",
//		Timeout:            15 * time.Second,
//		InsecureSkipVerify: true,
//		Logger:             customLogger,
//	}
//	client, err := wnc.NewClient(config)
func NewClient(config Config) (*Client, error) {
	return NewClientWithConfig(config)
}

// NewClientWithConfig creates a new WNC client using a configuration struct.
// This is the preferred method for creating clients as it follows the architectural guidelines.
// Additional configuration can still be provided through options for flexibility.
//
// Deprecated: Use wnc.New() from the wnc package instead.
//
// Example usage:
//
//	config := wnc.Config{
//		Controller:         "controller.example.com",
//		AccessToken:        "your-access-token",
//		Timeout:            15 * time.Second,
//		InsecureSkipVerify: true,
//		Logger:             customLogger,
//	}
//	client, err := wnc.NewClientWithConfig(config)
func NewClientWithConfig(config Config, options ...ClientOption) (*Client, error) {
	// Convert old Config to new client creation
	var opts []wnccore.Option

	// Handle timeout - set default if not specified
	finalTimeout := config.Timeout
	if finalTimeout <= 0 {
		finalTimeout = DefaultTimeout
	}
	if config.Timeout > 0 {
		opts = append(opts, wnccore.WithTimeout(config.Timeout))
	}

	// Handle InsecureSkipVerify
	if config.InsecureSkipVerify {
		opts = append(opts, wnccore.WithInsecureSkipVerify(true))
	}

	// Handle Logger - set default if not specified
	finalLogger := config.Logger
	if finalLogger == nil {
		finalLogger = slog.Default()
	}
	if config.Logger != nil {
		opts = append(opts, wnccore.WithLogger(config.Logger))
	}

	// Apply additional legacy options by converting them
	for _, option := range options {
		// Create a temporary client to extract the option values
		tempClient := &Client{
			timeout: DefaultTimeout,
			logger:  slog.Default(),
		}
		option(tempClient)

		// Convert to new options and track values for legacy fields
		if tempClient.timeout != DefaultTimeout {
			opts = append(opts, wnccore.WithTimeout(tempClient.timeout))
			finalTimeout = tempClient.timeout
		}
		if tempClient.logger != slog.Default() && tempClient.logger != nil {
			opts = append(opts, wnccore.WithLogger(tempClient.logger))
			finalLogger = tempClient.logger
		}
	}

	// Create new core client
	coreClient, err := wnccore.New(config.Controller, config.AccessToken, opts...)
	if err != nil {
		return nil, err
	}

	// Create legacy wrapper
	return &Client{
		coreClient:         coreClient,
		controller:         config.Controller,
		accessToken:        config.AccessToken,
		timeout:            finalTimeout,
		insecureSkipVerify: config.InsecureSkipVerify,
		logger:             finalLogger,
	}, nil
}

// SendAPIRequest sends an API request to the specified endpoint and unmarshals the response into the result.
// This is the core method used by all feature-specific methods in the individual packages.
//
// Deprecated: This method now delegates to the new core client implementation.
//
// Parameters:
//   - ctx: Context for request cancellation and timeouts
//   - endpoint: The API endpoint URL to call
//   - result: Pointer to a struct where the response should be unmarshaled
//
// Returns an error if the request fails or if the response cannot be unmarshaled.
func (c *Client) SendAPIRequest(ctx context.Context, endpoint string, result any) error {
	// Delegate to the new core client's Do method
	return c.coreClient.Do(ctx, "GET", endpoint, result)
}
