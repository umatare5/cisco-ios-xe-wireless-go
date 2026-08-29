package core_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil"
	mock "github.com/umatare5/cisco-ios-xe-wireless-go/pkg/testutil"
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

	t.Run("GetRaw with nil client", func(t *testing.T) {
		_, err := core.GetRaw(ctx, nil, "/test")
		testutil.AssertError(t, err, "GetRaw() with nil client should return error")
		testutil.AssertErrorContains(t, err, "client cannot be nil", "GetRaw() nil client message")
	})
}

// TestGetRawReturnsBodyUnchecked pins the seam GetData sits on: the body arrives as the
// controller sent it, with no envelope check and no decode.
//
// The body carries two top-level keys, which is what makes the assertion mean something:
// the same read through Get fails, so a byte-identical answer from GetRaw is the absence
// of the guard rather than a body the guard would have accepted anyway.
func TestGetRawReturnsBodyUnchecked(t *testing.T) {
	const body = `{"a:one":{"x":1},"a:two":{"y":2}}`

	mockServer := mock.NewMockServer(mock.WithSuccessResponses(map[string]string{"probe": body}))
	defer mockServer.Close()

	client, ok := mock.NewTestClient(mockServer).Core().(*core.Client)
	testutil.AssertTrue(t, ok, "test client should carry a *core.Client")

	t.Run("Get rejects the body", func(t *testing.T) {
		type probe struct{}
		_, err := core.Get[probe](context.Background(), client, "probe")
		testutil.AssertError(t, err, "Get() should reject a body carrying two top-level keys")
		testutil.AssertErrorContains(t, err, "top-level keys", "Get() envelope error")
	})

	t.Run("GetRaw returns the same body", func(t *testing.T) {
		got, err := core.GetRaw(context.Background(), client, "probe")
		testutil.AssertNoError(t, err, "GetRaw() should not check the envelope")
		testutil.AssertStringEquals(t, string(got), body, "GetRaw() body")
	})
}

// TestGetRawEmptyBodyIsNonNilAndEmpty pins the empty-2xx contract GetRaw documents: an
// answer with no body is a successful read, reported as a non-nil slice of length zero.
func TestGetRawEmptyBodyIsNonNilAndEmpty(t *testing.T) {
	mockServer := mock.NewMockServer(
		mock.WithTesting(t),
		mock.WithCustomResponse("probe", mock.ResponseConfig{StatusCode: 204, Body: ""}),
	)
	defer mockServer.Close()

	client, ok := mock.NewTestClient(mockServer).Core().(*core.Client)
	testutil.AssertTrue(t, ok, "test client should carry a *core.Client")

	got, err := core.GetRaw(context.Background(), client, "probe")
	testutil.AssertNoError(t, err, "GetRaw() with an empty body should succeed")
	testutil.AssertTrue(t, got != nil, "GetRaw() with an empty body should not return nil")
	testutil.AssertIntEquals(t, len(got), 0, "GetRaw() empty body length")
}

// TestGetRawAppliesGetOptions pins that GetRaw folds options the way Get does, observed on
// the wire rather than on the return value of applyGetOptions.
func TestGetRawAppliesGetOptions(t *testing.T) {
	server := mock.NewRESTCONFServer(t)
	defer server.Close()
	server.AddHandler(http.MethodGet, "probe", func() (int, string) { return http.StatusOK, `{}` })

	client, ok := mock.NewTestClient(mock.NewMockServerFromHTTP(server.Server)).Core().(*core.Client)
	testutil.AssertTrue(t, ok, "test client should carry a *core.Client")

	testCases := []struct {
		name     string
		opts     []core.GetOption
		expected string
	}{
		{name: "no option", opts: nil, expected: ""},
		{
			name:     "fields only",
			opts:     []core.GetOption{core.WithFields("probe(one;two)")},
			expected: "fields=probe(one;two)",
		},
		{name: "depth only", opts: []core.GetOption{core.WithDepth(3)}, expected: "depth=3"},
		{
			name: "all three fold in a fixed order",
			opts: []core.GetOption{
				core.WithDepth(3),
				core.WithFields("probe(one;two)"),
				core.WithDefaults(core.DefaultsReportAll),
			},
			expected: "with-defaults=report-all&fields=probe(one;two)&depth=3",
		},
	}

	for i, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			_, err := core.GetRaw(context.Background(), client, "probe", tt.opts...)
			testutil.AssertNoError(t, err, "GetRaw() should succeed")

			recorded := server.Requests()
			testutil.AssertIntEquals(t, len(recorded), i+1, "recorded request count")
			testutil.AssertStringEquals(t, recorded[i].RawQuery, tt.expected, "wire query")
			testutil.AssertStringEquals(t, recorded[i].Path, "/restconf/data/probe", "wire path")
		})
	}
}

// TestRequestFunctionsJSONUnmarshalError tests JSON unmarshal error paths.
func TestRequestFunctionsJSONUnmarshalError(t *testing.T) {
	ctx := context.Background()

	// Create a mock server that returns invalid JSON
	mockServer := mock.NewMockServer(mock.WithSuccessResponses(map[string]string{
		"test": `{"invalid": json`, // Invalid JSON that will cause unmarshal error
	}))
	defer mockServer.Close()

	testClient := mock.NewTestClient(mockServer)
	client := testClient.Core().(*core.Client)

	// Test JSON unmarshal error for Get
	t.Run("Get with JSON unmarshal error", func(t *testing.T) {
		type TestResponse struct{}
		_, err := core.Get[TestResponse](ctx, client, "/test")
		testutil.AssertError(t, err, "Get() with invalid JSON should return error")
		testutil.AssertStringContains(
			t, err.Error(), "failed to unmarshal response",
			"error should contain unmarshal message",
		)
	})

	// Test JSON unmarshal error for Post
	t.Run("Post with JSON unmarshal error", func(t *testing.T) {
		type TestResponse struct{}
		_, err := core.Post[TestResponse](ctx, client, "/test", map[string]string{"key": "value"})
		testutil.AssertError(t, err, "Post() with invalid JSON should return error")
		testutil.AssertStringContains(
			t, err.Error(), "failed to unmarshal response",
			"error should contain unmarshal message",
		)
	})

	// Test JSON unmarshal error for Put
	t.Run("Put with JSON unmarshal error", func(t *testing.T) {
		type TestResponse struct{}
		_, err := core.Put[TestResponse](ctx, client, "/test", map[string]string{"key": "value"})
		testutil.AssertError(t, err, "Put() with invalid JSON should return error")
		testutil.AssertStringContains(
			t, err.Error(), "failed to unmarshal response",
			"error should contain unmarshal message",
		)
	})

	// Test JSON unmarshal error for Patch
	t.Run("Patch with JSON unmarshal error", func(t *testing.T) {
		type TestResponse struct{}
		_, err := core.Patch[TestResponse](ctx, client, "/test", map[string]string{"key": "value"})
		testutil.AssertError(t, err, "Patch() with invalid JSON should return error")
		testutil.AssertStringContains(
			t, err.Error(), "failed to unmarshal response",
			"error should contain unmarshal message",
		)
	})
}

// TestRequestFunctionsEmptyBodyIsZeroValue tests that a read answered with no body succeeds at
// the zero value, which is how the controller answers a node it holds no data for.
func TestRequestFunctionsEmptyBodyIsZeroValue(t *testing.T) {
	ctx := context.Background()

	mockServer := mock.NewMockServer(
		mock.WithTesting(t),
		mock.WithCustomResponse("test", mock.ResponseConfig{StatusCode: 204, Body: ""}),
		mock.WithCustomResponse("post", mock.ResponseConfig{StatusCode: 204, Body: "", Method: "POST"}),
	)
	defer mockServer.Close()

	testClient := mock.NewTestClient(mockServer)
	client := testClient.Core().(*core.Client)

	t.Run("Get with empty body", func(t *testing.T) {
		type TestResponse struct{}
		result, err := core.Get[TestResponse](ctx, client, "/test")
		testutil.AssertNoError(t, err, "Get() with an empty body should succeed")
		testutil.AssertTrue(t, result != nil, "Get() with an empty body should return a zero result")
	})

	t.Run("Post with empty body", func(t *testing.T) {
		type TestResponse struct{}
		result, err := core.Post[TestResponse](ctx, client, "/post", map[string]string{"key": "value"})
		testutil.AssertNoError(t, err, "Post() with an empty body should succeed")
		testutil.AssertTrue(t, result != nil, "Post() with an empty body should return a zero result")
	})
}

// TestRequestSeamsCarryTheStatus holds the four untyped seams to what each returns: the three that
// keep handing back bytes, and RequestRaw, which hands back the status beside them.
//
// The three statuses are served on three methods because that is the only way to tell them apart:
// a create, a replace answered with nothing and a read of a node holding nothing all reach the
// caller with a zero-length body.
func TestRequestSeamsCarryTheStatus(t *testing.T) {
	const (
		patched = `{"a:one":{}}`
		output  = `{"a:output":{}}`
	)

	ctx := context.Background()

	server := mock.NewRESTCONFServer(t)
	defer server.Close()
	server.AddHandler(http.MethodPatch, "probe", func() (int, string) { return http.StatusOK, patched })
	server.AddHandler(http.MethodPut, "probe", func() (int, string) { return http.StatusNoContent, "" })
	server.AddHandler(http.MethodPost, "x:rpc", func() (int, string) { return http.StatusCreated, output })

	client, ok := mock.NewTestClient(mock.NewMockServerFromHTTP(server.Server)).Core().(*core.Client)
	testutil.AssertTrue(t, ok, "test client should carry a *core.Client")

	t.Run("EditRaw returns the body", func(t *testing.T) {
		got, err := core.EditRaw(ctx, client, http.MethodPatch, "probe", nil)
		testutil.AssertNoError(t, err, "EditRaw() should succeed")
		testutil.AssertStringEquals(t, string(got), patched, "EditRaw() body")
	})

	t.Run("CallRPCRaw returns the body", func(t *testing.T) {
		got, err := core.CallRPCRaw(ctx, client, "Cisco-IOS-XE-x:rpc", nil)
		testutil.AssertNoError(t, err, "CallRPCRaw() should succeed")
		testutil.AssertStringEquals(t, string(got), output, "CallRPCRaw() body")
	})

	t.Run("RequestRaw reports the status on the data root", func(t *testing.T) {
		resp, err := core.RequestRaw(ctx, client, http.MethodPut, "probe", nil)
		testutil.AssertNoError(t, err, "RequestRaw() should succeed")
		testutil.AssertIntEquals(t, resp.StatusCode, http.StatusNoContent, "status")
		testutil.AssertIntEquals(t, len(resp.Body), 0, "body length")
	})

	t.Run("RequestRaw reports the status on the operations root", func(t *testing.T) {
		resp, err := core.RequestRaw(ctx, client, http.MethodPost, "/restconf/operations/Cisco-IOS-XE-x:rpc", nil)
		testutil.AssertNoError(t, err, "RequestRaw() should succeed")
		testutil.AssertIntEquals(t, resp.StatusCode, http.StatusCreated, "status")
		testutil.AssertStringEquals(t, string(resp.Body), output, "RequestRaw() body")
	})

	t.Run("every pre-send fault is refused", func(t *testing.T) {
		_, err := core.RequestRaw(ctx, nil, http.MethodPut, "probe", nil)
		testutil.AssertError(t, err, "RequestRaw() with a nil client should be refused")

		// The payload is prepared before the path is routed, so an empty method on an operations
		// path reports the method rather than the POST-only rule.
		_, err = core.RequestRaw(ctx, client, "", "/restconf/operations/Cisco-IOS-XE-x:rpc", nil)
		testutil.AssertErrorContains(t, err, "HTTP method cannot be empty", "RequestRaw() empty method")

		_, err = core.RequestRaw(ctx, client, http.MethodPatch, "probe", []byte("not json"))
		testutil.AssertError(t, err, "RequestRaw() with non-JSON payload bytes should be refused")

		_, err = core.RequestRaw(ctx, client, http.MethodPut, "/restconf/operations/Cisco-IOS-XE-x:rpc", nil)
		testutil.AssertErrorContains(t, err, "only POST", "RequestRaw() with PUT on an operations path")
	})

	t.Run("every seam reports a request failure", func(t *testing.T) {
		gone := mock.NewRESTCONFServer(t)
		closed, ok := mock.NewTestClient(mock.NewMockServerFromHTTP(gone.Server)).Core().(*core.Client)
		testutil.AssertTrue(t, ok, "test client should carry a *core.Client")
		gone.Close()

		_, err := core.GetRaw(ctx, closed, "probe")
		testutil.AssertError(t, err, "GetRaw() against a closed server should fail")

		_, err = core.EditRaw(ctx, closed, http.MethodPatch, "probe", nil)
		testutil.AssertError(t, err, "EditRaw() against a closed server should fail")

		_, err = core.CallRPCRaw(ctx, closed, "Cisco-IOS-XE-x:rpc", nil)
		testutil.AssertError(t, err, "CallRPCRaw() against a closed server should fail")

		resp, err := core.RequestRaw(ctx, closed, http.MethodPut, "probe", nil)
		testutil.AssertError(t, err, "RequestRaw() against a closed server should fail")
		testutil.AssertTrue(t, resp == nil, "RequestRaw() should return no Response beside its error")
	})
}
