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
		return fiber.NewError(http.StatusUnprocessableEntity, err.Error())
	} else {
		// check the username already taken ?
		// dbUser := new(models.User)
		// var count int64 = 0
		// database.DB.Db.Where("username = ?", user.Username).Count(&count)
		// if count == 0 {
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
		err = database.DB.Db.Create(&User).Error
		if err != nil {
			return fiber.NewError(http.StatusConflict, "Username already taken.")
		}
		return c.Status(http.StatusOK).JSON(User)
	}
}

func Login(c *fiber.Ctx) error {
	user := new(User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	// validate the request body
	valid := validator.New()
	err := valid.Struct(user)
	if err != nil {
		return fiber.NewError(http.StatusUnprocessableEntity, err.Error())
	} else {
		dbUser := new(User)
		err := database.DB.Db.Where("username = ?", user.Username).Find(&dbUser).Error
		if err != nil {
			return fiber.NewError(http.StatusUnprocessableEntity, err.Error())
		} else {
			err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
			if err != nil {
				return fiber.NewError(http.StatusUnprocessableEntity, err.Error())
			} else {
				// create token here jwt
				return c.Status(200).JSON(fiber.Map{
					"message": "logged in",
					"success": true,
				})
			}
		}
	}
}
