package contract_test

import (
	"go/ast"
	"maps"
	"slices"
	"strings"
	"testing"
)

// payloadSuffix names the request bodies. Every type this SDK marshals into a PUT, PATCH, POST or
// RPC body is declared in an rpc_payload.go and named with this suffix — APTagPayload,
// CiscoIOSXEWirelessRFCfgRFTagsPayload, CiscoIOSXEWirelessSiteTagConfigsPayload,
// CiscoIOSXEWirelessWlanPolicyListEntriesPayload, APConfigRPCPayload, APSlotConfigRPCPayload —
// and every write call site under service/ passes one of those six. A request body is the one
// place a value struct is right: the node it names is required, and a nil pointer would marshal
// to null where the controller expects an object.
const payloadSuffix = "Payload"

// TestEveryLevelZeroNodeIsNilable holds every response wrapper this tree declares to the one shape
// that can tell an empty node from an absent one.
//
// It walks every struct type declared under service/, not only the types a core.Get call site
// names. A wrapper no accessor targets is invisible to a walk that starts from the call sites, and
// such a wrapper can hold its node as a value while every route-reachable gate passes.
//
// The base type is not examined. A module-qualified field holds either a struct or, where the route
// reads a single leaf, a qualified scalar; both fail the same way when held as a value, because a
// read answered with no body decodes to the zero value rather than to an absent reading.
func TestEveryLevelZeroNodeIsNilable(t *testing.T) {
	pkgs, _ := loadTree(t)

	examined := 0

	var excluded []string

	for _, pkg := range pkgs {
		requestBodies := writeBodies(pkg)

		for _, typeName := range slices.Sorted(maps.Keys(pkg.types)) {
			for _, field := range topTags(pkg.types[typeName]) {
				if moduleOf(field.tag) == "" {
					continue
				}

				if requestBodies[typeName] {
					if !field.nilable {
						excluded = append(excluded, pkg.dir+"/"+typeName+"."+field.tag)
					}

					continue
				}

				examined++

				if field.nilable {
					continue
				}

				t.Errorf("%s/%s: %s.%s is a value, want a slice or a pointer: a response that omits"+
					" the node decodes into a container of zeros no reader can tell from a reading",
					serviceDir, pkg.dir, typeName, field.tag)
			}
		}
	}

	if examined == 0 {
		t.Fatal("no module-qualified field was examined, so a pass here would mean nothing")
	}

	// The exclusions are logged rather than counted, because a count is a number that rots while a
	// list is something a reviewer can read. A request body added under a name that does not end in
	// Payload is reported as a finding instead of being excluded, which is the safe direction.
	t.Logf("examined %d module-qualified fields under %s; excluded %d request-body field(s): %s",
		examined, serviceDir, len(excluded), strings.Join(excluded, ", "))
}

// nodeKey identifies one YANG node as it is decoded inside one package: the local part of the JSON
// name, paired with the Go type the field decodes it into.
//
// The local part is what makes a parent and a wrapper comparable. A parent holds a child
// unqualified ("rogue-stats") while the wrapper that answers the child's own route holds it
// qualified ("Cisco-IOS-XE-wireless-rogue-oper:rogue-stats"), because RFC 7951 4 prefixes a member
// only where its module differs from its parent's.
//
// The Go type keeps the rule off a legitimate shape difference. service/ap declares both
// APTagPayload.ApTag of type APCfgApTagData and ApTags.ApTag of type ApTag under the local name
// "ap-tag"; pairing the name with the type is what stops those two from being compared.
type nodeKey struct {
	local string
	typ   string
}

// shape is one field reduced to what the rule compares: which type declares it, and whether it can
// hold the absence of the node it claims.
type shape struct {
	owner string
	// qualified is true when this field claimed the node with a module prefix, which is the shape
	// RFC 7951 4 gives the sole top-level member of an envelope: the evidence, and the only
	// evidence in the tree, that the node is readable on its own route.
	qualified bool
	// nilable is true for a pointer and for a slice alike: either can hold nil.
	nilable bool
	// pointer separates a deliberate "the controller may not have sent this" from a slice, whose
	// nil-ness comes free with the type.
	pointer bool
}

// TestEveryNodeIsNilableWhereverItIsDecoded holds the consistency half of the same property: a node
// that is nil-able in one decode path must be nil-able in every other decode path that shares its
// Go type.
//
// The rule needs no list of routes. A node readable on its own route has a wrapper whose sole
// member is a pointer, so within one package, once any field decodes that node into a pointer of
// type T, every other field decoding the same node into T carries the same absence contract. The
// predicate is symmetric and so covers both orders.
//
// Two exclusions. A node held as a value on every side is left alone, because value-held containers
// below the envelope are intended here. Scalar leaves are left to the absence gates, because a local
// name such as "state" names a different leaf under each parent and comparing by name would
// manufacture findings out of a coincidence of spelling.
func TestEveryNodeIsNilableWhereverItIsDecoded(t *testing.T) {
	pkgs, _ := loadTree(t)

	nodes, compared := 0, 0

	for _, pkg := range pkgs {
		byNode := nodesOf(pkg)

		for _, key := range slices.SortedFunc(maps.Keys(byNode), compareNodeKeys) {
			nodes++

			fields := byNode[key]
			if len(fields) < 2 {
				continue
			}

			compared++

			if !slices.ContainsFunc(fields, func(f shape) bool { return f.pointer }) {
				continue
			}

			// Without an anchor the group rests on nothing but a coincidence of spelling: nine
			// containers under service/ble each hold a "report" of one Go type, and none of those
			// nine nodes is readable on its own route.
			if !slices.ContainsFunc(fields, func(f shape) bool { return f.qualified }) {
				continue
			}

			for _, field := range fields {
				if field.nilable || strings.HasSuffix(field.owner, payloadSuffix) {
					continue
				}

				t.Errorf("%s/%s: %s.%s is a value, want a pointer: %s is decoded as a pointer at %s,"+
					" so this field cannot hold the absence the same node is allowed elsewhere",
					serviceDir, pkg.dir, field.owner, key.local, key.typ, pointerHolders(fields))
			}
		}
	}

	// Both counts are asserted because either can be emptied on its own. A refactor that stops
	// collecting fields drives nodes to zero; one that stops grouping them leaves nodes intact while
	// every group becomes a singleton, and a gate whose predicate never runs passes.
	if nodes == 0 {
		t.Fatal("no (node, type) pair was examined, so a pass here would mean nothing")
	}

	if compared == 0 {
		t.Fatal("no (node, type) pair is decoded in more than one place, so the predicate never ran")
	}

	t.Logf("examined %d (node, type) pairs under %s; %d of them are decoded in more than one place",
		nodes, serviceDir, compared)
}

// nodesOf groups every tagged field the package declares by the node it decodes and the type it
// decodes it into. Every struct type is walked at every nesting level, because a node held inside
// an inline anonymous struct is decoded no less than one held in a named type:
// service/mesh/global_oper.go holds mesh-global-stats that way.
//
// Two kinds of field are left out. A scalar is a leaf, whose absence contract absence_test.go owns
// and whose local name recurs under unrelated parents; time.Time is in scalarKinds for that reason,
// so a timestamp is excluded here as the leaf it is. A field whose type flatten leaves unnamed is
// skipped because an inline anonymous struct arrives with an empty type name, and two of those are
// not the same type merely because neither could be named — TestEveryLevelZeroNodeIsNilable is what
// covers an unnamed type at level 0.
func nodesOf(pkg *servicePkg) map[nodeKey][]shape {
	byNode := make(map[nodeKey][]shape)

	for _, owner := range slices.Sorted(maps.Keys(pkg.types)) {
		for _, field := range pkg.types[owner] {
			if field.tag == "" || field.tag == "-" || field.typeName == "" {
				continue
			}

			if scalarKinds[field.typeName] {
				continue
			}

			key := nodeKey{local: localName(field.tag), typ: field.typeName}
			byNode[key] = append(byNode[key], shape{
				owner:     owner,
				qualified: moduleOf(field.tag) != "",
				nilable:   field.nilable,
				pointer:   field.pointer,
			})
		}
	}

	return byNode
}

// localName returns the part of a JSON name that identifies the node, dropping the module prefix
// RFC 7951 4 puts on a member whose module differs from its parent's. It is the complement of
// moduleOf, which the envelope gates use to read that prefix.
func localName(tag string) string {
	if _, local, qualified := strings.Cut(tag, ":"); qualified {
		return local
	}

	return tag
}

// pointerHolders names the fields that already decode the node as a pointer, so a finding says
// which existing contract the offending field contradicts.
func pointerHolders(fields []shape) string {
	var owners []string

	for _, field := range fields {
		if field.pointer {
			owners = append(owners, field.owner)
		}
	}

	return strings.Join(owners, ", ")
}

// compareNodeKeys orders the keys so findings come out in the same order on every run.
func compareNodeKeys(a, b nodeKey) int {
	if a.local != b.local {
		return strings.Compare(a.local, b.local)
	}

	return strings.Compare(a.typ, b.typ)
}

// writeBodies names the types the package marshals into a request body, read off the write call
// sites rather than off a naming convention. The name is not evidence: service/rf declares
// CiscoIOSXEWirelessRFCfgRFTagPayload, whose own comment calls it a response structure and which
// service/rf/tag_service.go:37 hands to core.Get. Only a write call site distinguishes the two.
func writeBodies(pkg *servicePkg) map[string]bool {
	bodies := make(map[string]bool)

	for _, name := range slices.Sorted(maps.Keys(pkg.funcs)) {
		fn := pkg.funcs[name]

		ast.Inspect(fn.Body, func(node ast.Node) bool {
			call, ok := node.(*ast.CallExpr)
			if !ok {
				return true
			}

			if body, ok := writeCall(call); ok {
				if typeName, ok := bodyType(pkg, fn, body, 0); ok {
					bodies[typeName] = true
				}
			}

			return true
		})
	}

	return bodies
}

// writeMethods are the core entry points that carry a request body (internal/core/request.go).
var writeMethods = map[string]bool{
	"Post": true, "PostVoid": true, "PostRPCVoid": true,
	"Put": true, "PutVoid": true,
	"Patch": true, "PatchVoid": true,
}

// writeCall reports the body operand of a core write call. The body is the fourth operand, beside
// the context, the client and the endpoint.
func writeCall(call *ast.CallExpr) (ast.Expr, bool) {
	fun := call.Fun
	if index, ok := fun.(*ast.IndexExpr); ok {
		fun = index.X
	}

	sel, ok := fun.(*ast.SelectorExpr)
	if !ok || !writeMethods[sel.Sel.Name] {
		return nil, false
	}

	if pkgName, ok := sel.X.(*ast.Ident); !ok || pkgName.Name != "core" {
		return nil, false
	}

	if len(call.Args) < 4 {
		return nil, false
	}

	return call.Args[3], true
}

// bodyType resolves a body operand to the named type it is: a composite literal, a variable the
// same function assigned, or a call of a package-local builder, whose declared result names the
// type without the body having to be read.
func bodyType(pkg *servicePkg, fn *ast.FuncDecl, expr ast.Expr, hops int) (string, bool) {
	if hops > maxHops {
		return "", false
	}

	switch node := expr.(type) {
	case *ast.ParenExpr:
		return bodyType(pkg, fn, node.X, hops+1)

	case *ast.UnaryExpr:
		return bodyType(pkg, fn, node.X, hops+1)

	case *ast.CompositeLit:
		if ident, ok := node.Type.(*ast.Ident); ok {
			return ident.Name, true
		}

	case *ast.Ident:
		assigned, ok := assignedIn(fn, node.Name)
		if !ok {
			return "", false
		}

		return bodyType(pkg, fn, assigned, hops+1)

	case *ast.CallExpr:
		sel, ok := node.Fun.(*ast.SelectorExpr)
		if !ok {
			return "", false
		}

		builder, ok := pkg.funcs[sel.Sel.Name]
		if !ok || builder.Type.Results == nil || len(builder.Type.Results.List) != 1 {
			return "", false
		}

		if ident, ok := builder.Type.Results.List[0].Type.(*ast.Ident); ok {
			return ident.Name, true
		}
	}

	return "", false
}
