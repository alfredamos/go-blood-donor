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
	protectedRoutes(routesProtected.Group("/api"))

	//----> Admin routes middleware.
	routesOfAdmin := server.Use(middlewares.VerifyTokenJwt, middlewares.RolePermission("Admin"))

	//----> Admin routes
	adminRoutes(routesOfAdmin.Group("/api"))

	//----> Owner routes.
	routesOfOwner := server.Use(middlewares.VerifyTokenJwt, middlewares.SameUserAndAdmin)
	ownerRoutes(routesOfOwner.Group("/api"))
	
}
