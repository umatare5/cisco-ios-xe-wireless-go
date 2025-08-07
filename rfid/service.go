package rfid

import (
	"context"
	"net/http"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/constants"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
)

const (
	// RfidCfgBasePath defines the base path for RFID configuration endpoints
	RfidCfgBasePath = constants.YANGModelPrefix + "rfid-cfg:rfid-cfg-data"
	// RfidCfgEndpoint defines the endpoint for RFID configuration data
	RfidCfgEndpoint = RfidCfgBasePath
)

// Service provides RFID operations.
type Service struct {
	c *core.Client
}

// NewService creates a new service instance.
func NewService(client *core.Client) Service {
	return Service{c: client}
}

// Cfg returns RFID configuration data.
func (s Service) Cfg(ctx context.Context) (*model.RfidCfgResponse, error) {
	var result model.RfidCfgResponse
	return &result, s.c.Do(ctx, http.MethodGet, RfidCfgEndpoint, &result)
}
