// Package tests provides HTTP testing utilities for tests.
package tests

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
)

// MockHTTPServer creates a test HTTP server with predefined responses.
type MockHTTPServer struct {
	Server   *httptest.Server
	Handlers map[string]http.HandlerFunc
}

// NewMockHTTPServer creates a new mock HTTP server.
func NewMockHTTPServer() *MockHTTPServer {
	mock := &MockHTTPServer{
		Handlers: make(map[string]http.HandlerFunc),
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if handler, exists := mock.Handlers[r.URL.Path]; exists {
			handler(w, r)
			return
		}
		http.NotFound(w, r)
	})

	mock.Server = httptest.NewTLSServer(mux)
	return mock
}

// AddHandler adds a handler for a specific path.
func (m *MockHTTPServer) AddHandler(path string, handler http.HandlerFunc) {
	m.Handlers[path] = handler
}

// Close closes the mock server.
func (m *MockHTTPServer) Close() {
	if m.Server != nil {
		m.Server.Close()
	}
}

// CreateTestClientForMockServer creates a test client configured for the mock server.
func CreateTestClientForMockServer(t *testing.T, server *MockHTTPServer) *wnc.Client {
	t.Helper()

	if server == nil || server.Server == nil {
		t.Fatal("Mock server is nil")
	}

	config := wnc.Config{
		Controller:         server.Server.URL[8:], // Remove "https://" prefix
		AccessToken:        "dGVzdDp0ZXN0",        // "test:test" in base64
		Timeout:            5 * time.Second,
		InsecureSkipVerify: true,
	}

	client, err := wnc.NewClient(config)
	if err != nil {
		t.Fatalf("Failed to create test client: %v", err)
	}

	return client
}

// TestHTTPResponse represents expected HTTP response for testing.
type TestHTTPResponse struct {
	StatusCode int
	Body       string
	Headers    map[string]string
}

// CreateJSONResponse creates a JSON response handler.
func CreateJSONResponse(response TestHTTPResponse) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Set headers
		w.Header().Set("Content-Type", "application/json")
		for key, value := range response.Headers {
			w.Header().Set(key, value)
		}

		// Set status code
		w.WriteHeader(response.StatusCode)

		// Write body
		if response.Body != "" {
			w.Write([]byte(response.Body))
		}
	}
}

// TestAPIFunction tests a function that makes API calls with a mock server.
func TestAPIFunction(t *testing.T, testName string, setupMock func(*MockHTTPServer), testFunc func(*wnc.Client) error) {
	t.Helper()

	t.Run(testName, func(t *testing.T) {
		mock := NewMockHTTPServer()
		defer mock.Close()

		setupMock(mock)
		client := CreateTestClientForMockServer(t, mock)

		err := testFunc(client)
		if err != nil {
			t.Errorf("API function failed: %v", err)
		}
	})
}

// TestAPIFunctionWithContext tests a function that makes API calls with context.
func TestAPIFunctionWithContext(t *testing.T, testName string, setupMock func(*MockHTTPServer), testFunc func(context.Context, *wnc.Client) error) {
	t.Helper()

	t.Run(testName, func(t *testing.T) {
		mock := NewMockHTTPServer()
		defer mock.Close()

		setupMock(mock)
		client := CreateTestClientForMockServer(t, mock)

		ctx := context.Background()
		err := testFunc(ctx, client)
		if err != nil {
			t.Errorf("API function with context failed: %v", err)
		}
	})
}

// TestTimeoutAPI tests function behavior with timeout scenarios.
func TestTimeoutAPI(t *testing.T, testName string, testFunc func(context.Context, *wnc.Client) error) {
	t.Helper()

	t.Run(testName, func(t *testing.T) {
		mock := NewMockHTTPServer()
		defer mock.Close()

		// Add a handler that delays response
		mock.AddHandler("/slow", func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(200 * time.Millisecond)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"result": "success"}`))
		})

		client := CreateTestClientForMockServer(t, mock)

		// Create a context with very short timeout
		ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
		defer cancel()

		err := testFunc(ctx, client)
		if err == nil {
			t.Error("Expected timeout error, but got nil")
		}
	})
}
