// Package ap provides access point configuration management functionality for the Cisco Wireless Network Controller API.
package ap

import (
	"context"

	wnc "github.com/umatare5/cisco-xe-wireless-restconf-go"
)

const (
	// ApCfgBasePath defines the base path for access point configuration endpoints
	ApCfgBasePath = "/restconf/data/Cisco-IOS-XE-wireless-ap-cfg:ap-cfg-data"
	// ApCfgEndpoint retrieves complete access point configuration data
	ApCfgEndpoint = ApCfgBasePath
	// TagSourcePriorityConfigsEndpoint retrieves tag source priority configurations
	TagSourcePriorityConfigsEndpoint = ApCfgBasePath + "/tag-source-priority-configs"
	// ApTagsEndpoint retrieves access point tag configurations
	ApTagsEndpoint = ApCfgBasePath + "/ap-tags"
)

// ApCfgResponse represents the complete access point configuration response
type ApCfgResponse struct {
	CiscoIOSXEWirelessApCfgApCfgData struct {
		TagSourcePriorityConfigs TagSourcePriorityConfigs `json:"tag-source-priority-configs"`
		ApTags                   ApTags                   `json:"ap-tags"`
	} `json:"Cisco-IOS-XE-wireless-ap-cfg:ap-cfg-data"`
}

// ApCfgTagSourcePriorityConfigsResponse represents tag source priority configurations response
type ApCfgTagSourcePriorityConfigsResponse struct {
	TagSourcePriorityConfigs TagSourcePriorityConfigs `json:"Cisco-IOS-XE-wireless-ap-cfg:tag-source-priority-configs"`
}

// ApCfgApTagsResponse represents access point tags configuration response
type ApCfgApTagsResponse struct {
	ApTags ApTags `json:"Cisco-IOS-XE-wireless-ap-cfg:ap-tags"`
}

// TagSourcePriorityConfigs contains tag source priority configuration settings
type TagSourcePriorityConfigs struct {
	TagSourcePriorityConfig []struct {
		Priority int    `json:"priority"` // Priority level for tag source
		TagSrc   string `json:"tag-src"`  // Tag source identifier
	} `json:"tag-source-priority-config"`
}

// ApTags contains access point tag configuration data
type ApTags struct {
	ApTag []ApTag `json:"ap-tag"`
}

// ApTag represents tag assignments for a specific access point
type ApTag struct {
	ApMac     string `json:"ap-mac"`           // Access point MAC address
	PolicyTag string `json:"policy-tag"`       // Policy tag assigned to the AP
	SiteTag   string `json:"site-tag"`         // Site tag assigned to the AP
	RfTag     string `json:"rf-tag,omitempty"` // RF tag assigned to the AP (optional)
}

// GetApCfg retrieves complete access point configuration data.
func GetApCfg(client *wnc.Client, ctx context.Context) (*ApCfgResponse, error) {
	var data ApCfgResponse
	if err := client.SendAPIRequest(ctx, ApCfgEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetTagSourcePriorityConfigs retrieves tag source priority configurations.
func GetTagSourcePriorityConfigs(client *wnc.Client, ctx context.Context) (*TagSourcePriorityConfigs, error) {
	var data TagSourcePriorityConfigs
	if err := client.SendAPIRequest(ctx, TagSourcePriorityConfigsEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetApTagSourcePriorityConfigs retrieves tag source priority configurations with full response wrapper.
func GetApTagSourcePriorityConfigs(client *wnc.Client, ctx context.Context) (*ApCfgTagSourcePriorityConfigsResponse, error) {
	var data ApCfgTagSourcePriorityConfigsResponse
	if err := client.SendAPIRequest(ctx, TagSourcePriorityConfigsEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetApApTags retrieves access point tag configurations.
func GetApApTags(client *wnc.Client, ctx context.Context) (*ApCfgApTagsResponse, error) {
	var data ApCfgApTagsResponse
	if err := client.SendAPIRequest(ctx, ApTagsEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}
