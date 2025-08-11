package location

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/constants"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
)

const (
	// LocationCfgBasePath defines the base path for location configuration endpoints
	LocationCfgBasePath = constants.YANGModelPrefix + "location-cfg:location-cfg-data"
	// LocationCfgEndpoint defines the endpoint for location configuration data
	LocationCfgEndpoint = LocationCfgBasePath
)

// Service provides Location operations.
type Service struct {
	c *core.Client
}

// NewService creates a new service instance.
func NewService(client *core.Client) Service {
	return Service{c: client}
}

// GetCfg returns Location configuration data.
func (s Service) GetCfg(ctx context.Context) (*model.LocationCfgResponse, error) {
	return core.Get[model.LocationCfgResponse](ctx, s.c, LocationCfgEndpoint)
}
