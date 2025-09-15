package core

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"time"

	ierrors "github.com/umatare5/cisco-ios-xe-wireless-go/internal/errors"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/transport"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/validation"
)

// Default timeout constant.
const (
	// DefaultTimeout is the default timeout for API requests.
	DefaultTimeout = 60 * time.Second
)

// Client represents the core WNC API client with connection pooling and structured logging.
type Client struct {
	httpClient     *http.Client              // Reused HTTP client with connection pool
	rest           *restconf.Builder         // RESTCONF URL builder
	logger         *slog.Logger              // Structured logger
	token          string                    // Access token for authorization
	requestBuilder *transport.RequestBuilder // HTTP request builder
}

// Option represents a functional option for configuring the Client.
type Option func(*Client) error

// WithTimeout sets the timeout duration for HTTP requests.
func WithTimeout(timeout time.Duration) Option {
	return func(c *Client) error {
		if !validation.IsValidTimeout(timeout) {
			return fmt.Errorf("client configuration failed: %w",
				fmt.Errorf("timeout validation failed: timeout must be positive, got %v", timeout))
		}
		c.httpClient.Timeout = timeout
		return nil
	}
}

// WithInsecureSkipVerify configures TLS certificate verification.
func WithInsecureSkipVerify(skip bool) Option {
	return func(c *Client) error {
		c.httpClient.Transport = transport.NewTransport(skip)
		return nil
	}
}

// WithLogger sets a custom logger for the client.
func WithLogger(logger *slog.Logger) Option {
	return func(c *Client) error {
		if logger == nil {
			return errors.New("logger cannot be nil")
		}
		c.logger = logger
		return nil
	}
}

// WithUserAgent sets a custom User-Agent header.
func WithUserAgent(userAgent string) Option {
	return func(c *Client) error {
		// This will be handled in the headers when making requests
		// For now, we store it in the client context (not implemented yet)
		return nil
	}
}

// New creates a new WNC client with the specified host, token, and options.
func New(host, token string, opts ...Option) (*Client, error) {
	// Validate inputs using existing validation functions
	if !validation.IsValidController(host) {
		return nil, fmt.Errorf("client initialization failed: %w",
			fmt.Errorf("controller address validation failed: invalid format %s", host))
	}
	if !validation.IsValidAccessToken(token) {
		return nil, fmt.Errorf("client initialization failed: %w",
			errors.New("access token validation failed: token is empty or invalid format"))
	}

	// Create HTTP client with transport
	httpClient := &http.Client{
		Transport: transport.NewTransport(false), // Default to secure
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

	// Initialize request builder
	client.requestBuilder = transport.NewRequestBuilder(restBuilder, token, client.logger)

	// Apply options
	for _, opt := range opts {
		if err := opt(client); err != nil {
			return nil, fmt.Errorf("failed to apply option: %w", err)
		}
	}

	return client, nil
}

// Do executes an HTTP request and returns the response body.
func (c *Client) Do(ctx context.Context, method, path string) ([]byte, error) {
	if err := c.validateDoParameters(ctx); err != nil {
		return nil, err
	}

	req, err := c.requestBuilder.CreateRequest(ctx, method, path)
	if err != nil {
		return nil, err
	}

	resp, err := c.requestBuilder.ExecuteRequest(c.httpClient, req)
	if err != nil {
		return nil, err
	}
	defer c.closeResponseBody(resp)

	body, err := c.readResponseBody(resp)
	if err != nil {
		return nil, err
	}

	// Early return for HTTP errors
	if err := c.checkHTTPErrors(resp, body); err != nil {
		return nil, err
	}

	c.logger.Debug("Successfully processed API response", "path", path)
	return body, nil
}

// DoWithPayload performs an HTTP request with a payload and returns the response body.
func (c *Client) DoWithPayload(ctx context.Context, method, path string, payload any) ([]byte, error) {
	if err := c.validateDoParameters(ctx); err != nil {
		return nil, err
	}

	req, err := c.requestBuilder.CreateRequestWithPayload(ctx, method, path, payload)
	if err != nil {
		return nil, err
	}

	resp, err := c.requestBuilder.ExecuteRequest(c.httpClient, req)
	if err != nil {
		return nil, err
	}
	defer c.closeResponseBody(resp)

	body, err := c.readResponseBody(resp)
	if err != nil {
		return nil, err
	}

	if err := c.checkHTTPErrors(resp, body); err != nil {
		return nil, err
	}

	c.logger.Debug("Successfully processed API response", "path", path)
	return body, nil
}

// DoRPCWithPayload performs an HTTP RPC request with a payload and returns the response body.
func (c *Client) DoRPCWithPayload(ctx context.Context, method, rpcPath string, payload any) ([]byte, error) {
	if err := c.validateDoParameters(ctx); err != nil {
		return nil, err
	}

	req, err := c.requestBuilder.CreateRPCRequestWithPayload(ctx, method, rpcPath, payload)
	if err != nil {
		return nil, err
	}

	resp, err := c.requestBuilder.ExecuteRequest(c.httpClient, req)
	if err != nil {
		return nil, err
	}
	defer c.closeResponseBody(resp)

	body, err := c.readResponseBody(resp)
	if err != nil {
		return nil, err
	}

	if err := c.checkHTTPErrors(resp, body); err != nil {
		return nil, err
	}

	c.logger.Debug("Successfully processed RPC response", "rpcPath", rpcPath)
	return body, nil
}

// validateDoParameters validates input parameters for the Do method.
func (c *Client) validateDoParameters(ctx context.Context) error {
	if c == nil {
		return errors.New(ierrors.ErrClientNil)
	}
	if ctx == nil {
		return errors.New("context cannot be nil")
	}
	return nil
}

// closeResponseBody safely closes the response body with error logging.
func (c *Client) closeResponseBody(resp *http.Response) {
	if closeErr := resp.Body.Close(); closeErr != nil {
		c.logger.Error("Failed to close response body", "error", closeErr)
	}
}

// readResponseBody reads the complete response body.
func (c *Client) readResponseBody(resp *http.Response) ([]byte, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.logger.Error("Failed to read response body", "error", err)
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	c.logger.Debug("Received API response", "status", resp.StatusCode, "content_length", len(body))
	return body, nil
}

// checkHTTPErrors validates HTTP status codes and returns appropriate errors.
func (c *Client) checkHTTPErrors(resp *http.Response, body []byte) error {
	if resp.StatusCode >= http.StatusBadRequest {
		c.logger.Error("HTTP error response", "status", resp.StatusCode, "body", string(body))
		return &APIError{
			StatusCode: resp.StatusCode,
			Message:    string(body),
			Body:       body,
		}
	}
	return nil
}

// RESTCONFBuilder returns the RESTCONF URL builder for the client.
func (c *Client) RESTCONFBuilder() *restconf.Builder {
	if c == nil {
		return nil
	}
	return c.rest
}
