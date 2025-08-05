package wnc

// This file contains deprecated HTTP helper functions.
// These have been moved to internal/httpx package.
//
// Deprecated: Use internal/httpx package instead.

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"time"
)

// HTTP timeout constants - deprecated, moved to internal/httpx
const (
	// DefaultTLSHandshakeTimeout is the default timeout for TLS handshake
	// Deprecated: Use internal/httpx constants
	DefaultTLSHandshakeTimeout = 10 * time.Second

	// DefaultResponseHeaderTimeout is the default timeout for response headers
	// Deprecated: Use internal/httpx constants
	DefaultResponseHeaderTimeout = 10 * time.Second

	// DefaultIdleConnTimeout is the default timeout for idle connections
	// Deprecated: Use internal/httpx constants
	DefaultIdleConnTimeout = 90 * time.Second
)

// HTTP header key constants - deprecated, moved to internal/httpx
const (
	// HTTPHeaderKeyAuthorization defines the Authorization header key
	// Deprecated: Use internal/httpx constants
	HTTPHeaderKeyAuthorization = "Authorization"

	// HTTPHeaderKeyAccept defines the Accept header key
	// Deprecated: Use internal/httpx constants
	HTTPHeaderKeyAccept = "Accept"

	// HTTPHeaderKeyUserAgent defines the User-Agent header key
	// Deprecated: Use internal/httpx constants
	HTTPHeaderKeyUserAgent = "User-Agent"

	// HTTPHeaderKeyContentType defines the Content-Type header key
	// Deprecated: Use internal/httpx constants
	HTTPHeaderKeyContentType = "Content-Type"
)

// HTTP header value constants - deprecated, moved to internal/httpx
const (
	// HTTPHeaderValueBasicPrefix defines the Basic authentication prefix
	// Deprecated: Use internal/httpx constants
	HTTPHeaderValueBasicPrefix = "Basic "

	// HTTPHeaderValueYANGData defines the YANG data content type
	// Deprecated: Use internal/httpx constants
	HTTPHeaderValueYANGData = "application/yang-data+json"

	// HTTPHeaderUserAgent defines the User-Agent string
	// Deprecated: Use internal/httpx constants
	HTTPHeaderUserAgent = "wnc-go-client/1.0"

	// HTTPHeaderAccept defines the default Accept header value
	// Deprecated: Use internal/httpx constants
	HTTPHeaderAccept = HTTPHeaderValueYANGData

	// HTTPHeaderContentType defines the default Content-Type header value
	// Deprecated: Use internal/httpx constants
	HTTPHeaderContentType = HTTPHeaderValueYANGData
)

// Deprecated functions - these are kept for backward compatibility but delegate to new implementations

// createHTTPTransport creates and configures the HTTP transport
// Deprecated: Use internal/httpx.NewTransport instead
func (c *Client) createHTTPTransport() *http.Transport {
	return &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: c.insecureSkipVerify,
		},
		ForceAttemptHTTP2:     false,
		DisableKeepAlives:     false,
		DisableCompression:    false,
		TLSHandshakeTimeout:   DefaultTLSHandshakeTimeout,
		ResponseHeaderTimeout: DefaultResponseHeaderTimeout,
		IdleConnTimeout:       DefaultIdleConnTimeout,
	}
}

// buildRequestURL constructs the full request URL
// Deprecated: Use internal/restconf.Builder instead
func (c *Client) buildRequestURL(endpoint string) string {
	return fmt.Sprintf("https://%s%s", c.controller, endpoint)
}

// setRequestHeaders sets HTTP headers on the request
// Deprecated: Use internal/httpx.DefaultHeaders instead
func (c *Client) setRequestHeaders(req *http.Request, headers map[string]string) {
	for key, value := range headers {
		req.Header.Set(key, value)
	}
}

// buildHTTPHeaders constructs the HTTP headers required for API requests.
// Deprecated: Use internal/httpx.DefaultHeaders instead
func (c *Client) buildHTTPHeaders() map[string]string {
	return c.buildHTTPHeadersWithAcceptType(HTTPHeaderValueYANGData)
}

// buildHTTPHeadersWithAcceptType constructs HTTP headers with a specific Accept format.
// Deprecated: Use internal/httpx.DefaultHeaders instead
func (c *Client) buildHTTPHeadersWithAcceptType(acceptType string) map[string]string {
	return map[string]string{
		HTTPHeaderKeyAuthorization: HTTPHeaderValueBasicPrefix + c.accessToken,
		HTTPHeaderKeyAccept:        acceptType,
		HTTPHeaderKeyUserAgent:     HTTPHeaderUserAgent,
	}
}

// handleHTTPError processes HTTP response errors with early returns
// Deprecated: HTTP error handling is now done in the core client
func (c *Client) handleHTTPError(statusCode int, responseBody []byte, requestURL string) error {
	if isAuthenticationError(statusCode) {
		c.logger.Error("Authentication failed", "status", statusCode)
		return ErrAuthenticationFailed
	}

	if isAccessForbiddenError(statusCode) {
		c.logger.Error("Access forbidden", "status", statusCode)
		return ErrAccessForbidden
	}

	if isNotFoundError(statusCode) {
		c.logger.Error("Resource not found", "status", statusCode, "url", requestURL)
		return ErrResourceNotFound
	}

	if !isSuccessStatusCode(statusCode) {
		c.logger.Error("HTTP error", "status", statusCode, "body", string(responseBody))
		return fmt.Errorf("HTTP error %d: %s", statusCode, string(responseBody))
	}

	return nil
}
