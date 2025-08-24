// Package ap provides access point functionality for the Cisco IOS-XE Wireless Network Controller API.
package ap

import (
	"context"
	"fmt"
	"strconv"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/ap"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/validation"
)

// configOps provides high-level configuration operations for AP service
func (s Service) configOps() *core.ConfigOperations[model.ApCfg] {
	return core.NewConfigOperations[model.ApCfg](s.Client(), routes.APCfgBasePath)
}

const (

	// APCfgEndpoint retrieves complete access point configuration data
	APCfgEndpoint = routes.APCfgBasePath

	// TagSourcePriorityConfigsEndpoint retrieves tag source priority configurations
	TagSourcePriorityConfigsEndpoint = routes.APCfgBasePath + "/tag-source-priority-configs"

	// APTagsEndpoint retrieves access point tag configurations
	APTagsEndpoint = routes.APCfgBasePath + "/ap-tags"
)

// GetCfg retrieves the complete AP configuration data
func (s Service) GetCfg(ctx context.Context) (*model.ApCfg, error) {
	return s.configOps().GetAll(ctx)
}

// GetCfgApTagsOnly retrieves only AP tag configuration using fields parameter
func (s Service) GetCfgApTagsOnly(ctx context.Context) (*model.ApCfg, error) {
	return s.configOps().GetOnlyFields(ctx, "ap-tags")
}

// GetCfgTagSourcePriorityConfigsOnly retrieves only tag source priority configurations using fields parameter
func (s Service) GetCfgTagSourcePriorityConfigsOnly(ctx context.Context) (*model.ApCfg, error) {
	return s.configOps().GetOnlyFields(ctx, "tag-source-priority-configs")
}

// GetCfgApTagByMac retrieves AP tag configuration filtered by AP MAC address
func (s Service) GetCfgApTagByMac(ctx context.Context, apMacAddr string) (*model.ApCfgApTag, error) {
	if err := validation.ValidateAPMac(apMacAddr); err != nil {
		return nil, fmt.Errorf(ErrInvalidAPMacFormat, apMacAddr)
	}
	normalizedMAC := validation.NormalizeAPMac(apMacAddr)

	// Use configOps pattern for subresource access
	subOps := core.NewConfigOperations[model.ApCfgApTag](s.Client(), APTagsEndpoint)
	return subOps.GetByID(ctx, "ap-tag", normalizedMAC)
}

// GetApTagsCfg retrieves access point tag configurations
func (s Service) GetApTagsCfg(ctx context.Context) (*model.ApCfgApTags, error) {
	subOps := core.NewConfigOperations[model.ApCfgApTags](s.Client(), routes.APCfgBasePath)
	return subOps.GetSubRes(ctx, "ap-tags")
}

// GetCfgTagSourcePriorityByPriority retrieves tag source priority configuration filtered by priority
func (s Service) GetCfgTagSourcePriorityByPriority(
	ctx context.Context,
	priority int,
) (*model.ApCfgTagSourcePriorityConfigs, error) {
	subOps := core.NewConfigOperations[model.ApCfgTagSourcePriorityConfigs](
		s.Client(), TagSourcePriorityConfigsEndpoint,
	)
	return subOps.GetByID(ctx, "tag-source-priority", strconv.Itoa(priority))
}

// GetTagSourcePriorityCfg retrieves tag source priority configurations
func (s Service) GetTagSourcePriorityCfg(ctx context.Context) (*model.TagSourcePriorityConfigs, error) {
	subOps := core.NewConfigOperations[model.TagSourcePriorityConfigs](s.Client(), routes.APCfgBasePath)
	return subOps.GetSubRes(ctx, "tag-source-priority-configs")
}
