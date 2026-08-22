package core

import (
	"encoding/json"
	"fmt"
	"path"
	"reflect"
	"strings"
)

// decodeSoleKey decodes a read of one RESTCONF node into T, having checked that the body is
// the node the endpoint asked for and that T can consume it.
//
// RFC 7951 answers a container or list read with exactly one top-level member, qualified by
// the module that defines it. json.Unmarshal checks neither half of that: a body keyed for
// another node, and a T that declares no field for the key, both leave every field at its
// zero value and report success, which reaches the caller as a controller holding no data.
//
// Only the top level is checked here. A module-qualified tag below the top level is a source
// defect, not a response defect, and is caught statically by tests/contract.
func decodeSoleKey[T any](body []byte, endpoint string) (*T, error) {
	var envelope map[string]json.RawMessage
	if err := json.Unmarshal(body, &envelope); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	key, err := soleKey(envelope, endpoint)
	if err != nil {
		return nil, err
	}

	target := reflect.TypeFor[T]()
	if !declaresJSONField(target, key) {
		return nil, fmt.Errorf("%s: %s declares no field for response key %q", endpoint, target, key)
	}

	var out T
	if err := json.Unmarshal(body, &out); err != nil {
		return nil, fmt.Errorf("%s: failed to unmarshal response: %w", endpoint, err)
	}

	return &out, nil
}

// soleKey returns the one top-level key of a response body, having checked that it is
// module-qualified and names the node the endpoint asked for.
//
// The expectation comes from the endpoint, not from T's own struct tag: the tag is the
// artifact under test, so a tag naming the wrong node would validate against itself.
func soleKey(envelope map[string]json.RawMessage, endpoint string) (string, error) {
	if len(envelope) != 1 {
		return "", fmt.Errorf("%s: response carries %d top-level keys, want exactly 1", endpoint, len(envelope))
	}

	var key string
	for k := range envelope {
		key = k
	}

	want := nodeName(endpoint)
	module, local, qualified := strings.Cut(key, ":")
	if !qualified || module == "" || local != want {
		return "", fmt.Errorf("%s: response carries key %q, want a module-qualified %q", endpoint, key, want)
	}

	return key, nil
}

// nodeName returns the YANG node name the endpoint reads.
//
// The query and the list key are cut before the last segment is taken: a list key can hold
// both "/" and ":", so taking the segment or the module prefix first would read part of the
// key as the node name.
func nodeName(endpoint string) string {
	trimmed, _, _ := strings.Cut(endpoint, "?")
	trimmed, _, _ = strings.Cut(trimmed, "=")

	segment := path.Base(trimmed)
	if _, local, qualified := strings.Cut(segment, ":"); qualified {
		return local
	}

	return segment
}

// declaresJSONField reports whether t declares a JSON field named key.
//
// encoding/json drops a member no field claims, so this is what separates a response type
// pointed at the wrong node from one that decoded a node holding nothing.
func declaresJSONField(t reflect.Type, key string) bool {
	if t.Kind() != reflect.Struct {
		return false
	}

	for _, field := range reflect.VisibleFields(t) {
		if !field.IsExported() {
			continue
		}
		if name, _, _ := strings.Cut(field.Tag.Get("json"), ","); name == key {
			return true
		}
	}

	return false
}
