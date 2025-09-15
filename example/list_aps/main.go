//go:build example

// Package main in example/list_aps demonstrates how to list access points managed by a Cisco IOS-XE Wireless Network Controller.
package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
)

// APInfo holds basic AP information for display.
type APInfo struct {
	MAC    string
	Name   string
	IP     string
	Status string
}

// run performs the core logic; separated for testing.
func run(controller, token string, logger *slog.Logger) ([]APInfo, error) {
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

	// Get AP operational data which contains comprehensive AP information
	apData, err := client.AP().GetOperational(ctx)
	if err != nil {
		return nil, fmt.Errorf("AP oper request: %w", err)
	}

	var aps []APInfo

	// Create map for AP MAC to Name lookup
	nameMap := make(map[string]string)
	for _, nameMapping := range apData.CiscoIOSXEWirelessAccessPointOperAccessPointOperData.ApNameMACMap {
		nameMap[nameMapping.WtpMAC] = nameMapping.WtpName
	}

	// Extract AP information from CAPWAP data
	for _, capwapData := range apData.CiscoIOSXEWirelessAccessPointOperAccessPointOperData.CAPWAPData {
		apInfo := APInfo{
			MAC: capwapData.WtpMAC,
			IP:  capwapData.IPAddr,
		}

		// Get AP name from the name map
		apInfo.Name = nameMap[capwapData.WtpMAC]

		// Get status from AP state
		if capwapData.ApState.ApOperationState != "" {
			apInfo.Status = capwapData.ApState.ApOperationState
		} else {
			apInfo.Status = "N/A"
		}

		aps = append(aps, apInfo)
	}

	return aps, nil
}

// start returns an exit code allowing tests to exercise all branches.
func start() int {
	// Setup structured logging
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	logger.Info("starting AP listing tool",
		slog.String("operation", "list_aps"),
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

	aps, err := run(controller, token, logger)
	if err != nil {
		logger.Error("AP listing failed",
			slog.String("controller", controller),
			slog.String("error", err.Error()))
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	logger.Info("AP listing completed successfully",
		slog.String("controller", controller),
		slog.Int("ap_count", len(aps)))

	fmt.Printf("Successfully connected! Found %d APs\n\n", len(aps))
	fmt.Println("AP Name           | MAC Address         | IP Address       | Status")
	fmt.Println("------------------|---------------------|------------------|-----------------")
	for _, ap := range aps {
		fmt.Printf("%-17s | %-19s | %-16s | %-15s\n", ap.Name, ap.MAC, ap.IP, ap.Status)
	}

	return 0
}

// exitFunc allows tests to intercept the exit code.
var exitFunc = os.Exit

func main() { exitFunc(start()) }
