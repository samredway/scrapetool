package handlers

import (
	"bytes"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestSendEmail(t *testing.T) {
	// setup app (done only once)
	app := fiber.New()
	app.Post("/email", SendEmail)

	tests := []struct {
		name       string
		email      string
		wantStatus int
	}{
		{
			name:       "valid email",
			email:      "test@example.com",
			wantStatus: fiber.StatusOK,
		},
		{
			name:       "invalid email",
			email:      "invalid-email",
			wantStatus: fiber.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// setup request
			req := httptest.NewRequest("POST", "/email", bytes.NewBufferString(`{"email": "`+tt.email+`"}`))
			req.Header.Set("Content-Type", "application/json")

			// send request
			resp, err := app.Test(req)
			if err != nil {
				t.Fatalf("Error sending request: %v", err)
			}

			assert.Equal(t, tt.wantStatus, resp.StatusCode)
		})
	}
}
