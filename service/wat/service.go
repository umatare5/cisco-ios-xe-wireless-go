package wat

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/errors"
	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/wat"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/service"
)

// Service provides Wireless Application Templates (WAT) operations for Cisco IOS-XE Wireless LAN Controller.
type Service struct {
	service.BaseService
}

// NewService creates a new WAT service instance with the provided client.
func NewService(client *core.Client) Service {
	return Service{BaseService: service.NewBaseService(client)}
}

// GetConfig retrieves the complete WAT configuration from the controller.
// EXPERIMENTAL: Requires IOS-XE 17.18.1+.
func (s Service) GetConfig(ctx context.Context) (*model.WatConfig, error) {
	cfgData, err := core.Get[model.WatCfg](ctx, s.Client(), routes.WATCfgPath)
	if err != nil {
		return nil, err
	}
	if cfgData == nil || cfgData.WatCfgData == nil || cfgData.WatCfgData.WatConfig == nil {
		return &model.WatConfig{}, nil
	}
	return cfgData.WatCfgData.WatConfig, nil
}

// SetWatConfig configures the complete WAT configuration on the controller.
// EXPERIMENTAL: Requires IOS-XE 17.18.1+.
func (s Service) SetWatConfig(ctx context.Context, config *model.WatConfig) error {
	return core.PutVoid(ctx, s.Client(), routes.WATCfgPath, config)
}

// GetThousandeyesConfig retrieves the ThousandEyes integration configuration.
// EXPERIMENTAL: Requires IOS-XE 17.18.1+.
func (s Service) GetThousandeyesConfig(ctx context.Context) (*model.WatConfig, error) {
	return core.Get[model.WatConfig](ctx, s.Client(), routes.WATThousandeyesPath)
}

// SetThousandeyesConfig configures the ThousandEyes integration parameters.
// EXPERIMENTAL: Requires IOS-XE 17.18.1+.
func (s Service) SetThousandeyesConfig(ctx context.Context, config *model.WatConfig) error {
	return core.PutVoid(ctx, s.Client(), routes.WATThousandeyesPath, config)
}

// GetTestProfile retrieves a specific WAT test profile by name.
// EXPERIMENTAL: Requires IOS-XE 17.18.1+.
func (s Service) GetTestProfile(ctx context.Context, profileName string) (*model.WatTestProfile, error) {
	path := s.Client().RestconfBuilder().BuildQueryURL(routes.WATTestProfilePath, profileName)
	return core.Get[model.WatTestProfile](ctx, s.Client(), path)
}

// CreateTestProfile creates a new WAT test profile.
// EXPERIMENTAL: Requires IOS-XE 17.18.1+.
func (s Service) CreateTestProfile(ctx context.Context, profile *model.WatTestProfile) error {
	if profile.ProfileName == "" {
		return errors.RequiredParameterError("ProfileName")
	}
	path := s.Client().RestconfBuilder().BuildQueryURL(routes.WATTestProfilePath, profile.ProfileName)
	return core.PutVoid(ctx, s.Client(), path, profile)
}

// DeleteTestProfile removes a WAT test profile by name.
// EXPERIMENTAL: Requires IOS-XE 17.18.1+.
func (s Service) DeleteTestProfile(ctx context.Context, profileName string) error {
	path := s.Client().RestconfBuilder().BuildQueryURL(routes.WATTestProfilePath, profileName)
	return core.Delete(ctx, s.Client(), path)
}

// GetSchedule retrieves a specific WAT schedule by name.
// EXPERIMENTAL: Requires IOS-XE 17.18.1+.
func (s Service) GetSchedule(ctx context.Context, scheduleName string) (*model.WatSchedule, error) {
	path := s.Client().RestconfBuilder().BuildPathQueryURL(routes.WATCfgPath, "schedule", scheduleName)
	return core.Get[model.WatSchedule](ctx, s.Client(), path)
}

// CreateSchedule creates a new WAT schedule.
// EXPERIMENTAL: Requires IOS-XE 17.18.1+.
func (s Service) CreateSchedule(ctx context.Context, schedule *model.WatSchedule) error {
	if schedule.ScheduleName == "" {
		return errors.RequiredParameterError("ScheduleName")
	}
	path := s.Client().RestconfBuilder().BuildPathQueryURL(routes.WATCfgPath, "schedule", schedule.ScheduleName)
	return core.PutVoid(ctx, s.Client(), path, schedule)
}

// DeleteSchedule removes a WAT schedule by name.
// EXPERIMENTAL: Requires IOS-XE 17.18.1+.
func (s Service) DeleteSchedule(ctx context.Context, scheduleName string) error {
	path := s.Client().RestconfBuilder().BuildPathQueryURL(routes.WATCfgPath, "schedule", scheduleName)
	return core.Delete(ctx, s.Client(), path)
}

// GetReportTemplate retrieves a specific WAT report template by name.
// EXPERIMENTAL: Requires IOS-XE 17.18.1+.
func (s Service) GetReportTemplate(ctx context.Context, reportName string) (*model.WatReport, error) {
	path := s.Client().RestconfBuilder().BuildPathQueryURL(routes.WATCfgPath, "report", reportName)
	return core.Get[model.WatReport](ctx, s.Client(), path)
}

// CreateReportTemplate creates a new WAT report template.
// EXPERIMENTAL: Requires IOS-XE 17.18.1+.
func (s Service) CreateReportTemplate(ctx context.Context, report *model.WatReport) error {
	if report.ReportName == "" {
		return errors.RequiredParameterError("ReportName")
	}
	path := s.Client().RestconfBuilder().BuildPathQueryURL(routes.WATCfgPath, "report", report.ReportName)
	return core.PutVoid(ctx, s.Client(), path, report)
}

// DeleteReportTemplate removes a WAT report template by name.
// EXPERIMENTAL: Requires IOS-XE 17.18.1+.
func (s Service) DeleteReportTemplate(ctx context.Context, reportName string) error {
	path := s.Client().RestconfBuilder().BuildPathQueryURL(routes.WATCfgPath, "report", reportName)
	return core.Delete(ctx, s.Client(), path)
}
