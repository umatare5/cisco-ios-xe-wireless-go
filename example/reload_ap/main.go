//go:build example

// Package main in example/reload_ap demonstrates how to restart an access point managed by a Cisco IOS-XE Wireless Network Controller.
package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"strings"
	"time"

	wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
)

const (
	// Safety constants
	defaultTimeout       = 30 * time.Second
	defaultClientTimeout = 30 * time.Second
	confirmationPrompt   = "This will restart the specified Access Point(s). Type 'YES' to confirm: "

	// Environment variables
	envController = "WNC_CONTROLLER"
	envToken      = "WNC_ACCESS_TOKEN"
)

func main() {
	// Setup structured logging
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	logger.Info("starting AP reload tool",
		slog.String("operation", "reload_ap"),
		slog.String("version", "v1.0.0"))

	fmt.Println("=== Access Point Reload Tool ===")
	fmt.Println("WARNING: This tool will restart access points causing service interruption!")
	fmt.Println("Use only in controlled environments with proper authorization.")
	fmt.Println()

	// 1. Environment validation
	controller := os.Getenv(envController)
	token := os.Getenv(envToken)

	if controller == "" {
		logger.Error("environment validation failed",
			slog.String("missing_variable", envController))
		log.Fatalf("Error: %s environment variable not set", envController)
	}

	if token == "" {
		logger.Error("environment validation failed",
			slog.String("missing_variable", envToken))
		log.Fatalf("Error: %s environment variable not set", envToken)
	}

	logger.Info("environment validated",
		slog.String("controller", controller),
		slog.String("token_length", fmt.Sprintf("%d chars", len(token))))

	fmt.Printf("Target Controller: %s\n", controller)

	// 2. AP MAC address input
	fmt.Print("Enter AP MAC address (format: xx:xx:xx:xx:xx:xx or xx-xx-xx-xx-xx-xx): ")
	var apMacInput string
	if _, err := fmt.Scanln(&apMacInput); err != nil {
		logger.Error("failed to read MAC input",
			slog.String("error", err.Error()))
		log.Fatalf("Failed to read input: %v", err)
	}

	apMac := strings.TrimSpace(apMacInput)
	if apMac == "" {
		logger.Error("validation failed",
			slog.String("field", "ap_mac"),
			slog.String("reason", "empty"))
		log.Fatal("Error: AP MAC address is required")
	}

	logger.Info("AP MAC address provided",
		slog.String("ap_mac", apMac))

	fmt.Printf("Target AP MAC: %s\n", apMac)
	fmt.Println()

	// 3. Safety confirmation
	fmt.Print(confirmationPrompt)
	var confirmation string
	if _, err := fmt.Scanln(&confirmation); err != nil {
		logger.Error("failed to read confirmation",
			slog.String("error", err.Error()))
		log.Fatalf("Failed to read input: %v", err)
	}

	if confirmation != "YES" {
		logger.Info("operation canceled by user",
			slog.String("confirmation", confirmation))
		fmt.Println("Operation canceled.")
		os.Exit(0)
	}

	logger.Info("safety confirmation received")

	// 4. Client creation
	client, err := wnc.NewClient(controller, token,
		wnc.WithTimeout(defaultClientTimeout),
		wnc.WithInsecureSkipVerify(true), // lab only
		wnc.WithLogger(logger),
	)
	if err != nil {
		logger.Error("client creation failed",
			slog.String("controller", controller),
			slog.String("error", err.Error()))
		log.Fatalf("Failed to create WNC client: %v", err)
	}

	logger.Info("client created successfully",
		slog.String("controller", controller))
	fmt.Println("✓ WNC client created successfully")

	// 5. AP service setup
	apService := client.AP()

	// 6. Context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	// 7. Reload execution
	logger.Info("initiating AP reload",
		slog.String("ap_mac", apMac),
		slog.String("controller", controller))

	fmt.Printf("Executing AP reload for MAC %s\n", apMac)
	fmt.Println("WARNING: AP will become unavailable and disconnect all clients during restart...")

	err = apService.Reload(ctx, apMac)
	if err != nil {
		logger.Error("AP reload failed",
			slog.String("ap_mac", apMac),
			slog.String("controller", controller),
			slog.String("error", err.Error()))
		log.Printf("AP reload failed: %v", err)
		return
	}

	logger.Info("AP reload command sent successfully",
		slog.String("ap_mac", apMac),
		slog.String("controller", controller))

	fmt.Printf("✓ AP reload command sent successfully for MAC: %s\n", apMac)
	fmt.Println("Note: AP is now restarting and will be temporarily unavailable")
	fmt.Println("Clients will need to reconnect after AP restart completes")
}
