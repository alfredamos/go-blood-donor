package routes

import (
	"go-donor-list-backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func adminRoutes(router fiber.Router) {
	//----> Blood-stat routes.
	router.Get("/blood-stats", controllers.GetAllBloodStatsController)

	//----> Donor details routes.
	router.Get("/donor-details", controllers.GetAllDonorDetailsController)

	//----> Vital routes.
	router.Get("/vitals", controllers.GetAllVitalsController)

	//----> User routes
	router.Get("/users", controllers.GetAllUsersController)
}
