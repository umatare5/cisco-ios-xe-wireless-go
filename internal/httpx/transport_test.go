package httpx

import (
	"net/http"
	"testing"
	"time"
)

func TestNewTransport(t *testing.T) {
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

			if transport == nil {
				t.Fatal("NewTransport returned nil")
			}

			// Check TLS configuration
			if transport.TLSClientConfig == nil {
				t.Error("TLSClientConfig is nil")
			} else if transport.TLSClientConfig.InsecureSkipVerify != tt.skipVerify {
				t.Errorf("InsecureSkipVerify = %v, want %v", transport.TLSClientConfig.InsecureSkipVerify, tt.skipVerify)
			}

			// Check timeout settings
			if transport.TLSHandshakeTimeout != DefaultTLSHandshakeTimeout {
				t.Errorf("TLSHandshakeTimeout = %v, want %v", transport.TLSHandshakeTimeout, DefaultTLSHandshakeTimeout)
			}

			if transport.ResponseHeaderTimeout != DefaultResponseHeaderTimeout {
				t.Errorf(
					"ResponseHeaderTimeout = %v, want %v",
					transport.ResponseHeaderTimeout,
					DefaultResponseHeaderTimeout,
				)
			}

			if transport.IdleConnTimeout != DefaultIdleConnTimeout {
				t.Errorf("IdleConnTimeout = %v, want %v", transport.IdleConnTimeout, DefaultIdleConnTimeout)
			}

			// Check connection settings
			if transport.ForceAttemptHTTP2 {
				t.Error("ForceAttemptHTTP2 should be false")
			}

			if transport.DisableKeepAlives {
				t.Error("DisableKeepAlives should be false")
			}

			if transport.DisableCompression {
				t.Error("DisableCompression should be false")
			}

			if transport.MaxIdleConns != 100 {
				t.Errorf("MaxIdleConns = %d, want 100", transport.MaxIdleConns)
			}

			if transport.MaxIdleConnsPerHost != 10 {
				t.Errorf("MaxIdleConnsPerHost = %d, want 10", transport.MaxIdleConnsPerHost)
			}
		})
	}
}

func TestDefaultHeaders(t *testing.T) {
	token := "test-token-123"
	headers := DefaultHeaders(token)

	if headers == nil {
		t.Fatal("DefaultHeaders returned nil")
	}

	// Check Authorization header
	expectedAuth := HTTPHeaderValueBasicPrefix + token
	if auth := headers.Get(HTTPHeaderKeyAuthorization); auth != expectedAuth {
		t.Errorf("Authorization header = %q, want %q", auth, expectedAuth)
	}

	// Check Accept header
	if accept := headers.Get(HTTPHeaderKeyAccept); accept != HTTPHeaderValueYANGData {
		t.Errorf("Accept header = %q, want %q", accept, HTTPHeaderValueYANGData)
	}

	// Check User-Agent header
	if userAgent := headers.Get(HTTPHeaderKeyUserAgent); userAgent != HTTPHeaderUserAgent {
		t.Errorf("User-Agent header = %q, want %q", userAgent, HTTPHeaderUserAgent)
	}
}

func TestHTTPConstants(t *testing.T) {
	// Test timeout constants
	if DefaultTLSHandshakeTimeout != 10*time.Second {
		t.Errorf("DefaultTLSHandshakeTimeout = %v, want 10s", DefaultTLSHandshakeTimeout)
	}

	if DefaultResponseHeaderTimeout != 10*time.Second {
		t.Errorf("DefaultResponseHeaderTimeout = %v, want 10s", DefaultResponseHeaderTimeout)
	}

	if DefaultIdleConnTimeout != 90*time.Second {
		t.Errorf("DefaultIdleConnTimeout = %v, want 90s", DefaultIdleConnTimeout)
	}

	// Test header key constants
	expectedHeaderKeys := map[string]string{
		"HTTPHeaderKeyAuthorization": "Authorization",
		"HTTPHeaderKeyAccept":        "Accept",
		"HTTPHeaderKeyUserAgent":     "User-Agent",
		"HTTPHeaderKeyContentType":   "Content-Type",
	}

	actualHeaderKeys := map[string]string{
		"HTTPHeaderKeyAuthorization": HTTPHeaderKeyAuthorization,
		"HTTPHeaderKeyAccept":        HTTPHeaderKeyAccept,
		"HTTPHeaderKeyUserAgent":     HTTPHeaderKeyUserAgent,
		"HTTPHeaderKeyContentType":   HTTPHeaderKeyContentType,
	}

	for name, expected := range expectedHeaderKeys {
		if actual := actualHeaderKeys[name]; actual != expected {
			t.Errorf("%s = %q, want %q", name, actual, expected)
		}
	}

	// Test header value constants
	if HTTPHeaderValueBasicPrefix != "Basic " {
		t.Errorf("HTTPHeaderValueBasicPrefix = %q, want %q", HTTPHeaderValueBasicPrefix, "Basic ")
	}

	if HTTPHeaderValueYANGData != "application/yang-data+json" {
		t.Errorf("HTTPHeaderValueYANGData = %q, want %q", HTTPHeaderValueYANGData, "application/yang-data+json")
	}

	if HTTPHeaderUserAgent != "wnc-go-client/1.0" {
		t.Errorf("HTTPHeaderUserAgent = %q, want %q", HTTPHeaderUserAgent, "wnc-go-client/1.0")
	}

	if HTTPHeaderAccept != HTTPHeaderValueYANGData {
		t.Errorf("HTTPHeaderAccept = %q, want %q", HTTPHeaderAccept, HTTPHeaderValueYANGData)
	}

	if HTTPHeaderContentType != HTTPHeaderValueYANGData {
		t.Errorf("HTTPHeaderContentType = %q, want %q", HTTPHeaderContentType, HTTPHeaderValueYANGData)
	}
}

func TestTransportConfiguration(t *testing.T) {
	transport := NewTransport(false)

	// Verify the transport implements http.RoundTripper
	var _ http.RoundTripper = transport

	// Check that TLS config is properly set
	tlsConfig := transport.TLSClientConfig
	if tlsConfig == nil {
		t.Fatal("TLS config should not be nil")
	}

	// Test with different TLS settings
	secureTransport := NewTransport(false)
	insecureTransport := NewTransport(true)

	if secureTransport.TLSClientConfig.InsecureSkipVerify {
		t.Error("Secure transport should have InsecureSkipVerify = false")
	}

	if !insecureTransport.TLSClientConfig.InsecureSkipVerify {
		t.Error("Insecure transport should have InsecureSkipVerify = true")
	}
}

func TestDefaultHeadersWithEmptyToken(t *testing.T) {
	headers := DefaultHeaders("")

	auth := headers.Get(HTTPHeaderKeyAuthorization)
	expectedAuth := HTTPHeaderValueBasicPrefix + ""
	if auth != expectedAuth {
		t.Errorf("Authorization with empty token = %q, want %q", auth, expectedAuth)
	}

	// Other headers should still be set
	if headers.Get(HTTPHeaderKeyAccept) == "" {
		t.Error("Accept header should not be empty")
	}

	if headers.Get(HTTPHeaderKeyUserAgent) == "" {
		t.Error("User-Agent header should not be empty")
	}
}

func TestHeaderManipulation(t *testing.T) {
	token := "test-token"
	headers := DefaultHeaders(token)

	// Test that headers can be modified
	headers.Set(HTTPHeaderKeyContentType, "custom-type")
	if ct := headers.Get(HTTPHeaderKeyContentType); ct != "custom-type" {
		t.Errorf("Content-Type after modification = %q, want %q", ct, "custom-type")
	}

	// Test that original headers are still present
	expectedAuth := HTTPHeaderValueBasicPrefix + token
	if auth := headers.Get(HTTPHeaderKeyAuthorization); auth != expectedAuth {
		t.Errorf("Authorization after modification = %q, want %q", auth, expectedAuth)
	}
}
