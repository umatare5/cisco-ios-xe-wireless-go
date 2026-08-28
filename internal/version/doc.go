// Package version holds this module's version as a compile-time constant.
//
// It is its own package because internal/transport composes the default User-Agent from it, and
// internal/transport sits below the root package in the import graph — the root imports
// internal/core, which imports internal/transport — so it cannot read a constant the root declares.
// It is not that internal/transport imports stdlib only; it already imports internal/restconf.
// This package imports nothing at all, so importing it can make no cycle from anywhere.
package version
