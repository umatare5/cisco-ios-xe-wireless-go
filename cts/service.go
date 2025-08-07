package cts

import (
	"context"
	"net/http"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/constants"
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

// Cfg returns CTS configuration data.
func (s Service) Cfg(ctx context.Context) (*model.CtsCfgResponse, error) {
	var result model.CtsCfgResponse
	return &result, s.c.Do(ctx, http.MethodGet, CtsCfgEndpoint, &result)
}
