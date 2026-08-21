package core

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net"
	"net/http"
	"net/url"
	"strings"
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

// maxLoggedBodyBytes bounds the error body copied into the log line and into
// APIError.Message. APIError.Body keeps the whole document.
const maxLoggedBodyBytes = 512

// Client represents the core WNC API client with connection pooling and structured logging.
type Client struct {
	httpClient     *http.Client              // Reused HTTP client with connection pool
	httpTransport  *http.Transport           // Same transport as httpClient.Transport, mutated in place by options
	rest           *restconf.Builder         // RESTCONF URL builder
	logger         *slog.Logger              // Structured logger
	token          string                    // Access token for authorization
	userAgent      string                    // User-Agent header value; empty means the transport default
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
		c.httpTransport.TLSClientConfig.InsecureSkipVerify = skip //nolint:gosec
		return nil
	}
}

// WithProxy routes every request through the proxy the resolver returns.
// Proxying is off by default: no environment variable is consulted unless this
// option names a resolver. Pass http.ProxyFromEnvironment to honor HTTP_PROXY,
// HTTPS_PROXY and NO_PROXY, or http.ProxyURL to pin one proxy. A nil resolver,
// like a resolver returning a nil URL, connects directly.
func WithProxy(fn func(*http.Request) (*url.URL, error)) Option {
	return func(c *Client) error {
		c.httpTransport.Proxy = fn
		return nil
	}
}

// WithResponseHeaderTimeout bounds the wait from the end of the request write to the
// first byte of the response headers.
func WithResponseHeaderTimeout(timeout time.Duration) Option {
	return func(c *Client) error {
		if !validation.IsValidTimeout(timeout) {
			return fmt.Errorf("client configuration failed: %w",
				fmt.Errorf("response header timeout validation failed: timeout must be positive, got %v", timeout))
		}
		c.httpTransport.ResponseHeaderTimeout = timeout
		return nil
	}
}

// WithTLSHandshakeTimeout bounds the TLS handshake.
func WithTLSHandshakeTimeout(timeout time.Duration) Option {
	return func(c *Client) error {
		if !validation.IsValidTimeout(timeout) {
			return fmt.Errorf("client configuration failed: %w",
				fmt.Errorf("TLS handshake timeout validation failed: timeout must be positive, got %v", timeout))
		}
		c.httpTransport.TLSHandshakeTimeout = timeout
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
		if !validation.IsNonEmptyString(userAgent) {
			return errors.New("user agent cannot be empty")
		}
		c.userAgent = strings.TrimSpace(userAgent)
		return nil
	}
}

// New creates a new WNC client with the specified host, token, and options.
func New(host, token string, opts ...Option) (*Client, error) {
	// Normalize before validating: the trimmed and bracketed form is both what the
	// validator judges and what reaches the Authorization header and the request URL.
	host = validation.NormalizeHost(host)
	token = strings.TrimSpace(token)

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
	httpTransport := transport.NewTransport(false) // Default to secure
	httpClient := &http.Client{
		Transport: httpTransport,
		Timeout:   DefaultTimeout,
	}

	// Create RESTCONF URL builder
	restBuilder := restconf.NewBuilder(restconf.DefaultProtocol, host)

	// Create client with defaults
	client := &Client{
		httpClient:    httpClient,
		httpTransport: httpTransport,
		rest:          restBuilder,
		logger:        slog.Default(),
		token:         token,
	}

	// Apply options
	for _, opt := range opts {
		if err := opt(client); err != nil {
			return nil, fmt.Errorf("failed to apply option: %w", err)
		}
	}

	// Build the request builder last: WithLogger and WithUserAgent decide what it carries.
	client.requestBuilder = transport.NewRequestBuilder(restBuilder, token, client.userAgent, client.logger)

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

	body, err := c.execute(req)
	if err != nil {
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

	body, err := c.execute(req)
	if err != nil {
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

	body, err := c.execute(req)
	if err != nil {
		return nil, err
	}

	c.logger.Debug("Successfully processed RPC response", "rpcPath", rpcPath)
	return body, nil
}

// execute sends req and returns its body once the status has been checked.
func (c *Client) execute(req *http.Request) ([]byte, error) {
	resp, err := c.requestBuilder.ExecuteRequest(c.httpClient, req)
	if err != nil {
		return nil, classifyTransportError(err)
	}
	defer c.closeResponseBody(resp)

	body, err := c.readResponseBody(resp)
	if err != nil {
		return nil, err
	}

	if err := c.checkHTTPErrors(resp, body); err != nil {
		return nil, err
	}

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
		// The same client deadline can fire before the headers arrive or partway through the
		// body, and both are one event to whoever set the deadline.
		return nil, classifyTransportError(fmt.Errorf("failed to read response: %w", err))
	}

	c.logger.Debug("Received API response", "status", resp.StatusCode, "content_length", len(body))
	return body, nil
}

// checkHTTPErrors validates HTTP status codes and returns appropriate errors.
func (c *Client) checkHTTPErrors(resp *http.Response, body []byte) error {
	if resp.StatusCode < http.StatusBadRequest {
		return nil
	}

	summary := truncateBody(body)
	c.logger.Error("HTTP error response", "status", resp.StatusCode,
		"content_length", len(body), "body_prefix", summary)

	return &APIError{
		StatusCode: resp.StatusCode,
		Message:    summary,
		Body:       body,
	}
}

// truncateBody bounds a response body for logging, dropping the partial rune the cut
// may leave behind.
func truncateBody(body []byte) string {
	if len(body) <= maxLoggedBodyBytes {
		return string(body)
	}
	return strings.ToValidUTF8(string(body[:maxLoggedBodyBytes]), "") + "... (truncated)"
}

// classifyTransportError maps a transport failure onto the SDK error taxonomy so a
// caller can match it with errors.Is. The original error stays in the chain, so
// errors.As for *url.Error keeps working.
func classifyTransportError(err error) error {
	var netErr net.Error
	if (errors.As(err, &netErr) && netErr.Timeout()) || errors.Is(err, context.DeadlineExceeded) {
		return fmt.Errorf("%w: %w", ErrRequestTimeout, err)
	}
	return err
}

// RESTCONFBuilder returns the RESTCONF URL builder for the client.
func (c *Client) RESTCONFBuilder() *restconf.Builder {
	if c == nil {
		return nil
	}
	return c.rest
}
