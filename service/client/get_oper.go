// Package client provides wireless client operational data operations for the Cisco IOS-XE Wireless Network Controller API.
package client

import (
	"context"
	"fmt"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	ierrors "github.com/umatare5/cisco-ios-xe-wireless-go/internal/errors"
	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/client"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
)

// operOps provides high-level operational operations for Client service
func (s Service) operOps() *core.OperationalOperations[model.ClientOper] {
	return core.NewOperationalOperations[model.ClientOper](s.Client(), routes.ClientOperBasePath)
}

// GetOper retrieves the complete client operational data
func (s Service) GetOper(ctx context.Context) (*model.ClientOper, error) {
	return s.operOps().GetAll(ctx)
}

// GetOperCommonOperData retrieves common operational data for clients
func (s Service) GetOperCommonOperData(ctx context.Context) (*model.ClientOperCommonOperData, error) {
	return core.Get[model.ClientOperCommonOperData](ctx, s.Client(), routes.CommonOperDataEndpoint)
}

// GetOperByApName retrieves client operational data filtered by AP name
func (s Service) GetOperByApName(ctx context.Context, apName string) (*model.ClientOperCommonOperData, error) {
	if apName == "" {
		return nil, ierrors.EmptyParameterError("AP name")
	}

	if err := s.ValidateClient(); err != nil {
		return nil, err
	}
	url := s.Client().RestconfBuilder().BuildQueryURL(routes.CommonOperDataEndpoint, apName)
	return core.Get[model.ClientOperCommonOperData](ctx, s.Client(), url)
}

// GetOperByClientMac retrieves client operational data filtered by client MAC address
func (s Service) GetOperByClientMac(
	ctx context.Context, clientMac string,
) (*model.ClientOperCommonOperData, error) {
	if clientMac == "" {
		return nil, ierrors.EmptyParameterError("client MAC address")
	}

	if err := s.ValidateClient(); err != nil {
		return nil, err
	}
	url := s.Client().RestconfBuilder().BuildQueryURL(routes.CommonOperDataEndpoint, clientMac)
	return core.Get[model.ClientOperCommonOperData](ctx, s.Client(), url)
}

// GetOperByClientType retrieves client operational data filtered by client type
func (s Service) GetOperByClientType(
	ctx context.Context,
	clientType string,
) (*model.ClientOperCommonOperData, error) {
	if clientType == "" {
		return nil, ierrors.EmptyParameterError("client type")
	}

	if err := s.ValidateClient(); err != nil {
		return nil, err
	}
	url := s.Client().RestconfBuilder().BuildQueryURL(routes.CommonOperDataEndpoint, clientType)
	return core.Get[model.ClientOperCommonOperData](ctx, s.Client(), url)
}

// GetOperByCoState retrieves client operational data filtered by client operational state
func (s Service) GetOperByCoState(
	ctx context.Context, coState string,
) (*model.ClientOperCommonOperData, error) {
	if coState == "" {
		return nil, ierrors.EmptyParameterError("client operational state")
	}

	if err := s.ValidateClient(); err != nil {
		return nil, err
	}
	url := s.Client().RestconfBuilder().BuildQueryURL(routes.CommonOperDataEndpoint, coState)
	return core.Get[model.ClientOperCommonOperData](ctx, s.Client(), url)
}

// GetOperByMsRadioType retrieves operational data for a specified MS Radio Type from the Cisco Catalyst 9800 WLC
func (s Service) GetOperByMsRadioType(
	ctx context.Context, msRadioType string,
) (*model.ClientOperCommonOperData, error) {
	if msRadioType == "" {
		return nil, ierrors.EmptyParameterError("MS radio type")
	}

	if err := s.ValidateClient(); err != nil {
		return nil, err
	}
	url := s.Client().RestconfBuilder().BuildQueryURL(routes.CommonOperDataEndpoint, msRadioType)
	return core.Get[model.ClientOperCommonOperData](ctx, s.Client(), url)
}

// GetOperByUsername retrieves client operational data filtered by username
func (s Service) GetOperByUsername(
	ctx context.Context, username string,
) (*model.ClientOperCommonOperData, error) {
	if username == "" {
		return nil, ierrors.EmptyParameterError("username")
	}

	if err := s.ValidateClient(); err != nil {
		return nil, err
	}
	url := s.Client().RestconfBuilder().BuildQueryURL(routes.CommonOperDataEndpoint, username)
	return core.Get[model.ClientOperCommonOperData](ctx, s.Client(), url)
}

// GetOperByWlanID retrieves client operational data filtered by WLAN ID
func (s Service) GetOperByWlanID(ctx context.Context, wlanID int) (*model.ClientOperCommonOperData, error) {
	if wlanID < 0 {
		return nil, ierrors.ValidationError("WLAN ID", "cannot be negative")
	}

	url := fmt.Sprintf("%s=%d", routes.CommonOperDataEndpoint, wlanID)
	return core.Get[model.ClientOperCommonOperData](ctx, s.Client(), url)
}

// GetOperDcInfo retrieves discovery client information
func (s Service) GetOperDcInfo(ctx context.Context) (*model.ClientOperDcInfo, error) {
	return core.Get[model.ClientOperDcInfo](ctx, s.Client(), routes.DcInfoEndpoint)
}

// GetOperDot11OperData returns 802.11 operational data for clients
func (s Service) GetOperDot11OperData(ctx context.Context) (*model.ClientOperDot11OperData, error) {
	return core.Get[model.ClientOperDot11OperData](ctx, s.Client(), routes.Dot11OperDataEndpoint)
}

// GetOperMmIfClientHistory retrieves mobility manager interface client history
func (s Service) GetOperMmIfClientHistory(ctx context.Context) (*model.ClientOperMmIfClientHistory, error) {
	return core.Get[model.ClientOperMmIfClientHistory](ctx, s.Client(), routes.MmIfClientHistoryEndpoint)
}

// GetOperMmIfClientStats retrieves mobility manager interface client statistics
func (s Service) GetOperMmIfClientStats(ctx context.Context) (*model.ClientOperMmIfClientStats, error) {
	return core.Get[model.ClientOperMmIfClientStats](ctx, s.Client(), routes.MmIfClientStatsEndpoint)
}

// GetOperMobilityOperData retrieves mobility operational data for clients
func (s Service) GetOperMobilityOperData(ctx context.Context) (*model.ClientOperMobilityOperData, error) {
	return core.Get[model.ClientOperMobilityOperData](ctx, s.Client(), routes.MobilityOperDataEndpoint)
}

// GetOperPolicyData retrieves client policy data
func (s Service) GetOperPolicyData(ctx context.Context) (*model.ClientOperPolicyData, error) {
	return core.Get[model.ClientOperPolicyData](ctx, s.Client(), routes.PolicyDataEndpoint)
}

// GetOperSisfDBMac retrieves SISF database MAC information
func (s Service) GetOperSisfDBMac(ctx context.Context) (*model.ClientOperSisfDBMac, error) {
	return core.Get[model.ClientOperSisfDBMac](ctx, s.Client(), routes.SisfDBMacEndpoint)
}

// GetOperTrafficStats retrieves client traffic statistics
func (s Service) GetOperTrafficStats(ctx context.Context) (*model.ClientOperTrafficStats, error) {
	return core.Get[model.ClientOperTrafficStats](ctx, s.Client(), routes.TrafficStatsEndpoint)
}
