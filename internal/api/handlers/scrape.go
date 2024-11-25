package handlers

import (
	"log/slog"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/samredway/scrapeai/scrapeai"
	"github.com/samredway/scrapetool/internal/api/types"
)

// TODO doesnt really belong here. Should I have a validate package or is
// this a dependency?
var validate = validator.New()

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
	req, err := h.parseAndValidateRequest(c)
	if err != nil {
		return err
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

func (h *ScrapeHandler) parseAndValidateRequest(c *fiber.Ctx) (types.ScrapeRequest, error) {
	var req types.ScrapeRequest

	if err := c.BodyParser(&req); err != nil {
		return req, c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request format",
		})
	}

	if err := validate.Struct(req); err != nil {
		slog.Error("Invalid request", "error", err.Error())
		return req, c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return req, nil
}
