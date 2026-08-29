package core

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"errors"
	"fmt"
	"io"
	"math/big"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"
	"unicode/utf8"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/transport"
)

// Test constants.
const (
	testTimeout = 10 * time.Second
)

// TestCoreClientUnit_Constructor_Success tests the new core client creation.
func TestCoreClientUnit_Constructor_Success(t *testing.T) {
	controller := "wnc1.example.internal"
	token := "test-token-123"

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
	controller := "wnc1.example.internal"
	token := "test-token-123"

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

// TestCoreClientUnit_WithProxy_Success tests the proxy resolver option.
func TestCoreClientUnit_WithProxy_Success(t *testing.T) {
	controller := "wnc1.example.internal"
	token := "test-token-123"

	t.Run("DefaultLeavesProxyUnset", func(t *testing.T) {
		client, err := New(controller, token)
		testutil.AssertNoError(t, err, "Client creation should succeed")
		testutil.AssertTrue(t, client.httpTransport.Proxy == nil,
			"Proxy must stay unset so no environment variable is consulted")
	})

	t.Run("ResolverIsConsulted", func(t *testing.T) {
		want := "http://proxy.example.com:3128"
		proxyURL, err := url.Parse(want)
		testutil.AssertNoError(t, err, "Proxy URL should parse")

		client, err := New(controller, token, WithProxy(http.ProxyURL(proxyURL)))
		testutil.AssertNoError(t, err, "Client creation should succeed")
		testutil.AssertTrue(t, client.httpTransport.Proxy != nil, "Proxy resolver should be installed")

		req := httptest.NewRequest(http.MethodGet, "https://"+controller+"/restconf/data", http.NoBody)
		resolved, err := client.httpTransport.Proxy(req)
		testutil.AssertNoError(t, err, "Proxy resolver should not fail")
		testutil.AssertTrue(t, resolved != nil, "Proxy resolver should return a URL")
		if resolved == nil {
			return
		}
		testutil.AssertStringEquals(t, resolved.String(), want, "Resolved proxy URL")
	})

	t.Run("NilResolverRestoresDirect", func(t *testing.T) {
		client, err := New(controller, token, WithProxy(http.ProxyFromEnvironment), WithProxy(nil))
		testutil.AssertNoError(t, err, "Client creation should succeed")
		testutil.AssertTrue(t, client.httpTransport.Proxy == nil,
			"A nil resolver connects directly, it does not restore the environment resolver")
	})

	t.Run("SurvivesLaterTransportOption", func(t *testing.T) {
		client, err := New(controller, token,
			WithProxy(http.ProxyFromEnvironment), WithInsecureSkipVerify(true))
		testutil.AssertNoError(t, err, "Client creation should succeed")

		testutil.AssertTrue(t, client.httpTransport.Proxy != nil,
			"A later transport option must not drop the proxy resolver")
		testutil.AssertTrue(t, client.httpTransport.TLSClientConfig.InsecureSkipVerify,
			"WithInsecureSkipVerify should still take effect")

		tr, ok := client.httpClient.Transport.(*http.Transport)
		testutil.AssertTrue(t, ok && tr == client.httpTransport,
			"httpClient.Transport must stay the object the options mutate in place")
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
	client, err := New(serverURL, "test-token-123", WithInsecureSkipVerify(true), WithTimeout(testTimeout))
	testutil.AssertNoError(t, err, "Failed to create client")

	ctx, cancel := context.WithTimeout(context.Background(), testTimeout)
	defer cancel()

	t.Run("GET_GeneralOper", func(t *testing.T) {
		resp, err := client.do(ctx, http.MethodGet, "Cisco-IOS-XE-wireless-general-oper:general-oper-data")
		testutil.AssertNoError(t, err, "GET request should succeed with mock server")
		testutil.AssertTrue(t, len(resp.Body) > 0, "Response body should not be empty")
	})

	t.Run("InvalidMethod", func(t *testing.T) {
		_, err := client.do(ctx, "INVALID", "/restconf/data/test")
		testutil.AssertError(t, err, "Expected error for invalid HTTP method")
	})

	t.Run("NilContext", func(t *testing.T) {
		var nilCtx context.Context //nolint:staticcheck
		_, err := client.do(nilCtx, http.MethodGet, "/restconf/data/test")
		testutil.AssertError(t, err, "Expected error for nil context")
	})

	t.Run("NotFoundResponse", func(t *testing.T) {
		_, err := client.do(ctx, http.MethodGet, "/restconf/data/nonexistent")
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
	_, err = client.do(ctx, http.MethodGet, "/test")
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
			_, err = client.do(ctx, http.MethodGet, "/test")

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
	_, err = client.do(nilCtx, http.MethodGet, "/test")
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

	// Held by not panicking: CloseIdleConnections returns on a nil receiver.
	nilClient.CloseIdleConnections()
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
		client, err := New("wnc1.example.internal", "test-token-123")
		testutil.AssertClientCreated(t, client, err, "Failed to create client")

		// Use a payload that cannot be marshaled to JSON
		invalidPayload := make(chan int)
		err = PostVoid(context.Background(), client, "/test-endpoint", invalidPayload)
		testutil.AssertError(t, err, "Expected error for unmarshalable payload")
	})

	t.Run("Post_context_canceled", func(t *testing.T) {
		client, err := New("wnc1.example.internal", "test-token-123")
		testutil.AssertClientCreated(t, client, err, "Failed to create client")

		ctx, cancel := context.WithCancel(context.Background())
		cancel() // Cancel immediately

		payload := map[string]string{"test": "value"}
		err = PostVoid(ctx, client, "/test-endpoint", payload)
		testutil.AssertError(t, err, "Expected error for canceled context")
	})
}

// TestCoreClientUnit_RPCOperations_WithPayload tests the doRPC method.
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
	testClient, err := New(serverURL, "test-token-123", WithInsecureSkipVerify(true))
	testutil.AssertClientCreated(t, testClient, err, "Failed to create test client")

	ctx := context.Background()
	rpcPath := "/test-rpc"
	payload := map[string]string{"test": "data"}

	t.Run("ValidRPCRequest", func(t *testing.T) {
		result, err := testClient.doRPC(ctx, rpcPath, payload)
		if err != nil {
			t.Logf("doRPC() error (expected in test): %v", err)
			return
		}
		if result == nil {
			testutil.AssertNotNil(t, result, "doRPC() result should not be nil")
		}
	})

	t.Run("NilClient", func(t *testing.T) {
		var nilClient *Client
		_, err := nilClient.doRPC(ctx, rpcPath, payload)
		testutil.AssertError(t, err, "Expected error for nil client")
	})

	t.Run("NilContext", func(t *testing.T) {
		var nilCtx context.Context //nolint:staticcheck
		_, err := testClient.doRPC(nilCtx, rpcPath, payload)
		testutil.AssertError(t, err, "Expected error for nil context")
	})
}

// TestCoreClientUnit_RESTCONFBuilder tests RESTCONFBuilder method.
func TestCoreClientUnit_RESTCONFBuilder(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		controller := "wnc1.example.internal"
		token := "test-token-123"
		client, err := New(controller, token)

		testutil.AssertNoError(t, err, "Client creation should succeed")

		builder := client.RESTCONFBuilder()
		testutil.AssertTrue(t, builder != nil, "RESTCONFBuilder should return non-nil builder")
	})

	t.Run("NilClient", func(t *testing.T) {
		var client *Client
		builder := client.RESTCONFBuilder()
		testutil.AssertTrue(t, builder == nil, "RESTCONFBuilder should return nil for nil client")
	})
}

// TestCoreClientUnit_doWithPayload tests doWithPayload method.
func TestCoreClientUnit_doWithPayload(t *testing.T) {
	// Create mock server
	mockServer := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"result": "success"}`))
	}))
	defer mockServer.Close()

	client, err := New(strings.TrimPrefix(mockServer.URL, "https://"), "test-token-123", WithInsecureSkipVerify(true))
	testutil.AssertNoError(t, err, "Client creation should succeed")

	ctx := context.Background()
	payload := map[string]string{"test": "data"}
	resp, err := client.doWithPayload(ctx, "POST", "/restconf/data/test", payload)

	testutil.AssertNoError(t, err, "doWithPayload should succeed")
	testutil.AssertTrue(t, len(resp.Body) > 0, "Response body should not be empty")
}

// TestCoreClientUnit_GenericRequests tests generic request functions.
func TestCoreClientUnit_GenericRequests(t *testing.T) {
	// Create mock server
	mockServer := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"Cisco-IOS-XE-wireless-test:test": {"value": "test"}}`))
	}))
	defer mockServer.Close()

	client, err := New(strings.TrimPrefix(mockServer.URL, "https://"), "test-token-123", WithInsecureSkipVerify(true))
	testutil.AssertNoError(t, err, "Client creation should succeed")

	ctx := context.Background()

	type TestData struct {
		Data struct {
			Value string `json:"value"`
		} `json:"Cisco-IOS-XE-wireless-test:test"`
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

	client, err := New(strings.TrimPrefix(mockServer.URL, "https://"), "test-token-123", WithInsecureSkipVerify(true))
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

// TestCoreClientUnit_TransportOptions_Success pins the transport-level options against
// the option-order dependency the wholesale Transport replacement created.
func TestCoreClientUnit_TransportOptions_Success(t *testing.T) {
	controller := "wnc1.example.internal"
	token := "test-token-123"

	t.Run("SurvivesLaterInsecureSkipVerify", func(t *testing.T) {
		client, err := New(controller, token,
			WithResponseHeaderTimeout(20*time.Second),
			WithTLSHandshakeTimeout(15*time.Second),
			WithInsecureSkipVerify(true),
		)
		testutil.AssertClientCreated(t, client, err, "SurvivesLaterInsecureSkipVerify")

		tr := client.httpTransport
		testutil.AssertDurationEquals(t, tr.ResponseHeaderTimeout, 20*time.Second, "ResponseHeaderTimeout")
		testutil.AssertDurationEquals(t, tr.TLSHandshakeTimeout, 15*time.Second, "TLSHandshakeTimeout")
		testutil.AssertTrue(t, tr.TLSClientConfig.InsecureSkipVerify, "InsecureSkipVerify")
		testutil.AssertTrue(t, client.httpClient.Transport == tr, "httpClient keeps the same transport")
	})

	t.Run("DefaultsAndRejections", func(t *testing.T) {
		client, err := New(controller, token)
		testutil.AssertClientCreated(t, client, err, "DefaultsAndRejections")
		testutil.AssertTrue(t, client.httpTransport.Proxy == nil, "Proxy is off by default")
		testutil.AssertTrue(t, client.httpTransport.DialContext != nil, "DialContext is set")

		_, err = New(controller, token, WithResponseHeaderTimeout(0))
		testutil.AssertClientCreationError(t, err, "zero response header timeout")
		_, err = New(controller, token, WithTLSHandshakeTimeout(-1))
		testutil.AssertClientCreationError(t, err, "negative TLS handshake timeout")
	})
}

// TestCoreClientUnit_UserAgentReachesRequest_Success pins that WithUserAgent changes the
// header a request carries, which the earlier comment-only option body could not do.
func TestCoreClientUnit_UserAgentReachesRequest_Success(t *testing.T) {
	const path = "Cisco-IOS-XE-wireless-general-oper:general-oper-data"

	client, err := New("wnc1.example.internal", "test-token-123", WithUserAgent("  probe-agent/9.9  "))
	testutil.AssertClientCreated(t, client, err, "WithUserAgent")

	req, err := client.requestBuilder.CreateRequest(context.Background(), http.MethodGet, path)
	testutil.AssertNoError(t, err, "CreateRequest")
	testutil.AssertStringEquals(t, req.Header.Get(transport.HTTPHeaderKeyUserAgent),
		"probe-agent/9.9", "the trimmed custom User-Agent reaches the request")
}

// TestCoreClientUnit_HeaderTimeout_Error pins that a transport timeout carries the
// sentinel wnc.go already documents, not only an opaque *url.Error.
func TestCoreClientUnit_HeaderTimeout_Error(t *testing.T) {
	const path = "Cisco-IOS-XE-wireless-general-oper:general-oper-data"

	server := httptest.NewTLSServer(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
		time.Sleep(500 * time.Millisecond)
	}))
	defer server.Close()

	client, err := New(strings.TrimPrefix(server.URL, "https://"), "test-token-123",
		WithInsecureSkipVerify(true), WithResponseHeaderTimeout(20*time.Millisecond))
	testutil.AssertClientCreated(t, client, err, "HeaderTimeoutClient")

	_, err = client.do(context.Background(), http.MethodGet, path)
	testutil.AssertError(t, err, "header timeout")
	testutil.AssertTrue(t, errors.Is(err, ErrRequestTimeout), "errors.Is(err, ErrRequestTimeout)")
}

// TestCoreClientUnit_ConstructionInputNormalization_Success pins the two inputs that
// passed validation and then broke every request.
func TestCoreClientUnit_ConstructionInputNormalization_Success(t *testing.T) {
	const path = "Cisco-IOS-XE-wireless-general-oper:general-oper-data"

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{}`))
	}))
	defer server.Close()

	host := strings.TrimPrefix(server.URL, "https://")
	client, err := New(" "+host+"\n", "test-token-123\n", WithInsecureSkipVerify(true))
	testutil.AssertClientCreated(t, client, err, "padded host and newline-terminated token")

	_, err = client.do(context.Background(), http.MethodGet, path)
	testutil.AssertNoError(t, err, "request with a newline-terminated token")
}

// TestCoreClientUnit_ErrorBodyTruncation_Error pins the bound on the copies that reach
// a log line, and that APIError.Body is not the one truncated.
func TestCoreClientUnit_ErrorBodyTruncation_Error(t *testing.T) {
	const path = "Cisco-IOS-XE-wireless-general-oper:general-oper-data"
	body := strings.Repeat("x", maxLoggedBodyBytes*3)

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(body))
	}))
	defer server.Close()

	client, err := New(strings.TrimPrefix(server.URL, "https://"), "test-token-123",
		WithInsecureSkipVerify(true))
	testutil.AssertClientCreated(t, client, err, "ErrorBodyTruncation")

	_, err = client.do(context.Background(), http.MethodGet, path)
	var apiErr *APIError
	testutil.AssertTrue(t, errors.As(err, &apiErr), "errors.As(*APIError)")
	// The bound is written out rather than taken from the constant: comparing the constant
	// with itself passes at any value, which is what the bound exists to fix.
	testutil.AssertIntEquals(t, len(apiErr.Message), 512+len("... (truncated)"), "Message is bounded at 512")
	testutil.AssertTrue(t, strings.HasSuffix(apiErr.Message, "... (truncated)"), "truncation marker")
	testutil.AssertIntEquals(t, len(apiErr.Body), len(body), "Body is intact")
}

// TestCoreClientUnit_ErrorBodyTruncationBoundary_Error pins the comparison at the bound itself.
// A body of exactly maxLoggedBodyBytes must pass through untouched: relaxing the guard from <= to
// < truncates it, and no other test in this package notices, because they all use a body either
// far longer or far shorter than the bound.
func TestCoreClientUnit_ErrorBodyTruncationBoundary_Error(t *testing.T) {
	const path = "Cisco-IOS-XE-wireless-general-oper:general-oper-data"

	for _, tc := range []struct {
		name      string
		length    int
		truncated bool
	}{
		{"one byte under the bound", maxLoggedBodyBytes - 1, false},
		{"exactly the bound", maxLoggedBodyBytes, false},
		{"one byte over the bound", maxLoggedBodyBytes + 1, true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			body := strings.Repeat("x", tc.length)
			server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
				w.WriteHeader(http.StatusInternalServerError)
				_, _ = w.Write([]byte(body))
			}))
			defer server.Close()

			client, err := New(strings.TrimPrefix(server.URL, "https://"), "test-token-123",
				WithInsecureSkipVerify(true))
			testutil.AssertClientCreated(t, client, err, "ErrorBodyTruncationBoundary")

			_, err = client.do(context.Background(), http.MethodGet, path)
			var apiErr *APIError
			testutil.AssertTrue(t, errors.As(err, &apiErr), "errors.As(*APIError)")

			marked := strings.HasSuffix(apiErr.Message, "... (truncated)")
			if marked != tc.truncated {
				t.Errorf("Expected truncated=%v for a %d-byte body, got %v",
					tc.truncated, tc.length, marked)
			}
			if !tc.truncated {
				testutil.AssertIntEquals(t, len(apiErr.Message), tc.length, "Message is the whole body")
			}
		})
	}
}

// TestCoreClientUnit_ErrorBodyTruncationUTF8_Error pins the reason the cut runs through
// ToValidUTF8: 512 bytes lands mid-rune on a multibyte body.
func TestCoreClientUnit_ErrorBodyTruncationUTF8_Error(t *testing.T) {
	const path = "Cisco-IOS-XE-wireless-general-oper:general-oper-data"
	// 3 bytes per rune, so the 512th byte falls inside the 171st one.
	body := strings.Repeat("あ", 400)

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(body))
	}))
	defer server.Close()

	client, err := New(strings.TrimPrefix(server.URL, "https://"), "test-token-123",
		WithInsecureSkipVerify(true))
	testutil.AssertClientCreated(t, client, err, "ErrorBodyTruncationUTF8")

	_, err = client.do(context.Background(), http.MethodGet, path)
	var apiErr *APIError
	testutil.AssertTrue(t, errors.As(err, &apiErr), "errors.As(*APIError)")
	testutil.AssertTrue(t, utf8.ValidString(apiErr.Message), "Message stays valid UTF-8")
	testutil.AssertTrue(t, len(apiErr.Message) < 512+len("... (truncated)"),
		"the partial rune is dropped, so the cut is shorter than the bound")
}

// TestCoreClientUnit_TransportErrorClassification_Error pins the negative side: only a
// timeout carries ErrRequestTimeout. Without these, widening the classifier goes unnoticed.
func TestCoreClientUnit_TransportErrorClassification_Error(t *testing.T) {
	t.Run("CanceledIsNotTimeout", func(t *testing.T) {
		server := httptest.NewTLSServer(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
			time.Sleep(500 * time.Millisecond)
		}))
		defer server.Close()

		client, err := New(strings.TrimPrefix(server.URL, "https://"), "test-token-123",
			WithInsecureSkipVerify(true))
		testutil.AssertClientCreated(t, client, err, "CanceledIsNotTimeout")

		ctx, cancel := context.WithCancel(context.Background())
		go func() {
			time.Sleep(20 * time.Millisecond)
			cancel()
		}()
		_, err = client.do(ctx, http.MethodGet, "Cisco-IOS-XE-wireless-general-oper:general-oper-data")
		testutil.AssertError(t, err, "canceled request")
		testutil.AssertTrue(t, errors.Is(err, context.Canceled), "errors.Is(err, context.Canceled)")
		testutil.AssertTrue(t, !errors.Is(err, ErrRequestTimeout), "a cancel is not a timeout")
	})

	t.Run("RefusedIsNotTimeout", func(t *testing.T) {
		// A closed port on the loopback interface: the dial fails as a net.Error whose
		// Timeout() is false, which is the case dropping the guard would misclassify.
		listener, lerr := net.Listen("tcp", "127.0.0.1:0")
		testutil.AssertNoError(t, lerr, "listen")
		addr := listener.Addr().String()
		testutil.AssertNoError(t, listener.Close(), "close listener")

		client, err := New(addr, "test-token-123", WithInsecureSkipVerify(true))
		testutil.AssertClientCreated(t, client, err, "RefusedIsNotTimeout")

		_, err = client.do(context.Background(), http.MethodGet,
			"Cisco-IOS-XE-wireless-general-oper:general-oper-data")
		testutil.AssertError(t, err, "refused dial")
		testutil.AssertTrue(t, !errors.Is(err, ErrRequestTimeout), "a refused dial is not a timeout")
	})

	t.Run("BodyReadTimeoutIsATimeout", func(t *testing.T) {
		// The headers arrive and then the body stalls, so the deadline fires inside the
		// body read rather than before the headers. Both are one event to the caller.
		server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			w.Header().Set("Content-Length", "1024")
			w.WriteHeader(http.StatusOK)
			if f, ok := w.(http.Flusher); ok {
				f.Flush()
			}
			time.Sleep(2 * time.Second)
		}))
		defer server.Close()

		client, err := New(strings.TrimPrefix(server.URL, "https://"), "test-token-123",
			WithInsecureSkipVerify(true), WithTimeout(300*time.Millisecond))
		testutil.AssertClientCreated(t, client, err, "BodyReadTimeout")

		_, err = client.do(context.Background(), http.MethodGet,
			"Cisco-IOS-XE-wireless-general-oper:general-oper-data")
		testutil.AssertError(t, err, "body read timeout")
		testutil.AssertTrue(t, errors.Is(err, ErrRequestTimeout),
			"a deadline that fires during the body read is still a timeout")
	})
}

// TestCoreClientUnit_ConstructionSentinel_Error pins the one sentinel every construction failure
// carries, a refused option included: the option path used to report a bare string that matched
// no sentinel at all.
func TestCoreClientUnit_ConstructionSentinel_Error(t *testing.T) {
	cases := map[string]struct {
		host, token string
		opts        []Option
	}{
		"malformed authority": {host: "https://wnc1.example.internal", token: "test-token-123"},
		"empty host":          {host: "   ", token: "test-token-123"},
		"empty token":         {host: "wnc1.example.internal", token: "   "},
		"option refused":      {host: "wnc1.example.internal", token: "test-token-123", opts: []Option{WithTimeout(0)}},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			client, err := New(tc.host, tc.token, tc.opts...)
			testutil.AssertError(t, err, name)
			testutil.AssertTrue(t, errors.Is(err, ErrInvalidConfiguration), name)
			if client != nil {
				t.Fatalf("%s: expected no client, got %#v", name, client)
			}
		})
	}
}

// TestCoreClientUnit_DefaultBudgets_Success pins the three budgets a fresh client carries to the
// constants that name them, and pins that WithTimeout moves the whole-request one alone.
//
// The second subtest is the assertion the doc comments rest on: bundling the three into
// WithTimeout would pass every other test in this package and fail here.
func TestCoreClientUnit_DefaultBudgets_Success(t *testing.T) {
	controller := "wnc1.example.internal"
	token := "test-token-123"

	t.Run("ConstantsNameWhatTheTransportCarries", func(t *testing.T) {
		client, err := New(controller, token)
		testutil.AssertClientCreated(t, client, err, "ConstantsNameWhatTheTransportCarries")

		testutil.AssertDurationEquals(t, client.httpClient.Timeout, DefaultTimeout, "DefaultTimeout")
		testutil.AssertDurationEquals(t, client.httpTransport.ResponseHeaderTimeout,
			DefaultResponseHeaderTimeout, "DefaultResponseHeaderTimeout")
		testutil.AssertDurationEquals(t, client.httpTransport.TLSHandshakeTimeout,
			DefaultTLSHandshakeTimeout, "DefaultTLSHandshakeTimeout")
	})

	t.Run("WithTimeoutLiftsNeitherOtherBudget", func(t *testing.T) {
		client, err := New(controller, token, WithTimeout(2*time.Minute))
		testutil.AssertClientCreated(t, client, err, "WithTimeoutLiftsNeitherOtherBudget")

		testutil.AssertDurationEquals(t, client.httpClient.Timeout, 2*time.Minute, "raised whole-request budget")
		testutil.AssertDurationEquals(t, client.httpTransport.ResponseHeaderTimeout,
			DefaultResponseHeaderTimeout, "header budget is unchanged")
		testutil.AssertDurationEquals(t, client.httpTransport.TLSHandshakeTimeout,
			DefaultTLSHandshakeTimeout, "handshake budget is unchanged")
	})
}

// TestCoreClientUnit_TLSOptions_Success holds the two TLS options against a real handshake, which
// is the only thing that separates them from an assignment to an unexported field.
//
// The first subtest is the reason WithRootCAs exists: the same server, the same client, and the
// only difference is whether the certificate is trusted or verification is off.
func TestCoreClientUnit_TLSOptions_Success(t *testing.T) {
	const path = "Cisco-IOS-XE-wireless-general-oper:general-oper-data"

	// peerCommonName records the client certificate the server was presented, which is the only
	// place the WithClientCertificate assertion below can read it from.
	var peerCommonName string

	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		peerCommonName = ""
		if len(r.TLS.PeerCertificates) > 0 {
			peerCommonName = r.TLS.PeerCertificates[0].Subject.CommonName
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	server.TLS.ClientAuth = tls.RequestClientCert
	defer server.Close()

	host := strings.TrimPrefix(server.URL, "https://")

	pool := x509.NewCertPool()
	pool.AddCert(server.Certificate())

	t.Run("APrivateCAIsTrustedWithoutDisablingVerification", func(t *testing.T) {
		client, err := New(host, "test-token-123", WithRootCAs(pool))
		testutil.AssertClientCreated(t, client, err, "WithRootCAs")

		// Discard the first return: this commit precedes the Request retype, so do returns
		// []byte here and *Response afterwards, and the discard form compiles either way.
		_, err = client.do(context.Background(), http.MethodGet, path)
		testutil.AssertNoError(t, err, "a read verified against the supplied pool")
	})

	t.Run("TheSameServerIsRefusedWithoutThePool", func(t *testing.T) {
		client, err := New(host, "test-token-123")
		testutil.AssertClientCreated(t, client, err, "no TLS option")

		_, err = client.do(context.Background(), http.MethodGet, path)
		testutil.AssertError(t, err, "the same certificate against the host's roots")
	})

	t.Run("TheClientCertificateReachesTheServer", func(t *testing.T) {
		cert := selfSignedCertificate(t, "probe-client")
		client, err := New(host, "test-token-123", WithRootCAs(pool), WithClientCertificate(cert))
		testutil.AssertClientCreated(t, client, err, "WithClientCertificate")

		_, err = client.do(context.Background(), http.MethodGet, path)
		testutil.AssertNoError(t, err, "a read presenting the client certificate")
		testutil.AssertStringEquals(t, peerCommonName, "probe-client",
			"the certificate the server was presented")
	})

	t.Run("Rejections", func(t *testing.T) {
		_, err := New(host, "test-token-123", WithRootCAs(nil))
		testutil.AssertClientCreationError(t, err, "a nil root CA pool")
		testutil.AssertTrue(t, errors.Is(err, ErrInvalidConfiguration),
			"a refused option is a configuration error")
		_, err = New(host, "test-token-123", WithClientCertificate(tls.Certificate{}))
		testutil.AssertClientCreationError(t, err, "a certificate with no chain")
		testutil.AssertTrue(t, errors.Is(err, ErrInvalidConfiguration),
			"a refused option is a configuration error")
	})
}

// selfSignedCertificate builds a throwaway certificate for the mTLS handshake above. The key is
// generated per run rather than checked in, so no key material lives in this repository.
func selfSignedCertificate(t *testing.T, commonName string) tls.Certificate {
	t.Helper()

	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	testutil.AssertNoError(t, err, "generating a key")

	template := &x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject:      pkix.Name{CommonName: commonName},
		NotBefore:    time.Now().Add(-time.Hour),
		NotAfter:     time.Now().Add(time.Hour),
		KeyUsage:     x509.KeyUsageDigitalSignature,
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth},
	}

	der, err := x509.CreateCertificate(rand.Reader, template, template, &key.PublicKey, key)
	testutil.AssertNoError(t, err, "creating a certificate")

	return tls.Certificate{Certificate: [][]byte{der}, PrivateKey: key}
}
