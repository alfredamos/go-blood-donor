package middlewares

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/utils"
)

func RolePermission(roles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		//----> Get user role from context.
		userAuth := GetUserAuthFromContext(c)

		//----> Check for role in roles slice.
		if isValidRole := utils.Contains(roles, userAuth.Role); !isValidRole {
			//----> Invalid role.
			return c.Status(http.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": "You are not permitted to access this page!", "statusCode": http.StatusForbidden})
		}

		//----> The role is valid, user is authorized.

		return c.Next()

	}
}
