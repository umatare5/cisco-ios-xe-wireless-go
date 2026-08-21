package wnc

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"testing"

	mock "github.com/umatare5/cisco-ios-xe-wireless-go/pkg/testutil"
)

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
