package transport

import (
	"net/http"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/helper"
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

			helper.AssertNotNil(t, transport, "NewTransport result")

			// Check TLS configuration
			helper.AssertNotNil(t, transport.TLSClientConfig, "TLSClientConfig")
			helper.AssertBoolEquals(t, transport.TLSClientConfig.InsecureSkipVerify, tt.skipVerify,
				"TLSClientConfig.InsecureSkipVerify")

			// Check timeout settings
			helper.AssertDurationEquals(t, transport.TLSHandshakeTimeout,
				DefaultTLSHandshakeTimeout, "TLSHandshakeTimeout")

			helper.AssertDurationEquals(t, transport.ResponseHeaderTimeout,
				DefaultResponseHeaderTimeout, "ResponseHeaderTimeout")

			helper.AssertDurationEquals(t, transport.IdleConnTimeout, DefaultIdleConnTimeout, "IdleConnTimeout")

			// Check connection settings
			helper.AssertFalse(t, transport.ForceAttemptHTTP2, "ForceAttemptHTTP2")

			helper.AssertFalse(t, transport.DisableKeepAlives, "DisableKeepAlives")

			helper.AssertFalse(t, transport.DisableCompression, "DisableCompression")

			helper.AssertIntEquals(t, transport.MaxIdleConns, 100, "MaxIdleConns")

			helper.AssertIntEquals(t, transport.MaxIdleConnsPerHost, 10, "MaxIdleConnsPerHost")
		})
	}
}

func TestTransportUnit_DefaultHeaders_Success(t *testing.T) {
	token := "test-token-123"
	headers := DefaultHeaders(token)

	helper.AssertNotNil(t, headers, "DefaultHeaders result")

	// Check Authorization header
	expectedAuth := HTTPHeaderValueBasicPrefix + token
	auth := headers.Get(HTTPHeaderKeyAuthorization)
	helper.AssertStringEquals(t, auth, expectedAuth, "Authorization header")

	// Check Accept header
	accept := headers.Get(HTTPHeaderKeyAccept)
	helper.AssertStringEquals(t, accept, HTTPHeaderValueYANGData, "Accept header")

	// Check User-Agent header
	userAgent := headers.Get(HTTPHeaderKeyUserAgent)
	helper.AssertStringEquals(t, userAgent, HTTPHeaderUserAgent, "User-Agent header")
}

func TestTransportUnit_TransportConfiguration_Success(t *testing.T) {
	transport := NewTransport(false)

	// Verify the transport implements http.RoundTripper
	var _ http.RoundTripper = transport

	// Check that TLS config is properly set
	helper.AssertNotNil(t, transport.TLSClientConfig, "TLS config")

	// Test with different TLS settings
	secureTransport := NewTransport(false)
	insecureTransport := NewTransport(true)

	helper.AssertFalse(t, secureTransport.TLSClientConfig.InsecureSkipVerify, "Secure transport InsecureSkipVerify")

	helper.AssertTrue(t, insecureTransport.TLSClientConfig.InsecureSkipVerify, "Insecure transport InsecureSkipVerify")
}

func TestTransportUnit_DefaultHeadersWithEmptyToken_Success(t *testing.T) {
	headers := DefaultHeaders("")

	auth := headers.Get(HTTPHeaderKeyAuthorization)
	expectedAuth := HTTPHeaderValueBasicPrefix + ""
	helper.AssertStringEquals(t, auth, expectedAuth, "Authorization with empty token")

	// Other headers should still be set
	helper.AssertStringNotEmpty(t, headers.Get(HTTPHeaderKeyAccept), "Accept header")

	helper.AssertStringNotEmpty(t, headers.Get(HTTPHeaderKeyUserAgent), "User-Agent header")
}

func TestTransportUnit_HeaderManipulation_Success(t *testing.T) {
	token := "test-token"
	headers := DefaultHeaders(token)

	// Test that headers can be modified
	headers.Set(HTTPHeaderKeyContentType, "custom-type")
	ct := headers.Get(HTTPHeaderKeyContentType)
	helper.AssertStringEquals(t, ct, "custom-type", "Content-Type after modification")

	// Test that original headers are still present
	expectedAuth := HTTPHeaderValueBasicPrefix + token
	auth := headers.Get(HTTPHeaderKeyAuthorization)
	helper.AssertStringEquals(t, auth, expectedAuth, "Authorization after modification")
}

func TestTransportUnit_NewTransportDetailsConfiguration_Success(t *testing.T) {
	transport := NewTransport(true)

	// Test boolean flags
	helper.AssertTrue(t, transport.TLSClientConfig.InsecureSkipVerify, "InsecureSkipVerify")
	helper.AssertFalse(t, transport.ForceAttemptHTTP2, "ForceAttemptHTTP2")
	helper.AssertFalse(t, transport.DisableKeepAlives, "DisableKeepAlives")
	helper.AssertFalse(t, transport.DisableCompression, "DisableCompression")

	// Test numeric configurations
	helper.AssertIntEquals(t, transport.MaxIdleConns, DefaultMaxIdleConns, "MaxIdleConns")
	helper.AssertIntEquals(t, transport.MaxIdleConnsPerHost, DefaultMaxIdleConnsPerHost, "MaxIdleConnsPerHost")
}
