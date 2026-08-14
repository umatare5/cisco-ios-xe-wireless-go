package core

import "github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf"

// DefaultsMode selects the RFC 6243 retrieval mode for the with-defaults parameter.
type DefaultsMode string

const (
	// DefaultsReportAll asks the server to report the data nodes it holds for the
	// record, including those in force at their schema default.
	DefaultsReportAll DefaultsMode = "report-all"

	// DefaultsExplicit asks for client-set leaves only, which matches a plain GET.
	DefaultsExplicit DefaultsMode = "explicit"
)

// GetOption customizes a single GET request.
type GetOption func(*getConfig)

type getConfig struct {
	defaultsMode DefaultsMode
}

// WithDefaults requests the given with-defaults retrieval mode.
func WithDefaults(mode DefaultsMode) GetOption {
	return func(g *getConfig) { g.defaultsMode = mode }
}

// applyGetOptions folds opts into endpoint, returning the endpoint to request.
func applyGetOptions(endpoint string, opts []GetOption) string {
	var cfg getConfig
	for _, opt := range opts {
		if opt != nil {
			opt(&cfg)
		}
	}
	if cfg.defaultsMode == "" {
		return endpoint
	}
	return restconf.AppendQueryParam(endpoint, restconf.WithDefaultsParam, string(cfg.defaultsMode))
}
