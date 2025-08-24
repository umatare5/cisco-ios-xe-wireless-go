package data

// ControllerTestConstants represents controller service-specific test constants
type ControllerTestConstants struct {
	// WNC reload operations
	ValidReloadReason   string
	InvalidReloadReason string
	LongReloadReason    string
	EmptyReloadReason   string

	// Force reload scenarios
	ForceReloadTrue  bool
	ForceReloadFalse bool
}

// StandardControllerTestConstants returns default test constants for controller operations
func StandardControllerTestConstants() ControllerTestConstants {
	return ControllerTestConstants{
		ValidReloadReason:   "Maintenance reload",
		InvalidReloadReason: "",
		LongReloadReason:    "Very long reload reason that exceeds normal limits for testing edge cases in the system behavior",
		EmptyReloadReason:   "",
		ForceReloadTrue:     true,
		ForceReloadFalse:    false,
	}
}
