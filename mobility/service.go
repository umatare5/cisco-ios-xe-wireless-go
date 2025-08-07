package mobility

import (
	"context"
	"net/http"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/constants"
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

// Oper returns Mobility operational data.
func (s Service) Oper(ctx context.Context) (*model.MobilityOperResponse, error) {
	var result model.MobilityOperResponse
	return &result, s.c.Do(ctx, http.MethodGet, MobilityOperEndpoint, &result)
}
