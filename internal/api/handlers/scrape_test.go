package handlers

import (
	"bytes"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/samredway/scrapeai/scrapeai"
)

func TestHandleScrape(t *testing.T) {
	// Create a mock scrape function
	mockScrape := func(req *scrapeai.ScrapeAiRequest) (*scrapeai.ScrapeAiResult, error) {
		return &scrapeai.ScrapeAiResult{
			Results: "mocked result",
		}, nil
	}

	// Create a new handler with the mock function
	handler := NewScrapeHandler(mockScrape)

	// Create a new fiber app for testing
	app := fiber.New()

	// Setup the route with the handler method
	app.Post("/scrape", handler.HandleScrape)

	tests := []struct {
		name       string
		payload    string
		statusCode int
	}{
		{
			name:       "Valid request",
			payload:    `{"url": "https://example.com/", "prompt": "Test prompt", "responseStructure": ""}`,
			statusCode: fiber.StatusOK,
		},
		{
			name:       "Missing URL",
			payload:    `{"prompt": "Test prompt", "responseStructure": ""}`,
			statusCode: fiber.StatusBadRequest,
		},
		{
			name:       "Invalid URL",
			payload:    `{"url": "invalid-url", "prompt": "Test prompt", "responseStructure": ""}`,
			statusCode: fiber.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("POST", "/scrape", bytes.NewBufferString(tt.payload))
			req.Header.Set("Content-Type", "application/json")
			resp, err := app.Test(req)
			if err != nil {
				t.Errorf("Error getting response %v", err)
			}
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Errorf("Error reading response body")
			}
			if tt.statusCode != resp.StatusCode {
				t.Errorf("Incorrect status returned %v : %s", resp.StatusCode, body)
			}
		})
	}
}
