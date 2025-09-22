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
		mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(map[string]string{}))
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
	mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(map[string]string{
		"Cisco-IOS-XE-wireless-fabric-cfg:fabric-cfg-data": `{
			"Cisco-IOS-XE-wireless-fabric-cfg:fabric-cfg-data": {
				"fabric": {
					"fabric-enabled": true
				},
				"fabric-controlplane-names": {
					"fabric-controlplane-name": [
						{
							"control-plane-name": "default-control-plane",
							"description": "Preconfigured default control plane name"
						}
					]
				}
			}
		}`,
		"Cisco-IOS-XE-wireless-fabric-cfg:fabric-cfg-data/fabric": `{
			"Cisco-IOS-XE-wireless-fabric-cfg:fabric": {
				"fabric-enabled": true
			}
		}`,
		"Cisco-IOS-XE-wireless-fabric-cfg:fabric-cfg-data/fabric-profiles": `{
			"Cisco-IOS-XE-wireless-fabric-cfg:fabric-profiles": {
				"fabric-profile": []
			}
		}`,
		"Cisco-IOS-XE-wireless-fabric-cfg:fabric-cfg-data/fabric-controlplane-names": `{
			"Cisco-IOS-XE-wireless-fabric-cfg:fabric-controlplane-names": {
				"fabric-controlplane-name": [
					{
						"control-plane-name": "default-control-plane",
						"description": "Preconfigured default control plane name"
					}
				]
			}
		}`,
		"Cisco-IOS-XE-wireless-fabric-cfg:fabric-cfg-data/fabric/fabric-name-vnid-entries": `{
			"Cisco-IOS-XE-wireless-fabric-cfg:fabric-name-vnid-entries": {
				"fabric-name-vnid-entry": []
			}
		}`,
		"Cisco-IOS-XE-wireless-fabric-cfg:fabric-cfg-data/fabric-controlplane-names/fabric-controlplane-name/fabric-control-plane-ip-configs": `{
			"Cisco-IOS-XE-wireless-fabric-cfg:fabric-control-plane-ip-configs": {
				"fabric-control-plane-ip-config": []
			}
		}`,
	}))
	defer mockServer.Close()

	client := testutil.NewTestClient(mockServer)
	service := fabric.NewService(client.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	tests := []struct {
		name string
		fn   func() error
	}{
		{"GetConfig", func() error { _, err := service.GetConfig(ctx); return err }},
		{"ListCfgFabric", func() error { _, err := service.ListCfgFabric(ctx); return err }},
		{
			"ListCfgFabricProfiles",
			func() error { _, err := service.ListCfgFabricProfiles(ctx); return err },
		},
		{
			"ListCfgFabricControlplaneNames",
			func() error { _, err := service.ListCfgFabricControlplaneNames(ctx); return err },
		},
		{"ListFabricConfig", func() error { _, err := service.ListFabricConfig(ctx); return err }},
		{
			"ListFabricProfiles",
			func() error { _, err := service.ListFabricProfiles(ctx); return err },
		},
		{
			"ListFabricProfile",
			func() error { _, err := service.ListFabricProfile(ctx); return err },
		},
		{
			"ListFabricControlplanes",
			func() error { _, err := service.ListFabricControlplanes(ctx); return err },
		},
		{
			"ListFabricControlplaneName",
			func() error { _, err := service.ListFabricControlplaneName(ctx); return err },
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.fn(); err != nil {
				t.Errorf("Expected no error for %s, got: %v", tt.name, err)
			}
		})
	}
}

// TestFabricServiceUnit_GetOperations_ErrorHandling tests error scenarios using mock server.
func TestFabricServiceUnit_GetOperations_ErrorHandling(t *testing.T) {
	mockServer := testutil.NewMockServer(testutil.WithErrorResponses([]string{
		"Cisco-IOS-XE-wireless-fabric-cfg:fabric-cfg-data",
		"Cisco-IOS-XE-wireless-fabric-cfg:fabric-cfg-data/fabric",
		"Cisco-IOS-XE-wireless-fabric-cfg:fabric-cfg-data/fabric-profiles",
		"Cisco-IOS-XE-wireless-fabric-cfg:fabric-cfg-data/fabric-controlplane-names",
		"Cisco-IOS-XE-wireless-fabric-cfg:fabric-cfg-data/fabric/fabric-name-vnid-entries",
		"Cisco-IOS-XE-wireless-fabric-cfg:fabric-cfg-data/fabric-controlplane-names/fabric-controlplane-name/fabric-control-plane-ip-configs",
	}, 404))
	defer mockServer.Close()

	client := testutil.NewTestClient(mockServer)
	service := fabric.NewService(client.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	tests := []struct {
		name string
		fn   func() error
	}{
		{"GetConfig", func() error { _, err := service.GetConfig(ctx); return err }},
		{"ListCfgFabric", func() error { _, err := service.ListCfgFabric(ctx); return err }},
		{
			"ListCfgFabricProfiles",
			func() error { _, err := service.ListCfgFabricProfiles(ctx); return err },
		},
		{
			"ListCfgFabricControlplaneNames",
			func() error { _, err := service.ListCfgFabricControlplaneNames(ctx); return err },
		},
		{"ListFabricConfig", func() error { _, err := service.ListFabricConfig(ctx); return err }},
		{
			"ListFabricProfiles",
			func() error { _, err := service.ListFabricProfiles(ctx); return err },
		},
		{
			"ListFabricProfile",
			func() error { _, err := service.ListFabricProfile(ctx); return err },
		},
		{
			"ListFabricControlplanes",
			func() error { _, err := service.ListFabricControlplanes(ctx); return err },
		},
		{
			"ListFabricControlplaneName",
			func() error { _, err := service.ListFabricControlplaneName(ctx); return err },
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.fn(); err == nil {
				t.Errorf("Expected error for %s, got nil", tt.name)
			}
		})
	}
}
