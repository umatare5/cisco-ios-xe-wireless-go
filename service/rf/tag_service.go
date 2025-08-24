package rf

import (
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/helpers"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/service"
)

// RFTagConfig represents an RF tag configuration
type RFTagConfig struct {
	TagName             string              `json:"tag-name"`
	Description         string              `json:"description,omitempty"`
	Dot11ARfProfileName string              `json:"dot11a-rf-profile-name,omitempty"`
	Dot11BRfProfileName string              `json:"dot11b-rf-profile-name,omitempty"`
	Dot116GhzRfProfName string              `json:"dot11-6ghz-rf-prof-name,omitempty"`
	RfTagRadioProfiles  *RFTagRadioProfiles `json:"rf-tag-radio-profiles,omitempty"`
}

// RFTagRadioProfiles represents RF tag radio profiles
type RFTagRadioProfiles struct {
	RfTagRadioProfile []RFTagRadioProfile `json:"rf-tag-radio-profile"`
}

// RFTagRadioProfile represents an RF tag radio profile
type RFTagRadioProfile struct {
	SlotID string `json:"slot-id"`
	BandID string `json:"band-id"`
}

// RFTagCfgResponse represents the response structure for RF tag configuration data
type RFTagCfgResponse struct {
	RfTags *RFTagsCfgData `json:"Cisco-IOS-XE-wireless-rf-cfg:rf-tags,omitempty"`
	RfTag  []RFTagConfig  `json:"Cisco-IOS-XE-wireless-rf-cfg:rf-tag,omitempty"`
}

// RFTagsCfgData represents RF tags configuration data
type RFTagsCfgData struct {
	RfTag []RFTagConfig `json:"rf-tag"`
}

// RFTagService provides RF tag management operations
type RFTagService struct {
	service.BaseService
	tagOps *helpers.TagCRUDOperations
}

// NewRFTagService creates a new RF tag service
func NewRFTagService(client *core.Client) *RFTagService {
	config := helpers.TagCRUDConfig{
		BasePath:           routes.RfTagsEndpoint + "/rf-tag",
		ListPath:           routes.RfTagsEndpoint,
		YANGPrefix:         "rf-tag",
		ValidationErrorKey: "tag-name",
		ValidatorFunc:      nil, // No specific validation for RF tags
	}

	return &RFTagService{
		BaseService: service.NewBaseService(client),
		tagOps:      helpers.NewTagCRUDOperations(config, client),
	}
}
