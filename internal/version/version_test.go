package version

import (
	"os"
	"strings"
	"testing"
)

// TestVersionUnit_MatchesTheVersionFile_Success holds the constant to the VERSION file the release
// workflow tags from, which is the only thing keeping the User-Agent honest: a release that bumped
// VERSION alone would ship a header naming the previous release, with nothing else to notice.
//
// It fails make test-unit, so it fails the build and coverage jobs, and verify-version needs both
// — a release cannot tag through a mismatch.
func TestVersionUnit_MatchesTheVersionFile_Success(t *testing.T) {
	raw, err := os.ReadFile("../../VERSION")
	if err != nil {
		t.Fatalf("reading the VERSION file: %v", err)
	}

	if want := strings.TrimSpace(string(raw)); Version != want {
		t.Errorf("Version = %q but VERSION holds %q: bump both in the same commit", Version, want)
	}
}
