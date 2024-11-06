package handlers

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/samredway/scrapeai/scrapeai"
	"github.com/samredway/scrapeai/scraping"
	"github.com/samredway/scrapetool/internal/api/types"
)

func HandleScrape(c *fiber.Ctx) error {
	var req types.ScrapeRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request format",
		})
	}

	data, err := scrapeai.Scrape(scrapeai.ScrapeAiRequest{
		Url:       req.URL,
		Prompt:    req.Prompt,
		FetchFunc: scraping.Fetch,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	response := types.ScrapeResponse{
		Results: strings.Join(data.Results, "\n"),
		URL:     req.URL,
		Prompt:  req.Prompt,
	}

	return c.JSON(response)
}
