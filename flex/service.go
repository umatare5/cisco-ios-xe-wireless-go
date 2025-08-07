package flex

import (
	"context"
	"net/http"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
)

const (
	// FlexCfgBasePath defines the base path for FlexConnect configuration endpoints
	FlexCfgBasePath = "Cisco-IOS-XE-wireless-flex-cfg:flex-cfg-data"
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

// Cfg returns FlexConnect configuration data.
func (s Service) Cfg(ctx context.Context) (*model.FlexCfgResponse, error) {
	var result model.FlexCfgResponse
	return &result, s.c.Do(ctx, http.MethodGet, FlexCfgEndpoint, &result)
}
