package main

import (
	"log"

	"github.com/VSarcher/WebThumbnailGenerator/database"
	"github.com/VSarcher/WebThumbnailGenerator/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("DB connected")
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/image", handlers.SaveThumbnail)

	app.Listen(":3000")
}
