package routes

import (
	"go-donor-list-backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func ownerRoutes(router fiber.Router){
	//----> Blood-stats routes.
	router.Delete("/blood-stats", controllers.DeleteBloodStatController)
	router.Get("/blood-stats/:id", controllers.GetBloodStatByIdController)
	router.Patch("/blood-stats/:id", controllers.EditBloodStatController)

	//----> User routes.
	router.Delete("/users/:id", controllers.DeleteUserByIdController)
	router.Get("/users", controllers.GetAllUsersController)
}