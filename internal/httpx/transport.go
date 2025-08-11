// Package httpx provides HTTP transport and header utilities for the WNC client.
package httpx

import (
	"crypto/tls"
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

// NewTransport creates and configures a new HTTP transport with the specified TLS settings
func NewTransport(skipVerify bool) *http.Transport {
	return &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: skipVerify, //nolint:gosec // Required for test environments
		},
		ForceAttemptHTTP2:     false,
		DisableKeepAlives:     false,
		DisableCompression:    false,
		TLSHandshakeTimeout:   DefaultTLSHandshakeTimeout,
		ResponseHeaderTimeout: DefaultResponseHeaderTimeout,
		IdleConnTimeout:       DefaultIdleConnTimeout,
		MaxIdleConns:          100,
		MaxIdleConnsPerHost:   10,
	}
}

// DefaultHeaders returns a pre-configured header map with authentication and content type
func DefaultHeaders(token string) http.Header {
	headers := make(http.Header)
	headers.Set(HTTPHeaderKeyAuthorization, HTTPHeaderValueBasicPrefix+token)
	headers.Set(HTTPHeaderKeyAccept, HTTPHeaderValueYANGData)
	headers.Set(HTTPHeaderKeyUserAgent, HTTPHeaderUserAgent)
	return headers
}
