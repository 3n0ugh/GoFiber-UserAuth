package router

import (
	"github.com/3n0ugh/GoFiber-RestAPI-UserAuth/server/handler"
	"github.com/3n0ugh/GoFiber-RestAPI-UserAuth/server/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Middleware
	api := app.Group("/auth", middlewares.Protected())

	api.Post("/signup", handler.Signup)
	api.Post("/login", handler.Login)
}
