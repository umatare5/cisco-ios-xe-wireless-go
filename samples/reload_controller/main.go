//go:build samples || ignore

// Package main provides a sample program for WNC Controller reload operations.
//
// WARNING: This program will restart the wireless controller causing complete service interruption.
// Only use this in controlled environments with proper authorization.
//
// Usage:
//   export WNC_CONTROLLER="192.168.1.100"
//   export WNC_ACCESS_TOKEN="your-token-here"
//   go run examples/reload_controller/main.go
//
// This program is EXCLUDED from integration tests for safety reasons.

package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
)

const (
	// Safety constants
	defaultTimeout     = 30 * time.Second
	confirmationPrompt = "This will restart the WNC controller. Type 'YES' to confirm: "

	// Environment variables
	envController = "WNC_CONTROLLER"
	envToken      = "WNC_ACCESS_TOKEN"
)

func main() {
	fmt.Println("=== WNC Controller Reload Tool ===")
	fmt.Println("WARNING: This tool will restart the wireless controller!")
	fmt.Println("Use only in controlled environments with proper authorization.")
	fmt.Println()

	// 1. Environment validation
	controller := os.Getenv(envController)
	token := os.Getenv(envToken)

	if controller == "" {
		log.Fatalf("Error: %s environment variable not set", envController)
	}

	if token == "" {
		log.Fatalf("Error: %s environment variable not set", envToken)
	}

	fmt.Printf("Target Controller: %s\n", controller)
	fmt.Println()

	// 2. Safety confirmation
	fmt.Print(confirmationPrompt)
	var confirmation string
	fmt.Scanln(&confirmation)

	if confirmation != "YES" {
		fmt.Println("Operation cancelled.")
		os.Exit(0)
	}

	// 3. Client creation
	client, err := wnc.NewClient(controller, token)
	if err != nil {
		log.Fatalf("Failed to create WNC client: %v", err)
	}

	fmt.Println("✓ WNC client created successfully")

	// 4. Controller service setup
	controllerService := client.Controller()

	// 5. Context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	// 6. Reload execution
	reason := fmt.Sprintf("Manual reload via CLI tool at %s", time.Now().Format(time.RFC3339))

	fmt.Printf("Executing controller reload with reason: %s\n", reason)
	fmt.Println("WARNING: Controller will become unavailable during restart...")

	err = controllerService.ReloadWithReason(ctx, reason)
	if err != nil {
		log.Fatalf("Controller reload failed: %v", err)
	}

	fmt.Println("✓ Controller reload command sent successfully")
	fmt.Println("Note: Controller is now restarting and will be temporarily unavailable")
	fmt.Println("Wait for controller to complete restart before attempting reconnection")
}
