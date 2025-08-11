package tests

import "time"

// TestTimeouts
const (
	// DefaultTestTimeout is the default timeout for tests
	DefaultTestTimeout = 30 * time.Second
	// ExtendedTestTimeout is an extended timeout for longer tests
	ExtendedTestTimeout = 60 * time.Second
	// ShortTestTimeout is a short timeout for quick tests
	ShortTestTimeout = 5 * time.Second
)
