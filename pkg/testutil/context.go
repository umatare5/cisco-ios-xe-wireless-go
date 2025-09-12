package testutil

import (
	"context"
	"testing"
	"time"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/transport"
)

// TestContext creates a test context with appropriate timeout.
// Following Prometheus patterns for test context management.
func TestContext(t *testing.T) context.Context {
	ctx, cancel := context.WithTimeout(context.Background(), transport.QuickTimeout)
	t.Cleanup(cancel)
	return ctx
}

// TestContextWithTimeout creates a test context with custom timeout.
// Following Kubernetes patterns for configurable test timeouts.
func TestContextWithTimeout(t *testing.T, timeout time.Duration) context.Context {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	t.Cleanup(cancel)
	return ctx
}
