package dot11

import (
	"context"
	"fmt"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/dot11"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/validation"
)

// configOps provides high-level configuration operations for dot11 service
func (s Service) configOps() *core.ConfigOperations[model.Dot11Cfg] {
	return core.NewConfigOperations[model.Dot11Cfg](s.Client(), routes.Dot11CfgBasePath)
}

// GetCfg retrieves 802.11 configuration data.
func (s Service) GetCfg(ctx context.Context) (*model.Dot11Cfg, error) {
	return s.configOps().GetAll(ctx)
}

// GetCfgByCountryCode retrieves 802.11 configuration data filtered by country code.
func (s Service) GetCfgByCountryCode(ctx context.Context, countryCode string) (*model.Dot11CfgFilter, error) {
	if err := s.ValidateClient(); err != nil {
		return nil, err
	}

	if err := validation.ValidateNonEmptyString(countryCode, "country code"); err != nil {
		return nil, err
	}

	url := s.Client().RestconfBuilder().BuildPathQueryURL(
		routes.Dot11CfgConfiguredCountriesEndpoint, "configured-country", countryCode)
	return core.Get[model.Dot11CfgFilter](ctx, s.Client(), url)
}

// GetCfgByBand retrieves 802.11 configuration data filtered by band.
func (s Service) GetCfgByBand(ctx context.Context, band string) (*model.Dot11CfgFilter, error) {
	if err := s.ValidateClient(); err != nil {
		return nil, err
	}

	if err := validation.ValidateNonEmptyString(band, "band"); err != nil {
		return nil, err
	}

	url := s.Client().RestconfBuilder().BuildPathQueryURL(routes.Dot11CfgDot11EntriesEndpoint, "dot11-entry", band)
	return core.Get[model.Dot11CfgFilter](ctx, s.Client(), url)
}

// GetCfgBySpatialStreamAndIndex retrieves 802.11ac MCS configuration data filtered by spatial stream and index.
func (s Service) GetCfgBySpatialStreamAndIndex(
	ctx context.Context,
	spatialStream int,
	index string,
) (*model.Dot11CfgFilter, error) {
	if err := validation.ValidateSpatialStream(spatialStream); err != nil {
		return nil, err
	}

	if err := validation.ValidateNonEmptyString(index, "index"); err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/dot11ac-mcs-entries/dot11ac-mcs-entry=%d,%s", routes.Dot11CfgBasePath, spatialStream, index)
	return core.Get[model.Dot11CfgFilter](ctx, s.Client(), url)
}
