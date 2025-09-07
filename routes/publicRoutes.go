package routes

import (
	"go-donor-list-backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func publicRoutes(router fiber.Router) {
	//----> Auth-routes.
	router.Post("/auth/signup", controllers.SignupController)
	router.Post("/auth/login", controllers.LoginController)
}
