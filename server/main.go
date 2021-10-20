package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	// HTTP Logger
	app.Use(logger.New())
	// CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "localhost:8080",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		// return c.SendString("Hello, world!")
		return c.JSON(fiber.Map{
			"message": "Hello, world..!",
			"success": true,
		})
	})

	app.Listen(":8080")
}
