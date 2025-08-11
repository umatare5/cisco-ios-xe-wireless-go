// Minimal example: create a client and call one endpoint.
// Build: go build ./examples
// Run:   WNC_CONTROLLER="<controller-host-or-ip>" WNC_ACCESS_TOKEN="<base64-username:password>" go run ./examples/minimal
// Note: WithInsecureSkipVerify(true) is for lab/self-signed cert testing only.
package main

import (
	"context"
	"fmt"
	"os"
	"time"

	wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
)

// run performs the core logic; separated for testing.
func run(controller, token string) (int, error) {
	client, err := wnc.NewClient(controller, token,
		wnc.WithTimeout(30*time.Second),
		wnc.WithInsecureSkipVerify(true), // lab only
	)
	if err != nil {
		return 0, fmt.Errorf("create client: %w", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	apData, err := client.AP().GetOper(ctx)
	if err != nil {
		return 0, fmt.Errorf("AP oper request: %w", err)
	}
	return len(apData.CiscoIOSXEWirelessAccessPointOperAccessPointOperData.OperData), nil
}

// start returns an exit code allowing tests to exercise all branches.
func start() int {
	controller := os.Getenv("WNC_CONTROLLER")
	if controller == "" {
		fmt.Fprintln(os.Stderr, "WNC_CONTROLLER not set")
		return 1
	}
	token := os.Getenv("WNC_ACCESS_TOKEN")
	if token == "" {
		fmt.Fprintln(os.Stderr, "WNC_ACCESS_TOKEN not set")
		return 1
	}
	count, err := run(controller, token)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	fmt.Printf("Successfully connected! Found %d APs\n", count)
	return 0
}

// exitFunc allows tests to intercept the exit code.
var exitFunc = os.Exit

func main() { exitFunc(start()) }
