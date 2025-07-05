package router

import (
	"github.com/Satr10/portfolio-website/handlers"
	"github.com/gofiber/fiber/v2"
)

func Init(app *fiber.App) {
	app.Get("/", handlers.IndexPage)

	htmx := app.Group("/htmx")
	htmx.Post("/command", handlers.CommandHtmx)
}
