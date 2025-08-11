package tests

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf"
)

// NewRESTCONFSuccessServer creates an HTTPS test server that returns 200 OK with the provided
// JSON body for each RESTCONF endpoint given. Keys in the map are endpoint strings without the
// RESTCONF prefix (e.g., "Cisco-...:container/sub"). Values are raw JSON payloads to return.
// Any non-matching path returns 404.
func NewRESTCONFSuccessServer(endpoints map[string]string) *httptest.Server { //nolint:revive // test helper
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Trim the standard RESTCONF prefix so callers can pass pure endpoints.
		ep := strings.TrimPrefix(r.URL.Path, restconf.RESTCONFPathPrefix+"/")
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
// Non-listed paths return 404.
func NewRESTCONFErrorServer(paths []string, status int) *httptest.Server { //nolint:revive // test helper
	set := make(map[string]struct{}, len(paths))
	for _, p := range paths {
		set[p] = struct{}{}
	}
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ep := strings.TrimPrefix(r.URL.Path, restconf.RESTCONFPathPrefix+"/")
		if _, ok := set[ep]; ok {
			http.Error(w, http.StatusText(status), status)
			return
		}
		http.NotFound(w, r)
	})
	return httptest.NewTLSServer(handler)
}

// NewTLSClientForServer constructs a core.Client configured to talk to the given TLS test server.
// It sets a small timeout and disables certificate verification for the test server's self-signed cert.
func NewTLSClientForServer(t *testing.T, srv *httptest.Server) *core.Client { //nolint:revive // test helper
	t.Helper()
	// Allow tests to inject malformed URL or client creation failure for coverage.
	parse := url.Parse
	if parseURLHook != nil {
		parse = parseURLHook
	}
	u, err := parse(srv.URL)
	if err != nil {
		fatalfHook(t, "failed to parse server URL: %v", err)
		return nil
	}

	newCore := core.New
	if newCoreHook != nil {
		newCore = newCoreHook
	}
	c, err := newCore(u.Host, "token",
		core.WithInsecureSkipVerify(true),
		core.WithTimeout(5*time.Second),
	)
	if err != nil {
		fatalfHook(t, "failed to create core client: %v", err)
		return nil
	}
	return c
}

// Hooks for testing coverage of error paths; set only in tests and reset after.
var (
	parseURLHook func(rawURL string) (*url.URL, error)
	newCoreHook  func(controller, token string, opts ...core.Option) (*core.Client, error)
	fatalfHook   = func(t *testing.T, format string, args ...any) { t.Fatalf(format, args...) }
)
