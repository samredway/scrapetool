package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samredway/scrapeai/scrapeai"
	"github.com/samredway/scrapetool/internal/api/handlers"
)

func SetupRoutes(app *fiber.App) {
	// Initialize handlers with their specific dependencies
	scrapeHandler := handlers.NewScrapeHandler(scrapeai.Scrape)

	// Serve html templates
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{})
	})

	// api endpoints
	api := app.Group("/api/v1")
	api.Post("/scrape", scrapeHandler.HandleScrape)
}
