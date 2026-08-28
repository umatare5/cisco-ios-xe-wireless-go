package contract_test

import (
	"maps"
	"slices"
	"testing"
)

// publishedLeaves names the leaves a consumer turns straight into a metric, keyed by the package
// and the type that declares them. The list is written out rather than derived: the same predicate
// swept over the tree would condemn every value-held field below the envelope, and those stay as
// they are for now, so a closed list is what separates the contract this gate holds from the
// remainder it does not.
var publishedLeaves = map[string]map[string][]string{
	"ap": {
		"ApJoinInfo": {
			"is-joined",
			"num-join-req-recvd",
			"num-config-req-recvd",
			"num-succ-join-resp-sent",
			"num-unsucc-join-req-procn",
			"num-succ-conf-resp-sent",
			"num-unsucc-conf-req-procn",
		},
		"ApDiscoveryInfo": {
			"num-discovery-req-recvd",
			"num-succ-disc-resp-sent",
			"num-err-disc-req",
		},
		"DTLSSessInfo": {
			"data-dtls-setup-req",
			"data-dtls-success",
			"data-dtls-failure",
			"ctrl-dtls-setup-req",
			"ctrl-dtls-success",
			"ctrl-dtls-failure",
			"ctrl-dtls-decrypt-err",
			"ctrl-dtls-anti-replay-err",
			"data-dtls-decrypt-err",
			"data-dtls-anti-replay-err",
		},
		"PhyTxPwrLvlCfgData": {
			"tx-power-level-1",
			"curr-tx-power-in-dbm",
		},
		"TagInfo": {
			"is-ap-misconfigured",
		},
	},
}

// publishedLeafCount is how many leaves the list above holds. Checking it stops the gate from
// passing because an entry was dropped: every other assertion here is made per leaf, so a list
// that loses one loses the finding with it.
const publishedLeafCount = 23

// TestEveryPublishedLeafCanBeAbsent holds the leaves a consumer publishes as a metric to the one
// shape that can tell the controller's silence from a reading: a pointer with omitempty. A zero
// counter, an AP that has not joined and a radio at 0 dBm are all legitimate readings, so a value
// field there gives absence and measurement the same representation, and the consumer cannot
// withhold a series it has no reading for.
//
// The check runs at level 0 of each declaring type — the depth the type writes its own leaves at —
// and never descends. Depth is what separates this gate from a sweep: an inline anonymous struct
// or a nested container would put the predicate on fields no consumer publishes.
func TestEveryPublishedLeafCanBeAbsent(t *testing.T) {
	pkgs, _ := loadTree(t)

	byDir := make(map[string]*servicePkg, len(pkgs))
	for _, pkg := range pkgs {
		byDir[pkg.dir] = pkg
	}

	checked := 0

	for _, dir := range slices.Sorted(maps.Keys(publishedLeaves)) {
		pkg, ok := byDir[dir]
		if !ok {
			t.Errorf("%s/%s: the package the list names is not in the tree", serviceDir, dir)
			continue
		}

		for _, typeName := range slices.Sorted(maps.Keys(publishedLeaves[dir])) {
			fields, ok := pkg.types[typeName]
			if !ok {
				t.Errorf("%s/%s: %s is not declared", serviceDir, dir, typeName)
				continue
			}

			for _, tag := range publishedLeaves[dir][typeName] {
				field, ok := leafOf(fields, tag)
				if !ok {
					t.Errorf("%s/%s: %s declares no %q at its top level", serviceDir, dir, typeName, tag)
					continue
				}

				checked++

				if !field.nilable {
					t.Errorf("%s/%s: %s.%s is a value, want a pointer: its zero is a reading, so a value"+
						" cannot say the controller sent nothing", serviceDir, dir, typeName, tag)
				}

				if !field.omitempty {
					t.Errorf("%s/%s: %s.%s lacks omitempty, so an absent leaf is re-marshaled as null",
						serviceDir, dir, typeName, tag)
				}
			}
		}
	}

	if checked != publishedLeafCount {
		t.Errorf("checked %d published leaves, want %d", checked, publishedLeafCount)
	}
}

// leafOf returns the field a type declares under tag at its own top level.
func leafOf(fields []tagged, tag string) (tagged, bool) {
	for _, field := range fields {
		if field.level == 0 && field.tag == tag {
			return field, true
		}
	}

	return tagged{}, false
}

// scalarKinds is the set of types a leaf decodes into. A pointer to one of these is how this tree
// says "the controller may not have sent this"; anything else is a container or a slice, whose
// nil-ness the envelope gate owns.
//
// time.Time is here with the predeclared types because the wire carries it as one value, an RFC
// 3339 string, whatever Go spells it as. It is the only qualified type this tree binds to a leaf;
// the other qualified type in field position is the embedded service.BaseService, which carries no
// tag and is dropped before this map is consulted.
var scalarKinds = map[string]bool{
	"string": true, "bool": true,
	"int": true, "int8": true, "int16": true, "int32": true, "int64": true,
	"uint": true, "uint8": true, "uint16": true, "uint32": true, "uint64": true,
	"float32": true, "float64": true,
	"time.Time": true,
}

// TestEveryPointerLeafCanBeOmitted holds a property that needs no list to maintain: a leaf the
// tree already made a pointer must also carry omitempty, or an absent leaf is re-marshaled as
// null rather than disappearing.
//
// It says nothing about which leaves must be pointers, so it does not replace the gate above:
// reverting a publish-path leaf to a value passes here and fails there.
func TestEveryPointerLeafCanBeOmitted(t *testing.T) {
	pkgs, _ := loadTree(t)

	checked := 0

	for _, pkg := range pkgs {
		for _, typeName := range slices.Sorted(maps.Keys(pkg.types)) {
			for _, field := range pkg.types[typeName] {
				if !field.pointer || !scalarKinds[field.typeName] {
					continue
				}

				checked++

				if !field.omitempty {
					t.Errorf("%s/%s: %s.%s is a pointer leaf without omitempty, so an absent leaf"+
						" is re-marshaled as null", serviceDir, pkg.dir, typeName, field.tag)
				}
			}
		}
	}

	if checked == 0 {
		t.Fatal("no pointer leaf was checked, so a pass here would mean nothing")
	}

	t.Logf("checked %d pointer leaves", checked)
}

// trueByDefaultLeaves names every leaf this tree declares that the controller's own schema gives
// `default "true"`, keyed by the package and the type that declares it. Absence there is not the
// zero reading — it means the default is in force — so a value field inverts the answer instead
// of merely flattening it, which is the one decoding error this library exists to avoid.
//
// The list is written out because the arbiter is the device, not the source: it was built by reading
// the `default` substatement of each module the controller implements, taken through
// ietf-yang-library. It is therefore a closure over the modules measured so far, not a sample, and
// it grows only when another module is read.
var trueByDefaultLeaves = map[string]map[string][]string{
	"wlan": {
		"WlanCfgEntry": {
			"auth-key-mgmt-dot1x",
			"okc",
			"security-wpa",
			"wlan-11k-neigh-list",
			"wpa2-aes",
			"wpa2-enabled",
		},
		"APFVapIDData":       {"broadcast-ssid"},
		"APFVap80211vData":   {"dot11v-dms"},
		"Dot11beProfile":     {"eht-ofdma-downlink", "eht-ofdma-uplink"},
		"UmbrellaFlexParams": {"dhcp-dns-option-enable"},
		"WlanSwitchingPolicy": {
			"central-assoc-enable",
			"central-authentication",
			"central-dhcp",
			"central-switching",
		},
	},
}

// trueByDefaultLeafCount is how many leaves the list above holds, for the same reason
// publishedLeafCount exists: a dropped entry would take its finding with it.
const trueByDefaultLeafCount = 15

// TestEveryTrueByDefaultLeafIsNilable holds the leaves whose schema default is true to a pointer
// with omitempty. It is not the same property as TestEveryPublishedLeafCanBeAbsent: that gate asks
// whether a consumer can withhold a series, this one asks whether the SDK can report the value at
// all. A default-false leaf is deliberately absent from the list — decoding its absence as false
// gives the right answer, so pointerising it would be symmetry rather than a fix.
//
// Most of the leaves listed were already pointers before this gate existed, which is what makes a
// pass meaningful: they are the positive control that the predicate matches the shape the tree
// already treats as correct, so the gate cannot pass by asserting nothing.
func TestEveryTrueByDefaultLeafIsNilable(t *testing.T) {
	pkgs, _ := loadTree(t)

	byDir := make(map[string]*servicePkg, len(pkgs))
	for _, pkg := range pkgs {
		byDir[pkg.dir] = pkg
	}

	checked := 0

	for _, dir := range slices.Sorted(maps.Keys(trueByDefaultLeaves)) {
		pkg, ok := byDir[dir]
		if !ok {
			t.Errorf("%s/%s: the package the list names is not in the tree", serviceDir, dir)
			continue
		}

		for _, typeName := range slices.Sorted(maps.Keys(trueByDefaultLeaves[dir])) {
			fields, ok := pkg.types[typeName]
			if !ok {
				t.Errorf("%s/%s: %s is not declared", serviceDir, dir, typeName)
				continue
			}

			for _, tag := range trueByDefaultLeaves[dir][typeName] {
				field, ok := leafOf(fields, tag)
				if !ok {
					t.Errorf("%s/%s: %s declares no %q at its top level", serviceDir, dir, typeName, tag)
					continue
				}

				checked++

				if !field.nilable {
					t.Errorf("%s/%s: %s.%s is a value, want a pointer: the schema defaults it to true,"+
						" so decoding its absence as false inverts the reading",
						serviceDir, dir, typeName, tag)
				}

				if !field.omitempty {
					t.Errorf("%s/%s: %s.%s lacks omitempty, so an absent leaf is re-marshaled as null",
						serviceDir, dir, typeName, tag)
				}
			}
		}
	}

	if checked != trueByDefaultLeafCount {
		t.Errorf("checked %d default-true leaves, want %d", checked, trueByDefaultLeafCount)
	}
}
