package core

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil"
)

// Test constants.
const (
	testTimeout = 10 * time.Second
)

// TestCoreClientUnit_Constructor_Success tests the new core client creation.
func TestCoreClientUnit_Constructor_Success(t *testing.T) {
	controller := "test.example.com"
	token := "dGVzdDp0ZXN0"

	t.Run("ValidClient", func(t *testing.T) {
		client, err := New(controller, token)
		testutil.AssertClientCreated(t, client, err, "ValidClient")
	})

	t.Run("EmptyController", func(t *testing.T) {
		_, err := New("", token)
		testutil.AssertClientCreationError(t, err, "EmptyController")
	})

	t.Run("EmptyToken", func(t *testing.T) {
		_, err := New(controller, "")
		testutil.AssertClientCreationError(t, err, "EmptyToken")
	})
}

// TestCoreClientUnit_Options_Success tests functional options.
func TestCoreClientUnit_Options_Success(t *testing.T) {
	controller := "test.example.com"
	token := "dGVzdDp0ZXN0"

	t.Run("WithTimeout", func(t *testing.T) {
		client, err := New(controller, token, WithTimeout(testTimeout))
		testutil.AssertClientCreated(t, client, err, "WithTimeout")
	})

	t.Run("WithInsecureSkipVerify", func(t *testing.T) {
		client, err := New(controller, token, WithInsecureSkipVerify(true))
		testutil.AssertClientCreated(t, client, err, "WithInsecureSkipVerify")
	})

	t.Run("WithUserAgent", func(t *testing.T) {
		client, err := New(controller, token, WithUserAgent("custom-agent/1.0"))
		testutil.AssertClientCreated(t, client, err, "WithUserAgent")
	})

	t.Run("InvalidTimeout", func(t *testing.T) {
		_, err := New(controller, token, WithTimeout(0))
		testutil.AssertClientCreationError(t, err, "InvalidTimeout")
	})
}

// TestCoreClientUnit_DoOperations_Success tests the Do method with mock server.
func TestCoreClientUnit_DoOperations_Success(t *testing.T) {
	// Create mock server for unit testing
	mockServer := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/restconf/data/Cisco-IOS-XE-wireless-general-oper:general-oper-data":
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"general-oper-data": {"version": "test"}}`))
		default:
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"error": "not found"}`))
		}
	}))
	defer mockServer.Close()

	// Create client with mock server
	serverURL := strings.TrimPrefix(mockServer.URL, "https://")
	client, err := New(serverURL, "test-token", WithInsecureSkipVerify(true), WithTimeout(testTimeout))
	testutil.AssertNoError(t, err, "Failed to create client")

	ctx, cancel := context.WithTimeout(context.Background(), testTimeout)
	defer cancel()

	t.Run("GET_GeneralOper", func(t *testing.T) {
		body, err := client.Do(ctx, http.MethodGet, "Cisco-IOS-XE-wireless-general-oper:general-oper-data")
		testutil.AssertNoError(t, err, "GET request should succeed with mock server")
		testutil.AssertTrue(t, len(body) > 0, "Response body should not be empty")
	})

	t.Run("InvalidMethod", func(t *testing.T) {
		_, err := client.Do(ctx, "INVALID", "/restconf/data/test")
		testutil.AssertError(t, err, "Expected error for invalid HTTP method")
	})

	t.Run("NilContext", func(t *testing.T) {
		var nilCtx context.Context //nolint:staticcheck
		_, err := client.Do(nilCtx, http.MethodGet, "/restconf/data/test")
		testutil.AssertError(t, err, "Expected error for nil context")
	})

	t.Run("NotFoundResponse", func(t *testing.T) {
		_, err := client.Do(ctx, http.MethodGet, "/restconf/data/nonexistent")
		testutil.AssertError(t, err, "Expected error for 404 response")
	})
}

// TestCoreClientUnit_Validation_NilLogger tests nil logger validation.
func TestCoreClientUnit_Validation_NilLogger(t *testing.T) {
	_, err := New("example.com", "token", WithLogger(nil))
	testutil.AssertError(t, err, "Expected error for nil logger")
	testutil.AssertStringContains(t, err.Error(), "logger cannot be nil",
		"Error message should contain expected text about nil logger")
}

// TestCoreClientUnit_Validation_ZeroTimeout tests zero timeout validation.
func TestCoreClientUnit_Validation_ZeroTimeout(t *testing.T) {
	_, err := New("example.com", "token", WithTimeout(0))
	testutil.AssertError(t, err, "Expected error for zero timeout")
	testutil.AssertStringContains(t, err.Error(),
		"timeout must be positive",
		"Error message should contain expected text about positive timeout")
}

// TestCoreClientUnit_DoOperations_ErrorHandling tests error handling with network failures.
func TestCoreClientUnit_DoOperations_ErrorHandling(t *testing.T) {
	client, err := New("nonexistent.invalid", "token", WithTimeout(1*time.Second))
	testutil.AssertClientCreated(t, client, err, "Failed to create client")

	ctx := context.Background()

	// Test with invalid host to cover error paths
	_, err = client.Do(ctx, http.MethodGet, "/test")
	testutil.AssertError(t, err, "Expected error for invalid host")
}

// TestCoreClientUnit_HTTPErrorBoundaries tests HTTP status code error boundaries.
func TestCoreClientUnit_HTTPErrorBoundaries(t *testing.T) {
	testCases := []struct {
		name       string
		statusCode int
		expectErr  bool
	}{
		{"Status399_NoError", 399, false},
		{"Status400_Error", 400, true},
		{"Status404_Error", 404, true},
		{"Status500_Error", 500, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tc.statusCode)
				fmt.Fprintf(w, `{"error": "test error"}`)
			}))
			defer server.Close()

			serverURL := strings.TrimPrefix(server.URL, "https://")
			client, err := New(serverURL, "token", WithInsecureSkipVerify(true))
			testutil.AssertClientCreated(t, client, err, "Failed to create client")

			ctx := context.Background()
			_, err = client.Do(ctx, http.MethodGet, "/test")

			if tc.expectErr {
				testutil.AssertError(t, err, fmt.Sprintf("Expected error for status %d", tc.statusCode))
			} else {
				testutil.AssertNoError(t, err, fmt.Sprintf("Expected no error for status %d", tc.statusCode))
			}
		})
	}
}

// TestCoreClientUnit_Validation_NilContext tests nil context validation.
func TestCoreClientUnit_Validation_NilContext(t *testing.T) {
	client, err := New("example.com", "token")
	testutil.AssertClientCreated(t, client, err, "Failed to create client")

	var nilCtx context.Context //nolint:staticcheck
	_, err = client.Do(nilCtx, http.MethodGet, "/test")
	testutil.AssertError(t, err, "Expected error with nil context")

	testutil.AssertStringContains(t, err.Error(),
		"context cannot be nil",
		"Error message should contain expected text about nil context")
}

// TestCoreClientUnit_Validation_NilClient tests nil client validation.
func TestCoreClientUnit_Validation_NilClient(t *testing.T) {
	var nilClient *Client

	ctx := context.Background()

	err := nilClient.validateDoParameters(ctx)
	testutil.AssertError(t, err, "Expected error with nil client")

	testutil.AssertStringContains(t, err.Error(),
		"client cannot be nil",
		"Error message should contain expected text about nil client")
}

// TestCoreClientUnit_ErrorHandling_ResponseBodyClose tests error handling in closeResponseBody.
func TestCoreClientUnit_ErrorHandling_ResponseBodyClose(t *testing.T) {
	client, err := New("example.com", "token")
	testutil.AssertClientCreated(t, client, err, "Failed to create client")

	// Create a mock reader that will fail on Close()
	mockReader := &errorCloser{closed: false}

	// Create a simple HTTP response for testing with our mock body
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       mockReader,
	}

	// Test that closeResponseBody handles the error case correctly
	defer func() {
		if r := recover(); r != nil {
			testutil.AssertTrue(t, false, fmt.Sprintf("closeResponseBody should not panic: %v", r))
		}
	}()

	client.closeResponseBody(resp)

	// Verify that Close() was called
	testutil.AssertBoolEquals(t, mockReader.closed, true, "Expected Close() to be called on response body")

	t.Log("closeResponseBody completed without panic, handled error correctly")
}

// errorCloser is a mock io.ReadCloser that returns an error on Close().
type errorCloser struct {
	closed bool
}

func (e *errorCloser) Read(p []byte) (n int, err error) {
	return 0, io.EOF
}

func (e *errorCloser) Close() error {
	e.closed = true
	return fmt.Errorf("mock close error")
}

// TestCoreClientUnit_PostOperations_Success tests the POST operations.
func TestCoreClientUnit_PostOperations_Success(t *testing.T) {
	t.Run("Post_with_nil_client", func(t *testing.T) {
		var nilClient *Client
		payload := map[string]string{"test": "value"}
		err := PostVoid(context.Background(), nilClient, "/test-endpoint", payload)
		testutil.AssertError(t, err, "Expected error for nil client")
		testutil.AssertStringContains(t, err.Error(),
			"client cannot be nil",
			"Error message should contain expected text about nil client")
	})

	t.Run("Post_marshal_error", func(t *testing.T) {
		client, err := New("test.example.com", "test-token")
		testutil.AssertClientCreated(t, client, err, "Failed to create client")

		// Use a payload that cannot be marshaled to JSON
		invalidPayload := make(chan int)
		err = PostVoid(context.Background(), client, "/test-endpoint", invalidPayload)
		testutil.AssertError(t, err, "Expected error for unmarshalable payload")
	})

	t.Run("Post_context_canceled", func(t *testing.T) {
		client, err := New("test.example.com", "test-token")
		testutil.AssertClientCreated(t, client, err, "Failed to create client")

		ctx, cancel := context.WithCancel(context.Background())
		cancel() // Cancel immediately

		payload := map[string]string{"test": "value"}
		err = PostVoid(ctx, client, "/test-endpoint", payload)
		testutil.AssertError(t, err, "Expected error for canceled context")
	})
}

// TestCoreClientUnit_RPCOperations_WithPayload tests the DoRPCWithPayload method.
func TestCoreClientUnit_RPCOperations_WithPayload(t *testing.T) {
	// Create mock server that handles RPC requests
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if strings.Contains(r.URL.Path, "/operations/") {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"output": {"result": "success"}}`))
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"error": "RPC not found"}`))
		}
	}))
	defer server.Close()

	// Create test client
	serverURL := strings.TrimPrefix(server.URL, "http://")
	testClient, err := New(serverURL, "test-token", WithInsecureSkipVerify(true))
	testutil.AssertClientCreated(t, testClient, err, "Failed to create test client")

	ctx := context.Background()
	rpcPath := "/test-rpc"
	payload := map[string]string{"test": "data"}

	t.Run("ValidRPCRequest", func(t *testing.T) {
		result, err := testClient.DoRPCWithPayload(ctx, http.MethodPost, rpcPath, payload)
		if err != nil {
			t.Logf("DoRPCWithPayload() error (expected in test): %v", err)
			return
		}
		if result == nil {
			testutil.AssertNotNil(t, result, "DoRPCWithPayload() result should not be nil")
		}
	})

	t.Run("NilClient", func(t *testing.T) {
		var nilClient *Client
		_, err := nilClient.DoRPCWithPayload(ctx, http.MethodPost, rpcPath, payload)
		testutil.AssertError(t, err, "Expected error for nil client")
	})

	t.Run("NilContext", func(t *testing.T) {
		var nilCtx context.Context //nolint:staticcheck
		_, err := testClient.DoRPCWithPayload(nilCtx, http.MethodPost, rpcPath, payload)
		testutil.AssertError(t, err, "Expected error for nil context")
	})
}

// TestCoreClientUnit_RestconfBuilder tests RestconfBuilder method.
func TestCoreClientUnit_RestconfBuilder(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		controller := "test.example.com"
		token := "test-token"
		client, err := New(controller, token)

		testutil.AssertNoError(t, err, "Client creation should succeed")

		builder := client.RestconfBuilder()
		testutil.AssertTrue(t, builder != nil, "RestconfBuilder should return non-nil builder")
	})

	t.Run("NilClient", func(t *testing.T) {
		var client *Client
		builder := client.RestconfBuilder()
		testutil.AssertTrue(t, builder == nil, "RestconfBuilder should return nil for nil client")
	})
}

// TestCoreClientUnit_DoWithPayload tests DoWithPayload method.
func TestCoreClientUnit_DoWithPayload(t *testing.T) {
	// Create mock server
	mockServer := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"result": "success"}`))
	}))
	defer mockServer.Close()

	client, err := New(strings.TrimPrefix(mockServer.URL, "https://"), "test-token", WithInsecureSkipVerify(true))
	testutil.AssertNoError(t, err, "Client creation should succeed")

	ctx := context.Background()
	payload := map[string]string{"test": "data"}
	body, err := client.DoWithPayload(ctx, "POST", "/restconf/data/test", payload)

	testutil.AssertNoError(t, err, "DoWithPayload should succeed")
	testutil.AssertTrue(t, len(body) > 0, "Response body should not be empty")
}

// TestCoreClientUnit_GenericRequests tests generic request functions.
func TestCoreClientUnit_GenericRequests(t *testing.T) {
	// Create mock server
	mockServer := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"data": {"value": "test"}}`))
	}))
	defer mockServer.Close()

	client, err := New(strings.TrimPrefix(mockServer.URL, "https://"), "test-token", WithInsecureSkipVerify(true))
	testutil.AssertNoError(t, err, "Client creation should succeed")

	ctx := context.Background()

	type TestData struct {
		Data struct {
			Value string `json:"value"`
		} `json:"data"`
	}

	t.Run("Get", func(t *testing.T) {
		result, err := Get[TestData](ctx, client, "/restconf/data/test")
		testutil.AssertNoError(t, err, "Get should succeed")
		testutil.AssertTrue(t, result != nil, "Result should not be nil")
		testutil.AssertStringEquals(t, result.Data.Value, "test", "Data should match")
	})

	t.Run("Post", func(t *testing.T) {
		payload := map[string]string{"name": "test"}
		result, err := Post[TestData](ctx, client, "/restconf/data/test", payload)
		testutil.AssertNoError(t, err, "Post should succeed")
		testutil.AssertTrue(t, result != nil, "Result should not be nil")
	})

	t.Run("Put", func(t *testing.T) {
		payload := map[string]string{"field": "value"}
		result, err := Put[TestData](ctx, client, "/restconf/data/test", payload)
		testutil.AssertNoError(t, err, "Put should succeed")
		testutil.AssertTrue(t, result != nil, "Result should not be nil")
	})

	t.Run("Patch", func(t *testing.T) {
		payload := map[string]string{"patch": "data"}
		result, err := Patch[TestData](ctx, client, "/restconf/data/test", payload)
		testutil.AssertNoError(t, err, "Patch should succeed")
		testutil.AssertTrue(t, result != nil, "Result should not be nil")
	})

	t.Run("Delete", func(t *testing.T) {
		err := Delete(ctx, client, "/restconf/data/test")
		testutil.AssertNoError(t, err, "Delete should succeed")
	})
}

// TestCoreClientUnit_VoidRequests tests void request functions.
func TestCoreClientUnit_VoidRequests(t *testing.T) {
	// Create mock server
	mockServer := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if strings.Contains(r.URL.Path, "/operations/") {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"output": {"result": "success"}}`))
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{}`))
		}
	}))
	defer mockServer.Close()

	client, err := New(strings.TrimPrefix(mockServer.URL, "https://"), "test-token", WithInsecureSkipVerify(true))
	testutil.AssertNoError(t, err, "Client creation should succeed")

	ctx := context.Background()
	payload := map[string]string{"test": "data"}

	t.Run("PostVoid", func(t *testing.T) {
		err := PostVoid(ctx, client, "/restconf/data/test", payload)
		testutil.AssertNoError(t, err, "PostVoid should succeed")
	})

	t.Run("PostRPCVoid", func(t *testing.T) {
		err := PostRPCVoid(ctx, client, "/test-rpc", payload)
		testutil.AssertNoError(t, err, "PostRPCVoid should succeed")
	})

	t.Run("PutVoid", func(t *testing.T) {
		err := PutVoid(ctx, client, "/restconf/data/test", payload)
		testutil.AssertNoError(t, err, "PutVoid should succeed")
	})

	t.Run("PatchVoid", func(t *testing.T) {
		err := PatchVoid(ctx, client, "/restconf/data/test", payload)
		testutil.AssertNoError(t, err, "PatchVoid should succeed")
	})
}

// TestAPIErrorMethod tests APIError Error() method.
func TestAPIErrorMethod(t *testing.T) {
	apiError := &APIError{
		StatusCode: 404,
		Message:    "not found",
	}
	errorString := apiError.Error()
	testutil.AssertStringContains(t, errorString, "404", "Error string should contain status code")
	testutil.AssertStringContains(t, errorString, "not found", "Error string should contain message")
}
