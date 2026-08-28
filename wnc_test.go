package wnc

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
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

// probeEnvelope is the envelope shape GetDataInto takes: one field, tagged with the
// module-qualified node the path reads.
type probeEnvelope struct {
	Probe *struct {
		Leaf int `json:"leaf"`
	} `json:"a:probe"`
}

// wrongKeyEnvelope names a node the response does not carry.
type wrongKeyEnvelope struct {
	Other *struct{} `json:"a:other"`
}

// TestGetDataIntoAppliesTheEnvelopeCheck pins what GetDataInto adds over GetData. The two-key
// body is the same fixture GetData hands back intact, so the pair is the before and after.
func TestGetDataIntoAppliesTheEnvelopeCheck(t *testing.T) {
	t.Run("SoleKeyDecodes", func(t *testing.T) {
		client, _ := newRecordingClient(t, `{"a:probe":{"leaf":7}}`)

		got, err := GetDataInto[probeEnvelope](context.Background(), client, "probe")
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		if got.Probe == nil || got.Probe.Leaf != 7 {
			t.Errorf("Probe = %+v, want a node holding leaf 7", got.Probe)
		}
	})

	t.Run("TwoTopLevelKeysAreRefused", func(t *testing.T) {
		client, _ := newRecordingClient(t, `{"a:one":{"x":1},"a:two":{"y":2}}`)

		if _, err := GetDataInto[probeEnvelope](context.Background(), client, "probe"); err == nil {
			t.Error("Expected an error for a body carrying two top-level keys")
		}
	})

	t.Run("UnclaimedKeyIsRefused", func(t *testing.T) {
		client, _ := newRecordingClient(t, `{"a:probe":{"leaf":7}}`)

		if _, err := GetDataInto[wrongKeyEnvelope](context.Background(), client, "probe"); err == nil {
			t.Error("Expected an error for a T declaring no field for the response key")
		}
	})

	t.Run("NonStructIsRefused", func(t *testing.T) {
		client, _ := newRecordingClient(t, `{"a:probe":{"leaf":7}}`)

		if _, err := GetDataInto[map[string]any](context.Background(), client, "probe"); err == nil {
			t.Error("Expected an error for a non-struct T")
		}
	})

	t.Run("NilClientIsReported", func(t *testing.T) {
		if _, err := GetDataInto[probeEnvelope](context.Background(), nil, "probe"); err == nil {
			t.Error("Expected an error for a nil client")
		}
	})
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

// seenRequest is what the recorder keeps: the three things the untyped methods are responsible
// for, the body, and the Content-Type, which RecordedRequest deliberately does not carry.
type seenRequest struct {
	method      string
	path        string
	rawQuery    string
	body        string
	contentType string
}

// newRecorder starts a TLS server that answers every request with 200 and an empty body, and
// returns a client pointed at it beside the slice it records into.
func newRecorder(t *testing.T) (*Client, *[]seenRequest) {
	t.Helper()

	seen := &[]seenRequest{}

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		*seen = append(*seen, seenRequest{
			method:      r.Method,
			path:        r.URL.Path,
			rawQuery:    r.URL.RawQuery,
			body:        string(body),
			contentType: r.Header.Get("Content-Type"),
		})
		w.WriteHeader(http.StatusOK)
	}))
	t.Cleanup(server.Close)

	parsed, err := url.Parse(server.URL)
	if err != nil {
		t.Fatalf("parsing the recorder URL: %v", err)
	}

	client, err := NewClient(parsed.Host, "dGVzdDp0ZXN0", WithInsecureSkipVerify(true))
	if err != nil {
		t.Fatalf("building a client against the recorder: %v", err)
	}

	return client, seen
}

// last returns the request the recorder saw most recently.
func last(t *testing.T, seen *[]seenRequest) seenRequest {
	t.Helper()

	if len(*seen) == 0 {
		t.Fatal("the recorder saw no request")
	}

	return (*seen)[len(*seen)-1]
}

// TestClientUntyped_VerbMethods_ReachTheWire holds each verb method to the method and the RESTCONF
// root it names. The verb is the whole point of these methods: it is fixed at the call site rather
// than passed as a string, so a typo cannot reach the controller.
func TestClientUntyped_VerbMethods_ReachTheWire(t *testing.T) {
	client, seen := newRecorder(t)
	ctx := context.Background()

	tests := []struct {
		name     string
		call     func() ([]byte, error)
		method   string
		wantPath string
	}{
		{
			name:     "PostData",
			call:     func() ([]byte, error) { return client.PostData(ctx, "Cisco-IOS-XE-x:c", nil) },
			method:   http.MethodPost,
			wantPath: "/restconf/data/Cisco-IOS-XE-x:c",
		},
		{
			name:     "PutData",
			call:     func() ([]byte, error) { return client.PutData(ctx, "Cisco-IOS-XE-x:c", nil) },
			method:   http.MethodPut,
			wantPath: "/restconf/data/Cisco-IOS-XE-x:c",
		},
		{
			name:     "PatchData",
			call:     func() ([]byte, error) { return client.PatchData(ctx, "Cisco-IOS-XE-x:c", nil) },
			method:   http.MethodPatch,
			wantPath: "/restconf/data/Cisco-IOS-XE-x:c",
		},
		{
			name:     "DeleteData",
			call:     func() ([]byte, error) { return client.DeleteData(ctx, "Cisco-IOS-XE-x:c=key") },
			method:   http.MethodDelete,
			wantPath: "/restconf/data/Cisco-IOS-XE-x:c=key",
		},
		{
			name:     "PostRPC reaches the operations root",
			call:     func() ([]byte, error) { return client.PostRPC(ctx, "Cisco-IOS-XE-x:rpc", nil) },
			method:   http.MethodPost,
			wantPath: "/restconf/operations/Cisco-IOS-XE-x:rpc",
		},
		{
			name:     "an already-prefixed data path is not prefixed twice",
			call:     func() ([]byte, error) { return client.PutData(ctx, "/restconf/data/Cisco-IOS-XE-x:c", nil) },
			method:   http.MethodPut,
			wantPath: "/restconf/data/Cisco-IOS-XE-x:c",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := tt.call(); err != nil {
				t.Fatalf("Expected no error, got: %v", err)
			}

			got := last(t, seen)
			if got.method != tt.method {
				t.Errorf("Method mismatch: expected %s, got %s", tt.method, got.method)
			}
			if got.path != tt.wantPath {
				t.Errorf("Path mismatch: expected %s, got %s", tt.wantPath, got.path)
			}
		})
	}
}

// TestClientUntyped_Request_RoutesByPath holds the general method to its contract: the path chooses
// its RESTCONF root, the method reaches the wire unchanged on the data root, which is what makes an
// unforeseen verb reachable, and the operations root takes POST alone.
func TestClientUntyped_Request_RoutesByPath(t *testing.T) {
	client, seen := newRecorder(t)
	ctx := context.Background()

	t.Run("a bodiless HEAD reaches the data root", func(t *testing.T) {
		if _, err := client.Request(ctx, http.MethodHead, "/restconf/data/Cisco-IOS-XE-x:c", nil); err != nil {
			t.Fatalf("Expected no error, got: %v", err)
		}

		got := last(t, seen)
		if got.method != http.MethodHead {
			t.Errorf("Method mismatch: expected HEAD, got %s", got.method)
		}
		if got.body != "" {
			t.Errorf("Expected no body, got %q", got.body)
		}
	})

	t.Run("an operations path routes to the operations root", func(t *testing.T) {
		_, err := client.Request(ctx, http.MethodPost, "/restconf/operations/Cisco-IOS-XE-x:rpc", nil)
		if err != nil {
			t.Fatalf("Expected no error, got: %v", err)
		}

		if got := last(t, seen); got.path != "/restconf/operations/Cisco-IOS-XE-x:rpc" {
			t.Errorf("Path mismatch: got %s", got.path)
		}
	})

	t.Run("another method on an operations path is refused, not replaced", func(t *testing.T) {
		sent := len(*seen)

		if _, err := client.Request(ctx, http.MethodPut, "/restconf/operations/Cisco-IOS-XE-x:rpc", nil); err == nil {
			t.Fatal("Expected an error for PUT on an operations path, got nil")
		}

		// doRPC would have sent POST whatever the caller wrote, so a request that reached the wire
		// at all would have invoked the operation rather than replaced a node.
		if len(*seen) != sent {
			t.Errorf("Expected no request to reach the wire, got %s", last(t, seen).method)
		}
	})

	t.Run("a query the package has no option for is carried through", func(t *testing.T) {
		_, err := client.Request(ctx, http.MethodPost, "/restconf/data/Cisco-IOS-XE-x:c?insert=first", nil)
		if err != nil {
			t.Fatalf("Expected no error, got: %v", err)
		}

		if got := last(t, seen); got.rawQuery != "insert=first" {
			t.Errorf("Query mismatch: expected insert=first, got %q", got.rawQuery)
		}
	})
}

// TestClientUntyped_Payload_CarriesBytesVerbatim holds the payload rule that the whole hatch turns
// on. A body read with GetData has to survive being edited and sent back, and this platform sends
// a 64-bit counter as a bare number: decoding it into a Go value and re-encoding rounds it off, and
// handing the bytes to a parameter typed any base64s them. Only the verbatim path preserves it.
func TestClientUntyped_Payload_CarriesBytesVerbatim(t *testing.T) {
	client, seen := newRecorder(t)
	ctx := context.Background()

	const counter = `{"counter":18446744073709551615}`

	t.Run("bytes reach the wire unchanged", func(t *testing.T) {
		if _, err := client.PatchData(ctx, "x:c", []byte(counter)); err != nil {
			t.Fatalf("Expected no error, got: %v", err)
		}

		got := last(t, seen)
		if strings.TrimSpace(got.body) != counter {
			t.Errorf("Body mismatch: expected %s, got %s", counter, got.body)
		}
		if !strings.Contains(got.contentType, "json") {
			t.Errorf("Expected a JSON content type, got %q", got.contentType)
		}
	})

	t.Run("json.RawMessage reaches the wire unchanged", func(t *testing.T) {
		if _, err := client.PatchData(ctx, "x:c", json.RawMessage(counter)); err != nil {
			t.Fatalf("Expected no error, got: %v", err)
		}

		if got := last(t, seen); strings.TrimSpace(got.body) != counter {
			t.Errorf("Body mismatch: expected %s, got %s", counter, got.body)
		}
	})

	t.Run("a Go value is marshaled", func(t *testing.T) {
		if _, err := client.PostData(ctx, "x:c", map[string]string{"name": "a"}); err != nil {
			t.Fatalf("Expected no error, got: %v", err)
		}

		if got := last(t, seen); strings.TrimSpace(got.body) != `{"name":"a"}` {
			t.Errorf("Body mismatch: got %s", got.body)
		}
	})

	t.Run("a nil payload sends no body and no content type", func(t *testing.T) {
		if _, err := client.DeleteData(ctx, "x:c=key"); err != nil {
			t.Fatalf("Expected no error, got: %v", err)
		}

		got := last(t, seen)
		if got.body != "" {
			t.Errorf("Expected no body, got %q", got.body)
		}
		if got.contentType != "" {
			t.Errorf("Expected no content type, got %q", got.contentType)
		}
	})
}

// TestClientUntyped_Faults_AreReportedBeforeSending holds the two faults the hatch refuses rather
// than forwards. Both are checked before a request is built, so neither reaches the controller.
func TestClientUntyped_Faults_AreReportedBeforeSending(t *testing.T) {
	client, seen := newRecorder(t)
	ctx := context.Background()

	t.Run("an empty method is refused", func(t *testing.T) {
		before := len(*seen)

		if _, err := client.Request(ctx, "", "/restconf/data/x:c", nil); err == nil {
			t.Error("Expected an error for an empty method, got nil")
		}
		if len(*seen) != before {
			t.Error("Expected nothing to be sent for an empty method")
		}
	})

	t.Run("payload bytes that are not JSON are refused", func(t *testing.T) {
		before := len(*seen)

		if _, err := client.PatchData(ctx, "x:c", []byte("not json")); err == nil {
			t.Error("Expected an error for a non-JSON payload, got nil")
		}
		if len(*seen) != before {
			t.Error("Expected nothing to be sent for a non-JSON payload")
		}
	})

	t.Run("empty payload bytes are treated as no payload", func(t *testing.T) {
		if _, err := client.PatchData(ctx, "x:c", []byte{}); err != nil {
			t.Fatalf("Expected no error, got: %v", err)
		}

		if got := last(t, seen); got.body != "" {
			t.Errorf("Expected no body, got %q", got.body)
		}
	})
}

// TestClient_CloseIdleConnections_ReleasesThePool holds the one lever this package publishes over
// the connection pool, in the only terms a caller can observe it: whether the next request reuses
// the socket or dials a new one.
//
// Counting dials is what makes this more than a call that cannot fail. Without the close in the
// middle, the third request reuses the first connection and the count stays at 1.
func TestClient_CloseIdleConnections_ReleasesThePool(t *testing.T) {
	var (
		mu    sync.Mutex
		dials int
	)

	// Unstarted, because ConnState has to be in place before the first connection: setting it on
	// an already-serving httptest.Server races with the serving goroutine, which -race reports.
	server := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	}))
	server.Config.ConnState = func(_ net.Conn, state http.ConnState) {
		if state == http.StateNew {
			mu.Lock()
			dials++
			mu.Unlock()
		}
	}
	server.StartTLS()
	t.Cleanup(server.Close)

	count := func() int {
		mu.Lock()
		defer mu.Unlock()
		return dials
	}

	parsed, err := url.Parse(server.URL)
	if err != nil {
		t.Fatalf("parsing the server URL: %v", err)
	}

	client, err := NewClient(parsed.Host, "test-token-123", WithInsecureSkipVerify(true))
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}

	read := func(what string) {
		if _, err := client.GetData(context.Background(), "Cisco-IOS-XE-x:probe"); err != nil {
			t.Fatalf("%s: %v", what, err)
		}
	}

	read("first read")
	read("second read")

	if got := count(); got != 1 {
		t.Fatalf("dials after two reads = %d, want 1: the pool is not reusing the connection, so"+
			" this test cannot observe the close", got)
	}

	client.CloseIdleConnections()

	read("third read")

	if got := count(); got != 2 {
		t.Errorf("dials after the close and a third read = %d, want 2: the idle connection was"+
			" not released", got)
	}
}

// TestClient_TLSOptionWrappers_ReachTheCore holds the two one-line re-export wrappers, which would
// otherwise compile while pointing at the wrong core option and be exercised by nothing here.
func TestClient_TLSOptionWrappers_ReachTheCore(t *testing.T) {
	host := "wnc1.example.internal"

	if _, err := NewClient(host, "test-token-123", WithRootCAs(nil)); !errors.Is(err, ErrInvalidConfiguration) {
		t.Errorf("WithRootCAs(nil) error = %v, want ErrInvalidConfiguration", err)
	}

	_, err := NewClient(host, "test-token-123", WithClientCertificate(tls.Certificate{}))
	if !errors.Is(err, ErrInvalidConfiguration) {
		t.Errorf("WithClientCertificate(zero) error = %v, want ErrInvalidConfiguration", err)
	}

	if _, err := NewClient(host, "test-token-123", WithRootCAs(x509.NewCertPool())); err != nil {
		t.Errorf("WithRootCAs(pool) error = %v, want nil", err)
	}
}
