package core

import (
	"context"
	"crypto/tls"
	"crypto/x509"
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

// Default request budgets. A request is bounded by all three, and WithTimeout sets only the first.
const (
	// DefaultTimeout is the default timeout for API requests.
	DefaultTimeout = 60 * time.Second

	// DefaultResponseHeaderTimeout is the default budget for the response headers
	// (re-export of transport.DefaultResponseHeaderTimeout).
	DefaultResponseHeaderTimeout = transport.DefaultResponseHeaderTimeout

	// DefaultTLSHandshakeTimeout is the default budget for the TLS handshake
	// (re-export of transport.DefaultTLSHandshakeTimeout).
	DefaultTLSHandshakeTimeout = transport.DefaultTLSHandshakeTimeout
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
//
// It bounds the whole request and lifts neither of the transport's other two budgets:
// DefaultResponseHeaderTimeout still ends the wait for the response headers and
// DefaultTLSHandshakeTimeout still ends the handshake. Raise those with
// WithResponseHeaderTimeout and WithTLSHandshakeTimeout. They stay separate deliberately —
// every option here mutates one shared transport in place, so one option setting all three
// would overwrite a smaller budget an earlier option had set, with no compile error.
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

// WithRootCAs verifies the controller's certificate against pool instead of the host's roots.
//
// This is the option to reach for where WithInsecureSkipVerify would otherwise be used: a
// controller presenting a certificate from a private CA is verified rather than unverified. A nil
// pool is refused, because assigning one would silently mean "use the host's roots" and read at
// the call site as the opposite.
func WithRootCAs(pool *x509.CertPool) Option {
	return func(c *Client) error {
		if pool == nil {
			return errors.New("root CA pool cannot be nil")
		}
		c.httpTransport.TLSClientConfig.RootCAs = pool
		return nil
	}
}

// WithClientCertificate presents cert to the controller, for a deployment that authenticates the
// client with mTLS as well as with the Authorization header.
//
// A certificate carrying no chain is refused: appending the zero tls.Certificate would leave the
// handshake to fail on the wire rather than here.
func WithClientCertificate(cert tls.Certificate) Option {
	return func(c *Client) error {
		if len(cert.Certificate) == 0 {
			return errors.New("client certificate carries no certificate chain")
		}
		c.httpTransport.TLSClientConfig.Certificates = []tls.Certificate{cert}
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
//
// Unset, the client logs to slog.Default(), so a process that never passed a logger still gets
// this package's Debug and Error lines — an HTTP error line among them, carrying up to
// maxLoggedBodyBytes of the controller's rejection document. Pass
// WithLogger(slog.New(slog.DiscardHandler)) to silence the SDK; the default is left as
// slog.Default() because changing it would take logging away from an existing caller with no
// compile error and nothing to notice.
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
//
// host is an authority and nothing else — "wnc1.example.internal", "192.0.2.10:443" or a bare IPv6
// literal this brackets for you. Every error New returns matches ErrInvalidConfiguration.
func New(host, token string, opts ...Option) (*Client, error) {
	// Normalize before validating: the trimmed and bracketed form is both what the
	// validator judges and what reaches the Authorization header and the request URL.
	host = validation.NormalizeHost(host)
	token = strings.TrimSpace(token)

	if err := validation.ValidateHost(host); err != nil {
		return nil, fmt.Errorf("%w: %w", ErrInvalidConfiguration, err)
	}
	if !validation.IsValidAccessToken(token) {
		return nil, fmt.Errorf("%w: access token is empty", ErrInvalidConfiguration)
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
			return nil, fmt.Errorf("%w: %w", ErrInvalidConfiguration, err)
		}
	}

	// Build the request builder last: WithLogger and WithUserAgent decide what it carries.
	client.requestBuilder = transport.NewRequestBuilder(restBuilder, token, client.userAgent, client.logger)

	return client, nil
}

// do executes an HTTP request and returns the response body.
//
// The transport methods are unexported so that the root client's untyped methods are the only
// route out of the typed API. A value reached through a service constructor can build URLs and
// services with it, and nothing else.
func (c *Client) do(ctx context.Context, method, path string) (*Response, error) {
	if err := c.validateDoParameters(ctx); err != nil {
		return nil, err
	}

	req, err := c.requestBuilder.CreateRequest(ctx, method, path)
	if err != nil {
		return nil, err
	}

	resp, err := c.execute(req)
	if err != nil {
		return nil, err
	}

	c.logger.Debug("Successfully processed API response", "path", path)
	return resp, nil
}

// doWithPayload performs an HTTP request with a payload and returns the response body.
func (c *Client) doWithPayload(ctx context.Context, method, path string, payload any) (*Response, error) {
	if err := c.validateDoParameters(ctx); err != nil {
		return nil, err
	}

	req, err := c.requestBuilder.CreateRequestWithPayload(ctx, method, path, payload)
	if err != nil {
		return nil, err
	}

	resp, err := c.execute(req)
	if err != nil {
		return nil, err
	}

	c.logger.Debug("Successfully processed API response", "path", path)
	return resp, nil
}

// doRPC posts an RPC input to rpcPath and returns the output body. RFC 8040 4.4.2 invokes an
// operation with POST and nothing else, so the method is not a parameter; RequestRaw refuses
// another method on an operations path rather than routing it here.
func (c *Client) doRPC(ctx context.Context, rpcPath string, payload any) (*Response, error) {
	if err := c.validateDoParameters(ctx); err != nil {
		return nil, err
	}

	req, err := c.requestBuilder.CreateRPCRequestWithPayload(ctx, http.MethodPost, rpcPath, payload)
	if err != nil {
		return nil, err
	}

	resp, err := c.execute(req)
	if err != nil {
		return nil, err
	}

	c.logger.Debug("Successfully processed RPC response", "rpcPath", rpcPath)
	return resp, nil
}

// execute sends req and returns its status and body once the status has been checked.
func (c *Client) execute(req *http.Request) (*Response, error) {
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

	return &Response{StatusCode: resp.StatusCode, Body: body}, nil
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

// CloseIdleConnections closes the pooled connections that have no request on them, releasing the
// sockets a long-lived process would otherwise hold for DefaultIdleConnTimeout after its last read.
//
// A connection in use is left alone and the client stays usable afterwards: the next request dials
// again. It is the one lever this package publishes over the pool, because the transport itself
// stays unexported.
func (c *Client) CloseIdleConnections() {
	if c == nil {
		return
	}
	c.httpClient.CloseIdleConnections()
}

// RESTCONFBuilder returns the RESTCONF URL builder for the client.
func (c *Client) RESTCONFBuilder() *restconf.Builder {
	if c == nil {
		return nil
	}
	return c.rest
}
