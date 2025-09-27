package rf

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/service"
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
func (s *RFTagService) GetConfig(ctx context.Context) (*CiscoIOSXEWirelessRFCfg, error) {
	return core.Get[CiscoIOSXEWirelessRFCfg](ctx, s.Client(), routes.RFCfgPath)
}

// GetRFTag retrieves an RF tag configuration by name.
func (s *RFTagService) GetRFTag(ctx context.Context, tagName string) (*RFTag, error) {
	if err := s.validateTagName(tagName); err != nil {
		return nil, err
	}

	result, err := core.Get[CiscoIOSXEWirelessRFCfgRFTag](ctx, s.Client(), s.buildTagURL(tagName))
	if err != nil {
		return nil, err
	}

	if result == nil || result.RFTagList == nil || len(result.RFTagList) == 0 {
		return nil, nil
	}

	return &result.RFTagList[0], nil
}

// ListRFTags retrieves all RF tag configurations.
func (s *RFTagService) ListRFTags(ctx context.Context) ([]RFTag, error) {
	result, err := core.Get[CiscoIOSXEWirelessRFCfgRFTags](ctx, s.Client(), routes.RFTagsPath)
	if err != nil {
		return nil, err
	}

	if result == nil {
		return []RFTag{}, nil
	}

	if len(result.RFTags.RFTagList) == 0 {
		return []RFTag{}, nil
	}

	return result.RFTags.RFTagList, nil
}

// CreateRFTag creates a new RF tag configuration.
func (s *RFTagService) CreateRFTag(ctx context.Context, config *RFTag) error {
	if config == nil {
		return errors.New("RF tag config cannot be nil")
	}
	if config.TagName == "" {
		return errors.New("RF tag name cannot be empty")
	}

	if err := s.validateTagName(config.TagName); err != nil {
		return err
	}

	// Build payload directly from config
	payload := s.buildPayload(config)
	return core.PostVoid(ctx, s.Client(), routes.RFTagsPath, payload)
}

// DeleteRFTag deletes an RF tag configuration.
func (s *RFTagService) DeleteRFTag(ctx context.Context, tagName string) error {
	if err := s.validateTagName(tagName); err != nil {
		return err
	}
	return core.Delete(ctx, s.Client(), s.buildTagURL(tagName))
}

// SetDot11ARfProfile sets the 5GHz RF profile for an RF tag.
func (s *RFTagService) SetDot11ARfProfile(ctx context.Context, tagName, rfProfileName string) error {
	return s.updateTagField(ctx, tagName, func(payload *RFTag) {
		if payload != nil {
			payload.Dot11ARfProfileName = rfProfileName
		}
	})
}

// SetDot11BRfProfile sets the 2.4GHz RF profile for an RF tag.
func (s *RFTagService) SetDot11BRfProfile(ctx context.Context, tagName, rfProfileName string) error {
	return s.updateTagField(ctx, tagName, func(payload *RFTag) {
		if payload != nil {
			payload.Dot11BRfProfileName = rfProfileName
		}
	})
}

// SetDot116GhzRFProfile sets the 6GHz RF profile for an RF tag.
func (s *RFTagService) SetDot116GhzRFProfile(ctx context.Context, tagName, rfProfileName string) error {
	return s.updateTagField(ctx, tagName, func(payload *RFTag) {
		if payload != nil {
			payload.Dot116GhzRFProfName = rfProfileName
		}
	})
}

// SetDescription sets the description for an RF tag.
func (s *RFTagService) SetDescription(ctx context.Context, tagName, description string) error {
	return s.updateTagField(ctx, tagName, func(payload *RFTag) {
		if payload != nil {
			payload.Description = description
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
	if config.TagName == "" {
		return errors.New("RF tag name cannot be empty")
	}

	if err := s.validateTagName(config.TagName); err != nil {
		return err
	}

	// Build payload directly from config
	payload := s.buildPayload(config)
	return core.PutVoid(ctx, s.Client(), s.buildTagURL(config.TagName), payload)
}

// validateTagName validates RF tag name.
func (s *RFTagService) validateTagName(tagName string) error {
	if tagName == "" {
		return errors.New("RF tag name cannot be empty")
	}
	if strings.TrimSpace(tagName) == "" {
		return fmt.Errorf("RF tag validation failed: %w",
			fmt.Errorf("invalid tag name format: '%s'", tagName))
	}
	return nil
}

// buildTagURL builds URL for specific tag operations using RESTCONF builder.
func (s *RFTagService) buildTagURL(tagName string) string {
	return fmt.Sprintf("%s/rf-tag=%s", routes.RFTagsPath, tagName)
}

// buildPayload builds the payload for POST/PUT operations.
func (s *RFTagService) buildPayload(config *RFTag) map[string]any {
	return map[string]any{
		"Cisco-IOS-XE-wireless-rf-cfg:rf-tag": config,
	}
}
