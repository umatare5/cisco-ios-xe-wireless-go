package mobility

import (
	"context"
	"net/http"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/model"
	"github.com/umatare5/cisco-ios-xe-wireless-go/wnc"
)

const (
	// MobilityOperBasePath defines the base path for mobility operational data endpoints
	MobilityOperBasePath = "Cisco-IOS-XE-wireless-mobility-oper:mobility-oper-data"
	// MobilityOperEndpoint defines the endpoint for mobility operational data
	MobilityOperEndpoint = MobilityOperBasePath
)

// Service provides Mobility operations.
type Service struct {
	c *wnc.Client
}

// NewService creates a new service instance.
func NewService(client *wnc.Client) Service {
	return Service{c: client}
}

// Oper returns Mobility operational data.
func (s Service) Oper(ctx context.Context) (*model.MobilityOperResponse, error) {
	var result model.MobilityOperResponse
	return &result, s.c.Do(ctx, http.MethodGet, MobilityOperEndpoint, &result)
}
