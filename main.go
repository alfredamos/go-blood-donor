package main

import (
	"go-donor-list-backend/initializers"
	"go-donor-list-backend/routes"

	"github.com/gofiber/fiber/v2"
)

func init() {
	initializers.LoadEnvVariable()
	initializers.ConnectDB()
}

func main() {
	//----> Initialize fiber app.
	app := fiber.New()

	//----> Call all routes.
	routes.AllRoutes(app)

	app.Listen(":5000")
}
