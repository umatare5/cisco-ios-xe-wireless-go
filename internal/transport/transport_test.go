package transport

import (
	"testing"
	"time"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil"
)

func TestTransportUnit_NewTransport_Success(t *testing.T) {
	testCases := []struct {
		name       string
		skipVerify bool
	}{
		{"with TLS verification enabled", false},
		{"with TLS verification disabled", true},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			transport := NewTransport(tt.skipVerify)

			testutil.AssertNotNil(t, transport, "NewTransport result")

			// Check TLS configuration
			testutil.AssertNotNil(t, transport.TLSClientConfig, "TLSClientConfig")
			testutil.AssertBoolEquals(
				t,
				transport.TLSClientConfig.InsecureSkipVerify,
				tt.skipVerify,
				"TLSClientConfig.InsecureSkipVerify",
			)

			// Check timeout settings
			testutil.AssertDurationEquals(t, transport.TLSHandshakeTimeout,
				DefaultTLSHandshakeTimeout, "TLSHandshakeTimeout")

			testutil.AssertDurationEquals(t, transport.ResponseHeaderTimeout,
				DefaultResponseHeaderTimeout, "ResponseHeaderTimeout")

			testutil.AssertDurationEquals(
				t,
				transport.IdleConnTimeout,
				DefaultIdleConnTimeout,
				"IdleConnTimeout",
			)

			// Check connection settings
			testutil.AssertFalse(t, transport.ForceAttemptHTTP2, "ForceAttemptHTTP2")

			testutil.AssertFalse(t, transport.DisableKeepAlives, "DisableKeepAlives")

			testutil.AssertFalse(t, transport.DisableCompression, "DisableCompression")

			testutil.AssertIntEquals(t, transport.MaxIdleConns, 100, "MaxIdleConns")

			testutil.AssertIntEquals(t, transport.MaxIdleConnsPerHost, 10, "MaxIdleConnsPerHost")
		})
	}
}

func TestTransportUnit_DefaultHeaders_Success(t *testing.T) {
	token := "test-token-123"
	headers := DefaultHeaders(token, "")

	testutil.AssertNotNil(t, headers, "DefaultHeaders result")

	// Check Authorization header
	expectedAuth := HTTPHeaderValueBasicPrefix + token
	auth := headers.Get(HTTPHeaderKeyAuthorization)
	testutil.AssertStringEquals(t, auth, expectedAuth, "Authorization header")

	// Check Accept header
	accept := headers.Get(HTTPHeaderKeyAccept)
	testutil.AssertStringEquals(t, accept, HTTPHeaderValueYANGData, "Accept header")

	// Check User-Agent header
	userAgent := headers.Get(HTTPHeaderKeyUserAgent)
	testutil.AssertStringEquals(t, userAgent, HTTPHeaderUserAgent, "User-Agent header")
}

func TestTransportUnit_DefaultHeadersWithEmptyToken_Success(t *testing.T) {
	headers := DefaultHeaders("", "")

	auth := headers.Get(HTTPHeaderKeyAuthorization)
	expectedAuth := HTTPHeaderValueBasicPrefix + ""
	testutil.AssertStringEquals(t, auth, expectedAuth, "Authorization with empty token")

	// Other headers should still be set
	testutil.AssertStringNotEmpty(t, headers.Get(HTTPHeaderKeyAccept), "Accept header")

	testutil.AssertStringNotEmpty(t, headers.Get(HTTPHeaderKeyUserAgent), "User-Agent header")
}

func TestTransportUnit_NewTransportDetailsConfiguration_Success(t *testing.T) {
	transport := NewTransport(true)

	// Test boolean flags
	testutil.AssertTrue(t, transport.TLSClientConfig.InsecureSkipVerify, "InsecureSkipVerify")
	testutil.AssertFalse(t, transport.ForceAttemptHTTP2, "ForceAttemptHTTP2")
	testutil.AssertFalse(t, transport.DisableKeepAlives, "DisableKeepAlives")
	testutil.AssertFalse(t, transport.DisableCompression, "DisableCompression")

	// Test numeric configurations
	testutil.AssertIntEquals(t, transport.MaxIdleConns, DefaultMaxIdleConns, "MaxIdleConns")
	testutil.AssertIntEquals(
		t,
		transport.MaxIdleConnsPerHost,
		DefaultMaxIdleConnsPerHost,
		"MaxIdleConnsPerHost",
	)
}

func TestTransportUnit_DefaultHeadersUserAgent_Success(t *testing.T) {
	testutil.AssertStringEquals(t, HTTPHeaderUserAgent, "cisco-ios-xe-wireless-go",
		"the default User-Agent names the module")

	headers := DefaultHeaders("test-token-123", "cisco-wnc-exporter/1.2.3")
	testutil.AssertStringEquals(t, headers.Get(HTTPHeaderKeyUserAgent),
		"cisco-wnc-exporter/1.2.3", "custom User-Agent header")

	fallback := DefaultHeaders("test-token-123", "")
	testutil.AssertStringEquals(t, fallback.Get(HTTPHeaderKeyUserAgent),
		HTTPHeaderUserAgent, "User-Agent falls back to the constant")
}

func TestTransportUnit_NewTransportDialContext_Success(t *testing.T) {
	tr := NewTransport(false)
	testutil.AssertTrue(t, tr.DialContext != nil, "DialContext")
	testutil.AssertDurationEquals(t, DefaultDialTimeout, 30*time.Second, "DefaultDialTimeout")
	testutil.AssertDurationEquals(t, DefaultDialKeepAlive, 30*time.Second, "DefaultDialKeepAlive")
}

// TestTransportUnit_ConstantValues_Success pins the value behind every transport default, not the
// name. The assertions elsewhere in this file compare a built transport's field against the same
// constant that set it, so they hold whatever the constant becomes: changing QuickTimeout from 5 s
// to 500 s leaves them all green. These fail instead.
func TestTransportUnit_ConstantValues_Success(t *testing.T) {
	testutil.AssertDurationEquals(t, QuickTimeout, 5*time.Second, "QuickTimeout")
	testutil.AssertDurationEquals(t, StandardTimeout, 30*time.Second, "StandardTimeout")
	testutil.AssertDurationEquals(t, ExtendedTimeout, 90*time.Second, "ExtendedTimeout")

	testutil.AssertDurationEquals(t, DefaultTLSHandshakeTimeout, 5*time.Second,
		"DefaultTLSHandshakeTimeout")
	testutil.AssertDurationEquals(t, DefaultResponseHeaderTimeout, 5*time.Second,
		"DefaultResponseHeaderTimeout")
	testutil.AssertDurationEquals(t, DefaultIdleConnTimeout, 90*time.Second,
		"DefaultIdleConnTimeout")

	testutil.AssertIntEquals(t, DefaultMaxIdleConns, 100, "DefaultMaxIdleConns")
	testutil.AssertIntEquals(t, DefaultMaxIdleConnsPerHost, 10, "DefaultMaxIdleConnsPerHost")

	testutil.AssertStringEquals(t, HTTPHeaderValueYANGData, "application/yang-data+json",
		"HTTPHeaderValueYANGData")
	testutil.AssertStringEquals(t, HTTPHeaderValueBasicPrefix, "Basic ",
		"HTTPHeaderValueBasicPrefix")
}
