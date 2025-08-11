// Package main provides an advanced runnable example demonstrating robust client
// construction, context management, structured logging, basic error
// classification, and multi-service calls against the Cisco Wireless Network
// Controller (WNC) API. This is not part of the library API surface; it serves
// purely as documentation-through-code for experienced Go users.
//
// Build: go build ./examples/advanced
//
//	Run:   WNC_CONTROLLER="<controller-host-or-ip>" \
//	       WNC_ACCESS_TOKEN="<base64-username:password>" \
//	       go run ./examples/advanced
//
// Environment Variables:
//
//	WNC_CONTROLLER   Controller hostname or IP (required)
//	WNC_ACCESS_TOKEN RESTCONF access token (required)
//	WNC_TIMEOUT_SEC  Optional per-request timeout seconds (default: 15)
//
// Security Note:
//
//	The example enables wnc.WithInsecureSkipVerify(true) ONLY to simplify local
//	lab and self‑signed certificate testing. DO NOT use this in production.
//
// Demonstrated Concepts:
//   - Functional options (timeout, insecure TLS, custom logger, user agent)
//   - Context cancellation propagation
//   - Error wrapping + classification with errors.Is / errors.As
//   - Light introspection of returned data without binding to generated types
//   - Graceful resource cleanup and deterministic exit codes
//
// Exit Codes:
//
//	0 success
//	1 configuration / startup error
//	2 API (remote) error
//	3 unexpected internal error
package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"strings"
	"time"

	wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
)

// Injection points for tests (overridden in *_test.go).
var (
	exitFunc     = os.Exit
	buildClient  = newClient
	workflowFunc = runWorkflow
	// Individual service operation fetchers to allow targeted error simulation.
	fetchAPOper     func(ctx context.Context, c *wnc.Client) (any, error)
	fetchClientOper func(ctx context.Context, c *wnc.Client) (any, error)
	fetchRogueOper  func(ctx context.Context, c *wnc.Client) (any, error)
)

func main() {
	cfg, err := loadConfig()
	if err != nil {
		fatal(1, err)
		return
	}

	logger := newLogger(cfg.Verbose)
	logger.Info("starting advanced WNC example", "controller", cfg.Controller)

	client, err := buildClient(cfg, logger)
	if err != nil {
		fatal(1, fmt.Errorf("create client: %w", err))
		return
	}

	ctx, cancel := newRootContext(cfg.Timeout)
	defer cancel()

	// Collect a series of operations; each failure is classified.
	if err := workflowFunc(ctx, client, logger); err != nil {
		switch classifyError(err) {
		case errCategoryAuth, errCategoryForbidden:
			fatal(2, err)
		case errCategoryRemote, errCategoryNotFound, errCategoryTimeout:
			fatal(2, err)
		default:
			fatal(3, err)
		}
		return
	}

	logger.Info("workflow completed successfully")
	exitFunc(0)
}

// configuration holds runtime parameters.
type configuration struct {
	Controller  string
	AccessToken string
	Timeout     time.Duration
	Verbose     bool
}

// loadConfig sources configuration from flags and environment with precedence: flags > env > default.
func loadConfig() (*configuration, error) {
	var (
		flagTimeout = flag.Int("timeout", 0, "override request timeout in seconds (optional)")
		flagVerbose = flag.Bool("v", false, "enable verbose logging")
	)
	flag.Parse()

	controller := strings.TrimSpace(os.Getenv("WNC_CONTROLLER"))
	if controller == "" {
		return nil, errors.New("missing WNC_CONTROLLER environment variable")
	}
	accessToken := strings.TrimSpace(os.Getenv("WNC_ACCESS_TOKEN"))
	if accessToken == "" {
		return nil, errors.New("missing WNC_ACCESS_TOKEN environment variable")
	}

	// Base timeout: env or default 15s.
	timeout := 15 * time.Second
	if raw := os.Getenv("WNC_TIMEOUT_SEC"); raw != "" {
		if v, err := strconv.Atoi(raw); err == nil && v > 0 {
			timeout = time.Duration(v) * time.Second
		}
	}
	// Flag override highest precedence.
	if *flagTimeout > 0 {
		timeout = time.Duration(*flagTimeout) * time.Second
	}

	return &configuration{
		Controller:  controller,
		AccessToken: accessToken,
		Timeout:     timeout,
		Verbose:     *flagVerbose,
	}, nil
}

// newLogger builds a structured slog.Logger.
func newLogger(verbose bool) *slog.Logger {
	level := new(slog.LevelVar)
	if verbose {
		level.Set(slog.LevelDebug)
	} else {
		level.Set(slog.LevelInfo)
	}
	h := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: level})
	return slog.New(h)
}

// newClient creates a WNC client with functional options, including a custom logger and user agent.
func newClient(cfg *configuration, logger *slog.Logger) (*wnc.Client, error) {
	c, err := wnc.NewClient(cfg.Controller, cfg.AccessToken,
		wnc.WithTimeout(cfg.Timeout),
		wnc.WithInsecureSkipVerify(true), // lab only; remove for production
		wnc.WithLogger(logger),
		wnc.WithUserAgent("cisco-ios-xe-wireless-go-advanced-example/1.0"),
	)
	if err != nil {
		return nil, err
	}
	// Lazy default initialization for fetch functions (covered by tests calling newClient).
	if fetchAPOper == nil {
		fetchAPOper = func(ctx context.Context, c *wnc.Client) (any, error) { return c.AP().GetOper(ctx) }
	}
	if fetchClientOper == nil {
		fetchClientOper = func(ctx context.Context, c *wnc.Client) (any, error) { return c.Client().GetOper(ctx) }
	}
	if fetchRogueOper == nil {
		fetchRogueOper = func(ctx context.Context, c *wnc.Client) (any, error) { return c.Rogue().GetOper(ctx) }
	}
	return c, nil
}

// newRootContext constructs a base context with timeout; could be extended to support signals.
func newRootContext(d time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), d)
}

// runWorkflow performs multiple service calls demonstrating different response shapes.
func runWorkflow(ctx context.Context, client *wnc.Client, logger *slog.Logger) error {
	// AP operational data
	apOper, err := fetchAPOper(ctx, client)
	if err != nil {
		return fmt.Errorf("ap oper: %w", err)
	}
	logger.Info("retrieved AP operational data", "ptr", notNil(apOper))

	// Client operational summary (demonstrates another service)
	clientOper, err := fetchClientOper(ctx, client)
	if err != nil {
		return fmt.Errorf("client oper: %w", err)
	}
	logger.Info("retrieved Client operational data", "ptr", notNil(clientOper))

	// Rogue detection dataset (third service example)
	rogueOper, err := fetchRogueOper(ctx, client)
	if err != nil {
		return fmt.Errorf("rogue oper: %w", err)
	}
	logger.Info("retrieved Rogue operational data", "ptr", notNil(rogueOper))

	return nil
}

// notNil gives a concise bool for pointer presence to avoid logging large structures.
func notNil(v any) bool { return v != nil }

// Error classification categories for exit code mapping.
const (
	errCategoryAuth = iota + 1
	errCategoryForbidden
	errCategoryNotFound
	errCategoryTimeout
	errCategoryRemote
	errCategoryOther
)

// classifyError inspects wrapped errors for known sentinel or typed matches.
func classifyError(err error) int {
	if err == nil {
		return 0
	}
	// Unwrap chain checks.
	if errors.Is(err, wnc.ErrAuthenticationFailed) {
		return errCategoryAuth
	}
	if errors.Is(err, wnc.ErrAccessForbidden) {
		return errCategoryForbidden
	}
	if errors.Is(err, wnc.ErrResourceNotFound) {
		return errCategoryNotFound
	}
	if errors.Is(err, wnc.ErrRequestTimeout) || errors.Is(err, context.DeadlineExceeded) {
		return errCategoryTimeout
	}

	var apiErr *wnc.APIError
	if errors.As(err, &apiErr) {
		return errCategoryRemote
	}
	return errCategoryOther
}

// fatal prints an error and exits with the specified code.
func fatal(code int, err error) {
	if err == nil {
		return
	}
	_, _ = fmt.Fprintf(os.Stderr, "error: %v\n", err)
	exitFunc(code)
}
