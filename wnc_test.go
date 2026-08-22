package wnc

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"net/url"
	"sync"
	"testing"
	"time"

	mock "github.com/umatare5/cisco-ios-xe-wireless-go/pkg/testutil"
)

// TestNewClient tests the creation of a new unified client.
func TestNewClient(t *testing.T) {
	testCases := []struct {
		name        string
		host        string
		token       string
		opts        []Option
		expectError bool
	}{
		{
			name:        "ValidClient",
			host:        "192.168.1.100",
			token:       "YWRtaW46cGFzc3dvcmQ=", // base64 encoded "admin:password"
			opts:        nil,
			expectError: false,
		},
		{
			name:        "ValidClientWithOptions",
			host:        "wnc.example.internal",
			token:       "YWRtaW46cGFzc3dvcmQ=",
			opts:        []Option{WithTimeout(30 * time.Second), WithInsecureSkipVerify(true)},
			expectError: false,
		},
		{
			name:        "ValidClientWithLoggerAndUserAgent",
			host:        "controller.example.internal",
			token:       "YWRtaW46cGFzc3dvcmQ=",
			opts:        []Option{WithLogger(slog.New(slog.DiscardHandler)), WithUserAgent("custom-agent/1.0")},
			expectError: false,
		},
		{
			name:        "InvalidHost",
			host:        "",
			token:       "YWRtaW46cGFzc3dvcmQ=",
			opts:        nil,
			expectError: true,
		},
		{
			name:        "InvalidToken",
			host:        "controller.example.com",
			token:       "",
			opts:        nil,
			expectError: true,
		},
		{
			name:  "ValidClientWithTransportOptions",
			host:  "controller.example.internal",
			token: "test-token-123",
			opts: []Option{
				WithProxy(http.ProxyFromEnvironment),
				WithResponseHeaderTimeout(15 * time.Second),
				WithTLSHandshakeTimeout(10 * time.Second),
			},
			expectError: false,
		},
		{
			name:  "RejectsNonPositiveResponseHeaderTimeout",
			host:  "controller.example.internal",
			token: "test-token-123",
			opts: []Option{
				WithResponseHeaderTimeout(0),
			},
			expectError: true,
		},
		{
			name:  "RejectsNonPositiveTLSHandshakeTimeout",
			host:  "controller.example.internal",
			token: "test-token-123",
			opts: []Option{
				WithTLSHandshakeTimeout(-1 * time.Second),
			},
			expectError: true,
		},
		{
			name:  "NilProxyResolverIsAccepted",
			host:  "controller.example.internal",
			token: "test-token-123",
			opts: []Option{
				WithProxy(nil),
			},
			expectError: false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			client, err := NewClient(tt.host, tt.token, tt.opts...)

			if tt.expectError {
				if err == nil {
					t.Error("Expected error, but got none")
				}
				if client != nil {
					t.Error("Expected nil client on error")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if client == nil {
					t.Error("Expected client, but got nil")
				}
			}
		})
	}
}

// TestClientServiceAccessors tests that all service accessors return non-nil services.
func TestClientServiceAccessors(t *testing.T) {
	client, err := NewClient("controller.example.com", "dGVzdDp0ZXN0")
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	// Test all service accessors - verify they don't panic and return valid structs
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Service accessor panicked: %v", r)
		}
	}()

	// Test service accessors return valid services
	_ = client.AFC()           // Should not panic
	_ = client.AP()            // Should not panic
	_ = client.APF()           // Should not panic
	_ = client.AWIPS()         // Should not panic
	_ = client.BLE()           // Should not panic
	_ = client.Client()        // Should not panic
	_ = client.Controller()    // Should not panic
	_ = client.CTS()           // Should not panic
	_ = client.Dot11()         // Should not panic
	_ = client.Dot15()         // Should not panic
	_ = client.Fabric()        // Should not panic
	_ = client.Flex()          // Should not panic
	_ = client.General()       // Should not panic
	_ = client.Geolocation()   // Should not panic
	_ = client.Hyperlocation() // Should not panic
	_ = client.LISP()          // Should not panic
	_ = client.Location()      // Should not panic
	_ = client.Mcast()         // Should not panic
	_ = client.MDNS()          // Should not panic
	_ = client.Mesh()          // Should not panic
	_ = client.Mobility()      // Should not panic
	_ = client.NMSP()          // Should not panic
	_ = client.Radio()         // Should not panic
	_ = client.RF()            // Should not panic
	_ = client.RFID()          // Should not panic
	_ = client.Rogue()         // Should not panic
	_ = client.RRM()           // Should not panic
	_ = client.Site()          // Should not panic
	_ = client.Spaces()        // Should not panic
	_ = client.URWB()          // Should not panic
	_ = client.WAT()           // Should not panic
	_ = client.WLAN()          // Should not panic

	// Test tag service accessors
	_ = client.PolicyTag() // Should not panic
	_ = client.RFTag()     // Should not panic
	_ = client.SiteTag()   // Should not panic
}

// newRecordingClient returns a client pointed at a server that records every request it
// answers, which is the only way a query parameter is observable end to end.
func newRecordingClient(t *testing.T, body string) (*Client, *mock.RESTCONFServer) {
	t.Helper()

	server := mock.NewRESTCONFServer(t)
	t.Cleanup(server.Close)
	server.AddHandler(http.MethodGet, "probe", func() (int, string) { return http.StatusOK, body })

	parsed, err := url.Parse(server.URL)
	if err != nil {
		t.Fatalf("Failed to parse test server URL: %v", err)
	}

	client, err := NewClient(parsed.Host, "dGVzdDp0ZXN0", WithInsecureSkipVerify(true))
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	return client, server
}

// TestGetDataReturnsBodyUnchecked pins that GetData hands back the envelope intact. The
// body carries two top-level keys, which no typed accessor would have accepted.
func TestGetDataReturnsBodyUnchecked(t *testing.T) {
	const body = `{"a:one":{"x":1},"a:two":{"y":2}}`

	client, _ := newRecordingClient(t, body)

	got, err := client.GetData(context.Background(), "probe")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if string(got) != body {
		t.Errorf("GetData() body = %q, want %q", got, body)
	}
}

// TestGetDataWireQuery pins the re-exported options on the wire through the unified
// client. WithFields and WithDepth have no other route to a consumer.
func TestGetDataWireQuery(t *testing.T) {
	client, server := newRecordingClient(t, `{"a:probe":{}}`)

	testCases := []struct {
		name     string
		opts     []GetOption
		expected string
	}{
		{name: "without option", opts: nil, expected: ""},
		{
			name:     "with WithFields",
			opts:     []GetOption{WithFields("probe(one;two)")},
			expected: "fields=probe(one;two)",
		},
		{name: "with WithDepth", opts: []GetOption{WithDepth(3)}, expected: "depth=3"},
		{name: "with WithDepth below 1", opts: []GetOption{WithDepth(0)}, expected: ""},
		{
			name:     "with all three",
			opts:     []GetOption{WithDefaults(ReportAll), WithFields("probe(one;two)"), WithDepth(3)},
			expected: "with-defaults=report-all&fields=probe(one;two)&depth=3",
		},
	}

	for i, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := client.GetData(context.Background(), "probe", tt.opts...); err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			recorded := server.Requests()
			if len(recorded) != i+1 {
				t.Fatalf("Recorded %d requests, want %d", len(recorded), i+1)
			}
			if got := recorded[i].RawQuery; got != tt.expected {
				t.Errorf("RawQuery = %q, want %q", got, tt.expected)
			}
		})
	}
}

// TestGetDataPathPrefix pins the claim in GetData's doc that the prefix is optional: both
// spellings reach the same wire path.
//
// The body assertion holds the sole-top-level-key shape the doc describes: an answer with
// exactly one key is returned whole, envelope included, not unwrapped to its value.
func TestGetDataPathPrefix(t *testing.T) {
	const body = `{"a:probe":{}}`

	client, server := newRecordingClient(t, body)

	for _, path := range []string{"probe", "/restconf/data/probe"} {
		got, err := client.GetData(context.Background(), path)
		if err != nil {
			t.Fatalf("GetData(%q) unexpected error: %v", path, err)
		}
		if string(got) != body {
			t.Errorf("GetData(%q) body = %q, want %q", path, got, body)
		}
	}

	recorded := server.Requests()
	if len(recorded) != 2 {
		t.Fatalf("Recorded %d requests, want 2", len(recorded))
	}
	for _, req := range recorded {
		if req.Path != "/restconf/data/probe" {
			t.Errorf("Wire path = %q, want %q", req.Path, "/restconf/data/probe")
		}
	}
}

// TestGetDataErrorIsAPIError pins that GetData reports a controller error the same way a
// typed accessor does, so errors.As on *APIError keeps working on the untyped route.
func TestGetDataErrorIsAPIError(t *testing.T) {
	client, _ := newRecordingClient(t, `{"a:probe":{}}`)

	_, err := client.GetData(context.Background(), "absent")
	if err == nil {
		t.Fatal("GetData() on an unhandled path should return an error")
	}

	var apiErr *APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("GetData() error = %v, want an *APIError", err)
	}
	if apiErr.StatusCode != http.StatusNotFound {
		t.Errorf("StatusCode = %d, want %d", apiErr.StatusCode, http.StatusNotFound)
	}
}

// TestGetDataPathIsSentAsGiven pins the escaping half of GetData's doc. A list key value
// carrying "#" or "?" is not escaped for the caller, so the path ends early and a
// different node answers — with no error to show for it. The typed accessors escape the
// key value themselves, which is why the doc names the owner.
func TestGetDataPathIsSentAsGiven(t *testing.T) {
	client, server := newRecordingClient(t, `{"a:probe":{}}`)

	testCases := []struct {
		name     string
		path     string
		wantPath string
		wantRaw  string
	}{
		{name: "plain", path: "probe=one", wantPath: "/restconf/data/probe=one"},
		{name: "fragment truncates", path: "probe=one#two", wantPath: "/restconf/data/probe=one"},
		{name: "query truncates", path: "probe=one?two", wantPath: "/restconf/data/probe=one", wantRaw: "two"},
		{
			name:     "escaped by the caller",
			path:     "probe=" + url.PathEscape("one#two"),
			wantPath: "/restconf/data/probe=one#two",
		},
	}

	for i, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := client.GetData(context.Background(), tt.path); err != nil {
				t.Fatalf("GetData(%q) unexpected error: %v", tt.path, err)
			}

			recorded := server.Requests()
			if len(recorded) != i+1 {
				t.Fatalf("Recorded %d requests, want %d", len(recorded), i+1)
			}
			if got := recorded[i].Path; got != tt.wantPath {
				t.Errorf("Wire path = %q, want %q", got, tt.wantPath)
			}
			if got := recorded[i].RawQuery; got != tt.wantRaw {
				t.Errorf("RawQuery = %q, want %q", got, tt.wantRaw)
			}
		})
	}
}

// TestWithDefaultsWireQuery tests that the re-exported option reaches the wire through the
// unified client, and that a call without it still sends no query at all.
func TestWithDefaultsWireQuery(t *testing.T) {
	var (
		mu       sync.Mutex
		rawQuery string
	)
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		rawQuery = r.URL.RawQuery
		mu.Unlock()
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-entries": {}}`))
	}))
	defer server.Close()

	parsed, err := url.Parse(server.URL)
	if err != nil {
		t.Fatalf("Failed to parse test server URL: %v", err)
	}

	client, err := NewClient(parsed.Host, "dGVzdDp0ZXN0", WithInsecureSkipVerify(true))
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	observed := func() string {
		mu.Lock()
		defer mu.Unlock()
		return rawQuery
	}

	testCases := []struct {
		name     string
		opts     []GetOption
		expected string
	}{
		{
			name:     "without option",
			opts:     nil,
			expected: "",
		},
		{
			name:     "with ReportAll",
			opts:     []GetOption{WithDefaults(ReportAll)},
			expected: "with-defaults=report-all",
		},
		{
			name:     "with Explicit",
			opts:     []GetOption{WithDefaults(Explicit)},
			expected: "with-defaults=explicit",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := client.WLAN().ListWlanCfgEntries(context.Background(), tt.opts...); err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}
			if got := observed(); got != tt.expected {
				t.Errorf("RawQuery = %q, want %q", got, tt.expected)
			}
		})
	}
}
