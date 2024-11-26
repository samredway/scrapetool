package handlers

import (
	"fmt"
	"log/slog"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

// parseAndValidateRequest is a generic function to parse and validate any request type
func parseAndValidateRequest[T any](c *fiber.Ctx) (T, error) {
	var req T

	if err := c.BodyParser(&req); err != nil {
		var zero T
		slog.Error("Invalid request", "error", err.Error())
		return zero, fmt.Errorf("invalid request: %w", err)
	}

	if err := validate.Struct(req); err != nil {
		var zero T
		slog.Error("Invalid request", "error", err.Error())
		return zero, fmt.Errorf("invalid request: %w", err)
	}

	return req, nil
}
