package main

import (
	"github.com/gofiber/fiber/v2"

	"hajin-chung/templ-hi/views"
)

func main() {
	app := fiber.New()

	app.Static("/public", "public")

	app.Use("/", func(c *fiber.Ctx) error {
		c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
		return c.Next()
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Send(Render(views.IndexPage("sample page", "john")))
	})

	app.Listen(":3000")
}

