package handler

import "github.com/gofiber/fiber/v2"

func Hello(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Success..!",
	})
}

func Signup(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "signup",
		"success": true,
	})
}

func Login(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "login",
		"success": true,
	})
}
