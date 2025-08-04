// Package testutils provides common test utilities for the cisco-ios-xe-wireless-go library.
package testutils

import (
	"context"
	"testing"
	"time"

	wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil"
)

// GetTestClient returns a test client using environment variables.
// This is a common helper used across all test packages.
func GetTestClient(t *testing.T) *wnc.Client {
	t.Helper()

	// Check if we can create the client
	config := testutil.NewTestConfigFromEnv()
	if config == nil {
		t.Skip("Required environment variables not set - skipping test")
		return nil
	}

	return testutil.CreateTestClientFromEnv(t)
}

// GetTestClientWithTimeout returns a test client with specified timeout.
func GetTestClientWithTimeout(t *testing.T, timeout time.Duration) *wnc.Client {
	t.Helper()

	// Check if we can create the client
	config := testutil.NewTestConfigFromEnv()
	if config == nil {
		t.Skip("Required environment variables not set - skipping test")
		return nil
	}

	return testutil.CreateTestClientWithTimeout(t, timeout)
}

// GetTestClientWithContext returns a test client with context.
func GetTestClientWithContext(t *testing.T, ctx context.Context) *wnc.Client {
	t.Helper()

	// Check if we can create the client
	config := testutil.NewTestConfigFromEnv()
	if config == nil {
		t.Skip("Required environment variables not set - skipping test")
		return nil
	}

	client := testutil.CreateTestClientFromEnv(t)
	return client
}

// ValidateClient performs common client validation checks.
func ValidateClient(t *testing.T, client *wnc.Client) {
	t.Helper()
	if client == nil {
		t.Fatal("Client should not be nil")
	}
}

// SaveTestDataWithLogging saves test data to a file with logging.
func SaveTestDataWithLogging(filename string, data interface{}) {
	testutil.SaveTestDataWithLogging(filename, data)
}

// CreateTestContext creates a context with the specified timeout.
func CreateTestContext(timeout time.Duration) (context.Context, context.CancelFunc) {
	return testutil.CreateTestContext(timeout)
}

// CreateStandardTestContext creates a context with standard timeout (30 seconds).
func CreateStandardTestContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 30*time.Second)
}

// CreateQuickTestContext creates a context with quick timeout (1 microsecond for timeout tests).
func CreateQuickTestContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 1*time.Microsecond)
}

// CreateCancelledTestContext creates a cancelled context for testing cancellation scenarios.
func CreateCancelledTestContext() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // Cancel immediately
	return ctx, cancel
}
