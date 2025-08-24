package wlan

import (
	"context"
	"fmt"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	ierrors "github.com/umatare5/cisco-ios-xe-wireless-go/internal/errors"
	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/wlan"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
)

// configOps provides high-level configuration operations for WLAN service
func (s Service) configOps() *core.ConfigOperations[model.WlanCfg] {
	return core.NewConfigOperations[model.WlanCfg](s.Client(), routes.WlanCfgBasePath)
}

// GetCfg retrieves complete WLAN configuration data from the controller.
//
// This method provides access to all WLAN configuration settings including
// WLAN entries, policies, AAA configurations, and policy list entries.
//
// Parameters:
//   - ctx: Context for the request, allowing for timeout and cancellation control
//
// Returns:
//   - *model.WlanCfg: Complete WLAN configuration data structure
//   - error: Any error encountered during the operation
//
// Example:
//
//	wlanService := wlan.NewService(client)
//	cfg, err := wlanService.GetCfg(ctx)
//	if err != nil {
//		return fmt.Errorf("failed to get WLAN config: %w", err)
//	}
func (s Service) GetCfg(ctx context.Context) (*model.WlanCfg, error) {
	return s.configOps().GetAll(ctx)
}

// GetCfgEntries retrieves all WLAN configuration entries.
//
// This method provides access to the complete list of configured WLANs
// with their profile names, IDs, and basic configuration parameters.
//
// Parameters:
//   - ctx: Context for the request, allowing for timeout and cancellation control
//
// Returns:
//   - *model.WlanCfgEntries: All WLAN configuration entries
//   - error: Any error encountered during the operation
//
// Example:
//
//	entries, err := wlanService.GetCfgEntries(ctx)
//	if err != nil {
//		return err
//	}
//
//	for _, entry := range entries.WlanCfgEntries.WlanCfgEntry {
//	}
func (s Service) GetCfgEntries(ctx context.Context) (*model.WlanCfgEntries, error) {
	return core.Get[model.WlanCfgEntries](ctx, s.Client(), routes.WlanCfgEntriesEndpoint)
}

// GetCfgByProfileName retrieves WLAN configuration data filtered by profile name.
//
// This method allows retrieval of configuration data for a specific WLAN
// identified by its profile name.
//
// Parameters:
//   - ctx: Context for the request, allowing for timeout and cancellation control
//   - profileName: Name of the WLAN profile to retrieve
//
// Returns:
//   - *model.WlanCfgEntries: WLAN configuration for the specified profile
//   - error: Any error encountered during the operation
//
// Example:
//
//	profile, err := wlanService.GetCfgByProfileName(ctx, "corporate-wlan")
//	if err != nil {
//		return err
//	}
func (s Service) GetCfgByProfileName(
	ctx context.Context,
	profileName string,
) (*model.WlanCfgEntries, error) {
	url := fmt.Sprintf("%s=%s", routes.WlanCfgEntryEndpoint, profileName)
	return core.Get[model.WlanCfgEntries](ctx, s.Client(), url)
}

// GetCfgPolicies retrieves all WLAN policies.
//
// This method provides access to the complete list of WLAN policies
// configured on the controller.
//
// Parameters:
//   - ctx: Context for the request, allowing for timeout and cancellation control
//
// Returns:
//   - *model.WlanPolicies: All WLAN policies
//   - error: Any error encountered during the operation
//
// Example:
//
//	policies, err := wlanService.GetCfgPolicies(ctx)
//	if err != nil {
//		return err
//	}
func (s Service) GetCfgPolicies(ctx context.Context) (*model.WlanPolicies, error) {
	return core.Get[model.WlanPolicies](ctx, s.Client(), routes.WlanPoliciesEndpoint)
}

// GetCfgPoliciesByPolicyProfileName retrieves WLAN policies filtered by policy profile name.
//
// This method allows retrieval of policy data for a specific WLAN policy
// identified by its profile name.
//
// Parameters:
//   - ctx: Context for the request, allowing for timeout and cancellation control
//   - policyProfileName: Name of the policy profile to retrieve
//
// Returns:
//   - *model.WlanPolicies: WLAN policy for the specified profile
//   - error: Any error encountered during the operation
func (s Service) GetCfgPoliciesByPolicyProfileName(
	ctx context.Context,
	policyProfileName string,
) (*model.WlanPolicies, error) {
	url := fmt.Sprintf("%s=%s", routes.WlanPolicyEntryEndpoint, policyProfileName)
	return core.Get[model.WlanPolicies](ctx, s.Client(), url)
}

// GetCfgPolicyListEntries retrieves all policy list entries.
//
// This method provides access to the complete list of policy list entries
// configured on the controller.
//
// Parameters:
//   - ctx: Context for the request, allowing for timeout and cancellation control
//
// Returns:
//   - *model.PolicyListEntries: All policy list entries
//   - error: Any error encountered during the operation
func (s Service) GetCfgPolicyListEntries(ctx context.Context) (*model.PolicyListEntries, error) {
	return core.Get[model.PolicyListEntries](ctx, s.Client(), routes.PolicyListEntriesEndpoint)
}

// GetCfgWirelessAaaPolicyConfigs retrieves wireless AAA policy configurations.
//
// This method provides access to all wireless AAA policy configurations
// including authentication and authorization settings.
//
// Parameters:
//   - ctx: Context for the request, allowing for timeout and cancellation control
//
// Returns:
//   - *model.WirelessAaaPolicyConfigs: All wireless AAA policy configurations
//   - error: Any error encountered during the operation
func (s Service) GetCfgWirelessAaaPolicyConfigs(ctx context.Context) (*model.WirelessAaaPolicyConfigs, error) {
	return core.Get[model.WirelessAaaPolicyConfigs](ctx, s.Client(), routes.WirelessAaaPolicyConfigsEndpoint)
}

// GetCfgByID retrieves WLAN configuration data filtered by WLAN ID.
//
// This method allows retrieval of configuration data for a specific WLAN
// identified by its numeric ID.
//
// Parameters:
//   - ctx: Context for the request, allowing for timeout and cancellation control
//   - wlanID: Numeric ID of the WLAN to retrieve
//
// Returns:
//   - *model.WlanCfg: WLAN configuration for the specified ID
//   - error: Any error encountered during the operation
//
// Example:
//
//	cfg, err := wlanService.GetCfgByID(ctx, 1)
//	if err != nil {
//		return err
//	}
//
//	if len(cfg.WlanCfgData.WlanCfgEntries.WlanCfgEntry) > 0 {
//		entry := cfg.WlanCfgData.WlanCfgEntries.WlanCfgEntry[0]
//	}
func (s Service) GetCfgByID(ctx context.Context, wlanID int) (*model.WlanCfg, error) {
	// Get all WLAN entries first
	entries, err := s.GetCfgEntries(ctx)
	if err != nil {
		return nil, ierrors.ServiceOperationError("get", "WLAN", "configuration entries", err)
	}

	// Filter by WLAN ID
	var filteredEntries []model.WlanCfgEntry
	for _, entry := range entries.WlanCfgEntry {
		if entry.WlanID == wlanID {
			filteredEntries = append(filteredEntries, entry)
		}
	}

	return &model.WlanCfg{
		WlanCfgData: &model.WlanCfgData{
			WlanCfgEntries: &model.WlanCfgEntries{
				WlanCfgEntry: filteredEntries,
			},
		},
	}, nil
}
