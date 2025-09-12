package urwb

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/errors"
	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/urwb"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/service"
)

// Service provides access to URWB (Ultra Reliable Wireless Backhaul) operations.
type Service struct {
	service.BaseService
}

// NewService creates a new URWB service instance with the provided client.
func NewService(client *core.Client) Service {
	return Service{BaseService: service.NewBaseService(client)}
}

// GetConfig retrieves the complete URWB configuration from the controller
// EXPERIMENTAL: Requires IOS-XE 17.18.1+.
func (s Service) GetConfig(ctx context.Context) (*model.UrwbProfiles, error) {
	result, err := core.Get[model.UrwbCfgData](ctx, s.Client(), routes.URWBCfgPath)
	if err != nil {
		return nil, err
	}
	return result.UrwbProfiles, nil
}

// SetConfig configures the complete URWB configuration on the controller
// EXPERIMENTAL: Requires IOS-XE 17.18.1+.
func (s Service) SetConfig(ctx context.Context, profiles *model.UrwbProfiles) error {
	config := &model.UrwbCfgData{UrwbProfiles: profiles}
	return core.PutVoid(ctx, s.Client(), routes.URWBCfgPath, config)
}

// ListProfiles retrieves all URWB profiles
// EXPERIMENTAL: Requires IOS-XE 17.18.1+.
func (s Service) ListProfiles(ctx context.Context) ([]model.UrwbProfile, error) {
	result, err := core.Get[model.UrwbCfgData](ctx, s.Client(), routes.URWBCfgPath)
	if err != nil {
		return nil, err
	}
	if result.UrwbProfiles == nil {
		return nil, nil
	}
	return result.UrwbProfiles.UrwbProfile, nil
}

// GetProfile retrieves a specific URWB profile by name
// EXPERIMENTAL: Requires IOS-XE 17.18.1+.
func (s Service) GetProfile(ctx context.Context, profileName string) (*model.UrwbProfile, error) {
	if profileName == "" {
		return nil, errors.RequiredParameterError("profileName")
	}

	// Get all profiles and search for the specific one (no XPath filtering)
	profiles, err := s.ListProfiles(ctx)
	if err != nil {
		return nil, err
	}

	for _, profile := range profiles {
		if profile.ProfileName == profileName {
			return &profile, nil
		}
	}
	return nil, errors.NotFoundError("URWB profile", profileName)
}

// UpsertProfile configures a URWB profile
// EXPERIMENTAL: Requires IOS-XE 17.18.1+.
func (s Service) UpsertProfile(ctx context.Context, profile *model.UrwbProfile) error {
	if profile.ProfileName == "" {
		return errors.RequiredParameterError("ProfileName")
	}

	// Get current configuration
	currentProfiles, err := s.ListProfiles(ctx)
	if err != nil {
		return err
	}

	// Update or add the profile
	updated := false
	for i, currentProfile := range currentProfiles {
		if currentProfile.ProfileName == profile.ProfileName {
			currentProfiles[i] = *profile
			updated = true
			break
		}
	}

	if !updated {
		// Add new profile
		currentProfiles = append(currentProfiles, *profile)
	}

	// Set the complete configuration
	profilesContainer := &model.UrwbProfiles{
		UrwbProfile: currentProfiles,
	}
	return s.SetConfig(ctx, profilesContainer)
}

// DeleteProfile removes a URWB profile
// EXPERIMENTAL: Requires IOS-XE 17.18.1+.
func (s Service) DeleteProfile(ctx context.Context, profileName string) error {
	if profileName == "" {
		return errors.RequiredParameterError("profileName")
	}

	// Get current configuration
	currentProfiles, err := s.ListProfiles(ctx)
	if err != nil {
		return err
	}

	// Find and remove the profile
	newProfiles := make([]model.UrwbProfile, 0, len(currentProfiles))
	found := false
	for _, currentProfile := range currentProfiles {
		if currentProfile.ProfileName != profileName {
			newProfiles = append(newProfiles, currentProfile)
		} else {
			found = true
		}
	}

	if !found {
		return errors.NotFoundError("URWB profile", profileName)
	}

	// Set the updated configuration
	profilesContainer := &model.UrwbProfiles{
		UrwbProfile: newProfiles,
	}
	return s.SetConfig(ctx, profilesContainer)
}

// GetOperational retrieves the complete URWB operational data from the controller
// EXPERIMENTAL: Requires IOS-XE 17.18.1+.
func (s Service) GetOperational(ctx context.Context) (*model.UrwbnetOperData, error) {
	result, err := core.Get[model.UrwbnetOper](ctx, s.Client(), routes.URWBOperPath)
	if err != nil {
		return nil, err
	}
	return result.UrwbnetOperData, nil
}

// ListStats retrieves URWB network statistics for all coordinators
// EXPERIMENTAL: Requires IOS-XE 17.18.1+.
func (s Service) ListStats(ctx context.Context) ([]model.UrwbnetStats, error) {
	result, err := core.Get[model.UrwbnetOper](ctx, s.Client(), routes.URWBOperPath)
	if err != nil {
		return nil, err
	}
	if result.UrwbnetOperData == nil {
		return nil, nil
	}
	return result.UrwbnetOperData.UrwbnetStats, nil
}

// ListNodeGroups retrieves URWB node group information
// EXPERIMENTAL: Requires IOS-XE 17.18.1+.
func (s Service) ListNodeGroups(ctx context.Context) ([]model.UrwbnetNodeG, error) {
	result, err := core.Get[model.UrwbnetOper](ctx, s.Client(), routes.URWBOperPath)
	if err != nil {
		return nil, err
	}
	if result.UrwbnetOperData == nil {
		return nil, nil
	}
	return result.UrwbnetOperData.UrwbnetNodeG, nil
}
