package tests

import (
	"context"
	"testing"
	"time"
)

// TestContext creates a test context with timeout (noinline to ensure coverage accounting)
//
//go:noinline
func TestContext(t *testing.T) context.Context { return TestContextWithTimeout(t, DefaultTestTimeout) }

// TestContextWithTimeout creates a test context with custom timeout
func TestContextWithTimeout(t *testing.T, timeout time.Duration) context.Context {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	t.Cleanup(cancel)
	return ctx
}
