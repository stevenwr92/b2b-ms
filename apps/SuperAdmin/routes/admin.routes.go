package routes

import (
	"b2b/apps/SuperAdmin/handler"
	"b2b/middleware"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	api := app.Group("/api")
	admin := api.Group("/admin")

	admin.Post("/register", handler.Register)

	admin.Post("/login", handler.Login)

	admin.Get("/", middleware.ProtectedAdmin(), handler.HelloWorld)

}
