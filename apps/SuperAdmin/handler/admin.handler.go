package handler

import (
	"b2b/apps/SuperAdmin/dto"
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
	admin := models.Admin{}

	if err := c.BodyParser(&admin); err != nil {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, "Please Check Your Input")
	}

	hash, err := utils.HashPassword(admin.Password)
	if err != nil {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, "Failed to hash password.")
	}

	admin.Password = hash

	if err := models.DB.Create(&admin).Error; err != nil {
		return utils.SendErrorResponse(c, fiber.StatusInternalServerError, "Failed to create admin")
	}

	return c.Status(fiber.StatusCreated).JSON(admin)
}

func Login(c *fiber.Ctx) error {
	admin := models.Admin{}
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

	checkUser := models.DB.Where("email = ?", login.Email).First(&admin)

	if checkUser.Error != nil {
		return utils.SendErrorResponse(c, fiber.StatusUnauthorized, "Invalid credentials")
	}

	if !utils.CheckPasswordHash(login.Password, admin.Password) {
		return utils.SendErrorResponse(c, fiber.StatusUnauthorized, "Invalid credentials")
	}

	token, err := middleware.GenerateJWTTokenAdmin(admin)
	if err != nil {
		return utils.SendErrorResponse(c, fiber.StatusInternalServerError, "Error create JWT")
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}
