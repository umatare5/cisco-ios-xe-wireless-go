package testutil

import (
	"net/http"
	"net/http/httptest"
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

// foreignMockServer is a MockServer this package did not construct, which is the case the
// removed type assertion turned away.
type foreignMockServer struct{ server *httptest.Server }

func (f *foreignMockServer) URL() string { return f.server.URL }
func (f *foreignMockServer) Close()      { f.server.Close() }

// brokenMockServer reports a URL no parser accepts.
type brokenMockServer struct{}

func (brokenMockServer) URL() string { return "://" }
func (brokenMockServer) Close()      {}

// TestTestUtilUnit_NewTestClient_ForeignMockServer_Success pins that NewTestClient needs
// nothing but the URL. A test supplying its own handler used to be met with a panic.
func TestTestUtilUnit_NewTestClient_ForeignMockServer_Success(t *testing.T) {
	inner := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	server := &foreignMockServer{server: inner}
	defer server.Close()

	_, isOwn := MockServer(server).(*mockServerImpl)
	testutil.AssertFalse(t, isOwn, "the fixture must not be this package's own MockServer")

	client := NewTestClient(server)
	testutil.AssertTrue(t, client != nil, "NewTestClient should accept any MockServer")
	testutil.AssertTrue(t, client.Core() != nil, "Client core should not be nil")
}

// TestTestUtilUnit_NewMockServerFromHTTP_Success pins the adapter that lets a recording
// RESTCONFServer be handed to NewTestClient.
func TestTestUtilUnit_NewMockServerFromHTTP_Success(t *testing.T) {
	recorder := NewRESTCONFServer(t)
	defer recorder.Close()

	adapted := NewMockServerFromHTTP(recorder.Server)
	testutil.AssertTrue(t, adapted != nil, "NewMockServerFromHTTP should return a MockServer")
	testutil.AssertStringEquals(t, adapted.URL(), recorder.URL, "adapted URL")

	client := NewTestClient(adapted)
	testutil.AssertTrue(t, client.Core() != nil, "Client core should not be nil")
}

// TestTestUtilUnit_NewTestClient_UnparseableURL_Panics pins the boundary of what the
// removed type assertion cost: a MockServer of any type is accepted, but a URL that does
// not parse still stops the test rather than handing back a client aimed nowhere.
func TestTestUtilUnit_NewTestClient_UnparseableURL_Panics(t *testing.T) {
	var recovered any
	func() {
		defer func() { recovered = recover() }()
		NewTestClient(brokenMockServer{})
	}()

	testutil.AssertTrue(t, recovered != nil, "NewTestClient should panic on an unparseable URL")
	message, isString := recovered.(string)
	testutil.AssertTrue(t, isString, "the panic value should be a string")
	testutil.AssertStringContains(t, message, "failed to parse server URL", "panic message")
}
