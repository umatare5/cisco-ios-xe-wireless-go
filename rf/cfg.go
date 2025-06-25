// Package rf provides RF (Radio Frequency) configuration functionality for the Cisco Wireless Network Controller API.
package rf

import (
	"context"

	wnc "github.com/umatare5/cisco-xe-wireless-restconf-go"
)

const (
	// RfCfgBasePath defines the base path for RF configuration endpoints
	RfCfgBasePath = "/restconf/data/Cisco-IOS-XE-wireless-rf-cfg:rf-cfg-data"
	// RfCfgEndpoint retrieves complete RF configuration data
	RfCfgEndpoint = RfCfgBasePath
	// MultiBssidProfilesEndpoint retrieves multi-BSSID profiles configuration
	MultiBssidProfilesEndpoint = RfCfgBasePath + "/multi-bssid-profiles"
	// AtfPoliciesEndpoint retrieves ATF (Airtime Fairness) policies configuration
	AtfPoliciesEndpoint = RfCfgBasePath + "/atf-policies"
	// RfTagsEndpoint retrieves RF tags configuration
	RfTagsEndpoint = RfCfgBasePath + "/rf-tags"
	// RfProfilesEndpoint retrieves RF profiles configuration
	RfProfilesEndpoint = RfCfgBasePath + "/rf-profiles"
	// RfProfileDefaultEntriesEndpoint retrieves RF profile default entries
	RfProfileDefaultEntriesEndpoint = RfCfgBasePath + "/rf-profile-default-entries"
)

// RfCfgResponse represents the complete RF configuration response
type RfCfgResponse struct {
	CiscoIOSXEWirelessRfCfgRfCfgData struct {
		MultiBssidProfiles      MultiBssidProfiles      `json:"multi-bssid-profiles"`
		AtfPolicies             AtfPolicies             `json:"atf-policies"`
		RfTags                  RfTags                  `json:"rf-tags"`
		RfProfiles              RfProfiles              `json:"rf-profiles"`
		RfProfileDefaultEntries RfProfileDefaultEntries `json:"rf-profile-default-entries"`
	} `json:"Cisco-IOS-XE-wireless-rf-cfg:rf-cfg-data"`
}

// MultiBssidProfilesResponse represents the multi-BSSID profiles configuration response
type MultiBssidProfilesResponse struct {
	MultiBssidProfiles MultiBssidProfiles `json:"Cisco-IOS-XE-wireless-rf-cfg:multi-bssid-profiles"`
}

// AtfPoliciesResponse represents the ATF policies configuration response
type AtfPoliciesResponse struct {
	AtfPolicies AtfPolicies `json:"Cisco-IOS-XE-wireless-rf-cfg:atf-policies"`
}

// RfTagsResponse represents the RF tags configuration response
type RfTagsResponse struct {
	RfTags RfTags `json:"Cisco-IOS-XE-wireless-rf-cfg:rf-tags"`
}

// RfProfilesResponse represents the RF profiles configuration response
type RfProfilesResponse struct {
	RfProfiles RfProfiles `json:"Cisco-IOS-XE-wireless-rf-cfg:rf-profiles"`
}

// RfProfileDefaultEntriesResponse represents the RF profile default entries response
type RfProfileDefaultEntriesResponse struct {
	RfProfileDefaultEntries RfProfileDefaultEntries `json:"Cisco-IOS-XE-wireless-rf-cfg:rf-profile-default-entries"`
}

// MultiBssidProfiles contains multi-BSSID profile configuration entries
type MultiBssidProfiles struct {
	MultiBssidProfile []struct {
		ProfileName string `json:"profile-name"` // Multi-BSSID profile name
		Description string `json:"description"`  // Profile description
	} `json:"multi-bssid-profile"`
}

type AtfPolicies struct {
	AtfPolicy []AtfPolicy `json:"atf-policy"`
}

type AtfPolicy []struct {
	PolicyID      int    `json:"policy-id"`
	AtfpolicyName string `json:"atfpolicy-name"`
}

type RfTags struct {
	RfTag []RfTag `json:"rf-tag"`
}

type RfTag struct {
	TagName             string `json:"tag-name"`
	Dot11ARfProfileName string `json:"dot11a-rf-profile-name,omitempty"`
	Dot11BRfProfileName string `json:"dot11b-rf-profile-name,omitempty"`
	Dot116GhzRfProfName string `json:"dot11-6ghz-rf-prof-name,omitempty"`
	RfTagRadioProfiles  struct {
		RfTagRadioProfile []struct {
			SlotID string `json:"slot-id"`
			BandID string `json:"band-id"`
		} `json:"rf-tag-radio-profile"`
	} `json:"rf-tag-radio-profiles"`
	Description string `json:"description,omitempty"`
}

type RfProfiles struct {
	RfProfile []RfProfile `json:"rf-profile"`
}

type RfProfile struct {
	Name           string `json:"name"`
	Description    string `json:"description"`
	Status         bool   `json:"status"`
	Band           string `json:"band"`
	DataRate6M     string `json:"data-rate-6m"`
	DataRate12M    string `json:"data-rate-12m"`
	DataRate24M    string `json:"data-rate-24m"`
	RxSenSopCustom int    `json:"rx-sen-sop-custom,omitempty"`
	RfMcsEntries   struct {
		RfMcsEntry []struct {
			RfIndex int `json:"rf-index"`
		} `json:"rf-mcs-entry"`
	} `json:"rf-mcs-entries,omitempty"`
	RfdcaRemovedChannels struct {
		RfdcaRemovedChannel []struct {
			Channel int `json:"channel"`
		} `json:"rfdca-removed-channel"`
	} `json:"rfdca-removed-channels,omitempty"`
	RfDcaChanWidth                   string `json:"rf-dca-chan-width,omitempty"`
	TxPowerV1Threshold               int    `json:"tx-power-v1-threshold,omitempty"`
	CoverageDataPacketRssiThreshold  int    `json:"coverage-data-packet-rssi-threshold,omitempty"`
	MinNumClients                    int    `json:"min-num-clients,omitempty"`
	CoverageVoicePacketRssiThreshold int    `json:"coverage-voice-packet-rssi-threshold,omitempty"`
	RxSenSopThreshold                string `json:"rx-sen-sop-threshold,omitempty"`
	TxPowerMin                       int    `json:"tx-power-min,omitempty"`
	DataRate9M                       string `json:"data-rate-9m,omitempty"`
	DataRate1M                       string `json:"data-rate-1m,omitempty"`
	DataRate2M                       string `json:"data-rate-2m,omitempty"`
	DataRate55M                      string `json:"data-rate-5-5m,omitempty"`
	DataRate11M                      string `json:"data-rate-11m,omitempty"`
}

type RfProfileDefaultEntries struct {
	RfProfileDefaultEntry []RfProfileDefaultEntry `json:"rf-profile-default-entry"`
}

type RfProfileDefaultEntry struct {
	Band                    string `json:"band"`
	Name                    string `json:"name"`
	Description             string `json:"description"`
	DataRate6M              string `json:"data-rate-6m"`
	DataRate12M             string `json:"data-rate-12m"`
	DataRate24M             string `json:"data-rate-24m"`
	BandSelectProbeResponse bool   `json:"band-select-probe-response,omitempty"`
	RfMcsDefaultEntries     struct {
		RfMcsDefaultEntry []struct {
			RfIndex     int `json:"rf-index"`
			McsDataRate int `json:"mcs-data-rate"`
		} `json:"rf-mcs-default-entry"`
	} `json:"rf-mcs-default-entries"`
}

func GetRfCfg(client *wnc.Client, ctx context.Context) (*RfCfgResponse, error) {
	var data RfCfgResponse
	if err := client.SendAPIRequest(ctx, RfCfgEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

func GetRfMultiBssidProfiles(client *wnc.Client, ctx context.Context) (*MultiBssidProfilesResponse, error) {
	var data MultiBssidProfilesResponse
	if err := client.SendAPIRequest(ctx, MultiBssidProfilesEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

func GetRfAtfPolicies(client *wnc.Client, ctx context.Context) (*AtfPoliciesResponse, error) {
	var data AtfPoliciesResponse
	if err := client.SendAPIRequest(ctx, AtfPoliciesEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

func GetRfTags(client *wnc.Client, ctx context.Context) (*RfTagsResponse, error) {
	var data RfTagsResponse
	if err := client.SendAPIRequest(ctx, RfTagsEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

func GetRfProfiles(client *wnc.Client, ctx context.Context) (*RfProfilesResponse, error) {
	var data RfProfilesResponse
	if err := client.SendAPIRequest(ctx, RfProfilesEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

func GetRfProfileDefaultEntries(client *wnc.Client, ctx context.Context) (*RfProfileDefaultEntriesResponse, error) {
	var data RfProfileDefaultEntriesResponse
	if err := client.SendAPIRequest(ctx, RfProfileDefaultEntriesEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}
