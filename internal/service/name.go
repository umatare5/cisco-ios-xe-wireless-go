package service

import (
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/validation"
)

// RequireAPName rejects a blank access point name with core.ErrResourceNotFound, the class
// RequireMACAddress uses for a blank address.
//
// Nothing but the error is returned: the controller keys ap-name-mac-map and the RPC name arms by
// the name as it is typed, so there is no normalized form to hand back.
func RequireAPName(apName string) error {
	if !validation.IsNonEmptyString(apName) {
		return core.ErrResourceNotFound
	}

	return nil
}
