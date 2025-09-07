package main

import (
	"go-donor-list-backend/initializers"
	"go-donor-list-backend/routes"

	"github.com/gofiber/fiber/v2"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

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
