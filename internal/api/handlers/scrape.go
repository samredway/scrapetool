package handlers

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/samredway/scrapeai/scrapeai"
	"github.com/samredway/scrapetool/internal/api/types"
)

// ScrapeFunc is a function that scrapes a URL and returns a ScrapeAiResult
type ScrapeFunc func(*scrapeai.ScrapeAiRequest) (*scrapeai.ScrapeAiResult, error)

// Scrape function to handle scraping requests
func Scrape(scrapeFunc ScrapeFunc) fiber.Handler {
	return func(c *fiber.Ctx) error {
		req, err := parseAndValidateRequest[types.ScrapeRequest](c)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		saReq, err := scrapeai.NewScrapeAiRequest(
			req.URL,
			req.Prompt,
			scrapeai.WithSchema(req.ResponseStructure),
		)
		if err != nil {
			slog.Error("Error creating scrapeai request", "error", err.Error())
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		result, err := scrapeFunc(saReq)
		if err != nil {
			slog.Error("Error scraping", "error", err.Error())
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.JSON(types.ScrapeResponse{
			URL:     req.URL,
			Prompt:  req.Prompt,
			Results: result.Results,
		})
	}
}
