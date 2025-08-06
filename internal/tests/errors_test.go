package tests

import (
	"context"
	"errors"
	"fmt"
	"os"
	"testing"
	"time"

	wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
)

func TestRunCommonErrorTestsFunction(t *testing.T) {
	testCases := []ErrorTestCase{
		{
			Name: "TestError",
			TestFunc: func(client *wnc.Client) error {
				return errors.New("test error")
			},
			ExpectedError: "test error",
		},
	}

	RunCommonErrorTests(t, "CommonErrorTests", testCases)
}

func TestWithCancelledContextFunction(t *testing.T) {
	testFunc := func(ctx context.Context, client *wnc.Client) error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			return nil
		}
	}

	TestWithCancelledContext(t, testFunc)
}

func TestWithTimeoutFunction(t *testing.T) {
	testFunc := func(ctx context.Context, client *wnc.Client) error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(10 * time.Millisecond):
			return nil
		}
	}

	TestWithTimeout(t, testFunc, 1*time.Microsecond)
}

func TestContextHandlingFunction(t *testing.T) {
	testFunc := func(ctx context.Context, client *wnc.Client) error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(10 * time.Millisecond):
			return nil
		}
	}

	TestContextHandling(t, testFunc)
}

// Additional tests to improve coverage

func TestGetNilClientErrorTestsExtended(t *testing.T) {
	tests := GetNilClientErrorTests()

	// Test that all test cases handle nil client correctly
	for i, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			if test.TestFunc == nil {
				t.Errorf("Test case %d: TestFunc is nil", i)
				return
			}

			// Test with nil client
			err := test.TestFunc(nil)
			if err == nil {
				t.Errorf("Test case %d: Expected error but got nil", i)
				return
			}

			// Verify error message contains expected content
			if test.ExpectedError != "" && err.Error() != test.ExpectedError {
				t.Errorf("Test case %d: Expected error '%s', got '%s'", i, test.ExpectedError, err.Error())
			}
		})
	}
}

// TestErrorTestCaseAdditionalCoverage tests additional error test case patterns
func TestErrorTestCaseAdditionalCoverage(t *testing.T) {
	t.Run("GetNilClientErrorTests_Coverage", func(t *testing.T) {
		testCases := GetNilClientErrorTests()
		if len(testCases) == 0 {
			t.Error("Expected non-empty test cases")
		}

		// Test the actual function logic
		for _, tc := range testCases {
			t.Run(tc.Name, func(t *testing.T) {
				// Test with nil client
				err := tc.TestFunc(nil)
				if err == nil {
					t.Error("Expected error with nil client")
				}
				if err.Error() != tc.ExpectedError {
					t.Errorf("Expected error %q, got %q", tc.ExpectedError, err.Error())
				}

				// Test with valid client (if environment allows)
				if os.Getenv("WNC_CONTROLLER") != "" && os.Getenv("WNC_ACCESS_TOKEN") != "" {
					client := GetTestClient(t)
					err = tc.TestFunc(client)
					if err != nil {
						t.Logf("Got error with valid client (acceptable): %v", err)
					}
				}
			})
		}
	})

	t.Run("RunCommonErrorTests_Coverage", func(t *testing.T) {
		testCases := []ErrorTestCase{
			{
				Name: "TestCase1",
				TestFunc: func(client *wnc.Client) error {
					if client == nil {
						return fmt.Errorf("test error")
					}
					// Return a different error for non-nil client that still validates our error handling
					return fmt.Errorf("test error with valid client")
				},
				ExpectedError: "test error",
			},
		}

		RunCommonErrorTests(t, "TestCoverage", testCases)
	})

	t.Run("TestWithCancelledContext_Coverage", func(t *testing.T) {
		TestWithCancelledContext(t, func(ctx context.Context, c *wnc.Client) error {
			// Simulate checking context cancellation
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
				return fmt.Errorf("context not cancelled")
			}
		})
	})

	t.Run("TestWithTimeout_Coverage", func(t *testing.T) {
		TestWithTimeout(t, func(ctx context.Context, c *wnc.Client) error {
			// Simple function that should complete within timeout
			return nil
		}, 100*time.Millisecond)
	})

	t.Run("TestContextHandling_Coverage", func(t *testing.T) {
		TestContextHandling(t, func(ctx context.Context, c *wnc.Client) error {
			// Function that simulates work and checks for context cancellation
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(50 * time.Millisecond):
				return nil
			}
		})
	})
}

func TestRunCommonErrorTestsWithMultipleCases(t *testing.T) {
	testCases := []ErrorTestCase{
		{
			Name: "ConnectionError",
			TestFunc: func(client *wnc.Client) error {
				return errors.New("connection refused")
			},
			ExpectedError: "connection refused",
		},
		{
			Name: "TimeoutError",
			TestFunc: func(client *wnc.Client) error {
				return errors.New("timeout occurred")
			},
			ExpectedError: "timeout occurred",
		},
		{
			Name: "NotFoundError",
			TestFunc: func(client *wnc.Client) error {
				return errors.New("404 Not Found")
			},
			ExpectedError: "404 Not Found",
		},
	}

	RunCommonErrorTests(t, "MultipleErrorCases", testCases)
}

func TestWithCancelledContextEdgeCases(t *testing.T) {
	// Test with function that ignores context
	t.Run("IgnoresCancelledContext", func(t *testing.T) {
		testFunc := func(ctx context.Context, client *wnc.Client) error {
			// Function that doesn't check context cancellation
			// but still might detect cancellation in some cases
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(1 * time.Millisecond):
				return nil
			}
		}

		// This test exercises the code path but may not always fail as expected
		// since context cancellation timing is non-deterministic
		TestWithCancelledContext(t, testFunc)
	})

	// Test with function that returns immediate error
	t.Run("ImmediateError", func(t *testing.T) {
		testFunc := func(ctx context.Context, client *wnc.Client) error {
			// Check if context is already cancelled first
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
				return errors.New("immediate error")
			}
		}

		// This will exercise the error handling path
		TestWithCancelledContext(t, testFunc)
	})
}

func TestWithTimeoutEdgeCases(t *testing.T) {
	// Test with very short timeout
	t.Run("VeryShortTimeout", func(t *testing.T) {
		testFunc := func(ctx context.Context, client *wnc.Client) error {
			time.Sleep(100 * time.Millisecond) // Longer than timeout
			return nil
		}

		TestWithTimeout(t, testFunc, 1*time.Nanosecond)
	})

	// Test with function that finishes before timeout
	t.Run("FastFunction", func(t *testing.T) {
		testFunc := func(ctx context.Context, client *wnc.Client) error {
			return nil // Immediate return
		}

		TestWithTimeout(t, testFunc, 1*time.Second)
	})

	// Test with function that returns error
	t.Run("FunctionWithError", func(t *testing.T) {
		testFunc := func(ctx context.Context, client *wnc.Client) error {
			return errors.New("function error")
		}

		TestWithTimeout(t, testFunc, 1*time.Second)
	})
}

func TestContextHandlingEdgeCases(t *testing.T) {
	// Test with function that returns context error
	t.Run("ContextError", func(t *testing.T) {
		testFunc := func(ctx context.Context, client *wnc.Client) error {
			<-ctx.Done()
			return ctx.Err()
		}

		TestContextHandling(t, testFunc)
	})

	// Test with function that times out
	t.Run("TimeoutFunction", func(t *testing.T) {
		testFunc := func(ctx context.Context, client *wnc.Client) error {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(200 * time.Millisecond):
				return errors.New("operation completed")
			}
		}

		TestContextHandling(t, testFunc)
	})

	// Test with function that ignores context entirely
	t.Run("IgnoresContext", func(t *testing.T) {
		testFunc := func(ctx context.Context, client *wnc.Client) error {
			// Function that doesn't use context at all
			return errors.New("context ignored")
		}

		TestContextHandling(t, testFunc)
	})
}
