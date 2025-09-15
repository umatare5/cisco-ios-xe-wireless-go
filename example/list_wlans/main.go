//go:build example

// Package main in example/list_wlans demonstrates how to list comprehensive WLAN configuration including Radio, WLAN, and BSSID information.
package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"sort"
	"time"

	wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
)

// WLANInfo holds comprehensive WLAN information for display.
type WLANInfo struct {
	APMAC       string // Access Point MAC address
	APName      string // Access Point name
	RadioSlotID int    // Radio slot identifier
	WlanID      int    // WLAN identifier
	BSSID       string // Basic Service Set Identifier
	SSID        string // Service Set Identifier
	Enabled     bool   // Whether the WLAN is enabled
}

// run performs the core logic; separated for testing.
func run(controller, token string, logger *slog.Logger) ([]WLANInfo, error) {
	client, err := wnc.NewClient(controller, token,
		wnc.WithTimeout(30*time.Second),
		wnc.WithInsecureSkipVerify(true), // lab only
		wnc.WithLogger(logger),
	)
	if err != nil {
		return nil, fmt.Errorf("create client: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// Get AP operational data which contains comprehensive WLAN information
	apData, err := client.AP().GetOperational(ctx)
	if err != nil {
		return nil, fmt.Errorf("AP oper request: %w", err)
	}

	// Create map for AP MAC to Name lookup
	nameMap := make(map[string]string)
	for _, nameMapping := range apData.CiscoIOSXEWirelessAccessPointOperAccessPointOperData.ApNameMACMap {
		nameMap[nameMapping.WtpMAC] = nameMapping.WtpName
	}

	var wlans []WLANInfo

	// Process all WLAN statistics which contains Radio + WLAN + BSSID information
	for _, wlanStats := range apData.CiscoIOSXEWirelessAccessPointOperAccessPointOperData.WtpSlotWlanStats {
		if wlanStats.BssidMAC == "" {
			continue // Skip entries without BSSID
		}

		wlanInfo := WLANInfo{
			WlanID:      wlanStats.WlanID,
			APMAC:       wlanStats.WtpMAC,
			RadioSlotID: wlanStats.SlotID,
			BSSID:       wlanStats.BssidMAC,
			SSID:        wlanStats.Ssid,
		}

		// Get AP name from the name map
		wlanInfo.APName = nameMap[wlanStats.WtpMAC]

		wlans = append(wlans, wlanInfo)
	}

	// Sort by AP MAC, then by Radio Slot, then by WLAN ID for consistent display
	sort.Slice(wlans, func(i, j int) bool {
		if wlans[i].APMAC != wlans[j].APMAC {
			return wlans[i].APMAC < wlans[j].APMAC
		}
		if wlans[i].RadioSlotID != wlans[j].RadioSlotID {
			return wlans[i].RadioSlotID < wlans[j].RadioSlotID
		}
		return wlans[i].WlanID < wlans[j].WlanID
	})

	return wlans, nil
}

// start returns an exit code allowing tests to exercise all branches.
func start() int {
	// Setup structured logging
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	logger.Info("starting WLAN listing tool",
		slog.String("operation", "list_wlans"),
		slog.String("version", "v1.0.0"))

	controller := os.Getenv("WNC_CONTROLLER")
	if controller == "" {
		logger.Error("environment validation failed",
			slog.String("missing_variable", "WNC_CONTROLLER"))
		fmt.Fprintln(os.Stderr, "WNC_CONTROLLER not set")
		return 1
	}

	token := os.Getenv("WNC_ACCESS_TOKEN")
	if token == "" {
		logger.Error("environment validation failed",
			slog.String("missing_variable", "WNC_ACCESS_TOKEN"))
		fmt.Fprintln(os.Stderr, "WNC_ACCESS_TOKEN not set")
		return 1
	}

	logger.Info("environment validated",
		slog.String("controller", controller),
		slog.String("token_length", fmt.Sprintf("%d chars", len(token))))

	wlans, err := run(controller, token, logger)
	if err != nil {
		logger.Error("WLAN listing failed",
			slog.String("controller", controller),
			slog.String("error", err.Error()))
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	logger.Info("WLAN listing completed successfully",
		slog.String("controller", controller),
		slog.Int("wlan_count", len(wlans)))

	fmt.Printf("Successfully connected! Found %d WLANs across all APs\n\n", len(wlans))

	if len(wlans) == 0 {
		fmt.Println("No WLANs found.")
		fmt.Println("\nTip: Make sure APs are operational and WLANs are configured.")
		return 0
	}

	fmt.Println("AP Name             | AP MAC Address    | Radio | WLAN | BSSID             | SSID")
	fmt.Println("--------------------|-------------------|-------|------|-------------------|-------------------------")

	for _, wlan := range wlans {
		fmt.Printf("%-19s | %-17s | %5d | %4d | %-17s | %s\n",
			wlan.APName,
			wlan.APMAC,
			wlan.RadioSlotID,
			wlan.WlanID,
			wlan.BSSID,
			wlan.SSID)
	}

	return 0
}

// exitFunc allows tests to intercept the exit code.
var exitFunc = os.Exit

func main() { exitFunc(start()) }
