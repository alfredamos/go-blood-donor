package routes

import (
	"go-donor-list-backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func adminRoutes(router fiber.Router) {
	//----> Blood-stat routes.
	router.Get("/blood-stats", controllers.GetAllBloodStatsController)

	//----> User routes
	router.Delete("/users/:id", controllers.DeleteUserByIdController)
	router.Get("/users", controllers.GetAllUsersController)
	router.Get("/users/:id", controllers.GetUserByIdController)
}
