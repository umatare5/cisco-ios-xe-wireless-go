package awips

import (
	"os"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/client"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil/framework"
)

// TestService tests the AWIPS service using the minimal test pattern for oper-only service
func TestService(t *testing.T) {
	// Create and run minimal service test suite
	suite := framework.MinimalServiceTestSuite{
		ServiceName: "AWIPS",
		NewServiceFunc: func(client any) any {
			if client == nil {
				return NewService(nil)
			}
			return NewService(client.(*core.Client))
		},
		GetClientFunc: func(service any) any {
			if service == nil {
				return nil
			}
			s := service.(Service)
			return s.Client()
		},
	}

	framework.RunMinimalServiceTests(t, suite)
}

// TestAwipsServiceConstants tests service constants and configuration
func TestAwipsServiceConstants(t *testing.T) {
	// Test service constants
	t.Run("ServiceConstants", func(t *testing.T) {
		if EntityTypeAWIPS == "" {
			t.Error("EntityTypeAWIPS should not be empty")
		}
		if routes.AWIPSOperBasePath == "" {
			t.Error("routes.AWIPSOperBasePath should not be empty")
		}
		if routes.AWIPSOperEndpoint == "" {
			t.Error("routes.AWIPSOperEndpoint should not be empty")
		}
	})
}

// TestAwipsServiceIntegration tests integration with specific AP MAC when available
func TestAwipsServiceIntegration(t *testing.T) {
	// Integration test only if client is available
	setup := client.SetupOptionalClient(t)
	if setup.Client == nil {
		t.Log("Skipping integration tests: no client available")
		return
	}
	client.SkipIfNoConnection(t, setup.Client)

	service := NewService(setup.Client)

	t.Run("GetOper", func(t *testing.T) {
		data, err := service.GetOper(setup.Context)
		if err != nil {
			t.Logf("GetOper returned error (may be expected): %v", err)
			return
		}
		if data != nil {
			t.Log("Successfully retrieved AWIPS operational data")
		}
	})

	// AP-specific tests with WNC_TEST_AP_MAC
	testApMac := os.Getenv("WNC_TEST_AP_MAC")
	if testApMac != "" {
		t.Run("GetOperByApMac", func(t *testing.T) {
			data, err := service.GetOperByApMac(setup.Context, testApMac)
			if err != nil {
				t.Logf("GetOperByApMac for %s returned error (may be expected): %v", testApMac, err)
				return
			}
			if data != nil {
				t.Logf("Successfully retrieved AWIPS operational data for AP %s", testApMac)
			}
		})

		t.Run("GetOperByApMacDownloadStatus", func(t *testing.T) {
			data, err := service.GetOperByApMacDownloadStatus(setup.Context, testApMac)
			if err != nil {
				t.Logf("GetOperByApMacDownloadStatus for %s returned error (may be expected): %v", testApMac, err)
				return
			}
			if data != nil {
				t.Logf("Successfully retrieved AWIPS download status for AP %s", testApMac)
			}
		})
	} else {
		t.Log("WNC_TEST_AP_MAC not set, skipping AP-specific tests")
	}
}
