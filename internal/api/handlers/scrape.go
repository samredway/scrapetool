package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samredway/scrapeai/scrapeai"
	"github.com/samredway/scrapetool/internal/api/types"
)

func HandleScrape(c *fiber.Ctx) error {
	var req types.ScrapeRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request format",
		})
	}

	data, err := scrapeai.Scrape(scrapeai.NewScrapeAiRequest(req.URL, req.Prompt))

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
