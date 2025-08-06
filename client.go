// Package wnc provides the components for interacting with the Cisco Wireless Network Controller API.
package wnc

import (
	"context"
	"log/slog"
	"net/http"
	"strings"

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
	// Remove "/restconf/data/" prefix if present to avoid duplication
	// The core client's BuildRESTCONFURL method will add the proper prefix
	cleanEndpoint := endpoint
	if strings.HasPrefix(endpoint, "/restconf/data/") {
		cleanEndpoint = strings.TrimPrefix(endpoint, "/restconf/data/")
	}

	// Delegate to the new core client's Do method
	return c.coreClient.Do(ctx, "GET", cleanEndpoint, result)
}

// CoreClient returns the underlying wnc.Client for use with domain services.
// This provides access to the new three-layer architecture components.
//
// Example usage with domain services:
//
//	import (
//	  "github.com/umatare5/cisco-ios-xe-wireless-go/afc"
//	  "github.com/umatare5/cisco-ios-xe-wireless-go/general"
//	)
//
//	// Create services using the core client
//	afcService := afc.NewService(client.CoreClient())
//	generalService := general.NewService(client.CoreClient())
//	geolocationService := geolocation.NewService(client.CoreClient())
//	apService := ap.NewService(client.CoreClient())
//
//	// Use the new service-based API
//	afcOper, err := afcService.Oper(ctx)
//	generalCfg, err := generalService.Cfg(ctx)
//	apCfg, err := apService.Cfg(ctx)
func (c *Client) CoreClient() *wnccore.Client {
	return c.coreClient
}

// Service domain methods for direct access
// These create service instances from domain packages

// AFC returns the AFC service
func (c *Client) AFC() interface{} {
	// Import dynamically to avoid import cycles
	// This requires reflection but is cleaner
	return struct {
		Oper       func(context.Context) (interface{}, error)
		CloudStats func(context.Context) (interface{}, error)
	}{
		Oper: func(ctx context.Context) (interface{}, error) {
			var result interface{}
			err := c.coreClient.Do(ctx, "GET", "Cisco-IOS-XE-wireless-afc-oper:afc-oper-data", &result)
			return result, err
		},
		CloudStats: func(ctx context.Context) (interface{}, error) {
			var result interface{}
			err := c.coreClient.Do(ctx, "GET", "Cisco-IOS-XE-wireless-afc-oper:afc-oper-data/afc-cloud-stats", &result)
			return result, err
		},
	}
}

// General returns the General service
func (c *Client) General() interface{} {
	return struct {
		Oper func(context.Context) (interface{}, error)
		Cfg  func(context.Context) (interface{}, error)
	}{
		Oper: func(ctx context.Context) (interface{}, error) {
			var result interface{}
			err := c.coreClient.Do(ctx, "GET", "Cisco-IOS-XE-wireless-general-oper:general-oper-data", &result)
			return result, err
		},
		Cfg: func(ctx context.Context) (interface{}, error) {
			var result interface{}
			err := c.coreClient.Do(ctx, "GET", "Cisco-IOS-XE-wireless-general-cfg:general-cfg-data", &result)
			return result, err
		},
	}
}

// New creates a new WNC client using functional options.
// This is the recommended constructor for the service-based API.
//
// Example usage:
//
//	client, err := wnc.New("192.168.1.100", "token",
//		wnc.WithTimeout(30*time.Second),
//		wnc.WithInsecureSkipVerify(true))
func New(host, token string, opts ...wnccore.Option) (*Client, error) {
	// Create new core client
	coreClient, err := wnccore.New(host, token, opts...)
	if err != nil {
		return nil, err
	}

	// Create legacy wrapper
	return &Client{
		coreClient:         coreClient,
		controller:         host,
		accessToken:        token,
		timeout:            DefaultTimeout,
		insecureSkipVerify: false,
		logger:             slog.Default(),
	}, nil
}
