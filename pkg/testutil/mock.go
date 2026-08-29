package testutil

import (
	"io"
	"net/http"
	"net/http/httptest"
	"slices"
	"strings"
	"sync"
	"testing"
)

// HandlerMap represents the handler mapping structure for mock servers.
type HandlerMap[T any] = map[string]map[string]T

// RecordedRequest is one request a RESTCONFServer answered.
//
// Body is recorded although no handler reads it: whether a replacing write carried a leaf the
// caller never named is decided by the body, and the handler signature cannot report it. It is a
// string rather than a []byte so Requests' copy cannot be edited through the slice it returns.
type RecordedRequest struct {
	Method   string
	Path     string
	RawQuery string
	Body     string
}

// RESTCONFServer provides a flexible mock RESTCONF server for testing.
type RESTCONFServer struct {
	*httptest.Server
	handlers HandlerMap[func() (int, string)] // method -> path -> handler
	mu       sync.Mutex
	requests []RecordedRequest
}

// NewRESTCONFSuccessServer creates an HTTPS test server that returns 200 OK with the provided
// JSON body for each RESTCONF endpoint given. Keys in the map are endpoint strings without the
// RESTCONF prefix (e.g., "Cisco-...:container/sub" or "operations/Cisco-...:rpc"). Values are raw JSON payloads to return.
// Any non-matching path returns 404.
func NewRESTCONFSuccessServer(endpoints map[string]string) *httptest.Server {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ep := r.URL.Path
		ep = strings.TrimPrefix(ep, "/restconf/data/")
		ep = strings.TrimPrefix(ep, "/restconf/operations/")

		if body, ok := endpoints[ep]; ok {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(body))
			return
		}
		http.NotFound(w, r)
	})
	return httptest.NewTLSServer(handler)
}

// NewRESTCONFErrorServer creates an HTTPS test server that returns the provided status code
// for all listed RESTCONF endpoints. Endpoints should be provided without the RESTCONF prefix.
// Supports both data operations and RPC operations paths.
// Non-listed paths return 404.
func NewRESTCONFErrorServer(paths []string, status int) *httptest.Server { //nolint:revive // test helper
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Handle both data operations (/restconf/data/) and RPC operations (/restconf/operations/)
		ep := r.URL.Path
		ep = strings.TrimPrefix(ep, "/restconf/data/")
		ep = strings.TrimPrefix(ep, "/restconf/operations/")

		for _, p := range paths {
			if ep == p {
				http.Error(w, http.StatusText(status), status)
				return
			}
		}
		http.NotFound(w, r)
	})
	return httptest.NewTLSServer(handler)
}

// NewRESTCONFServer creates a new flexible RESTCONF server for testing.
func NewRESTCONFServer(t *testing.T) *RESTCONFServer {
	t.Helper()
	server := &RESTCONFServer{
		handlers: nil,
	}

	// Helper functions for cleaner handler logic
	normalizePath := func(path string) string {
		path = strings.TrimPrefix(path, "/restconf/data/")
		path = strings.TrimPrefix(path, "/restconf/operations/")
		return path
	}

	writeResponse := func(w http.ResponseWriter, status int, body string) {
		w.WriteHeader(status)
		if body != "" {
			_, _ = w.Write([]byte(body))
		}
	}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		server.record(r)

		if handler := server.handlerFor(r.Method, normalizePath(r.URL.Path)); handler != nil {
			status, body := handler()
			writeResponse(w, status, body)
			return
		}

		http.NotFound(w, r)
	})

	server.Server = httptest.NewTLSServer(handler)
	return server
}

// AddHandler adds a handler for a specific HTTP method and path pattern.
//
// The handler map is written under the same lock the request log uses, so a server may be
// given handlers while it is answering: guarding only one of the two would leave the type
// looking concurrency-safe while a race remained.
func (s *RESTCONFServer) AddHandler(method, pathPrefix string, handler func() (int, string)) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.handlers == nil {
		s.handlers = make(HandlerMap[func() (int, string)])
	}
	if s.handlers[method] == nil {
		s.handlers[method] = make(map[string]func() (int, string))
	}
	s.handlers[method][pathPrefix] = handler
}

// handlerFor resolves a request to its handler, holding the lock across the lookup and
// releasing it before the handler runs so a handler may call back into the server.
func (s *RESTCONFServer) handlerFor(method, path string) func() (int, string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	methodHandlers, ok := s.handlers[method]
	if !ok {
		return nil
	}
	return findHandler(path, methodHandlers)
}

// findHandler returns the handler whose registered prefix matches path most closely.
//
// The longest match wins because map iteration order is unspecified: with both a parent
// and a child path registered, a first-match loop serves a different body per run.
func findHandler(path string, methodHandlers map[string]func() (int, string)) func() (int, string) {
	var best string
	var found func() (int, string)

	for pathPrefix, handler := range methodHandlers {
		if strings.Contains(path, pathPrefix) && len(pathPrefix) > len(best) {
			best, found = pathPrefix, handler
		}
	}

	return found
}

// record keeps a request for a later assertion. The query is recorded although the
// handler lookup ignores it, which is what makes a GetOption observable.
//
// The body is read before the lock is taken, so the mutex the handler map shares is not held
// across a read from the connection. A read failure records an empty body, which the assertion
// on Body is what reports.
func (s *RESTCONFServer) record(r *http.Request) {
	body, _ := io.ReadAll(r.Body)

	s.mu.Lock()
	defer s.mu.Unlock()
	s.requests = append(s.requests, RecordedRequest{
		Method:   r.Method,
		Path:     r.URL.Path,
		RawQuery: r.URL.RawQuery,
		Body:     string(body),
	})
}

// Requests returns the requests the server answered, oldest first.
func (s *RESTCONFServer) Requests() []RecordedRequest {
	s.mu.Lock()
	defer s.mu.Unlock()
	return slices.Clone(s.requests)
}
