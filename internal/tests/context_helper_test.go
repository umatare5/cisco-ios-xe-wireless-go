package tests

import (
	"context"
	"testing"
	"time"
)

func TestTestContext(t *testing.T) {
	ctx := TestContext(t)
	if ctx == nil {
		t.Error("expected non-nil context")
	}
	if deadline, ok := ctx.Deadline(); !ok || deadline.Before(time.Now()) {
		t.Error("expected future deadline")
	}
}

func TestTestContextNilDefensive(t *testing.T) {
	if ctx := TestContext(t); ctx == nil {
		t.Fatalf("unexpected nil context")
	}
}

func TestTestContextDirect(t *testing.T) {
	if ctx := TestContext(t); ctx == nil {
		t.Fatal("expected non-nil context")
	}
}

func TestContextOperations(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	if deadline, ok := ctx.Deadline(); !ok || deadline.Before(time.Now()) {
		t.Error("deadline issue")
	}
	cancel()
	select {
	case <-ctx.Done():
	default:
		t.Error("expected canceled")
	}
}

func TestContextTimeoutSeparate(t *testing.T) {
	ctx := TestContextWithTimeout(t, 1*time.Millisecond)
	<-ctx.Done()
}
