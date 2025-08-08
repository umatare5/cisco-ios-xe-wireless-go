package apf

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/constants"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
)

const (
	// ApfCfgBasePath defines the base path for APF configuration endpoints
	ApfCfgBasePath = constants.YANGModelPrefix + "apf-cfg:apf-cfg-data"
	// ApfCfgEndpoint defines the endpoint for APF configuration data
	ApfCfgEndpoint = ApfCfgBasePath
)

// Service provides APF operations.
type Service struct {
	c *core.Client
}

// NewService creates a new service instance.
func NewService(client *core.Client) Service {
	return Service{c: client}
}

// Cfg returns APF configuration data.
func (s Service) Cfg(ctx context.Context) (*model.ApfCfgResponse, error) {
	return core.Get[model.ApfCfgResponse](ctx, s.c, ApfCfgEndpoint)
}
