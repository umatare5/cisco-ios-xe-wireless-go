package site

import (
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/helpers"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/service"
)

// SiteTagConfig represents a site tag configuration
type SiteTagConfig struct {
	SiteTagName   string `json:"site-tag-name"`
	Description   string `json:"description,omitempty"`
	FlexProfile   string `json:"flex-profile,omitempty"`
	APJoinProfile string `json:"ap-join-profile,omitempty"`
	IsLocalSite   *bool  `json:"is-local-site,omitempty"`
}

// SiteTagService provides site tag management operations
type SiteTagService struct {
	service.BaseService
	tagOps *helpers.TagCRUDOperations
}

// NewSiteTagService creates a new site tag service
func NewSiteTagService(client *core.Client) *SiteTagService {
	config := helpers.TagCRUDConfig{
		BasePath:           routes.SiteTagConfigsEndpoint + "/site-tag-config",
		ListPath:           routes.SiteTagConfigsEndpoint,
		YANGPrefix:         "Cisco-IOS-XE-wireless-site-cfg:site-tag-config",
		ValidationErrorKey: "site-tag-name",
		ValidatorFunc:      nil, // No specific validation for site tags
	}

	return &SiteTagService{
		BaseService: service.NewBaseService(client),
		tagOps:      helpers.NewTagCRUDOperations(config, client),
	}
}

const (
	// SiteTagBasePath is the base path for site tag operations
	SiteTagBasePath = routes.SiteTagConfigsEndpoint
)
