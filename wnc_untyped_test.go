package wnc

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

// seenRequest is what the recorder keeps: the three things the untyped methods are responsible
// for, and the body, which the shared mock server in pkg/testutil does not record.
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

// TestClientUntyped_Request_SendsTheMethodAsGiven holds the general method to its contract: the
// method reaches the wire unchanged, which is what makes an unforeseen verb reachable, and the
// path chooses its RESTCONF root.
func TestClientUntyped_Request_SendsTheMethodAsGiven(t *testing.T) {
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
