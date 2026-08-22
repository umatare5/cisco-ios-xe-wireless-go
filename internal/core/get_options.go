package core

import (
	"strconv"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf"
)

// DefaultsMode selects the RFC 6243 retrieval mode for the with-defaults parameter.
type DefaultsMode string

const (
	// DefaultsReportAll asks the server to report the data nodes it holds for the
	// record, including those in force at their schema default.
	DefaultsReportAll DefaultsMode = "report-all"

	// DefaultsExplicit asks for the leaves a client has explicitly set, including any a
	// client set to its schema default value (RFC 6243 3.3). This is not the same answer
	// as a plain read, which omits every leaf whose value equals its default whoever set it.
	DefaultsExplicit DefaultsMode = "explicit"
)

// GetOption customizes a single GET request.
type GetOption func(*getConfig)

type getConfig struct {
	defaultsMode DefaultsMode
	fields       string
	depth        int
}

// WithDefaults requests the given with-defaults retrieval mode.
func WithDefaults(mode DefaultsMode) GetOption {
	return func(g *getConfig) { g.defaultsMode = mode }
}

// WithFields limits the answer to the RFC 8040 4.8.3 fields expression given, for
// example "wlan-cfg-entries/wlan-cfg-entry(profile-name;wlan-id)".
//
// A leaf the expression prunes is absent from the answer, and a non-pointer field
// decodes an absent leaf as its zero value, so prune only what the caller reads: a
// pruned counter is indistinguishable from a counter reading zero.
func WithFields(expression string) GetOption {
	return func(g *getConfig) { g.fields = expression }
}

// WithDepth limits the answer to the RFC 8040 4.8.2 subtree depth given (1 to 65535).
// A value below 1 is ignored, as an unset mode is.
//
// A node the depth limit cuts is absent from the answer exactly as a pruned leaf is,
// so the same hazard applies: a non-pointer field decodes an absent leaf as its zero
// value, and a leaf cut off by depth is indistinguishable from a leaf reading zero.
func WithDepth(levels int) GetOption {
	return func(g *getConfig) { g.depth = levels }
}

// applyGetOptions folds opts into endpoint, returning the endpoint to request.
func applyGetOptions(endpoint string, opts []GetOption) string {
	var cfg getConfig
	for _, opt := range opts {
		if opt != nil {
			opt(&cfg)
		}
	}

	if cfg.defaultsMode != "" {
		endpoint = restconf.AppendQueryParam(endpoint, restconf.WithDefaultsParam, string(cfg.defaultsMode))
	}
	if cfg.fields != "" {
		endpoint = restconf.AppendQueryParam(endpoint, restconf.FieldsParam, cfg.fields)
	}
	if cfg.depth > 0 {
		endpoint = restconf.AppendQueryParam(endpoint, restconf.DepthParam, strconv.Itoa(cfg.depth))
	}

	return endpoint
}
