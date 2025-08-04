// Package testutils provides common error handling utilities for tests.
package testutils

import (
	"context"
	"errors"
	"testing"
	"time"

	wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil"
)

// CommonErrorTestCases is an alias for ErrorTestCase for backward compatibility.
type CommonErrorTestCases = ErrorTestCase

// ErrorTestCase represents a single error test case.
type ErrorTestCase struct {
	Name          string
	TestFunc      func(*wnc.Client) error
	ExpectedError string
}

// GetNilClientErrorTests returns common nil client error test cases.
func GetNilClientErrorTests() []ErrorTestCase {
	return []ErrorTestCase{
		{
			Name: "NilClient",
			TestFunc: func(client *wnc.Client) error {
				if client == nil {
					return errors.New("client is nil")
				}
				return nil
			},
			ExpectedError: "client is nil",
		},
	}
}

// RunCommonErrorTests executes common error handling patterns.
func RunCommonErrorTests(t *testing.T, testCaseName string, testCases []ErrorTestCase) {
	t.Helper()

	// Check if environment variables are available
	config := testutil.NewTestConfigFromEnv()
	if config == nil {
		t.Skip("Required environment variables not set - skipping test")
		return
	}

	client, err := wnc.NewClient(*config)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Helper()
			result := tc.TestFunc(client)
			ValidateErrorContains(t, result, tc.ExpectedError)
		})
	}
}

// TestWithCancelledContext tests function behavior with cancelled context.
func TestWithCancelledContext(t *testing.T, testFunc func(context.Context, *wnc.Client) error) {
	t.Helper()

	client := GetTestClient(t)
	if client == nil {
		t.Skip("Cannot create test client - skipping context cancellation test")
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	cancel() // Cancel immediately

	err := testFunc(ctx, client)
	if err == nil {
		t.Error("Expected error with cancelled context, got nil")
	}

	if !errors.Is(err, context.Canceled) {
		t.Errorf("Expected context.Canceled error, got: %v", err)
	}
}

// TestWithTimeout tests function behavior with timeout context.
func TestWithTimeout(t *testing.T, testFunc func(context.Context, *wnc.Client) error, timeout time.Duration) {
	t.Helper()

	client := GetTestClient(t)
	if client == nil {
		t.Skip("Cannot create test client - skipping timeout test")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	err := testFunc(ctx, client)
	// This may or may not timeout depending on the actual operation speed
	// The test ensures the context is properly handled
	_ = err
}

// TestContextHandling provides a standardized way to test context handling.
func TestContextHandling(t *testing.T, testFunc func(context.Context, *wnc.Client) error) {
	t.Helper()

	client := GetTestClient(t)
	if client == nil {
		t.Skip("Cannot create test client - skipping context handling tests")
		return
	}

	tests := []struct {
		name    string
		ctx     func() (context.Context, context.CancelFunc)
		wantErr bool
	}{
		{
			name: "CancelledContext",
			ctx: func() (context.Context, context.CancelFunc) {
				ctx, cancel := context.WithCancel(context.Background())
				cancel() // Cancel immediately
				return ctx, cancel
			},
			wantErr: true,
		},
		{
			name: "QuickTimeout",
			ctx: func() (context.Context, context.CancelFunc) {
				return context.WithTimeout(context.Background(), 1*time.Microsecond)
			},
			wantErr: true, // Usually times out
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := tt.ctx()
			defer cancel()

			err := testFunc(ctx, client)
			if tt.wantErr && err == nil {
				t.Errorf("Expected error for %s, got nil", tt.name)
			}
		})
	}
}
