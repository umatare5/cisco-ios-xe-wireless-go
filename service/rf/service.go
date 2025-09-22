package rf

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/service"
)

// Service provides RF (Radio Frequency) management operations for Cisco IOS-XE Wireless LAN Controller.
type Service struct {
	service.BaseService
}

// NewService creates a new RF service instance with the provided client.
func NewService(client *core.Client) Service {
	return Service{BaseService: service.NewBaseService(client)}
}

// RFTag returns an RF tag service instance for RF tag management operations.
func (s Service) RFTag() *RFTagService {
	return NewRFTagService(s.Client())
}

// GetConfig retrieves RF configuration data including RF profiles and power settings.
func (s Service) GetConfig(ctx context.Context) (*RFCfg, error) {
	return core.Get[RFCfg](ctx, s.Client(), routes.RFCfgPath)
}

// ListRFTags retrieves RF tags collection from the configuration.
func (s Service) ListRFTags(ctx context.Context) (*RFCfgRFTags, error) {
	return core.Get[RFCfgRFTags](ctx, s.Client(), routes.RFTagsPath)
}

// ListRFProfiles retrieves RF profiles collection from the configuration.
func (s Service) ListRFProfiles(ctx context.Context) (*RFProfiles, error) {
	return core.Get[RFProfiles](ctx, s.Client(), routes.RFProfilesPath)
}

// ListMultiBssidProfiles retrieves Multi-BSSID profiles from the configuration.
func (s Service) ListMultiBssidProfiles(ctx context.Context) (*MultiBssidProfiles, error) {
	return core.Get[MultiBssidProfiles](ctx, s.Client(), routes.MultiBssidProfilesPath)
}

// ListAtfPolicies retrieves Air Time Fairness policies from the configuration.
func (s Service) ListAtfPolicies(ctx context.Context) (*AtfPolicies, error) {
	return core.Get[AtfPolicies](ctx, s.Client(), routes.AtfPoliciesPath)
}

// ListRFProfileDefaultEntries retrieves RF profile default entries from the configuration.
func (s Service) ListRFProfileDefaultEntries(ctx context.Context) (*RFProfileDefaultEntries, error) {
	return core.Get[RFProfileDefaultEntries](ctx, s.Client(), routes.RFProfileDefaultEntriesPath)
}
