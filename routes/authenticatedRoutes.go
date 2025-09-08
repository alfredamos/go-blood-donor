package routes

import (
	"go-donor-list-backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func protectedRoutes(router fiber.Router) {
	//----> Auth routes.
	router.Get("/auth/me", controllers.GetCurrentUserController)
	router.Patch("/auth/change-password", controllers.ChangePasswordController)
	router.Patch("/auth/edit-profile", controllers.EditProfileController)
	router.Post("/auth/logout", controllers.LogoutController)

	//----> Blood-stat routes.
	router.Post("/blood-stats", controllers.CreateBloodStatController)
	router.Delete("/blood-stats/:id", controllers.DeleteBloodStatByIdController)
	router.Get("/blood-stats/:id", controllers.GetBloodStatByIdController)
	router.Patch("/blood-stats/:id", controllers.EditBloodStatByIdController)
	
	//----> Donor details routes.
	router.Post("/donor-details", controllers.CreateDonorDetailController)
	router.Delete("/donor-details/:id", controllers.DeleteDonorDetailByIdController)
	router.Get("/donor-details/:id", controllers.GetDonorDetailByIdController)
	router.Patch("/donor-details/:id", controllers.EditDonorDetailByIdController)

	//----> Vital routes.
	router.Post("/vitals", controllers.CreateVitalController)
	router.Delete("/vitals/:id", controllers.DeleteVitalByIdController)
	router.Get("/vitals/:id", controllers.GetVitalByIdController)
	router.Patch("/vitals/:id", controllers.EditVitalByIdController)

}
