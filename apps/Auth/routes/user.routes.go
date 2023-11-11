package routes

import (
	"b2b/apps/Auth/handler"
	"b2b/middleware"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {

	api := app.Group("/api")
	user := api.Group("/user")

	app.Get("/", handler.HelloAuth)

	user.Post("/register", handler.Register)

	user.Post("/login", handler.Login)

	user.Get("/", handler.HelloAuth)

	user.Get("/protected", middleware.Protected(), handler.HelloWorld)

}
