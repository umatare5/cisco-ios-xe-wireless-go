package dot15_test

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/pkg/testutil"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/dot15"
)

// TestDot15ServiceUnit_Constructor_Success tests service constructor.
func TestDot15ServiceUnit_Constructor_Success(t *testing.T) {
	t.Run("NewServiceWithValidClient", func(t *testing.T) {
		mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(map[string]string{}))
		defer mockServer.Close()

		client := testutil.NewTestClient(mockServer)
		service := dot15.NewService(client.Core().(*core.Client))
		if service.Client() == nil {
			t.Error("Expected service to have a client")
		}
	})

	t.Run("NewServiceWithNilClient", func(t *testing.T) {
		service := dot15.NewService(nil)
		if service.Client() != nil {
			t.Error("Expected service client to be nil")
		}
	})
}

// TestDot15ServiceUnit_GetOperations_MockSuccess tests Get operations using mock server.
func TestDot15ServiceUnit_GetOperations_MockSuccess(t *testing.T) {
	mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(map[string]string{
		"Cisco-IOS-XE-wireless-dot15-cfg:dot15-cfg-data": `{
			"Cisco-IOS-XE-wireless-dot15-cfg:dot15-cfg-data": {
				"dot15-global-config": {
					"global-radio-shut": false
				}
			}
		}`,
		"Cisco-IOS-XE-wireless-dot15-cfg:dot15-cfg-data/dot15-global-config": `{
			"Cisco-IOS-XE-wireless-dot15-cfg:dot15-global-config": {
				"global-radio-shut": false
			}
		}`,
	}))
	defer mockServer.Close()

	client := testutil.NewTestClient(mockServer)
	service := dot15.NewService(client.Core().(*core.Client))
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

	t.Run("ListDot15GlobalConfigs", func(t *testing.T) {
		result, err := service.ListDot15GlobalConfigs(ctx)
		if err != nil {
			t.Errorf("Expected no error for ListDot15GlobalConfigs, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for ListDot15GlobalConfigs, got nil")
		}
	})
}

// TestDot15ServiceUnit_GetOperations_ErrorHandling tests error scenarios using mock server.
func TestDot15ServiceUnit_GetOperations_ErrorHandling(t *testing.T) {
	mockServer := testutil.NewMockServer(testutil.WithErrorResponses([]string{
		"Cisco-IOS-XE-wireless-dot15-cfg:dot15-cfg-data",
		"Cisco-IOS-XE-wireless-dot15-cfg:dot15-cfg-data/dot15-global-config",
	}, 404))
	defer mockServer.Close()

	client := testutil.NewTestClient(mockServer)
	service := dot15.NewService(client.Core().(*core.Client))
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

	t.Run("ListDot15GlobalConfigs_404Error", func(t *testing.T) {
		_, err := service.ListDot15GlobalConfigs(ctx)
		if err == nil {
			t.Error("Expected error for 404 response, got nil")
		}
		if !core.IsNotFoundError(err) {
			t.Errorf("Expected NotFound error, got: %v", err)
		}
	})
}

// TestDot15ServiceUnit_ErrorHandling_NilClient tests operations with nil client.
func TestDot15ServiceUnit_ErrorHandling_NilClient(t *testing.T) {
	service := dot15.NewService(nil)
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

	t.Run("ListDot15GlobalConfigs_NilClient", func(t *testing.T) {
		result, err := service.ListDot15GlobalConfigs(ctx)
		if err == nil {
			t.Error("Expected error for nil client, got nil")
		}
		if result != nil {
			t.Error("Expected nil result for nil client")
		}
	})
}
