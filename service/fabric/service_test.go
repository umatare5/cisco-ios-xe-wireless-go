package fabric_test

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/pkg/testutil"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/fabric"
)

// TestFabricServiceUnit_Constructor_Success tests service constructor.
func TestFabricServiceUnit_Constructor_Success(t *testing.T) {
	t.Run("NewServiceWithValidClient", func(t *testing.T) {
		mockServer := testutil.NewMockServer(map[string]string{})
		defer mockServer.Close()

		client := testutil.NewTestClient(mockServer)
		service := fabric.NewService(client.Core().(*core.Client))
		if service.Client() == nil {
			t.Error("Expected service to have a client")
		}
	})

	t.Run("NewServiceWithNilClient", func(t *testing.T) {
		service := fabric.NewService(nil)
		if service.Client() != nil {
			t.Error("Expected service client to be nil")
		}
	})
}

// TestFabricServiceUnit_GetOperations_MockSuccess tests Get operations using mock server.
func TestFabricServiceUnit_GetOperations_MockSuccess(t *testing.T) {
	mockServer := testutil.NewMockServer(map[string]string{
		"Cisco-IOS-XE-wireless-fabric-cfg:fabric-cfg-data": `{
			"Cisco-IOS-XE-wireless-fabric-cfg:fabric-cfg-data": {
				"config": {
					"enable": true
				}
			}
		}`,
	})
	defer mockServer.Close()

	client := testutil.NewTestClient(mockServer)
	service := fabric.NewService(client.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	_, err := service.GetConfig(ctx)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
}

// TestFabricServiceUnit_GetOperations_ErrorHandling tests error scenarios using mock server.
func TestFabricServiceUnit_GetOperations_ErrorHandling(t *testing.T) {
	mockServer := testutil.NewMockErrorServer([]string{
		"Cisco-IOS-XE-wireless-fabric-cfg:fabric-cfg-data",
	}, 404)
	defer mockServer.Close()

	client := testutil.NewTestClient(mockServer)
	service := fabric.NewService(client.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	_, err := service.GetConfig(ctx)
	if err == nil {
		t.Error("Expected error for 404 response, got nil")
	}
}
