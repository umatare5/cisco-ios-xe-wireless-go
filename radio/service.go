package radio

import (
	"context"
	"net/http"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
	"github.com/umatare5/cisco-ios-xe-wireless-go/wnc"
)

const (
	// RadioCfgBasePath defines the base path for radio configuration endpoints
	RadioCfgBasePath = "Cisco-IOS-XE-wireless-radio-cfg:radio-cfg-data"
	// RadioCfgEndpoint defines the endpoint for radio configuration data
	RadioCfgEndpoint = RadioCfgBasePath
)

// Service provides Radio operations.
type Service struct {
	c *wnc.Client
}

// NewService creates a new service instance.
func NewService(client *wnc.Client) Service {
	return Service{c: client}
}

// Cfg returns Radio configuration data.
func (s Service) Cfg(ctx context.Context) (*model.RadioCfgResponse, error) {
	var result model.RadioCfgResponse
	return &result, s.c.Do(ctx, http.MethodGet, RadioCfgEndpoint, &result)
}
