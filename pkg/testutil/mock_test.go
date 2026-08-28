package testutil

import (
	"net/http"
	"strconv"
	"strings"
	"sync"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil"
)

func TestTestUtilUnit_NewRESTCONFSuccessServer_Success(t *testing.T) {
	endpoints := map[string]string{
		"test/endpoint": `{"result": "success"}`,
		"ap/config":     `{"ap": {"name": "test-ap"}}`,
	}

	server := NewRESTCONFSuccessServer(endpoints)
	defer server.Close()

	if server == nil {
		testutil.AssertNotNil(t, server, "NewRESTCONFSuccessServer should return a non-nil server")
		return
	}

	// Test that server responds correctly
	client := server.Client()
	if client == nil {
		testutil.AssertNotNil(t, client, "Server client should not be nil")
		return
	}
	resp, err := client.Get(server.URL + "/restconf/data/test/endpoint")
	if err != nil {
		testutil.AssertNoError(t, err, "Failed to make GET request")
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		testutil.AssertIntEquals(t, resp.StatusCode, http.StatusOK, "Expected status 200")
	}
}

func TestTestUtilUnit_NewRESTCONFErrorServer_Success(t *testing.T) {
	paths := []string{"test/error", "ap/error"}
	status := http.StatusInternalServerError

	server := NewRESTCONFErrorServer(paths, status)
	defer server.Close()

	if server == nil {
		testutil.AssertNotNil(t, server, "NewRESTCONFErrorServer should return a non-nil server")
		return
	}

	// Test that server returns expected error
	client := server.Client()
	if client == nil {
		testutil.AssertNotNil(t, client, "Server client should not be nil")
		return
	}
	resp, err := client.Get(server.URL + "/restconf/data/test/error")
	if err != nil {
		testutil.AssertNoError(t, err, "Failed to make GET request")
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != status {
		testutil.AssertIntEquals(t, resp.StatusCode, status, "Expected status to match")
	}
}

func TestTestUtilUnit_NewRESTCONFServer_Success(t *testing.T) {
	server := NewRESTCONFServer(t)
	if server == nil {
		testutil.AssertNotNil(t, server, "NewRESTCONFServer should return a non-nil server")
		return
	}
	defer server.Close()

	if server.handlers == nil {
		testutil.AssertNotNil(t, server.handlers, "Server handlers should be initialized")
	}
}

func TestTestUtilUnit_RESTCONFServerAddHandler_Success(t *testing.T) {
	server := NewRESTCONFServer(t)
	if server == nil {
		testutil.AssertNotNil(t, server, "NewRESTCONFServer should return a non-nil server")
		return
	}
	defer server.Close()

	// Add a handler
	server.AddHandler("GET", "test/handler", func() (int, string) {
		return http.StatusOK, `{"test": "handler"}`
	})

	// Verify handler was added
	if server.handlers["GET"] == nil {
		testutil.AssertNotNil(t, server.handlers["GET"], "GET handlers should be initialized")
	}
	if server.handlers["GET"]["test/handler"] == nil {
		testutil.AssertNotNil(t, server.handlers["GET"]["test/handler"], "Handler should be registered")
	}

	// Test the handler
	client := server.Client()
	if client == nil {
		testutil.AssertNotNil(t, client, "Server client should not be nil")
		return
	}
	resp, err := client.Get(server.URL + "/restconf/data/test/handler")
	if err != nil {
		testutil.AssertNoError(t, err, "Failed to make GET request")
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		testutil.AssertIntEquals(t, resp.StatusCode, http.StatusOK, "Expected status 200")
	}
}

func TestTestUtilUnit_RESTCONFServerAddNotFoundHandler_Success(t *testing.T) {
	server := NewRESTCONFServer(t)
	testutil.AssertNotNil(t, server, "NewRESTCONFServer should return a non-nil server")
	if server == nil {
		return
	}
	defer server.Close()

	// Test the handler
	client := server.Client()
	testutil.AssertNotNil(t, client, "Server client should not be nil")
	if client == nil {
		return
	}
	resp, err := client.Get(server.URL + "/restconf/data/test/notfound")
	testutil.AssertNoError(t, err, "Failed to make GET request")
	defer resp.Body.Close()

	testutil.AssertIntEquals(t, resp.StatusCode, http.StatusNotFound, "Expected status 404")
}

func TestTestUtilUnit_RESTCONFServerDefaultNotFound_Success(t *testing.T) {
	server := NewRESTCONFServer(t)
	testutil.AssertNotNil(t, server, "NewRESTCONFServer should return a non-nil server")
	if server == nil {
		return
	}
	defer server.Close()

	// Test unregistered endpoint returns 404
	client := server.Client()
	testutil.AssertNotNil(t, client, "Server client should not be nil")
	if client == nil {
		return
	}
	resp, err := client.Get(server.URL + "/restconf/data/nonexistent")
	testutil.AssertNoError(t, err, "Failed to make GET request")
	defer resp.Body.Close()

	testutil.AssertIntEquals(t, resp.StatusCode, http.StatusNotFound, "Expected status 404")
}

func TestTestUtilUnit_RESTCONFSuccessServerPathHandling_Success(t *testing.T) {
	endpoints := map[string]string{
		"exact/match": `{"type": "exact"}`,
	}

	server := NewRESTCONFSuccessServer(endpoints)
	defer server.Close()

	testutil.AssertNotNil(t, server, "NewRESTCONFSuccessServer should return a non-nil server")
	if server == nil {
		return
	}

	client := server.Client()
	testutil.AssertNotNil(t, client, "Server client should not be nil")
	if client == nil {
		return
	}

	// Test exact match
	resp, err := client.Get(server.URL + "/restconf/data/exact/match")
	testutil.AssertNoError(t, err, "Failed to make GET request")
	resp.Body.Close()

	testutil.AssertIntEquals(t, resp.StatusCode, http.StatusOK, "Expected status 200 for exact match")

	// Test non-match returns 404
	resp, err = client.Get(server.URL + "/restconf/data/no/match")
	testutil.AssertNoError(t, err, "Failed to make GET request")
	resp.Body.Close()

	testutil.AssertIntEquals(t, resp.StatusCode, http.StatusNotFound, "Expected status 404 for non-match")
}

// TestTestUtilUnit_RESTCONFServer_Requests_Success pins what the recorder observes: the
// query the handler lookup ignores, a request no handler answered, and arrival order.
func TestTestUtilUnit_RESTCONFServer_Requests_Success(t *testing.T) {
	server := NewRESTCONFServer(t)
	defer server.Close()
	server.AddHandler(http.MethodGet, "probe", func() (int, string) { return http.StatusOK, `{}` })

	testutil.AssertIntEquals(t, len(server.Requests()), 0, "Requests() before any request")

	client := server.Client()
	for _, suffix := range []string{"probe?depth=3&fields=probe(one)", "absent"} {
		resp, err := client.Get(server.URL + "/restconf/data/" + suffix)
		testutil.AssertNoError(t, err, "Failed to make GET request")
		resp.Body.Close()
	}

	recorded := server.Requests()
	testutil.AssertIntEquals(t, len(recorded), 2, "Requests() count")

	testutil.AssertStringEquals(t, recorded[0].Method, http.MethodGet, "first method")
	testutil.AssertStringEquals(t, recorded[0].Path, "/restconf/data/probe", "first path")
	testutil.AssertStringEquals(t, recorded[0].RawQuery, "depth=3&fields=probe(one)", "first query")

	// A 404 is recorded too: record runs before the handler lookup, so a test can tell a
	// request that missed every handler from one that was never sent.
	testutil.AssertStringEquals(t, recorded[1].Path, "/restconf/data/absent", "second path")
	testutil.AssertStringEquals(t, recorded[1].RawQuery, "", "second query")
}

// TestTestUtilUnit_RESTCONFServer_RequestsIsACopy_Success pins that a caller cannot edit
// the server's record through the slice it was handed.
func TestTestUtilUnit_RESTCONFServer_RequestsIsACopy_Success(t *testing.T) {
	server := NewRESTCONFServer(t)
	defer server.Close()
	server.AddHandler(http.MethodGet, "probe", func() (int, string) { return http.StatusOK, `{}` })

	resp, err := server.Client().Get(server.URL + "/restconf/data/probe")
	testutil.AssertNoError(t, err, "Failed to make GET request")
	resp.Body.Close()

	first := server.Requests()
	testutil.AssertIntEquals(t, len(first), 1, "Requests() count")
	first[0].RawQuery = "tampered"

	testutil.AssertStringEquals(t, server.Requests()[0].RawQuery, "", "query after caller edit")
}

// TestTestUtilUnit_RESTCONFServer_ConcurrentAddHandler_Success pins that handlers and the
// request log are guarded by the same lock, so a server may be given handlers while it is
// answering. Guarding only one of the two leaves the type looking concurrency-safe.
//
// The failure this pins is a data race, which only the race detector reports. CI runs the
// suite with -race, so the guard is live there.
func TestTestUtilUnit_RESTCONFServer_ConcurrentAddHandler_Success(t *testing.T) {
	server := NewRESTCONFServer(t)
	defer server.Close()

	var wg sync.WaitGroup
	for i := range 24 {
		wg.Add(2)

		go func(n int) {
			defer wg.Done()
			prefix := "probe" + strconv.Itoa(n)
			server.AddHandler(http.MethodGet, prefix, func() (int, string) { return http.StatusOK, `{}` })
		}(i)

		go func() {
			defer wg.Done()
			resp, err := server.Client().Get(server.URL + "/restconf/data/probe0")
			if err == nil {
				resp.Body.Close()
			}
		}()
	}
	wg.Wait()

	testutil.AssertTrue(t, len(server.Requests()) == 24, "every request is recorded")
}

// TestTestUtilUnit_RESTCONFServer_RecordsBody_Success pins the body the recorder keeps. The
// second assertion is the load-bearing one: a request that carried no body records an empty
// one rather than inheriting the body of the request before it.
func TestTestUtilUnit_RESTCONFServer_RecordsBody_Success(t *testing.T) {
	const payload = `{"leaf": "value"}`

	server := NewRESTCONFServer(t)
	defer server.Close()
	server.AddHandler(http.MethodPut, "probe", func() (int, string) { return http.StatusNoContent, "" })
	server.AddHandler(http.MethodGet, "probe", func() (int, string) { return http.StatusOK, `{}` })

	client := server.Client()
	endpoint := server.URL + "/restconf/data/probe"

	written, err := http.NewRequestWithContext(t.Context(), http.MethodPut, endpoint, strings.NewReader(payload))
	testutil.AssertNoError(t, err, "Failed to build the PUT request")
	resp, err := client.Do(written)
	testutil.AssertNoError(t, err, "Failed to make PUT request")
	resp.Body.Close()

	read, err := http.NewRequestWithContext(t.Context(), http.MethodGet, endpoint, http.NoBody)
	testutil.AssertNoError(t, err, "Failed to build the GET request")
	resp, err = client.Do(read)
	testutil.AssertNoError(t, err, "Failed to make GET request")
	resp.Body.Close()

	recorded := server.Requests()
	testutil.AssertIntEquals(t, len(recorded), 2, "Requests() count")
	testutil.AssertStringEquals(t, recorded[0].Body, payload, "recorded PUT body")
	testutil.AssertStringEquals(t, recorded[1].Body, "", "recorded GET body")
}
