package handlers

import (
	"log/slog"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/samredway/scrapeai/scrapeai"
	"github.com/samredway/scrapetool/internal/api/types"
)

var validate = validator.New()

type ScrapeFunc func(*scrapeai.ScrapeAiRequest) (*scrapeai.ScrapeAiResult, error)

type ScrapeHandler struct {
	scrapeFunc ScrapeFunc
}

func NewScrapeHandler(scrapeFunc ScrapeFunc) *ScrapeHandler {
	return &ScrapeHandler{
		scrapeFunc: scrapeFunc,
	}
}

// HandleScrape is now a method on ScrapeHandler
func (h *ScrapeHandler) HandleScrape(c *fiber.Ctx) error {
	var req types.ScrapeRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request format",
		})
	}

	slog.Info(
		"Scrape request",
		"url:", req.URL,
		"prompt:", req.Prompt,
		"responseStructure:", req.ResponseStructure,
	)

	err := validate.Struct(req)
	if err != nil {
		slog.Error("Invalid request", "error", err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	data, err := h.scrapeFunc(
		scrapeai.NewScrapeAiRequest(
			req.URL,
			req.Prompt,
			scrapeai.WithSchema(req.ResponseStructure),
		),
	)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	response := types.ScrapeResponse{
		URL:     req.URL,
		Prompt:  req.Prompt,
		Results: data.Results.(string),
	}

	return c.JSON(response)
}
