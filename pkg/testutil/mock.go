package testutil

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
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

// NewTLSClientForServer constructs a core.Client configured to talk to the given TLS test server.
// It sets a small timeout and disables certificate verification for the test server's self-signed cert.
func NewTLSClientForServer(t *testing.T, srv *httptest.Server) *core.Client { //nolint:revive // test helper
	t.Helper()
	u, err := url.Parse(srv.URL)
	if err != nil {
		t.Fatalf("failed to parse server URL: %v", err)
		return nil
	}

	c, err := core.New(u.Host, "token",
		core.WithInsecureSkipVerify(true),
		core.WithTimeout(5*time.Second),
	)
	if err != nil {
		t.Fatalf("failed to create core client: %v", err)
		return nil
	}
	return c
}

// NewRESTCONFServer creates a new flexible RESTCONF server for testing.
func NewRESTCONFServer(t *testing.T) *RESTCONFServer {
	t.Helper()
	server := &RESTCONFServer{
		handlers: nil,
	}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		method := r.Method
		path := r.URL.Path
		path = strings.TrimPrefix(path, "/restconf/data/")
		path = strings.TrimPrefix(path, "/restconf/operations/")

		methodHandlers, ok := server.handlers[method]
		if !ok {
			http.NotFound(w, r)
			return
		}

		for pathPrefix, handler := range methodHandlers {
			if !strings.Contains(path, pathPrefix) {
				continue
			}
			status, body := handler()
			w.WriteHeader(status)
			if body != "" {
				_, _ = w.Write([]byte(body))
			}
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
