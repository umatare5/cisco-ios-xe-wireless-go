package wnc

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"sync"
	"testing"
)

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
		_, _ = w.Write([]byte(`{}`))
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
