// Package awips provides AWIPS (Advanced Weather Interactive Processing System) operational data functionality for the Cisco Wireless Network Controller API.
package awips

import (
	"context"

	wnc "github.com/umatare5/cisco-xe-wireless-restconf-go"
)

const (
	// AwipsOperBasePath defines the base path for AWIPS operational data endpoints
	AwipsOperBasePath = "/restconf/data/Cisco-IOS-XE-wireless-awips-oper:awips-oper-data"
	// AwipsOperEndpoint retrieves complete AWIPS operational data
	AwipsOperEndpoint = AwipsOperBasePath
	// AwipsPerApInfoEndpoint retrieves AWIPS per-AP information
	AwipsPerApInfoEndpoint = AwipsOperBasePath + "/awips-per-ap-info"
	// AwipsDwldStatusEndpoint retrieves AWIPS download status
	AwipsDwldStatusEndpoint = AwipsOperBasePath + "/awips-dwld-status"
	// AwipsApDwldStatusEndpoint retrieves AWIPS AP download status
	AwipsApDwldStatusEndpoint = AwipsOperBasePath + "/awips-ap-dwld-status"
)

// AwipsOperResponse represents the complete AWIPS operational data response
type AwipsOperResponse struct {
	CiscoIOSXEWirelessAwipsOperAwipsOperData struct {
		AwipsPerApInfo    []AwipsPerApInfo    `json:"awips-per-ap-info"`
		AwipsDwldStatus   AwipsDwldStatus     `json:"awips-dwld-status"`
		AwipsApDwldStatus []AwipsApDwldStatus `json:"awips-ap-dwld-status"`
	} `json:"Cisco-IOS-XE-wireless-awips-oper:awips-oper-data"`
}

// AwipsOperPerApInfoResponse represents the AWIPS per-AP information response
type AwipsOperPerApInfoResponse struct {
	AwipsPerApInfo []AwipsPerApInfo `json:"Cisco-IOS-XE-wireless-awips-oper:awips-per-ap-info"`
}

// AwipsOperDwldStatusResponse represents the AWIPS download status response
type AwipsOperDwldStatusResponse struct {
	AwipsDwldStatus AwipsDwldStatus `json:"Cisco-IOS-XE-wireless-awips-oper:awips-dwld-status"`
}

// AwipsOperApDwldStatusResponse represents the AWIPS AP download status response
type AwipsOperApDwldStatusResponse struct {
	AwipsApDwldStatus []AwipsApDwldStatus `json:"Cisco-IOS-XE-wireless-awips-oper:awips-ap-dwld-status"`
}

// AwipsDwldStatus contains AWIPS download status information
type AwipsDwldStatus struct {
	LastSuccessTimestamp string `json:"last-success-timestamp"`  // Timestamp of last successful download
	LastFailedTimestamp  string `json:"last-failed-timestamp"`   // Timestamp of last failed download
	NumOfFailureAttempts int    `json:"num-of-failure-attempts"` // Number of failure attempts
	LastFailureReason    int    `json:"last-failure-reason"`     // Reason code for last failure
	WlcVersion           string `json:"wlc-version"`             // WLC version
	MaxFileVer           int    `json:"max-file-ver"`            // Maximum file version
	LatestFileVersion    int    `json:"latest-file-version"`     // Latest file version
	DownloadStatus       string `json:"download-status"`         // Current download status
	FileHash             string `json:"file-hash"`               // File hash for verification
}

// AwipsPerApInfo contains AWIPS information for a specific access point
type AwipsPerApInfo struct {
	ApMac                 string `json:"ap-mac"`                  // Access point MAC address
	AwipsStatus           string `json:"awips-status"`            // AWIPS status for this AP
	AlarmCount            string `json:"alarm-count"`             // Number of alarms
	ForensicCaptureStatus string `json:"forensic-capture-status"` // Status of forensic capture
}

// AwipsApDwldStatus contains AWIPS download status for a specific access point
type AwipsApDwldStatus struct {
	ApMac       string `json:"ap-mac"`       // Access point MAC address
	DwldStatus  string `json:"dwld-status"`  // Download status
	FileVersion int    `json:"file-version"` // File version
	FileHash    string `json:"file-hash"`    // File hash for verification
}

// GetAwipsOper retrieves complete AWIPS operational data.
func GetAwipsOper(client *wnc.Client, ctx context.Context) (*AwipsOperResponse, error) {
	var data AwipsOperResponse
	if err := client.SendAPIRequest(ctx, AwipsOperEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetAwipsPerApInfo retrieves AWIPS per-AP information.
func GetAwipsPerApInfo(client *wnc.Client, ctx context.Context) (*AwipsOperPerApInfoResponse, error) {
	var data AwipsOperPerApInfoResponse
	if err := client.SendAPIRequest(ctx, AwipsPerApInfoEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetAwipsDwldStatus retrieves AWIPS download status.
func GetAwipsDwldStatus(client *wnc.Client, ctx context.Context) (*AwipsOperDwldStatusResponse, error) {
	var data AwipsOperDwldStatusResponse
	if err := client.SendAPIRequest(ctx, AwipsDwldStatusEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// GetAwipsApDwldStatus retrieves AWIPS AP download status.
func GetAwipsApDwldStatus(client *wnc.Client, ctx context.Context) (*AwipsOperApDwldStatusResponse, error) {
	var data AwipsOperApDwldStatusResponse
	if err := client.SendAPIRequest(ctx, AwipsApDwldStatusEndpoint, &data); err != nil {
		return nil, err
	}
	return &data, nil
}
