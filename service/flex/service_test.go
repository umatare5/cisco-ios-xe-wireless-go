package flex_test

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/pkg/testutil"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/flex"
)

// TestFlexServiceUnit_Constructor_Success tests service constructor.
func TestFlexServiceUnit_Constructor_Success(t *testing.T) {
	t.Run("NewServiceWithValidClient", func(t *testing.T) {
		mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(map[string]string{}))
		defer mockServer.Close()

		client := testutil.NewTestClient(mockServer)
		service := flex.NewService(client.Core().(*core.Client))
		if service.Client() == nil {
			t.Error("Expected service to have a client")
		}
	})

	t.Run("NewServiceWithNilClient", func(t *testing.T) {
		service := flex.NewService(nil)
		if service.Client() != nil {
			t.Error("Expected service client to be nil")
		}
	})
}

// TestFlexServiceUnit_GetOperations_MockSuccess tests Get operations using mock server.
func TestFlexServiceUnit_GetOperations_MockSuccess(t *testing.T) {
	mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(map[string]string{
		"Cisco-IOS-XE-wireless-flex-cfg:flex-cfg-data": `{
			"Cisco-IOS-XE-wireless-flex-cfg:flex-cfg-data": {
				"flex-policy-entries": {
					"flex-policy-entry": [
						{
							"policy-name": "test-policy",
							"description": "Test FlexConnect policy"
						}
					]
				}
			}
		}`,
		"Cisco-IOS-XE-wireless-flex-cfg:flex-cfg-data/flex-policy-entries": `{
			"Cisco-IOS-XE-wireless-flex-cfg:flex-policy-entries": {
				"flex-policy-entry": [
					{
						"policy-name": "test-policy",
						"description": "Test FlexConnect policy"
					}
				]
			}
		}`,
	}))
	defer mockServer.Close()

	client := testutil.NewTestClient(mockServer)
	service := flex.NewService(client.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	t.Run("GetConfig", func(t *testing.T) {
		result, err := service.GetConfig(ctx)
		if err != nil {
			t.Errorf("Expected no error for GetConfig, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for GetConfig, got nil")
		}
	})

	t.Run("ListFlexPolicyEntries", func(t *testing.T) {
		result, err := service.ListFlexPolicyEntries(ctx)
		if err != nil {
			t.Errorf("Expected no error for ListFlexPolicyEntries, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for ListFlexPolicyEntries, got nil")
		}
	})
}

// TestFlexServiceUnit_GetOperations_ErrorHandling tests error scenarios using mock server.
func TestFlexServiceUnit_GetOperations_ErrorHandling(t *testing.T) {
	mockServer := testutil.NewMockServer(testutil.WithErrorResponses([]string{
		"Cisco-IOS-XE-wireless-flex-cfg:flex-cfg-data",
		"Cisco-IOS-XE-wireless-flex-cfg:flex-cfg-data/flex-policy-entries",
	}, 404))
	defer mockServer.Close()

	client := testutil.NewTestClient(mockServer)
	service := flex.NewService(client.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	t.Run("GetConfig_404Error", func(t *testing.T) {
		_, err := service.GetConfig(ctx)
		if err == nil {
			t.Error("Expected error for 404 response, got nil")
		}
		if !core.IsNotFoundError(err) {
			t.Errorf("Expected NotFound error, got: %v", err)
		}
	})

	t.Run("ListFlexPolicyEntries_404Error", func(t *testing.T) {
		_, err := service.ListFlexPolicyEntries(ctx)
		if err == nil {
			t.Error("Expected error for 404 response, got nil")
		}
		if !core.IsNotFoundError(err) {
			t.Errorf("Expected NotFound error, got: %v", err)
		}
	})
}

// TestFlexServiceUnit_ErrorHandling_NilClient tests operations with nil client.
func TestFlexServiceUnit_ErrorHandling_NilClient(t *testing.T) {
	service := flex.NewService(nil)
	ctx := testutil.TestContext(t)

	t.Run("GetConfig_NilClient", func(t *testing.T) {
		result, err := service.GetConfig(ctx)
		if err == nil {
			t.Error("Expected error for nil client, got nil")
		}
		if result != nil {
			t.Error("Expected nil result for nil client")
		}
	})

	t.Run("ListFlexPolicyEntries_NilClient", func(t *testing.T) {
		result, err := service.ListFlexPolicyEntries(ctx)
		if err == nil {
			t.Error("Expected error for nil client, got nil")
		}
		if result != nil {
			t.Error("Expected nil result for nil client")
		}
	})
}
