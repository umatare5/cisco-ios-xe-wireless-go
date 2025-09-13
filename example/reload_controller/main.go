//go:build example

// Package main in example/reload_controller demonstrates how to restart a Cisco IOS-XE Wireless Network Controller using the provided Go client library.
package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"time"

	wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
)

const (
	// Safety constants
	defaultTimeout       = 30 * time.Second
	defaultClientTimeout = 30 * time.Second
	confirmationPrompt   = "This will restart the WNC controller. Type 'YES' to confirm: "

	// Environment variables
	envController = "WNC_CONTROLLER"
	envToken      = "WNC_ACCESS_TOKEN"
)

func main() {
	// Setup structured logging
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	logger.Info("starting controller reload tool",
		slog.String("operation", "controller_reload"),
		slog.String("version", "v1.0.0"))

	fmt.Println("=== WNC Controller Reload Tool ===")
	fmt.Println("WARNING: This tool will restart the wireless controller!")
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
	fmt.Println()

	// 2. Safety confirmation
	fmt.Print(confirmationPrompt)
	var confirmation string
	if _, err := fmt.Scanln(&confirmation); err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}

	if confirmation != "YES" {
		fmt.Println("Operation canceled.")
		os.Exit(0)
	}

	// 3. Client creation with structured logging
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

	// 4. Controller service setup
	controllerService := client.Controller()

	// 5. Context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	// 6. Reload execution with structured logging
	reason := "Manual reload via CLI tool at " + time.Now().Format(time.RFC3339)

	logger.Info("starting controller reload",
		slog.String("reason", reason),
		slog.String("controller", controller))

	fmt.Printf("Executing controller reload with reason: %s\n", reason)
	fmt.Println("WARNING: Controller will become unavailable during restart...")

	err = controllerService.ReloadWithReason(ctx, reason)
	if err != nil {
		logger.Error("controller reload failed",
			slog.String("controller", controller),
			slog.String("reason", reason),
			slog.String("error", err.Error()))
		log.Printf("Controller reload failed: %v", err)
		return
	}

	logger.Info("controller reload command sent successfully",
		slog.String("controller", controller),
		slog.String("reason", reason))

	fmt.Println("✓ Controller reload command sent successfully")
	fmt.Println("Note: Controller is now restarting and will be temporarily unavailable")
	fmt.Println("Wait for controller to complete restart before attempting reconnection")
}
