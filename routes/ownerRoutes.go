package routes

import (
	"go-donor-list-backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func ownerRoutes(router fiber.Router){
	//----> Donor-details routes.
	router.Delete("/blood-stats/:id", controllers.DeleteBloodStatByIdController)
	router.Get("/blood-stats/:id", controllers.GetBloodStatByIdController)
	router.Patch("/blood-stats/:id", controllers.EditBloodStatByIdController)

	//----> Donor-details routes.
	router.Delete("/donor-details/:id", controllers.DeleteDonorDetailByIdController)
	router.Get("/donor-details/:id", controllers.GetDonorDetailByIdController)
	router.Patch("/donor-details/:id", controllers.EditDonorDetailByIdController)

	//----> Vital routes.
	router.Delete("/vitals/:id", controllers.DeleteVitalByIdController)
	router.Get("/vitals/:id", controllers.GetVitalByIdController)
	router.Patch("/vitals/:id", controllers.EditVitalByIdController)

	//----> User routes.
	router.Delete("/users/:id", controllers.DeleteUserByIdController)
	router.Get("/users/:id", controllers.GetUserByIdController)
}