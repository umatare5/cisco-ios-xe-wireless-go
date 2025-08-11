package main

import (
	"context"
	"errors"
	"flag"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
	"time"

	wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
)

// dummy client creator returning a lightweight *wnc.Client via real constructor with safe inputs.
func newTestClient(t *testing.T) *wnc.Client {
	t.Helper()
	c, err := wnc.NewClient("controller", "token", wnc.WithTimeout(1*time.Second))
	if err != nil {
		t.Fatalf("newTestClient: %v", err)
	}
	return c
}

func TestLoadConfigMissing(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"cmd"}
	oldC, oldT := os.Getenv("WNC_CONTROLLER"), os.Getenv("WNC_ACCESS_TOKEN")
	defer os.Setenv("WNC_CONTROLLER", oldC)
	defer os.Setenv("WNC_ACCESS_TOKEN", oldT)
	os.Unsetenv("WNC_CONTROLLER")
	os.Setenv("WNC_ACCESS_TOKEN", "tok")
	flag.CommandLine = flag.NewFlagSet("test", flag.ContinueOnError)
	if _, err := loadConfig(); err == nil {
		t.Fatal("expected controller error")
	}
	os.Setenv("WNC_CONTROLLER", "ctrl")
	os.Unsetenv("WNC_ACCESS_TOKEN")
	flag.CommandLine = flag.NewFlagSet("test", flag.ContinueOnError)
	if _, err := loadConfig(); err == nil {
		t.Fatal("expected token error")
	}
}

func TestLoadConfigTimeoutPrecedence(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	// Provide flag to override env timeout.
	os.Args = []string{"cmd", "-timeout", "7"}
	oldC, oldT, oldTO := os.Getenv("WNC_CONTROLLER"), os.Getenv("WNC_ACCESS_TOKEN"), os.Getenv("WNC_TIMEOUT_SEC")
	defer os.Setenv("WNC_CONTROLLER", oldC)
	defer os.Setenv("WNC_ACCESS_TOKEN", oldT)
	defer os.Setenv("WNC_TIMEOUT_SEC", oldTO)
	os.Setenv("WNC_CONTROLLER", "ctrl")
	os.Setenv("WNC_ACCESS_TOKEN", "tok")
	os.Setenv("WNC_TIMEOUT_SEC", "3")
	// Reset flag state between tests.
	flag.CommandLine = flag.NewFlagSet("test", flag.ContinueOnError)
	cfg, err := loadConfig()
	if err != nil {
		t.Fatalf("loadConfig error: %v", err)
	}
	if cfg.Timeout != 7*time.Second {
		t.Fatalf("expected flag override 7s, got %v", cfg.Timeout)
	}
}

func TestNewLoggerLevels(t *testing.T) {
	lInfo := newLogger(false)
	lDebug := newLogger(true)
	if lInfo == nil || lDebug == nil {
		t.Fatal("expected non-nil loggers")
	}
}

func TestRunWorkflowSuccess(t *testing.T) {
	c := newTestClient(t)
	// Stub service fetchers to avoid real network.
	oldAP, oldClient, oldRogue := fetchAPOper, fetchClientOper, fetchRogueOper
	defer func() { fetchAPOper, fetchClientOper, fetchRogueOper = oldAP, oldClient, oldRogue }()
	fetchAPOper = func(_ context.Context, _ *wnc.Client) (any, error) { return struct{}{}, nil }
	fetchClientOper = func(_ context.Context, _ *wnc.Client) (any, error) { return struct{}{}, nil }
	fetchRogueOper = func(_ context.Context, _ *wnc.Client) (any, error) { return struct{}{}, nil }
	if err := runWorkflow(context.Background(), c, newLogger(false)); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestRunWorkflowFailures(t *testing.T) {
	c := newTestClient(t)
	// AP failure
	oldAP := fetchAPOper
	fetchAPOper = func(_ context.Context, _ *wnc.Client) (any, error) { return nil, errors.New("ap") }
	if err := runWorkflow(context.Background(), c, newLogger(false)); err == nil || err.Error() == "" {
		t.Fatal("expected ap failure")
	}
	fetchAPOper = oldAP
	// Client failure
	oldClient := fetchClientOper
	fetchAPOper = func(_ context.Context, _ *wnc.Client) (any, error) { return struct{}{}, nil }
	fetchClientOper = func(_ context.Context, _ *wnc.Client) (any, error) { return nil, errors.New("client") }
	if err := runWorkflow(context.Background(), c, newLogger(false)); err == nil || err.Error() == "" {
		t.Fatal("expected client failure")
	}
	fetchClientOper = oldClient
	// Rogue failure
	oldRogue := fetchRogueOper
	fetchClientOper = func(_ context.Context, _ *wnc.Client) (any, error) { return struct{}{}, nil }
	fetchRogueOper = func(_ context.Context, _ *wnc.Client) (any, error) { return nil, errors.New("rogue") }
	if err := runWorkflow(context.Background(), c, newLogger(false)); err == nil || err.Error() == "" {
		t.Fatal("expected rogue failure")
	}
	fetchRogueOper = oldRogue
}

func TestClassifyError(t *testing.T) {
	if classifyError(nil) != 0 {
		t.Fatal("nil error should classify to 0")
	}
	if classifyError(wnc.ErrAuthenticationFailed) != errCategoryAuth {
		t.Fatal("auth classification failed")
	}
	if classifyError(wnc.ErrAccessForbidden) != errCategoryForbidden {
		t.Fatal("forbidden classification failed")
	}
	if classifyError(wnc.ErrResourceNotFound) != errCategoryNotFound {
		t.Fatal("notfound classification failed")
	}
	if classifyError(wnc.ErrRequestTimeout) != errCategoryTimeout {
		t.Fatal("timeout classification failed")
	}
	if classifyError(context.DeadlineExceeded) != errCategoryTimeout {
		t.Fatal("deadline should map to timeout")
	}
	// APIError mapping
	if classifyError(&wnc.APIError{StatusCode: 500}) != errCategoryRemote {
		t.Fatal("api error classification failed")
	}
	// Other
	if classifyError(errors.New("x")) != errCategoryOther {
		t.Fatal("generic error classification failed")
	}
}

func TestNotNil(t *testing.T) {
	if notNil(nil) {
		t.Fatal("expected false for nil")
	}
	if !notNil(struct{}{}) {
		t.Fatal("expected true for non-nil")
	}
}

func TestFatalAndMainFlows(t *testing.T) {
	// Intercept exit codes.
	oldExit := exitFunc
	defer func() { exitFunc = oldExit }()
	codeCh := make(chan int, 3)
	exitFunc = func(code int) { codeCh <- code }

	// Prepare config + stub workflow success.
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"cmd"}
	oldC, oldT := os.Getenv("WNC_CONTROLLER"), os.Getenv("WNC_ACCESS_TOKEN")
	defer os.Setenv("WNC_CONTROLLER", oldC)
	defer os.Setenv("WNC_ACCESS_TOKEN", oldT)
	os.Setenv("WNC_CONTROLLER", "ctrl")
	os.Setenv("WNC_ACCESS_TOKEN", "tok")
	flag.CommandLine = flag.NewFlagSet("test", flag.ContinueOnError)

	// Stub client builder to use test client quickly.
	oldBuilder := buildClient
	buildClient = func(_ *configuration, _ *slog.Logger) (*wnc.Client, error) { return newTestClient(t), nil }
	defer func() { buildClient = oldBuilder }()

	// Stub workflow success.
	oldWF := workflowFunc
	workflowFunc = func(_ context.Context, _ *wnc.Client, _ *slog.Logger) error { return nil }
	defer func() { workflowFunc = oldWF }()

	flag.CommandLine = flag.NewFlagSet("test", flag.ContinueOnError)
	main()
	select {
	case c := <-codeCh:
		if c != 0 {
			t.Fatalf("expected 0 got %d", c)
		}
	case <-time.After(time.Second):
		t.Fatal("no exit code")
	}

	// Now force workflow error classification paths.
	// Each iteration resets flags to avoid re-parse issues.
	cases := []struct {
		wf   func(context.Context, *wnc.Client, *slog.Logger) error
		want int
	}{
		{func(context.Context, *wnc.Client, *slog.Logger) error { return wnc.ErrAuthenticationFailed }, 2},
		{func(context.Context, *wnc.Client, *slog.Logger) error { return wnc.ErrAccessForbidden }, 2},
		{func(context.Context, *wnc.Client, *slog.Logger) error { return wnc.ErrResourceNotFound }, 2},
		{func(context.Context, *wnc.Client, *slog.Logger) error { return wnc.ErrRequestTimeout }, 2},
		{func(context.Context, *wnc.Client, *slog.Logger) error { return &wnc.APIError{StatusCode: 500} }, 2},
		{func(context.Context, *wnc.Client, *slog.Logger) error { return errors.New("other") }, 3},
	}
	for _, tc := range cases {
		workflowFunc = tc.wf
		flag.CommandLine = flag.NewFlagSet("test", flag.ContinueOnError)
		main()
		select {
		case c := <-codeCh:
			if c != tc.want {
				t.Fatalf("expected %d got %d", tc.want, c)
			}
		case <-time.After(time.Second):
			t.Fatal("no exit code in loop")
		}
	}
}

func TestMainConfigError(t *testing.T) {
	oldExit := exitFunc
	defer func() { exitFunc = oldExit }()
	codeCh := make(chan int, 1)
	exitFunc = func(code int) { codeCh <- code }
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"cmd"}
	oldC, oldT := os.Getenv("WNC_CONTROLLER"), os.Getenv("WNC_ACCESS_TOKEN")
	defer os.Setenv("WNC_CONTROLLER", oldC)
	defer os.Setenv("WNC_ACCESS_TOKEN", oldT)
	os.Unsetenv("WNC_CONTROLLER")
	os.Setenv("WNC_ACCESS_TOKEN", "tok")
	flag.CommandLine = flag.NewFlagSet("test", flag.ContinueOnError)
	main()
	select {
	case c := <-codeCh:
		if c != 1 {
			t.Fatalf("expected 1 got %d", c)
		}
	case <-time.After(time.Second):
		t.Fatal("no exit code")
	}
}

func TestMainClientCreationError(t *testing.T) {
	oldExit := exitFunc
	defer func() { exitFunc = oldExit }()
	codeCh := make(chan int, 1)
	exitFunc = func(code int) { codeCh <- code }
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"cmd"}
	oldC, oldT := os.Getenv("WNC_CONTROLLER"), os.Getenv("WNC_ACCESS_TOKEN")
	defer os.Setenv("WNC_CONTROLLER", oldC)
	defer os.Setenv("WNC_ACCESS_TOKEN", oldT)
	os.Setenv("WNC_CONTROLLER", "ctrl")
	os.Setenv("WNC_ACCESS_TOKEN", "tok")
	flag.CommandLine = flag.NewFlagSet("test", flag.ContinueOnError)
	// Stub buildClient to return error
	oldBuilder := buildClient
	buildClient = func(_ *configuration, _ *slog.Logger) (*wnc.Client, error) { return nil, errors.New("boom") }
	defer func() { buildClient = oldBuilder }()
	main()
	select {
	case c := <-codeCh:
		if c != 1 {
			t.Fatalf("expected 1 got %d", c)
		}
	case <-time.After(time.Second):
		t.Fatal("no exit code")
	}
}

func TestNewClientDirectSuccess(t *testing.T) {
	cfg := &configuration{Controller: "ctrl", AccessToken: "tok", Timeout: 1 * time.Second, Verbose: false}
	c, err := newClient(cfg, newLogger(false))
	if err != nil || c == nil {
		t.Fatalf("expected newClient success, got %v, %v", c, err)
	}
}

// Ensure lazy default initialization of fetch functions occurs inside newClient.
func TestNewClientAssignsDefaultFetchers(t *testing.T) {
	// Save and clear existing fetchers to force assignment paths.
	oldAP, oldClient, oldRogue := fetchAPOper, fetchClientOper, fetchRogueOper
	fetchAPOper, fetchClientOper, fetchRogueOper = nil, nil, nil
	t.Cleanup(func() { fetchAPOper, fetchClientOper, fetchRogueOper = oldAP, oldClient, oldRogue })

	cfg := &configuration{Controller: "ctrl", AccessToken: "tok", Timeout: 1 * time.Second}
	if _, err := newClient(cfg, newLogger(false)); err != nil {
		t.Fatalf("newClient error: %v", err)
	}
	if fetchAPOper == nil || fetchClientOper == nil || fetchRogueOper == nil {
		t.Fatal("expected default fetchers to be assigned")
	}
}

// Cover the error branch in newClient by providing an invalid timeout (0).
func TestNewClientErrorPath(t *testing.T) {
	cfg := &configuration{Controller: "ctrl", AccessToken: "tok", Timeout: 0}
	if _, err := newClient(cfg, newLogger(false)); err == nil {
		t.Fatal("expected error for invalid timeout")
	}
}

// Exercise default fetcher functions end-to-end against a local TLS server.
func TestNewClientDefaultFetchersExecute(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/restconf/data/Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data":{}}`))
	})
	mux.HandleFunc("/restconf/data/Cisco-IOS-XE-wireless-client-oper:client-oper-data", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"Cisco-IOS-XE-wireless-client-oper:client-oper-data":{}}`))
	})
	mux.HandleFunc("/restconf/data/Cisco-IOS-XE-wireless-rogue-oper:rogue-oper-data", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"Cisco-IOS-XE-wireless-rogue-oper:rogue-oper-data":{}}`))
	})
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	u, err := url.Parse(ts.URL)
	if err != nil {
		t.Fatalf("parse url: %v", err)
	}
	// Clear fetchers so newClient assigns defaults.
	oldAP, oldClient, oldRogue := fetchAPOper, fetchClientOper, fetchRogueOper
	fetchAPOper, fetchClientOper, fetchRogueOper = nil, nil, nil
	t.Cleanup(func() { fetchAPOper, fetchClientOper, fetchRogueOper = oldAP, oldClient, oldRogue })

	cfg := &configuration{Controller: u.Host, AccessToken: "tok", Timeout: 1 * time.Second}
	c, err := newClient(cfg, newLogger(false))
	if err != nil {
		t.Fatalf("newClient error: %v", err)
	}

	// Now call each default fetcher to ensure their function literal bodies execute.
	if _, err := fetchAPOper(context.Background(), c); err != nil {
		t.Fatalf("fetchAPOper error: %v", err)
	}
	if _, err := fetchClientOper(context.Background(), c); err != nil {
		t.Fatalf("fetchClientOper error: %v", err)
	}
	if _, err := fetchRogueOper(context.Background(), c); err != nil {
		t.Fatalf("fetchRogueOper error: %v", err)
	}
}

func TestFatalNil(t *testing.T) {
	oldExit := exitFunc
	defer func() { exitFunc = oldExit }()
	called := false
	exitFunc = func(code int) { called = true }
	fatal(9, nil) // should not call exitFunc
	if called {
		t.Fatal("exitFunc should not be called for nil error")
	}
}
