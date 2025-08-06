// Package wnc provides a multi-layer client for Cisco Catalyst 9800 WNC.
//
// This library implements a three-layer architecture separating the core HTTP client,
// domain-specific services, and generated type definitions. All interactions follow
// the pattern client.<Domain>().<Method>().
//
// Quick example:
//
//	c, _ := wnc.New("host", "token")
//	radios, _ := c.RRM().Oper(ctx)
//	stats, _ := c.AFC().CloudStats(ctx)
//
// Supported domains include AFC, AP, Client, General, RRM, WLAN, Rogue, and more.
// Each domain provides typed methods for configuration and operational data retrieval.
package wnc
