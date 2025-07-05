package main

import (
	"github.com/Satr10/portfolio-website/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django/v3"
)

func main() {
	engine := django.New("./views", ".html")
	appConfig := fiber.Config{
		Views: engine,
	}

	app := fiber.New(appConfig)
	app.Static("/", "./public")

	router.Init(app)

	app.Listen(":5001")
}
