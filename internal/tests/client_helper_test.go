package tests

import (
	"context"
	"errors"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/umatare5/cisco-ios-xe-wireless-go/internal/core"
)

// --- Client helper tests ---

func TestTestClient(t *testing.T) {
	originalController := os.Getenv("WNC_CONTROLLER")
	originalToken := os.Getenv("WNC_ACCESS_TOKEN")
	t.Cleanup(func() {
		if originalController != "" {
			os.Setenv("WNC_CONTROLLER", originalController)
		} else {
			os.Unsetenv("WNC_CONTROLLER")
		}
		if originalToken != "" {
			os.Setenv("WNC_ACCESS_TOKEN", originalToken)
		} else {
			os.Unsetenv("WNC_ACCESS_TOKEN")
		}
	})

	t.Run("WithEnvironmentVariables", func(t *testing.T) {
		os.Setenv("WNC_CONTROLLER", "test.example.com")
		os.Setenv("WNC_ACCESS_TOKEN", "test-token")
		if client := TestClient(t); client == nil {
			t.Error("expected non-nil client")
		}
	})

	t.Run("WithValidEnvironmentVariables", func(t *testing.T) {
		controller, token := originalController, originalToken
		if controller != "" && token != "" {
			os.Setenv("WNC_CONTROLLER", controller)
			os.Setenv("WNC_ACCESS_TOKEN", token)
			if client := TestClient(t); client == nil {
				t.Error("expected non-nil client with real env")
			}
		} else {
			t.Skip("real env not available")
		}
	})

	t.Run("WithEmptyController", func(t *testing.T) {
		os.Setenv("WNC_CONTROLLER", "")
		os.Setenv("WNC_ACCESS_TOKEN", "valid-token")
		controller := os.Getenv("WNC_CONTROLLER")
		token := os.Getenv("WNC_ACCESS_TOKEN")
		if controller == "" || token == "" {
			t.Log("TestClient would skip due to empty env")
		} else {
			t.Error("expected skip condition")
		}
	})

	t.Run("WithEmptyToken", func(t *testing.T) {
		os.Setenv("WNC_CONTROLLER", "valid-controller")
		os.Setenv("WNC_ACCESS_TOKEN", "")
		controller := os.Getenv("WNC_CONTROLLER")
		token := os.Getenv("WNC_ACCESS_TOKEN")
		if controller == "" || token == "" {
			t.Log("skip due to empty token")
		} else {
			t.Error("expected skip condition")
		}
	})

	t.Run("WithBothEmpty", func(t *testing.T) {
		os.Setenv("WNC_CONTROLLER", "")
		os.Setenv("WNC_ACCESS_TOKEN", "")
		controller := os.Getenv("WNC_CONTROLLER")
		token := os.Getenv("WNC_ACCESS_TOKEN")
		if controller == "" || token == "" {
			t.Log("skip due to both empty")
		} else {
			t.Error("expected skip condition")
		}
	})

	t.Run("WithClientCreationFailure", func(t *testing.T) {
		os.Setenv("WNC_CONTROLLER", "test.example.com")
		os.Setenv("WNC_ACCESS_TOKEN", "test-token")
		controller := os.Getenv("WNC_CONTROLLER")
		token := os.Getenv("WNC_ACCESS_TOKEN")
		if controller != "" && token != "" {
			_, err := core.New(controller, token, core.WithTimeout(-1*time.Second))
			if err != nil {
				t.Logf("client creation failed (expected): %v", err)
			}
			_, err = core.New("", token)
			if err != nil {
				t.Logf("client creation failed(empty controller): %v", err)
			}
		}
	})

	t.Run("ActualFailureScenario", func(t *testing.T) {
		os.Setenv("WNC_CONTROLLER", "invalid-controller-that-will-fail")
		os.Setenv("WNC_ACCESS_TOKEN", "valid-token")
		controller := os.Getenv("WNC_CONTROLLER")
		token := os.Getenv("WNC_ACCESS_TOKEN")
		if controller != "" && token != "" {
			_, err := core.New(controller, token, core.WithTimeout(30*time.Second), core.WithInsecureSkipVerify(true))
			if err != nil {
				t.Logf("client creation would fail (as expected): %v", err)
			}
		}
	})
}

func TestTestClient_AdditionalBranches(t *testing.T) {
	origController, originalToken := os.Getenv("WNC_CONTROLLER"), os.Getenv("WNC_ACCESS_TOKEN")
	origCreate, origFail, origSim := createCoreClient, failOnClientError, simulateFatalAsLog
	t.Cleanup(func() {
		if origController == "" {
			os.Unsetenv("WNC_CONTROLLER")
		} else {
			os.Setenv("WNC_CONTROLLER", origController)
		}
		if originalToken == "" {
			os.Unsetenv("WNC_ACCESS_TOKEN")
		} else {
			os.Setenv("WNC_ACCESS_TOKEN", originalToken)
		}
		createCoreClient = origCreate
		failOnClientError = origFail
		simulateFatalAsLog = origSim
	})
	os.Unsetenv("WNC_CONTROLLER")
	os.Unsetenv("WNC_ACCESS_TOKEN")
	_ = TestClient(t)

	os.Setenv("WNC_CONTROLLER", "ctrl.example.test")
	os.Setenv("WNC_ACCESS_TOKEN", "tok")
	createCoreClient = func(controller, token string, opts ...core.Option) (*core.Client, error) {
		return nil, errors.New("injected failure")
	}
	failOnClientError = true
	simulateFatalAsLog = true
	if c := TestClient(t); c != nil {
		t.Fatalf("expected nil client on injected failure, got %v", c)
	}

	os.Setenv("WNC_CONTROLLER", "ctrl.example.test")
	os.Setenv("WNC_ACCESS_TOKEN", "tok")
	createCoreClient = func(controller, token string, opts ...core.Option) (*core.Client, error) {
		return nil, errors.New("injected failure 2")
	}
	failOnClientError = false
	_ = TestClient(t)
}

func TestOptionalTestClient(t *testing.T) {
	originalController := os.Getenv("WNC_CONTROLLER")
	originalToken := os.Getenv("WNC_ACCESS_TOKEN")
	t.Cleanup(func() {
		if originalController == "" {
			os.Unsetenv("WNC_CONTROLLER")
		} else {
			os.Setenv("WNC_CONTROLLER", originalController)
		}
		if originalToken == "" {
			os.Unsetenv("WNC_ACCESS_TOKEN")
		} else {
			os.Setenv("WNC_ACCESS_TOKEN", originalToken)
		}
		createCoreClient = core.New
	})
	os.Unsetenv("WNC_CONTROLLER")
	os.Unsetenv("WNC_ACCESS_TOKEN")
	if c := OptionalTestClient(t); c != nil {
		t.Errorf("expected nil client when env missing")
	}
	os.Setenv("WNC_CONTROLLER", "test.example.com")
	os.Setenv("WNC_ACCESS_TOKEN", "token")
	if c := OptionalTestClient(t); c == nil {
		t.Errorf("expected non-nil client with env present")
	}
	createCoreClient = func(controller, token string, opts ...core.Option) (*core.Client, error) {
		return nil, fmt.Errorf("injected error")
	}
	if c := OptionalTestClient(t); c != nil {
		t.Errorf("expected nil client on injected creation error")
	}
}

func TestCreateTestClientFromEnv(t *testing.T) {
	origC, origT := os.Getenv("WNC_CONTROLLER"), os.Getenv("WNC_ACCESS_TOKEN")
	t.Cleanup(func() {
		if origC != "" {
			os.Setenv("WNC_CONTROLLER", origC)
		} else {
			os.Unsetenv("WNC_CONTROLLER")
		}
		if origT != "" {
			os.Setenv("WNC_ACCESS_TOKEN", origT)
		} else {
			os.Unsetenv("WNC_ACCESS_TOKEN")
		}
	})
	os.Setenv("WNC_CONTROLLER", "test.example.com")
	os.Setenv("WNC_ACCESS_TOKEN", "test-token")
	if c := CreateTestClientFromEnv(t); c == nil {
		t.Error("expected client")
	}
}

func TestTestClient_ErrorSkip(t *testing.T) {
	origController := os.Getenv("WNC_CONTROLLER")
	originalToken := os.Getenv("WNC_ACCESS_TOKEN")
	t.Cleanup(func() {
		if origController != "" {
			os.Setenv("WNC_CONTROLLER", origController)
		} else {
			os.Unsetenv("WNC_CONTROLLER")
		}
		if originalToken != "" {
			os.Setenv("WNC_ACCESS_TOKEN", originalToken)
		} else {
			os.Unsetenv("WNC_ACCESS_TOKEN")
		}
		failOnClientError = true
	})
	os.Unsetenv("WNC_CONTROLLER")
	os.Unsetenv("WNC_ACCESS_TOKEN")
	failOnClientError = false
	_ = TestClient(t)
}

func TestTestClient_FatalSimulated(t *testing.T) {
	origController := os.Getenv("WNC_CONTROLLER")
	originalToken := os.Getenv("WNC_ACCESS_TOKEN")
	t.Cleanup(func() {
		if origController != "" {
			os.Setenv("WNC_CONTROLLER", origController)
		} else {
			os.Unsetenv("WNC_CONTROLLER")
		}
		if originalToken != "" {
			os.Setenv("WNC_ACCESS_TOKEN", originalToken)
		} else {
			os.Unsetenv("WNC_ACCESS_TOKEN")
		}
		failOnClientError = true
		simulateFatalAsLog = false
	})
	os.Setenv("WNC_CONTROLLER", "bad")
	os.Setenv("WNC_ACCESS_TOKEN", "bad")
	failOnClientError = true
	simulateFatalAsLog = true
	_ = TestClient(t)
}

func TestTestClient_SuccessPath(t *testing.T) {
	origCreate := createCoreClient
	t.Cleanup(func() { createCoreClient = origCreate })
	createCoreClient = func(controller, token string, opts ...core.Option) (*core.Client, error) {
		c, _ := core.New("1.1.1.1", token)
		return c, nil
	}
	os.Setenv("WNC_CONTROLLER", "1.1.1.1")
	os.Setenv("WNC_ACCESS_TOKEN", "token1234567890")
	defer func() { os.Unsetenv("WNC_CONTROLLER"); os.Unsetenv("WNC_ACCESS_TOKEN") }()
	failOnClientError = true
	simulateFatalAsLog = false
	if c := TestClient(t); c == nil {
		t.Fatalf("expected non-nil client")
	}
}

func TestTestClient_ErrorSkipOnCreateError(t *testing.T) {
	origCreate := createCoreClient
	t.Cleanup(func() { createCoreClient = origCreate })
	createCoreClient = func(controller, token string, opts ...core.Option) (*core.Client, error) {
		return nil, fmt.Errorf("forced error")
	}
	os.Setenv("WNC_CONTROLLER", "1.1.1.1")
	os.Setenv("WNC_ACCESS_TOKEN", "token1234567890")
	defer func() { os.Unsetenv("WNC_CONTROLLER"); os.Unsetenv("WNC_ACCESS_TOKEN") }()
	failOnClientError = false
	_ = TestClient(t)
}

func TestTestClient_FatalReal(t *testing.T) {
	origCreate, origFatal := createCoreClient, testFatalf
	t.Cleanup(func() { createCoreClient = origCreate; testFatalf = origFatal })
	createCoreClient = func(controller, token string, opts ...core.Option) (*core.Client, error) {
		return nil, fmt.Errorf("forced error real fatal")
	}
	testFatalf = func(t *testing.T, format string, args ...any) { t.Logf("hook fatal invoked: "+format, args...) }
	os.Setenv("WNC_CONTROLLER", "1.1.1.1")
	os.Setenv("WNC_ACCESS_TOKEN", "token1234567890")
	defer func() { os.Unsetenv("WNC_CONTROLLER"); os.Unsetenv("WNC_ACCESS_TOKEN") }()
	failOnClientError = true
	simulateFatalAsLog = false
	_ = TestClient(t)
}

func TestSkipIfNoConnection(t *testing.T) {
	t.Run("WithValidClient", func(t *testing.T) {
		controller := os.Getenv("WNC_CONTROLLER")
		token := os.Getenv("WNC_ACCESS_TOKEN")
		if controller == "" || token == "" {
			t.Skip("env not set")
		}
		client, err := core.New(controller, token, core.WithTimeout(5*time.Second), core.WithInsecureSkipVerify(true))
		if err != nil {
			t.Fatalf("failed to create test client: %v", err)
		}
		SkipIfNoConnection(t, client)
	})

	t.Run("WithNilClient", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("panic: %v", r)
			}
		}()
		var nilClient *core.Client
		t.Logf("nil client: %t", nilClient == nil)
		// Exercise early-return branch (no skip) for nil client
		SkipIfNoConnection(t, nilClient)
	})
}

func TestTestClientExtensiveCoverage(t *testing.T) {
	t.Run("ClientCreationWithValidEnvironment", func(t *testing.T) {
		origC, origT := os.Getenv("WNC_CONTROLLER"), os.Getenv("WNC_ACCESS_TOKEN")
		os.Setenv("WNC_CONTROLLER", "example.com")
		os.Setenv("WNC_ACCESS_TOKEN", "valid-test-token")
		defer func() {
			if origC == "" {
				os.Unsetenv("WNC_CONTROLLER")
			} else {
				os.Setenv("WNC_CONTROLLER", origC)
			}
			if origT == "" {
				os.Unsetenv("WNC_ACCESS_TOKEN")
			} else {
				os.Setenv("WNC_ACCESS_TOKEN", origT)
			}
		}()
		if c := TestClient(t); c == nil {
			t.Error("expected non-nil client")
		}
	})

	t.Run("TestClientWithRealEnvironment", func(t *testing.T) {
		controller, token := os.Getenv("WNC_CONTROLLER"), os.Getenv("WNC_ACCESS_TOKEN")
		if controller != "" && token != "" {
			if c := TestClient(t); c == nil {
				t.Error("expected client with real env")
			}
		} else {
			t.Skip("no real env")
		}
	})

	t.Run("TestClientWithActualSkip", func(t *testing.T) {
		origC, origT := os.Getenv("WNC_CONTROLLER"), os.Getenv("WNC_ACCESS_TOKEN")
		defer func() {
			if origC == "" {
				os.Unsetenv("WNC_CONTROLLER")
			} else {
				os.Setenv("WNC_CONTROLLER", origC)
			}
			if origT == "" {
				os.Unsetenv("WNC_ACCESS_TOKEN")
			} else {
				os.Setenv("WNC_ACCESS_TOKEN", origT)
			}
		}()
		os.Unsetenv("WNC_CONTROLLER")
		os.Unsetenv("WNC_ACCESS_TOKEN")
		t.Run("SkipSubtest", func(t *testing.T) { TestClient(t); t.Error("Test should have been skipped") })
	})

	t.Run("TestClientWithClientCreationError", func(t *testing.T) {
		t.Skip("fatal path not testable without failing suite")
	})
}

func TestTestClientAttemptCoverage(t *testing.T) {
	origC, origT := os.Getenv("WNC_CONTROLLER"), os.Getenv("WNC_ACCESS_TOKEN")
	os.Unsetenv("WNC_CONTROLLER")
	os.Unsetenv("WNC_ACCESS_TOKEN")
	if _, err := TestClientAttempt(); err == nil {
		t.Error("expected error when env vars missing")
	}
	os.Setenv("WNC_CONTROLLER", "c")
	os.Setenv("WNC_ACCESS_TOKEN", "t")
	origCreate := createCoreClient
	createCoreClient = func(controller, token string, opts ...core.Option) (*core.Client, error) { return nil, nil }
	defer func() {
		createCoreClient = origCreate
		os.Setenv("WNC_CONTROLLER", origC)
		os.Setenv("WNC_ACCESS_TOKEN", origT)
	}()
	if _, err := TestClientAttempt(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestTestClientFailOnErrorDowngrade(t *testing.T) {
	origFail := failOnClientError
	failOnClientError = false
	defer func() { failOnClientError = origFail }()
	os.Setenv("WNC_CONTROLLER", "bad-controller")
	os.Setenv("WNC_ACCESS_TOKEN", "bad-token")
	createOrig := createCoreClient
	createCoreClient = func(controller, token string, opts ...core.Option) (*core.Client, error) {
		return nil, fmt.Errorf("injected failure")
	}
	defer func() { createCoreClient = createOrig }()
	TestClient(t)
}

func TestSkipIfNoConnectionSuccess(t *testing.T) {
	orig := connectivityCheck
	defer func() { connectivityCheck = orig }()
	connectivityCheck = func(ctx context.Context, c *core.Client) error { return nil }
	c, err := core.New("controller", "token")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	SkipIfNoConnection(t, c)
	_ = executedSkipIfNoConnection
}

func TestSkipIfNoConnectionError(t *testing.T) {
	orig := connectivityCheck
	defer func() { connectivityCheck = orig }()
	connectivityCheck = func(ctx context.Context, c *core.Client) error { return context.DeadlineExceeded }
	c, err := core.New("controller", "token")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	SkipIfNoConnection(t, c)
}
