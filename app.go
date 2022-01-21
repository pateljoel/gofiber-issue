package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/handlebars"
)

func main() {
	engine := handlebars.New("./views", ".hbs").Layout("body")

	app := fiber.New(fiber.Config{
		Views:                 engine,
		ViewsLayout:           "layouts/default",
		DisableStartupMessage: true,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"page":  "index",
			"title": "Homepage",
		}, "layouts/default")
	})

	app.Get("/test", func(c *fiber.Ctx) error {
		return c.Render("test", fiber.Map{
			"page":  "test",
			"title": "Test",
		}, "layouts/default")
	})

	log.Fatal(app.Listen(":3000"))

}