package rf

import (
	"encoding/json"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/pkg/testutil"
)

func TestRfServiceUnit_Constructor_Success(t *testing.T) {
	t.Parallel()

	server := testutil.NewMockServer(testutil.WithSuccessResponses(map[string]string{}))
	defer server.Close()
	testClient := testutil.NewTestClient(server)
	service := NewService(testClient.Core().(*core.Client))
	if service.Client() == nil {
		t.Error("Expected valid client, got nil")
	}
}

func TestRfServiceUnit_GetOperations_MockSuccess(t *testing.T) {
	t.Parallel()

	// Mock responses based on real WNC data from RF configuration
	responses := map[string]string{
		"Cisco-IOS-XE-wireless-rf-cfg:rf-cfg-data": `{
			"Cisco-IOS-XE-wireless-rf-cfg:rf-cfg-data": {
				"rf-tags": {
					"rf-tag": [
						{
							"tag-name": "labo-inside",
							"dot11a-rf-profile-name": "labo-rf-5gh-inside",
							"dot11b-rf-profile-name": "labo-rf-24gh",
							"dot11-6ghz-rf-prof-name": "labo-rf-6gh"
						},
						{
							"tag-name": "default-rf-tag",
							"description": "Preconfigured default RF tag"
						}
					]
				},
				"rf-profiles": {
					"rf-profile": [
						{
							"profile-name": "labo-rf-5gh-inside",
							"rf-band": "dot11-5ghz-band",
							"description": "RF profile for 5GHz indoor"
						}
					]
				},
				"multi-bssid-profiles": {
					"multi-bssid-profile": []
				},
				"atf-policies": {
					"atf-policy": []
				},
				"rf-profile-default-entries": {
					"rf-profile-default-entry": []
				}
			}
		}`,
		// Individual wrapper endpoints
		"Cisco-IOS-XE-wireless-rf-cfg:rf-cfg-data/rf-tags": `{
			"Cisco-IOS-XE-wireless-rf-cfg:rf-tags": {
				"rf-tag": [
					{
						"tag-name": "labo-inside",
						"dot11a-rf-profile-name": "labo-rf-5gh-inside",
						"dot11b-rf-profile-name": "labo-rf-24gh",
						"dot11-6ghz-rf-prof-name": "labo-rf-6gh"
					},
					{
						"tag-name": "default-rf-tag",
						"description": "Preconfigured default RF tag"
					}
				]
			}
		}`,
		"Cisco-IOS-XE-wireless-rf-cfg:rf-cfg-data/rf-profiles": `{
			"Cisco-IOS-XE-wireless-rf-cfg:rf-profiles": {
				"rf-profile": [
					{
						"profile-name": "labo-rf-5gh-inside",
						"rf-band": "dot11-5ghz-band",
						"description": "RF profile for 5GHz indoor"
					}
				]
			}
		}`,
		"Cisco-IOS-XE-wireless-rf-cfg:rf-cfg-data/multi-bssid-profiles": `{
			"Cisco-IOS-XE-wireless-rf-cfg:multi-bssid-profiles": {
				"multi-bssid-profile": []
			}
		}`,
		"Cisco-IOS-XE-wireless-rf-cfg:rf-cfg-data/atf-policies": `{
			"Cisco-IOS-XE-wireless-rf-cfg:atf-policies": {
				"atf-policy": []
			}
		}`,
		"Cisco-IOS-XE-wireless-rf-cfg:rf-cfg-data/rf-profile-default-entries": `{
			"Cisco-IOS-XE-wireless-rf-cfg:rf-profile-default-entries": {
				"rf-profile-default-entry": []
			}
		}`,
		// RRM operational data responses
		"Cisco-IOS-XE-wireless-rrm-oper:rrm-oper-data": `{
			"Cisco-IOS-XE-wireless-rrm-oper:rrm-oper-data": {
				"ap-auto-rf-dot11-data": [
					{
						"wtp-mac": "28:ac:9e:bb:3c:80",
						"radio-slot-id": 0,
						"neighbor-radio-info": {
							"neighbor-radio-list": [
								{
									"neighbor-radio-info": {
										"neighbor-radio-mac": "f0:d8:05:2c:41:20",
										"neighbor-radio-slot-id": 0,
										"rssi": -21,
										"snr": 62,
										"channel": 11,
										"power": 18,
										"group-leader-ip": "192.168.255.4",
										"chan-width": "radio-neighbor-chan-width-20-mhz",
										"sensor-covered": false
									}
								}
							]
						}
					}
				],
				"ap-dot11-radar-data": [
					{
						"wtp-mac": "28:ac:9e:bb:3c:80",
						"radio-slot-id": 0,
						"last-radar-on-radio": "1970-01-01T00:00:00+00:00"
					}
				]
			}
		}`,
		"Cisco-IOS-XE-wireless-rrm-oper:rrm-oper-data/ap-auto-rf-dot11-data": `{
			"Cisco-IOS-XE-wireless-rrm-oper:ap-auto-rf-dot11-data": [
				{
					"wtp-mac": "28:ac:9e:bb:3c:80",
					"radio-slot-id": 0,
					"neighbor-radio-info": {
						"neighbor-radio-list": [
							{
								"neighbor-radio-info": {
									"neighbor-radio-mac": "f0:d8:05:2c:41:20",
									"neighbor-radio-slot-id": 0,
									"rssi": -21,
									"snr": 62,
									"channel": 11,
									"power": 18,
									"group-leader-ip": "192.168.255.4",
									"chan-width": "radio-neighbor-chan-width-20-mhz",
									"sensor-covered": false
								}
							}
						]
					}
				}
			]
		}`,
		"Cisco-IOS-XE-wireless-rrm-oper:rrm-oper-data/ap-dot11-radar-data": `{
			"Cisco-IOS-XE-wireless-rrm-oper:ap-dot11-radar-data": [
				{
					"wtp-mac": "28:ac:9e:bb:3c:80",
					"radio-slot-id": 0,
					"last-radar-on-radio": "1970-01-01T00:00:00+00:00"
				}
			]
		}`,
	}

	mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(responses))
	defer mockServer.Close()

	testClient := testutil.NewTestClient(mockServer)
	service := NewService(testClient.Core().(*core.Client))
	ctx := testutil.TestContext(t)

	t.Run("GetConfig", func(t *testing.T) {
		result, err := service.GetConfig(ctx)
		if err != nil {
			t.Errorf("GetConfig returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("GetConfig returned nil result")
		}
	})

	t.Run("RFTag", func(t *testing.T) {
		rfTagService := service.RFTag()
		if rfTagService == nil {
			t.Error("RFTag service returned nil")
		}
	})
	t.Run("ListRFTags", func(t *testing.T) {
		result, err := service.ListRFTags(ctx)
		if err != nil {
			t.Errorf("ListRFTags returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("ListRFTags returned nil result")
		}
	})

	t.Run("ListRFProfiles", func(t *testing.T) {
		result, err := service.ListRFProfiles(ctx)
		if err != nil {
			t.Errorf("ListRFProfiles returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("ListRFProfiles returned nil result")
		}
	})

	t.Run("ListMultiBssidProfiles", func(t *testing.T) {
		result, err := service.ListMultiBssidProfiles(ctx)
		if err != nil {
			t.Errorf("ListMultiBssidProfiles returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("ListMultiBssidProfiles returned nil result")
		}
	})

	t.Run("ListAtfPolicies", func(t *testing.T) {
		result, err := service.ListAtfPolicies(ctx)
		if err != nil {
			t.Errorf("ListAtfPolicies returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("ListAtfPolicies returned nil result")
		}
	})

	t.Run("ListRFProfileDefaultEntries", func(t *testing.T) {
		result, err := service.ListRFProfileDefaultEntries(ctx)
		if err != nil {
			t.Errorf("ListRFProfileDefaultEntries returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("ListRFProfileDefaultEntries returned nil result")
		}
	})

	t.Run("GetOperational", func(t *testing.T) {
		result, err := service.GetOperational(ctx)
		if err != nil {
			t.Errorf("GetOperational returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("GetOperational returned nil result")
		}
	})

	t.Run("GetAutoRFDot11Data", func(t *testing.T) {
		result, err := service.GetAutoRFDot11Data(ctx)
		if err != nil {
			t.Errorf("GetAutoRFDot11Data returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("GetAutoRFDot11Data returned nil result")
		}
	})

	t.Run("GetRadarDetectionData", func(t *testing.T) {
		result, err := service.GetRadarDetectionData(ctx)
		if err != nil {
			t.Errorf("GetRadarDetectionData returned unexpected error: %v", err)
		}
		if result == nil {
			t.Error("GetRadarDetectionData returned nil result")
		}
	})
}

// TestRfServiceUnit_OmittedRadioProfileName_MockSuccess tests that a slot override read without
// with-defaults stays nil while an explicit name decodes, and that a nil one is not written.
//
// No contract gate holds this field: it is neither a leaf a consumer publishes nor one the schema
// defaults to true, so this test is the only thing a revert to a value string fails.
func TestRfServiceUnit_OmittedRadioProfileName_MockSuccess(t *testing.T) {
	t.Parallel()

	mockServer := testutil.NewMockServer(testutil.WithSuccessResponses(map[string]string{
		"Cisco-IOS-XE-wireless-rf-cfg:rf-cfg-data/rf-tags": `{
			"Cisco-IOS-XE-wireless-rf-cfg:rf-tags": {
				"rf-tag": [
					{
						"tag-name": "tag-plain-read",
						"rf-tag-radio-profiles": {
							"rf-tag-radio-profile": [
								{
									"slot-id": "slot-0",
									"band-id": "band-2-dot-4-ghz"
								}
							]
						}
					},
					{
						"tag-name": "tag-report-all",
						"rf-tag-radio-profiles": {
							"rf-tag-radio-profile": [
								{
									"slot-id": "slot-0",
									"band-id": "band-2-dot-4-ghz",
									"radio-profile-name": "profile-slot-0"
								}
							]
						}
					}
				]
			}
		}`,
	}))
	defer mockServer.Close()

	testClient := testutil.NewTestClient(mockServer)
	service := NewService(testClient.Core().(*core.Client))

	result, err := service.ListRFTags(testutil.TestContext(t))
	if err != nil {
		t.Fatalf("ListRFTags failed: %v", err)
	}
	if result.RFTags == nil || len(result.RFTags.RFTags) != 2 {
		t.Fatalf("Expected 2 RF tags, got %+v", result.RFTags)
	}

	omitted := result.RFTags.RFTags[0].RFTagRadioProfiles.RFTagRadioProfile[0]
	if omitted.SlotID != "slot-0" {
		t.Fatalf("slot-id = %q, so the list element was not decoded", omitted.SlotID)
	}
	if omitted.RadioProfileName != nil {
		t.Error("radio-profile-name: want nil, which is a plain read rather than a slot with no override")
	}

	sent := result.RFTags.RFTags[1].RFTagRadioProfiles.RFTagRadioProfile[0]
	if sent.RadioProfileName == nil || *sent.RadioProfileName != "profile-slot-0" {
		t.Errorf("radio-profile-name = %v, want profile-slot-0", sent.RadioProfileName)
	}

	// The write side: a nil pointer marshals away, so declaring the field sends no new leaf on the
	// replacing PUT. Asserted on the bytes, because that is what the controller receives.
	body, err := json.Marshal(omitted)
	if err != nil {
		t.Fatalf("json.Marshal failed: %v", err)
	}
	if got := string(body); got != `{"slot-id":"slot-0","band-id":"band-2-dot-4-ghz"}` {
		t.Errorf("a nil radio-profile-name reached the wire: %s", got)
	}
}

func TestRfServiceUnit_ErrorHandling_NilClient(t *testing.T) {
	t.Parallel()

	service := NewService(nil)
	ctx := testutil.TestContext(t)

	t.Run("GetConfig_NilClient", func(t *testing.T) {
		result, err := service.GetConfig(ctx)
		if err == nil {
			t.Error("Expected error for nil client")
		}
		if result != nil {
			t.Error("Expected nil result for error case")
		}
	})

	t.Run("ListRFTags_NilClient", func(t *testing.T) {
		result, err := service.ListRFTags(ctx)
		if err == nil {
			t.Error("Expected error for nil client")
		}
		if result != nil {
			t.Error("Expected nil result for error case")
		}
	})

	t.Run("ListRFProfiles_NilClient", func(t *testing.T) {
		result, err := service.ListRFProfiles(ctx)
		if err == nil {
			t.Error("Expected error for nil client")
		}
		if result != nil {
			t.Error("Expected nil result for error case")
		}
	})

	t.Run("ListMultiBssidProfiles_NilClient", func(t *testing.T) {
		result, err := service.ListMultiBssidProfiles(ctx)
		if err == nil {
			t.Error("Expected error for nil client")
		}
		if result != nil {
			t.Error("Expected nil result for error case")
		}
	})

	t.Run("ListAtfPolicies_NilClient", func(t *testing.T) {
		result, err := service.ListAtfPolicies(ctx)
		if err == nil {
			t.Error("Expected error for nil client")
		}
		if result != nil {
			t.Error("Expected nil result for error case")
		}
	})

	t.Run("ListRFProfileDefaultEntries_NilClient", func(t *testing.T) {
		result, err := service.ListRFProfileDefaultEntries(ctx)
		if err == nil {
			t.Error("Expected error for nil client")
		}
		if result != nil {
			t.Error("Expected nil result for error case")
		}
	})

	t.Run("GetOperational_NilClient", func(t *testing.T) {
		result, err := service.GetOperational(ctx)
		if err == nil {
			t.Error("Expected error for nil client")
		}
		if result != nil {
			t.Error("Expected nil result for error case")
		}
	})

	t.Run("GetAutoRFDot11Data_NilClient", func(t *testing.T) {
		result, err := service.GetAutoRFDot11Data(ctx)
		if err == nil {
			t.Error("Expected error for nil client")
		}
		if result != nil {
			t.Error("Expected nil result for error case")
		}
	})

	t.Run("GetRadarDetectionData_NilClient", func(t *testing.T) {
		result, err := service.GetRadarDetectionData(ctx)
		if err == nil {
			t.Error("Expected error for nil client")
		}
		if result != nil {
			t.Error("Expected nil result for error case")
		}
	})
}
