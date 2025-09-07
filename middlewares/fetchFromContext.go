package middlewares

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func GetUserAuthFromContext(c *fiber.Ctx) (string, string, bool) {
	//----> Get user role from context.
	role := c.Locals("role")

	//----> Get the user-id from context.
	userId := c.Locals("userId")

	//----> Check for admin role.
	isAdmin := fmt.Sprintf("%v", role) == "Admin"

	//----> Send back the role.
	return fmt.Sprintf("%v", role), fmt.Sprintf("%v", userId), isAdmin

}

func GetUserEmailFromContext(c *fiber.Ctx)string{
	//----> Get user-id from context.
	email := c.Locals("email")

	//----> Send back the user-id.
	return fmt.Sprintf("%v", email)
}

func GetUserIdFromContext(c *fiber.Ctx) string {
	//----> Get user-id from context.
	userId := c.Locals("userId")

	//----> Send back the user-id.
	return fmt.Sprintf("%v", userId)
}
