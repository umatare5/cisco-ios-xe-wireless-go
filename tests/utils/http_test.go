package testutils

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
)

func TestNewMockHTTPServer(t *testing.T) {
	mock := NewMockHTTPServer()
	defer mock.Close()

	if mock.Server == nil {
		t.Error("Expected server to be created, got nil")
	}

	// Test that URL is properly set
	if mock.Server.URL == "" {
		t.Error("Expected server URL to be set")
	}

	// Test that Handlers map is initialized
	if mock.Handlers == nil {
		t.Error("Expected handlers map to be initialized")
	}

	// Test the default handler behavior
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/non-existent", nil)

	// Call the default handler directly
	mock.Server.Config.Handler.ServeHTTP(w, r)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status %d for non-existent path, got %d", http.StatusNotFound, w.Code)
	}
}

func TestMockHTTPServer_AddHandler(t *testing.T) {
	mock := NewMockHTTPServer()
	defer mock.Close()

	testPath := "/test"
	testResponse := "test response"

	mock.AddHandler(testPath, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(testResponse))
	})

	if len(mock.Handlers) != 1 {
		t.Errorf("Expected 1 handler, got %d", len(mock.Handlers))
	}

	if _, exists := mock.Handlers[testPath]; !exists {
		t.Errorf("Handler for path %s not found", testPath)
	}
}

func TestCreateTestClientForMockServer(t *testing.T) {
	mock := NewMockHTTPServer()
	defer mock.Close()

	client := CreateTestClientForMockServer(t, mock)
	if client == nil {
		t.Error("Expected client to be created, got nil")
	}
}

func TestCreateJSONResponse(t *testing.T) {
	response := TestHTTPResponse{
		StatusCode: http.StatusOK,
		Body:       `{"test": "data"}`,
		Headers:    map[string]string{"X-Test": "value"},
	}

	handler := CreateJSONResponse(response)
	if handler == nil {
		t.Error("Expected handler to be created, got nil")
	}

	// Test the handler
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/test", nil)

	handler(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

	if w.Header().Get("X-Test") != "value" {
		t.Errorf("Expected header X-Test: value, got %s", w.Header().Get("X-Test"))
	}
}

func TestCreateJSONResponseFromStruct(t *testing.T) {
	// Test creating a JSON response for struct data
	response := TestHTTPResponse{
		StatusCode: http.StatusOK,
		Body:       `{"name":"test","id":123}`,
		Headers:    map[string]string{"Content-Type": "application/json"},
	}

	handler := CreateJSONResponse(response)
	if handler == nil {
		t.Error("Expected handler to be created, got nil")
	}

	// Test the handler
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/test", nil)

	handler(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

	var result struct {
		Name string `json:"name"`
		ID   int    `json:"id"`
	}
	err := json.Unmarshal(w.Body.Bytes(), &result)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if result.Name != "test" {
		t.Errorf("Expected name 'test', got %s", result.Name)
	}
}

func TestCreateErrorResponse(t *testing.T) {
	// Create an error response manually since CreateErrorResponse doesn't exist
	response := TestHTTPResponse{
		StatusCode: http.StatusBadRequest,
		Body:       `{"error":"test error"}`,
		Headers:    map[string]string{"Content-Type": "application/json"},
	}

	handler := CreateJSONResponse(response)
	if handler == nil {
		t.Error("Expected handler to be created, got nil")
	}

	// Test the handler
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/test", nil)

	handler(w, r)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status %d, got %d", http.StatusBadRequest, w.Code)
	}

	var result map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &result)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if result["error"] != "test error" {
		t.Errorf("Expected error 'test error', got %v", result["error"])
	}
}

func TestTestAPIFunction(t *testing.T) {
	TestAPIFunction(t, "BasicAPITest",
		func(mock *MockHTTPServer) {
			mock.AddHandler("/restconf/data/test", CreateJSONResponse(TestHTTPResponse{
				StatusCode: http.StatusOK,
				Body:       `{"result": "success"}`,
			}))
		},
		func(client *wnc.Client) error {
			// This would be a real API call in actual usage
			return nil
		},
	)
}

func TestTestAPIFunctionWithContext(t *testing.T) {
	TestAPIFunctionWithContext(t, "ContextAPITest",
		func(mock *MockHTTPServer) {
			mock.AddHandler("/restconf/data/test", CreateJSONResponse(TestHTTPResponse{
				StatusCode: http.StatusOK,
				Body:       `{"result": "success"}`,
			}))
		},
		func(ctx context.Context, client *wnc.Client) error {
			// This would be a real API call with context in actual usage
			return nil
		},
	)
}

func TestTestTimeout(t *testing.T) {
	TestTimeout(t, "TimeoutTest",
		func(ctx context.Context, client *wnc.Client) error {
			// Simulate a slow operation
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(100 * time.Millisecond):
				return nil
			}
		},
	)
}

func TestHTTPResponseHeaders(t *testing.T) {
	response := TestHTTPResponse{
		StatusCode: http.StatusOK,
		Body:       `{"test": "data"}`,
		Headers:    map[string]string{"X-Test": "value", "X-Custom": "custom"},
	}

	handler := CreateJSONResponse(response)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/test", nil)

	handler(w, r)

	if w.Header().Get("X-Test") != "value" {
		t.Errorf("Expected header X-Test: value, got %s", w.Header().Get("X-Test"))
	}

	if w.Header().Get("X-Custom") != "custom" {
		t.Errorf("Expected header X-Custom: custom, got %s", w.Header().Get("X-Custom"))
	}

	if w.Header().Get("Content-Type") != "application/json" {
		t.Errorf("Expected Content-Type: application/json, got %s", w.Header().Get("Content-Type"))
	}
}
