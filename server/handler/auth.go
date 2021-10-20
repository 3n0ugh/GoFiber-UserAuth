package handler

import (
	"net/http"
	"regexp"

	"github.com/3n0ugh/GoFiber-RestAPI-UserAuth/server/database"
	"github.com/3n0ugh/GoFiber-RestAPI-UserAuth/server/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

var regexpUsername = regexp.MustCompile("^[a-zA-Z]+[0-9]{5,12}$")

var regexpPassword = regexp.MustCompile("^[a-zA-Z]+[0-9]{8,32}$")

type User struct {
	Username   string `json:"username" validate:"required"`
	Password   string `json:"password"`
	RePassword string `json:"repassword"`
}

func Signup(c *fiber.Ctx) error {
	user := new(User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	// sonradan duruma gore duzenlenicek
	if !regexpUsername.MatchString(user.Username) {
		return fiber.NewError(http.StatusBadRequest, "Check Username!!")
	} else if regexpPassword.MatchString(user.Password) {
		return fiber.NewError(http.StatusBadRequest, "Check Password!!")
	} else if user.RePassword != user.Password {
		return fiber.NewError(http.StatusBadRequest, "Passwords Not Match!!")
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
