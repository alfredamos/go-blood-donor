package main

import (
	"go-donor-list-backend/initializers"
	"go-donor-list-backend/middlewares"
	"go-donor-list-backend/routes"

	"github.com/gofiber/fiber/v2"
)

func init() {
	initializers.LoadEnvVariable()
	initializers.ConnectDB()
}

func main() {
	//----> Initialize fiber app.
	app := fiber.New()

	//----> Public routes.
	unprotectedRoutes := app.Group("/api/auth")
	routes.PublicRoutes(unprotectedRoutes)

	//----> Protected routes.
	protectedRoutes := app.Group("/api")
	protectedRoutes.Use(middlewares.VerifyTokenJwtMiddleware)
	routes.ProtectedRoutes(protectedRoutes)

	//----> Owners routes.
	routes.OwnerRoutes(protectedRoutes)

	//----> Admin routes.
	adminRoutes := app.Group("/api")
	adminRoutes.Use(middlewares.VerifyTokenJwtMiddleware, middlewares.RolePermissionMiddleware)
	routes.AdminRoutes(adminRoutes)


	app.Listen(":5000")
}
