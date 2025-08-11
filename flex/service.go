package flex

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/constants"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
)

const (
	// FlexCfgBasePath defines the base path for FlexConnect configuration endpoints
	FlexCfgBasePath = constants.YANGModelPrefix + "flex-cfg:flex-cfg-data"
	// FlexCfgEndpoint defines the endpoint for FlexConnect configuration data
	FlexCfgEndpoint = FlexCfgBasePath
)

// Service provides FlexConnect operations.
type Service struct {
	c *core.Client
}

// NewService creates a new service instance.
func NewService(client *core.Client) Service {
	return Service{c: client}
}

// GetCfg returns FlexConnect configuration data.
func (s Service) GetCfg(ctx context.Context) (*model.FlexCfgResponse, error) {
	return core.Get[model.FlexCfgResponse](ctx, s.c, FlexCfgEndpoint)
}
