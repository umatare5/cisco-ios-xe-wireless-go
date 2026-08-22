package contract_test

import (
	"go/parser"
	"go/token"
	"maps"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"
)

// docFile is the file a package's own documentation belongs in. Every package in this tree that
// documents itself at all puts its summary here.
const docFile = "doc.go"

// TestOnlyDocGoCarriesThePackageComment holds the one property that keeps a rendered package doc
// readable: godoc concatenates every package comment in a package, in file order, so a second one
// appends a second summary to the page and the reader finishes on whichever file sorts last.
//
// The rule applies only where a doc.go exists. A package whose only package comment sits in another
// file is documented, not defective, and condemning it would delete documentation to satisfy a
// naming convention.
//
// Nothing in the toolchain reports this: staticcheck's ST1000 finds a package with no comment at all
// and says nothing about a package with two.
func TestOnlyDocGoCarriesThePackageComment(t *testing.T) {
	byDir := documentedPackages(t)

	if len(byDir) == 0 {
		t.Fatal("no package with a doc.go was found, so a pass here would mean nothing")
	}

	examined := 0

	for _, dir := range slices.Sorted(maps.Keys(byDir)) {
		for _, rel := range byDir[dir] {
			if filepath.Base(rel) == docFile {
				continue
			}

			examined++

			pkg, documented := packageComment(t, rel)
			if !documented {
				continue
			}

			t.Errorf("%s: carries a package comment beside %s, so godoc renders two summaries for"+
				" package %s and the page ends on whichever file sorts last",
				rel, filepath.Join(dir, docFile), pkg)
		}
	}

	t.Logf("examined %d non-doc.go file(s) across %d package(s) that have a %s",
		examined, len(byDir), docFile)
}

// documentedPackages lists the Go files of every directory holding a doc.go, keyed by that
// directory. Test files are included: a package comment in one renders the same as in any other.
func documentedPackages(t *testing.T) map[string][]string {
	t.Helper()

	byDir := make(map[string][]string)

	err := filepath.WalkDir(repoRoot, func(p string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			// tmp holds worktrees of this same repository, which would be counted twice.
			if d.Name() == "tmp" || d.Name() == ".git" {
				return filepath.SkipDir
			}

			return nil
		}
		if !strings.HasSuffix(p, ".go") {
			return nil
		}

		rel, err := filepath.Rel(repoRoot, p)
		if err != nil {
			return err
		}
		dir := filepath.Dir(rel)
		byDir[dir] = append(byDir[dir], rel)

		return nil
	})
	if err != nil {
		t.Fatalf("walking %s: %v", repoRoot, err)
	}

	for dir, files := range byDir {
		if !slices.ContainsFunc(files, func(f string) bool { return filepath.Base(f) == docFile }) {
			delete(byDir, dir)
		}
	}

	return byDir
}

// packageComment returns the file's package name and whether it attaches a doc comment to its
// package clause. The parser is the arbiter rather than a line match: a comment separated from the
// clause by a blank line is not a package comment, and godoc does not render it.
func packageComment(t *testing.T, rel string) (string, bool) {
	t.Helper()

	fset := token.NewFileSet()

	const mode = parser.ParseComments | parser.SkipObjectResolution

	file, err := parser.ParseFile(fset, filepath.Join(repoRoot, rel), nil, mode)
	if err != nil {
		t.Fatalf("parsing %s: %v", rel, err)
	}

	return file.Name.Name, file.Doc != nil
}
