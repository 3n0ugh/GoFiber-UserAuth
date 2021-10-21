package handler

import (
	"time"

	"github.com/3n0ugh/GoFiber-RestAPI-UserAuth/server/config"
	"github.com/3n0ugh/GoFiber-RestAPI-UserAuth/server/database"
	"github.com/3n0ugh/GoFiber-RestAPI-UserAuth/server/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	// validate the request body
	valid := validator.New()
	err := valid.Struct(user)
	if err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	} else {
		// hash the password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		// inserting into User model
		User := new(models.User)
		User.Username = user.Username
		User.Password = string(hashedPassword)
		// adding user to db
		err = database.DB.Db.Create(&User).Error
		// if user exist in db got an error
		if err != nil {
			return fiber.NewError(fiber.StatusConflict, "username already taken")
		}
		return createTokenSendResponse(c, User)
	}
}

func Login(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	// validate the request body
	valid := validator.New()
	err := valid.Struct(user)
	if err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	} else {
		dbUser := new(models.User)
		// find the user from the database
		database.DB.Db.Where("username = ?", user.Username).Find(&dbUser)
		if dbUser.Username == "" {
			return fiber.NewError(fiber.StatusUnprocessableEntity, "wrong username or password")
		} else {
			// checking password with hashedPassword
			err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
			if err != nil {
				return fiber.NewError(fiber.StatusUnprocessableEntity, "wrong username or password")
			} else {
				return createTokenSendResponse(c, user)
			}
		}
	}
}

func createTokenSendResponse(c *fiber.Ctx, user *models.User) error {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	// Get the jwt configs from dotenv
	config, err := config.GetConfig()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	claims["username"] = user.Username
	claims["exp"] = time.Now().Add(time.Minute * config.JwtExpireTime).Unix()

	// create java web token
	t, err := token.SignedString([]byte(config.JwtSecretKey))
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(fiber.Map{
		"message": "logged in",
		"success": true,
		"data":    t,
	})
}
