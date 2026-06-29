package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey_ReturnsKey(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey my-secret-key")

	got, err := GetAPIKey(headers)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != "my-secret-key" {
		t.Fatalf("expected %q, got %q", "my-secret-key", got)
	}
}

func TestGetAPIKey_ErrorWhenNoHeader(t *testing.T) {
	headers := http.Header{}

	_, err := GetAPIKey(headers)
	if err == nil {
		t.Fatal("expected an error when Authorization header is missing")
	}
}

func TestGetAPIKey_ErrorWhenMalformed(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "Bearer my-secret-key")

	_, err := GetAPIKey(headers)
	if err == nil {
		t.Fatal("expected an error for malformed authorization header")
	}
}
