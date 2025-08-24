package site

import (
	"context"
)

// DeleteSiteTag deletes a site tag configuration.
// This function removes a site tag. The tag must not be associated with any APs.
//
// **Parameters:**
//   - ctx: Context for the request
//   - siteTagName: Name of the site tag to delete
//
// **Returns:**
//   - error: nil on success, error otherwise
func (s *SiteTagService) DeleteSiteTag(ctx context.Context, siteTagName string) error {
	if err := s.ValidateClient(); err != nil {
		return err
	}

	return s.tagOps.Delete(ctx, siteTagName)
}
