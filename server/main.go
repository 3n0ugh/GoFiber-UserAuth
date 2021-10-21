package main

import (
	"github.com/3n0ugh/GoFiber-RestAPI-UserAuth/server/database"
	"github.com/3n0ugh/GoFiber-RestAPI-UserAuth/server/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	// HTTP Logger
	app.Use(logger.New())
	// CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:8081, http://localhost:8080",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		// return c.SendString("Hello, world!")
		return c.JSON(fiber.Map{
			"message": "Hello, world..!",
			"success": true,
		})
	})

	// Setup Routes
	router.SetupRoutes(app)

	// Not Found Message
	app.Use(notFound)

	app.Listen(":8080")
}

func notFound(c *fiber.Ctx) error {
	return c.SendString("Not Found - " + c.OriginalURL())
}
