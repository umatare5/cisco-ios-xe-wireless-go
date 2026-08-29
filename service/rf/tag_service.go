package rf

import (
	"context"
	"errors"
	"fmt"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/service"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/validation"
)

// RFTagService provides RF tag management functionality.
type RFTagService struct {
	service.BaseService
}

// NewRFTagService creates a new RF tag service.
func NewRFTagService(client *core.Client) *RFTagService {
	return &RFTagService{
		BaseService: service.NewBaseService(client),
	}
}

// GetConfig retrieves the RF configuration.
func (s *RFTagService) GetConfig(ctx context.Context, opts ...core.GetOption) (*CiscoIOSXEWirelessRFCfg, error) {
	return core.Get[CiscoIOSXEWirelessRFCfg](ctx, s.Client(), routes.RFCfgPath, opts...)
}

// GetRFTag retrieves an RF tag configuration by name.
func (s *RFTagService) GetRFTag(ctx context.Context, tagName string, opts ...core.GetOption) (*RFTag, error) {
	if err := validation.ValidateTagName(tagName); err != nil {
		return nil, err
	}

	result, err := core.Get[CiscoIOSXEWirelessRFCfgRFTagPayload](ctx, s.Client(), s.buildTagURL(tagName), opts...)
	if err != nil {
		return nil, err
	}
	if result == nil || result.RFTags == nil || len(result.RFTags) == 0 {
		return nil, nil
	}

	return &result.RFTags[0], nil
}

// ListRFTags retrieves all RF tag configurations.
func (s *RFTagService) ListRFTags(ctx context.Context, opts ...core.GetOption) ([]RFTag, error) {
	rfService := NewService(s.Client())

	result, err := rfService.ListRFTags(ctx, opts...)
	if err != nil {
		return nil, err
	}
	if result == nil || result.RFTags == nil {
		return nil, nil
	}
	if len(result.RFTags.RFTags) == 0 {
		return []RFTag{}, nil
	}

	return result.RFTags.RFTags, nil
}

// CreateRFTag creates a new RF tag configuration.
func (s *RFTagService) CreateRFTag(ctx context.Context, config *RFTag) error {
	if config == nil {
		return errors.New("RF tag config cannot be nil")
	}
	if err := validation.ValidateTagName(config.TagName); err != nil {
		return err
	}

	payload := s.buildPayload(config)
	return core.PostVoid(ctx, s.Client(), routes.RFTagsPath, payload)
}

// DeleteRFTag deletes an RF tag configuration.
func (s *RFTagService) DeleteRFTag(ctx context.Context, tagName string) error {
	if err := validation.ValidateTagName(tagName); err != nil {
		return err
	}
	return core.Delete(ctx, s.Client(), s.buildTagURL(tagName))
}

// SetDot11ARfProfile sets the 5 GHz RF profile for an RF tag, the dot11a-rf-profile-name leaf.
func (s *RFTagService) SetDot11ARfProfile(ctx context.Context, tagName, rfProfileName string) error {
	return s.updateTagField(ctx, tagName, func(payload *RFTag) {
		if payload != nil {
			payload.Dot11ARfProfileName = &rfProfileName
		}
	})
}

// SetDot11BRfProfile sets the 2.4 GHz RF profile for an RF tag, the dot11b-rf-profile-name leaf.
func (s *RFTagService) SetDot11BRfProfile(ctx context.Context, tagName, rfProfileName string) error {
	return s.updateTagField(ctx, tagName, func(payload *RFTag) {
		if payload != nil {
			payload.Dot11BRfProfileName = &rfProfileName
		}
	})
}

// SetDot116GhzRFProfile sets the 6 GHz RF profile for an RF tag, the dot11-6ghz-rf-prof-name leaf.
func (s *RFTagService) SetDot116GhzRFProfile(ctx context.Context, tagName, rfProfileName string) error {
	return s.updateTagField(ctx, tagName, func(payload *RFTag) {
		if payload != nil {
			payload.Dot116GhzRFProfName = &rfProfileName
		}
	})
}

// SetDescription sets the description for an RF tag.
func (s *RFTagService) SetDescription(ctx context.Context, tagName, description string) error {
	return s.updateTagField(ctx, tagName, func(payload *RFTag) {
		if payload != nil {
			payload.Description = &description
		}
	})
}

// updateTagField updates a specific field of an RF tag using the provided update function.
func (s *RFTagService) updateTagField(ctx context.Context, tagName string,
	updateFunc func(*RFTag),
) error {
	if updateFunc == nil {
		return errors.New("update function cannot be nil")
	}

	tag, err := s.GetRFTag(ctx, tagName)
	if err != nil {
		return fmt.Errorf("RF tag operation failed: %w",
			fmt.Errorf("tag retrieval failed for '%s': %w", tagName, err))
	}
	if tag == nil {
		return fmt.Errorf("RF tag operation failed: %w",
			fmt.Errorf("tag '%s' not found in controller configuration", tagName))
	}

	updateFunc(tag)
	return s.setRFTag(ctx, tag)
}

// setRFTag sets/updates an existing RF tag configuration.
func (s *RFTagService) setRFTag(ctx context.Context, config *RFTag) error {
	if config == nil {
		return errors.New("RF tag config cannot be nil")
	}
	if err := validation.ValidateTagName(config.TagName); err != nil {
		return err
	}

	// A merge PATCH, as the site and policy tag services already use. Measured on 17.12.8 as a
	// paired contrast on one probe tag: PATCH left a top-level leaf and a nested-list non-key leaf
	// at the values that had been set, while a PUT with identical omissions demoted both to their
	// schema defaults. A node this struct cannot represent needs that — 17.18 serves ap-beam-state
	// and a urwb container on every rf-tag and RFTag declares neither, so the replacing PUT demoted
	// them on every write. The contrast has not been repeated on 17.15.6 or 17.18.4a.
	payload := s.buildPayload(config)
	return core.PatchVoid(ctx, s.Client(), s.buildTagURL(config.TagName), payload)
}

// buildTagURL builds URL for specific tag operations using RESTCONF builder.
func (s *RFTagService) buildTagURL(tagName string) string {
	return s.Client().RESTCONFBuilder().BuildQueryURL(routes.RFTagByNamePath, tagName)
}

// buildPayload builds a payload for tag operations using the request.
func (s *RFTagService) buildPayload(config *RFTag) CiscoIOSXEWirelessRFCfgRFTagsPayload {
	return CiscoIOSXEWirelessRFCfgRFTagsPayload{
		RFTag: *config,
	}
}
