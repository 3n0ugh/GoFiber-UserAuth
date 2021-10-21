package middlewares

import (
	"github.com/3n0ugh/GoFiber-RestAPI-UserAuth/server/config"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

func Protected() func(c *fiber.Ctx) error {
	jwtconfig, _ := config.GetConfig()

	return jwtware.New(jwtware.Config{
		SigningKey:   jwtconfig.JwtSecretKey,
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		})
	} else {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid or expired JWT",
			"data":    nil,
		})
	}
}
