package testutil

import (
	"net/http"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil"
)

// TestTestUtilUnit_NewTestClient_Success tests NewTestClient function.
func TestTestUtilUnit_NewTestClient_Success(t *testing.T) {
	responses := map[string]string{
		"test/data": `{"result": "success"}`,
	}
	server := NewMockServer(WithSuccessResponses(responses))
	defer server.Close()

	client := NewTestClient(server)
	testutil.AssertNotNil(t, client, "NewTestClient should return a non-nil client")
	testutil.AssertNotNil(t, client.Core(), "Client core should not be nil")
}

// TestTestUtilUnit_NewMockServer_Success tests NewMockServer function.
func TestTestUtilUnit_NewMockServer_Success(t *testing.T) {
	responses := map[string]string{
		"test/endpoint": `{"status": "ok"}`,
	}
	server := NewMockServer(WithSuccessResponses(responses))
	defer server.Close()

	testutil.AssertNotNil(t, server, "NewMockServer should return a non-nil server")
	testutil.AssertStringNotEmpty(t, server.URL(), "Server URL should not be empty")
}

// TestTestUtilUnit_NewMockErrorServer_Success tests NewMockErrorServer function.
func TestTestUtilUnit_NewMockErrorServer_Success(t *testing.T) {
	errorPaths := []string{"error/path"}
	server := NewMockServer(WithErrorResponses(errorPaths, http.StatusInternalServerError))
	defer server.Close()

	testutil.AssertNotNil(t, server, "NewMockErrorServer should return a non-nil server")
	testutil.AssertStringNotEmpty(t, server.URL(), "Server URL should not be empty")
}

// TestTestUtilUnit_NewMockServerWithCustomResponses_Success tests NewMockServerWithCustomResponses function.
func TestTestUtilUnit_NewMockServerWithCustomResponses_Success(t *testing.T) {
	responseConfigs := []MockServerOption{
		WithTesting(t),
		WithCustomResponse("custom/response", ResponseConfig{
			StatusCode: http.StatusAccepted,
			Body:       `{"custom": "response"}`,
			Method:     "GET",
		}),
	}
	server := NewMockServer(responseConfigs...)
	defer server.Close()

	testutil.AssertNotNil(t, server, "NewMockServerWithCustomResponses should return a non-nil server")
	testutil.AssertStringNotEmpty(t, server.URL(), "Server URL should not be empty")
}

// TestTestUtilUnit_NewMockServerWithCustomErrors_Success tests NewMockServerWithCustomErrors function.
func TestTestUtilUnit_NewMockServerWithCustomErrors_Success(t *testing.T) {
	errorConfigs := []MockServerOption{
		WithTesting(t),
		WithCustomResponse("custom/error", ResponseConfig{
			StatusCode: http.StatusBadRequest,
			Body:       `{"error": "custom error"}`,
			Method:     "GET",
		}),
	}
	server := NewMockServer(errorConfigs...)
	defer server.Close()

	testutil.AssertNotNil(t, server, "NewMockServerWithCustomErrors should return a non-nil server")
	testutil.AssertStringNotEmpty(t, server.URL(), "Server URL should not be empty")
}

// TestTestUtilUnit_MockServerImpl_Success tests mockServerImpl methods.
func TestTestUtilUnit_MockServerImpl_Success(t *testing.T) {
	responses := map[string]string{
		"test/impl": `{"impl": "test"}`,
	}
	server := NewMockServer(WithSuccessResponses(responses))
	defer server.Close()

	// Test URL() method
	url := server.URL()
	testutil.AssertStringNotEmpty(t, url, "URL should not be empty")

	// Test NewTestClient() method
	client := NewTestClient(server)
	testutil.AssertNotNil(t, client, "NewTestClient should return a non-nil client")
}

// TestTestUtilUnit_TestClientImpl_Success tests testClientImpl methods.
func TestTestUtilUnit_TestClientImpl_Success(t *testing.T) {
	responses := map[string]string{
		"test/client": `{"client": "test"}`,
	}
	server := NewMockServer(WithSuccessResponses(responses))
	defer server.Close()

	client := NewTestClient(server)
	testutil.AssertNotNil(t, client, "NewTestClient should return a non-nil client")

	// Test Core() method
	core := client.Core()
	testutil.AssertNotNil(t, core, "Core should return a non-nil interface")
}

// Test the new unified API

// TestTestUtilUnit_NewMockServer_WithSuccessResponses_Success tests NewMockServer with success responses.
func TestTestUtilUnit_NewMockServer_WithSuccessResponses_Success(t *testing.T) {
	responses := map[string]string{"test-endpoint": `{"test": "data"}`}
	server := NewMockServer(WithSuccessResponses(responses))
	defer server.Close()

	testutil.AssertNotNil(t, server, "NewMockServer should return a non-nil server")
	testutil.AssertStringContains(t, server.URL(), "https://", "Server URL should be HTTPS")
}

// TestTestUtilUnit_NewMockServer_WithErrorResponses_Success tests NewMockServer with error responses.
func TestTestUtilUnit_NewMockServer_WithErrorResponses_Success(t *testing.T) {
	errorPaths := []string{"error-endpoint"}
	server := NewMockServer(WithErrorResponses(errorPaths, 500))
	defer server.Close()

	testutil.AssertNotNil(t, server, "NewMockServer should return a non-nil server")
	testutil.AssertStringContains(t, server.URL(), "https://", "Server URL should be HTTPS")
}

// TestTestUtilUnit_NewMockServer_Mixed_Success tests NewMockServer with mixed responses.
func TestTestUtilUnit_NewMockServer_Mixed_Success(t *testing.T) {
	server := NewMockServer(
		WithSuccessResponse("success-path", `{"status": "ok"}`),
		WithErrorResponse("error-path", 404),
		WithCustomResponse("custom-path", ResponseConfig{
			StatusCode: 202,
			Body:       `{"custom": "response"}`,
			Method:     "POST",
		}),
	)
	defer server.Close()

	testutil.AssertNotNil(t, server, "NewMockServer should return a non-nil server")
	testutil.AssertStringContains(t, server.URL(), "https://", "Server URL should be HTTPS")
}

// TestTestUtilUnit_NewMockServer_WithTesting_Success tests NewMockServer with testing context.
func TestTestUtilUnit_NewMockServer_WithTesting_Success(t *testing.T) {
	server := NewMockServer(
		WithTesting(t),
		WithSuccessResponse("test-with-context", `{"context": "provided"}`),
	)
	defer server.Close()

	testutil.AssertNotNil(t, server, "NewMockServer should return a non-nil server")
}
