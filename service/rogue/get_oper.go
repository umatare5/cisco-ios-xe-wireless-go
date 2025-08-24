package rogue

import (
	"context"
	"fmt"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/rogue"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/validation"
)

// operOps provides high-level operational operations for rogue service
func (s Service) operOps() *core.OperationalOperations[model.RogueOper] {
	return core.NewOperationalOperations[model.RogueOper](s.Client(), routes.RogueOperBasePath)
}

// GetOper retrieves rogue operational data.
func (s Service) GetOper(ctx context.Context) (*model.RogueOper, error) {
	return s.operOps().GetAll(ctx)
}

// GetOperByRogueAddress retrieves rogue data filtered by rogue address.
func (s Service) GetOperByRogueAddress(ctx context.Context, rogueAddress string) (*model.RogueData, error) {
	if err := s.ValidateClient(); err != nil {
		return nil, err
	}

	if err := validation.ValidateRogueAddress(rogueAddress); err != nil {
		return nil, err
	}

	url := s.Client().RestconfBuilder().BuildPathQueryURL(routes.RogueOperBasePath, "rogue-data", rogueAddress)
	return core.Get[model.RogueData](ctx, s.Client(), url)
}

// GetOperByRogueClientAddress retrieves rogue client data filtered by rogue client address.
func (s Service) GetOperByRogueClientAddress(
	ctx context.Context,
	rogueClientAddress string,
) (*model.RogueClientData, error) {
	if err := s.ValidateClient(); err != nil {
		return nil, err
	}

	if err := validation.ValidateRogueClientAddress(rogueClientAddress); err != nil {
		return nil, err
	}

	url := s.Client().RestconfBuilder().BuildPathQueryURL(
		routes.RogueOperBasePath, "rogue-client-data", rogueClientAddress)
	return core.Get[model.RogueClientData](ctx, s.Client(), url)
}

// GetOperClientData retrieves rogue client data.
func (s Service) GetOperClientData(ctx context.Context) (*model.RogueClientData, error) {
	return core.Get[model.RogueClientData](ctx, s.Client(), routes.RogueClientDataEndpoint)
}

// GetOperData retrieves rogue data.
func (s Service) GetOperData(ctx context.Context) (*model.RogueData, error) {
	return core.Get[model.RogueData](ctx, s.Client(), routes.RogueDataEndpoint)
}

// GetOperRldpStats retrieves RLDP statistics.
func (s Service) GetOperRldpStats(ctx context.Context) (*model.RldpStats, error) {
	return core.Get[model.RldpStats](ctx, s.Client(), routes.RldpStatsEndpoint)
}

// GetOperStats retrieves rogue statistics.
func (s Service) GetOperStats(ctx context.Context) (*model.RogueStats, error) {
	return core.Get[model.RogueStats](ctx, s.Client(), routes.RogueStatsEndpoint)
}

// GetOperByClassType retrieves rogue data filtered by class type.
func (s Service) GetOperByClassType(ctx context.Context, classType string) (*model.RogueData, error) {
	if err := validation.ValidateClassType(classType); err != nil {
		return nil, err
	}

	// Note: This is a conceptual implementation - actual API endpoint may vary
	url := fmt.Sprintf("%s?class-type=%s", routes.RogueDataEndpoint, classType)
	return core.Get[model.RogueData](ctx, s.Client(), url)
}

// GetOperByContainmentLevel retrieves rogue data filtered by containment level.
func (s Service) GetOperByContainmentLevel(ctx context.Context, level int) (*model.RogueData, error) {
	if err := validation.ValidateContainmentLevel(level); err != nil {
		return nil, err
	}

	// Note: This is a conceptual implementation - actual API endpoint may vary
	url := fmt.Sprintf("%s?containment-level=%d", routes.RogueDataEndpoint, level)
	return core.Get[model.RogueData](ctx, s.Client(), url)
}

// GetOperWithFields retrieves rogue operational data with specific fields.
func (s Service) GetOperWithFields(ctx context.Context, fields []string) (*model.RogueOper, error) {
	if err := s.ValidateClient(); err != nil {
		return nil, err
	}

	url := s.Client().RestconfBuilder().BuildFieldsURLMultiple(routes.RogueOperEndpoint, fields)
	return core.Get[model.RogueOper](ctx, s.Client(), url)
}

// GetOperStatsWithFields retrieves rogue statistics with specific fields.
func (s Service) GetOperStatsWithFields(ctx context.Context, fields []string) (*model.RogueStats, error) {
	if err := s.ValidateClient(); err != nil {
		return nil, err
	}

	url := s.Client().RestconfBuilder().BuildFieldsURLMultiple(routes.RogueStatsEndpoint, fields)
	return core.Get[model.RogueStats](ctx, s.Client(), url)
}

// GetOperDataWithFields retrieves rogue data with specific fields.
func (s Service) GetOperDataWithFields(ctx context.Context, fields []string) (*model.RogueData, error) {
	if err := s.ValidateClient(); err != nil {
		return nil, err
	}

	url := s.Client().RestconfBuilder().BuildFieldsURLMultiple(routes.RogueDataEndpoint, fields)
	return core.Get[model.RogueData](ctx, s.Client(), url)
}

// GetOperClientDataWithFields retrieves rogue client data with specific fields.
func (s Service) GetOperClientDataWithFields(
	ctx context.Context,
	fields []string,
) (*model.RogueClientData, error) {
	if err := s.ValidateClient(); err != nil {
		return nil, err
	}

	url := s.Client().RestconfBuilder().BuildFieldsURLMultiple(routes.RogueClientDataEndpoint, fields)
	return core.Get[model.RogueClientData](ctx, s.Client(), url)
}

// GetOperRldpStatsWithFields retrieves RLDP statistics with specific fields.
func (s Service) GetOperRldpStatsWithFields(ctx context.Context, fields []string) (*model.RldpStats, error) {
	if err := s.ValidateClient(); err != nil {
		return nil, err
	}

	url := s.Client().RestconfBuilder().BuildFieldsURLMultiple(routes.RldpStatsEndpoint, fields)
	return core.Get[model.RldpStats](ctx, s.Client(), url)
}
