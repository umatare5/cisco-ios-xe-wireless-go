package mobility

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/constants"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
)

const (
	// MobilityOperBasePath defines the base path for mobility operational data endpoints
	MobilityOperBasePath = constants.YANGModelPrefix + "mobility-oper:mobility-oper-data"
	// MobilityOperEndpoint defines the endpoint for mobility operational data
	MobilityOperEndpoint = MobilityOperBasePath
)

// Service provides Mobility operations.
type Service struct {
	c *core.Client
}

// NewService creates a new service instance.
func NewService(client *core.Client) Service {
	return Service{c: client}
}

// GetOper returns Mobility operational data.
func (s Service) GetOper(ctx context.Context) (*model.MobilityOperResponse, error) {
	return core.Get[model.MobilityOperResponse](ctx, s.c, MobilityOperEndpoint)
}
