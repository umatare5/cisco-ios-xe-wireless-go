package service

import (
	"errors"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil"
)

// TestServiceNameUnit_RequireAPName_Sentinel pins which sentinel a blank name carries. It is the
// class RequireMACAddress uses for a blank address, so a caller can ask "was this record absent"
// with one predicate for either identifier.
func TestServiceNameUnit_RequireAPName_Sentinel(t *testing.T) {
	t.Parallel()

	for _, apName := range []string{"", " ", "\t\n"} {
		err := RequireAPName(apName)

		testutil.AssertTrue(
			t,
			errors.Is(err, core.ErrResourceNotFound),
			"RequireAPName("+apName+") must be ErrResourceNotFound, got: "+errorText(err),
		)
	}
}

// TestServiceNameUnit_RequireAPName_Accepted pins that a name is taken as typed. The controller
// keys ap-name-mac-map and the RPC name arms by the name itself, so a surrounding or interior
// space is the caller's to own and must not be trimmed here.
func TestServiceNameUnit_RequireAPName_Accepted(t *testing.T) {
	t.Parallel()

	for _, apName := range []string{"TEST-AP01", " TEST-AP01", "TEST AP01"} {
		testutil.AssertNoError(t, RequireAPName(apName), "RequireAPName("+apName+")")
	}
}
