// Package wnc provides the components for interacting with the Cisco Wireless Network Controller API.
package wnc

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
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
	// Early return for invalid controller
	if !isValidController(config.Controller) {
		return nil, fmt.Errorf("%w: controller address is required", ErrInvalidConfiguration)
	}

	// Early return for invalid access token
	if !isValidAccessToken(config.AccessToken) {
		return nil, fmt.Errorf("%w: access token is required", ErrInvalidConfiguration)
	}

	// Apply defaults for missing configuration values
	if config.Timeout <= ZeroTimeoutSeconds {
		config.Timeout = DefaultTimeout
	}

	if config.Logger == nil {
		config.Logger = slog.Default()
	}

	client := &Client{
		controller:         config.Controller,
		accessToken:        config.AccessToken,
		timeout:            config.Timeout,
		insecureSkipVerify: config.InsecureSkipVerify,
		logger:             config.Logger,
	}

	// Apply additional options
	for _, option := range options {
		option(client)
	}

	// Final validation with early return for invalid timeout
	if !isPositiveTimeout(client.timeout) {
		return nil, fmt.Errorf("%w: timeout must be positive", ErrInvalidConfiguration)
	}

	return client, nil
}

// SendAPIRequest sends an API request to the specified endpoint and unmarshals the response into the result.
// This is the core method used by all feature-specific methods in the individual packages.
//
// Parameters:
//   - ctx: Context for request cancellation and timeouts
//   - endpoint: The API endpoint URL to call
//   - result: Pointer to a struct where the response should be unmarshaled
//
// Returns an error if the request fails or if the response cannot be unmarshaled.
func (c *Client) SendAPIRequest(ctx context.Context, endpoint string, result any) error {
	if ctx == nil {
		ctx = context.Background()
	}

	requestContext, cancelFunc := context.WithTimeout(ctx, c.timeout)
	defer cancelFunc()

	httpRequest, err := c.createHTTPRequest(requestContext, endpoint)
	if err != nil {
		return err
	}

	httpResponse, err := c.executeHTTPRequest(httpRequest)
	if err != nil {
		return err
	}
	defer httpResponse.Body.Close()

	return c.processHTTPResponse(httpResponse, endpoint, result)
}

// createHTTPRequest creates and configures an HTTP request
func (c *Client) createHTTPRequest(ctx context.Context, endpoint string) (*http.Request, error) {
	requestURL := c.buildRequestURL(endpoint)

	httpRequest, err := http.NewRequestWithContext(ctx, HTTPMethodGet, requestURL, nil)
	if err != nil {
		c.logger.Error("Failed to create HTTP request", "error", err, "url", requestURL)
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	requestHeaders := c.buildHTTPHeaders()
	c.setRequestHeaders(httpRequest, requestHeaders)

	c.logger.Debug("Sending API request", "method", httpRequest.Method, "url", requestURL)
	return httpRequest, nil
}

// executeHTTPRequest executes an HTTP request and handles common errors
func (c *Client) executeHTTPRequest(httpRequest *http.Request) (*http.Response, error) {
	httpTransport := c.createHTTPTransport()
	httpClient := &http.Client{
		Transport: httpTransport,
	}

	httpResponse, err := httpClient.Do(httpRequest)
	if err != nil {
		c.logger.Error("HTTP request failed", "error", err, "url", httpRequest.URL.String())

		// Early return for timeout errors
		if isDeadlineExceededError(err) {
			return nil, ErrRequestTimeout
		}
		return nil, fmt.Errorf("request failed: %w", err)
	}

	return httpResponse, nil
}

// processHTTPResponse processes the HTTP response and unmarshals the result
func (c *Client) processHTTPResponse(httpResponse *http.Response, endpoint string, result any) error {
	responseBody, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		c.logger.Error("Failed to read response body", "error", err)
		return fmt.Errorf("failed to read response: %w", err)
	}

	c.logger.Debug("Received API response", "status", httpResponse.StatusCode, "content_length", len(responseBody))

	if err := c.handleHTTPError(httpResponse.StatusCode, responseBody, httpResponse.Request.URL.String()); err != nil {
		return err
	}

	if err := json.Unmarshal(responseBody, result); err != nil {
		c.logger.Error("Failed to unmarshal JSON response", "error", err, "body_length", len(responseBody))
		return fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.logger.Debug("Successfully processed API response", "endpoint", endpoint)
	return nil
}
