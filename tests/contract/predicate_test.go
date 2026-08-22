package contract_test

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

// TestTagMatchesRouteStripsTheListKey pins the predicate the first gate rests on. A keyed read
// ends its URL at "node=key" while the envelope it answers with is keyed "Module:node", and a key
// carries ":" and "," and "/" of its own, so the node name survives only if the key and the query
// are cut before the last segment is taken.
func TestTagMatchesRouteStripsTheListKey(t *testing.T) {
	const (
		apOper   = "/restconf/data/Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data"
		wlanOper = "/restconf/data/Cisco-IOS-XE-wireless-wlan-global-oper:wlan-global-oper-data"
		rfCfg    = "/restconf/data/Cisco-IOS-XE-wireless-rf-cfg:rf-cfg-data"
	)

	cases := []struct {
		name  string
		tag   string
		route string
		want  bool
	}{
		{
			"root container matches its own qualified tag",
			"Cisco-IOS-XE-wireless-rf-cfg:rf-cfg-data", rfCfg, true,
		},
		{
			"whole-container tag on a sub-container route",
			"Cisco-IOS-XE-wireless-wlan-global-oper:wlan-global-oper-data", wlanOper + "/wlan-info", false,
		},
		{
			"sub-container tag on its own route",
			"Cisco-IOS-XE-wireless-wlan-global-oper:wlan-info", wlanOper + "/wlan-info", true,
		},
		{
			"colon-separated key",
			"Cisco-IOS-XE-wireless-access-point-oper:capwap-data", apOper + "/capwap-data=a:b:c", true,
		},
		{
			"key holding a slash",
			"Cisco-IOS-XE-wireless-access-point-oper:capwap-data", apOper + "/capwap-data=a/b", true,
		},
		{
			"composite key",
			"Cisco-IOS-XE-wireless-access-point-oper:radio-oper-data", apOper + "/radio-oper-data=a:b:c,1", true,
		},
		{
			"query appended to a keyed read",
			"Cisco-IOS-XE-wireless-rf-cfg:rf-tag", rfCfg + "/rf-tags/rf-tag=name?with-defaults=report-all", true,
		},
		{
			"query appended to a container read",
			"Cisco-IOS-XE-wireless-rf-cfg:rf-cfg-data", rfCfg + "?with-defaults=report-all", true,
		},
		{
			"unqualified tag",
			"capwap-data", apOper + "/capwap-data", false,
		},
		{
			"element type of the keyed list",
			"ap-mac", apOper + "/capwap-data=a:b:c", false,
		},
		{
			"right node under the wrong module",
			"Cisco-IOS-XE-wireless-rrm-oper:capwap-data", apOper + "/capwap-data", false,
		},
		{
			"key holding a colon does not supply the module",
			"Cisco-IOS-XE-wireless-access-point-oper:capwap-data", "/restconf/data/capwap-data=a:b:c", false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := tagMatchesRoute(tc.tag, tc.route); got != tc.want {
				t.Errorf("tagMatchesRoute(%q, %q) = %v, want %v", tc.tag, tc.route, got, tc.want)
			}
		})
	}
}

// TestWalkTagsFindsOnlyTheSameModuleRepeat pins the predicate the second gate rests on. Testing a
// nested tag for ":" would condemn a child that a future augment puts in another module, which
// RFC 7951 4 requires to be qualified; only a comparison against the module in force at the
// parent separates the two.
func TestWalkTagsFindsOnlyTheSameModuleRepeat(t *testing.T) {
	const (
		general = "Cisco-IOS-XE-wireless-general-cfg"
		wat     = "Cisco-IOS-XE-wireless-wat-cfg"
	)

	types := map[string][]tagged{
		// An inline wrapper whose child repeats the module, as service/wat/cfg.go did.
		"RepeatingInline": {
			{tag: wat + ":wat-cfg-data", level: 0, nilable: true},
			{typeName: "WATConfig", tag: wat + ":wat-config", level: 1, nilable: true},
		},
		"WATConfig": {{tag: "wat-enable", level: 0}},

		// A named wrapper whose child repeats the module, as service/general/cfg.go did.
		"RepeatingNamed":  {{typeName: "RepeatedChild", tag: general + ":mewlc-config", level: 0, nilable: true}},
		"RepeatedChild":   {{typeName: "MewlcConfigData", tag: general + ":mewlc-config", level: 0, nilable: true}},
		"MewlcConfigData": {{tag: "mewlc-platform", level: 0}},

		// The same wrapper pointed straight at the child that decodes.
		"Repaired": {{typeName: "MewlcConfigData", tag: general + ":mewlc-config", level: 0, nilable: true}},

		// A child in another module is qualified legitimately.
		"Augmented": {{typeName: "AugmentedChild", tag: general + ":laginfo", level: 0, nilable: true}},
		"AugmentedChild": {
			{tag: "Cisco-IOS-XE-wireless-other-cfg:added-leaf", level: 0, nilable: true},
		},
	}

	cases := []struct {
		decode string
		want   int
	}{
		{"RepeatingInline", 1},
		{"RepeatingNamed", 1},
		{"Repaired", 0},
		{"Augmented", 0},
	}

	for _, tc := range cases {
		t.Run(tc.decode, func(t *testing.T) {
			var found []string
			walkTags(types, tc.decode, "", 0, make(map[string]bool), &found)

			if len(found) != tc.want {
				t.Errorf("walkTags(%s) reported %v, want %d finding(s)", tc.decode, found, tc.want)
			}
		})
	}
}

// TestEvalConstFoldsTheRoutesShapes pins the fold the first gate resolves its routes with. The
// route constants are written as concatenations of each other, so a resolver that reads only
// string literals sees no route at all.
func TestEvalConstFoldsTheRoutesShapes(t *testing.T) {
	src := "package routes\n" +
		"const (\n" +
		"\tRESTCONFDataPath = \"/restconf/data\"\n" +
		"\tRFCfgPath = RESTCONFDataPath + \"/Cisco-IOS-XE-wireless-rf-cfg:rf-cfg-data\"\n" +
		"\tRFTagsPath = RFCfgPath + \"/rf-tags\"\n" +
		"\tNotAPath = 42\n" +
		"\tElsewhere = OtherPath + \"/leaf\"\n" +
		")\n"

	fset := token.NewFileSet()

	file, err := parser.ParseFile(fset, "routes.go", src, parser.SkipObjectResolution)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}

	raw := make(map[string]ast.Expr)
	collectConsts(file, raw)

	cases := []struct {
		name string
		want string
		ok   bool
	}{
		{"RESTCONFDataPath", "/restconf/data", true},
		{"RFCfgPath", "/restconf/data/Cisco-IOS-XE-wireless-rf-cfg:rf-cfg-data", true},
		{"RFTagsPath", "/restconf/data/Cisco-IOS-XE-wireless-rf-cfg:rf-cfg-data/rf-tags", true},
		{"NotAPath", "", false},
		{"Elsewhere", "", false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, ok := evalConst(raw, raw[tc.name], 0)
			if got != tc.want || ok != tc.ok {
				t.Errorf("evalConst(%s) = %q, %v; want %q, %v", tc.name, got, ok, tc.want, tc.ok)
			}
		})
	}
}
