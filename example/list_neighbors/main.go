//go:build example

// Package main in example/list_neighbors demonstrates how to list AP radio neighbors detected by a Cisco IOS-XE Wireless Network Controller.
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

// NeighborInfo holds AP radio neighbor information for display.
type NeighborInfo struct {
	DetectingAPMAC  string    // The AP that detected this neighbor
	DetectingAPName string    // Name of the detecting AP
	DetectingSlotID int       // Radio slot that detected the neighbor
	NeighborBSSID   string    // BSSID of the neighboring AP
	NeighborSSID    string    // SSID of the neighboring AP
	RSSI            int       // Signal strength
	PrimaryChannel  int       // Primary channel
	LastUpdateRcvd  time.Time // Last update timestamp
}

// run performs the core logic; separated for testing.
func run(controller, token string, logger *slog.Logger) ([]NeighborInfo, error) {
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

	// Get all radio neighbor information
	neighborData, err := client.AP().ListRadioNeighbors(ctx)
	if err != nil {
		return nil, fmt.Errorf("neighbor request: %w", err)
	}

	// Get AP name mapping for better display
	apData, err := client.AP().GetOperational(ctx)
	if err != nil {
		return nil, fmt.Errorf("AP oper request: %w", err)
	}

	// Create map for AP MAC to Name lookup
	nameMap := make(map[string]string)
	for _, nameMapping := range apData.CiscoIOSXEWirelessAPOperData.ApNameMACMap {
		nameMap[nameMapping.WtpMAC] = nameMapping.WtpName
	}

	var neighbors []NeighborInfo

	// Process all neighbor entries
	for _, neighbor := range neighborData.ApRadioNeighbor {
		neighborInfo := NeighborInfo{
			DetectingAPMAC:  neighbor.ApMAC,
			DetectingSlotID: neighbor.SlotID,
			NeighborBSSID:   neighbor.Bssid,
			NeighborSSID:    neighbor.Ssid,
			RSSI:            neighbor.RSSI,
			PrimaryChannel:  neighbor.PrimaryChannel,
			LastUpdateRcvd:  neighbor.LastUpdateRcvd,
		}

		// Get AP name from the name map
		neighborInfo.DetectingAPName = nameMap[neighbor.ApMAC]

		neighbors = append(neighbors, neighborInfo)
	}

	// Sort neighbors by detecting AP MAC and then by neighbor BSSID for consistent display
	sort.Slice(neighbors, func(i, j int) bool {
		if neighbors[i].DetectingAPMAC != neighbors[j].DetectingAPMAC {
			return neighbors[i].DetectingAPMAC < neighbors[j].DetectingAPMAC
		}
		return neighbors[i].NeighborBSSID < neighbors[j].NeighborBSSID
	})

	return neighbors, nil
}

// start returns an exit code allowing tests to exercise all branches.
func start() int {
	// Setup structured logging
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	logger.Info("starting neighbor listing tool",
		slog.String("operation", "list_neighbors"),
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

	neighbors, err := run(controller, token, logger)
	if err != nil {
		logger.Error("neighbor listing failed",
			slog.String("controller", controller),
			slog.String("error", err.Error()))
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	logger.Info("neighbor listing completed successfully",
		slog.String("controller", controller),
		slog.Int("neighbor_count", len(neighbors)))

	fmt.Printf("Successfully connected! Found %d AP neighbors\n\n", len(neighbors))

	if len(neighbors) == 0 {
		fmt.Println("No AP neighbors detected.")
		fmt.Println("\nTip: Neighbors are discovered when APs can hear each other's beacon frames.")
		fmt.Println("     Make sure you have multiple APs deployed and operational.")
		return 0
	}

	fmt.Println("AP Name           | Slot | Neighbor BSSID    | Neighbor SSID          | RSSI  | Channel | Last Heard At")
	fmt.Println("------------------|------|-------------------|------------------------|-------|---------|--------------------------")

	for _, neighbor := range neighbors {
		// Format timestamp for better readability
		timeStr := neighbor.LastUpdateRcvd.Format("2006-01-02 15:04:05")

		fmt.Printf("%-17s | %4d | %-17s | %-22s | %5d | %7d | %s\n",
			neighbor.DetectingAPName,
			neighbor.DetectingSlotID,
			neighbor.NeighborBSSID,
			neighbor.NeighborSSID,
			neighbor.RSSI,
			neighbor.PrimaryChannel,
			timeStr)
	}

	return 0
}

// exitFunc allows tests to intercept the exit code.
var exitFunc = os.Exit

func main() { exitFunc(start()) }
