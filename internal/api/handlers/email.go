package handlers

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/samredway/scrapetool/internal/api/types"
)

func SendEmail(c *fiber.Ctx) error {
	req, err := parseAndValidateRequest[types.EmailRequest](c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	slog.Info("Sending email", "email", req.Email)

	return c.SendStatus(fiber.StatusOK)
}
