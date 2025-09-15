// Package ap provides AP-specific errors for the Cisco IOS-XE Wireless Network Controller API.
package ap

import (
	ierrors "github.com/umatare5/cisco-ios-xe-wireless-go/internal/errors"
)

// Common error messages for AP operations.
const (
	// ErrAPNameEmpty is the error message when AP name is empty or whitespace-only.
	ErrAPNameEmpty = "AP name cannot be empty"

	// ErrSiteTagEmpty is the error message when site tag is empty.
	ErrSiteTagEmpty = "site tag cannot be empty"

	// ErrPolicyTagEmpty is the error message when policy tag is empty.
	ErrPolicyTagEmpty = "policy tag cannot be empty"

	// ErrRFTagEmpty is the error message when RF tag is empty.
	ErrRFTagEmpty = "RF tag cannot be empty"

	// ErrAtLeastOneTagRequired is the error message when no tags are specified.
	ErrAtLeastOneTagRequired = ierrors.ErrAtLeastOneTagRequired

	// ErrRadioBandNegative is the error message when radio band is negative.
	ErrRadioBandNegative = "radio band cannot be negative"

	// ErrEitherMACOrNameRequired is the error message when neither MAC nor name is provided.
	ErrEitherMACOrNameRequired = "either AP MAC address or AP name must be provided"

	// ErrOnlyOneIdentifierAllowed is the error message when both MAC and name are provided.
	ErrOnlyOneIdentifierAllowed = "only one of AP MAC address or AP name should be provided"

	// ErrAPMacRequired is the error message when AP MAC address is required.
	ErrAPMacRequired = "AP MAC address is required"

	// ErrRadioBandRequired is the error message when radio band is required.
	ErrRadioBandRequired = "radio band is required for radio operations"

	// ErrCAPWAPDataUnavailable is the error message when CAPWAP data is not available.
	ErrCAPWAPDataUnavailable = "no control protocol data available" // nosec: G101 - This is not a credential

	// ErrNameMACMapDataUnavailable is the error message when name-MAC mapping data is not available.
	ErrNameMACMapDataUnavailable = "no name-MAC mapping data available"

	// ErrAPNotFoundByName is the error message when AP is not found by name.
	ErrAPNotFoundByName = "AP with name %q not found"

	// ErrAPNotFoundByMAC is the error message when AP is not found by MAC address.
	ErrAPNotFoundByMAC = "AP with MAC address %q not found"

	// ErrFailedGetCAPWAPData is the error message when getting CAPWAP data fails.
	ErrFailedGetCAPWAPData = "failed to get CAPWAP data: %w"

	// ErrFailedGetNameMACMapData is the error message when getting name-MAC mapping data fails.
	ErrFailedGetNameMACMapData = "failed to get name-MAC mapping data: %w"

	// ErrInvalidAPMac is the error message for invalid MAC address format.
	ErrInvalidAPMac = "invalid AP MAC address %s"

	// ErrUnsupportedStateType is the error message for unsupported state types.
	ErrUnsupportedStateType = "unsupported state type: %d"

	// ErrOperationFailed is the generic operation failure message template.
	ErrOperationFailed = ierrors.ErrOperationFailedTemplate

	// ErrRadioOperationFailed is the radio operation failure message template.
	ErrRadioOperationFailed = "failed to %s radio %s on AP %s: %w"

	// ErrStateOperationFailed is the state operation failure message template.
	ErrStateOperationFailed = "failed to %s state: %w"

	// ErrInvalidAPMacFormat is the error message for invalid MAC address format.
	ErrInvalidAPMacFormat = "invalid AP MAC address: %s"

	// ErrFailedResolveAPName is the error message when resolving AP name fails.
	ErrFailedResolveAPName = "failed to resolve AP name for MAC %s: %w"

	// ErrFailedAssignTags is the error message when assigning tags fails.
	ErrFailedAssignTags = "failed to assign tags to AP %s: %w"
)
