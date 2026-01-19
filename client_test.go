package relaywarden

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/relaywarden/go-sdk/errors"
)

func TestNewClient(t *testing.T) {
	client := NewClient("https://api.relaywarden.eu/api/v1", "test-token")
	if client == nil {
		t.Fatal("Expected client to be created")
	}
}

func TestSetProjectID(t *testing.T) {
	client := NewClient("https://api.relaywarden.eu/api/v1", "test-token")
	client.SetProjectID("project-123")
	if id := client.GetProjectID(); id == nil || *id != "project-123" {
		t.Errorf("Expected project ID to be 'project-123', got %v", id)
	}
}

func TestGetRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		if r.Header.Get("Authorization") != "Bearer test-token" {
			t.Errorf("Expected Authorization header")
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"data": map[string]interface{}{
				"id":   "123",
				"name": "Test",
			},
			"meta": map[string]interface{}{
				"request_id": "req-123",
			},
		})
	}))
	defer server.Close()

	client := NewClient(server.URL, "test-token")
	result, err := client.Get(context.Background(), "/test", nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result == nil {
		t.Fatal("Expected result to be non-nil")
	}
	if data, ok := result["data"].(map[string]interface{}); ok {
		if id, ok := data["id"].(string); !ok || id != "123" {
			t.Errorf("Expected id to be '123', got %v", id)
		}
	} else {
		t.Error("Expected data field in response")
	}
}

func TestAuthenticationError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": map[string]interface{}{
				"code":    "unauthorized",
				"message": "Unauthenticated",
			},
			"meta": map[string]interface{}{
				"request_id": "req-123",
			},
		})
	}))
	defer server.Close()

	client := NewClient(server.URL, "test-token")
	_, err := client.Get(context.Background(), "/test", nil)
	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if _, ok := err.(*errors.AuthenticationError); !ok {
		t.Errorf("Expected AuthenticationError, got %T", err)
	}
}

func TestRateLimitError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusTooManyRequests)
		w.Header().Set("Retry-After", "60")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": map[string]interface{}{
				"code":    "rate_limit_exceeded",
				"message": "Rate limit exceeded",
			},
			"meta": map[string]interface{}{
				"request_id": "req-123",
			},
		})
	}))
	defer server.Close()

	client := NewClient(server.URL, "test-token")
	_, err := client.Get(context.Background(), "/test", nil)
	if err == nil {
		t.Fatal("Expected error, got nil")
	}
	if rateLimitErr, ok := err.(*errors.RateLimitError); ok {
		if rateLimitErr.RetryAfter != 60 {
			t.Errorf("Expected RetryAfter to be 60, got %d", rateLimitErr.RetryAfter)
		}
	} else {
		t.Errorf("Expected RateLimitError, got %T", err)
	}
}
