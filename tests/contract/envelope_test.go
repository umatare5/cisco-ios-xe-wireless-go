// Contract tests hold the decode types to the routes that read them. They parse the tree instead
// of asking the runtime: what they check lives in a struct tag and in a route constant, and a
// mismatch between the two decodes to a zero value that no test of a response can see.
package contract_test

import (
	"go/ast"
	"go/parser"
	"go/token"
	"maps"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"slices"
	"strconv"
	"strings"
	"testing"
)

const (
	// repoRoot is the module root relative to this package: a test runs in its own directory.
	repoRoot = "../.."

	serviceDir = "service"
	routesDir  = "internal/restconf/routes"

	// routesPkg is the name every service file imports the route constants under.
	routesPkg = "routes"

	// maxHops bounds how far a route is traced through constants, variables and helpers.
	// Exceeding it leaves the endpoint unresolved, which the gate reports: a bound placed here
	// can add a finding, never hide one.
	maxHops = 8
)

// tagged is one field of a decode type: the named type it declares, the name its JSON tag gives
// it on the wire, the nesting an inline anonymous struct puts it under, and whether it can hold
// the absence of the node.
type tagged struct {
	typeName string
	tag      string
	level    int
	// nilable is true for a pointer and for a slice alike, because either can be nil.
	// pointer separates the two: a slice is nil-able by nature, while a pointer is a
	// deliberate statement that the controller may not have sent the node.
	nilable   bool
	pointer   bool
	omitempty bool
}

// site is one core.Get call: where it is written, the type it decodes into, the expression that
// says which node it reads, and the function that expression is resolved in.
type site struct {
	pos      string
	decode   string
	endpoint ast.Expr
	fn       *ast.FuncDecl
}

// servicePkg is one package under service/: the struct types it declares, the functions a route
// can be traced through, and the reads it makes.
type servicePkg struct {
	dir   string
	types map[string][]tagged
	funcs map[string]*ast.FuncDecl
	sites []site
}

// TestEveryGetDecodesTheNodeItsRouteReads holds every read to the envelope RFC 7951 answers it
// with: one top-level member, qualified by the module that defines the node the route ends at.
// A type that declares a different node, or more than one, decodes the body to a zero value and
// reports success.
func TestEveryGetDecodesTheNodeItsRouteReads(t *testing.T) {
	pkgs, routes := loadTree(t)

	reads := 0

	for _, pkg := range pkgs {
		for _, s := range pkg.sites {
			reads++

			route, ok := resolveEndpoint(pkg, routes, s.fn, s.endpoint, 0)
			if !ok {
				t.Errorf("%s: core.Get[%s] reads an endpoint the gate cannot resolve to a route", s.pos, s.decode)
				continue
			}

			top := topTags(pkg.types[s.decode])
			if len(top) != 1 {
				t.Errorf("%s: %s declares %d top-level JSON tags, want exactly 1", s.pos, s.decode, len(top))
				continue
			}

			if !tagMatchesRoute(top[0].tag, route) {
				t.Errorf("%s: %s is tagged %q, but %s reads %q under module %q",
					s.pos, s.decode, top[0].tag, route, nodeName(route), moduleOfRoute(route))
			}
		}
	}

	if reads == 0 {
		t.Fatalf("no core.Get call sites found under %s", serviceDir)
	}
}

// TestNoDecodeTypeRepeatsItsModuleBelowTheTop holds every decode type to RFC 7951 4, which
// qualifies a child only where its module differs from its parent's. A repeat of the module
// already in force names a key the controller never sends at that depth, so the child it names
// is dropped whole.
func TestNoDecodeTypeRepeatsItsModuleBelowTheTop(t *testing.T) {
	pkgs, _ := loadTree(t)

	for _, pkg := range pkgs {
		var repeats []string
		for _, decode := range decodeTypes(pkg) {
			walkTags(pkg.types, decode, "", 0, make(map[string]bool), &repeats)
		}

		// A wrapper that several reads reach is one defect, not one per read.
		reported := make(map[string]bool, len(repeats))
		for _, repeat := range repeats {
			reported[repeat] = true
		}

		for _, repeat := range slices.Sorted(maps.Keys(reported)) {
			t.Errorf("%s/%s: %s repeats the module already in force at its parent", serviceDir, pkg.dir, repeat)
		}
	}
}

// TestEveryEnvelopeFieldCanBeAbsent holds every decode type to the one shape that can tell an
// empty node from an absent one: a slice or a pointer. A value struct there decodes a controller
// that sent nothing into a container full of zeros.
func TestEveryEnvelopeFieldCanBeAbsent(t *testing.T) {
	pkgs, _ := loadTree(t)

	for _, pkg := range pkgs {
		for _, decode := range decodeTypes(pkg) {
			for _, field := range topTags(pkg.types[decode]) {
				if moduleOf(field.tag) == "" || field.nilable {
					continue
				}

				t.Errorf("%s/%s: %s.%s is a value, want a slice or a pointer",
					serviceDir, pkg.dir, decode, field.tag)
			}
		}
	}
}

// loadTree parses every non-test file under service/ and the route constants those files read.
// It fails when either comes back empty, because a gate that examines nothing passes.
func loadTree(t *testing.T) (pkgs []*servicePkg, routes map[string]string) {
	t.Helper()

	fset := token.NewFileSet()
	pkgs = loadServices(t, fset)
	routes = loadRoutes(t, fset)

	if len(pkgs) == 0 {
		t.Fatalf("no packages parsed under %s", serviceDir)
	}

	if len(routes) == 0 {
		t.Fatalf("no route constants parsed under %s", routesDir)
	}

	return pkgs, routes
}

// loadServices groups the files under service/ by their directory, which is the package a decode
// type and the read that uses it always share.
func loadServices(t *testing.T, fset *token.FileSet) []*servicePkg {
	t.Helper()

	byDir := make(map[string]*servicePkg)

	for _, file := range goFiles(t, serviceDir) {
		dir := filepath.Base(filepath.Dir(file))

		pkg, ok := byDir[dir]
		if !ok {
			pkg = &servicePkg{
				dir:   dir,
				types: make(map[string][]tagged),
				funcs: make(map[string]*ast.FuncDecl),
			}
			byDir[dir] = pkg
		}

		parsed := parseFile(t, fset, file)
		collectTypes(parsed, pkg.types)
		collectFuncs(parsed, pkg.funcs)
		pkg.sites = append(pkg.sites, collectSites(fset, parsed)...)
	}

	pkgs := make([]*servicePkg, 0, len(byDir))
	for _, dir := range slices.Sorted(maps.Keys(byDir)) {
		pkgs = append(pkgs, byDir[dir])
	}

	return pkgs
}

// loadRoutes folds the route constants into the paths they denote. The constants are written as
// concatenations of each other and spread over one file per service, so the whole directory is
// collected before any of it is folded.
func loadRoutes(t *testing.T, fset *token.FileSet) map[string]string {
	t.Helper()

	raw := make(map[string]ast.Expr)
	for _, file := range goFiles(t, routesDir) {
		collectConsts(parseFile(t, fset, file), raw)
	}

	routes := make(map[string]string, len(raw))
	for name, expr := range raw {
		if folded, ok := evalConst(raw, expr, 0); ok {
			routes[name] = folded
		}
	}

	return routes
}

// goFiles lists the non-test Go files under dir, named relative to the module root so that a
// finding points at the path the repository uses.
func goFiles(t *testing.T, dir string) []string {
	t.Helper()

	var files []string

	err := filepath.WalkDir(filepath.Join(repoRoot, dir), func(p string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || !strings.HasSuffix(p, ".go") || strings.HasSuffix(p, "_test.go") {
			return nil
		}

		rel, err := filepath.Rel(repoRoot, p)
		if err != nil {
			return err
		}
		files = append(files, rel)

		return nil
	})
	if err != nil {
		t.Fatalf("walk %s: %v", dir, err)
	}

	slices.Sort(files)

	return files
}

// parseFile parses one file, keeping the relative name for its positions.
func parseFile(t *testing.T, fset *token.FileSet, rel string) *ast.File {
	t.Helper()

	src, err := os.ReadFile(filepath.Join(repoRoot, rel))
	if err != nil {
		t.Fatalf("read %s: %v", rel, err)
	}

	parsed, err := parser.ParseFile(fset, rel, src, parser.SkipObjectResolution)
	if err != nil {
		t.Fatalf("parse %s: %v", rel, err)
	}

	return parsed
}

// collectTypes records the fields of every struct type the file declares.
func collectTypes(file *ast.File, into map[string][]tagged) {
	for _, decl := range file.Decls {
		gen, ok := decl.(*ast.GenDecl)
		if !ok || gen.Tok != token.TYPE {
			continue
		}

		for _, spec := range gen.Specs {
			typ, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}
			if structType, ok := typ.Type.(*ast.StructType); ok {
				into[typ.Name.Name] = flatten(structType, 0)
			}
		}
	}
}

// collectFuncs records the file's functions by name, which is how a read that builds its URL in a
// helper is traced back to the route the helper names.
func collectFuncs(file *ast.File, into map[string]*ast.FuncDecl) {
	for _, decl := range file.Decls {
		if fn, ok := decl.(*ast.FuncDecl); ok && fn.Body != nil {
			into[fn.Name.Name] = fn
		}
	}
}

// collectConsts records the file's constant expressions unfolded, because a constant may be
// written in terms of one the caller has not read yet.
func collectConsts(file *ast.File, into map[string]ast.Expr) {
	for _, decl := range file.Decls {
		gen, ok := decl.(*ast.GenDecl)
		if !ok || gen.Tok != token.CONST {
			continue
		}

		for _, spec := range gen.Specs {
			value, ok := spec.(*ast.ValueSpec)
			if !ok {
				continue
			}

			for i, name := range value.Names {
				if i < len(value.Values) {
					into[name.Name] = value.Values[i]
				}
			}
		}
	}
}

// collectSites records every core.Get call in the file together with the function it sits in,
// which is the scope the endpoint expression is resolved in.
func collectSites(fset *token.FileSet, file *ast.File) []site {
	var sites []site

	for _, decl := range file.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if !ok || fn.Body == nil {
			continue
		}

		ast.Inspect(fn.Body, func(node ast.Node) bool {
			call, ok := node.(*ast.CallExpr)
			if !ok {
				return true
			}

			decode, endpoint, ok := getCall(call)
			if ok {
				sites = append(sites, site{
					pos:      fset.Position(call.Pos()).String(),
					decode:   decode,
					endpoint: endpoint,
					fn:       fn,
				})
			}

			return true
		})
	}

	return sites
}

// getCall reports the type a core.Get call decodes into and the expression that gives it its
// endpoint. The type is the call's type argument and the endpoint its third operand
// (internal/core/request.go).
func getCall(call *ast.CallExpr) (string, ast.Expr, bool) {
	index, ok := call.Fun.(*ast.IndexExpr)
	if !ok {
		return "", nil, false
	}

	sel, ok := index.X.(*ast.SelectorExpr)
	if !ok || sel.Sel.Name != "Get" {
		return "", nil, false
	}

	if pkg, ok := sel.X.(*ast.Ident); !ok || pkg.Name != "core" {
		return "", nil, false
	}

	decode, ok := index.Index.(*ast.Ident)
	if !ok || len(call.Args) < 3 {
		return "", nil, false
	}

	return decode.Name, call.Args[2], true
}

// flatten lists a struct's fields, recording for each the nesting an inline anonymous struct adds:
// such a struct contributes its fields to the same list, because the wire nests them under the tag
// of the field that declares it.
func flatten(structType *ast.StructType, level int) []tagged {
	var fields []tagged

	for _, field := range structType.Fields.List {
		base, nilable := baseType(field.Type)
		_, isPointer := field.Type.(*ast.StarExpr)
		entry := tagged{
			tag:       jsonName(field.Tag),
			level:     level,
			nilable:   nilable,
			pointer:   isPointer,
			omitempty: hasOmitEmpty(field.Tag),
		}

		switch typ := base.(type) {
		case *ast.Ident:
			entry.typeName = typ.Name
		case *ast.SelectorExpr:
			// A qualified type is named rather than left blank because time.Time is a leaf on the
			// wire whatever Go spells it as. An unnamed one is invisible to scalarKinds, which is
			// how a pointer to it escaped the omitempty gate.
			if qualifier, ok := typ.X.(*ast.Ident); ok {
				entry.typeName = qualifier.Name + "." + typ.Sel.Name
			}
		case *ast.StructType:
			fields = append(fields, entry)
			fields = append(fields, flatten(typ, level+1)...)

			continue
		}

		fields = append(fields, entry)
	}

	return fields
}

// baseType unwraps the pointers and slices around a field's type, reporting whether any of them
// makes the field nil-able: only a slice or a pointer can hold the absence of a node.
func baseType(expr ast.Expr) (ast.Expr, bool) {
	nilable := false

	for {
		switch typ := expr.(type) {
		case *ast.StarExpr:
			nilable, expr = true, typ.X
		case *ast.ArrayType:
			nilable, expr = true, typ.Elt
		default:
			return expr, nilable
		}
	}
}

// jsonName returns the name a field's JSON tag gives it on the wire, or "" when it has none.
func jsonName(tag *ast.BasicLit) string {
	if tag == nil {
		return ""
	}

	unquoted, err := strconv.Unquote(tag.Value)
	if err != nil {
		return ""
	}

	name, _, _ := strings.Cut(reflect.StructTag(unquoted).Get("json"), ",")

	return name
}

// hasOmitEmpty reports whether a field's JSON tag carries omitempty. It is the other half of an
// absence contract a pointer opens: without it a re-marshaled record claims the node with a null
// the controller never sent.
func hasOmitEmpty(tag *ast.BasicLit) bool {
	if tag == nil {
		return false
	}

	unquoted, err := strconv.Unquote(tag.Value)
	if err != nil {
		return false
	}

	_, options, _ := strings.Cut(reflect.StructTag(unquoted).Get("json"), ",")

	return slices.Contains(strings.Split(options, ","), "omitempty")
}

// topTags returns the tags a decode type declares at its top level, which is where the sole
// member of the envelope has to be claimed.
func topTags(fields []tagged) []tagged {
	var top []tagged

	for _, field := range fields {
		if field.level == 0 && field.tag != "" && field.tag != "-" {
			top = append(top, field)
		}
	}

	return top
}

// decodeTypes lists the types the package decodes into, once each and in a fixed order: a type
// several reads share would otherwise be reported once per read.
func decodeTypes(pkg *servicePkg) []string {
	seen := make(map[string]bool, len(pkg.sites))
	for _, s := range pkg.sites {
		seen[s.decode] = true
	}

	return slices.Sorted(maps.Keys(seen))
}

// resolveEndpoint resolves the endpoint expression of a read to the route it names. Three shapes
// occur under service/: a route constant, a local variable the same function assigns from a
// RESTCONF URL builder, and a package-local helper that returns such a call. The builders take
// the route as their first argument and append the list key to it, so the first argument is the
// route the read asks for.
func resolveEndpoint(
	pkg *servicePkg,
	routes map[string]string,
	fn *ast.FuncDecl,
	expr ast.Expr,
	hops int,
) (string, bool) {
	if hops > maxHops {
		return "", false
	}

	switch node := expr.(type) {
	case *ast.ParenExpr:
		return resolveEndpoint(pkg, routes, fn, node.X, hops+1)

	case *ast.BasicLit:
		return stringLit(node)

	case *ast.BinaryExpr:
		if node.Op != token.ADD {
			return "", false
		}
		left, leftOK := resolveEndpoint(pkg, routes, fn, node.X, hops+1)
		right, rightOK := resolveEndpoint(pkg, routes, fn, node.Y, hops+1)
		if !leftOK || !rightOK {
			return "", false
		}

		return left + right, true

	case *ast.SelectorExpr:
		qualifier, ok := node.X.(*ast.Ident)
		if !ok || qualifier.Name != routesPkg {
			return "", false
		}
		route, ok := routes[node.Sel.Name]

		return route, ok

	case *ast.Ident:
		assigned, ok := assignedIn(fn, node.Name)
		if !ok {
			return "", false
		}

		return resolveEndpoint(pkg, routes, fn, assigned, hops+1)

	case *ast.CallExpr:
		return resolveCall(pkg, routes, fn, node, hops)
	}

	return "", false
}

// resolveCall resolves a call that produces an endpoint: either a RESTCONF URL builder, whose
// first argument is the route, or a helper in the same package that returns one.
func resolveCall(
	pkg *servicePkg,
	routes map[string]string,
	fn *ast.FuncDecl,
	call *ast.CallExpr,
	hops int,
) (string, bool) {
	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return "", false
	}

	if isURLBuilder(sel.Sel.Name) && len(call.Args) > 0 {
		return resolveEndpoint(pkg, routes, fn, call.Args[0], hops+1)
	}

	helper, ok := pkg.funcs[sel.Sel.Name]
	if !ok {
		return "", false
	}

	returned, ok := soleReturn(helper)
	if !ok {
		return "", false
	}

	return resolveEndpoint(pkg, routes, helper, returned, hops+1)
}

// isURLBuilder reports whether name is a RESTCONF builder that takes its route first
// (internal/restconf/builder.go).
func isURLBuilder(name string) bool {
	return name == "BuildQueryURL" || name == "BuildQueryCompositeURL"
}

// assignedIn returns the expression a function assigns to name, which is where a keyed read holds
// the URL its builder produced.
func assignedIn(fn *ast.FuncDecl, name string) (ast.Expr, bool) {
	if fn == nil || fn.Body == nil {
		return nil, false
	}

	var found ast.Expr

	ast.Inspect(fn.Body, func(node ast.Node) bool {
		assign, ok := node.(*ast.AssignStmt)
		if !ok || len(assign.Lhs) != len(assign.Rhs) {
			return true
		}

		for i, lhs := range assign.Lhs {
			if ident, ok := lhs.(*ast.Ident); ok && ident.Name == name {
				found = assign.Rhs[i]
			}
		}

		return true
	})

	return found, found != nil
}

// soleReturn returns the expression a helper returns, when it returns exactly one from exactly
// one statement. A helper written any other way is left unresolved rather than guessed at.
func soleReturn(fn *ast.FuncDecl) (ast.Expr, bool) {
	if fn.Body == nil || len(fn.Body.List) != 1 {
		return nil, false
	}

	ret, ok := fn.Body.List[0].(*ast.ReturnStmt)
	if !ok || len(ret.Results) != 1 {
		return nil, false
	}

	return ret.Results[0], true
}

// evalConst folds a constant expression into the string it denotes, following identifiers through
// the constants collected with it. Anything that is not a string, or that names a constant from
// elsewhere, is reported unresolved.
func evalConst(consts map[string]ast.Expr, expr ast.Expr, hops int) (string, bool) {
	if hops > maxHops {
		return "", false
	}

	switch node := expr.(type) {
	case *ast.ParenExpr:
		return evalConst(consts, node.X, hops+1)

	case *ast.BasicLit:
		return stringLit(node)

	case *ast.BinaryExpr:
		if node.Op != token.ADD {
			return "", false
		}
		left, leftOK := evalConst(consts, node.X, hops+1)
		right, rightOK := evalConst(consts, node.Y, hops+1)
		if !leftOK || !rightOK {
			return "", false
		}

		return left + right, true

	case *ast.Ident:
		referenced, ok := consts[node.Name]
		if !ok {
			return "", false
		}

		return evalConst(consts, referenced, hops+1)
	}

	return "", false
}

// stringLit returns the string a literal denotes. A route is built from string literals alone, so
// a literal of any other kind leaves the expression unresolved.
func stringLit(lit *ast.BasicLit) (string, bool) {
	if lit.Kind != token.STRING {
		return "", false
	}

	unquoted, err := strconv.Unquote(lit.Value)

	return unquoted, err == nil
}

// tagMatchesRoute reports whether tag is the envelope key a read of route answers with.
func tagMatchesRoute(tag, route string) bool {
	module, local, qualified := strings.Cut(tag, ":")

	return qualified && module == moduleOfRoute(route) && local == nodeName(route)
}

// moduleOfRoute returns the module the route names. Testing the tag's module for non-emptiness
// only, which this predicate did before, let a decode type name the right node under the wrong
// module and still pass: a type tagged rrm-oper:rogue-stats on the rogue-oper route was accepted.
// Every route constant in this tree names exactly one module, at its data root, and RFC 7951 4
// qualifies the sole top-level member with the module that defines the node, so the two are the
// same string and comparing them is exact rather than heuristic.
func moduleOfRoute(route string) string {
	trimmed, _, _ := strings.Cut(route, "?")
	trimmed, _, _ = strings.Cut(trimmed, "=")

	module := ""

	for segment := range strings.SplitSeq(trimmed, "/") {
		if prefix, _, qualified := strings.Cut(segment, ":"); qualified {
			module = prefix
		}
	}

	return module
}

// nodeName returns the YANG node a route reads. The query and the list key are cut before the
// last segment is taken: a key can hold both "/" and ":", so taking the segment or the module
// prefix first would read part of the key as the node name.
func nodeName(route string) string {
	trimmed, _, _ := strings.Cut(route, "?")
	trimmed, _, _ = strings.Cut(trimmed, "=")

	segment := path.Base(trimmed)
	if _, local, qualified := strings.Cut(segment, ":"); qualified {
		return local
	}

	return segment
}

// moduleOf returns the YANG module prefix of a module-qualified name, or "" when it has none.
func moduleOf(name string) string {
	module, _, found := strings.Cut(name, ":")
	if !found {
		return ""
	}

	return module
}

// walkTags collects every module-qualified tag that repeats the module already in force at its
// parent node. RFC 7951 4 qualifies a child only when its module differs from its parent's, so a
// repeat is a key the wire never sends at that depth while a genuine cross-module child is left
// alone. inForce is indexed by nesting level, because an anonymous nested struct contributes its
// fields to the same slice.
func walkTags(types map[string][]tagged, name, parent string, level int, seen map[string]bool, out *[]string) {
	if name == "" || seen[name] {
		return
	}

	seen[name] = true
	defer delete(seen, name)

	inForce := []string{parent}

	for _, field := range types[name] {
		inForce = inForce[:min(field.level+1, len(inForce))]
		module := moduleOf(field.tag)

		if module != "" && level+field.level > 0 && module == inForce[len(inForce)-1] {
			*out = append(*out, name+"."+field.tag)
		}

		child := inForce[len(inForce)-1]
		if module != "" {
			child = module
		}
		inForce = append(inForce, child)

		next := level + field.level
		if field.tag != "" {
			next++
		}

		walkTags(types, field.typeName, child, next, seen, out)
	}
}
