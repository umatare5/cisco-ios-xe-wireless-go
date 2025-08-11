package radio

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/constants"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
)

const (
	// RadioCfgBasePath defines the base path for radio configuration endpoints
	RadioCfgBasePath = constants.YANGModelPrefix + "radio-cfg:radio-cfg-data"
	// RadioCfgEndpoint defines the endpoint for radio configuration data
	RadioCfgEndpoint = RadioCfgBasePath
)

// Service provides Radio operations.
type Service struct {
	c *core.Client
}

// NewService creates a new service instance.
func NewService(client *core.Client) Service {
	return Service{c: client}
}

// GetCfg returns Radio configuration data.
func (s Service) GetCfg(ctx context.Context) (*model.RadioCfgResponse, error) {
	return core.Get[model.RadioCfgResponse](ctx, s.c, RadioCfgEndpoint)
}
