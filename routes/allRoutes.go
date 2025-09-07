package routes

import (
	"go-donor-list-backend/middlewares"

	"github.com/gofiber/fiber/v2"
)

func AllRoutes(server *fiber.App) {
	//----> Unprotected routes.
	unAuthenticatedRoutes := server.Group("/api")

	publicRoutes(unAuthenticatedRoutes)

	//----> Apply middleware for protected routes
	routesProtected := server.Use(middlewares.VerifyTokenJwt)

	//----> Protected routes.
	protectedRoutes(routesProtected.Group("/api/auth"))

	//----> Admin role permitted routes middleware.
	routesOfAdmin := server.Use(middlewares.VerifyTokenJwt, middlewares.RolePermission("Admin"))

	//----> Admin routes
	adminRoutes(routesOfAdmin.Group("/api"))

	//----> Owner and admin routes.
	//adminAndOwnerRoutes := server.Group("/api").Use(authenticate.VerifyTokenJwt, controllers.OwnerAndAdmin)
	//ownerAndAdminRoutes(adminAndOwnerRoutes)

	//----> Same user and admin routes.
	//userSameAndAdminRoutes := server.Group("/api").Use(authenticate.VerifyTokenJwt, authenticate.SameUserAndAdmin)
	//sameUserAndAdminRoutes(userSameAndAdminRoutes)

}
