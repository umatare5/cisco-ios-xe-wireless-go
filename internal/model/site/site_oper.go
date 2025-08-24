// Package model provides data models for site operational data.
package model

// SiteOper  represents the Site operational data.
type SiteOper struct {
	SiteOperData SiteOperData `json:"Cisco-IOS-XE-wireless-site-oper:site-oper-data"`
}

// SiteOperSites  represents the Sites data.
type SiteOperSites struct {
	Sites []SiteInfo `json:"Cisco-IOS-XE-wireless-site-oper:sites"`
}

type SiteOperData struct {
	Sites []SiteInfo `json:"sites"`
}

type SiteInfo struct {
	SiteID          string          `json:"site-id"`
	SiteName        string          `json:"site-name"`
	SiteType        string          `json:"site-type"`
	ApCount         int             `json:"ap-count"`
	ClientCount     int             `json:"client-count"`
	SiteStatus      string          `json:"site-status"`
	LocationDetails LocationDetails `json:"location-details"`
}

type LocationDetails struct {
	Building  string    `json:"building"`
	Floor     string    `json:"floor"`
	Address   string    `json:"address"`
	GpsCoords GpsCoords `json:"gps-coords"`
}

type GpsCoords struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
