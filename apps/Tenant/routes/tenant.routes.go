package routes

import (
	"b2b/apps/Tenant/handler"
	"b2b/middleware"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	api := app.Group("/api")
	user := api.Group("/tenant")

	user.Post("/register", handler.Register)

	user.Post("/login", handler.Login)

	user.Get("/", middleware.Protected(), handler.HelloWorld)
}
