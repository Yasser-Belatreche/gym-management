package tests

import (
	"encoding/json"
	"net/http"
	"testing"
)

func TestSmoke(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/api/v1/health")
	if err != nil {
		t.Errorf("Failed to send GET request: %v", err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
		return
	}

	body := make(map[string]interface{})
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		t.Errorf("Failed to decode response body: %v", err)
		return
	}

	if body["status"] != "UP" {
		t.Errorf("Expected status=UP, got status=%s", body["status"])
		return
	}
}
