package core_test

import (
	"context"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil"
)

// TestRequestFunctions tests the request helper functions with nil client validation.
func TestRequestFunctions(t *testing.T) {
	ctx := context.Background()

	// Test all functions with nil client
	t.Run("Get with nil client", func(t *testing.T) {
		type TestResponse struct{}
		_, err := core.Get[TestResponse](ctx, nil, "/test")
		testutil.AssertError(t, err, "Get() with nil client should return error")
	})

	t.Run("Post with nil client", func(t *testing.T) {
		type TestResponse struct{}
		_, err := core.Post[TestResponse](ctx, nil, "/test", nil)
		testutil.AssertError(t, err, "Post() with nil client should return error")
	})

	t.Run("PostVoid with nil client", func(t *testing.T) {
		err := core.PostVoid(ctx, nil, "/test", nil)
		testutil.AssertError(t, err, "PostVoid() with nil client should return error")
	})

	t.Run("PostRPCVoid with nil client", func(t *testing.T) {
		err := core.PostRPCVoid(ctx, nil, "/test", nil)
		testutil.AssertError(t, err, "PostRPCVoid() with nil client should return error")
	})

	t.Run("Put with nil client", func(t *testing.T) {
		type TestResponse struct{}
		_, err := core.Put[TestResponse](ctx, nil, "/test", nil)
		testutil.AssertError(t, err, "Put() with nil client should return error")
	})

	t.Run("PutVoid with nil client", func(t *testing.T) {
		err := core.PutVoid(ctx, nil, "/test", nil)
		testutil.AssertError(t, err, "PutVoid() with nil client should return error")
	})

	t.Run("Patch with nil client", func(t *testing.T) {
		type TestResponse struct{}
		_, err := core.Patch[TestResponse](ctx, nil, "/test", nil)
		testutil.AssertError(t, err, "Patch() with nil client should return error")
	})

	t.Run("PatchVoid with nil client", func(t *testing.T) {
		err := core.PatchVoid(ctx, nil, "/test", nil)
		testutil.AssertError(t, err, "PatchVoid() with nil client should return error")
	})

	t.Run("Delete with nil client", func(t *testing.T) {
		err := core.Delete(ctx, nil, "/test")
		testutil.AssertError(t, err, "Delete() with nil client should return error")
	})
}
