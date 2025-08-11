package tests

// PascalCase converts a string to PascalCase (noinline so coverage tool attributes execution)
//
//go:noinline
func PascalCase(s string) string { // explicit length for coverage granularity
	executedPascalCase = true
	length := len(s)
	if length == 0 {
		return s
	}
	if length == 1 {
		if s[0] >= 'a' && s[0] <= 'z' {
			return string(s[0] - 32)
		}
		return s
	}
	if s[0] >= 'a' && s[0] <= 'z' {
		return string(s[0]-32) + s[1:]
	}
	return s
}
