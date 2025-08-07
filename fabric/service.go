package fabric

import (
	"context"
	"net/http"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/constants"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
)

const (
	// FabricCfgBasePath defines the base path for fabric configuration endpoints
	FabricCfgBasePath = constants.YANGModelPrefix + "fabric-cfg:fabric-cfg-data"
	// FabricCfgEndpoint defines the endpoint for fabric configuration data
	FabricCfgEndpoint = FabricCfgBasePath
)

// Service provides Fabric operations.
type Service struct {
	c *core.Client
}

// NewService creates a new service instance.
func NewService(client *core.Client) Service {
	return Service{c: client}
}

// Cfg returns Fabric configuration data.
func (s Service) Cfg(ctx context.Context) (*model.FabricCfgResponse, error) {
	var result model.FabricCfgResponse
	return &result, s.c.Do(ctx, http.MethodGet, FabricCfgEndpoint, &result)
}
