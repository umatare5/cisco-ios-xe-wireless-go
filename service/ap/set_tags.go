package ap

import (
	"context"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/builders"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
	ierrors "github.com/umatare5/cisco-ios-xe-wireless-go/internal/errors"
	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/ap"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/restconf/routes"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/validation"
)

// Tag type enumeration for unified handling
type tagType int

const (
	siteTagType tagType = iota
	policyTagType
	rfTagType
)

// AssignSiteTag assigns a site tag to an Access Point using MAC address.
func (s Service) AssignSiteTag(ctx context.Context, apMac, siteTag string) error {
	return s.assignSingleTag(ctx, apMac, siteTag, siteTagType)
}

// AssignPolicyTag assigns a policy tag to an Access Point using MAC address.
func (s Service) AssignPolicyTag(ctx context.Context, apMac, policyTag string) error {
	return s.assignSingleTag(ctx, apMac, policyTag, policyTagType)
}

// AssignRFTag assigns an RF tag to an Access Point using MAC address.
func (s Service) AssignRFTag(ctx context.Context, apMac, rfTag string) error {
	return s.assignSingleTag(ctx, apMac, rfTag, rfTagType)
}

// assignSingleTag is a unified helper for single tag assignment
func (s Service) assignSingleTag(ctx context.Context, apMac, tagValue string, tType tagType) error {
	var tagTypeStr string
	var builderType builders.TagType

	switch tType {
	case siteTagType:
		tagTypeStr = "site"
		builderType = builders.SiteTagType
	case policyTagType:
		tagTypeStr = "policy"
		builderType = builders.PolicyTagType
	case rfTagType:
		tagTypeStr = "rf"
		builderType = builders.RFTagType
	default:
		return ierrors.GetTagValidationError("unknown")
	}

	if !validation.IsValidTagAssignment(tagValue, tagTypeStr) {
		return ierrors.GetTagValidationError(tagTypeStr)
	}

	return s.assignTags(ctx, apMac, builders.CreateSingleTagRequest(tagValue, builderType))
}

// assignTags assigns multiple tags to an Access Point (internal implementation).
func (s Service) assignTags(ctx context.Context, apMac string, tags model.ApCfgApTagsParams) error {
	if !validation.IsValidAPMacFormat(apMac) {
		return ierrors.ValidationError("AP MAC address", apMac)
	}
	if !validation.HasValidTags(tags.SiteTag, tags.PolicyTag, tags.RFTag) {
		return ierrors.RequiredParameterError("at least one tag")
	}

	normalizedMAC := validation.NormalizeAPMac(apMac)
	url := s.Client().RestconfBuilder().BuildQueryURL(routes.APCfgBasePath+"/ap-tags/ap-tag", normalizedMAC)
	tagData := builders.BuildAPCfgApTagData(normalizedMAC, tags)

	// Execute operation with direct error propagation
	if err := core.PutVoid(ctx, s.Client(), url, model.APTagPayload{ApTag: tagData}); validation.HasError(err) {
		return ierrors.ServiceOperationError("assign", "AP", "tags", err)
	}
	return nil
}
