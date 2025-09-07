package routes

import (
	"go-donor-list-backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func protectedRoutes(router fiber.Router) {
	//----> Auth routes.
	router.Get("/me", controllers.GetCurrentUserController)
	router.Patch("/change-password", controllers.ChangePasswordController)
	router.Patch("/edit-profile", controllers.EditProfileController)
	router.Post("/logout", controllers.LogoutController)
}
