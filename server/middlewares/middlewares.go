package middlewares

import (
	"os"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/golobby/dotenv"
)

func Protected() func(c *fiber.Ctx) error {
	type config struct {
		jwtsecret string `env:"JWT_SECRET_KEY"`
	}
	file, _ := os.Open(".env")
	// if err != nil {
	// 	return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	// }
	defer file.Close()
	jwtconfig := &config{}
	dotenv.NewDecoder(file).Decode(jwtconfig)
	// if err != nil {
	// 	return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	// }
	return jwtware.New(jwtware.Config{
		SigningKey:   jwtconfig.jwtsecret,
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
