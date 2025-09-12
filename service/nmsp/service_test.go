package nmsp_test

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/pkg/testutil"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/nmsp"
)

func TestNmspServiceUnit_Constructor_Success(t *testing.T) {
	t.Run("NewServiceWithValidClient", func(t *testing.T) {
		server := testutil.NewMockServer(map[string]string{})
		defer server.Close()
		testClient := testutil.NewTestClient(server)
		service := nmsp.NewService(testClient.Core().(*core.Client))
		if service.Client() == nil {
			t.Error("Expected valid client, got nil")
		}
	})
	t.Run("NewServiceWithNilClient", func(t *testing.T) {
		service := nmsp.NewService(nil)
		if service.Client() != nil {
			t.Error("Expected nil client, got non-nil")
		}
	})
}

func TestNmspServiceUnit_GetOperations_MockSuccess(t *testing.T) {
	t.Parallel()

	// Mock responses based on live WNC NMSP data structure - simplified for testing
	responses := map[string]string{
		"Cisco-IOS-XE-wireless-nmsp-oper:nmsp-oper-data": `{
			"Cisco-IOS-XE-wireless-nmsp-oper:nmsp-oper-data": {
				"client-registration": [
					{
						"client-id": 0,
						"services": {
							"mask": "6876135"
						}
					}
				],
				"cmx-connection": [
					{
						"peer-ip": "0.0.0.0",
						"connection-id": "0",
						"active": false
					}
				],
				"cmx-cloud-info": {
					"cloud-status": {
						"ip-address": "0.0.0.0",
						"connectivity": "nmsp-connectivity-down",
						"service-up": false
					}
				}
			}
		}`,
		"Cisco-IOS-XE-wireless-nmsp-oper:nmsp-oper-data/client-registration": `{
			"Cisco-IOS-XE-wireless-nmsp-oper:client-registration": [
				{
					"client-id": 0,
					"services": {
						"mask": "6876135"
					}
				}
			]
		}`,
		"Cisco-IOS-XE-wireless-nmsp-oper:nmsp-oper-data/cmx-connection": `{
			"Cisco-IOS-XE-wireless-nmsp-oper:cmx-connection": [
				{
					"peer-ip": "0.0.0.0",
					"connection-id": "0",
					"active": false
				}
			]
		}`,
		"Cisco-IOS-XE-wireless-nmsp-oper:nmsp-oper-data/cmx-cloud-info": `{
			"Cisco-IOS-XE-wireless-nmsp-oper:cmx-cloud-info": {
				"cloud-status": {
					"connectivity": "nmsp-connectivity-down",
					"service-up": false
				}
			}
		}`,
	}

	mockServer := testutil.NewMockServer(responses)
	defer mockServer.Close()
	testClient := testutil.NewTestClient(mockServer)
	service := nmsp.NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	t.Run("GetOperational", func(t *testing.T) {
		result, err := service.GetOperational(ctx)
		if err != nil {
			t.Errorf("GetOperational returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("GetOperational returned nil result")
		}
	})

	t.Run("ListClientRegistrations", func(t *testing.T) {
		result, err := service.ListClientRegistrations(ctx)
		if err != nil {
			t.Errorf("ListClientRegistrations returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("ListClientRegistrations returned nil result")
		}
	})

	t.Run("GetCMXConnectionInfo", func(t *testing.T) {
		result, err := service.GetCMXConnectionInfo(ctx)
		if err != nil {
			t.Errorf("GetCMXConnectionInfo returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("GetCMXConnectionInfo returned nil result")
		}
	})

	t.Run("GetCMXCloudInfo", func(t *testing.T) {
		result, err := service.GetCMXCloudInfo(ctx)
		if err != nil {
			t.Errorf("GetCMXCloudInfo returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("GetCMXCloudInfo returned nil result")
		}
	})
}
