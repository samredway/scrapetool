package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
	"github.com/samredway/scrapetool/internal/api/routes"
)

const (
	VIEWS_PATH  = "./web/views"
	STATIC_PATH = "./web/static"
	VIEW_EXT    = ".html"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	engine := html.New(VIEWS_PATH, VIEW_EXT)
	app := fiber.New(fiber.Config{Views: engine})
	app.Static("/", STATIC_PATH)
	routes.SetupRoutes(app)
	if err := app.Listen(":3000"); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
