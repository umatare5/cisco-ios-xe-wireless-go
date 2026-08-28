package version

// Version is this module's version, spelled as the VERSION file spells it: no leading "v".
//
// Nothing generates this line. A go:embed of ../../VERSION is rejected as an invalid pattern
// because the path leaves the package directory, and no build or CI step runs a generator, so the
// constant and the file are held together by version_test.go instead — bump both in one commit.
const Version = "0.9.2"
