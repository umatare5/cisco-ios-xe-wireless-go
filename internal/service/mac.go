package service

import (
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/validation"
)

// RequireMACAddress validates a MAC address and returns it in the colon-separated form
// the controller keys its lists by, so a dotted or dashed address reaches the same
// record rather than a 404.
//
// An empty address is core.ErrResourceNotFound: the caller asked for a record that
// cannot exist. A malformed one carries the validation error.
func RequireMACAddress(mac string) (string, error) {
	if !validation.IsNonEmptyString(mac) {
		return "", core.ErrResourceNotFound
	}

	return validation.NormalizeMACAddress(mac)
}
