package routes

import (
	"fmt"
	"go-donor-list-backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func ownerRoutes(router fiber.Router){
	fmt.Println("In owner-routes")
	//----> Blood-stats routes.
	router.Get("/blood-stats/get-by-user-id/:userId", controllers.GetBloodStatByUserIdController)
	router.Delete("/blood-stats/delete-by-user-id/:userId", controllers.DeleteBloodStatByUserIdController)
	
	
	//----> Donor-details routes.
	router.Get("/donor-details/get-by-user-id/:userId", controllers.GetAllDonorDetailsByUserIdController)
	router.Delete("/donor-details/delete-by-user-id/:userId", controllers.DeleteAllDonorsByUserIdController)

	//----> vitals routes.
	router.Get("/vitals/get-by-user-id/:userId", controllers.GetAllVitalsByUserIdController)
	router.Delete("/vitals/delete-by-user-id/:userId", controllers.DeleteAllVitalsByUserIdController)
	
}