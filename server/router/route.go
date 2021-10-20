package router

import (
	"github.com/3n0ugh/GoFiber-RestAPI-UserAuth/server/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Middleware
	api := app.Group("/auth")

	api.Get("/", handler.Hello)
	api.Post("/signup", handler.Signup)
	api.Post("/login", handler.Login)
}
