package mock

import (
	"net/http"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/helper"
)

func TestTestUtilMockUnit_NewRESTCONFSuccessServer_Success(t *testing.T) {
	endpoints := map[string]string{
		"test/endpoint": `{"result": "success"}`,
		"ap/config":     `{"ap": {"name": "test-ap"}}`,
	}

	server := NewRESTCONFSuccessServer(endpoints)
	defer server.Close()

	if server == nil {
		helper.AssertNotNil(t, server, "NewRESTCONFSuccessServer should return a non-nil server")
		return
	}

	// Test that server responds correctly
	client := server.Client()
	if client == nil {
		helper.AssertNotNil(t, client, "Server client should not be nil")
		return
	}
	resp, err := client.Get(server.URL + "/restconf/data/test/endpoint")
	if err != nil {
		helper.AssertNoError(t, err, "Failed to make GET request")
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		helper.AssertIntEquals(t, resp.StatusCode, http.StatusOK, "Expected status 200")
	}
}

func TestTestUtilMockUnit_NewRESTCONFErrorServer_Success(t *testing.T) {
	paths := []string{"test/error", "ap/error"}
	status := http.StatusInternalServerError

	server := NewRESTCONFErrorServer(paths, status)
	defer server.Close()

	if server == nil {
		helper.AssertNotNil(t, server, "NewRESTCONFErrorServer should return a non-nil server")
		return
	}

	// Test that server returns expected error
	client := server.Client()
	if client == nil {
		helper.AssertNotNil(t, client, "Server client should not be nil")
		return
	}
	resp, err := client.Get(server.URL + "/restconf/data/test/error")
	if err != nil {
		helper.AssertNoError(t, err, "Failed to make GET request")
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != status {
		helper.AssertIntEquals(t, resp.StatusCode, status, "Expected status to match")
	}
}

func TestTestUtilMockUnit_NewTLSClientForServer_Success(t *testing.T) {
	endpoints := map[string]string{
		"test": `{"result": "test"}`,
	}
	server := NewRESTCONFSuccessServer(endpoints)
	if server == nil {
		helper.AssertNotNil(t, server, "NewRESTCONFSuccessServer should return a non-nil server")
		return
	}
	defer server.Close()

	client := NewTLSClientForServer(t, server)
	if client == nil {
		helper.AssertNotNil(t, client, "NewTLSClientForServer should return a non-nil client")
		return
	}
}

func TestTestUtilMockUnit_NewRESTCONFServer_Success(t *testing.T) {
	server := NewRESTCONFServer(t)
	if server == nil {
		helper.AssertNotNil(t, server, "NewRESTCONFServer should return a non-nil server")
		return
	}
	defer server.Close()

	if server.handlers == nil {
		helper.AssertNotNil(t, server.handlers, "Server handlers should be initialized")
	}
}

func TestTestUtilMockUnit_RESTCONFServerAddHandler_Success(t *testing.T) {
	server := NewRESTCONFServer(t)
	if server == nil {
		helper.AssertNotNil(t, server, "NewRESTCONFServer should return a non-nil server")
		return
	}
	defer server.Close()

	// Add a handler
	server.AddHandler("GET", "test/handler", func() (int, string) {
		return http.StatusOK, `{"test": "handler"}`
	})

	// Verify handler was added
	if server.handlers["GET"] == nil {
		helper.AssertNotNil(t, server.handlers["GET"], "GET handlers should be initialized")
	}
	if server.handlers["GET"]["test/handler"] == nil {
		helper.AssertNotNil(t, server.handlers["GET"]["test/handler"], "Handler should be registered")
	}

	// Test the handler
	client := server.Client()
	if client == nil {
		helper.AssertNotNil(t, client, "Server client should not be nil")
		return
	}
	resp, err := client.Get(server.URL + "/restconf/data/test/handler")
	if err != nil {
		helper.AssertNoError(t, err, "Failed to make GET request")
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		helper.AssertIntEquals(t, resp.StatusCode, http.StatusOK, "Expected status 200")
	}
}

func TestTestUtilMockUnit_RESTCONFServerAddNotFoundHandler_Success(t *testing.T) {
	server := NewRESTCONFServer(t)
	helper.AssertNotNil(t, server, "NewRESTCONFServer should return a non-nil server")
	if server == nil {
		return
	}
	defer server.Close()

	// Add a not found handler
	server.AddNotFoundHandler("GET", "test/notfound")

	// Test the handler
	client := server.Client()
	helper.AssertNotNil(t, client, "Server client should not be nil")
	if client == nil {
		return
	}
	resp, err := client.Get(server.URL + "/restconf/data/test/notfound")
	helper.AssertNoError(t, err, "Failed to make GET request")
	defer resp.Body.Close()

	helper.AssertIntEquals(t, resp.StatusCode, http.StatusNotFound, "Expected status 404")
}

func TestTestUtilMockUnit_RESTCONFServerDefaultNotFound_Success(t *testing.T) {
	server := NewRESTCONFServer(t)
	helper.AssertNotNil(t, server, "NewRESTCONFServer should return a non-nil server")
	if server == nil {
		return
	}
	defer server.Close()

	// Test unregistered endpoint returns 404
	client := server.Client()
	helper.AssertNotNil(t, client, "Server client should not be nil")
	if client == nil {
		return
	}
	resp, err := client.Get(server.URL + "/restconf/data/nonexistent")
	helper.AssertNoError(t, err, "Failed to make GET request")
	defer resp.Body.Close()

	helper.AssertIntEquals(t, resp.StatusCode, http.StatusNotFound, "Expected status 404")
}

func TestTestUtilMockUnit_RESTCONFSuccessServerPathHandling_Success(t *testing.T) {
	endpoints := map[string]string{
		"exact/match": `{"type": "exact"}`,
	}

	server := NewRESTCONFSuccessServer(endpoints)
	defer server.Close()

	helper.AssertNotNil(t, server, "NewRESTCONFSuccessServer should return a non-nil server")
	if server == nil {
		return
	}

	client := server.Client()
	helper.AssertNotNil(t, client, "Server client should not be nil")
	if client == nil {
		return
	}

	// Test exact match
	resp, err := client.Get(server.URL + "/restconf/data/exact/match")
	helper.AssertNoError(t, err, "Failed to make GET request")
	resp.Body.Close()

	helper.AssertIntEquals(t, resp.StatusCode, http.StatusOK, "Expected status 200 for exact match")

	// Test non-match returns 404
	resp, err = client.Get(server.URL + "/restconf/data/no/match")
	helper.AssertNoError(t, err, "Failed to make GET request")
	resp.Body.Close()

	helper.AssertIntEquals(t, resp.StatusCode, http.StatusNotFound, "Expected status 404 for non-match")
}
