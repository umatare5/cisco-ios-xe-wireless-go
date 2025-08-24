package rf

import (
	"context"
)

// DeleteRFTag deletes an RF tag configuration.
// This function removes an RF tag configuration from the wireless controller.
//
// **Parameters:**
//   - ctx: Context for the request
//   - tagName: Name of the RF tag to delete
//
// **Returns:**
//   - error: nil on success, error otherwise
//
// **YANG Path:** /Cisco-IOS-XE-wireless-rf-cfg:rf-cfg-data/rf-tags/rf-tag={tag-name}
func (s *RFTagService) DeleteRFTag(ctx context.Context, tagName string) error {
	if err := s.ValidateClient(); err != nil {
		return err
	}

	return s.tagOps.Delete(ctx, tagName)
}
