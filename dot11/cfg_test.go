// Package dot11 provides 802.11 configuration test functionality for the Cisco Wireless Network Controller API.
package dot11

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	testutils "github.com/umatare5/cisco-ios-xe-wireless-go/internal/tests"
)

// =============================================================================
// 1. UNIT TESTS (Structure/Type Validation & JSON Serialization/Deserialization)
// =============================================================================

func TestDot11CfgDataStructures(t *testing.T) {
	// Test Dot11CfgResponse structure
	t.Run("Dot11CfgResponse", func(t *testing.T) {
		sampleJSON := `{
			"Cisco-IOS-XE-wireless-dot11-cfg:dot11-cfg-data": {
				"configured-countries": {
					"configured-country": [
						{"country-code": "US"},
						{"country-code": "CA"}
					]
				},
				"dot11ac-mcs-entries": {
					"dot11ac-mcs-entry": [
						{"spatial-stream": 1, "index": "0"}
					]
				},
				"dot11-entries": {
					"dot11-entry": [
						{
							"band": "5GHz",
							"voice-adm-ctrl-support": true
						}
					]
				}
			}
		}`

		var response Dot11CfgResponse
		err := json.Unmarshal([]byte(sampleJSON), &response)
		if err != nil {
			t.Fatalf("Failed to unmarshal Dot11CfgResponse: %v", err)
		}

		if len(response.CiscoIOSXEWirelessDot11CfgDot11CfgData.ConfiguredCountries.ConfiguredCountry) != 2 {
			t.Errorf("Expected 2 configured countries, got %d",
				len(response.CiscoIOSXEWirelessDot11CfgDot11CfgData.ConfiguredCountries.ConfiguredCountry))
		}

		if response.CiscoIOSXEWirelessDot11CfgDot11CfgData.ConfiguredCountries.ConfiguredCountry[0].CountryCode != "US" {
			t.Errorf("Expected first country code 'US', got '%s'",
				response.CiscoIOSXEWirelessDot11CfgDot11CfgData.ConfiguredCountries.ConfiguredCountry[0].CountryCode)
		}
	})

	// Test Dot11ConfiguredCountriesResponse structure
	t.Run("Dot11ConfiguredCountriesResponse", func(t *testing.T) {
		sampleJSON := `{
			"Cisco-IOS-XE-wireless-dot11-cfg:configured-countries": {
				"configured-country": [
					{"country-code": "JP"},
					{"country-code": "DE"}
				]
			}
		}`

		var response Dot11ConfiguredCountriesResponse
		err := json.Unmarshal([]byte(sampleJSON), &response)
		if err != nil {
			t.Fatalf("Failed to unmarshal Dot11ConfiguredCountriesResponse: %v", err)
		}

		if len(response.ConfiguredCountries.ConfiguredCountry) != 2 {
			t.Errorf("Expected 2 configured countries, got %d",
				len(response.ConfiguredCountries.ConfiguredCountry))
		}

		if response.ConfiguredCountries.ConfiguredCountry[0].CountryCode != "JP" {
			t.Errorf("Expected first country code 'JP', got '%s'",
				response.ConfiguredCountries.ConfiguredCountry[0].CountryCode)
		}
	})

	// Test Dot11acMcsEntriesResponse structure
	t.Run("Dot11acMcsEntriesResponse", func(t *testing.T) {
		sampleJSON := `{
			"Cisco-IOS-XE-wireless-dot11-cfg:dot11ac-mcs-entries": {
				"dot11ac-mcs-entry": [
					{"spatial-stream": 2, "index": "5"},
					{"spatial-stream": 4, "index": "9"}
				]
			}
		}`

		var response Dot11acMcsEntriesResponse
		err := json.Unmarshal([]byte(sampleJSON), &response)
		if err != nil {
			t.Fatalf("Failed to unmarshal Dot11acMcsEntriesResponse: %v", err)
		}

		if len(response.Dot11acMcsEntries.Dot11acMcsEntry) != 2 {
			t.Errorf("Expected 2 MCS entries, got %d",
				len(response.Dot11acMcsEntries.Dot11acMcsEntry))
		}

		entry := response.Dot11acMcsEntries.Dot11acMcsEntry[0]
		if entry.SpatialStream != 2 {
			t.Errorf("Expected spatial stream 2, got %d", entry.SpatialStream)
		}

		if entry.Index != "5" {
			t.Errorf("Expected index '5', got '%s'", entry.Index)
		}
	})

	// Test Dot11EntriesResponse structure
	t.Run("Dot11EntriesResponse", func(t *testing.T) {
		sampleJSON := `{
			"Cisco-IOS-XE-wireless-dot11-cfg:dot11-entries": {
				"dot11-entry": [
					{
						"band": "2.4GHz",
						"voice-adm-ctrl-support": false,
						"dot11ax-cfg": {
							"he-bss-color": true
						},
						"ampdu-entries": {
							"ampdu-entry": [
								{
									"index": 1,
									"apf-80211n-ampdu-tx-priority": "high"
								}
							]
						}
					}
				]
			}
		}`

		var response Dot11EntriesResponse
		err := json.Unmarshal([]byte(sampleJSON), &response)
		if err != nil {
			t.Fatalf("Failed to unmarshal Dot11EntriesResponse: %v", err)
		}

		if len(response.Dot11Entries.Dot11Entry) != 1 {
			t.Errorf("Expected 1 dot11 entry, got %d", len(response.Dot11Entries.Dot11Entry))
		}

		entry := response.Dot11Entries.Dot11Entry[0]
		if entry.Band != "2.4GHz" {
			t.Errorf("Expected band '2.4GHz', got '%s'", entry.Band)
		}

		if entry.VoiceAdmCtrlSupport {
			t.Error("Expected voice admission control support to be false")
		}

		if entry.Dot11axCfg == nil || !entry.Dot11axCfg.HeBssColor {
			t.Error("Expected HE BSS Color to be true")
		}

		if entry.AmpduEntries == nil || len(entry.AmpduEntries.AmpduEntry) != 1 {
			t.Error("Expected 1 AMPDU entry")
		} else {
			ampdu := entry.AmpduEntries.AmpduEntry[0]
			if ampdu.Index != 1 {
				t.Errorf("Expected AMPDU index 1, got %d", ampdu.Index)
			}
			if ampdu.Apf80211nAmpduTxPriority != "high" {
				t.Errorf("Expected AMPDU priority 'high', got '%s'", ampdu.Apf80211nAmpduTxPriority)
			}
		}
	})

	// Test ConfiguredCountry structure
	t.Run("ConfiguredCountry", func(t *testing.T) {
		sampleJSON := `{"country-code": "FR"}`

		var country ConfiguredCountry
		err := json.Unmarshal([]byte(sampleJSON), &country)
		if err != nil {
			t.Fatalf("Failed to unmarshal ConfiguredCountry: %v", err)
		}

		if country.CountryCode != "FR" {
			t.Errorf("Expected country code 'FR', got '%s'", country.CountryCode)
		}
	})

	// Test Dot11acMcsEntry structure
	t.Run("Dot11acMcsEntry", func(t *testing.T) {
		sampleJSON := `{"spatial-stream": 3, "index": "7"}`

		var entry Dot11acMcsEntry
		err := json.Unmarshal([]byte(sampleJSON), &entry)
		if err != nil {
			t.Fatalf("Failed to unmarshal Dot11acMcsEntry: %v", err)
		}

		if entry.SpatialStream != 3 {
			t.Errorf("Expected spatial stream 3, got %d", entry.SpatialStream)
		}

		if entry.Index != "7" {
			t.Errorf("Expected index '7', got '%s'", entry.Index)
		}
	})
}

// =============================================================================
// 2. INTEGRATION TESTS (API Endpoint Testing with Live Data Validation)
// =============================================================================

// TestDot11ConfigurationFunctions tests all Dot11 configuration functions with real WNC data collection
func TestDot11ConfigurationFunctions(t *testing.T) {
	client := testutils.CreateTestClientFromEnv(t)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Create a comprehensive test data collection
	collector := testutils.NewTestDataCollector()
	endpointMapping := map[string]string{
		"Dot11CfgEndpoint":                 "/restconf/data/Cisco-IOS-XE-wireless-dot11-cfg:dot11-cfg-data",
		"Dot11ConfiguredCountriesEndpoint": "/restconf/data/Cisco-IOS-XE-wireless-dot11-cfg:dot11-cfg-data/configured-countries",
		"Dot11acMcsEntriesEndpoint":        "/restconf/data/Cisco-IOS-XE-wireless-dot11-cfg:dot11-cfg-data/dot11ac-mcs-entries",
		"Dot11EntriesEndpoint":             "/restconf/data/Cisco-IOS-XE-wireless-dot11-cfg:dot11-cfg-data/dot11-entries",
	}

	t.Run("GetDot11Cfg", func(t *testing.T) {
		result, err := GetDot11Cfg(client, ctx)
		testutils.CollectTestResult(collector, "GetDot11Cfg", endpointMapping["Dot11CfgEndpoint"], result, err)
		if err != nil {
			t.Logf("GetDot11Cfg failed: %v", err)
		}
	})

	t.Run("GetDot11ConfiguredCountries", func(t *testing.T) {
		result, err := GetDot11ConfiguredCountries(client, ctx)
		testutils.CollectTestResult(collector, "GetDot11ConfiguredCountries", endpointMapping["Dot11ConfiguredCountriesEndpoint"], result, err)
		if err != nil {
			t.Logf("GetDot11ConfiguredCountries failed: %v", err)
		}
	})

	t.Run("GetDot11acMcsEntries", func(t *testing.T) {
		result, err := GetDot11acMcsEntries(client, ctx)
		testutils.CollectTestResult(collector, "GetDot11acMcsEntries", endpointMapping["Dot11acMcsEntriesEndpoint"], result, err)
		if err != nil {
			t.Logf("GetDot11acMcsEntries failed: %v", err)
		}
	})

	t.Run("GetDot11Entries", func(t *testing.T) {
		result, err := GetDot11Entries(client, ctx)
		testutils.CollectTestResult(collector, "GetDot11Entries", endpointMapping["Dot11EntriesEndpoint"], result, err)
		if err != nil {
			t.Logf("GetDot11Entries failed: %v", err)
		}
	})

	// Save collected test data to JSON file
	testutils.SaveCollectedTestData(t, collector, "dot11_cfg_test_data_collected.json")

	// Test error handling with nil client
	t.Run("GetDot11CfgWithNilClient", func(t *testing.T) {
		_, err := GetDot11Cfg(nil, ctx)
		if err == nil {
			t.Error("Expected error with nil client, got nil")
		}
		if err.Error() != "client is nil" {
			t.Errorf("Expected 'client is nil' error, got: %v", err)
		}
	})

	t.Run("GetDot11ConfiguredCountriesWithNilClient", func(t *testing.T) {
		_, err := GetDot11ConfiguredCountries(nil, ctx)
		if err == nil {
			t.Error("Expected error with nil client, got nil")
		}
		if err.Error() != "client is nil" {
			t.Errorf("Expected 'client is nil' error, got: %v", err)
		}
	})

	t.Run("GetDot11acMcsEntriesWithNilClient", func(t *testing.T) {
		_, err := GetDot11acMcsEntries(nil, ctx)
		if err == nil {
			t.Error("Expected error with nil client, got nil")
		}
		if err.Error() != "client is nil" {
			t.Errorf("Expected 'client is nil' error, got: %v", err)
		}
	})

	t.Run("GetDot11EntriesWithNilClient", func(t *testing.T) {
		_, err := GetDot11Entries(nil, ctx)
		if err == nil {
			t.Error("Expected error with nil client, got nil")
		}
		if err.Error() != "client is nil" {
			t.Errorf("Expected 'client is nil' error, got: %v", err)
		}
	})
}

// TestDot11ConfigurationEndpoints tests that all 802.11 configuration endpoints are correctly defined
func TestDot11ConfigurationEndpoints(t *testing.T) {
	endpoints := map[string]string{
		"Dot11CfgBasePath":                 Dot11CfgBasePath,
		"Dot11CfgEndpoint":                 "/restconf/data/Cisco-IOS-XE-wireless-dot11-cfg:dot11-cfg-data",
		"Dot11ConfiguredCountriesEndpoint": "/restconf/data/Cisco-IOS-XE-wireless-dot11-cfg:dot11-cfg-data/configured-countries",
		"Dot11acMcsEntriesEndpoint":        "/restconf/data/Cisco-IOS-XE-wireless-dot11-cfg:dot11-cfg-data/dot11ac-mcs-entries",
		"Dot11EntriesEndpoint":             "/restconf/data/Cisco-IOS-XE-wireless-dot11-cfg:dot11-cfg-data/dot11-entries",
	}

	testutils.ValidateEndpoints(t, endpoints)
}
