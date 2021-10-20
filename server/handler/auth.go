package handler

import (
	"net/http"

	"github.com/3n0ugh/GoFiber-RestAPI-UserAuth/server/database"
	"github.com/3n0ugh/GoFiber-RestAPI-UserAuth/server/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username   string `json:"username" validate:"required,min=5,max=12,alphanum"`
	Password   string `json:"password" validate:"required,min=8,max=32"`
	RePassword string `json:"repassword" validate:"required"`
}

func Signup(c *fiber.Ctx) error {
	user := new(User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	// validate the request body
	valid := validator.New()
	err := valid.Struct(user)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	} else {
		// check the username already taken ?
		if err := database.DB.Db.Where("username = ?", user.Username).Limit(0).Find(&user.Username); err != nil {
			// hashing password
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
			if err != nil {
				return fiber.NewError(http.StatusInternalServerError, err.Error())
			}
			// inserting into User model
			User := new(models.User)
			User.Username = user.Username
			User.Password = string(hashedPassword)
			// adding user to db
			database.DB.Db.Create(&User)
			return c.Status(http.StatusOK).JSON(User)
		} else {
			return fiber.NewError(http.StatusInternalServerError, "Username already taken.")
		}
	}
}

func Login(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "login",
		"success": true,
	})
}
