package rf

import (
	"context"

	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/rf"
)

// CreateRFTag creates a new RF tag configuration.
// This function creates a new RF tag configuration on the wireless controller.
//
// **Parameters:**
//   - ctx: Context for the request
//   - config: RF tag configuration to create
//
// **Returns:**
//   - error: nil on success, error otherwise
//
// **YANG Path:** /Cisco-IOS-XE-wireless-rf-cfg:rf-cfg-data/rf-tags
func (s *RFTagService) CreateRFTag(ctx context.Context, config RFTagConfig) error {
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

	return s.tagOps.Create(ctx, internalConfig, config.TagName)
}
