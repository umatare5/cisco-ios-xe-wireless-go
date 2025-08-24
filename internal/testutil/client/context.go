package client

import (
	"context"
	"testing"
	"time"
)

// ShortTestTimeout is the standard timeout for short tests
const ShortTestTimeout = 5 * time.Second

// DefaultTestTimeout is the standard timeout for tests
const DefaultTestTimeout = 30 * time.Second

// ExtendedTestTimeout is the extended timeout for long-running tests
const ExtendedTestTimeout = 60 * time.Second

// executedSkipIfNoConnection tracks if skip connection check was executed (for testing coverage)
var executedSkipIfNoConnection bool

// TestContext creates a test context with timeout (noinline to ensure coverage accounting)
//
//go:noinline
func TestContext(t *testing.T) context.Context {
	return TestContextWithTimeout(t, DefaultTestTimeout)
}

// TestContextWithTimeout creates a test context with custom timeout
func TestContextWithTimeout(t *testing.T, timeout time.Duration) context.Context {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	t.Cleanup(cancel)
	return ctx
}
