package core

// Radio administrative state constants.
const (
	// AdminStateEnabled represents the enabled admin state.
	AdminStateEnabled = "admin-state-enabled"

	// AdminStateDisabled represents the disabled admin state.
	AdminStateDisabled = "admin-state-disabled"
)

// GetAdminStateMode returns the admin state mode string based on enabled flag.
func GetAdminStateMode(enabled bool) string {
	if enabled {
		return AdminStateEnabled
	}
	return AdminStateDisabled
}
