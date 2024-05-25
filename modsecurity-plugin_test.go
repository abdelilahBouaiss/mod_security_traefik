package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestModSecurityPlugin(t *testing.T) {
	// Define the configuration
	config := &Config{
		RuleFile: "path/to/rules.conf",
	}

	// Create the plugin
	handler, err := New(context.Background(), http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
	}), config, "modsecurity-plugin")

	if err != nil {
		t.Fatalf("Failed to create plugin: %v", err)
	}

	// Create a test request
	req := httptest.NewRequest(http.MethodGet, "http://localhost", nil)
	rw := httptest.NewRecorder()

	// Serve the test request
	handler.ServeHTTP(rw, req)

	// Check the response
	if rw.Code != http.StatusOK {
		t.Errorf("Expected status OK, got %v", rw.Code)
	}
}

