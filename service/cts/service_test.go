package cts_test

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/pkg/testutil"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/cts"
)

// TestCtsServiceUnit_Constructor_Success tests service constructor.
func TestCtsServiceUnit_Constructor_Success(t *testing.T) {
	t.Run("NewServiceWithValidClient", func(t *testing.T) {
		mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(map[string]string{}))
		defer mockServer.Close()

		client := testutil.NewTestClient(mockServer)
		service := cts.NewService(client.Core().(*core.Client))
		if service.Client() == nil {
			t.Error("Expected service to have a client")
		}
	})

	t.Run("NewServiceWithNilClient", func(t *testing.T) {
		service := cts.NewService(nil)
		if service.Client() != nil {
			t.Error("Expected service client to be nil")
		}
	})
}

// TestCtsServiceUnit_GetOperations_MockSuccess tests Get operations using mock server.
func TestCtsServiceUnit_GetOperations_MockSuccess(t *testing.T) {
	mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(map[string]string{
		"Cisco-IOS-XE-wireless-cts-sxp-cfg:cts-sxp-cfg-data": `{
			"Cisco-IOS-XE-wireless-cts-sxp-cfg:cts-sxp-cfg-data": {
				"cts-sxp-configuration": {
					"cts-sxp-config": [
						{
							"sxp-profile-name": "test-profile",
							"enable": true
						}
					]
				}
			}
		}`,
		"Cisco-IOS-XE-wireless-cts-sxp-cfg:cts-sxp-cfg-data/cts-sxp-configuration/cts-sxp-config": `{
			"Cisco-IOS-XE-wireless-cts-sxp-cfg:cts-sxp-config": [
				{
					"sxp-profile-name": "test-profile",
					"enable": true
				}
			]
		}`,
		"Cisco-IOS-XE-wireless-cts-sxp-oper:cts-sxp-oper-data": `{
			"Cisco-IOS-XE-wireless-cts-sxp-oper:cts-sxp-oper-data": {
				"flex-mode-ap-sxp-connection-status": [
					{
						"wtp-mac": "00:11:22:33:44:55",
						"peer-ip": "192.168.1.1",
						"conn-mode": "listener"
					}
				]
			}
		}`,
		"Cisco-IOS-XE-wireless-cts-sxp-oper:cts-sxp-oper-data/flex-mode-ap-sxp-connection-status": `{
			"Cisco-IOS-XE-wireless-cts-sxp-oper:flex-mode-ap-sxp-connection-status": [
				{
					"wtp-mac": "00:11:22:33:44:55",
					"peer-ip": "192.168.1.1",
					"conn-mode": "listener"
				}
			]
		}`,
	}))
	defer mockServer.Close()

	client := testutil.NewTestClient(mockServer)
	service := cts.NewService(client.Core().(*core.Client))
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

	t.Run("GetOperational", func(t *testing.T) {
		result, err := service.GetOperational(ctx)
		if err != nil {
			t.Errorf("Expected no error for GetOperational, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for GetOperational, got nil")
		}
	})

	t.Run("ListFlexModeApSxpConnectionStatus", func(t *testing.T) {
		result, err := service.ListFlexModeApSxpConnectionStatus(ctx)
		if err != nil {
			t.Errorf("Expected no error for ListFlexModeApSxpConnectionStatus, got: %v", err)
		}
		if result == nil {
			t.Error("Expected result for ListFlexModeApSxpConnectionStatus, got nil")
		}
	})
}

// TestCtsServiceUnit_GetOperations_ErrorHandling tests error scenarios using mock server.
func TestCtsServiceUnit_GetOperations_ErrorHandling(t *testing.T) {
	mockServer := testutil.NewMockServer(testutil.WithErrorResponses([]string{
		"Cisco-IOS-XE-wireless-cts-sxp-cfg:cts-sxp-cfg-data",
		"Cisco-IOS-XE-wireless-cts-sxp-cfg:cts-sxp-cfg-data/cts-sxp-configuration/cts-sxp-config",
		"Cisco-IOS-XE-wireless-cts-sxp-oper:cts-sxp-oper-data",
		"Cisco-IOS-XE-wireless-cts-sxp-oper:cts-sxp-oper-data/flex-mode-ap-sxp-connection-status",
	}, 404))
	defer mockServer.Close()

	client := testutil.NewTestClient(mockServer)
	service := cts.NewService(client.Core().(*core.Client))
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

	t.Run("GetOperational_404Error", func(t *testing.T) {
		_, err := service.GetOperational(ctx)
		if err == nil {
			t.Error("Expected error for 404 response, got nil")
		}
		if !core.IsNotFoundError(err) {
			t.Errorf("Expected NotFound error, got: %v", err)
		}
	})

	t.Run("ListFlexModeApSxpConnectionStatus_404Error", func(t *testing.T) {
		_, err := service.ListFlexModeApSxpConnectionStatus(ctx)
		if err == nil {
			t.Error("Expected error for 404 response, got nil")
		}
		if !core.IsNotFoundError(err) {
			t.Errorf("Expected NotFound error, got: %v", err)
		}
	})
}

// TestCtsServiceUnit_ErrorHandling_NilClient tests operations with nil client.
func TestCtsServiceUnit_ErrorHandling_NilClient(t *testing.T) {
	service := cts.NewService(nil)
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

	t.Run("GetOperational_NilClient", func(t *testing.T) {
		result, err := service.GetOperational(ctx)
		if err == nil {
			t.Error("Expected error for nil client, got nil")
		}
		if result != nil {
			t.Error("Expected nil result for nil client")
		}
	})

	t.Run("ListFlexModeApSxpConnectionStatus_NilClient", func(t *testing.T) {
		result, err := service.ListFlexModeApSxpConnectionStatus(ctx)
		if err == nil {
			t.Error("Expected error for nil client, got nil")
		}
		if result != nil {
			t.Error("Expected nil result for nil client")
		}
	})
}
