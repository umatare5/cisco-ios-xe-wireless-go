package rfid

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/constants"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
)

const (
	// RFIDCfgBasePath defines the base path for RFID configuration endpoints
	RFIDCfgBasePath = constants.YANGModelPrefix + "rfid-cfg:rfid-cfg-data"
	// RFIDCfgEndpoint defines the endpoint for RFID configuration data
	RFIDCfgEndpoint = RFIDCfgBasePath
)

// Service provides RFID operations.
type Service struct {
	c *core.Client
}

// NewService creates a new service instance.
func NewService(client *core.Client) Service {
	return Service{c: client}
}

// GetCfg returns RFID configuration data.
func (s Service) GetCfg(ctx context.Context) (*model.RfidCfgResponse, error) {
	return core.Get[model.RfidCfgResponse](ctx, s.c, RFIDCfgEndpoint)
}
