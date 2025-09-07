package middlewares

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// CorsMiddleware CORS middleware function definition
func CorsMiddleware() fiber.Handler {
	// Define allowed origins as a comma-separated string
	originsString := "http://localhost:3000,http://localhost:4200,http://localhost:5173,http://localhost:5174"
	var allowedOrigins []string
	// Split the originsString into individual origins and store them in allowedOrigins slice
	allowedOrigins = strings.Split(originsString, ",")

	// Return the actual middleware handler function
	return func(c *fiber.Ctx) error {
		//----> Get the Origin header from the request
		origin := GetOrigin(c)

		//----> Function to check if a given origin is allowed
		isOriginAllowed := getAllAllowedOrigins(origin, allowedOrigins)

		//----> Check if the origin is allowed
		if isOriginAllowed {
			//----> If the origin is allowed, set CORS headers in the response
			setCorsHeaders(c, origin)
		}

		fmt.Println("origin : ", origin)

		fmt.Println("isOriginAllowed : ", isOriginAllowed)

		//----> Handle preflight OPTIONS requests by aborting with status 204
		if c.Method() == "OPTIONS" {
			return c.Status(fiber.StatusNoContent).JSON(fiber.Map{"message": "wrongly configured"})
		}

		//----> Call the next handler
		return c.Next()
	}
}

func getAllAllowedOrigins(origin string, allowedOrigins []string) bool {
	for _, allowedOrigin := range allowedOrigins {
		if origin == allowedOrigin {
			return true
		}
	}
	return false
}

func setCorsHeaders(c *fiber.Ctx, origin string) {
	c.Set("Access-Control-Allow-Origin", origin)
	c.Set("Access-Control-Allow-Credentials", "true")
	c.Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Set("Access-Control-Allow-Methods", "POST, PATCH, DELETE, OPTIONS, GET, PUT")
}

func GetOrigin(c *fiber.Ctx) string {
	return c.Get("Origin")
}
