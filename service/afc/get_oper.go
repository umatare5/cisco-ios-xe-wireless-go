// Package afc provides Automated Frequency Coordination functionality for the Cisco IOS-XE Wireless Network Controller API.
package afc

import (
	"context"
	"fmt"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	ierrors "github.com/umatare5/cisco-ios-xe-wireless-go/internal/errors"
	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/afc"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/validation"
)

// operOps provides high-level operational operations for AFC service
func (s Service) operOps() *core.OperationalOperations[model.AfcOper] {
	return core.NewOperationalOperations[model.AfcOper](s.Client(), routes.AFCOperBasePath)
}

// GetOper retrieves the complete AFC operational data
func (s Service) GetOper(ctx context.Context) (*model.AfcOper, error) {
	return s.operOps().GetAll(ctx)
}

// GetOperAPResp retrieves AFC AP response data
func (s Service) GetOperAPResp(ctx context.Context) (*model.AfcOperEwlcAfcApResp, error) {
	return core.Get[model.AfcOperEwlcAfcApResp](ctx, s.Client(), routes.ApRespEndpoint)
}

// GetOperCloudOper retrieves AFC cloud operational data
func (s Service) GetOperCloudOper(ctx context.Context) (*model.AfcCloudOper, error) {
	return core.Get[model.AfcCloudOper](ctx, s.Client(), routes.CloudOperEndpoint)
}

// GetOperCloudStats retrieves AFC cloud statistics
func (s Service) GetOperCloudStats(ctx context.Context) (*model.AfcCloudOperAfcCloudStats, error) {
	return core.Get[model.AfcCloudOperAfcCloudStats](ctx, s.Client(), routes.CloudStatsEndpoint)
}

// GetOperByApMac retrieves AFC operational data filtered by AP MAC address
func (s Service) GetOperByApMac(ctx context.Context, apMac string) (*model.AfcOperEwlcAfcApResp, error) {
	if err := s.ValidateClient(); err != nil {
		return nil, err
	}

	if err := validation.ValidateAPMac(apMac); err != nil {
		errMsg := "data for MAC " + apMac
		return nil, ierrors.ServiceOperationError(ierrors.ActionGet, EntityTypeAFC, errMsg, err)
	}

	url := s.Client().RestconfBuilder().BuildQueryURL(routes.ApRespEndpoint, apMac)
	return core.Get[model.AfcOperEwlcAfcApResp](ctx, s.Client(), url)
}

// GetOperByApMacAndSlot retrieves AFC operational data filtered by AP MAC address and slot
func (s Service) GetOperByApMacAndSlot(
	ctx context.Context, apMac string, slot int,
) (*model.AfcOperEwlcAfcApResp, error) {
	if err := validation.ValidateAPMac(apMac); err != nil {
		errMsg := fmt.Sprintf("data for MAC %s slot %d", apMac, slot)
		return nil, ierrors.ServiceOperationError(ierrors.ActionGet, EntityTypeAFC, errMsg, err)
	}
	if err := validation.ValidateSlotID(slot); err != nil {
		errMsg := fmt.Sprintf("data for slot %d", slot)
		return nil, ierrors.ServiceOperationError(ierrors.ActionGet, EntityTypeAFC, errMsg, err)
	}

	url := fmt.Sprintf("%s=%s,%d", routes.ApRespEndpoint, apMac, slot)
	return core.Get[model.AfcOperEwlcAfcApResp](ctx, s.Client(), url)
}

// GetOperByApMacAndRequestID retrieves AFC operational data filtered by AP MAC address and request ID
func (s Service) GetOperByApMacAndRequestID(
	ctx context.Context, apMac, requestID string,
) (*model.AfcOperEwlcAfcApResp, error) {
	if err := validation.ValidateAPMac(apMac); err != nil {
		errMsg := "data for MAC " + apMac
		return nil, ierrors.ServiceOperationError(ierrors.ActionGet, EntityTypeAFC, errMsg, err)
	}
	if err := validation.ValidateNonEmptyString(requestID, "request ID"); err != nil {
		return nil, ierrors.ServiceOperationError(ierrors.ActionValidate, EntityTypeAFC, "request ID", err)
	}

	url := fmt.Sprintf("%s=%s/resp-data/request-id=%s", routes.ApRespEndpoint, apMac, requestID)
	return core.Get[model.AfcOperEwlcAfcApResp](ctx, s.Client(), url)
}

// GetOperBySlot retrieves AFC operational data filtered by slot ID
func (s Service) GetOperBySlot(
	ctx context.Context, apMac string, slot int,
) (*model.AfcOperEwlcAfcApResp, error) {
	if err := validation.ValidateAPMac(apMac); err != nil {
		errMsg := "data for MAC " + apMac
		return nil, ierrors.ServiceOperationError(ierrors.ActionGet, EntityTypeAFC, errMsg, err)
	}
	if err := validation.ValidateSlotID(slot); err != nil {
		errMsg := fmt.Sprintf("data for slot %d", slot)
		return nil, ierrors.ServiceOperationError(ierrors.ActionGet, EntityTypeAFC, errMsg, err)
	}

	url := fmt.Sprintf("%s=%s,%d", routes.ApRespEndpoint, apMac, slot)
	return core.Get[model.AfcOperEwlcAfcApResp](ctx, s.Client(), url)
}

// GetOperByRequestID retrieves AFC operational data filtered by request ID
func (s Service) GetOperByRequestID(
	ctx context.Context, apMac, requestID string,
) (*model.AfcOperEwlcAfcApResp, error) {
	if err := validation.ValidateAPMac(apMac); err != nil {
		errMsg := "data for MAC " + apMac
		return nil, ierrors.ServiceOperationError(ierrors.ActionGet, EntityTypeAFC, errMsg, err)
	}
	if err := validation.ValidateNonEmptyString(requestID, "request ID"); err != nil {
		return nil, ierrors.ServiceOperationError(ierrors.ActionValidate, EntityTypeAFC, "request ID", err)
	}

	url := fmt.Sprintf("%s=%s/resp-data/request-id=%s", routes.ApRespEndpoint, apMac, requestID)
	return core.Get[model.AfcOperEwlcAfcApResp](ctx, s.Client(), url)
}
