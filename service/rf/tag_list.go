package rf

import (
	"context"
	"encoding/json"

	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/rf"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
)

// ListRFTags retrieves all RF tag configurations.
// This function retrieves all RF tag configurations from the wireless controller.
//
// **Parameters:**
//   - ctx: Context for the request
//
// **Returns:**
//   - *RFTagCfgResponse: RF tag configurations response
//   - error: nil on success, error otherwise
//
// **YANG Path:** /Cisco-IOS-XE-wireless-rf-cfg:rf-cfg-data/rf-tags
func (s *RFTagService) ListRFTags(ctx context.Context) (*RFTagCfgResponse, error) {
	if err := s.ValidateClient(); err != nil {
		return nil, err
	}

	// Get raw data using REST client
	body, err := s.Client().Do(ctx, "GET", routes.RfTagsEndpoint)
	if err != nil {
		return nil, err
	}

	if len(body) == 0 {
		return &RFTagCfgResponse{}, nil
	}

	// Parse response into internal model first
	var internalResp model.RfTags
	if err := json.Unmarshal(body, &internalResp); err != nil {
		return nil, err
	}

	// Convert to service model
	response := &RFTagCfgResponse{}
	if len(internalResp.RfTagList) > 0 {
		tags := make([]RFTagConfig, len(internalResp.RfTagList))
		for i, tag := range internalResp.RfTagList {
			config := RFTagConfig{
				TagName:             tag.TagName,
				Description:         tag.Description,
				Dot11ARfProfileName: tag.Dot11ARfProfileName,
				Dot11BRfProfileName: tag.Dot11BRfProfileName,
				Dot116GhzRfProfName: tag.Dot116GhzRfProfName,
			}

			// Convert radio profiles if present
			if tag.RfTagRadioProfiles != nil {
				radioProfiles := make(
					[]RFTagRadioProfile,
					len(tag.RfTagRadioProfiles.RfTagRadioProfile),
				)
				for j, profile := range tag.RfTagRadioProfiles.RfTagRadioProfile {
					radioProfiles[j] = RFTagRadioProfile{
						SlotID: profile.SlotID,
						BandID: profile.BandID,
					}
				}
				config.RfTagRadioProfiles = &RFTagRadioProfiles{
					RfTagRadioProfile: radioProfiles,
				}
			}

			tags[i] = config
		}

		response.RfTags = &RFTagsCfgData{
			RfTag: tags,
		}
	}

	return response, nil
}
