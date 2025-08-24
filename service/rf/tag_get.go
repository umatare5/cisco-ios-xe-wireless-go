package rf

import (
	"context"
	"encoding/json"

	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/rf"
)

// GetRFTag retrieves an RF tag configuration.
// This function retrieves a specific RF tag configuration from the wireless controller.
//
// **Parameters:**
//   - ctx: Context for the request
//   - tagName: Name of the RF tag to retrieve
//
// **Returns:**
//   - *RFTagConfig: RF tag configuration if found, nil otherwise
//   - error: nil on success, error otherwise
//
// **YANG Path:** /Cisco-IOS-XE-wireless-rf-cfg:rf-cfg-data/rf-tags/rf-tag={tag-name}
func (s *RFTagService) GetRFTag(ctx context.Context, tagName string) (*RFTagConfig, error) {
	if err := s.ValidateClient(); err != nil {
		return nil, err
	}

	// Get raw data from the helper
	body, err := s.tagOps.Get(ctx, tagName)
	if err != nil {
		return nil, err
	}

	if len(body) == 0 {
		return nil, nil
	}

	// Direct parsing approach for RF tag response (array with single object)
	var response struct {
		RfTag []model.RfTag `json:"Cisco-IOS-XE-wireless-rf-cfg:rf-tag"`
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	// Check if the tag was found
	if len(response.RfTag) == 0 || response.RfTag[0].TagName == "" {
		return nil, nil // Not found
	}

	// Convert internal model to service model
	internalTag := response.RfTag[0]

	// Convert internal model to service model
	config := &RFTagConfig{
		TagName:             internalTag.TagName,
		Description:         internalTag.Description,
		Dot11ARfProfileName: internalTag.Dot11ARfProfileName,
		Dot11BRfProfileName: internalTag.Dot11BRfProfileName,
		Dot116GhzRfProfName: internalTag.Dot116GhzRfProfName,
	}

	// Convert radio profiles if present
	if internalTag.RfTagRadioProfiles != nil {
		radioProfiles := make(
			[]RFTagRadioProfile,
			len(internalTag.RfTagRadioProfiles.RfTagRadioProfile),
		)
		for i, profile := range internalTag.RfTagRadioProfiles.RfTagRadioProfile {
			radioProfiles[i] = RFTagRadioProfile{
				SlotID: profile.SlotID,
				BandID: profile.BandID,
			}
		}
		config.RfTagRadioProfiles = &RFTagRadioProfiles{
			RfTagRadioProfile: radioProfiles,
		}
	}

	return config, nil
}
