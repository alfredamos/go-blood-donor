package middlewares

import (
	"net/http"
	"github.com/gofiber/fiber/v2"
)

func SameUserAndAdmin(c *fiber.Ctx) error{
	//----> Get the user-id from param.
	userIdFromContext := c.Params("userId")

	//----> Get user role from context.
	_, userId, isAdmin := GetUserAuthFromContext(c)

	//----> Check for same user.
	isUserSame := IsSameUser(userId, userIdFromContext)

	//----> Check for same user and admin privilege.
	if !isUserSame && !isAdmin {
		//----> Invalid user.
		return c.Status(http.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": "You are not permitted to access this page!", "statusCode": http.StatusForbidden})
	}

	//----> Same user and admin are allowed to gain access.
	return c.Next()
}

func IsSameUser(userId1, userId2 string) bool {
	return userId1 == userId2
}
