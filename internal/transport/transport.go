// Package transport provides HTTP transport and header utilities for the WNC client.
package transport

import (
	"crypto/tls"
	"net/http"
	"time"
)

// HTTP header key constants.
const (
	// HTTPHeaderKeyAuthorization defines the Authorization header key.
	HTTPHeaderKeyAuthorization = "Authorization"
	// HTTPHeaderKeyAccept defines the Accept header key.
	HTTPHeaderKeyAccept = "Accept"
	// HTTPHeaderKeyUserAgent defines the User-Agent header key.
	HTTPHeaderKeyUserAgent = "User-Agent"
	// HTTPHeaderKeyContentType defines the Content-Type header key.
	HTTPHeaderKeyContentType = "Content-Type"
)

// HTTP header value constants.
const (
	// HTTPHeaderValueBasicPrefix defines the Basic authentication prefix.
	HTTPHeaderValueBasicPrefix = "Basic "
	// HTTPHeaderValueYANGData defines the YANG data content type.
	HTTPHeaderValueYANGData = "application/yang-data+json"
	// HTTPHeaderUserAgent defines the User-Agent string.
	HTTPHeaderUserAgent = "wnc-go-client/1.0"
	// HTTPHeaderAccept defines the default Accept header value.
	HTTPHeaderAccept = HTTPHeaderValueYANGData
	// HTTPHeaderContentType defines the default Content-Type header value.
	HTTPHeaderContentType = HTTPHeaderValueYANGData
)

// HTTP connection pool constants.
const (
	// DefaultMaxIdleConns is the maximum number of idle connections across all hosts.
	DefaultMaxIdleConns = 100
	// DefaultMaxIdleConnsPerHost is the maximum number of idle connections per host.
	DefaultMaxIdleConnsPerHost = 10
	// DefaultTLSHandshakeTimeout is the default timeout for TLS handshake.
	DefaultTLSHandshakeTimeout = QuickTimeout
	// DefaultResponseHeaderTimeout is the default timeout for response headers.
	DefaultResponseHeaderTimeout = QuickTimeout
	// DefaultIdleConnTimeout is the default timeout for idle connections.
	DefaultIdleConnTimeout = ExtendedTimeout
	// ExtendedTimeout for longer operations (idle connections).
	ExtendedTimeout = 90 * time.Second
	// QuickTimeout for fast operations (TLS handshake, response headers).
	QuickTimeout = 5 * time.Second
)

// NewTransport creates and configures a new HTTP transport with the specified TLS settings.
func NewTransport(skipVerify bool) *http.Transport {
	return &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: skipVerify, //nolint:gosec
		},
		ForceAttemptHTTP2:     false,
		DisableKeepAlives:     false,
		DisableCompression:    false,
		TLSHandshakeTimeout:   DefaultTLSHandshakeTimeout,
		ResponseHeaderTimeout: DefaultResponseHeaderTimeout,
		IdleConnTimeout:       DefaultIdleConnTimeout,
		MaxIdleConns:          DefaultMaxIdleConns,
		MaxIdleConnsPerHost:   DefaultMaxIdleConnsPerHost,
	}
}

// DefaultHeaders returns a pre-configured header map with authentication and content type.
func DefaultHeaders(token string) http.Header {
	headers := make(http.Header)
	headers.Set(HTTPHeaderKeyAuthorization, HTTPHeaderValueBasicPrefix+token)
	headers.Set(HTTPHeaderKeyAccept, HTTPHeaderValueYANGData)
	headers.Set(HTTPHeaderKeyUserAgent, HTTPHeaderUserAgent)
	return headers
}
