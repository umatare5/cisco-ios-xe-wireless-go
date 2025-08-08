// Minimal example: create a client and call one endpoint.
// Build: go build ./examples
// Run:   WNC_CONTROLLER="https://controller.example.com" WNC_ACCESS_TOKEN="<token>" go run ./examples/minimal
// Only use WithInsecureSkipVerify(true) for lab / self‑signed cert testing.
package main

import (
	"context"
	"fmt"
	"os"
	"time"

	wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
)

// fetchAPOper is an injectable function (overridden in tests) that returns the AP count.
var fetchAPOper = func(ctx context.Context, c *wnc.Client) (int, error) {
	apData, err := c.AP().GetOper(ctx)
	if err != nil {
		return 0, err
	}
	if apData == nil {
		return 0, fmt.Errorf("nil AP oper data")
	}
	return len(apData.CiscoIOSXEWirelessAccessPointOperAccessPointOperData.OperData), nil
}

// run performs the core logic; separated for testing.
func run(controller, token string, timeout time.Duration) (int, error) {
	client, err := wnc.NewClient(controller, token,
		wnc.WithTimeout(timeout),
		wnc.WithInsecureSkipVerify(true), // lab only
	)
	if err != nil {
		return 0, fmt.Errorf("create client: %w", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	count, err := fetchAPOper(ctx, client)
	if err != nil {
		return 0, fmt.Errorf("AP oper request: %w", err)
	}
	return count, nil
}

// realMain returns an exit code allowing tests to exercise all branches.
func realMain() int {
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
	count, err := run(controller, token, 30*time.Second)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	fmt.Printf("Successfully connected! Found %d APs\n", count)
	return 0
}

// exitFunc allows tests to intercept the exit code.
var exitFunc = os.Exit

func main() { exitFunc(realMain()) }
