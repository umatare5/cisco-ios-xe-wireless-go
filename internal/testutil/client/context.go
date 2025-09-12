package client

import (
	"context"
	"testing"
	"time"
)

// TestContext creates a test context with timeout.
func TestContext(t *testing.T) context.Context {
	return TestContextWithTimeout(t, 30*time.Second)
}

// TestContextWithTimeout creates a test context with custom timeout.
func TestContextWithTimeout(t *testing.T, timeout time.Duration) context.Context {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	t.Cleanup(cancel)
	return ctx
}
