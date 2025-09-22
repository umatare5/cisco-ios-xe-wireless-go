package dot11_test

import (
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/pkg/testutil"
	"github.com/umatare5/cisco-ios-xe-wireless-go/service/dot11"
)

// TestDot11ServiceUnit_Constructor_Success validates the DOT11 service constructor.
func TestDot11ServiceUnit_Constructor_Success(t *testing.T) {
	tests := []struct {
		name         string
		clientExists bool
	}{
		{"NewServiceWithValidClient", true},
		{"NewServiceWithNilClient", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var service dot11.Service
			if tt.clientExists {
				mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(map[string]string{}))
				defer mockServer.Close()
				client := testutil.NewTestClient(mockServer)
				service = dot11.NewService(client.Core().(*core.Client))
			} else {
				service = dot11.NewService(nil)
			}

			if tt.clientExists && service.Client() == nil {
				t.Error("Expected valid client but got nil")
			}
			if !tt.clientExists && service.Client() != nil {
				t.Error("Expected nil client but got valid client")
			}
		})
	}
}

// TestDot11ServiceUnit_GetOperations_MockSuccess validates successful operations.
func TestDot11ServiceUnit_GetOperations_MockSuccess(t *testing.T) {
	t.Parallel()

	mockServer := testutil.NewMockServer(
		testutil.WithSuccessResponses(map[string]string{
			"Cisco-IOS-XE-wireless-dot11-cfg:dot11-cfg-data":                                         `{"Cisco-IOS-XE-wireless-dot11-cfg:dot11-cfg-data":{}}`,
			"Cisco-IOS-XE-wireless-dot11-cfg:dot11-cfg-data/configured-countries":                    `{"Cisco-IOS-XE-wireless-dot11-cfg:configured-countries":{}}`,
			"Cisco-IOS-XE-wireless-dot11-cfg:dot11-cfg-data/dot11-entries":                           `{"Cisco-IOS-XE-wireless-dot11-cfg:dot11-entries":{}}`,
			"Cisco-IOS-XE-wireless-dot11-cfg:dot11-cfg-data/dot11ac-mcs-entries":                     `{"Cisco-IOS-XE-wireless-dot11-cfg:dot11ac-mcs-entries":{}}`,
			"Cisco-IOS-XE-wireless-dot11-cfg:dot11-cfg-data/configured-countries/configured-country": `{"Cisco-IOS-XE-wireless-dot11-cfg:configured-country":{}}`,
			"Cisco-IOS-XE-wireless-dot11-cfg:dot11-cfg-data/dot11-entries/dot11-entry":               `{"Cisco-IOS-XE-wireless-dot11-cfg:dot11-entry":{}}`,
		}),
	)
	defer mockServer.Close()

	client := testutil.NewTestClient(mockServer)
	service := dot11.NewService(client.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	tests := []struct {
		name string
		fn   func() error
	}{
		{"GetConfig", func() error { _, err := service.GetConfig(ctx); return err }},
		{"ListCfgFilters", func() error { _, err := service.ListCfgFilters(ctx); return err }},
		{"ListCfgConfiguredCountries", func() error { _, err := service.ListCfgConfiguredCountries(ctx); return err }},
		{"ListCfgDot11Entries", func() error { _, err := service.ListCfgDot11Entries(ctx); return err }},
		{"ListCfgDot11acMcsEntries", func() error { _, err := service.ListCfgDot11acMcsEntries(ctx); return err }},
		{"ListConfiguredCountries", func() error { _, err := service.ListConfiguredCountries(ctx); return err }},
		{"ListDot11Entries", func() error { _, err := service.ListDot11Entries(ctx); return err }},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.fn(); err != nil {
				t.Errorf("Expected no error for %s, got: %v", tt.name, err)
			}
		})
	}
}

// TestDot11ServiceUnit_GetOperations_ErrorHandling validates error handling.
func TestDot11ServiceUnit_GetOperations_ErrorHandling(t *testing.T) {
	t.Parallel()

	mockServer := testutil.NewMockServer(
		testutil.WithErrorResponses([]string{
			"Cisco-IOS-XE-wireless-dot11-cfg:dot11-cfg-data",
			"Cisco-IOS-XE-wireless-dot11-cfg:dot11-cfg-data/configured-countries",
			"Cisco-IOS-XE-wireless-dot11-cfg:dot11-cfg-data/dot11-entries",
			"Cisco-IOS-XE-wireless-dot11-cfg:dot11-cfg-data/dot11ac-mcs-entries",
			"Cisco-IOS-XE-wireless-dot11-cfg:dot11-cfg-data/configured-countries/configured-country",
			"Cisco-IOS-XE-wireless-dot11-cfg:dot11-cfg-data/dot11-entries/dot11-entry",
		}, 404),
	)
	defer mockServer.Close()

	client := testutil.NewTestClient(mockServer)
	service := dot11.NewService(client.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	tests := []struct {
		name string
		fn   func() error
	}{
		{"GetConfig", func() error { _, err := service.GetConfig(ctx); return err }},
		{"ListCfgFilters", func() error { _, err := service.ListCfgFilters(ctx); return err }},
		{"ListCfgConfiguredCountries", func() error { _, err := service.ListCfgConfiguredCountries(ctx); return err }},
		{"ListCfgDot11Entries", func() error { _, err := service.ListCfgDot11Entries(ctx); return err }},
		{"ListCfgDot11acMcsEntries", func() error { _, err := service.ListCfgDot11acMcsEntries(ctx); return err }},
		{"ListConfiguredCountries", func() error { _, err := service.ListConfiguredCountries(ctx); return err }},
		{"ListDot11Entries", func() error { _, err := service.ListDot11Entries(ctx); return err }},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.fn(); err == nil {
				t.Errorf("Expected error for %s, got nil", tt.name)
			}
		})
	}
}

// TestDot11ServiceUnit_ErrorHandling_NilClient validates nil client handling.
func TestDot11ServiceUnit_ErrorHandling_NilClient(t *testing.T) {
	t.Parallel()

	service := dot11.NewService(nil)
	ctx := testutil.TestContext(t)

	tests := []struct {
		name string
		fn   func() error
	}{
		{"GetConfig_NilClient", func() error { _, err := service.GetConfig(ctx); return err }},
		{"ListCfgFilters_NilClient", func() error { _, err := service.ListCfgFilters(ctx); return err }},
		{"ListDot11Entries_NilClient", func() error { _, err := service.ListDot11Entries(ctx); return err }},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.fn(); err == nil {
				t.Errorf("Expected error for %s with nil client, got nil", tt.name)
			}
		})
	}
}
