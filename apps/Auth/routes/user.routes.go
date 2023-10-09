package routes

import (
	"b2b/apps/Auth/handler"
	"b2b/middleware"

	"github.com/gofiber/fiber/v2"
)

func ForbiddenError(c *fiber.Ctx) error {
	return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
		"error":   "Forbidden",
		"message": "You don't have permission to access this resource.",
	})
}

func UserRoutes(app *fiber.App) {
	api := app.Group("/api")
	user := api.Group("/user")

	user.Post("/register", handler.Register)

	user.Post("/login", handler.Login)

	user.Get("/", middleware.Protected(), handler.HelloWorld)
}
