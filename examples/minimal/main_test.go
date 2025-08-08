package main

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	wnc "github.com/umatare5/cisco-ios-xe-wireless-go"
)

// TestRealMainEnvMissing covers missing env branches.
func TestRealMainEnvMissing(t *testing.T) {
	oldC, oldT := os.Getenv("WNC_CONTROLLER"), os.Getenv("WNC_ACCESS_TOKEN")
	defer os.Setenv("WNC_CONTROLLER", oldC)
	defer os.Setenv("WNC_ACCESS_TOKEN", oldT)
	os.Unsetenv("WNC_CONTROLLER")
	os.Setenv("WNC_ACCESS_TOKEN", "token")
	if code := realMain(); code != 1 {
		t.Fatalf("expected exit 1 for missing controller, got %d", code)
	}
	os.Setenv("WNC_CONTROLLER", "ctrl")
	os.Unsetenv("WNC_ACCESS_TOKEN")
	if code := realMain(); code != 1 {
		t.Fatalf("expected exit 1 for missing token, got %d", code)
	}
}

// TestRunSuccess covers success path using stubbed fetch function.
func TestRunSuccess(t *testing.T) {
	old := fetchAPOper
	defer func() { fetchAPOper = old }()
	fetchAPOper = func(_ context.Context, _ *wnc.Client) (int, error) { return 5, nil }
	count, err := run("controller.example.com", "abcd", 5*time.Second)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if count != 5 {
		t.Fatalf("expected 5, got %d", count)
	}
}

// TestRunFailure ensures AP oper failure propagates.
func TestRunFailure(t *testing.T) {
	old := fetchAPOper
	defer func() { fetchAPOper = old }()
	fetchAPOper = func(_ context.Context, _ *wnc.Client) (int, error) { return 0, errors.New("boom") }
	_, err := run("controller.example.com", "abcd", 5*time.Second)
	if err == nil || err.Error() == "" {
		t.Fatalf("expected wrapped error, got %v", err)
	}
}

// TestRunClientCreationError covers the error path when WithTimeout returns error (timeout <=0).
func TestRunClientCreationError(t *testing.T) {
	// Negative or zero timeout triggers option error.
	_, err := run("controller.example.com", "abcd", 0)
	if err == nil || err.Error() == "" || !contains(err.Error(), "create client") {
		t.Fatalf("expected client creation error, got %v", err)
	}
}

// TestRealMainSuccess covers fully successful execution.
func TestRealMainSuccess(t *testing.T) {
	old := fetchAPOper
	defer func() { fetchAPOper = old }()
	fetchAPOper = func(_ context.Context, _ *wnc.Client) (int, error) { return 1, nil }
	oldC, oldT := os.Getenv("WNC_CONTROLLER"), os.Getenv("WNC_ACCESS_TOKEN")
	defer os.Setenv("WNC_CONTROLLER", oldC)
	defer os.Setenv("WNC_ACCESS_TOKEN", oldT)
	os.Setenv("WNC_CONTROLLER", "controller.example.com")
	os.Setenv("WNC_ACCESS_TOKEN", "token")
	if code := realMain(); code != 0 {
		t.Fatalf("expected exit 0, got %d", code)
	}
}

// TestRealMainRunFailure covers failure after run() returns error.
func TestRealMainRunFailure(t *testing.T) {
	old := fetchAPOper
	defer func() { fetchAPOper = old }()
	fetchAPOper = func(_ context.Context, _ *wnc.Client) (int, error) { return 0, errors.New("ap fail") }
	oldC, oldT := os.Getenv("WNC_CONTROLLER"), os.Getenv("WNC_ACCESS_TOKEN")
	defer os.Setenv("WNC_CONTROLLER", oldC)
	defer os.Setenv("WNC_ACCESS_TOKEN", oldT)
	os.Setenv("WNC_CONTROLLER", "controller.example.com")
	os.Setenv("WNC_ACCESS_TOKEN", "token")
	if code := realMain(); code != 1 {
		t.Fatalf("expected exit 1, got %d", code)
	}
}

// TestMainFunction covers the main() function via exitFunc injection.
func TestMainFunction(t *testing.T) {
	oldExit := exitFunc
	defer func() { exitFunc = oldExit }()
	codeCh := make(chan int, 1)
	exitFunc = func(code int) { codeCh <- code }
	old := fetchAPOper
	fetchAPOper = func(_ context.Context, _ *wnc.Client) (int, error) { return 2, nil }
	defer func() { fetchAPOper = old }()
	oldC, oldT := os.Getenv("WNC_CONTROLLER"), os.Getenv("WNC_ACCESS_TOKEN")
	defer os.Setenv("WNC_CONTROLLER", oldC)
	defer os.Setenv("WNC_ACCESS_TOKEN", oldT)
	os.Setenv("WNC_CONTROLLER", "controller.example.com")
	os.Setenv("WNC_ACCESS_TOKEN", "token")
	main()
	select {
	case code := <-codeCh:
		if code != 0 {
			t.Fatalf("expected exit code 0, got %d", code)
		}
	case <-time.After(1 * time.Second):
		t.Fatal("main did not call exitFunc")
	}
}

// contains is a small helper to avoid importing strings repeatedly.
func contains(s, sub string) bool {
	if len(sub) == 0 {
		return true
	}
	if len(sub) > len(s) {
		return false
	}
	for i := 0; i+len(sub) <= len(s); i++ {
		if s[i:i+len(sub)] == sub {
			return true
		}
	}
	return false
}
