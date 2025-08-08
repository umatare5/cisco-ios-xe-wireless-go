package cts

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/constants"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
)

const (
	// CtsCfgBasePath defines the base path for CTS configuration endpoints
	CtsCfgBasePath = constants.YANGModelPrefix + "cts-cfg:cts-cfg-data"
	// CtsCfgEndpoint defines the endpoint for CTS configuration data
	CtsCfgEndpoint = CtsCfgBasePath
)

// Service provides CTS operations.
type Service struct {
	c *core.Client
}

// NewService creates a new service instance.
func NewService(client *core.Client) Service {
	return Service{c: client}
}

// GetCfg returns CTS configuration data.
func (s Service) GetCfg(ctx context.Context) (*model.CtsCfgResponse, error) {
	return core.Get[model.CtsCfgResponse](ctx, s.c, CtsCfgEndpoint)
}
