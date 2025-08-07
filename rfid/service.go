package rfid

import (
	"context"
	"net/http"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
	"github.com/umatare5/cisco-ios-xe-wireless-go/wnc"
)

const (
	// RfidCfgBasePath defines the base path for RFID configuration endpoints
	RfidCfgBasePath = "Cisco-IOS-XE-wireless-rfid-cfg:rfid-cfg-data"
	// RfidCfgEndpoint defines the endpoint for RFID configuration data
	RfidCfgEndpoint = RfidCfgBasePath
)

// Service provides RFID operations.
type Service struct {
	c *wnc.Client
}

// NewService creates a new service instance.
func NewService(client *wnc.Client) Service {
	return Service{c: client}
}

// Cfg returns RFID configuration data.
func (s Service) Cfg(ctx context.Context) (*model.RfidCfgResponse, error) {
	var result model.RfidCfgResponse
	return &result, s.c.Do(ctx, http.MethodGet, RfidCfgEndpoint, &result)
}
