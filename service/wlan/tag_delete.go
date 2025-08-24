package wlan

import (
	"context"
)

// DeletePolicyTag deletes a policy tag configuration.
// This function removes a policy tag. The tag must not be associated with any APs.
//
// **Parameters:**
//   - ctx: Context for the request
//   - tagName: Name of the policy tag to delete
//
// **Returns:**
//   - error: nil on success, error otherwise
func (s *PolicyTagService) DeletePolicyTag(ctx context.Context, tagName string) error {
	if err := s.ValidateClient(); err != nil {
		return err
	}

	return s.tagOps.Delete(ctx, tagName)
}
