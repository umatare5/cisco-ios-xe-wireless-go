//go:build samples || ignore

// Package main provides a sample program for Access Point reload operations.
//
// WARNING: This program will restart specific access points causing service interruption
// for clients connected to those APs. Only use this in controlled environments with proper authorization.
//
// Usage:
//
//	export WNC_CONTROLLER="192.168.1.100"
//	export WNC_ACCESS_TOKEN="your-token-here"
//	go run samples/reload_ap/main.go
//
// This program is EXCLUDED from integration tests for safety reasons.
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
)

const (
	// Safety constants
	defaultTimeout     = 30 * time.Second
	confirmationPrompt = "This will restart the specified Access Point(s). Type 'YES' to confirm: "

	// Environment variables
	envController = "WNC_CONTROLLER"
	envToken      = "WNC_ACCESS_TOKEN"
)

func main() {
	fmt.Println("=== Access Point Reload Tool ===")
	fmt.Println("WARNING: This tool will restart access points causing service interruption!")
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

	// 2. AP MAC address input
	fmt.Print("Enter AP MAC address (format: xx:xx:xx:xx:xx:xx or xx-xx-xx-xx-xx-xx): ")
	var apMacInput string
	fmt.Scanln(&apMacInput)

	apMac := strings.TrimSpace(apMacInput)
	if apMac == "" {
		log.Fatal("Error: AP MAC address is required")
	}

	fmt.Printf("Target AP MAC: %s\n", apMac)
	fmt.Println()

	// 3. Safety confirmation
	fmt.Print(confirmationPrompt)
	var confirmation string
	fmt.Scanln(&confirmation)

	if confirmation != "YES" {
		fmt.Println("Operation cancelled.")
		os.Exit(0)
	}

	// 4. Client creation
	client, err := wnc.NewClient(controller, token)
	if err != nil {
		log.Fatalf("Failed to create WNC client: %v", err)
	}

	fmt.Println("✓ WNC client created successfully")

	// 5. AP service setup
	apService := client.AP()

	// 6. Context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	// 7. Reload execution
	fmt.Printf("Executing AP reload for MAC %s\n", apMac)
	fmt.Println("WARNING: AP will become unavailable and disconnect all clients during restart...")

	err = apService.Reload(ctx, apMac)
	if err != nil {
		log.Fatalf("AP reload failed: %v", err)
	}

	fmt.Printf("✓ AP reload command sent successfully for MAC: %s\n", apMac)
	fmt.Println("Note: AP is now restarting and will be temporarily unavailable")
	fmt.Println("Clients will need to reconnect after AP restart completes")
}
