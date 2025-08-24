package geolocation

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/geolocation"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
)

// operOps provides high-level operational operations for geolocation service
func (s Service) operOps() *core.OperationalOperations[model.GeolocationOper] {
	return core.NewOperationalOperations[model.GeolocationOper](s.Client(), routes.GeolocationOperBasePath)
}

// GetOper retrieves geolocation operational data for geographic positioning and location mapping.
func (s Service) GetOper(ctx context.Context) (*model.GeolocationOper, error) {
	return s.operOps().GetAll(ctx)
}

// GetOperApGeoLocStats retrieves AP geolocation statistics.
func (s Service) GetOperApGeoLocStats(ctx context.Context) (*model.GeolocationOperApGeoLocStats, error) {
	return core.Get[model.GeolocationOperApGeoLocStats](ctx, s.Client(), routes.ApGeoLocStatsEndpoint)
}
