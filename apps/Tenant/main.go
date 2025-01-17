package main

import (
	"b2b/apps/Tenant/routes"
	"b2b/models"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	app := fiber.New()
	models.ConnectDatabase()
	routes.UserRoutes(app)

	err := app.Listen(":7001")
	if err != nil {
		fmt.Printf("Error :%v", err)
	}
}
