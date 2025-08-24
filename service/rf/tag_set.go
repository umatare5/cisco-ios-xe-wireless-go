package rf

import (
	"context"
	"fmt"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/rf"
)

// SetRFTag sets/updates an existing RF tag configuration.
// This function modifies an existing RF tag configuration.
//
// **Parameters:**
//   - ctx: Context for the request
//   - config: Updated RF tag configuration
//
// **Returns:**
//   - error: nil on success, error otherwise
func (s *RFTagService) SetRFTag(ctx context.Context, config RFTagConfig) error {
	if err := s.ValidateClient(); err != nil {
		return err
	}

	// Convert service model to internal model
	internalConfig := model.RfTag{
		TagName:             config.TagName,
		Description:         config.Description,
		Dot11ARfProfileName: config.Dot11ARfProfileName,
		Dot11BRfProfileName: config.Dot11BRfProfileName,
		Dot116GhzRfProfName: config.Dot116GhzRfProfName,
	}

	// Convert radio profiles if provided
	if config.RfTagRadioProfiles != nil {
		internalRadioProfiles := make(
			[]model.RfTagRadioProfile,
			len(config.RfTagRadioProfiles.RfTagRadioProfile),
		)
		for i, profile := range config.RfTagRadioProfiles.RfTagRadioProfile {
			internalRadioProfiles[i] = model.RfTagRadioProfile{
				SlotID: profile.SlotID,
				BandID: profile.BandID,
			}
		}
		internalConfig.RfTagRadioProfiles = &model.RfTagRadioProfiles{
			RfTagRadioProfile: internalRadioProfiles,
		}
	}

	return s.tagOps.Update(ctx, internalConfig, config.TagName)
}

// SetDot11ARfProfile sets the 5GHz RF profile for an RF tag.
//
// **Parameters:**
//   - ctx: Context for the request
//   - tagName: Name of the RF tag
//   - rfProfileName: 5GHz RF profile name to set
//
// **Returns:**
//   - error: nil on success, error otherwise
func (s *RFTagService) SetDot11ARfProfile(
	ctx context.Context,
	tagName, rfProfileName string,
) error {
	// Get existing RF tag
	config, err := s.GetRFTag(ctx, tagName)
	if err != nil {
		return fmt.Errorf("failed to get RF tag: %w", err)
	}

	if config == nil {
		return fmt.Errorf("RF tag %s not found", tagName)
	}

	// Update 5GHz RF profile
	config.Dot11ARfProfileName = rfProfileName
	return s.SetRFTag(ctx, *config)
}

// SetDot11BRfProfile sets the 2.4GHz RF profile for an RF tag.
//
// **Parameters:**
//   - ctx: Context for the request
//   - tagName: Name of the RF tag
//   - rfProfileName: 2.4GHz RF profile name to set
//
// **Returns:**
//   - error: nil on success, error otherwise
func (s *RFTagService) SetDot11BRfProfile(
	ctx context.Context,
	tagName, rfProfileName string,
) error {
	// Get existing RF tag
	config, err := s.GetRFTag(ctx, tagName)
	if err != nil {
		return fmt.Errorf("failed to get RF tag: %w", err)
	}

	if config == nil {
		return fmt.Errorf("RF tag %s not found", tagName)
	}

	// Update 2.4GHz RF profile
	config.Dot11BRfProfileName = rfProfileName
	return s.SetRFTag(ctx, *config)
}

// SetDot116GhzRfProfile sets the 6GHz RF profile for an RF tag.
//
// **Parameters:**
//   - ctx: Context for the request
//   - tagName: Name of the RF tag
//   - rfProfileName: 6GHz RF profile name to set
//
// **Returns:**
//   - error: nil on success, error otherwise
func (s *RFTagService) SetDot116GhzRfProfile(
	ctx context.Context,
	tagName, rfProfileName string,
) error {
	// Get existing RF tag
	config, err := s.GetRFTag(ctx, tagName)
	if err != nil {
		return fmt.Errorf("failed to get RF tag: %w", err)
	}

	if config == nil {
		return fmt.Errorf("RF tag %s not found", tagName)
	}

	// Update 6GHz RF profile
	config.Dot116GhzRfProfName = rfProfileName
	return s.SetRFTag(ctx, *config)
}

// SetDescription sets the description for an RF tag.
//
// **Parameters:**
//   - ctx: Context for the request
//   - tagName: Name of the RF tag
//   - description: Description to set
//
// **Returns:**
//   - error: nil on success, error otherwise
func (s *RFTagService) SetDescription(ctx context.Context, tagName, description string) error {
	// Get existing RF tag
	config, err := s.GetRFTag(ctx, tagName)
	if err != nil {
		return fmt.Errorf("failed to get RF tag: %w", err)
	}

	if config == nil {
		return fmt.Errorf("RF tag %s not found", tagName)
	}

	// Update description
	config.Description = description
	return s.SetRFTag(ctx, *config)
}

// ConfigureRFTag provides a complete configuration interface for RF tags.
// This function either creates a new RF tag or updates an existing one.
//
// **Parameters:**
//   - ctx: Context for the request
//   - config: Complete RF tag configuration
//
// **Returns:**
//   - error: nil on success, error otherwise
func (s *RFTagService) ConfigureRFTag(ctx context.Context, config RFTagConfig) error {
	// Try to set first, if fails with 404, create new
	err := s.SetRFTag(ctx, config)
	if err != nil {
		// Check if error is 404 (not found)
		if core.IsNotFoundError(err) {
			return s.CreateRFTag(ctx, config)
		}
		return err
	}

	return nil
}
