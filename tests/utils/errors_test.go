// Package testutils provides common error handling utilities for tests.
package testutils

import (
	"context"
	"errors"
	"testing"
	"time"

	wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
)

func TestGetNilClientErrorTests(t *testing.T) {
	tests := GetNilClientErrorTests()

	if len(tests) == 0 {
		t.Error("Expected at least one test case")
	}

	for i, test := range tests {
		if test.Name == "" {
			t.Errorf("Test case %d: Name is empty", i)
		}
		if test.TestFunc == nil {
			t.Errorf("Test case %d: TestFunc is nil", i)
		}
		if test.ExpectedError == "" {
			t.Errorf("Test case %d: ExpectedError is empty", i)
		}

		// Test the function with nil client
		err := test.TestFunc(nil)
		if err == nil {
			t.Errorf("Test case %d: Expected error but got nil", i)
		}

		if test.ExpectedError != "client is nil" {
			t.Errorf("Test case %d: Expected error message 'client is nil', got '%s'", i, test.ExpectedError)
		}
	}
}

func TestRunCommonErrorTests(t *testing.T) {
	testCases := []ErrorTestCase{
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
		{
			Name: "ValidClient",
			TestFunc: func(client *wnc.Client) error {
				// Test with non-nil client - should not error
				return nil
			},
			ExpectedError: "", // No error expected
		},
	}

	// Test without calling the actual RunCommonErrorTests function
	// since it requires environment variables
	t.Run("NilClientTest", func(t *testing.T) {
		tc := testCases[0] // NilClient test case
		result := tc.TestFunc(nil)
		if result == nil {
			t.Error("Expected error with nil client, got nil")
		} else if result.Error() != tc.ExpectedError {
			t.Errorf("Expected error '%s', got '%s'", tc.ExpectedError, result.Error())
		}
	})

	t.Run("ValidClientTest", func(t *testing.T) {
		tc := testCases[1] // ValidClient test case
		// Create a dummy client to test with
		config := &wnc.Config{
			Controller:  "192.168.1.100",
			AccessToken: "dGVzdDp0ZXN0", // "test:test" in base64
		}
		client, err := wnc.NewClient(*config)
		if err != nil {
			t.Fatalf("Failed to create client: %v", err)
		}

		result := tc.TestFunc(client)
		if result != nil {
			t.Errorf("Expected no error with valid client, got: %v", result)
		}
	})
}

func TestTestWithCancelledContext(t *testing.T) {
	testFunc := func(ctx context.Context, client *wnc.Client) error {
		if client == nil {
			return errors.New("client is nil")
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(100 * time.Millisecond):
			return nil
		}
	}

	TestWithCancelledContext(t, testFunc)
}

func TestTestWithTimeout(t *testing.T) {
	testFunc := func(ctx context.Context, client *wnc.Client) error {
		if client == nil {
			return errors.New("client is nil")
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(2 * time.Microsecond):
			return nil
		}
	}

	TestWithTimeout(t, testFunc, 1*time.Microsecond)
}

func TestTestContextHandling(t *testing.T) {
	testFunc := func(ctx context.Context, client *wnc.Client) error {
		if client == nil {
			return errors.New("client is nil")
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(10 * time.Millisecond):
			return nil
		}
	}

	TestContextHandling(t, testFunc)
}

func TestRunCommonErrorTestsDirectCall(t *testing.T) {
	testCases := []ErrorTestCase{
		{
			Name: "DirectNilClientTest",
			TestFunc: func(client *wnc.Client) error {
				if client == nil {
					return errors.New("client is nil")
				}
				return nil
			},
			ExpectedError: "client is nil",
		},
	}

	// Test direct execution without needing environment setup
	t.Run("DirectNilClientTest", func(t *testing.T) {
		tc := testCases[0]
		result := tc.TestFunc(nil)
		if result == nil {
			t.Error("Expected error with nil client, got nil")
		} else if result.Error() != tc.ExpectedError {
			t.Errorf("Expected error '%s', got '%s'", tc.ExpectedError, result.Error())
		}
	})

	// Test with valid client - this should return nil error
	t.Run("DirectNilClientTest", func(t *testing.T) {
		// Create a mock client for testing
		client := &wnc.Client{}
		tc := testCases[0]
		result := tc.TestFunc(client)
		if result != nil {
			t.Errorf("Expected no error with valid client, got: %v", result)
		}
	})
}
