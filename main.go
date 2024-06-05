package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Static("/", "./public")

	app.Get("/api", func(c *fiber.Ctx) error {
		return c.SendString("What up")
	})

	log.Fatal(app.Listen(":3000"))
}
