package routes

import (
	"fmt"
	"go-donor-list-backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func adminRoutes(router fiber.Router) {
	fmt.Println("In admin routes")
	//----> Blood-stat routes.
	router.Get("/blood-stats", controllers.GetAllBloodStatsController)
	router.Delete("/blood-stats/all/delete-all", controllers.DeleteAllBloodStatController)

	//----> Donor details routes.
	router.Get("/donor-details", controllers.GetAllDonorDetailsController)
	router.Delete("/donor-details/all/delete-all", controllers.DeleteAllDonorsController)

	//----> Vital routes.
	router.Get("/vitals", controllers.GetAllVitalsController)
	router.Delete("/vitals/all/delete-all", controllers.DeleteAllVitalsController)
	//----> User routes
	router.Get("/users", controllers.GetAllUsersController)
}
