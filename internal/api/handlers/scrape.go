package handlers

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/samredway/scrapeai/scrapeai"
	"github.com/samredway/scrapetool/internal/api/types"
)

// ScrapeFunc is a function that scrapes a URL and returns a ScrapeAiResult
type ScrapeFunc func(*scrapeai.ScrapeAiRequest) (*scrapeai.ScrapeAiResult, error)

// ScrapeHandler allows passing in a scraper function as a dependency injection
type ScrapeHandler struct {
	scrapeFunc ScrapeFunc
}

// NewScrapeHandler creates a new ScrapeHandler with a given scrape function
func NewScrapeHandler(scrapeFunc ScrapeFunc) *ScrapeHandler {
	return &ScrapeHandler{
		scrapeFunc: scrapeFunc,
	}
}

// Scrape calles underlying scrapeai Scrape function to collect return data
func (h *ScrapeHandler) Scrape(c *fiber.Ctx) error {
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

	result, err := h.scrapeFunc(saReq)
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
