package rf

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/constants"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
)

const (
	// RfCfgBasePath defines the base path for RF configuration endpoints
	RfCfgBasePath = constants.YANGModelPrefix + "rf-cfg:rf-cfg-data"
	// RfCfgEndpoint defines the endpoint for RF configuration data
	RfCfgEndpoint = RfCfgBasePath
)

// Service provides RF operations.
type Service struct {
	c *core.Client
}

// NewService creates a new service instance.
func NewService(client *core.Client) Service {
	return Service{c: client}
}

// GetCfg returns RF configuration data.
func (s Service) GetCfg(ctx context.Context) (*model.RfCfgResponse, error) {
	return core.Get[model.RfCfgResponse](ctx, s.c, RfCfgEndpoint)
}
