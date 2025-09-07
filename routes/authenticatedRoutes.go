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
}
