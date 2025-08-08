package dot15

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/constants"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
)

const (
	// Dot15CfgBasePath defines the base path for 802.15 configuration endpoints
	Dot15CfgBasePath = constants.YANGModelPrefix + "dot15-cfg:dot15-cfg-data"
	// Dot15CfgEndpoint defines the endpoint for 802.15 configuration data
	Dot15CfgEndpoint = Dot15CfgBasePath
)

// Service provides 802.15 operations.
type Service struct {
	c *core.Client
}

// NewService creates a new service instance.
func NewService(client *core.Client) Service {
	return Service{c: client}
}

// GetCfg returns 802.15 configuration data.
func (s Service) GetCfg(ctx context.Context) (*model.Dot15CfgResponse, error) {
	return core.Get[model.Dot15CfgResponse](ctx, s.c, Dot15CfgEndpoint)
}
