package tests

import (
	"context"
	"errors"
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
