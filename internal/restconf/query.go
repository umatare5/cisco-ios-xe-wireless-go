package restconf

import "strings"

// WithDefaultsParam is the RFC 8040 4.8.9 query parameter name.
const WithDefaultsParam = "with-defaults"

// AppendQueryParam appends name=value to the endpoint, using "?" or "&" as needed.
func AppendQueryParam(endpointPath, name, value string) string {
	sep := "?"
	if strings.Contains(endpointPath, "?") {
		sep = "&"
	}
	return endpointPath + sep + name + "=" + value
}
