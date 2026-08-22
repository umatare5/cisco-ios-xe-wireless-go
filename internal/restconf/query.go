package restconf

import "strings"

// RFC 8040 query parameter names.
const (
	// WithDefaultsParam is the RFC 8040 4.8.9 query parameter name.
	WithDefaultsParam = "with-defaults"
	// FieldsParam is the RFC 8040 4.8.3 query parameter name.
	FieldsParam = "fields"
	// DepthParam is the RFC 8040 4.8.2 query parameter name.
	DepthParam = "depth"
)

// upperhex is the alphabet of a percent-encoded byte.
const upperhex = "0123456789ABCDEF"

// AppendQueryParam appends name=value to the endpoint, using "?" or "&" as needed.
func AppendQueryParam(endpointPath, name, value string) string {
	sep := "?"
	if strings.Contains(endpointPath, "?") {
		sep = "&"
	}
	return endpointPath + sep + name + "=" + escapeQueryValue(value)
}

// escapeQueryValue percent-encodes the bytes a query value must not carry literally.
//
// url.QueryEscape is the wrong escaper here: it encodes the "/", ";", "(" and ")" that
// the RFC 8040 4.8.3 fields grammar is built from, which the controller then reads as
// one node name instead of an expression.
func escapeQueryValue(value string) string {
	var escaped strings.Builder
	escaped.Grow(len(value))

	for i := range len(value) {
		c := value[i]
		if isQueryValueByte(c) {
			escaped.WriteByte(c)
			continue
		}

		escaped.WriteByte('%')
		escaped.WriteByte(upperhex[c>>4])
		escaped.WriteByte(upperhex[c&0x0f])
	}

	return escaped.String()
}

// isQueryValueByte reports whether c stands for itself in a query value: RFC 3986
// unreserved, plus the characters the RFC 8040 fields and depth grammars use as syntax
// — "/" descends, ";" separates siblings, "(" ")" nest and ":" qualifies a node with
// its module. "&", "=" and "#" are not in the set, so a value cannot open a parameter.
func isQueryValueByte(c byte) bool {
	if 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z' || '0' <= c && c <= '9' {
		return true
	}

	switch c {
	case '-', '.', '_', '~', '/', ';', '(', ')', ':':
		return true
	default:
		return false
	}
}
