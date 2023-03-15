package main

import (
	"github.com/VSarcher/WebThumbnailGenerator/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/image", handlers.SaveThumbnail)

	app.Listen(":3000")
}
