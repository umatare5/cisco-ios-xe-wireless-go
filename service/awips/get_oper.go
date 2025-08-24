// Package awips provides Automated Wireless IPS operational data operations for the Cisco IOS-XE Wireless Network Controller API.
package awips

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	ierrors "github.com/umatare5/cisco-ios-xe-wireless-go/internal/errors"
	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/awips"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/validation"
)

// operOps provides high-level operational operations for awips service
func (s Service) operOps() *core.OperationalOperations[model.AwipsOper] {
	return core.NewOperationalOperations[model.AwipsOper](s.Client(), routes.AWIPSOperBasePath)
}

// GetOper retrieves the complete AWIPS operational data
func (s Service) GetOper(ctx context.Context) (*model.AwipsOper, error) {
	return s.operOps().GetAll(ctx)
}

// GetOperByApMac retrieves AWIPS operational data filtered by AP MAC address for both awips-per-ap-info and awips-ap-dwld-status arrays
func (s Service) GetOperByApMac(ctx context.Context, apMac string) (*model.AwipsOper, error) {
	if err := validation.ValidateAPMac(apMac); err != nil {
		return nil, ierrors.ServiceOperationError(ierrors.ActionGet, EntityTypeAWIPS, "data by AP MAC", err)
	}

	if err := s.ValidateClient(); err != nil {
		return nil, err
	}
	url := s.Client().RestconfBuilder().BuildPathQueryURL(routes.AWIPSOperBasePath, "awips-per-ap-info", apMac)
	return core.Get[model.AwipsOper](ctx, s.Client(), url)
}

// GetOperByApMacDownloadStatus retrieves AWIPS AP download status data filtered by AP MAC address for the awips-ap-dwld-status array
func (s Service) GetOperByApMacDownloadStatus(ctx context.Context, apMac string) (*model.AwipsOper, error) {
	if err := validation.ValidateAPMac(apMac); err != nil {
		return nil, ierrors.ServiceOperationError(ierrors.ActionGet, EntityTypeAWIPS, "download status by AP MAC", err)
	}

	if err := s.ValidateClient(); err != nil {
		return nil, err
	}
	url := s.Client().RestconfBuilder().BuildPathQueryURL(routes.AWIPSOperBasePath, "awips-ap-dwld-status", apMac)
	return core.Get[model.AwipsOper](ctx, s.Client(), url)
}
