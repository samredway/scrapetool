package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samredway/scrapetool/internal/api/handlers"
)

func SetupRoutes(app *fiber.App) {
	// Serve html templates
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{})
	})

	// data api
	api := app.Group("/api/v1")
	api.Post("/scrape", handlers.HandleScrape)
}
