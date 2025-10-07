package routes

import (
	"go-donor-list-backend/controllers"
	"go-donor-list-backend/middlewares"

	"github.com/gofiber/fiber/v2"
)

func OwnerRoutes(router fiber.Router){
	//----> Blood-stats routes.
	router.Get("/blood-stats/get-by-user-id/:userId", middlewares.SameUserAndAdminMiddleware, controllers.GetBloodStatByUserIdController)
	router.Delete("/blood-stats/delete-by-user-id/:userId", middlewares.SameUserAndAdminMiddleware, controllers.DeleteBloodStatByUserIdController)
	
	
	//----> Donor-details routes.
	router.Get("/donor-details/get-by-user-id/:userId", controllers.GetAllDonorDetailsByUserIdController)
	router.Delete("/donor-details/delete-by-user-id/:userId", controllers.DeleteAllDonorsByUserIdController)

	//----> vitals routes.
	router.Get("/vitals/get-by-user-id/:userId", controllers.GetAllVitalsByUserIdController)
	router.Delete("/vitals/delete-by-user-id/:userId", controllers.DeleteAllVitalsByUserIdController)
	
}