package handlers

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/samredway/scrapetool/internal/api/storage"
	"github.com/samredway/scrapetool/internal/api/types"
)

func SendEmail(writer storage.EmailWriter) fiber.Handler {
	return func(c *fiber.Ctx) error {
		req, err := parseAndValidateRequest[types.EmailRequest](c)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if err := writer.Write(req.Email); err != nil {
			slog.Error("Failed to save email", "error", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to process email",
			})
		}

		slog.Info("Email saved", "email", req.Email)
		return c.SendStatus(fiber.StatusOK)
	}
}
