package testutil

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// HandlerMap represents the handler mapping structure for mock servers.
type HandlerMap[T any] = map[string]map[string]T

// RESTCONFServer provides a flexible mock RESTCONF server for testing.
type RESTCONFServer struct {
	*httptest.Server
	handlers HandlerMap[func() (int, string)] // method -> path -> handler
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

	findHandler := func(path string, methodHandlers map[string]func() (int, string)) func() (int, string) {
		for pathPrefix, handler := range methodHandlers {
			if strings.Contains(path, pathPrefix) {
				return handler
			}
		}
		return nil
	}

	writeResponse := func(w http.ResponseWriter, status int, body string) {
		w.WriteHeader(status)
		if body != "" {
			_, _ = w.Write([]byte(body))
		}
	}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := normalizePath(r.URL.Path)
		methodHandlers, ok := server.handlers[r.Method]
		if !ok {
			http.NotFound(w, r)
			return
		}

		if handler := findHandler(path, methodHandlers); handler != nil {
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
func (s *RESTCONFServer) AddHandler(method, pathPrefix string, handler func() (int, string)) {
	if s.handlers == nil {
		s.handlers = make(HandlerMap[func() (int, string)])
	}
	if s.handlers[method] == nil {
		s.handlers[method] = make(map[string]func() (int, string))
	}
	s.handlers[method][pathPrefix] = handler
}
