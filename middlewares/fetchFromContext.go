package middlewares

import (
	"fmt"
	"go-donor-list-backend/utils"

	"github.com/gofiber/fiber/v2"
)


func GetUserAuthFromContext(c *fiber.Ctx) utils.UserAuth{
	//----> Get user role from context.
	role := fmt.Sprintf("%v", c.Locals("role"))

	//----> Get the user-id from context.
	userId := fmt.Sprintf("%v", c.Locals("userId"))
	
	isAdmin := role == "Admin"

	//----> Send back the role.
	return utils.UserAuth{IsAdmin: isAdmin, UserId: userId, Role: role}

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
