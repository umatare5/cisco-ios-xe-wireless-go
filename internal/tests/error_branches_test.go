package tests

import (
	"errors"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
)

// TestAssertNonNilResult_NilBranch exercises the nil path using simulation flag to avoid failing.
func TestAssertNonNilResult_NilBranch(t *testing.T) {
	prev := simulateAssertErrorAsLog
	simulateAssertErrorAsLog = true
	defer func() { simulateAssertErrorAsLog = prev }()
	AssertNonNilResult(t, nil, "NilBranchMethod")
}

// TestNewTLSClientForServer_ErrorPaths covers URL parse failure and client creation failure via hooks.
func TestNewTLSClientForServer_ErrorPaths(t *testing.T) {
	// 1) URL parse error
	ts := httptest.NewTLSServer(httptest.NewServer(nil).Config.Handler)
	defer ts.Close()

	prevParse, prevFatal := parseURLHook, fatalfHook
	parseURLHook = func(string) (*url.URL, error) { return nil, errors.New("bad url") }
	fatalCalled := false
	fatalfHook = func(t *testing.T, format string, args ...any) { fatalCalled = true }
	_ = NewTLSClientForServer(t, ts)
	if !fatalCalled {
		t.Fatal("expected fatalfHook to be called for parse error")
	}
	parseURLHook, fatalfHook = prevParse, prevFatal

	// 2) client creation error
	prevNew, prevFatal2 := newCoreHook, fatalfHook
	newCoreHook = func(controller, token string, opts ...core.Option) (*core.Client, error) {
		return nil, errors.New("create error")
	}
	fatalCalled = false
	fatalfHook = func(t *testing.T, format string, args ...any) { fatalCalled = true }
	_ = NewTLSClientForServer(t, ts)
	if !fatalCalled {
		t.Fatal("expected fatalfHook to be called for client create error")
	}
	newCoreHook, fatalfHook = prevNew, prevFatal2
}
