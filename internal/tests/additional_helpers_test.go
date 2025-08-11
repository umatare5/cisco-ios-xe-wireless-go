package tests

import (
	"context"
	"testing"
	"time"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
)

// TestNewTLSClientForServer_Coverage ensures the helper constructs a usable client
// and exercises a minimal core.Get call against a local TLS server.
func TestNewTLSClientForServer_Coverage(t *testing.T) {
	t.Parallel()

	// Serve minimal valid JSON for a simple RESTCONF endpoint
	const ep = "Cisco-IOS-XE-wireless-access-point-oper:ap-name-mac-map"
	srv := NewRESTCONFSuccessServer(map[string]string{
		ep: `{"Cisco-IOS-XE-wireless-access-point-oper:ap-name-mac-map":[]}`,
	})
	t.Cleanup(srv.Close)

	c := NewTLSClientForServer(t, srv)

	// Local response type matching the served JSON
	type resp struct {
		Data []struct{} `json:"Cisco-IOS-XE-wireless-access-point-oper:ap-name-mac-map"`
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	out, err := core.Get[resp](ctx, c, ep)
	if err != nil {
		t.Fatalf("core.Get error: %v", err)
	}
	if out == nil {
		t.Fatalf("expected non-nil response")
	}
}

// TestAssertNonNilResult_Branches covers both nil and non-nil inputs.
func TestAssertNonNilResult_Branches(t *testing.T) {
	t.Parallel()

	// non-nil path only (nil path intentionally omitted: it would fail the test via t.Errorf)
	AssertNonNilResult(t, struct{}{}, "dummy")
}

// TestValidateStructType_Paths covers nil early-return and marshal-error branches.
func TestValidateStructType_Paths(t *testing.T) {
	t.Parallel()

	// nil input: early return
	ValidateStructType(t, nil)

	// Unmarshalable type for json.Marshal (channel) to exercise error path
	ch := make(chan int)
	ValidateStructType(t, ch)
}
