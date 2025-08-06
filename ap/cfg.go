// Package ap provides access point configuration management functionality for the Cisco Wireless Network Controller API.
package ap

import (
	"context"
	"errors"

	wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
)

const (
	// ApCfgBasePath defines the base path for access point configuration endpoints
	ApCfgBasePath = "Cisco-IOS-XE-wireless-ap-cfg:ap-cfg-data"
	// ApCfgEndpoint retrieves complete access point configuration data
	ApCfgEndpoint = ApCfgBasePath
	// TagSourcePriorityConfigsEndpoint retrieves tag source priority configurations
	TagSourcePriorityConfigsEndpoint = ApCfgBasePath + "/tag-source-priority-configs"
	// ApTagsEndpoint retrieves access point tag configurations
	ApTagsEndpoint = ApCfgBasePath + "/ap-tags"
)

// ApCfgResponse represents the complete access point configuration response
type ApCfgResponse = model.ApCfgResponse

// ApCfgTagSourcePriorityConfigsResponse represents tag source priority configurations response
type ApCfgTagSourcePriorityConfigsResponse = model.ApCfgTagSourcePriorityConfigsResponse

// ApCfgApTagsResponse represents access point tags configuration response
type ApCfgApTagsResponse = model.ApCfgApTagsResponse

// TagSourcePriorityConfigs contains tag source priority configuration settings
type TagSourcePriorityConfigs = model.TagSourcePriorityConfigs

// ApTags contains access point tag configuration data
type ApTags = model.ApTags

// ApTag represents tag assignments for a specific access point
type ApTag = model.ApTag

// GetApCfg retrieves complete access point configuration data.
//
// Deprecated: Use ap.NewService(client.CoreClient()).Cfg(ctx) instead.
func GetApCfg(client *wnc.Client, ctx context.Context) (*ApCfgResponse, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	service := NewService(client.CoreClient())
	return service.Cfg(ctx)
}

// GetTagSourcePriorityConfigs retrieves tag source priority configurations.
//
// Deprecated: Use ap.NewService(client.CoreClient()).TagSourcePriorityConfigs(ctx) instead.
func GetTagSourcePriorityConfigs(client *wnc.Client, ctx context.Context) (*TagSourcePriorityConfigs, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	service := NewService(client.CoreClient())
	return service.TagSourcePriorityConfigs(ctx)
}

// GetApTagSourcePriorityConfigs retrieves tag source priority configurations with full response wrapper.
//
// Deprecated: Use ap.NewService(client.CoreClient()).ApTags(ctx) instead.
func GetApTagSourcePriorityConfigs(client *wnc.Client, ctx context.Context) (*ApCfgTagSourcePriorityConfigsResponse, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	var data ApCfgTagSourcePriorityConfigsResponse
	if err := client.SendAPIRequest(ctx, TagSourcePriorityConfigsEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetApApTags retrieves access point tag configurations.
//
// Deprecated: Use ap.NewService(client.CoreClient()).ApTags(ctx) instead.
func GetApApTags(client *wnc.Client, ctx context.Context) (*ApCfgApTagsResponse, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	service := NewService(client.CoreClient())
	return service.ApTags(ctx)
}
