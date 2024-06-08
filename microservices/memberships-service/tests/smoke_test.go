package tests

import (
	"encoding/json"
	"gym-management-memberships/src/lib/primitives/application_specific"
	"net/http"
	"os"
	"testing"
)

func TestSmoke(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:8080/api/v1/health", nil)
	if err != nil {
		t.Errorf("Failed to create GET request: %v", err)
		return
	}

	session, err := application_specific.NewSession().ToBase64()
	if err != nil {
		t.Errorf("Failed to create session: %v", err)
		return
	}

	req.Header.Set("X-Session", session)
	req.Header.Set("X-Api-Secret", os.Getenv("API_SECRET"))

	resp, err := http.DefaultClient.Do(req)
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
