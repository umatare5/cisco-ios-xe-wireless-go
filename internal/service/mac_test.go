package service

import (
	"errors"
	"testing"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/testutil"
)

// TestServiceMACUnit_RequireMACAddress_Sentinel pins which sentinel an unusable list key
// carries. The empty and whitespace-only forms answer core.ErrResourceNotFound, which
// core.IsNotFoundError reports true for, and a malformed address must not: a caller asking
// "was this record absent" would otherwise read a typo as an absence.
func TestServiceMACUnit_RequireMACAddress_Sentinel(t *testing.T) {
	t.Parallel()

	for _, mac := range []string{"", " ", "\t\n"} {
		got, err := RequireMACAddress(mac)

		testutil.AssertTrue(
			t,
			errors.Is(err, core.ErrResourceNotFound),
			"RequireMACAddress("+mac+") must be ErrResourceNotFound, got: "+errorText(err),
		)
		testutil.AssertStringEquals(t, got, "", "empty key returns no address")
	}

	_, err := RequireMACAddress("00:11:22:33:44")
	testutil.AssertError(t, err, "a short address is rejected")
	testutil.AssertFalse(
		t,
		errors.Is(err, core.ErrResourceNotFound),
		"a malformed address must not be reported as an absent record",
	)
	testutil.AssertTrue(
		t,
		core.IsNotFoundError(core.ErrResourceNotFound),
		"positive control: the empty-key sentinel is what IsNotFoundError matches",
	)
}

// TestServiceMACUnit_RequireMACAddress_Normalization pins the wire form. The controller
// keys its lists by lower-case colon-separated addresses, so every separator spelling and
// either case has to arrive as the same string.
func TestServiceMACUnit_RequireMACAddress_Normalization(t *testing.T) {
	t.Parallel()

	const want = "00:11:22:33:44:55"

	for _, mac := range []string{
		"00:11:22:33:44:55",
		"00-11-22-33-44-55",
		"0011.2233.4455",
		"001122334455",
	} {
		got, err := RequireMACAddress(mac)
		testutil.AssertNoError(t, err, "RequireMACAddress("+mac+")")
		testutil.AssertStringEquals(t, got, want, "normalized form of "+mac)
	}

	got, err := RequireMACAddress("AA:BB:CC:DD:EE:FF")
	testutil.AssertNoError(t, err, "RequireMACAddress on an upper-case address")
	testutil.AssertStringEquals(t, got, "aa:bb:cc:dd:ee:ff", "upper case is folded down")
}

// TestServiceMACUnit_RequireMACAddress_Malformed pins that a malformed address is rejected
// before it can be concatenated into a URL, rather than reaching the controller as a key
// no list can hold.
func TestServiceMACUnit_RequireMACAddress_Malformed(t *testing.T) {
	t.Parallel()

	for _, mac := range []string{
		"invalid-mac",
		"00:11:22:33:44:55:66",
		"00:11:22:33:44:5g",
		"zz:zz:zz:zz:zz:zz",
		"00:11:22:33:44:55 ",
	} {
		got, err := RequireMACAddress(mac)
		testutil.AssertError(t, err, "RequireMACAddress("+mac+") must fail")
		testutil.AssertStringEquals(t, got, "", "a rejected address returns no string")
	}
}

// errorText renders an error for a failure message without dereferencing nil.
func errorText(err error) string {
	if err == nil {
		return "<nil>"
	}
	return err.Error()
}
