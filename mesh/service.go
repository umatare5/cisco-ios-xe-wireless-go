package mesh

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/constants"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
)

const (
	// MeshOperBasePath defines the base path for mesh operational endpoints
	MeshOperBasePath = constants.YANGModelPrefix + "mesh-oper:mesh-oper-data"
	// MeshOperEndpoint defines the endpoint for mesh operational data
	MeshOperEndpoint = MeshOperBasePath

	// MeshCfgBasePath defines the base path for mesh configuration endpoints
	MeshCfgBasePath = constants.YANGModelPrefix + "mesh-cfg:mesh-cfg-data"
	// MeshCfgEndpoint defines the endpoint for mesh configuration data
	MeshCfgEndpoint = MeshCfgBasePath
)

// Service provides Mesh operations.
type Service struct {
	c *core.Client
}

// NewService creates a new service instance.
func NewService(client *core.Client) Service {
	return Service{c: client}
}

// GetOper returns Mesh operational data.
func (s Service) GetOper(ctx context.Context) (*model.MeshOperResponse, error) {
	return core.Get[model.MeshOperResponse](ctx, s.c, MeshOperEndpoint)
}

// GetCfg returns Mesh configuration data.
func (s Service) GetCfg(ctx context.Context) (*model.MeshCfgResponse, error) {
	return core.Get[model.MeshCfgResponse](ctx, s.c, MeshCfgEndpoint)
}
