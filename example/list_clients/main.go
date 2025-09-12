//go:build sample

// Package main in example/list_clients demonstrates how to list wireless clients (MAC and IP addresses) managed by a Cisco IOS-XE Wireless Network Controller.
package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
)

// ClientInfo holds basic client information for display.
type ClientInfo struct {
	MAC string
	IP  string
}

// run performs the core logic; separated for testing.
func run(controller, token string, logger *slog.Logger) ([]ClientInfo, error) {
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

	// Get SISF database which contains MAC to IP address bindings
	clientData, err := client.Client().GetOperational(ctx)
	if err != nil {
		return nil, fmt.Errorf("client oper request: %w", err)
	}

	var clients []ClientInfo
	for _, sisfData := range clientData.CiscoIOSXEWirelessClientOperClientOperData.SisfDBMac {
		clientInfo := ClientInfo{
			MAC: sisfData.MacAddr,
		}

		// Extract IPv4 address if available
		if sisfData.Ipv4Binding.IPKey.IPAddr != "" {
			clientInfo.IP = sisfData.Ipv4Binding.IPKey.IPAddr
		} else {
			clientInfo.IP = "N/A"
		}

		clients = append(clients, clientInfo)
	}

	return clients, nil
}

// start returns an exit code allowing tests to exercise all branches.
func start() int {
	// Setup structured logging
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	logger.Info("starting client listing tool",
		slog.String("operation", "list_clients"),
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

	clients, err := run(controller, token, logger)
	if err != nil {
		logger.Error("client listing failed",
			slog.String("controller", controller),
			slog.String("error", err.Error()))
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	logger.Info("client listing completed successfully",
		slog.String("controller", controller),
		slog.Int("client_count", len(clients)))

	fmt.Printf("Successfully connected! Found %d clients\n\n", len(clients))
	fmt.Println("MAC Address           | IP Address")
	fmt.Println("----------------------|----------------")
	for _, client := range clients {
		fmt.Printf("%-21s | %s\n", client.MAC, client.IP)
	}

	return 0
}

// exitFunc allows tests to intercept the exit code.
var exitFunc = os.Exit

func main() { exitFunc(start()) }
