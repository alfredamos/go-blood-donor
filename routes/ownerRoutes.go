package routes

import (
	"fmt"
	"go-donor-list-backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func ownerRoutes(router fiber.Router){
	fmt.Println("In owner-routes")
	//----> Blood-stats routes.
	router.Get("/blood-stats/get-all-by-user-id/:userId", controllers.GetBloodStatByUserIdController)
	router.Delete("/blood-stats/delete-by-user-id/:userId", controllers.DeleteBloodStatByUserIdController)
	router.Delete("/blood-stats/delete-all", controllers.DeleteAllBloodStatController)
	
	//----> Donor-details routes.
	router.Get("/donor-details/get-all-by-user-id/:userId", controllers.GetAllDonorDetailsByUserIdController)
	router.Delete("/donor-details/delete-by-user-id/:userId", controllers.DeleteAllDonorsByUserIdController)
	router.Delete("/blood-stats/delete-all", controllers.DeleteAllDonorsController)
	
	//----> vitals routes.
	router.Get("/vitals/get-all-by-user-id/:userId", controllers.GetAllVitalsByUserIdController)
	router.Delete("/vitals/delete-by-user-id/:userId", controllers.DeleteAllVitalsByUserIdController)
	router.Delete("/blood-stats/delete-all", controllers.DeleteAllVitalsController)
}