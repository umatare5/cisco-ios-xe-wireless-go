package mobility

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/mobility"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/validation"
)

// operOps provides high-level operational operations for mobility service
func (s Service) operOps() *core.OperationalOperations[model.MobilityOper] {
	return core.NewOperationalOperations[model.MobilityOper](s.Client(), routes.MobilityOperBasePath)
}

// GetOper retrieves all mobility operational data.
func (s Service) GetOper(ctx context.Context) (*model.MobilityOper, error) {
	return s.operOps().GetAll(ctx)
}

// GetOperApCache retrieves AP cache data.
func (s Service) GetOperApCache(ctx context.Context) (*model.MobilityOperApCache, error) {
	return core.Get[model.MobilityOperApCache](ctx, s.Client(), routes.MobilityOperApCacheEndpoint)
}

// GetOperApPeerList retrieves AP peer list data.
func (s Service) GetOperApPeerList(ctx context.Context) (*model.MobilityOperApPeerList, error) {
	return core.Get[model.MobilityOperApPeerList](ctx, s.Client(), routes.MobilityOperApPeerListEndpoint)
}

// GetOperMmGlobalData retrieves MM global data.
func (s Service) GetOperMmGlobalData(ctx context.Context) (*model.MobilityOperMmGlobalData, error) {
	return core.Get[model.MobilityOperMmGlobalData](ctx, s.Client(), routes.MobilityOperMmGlobalDataEndpoint)
}

// GetOperMmIfGlobalStats retrieves MM interface global statistics.
func (s Service) GetOperMmIfGlobalStats(ctx context.Context) (*model.MobilityOperMmIfGlobalStats, error) {
	return core.Get[model.MobilityOperMmIfGlobalStats](
		ctx, s.Client(), routes.MobilityOperMmIfGlobalStatsEndpoint)
}

// GetOperMobilityClientData retrieves mobility client data.
func (s Service) GetOperMobilityClientData(ctx context.Context) (*model.MobilityOperMobilityClientData, error) {
	return core.Get[model.MobilityOperMobilityClientData](
		ctx, s.Client(), routes.MobilityOperMobilityClientDataEndpoint)
}

// GetOperMobilityGlobalStats retrieves mobility global statistics.
func (s Service) GetOperMobilityGlobalStats(
	ctx context.Context,
) (*model.MobilityOperMobilityGlobalStats, error) {
	return core.Get[model.MobilityOperMobilityGlobalStats](
		ctx, s.Client(), routes.MobilityOperMobilityGlobalStatsEndpoint,
	)
}

// GetOperByAPMac retrieves mobility operational data filtered by Access Point MAC address.
func (s Service) GetOperByAPMac(ctx context.Context, apMac string) (*model.MobilityOper, error) {
	if err := validation.ValidateAPMac(apMac); err != nil {
		return nil, err
	}
	return s.operOps().GetByCompositeKey(ctx, "ap-cache", apMac)
}

// GetOperByClientMAC retrieves mobility operational data filtered by client MAC address.
func (s Service) GetOperByClientMAC(ctx context.Context, clientMAC string) (*model.MobilityOper, error) {
	if err := validation.ValidateAPMac(clientMAC); err != nil {
		return nil, err
	}
	return s.operOps().GetByCompositeKey(ctx, "mobility-client-data", clientMAC)
}
