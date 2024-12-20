package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samredway/scrapeai/scrapeai"
	"github.com/samredway/scrapetool/internal/api/handlers"
	"github.com/samredway/scrapetool/internal/api/storage"
)

func SetupRoutes(app *fiber.App) {
	// Serve html templates
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{})
	})
	app.Get("/demo", func(c *fiber.Ctx) error {
		return c.Render("demo", fiber.Map{})
	})

	// api endpoints
	api := app.Group("/api/v1")
	api.Post("/scrape", handlers.Scrape(scrapeai.Scrape))
	api.Post("/email", handlers.SendEmail(storage.NewFileEmailWriter()))
}
