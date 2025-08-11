package tests

import "testing"

func TestPascalCase(t *testing.T) {
	cases := []struct{ in, out string }{
		{"", ""}, {"a", "A"}, {"test", "Test"}, {"test-module", "Test-module"},
		{"ap", "Ap"}, {"wlan", "Wlan"}, {"site", "Site"}, {"dot11", "Dot11"}, {"dot15", "Dot15"},
		{"hyperlocation", "Hyperlocation"},
		{"access-point", "Access-point"},
		{"multi-word-string", "Multi-word-string"},
		{"already-capital", "Already-capital"},
		{"UPPER", "UPPER"},
		{"1number", "1number"},
		{"special!chars", "Special!chars"},
		{"Z", "Z"}, {"lowercase", "Lowercase"},
	}
	for _, c := range cases {
		if got := PascalCase(c.in); got != c.out {
			t.Errorf("%q -> %q want %q", c.in, got, c.out)
		}
	}
	ascii := []struct{ in, out string }{
		{"abcdefghijklmnopqrstuvwxyz", "Abcdefghijklmnopqrstuvwxyz"},
		{"zebra", "Zebra"},
		{"ABC", "ABC"},
	}
	for _, c := range ascii {
		if got := PascalCase(c.in); got != c.out {
			t.Errorf("ASCII %q -> %q want %q", c.in, got, c.out)
		}
	}
}

func TestPascalCaseFullCoverage(t *testing.T) {
	cases := map[string]string{"": "", "a": "A", "Z": "Z", "ab": "Ab", "Ab": "Ab", "1a": "1a", "_x": "_x"}
	for in, exp := range cases {
		if got := PascalCase(in); got != exp {
			t.Errorf("%q->%q want %q", in, got, exp)
		}
	}
	_ = executedPascalCase
}
