package model

// SiteOperResponse represents the response structure for Site operational data.
type SiteOperResponse struct {
	SiteOperData SiteOperData `json:"Cisco-IOS-XE-wireless-site-oper:site-oper-data"`
}

// SiteOperData contains Site operational data
type SiteOperData struct {
	Sites []SiteInfo `json:"sites"`
}

// SiteInfo represents site information
type SiteInfo struct {
	SiteID          string          `json:"site-id"`
	SiteName        string          `json:"site-name"`
	SiteType        string          `json:"site-type"`
	ApCount         int             `json:"ap-count"`
	ClientCount     int             `json:"client-count"`
	SiteStatus      string          `json:"site-status"`
	LocationDetails LocationDetails `json:"location-details"`
}

// LocationDetails represents site location details
type LocationDetails struct {
	Building  string    `json:"building"`
	Floor     string    `json:"floor"`
	Address   string    `json:"address"`
	GpsCoords GpsCoords `json:"gps-coords"`
}

// GpsCoords represents GPS coordinates
type GpsCoords struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
