package routes

import (
	"fmt"
	"go-donor-list-backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func publicRoutes(router fiber.Router) {
	fmt.Println("In public-routes")
	//----> Auth-routes.
	router.Post("/auth/signup", controllers.SignupController)
	router.Post("/auth/login", controllers.LoginController)
}
