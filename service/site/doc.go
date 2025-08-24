// Package site provides site-specific configuration and operational operations for Cisco IOS-XE wireless controllers.
//
// This package implements comprehensive methods to interact with site-specific configuration and operational data
// through the RESTCONF API on Cisco IOS-XE controllers. It provides capabilities for site-level management,
// location-based policies, AP configuration profiles, and geographic deployments.
//
// The package leverages the YANG models: Cisco-IOS-XE-wireless-site-cfg and Cisco-IOS-XE-wireless-site-oper for complete site management.
//
// # Main Functions
//
// The Site service provides three main categories of operations:
//
// 1. GetCfg - Configuration Data
// 2. GetOper - Operational Data
// 3. Site Tag Management - CRUD operations for site tags
//
// # Site Tag Configuration Management
//
// The service provides comprehensive site tag configuration operations:
//
// **Site Tag CRUD Operations:**
//   - CreateSiteTag(ctx, config): Creates a new site tag
//   - GetSiteTag(ctx, siteTagName): Retrieves a specific site tag
//   - GetAllSiteTags(ctx): Lists all site tags
//   - SetSiteTag(ctx, config): Sets/updates an existing site tag
//   - DeleteSiteTag(ctx, siteTagName): Removes a site tag
//
// **Individual Attribute Setters:**
//   - SetAPJoinProfile(ctx, siteTagName, apProfile): Sets AP join profile
//   - SetFlexProfile(ctx, siteTagName, flexProfile): Sets flex profile
//   - SetLocalSite(ctx, siteTagName, enabled): Configures local site mode
//   - SetDescription(ctx, siteTagName, description): Sets description
//
// **Complete Configuration:**
//   - ConfigureSiteTag(ctx, ...): Creates or updates with all attributes
//
// Site tags define site-specific configurations that are applied to
// Access Points. They include AP profiles, flex profiles, and local
// site settings that control AP behavior and capabilities.
//
// # Filter Functions
//
// This package provides 2 filter functions for precise data retrieval:
//
// Configuration Filters:
//   - GetCfgByApProfile - Filter by AP profile name
//   - GetCfgBySiteTag - Filter by site tag name
package site
