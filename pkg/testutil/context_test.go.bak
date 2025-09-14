package testutil

import (
	"testing"
	"time"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil"
)

// TestTestUtilUnit_TestContext_Success tests TestContext function.
func TestTestUtilUnit_TestContext_Success(t *testing.T) {
	ctx := TestContext(t)
	testutil.AssertNotNil(t, ctx, "TestContext should return a non-nil context")

	// Check that context has a deadline
	deadline, ok := ctx.Deadline()
	testutil.AssertBoolEquals(t, ok, true, "Context should have a deadline")
	testutil.AssertBoolEquals(t, deadline.IsZero(), false, "Deadline should not be zero")
}

// TestTestUtilUnit_TestContextWithTimeout_Success tests TestContextWithTimeout function.
func TestTestUtilUnit_TestContextWithTimeout_Success(t *testing.T) {
	timeout := 5 * time.Second
	ctx := TestContextWithTimeout(t, timeout)
	testutil.AssertNotNil(t, ctx, "TestContextWithTimeout should return a non-nil context")

	// Check that context has a deadline
	deadline, ok := ctx.Deadline()
	testutil.AssertBoolEquals(t, ok, true, "Context should have a deadline")
	testutil.AssertBoolEquals(t, deadline.IsZero(), false, "Deadline should not be zero")
}
