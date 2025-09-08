package routes

import (
	"go-donor-list-backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func ownerRoutes(router fiber.Router){
	//----> User routes.
	router.Delete("/users/:id", controllers.DeleteUserByIdController)
	router.Get("/users/:id", controllers.GetUserByIdController)
}