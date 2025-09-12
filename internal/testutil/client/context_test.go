package client

import (
	"context"
	"testing"
	"time"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/helper"
)

func TestTestUtilClientUnit_ContextHelper_Success(t *testing.T) {
	ctx := TestContext(t)
	if ctx == nil {
		helper.AssertNotNil(t, ctx, "expected non-nil context")
	}
	if deadline, ok := ctx.Deadline(); !ok || deadline.Before(time.Now()) {
		helper.AssertTrue(t, ok && deadline.After(time.Now()), "expected future deadline")
	}
}

func TestTestUtilClientUnit_ContextHelper_NilDefensive(t *testing.T) {
	if ctx := TestContext(t); ctx == nil {
		helper.AssertNotNil(t, ctx, "unexpected nil context")
		return
	}
}

func TestTestUtilClientUnit_ContextDirect_Success(t *testing.T) {
	ctx := TestContext(t)
	if ctx == nil {
		t.Errorf("Context is nil")
	}
}

func TestTestUtilClientUnit_ContextOperations_Success(t *testing.T) {
	ctx, cancel := context.WithTimeout(TestContext(t), time.Minute) // Modern Go style - no explicit multiplication by 1
	defer cancel()

	if ctx == nil {
		t.Errorf("Context is nil")
	}

	select {
	case <-ctx.Done():
		t.Errorf("Context is done")
	default:
		// Context is not done
	}
}

func TestTestUtilClientUnit_ContextTimeoutSeparate_Success(t *testing.T) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 100*time.Millisecond)
	defer cancel()

	// Sleep for longer than timeout
	time.Sleep(200 * time.Millisecond)

	// Check if context is done
	select {
	case <-ctx.Done():
		// Context is done, which is expected
	default:
		t.Errorf("Context should be done due to timeout")
	}
}
