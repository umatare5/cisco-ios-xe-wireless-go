package transport

import (
	"context"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil"
)

func TestClientUnit_NewRequestBuilder_Success(t *testing.T) {
	t.Run("ValidParams", func(t *testing.T) {
		restBuilder := restconf.NewBuilder("https", "controller.example.com")
		rb := NewRequestBuilder(restBuilder, "token", slog.Default())
		testutil.AssertNotNil(t, rb, "NewRequestBuilder result")
	})

	t.Run("NilRESTCONFBuilder", func(t *testing.T) {
		rb := NewRequestBuilder(nil, "token", slog.Default())
		testutil.AssertNotNil(t, rb, "NewRequestBuilder result for nil RESTCONF builder")
	})
}

func TestClientUnit_RequestBuilderCreateRequestWithPayload_Success(t *testing.T) {
	t.Run("ValidPayload", func(t *testing.T) {
		restBuilder := restconf.NewBuilder("https", "controller.example.com")
		logger := slog.Default()
		rb := NewRequestBuilder(restBuilder, "token", logger)

		req, err := rb.CreateRequestWithPayload(
			context.Background(),
			http.MethodPost,
			"test/path",
			map[string]string{"test": "data"},
		)
		testutil.AssertNoError(t, err, "CreateRequestWithPayload")

		testutil.AssertStringEquals(t, req.Method, http.MethodPost, "HTTP method")
	})

	t.Run("NilRESTCONFBuilder", func(t *testing.T) {
		rb := NewRequestBuilder(nil, "token", slog.Default())
		_, err := rb.CreateRequestWithPayload(context.Background(), http.MethodGet, "test/path", nil)
		testutil.AssertError(t, err, "Expected error for nil RESTCONF builder")
		testutil.AssertStringContains(t, err.Error(), "RESTCONF builder is not properly initialized", "error message")
	})
}

func TestClientUnit_RequestBuilderCreateRPCRequestWithPayload_Success(t *testing.T) {
	t.Run("ValidRPCPayload", func(t *testing.T) {
		restBuilder := restconf.NewBuilder("https", "controller.example.com")
		logger := slog.Default()
		rb := NewRequestBuilder(restBuilder, "token", logger)

		req, err := rb.CreateRPCRequestWithPayload(
			context.Background(),
			http.MethodPost,
			"test/rpc",
			map[string]string{"test": "data"},
		)
		testutil.AssertNoError(t, err, "CreateRPCRequestWithPayload")

		testutil.AssertStringEquals(t, req.Method, http.MethodPost, "HTTP method")
	})

	t.Run("NilRESTCONFBuilder", func(t *testing.T) {
		rb := NewRequestBuilder(nil, "token", slog.Default())
		_, err := rb.CreateRPCRequestWithPayload(context.Background(), http.MethodPost, "test/rpc", nil)
		testutil.AssertError(t, err, "Expected error for nil RESTCONF builder")
		testutil.AssertStringContains(t, err.Error(), "RESTCONF builder is not properly initialized", "error message")
	})

	t.Run("InvalidHTTPMethodWithNilPayload", func(t *testing.T) {
		restBuilder := restconf.NewBuilder("https", "controller.example.com")
		logger := slog.Default()
		rb := NewRequestBuilder(restBuilder, "token", logger)

		// Test with invalid method and nil payload
		_, err := rb.CreateRPCRequestWithPayload(context.Background(), "\n", "test/rpc", nil)
		testutil.AssertError(t, err, "Expected error for invalid HTTP method")
		testutil.AssertStringContains(t, err.Error(), "failed to create RPC request", "error message")
	})
}

func TestClientUnit_RequestBuilderExecuteRequest_Success(t *testing.T) {
	t.Run("NilRequest", func(t *testing.T) {
		restBuilder := restconf.NewBuilder("https", "controller.example.com")
		logger := slog.Default()
		rb := NewRequestBuilder(restBuilder, "token", logger)
		client := &http.Client{}

		// Execute with nil request should handle gracefully
		resp, err := rb.ExecuteRequest(client, nil)
		if resp != nil {
			resp.Body.Close()
		}
		testutil.AssertError(t, err, "Expected error for nil request")
	})
}

// Test CreateRequest function.
func TestClientUnit_RequestBuilderCreateRequest_Success(t *testing.T) {
	t.Run("ValidRequest", func(t *testing.T) {
		restBuilder := restconf.NewBuilder("https", "controller.example.com")
		logger := slog.Default()
		rb := NewRequestBuilder(restBuilder, "token", logger)

		req, err := rb.CreateRequest(context.Background(), http.MethodGet, "test/path")
		testutil.AssertNoError(t, err, "CreateRequest")
		testutil.AssertNotNil(t, req, "Request should not be nil")
		testutil.AssertStringEquals(t, req.Method, http.MethodGet, "HTTP method")
		testutil.AssertStringNotEmpty(t, req.Header.Get("Authorization"), "Authorization header")
	})

	t.Run("NilRESTCONFBuilder", func(t *testing.T) {
		rb := NewRequestBuilder(nil, "token", slog.Default())
		_, err := rb.CreateRequest(context.Background(), http.MethodGet, "test/path")
		testutil.AssertError(t, err, "Expected error for nil RESTCONF builder")
		testutil.AssertStringContains(t, err.Error(), "RESTCONF builder is not properly initialized", "error message")
	})

	t.Run("InvalidHTTPMethod", func(t *testing.T) {
		restBuilder := restconf.NewBuilder("https", "controller.example.com")
		logger := slog.Default()
		rb := NewRequestBuilder(restBuilder, "token", logger)

		// Test with invalid method that causes http.NewRequestWithContext to fail
		_, err := rb.CreateRequest(context.Background(), "\n", "test/path")
		testutil.AssertError(t, err, "Expected error for invalid HTTP method")
		testutil.AssertStringContains(t, err.Error(), "failed to create request", "error message")
	})
}

// Test ExecuteRequest with mock server.
func TestClientUnit_ExecuteRequestWithMockServer_Success(t *testing.T) {
	// Create a mock server for testing
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/yang-data+json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"test": "response"}`))
	}))
	defer mockServer.Close()

	restBuilder := restconf.NewBuilder("http", mockServer.URL[7:]) // Remove "http://"
	logger := slog.Default()
	rb := NewRequestBuilder(restBuilder, "test-token", logger)
	client := mockServer.Client()

	req, err := rb.CreateRequest(context.Background(), http.MethodGet, "test-endpoint")
	testutil.AssertNoError(t, err, "CreateRequest")

	resp, err := rb.ExecuteRequest(client, req)
	testutil.AssertNoError(t, err, "ExecuteRequest")
	testutil.AssertNotNil(t, resp, "Response should not be nil")
	defer resp.Body.Close()

	testutil.AssertIntEquals(t, resp.StatusCode, http.StatusOK, "Status code")
}

// Test CreateRequestWithPayload error scenarios.
func TestClientUnit_CreateRequestWithPayload_ErrorScenarios(t *testing.T) {
	t.Run("InvalidPayloadSerialization", func(t *testing.T) {
		restBuilder := restconf.NewBuilder("https", "controller.example.com")
		logger := slog.Default()
		rb := NewRequestBuilder(restBuilder, "token", logger)

		// Create an invalid payload that cannot be marshaled
		invalidPayload := make(chan int) // channels cannot be JSON marshaled

		_, err := rb.CreateRequestWithPayload(
			context.Background(),
			http.MethodPost,
			"test/path",
			invalidPayload,
		)
		testutil.AssertError(t, err, "Expected error for invalid payload")
		testutil.AssertStringContains(t, err.Error(), "failed to marshal payload", "error message")
	})

	t.Run("NilPayload", func(t *testing.T) {
		restBuilder := restconf.NewBuilder("https", "controller.example.com")
		logger := slog.Default()
		rb := NewRequestBuilder(restBuilder, "token", logger)

		req, err := rb.CreateRequestWithPayload(
			context.Background(),
			http.MethodPost,
			"test/path",
			nil, // nil payload
		)
		testutil.AssertNoError(t, err, "Should handle nil payload")
		testutil.AssertStringEquals(t, req.Method, http.MethodPost, "HTTP method")
		testutil.AssertStringEquals(
			t,
			req.Header.Get("Content-Type"),
			"",
			"Content-Type should not be set for nil payload",
		)
	})
}

// Test CreateRPCRequestWithPayload error scenarios.
func TestClientUnit_CreateRPCRequestWithPayload_ErrorScenarios(t *testing.T) {
	t.Run("InvalidPayloadSerialization", func(t *testing.T) {
		restBuilder := restconf.NewBuilder("https", "controller.example.com")
		logger := slog.Default()
		rb := NewRequestBuilder(restBuilder, "token", logger)

		// Create an invalid payload that cannot be marshaled
		invalidPayload := make(chan int) // channels cannot be JSON marshaled

		_, err := rb.CreateRPCRequestWithPayload(
			context.Background(),
			http.MethodPost,
			"test/rpc",
			invalidPayload,
		)
		testutil.AssertError(t, err, "Expected error for invalid payload")
		testutil.AssertStringContains(t, err.Error(), "failed to marshal payload", "error message")
	})

	t.Run("NilPayload", func(t *testing.T) {
		restBuilder := restconf.NewBuilder("https", "controller.example.com")
		logger := slog.Default()
		rb := NewRequestBuilder(restBuilder, "token", logger)

		req, err := rb.CreateRPCRequestWithPayload(
			context.Background(),
			http.MethodPost,
			"test/rpc",
			nil, // nil payload
		)
		testutil.AssertNoError(t, err, "Should handle nil payload")
		testutil.AssertStringEquals(t, req.Method, http.MethodPost, "HTTP method")
		testutil.AssertStringEquals(
			t,
			req.Header.Get("Content-Type"),
			"",
			"Content-Type should not be set for nil payload",
		)
	})
}

// Test ExecuteRequest with mock server error scenarios.
func TestClientUnit_ExecuteRequest_ErrorScenarios(t *testing.T) {
	t.Run("HTTPErrorStatus", func(t *testing.T) {
		// Create a mock server that returns 404
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"error": "not found"}`))
		}))
		defer mockServer.Close()

		restBuilder := restconf.NewBuilder("http", mockServer.URL[7:]) // Remove "http://"
		logger := slog.Default()
		rb := NewRequestBuilder(restBuilder, "test-token", logger)
		client := mockServer.Client()

		req, err := rb.CreateRequest(context.Background(), http.MethodGet, "test-endpoint")
		testutil.AssertNoError(t, err, "CreateRequest")

		// ExecuteRequest should succeed but return 404 status
		resp, err := rb.ExecuteRequest(client, req)
		testutil.AssertNoError(t, err, "ExecuteRequest should not fail on HTTP error status")
		testutil.AssertNotNil(t, resp, "Response should not be nil")
		defer resp.Body.Close()

		testutil.AssertIntEquals(t, resp.StatusCode, http.StatusNotFound, "Status code should be 404")
	})
}
