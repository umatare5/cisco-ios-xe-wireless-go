//go:build example

// Package main in example/save_config demonstrates how to copy the running configuration of a Cisco IOS-XE Wireless Network Controller to its startup configuration using the provided Go SDK.
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
	defaultTimeout = 30 * time.Second
	// The save took 2.5 to 3.7 seconds when measured, and WithTimeout does not lift the
	// five-second response-header budget, so this raises that one on its own.
	defaultHeaderTimeout = 30 * time.Second
	defaultClientTimeout = 30 * time.Second
	confirmationPrompt   = "This will overwrite the startup configuration. Type 'YES' to confirm: "

	// Environment variables
	envController = "WNC_CONTROLLER"
	envToken      = "WNC_ACCESS_TOKEN"
)

func main() {
	// Setup structured logging
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	logger.Info("starting configuration save tool",
		slog.String("operation", "save_config"))

	fmt.Println("=== WNC Configuration Save Tool ===")
	fmt.Println("WARNING: This tool overwrites the startup configuration and cannot be undone!")
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
		wnc.WithResponseHeaderTimeout(defaultHeaderTimeout),
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

	// 6. Save execution with structured logging
	logger.Info("starting configuration save",
		slog.String("controller", controller))

	fmt.Println("Executing configuration save...")

	out, err := controllerService.SaveConfig(ctx)
	if err != nil {
		logger.Error("configuration save failed",
			slog.String("controller", controller),
			slog.String("error", err.Error()))
		log.Printf("Configuration save failed: %v", err)
		return
	}

	// The controller's account is the only evidence the save happened: no release exposes the
	// startup datastore over RESTCONF, so an answer carrying none leaves the outcome unknown.
	if out.Output == nil || out.Output.Result == "" {
		logger.Warn("configuration save reported no result",
			slog.String("controller", controller))
		fmt.Println("✓ Save accepted, but the controller reported no result")
		return
	}

	logger.Info("configuration save reported by the controller",
		slog.String("controller", controller),
		slog.String("result", out.Output.Result))

	fmt.Printf("✓ %s\n", out.Output.Result)
}
