package wnc

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"time"
)

// HTTP timeout constants
const (
	// DefaultTLSHandshakeTimeout is the default timeout for TLS handshake
	DefaultTLSHandshakeTimeout = 10 * time.Second

	// DefaultResponseHeaderTimeout is the default timeout for response headers
	DefaultResponseHeaderTimeout = 10 * time.Second

	// DefaultIdleConnTimeout is the default timeout for idle connections
	DefaultIdleConnTimeout = 90 * time.Second
)

// HTTP header key constants
const (
	// HTTPHeaderKeyAuthorization defines the Authorization header key
	HTTPHeaderKeyAuthorization = "Authorization"

	// HTTPHeaderKeyAccept defines the Accept header key
	HTTPHeaderKeyAccept = "Accept"

	// HTTPHeaderKeyUserAgent defines the User-Agent header key
	HTTPHeaderKeyUserAgent = "User-Agent"

	// HTTPHeaderKeyContentType defines the Content-Type header key
	HTTPHeaderKeyContentType = "Content-Type"
)

// HTTP header value constants
const (
	// HTTPHeaderValueBasicPrefix defines the Basic authentication prefix
	HTTPHeaderValueBasicPrefix = "Basic "

	// HTTPHeaderValueYANGData defines the YANG data content type
	HTTPHeaderValueYANGData = "application/yang-data+json"

	// HTTPHeaderUserAgent defines the User-Agent string
	HTTPHeaderUserAgent = "wnc-go-client/1.0"

	// HTTPHeaderAccept defines the default Accept header value
	HTTPHeaderAccept = HTTPHeaderValueYANGData

	// HTTPHeaderContentType defines the default Content-Type header value
	HTTPHeaderContentType = HTTPHeaderValueYANGData
)

// createHTTPTransport creates and configures the HTTP transport
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
func (c *Client) buildRequestURL(endpoint string) string {
	return fmt.Sprintf("%s://%s%s", HTTPSScheme, c.controller, endpoint)
}

// setRequestHeaders sets HTTP headers on the request
func (c *Client) setRequestHeaders(req *http.Request, headers map[string]string) {
	for key, value := range headers {
		req.Header.Set(key, value)
	}
}

// buildHTTPHeaders constructs the HTTP headers required for API requests.
// Returns a map containing only essential headers like curl does.
func (c *Client) buildHTTPHeaders() map[string]string {
	return c.buildHTTPHeadersWithAcceptType(HTTPHeaderValueYANGData)
}

// buildHTTPHeadersWithAcceptType constructs HTTP headers with a specific Accept format.
// Only includes essential headers to match curl behavior.
func (c *Client) buildHTTPHeadersWithAcceptType(acceptType string) map[string]string {
	return map[string]string{
		HTTPHeaderKeyAuthorization: HTTPHeaderValueBasicPrefix + c.accessToken,
		HTTPHeaderKeyAccept:        acceptType,
		HTTPHeaderKeyUserAgent:     HTTPHeaderUserAgent,
	}
}

// handleHTTPError processes HTTP response errors with early returns
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
