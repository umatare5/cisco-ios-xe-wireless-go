package wlan

import (
	"errors"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/helpers"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/service"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/validation"
)

// PolicyTagConfig represents a policy tag configuration
type PolicyTagConfig struct {
	TagName      string        `json:"tag-name"`
	Description  string        `json:"description,omitempty"`
	WLANPolicies *WLANPolicies `json:"wlan-policies,omitempty"`
}

// GetTagName returns the tag name (implements testutil.TagConfig)
func (c *PolicyTagConfig) GetTagName() string {
	return c.TagName
}

// SetTagName sets the tag name (implements testutil.TagConfig)
func (c *PolicyTagConfig) SetTagName(name string) {
	c.TagName = name
}

// Validate validates the policy tag configuration (implements testutil.TagConfig)
func (c *PolicyTagConfig) Validate() error {
	if c.TagName == "" {
		return errors.New("policy tag name cannot be empty")
	}
	return nil
}

// WLANPolicies represents the container for WLAN policy mappings
type WLANPolicies struct {
	WLANPolicy []WLANPolicyMap `json:"wlan-policy,omitempty"`
}

// WLANPolicyMap represents a WLAN to policy profile mapping
type WLANPolicyMap struct {
	WLANProfileName   string `json:"wlan-profile-name"`
	PolicyProfileName string `json:"policy-profile-name"`
}

// PolicyTagService provides Policy Tag management operations
type PolicyTagService struct {
	service.BaseService
	tagOps *helpers.TagCRUDOperations
}

// NewPolicyTagService creates a new PolicyTagService instance
func NewPolicyTagService(c *core.Client) *PolicyTagService {
	config := helpers.TagCRUDConfig{
		BasePath:           PolicyTagBasePath + "/policy-list-entry",
		ListPath:           PolicyTagBasePath,
		YANGPrefix:         "Cisco-IOS-XE-wireless-wlan-cfg:policy-list-entry",
		TagNameField:       "tag-name",
		ValidatorFunc:      validation.IsValidPolicyTag,
		ValidationErrorKey: "policy",
	}

	return &PolicyTagService{
		BaseService: service.NewBaseService(c),
		tagOps:      helpers.NewTagCRUDOperations(config, c),
	}
}

const (
	// PolicyTagBasePath is the base path for policy tag operations
	PolicyTagBasePath = restconf.YANGModelPrefix + "wlan-cfg:wlan-cfg-data/policy-list-entries"
)
