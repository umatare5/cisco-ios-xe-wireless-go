// Package builders provides data structure builders for various Cisco Wireless Network Controller operations.
package builders

import (
	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/ap"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/validation"
)

// TagType represents the type of tag being assigned
type TagType int

const (
	// SiteTagType represents a site tag
	SiteTagType TagType = iota
	// PolicyTagType represents a policy tag
	PolicyTagType
	// RFTagType represents an RF tag
	RFTagType
)

// CreateSingleTagRequest creates tag request with single tag value
func CreateSingleTagRequest(tagValue string, tType TagType) model.ApCfgApTagsParams {
	switch tType {
	case SiteTagType:
		return model.ApCfgApTagsParams{SiteTag: tagValue}
	case PolicyTagType:
		return model.ApCfgApTagsParams{PolicyTag: tagValue}
	case RFTagType:
		return model.ApCfgApTagsParams{RFTag: tagValue}
	default:
		return model.ApCfgApTagsParams{}
	}
}

// BuildAPCfgApTagData constructs the payload for tag assignment requests
func BuildAPCfgApTagData(normalizedMAC string, tags model.ApCfgApTagsParams) model.APCfgApTagData {
	return model.APCfgApTagData{
		APMac:     normalizedMAC,
		SiteTag:   validation.SelectNonEmptyValue(tags.SiteTag, validation.DefaultSiteTag),
		PolicyTag: validation.SelectNonEmptyValue(tags.PolicyTag, validation.DefaultPolicyTag),
		RFTag:     validation.SelectNonEmptyValue(tags.RFTag, validation.DefaultRFTag),
	}
}
