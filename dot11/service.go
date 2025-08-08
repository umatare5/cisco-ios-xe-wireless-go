package dot11

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/constants"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
)

const (
	// Dot11CfgBasePath defines the base path for 802.11 configuration endpoints
	Dot11CfgBasePath = constants.YANGModelPrefix + "dot11-cfg:dot11-cfg-data"
	// Dot11CfgEndpoint defines the endpoint for 802.11 configuration data
	Dot11CfgEndpoint = Dot11CfgBasePath
)

// Service provides 802.11 operations.
type Service struct {
	c *core.Client
}

// NewService creates a new service instance.
func NewService(client *core.Client) Service {
	return Service{c: client}
}

// GetCfg returns 802.11 configuration data.
func (s Service) GetCfg(ctx context.Context) (*model.Dot11CfgResponse, error) {
	return core.Get[model.Dot11CfgResponse](ctx, s.c, Dot11CfgEndpoint)
}
