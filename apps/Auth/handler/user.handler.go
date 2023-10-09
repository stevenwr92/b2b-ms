package handler

import (
	"b2b/apps/Auth/dto"
	"b2b/middleware"
	"b2b/models"
	"b2b/utils"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func HelloWorld(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Message": "Hello World",
	})
}

func Register(c *fiber.Ctx) error {
	user := models.User{}

	if err := c.BodyParser(&user); err != nil {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, "Please Check Your Input")
	}

	hash, err := utils.HashPassword(user.Password)
	if err != nil {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, "Failed to hash password.")
	}

	user.Password = hash

	if err := models.DB.Create(&user).Error; err != nil {
		return utils.SendErrorResponse(c, fiber.StatusInternalServerError, "Failed to create user")
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

func Login(c *fiber.Ctx) error {
	user := models.User{}
	login := dto.Login{}

	if err := c.BodyParser(&login); err != nil {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, "Please Check Your Input")
	}

	err := validate.Struct(login)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("%s is %s", err.Field(), err.Tag())
			return utils.SendErrorResponse(c, fiber.StatusBadRequest, errorMessage)
		}
	}

	checkUser := models.DB.Where("email = ?", login.Email).First(&user)

	if checkUser.Error != nil {
		return utils.SendErrorResponse(c, fiber.StatusUnauthorized, "Invalid credentials")
	}

	if !utils.CheckPasswordHash(login.Password, user.Password) {
		return utils.SendErrorResponse(c, fiber.StatusUnauthorized, "Invalid credentials")
	}

	token, err := middleware.GenerateJWTToken(user)
	if err != nil {
		return utils.SendErrorResponse(c, fiber.StatusInternalServerError, "Error create JWT")
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}
