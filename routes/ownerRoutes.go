package routes

import (
	"go-donor-list-backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func ownerRoutes(router fiber.Router){
	//----> Donor-details routes.
	router.Delete("/donor-details", controllers.DeleteBloodStatByIdController)
	router.Get("/donor-details/:id", controllers.GetBloodStatByIdController)
	router.Patch("/donor-details/:id", controllers.EditBloodStatByIdController)

	//----> Donor-details routes.
	router.Delete("/donor-details", controllers.DeleteDonorDetailByIdController)
	router.Get("/donor-details/:id", controllers.GetDonorDetailByIdController)
	router.Patch("/donor-details/:id", controllers.EditDonorDetailByIdController)

	//----> User routes.
	router.Delete("/users/:id", controllers.DeleteUserByIdController)
	router.Get("/users", controllers.GetAllUsersController)
}