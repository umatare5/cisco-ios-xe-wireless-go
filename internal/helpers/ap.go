// Package helpers provides helper functions for ap service
package helpers

import (
	model "github.com/umatare5/cisco-ios-xe-wireless-go/internal/model/ap"
	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/validation"
)

// FindAPByMAC searches for an AP with the given MAC address in CAPWAP data
func FindAPByMAC(capwapData *model.ApOperCapwapData, apMac string) (string, bool) {
	if capwapData == nil {
		return "", false
	}
	for _, data := range capwapData.CapwapData {
		if validation.FieldMatches(data.WtpMac, apMac) {
			return data.Name, true
		}
	}
	return "", false
}
