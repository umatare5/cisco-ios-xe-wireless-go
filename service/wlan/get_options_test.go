package wlan

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"sync"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil"
)

// queryRecorder stores the raw query string of the request the server received last.
type queryRecorder struct {
	mu  sync.Mutex
	raw string
}

func (q *queryRecorder) set(raw string) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.raw = raw
}

func (q *queryRecorder) get() string {
	q.mu.Lock()
	defer q.mu.Unlock()
	return q.raw
}

// newRecordingService starts a server that records the query of each request and answers
// with body, then returns a service bound to it. The mocks in pkg/testutil match on the path
// alone, so they cannot observe the query. body must be the envelope of the node under read:
// core.Get holds the sole top-level key against the endpoint.
func newRecordingService(t *testing.T, body string) (Service, *queryRecorder) {
	t.Helper()

	recorder := &queryRecorder{}
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		recorder.set(r.URL.RawQuery)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(body))
	}))
	t.Cleanup(server.Close)

	parsed, err := url.Parse(server.URL)
	testutil.AssertNoError(t, err, "parse test server URL")

	client, err := core.New(parsed.Host, "test-token", core.WithInsecureSkipVerify(true))
	testutil.AssertNoError(t, err, "create core client")

	return NewService(client), recorder
}

// assertQueryBothDirections calls read twice: once without an option, where the wire must
// carry no query, and once with WithDefaults, where the wire must carry the parameter.
func assertQueryBothDirections(t *testing.T, read func(ctx context.Context, opts ...core.GetOption) error,
	recorder *queryRecorder,
) {
	t.Helper()
	ctx := context.Background()

	testutil.AssertNoError(t, read(ctx), "read without option")
	testutil.AssertStringEquals(t, recorder.get(), "", "RawQuery without option")

	testutil.AssertNoError(t, read(ctx, core.WithDefaults(core.DefaultsReportAll)), "read with option")
	testutil.AssertStringEquals(t, recorder.get(), "with-defaults=report-all", "RawQuery with option")
}

func TestWlanServiceUnit_GetOptions_ConfigRouteWireQuery(t *testing.T) {
	tests := []struct {
		name string
		body string
		read func(s Service) func(ctx context.Context, opts ...core.GetOption) error
	}{
		{
			name: "GetConfig",
			body: `{"Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-data": {}}`,
			read: func(s Service) func(ctx context.Context, opts ...core.GetOption) error {
				return func(ctx context.Context, opts ...core.GetOption) error {
					_, err := s.GetConfig(ctx, opts...)
					return err
				}
			},
		},
		{
			name: "ListWlanCfgEntries",
			body: `{"Cisco-IOS-XE-wireless-wlan-cfg:wlan-cfg-entries": {}}`,
			read: func(s Service) func(ctx context.Context, opts ...core.GetOption) error {
				return func(ctx context.Context, opts ...core.GetOption) error {
					_, err := s.ListWlanCfgEntries(ctx, opts...)
					return err
				}
			},
		},
		{
			name: "ListWlanPolicies",
			body: `{"Cisco-IOS-XE-wireless-wlan-cfg:wlan-policies": {}}`,
			read: func(s Service) func(ctx context.Context, opts ...core.GetOption) error {
				return func(ctx context.Context, opts ...core.GetOption) error {
					_, err := s.ListWlanPolicies(ctx, opts...)
					return err
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service, recorder := newRecordingService(t, tt.body)
			assertQueryBothDirections(t, tt.read(service), recorder)
		})
	}
}

func TestWlanServiceUnit_GetOptions_OperRouteWireQuery(t *testing.T) {
	service, recorder := newRecordingService(t,
		`{"Cisco-IOS-XE-wireless-wlan-global-oper:wlan-global-oper-data": {}}`)

	assertQueryBothDirections(t, func(ctx context.Context, opts ...core.GetOption) error {
		_, err := service.GetOperational(ctx, opts...)
		return err
	}, recorder)
}
